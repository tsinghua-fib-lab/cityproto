// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/water/input/v1/water.proto

package inputv1

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

type RainPeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 起始时间点，单位为秒，但必须整小时
	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty" yaml:"start" bson:"start" db:"start"`
	// 降雨量：单位mm
	Rainfall float64 `protobuf:"fixed64,2,opt,name=rainfall,proto3" json:"rainfall,omitempty" yaml:"rainfall" bson:"rainfall" db:"rainfall"`
}

func (x *RainPeriod) Reset() {
	*x = RainPeriod{}
	mi := &file_city_water_input_v1_water_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RainPeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RainPeriod) ProtoMessage() {}

func (x *RainPeriod) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_input_v1_water_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RainPeriod.ProtoReflect.Descriptor instead.
func (*RainPeriod) Descriptor() ([]byte, []int) {
	return file_city_water_input_v1_water_proto_rawDescGZIP(), []int{0}
}

func (x *RainPeriod) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *RainPeriod) GetRainfall() float64 {
	if x != nil {
		return x.Rainfall
	}
	return 0
}

// 全天降雨情况，在数据库中体现为一条数据
type Rain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rains []*RainPeriod `protobuf:"bytes,1,rep,name=rains,proto3" json:"rains,omitempty" db:"rains" yaml:"rains" bson:"rains"`
}

func (x *Rain) Reset() {
	*x = Rain{}
	mi := &file_city_water_input_v1_water_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rain) ProtoMessage() {}

func (x *Rain) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_input_v1_water_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rain.ProtoReflect.Descriptor instead.
func (*Rain) Descriptor() ([]byte, []int) {
	return file_city_water_input_v1_water_proto_rawDescGZIP(), []int{1}
}

func (x *Rain) GetRains() []*RainPeriod {
	if x != nil {
		return x.Rains
	}
	return nil
}

var File_city_water_input_v1_water_proto protoreflect.FileDescriptor

var file_city_water_input_v1_water_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x3e, 0x0a, 0x0a, 0x52, 0x61, 0x69, 0x6e, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61,
	0x69, 0x6e, 0x66, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x61,
	0x69, 0x6e, 0x66, 0x61, 0x6c, 0x6c, 0x22, 0x3d, 0x0a, 0x04, 0x52, 0x61, 0x69, 0x6e, 0x12, 0x35,
	0x0a, 0x05, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x69, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x05,
	0x72, 0x61, 0x69, 0x6e, 0x73, 0x42, 0xd1, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x0a, 0x57, 0x61, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f,
	0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x57, 0x49, 0xaa, 0x02, 0x13, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x43, 0x69, 0x74, 0x79, 0x5c,
	0x57, 0x61, 0x74, 0x65, 0x72, 0x5c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1f, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x57, 0x61, 0x74, 0x65, 0x72, 0x5c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x16, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x57, 0x61, 0x74, 0x65, 0x72, 0x3a, 0x3a,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_city_water_input_v1_water_proto_rawDescOnce sync.Once
	file_city_water_input_v1_water_proto_rawDescData = file_city_water_input_v1_water_proto_rawDesc
)

func file_city_water_input_v1_water_proto_rawDescGZIP() []byte {
	file_city_water_input_v1_water_proto_rawDescOnce.Do(func() {
		file_city_water_input_v1_water_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_water_input_v1_water_proto_rawDescData)
	})
	return file_city_water_input_v1_water_proto_rawDescData
}

var file_city_water_input_v1_water_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_water_input_v1_water_proto_goTypes = []any{
	(*RainPeriod)(nil), // 0: city.water.input.v1.RainPeriod
	(*Rain)(nil),       // 1: city.water.input.v1.Rain
}
var file_city_water_input_v1_water_proto_depIdxs = []int32{
	0, // 0: city.water.input.v1.Rain.rains:type_name -> city.water.input.v1.RainPeriod
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_city_water_input_v1_water_proto_init() }
func file_city_water_input_v1_water_proto_init() {
	if File_city_water_input_v1_water_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_water_input_v1_water_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_water_input_v1_water_proto_goTypes,
		DependencyIndexes: file_city_water_input_v1_water_proto_depIdxs,
		MessageInfos:      file_city_water_input_v1_water_proto_msgTypes,
	}.Build()
	File_city_water_input_v1_water_proto = out.File
	file_city_water_input_v1_water_proto_rawDesc = nil
	file_city_water_input_v1_water_proto_goTypes = nil
	file_city_water_input_v1_water_proto_depIdxs = nil
}
