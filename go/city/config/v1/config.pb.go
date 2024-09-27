// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/config/v1/config.proto

package configv1

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

// MongoDB配置
type MongoPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据库名
	Db string `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty" bson:"db" db:"db" yaml:"db"`
	// 集合名
	Col string `protobuf:"bytes,2,opt,name=col,proto3" json:"col,omitempty" bson:"col" db:"col" yaml:"col"`
}

func (x *MongoPath) Reset() {
	*x = MongoPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_config_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoPath) ProtoMessage() {}

func (x *MongoPath) ProtoReflect() protoreflect.Message {
	mi := &file_city_config_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoPath.ProtoReflect.Descriptor instead.
func (*MongoPath) Descriptor() ([]byte, []int) {
	return file_city_config_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *MongoPath) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *MongoPath) GetCol() string {
	if x != nil {
		return x.Col
	}
	return ""
}

// 输出目标PostgreSQL
type OutputTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sql string `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty" bson:"sql" db:"sql" yaml:"sql"`
}

func (x *OutputTarget) Reset() {
	*x = OutputTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_config_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputTarget) ProtoMessage() {}

func (x *OutputTarget) ProtoReflect() protoreflect.Message {
	mi := &file_city_config_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputTarget.ProtoReflect.Descriptor instead.
func (*OutputTarget) Descriptor() ([]byte, []int) {
	return file_city_config_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *OutputTarget) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

var File_city_config_v1_config_proto protoreflect.FileDescriptor

var file_city_config_v1_config_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x2d, 0x0a,
	0x09, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6f, 0x6c, 0x22, 0x20, 0x0a, 0x0c,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x71, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x42, 0xb1,
	0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62,
	0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x58,
	0xaa, 0x02, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1a, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x10, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_config_v1_config_proto_rawDescOnce sync.Once
	file_city_config_v1_config_proto_rawDescData = file_city_config_v1_config_proto_rawDesc
)

func file_city_config_v1_config_proto_rawDescGZIP() []byte {
	file_city_config_v1_config_proto_rawDescOnce.Do(func() {
		file_city_config_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_config_v1_config_proto_rawDescData)
	})
	return file_city_config_v1_config_proto_rawDescData
}

var file_city_config_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_config_v1_config_proto_goTypes = []any{
	(*MongoPath)(nil),    // 0: city.config.v1.MongoPath
	(*OutputTarget)(nil), // 1: city.config.v1.OutputTarget
}
var file_city_config_v1_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_config_v1_config_proto_init() }
func file_city_config_v1_config_proto_init() {
	if File_city_config_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_config_v1_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MongoPath); i {
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
		file_city_config_v1_config_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OutputTarget); i {
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
			RawDescriptor: file_city_config_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_config_v1_config_proto_goTypes,
		DependencyIndexes: file_city_config_v1_config_proto_depIdxs,
		MessageInfos:      file_city_config_v1_config_proto_msgTypes,
	}.Build()
	File_city_config_v1_config_proto = out.File
	file_city_config_v1_config_proto_rawDesc = nil
	file_city_config_v1_config_proto_goTypes = nil
	file_city_config_v1_config_proto_depIdxs = nil
}
