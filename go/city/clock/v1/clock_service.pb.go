// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: city/clock/v1/clock_service.proto

package clockv1

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

// 获取当前的模拟时间请求
// request of getting current simulation clock
type NowRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NowRequest) Reset() {
	*x = NowRequest{}
	mi := &file_city_clock_v1_clock_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowRequest) ProtoMessage() {}

func (x *NowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_clock_v1_clock_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowRequest.ProtoReflect.Descriptor instead.
func (*NowRequest) Descriptor() ([]byte, []int) {
	return file_city_clock_v1_clock_service_proto_rawDescGZIP(), []int{0}
}

// 获取当前的模拟时间响应
// response of getting current simulation clock
type NowResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 当前模拟的天数
	// current simulation day
	Day *int32 `protobuf:"varint,2,opt,name=day,proto3,oneof" json:"day,omitempty"`
	// 当前的模拟时间，单位为秒
	// current simulation clock, in seconds
	T             float64 `protobuf:"fixed64,1,opt,name=t,proto3" json:"t,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NowResponse) Reset() {
	*x = NowResponse{}
	mi := &file_city_clock_v1_clock_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowResponse) ProtoMessage() {}

func (x *NowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_clock_v1_clock_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowResponse.ProtoReflect.Descriptor instead.
func (*NowResponse) Descriptor() ([]byte, []int) {
	return file_city_clock_v1_clock_service_proto_rawDescGZIP(), []int{1}
}

func (x *NowResponse) GetDay() int32 {
	if x != nil && x.Day != nil {
		return *x.Day
	}
	return 0
}

func (x *NowResponse) GetT() float64 {
	if x != nil {
		return x.T
	}
	return 0
}

var File_city_clock_v1_clock_service_proto protoreflect.FileDescriptor

var file_city_clock_v1_clock_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x22, 0x0c, 0x0a, 0x0a, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3a, 0x0a, 0x0b, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x15, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03,
	0x64, 0x61, 0x79, 0x88, 0x01, 0x01, 0x12, 0x0c, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x01, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64, 0x61, 0x79, 0x32, 0x4c, 0x0a, 0x0c,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03,
	0x4e, 0x6f, 0x77, 0x12, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb3, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x42, 0x11, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61,
	0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x43, 0x58, 0xaa, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x19, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0f, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_clock_v1_clock_service_proto_rawDescOnce sync.Once
	file_city_clock_v1_clock_service_proto_rawDescData = file_city_clock_v1_clock_service_proto_rawDesc
)

func file_city_clock_v1_clock_service_proto_rawDescGZIP() []byte {
	file_city_clock_v1_clock_service_proto_rawDescOnce.Do(func() {
		file_city_clock_v1_clock_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_clock_v1_clock_service_proto_rawDescData)
	})
	return file_city_clock_v1_clock_service_proto_rawDescData
}

var file_city_clock_v1_clock_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_clock_v1_clock_service_proto_goTypes = []any{
	(*NowRequest)(nil),  // 0: city.clock.v1.NowRequest
	(*NowResponse)(nil), // 1: city.clock.v1.NowResponse
}
var file_city_clock_v1_clock_service_proto_depIdxs = []int32{
	0, // 0: city.clock.v1.ClockService.Now:input_type -> city.clock.v1.NowRequest
	1, // 1: city.clock.v1.ClockService.Now:output_type -> city.clock.v1.NowResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_clock_v1_clock_service_proto_init() }
func file_city_clock_v1_clock_service_proto_init() {
	if File_city_clock_v1_clock_service_proto != nil {
		return
	}
	file_city_clock_v1_clock_service_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_clock_v1_clock_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_clock_v1_clock_service_proto_goTypes,
		DependencyIndexes: file_city_clock_v1_clock_service_proto_depIdxs,
		MessageInfos:      file_city_clock_v1_clock_service_proto_msgTypes,
	}.Build()
	File_city_clock_v1_clock_service_proto = out.File
	file_city_clock_v1_clock_service_proto_rawDesc = nil
	file_city_clock_v1_clock_service_proto_goTypes = nil
	file_city_clock_v1_clock_service_proto_depIdxs = nil
}
