// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: city/map/v2/road_service.proto

package mapv2

import (
	v1 "git.fiblab.net/sim/protos/v2/go/city/event/v1"
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

// 道路拥堵情况
// road congestion level
type RoadLevel int32

const (
	// 未指定
	// unspecified
	RoadLevel_ROAD_LEVEL_UNSPECIFIED RoadLevel = 0
	// 畅通
	// clear
	RoadLevel_ROAD_LEVEL_CLEAR RoadLevel = 1
	// 轻度拥堵
	// light load
	RoadLevel_ROAD_LEVEL_LIGHT_LOAD RoadLevel = 2
	// 中度拥堵
	// medium load
	RoadLevel_ROAD_LEVEL_MEDIUM_LOAD RoadLevel = 3
	// 重度拥堵
	// heavy load
	RoadLevel_ROAD_LEVEL_HEAVY_LOAD RoadLevel = 4
	// 极端拥堵
	// overload
	RoadLevel_ROAD_LEVEL_OVERLOAD RoadLevel = 5
	// 限行
	// restricted
	RoadLevel_ROAD_LEVEL_RESTRICTED RoadLevel = 6
)

// Enum value maps for RoadLevel.
var (
	RoadLevel_name = map[int32]string{
		0: "ROAD_LEVEL_UNSPECIFIED",
		1: "ROAD_LEVEL_CLEAR",
		2: "ROAD_LEVEL_LIGHT_LOAD",
		3: "ROAD_LEVEL_MEDIUM_LOAD",
		4: "ROAD_LEVEL_HEAVY_LOAD",
		5: "ROAD_LEVEL_OVERLOAD",
		6: "ROAD_LEVEL_RESTRICTED",
	}
	RoadLevel_value = map[string]int32{
		"ROAD_LEVEL_UNSPECIFIED": 0,
		"ROAD_LEVEL_CLEAR":       1,
		"ROAD_LEVEL_LIGHT_LOAD":  2,
		"ROAD_LEVEL_MEDIUM_LOAD": 3,
		"ROAD_LEVEL_HEAVY_LOAD":  4,
		"ROAD_LEVEL_OVERLOAD":    5,
		"ROAD_LEVEL_RESTRICTED":  6,
	}
)

func (x RoadLevel) Enum() *RoadLevel {
	p := new(RoadLevel)
	*p = x
	return p
}

func (x RoadLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoadLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_city_map_v2_road_service_proto_enumTypes[0].Descriptor()
}

func (RoadLevel) Type() protoreflect.EnumType {
	return &file_city_map_v2_road_service_proto_enumTypes[0]
}

func (x RoadLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoadLevel.Descriptor instead.
func (RoadLevel) EnumDescriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{0}
}

// 道路中断原因
// road interruption reason
type InterruptionReason int32

const (
	InterruptionReason_INTERRUPTION_REASON_UNSPECIFIED InterruptionReason = 0
	InterruptionReason_INTERRUPTION_REASON_RUINED      InterruptionReason = 1
	InterruptionReason_INTERRUPTION_REASON_CASCADE     InterruptionReason = 2
	InterruptionReason_INTERRUPTION_REASON_CONGESTION  InterruptionReason = 3
)

// Enum value maps for InterruptionReason.
var (
	InterruptionReason_name = map[int32]string{
		0: "INTERRUPTION_REASON_UNSPECIFIED",
		1: "INTERRUPTION_REASON_RUINED",
		2: "INTERRUPTION_REASON_CASCADE",
		3: "INTERRUPTION_REASON_CONGESTION",
	}
	InterruptionReason_value = map[string]int32{
		"INTERRUPTION_REASON_UNSPECIFIED": 0,
		"INTERRUPTION_REASON_RUINED":      1,
		"INTERRUPTION_REASON_CASCADE":     2,
		"INTERRUPTION_REASON_CONGESTION":  3,
	}
)

func (x InterruptionReason) Enum() *InterruptionReason {
	p := new(InterruptionReason)
	*p = x
	return p
}

func (x InterruptionReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InterruptionReason) Descriptor() protoreflect.EnumDescriptor {
	return file_city_map_v2_road_service_proto_enumTypes[1].Descriptor()
}

func (InterruptionReason) Type() protoreflect.EnumType {
	return &file_city_map_v2_road_service_proto_enumTypes[1]
}

func (x InterruptionReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InterruptionReason.Descriptor instead.
func (InterruptionReason) EnumDescriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{1}
}

// 查询道路信息请求
// Request for getting road information
type GetRoadRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 指定查询的道路ID列表，为空代表查询所有道路
	// List of targeted road IDs. If empty, it means querying all roads.
	RoadIds []int32 `protobuf:"varint,1,rep,packed,name=road_ids,json=roadIds,proto3" json:"road_ids,omitempty" bson:"road_ids" db:"road_ids" yaml:"road_ids"`
	// 是否要排除车道信息
	// Whether to exclude lane information
	ExcludeLane bool `protobuf:"varint,2,opt,name=exclude_lane,json=excludeLane,proto3" json:"exclude_lane,omitempty" yaml:"exclude_lane" bson:"exclude_lane" db:"exclude_lane"`
	// 是否要排除车道上的人的信息（仅在包含车道信息时有效）
	// Whether to exclude information about person in the lane (only valid when lane information is included)
	ExcludePerson bool `protobuf:"varint,3,opt,name=exclude_person,json=excludePerson,proto3" json:"exclude_person,omitempty" yaml:"exclude_person" bson:"exclude_person" db:"exclude_person"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoadRequest) Reset() {
	*x = GetRoadRequest{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoadRequest) ProtoMessage() {}

func (x *GetRoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoadRequest.ProtoReflect.Descriptor instead.
func (*GetRoadRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRoadRequest) GetRoadIds() []int32 {
	if x != nil {
		return x.RoadIds
	}
	return nil
}

func (x *GetRoadRequest) GetExcludeLane() bool {
	if x != nil {
		return x.ExcludeLane
	}
	return false
}

func (x *GetRoadRequest) GetExcludePerson() bool {
	if x != nil {
		return x.ExcludePerson
	}
	return false
}

// 查询道路信息响应
// Response of getting road information
type GetRoadResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 道路信息列表
	// List of road information
	States        []*RoadState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty" yaml:"states" bson:"states" db:"states"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoadResponse) Reset() {
	*x = GetRoadResponse{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoadResponse) ProtoMessage() {}

func (x *GetRoadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoadResponse.ProtoReflect.Descriptor instead.
func (*GetRoadResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRoadResponse) GetStates() []*RoadState {
	if x != nil {
		return x.States
	}
	return nil
}

type GetRuinInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRuinInfoRequest) Reset() {
	*x = GetRuinInfoRequest{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoRequest) ProtoMessage() {}

func (x *GetRuinInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuinInfoRequest.ProtoReflect.Descriptor instead.
func (*GetRuinInfoRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{2}
}

type RuinInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Num           int32                  `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty" yaml:"num" bson:"num" db:"num"`            // 损坏数量。Ruined number
	Ratio         float64                `protobuf:"fixed64,2,opt,name=ratio,proto3" json:"ratio,omitempty" yaml:"ratio" bson:"ratio" db:"ratio"` // 损坏占比。Ruined ratio
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuinInfo) Reset() {
	*x = RuinInfo{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuinInfo) ProtoMessage() {}

func (x *RuinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuinInfo.ProtoReflect.Descriptor instead.
func (*RuinInfo) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{3}
}

func (x *RuinInfo) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *RuinInfo) GetRatio() float64 {
	if x != nil {
		return x.Ratio
	}
	return 0
}

type GetRuinInfoResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 三级损伤信息
	// Three-level ruin information
	One           *RuinInfo `protobuf:"bytes,1,opt,name=one,proto3" json:"one,omitempty" yaml:"one" bson:"one" db:"one"`
	Two           *RuinInfo `protobuf:"bytes,2,opt,name=two,proto3" json:"two,omitempty" yaml:"two" bson:"two" db:"two"`
	Three         *RuinInfo `protobuf:"bytes,3,opt,name=three,proto3" json:"three,omitempty" yaml:"three" bson:"three" db:"three"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRuinInfoResponse) Reset() {
	*x = GetRuinInfoResponse{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoResponse) ProtoMessage() {}

func (x *GetRuinInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuinInfoResponse.ProtoReflect.Descriptor instead.
func (*GetRuinInfoResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetRuinInfoResponse) GetOne() *RuinInfo {
	if x != nil {
		return x.One
	}
	return nil
}

func (x *GetRuinInfoResponse) GetTwo() *RuinInfo {
	if x != nil {
		return x.Two
	}
	return nil
}

func (x *GetRuinInfoResponse) GetThree() *RuinInfo {
	if x != nil {
		return x.Three
	}
	return nil
}

type GetEventsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{5}
}

type GetEventsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Events        *v1.Events             `protobuf:"bytes,1,opt,name=events,proto3" json:"events,omitempty" yaml:"events" bson:"events" db:"events"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetEventsResponse) GetEvents() *v1.Events {
	if x != nil {
		return x.Events
	}
	return nil
}

// 道路状态
// road state
type RoadState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 道路ID
	// road ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" yaml:"id" bson:"id"`
	// 道路平均速度（m/s）
	// road average speed (m/s)
	AvgV float64 `protobuf:"fixed64,4,opt,name=avg_v,json=avgV,proto3" json:"avg_v,omitempty" db:"avg_v" yaml:"avg_v" bson:"avg_v"`
	// 道路拥堵情况
	// road congestion level
	Level RoadLevel `protobuf:"varint,2,opt,name=level,proto3,enum=city.map.v2.RoadLevel" json:"level,omitempty" yaml:"level" bson:"level" db:"level"`
	// 道路中断原因
	// road interruption reason
	Reason InterruptionReason `protobuf:"varint,3,opt,name=reason,proto3,enum=city.map.v2.InterruptionReason" json:"reason,omitempty" yaml:"reason" bson:"reason" db:"reason"`
	// 车道情况
	// lane state
	Lanes         []*LaneState `protobuf:"bytes,5,rep,name=lanes,proto3" json:"lanes,omitempty" yaml:"lanes" bson:"lanes" db:"lanes"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoadState) Reset() {
	*x = RoadState{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoadState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoadState) ProtoMessage() {}

func (x *RoadState) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoadState.ProtoReflect.Descriptor instead.
func (*RoadState) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{7}
}

func (x *RoadState) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoadState) GetAvgV() float64 {
	if x != nil {
		return x.AvgV
	}
	return 0
}

func (x *RoadState) GetLevel() RoadLevel {
	if x != nil {
		return x.Level
	}
	return RoadLevel_ROAD_LEVEL_UNSPECIFIED
}

func (x *RoadState) GetReason() InterruptionReason {
	if x != nil {
		return x.Reason
	}
	return InterruptionReason_INTERRUPTION_REASON_UNSPECIFIED
}

func (x *RoadState) GetLanes() []*LaneState {
	if x != nil {
		return x.Lanes
	}
	return nil
}

var File_city_map_v2_road_service_proto protoreflect.FileDescriptor

var file_city_map_v2_road_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f,
	0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x1a, 0x19, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d,
	0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x61, 0x64,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x61, 0x64,
	0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x6c,
	0x61, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x4c, 0x61, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x41, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x08, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x74,
	0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x03, 0x74, 0x77, 0x6f, 0x12, 0x2b, 0x0a, 0x05, 0x74, 0x68, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76,
	0x32, 0x2e, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x68, 0x72, 0x65,
	0x65, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x09, 0x52, 0x6f,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x76, 0x67, 0x5f, 0x76,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x61, 0x76, 0x67, 0x56, 0x12, 0x2c, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x61, 0x64, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x37, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x05, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32,
	0x2e, 0x4c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x6c, 0x61, 0x6e, 0x65,
	0x73, 0x2a, 0xc3, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x61, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1a, 0x0a, 0x16, 0x52, 0x4f, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x52,
	0x4f, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x4f, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16,
	0x52, 0x4f, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55,
	0x4d, 0x5f, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x4f, 0x41, 0x44,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x48, 0x45, 0x41, 0x56, 0x59, 0x5f, 0x4c, 0x4f, 0x41,
	0x44, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x4f, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15,
	0x52, 0x4f, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x52,
	0x49, 0x43, 0x54, 0x45, 0x44, 0x10, 0x06, 0x2a, 0x9e, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x1f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x55, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x53, 0x43, 0x41,
	0x44, 0x45, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x47,
	0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x32, 0xf1, 0x01, 0x0a, 0x0b, 0x52, 0x6f, 0x61,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76,
	0x32, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa4, 0x01, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32,
	0x42, 0x10, 0x52, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62,
	0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76,
	0x32, 0x3b, 0x6d, 0x61, 0x70, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x0b,
	0x43, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0b, 0x43, 0x69,
	0x74, 0x79, 0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x17, 0x43, 0x69, 0x74, 0x79,
	0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x3a,
	0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_map_v2_road_service_proto_rawDescOnce sync.Once
	file_city_map_v2_road_service_proto_rawDescData = file_city_map_v2_road_service_proto_rawDesc
)

func file_city_map_v2_road_service_proto_rawDescGZIP() []byte {
	file_city_map_v2_road_service_proto_rawDescOnce.Do(func() {
		file_city_map_v2_road_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_map_v2_road_service_proto_rawDescData)
	})
	return file_city_map_v2_road_service_proto_rawDescData
}

var file_city_map_v2_road_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_city_map_v2_road_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_city_map_v2_road_service_proto_goTypes = []any{
	(RoadLevel)(0),              // 0: city.map.v2.RoadLevel
	(InterruptionReason)(0),     // 1: city.map.v2.InterruptionReason
	(*GetRoadRequest)(nil),      // 2: city.map.v2.GetRoadRequest
	(*GetRoadResponse)(nil),     // 3: city.map.v2.GetRoadResponse
	(*GetRuinInfoRequest)(nil),  // 4: city.map.v2.GetRuinInfoRequest
	(*RuinInfo)(nil),            // 5: city.map.v2.RuinInfo
	(*GetRuinInfoResponse)(nil), // 6: city.map.v2.GetRuinInfoResponse
	(*GetEventsRequest)(nil),    // 7: city.map.v2.GetEventsRequest
	(*GetEventsResponse)(nil),   // 8: city.map.v2.GetEventsResponse
	(*RoadState)(nil),           // 9: city.map.v2.RoadState
	(*v1.Events)(nil),           // 10: city.event.v1.Events
	(*LaneState)(nil),           // 11: city.map.v2.LaneState
}
var file_city_map_v2_road_service_proto_depIdxs = []int32{
	9,  // 0: city.map.v2.GetRoadResponse.states:type_name -> city.map.v2.RoadState
	5,  // 1: city.map.v2.GetRuinInfoResponse.one:type_name -> city.map.v2.RuinInfo
	5,  // 2: city.map.v2.GetRuinInfoResponse.two:type_name -> city.map.v2.RuinInfo
	5,  // 3: city.map.v2.GetRuinInfoResponse.three:type_name -> city.map.v2.RuinInfo
	10, // 4: city.map.v2.GetEventsResponse.events:type_name -> city.event.v1.Events
	0,  // 5: city.map.v2.RoadState.level:type_name -> city.map.v2.RoadLevel
	1,  // 6: city.map.v2.RoadState.reason:type_name -> city.map.v2.InterruptionReason
	11, // 7: city.map.v2.RoadState.lanes:type_name -> city.map.v2.LaneState
	2,  // 8: city.map.v2.RoadService.GetRoad:input_type -> city.map.v2.GetRoadRequest
	4,  // 9: city.map.v2.RoadService.GetRuinInfo:input_type -> city.map.v2.GetRuinInfoRequest
	7,  // 10: city.map.v2.RoadService.GetEvents:input_type -> city.map.v2.GetEventsRequest
	3,  // 11: city.map.v2.RoadService.GetRoad:output_type -> city.map.v2.GetRoadResponse
	6,  // 12: city.map.v2.RoadService.GetRuinInfo:output_type -> city.map.v2.GetRuinInfoResponse
	8,  // 13: city.map.v2.RoadService.GetEvents:output_type -> city.map.v2.GetEventsResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_city_map_v2_road_service_proto_init() }
func file_city_map_v2_road_service_proto_init() {
	if File_city_map_v2_road_service_proto != nil {
		return
	}
	file_city_map_v2_lane_state_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_map_v2_road_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_map_v2_road_service_proto_goTypes,
		DependencyIndexes: file_city_map_v2_road_service_proto_depIdxs,
		EnumInfos:         file_city_map_v2_road_service_proto_enumTypes,
		MessageInfos:      file_city_map_v2_road_service_proto_msgTypes,
	}.Build()
	File_city_map_v2_road_service_proto = out.File
	file_city_map_v2_road_service_proto_rawDesc = nil
	file_city_map_v2_road_service_proto_goTypes = nil
	file_city_map_v2_road_service_proto_depIdxs = nil
}
