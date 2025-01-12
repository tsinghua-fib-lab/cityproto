// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: city/map/v2/aoi_service.proto

package mapv2

import (
	v2 "git.fiblab.net/sim/protos/v2/go/city/person/v2"
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

// AOI状态
// AOI state
type AoiState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// AOI ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// AOI内的人
	// Persons in AOI
	Persons       []*v2.PersonMotion `protobuf:"bytes,2,rep,name=persons,proto3" json:"persons,omitempty" bson:"persons" db:"persons" yaml:"persons"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AoiState) Reset() {
	*x = AoiState{}
	mi := &file_city_map_v2_aoi_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AoiState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AoiState) ProtoMessage() {}

func (x *AoiState) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_aoi_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AoiState.ProtoReflect.Descriptor instead.
func (*AoiState) Descriptor() ([]byte, []int) {
	return file_city_map_v2_aoi_service_proto_rawDescGZIP(), []int{0}
}

func (x *AoiState) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AoiState) GetPersons() []*v2.PersonMotion {
	if x != nil {
		return x.Persons
	}
	return nil
}

// 获取AOI信息请求
// Request for getting AOI information
type GetAoiRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 指定AOI ID列表，如果为空，则返回所有AOI信息
	// List of targeted AOI IDs, if empty, returns all information of AOIs
	AoiIds        []int32 `protobuf:"varint,1,rep,packed,name=aoi_ids,json=aoiIds,proto3" json:"aoi_ids,omitempty" bson:"aoi_ids" db:"aoi_ids" yaml:"aoi_ids"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAoiRequest) Reset() {
	*x = GetAoiRequest{}
	mi := &file_city_map_v2_aoi_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAoiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAoiRequest) ProtoMessage() {}

func (x *GetAoiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_aoi_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAoiRequest.ProtoReflect.Descriptor instead.
func (*GetAoiRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_aoi_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAoiRequest) GetAoiIds() []int32 {
	if x != nil {
		return x.AoiIds
	}
	return nil
}

// 获取AOI信息响应
// Response for getting AOI information
type GetAoiResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// AOI信息列表
	// Lis of AOIs information
	States        []*AoiState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty" bson:"states" db:"states" yaml:"states"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAoiResponse) Reset() {
	*x = GetAoiResponse{}
	mi := &file_city_map_v2_aoi_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAoiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAoiResponse) ProtoMessage() {}

func (x *GetAoiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_aoi_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAoiResponse.ProtoReflect.Descriptor instead.
func (*GetAoiResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_aoi_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAoiResponse) GetStates() []*AoiState {
	if x != nil {
		return x.States
	}
	return nil
}

var File_city_map_v2_aoi_service_proto protoreflect.FileDescriptor

var file_city_map_v2_aoi_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x6f,
	0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x63, 0x69,
	0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x08, 0x41, 0x6f, 0x69,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x28, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6f, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x61, 0x6f, 0x69, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x6f, 0x69, 0x49, 0x64, 0x73, 0x22, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6f,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6f, 0x69, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x32, 0x4f, 0x0a, 0x0a, 0x41, 0x6f, 0x69, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6f, 0x69,
	0x12, 0x1a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6f, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6f,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa3, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x42, 0x0f, 0x41,
	0x6f, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74,
	0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x3b, 0x6d, 0x61,
	0x70, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x0b, 0x43, 0x69, 0x74, 0x79,
	0x2e, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x4d,
	0x61, 0x70, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x17, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x4d, 0x61, 0x70,
	0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x3a, 0x3a, 0x56, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_map_v2_aoi_service_proto_rawDescOnce sync.Once
	file_city_map_v2_aoi_service_proto_rawDescData = file_city_map_v2_aoi_service_proto_rawDesc
)

func file_city_map_v2_aoi_service_proto_rawDescGZIP() []byte {
	file_city_map_v2_aoi_service_proto_rawDescOnce.Do(func() {
		file_city_map_v2_aoi_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_map_v2_aoi_service_proto_rawDescData)
	})
	return file_city_map_v2_aoi_service_proto_rawDescData
}

var file_city_map_v2_aoi_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_city_map_v2_aoi_service_proto_goTypes = []any{
	(*AoiState)(nil),        // 0: city.map.v2.AoiState
	(*GetAoiRequest)(nil),   // 1: city.map.v2.GetAoiRequest
	(*GetAoiResponse)(nil),  // 2: city.map.v2.GetAoiResponse
	(*v2.PersonMotion)(nil), // 3: city.person.v2.PersonMotion
}
var file_city_map_v2_aoi_service_proto_depIdxs = []int32{
	3, // 0: city.map.v2.AoiState.persons:type_name -> city.person.v2.PersonMotion
	0, // 1: city.map.v2.GetAoiResponse.states:type_name -> city.map.v2.AoiState
	1, // 2: city.map.v2.AoiService.GetAoi:input_type -> city.map.v2.GetAoiRequest
	2, // 3: city.map.v2.AoiService.GetAoi:output_type -> city.map.v2.GetAoiResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_city_map_v2_aoi_service_proto_init() }
func file_city_map_v2_aoi_service_proto_init() {
	if File_city_map_v2_aoi_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_map_v2_aoi_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_map_v2_aoi_service_proto_goTypes,
		DependencyIndexes: file_city_map_v2_aoi_service_proto_depIdxs,
		MessageInfos:      file_city_map_v2_aoi_service_proto_msgTypes,
	}.Build()
	File_city_map_v2_aoi_service_proto = out.File
	file_city_map_v2_aoi_service_proto_rawDesc = nil
	file_city_map_v2_aoi_service_proto_goTypes = nil
	file_city_map_v2_aoi_service_proto_depIdxs = nil
}
