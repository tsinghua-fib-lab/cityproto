// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/comm/interaction/gateway/v1/gateway.proto

package gatewayv1

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

type Reason int32

const (
	Reason_REASON_UNSPECIFIED Reason = 0
	Reason_REASON_RUIN        Reason = 1
	Reason_REASON_CASCADE     Reason = 2
)

// Enum value maps for Reason.
var (
	Reason_name = map[int32]string{
		0: "REASON_UNSPECIFIED",
		1: "REASON_RUIN",
		2: "REASON_CASCADE",
	}
	Reason_value = map[string]int32{
		"REASON_UNSPECIFIED": 0,
		"REASON_RUIN":        1,
		"REASON_CASCADE":     2,
	}
)

func (x Reason) Enum() *Reason {
	p := new(Reason)
	*p = x
	return p
}

func (x Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_city_comm_interaction_gateway_v1_gateway_proto_enumTypes[0].Descriptor()
}

func (Reason) Type() protoreflect.EnumType {
	return &file_city_comm_interaction_gateway_v1_gateway_proto_enumTypes[0]
}

func (x Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reason.Descriptor instead.
func (Reason) EnumDescriptor() ([]byte, []int) {
	return file_city_comm_interaction_gateway_v1_gateway_proto_rawDescGZIP(), []int{0}
}

type Station struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Reason Reason `protobuf:"varint,3,opt,name=reason,proto3,enum=city.comm.interaction.gateway.v1.Reason" json:"reason,omitempty"`
}

func (x *Station) Reset() {
	*x = Station{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Station) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Station) ProtoMessage() {}

func (x *Station) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Station.ProtoReflect.Descriptor instead.
func (*Station) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_gateway_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *Station) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Station) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Station) GetReason() Reason {
	if x != nil {
		return x.Reason
	}
	return Reason_REASON_UNSPECIFIED
}

var File_city_comm_interaction_gateway_v1_gateway_proto protoreflect.FileDescriptor

var file_city_comm_interaction_gateway_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x22, 0x73, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0x45, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x55, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x53, 0x43, 0x41, 0x44, 0x45, 0x10, 0x02, 0x42, 0xa2,
	0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62,
	0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x43, 0x43, 0x49, 0x47, 0xaa, 0x02, 0x20, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20, 0x43, 0x69, 0x74,
	0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2c,
	0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x43,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_comm_interaction_gateway_v1_gateway_proto_rawDescOnce sync.Once
	file_city_comm_interaction_gateway_v1_gateway_proto_rawDescData = file_city_comm_interaction_gateway_v1_gateway_proto_rawDesc
)

func file_city_comm_interaction_gateway_v1_gateway_proto_rawDescGZIP() []byte {
	file_city_comm_interaction_gateway_v1_gateway_proto_rawDescOnce.Do(func() {
		file_city_comm_interaction_gateway_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_comm_interaction_gateway_v1_gateway_proto_rawDescData)
	})
	return file_city_comm_interaction_gateway_v1_gateway_proto_rawDescData
}

var file_city_comm_interaction_gateway_v1_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_city_comm_interaction_gateway_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_city_comm_interaction_gateway_v1_gateway_proto_goTypes = []any{
	(Reason)(0),     // 0: city.comm.interaction.gateway.v1.Reason
	(*Station)(nil), // 1: city.comm.interaction.gateway.v1.Station
}
var file_city_comm_interaction_gateway_v1_gateway_proto_depIdxs = []int32{
	0, // 0: city.comm.interaction.gateway.v1.Station.reason:type_name -> city.comm.interaction.gateway.v1.Reason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_city_comm_interaction_gateway_v1_gateway_proto_init() }
func file_city_comm_interaction_gateway_v1_gateway_proto_init() {
	if File_city_comm_interaction_gateway_v1_gateway_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_comm_interaction_gateway_v1_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_comm_interaction_gateway_v1_gateway_proto_goTypes,
		DependencyIndexes: file_city_comm_interaction_gateway_v1_gateway_proto_depIdxs,
		EnumInfos:         file_city_comm_interaction_gateway_v1_gateway_proto_enumTypes,
		MessageInfos:      file_city_comm_interaction_gateway_v1_gateway_proto_msgTypes,
	}.Build()
	File_city_comm_interaction_gateway_v1_gateway_proto = out.File
	file_city_comm_interaction_gateway_v1_gateway_proto_rawDesc = nil
	file_city_comm_interaction_gateway_v1_gateway_proto_goTypes = nil
	file_city_comm_interaction_gateway_v1_gateway_proto_depIdxs = nil
}
