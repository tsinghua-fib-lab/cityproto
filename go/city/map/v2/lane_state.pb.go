// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: city/map/v2/lane_state.proto

package mapv2

import (
	v2 "git.fiblab.net/sim/protos/v2/go/city/person/v2"
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

// Lane状态
// Lane state
type LaneState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Lane ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// Lane上的人/车
	// person/vehicle on the lane
	Persons []*v2.PersonMotion `protobuf:"bytes,2,rep,name=persons,proto3" json:"persons,omitempty" bson:"persons" db:"persons" yaml:"persons"`
	// 平均速度（m/s）
	// average speed (m/s)
	AvgV float64 `protobuf:"fixed64,3,opt,name=avg_v,json=avgV,proto3" json:"avg_v,omitempty" bson:"avg_v" db:"avg_v" yaml:"avg_v"`
	// 是否限行
	// whether restricted
	Restriction bool `protobuf:"varint,4,opt,name=restriction,proto3" json:"restriction,omitempty" bson:"restriction" db:"restriction" yaml:"restriction"`
	// 交通灯状态
	// traffic light state
	LightState    LightState `protobuf:"varint,5,opt,name=light_state,json=lightState,proto3,enum=city.map.v2.LightState" json:"light_state,omitempty" bson:"light_state" db:"light_state" yaml:"light_state"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LaneState) Reset() {
	*x = LaneState{}
	mi := &file_city_map_v2_lane_state_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LaneState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneState) ProtoMessage() {}

func (x *LaneState) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_lane_state_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneState.ProtoReflect.Descriptor instead.
func (*LaneState) Descriptor() ([]byte, []int) {
	return file_city_map_v2_lane_state_proto_rawDescGZIP(), []int{0}
}

func (x *LaneState) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LaneState) GetPersons() []*v2.PersonMotion {
	if x != nil {
		return x.Persons
	}
	return nil
}

func (x *LaneState) GetAvgV() float64 {
	if x != nil {
		return x.AvgV
	}
	return 0
}

func (x *LaneState) GetRestriction() bool {
	if x != nil {
		return x.Restriction
	}
	return false
}

func (x *LaneState) GetLightState() LightState {
	if x != nil {
		return x.LightState
	}
	return LightState_LIGHT_STATE_UNSPECIFIED
}

var File_city_map_v2_lane_state_proto protoreflect.FileDescriptor

var file_city_map_v2_lane_state_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x61,
	0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x63, 0x69, 0x74,
	0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x09, 0x4c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x36, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x76, 0x67, 0x5f, 0x76,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x61, 0x76, 0x67, 0x56, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x0a, 0x0b, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76,
	0x32, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0xa2, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x42, 0x0e, 0x4c, 0x61,
	0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73,
	0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x3b, 0x6d, 0x61, 0x70, 0x76,
	0x32, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x4d,
	0x61, 0x70, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x4d, 0x61, 0x70,
	0x5c, 0x56, 0x32, 0xe2, 0x02, 0x17, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56,
	0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d,
	0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_map_v2_lane_state_proto_rawDescOnce sync.Once
	file_city_map_v2_lane_state_proto_rawDescData = file_city_map_v2_lane_state_proto_rawDesc
)

func file_city_map_v2_lane_state_proto_rawDescGZIP() []byte {
	file_city_map_v2_lane_state_proto_rawDescOnce.Do(func() {
		file_city_map_v2_lane_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_map_v2_lane_state_proto_rawDescData)
	})
	return file_city_map_v2_lane_state_proto_rawDescData
}

var file_city_map_v2_lane_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_city_map_v2_lane_state_proto_goTypes = []any{
	(*LaneState)(nil),       // 0: city.map.v2.LaneState
	(*v2.PersonMotion)(nil), // 1: city.person.v2.PersonMotion
	(LightState)(0),         // 2: city.map.v2.LightState
}
var file_city_map_v2_lane_state_proto_depIdxs = []int32{
	1, // 0: city.map.v2.LaneState.persons:type_name -> city.person.v2.PersonMotion
	2, // 1: city.map.v2.LaneState.light_state:type_name -> city.map.v2.LightState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_city_map_v2_lane_state_proto_init() }
func file_city_map_v2_lane_state_proto_init() {
	if File_city_map_v2_lane_state_proto != nil {
		return
	}
	file_city_map_v2_light_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_map_v2_lane_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_map_v2_lane_state_proto_goTypes,
		DependencyIndexes: file_city_map_v2_lane_state_proto_depIdxs,
		MessageInfos:      file_city_map_v2_lane_state_proto_msgTypes,
	}.Build()
	File_city_map_v2_lane_state_proto = out.File
	file_city_map_v2_lane_state_proto_rawDesc = nil
	file_city_map_v2_lane_state_proto_goTypes = nil
	file_city_map_v2_lane_state_proto_depIdxs = nil
}
