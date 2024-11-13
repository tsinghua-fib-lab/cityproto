// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/routing/v2/routing.proto

package routingv2

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

// 导航请求类型
// routing type
type RouteType int32

const (
	// 未指定
	RouteType_ROUTE_TYPE_UNSPECIFIED RouteType = 0
	// 驾车
	RouteType_ROUTE_TYPE_DRIVING RouteType = 1
	// 步行
	RouteType_ROUTE_TYPE_WALKING RouteType = 2
	// 公交
	RouteType_ROUTE_TYPE_BUS RouteType = 3
	// 地铁
	RouteType_ROUTE_TYPE_SUBWAY RouteType = 4
	// 公交+地铁
	RouteType_ROUTE_TYPE_BUS_SUBWAY RouteType = 5
)

// Enum value maps for RouteType.
var (
	RouteType_name = map[int32]string{
		0: "ROUTE_TYPE_UNSPECIFIED",
		1: "ROUTE_TYPE_DRIVING",
		2: "ROUTE_TYPE_WALKING",
		3: "ROUTE_TYPE_BUS",
		4: "ROUTE_TYPE_SUBWAY",
		5: "ROUTE_TYPE_BUS_SUBWAY",
	}
	RouteType_value = map[string]int32{
		"ROUTE_TYPE_UNSPECIFIED": 0,
		"ROUTE_TYPE_DRIVING":     1,
		"ROUTE_TYPE_WALKING":     2,
		"ROUTE_TYPE_BUS":         3,
		"ROUTE_TYPE_SUBWAY":      4,
		"ROUTE_TYPE_BUS_SUBWAY":  5,
	}
)

func (x RouteType) Enum() *RouteType {
	p := new(RouteType)
	*p = x
	return p
}

func (x RouteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteType) Descriptor() protoreflect.EnumDescriptor {
	return file_city_routing_v2_routing_proto_enumTypes[0].Descriptor()
}

func (RouteType) Type() protoreflect.EnumType {
	return &file_city_routing_v2_routing_proto_enumTypes[0]
}

func (x RouteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteType.Descriptor instead.
func (RouteType) EnumDescriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{0}
}

// 移动方式
// travelling mode
// Journey用以描述采用一种特定交通方式从一点出发到达另一点的路径。
// Journey is used to describe the path from one point to another using one specific travelling mode
// 一般来说，多个Journey是一个Trip的“实现”。
// Generally, multiple Journeys are used to "implement" a Trip
// 例如：Trip(从清华乘地铁到天安门):
// For example: Trip (taking the subway from Tsinghua to Tiananmen Square):
// Journey(步行到地铁站)->Journey(地铁)->Journey(步行到天安门)
// Journey (walking to subway station) -> Journey (subway) -> Journey (walking to Tiananmen Square)
type JourneyType int32

const (
	// 未指定
	// unspecified
	JourneyType_JOURNEY_TYPE_UNSPECIFIED JourneyType = 0
	// 驾车
	// driving
	JourneyType_JOURNEY_TYPE_DRIVING JourneyType = 1
	// 步行
	// walking
	JourneyType_JOURNEY_TYPE_WALKING JourneyType = 2
	// 公交
	// taking bus
	JourneyType_JOURNEY_TYPE_BY_BUS JourneyType = 3
)

// Enum value maps for JourneyType.
var (
	JourneyType_name = map[int32]string{
		0: "JOURNEY_TYPE_UNSPECIFIED",
		1: "JOURNEY_TYPE_DRIVING",
		2: "JOURNEY_TYPE_WALKING",
		3: "JOURNEY_TYPE_BY_BUS",
	}
	JourneyType_value = map[string]int32{
		"JOURNEY_TYPE_UNSPECIFIED": 0,
		"JOURNEY_TYPE_DRIVING":     1,
		"JOURNEY_TYPE_WALKING":     2,
		"JOURNEY_TYPE_BY_BUS":      3,
	}
)

func (x JourneyType) Enum() *JourneyType {
	p := new(JourneyType)
	*p = x
	return p
}

func (x JourneyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JourneyType) Descriptor() protoreflect.EnumDescriptor {
	return file_city_routing_v2_routing_proto_enumTypes[1].Descriptor()
}

func (JourneyType) Type() protoreflect.EnumType {
	return &file_city_routing_v2_routing_proto_enumTypes[1]
}

func (x JourneyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JourneyType.Descriptor instead.
func (JourneyType) EnumDescriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{1}
}

// 步行移动方向
// Walking direction
// 行人前进的方向与Lane的正方向（s增大的方向）的关系
// The relationship between the direction of pedestrian movement and the positive direction of Lane (the direction where s increases)
type MovingDirection int32

const (
	// 未指定
	// unspecified
	MovingDirection_MOVING_DIRECTION_UNSPECIFIED MovingDirection = 0
	// 与正方向同向
	// In the same direction as the positive lane direction
	MovingDirection_MOVING_DIRECTION_FORWARD MovingDirection = 1
	// 与正方向反向
	// In the opposite direction as the positive lane direction
	MovingDirection_MOVING_DIRECTION_BACKWARD MovingDirection = 2
)

// Enum value maps for MovingDirection.
var (
	MovingDirection_name = map[int32]string{
		0: "MOVING_DIRECTION_UNSPECIFIED",
		1: "MOVING_DIRECTION_FORWARD",
		2: "MOVING_DIRECTION_BACKWARD",
	}
	MovingDirection_value = map[string]int32{
		"MOVING_DIRECTION_UNSPECIFIED": 0,
		"MOVING_DIRECTION_FORWARD":     1,
		"MOVING_DIRECTION_BACKWARD":    2,
	}
)

func (x MovingDirection) Enum() *MovingDirection {
	p := new(MovingDirection)
	*p = x
	return p
}

func (x MovingDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MovingDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_city_routing_v2_routing_proto_enumTypes[2].Descriptor()
}

func (MovingDirection) Type() protoreflect.EnumType {
	return &file_city_routing_v2_routing_proto_enumTypes[2]
}

func (x MovingDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MovingDirection.Descriptor instead.
func (MovingDirection) EnumDescriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{2}
}

// 驾车出行方式的路径规划结果
// Routing results for driving journey
type DrivingJourneyBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 从起点到终点的道路序列
	// Road sequence from origin to destination
	// 采用道路序列的原因是主动变道对车道级的导航需要频繁修改
	// The reason for using road sequences is that active lane changing requires frequent modifications to lane level navigation
	// 优先使用road_ids，如果road_ids为空，则使用route（也可以直接忽略route）
	// Prioritize using road_ids. If road_ids is empty, use route (or simply ignore route)
	RoadIds []int32 `protobuf:"varint,2,rep,packed,name=road_ids,json=roadIds,proto3" json:"road_ids,omitempty" yaml:"road_ids" bson:"road_ids" db:"road_ids"`
	// 从起点到终点预计的时间(estimation time of arrival)
	// estimation time of arrival
	Eta float64 `protobuf:"fixed64,3,opt,name=eta,proto3" json:"eta,omitempty" db:"eta" yaml:"eta" bson:"eta"`
}

func (x *DrivingJourneyBody) Reset() {
	*x = DrivingJourneyBody{}
	mi := &file_city_routing_v2_routing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DrivingJourneyBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrivingJourneyBody) ProtoMessage() {}

func (x *DrivingJourneyBody) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrivingJourneyBody.ProtoReflect.Descriptor instead.
func (*DrivingJourneyBody) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{0}
}

func (x *DrivingJourneyBody) GetRoadIds() []int32 {
	if x != nil {
		return x.RoadIds
	}
	return nil
}

func (x *DrivingJourneyBody) GetEta() float64 {
	if x != nil {
		return x.Eta
	}
	return 0
}

// 步行出行方式的路径规划结果中的一段
// A segment in the routing results of walking journey
type WalkingRouteSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lane ID
	LaneId int32 `protobuf:"varint,1,opt,name=lane_id,json=laneId,proto3" json:"lane_id,omitempty" yaml:"lane_id" bson:"lane_id" db:"lane_id"`
	// 移动方向
	// moving direction
	MovingDirection MovingDirection `protobuf:"varint,2,opt,name=moving_direction,json=movingDirection,proto3,enum=city.routing.v2.MovingDirection" json:"moving_direction,omitempty" db:"moving_direction" yaml:"moving_direction" bson:"moving_direction"`
}

func (x *WalkingRouteSegment) Reset() {
	*x = WalkingRouteSegment{}
	mi := &file_city_routing_v2_routing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WalkingRouteSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalkingRouteSegment) ProtoMessage() {}

func (x *WalkingRouteSegment) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalkingRouteSegment.ProtoReflect.Descriptor instead.
func (*WalkingRouteSegment) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{1}
}

func (x *WalkingRouteSegment) GetLaneId() int32 {
	if x != nil {
		return x.LaneId
	}
	return 0
}

func (x *WalkingRouteSegment) GetMovingDirection() MovingDirection {
	if x != nil {
		return x.MovingDirection
	}
	return MovingDirection_MOVING_DIRECTION_UNSPECIFIED
}

// 步行出行方式的路径规划结果
// Routing results of walking journey
type WalkingJourneyBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 从起点到终点的（Lane+方向）序列
	// The (Lane+direction) sequence from the origin to destination
	Route []*WalkingRouteSegment `protobuf:"bytes,1,rep,name=route,proto3" json:"route,omitempty" yaml:"route" bson:"route" db:"route"`
	// 从起点到终点预计的时间(estimation time of arrival)
	// estimation time of arrival
	Eta float64 `protobuf:"fixed64,2,opt,name=eta,proto3" json:"eta,omitempty" yaml:"eta" bson:"eta" db:"eta"`
}

func (x *WalkingJourneyBody) Reset() {
	*x = WalkingJourneyBody{}
	mi := &file_city_routing_v2_routing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WalkingJourneyBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalkingJourneyBody) ProtoMessage() {}

func (x *WalkingJourneyBody) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalkingJourneyBody.ProtoReflect.Descriptor instead.
func (*WalkingJourneyBody) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{2}
}

func (x *WalkingJourneyBody) GetRoute() []*WalkingRouteSegment {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *WalkingJourneyBody) GetEta() float64 {
	if x != nil {
		return x.Eta
	}
	return 0
}

//	message BusJourneyBody {
//	  int32 line_id = 1;
//	  int32 start_station_id = 2;
//	  int32 end_station_id = 3;
//	}
type TransferSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SublineId      int32 `protobuf:"varint,1,opt,name=subline_id,json=sublineId,proto3" json:"subline_id,omitempty" yaml:"subline_id" bson:"subline_id" db:"subline_id"`
	StartStationId int32 `protobuf:"varint,2,opt,name=start_station_id,json=startStationId,proto3" json:"start_station_id,omitempty" yaml:"start_station_id" bson:"start_station_id" db:"start_station_id"`
	EndStationId   int32 `protobuf:"varint,3,opt,name=end_station_id,json=endStationId,proto3" json:"end_station_id,omitempty" yaml:"end_station_id" bson:"end_station_id" db:"end_station_id"`
}

func (x *TransferSegment) Reset() {
	*x = TransferSegment{}
	mi := &file_city_routing_v2_routing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferSegment) ProtoMessage() {}

func (x *TransferSegment) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferSegment.ProtoReflect.Descriptor instead.
func (*TransferSegment) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{3}
}

func (x *TransferSegment) GetSublineId() int32 {
	if x != nil {
		return x.SublineId
	}
	return 0
}

func (x *TransferSegment) GetStartStationId() int32 {
	if x != nil {
		return x.StartStationId
	}
	return 0
}

func (x *TransferSegment) GetEndStationId() int32 {
	if x != nil {
		return x.EndStationId
	}
	return 0
}

type BusJourneyBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transfers []*TransferSegment `protobuf:"bytes,1,rep,name=transfers,proto3" json:"transfers,omitempty" bson:"transfers" db:"transfers" yaml:"transfers"`
	// 从起点到终点预计的时间(estimation time of arrival)
	// estimation time of arrival
	Eta float64 `protobuf:"fixed64,2,opt,name=eta,proto3" json:"eta,omitempty" yaml:"eta" bson:"eta" db:"eta"`
}

func (x *BusJourneyBody) Reset() {
	*x = BusJourneyBody{}
	mi := &file_city_routing_v2_routing_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BusJourneyBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusJourneyBody) ProtoMessage() {}

func (x *BusJourneyBody) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusJourneyBody.ProtoReflect.Descriptor instead.
func (*BusJourneyBody) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{4}
}

func (x *BusJourneyBody) GetTransfers() []*TransferSegment {
	if x != nil {
		return x.Transfers
	}
	return nil
}

func (x *BusJourneyBody) GetEta() float64 {
	if x != nil {
		return x.Eta
	}
	return 0
}

// 路径规划结果的一部分，含且仅含采用一种交通出行方式的完整出行序列
// Part of the routing results, including a complete travel sequence using exactly one travelling mode
type Journey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 出行方式
	// journey travelling mode
	Type JourneyType `protobuf:"varint,1,opt,name=type,proto3,enum=city.routing.v2.JourneyType" json:"type,omitempty" bson:"type" db:"type" yaml:"type"`
	// 驾车
	// Routing results for driving journey
	Driving *DrivingJourneyBody `protobuf:"bytes,2,opt,name=driving,proto3,oneof" json:"driving,omitempty" yaml:"driving" bson:"driving" db:"driving"`
	// 步行
	// Routing results of walking journey
	Walking *WalkingJourneyBody `protobuf:"bytes,3,opt,name=walking,proto3,oneof" json:"walking,omitempty" yaml:"walking" bson:"walking" db:"walking"`
	// 公交
	// Routing results of bus journey
	ByBus *BusJourneyBody `protobuf:"bytes,4,opt,name=by_bus,json=byBus,proto3,oneof" json:"by_bus,omitempty" db:"by_bus" yaml:"by_bus" bson:"by_bus"`
}

func (x *Journey) Reset() {
	*x = Journey{}
	mi := &file_city_routing_v2_routing_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Journey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Journey) ProtoMessage() {}

func (x *Journey) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Journey.ProtoReflect.Descriptor instead.
func (*Journey) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{5}
}

func (x *Journey) GetType() JourneyType {
	if x != nil {
		return x.Type
	}
	return JourneyType_JOURNEY_TYPE_UNSPECIFIED
}

func (x *Journey) GetDriving() *DrivingJourneyBody {
	if x != nil {
		return x.Driving
	}
	return nil
}

func (x *Journey) GetWalking() *WalkingJourneyBody {
	if x != nil {
		return x.Walking
	}
	return nil
}

func (x *Journey) GetByBus() *BusJourneyBody {
	if x != nil {
		return x.ByBus
	}
	return nil
}

// 预计算路况信息
// Pre calculate road condition information
type RoadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 车道ID
	// Lane ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id" bson:"id" db:"id"`
	// 车道在各个时间片（每个5min）的速度
	// The speed of the lane at each time slot (5 minutes each)
	Speed []float64 `protobuf:"fixed64,2,rep,packed,name=speed,proto3" json:"speed,omitempty" yaml:"speed" bson:"speed" db:"speed"`
}

func (x *RoadStatus) Reset() {
	*x = RoadStatus{}
	mi := &file_city_routing_v2_routing_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoadStatus) ProtoMessage() {}

func (x *RoadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoadStatus.ProtoReflect.Descriptor instead.
func (*RoadStatus) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{6}
}

func (x *RoadStatus) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoadStatus) GetSpeed() []float64 {
	if x != nil {
		return x.Speed
	}
	return nil
}

// 预计算道路路况信息集合，对应一个预计算道路况信息pb文件或一个预计算路况信息mongodb collection
// Pre calculated road condition information set, corresponding to a pre calculated road condition information PB file or a pre calculated road condition information mongodb collection
type RoadStatuses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoadStatuses []*RoadStatus `protobuf:"bytes,1,rep,name=road_statuses,json=roadStatuses,proto3" json:"road_statuses,omitempty" yaml:"road_statuses" bson:"road_statuses" db:"road_statuses"`
}

func (x *RoadStatuses) Reset() {
	*x = RoadStatuses{}
	mi := &file_city_routing_v2_routing_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoadStatuses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoadStatuses) ProtoMessage() {}

func (x *RoadStatuses) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoadStatuses.ProtoReflect.Descriptor instead.
func (*RoadStatuses) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_proto_rawDescGZIP(), []int{7}
}

func (x *RoadStatuses) GetRoadStatuses() []*RoadStatus {
	if x != nil {
		return x.RoadStatuses
	}
	return nil
}

var File_city_routing_v2_routing_proto protoreflect.FileDescriptor

var file_city_routing_v2_routing_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32,
	0x22, 0x41, 0x0a, 0x12, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x65, 0x79, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x61, 0x64, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x61, 0x64, 0x49, 0x64,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x65, 0x74, 0x61, 0x22, 0x7b, 0x0a, 0x13, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x61, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x10, 0x6d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x6d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x62, 0x0a, 0x12, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x65, 0x79, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x65, 0x74, 0x61, 0x22, 0x80, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x75,
	0x62, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x0e, 0x42, 0x75, 0x73, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x65, 0x79, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x3e, 0x0a, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x65, 0x74, 0x61, 0x22, 0xa3, 0x02, 0x0a, 0x07,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x64, 0x72, 0x69,
	0x76, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x72, 0x69,
	0x76, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x42, 0x6f, 0x64, 0x79, 0x48,
	0x00, 0x52, 0x07, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a,
	0x07, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32,
	0x2e, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x42,
	0x6f, 0x64, 0x79, 0x48, 0x01, 0x52, 0x07, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x88, 0x01,
	0x01, 0x12, 0x3b, 0x0a, 0x06, 0x62, 0x79, 0x5f, 0x62, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x32, 0x2e, 0x42, 0x75, 0x73, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x42, 0x6f,
	0x64, 0x79, 0x48, 0x02, 0x52, 0x05, 0x62, 0x79, 0x42, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x77,
	0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x79, 0x5f, 0x62, 0x75,
	0x73, 0x22, 0x32, 0x0a, 0x0a, 0x52, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x05,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x22, 0x50, 0x0a, 0x0c, 0x52, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0d, 0x72, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x72, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x2a, 0x9d, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x52, 0x49, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f, 0x55,
	0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x41, 0x4c, 0x4b, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x42, 0x55, 0x53, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x55, 0x42, 0x57, 0x41, 0x59, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15,
	0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x53, 0x5f, 0x53,
	0x55, 0x42, 0x57, 0x41, 0x59, 0x10, 0x05, 0x2a, 0x78, 0x0a, 0x0b, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x4a, 0x4f, 0x55, 0x52, 0x4e, 0x45,
	0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f, 0x55, 0x52, 0x4e, 0x45, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x18,
	0x0a, 0x14, 0x4a, 0x4f, 0x55, 0x52, 0x4e, 0x45, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57,
	0x41, 0x4c, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4a, 0x4f, 0x55, 0x52,
	0x4e, 0x45, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x59, 0x5f, 0x42, 0x55, 0x53, 0x10,
	0x03, 0x2a, 0x70, 0x0a, 0x0f, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x4f, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x44,
	0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x4f, 0x56, 0x49, 0x4e, 0x47,
	0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41,
	0x52, 0x44, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x4f, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x44,
	0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x57, 0x41, 0x52,
	0x44, 0x10, 0x02, 0x42, 0xbc, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74,
	0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x52, 0x58, 0xaa, 0x02, 0x0f, 0x43,
	0x69, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca, 0x02,
	0x0f, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x32,
	0xe2, 0x02, 0x1b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c,
	0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a,
	0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_routing_v2_routing_proto_rawDescOnce sync.Once
	file_city_routing_v2_routing_proto_rawDescData = file_city_routing_v2_routing_proto_rawDesc
)

func file_city_routing_v2_routing_proto_rawDescGZIP() []byte {
	file_city_routing_v2_routing_proto_rawDescOnce.Do(func() {
		file_city_routing_v2_routing_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_routing_v2_routing_proto_rawDescData)
	})
	return file_city_routing_v2_routing_proto_rawDescData
}

var file_city_routing_v2_routing_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_city_routing_v2_routing_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_city_routing_v2_routing_proto_goTypes = []any{
	(RouteType)(0),              // 0: city.routing.v2.RouteType
	(JourneyType)(0),            // 1: city.routing.v2.JourneyType
	(MovingDirection)(0),        // 2: city.routing.v2.MovingDirection
	(*DrivingJourneyBody)(nil),  // 3: city.routing.v2.DrivingJourneyBody
	(*WalkingRouteSegment)(nil), // 4: city.routing.v2.WalkingRouteSegment
	(*WalkingJourneyBody)(nil),  // 5: city.routing.v2.WalkingJourneyBody
	(*TransferSegment)(nil),     // 6: city.routing.v2.TransferSegment
	(*BusJourneyBody)(nil),      // 7: city.routing.v2.BusJourneyBody
	(*Journey)(nil),             // 8: city.routing.v2.Journey
	(*RoadStatus)(nil),          // 9: city.routing.v2.RoadStatus
	(*RoadStatuses)(nil),        // 10: city.routing.v2.RoadStatuses
}
var file_city_routing_v2_routing_proto_depIdxs = []int32{
	2, // 0: city.routing.v2.WalkingRouteSegment.moving_direction:type_name -> city.routing.v2.MovingDirection
	4, // 1: city.routing.v2.WalkingJourneyBody.route:type_name -> city.routing.v2.WalkingRouteSegment
	6, // 2: city.routing.v2.BusJourneyBody.transfers:type_name -> city.routing.v2.TransferSegment
	1, // 3: city.routing.v2.Journey.type:type_name -> city.routing.v2.JourneyType
	3, // 4: city.routing.v2.Journey.driving:type_name -> city.routing.v2.DrivingJourneyBody
	5, // 5: city.routing.v2.Journey.walking:type_name -> city.routing.v2.WalkingJourneyBody
	7, // 6: city.routing.v2.Journey.by_bus:type_name -> city.routing.v2.BusJourneyBody
	9, // 7: city.routing.v2.RoadStatuses.road_statuses:type_name -> city.routing.v2.RoadStatus
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_city_routing_v2_routing_proto_init() }
func file_city_routing_v2_routing_proto_init() {
	if File_city_routing_v2_routing_proto != nil {
		return
	}
	file_city_routing_v2_routing_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_routing_v2_routing_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_routing_v2_routing_proto_goTypes,
		DependencyIndexes: file_city_routing_v2_routing_proto_depIdxs,
		EnumInfos:         file_city_routing_v2_routing_proto_enumTypes,
		MessageInfos:      file_city_routing_v2_routing_proto_msgTypes,
	}.Build()
	File_city_routing_v2_routing_proto = out.File
	file_city_routing_v2_routing_proto_rawDesc = nil
	file_city_routing_v2_routing_proto_goTypes = nil
	file_city_routing_v2_routing_proto_depIdxs = nil
}
