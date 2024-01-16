package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/thor-electronics/core/hub/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createHub(client pb.HubClusterClient, hub *pb.Hub) {
	log.Printf("Creating hub %v...", hub.GroupId)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	h, err := client.CreateHub(ctx, hub)
	if err != nil {
		log.Fatalf("client.CreateHub failed: %v", err)
	}
	log.Printf("%+v", h)
}

func listHubs(client pb.HubClusterClient, empty *pb.Empty) {
	log.Printf("listHubs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ListHubs(ctx, empty)
	if err != nil {
		log.Fatalf("client.ListHubs failed: %v", err)
	}
	for {
		hub, err := stream.Recv()
		// log.Printf("Got hub %v", hub.GetGroupId())
		if err == io.EOF {
			log.Printf("EOF")
			break
		}
		if err != nil {
			log.Fatalf("client.ListHubs failed: %v", err)
		}
		log.Printf("Hub: %+v", hub)
	}
}

func terminateHub(client pb.HubClusterClient, hub *pb.Hub) {
	// log.Printf("Terminating hub %v...", hub.GroupId)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	h, err := client.TerminateHub(ctx, hub)
	if err != nil {
		log.Fatalf("client.TerminateHub failed: %v", err)
	}
	log.Printf("Terminated hub %v", h)
}

func main() {
	conn, err := grpc.Dial("localhost:3990", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial localhost:3990: %v", err)
	}
	defer conn.Close()
	client := pb.NewHubClusterClient(conn)

	createHub(client, &pb.Hub{GroupId: "Group_ID_01"})
	createHub(client, &pb.Hub{GroupId: "Group_ID_02"})
	createHub(client, &pb.Hub{GroupId: "Group_ID_03"})

	listHubs(client, &pb.Empty{})

	createHub(client, &pb.Hub{GroupId: "Group_ID_04"})
	createHub(client, &pb.Hub{GroupId: "Group_ID_05"})
	createHub(client, &pb.Hub{GroupId: "Group_ID_06"})
	createHub(client, &pb.Hub{GroupId: "Group_ID_07"})

	terminateHub(client, &pb.Hub{GroupId: "Group_ID_04"})
	terminateHub(client, &pb.Hub{GroupId: "Group_ID_06"})

	listHubs(client, &pb.Empty{})
}
