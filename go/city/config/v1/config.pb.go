// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/config/v1/config.proto

package configv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MongoDB配置
type MongoPath struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 数据库名
	Db string `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty" bson:"db" db:"db" yaml:"db"`
	// 集合名
	Col           string `protobuf:"bytes,2,opt,name=col,proto3" json:"col,omitempty" bson:"col" db:"col" yaml:"col"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MongoPath) Reset() {
	*x = MongoPath{}
	mi := &file_city_config_v1_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MongoPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoPath) ProtoMessage() {}

func (x *MongoPath) ProtoReflect() protoreflect.Message {
	mi := &file_city_config_v1_config_proto_msgTypes[0]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sql           string                 `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty" bson:"sql" db:"sql" yaml:"sql"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutputTarget) Reset() {
	*x = OutputTarget{}
	mi := &file_city_config_v1_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutputTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputTarget) ProtoMessage() {}

func (x *OutputTarget) ProtoReflect() protoreflect.Message {
	mi := &file_city_config_v1_config_proto_msgTypes[1]
	if x != nil {
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

const file_city_config_v1_config_proto_rawDesc = "" +
	"\n" +
	"\x1bcity/config/v1/config.proto\x12\x0ecity.config.v1\"-\n" +
	"\tMongoPath\x12\x0e\n" +
	"\x02db\x18\x01 \x01(\tR\x02db\x12\x10\n" +
	"\x03col\x18\x02 \x01(\tR\x03col\" \n" +
	"\fOutputTarget\x12\x10\n" +
	"\x03sql\x18\x01 \x01(\tR\x03sqlB\xb4\x01\n" +
	"\x12com.city.config.v1B\vConfigProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/config/v1;configv1\xa2\x02\x03CCX\xaa\x02\x0eCity.Config.V1\xca\x02\x0eCity\\Config\\V1\xe2\x02\x1aCity\\Config\\V1\\GPBMetadata\xea\x02\x10City::Config::V1b\x06proto3"

var (
	file_city_config_v1_config_proto_rawDescOnce sync.Once
	file_city_config_v1_config_proto_rawDescData []byte
)

func file_city_config_v1_config_proto_rawDescGZIP() []byte {
	file_city_config_v1_config_proto_rawDescOnce.Do(func() {
		file_city_config_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_config_v1_config_proto_rawDesc), len(file_city_config_v1_config_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_config_v1_config_proto_rawDesc), len(file_city_config_v1_config_proto_rawDesc)),
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
	file_city_config_v1_config_proto_goTypes = nil
	file_city_config_v1_config_proto_depIdxs = nil
}
