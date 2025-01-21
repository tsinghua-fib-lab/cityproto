// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: city/comm/interaction/aoi/v1/aoi_service.proto

package aoiv1

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

type GetBadAoiIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBadAoiIDRequest) Reset() {
	*x = GetBadAoiIDRequest{}
	mi := &file_city_comm_interaction_aoi_v1_aoi_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBadAoiIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBadAoiIDRequest) ProtoMessage() {}

func (x *GetBadAoiIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_aoi_v1_aoi_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBadAoiIDRequest.ProtoReflect.Descriptor instead.
func (*GetBadAoiIDRequest) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDescGZIP(), []int{0}
}

type GetBadAoiIDResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ids           []int32                `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBadAoiIDResponse) Reset() {
	*x = GetBadAoiIDResponse{}
	mi := &file_city_comm_interaction_aoi_v1_aoi_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBadAoiIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBadAoiIDResponse) ProtoMessage() {}

func (x *GetBadAoiIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_aoi_v1_aoi_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBadAoiIDResponse.ProtoReflect.Descriptor instead.
func (*GetBadAoiIDResponse) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetBadAoiIDResponse) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_city_comm_interaction_aoi_v1_aoi_service_proto protoreflect.FileDescriptor

var file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x6f, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x6f, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x6f, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x14,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x64, 0x41, 0x6f, 0x69, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x61, 0x64, 0x41, 0x6f,
	0x69, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x32, 0x80, 0x01,
	0x0a, 0x0a, 0x41, 0x6f, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x64, 0x41, 0x6f, 0x69, 0x49, 0x44, 0x12, 0x30, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x6f, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x64, 0x41, 0x6f, 0x69, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x6f, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x64, 0x41, 0x6f, 0x69, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x8c, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61,
	0x6f, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x41, 0x6f, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69,
	0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x6f, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x6f, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x43,
	0x43, 0x49, 0x41, 0xaa, 0x02, 0x1c, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6f, 0x69, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x1c, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x41, 0x6f, 0x69, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x28, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x41, 0x6f, 0x69, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x43,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x6f, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDescOnce sync.Once
	file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDescData = file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDesc
)

func file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDescGZIP() []byte {
	file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDescOnce.Do(func() {
		file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDescData)
	})
	return file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDescData
}

var file_city_comm_interaction_aoi_v1_aoi_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_comm_interaction_aoi_v1_aoi_service_proto_goTypes = []any{
	(*GetBadAoiIDRequest)(nil),  // 0: city.comm.interaction.aoi.v1.GetBadAoiIDRequest
	(*GetBadAoiIDResponse)(nil), // 1: city.comm.interaction.aoi.v1.GetBadAoiIDResponse
}
var file_city_comm_interaction_aoi_v1_aoi_service_proto_depIdxs = []int32{
	0, // 0: city.comm.interaction.aoi.v1.AoiService.GetBadAoiID:input_type -> city.comm.interaction.aoi.v1.GetBadAoiIDRequest
	1, // 1: city.comm.interaction.aoi.v1.AoiService.GetBadAoiID:output_type -> city.comm.interaction.aoi.v1.GetBadAoiIDResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_comm_interaction_aoi_v1_aoi_service_proto_init() }
func file_city_comm_interaction_aoi_v1_aoi_service_proto_init() {
	if File_city_comm_interaction_aoi_v1_aoi_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_comm_interaction_aoi_v1_aoi_service_proto_goTypes,
		DependencyIndexes: file_city_comm_interaction_aoi_v1_aoi_service_proto_depIdxs,
		MessageInfos:      file_city_comm_interaction_aoi_v1_aoi_service_proto_msgTypes,
	}.Build()
	File_city_comm_interaction_aoi_v1_aoi_service_proto = out.File
	file_city_comm_interaction_aoi_v1_aoi_service_proto_rawDesc = nil
	file_city_comm_interaction_aoi_v1_aoi_service_proto_goTypes = nil
	file_city_comm_interaction_aoi_v1_aoi_service_proto_depIdxs = nil
}
