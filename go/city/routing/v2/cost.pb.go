// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: city/routing/v2/cost.proto

package routingv2

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

// 路径成本设置
type Cost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 目标拓扑元素（只支持道路Road）
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id" bson:"id" db:"id"`
	// 路径成本（单位：秒）
	Cost float64 `protobuf:"fixed64,2,opt,name=cost,proto3" json:"cost,omitempty" yaml:"cost" bson:"cost" db:"cost"`
	// 设置的时间（单位：秒）
	// 即设置几点几分的道路通行成本为cost
	// 为空表示设置全天通行成本均为cost
	Time *float64 `protobuf:"fixed64,3,opt,name=time,proto3,oneof" json:"time,omitempty" db:"time" yaml:"time" bson:"time"`
}

func (x *Cost) Reset() {
	*x = Cost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_routing_v2_cost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_cost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cost.ProtoReflect.Descriptor instead.
func (*Cost) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_cost_proto_rawDescGZIP(), []int{0}
}

func (x *Cost) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cost) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Cost) GetTime() float64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

var File_city_routing_v2_cost_proto protoreflect.FileDescriptor

var file_city_routing_v2_cost_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x22, 0x4c, 0x0a,
	0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0xb6, 0x01, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x32, 0x42, 0x09, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74,
	0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x52, 0x58, 0xaa, 0x02,
	0x0f, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x32,
	0xca, 0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c,
	0x56, 0x32, 0xe2, 0x02, 0x1b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x11, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_routing_v2_cost_proto_rawDescOnce sync.Once
	file_city_routing_v2_cost_proto_rawDescData = file_city_routing_v2_cost_proto_rawDesc
)

func file_city_routing_v2_cost_proto_rawDescGZIP() []byte {
	file_city_routing_v2_cost_proto_rawDescOnce.Do(func() {
		file_city_routing_v2_cost_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_routing_v2_cost_proto_rawDescData)
	})
	return file_city_routing_v2_cost_proto_rawDescData
}

var file_city_routing_v2_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_city_routing_v2_cost_proto_goTypes = []interface{}{
	(*Cost)(nil), // 0: city.routing.v2.Cost
}
var file_city_routing_v2_cost_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_routing_v2_cost_proto_init() }
func file_city_routing_v2_cost_proto_init() {
	if File_city_routing_v2_cost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_routing_v2_cost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cost); i {
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
	file_city_routing_v2_cost_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_routing_v2_cost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_routing_v2_cost_proto_goTypes,
		DependencyIndexes: file_city_routing_v2_cost_proto_depIdxs,
		MessageInfos:      file_city_routing_v2_cost_proto_msgTypes,
	}.Build()
	File_city_routing_v2_cost_proto = out.File
	file_city_routing_v2_cost_proto_rawDesc = nil
	file_city_routing_v2_cost_proto_goTypes = nil
	file_city_routing_v2_cost_proto_depIdxs = nil
}
