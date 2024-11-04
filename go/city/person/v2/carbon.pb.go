// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/person/v2/carbon.proto

package personv2

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

// 车辆瞬时碳排放信息
// Vehicle instantaneous carbon emission information
type VehicleCarbon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// delta distance (m)
	Ds float64 `protobuf:"fixed64,2,opt,name=ds,proto3" json:"ds,omitempty" bson:"ds" db:"ds" yaml:"ds"`
	// vehicle speed (m/s)
	V float64 `protobuf:"fixed64,3,opt,name=v,proto3" json:"v,omitempty" bson:"v" db:"v" yaml:"v"`
	// vehicle acceleration (m/s^2)
	A float64 `protobuf:"fixed64,4,opt,name=a,proto3" json:"a,omitempty" bson:"a" db:"a" yaml:"a"`
	// energy for acceleration (J)
	UAcc float64 `protobuf:"fixed64,5,opt,name=u_acc,json=uAcc,proto3" json:"u_acc,omitempty" bson:"u_acc" db:"u_acc" yaml:"u_acc"`
	// energy for rolling resistance (J)
	URoll float64 `protobuf:"fixed64,6,opt,name=u_roll,json=uRoll,proto3" json:"u_roll,omitempty" bson:"u_roll" db:"u_roll" yaml:"u_roll"`
	// energy for air resistance (J)
	UAero float64 `protobuf:"fixed64,7,opt,name=u_aero,json=uAero,proto3" json:"u_aero,omitempty" bson:"u_aero" db:"u_aero" yaml:"u_aero"`
	// C_D: drag coefficient
	CD float64 `protobuf:"fixed64,8,opt,name=c_d,json=cD,proto3" json:"c_d,omitempty" bson:"cd" db:"cd" yaml:"cd"`
}

func (x *VehicleCarbon) Reset() {
	*x = VehicleCarbon{}
	mi := &file_city_person_v2_carbon_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VehicleCarbon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleCarbon) ProtoMessage() {}

func (x *VehicleCarbon) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v2_carbon_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleCarbon.ProtoReflect.Descriptor instead.
func (*VehicleCarbon) Descriptor() ([]byte, []int) {
	return file_city_person_v2_carbon_proto_rawDescGZIP(), []int{0}
}

func (x *VehicleCarbon) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VehicleCarbon) GetDs() float64 {
	if x != nil {
		return x.Ds
	}
	return 0
}

func (x *VehicleCarbon) GetV() float64 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *VehicleCarbon) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *VehicleCarbon) GetUAcc() float64 {
	if x != nil {
		return x.UAcc
	}
	return 0
}

func (x *VehicleCarbon) GetURoll() float64 {
	if x != nil {
		return x.URoll
	}
	return 0
}

func (x *VehicleCarbon) GetUAero() float64 {
	if x != nil {
		return x.UAero
	}
	return 0
}

func (x *VehicleCarbon) GetCD() float64 {
	if x != nil {
		return x.CD
	}
	return 0
}

var File_city_person_v2_carbon_proto protoreflect.FileDescriptor

var file_city_person_v2_carbon_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x22, 0x9f, 0x01,
	0x0a, 0x0d, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x64, 0x73, 0x12,
	0x0c, 0x0a, 0x01, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a,
	0x01, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x61, 0x12, 0x13, 0x0a, 0x05, 0x75,
	0x5f, 0x61, 0x63, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x75, 0x41, 0x63, 0x63,
	0x12, 0x15, 0x0a, 0x06, 0x75, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x75, 0x52, 0x6f, 0x6c, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x75, 0x5f, 0x61, 0x65, 0x72,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x75, 0x41, 0x65, 0x72, 0x6f, 0x12, 0x0f,
	0x0a, 0x03, 0x63, 0x5f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x63, 0x44, 0x42,
	0xb4, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x0b, 0x43, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61,
	0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x76, 0x32, 0xa2, 0x02,
	0x03, 0x43, 0x50, 0x58, 0xaa, 0x02, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1a, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_person_v2_carbon_proto_rawDescOnce sync.Once
	file_city_person_v2_carbon_proto_rawDescData = file_city_person_v2_carbon_proto_rawDesc
)

func file_city_person_v2_carbon_proto_rawDescGZIP() []byte {
	file_city_person_v2_carbon_proto_rawDescOnce.Do(func() {
		file_city_person_v2_carbon_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_person_v2_carbon_proto_rawDescData)
	})
	return file_city_person_v2_carbon_proto_rawDescData
}

var file_city_person_v2_carbon_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_city_person_v2_carbon_proto_goTypes = []any{
	(*VehicleCarbon)(nil), // 0: city.person.v2.VehicleCarbon
}
var file_city_person_v2_carbon_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_person_v2_carbon_proto_init() }
func file_city_person_v2_carbon_proto_init() {
	if File_city_person_v2_carbon_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_person_v2_carbon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_person_v2_carbon_proto_goTypes,
		DependencyIndexes: file_city_person_v2_carbon_proto_depIdxs,
		MessageInfos:      file_city_person_v2_carbon_proto_msgTypes,
	}.Build()
	File_city_person_v2_carbon_proto = out.File
	file_city_person_v2_carbon_proto_rawDesc = nil
	file_city_person_v2_carbon_proto_goTypes = nil
	file_city_person_v2_carbon_proto_depIdxs = nil
}
