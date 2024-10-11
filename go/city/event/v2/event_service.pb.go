// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/event/v2/event_service.proto

package eventv2

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

// 按照topic查询事件请求
type GetEventsByTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主题
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *GetEventsByTopicRequest) Reset() {
	*x = GetEventsByTopicRequest{}
	mi := &file_city_event_v2_event_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsByTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsByTopicRequest) ProtoMessage() {}

func (x *GetEventsByTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v2_event_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsByTopicRequest.ProtoReflect.Descriptor instead.
func (*GetEventsByTopicRequest) Descriptor() ([]byte, []int) {
	return file_city_event_v2_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetEventsByTopicRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

// 按照topic查询事件响应
type GetEventsByTopicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件列表
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventsByTopicResponse) Reset() {
	*x = GetEventsByTopicResponse{}
	mi := &file_city_event_v2_event_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsByTopicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsByTopicResponse) ProtoMessage() {}

func (x *GetEventsByTopicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v2_event_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsByTopicResponse.ProtoReflect.Descriptor instead.
func (*GetEventsByTopicResponse) Descriptor() ([]byte, []int) {
	return file_city_event_v2_event_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventsByTopicResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// 确认事件已被处理请求
type ResolveEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件列表
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ResolveEventsRequest) Reset() {
	*x = ResolveEventsRequest{}
	mi := &file_city_event_v2_event_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveEventsRequest) ProtoMessage() {}

func (x *ResolveEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v2_event_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveEventsRequest.ProtoReflect.Descriptor instead.
func (*ResolveEventsRequest) Descriptor() ([]byte, []int) {
	return file_city_event_v2_event_service_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveEventsRequest) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// 确认事件已被处理响应
type ResolveEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResolveEventsResponse) Reset() {
	*x = ResolveEventsResponse{}
	mi := &file_city_event_v2_event_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveEventsResponse) ProtoMessage() {}

func (x *ResolveEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v2_event_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveEventsResponse.ProtoReflect.Descriptor instead.
func (*ResolveEventsResponse) Descriptor() ([]byte, []int) {
	return file_city_event_v2_event_service_proto_rawDescGZIP(), []int{3}
}

var File_city_event_v2_event_service_proto protoreflect.FileDescriptor

var file_city_event_v2_event_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x32, 0x1a, 0x19, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x48,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x44, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x17,
	0x0a, 0x15, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xcf, 0x01, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x26, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x0d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x23,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb3, 0x01, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x42,
	0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62,
	0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x32, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x45,
	0x58, 0xaa, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x56,
	0x32, 0xca, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x56,
	0x32, 0xe2, 0x02, 0x19, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x56,
	0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f,
	0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_event_v2_event_service_proto_rawDescOnce sync.Once
	file_city_event_v2_event_service_proto_rawDescData = file_city_event_v2_event_service_proto_rawDesc
)

func file_city_event_v2_event_service_proto_rawDescGZIP() []byte {
	file_city_event_v2_event_service_proto_rawDescOnce.Do(func() {
		file_city_event_v2_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_event_v2_event_service_proto_rawDescData)
	})
	return file_city_event_v2_event_service_proto_rawDescData
}

var file_city_event_v2_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_city_event_v2_event_service_proto_goTypes = []any{
	(*GetEventsByTopicRequest)(nil),  // 0: city.event.v2.GetEventsByTopicRequest
	(*GetEventsByTopicResponse)(nil), // 1: city.event.v2.GetEventsByTopicResponse
	(*ResolveEventsRequest)(nil),     // 2: city.event.v2.ResolveEventsRequest
	(*ResolveEventsResponse)(nil),    // 3: city.event.v2.ResolveEventsResponse
	(*Event)(nil),                    // 4: city.event.v2.Event
}
var file_city_event_v2_event_service_proto_depIdxs = []int32{
	4, // 0: city.event.v2.GetEventsByTopicResponse.events:type_name -> city.event.v2.Event
	4, // 1: city.event.v2.ResolveEventsRequest.events:type_name -> city.event.v2.Event
	0, // 2: city.event.v2.EventService.GetEventsByTopic:input_type -> city.event.v2.GetEventsByTopicRequest
	2, // 3: city.event.v2.EventService.ResolveEvents:input_type -> city.event.v2.ResolveEventsRequest
	1, // 4: city.event.v2.EventService.GetEventsByTopic:output_type -> city.event.v2.GetEventsByTopicResponse
	3, // 5: city.event.v2.EventService.ResolveEvents:output_type -> city.event.v2.ResolveEventsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_city_event_v2_event_service_proto_init() }
func file_city_event_v2_event_service_proto_init() {
	if File_city_event_v2_event_service_proto != nil {
		return
	}
	file_city_event_v2_event_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_event_v2_event_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_event_v2_event_service_proto_goTypes,
		DependencyIndexes: file_city_event_v2_event_service_proto_depIdxs,
		MessageInfos:      file_city_event_v2_event_service_proto_msgTypes,
	}.Build()
	File_city_event_v2_event_service_proto = out.File
	file_city_event_v2_event_service_proto_rawDesc = nil
	file_city_event_v2_event_service_proto_goTypes = nil
	file_city_event_v2_event_service_proto_depIdxs = nil
}
