package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/thor-electronics/core/hub/proto"
	"google.golang.org/grpc"
)

type hubServer struct {
	pb.UnimplementedHubClusterServer
	// hubs []*pb.Hub // todo: read-only after initialized?

	mu sync.Mutex // Protects hubs
	// hubs map[*pb.Hub]bool
	hubs map[string]*pb.Hub
	// hubs map[string][]*pb.Hub
}

func (s *hubServer) GetHub(ctx context.Context, hub *pb.Hub) (*pb.Hub, error) {
	log.Printf("GetHub %s", hub.GroupId)
	for groupId := range s.hubs { // , active
		if hub.GroupId == groupId {
			// if hub.GroupId == h.GroupId {
			// if proto.Equal(hub.GroupId, h.GroupId) {
			return s.hubs[groupId], nil
		}
	}
	return nil, fmt.Errorf("no hub found with group id %v", hub.GroupId)
}

func (s *hubServer) ListHubs(empty *pb.Empty, stream pb.HubCluster_ListHubsServer) error {
	log.Printf("GetHubs")
	for groupId, hub := range s.hubs {
		log.Printf("Sending hub %s...", groupId)
		if err := stream.Send(hub); err != nil {
			return err
		}
	}
	return nil
}

func (s *hubServer) CreateHub(ctx context.Context, hub *pb.Hub) (*pb.Hub, error) {
	log.Printf("CreateHub %+v", hub)
	s.hubs[hub.GroupId] = hub
	return s.hubs[hub.GroupId], nil
}

func (s *hubServer) TerminateHub(ctx context.Context, hub *pb.Hub) (*pb.Hub, error) {
	log.Printf("TerminateHub %+v", hub)
	delete(s.hubs, hub.GroupId)
	return hub, nil
}

func newServer() *hubServer {
	s := &hubServer{hubs: make(map[string]*pb.Hub)}
	return s
}

func main() {
	fmt.Println("Starting hub service on localhost:3990...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3990))
	if err != nil {
		log.Fatalf("failed to listen to 3990: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHubClusterServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
