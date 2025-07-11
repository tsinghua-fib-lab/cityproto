// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/map/v2/road_service.proto

package mapv2

import (
	v1 "git.fiblab.net/sim/protos/v2/go/city/event/v1"
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
	ExcludeLane bool `protobuf:"varint,2,opt,name=exclude_lane,json=excludeLane,proto3" json:"exclude_lane,omitempty" bson:"exclude_lane" db:"exclude_lane" yaml:"exclude_lane"`
	// 是否要排除车道上的人的信息（仅在包含车道信息时有效）
	// Whether to exclude information about person in the lane (only valid when lane information is included)
	ExcludePerson bool `protobuf:"varint,3,opt,name=exclude_person,json=excludePerson,proto3" json:"exclude_person,omitempty" bson:"exclude_person" db:"exclude_person" yaml:"exclude_person"`
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
	States        []*RoadState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty" bson:"states" db:"states" yaml:"states"`
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

// 获取Road全局统计信息请求
// Request for getting road global statistics
type GetRoadGlobalStatisticsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoadGlobalStatisticsRequest) Reset() {
	*x = GetRoadGlobalStatisticsRequest{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoadGlobalStatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoadGlobalStatisticsRequest) ProtoMessage() {}

func (x *GetRoadGlobalStatisticsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetRoadGlobalStatisticsRequest.ProtoReflect.Descriptor instead.
func (*GetRoadGlobalStatisticsRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{2}
}

// 获取Road全局统计信息响应
// Response of getting road global statistics
type GetRoadGlobalStatisticsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 平均道路拥堵指数
	// average congestion index
	AvgRoadCongestionIndex float64 `protobuf:"fixed64,1,opt,name=avg_road_congestion_index,json=avgRoadCongestionIndex,proto3" json:"avg_road_congestion_index,omitempty" bson:"avg_road_congestion_index" db:"avg_road_congestion_index" yaml:"avg_road_congestion_index"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetRoadGlobalStatisticsResponse) Reset() {
	*x = GetRoadGlobalStatisticsResponse{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoadGlobalStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoadGlobalStatisticsResponse) ProtoMessage() {}

func (x *GetRoadGlobalStatisticsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetRoadGlobalStatisticsResponse.ProtoReflect.Descriptor instead.
func (*GetRoadGlobalStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoadGlobalStatisticsResponse) GetAvgRoadCongestionIndex() float64 {
	if x != nil {
		return x.AvgRoadCongestionIndex
	}
	return 0
}

type GetRuinInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRuinInfoRequest) Reset() {
	*x = GetRuinInfoRequest{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoRequest) ProtoMessage() {}

func (x *GetRuinInfoRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetRuinInfoRequest.ProtoReflect.Descriptor instead.
func (*GetRuinInfoRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{4}
}

type RuinInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Num           int32                  `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty" bson:"num" db:"num" yaml:"num"`            // 损坏数量。Ruined number
	Ratio         float64                `protobuf:"fixed64,2,opt,name=ratio,proto3" json:"ratio,omitempty" bson:"ratio" db:"ratio" yaml:"ratio"` // 损坏占比。Ruined ratio
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuinInfo) Reset() {
	*x = RuinInfo{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuinInfo) ProtoMessage() {}

func (x *RuinInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RuinInfo.ProtoReflect.Descriptor instead.
func (*RuinInfo) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{5}
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
	One           *RuinInfo `protobuf:"bytes,1,opt,name=one,proto3" json:"one,omitempty" bson:"one" db:"one" yaml:"one"`
	Two           *RuinInfo `protobuf:"bytes,2,opt,name=two,proto3" json:"two,omitempty" bson:"two" db:"two" yaml:"two"`
	Three         *RuinInfo `protobuf:"bytes,3,opt,name=three,proto3" json:"three,omitempty" bson:"three" db:"three" yaml:"three"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRuinInfoResponse) Reset() {
	*x = GetRuinInfoResponse{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoResponse) ProtoMessage() {}

func (x *GetRuinInfoResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetRuinInfoResponse.ProtoReflect.Descriptor instead.
func (*GetRuinInfoResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{6}
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
	mi := &file_city_map_v2_road_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{7}
}

type GetEventsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Events        *v1.Events             `protobuf:"bytes,1,opt,name=events,proto3" json:"events,omitempty" bson:"events" db:"events" yaml:"events"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[8]
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
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{8}
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
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// 当前进入道路的车辆数
	// current entering road vehicle count
	InVehicleCnt int32 `protobuf:"varint,7,opt,name=in_vehicle_cnt,json=inVehicleCnt,proto3" json:"in_vehicle_cnt,omitempty" bson:"in_vehicle_cnt" db:"in_vehicle_cnt" yaml:"in_vehicle_cnt"`
	// 当前离开道路的车辆数
	// current leaving road vehicle count
	OutVehicleCnt int32 `protobuf:"varint,8,opt,name=out_vehicle_cnt,json=outVehicleCnt,proto3" json:"out_vehicle_cnt,omitempty" bson:"out_vehicle_cnt" db:"out_vehicle_cnt" yaml:"out_vehicle_cnt"`
	// 当前道路车辆数
	// current road vehicle count
	VehicleCnt int32 `protobuf:"varint,9,opt,name=vehicle_cnt,json=vehicleCnt,proto3" json:"vehicle_cnt,omitempty" bson:"vehicle_cnt" db:"vehicle_cnt" yaml:"vehicle_cnt"`
	// 累计进入道路的车辆数
	// cumulative entering road vehicle count
	CumInVehicleCnt int32 `protobuf:"varint,10,opt,name=cum_in_vehicle_cnt,json=cumInVehicleCnt,proto3" json:"cum_in_vehicle_cnt,omitempty" bson:"cum_in_vehicle_cnt" db:"cum_in_vehicle_cnt" yaml:"cum_in_vehicle_cnt"`
	// 累计离开道路的车辆数
	// cumulative leaving road vehicle count
	CumOutVehicleCnt int32 `protobuf:"varint,11,opt,name=cum_out_vehicle_cnt,json=cumOutVehicleCnt,proto3" json:"cum_out_vehicle_cnt,omitempty" bson:"cum_out_vehicle_cnt" db:"cum_out_vehicle_cnt" yaml:"cum_out_vehicle_cnt"`
	// 道路平均速度（m/s）
	// road average speed (m/s)
	AvgV float64 `protobuf:"fixed64,4,opt,name=avg_v,json=avgV,proto3" json:"avg_v,omitempty" bson:"avg_v" db:"avg_v" yaml:"avg_v"`
	// 当前平均通行时间（s）
	// current average travel time (s)
	AvgTravelTime float64 `protobuf:"fixed64,12,opt,name=avg_travel_time,json=avgTravelTime,proto3" json:"avg_travel_time,omitempty" bson:"avg_travel_time" db:"avg_travel_time" yaml:"avg_travel_time"`
	// 道路拥堵情况
	// road congestion level
	Level RoadLevel `protobuf:"varint,2,opt,name=level,proto3,enum=city.map.v2.RoadLevel" json:"level,omitempty" bson:"level" db:"level" yaml:"level"`
	// 拥堵指数（最大限速/平均车速）
	// congestion index (max speed / average speed)
	CongestionIndex float64 `protobuf:"fixed64,13,opt,name=congestion_index,json=congestionIndex,proto3" json:"congestion_index,omitempty" bson:"congestion_index" db:"congestion_index" yaml:"congestion_index"`
	// 道路中断原因
	// road interruption reason
	Reason InterruptionReason `protobuf:"varint,3,opt,name=reason,proto3,enum=city.map.v2.InterruptionReason" json:"reason,omitempty" bson:"reason" db:"reason" yaml:"reason"`
	// 车道情况
	// lane state
	Lanes []*LaneState `protobuf:"bytes,5,rep,name=lanes,proto3" json:"lanes,omitempty" bson:"lanes" db:"lanes" yaml:"lanes"`
	// 道路最大限速
	// road max speed
	MaxV          float64 `protobuf:"fixed64,6,opt,name=max_v,json=maxV,proto3" json:"max_v,omitempty" bson:"max_v" db:"max_v" yaml:"max_v"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoadState) Reset() {
	*x = RoadState{}
	mi := &file_city_map_v2_road_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoadState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoadState) ProtoMessage() {}

func (x *RoadState) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_road_service_proto_msgTypes[9]
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
	return file_city_map_v2_road_service_proto_rawDescGZIP(), []int{9}
}

func (x *RoadState) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoadState) GetInVehicleCnt() int32 {
	if x != nil {
		return x.InVehicleCnt
	}
	return 0
}

func (x *RoadState) GetOutVehicleCnt() int32 {
	if x != nil {
		return x.OutVehicleCnt
	}
	return 0
}

func (x *RoadState) GetVehicleCnt() int32 {
	if x != nil {
		return x.VehicleCnt
	}
	return 0
}

func (x *RoadState) GetCumInVehicleCnt() int32 {
	if x != nil {
		return x.CumInVehicleCnt
	}
	return 0
}

func (x *RoadState) GetCumOutVehicleCnt() int32 {
	if x != nil {
		return x.CumOutVehicleCnt
	}
	return 0
}

func (x *RoadState) GetAvgV() float64 {
	if x != nil {
		return x.AvgV
	}
	return 0
}

func (x *RoadState) GetAvgTravelTime() float64 {
	if x != nil {
		return x.AvgTravelTime
	}
	return 0
}

func (x *RoadState) GetLevel() RoadLevel {
	if x != nil {
		return x.Level
	}
	return RoadLevel_ROAD_LEVEL_UNSPECIFIED
}

func (x *RoadState) GetCongestionIndex() float64 {
	if x != nil {
		return x.CongestionIndex
	}
	return 0
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

func (x *RoadState) GetMaxV() float64 {
	if x != nil {
		return x.MaxV
	}
	return 0
}

var File_city_map_v2_road_service_proto protoreflect.FileDescriptor

const file_city_map_v2_road_service_proto_rawDesc = "" +
	"\n" +
	"\x1ecity/map/v2/road_service.proto\x12\vcity.map.v2\x1a\x19city/event/v1/event.proto\x1a\x1ccity/map/v2/lane_state.proto\"u\n" +
	"\x0eGetRoadRequest\x12\x19\n" +
	"\broad_ids\x18\x01 \x03(\x05R\aroadIds\x12!\n" +
	"\fexclude_lane\x18\x02 \x01(\bR\vexcludeLane\x12%\n" +
	"\x0eexclude_person\x18\x03 \x01(\bR\rexcludePerson\"A\n" +
	"\x0fGetRoadResponse\x12.\n" +
	"\x06states\x18\x01 \x03(\v2\x16.city.map.v2.RoadStateR\x06states\" \n" +
	"\x1eGetRoadGlobalStatisticsRequest\"\\\n" +
	"\x1fGetRoadGlobalStatisticsResponse\x129\n" +
	"\x19avg_road_congestion_index\x18\x01 \x01(\x01R\x16avgRoadCongestionIndex\"\x14\n" +
	"\x12GetRuinInfoRequest\"2\n" +
	"\bRuinInfo\x12\x10\n" +
	"\x03num\x18\x01 \x01(\x05R\x03num\x12\x14\n" +
	"\x05ratio\x18\x02 \x01(\x01R\x05ratio\"\x94\x01\n" +
	"\x13GetRuinInfoResponse\x12'\n" +
	"\x03one\x18\x01 \x01(\v2\x15.city.map.v2.RuinInfoR\x03one\x12'\n" +
	"\x03two\x18\x02 \x01(\v2\x15.city.map.v2.RuinInfoR\x03two\x12+\n" +
	"\x05three\x18\x03 \x01(\v2\x15.city.map.v2.RuinInfoR\x05three\"\x12\n" +
	"\x10GetEventsRequest\"B\n" +
	"\x11GetEventsResponse\x12-\n" +
	"\x06events\x18\x01 \x01(\v2\x15.city.event.v1.EventsR\x06events\"\xf8\x03\n" +
	"\tRoadState\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12$\n" +
	"\x0ein_vehicle_cnt\x18\a \x01(\x05R\finVehicleCnt\x12&\n" +
	"\x0fout_vehicle_cnt\x18\b \x01(\x05R\routVehicleCnt\x12\x1f\n" +
	"\vvehicle_cnt\x18\t \x01(\x05R\n" +
	"vehicleCnt\x12+\n" +
	"\x12cum_in_vehicle_cnt\x18\n" +
	" \x01(\x05R\x0fcumInVehicleCnt\x12-\n" +
	"\x13cum_out_vehicle_cnt\x18\v \x01(\x05R\x10cumOutVehicleCnt\x12\x13\n" +
	"\x05avg_v\x18\x04 \x01(\x01R\x04avgV\x12&\n" +
	"\x0favg_travel_time\x18\f \x01(\x01R\ravgTravelTime\x12,\n" +
	"\x05level\x18\x02 \x01(\x0e2\x16.city.map.v2.RoadLevelR\x05level\x12)\n" +
	"\x10congestion_index\x18\r \x01(\x01R\x0fcongestionIndex\x127\n" +
	"\x06reason\x18\x03 \x01(\x0e2\x1f.city.map.v2.InterruptionReasonR\x06reason\x12,\n" +
	"\x05lanes\x18\x05 \x03(\v2\x16.city.map.v2.LaneStateR\x05lanes\x12\x13\n" +
	"\x05max_v\x18\x06 \x01(\x01R\x04maxV*\xc3\x01\n" +
	"\tRoadLevel\x12\x1a\n" +
	"\x16ROAD_LEVEL_UNSPECIFIED\x10\x00\x12\x14\n" +
	"\x10ROAD_LEVEL_CLEAR\x10\x01\x12\x19\n" +
	"\x15ROAD_LEVEL_LIGHT_LOAD\x10\x02\x12\x1a\n" +
	"\x16ROAD_LEVEL_MEDIUM_LOAD\x10\x03\x12\x19\n" +
	"\x15ROAD_LEVEL_HEAVY_LOAD\x10\x04\x12\x17\n" +
	"\x13ROAD_LEVEL_OVERLOAD\x10\x05\x12\x19\n" +
	"\x15ROAD_LEVEL_RESTRICTED\x10\x06*\x9e\x01\n" +
	"\x12InterruptionReason\x12#\n" +
	"\x1fINTERRUPTION_REASON_UNSPECIFIED\x10\x00\x12\x1e\n" +
	"\x1aINTERRUPTION_REASON_RUINED\x10\x01\x12\x1f\n" +
	"\x1bINTERRUPTION_REASON_CASCADE\x10\x02\x12\"\n" +
	"\x1eINTERRUPTION_REASON_CONGESTION\x10\x032\xe7\x02\n" +
	"\vRoadService\x12D\n" +
	"\aGetRoad\x12\x1b.city.map.v2.GetRoadRequest\x1a\x1c.city.map.v2.GetRoadResponse\x12t\n" +
	"\x17GetRoadGlobalStatistics\x12+.city.map.v2.GetRoadGlobalStatisticsRequest\x1a,.city.map.v2.GetRoadGlobalStatisticsResponse\x12P\n" +
	"\vGetRuinInfo\x12\x1f.city.map.v2.GetRuinInfoRequest\x1a .city.map.v2.GetRuinInfoResponse\x12J\n" +
	"\tGetEvents\x12\x1d.city.map.v2.GetEventsRequest\x1a\x1e.city.map.v2.GetEventsResponseB\xa4\x01\n" +
	"\x0fcom.city.map.v2B\x10RoadServiceProtoP\x01Z1git.fiblab.net/sim/protos/v2/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\vCity.Map.V2\xca\x02\vCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3"

var (
	file_city_map_v2_road_service_proto_rawDescOnce sync.Once
	file_city_map_v2_road_service_proto_rawDescData []byte
)

func file_city_map_v2_road_service_proto_rawDescGZIP() []byte {
	file_city_map_v2_road_service_proto_rawDescOnce.Do(func() {
		file_city_map_v2_road_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_map_v2_road_service_proto_rawDesc), len(file_city_map_v2_road_service_proto_rawDesc)))
	})
	return file_city_map_v2_road_service_proto_rawDescData
}

var file_city_map_v2_road_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_city_map_v2_road_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_city_map_v2_road_service_proto_goTypes = []any{
	(RoadLevel)(0),                          // 0: city.map.v2.RoadLevel
	(InterruptionReason)(0),                 // 1: city.map.v2.InterruptionReason
	(*GetRoadRequest)(nil),                  // 2: city.map.v2.GetRoadRequest
	(*GetRoadResponse)(nil),                 // 3: city.map.v2.GetRoadResponse
	(*GetRoadGlobalStatisticsRequest)(nil),  // 4: city.map.v2.GetRoadGlobalStatisticsRequest
	(*GetRoadGlobalStatisticsResponse)(nil), // 5: city.map.v2.GetRoadGlobalStatisticsResponse
	(*GetRuinInfoRequest)(nil),              // 6: city.map.v2.GetRuinInfoRequest
	(*RuinInfo)(nil),                        // 7: city.map.v2.RuinInfo
	(*GetRuinInfoResponse)(nil),             // 8: city.map.v2.GetRuinInfoResponse
	(*GetEventsRequest)(nil),                // 9: city.map.v2.GetEventsRequest
	(*GetEventsResponse)(nil),               // 10: city.map.v2.GetEventsResponse
	(*RoadState)(nil),                       // 11: city.map.v2.RoadState
	(*v1.Events)(nil),                       // 12: city.event.v1.Events
	(*LaneState)(nil),                       // 13: city.map.v2.LaneState
}
var file_city_map_v2_road_service_proto_depIdxs = []int32{
	11, // 0: city.map.v2.GetRoadResponse.states:type_name -> city.map.v2.RoadState
	7,  // 1: city.map.v2.GetRuinInfoResponse.one:type_name -> city.map.v2.RuinInfo
	7,  // 2: city.map.v2.GetRuinInfoResponse.two:type_name -> city.map.v2.RuinInfo
	7,  // 3: city.map.v2.GetRuinInfoResponse.three:type_name -> city.map.v2.RuinInfo
	12, // 4: city.map.v2.GetEventsResponse.events:type_name -> city.event.v1.Events
	0,  // 5: city.map.v2.RoadState.level:type_name -> city.map.v2.RoadLevel
	1,  // 6: city.map.v2.RoadState.reason:type_name -> city.map.v2.InterruptionReason
	13, // 7: city.map.v2.RoadState.lanes:type_name -> city.map.v2.LaneState
	2,  // 8: city.map.v2.RoadService.GetRoad:input_type -> city.map.v2.GetRoadRequest
	4,  // 9: city.map.v2.RoadService.GetRoadGlobalStatistics:input_type -> city.map.v2.GetRoadGlobalStatisticsRequest
	6,  // 10: city.map.v2.RoadService.GetRuinInfo:input_type -> city.map.v2.GetRuinInfoRequest
	9,  // 11: city.map.v2.RoadService.GetEvents:input_type -> city.map.v2.GetEventsRequest
	3,  // 12: city.map.v2.RoadService.GetRoad:output_type -> city.map.v2.GetRoadResponse
	5,  // 13: city.map.v2.RoadService.GetRoadGlobalStatistics:output_type -> city.map.v2.GetRoadGlobalStatisticsResponse
	8,  // 14: city.map.v2.RoadService.GetRuinInfo:output_type -> city.map.v2.GetRuinInfoResponse
	10, // 15: city.map.v2.RoadService.GetEvents:output_type -> city.map.v2.GetEventsResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_map_v2_road_service_proto_rawDesc), len(file_city_map_v2_road_service_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_map_v2_road_service_proto_goTypes,
		DependencyIndexes: file_city_map_v2_road_service_proto_depIdxs,
		EnumInfos:         file_city_map_v2_road_service_proto_enumTypes,
		MessageInfos:      file_city_map_v2_road_service_proto_msgTypes,
	}.Build()
	File_city_map_v2_road_service_proto = out.File
	file_city_map_v2_road_service_proto_goTypes = nil
	file_city_map_v2_road_service_proto_depIdxs = nil
}
