// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/map/v2/light.proto

package mapv2

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

// 交通灯的状态
// traffic light state
type LightState int32

const (
	// 未指定
	// unspecified
	LightState_LIGHT_STATE_UNSPECIFIED LightState = 0
	// 红灯
	// red light
	LightState_LIGHT_STATE_RED LightState = 1
	// 绿灯
	// green light
	LightState_LIGHT_STATE_GREEN LightState = 2
	// 黄灯
	// yellow light
	LightState_LIGHT_STATE_YELLOW LightState = 3
)

// Enum value maps for LightState.
var (
	LightState_name = map[int32]string{
		0: "LIGHT_STATE_UNSPECIFIED",
		1: "LIGHT_STATE_RED",
		2: "LIGHT_STATE_GREEN",
		3: "LIGHT_STATE_YELLOW",
	}
	LightState_value = map[string]int32{
		"LIGHT_STATE_UNSPECIFIED": 0,
		"LIGHT_STATE_RED":         1,
		"LIGHT_STATE_GREEN":       2,
		"LIGHT_STATE_YELLOW":      3,
	}
)

func (x LightState) Enum() *LightState {
	p := new(LightState)
	*p = x
	return p
}

func (x LightState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LightState) Descriptor() protoreflect.EnumDescriptor {
	return file_city_map_v2_light_proto_enumTypes[0].Descriptor()
}

func (LightState) Type() protoreflect.EnumType {
	return &file_city_map_v2_light_proto_enumTypes[0]
}

func (x LightState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LightState.Descriptor instead.
func (LightState) EnumDescriptor() ([]byte, []int) {
	return file_city_map_v2_light_proto_rawDescGZIP(), []int{0}
}

// 交通灯相位
// traffic light phase
type Phase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 相位持续时间，单位秒
	// Phase duration in seconds
	Duration float64 `protobuf:"fixed64,1,opt,name=duration,proto3" json:"duration,omitempty" db:"duration" yaml:"duration" bson:"duration"`
	// 描述该相位下每个lane的灯控情况，lane与Junction.lane_ids一一对应
	// The lighting control situation of each lane in this phase, and the lane
	// corresponds one-to-one with junction.lane_ids
	States []LightState `protobuf:"varint,2,rep,packed,name=states,proto3,enum=city.map.v2.LightState" json:"states,omitempty" yaml:"states" bson:"states" db:"states"`
}

func (x *Phase) Reset() {
	*x = Phase{}
	mi := &file_city_map_v2_light_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Phase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phase) ProtoMessage() {}

func (x *Phase) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_light_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phase.ProtoReflect.Descriptor instead.
func (*Phase) Descriptor() ([]byte, []int) {
	return file_city_map_v2_light_proto_rawDescGZIP(), []int{0}
}

func (x *Phase) GetDuration() float64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Phase) GetStates() []LightState {
	if x != nil {
		return x.States
	}
	return nil
}

type AvailablePhase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 描述最大压力信控的可行相位，由每个lane的灯控情况组成，lane与Junction.lane_ids一一对应
	// Describes the feasible phase for max pressure algorithm, consisting of the
	// lighting control situation for each lane in the junction, nd the lane
	// corresponds one-to-one with junction.lane_ids
	States []LightState `protobuf:"varint,1,rep,packed,name=states,proto3,enum=city.map.v2.LightState" json:"states,omitempty" yaml:"states" bson:"states" db:"states"`
}

func (x *AvailablePhase) Reset() {
	*x = AvailablePhase{}
	mi := &file_city_map_v2_light_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AvailablePhase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailablePhase) ProtoMessage() {}

func (x *AvailablePhase) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_light_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailablePhase.ProtoReflect.Descriptor instead.
func (*AvailablePhase) Descriptor() ([]byte, []int) {
	return file_city_map_v2_light_proto_rawDescGZIP(), []int{1}
}

func (x *AvailablePhase) GetStates() []LightState {
	if x != nil {
		return x.States
	}
	return nil
}

// 交通灯
// traffic light
type TrafficLight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所在路口id
	// ID of the junction where the traffic light is at
	JunctionId int32 `protobuf:"varint,1,opt,name=junction_id,json=junctionId,proto3" json:"junction_id,omitempty" db:"junction_id" yaml:"junction_id" bson:"junction_id"`
	// 相位循环的一个循环周期
	// One cycle of phase cycling
	Phases []*Phase `protobuf:"bytes,2,rep,name=phases,proto3" json:"phases,omitempty" yaml:"phases" bson:"phases" db:"phases"`
}

func (x *TrafficLight) Reset() {
	*x = TrafficLight{}
	mi := &file_city_map_v2_light_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrafficLight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficLight) ProtoMessage() {}

func (x *TrafficLight) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_light_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficLight.ProtoReflect.Descriptor instead.
func (*TrafficLight) Descriptor() ([]byte, []int) {
	return file_city_map_v2_light_proto_rawDescGZIP(), []int{2}
}

func (x *TrafficLight) GetJunctionId() int32 {
	if x != nil {
		return x.JunctionId
	}
	return 0
}

func (x *TrafficLight) GetPhases() []*Phase {
	if x != nil {
		return x.Phases
	}
	return nil
}

var File_city_map_v2_light_proto protoreflect.FileDescriptor

var file_city_map_v2_light_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x22, 0x54, 0x0a, 0x05, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x0e,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22,
	0x5b, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x06, 0x70, 0x68, 0x61, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x68, 0x61, 0x73, 0x65, 0x52, 0x06, 0x70, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2a, 0x6d, 0x0a, 0x0a,
	0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x49,
	0x47, 0x48, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x47, 0x48, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x52, 0x45, 0x45,
	0x4e, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x59, 0x45, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x03, 0x42, 0x9e, 0x01, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x42,
	0x0a, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x3b, 0x6d, 0x61, 0x70, 0x76, 0x32,
	0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x61,
	0x70, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x4d, 0x61, 0x70, 0x5c,
	0x56, 0x32, 0xe2, 0x02, 0x17, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56, 0x32,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x43,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_map_v2_light_proto_rawDescOnce sync.Once
	file_city_map_v2_light_proto_rawDescData = file_city_map_v2_light_proto_rawDesc
)

func file_city_map_v2_light_proto_rawDescGZIP() []byte {
	file_city_map_v2_light_proto_rawDescOnce.Do(func() {
		file_city_map_v2_light_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_map_v2_light_proto_rawDescData)
	})
	return file_city_map_v2_light_proto_rawDescData
}

var file_city_map_v2_light_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_city_map_v2_light_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_city_map_v2_light_proto_goTypes = []any{
	(LightState)(0),        // 0: city.map.v2.LightState
	(*Phase)(nil),          // 1: city.map.v2.Phase
	(*AvailablePhase)(nil), // 2: city.map.v2.AvailablePhase
	(*TrafficLight)(nil),   // 3: city.map.v2.TrafficLight
}
var file_city_map_v2_light_proto_depIdxs = []int32{
	0, // 0: city.map.v2.Phase.states:type_name -> city.map.v2.LightState
	0, // 1: city.map.v2.AvailablePhase.states:type_name -> city.map.v2.LightState
	1, // 2: city.map.v2.TrafficLight.phases:type_name -> city.map.v2.Phase
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_city_map_v2_light_proto_init() }
func file_city_map_v2_light_proto_init() {
	if File_city_map_v2_light_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_map_v2_light_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_map_v2_light_proto_goTypes,
		DependencyIndexes: file_city_map_v2_light_proto_depIdxs,
		EnumInfos:         file_city_map_v2_light_proto_enumTypes,
		MessageInfos:      file_city_map_v2_light_proto_msgTypes,
	}.Build()
	File_city_map_v2_light_proto = out.File
	file_city_map_v2_light_proto_rawDesc = nil
	file_city_map_v2_light_proto_goTypes = nil
	file_city_map_v2_light_proto_depIdxs = nil
}
