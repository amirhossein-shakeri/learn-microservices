// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: hub.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{0}
}

type Hub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"` // optional string password = 2;
}

func (x *Hub) Reset() {
	*x = Hub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hub) ProtoMessage() {}

func (x *Hub) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hub.ProtoReflect.Descriptor instead.
func (*Hub) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{1}
}

func (x *Hub) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type Hubs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hubs  []*Hub `protobuf:"bytes,1,rep,name=hubs,proto3" json:"hubs,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"` // optional int32 count = 2;
}

func (x *Hubs) Reset() {
	*x = Hubs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hubs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hubs) ProtoMessage() {}

func (x *Hubs) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hubs.ProtoReflect.Descriptor instead.
func (*Hubs) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{2}
}

func (x *Hubs) GetHubs() []*Hub {
	if x != nil {
		return x.Hubs
	}
	return nil
}

func (x *Hubs) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_hub_proto protoreflect.FileDescriptor

var file_hub_proto_rawDesc = []byte{
	0x0a, 0x09, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x03, 0x48, 0x75, 0x62, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x04, 0x48, 0x75, 0x62, 0x73, 0x12, 0x18,
	0x0a, 0x04, 0x68, 0x75, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x48,
	0x75, 0x62, 0x52, 0x04, 0x68, 0x75, 0x62, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x7b,
	0x0a, 0x0a, 0x48, 0x75, 0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x48, 0x75, 0x62, 0x12, 0x04, 0x2e, 0x48, 0x75, 0x62, 0x1a, 0x04, 0x2e, 0x48,
	0x75, 0x62, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x62, 0x73,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x04, 0x2e, 0x48, 0x75, 0x62, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x19, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x75, 0x62, 0x12,
	0x04, 0x2e, 0x48, 0x75, 0x62, 0x1a, 0x04, 0x2e, 0x48, 0x75, 0x62, 0x22, 0x00, 0x12, 0x1c, 0x0a,
	0x0c, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x48, 0x75, 0x62, 0x12, 0x04, 0x2e,
	0x48, 0x75, 0x62, 0x1a, 0x04, 0x2e, 0x48, 0x75, 0x62, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x2d, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x68, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_hub_proto_rawDescOnce sync.Once
	file_hub_proto_rawDescData = file_hub_proto_rawDesc
)

func file_hub_proto_rawDescGZIP() []byte {
	file_hub_proto_rawDescOnce.Do(func() {
		file_hub_proto_rawDescData = protoimpl.X.CompressGZIP(file_hub_proto_rawDescData)
	})
	return file_hub_proto_rawDescData
}

var file_hub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hub_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: Empty
	(*Hub)(nil),   // 1: Hub
	(*Hubs)(nil),  // 2: Hubs
}
var file_hub_proto_depIdxs = []int32{
	1, // 0: Hubs.hubs:type_name -> Hub
	1, // 1: HubCluster.GetHub:input_type -> Hub
	0, // 2: HubCluster.ListHubs:input_type -> Empty
	1, // 3: HubCluster.CreateHub:input_type -> Hub
	1, // 4: HubCluster.TerminateHub:input_type -> Hub
	1, // 5: HubCluster.GetHub:output_type -> Hub
	1, // 6: HubCluster.ListHubs:output_type -> Hub
	1, // 7: HubCluster.CreateHub:output_type -> Hub
	1, // 8: HubCluster.TerminateHub:output_type -> Hub
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hub_proto_init() }
func file_hub_proto_init() {
	if File_hub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hub); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hubs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hub_proto_goTypes,
		DependencyIndexes: file_hub_proto_depIdxs,
		MessageInfos:      file_hub_proto_msgTypes,
	}.Build()
	File_hub_proto = out.File
	file_hub_proto_rawDesc = nil
	file_hub_proto_goTypes = nil
	file_hub_proto_depIdxs = nil
}