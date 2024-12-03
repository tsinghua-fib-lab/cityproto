// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: city/event/v1/event_service.proto

package eventv1

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

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty" yaml:"event" bson:"event" db:"event"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	mi := &file_city_event_v1_event_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v1_event_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_city_event_v1_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	mi := &file_city_event_v1_event_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v1_event_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_city_event_v1_event_service_proto_rawDescGZIP(), []int{1}
}

type PullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullRequest) Reset() {
	*x = PullRequest{}
	mi := &file_city_event_v1_event_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRequest) ProtoMessage() {}

func (x *PullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v1_event_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRequest.ProtoReflect.Descriptor instead.
func (*PullRequest) Descriptor() ([]byte, []int) {
	return file_city_event_v1_event_service_proto_rawDescGZIP(), []int{2}
}

type PullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty" yaml:"events" bson:"events" db:"events"`
}

func (x *PullResponse) Reset() {
	*x = PullResponse{}
	mi := &file_city_event_v1_event_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullResponse) ProtoMessage() {}

func (x *PullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v1_event_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullResponse.ProtoReflect.Descriptor instead.
func (*PullResponse) Descriptor() ([]byte, []int) {
	return file_city_event_v1_event_service_proto_rawDescGZIP(), []int{3}
}

func (x *PullResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_city_event_v1_event_service_proto protoreflect.FileDescriptor

var file_city_event_v1_event_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x19, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a,
	0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d,
	0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a,
	0x0c, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x9d, 0x01, 0x0a, 0x0c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x07,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c,
	0x12, 0x1a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb3, 0x01, 0x0a, 0x11,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c,
	0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x45, 0x58, 0xaa, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_event_v1_event_service_proto_rawDescOnce sync.Once
	file_city_event_v1_event_service_proto_rawDescData = file_city_event_v1_event_service_proto_rawDesc
)

func file_city_event_v1_event_service_proto_rawDescGZIP() []byte {
	file_city_event_v1_event_service_proto_rawDescOnce.Do(func() {
		file_city_event_v1_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_event_v1_event_service_proto_rawDescData)
	})
	return file_city_event_v1_event_service_proto_rawDescData
}

var file_city_event_v1_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_city_event_v1_event_service_proto_goTypes = []any{
	(*PublishRequest)(nil),  // 0: city.event.v1.PublishRequest
	(*PublishResponse)(nil), // 1: city.event.v1.PublishResponse
	(*PullRequest)(nil),     // 2: city.event.v1.PullRequest
	(*PullResponse)(nil),    // 3: city.event.v1.PullResponse
	(*Event)(nil),           // 4: city.event.v1.Event
}
var file_city_event_v1_event_service_proto_depIdxs = []int32{
	4, // 0: city.event.v1.PublishRequest.event:type_name -> city.event.v1.Event
	4, // 1: city.event.v1.PullResponse.events:type_name -> city.event.v1.Event
	0, // 2: city.event.v1.EventService.Publish:input_type -> city.event.v1.PublishRequest
	2, // 3: city.event.v1.EventService.Pull:input_type -> city.event.v1.PullRequest
	1, // 4: city.event.v1.EventService.Publish:output_type -> city.event.v1.PublishResponse
	3, // 5: city.event.v1.EventService.Pull:output_type -> city.event.v1.PullResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_city_event_v1_event_service_proto_init() }
func file_city_event_v1_event_service_proto_init() {
	if File_city_event_v1_event_service_proto != nil {
		return
	}
	file_city_event_v1_event_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_event_v1_event_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_event_v1_event_service_proto_goTypes,
		DependencyIndexes: file_city_event_v1_event_service_proto_depIdxs,
		MessageInfos:      file_city_event_v1_event_service_proto_msgTypes,
	}.Build()
	File_city_event_v1_event_service_proto = out.File
	file_city_event_v1_event_service_proto_rawDesc = nil
	file_city_event_v1_event_service_proto_goTypes = nil
	file_city_event_v1_event_service_proto_depIdxs = nil
}
