// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: city/agent/v2/agent.proto

package agentv2

import (
	v2 "git.fiblab.net/sim/protos/go/city/geo/v2"
	v21 "git.fiblab.net/sim/protos/go/city/trip/v2"
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

type AgentType int32

const (
	// 未指定
	AgentType_AGENT_TYPE_UNSPECIFIED AgentType = 0
	// 人
	AgentType_AGENT_TYPE_PERSON AgentType = 1
	// 私家车
	AgentType_AGENT_TYPE_PRIVATE_CAR AgentType = 2
	// 公交车
	AgentType_AGENT_TYPE_BUS AgentType = 3
)

// Enum value maps for AgentType.
var (
	AgentType_name = map[int32]string{
		0: "AGENT_TYPE_UNSPECIFIED",
		1: "AGENT_TYPE_PERSON",
		2: "AGENT_TYPE_PRIVATE_CAR",
		3: "AGENT_TYPE_BUS",
	}
	AgentType_value = map[string]int32{
		"AGENT_TYPE_UNSPECIFIED": 0,
		"AGENT_TYPE_PERSON":      1,
		"AGENT_TYPE_PRIVATE_CAR": 2,
		"AGENT_TYPE_BUS":         3,
	}
)

func (x AgentType) Enum() *AgentType {
	p := new(AgentType)
	*p = x
	return p
}

func (x AgentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentType) Descriptor() protoreflect.EnumDescriptor {
	return file_city_agent_v2_agent_proto_enumTypes[0].Descriptor()
}

func (AgentType) Type() protoreflect.EnumType {
	return &file_city_agent_v2_agent_proto_enumTypes[0]
}

func (x AgentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentType.Descriptor instead.
func (AgentType) EnumDescriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_proto_rawDescGZIP(), []int{0}
}

// 智能体属性（通用）
//
// Deprecated: Marked as deprecated in city/agent/v2/agent.proto.
type AgentAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 智能体类型
	Type AgentType `protobuf:"varint,1,opt,name=type,proto3,enum=city.agent.v2.AgentType" json:"type,omitempty" yaml:"type" bson:"type" db:"type"`
	// 单位: m，长度
	Length float64 `protobuf:"fixed64,2,opt,name=length,proto3" json:"length,omitempty" yaml:"length" bson:"length" db:"length"`
	// 单位: m，宽度
	Width float64 `protobuf:"fixed64,3,opt,name=width,proto3" json:"width,omitempty" yaml:"width" bson:"width" db:"width"`
	// 单位: m/s
	MaxSpeed float64 `protobuf:"fixed64,4,opt,name=max_speed,json=maxSpeed,proto3" json:"max_speed,omitempty" bson:"max_speed" db:"max_speed" yaml:"max_speed"`
	// 单位: m/s^2, 最大加速度（正值）
	MaxAcceleration float64 `protobuf:"fixed64,5,opt,name=max_acceleration,json=maxAcceleration,proto3" json:"max_acceleration,omitempty" yaml:"max_acceleration" bson:"max_acceleration" db:"max_acceleration"`
	// 单位: m/s^2, 最大减速度（负值）
	MaxBrakingAcceleration float64 `protobuf:"fixed64,6,opt,name=max_braking_acceleration,json=maxBrakingAcceleration,proto3" json:"max_braking_acceleration,omitempty" bson:"max_braking_acceleration" db:"max_braking_acceleration" yaml:"max_braking_acceleration"`
	// 单位: m/s^2, 一般加速度（正值），要求小于最大加速度
	UsualAcceleration float64 `protobuf:"fixed64,7,opt,name=usual_acceleration,json=usualAcceleration,proto3" json:"usual_acceleration,omitempty" yaml:"usual_acceleration" bson:"usual_acceleration" db:"usual_acceleration"`
	// 单位: m/s^2, 一般减速度（负值），要求大于最大减速度
	UsualBrakingAcceleration float64 `protobuf:"fixed64,8,opt,name=usual_braking_acceleration,json=usualBrakingAcceleration,proto3" json:"usual_braking_acceleration,omitempty" yaml:"usual_braking_acceleration" bson:"usual_braking_acceleration" db:"usual_braking_acceleration"`
}

func (x *AgentAttribute) Reset() {
	*x = AgentAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentAttribute) ProtoMessage() {}

func (x *AgentAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentAttribute.ProtoReflect.Descriptor instead.
func (*AgentAttribute) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_proto_rawDescGZIP(), []int{0}
}

func (x *AgentAttribute) GetType() AgentType {
	if x != nil {
		return x.Type
	}
	return AgentType_AGENT_TYPE_UNSPECIFIED
}

func (x *AgentAttribute) GetLength() float64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *AgentAttribute) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *AgentAttribute) GetMaxSpeed() float64 {
	if x != nil {
		return x.MaxSpeed
	}
	return 0
}

func (x *AgentAttribute) GetMaxAcceleration() float64 {
	if x != nil {
		return x.MaxAcceleration
	}
	return 0
}

func (x *AgentAttribute) GetMaxBrakingAcceleration() float64 {
	if x != nil {
		return x.MaxBrakingAcceleration
	}
	return 0
}

func (x *AgentAttribute) GetUsualAcceleration() float64 {
	if x != nil {
		return x.UsualAcceleration
	}
	return 0
}

func (x *AgentAttribute) GetUsualBrakingAcceleration() float64 {
	if x != nil {
		return x.UsualBrakingAcceleration
	}
	return 0
}

// 车辆附加属性
type VehicleAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 单位: m, 完成变道所需路程
	LaneChangeLength float64 `protobuf:"fixed64,1,opt,name=lane_change_length,json=laneChangeLength,proto3" json:"lane_change_length,omitempty" db:"lane_change_length" yaml:"lane_change_length" bson:"lane_change_length"`
	// 单位：米，本车距离前车的最小距离
	MinGap float64 `protobuf:"fixed64,2,opt,name=min_gap,json=minGap,proto3" json:"min_gap,omitempty" yaml:"min_gap" bson:"min_gap" db:"min_gap"`
}

func (x *VehicleAttribute) Reset() {
	*x = VehicleAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleAttribute) ProtoMessage() {}

func (x *VehicleAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleAttribute.ProtoReflect.Descriptor instead.
func (*VehicleAttribute) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_proto_rawDescGZIP(), []int{1}
}

func (x *VehicleAttribute) GetLaneChangeLength() float64 {
	if x != nil {
		return x.LaneChangeLength
	}
	return 0
}

func (x *VehicleAttribute) GetMinGap() float64 {
	if x != nil {
		return x.MinGap
	}
	return 0
}

// 公交车附加属性
type BusAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 公交线路ID
	LineId int32 `protobuf:"varint,1,opt,name=line_id,json=lineId,proto3" json:"line_id,omitempty" yaml:"line_id" bson:"line_id" db:"line_id"`
	// 公交车容量
	Capacity int32 `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty" yaml:"capacity" bson:"capacity" db:"capacity"`
}

func (x *BusAttribute) Reset() {
	*x = BusAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusAttribute) ProtoMessage() {}

func (x *BusAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusAttribute.ProtoReflect.Descriptor instead.
func (*BusAttribute) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_proto_rawDescGZIP(), []int{2}
}

func (x *BusAttribute) GetLineId() int32 {
	if x != nil {
		return x.LineId
	}
	return 0
}

func (x *BusAttribute) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

// 自行车附加属性
type BikeAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BikeAttribute) Reset() {
	*x = BikeAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BikeAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BikeAttribute) ProtoMessage() {}

func (x *BikeAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BikeAttribute.ProtoReflect.Descriptor instead.
func (*BikeAttribute) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_proto_rawDescGZIP(), []int{3}
}

// 智能体
//
// Deprecated: Marked as deprecated in city/agent/v2/agent.proto.
type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 智能体ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id" bson:"id" db:"id"`
	// 参数
	Attribute *AgentAttribute `protobuf:"bytes,2,opt,name=attribute,proto3" json:"attribute,omitempty" bson:"attribute" db:"attribute" yaml:"attribute"`
	// 初始位置
	Home *v2.Position `protobuf:"bytes,3,opt,name=home,proto3" json:"home,omitempty" yaml:"home" bson:"home" db:"home"`
	// 初始日程
	Schedules []*v21.Schedule `protobuf:"bytes,4,rep,name=schedules,proto3" json:"schedules,omitempty" yaml:"schedules" bson:"schedules" db:"schedules"`
	// 车辆附加属性
	VehicleAttribute *VehicleAttribute `protobuf:"bytes,7,opt,name=vehicle_attribute,json=vehicleAttribute,proto3,oneof" json:"vehicle_attribute,omitempty" yaml:"vehicle_attribute" bson:"vehicle_attribute" db:"vehicle_attribute"`
	// 公交车附加属性
	BusAttribute *BusAttribute `protobuf:"bytes,8,opt,name=bus_attribute,json=busAttribute,proto3,oneof" json:"bus_attribute,omitempty" yaml:"bus_attribute" bson:"bus_attribute" db:"bus_attribute"`
	// 自行车附加属性
	BikeAttribute *BikeAttribute `protobuf:"bytes,9,opt,name=bike_attribute,json=bikeAttribute,proto3,oneof" json:"bike_attribute,omitempty" yaml:"bike_attribute" bson:"bike_attribute" db:"bike_attribute"`
	// [可空] 额外的标签（例如：抢修车类型->电网）
	Labels map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"labels" bson:"labels" db:"labels"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_proto_rawDescGZIP(), []int{4}
}

func (x *Agent) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Agent) GetAttribute() *AgentAttribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

func (x *Agent) GetHome() *v2.Position {
	if x != nil {
		return x.Home
	}
	return nil
}

func (x *Agent) GetSchedules() []*v21.Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

func (x *Agent) GetVehicleAttribute() *VehicleAttribute {
	if x != nil {
		return x.VehicleAttribute
	}
	return nil
}

func (x *Agent) GetBusAttribute() *BusAttribute {
	if x != nil {
		return x.BusAttribute
	}
	return nil
}

func (x *Agent) GetBikeAttribute() *BikeAttribute {
	if x != nil {
		return x.BikeAttribute
	}
	return nil
}

func (x *Agent) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// 智能体集合，对应一个智能体pb文件或一个智能体mongodb collection
//
// Deprecated: Marked as deprecated in city/agent/v2/agent.proto.
type Agents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agents []*Agent `protobuf:"bytes,1,rep,name=agents,proto3" json:"agents,omitempty" yaml:"agents" bson:"agents" db:"agents"`
}

func (x *Agents) Reset() {
	*x = Agents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agents) ProtoMessage() {}

func (x *Agents) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agents.ProtoReflect.Descriptor instead.
func (*Agents) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_proto_rawDescGZIP(), []int{5}
}

func (x *Agents) GetAgents() []*Agent {
	if x != nil {
		return x.Agents
	}
	return nil
}

var File_city_agent_v2_agent_proto protoreflect.FileDescriptor

var file_city_agent_v2_agent_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x15, 0x63, 0x69, 0x74, 0x79,
	0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x2f, 0x76, 0x32, 0x2f,
	0x74, 0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x02, 0x0a, 0x0e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x2c, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0f, 0x6d, 0x61, 0x78, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x18, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x72, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x42, 0x72, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41,
	0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x75,
	0x73, 0x75, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x41, 0x63,
	0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x1a, 0x75, 0x73,
	0x75, 0x61, 0x6c, 0x5f, 0x62, 0x72, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x18,
	0x75, 0x73, 0x75, 0x61, 0x6c, 0x42, 0x72, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x65,
	0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x59, 0x0a, 0x10,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x6c, 0x61,
	0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x6d, 0x69, 0x6e, 0x47, 0x61, 0x70, 0x22, 0x43, 0x0a, 0x0c, 0x42, 0x75, 0x73, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x0f, 0x0a, 0x0d,
	0x42, 0x69, 0x6b, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0xcd, 0x04,
	0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x76,
	0x32, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x11, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x48, 0x00, 0x52, 0x10, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x42, 0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x48, 0x01, 0x52, 0x0c,
	0x62, 0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x48, 0x0a, 0x0e, 0x62, 0x69, 0x6b, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x69, 0x6b, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x48, 0x02, 0x52, 0x0d, 0x62, 0x69, 0x6b, 0x65, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x02,
	0x18, 0x01, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x62, 0x75, 0x73,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x62,
	0x69, 0x6b, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x3a, 0x0a,
	0x06, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x02, 0x18, 0x01, 0x2a, 0x6e, 0x0a, 0x09, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x47, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x41, 0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x53, 0x10, 0x03, 0x42, 0xa9, 0x01, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x42,
	0x0a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79,
	0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x76,
	0x32, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x19, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_agent_v2_agent_proto_rawDescOnce sync.Once
	file_city_agent_v2_agent_proto_rawDescData = file_city_agent_v2_agent_proto_rawDesc
)

func file_city_agent_v2_agent_proto_rawDescGZIP() []byte {
	file_city_agent_v2_agent_proto_rawDescOnce.Do(func() {
		file_city_agent_v2_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_agent_v2_agent_proto_rawDescData)
	})
	return file_city_agent_v2_agent_proto_rawDescData
}

var file_city_agent_v2_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_city_agent_v2_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_city_agent_v2_agent_proto_goTypes = []interface{}{
	(AgentType)(0),           // 0: city.agent.v2.AgentType
	(*AgentAttribute)(nil),   // 1: city.agent.v2.AgentAttribute
	(*VehicleAttribute)(nil), // 2: city.agent.v2.VehicleAttribute
	(*BusAttribute)(nil),     // 3: city.agent.v2.BusAttribute
	(*BikeAttribute)(nil),    // 4: city.agent.v2.BikeAttribute
	(*Agent)(nil),            // 5: city.agent.v2.Agent
	(*Agents)(nil),           // 6: city.agent.v2.Agents
	nil,                      // 7: city.agent.v2.Agent.LabelsEntry
	(*v2.Position)(nil),      // 8: city.geo.v2.Position
	(*v21.Schedule)(nil),     // 9: city.trip.v2.Schedule
}
var file_city_agent_v2_agent_proto_depIdxs = []int32{
	0, // 0: city.agent.v2.AgentAttribute.type:type_name -> city.agent.v2.AgentType
	1, // 1: city.agent.v2.Agent.attribute:type_name -> city.agent.v2.AgentAttribute
	8, // 2: city.agent.v2.Agent.home:type_name -> city.geo.v2.Position
	9, // 3: city.agent.v2.Agent.schedules:type_name -> city.trip.v2.Schedule
	2, // 4: city.agent.v2.Agent.vehicle_attribute:type_name -> city.agent.v2.VehicleAttribute
	3, // 5: city.agent.v2.Agent.bus_attribute:type_name -> city.agent.v2.BusAttribute
	4, // 6: city.agent.v2.Agent.bike_attribute:type_name -> city.agent.v2.BikeAttribute
	7, // 7: city.agent.v2.Agent.labels:type_name -> city.agent.v2.Agent.LabelsEntry
	5, // 8: city.agent.v2.Agents.agents:type_name -> city.agent.v2.Agent
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_city_agent_v2_agent_proto_init() }
func file_city_agent_v2_agent_proto_init() {
	if File_city_agent_v2_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_agent_v2_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentAttribute); i {
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
		file_city_agent_v2_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleAttribute); i {
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
		file_city_agent_v2_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusAttribute); i {
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
		file_city_agent_v2_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BikeAttribute); i {
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
		file_city_agent_v2_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
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
		file_city_agent_v2_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agents); i {
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
	file_city_agent_v2_agent_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_agent_v2_agent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_agent_v2_agent_proto_goTypes,
		DependencyIndexes: file_city_agent_v2_agent_proto_depIdxs,
		EnumInfos:         file_city_agent_v2_agent_proto_enumTypes,
		MessageInfos:      file_city_agent_v2_agent_proto_msgTypes,
	}.Build()
	File_city_agent_v2_agent_proto = out.File
	file_city_agent_v2_agent_proto_rawDesc = nil
	file_city_agent_v2_agent_proto_goTypes = nil
	file_city_agent_v2_agent_proto_depIdxs = nil
}
