// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/person/v2/pedestrian.proto

package personv2

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

// 行人行为
// Pedestrian behavior
type PedestrianAction struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 行人id
	// Pedestrian id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// x方向速度（单位：米/秒）
	// x direction speed (unit: m/s)
	Vx float64 `protobuf:"fixed64,2,opt,name=vx,proto3" json:"vx,omitempty" bson:"vx" db:"vx" yaml:"vx"`
	// y方向速度（单位：米/秒）
	// y direction speed (unit: m/s)
	Vy            float64 `protobuf:"fixed64,3,opt,name=vy,proto3" json:"vy,omitempty" bson:"vy" db:"vy" yaml:"vy"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PedestrianAction) Reset() {
	*x = PedestrianAction{}
	mi := &file_city_person_v2_pedestrian_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PedestrianAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PedestrianAction) ProtoMessage() {}

func (x *PedestrianAction) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v2_pedestrian_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PedestrianAction.ProtoReflect.Descriptor instead.
func (*PedestrianAction) Descriptor() ([]byte, []int) {
	return file_city_person_v2_pedestrian_proto_rawDescGZIP(), []int{0}
}

func (x *PedestrianAction) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PedestrianAction) GetVx() float64 {
	if x != nil {
		return x.Vx
	}
	return 0
}

func (x *PedestrianAction) GetVy() float64 {
	if x != nil {
		return x.Vy
	}
	return 0
}

var File_city_person_v2_pedestrian_proto protoreflect.FileDescriptor

const file_city_person_v2_pedestrian_proto_rawDesc = "" +
	"\n" +
	"\x1fcity/person/v2/pedestrian.proto\x12\x0ecity.person.v2\"B\n" +
	"\x10PedestrianAction\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x0e\n" +
	"\x02vx\x18\x02 \x01(\x01R\x02vx\x12\x0e\n" +
	"\x02vy\x18\x03 \x01(\x01R\x02vyB\xb8\x01\n" +
	"\x12com.city.person.v2B\x0fPedestrianProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2b\x06proto3"

var (
	file_city_person_v2_pedestrian_proto_rawDescOnce sync.Once
	file_city_person_v2_pedestrian_proto_rawDescData []byte
)

func file_city_person_v2_pedestrian_proto_rawDescGZIP() []byte {
	file_city_person_v2_pedestrian_proto_rawDescOnce.Do(func() {
		file_city_person_v2_pedestrian_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_person_v2_pedestrian_proto_rawDesc), len(file_city_person_v2_pedestrian_proto_rawDesc)))
	})
	return file_city_person_v2_pedestrian_proto_rawDescData
}

var file_city_person_v2_pedestrian_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_city_person_v2_pedestrian_proto_goTypes = []any{
	(*PedestrianAction)(nil), // 0: city.person.v2.PedestrianAction
}
var file_city_person_v2_pedestrian_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_person_v2_pedestrian_proto_init() }
func file_city_person_v2_pedestrian_proto_init() {
	if File_city_person_v2_pedestrian_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_person_v2_pedestrian_proto_rawDesc), len(file_city_person_v2_pedestrian_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_person_v2_pedestrian_proto_goTypes,
		DependencyIndexes: file_city_person_v2_pedestrian_proto_depIdxs,
		MessageInfos:      file_city_person_v2_pedestrian_proto_msgTypes,
	}.Build()
	File_city_person_v2_pedestrian_proto = out.File
	file_city_person_v2_pedestrian_proto_goTypes = nil
	file_city_person_v2_pedestrian_proto_depIdxs = nil
}
