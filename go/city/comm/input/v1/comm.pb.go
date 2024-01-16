// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: city/comm/input/v1/comm.proto

package inputv1

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

// 本文件描述通信部分拓扑结构
// 三种节点类型
type NodeType int32

const (
	NodeType_NODE_TYPE_UNSPECIFIED  NodeType = 0
	NodeType_NODE_TYPE_INTERNET     NodeType = 1
	NodeType_NODE_TYPE_GATEWAY      NodeType = 2
	NodeType_NODE_TYPE_BASE_STATION NodeType = 3
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		0: "NODE_TYPE_UNSPECIFIED",
		1: "NODE_TYPE_INTERNET",
		2: "NODE_TYPE_GATEWAY",
		3: "NODE_TYPE_BASE_STATION",
	}
	NodeType_value = map[string]int32{
		"NODE_TYPE_UNSPECIFIED":  0,
		"NODE_TYPE_INTERNET":     1,
		"NODE_TYPE_GATEWAY":      2,
		"NODE_TYPE_BASE_STATION": 3,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_city_comm_input_v1_comm_proto_enumTypes[0].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_city_comm_input_v1_comm_proto_enumTypes[0]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_city_comm_input_v1_comm_proto_rawDescGZIP(), []int{0}
}

type BaseStationType int32

const (
	BaseStationType_BASE_STATION_TYPE_UNSPECIFIED BaseStationType = 0
	BaseStationType_BASE_STATION_TYPE_INDOOR      BaseStationType = 1
	BaseStationType_BASE_STATION_TYPE_OUTDOOR     BaseStationType = 2
)

// Enum value maps for BaseStationType.
var (
	BaseStationType_name = map[int32]string{
		0: "BASE_STATION_TYPE_UNSPECIFIED",
		1: "BASE_STATION_TYPE_INDOOR",
		2: "BASE_STATION_TYPE_OUTDOOR",
	}
	BaseStationType_value = map[string]int32{
		"BASE_STATION_TYPE_UNSPECIFIED": 0,
		"BASE_STATION_TYPE_INDOOR":      1,
		"BASE_STATION_TYPE_OUTDOOR":     2,
	}
)

func (x BaseStationType) Enum() *BaseStationType {
	p := new(BaseStationType)
	*p = x
	return p
}

func (x BaseStationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseStationType) Descriptor() protoreflect.EnumDescriptor {
	return file_city_comm_input_v1_comm_proto_enumTypes[1].Descriptor()
}

func (BaseStationType) Type() protoreflect.EnumType {
	return &file_city_comm_input_v1_comm_proto_enumTypes[1]
}

func (x BaseStationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseStationType.Descriptor instead.
func (BaseStationType) EnumDescriptor() ([]byte, []int) {
	return file_city_comm_input_v1_comm_proto_rawDescGZIP(), []int{1}
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id" bson:"id" db:"id"`
	Type NodeType `protobuf:"varint,2,opt,name=type,proto3,enum=city.comm.input.v1.NodeType" json:"type,omitempty" yaml:"type" bson:"type" db:"type"`
	// 父节点
	ParentId int32 `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty" yaml:"parent_id" bson:"parent_id" db:"parent_id"`
	// 子节点
	ChildrenIds []int32 `protobuf:"varint,4,rep,packed,name=children_ids,json=childrenIds,proto3" json:"children_ids,omitempty" yaml:"children_ids" bson:"children_ids" db:"children_ids"`
	// 节点经纬度位置
	Position *v2.Position `protobuf:"bytes,5,opt,name=position,proto3,oneof" json:"position,omitempty" yaml:"position" bson:"position" db:"position"`
	// 节点所在aoi
	AoiId *int32 `protobuf:"varint,6,opt,name=aoi_id,json=aoiId,proto3,oneof" json:"aoi_id,omitempty" bson:"aoi_id" db:"aoi_id" yaml:"aoi_id"`
	// 基站频段id
	FreqRangeId *int32 `protobuf:"varint,7,opt,name=freq_range_id,json=freqRangeId,proto3,oneof" json:"freq_range_id,omitempty" yaml:"freq_range_id" bson:"freq_range_id" db:"freq_range_id"`
	// 室内外基站类型
	BaseStationType *BaseStationType `protobuf:"varint,8,opt,name=base_station_type,json=baseStationType,proto3,enum=city.comm.input.v1.BaseStationType,oneof" json:"base_station_type,omitempty" yaml:"base_station_type" bson:"base_station_type" db:"base_station_type"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_input_v1_comm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_input_v1_comm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_city_comm_input_v1_comm_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetType() NodeType {
	if x != nil {
		return x.Type
	}
	return NodeType_NODE_TYPE_UNSPECIFIED
}

func (x *Node) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Node) GetChildrenIds() []int32 {
	if x != nil {
		return x.ChildrenIds
	}
	return nil
}

func (x *Node) GetPosition() *v2.Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Node) GetAoiId() int32 {
	if x != nil && x.AoiId != nil {
		return *x.AoiId
	}
	return 0
}

func (x *Node) GetFreqRangeId() int32 {
	if x != nil && x.FreqRangeId != nil {
		return *x.FreqRangeId
	}
	return 0
}

func (x *Node) GetBaseStationType() BaseStationType {
	if x != nil && x.BaseStationType != nil {
		return *x.BaseStationType
	}
	return BaseStationType_BASE_STATION_TYPE_UNSPECIFIED
}

// 抢修站
type RepairStation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" yaml:"id" bson:"id"`
	AoiId    int32        `protobuf:"varint,2,opt,name=aoi_id,json=aoiId,proto3" json:"aoi_id,omitempty" yaml:"aoi_id" bson:"aoi_id" db:"aoi_id"`
	Position *v2.Position `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty" yaml:"position" bson:"position" db:"position"`
}

func (x *RepairStation) Reset() {
	*x = RepairStation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_input_v1_comm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepairStation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepairStation) ProtoMessage() {}

func (x *RepairStation) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_input_v1_comm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepairStation.ProtoReflect.Descriptor instead.
func (*RepairStation) Descriptor() ([]byte, []int) {
	return file_city_comm_input_v1_comm_proto_rawDescGZIP(), []int{1}
}

func (x *RepairStation) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RepairStation) GetAoiId() int32 {
	if x != nil {
		return x.AoiId
	}
	return 0
}

func (x *RepairStation) GetPosition() *v2.Position {
	if x != nil {
		return x.Position
	}
	return nil
}

// 水泵
type Pump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id" bson:"id" db:"id"`
	Position *v2.Position `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty" yaml:"position" bson:"position" db:"position"`
}

func (x *Pump) Reset() {
	*x = Pump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_input_v1_comm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pump) ProtoMessage() {}

func (x *Pump) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_input_v1_comm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pump.ProtoReflect.Descriptor instead.
func (*Pump) Descriptor() ([]byte, []int) {
	return file_city_comm_input_v1_comm_proto_rawDescGZIP(), []int{2}
}

func (x *Pump) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pump) GetPosition() *v2.Position {
	if x != nil {
		return x.Position
	}
	return nil
}

// 终端通信需求
type CommDemand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	Demands []float64 `protobuf:"fixed64,2,rep,packed,name=demands,proto3" json:"demands,omitempty" yaml:"demands" bson:"demands" db:"demands"`
}

func (x *CommDemand) Reset() {
	*x = CommDemand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_input_v1_comm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommDemand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommDemand) ProtoMessage() {}

func (x *CommDemand) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_input_v1_comm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommDemand.ProtoReflect.Descriptor instead.
func (*CommDemand) Descriptor() ([]byte, []int) {
	return file_city_comm_input_v1_comm_proto_rawDescGZIP(), []int{3}
}

func (x *CommDemand) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommDemand) GetDemands() []float64 {
	if x != nil {
		return x.Demands
	}
	return nil
}

type Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes          []*Node          `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty" yaml:"nodes" bson:"nodes" db:"nodes"`
	RepairStations []*RepairStation `protobuf:"bytes,2,rep,name=repair_stations,json=repairStations,proto3" json:"repair_stations,omitempty" yaml:"repair_stations" bson:"repair_stations" db:"repair_stations"`
	Pumps          []*Pump          `protobuf:"bytes,3,rep,name=pumps,proto3" json:"pumps,omitempty" yaml:"pumps" bson:"pumps" db:"pumps"`
}

func (x *Nodes) Reset() {
	*x = Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_input_v1_comm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nodes) ProtoMessage() {}

func (x *Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_input_v1_comm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nodes.ProtoReflect.Descriptor instead.
func (*Nodes) Descriptor() ([]byte, []int) {
	return file_city_comm_input_v1_comm_proto_rawDescGZIP(), []int{4}
}

func (x *Nodes) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Nodes) GetRepairStations() []*RepairStation {
	if x != nil {
		return x.RepairStations
	}
	return nil
}

func (x *Nodes) GetPumps() []*Pump {
	if x != nil {
		return x.Pumps
	}
	return nil
}

type CommDemands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommDemands []*CommDemand `protobuf:"bytes,1,rep,name=comm_demands,json=commDemands,proto3" json:"comm_demands,omitempty" yaml:"comm_demands" bson:"comm_demands" db:"comm_demands"`
}

func (x *CommDemands) Reset() {
	*x = CommDemands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_input_v1_comm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommDemands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommDemands) ProtoMessage() {}

func (x *CommDemands) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_input_v1_comm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommDemands.ProtoReflect.Descriptor instead.
func (*CommDemands) Descriptor() ([]byte, []int) {
	return file_city_comm_input_v1_comm_proto_rawDescGZIP(), []int{5}
}

func (x *CommDemands) GetCommDemands() []*CommDemand {
	if x != nil {
		return x.CommDemands
	}
	return nil
}

var File_city_comm_input_v1_comm_proto protoreflect.FileDescriptor

var file_city_comm_input_v1_comm_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x76, 0x32,
	0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x03, 0x0a, 0x04, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67,
	0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a,
	0x06, 0x61, 0x6f, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52,
	0x05, 0x61, 0x6f, 0x69, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x66, 0x72, 0x65,
	0x71, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x02, 0x52, 0x0b, 0x66, 0x72, 0x65, 0x71, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x54, 0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x48, 0x03, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6f, 0x69, 0x5f, 0x69, 0x64,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x69, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x61,
	0x69, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x6f, 0x69,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x6f, 0x69, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x04, 0x50, 0x75, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x36,
	0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x2e, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x4a, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x70, 0x61, 0x69, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x05,
	0x70, 0x75, 0x6d, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x6d, 0x70, 0x52, 0x05, 0x70, 0x75, 0x6d, 0x70, 0x73, 0x22, 0x50, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x6d, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x5f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x44, 0x65, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2a, 0x70,
	0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x4f,
	0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57,
	0x41, 0x59, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03,
	0x2a, 0x71, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x44, 0x4f,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x44, 0x4f, 0x4f,
	0x52, 0x10, 0x02, 0x42, 0xc7, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x09,
	0x43, 0x6f, 0x6d, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x49, 0xaa, 0x02, 0x12, 0x43, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x12, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x5c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f,
	0x6d, 0x6d, 0x3a, 0x3a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_comm_input_v1_comm_proto_rawDescOnce sync.Once
	file_city_comm_input_v1_comm_proto_rawDescData = file_city_comm_input_v1_comm_proto_rawDesc
)

func file_city_comm_input_v1_comm_proto_rawDescGZIP() []byte {
	file_city_comm_input_v1_comm_proto_rawDescOnce.Do(func() {
		file_city_comm_input_v1_comm_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_comm_input_v1_comm_proto_rawDescData)
	})
	return file_city_comm_input_v1_comm_proto_rawDescData
}

var file_city_comm_input_v1_comm_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_city_comm_input_v1_comm_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_city_comm_input_v1_comm_proto_goTypes = []interface{}{
	(NodeType)(0),         // 0: city.comm.input.v1.NodeType
	(BaseStationType)(0),  // 1: city.comm.input.v1.BaseStationType
	(*Node)(nil),          // 2: city.comm.input.v1.Node
	(*RepairStation)(nil), // 3: city.comm.input.v1.RepairStation
	(*Pump)(nil),          // 4: city.comm.input.v1.Pump
	(*CommDemand)(nil),    // 5: city.comm.input.v1.CommDemand
	(*Nodes)(nil),         // 6: city.comm.input.v1.Nodes
	(*CommDemands)(nil),   // 7: city.comm.input.v1.CommDemands
	(*v2.Position)(nil),   // 8: city.geo.v2.Position
}
var file_city_comm_input_v1_comm_proto_depIdxs = []int32{
	0, // 0: city.comm.input.v1.Node.type:type_name -> city.comm.input.v1.NodeType
	8, // 1: city.comm.input.v1.Node.position:type_name -> city.geo.v2.Position
	1, // 2: city.comm.input.v1.Node.base_station_type:type_name -> city.comm.input.v1.BaseStationType
	8, // 3: city.comm.input.v1.RepairStation.position:type_name -> city.geo.v2.Position
	8, // 4: city.comm.input.v1.Pump.position:type_name -> city.geo.v2.Position
	2, // 5: city.comm.input.v1.Nodes.nodes:type_name -> city.comm.input.v1.Node
	3, // 6: city.comm.input.v1.Nodes.repair_stations:type_name -> city.comm.input.v1.RepairStation
	4, // 7: city.comm.input.v1.Nodes.pumps:type_name -> city.comm.input.v1.Pump
	5, // 8: city.comm.input.v1.CommDemands.comm_demands:type_name -> city.comm.input.v1.CommDemand
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_city_comm_input_v1_comm_proto_init() }
func file_city_comm_input_v1_comm_proto_init() {
	if File_city_comm_input_v1_comm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_comm_input_v1_comm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_city_comm_input_v1_comm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepairStation); i {
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
		file_city_comm_input_v1_comm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pump); i {
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
		file_city_comm_input_v1_comm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommDemand); i {
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
		file_city_comm_input_v1_comm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nodes); i {
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
		file_city_comm_input_v1_comm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommDemands); i {
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
	file_city_comm_input_v1_comm_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_comm_input_v1_comm_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_comm_input_v1_comm_proto_goTypes,
		DependencyIndexes: file_city_comm_input_v1_comm_proto_depIdxs,
		EnumInfos:         file_city_comm_input_v1_comm_proto_enumTypes,
		MessageInfos:      file_city_comm_input_v1_comm_proto_msgTypes,
	}.Build()
	File_city_comm_input_v1_comm_proto = out.File
	file_city_comm_input_v1_comm_proto_rawDesc = nil
	file_city_comm_input_v1_comm_proto_goTypes = nil
	file_city_comm_input_v1_comm_proto_depIdxs = nil
}
