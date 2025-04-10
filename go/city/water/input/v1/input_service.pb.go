// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/water/input/v1/input_service.proto

package inputv1

import (
	v2 "git.fiblab.net/sim/protos/v2/go/city/map/v2"
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

type InitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	mi := &file_city_water_input_v1_input_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_input_v1_input_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_city_water_input_v1_input_service_proto_rawDescGZIP(), []int{0}
}

type InitResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 模拟器gRPC监听地址
	Address string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" bson:"address" db:"address" yaml:"address"`
	Control *Control `protobuf:"bytes,3,opt,name=control,proto3" json:"control,omitempty" bson:"control" db:"control" yaml:"control"`
	Rain    *Rain    `protobuf:"bytes,1,opt,name=rain,proto3" json:"rain,omitempty" bson:"rain" db:"rain" yaml:"rain"`
	// 仅包括header与roads
	Map           *v2.Map `protobuf:"bytes,4,opt,name=map,proto3" json:"map,omitempty" bson:"map" db:"map" yaml:"map"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	mi := &file_city_water_input_v1_input_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_input_v1_input_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_city_water_input_v1_input_service_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InitResponse) GetControl() *Control {
	if x != nil {
		return x.Control
	}
	return nil
}

func (x *InitResponse) GetRain() *Rain {
	if x != nil {
		return x.Rain
	}
	return nil
}

func (x *InitResponse) GetMap() *v2.Map {
	if x != nil {
		return x.Map
	}
	return nil
}

var File_city_water_input_v1_input_service_proto protoreflect.FileDescriptor

const file_city_water_input_v1_input_service_proto_rawDesc = "" +
	"\n" +
	"'city/water/input/v1/input_service.proto\x12\x13city.water.input.v1\x1a\x15city/map/v2/map.proto\x1a city/water/input/v1/config.proto\x1a\x1fcity/water/input/v1/water.proto\"\r\n" +
	"\vInitRequest\"\xb3\x01\n" +
	"\fInitResponse\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\x126\n" +
	"\acontrol\x18\x03 \x01(\v2\x1c.city.water.input.v1.ControlR\acontrol\x12-\n" +
	"\x04rain\x18\x01 \x01(\v2\x19.city.water.input.v1.RainR\x04rain\x12\"\n" +
	"\x03map\x18\x04 \x01(\v2\x10.city.map.v2.MapR\x03map2]\n" +
	"\fInputService\x12M\n" +
	"\x04Init\x12 .city.water.input.v1.InitRequest\x1a!.city.water.input.v1.InitResponse\"\x00B\xd8\x01\n" +
	"\x17com.city.water.input.v1B\x11InputServiceProtoP\x01Z;git.fiblab.net/sim/protos/v2/go/city/water/input/v1;inputv1\xa2\x02\x03CWI\xaa\x02\x13City.Water.Input.V1\xca\x02\x13City\\Water\\Input\\V1\xe2\x02\x1fCity\\Water\\Input\\V1\\GPBMetadata\xea\x02\x16City::Water::Input::V1b\x06proto3"

var (
	file_city_water_input_v1_input_service_proto_rawDescOnce sync.Once
	file_city_water_input_v1_input_service_proto_rawDescData []byte
)

func file_city_water_input_v1_input_service_proto_rawDescGZIP() []byte {
	file_city_water_input_v1_input_service_proto_rawDescOnce.Do(func() {
		file_city_water_input_v1_input_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_water_input_v1_input_service_proto_rawDesc), len(file_city_water_input_v1_input_service_proto_rawDesc)))
	})
	return file_city_water_input_v1_input_service_proto_rawDescData
}

var file_city_water_input_v1_input_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_water_input_v1_input_service_proto_goTypes = []any{
	(*InitRequest)(nil),  // 0: city.water.input.v1.InitRequest
	(*InitResponse)(nil), // 1: city.water.input.v1.InitResponse
	(*Control)(nil),      // 2: city.water.input.v1.Control
	(*Rain)(nil),         // 3: city.water.input.v1.Rain
	(*v2.Map)(nil),       // 4: city.map.v2.Map
}
var file_city_water_input_v1_input_service_proto_depIdxs = []int32{
	2, // 0: city.water.input.v1.InitResponse.control:type_name -> city.water.input.v1.Control
	3, // 1: city.water.input.v1.InitResponse.rain:type_name -> city.water.input.v1.Rain
	4, // 2: city.water.input.v1.InitResponse.map:type_name -> city.map.v2.Map
	0, // 3: city.water.input.v1.InputService.Init:input_type -> city.water.input.v1.InitRequest
	1, // 4: city.water.input.v1.InputService.Init:output_type -> city.water.input.v1.InitResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_city_water_input_v1_input_service_proto_init() }
func file_city_water_input_v1_input_service_proto_init() {
	if File_city_water_input_v1_input_service_proto != nil {
		return
	}
	file_city_water_input_v1_config_proto_init()
	file_city_water_input_v1_water_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_water_input_v1_input_service_proto_rawDesc), len(file_city_water_input_v1_input_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_water_input_v1_input_service_proto_goTypes,
		DependencyIndexes: file_city_water_input_v1_input_service_proto_depIdxs,
		MessageInfos:      file_city_water_input_v1_input_service_proto_msgTypes,
	}.Build()
	File_city_water_input_v1_input_service_proto = out.File
	file_city_water_input_v1_input_service_proto_goTypes = nil
	file_city_water_input_v1_input_service_proto_depIdxs = nil
}
