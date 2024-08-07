// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/agent/v2/motion.proto

package agentv2

import (
	v2 "git.fiblab.net/sim/protos/go/city/geo/v2"
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

// Agent（人）的运行时状态
type Status int32

const (
	// 未指定
	Status_STATUS_UNSPECIFIED Status = 0
	// 没有移动行为
	Status_STATUS_SLEEP Status = 1
	// 开车
	Status_STATUS_DRIVING Status = 2
	// 步行
	Status_STATUS_WALKING Status = 3
	// 室内行人
	Status_STATUS_CROWD Status = 4
	// 乘客
	Status_STATUS_PASSENGER Status = 5
	// 等待路径规划
	Status_STATUS_WAIT_ROUTE Status = 6
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SLEEP",
		2: "STATUS_DRIVING",
		3: "STATUS_WALKING",
		4: "STATUS_CROWD",
		5: "STATUS_PASSENGER",
		6: "STATUS_WAIT_ROUTE",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SLEEP":       1,
		"STATUS_DRIVING":     2,
		"STATUS_WALKING":     3,
		"STATUS_CROWD":       4,
		"STATUS_PASSENGER":   5,
		"STATUS_WAIT_ROUTE":  6,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_city_agent_v2_motion_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_city_agent_v2_motion_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_city_agent_v2_motion_proto_rawDescGZIP(), []int{0}
}

// Agent（人）的运动状态
//
// Deprecated: Marked as deprecated in city/agent/v2/motion.proto.
type AgentMotion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id" bson:"id" db:"id"`
	// 状态
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=city.agent.v2.Status" json:"status,omitempty" yaml:"status" bson:"status" db:"status"`
	// 位置（包含逻辑位置、XY位置、经纬度位置）
	Position *v2.Position `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty" yaml:"position" bson:"position" db:"position"`
	// 速度
	V float64 `protobuf:"fixed64,4,opt,name=v,proto3" json:"v,omitempty" bson:"v" db:"v" yaml:"v"`
	// 方向角（atan2计算得到的弧度）
	Direction float64 `protobuf:"fixed64,5,opt,name=direction,proto3" json:"direction,omitempty" yaml:"direction" bson:"direction" db:"direction"`
	// 活动描述
	Activity string `protobuf:"bytes,6,opt,name=activity,proto3" json:"activity,omitempty" yaml:"activity" bson:"activity" db:"activity"`
}

func (x *AgentMotion) Reset() {
	*x = AgentMotion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_motion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentMotion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentMotion) ProtoMessage() {}

func (x *AgentMotion) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_motion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentMotion.ProtoReflect.Descriptor instead.
func (*AgentMotion) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_motion_proto_rawDescGZIP(), []int{0}
}

func (x *AgentMotion) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AgentMotion) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *AgentMotion) GetPosition() *v2.Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *AgentMotion) GetV() float64 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *AgentMotion) GetDirection() float64 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *AgentMotion) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

var File_city_agent_v2_motion_proto protoreflect.FileDescriptor

var file_city_agent_v2_motion_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x15, 0x63, 0x69, 0x74,
	0x79, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x01, 0x76, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x3a, 0x02, 0x18, 0x01,
	0x2a, 0x99, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x4c,
	0x45, 0x45, 0x50, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x44, 0x52, 0x49, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x4c, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x4f, 0x57, 0x44, 0x10, 0x04, 0x12,
	0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x4e,
	0x47, 0x45, 0x52, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x57, 0x41, 0x49, 0x54, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x10, 0x06, 0x42, 0xaa, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x32, 0x42, 0x0b, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65,
	0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x0d, 0x43, 0x69,
	0x74, 0x79, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0d, 0x43, 0x69,
	0x74, 0x79, 0x5c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x19, 0x43, 0x69,
	0x74, 0x79, 0x5c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_city_agent_v2_motion_proto_rawDescOnce sync.Once
	file_city_agent_v2_motion_proto_rawDescData = file_city_agent_v2_motion_proto_rawDesc
)

func file_city_agent_v2_motion_proto_rawDescGZIP() []byte {
	file_city_agent_v2_motion_proto_rawDescOnce.Do(func() {
		file_city_agent_v2_motion_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_agent_v2_motion_proto_rawDescData)
	})
	return file_city_agent_v2_motion_proto_rawDescData
}

var file_city_agent_v2_motion_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_city_agent_v2_motion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_city_agent_v2_motion_proto_goTypes = []any{
	(Status)(0),         // 0: city.agent.v2.Status
	(*AgentMotion)(nil), // 1: city.agent.v2.AgentMotion
	(*v2.Position)(nil), // 2: city.geo.v2.Position
}
var file_city_agent_v2_motion_proto_depIdxs = []int32{
	0, // 0: city.agent.v2.AgentMotion.status:type_name -> city.agent.v2.Status
	2, // 1: city.agent.v2.AgentMotion.position:type_name -> city.geo.v2.Position
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_city_agent_v2_motion_proto_init() }
func file_city_agent_v2_motion_proto_init() {
	if File_city_agent_v2_motion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_agent_v2_motion_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AgentMotion); i {
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
			RawDescriptor: file_city_agent_v2_motion_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_agent_v2_motion_proto_goTypes,
		DependencyIndexes: file_city_agent_v2_motion_proto_depIdxs,
		EnumInfos:         file_city_agent_v2_motion_proto_enumTypes,
		MessageInfos:      file_city_agent_v2_motion_proto_msgTypes,
	}.Build()
	File_city_agent_v2_motion_proto = out.File
	file_city_agent_v2_motion_proto_rawDesc = nil
	file_city_agent_v2_motion_proto_goTypes = nil
	file_city_agent_v2_motion_proto_depIdxs = nil
}
