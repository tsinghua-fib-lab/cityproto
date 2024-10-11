// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/elec/output/v1/output.proto

package outputv1

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

type Aoi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用电需求不满足比例（没电100%/有电0）
	UnsatisfiedDemandRatio float64 `protobuf:"fixed64,2,opt,name=unsatisfied_demand_ratio,json=unsatisfiedDemandRatio,proto3" json:"unsatisfied_demand_ratio,omitempty"`
	// 用电需求不满足人数 （没电就是aoi里的人都不满足）
	UnsatisfiedDemandNum int32 `protobuf:"varint,3,opt,name=unsatisfied_demand_num,json=unsatisfiedDemandNum,proto3" json:"unsatisfied_demand_num,omitempty"`
	// 该aoi当前时刻的用电需求,单位为KW
	Demand float64 `protobuf:"fixed64,4,opt,name=demand,proto3" json:"demand,omitempty"`
	// 电力系统向该aoi供应的电力,单位为KW
	Supply float64 `protobuf:"fixed64,5,opt,name=supply,proto3" json:"supply,omitempty"`
}

func (x *Aoi) Reset() {
	*x = Aoi{}
	mi := &file_city_elec_output_v1_output_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Aoi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aoi) ProtoMessage() {}

func (x *Aoi) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_output_v1_output_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aoi.ProtoReflect.Descriptor instead.
func (*Aoi) Descriptor() ([]byte, []int) {
	return file_city_elec_output_v1_output_proto_rawDescGZIP(), []int{0}
}

func (x *Aoi) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Aoi) GetUnsatisfiedDemandRatio() float64 {
	if x != nil {
		return x.UnsatisfiedDemandRatio
	}
	return 0
}

func (x *Aoi) GetUnsatisfiedDemandNum() int32 {
	if x != nil {
		return x.UnsatisfiedDemandNum
	}
	return 0
}

func (x *Aoi) GetDemand() float64 {
	if x != nil {
		return x.Demand
	}
	return 0
}

func (x *Aoi) GetSupply() float64 {
	if x != nil {
		return x.Supply
	}
	return 0
}

type AveragePower struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前各类变压器的平均承受功率，单位为MW
	Transformer_500 float64 `protobuf:"fixed64,1,opt,name=transformer_500,json=transformer500,proto3" json:"transformer_500,omitempty"`
	Transformer_220 float64 `protobuf:"fixed64,2,opt,name=transformer_220,json=transformer220,proto3" json:"transformer_220,omitempty"`
	Transformer_110 float64 `protobuf:"fixed64,3,opt,name=transformer_110,json=transformer110,proto3" json:"transformer_110,omitempty"`
	Transformer_10  float64 `protobuf:"fixed64,4,opt,name=transformer_10,json=transformer10,proto3" json:"transformer_10,omitempty"`
}

func (x *AveragePower) Reset() {
	*x = AveragePower{}
	mi := &file_city_elec_output_v1_output_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AveragePower) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AveragePower) ProtoMessage() {}

func (x *AveragePower) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_output_v1_output_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AveragePower.ProtoReflect.Descriptor instead.
func (*AveragePower) Descriptor() ([]byte, []int) {
	return file_city_elec_output_v1_output_proto_rawDescGZIP(), []int{1}
}

func (x *AveragePower) GetTransformer_500() float64 {
	if x != nil {
		return x.Transformer_500
	}
	return 0
}

func (x *AveragePower) GetTransformer_220() float64 {
	if x != nil {
		return x.Transformer_220
	}
	return 0
}

func (x *AveragePower) GetTransformer_110() float64 {
	if x != nil {
		return x.Transformer_110
	}
	return 0
}

func (x *AveragePower) GetTransformer_10() float64 {
	if x != nil {
		return x.Transformer_10
	}
	return 0
}

var File_city_elec_output_v1_output_proto protoreflect.FileDescriptor

var file_city_elec_output_v1_output_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xb5, 0x01, 0x0a, 0x03, 0x41, 0x6f, 0x69, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x38, 0x0a, 0x18, 0x75, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x16, 0x75, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x34, 0x0a, 0x16, 0x75, 0x6e, 0x73,
	0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x75, 0x6e, 0x73, 0x61, 0x74,
	0x69, 0x73, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x22,
	0xb0, 0x01, 0x0a, 0x0c, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x5f,
	0x35, 0x30, 0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x35, 0x30, 0x30, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x32, 0x32, 0x30, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x32,
	0x32, 0x30, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x5f, 0x31, 0x31, 0x30, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x31, 0x31, 0x30, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x31, 0x30, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x31, 0x30, 0x42, 0xd3, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x65, 0x6c, 0x65, 0x63, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x45,
	0x4f, 0xaa, 0x02, 0x13, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x2e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45,
	0x6c, 0x65, 0x63, 0x5c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f,
	0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x6c, 0x65, 0x63, 0x5c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x16, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x6c, 0x65, 0x63, 0x3a, 0x3a, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_elec_output_v1_output_proto_rawDescOnce sync.Once
	file_city_elec_output_v1_output_proto_rawDescData = file_city_elec_output_v1_output_proto_rawDesc
)

func file_city_elec_output_v1_output_proto_rawDescGZIP() []byte {
	file_city_elec_output_v1_output_proto_rawDescOnce.Do(func() {
		file_city_elec_output_v1_output_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_elec_output_v1_output_proto_rawDescData)
	})
	return file_city_elec_output_v1_output_proto_rawDescData
}

var file_city_elec_output_v1_output_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_elec_output_v1_output_proto_goTypes = []any{
	(*Aoi)(nil),          // 0: city.elec.output.v1.Aoi
	(*AveragePower)(nil), // 1: city.elec.output.v1.AveragePower
}
var file_city_elec_output_v1_output_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_elec_output_v1_output_proto_init() }
func file_city_elec_output_v1_output_proto_init() {
	if File_city_elec_output_v1_output_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_elec_output_v1_output_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_elec_output_v1_output_proto_goTypes,
		DependencyIndexes: file_city_elec_output_v1_output_proto_depIdxs,
		MessageInfos:      file_city_elec_output_v1_output_proto_msgTypes,
	}.Build()
	File_city_elec_output_v1_output_proto = out.File
	file_city_elec_output_v1_output_proto_rawDesc = nil
	file_city_elec_output_v1_output_proto_goTypes = nil
	file_city_elec_output_v1_output_proto_depIdxs = nil
}
