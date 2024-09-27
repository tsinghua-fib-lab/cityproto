// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/person/v1/person_runtime.proto

package personv1

import (
	v2 "git.fiblab.net/sim/protos/go/city/event/v2"
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

// 智能体运行时信息
type PersonRuntime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// person信息
	// person information
	Base *Person `protobuf:"bytes,1,opt,name=base,proto3,oneof" json:"base,omitempty" yaml:"base" bson:"base" db:"base"`
	// person运动信息
	// person motion information
	Motion *PersonMotion `protobuf:"bytes,2,opt,name=motion,proto3" json:"motion,omitempty" yaml:"motion" bson:"motion" db:"motion"`
	// 事件信息
	// event information
	Events []*v2.Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty" yaml:"events" bson:"events" db:"events"`
}

func (x *PersonRuntime) Reset() {
	*x = PersonRuntime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_runtime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonRuntime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRuntime) ProtoMessage() {}

func (x *PersonRuntime) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_runtime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRuntime.ProtoReflect.Descriptor instead.
func (*PersonRuntime) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_runtime_proto_rawDescGZIP(), []int{0}
}

func (x *PersonRuntime) GetBase() *Person {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *PersonRuntime) GetMotion() *PersonMotion {
	if x != nil {
		return x.Motion
	}
	return nil
}

func (x *PersonRuntime) GetEvents() []*v2.Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_city_person_v1_person_runtime_proto protoreflect.FileDescriptor

var file_city_person_v1_person_runtime_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a,
	0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x42, 0xb8, 0x01, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x42, 0x12, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62,
	0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x50, 0x58, 0xaa, 0x02, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_person_v1_person_runtime_proto_rawDescOnce sync.Once
	file_city_person_v1_person_runtime_proto_rawDescData = file_city_person_v1_person_runtime_proto_rawDesc
)

func file_city_person_v1_person_runtime_proto_rawDescGZIP() []byte {
	file_city_person_v1_person_runtime_proto_rawDescOnce.Do(func() {
		file_city_person_v1_person_runtime_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_person_v1_person_runtime_proto_rawDescData)
	})
	return file_city_person_v1_person_runtime_proto_rawDescData
}

var file_city_person_v1_person_runtime_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_city_person_v1_person_runtime_proto_goTypes = []any{
	(*PersonRuntime)(nil), // 0: city.person.v1.PersonRuntime
	(*Person)(nil),        // 1: city.person.v1.Person
	(*PersonMotion)(nil),  // 2: city.person.v1.PersonMotion
	(*v2.Event)(nil),      // 3: city.event.v2.Event
}
var file_city_person_v1_person_runtime_proto_depIdxs = []int32{
	1, // 0: city.person.v1.PersonRuntime.base:type_name -> city.person.v1.Person
	2, // 1: city.person.v1.PersonRuntime.motion:type_name -> city.person.v1.PersonMotion
	3, // 2: city.person.v1.PersonRuntime.events:type_name -> city.event.v2.Event
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_city_person_v1_person_runtime_proto_init() }
func file_city_person_v1_person_runtime_proto_init() {
	if File_city_person_v1_person_runtime_proto != nil {
		return
	}
	file_city_person_v1_motion_proto_init()
	file_city_person_v1_person_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_person_v1_person_runtime_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PersonRuntime); i {
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
	file_city_person_v1_person_runtime_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_person_v1_person_runtime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_person_v1_person_runtime_proto_goTypes,
		DependencyIndexes: file_city_person_v1_person_runtime_proto_depIdxs,
		MessageInfos:      file_city_person_v1_person_runtime_proto_msgTypes,
	}.Build()
	File_city_person_v1_person_runtime_proto = out.File
	file_city_person_v1_person_runtime_proto_rawDesc = nil
	file_city_person_v1_person_runtime_proto_goTypes = nil
	file_city_person_v1_person_runtime_proto_depIdxs = nil
}
