// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/person/v1/vehicle.proto

package personv1

import (
	v2 "git.fiblab.net/sim/protos/v2/go/city/routing/v2"
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

type VehicleRelation int32

const (
	// 未指定
	// unspecified
	VehicleRelation_VEHICLE_RELATION_UNSPECIFIED VehicleRelation = 0
	// 当前车道前车
	// vehicle ahead in the current lane
	VehicleRelation_VEHICLE_RELATION_AHEAD VehicleRelation = 1
	// 当前车道后车
	// vehicle behind in the current lane
	VehicleRelation_VEHICLE_RELATION_BEHIND VehicleRelation = 2
	// 影子车道前车
	// vehicle ahead in the shadow lane
	VehicleRelation_VEHICLE_RELATION_SHADOW_AHEAD VehicleRelation = 3
	// 影子车道后车
	// vehicle behind in the shadow lane
	VehicleRelation_VEHICLE_RELATION_SHADOW_BEHIND VehicleRelation = 4
	// 当前车道左侧车道前车
	// vehicle ahead in the left lane
	VehicleRelation_VEHICLE_RELATION_LEFT_AHEAD VehicleRelation = 5
	// 当前车道右侧车道前车
	// vehicle ahead in the right lane
	VehicleRelation_VEHICLE_RELATION_RIGHT_AHEAD VehicleRelation = 6
	// 当前车道左侧车道后车
	// vehicle behind in the left lane
	VehicleRelation_VEHICLE_RELATION_LEFT_BEHIND VehicleRelation = 7
	// 当前车道右侧车道后车
	// vehicle behind in the right lane
	VehicleRelation_VEHICLE_RELATION_RIGHT_BEHIND VehicleRelation = 8
)

// Enum value maps for VehicleRelation.
var (
	VehicleRelation_name = map[int32]string{
		0: "VEHICLE_RELATION_UNSPECIFIED",
		1: "VEHICLE_RELATION_AHEAD",
		2: "VEHICLE_RELATION_BEHIND",
		3: "VEHICLE_RELATION_SHADOW_AHEAD",
		4: "VEHICLE_RELATION_SHADOW_BEHIND",
		5: "VEHICLE_RELATION_LEFT_AHEAD",
		6: "VEHICLE_RELATION_RIGHT_AHEAD",
		7: "VEHICLE_RELATION_LEFT_BEHIND",
		8: "VEHICLE_RELATION_RIGHT_BEHIND",
	}
	VehicleRelation_value = map[string]int32{
		"VEHICLE_RELATION_UNSPECIFIED":   0,
		"VEHICLE_RELATION_AHEAD":         1,
		"VEHICLE_RELATION_BEHIND":        2,
		"VEHICLE_RELATION_SHADOW_AHEAD":  3,
		"VEHICLE_RELATION_SHADOW_BEHIND": 4,
		"VEHICLE_RELATION_LEFT_AHEAD":    5,
		"VEHICLE_RELATION_RIGHT_AHEAD":   6,
		"VEHICLE_RELATION_LEFT_BEHIND":   7,
		"VEHICLE_RELATION_RIGHT_BEHIND":  8,
	}
)

func (x VehicleRelation) Enum() *VehicleRelation {
	p := new(VehicleRelation)
	*p = x
	return p
}

func (x VehicleRelation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VehicleRelation) Descriptor() protoreflect.EnumDescriptor {
	return file_city_person_v1_vehicle_proto_enumTypes[0].Descriptor()
}

func (VehicleRelation) Type() protoreflect.EnumType {
	return &file_city_person_v1_vehicle_proto_enumTypes[0]
}

func (x VehicleRelation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VehicleRelation.Descriptor instead.
func (VehicleRelation) EnumDescriptor() ([]byte, []int) {
	return file_city_person_v1_vehicle_proto_rawDescGZIP(), []int{0}
}

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
	return file_city_person_v1_vehicle_proto_enumTypes[1].Descriptor()
}

func (LightState) Type() protoreflect.EnumType {
	return &file_city_person_v1_vehicle_proto_enumTypes[1]
}

func (x LightState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LightState.Descriptor instead.
func (LightState) EnumDescriptor() ([]byte, []int) {
	return file_city_person_v1_vehicle_proto_rawDescGZIP(), []int{1}
}

// 变道相关的信息
// lane change related information
type LC struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 影子车道ID（变道前的车道）
	// shadow lane id (lane before lane change)
	ShadowLaneId int32 `protobuf:"varint,1,opt,name=shadow_lane_id,json=shadowLaneId,proto3" json:"shadow_lane_id,omitempty" bson:"shadow_lane_id" db:"shadow_lane_id" yaml:"shadow_lane_id"`
	// 投影到影子车道的坐标
	// s coordinate projected to shadow lane
	ShadowS float64 `protobuf:"fixed64,2,opt,name=shadow_s,json=shadowS,proto3" json:"shadow_s,omitempty" bson:"shadow_s" db:"shadow_s" yaml:"shadow_s"`
	// 变道过程车头相对于前进方向的偏转角（弧度，总是为正，0代表不转向）
	// deviation angle of the vehicle head relative to the forward direction during lane change (radians, always positive, 0 means no steering)
	Angle float64 `protobuf:"fixed64,3,opt,name=angle,proto3" json:"angle,omitempty" bson:"angle" db:"angle" yaml:"angle"`
	// 已完成的变道比例
	// completed ratio of lane change
	CompletedRatio float64 `protobuf:"fixed64,4,opt,name=completed_ratio,json=completedRatio,proto3" json:"completed_ratio,omitempty" bson:"completed_ratio" db:"completed_ratio" yaml:"completed_ratio"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LC) Reset() {
	*x = LC{}
	mi := &file_city_person_v1_vehicle_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LC) ProtoMessage() {}

func (x *LC) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_vehicle_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LC.ProtoReflect.Descriptor instead.
func (*LC) Descriptor() ([]byte, []int) {
	return file_city_person_v1_vehicle_proto_rawDescGZIP(), []int{0}
}

func (x *LC) GetShadowLaneId() int32 {
	if x != nil {
		return x.ShadowLaneId
	}
	return 0
}

func (x *LC) GetShadowS() float64 {
	if x != nil {
		return x.ShadowS
	}
	return 0
}

func (x *LC) GetAngle() float64 {
	if x != nil {
		return x.Angle
	}
	return 0
}

func (x *LC) GetCompletedRatio() float64 {
	if x != nil {
		return x.CompletedRatio
	}
	return 0
}

// 车辆控制信息
// vehicle control information
type VehicleAction struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 车辆编号
	// vehicle id
	Id int32 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// 本轮更新中设定的加速度
	// acceleration set in this step
	Acc float64 `protobuf:"fixed64,1,opt,name=acc,proto3" json:"acc,omitempty" bson:"acc" db:"acc" yaml:"acc"`
	// 变道目标（可选，不设置代表不变道或保持变道状态）
	// lane change target (optional, not set means no lane change)
	LcTargetId *int32 `protobuf:"varint,2,opt,name=lc_target_id,json=lcTargetId,proto3,oneof" json:"lc_target_id,omitempty" bson:"lc_target_id" db:"lc_target_id" yaml:"lc_target_id"`
	// 变道过程的转向角度
	// steering angle during lane change
	Angle         float64 `protobuf:"fixed64,3,opt,name=angle,proto3" json:"angle,omitempty" bson:"angle" db:"angle" yaml:"angle"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VehicleAction) Reset() {
	*x = VehicleAction{}
	mi := &file_city_person_v1_vehicle_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VehicleAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleAction) ProtoMessage() {}

func (x *VehicleAction) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_vehicle_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleAction.ProtoReflect.Descriptor instead.
func (*VehicleAction) Descriptor() ([]byte, []int) {
	return file_city_person_v1_vehicle_proto_rawDescGZIP(), []int{1}
}

func (x *VehicleAction) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VehicleAction) GetAcc() float64 {
	if x != nil {
		return x.Acc
	}
	return 0
}

func (x *VehicleAction) GetLcTargetId() int32 {
	if x != nil && x.LcTargetId != nil {
		return *x.LcTargetId
	}
	return 0
}

func (x *VehicleAction) GetAngle() float64 {
	if x != nil {
		return x.Angle
	}
	return 0
}

type VehicleRuntime struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 基本运行时信息
	// basic runtime information
	Base *PersonMotion `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty" bson:"base" db:"base" yaml:"base"`
	// 变道信息
	// lane change information
	Lc *LC `protobuf:"bytes,4,opt,name=lc,proto3,oneof" json:"lc,omitempty" bson:"lc" db:"lc" yaml:"lc"`
	// 本轮车辆行为（获取车辆环境信息时不返回）
	// vehicle action in the step (not returned when getting vehicle environment information)
	Action *VehicleAction `protobuf:"bytes,5,opt,name=action,proto3,oneof" json:"action,omitempty" bson:"action" db:"action" yaml:"action"`
	// 走过的里程
	// running distance
	RunningDistance float64 `protobuf:"fixed64,6,opt,name=running_distance,json=runningDistance,proto3" json:"running_distance,omitempty" bson:"running_distance" db:"running_distance" yaml:"running_distance"`
	// 走错路次数
	// number of going astray
	NumGoingAstray int32 `protobuf:"varint,7,opt,name=num_going_astray,json=numGoingAstray,proto3" json:"num_going_astray,omitempty" bson:"num_going_astray" db:"num_going_astray" yaml:"num_going_astray"`
	// 出发时刻
	// departure time
	DepartureTime float64 `protobuf:"fixed64,8,opt,name=departure_time,json=departureTime,proto3" json:"departure_time,omitempty" bson:"departure_time" db:"departure_time" yaml:"departure_time"`
	// 预计到达时刻（导航返回的eta+出发时刻）
	// estimated arrival time (eta returned by routing + departure time)
	Eta float64 `protobuf:"fixed64,9,opt,name=eta,proto3" json:"eta,omitempty" bson:"eta" db:"eta" yaml:"eta"`
	// 自由流下的预计到达时刻
	// estimated arrival time under free flow
	EtaFreeFlow   float64 `protobuf:"fixed64,10,opt,name=eta_free_flow,json=etaFreeFlow,proto3" json:"eta_free_flow,omitempty" bson:"eta_free_flow" db:"eta_free_flow" yaml:"eta_free_flow"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VehicleRuntime) Reset() {
	*x = VehicleRuntime{}
	mi := &file_city_person_v1_vehicle_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VehicleRuntime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleRuntime) ProtoMessage() {}

func (x *VehicleRuntime) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_vehicle_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleRuntime.ProtoReflect.Descriptor instead.
func (*VehicleRuntime) Descriptor() ([]byte, []int) {
	return file_city_person_v1_vehicle_proto_rawDescGZIP(), []int{2}
}

func (x *VehicleRuntime) GetBase() *PersonMotion {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *VehicleRuntime) GetLc() *LC {
	if x != nil {
		return x.Lc
	}
	return nil
}

func (x *VehicleRuntime) GetAction() *VehicleAction {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *VehicleRuntime) GetRunningDistance() float64 {
	if x != nil {
		return x.RunningDistance
	}
	return 0
}

func (x *VehicleRuntime) GetNumGoingAstray() int32 {
	if x != nil {
		return x.NumGoingAstray
	}
	return 0
}

func (x *VehicleRuntime) GetDepartureTime() float64 {
	if x != nil {
		return x.DepartureTime
	}
	return 0
}

func (x *VehicleRuntime) GetEta() float64 {
	if x != nil {
		return x.Eta
	}
	return 0
}

func (x *VehicleRuntime) GetEtaFreeFlow() float64 {
	if x != nil {
		return x.EtaFreeFlow
	}
	return 0
}

// 观测到的车辆
// observed vehicles
type ObservedVehicle struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 车辆编号
	// vehicle id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// 当前的车辆运行时信息
	// current vehicle runtime information
	Motion *PersonMotion `protobuf:"bytes,2,opt,name=motion,proto3" json:"motion,omitempty" bson:"motion" db:"motion" yaml:"motion"`
	// 相对距离
	// relative distance
	RelativeDistance float64 `protobuf:"fixed64,3,opt,name=relative_distance,json=relativeDistance,proto3" json:"relative_distance,omitempty" bson:"relative_distance" db:"relative_distance" yaml:"relative_distance"`
	// 关系枚举
	// relation enumeration
	Relation      VehicleRelation `protobuf:"varint,4,opt,name=relation,proto3,enum=city.person.v1.VehicleRelation" json:"relation,omitempty" bson:"relation" db:"relation" yaml:"relation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObservedVehicle) Reset() {
	*x = ObservedVehicle{}
	mi := &file_city_person_v1_vehicle_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObservedVehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservedVehicle) ProtoMessage() {}

func (x *ObservedVehicle) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_vehicle_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservedVehicle.ProtoReflect.Descriptor instead.
func (*ObservedVehicle) Descriptor() ([]byte, []int) {
	return file_city_person_v1_vehicle_proto_rawDescGZIP(), []int{3}
}

func (x *ObservedVehicle) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObservedVehicle) GetMotion() *PersonMotion {
	if x != nil {
		return x.Motion
	}
	return nil
}

func (x *ObservedVehicle) GetRelativeDistance() float64 {
	if x != nil {
		return x.RelativeDistance
	}
	return 0
}

func (x *ObservedVehicle) GetRelation() VehicleRelation {
	if x != nil {
		return x.Relation
	}
	return VehicleRelation_VEHICLE_RELATION_UNSPECIFIED
}

// 观测到的车道
// observed lane
type ObservedLane struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Lane ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// 是否限行
	// whether restricted
	Restriction bool `protobuf:"varint,2,opt,name=restriction,proto3" json:"restriction,omitempty" bson:"restriction" db:"restriction" yaml:"restriction"`
	// 交通灯状态
	// traffic light state
	LightState LightState `protobuf:"varint,3,opt,name=light_state,json=lightState,proto3,enum=city.person.v1.LightState" json:"light_state,omitempty" bson:"light_state" db:"light_state" yaml:"light_state"`
	// 交通灯剩余时间
	// remaining time of traffic light
	LightRemainingTime float64 `protobuf:"fixed64,4,opt,name=light_remaining_time,json=lightRemainingTime,proto3" json:"light_remaining_time,omitempty" bson:"light_remaining_time" db:"light_remaining_time" yaml:"light_remaining_time"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ObservedLane) Reset() {
	*x = ObservedLane{}
	mi := &file_city_person_v1_vehicle_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObservedLane) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservedLane) ProtoMessage() {}

func (x *ObservedLane) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_vehicle_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservedLane.ProtoReflect.Descriptor instead.
func (*ObservedLane) Descriptor() ([]byte, []int) {
	return file_city_person_v1_vehicle_proto_rawDescGZIP(), []int{4}
}

func (x *ObservedLane) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObservedLane) GetRestriction() bool {
	if x != nil {
		return x.Restriction
	}
	return false
}

func (x *ObservedLane) GetLightState() LightState {
	if x != nil {
		return x.LightState
	}
	return LightState_LIGHT_STATE_UNSPECIFIED
}

func (x *ObservedLane) GetLightRemainingTime() float64 {
	if x != nil {
		return x.LightRemainingTime
	}
	return 0
}

type VehicleEnv struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 车辆编号
	// vehicle id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// 当前的车辆运行时信息
	// current vehicle runtime information
	Runtime *VehicleRuntime `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty" bson:"runtime" db:"runtime" yaml:"runtime"`
	// 当前的路径规划结果
	// journey: current routing result
	Journey *v2.Journey `protobuf:"bytes,3,opt,name=journey,proto3" json:"journey,omitempty" bson:"journey" db:"journey" yaml:"journey"`
	// 观测到的车辆
	// observed vehicles
	ObservedVehicles []*ObservedVehicle `protobuf:"bytes,4,rep,name=observed_vehicles,json=observedVehicles,proto3" json:"observed_vehicles,omitempty" bson:"observed_vehicles" db:"observed_vehicles" yaml:"observed_vehicles"`
	// 观测到的车道状态
	// observed lane states
	ObservedLanes []*ObservedLane `protobuf:"bytes,5,rep,name=observed_lanes,json=observedLanes,proto3" json:"observed_lanes,omitempty" bson:"observed_lanes" db:"observed_lanes" yaml:"observed_lanes"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VehicleEnv) Reset() {
	*x = VehicleEnv{}
	mi := &file_city_person_v1_vehicle_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VehicleEnv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleEnv) ProtoMessage() {}

func (x *VehicleEnv) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_vehicle_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleEnv.ProtoReflect.Descriptor instead.
func (*VehicleEnv) Descriptor() ([]byte, []int) {
	return file_city_person_v1_vehicle_proto_rawDescGZIP(), []int{5}
}

func (x *VehicleEnv) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VehicleEnv) GetRuntime() *VehicleRuntime {
	if x != nil {
		return x.Runtime
	}
	return nil
}

func (x *VehicleEnv) GetJourney() *v2.Journey {
	if x != nil {
		return x.Journey
	}
	return nil
}

func (x *VehicleEnv) GetObservedVehicles() []*ObservedVehicle {
	if x != nil {
		return x.ObservedVehicles
	}
	return nil
}

func (x *VehicleEnv) GetObservedLanes() []*ObservedLane {
	if x != nil {
		return x.ObservedLanes
	}
	return nil
}

var File_city_person_v1_vehicle_proto protoreflect.FileDescriptor

const file_city_person_v1_vehicle_proto_rawDesc = "" +
	"\n" +
	"\x1ccity/person/v1/vehicle.proto\x12\x0ecity.person.v1\x1a\x1bcity/person/v1/motion.proto\x1a\x1dcity/routing/v2/routing.proto\"\x84\x01\n" +
	"\x02LC\x12$\n" +
	"\x0eshadow_lane_id\x18\x01 \x01(\x05R\fshadowLaneId\x12\x19\n" +
	"\bshadow_s\x18\x02 \x01(\x01R\ashadowS\x12\x14\n" +
	"\x05angle\x18\x03 \x01(\x01R\x05angle\x12'\n" +
	"\x0fcompleted_ratio\x18\x04 \x01(\x01R\x0ecompletedRatio\"\x7f\n" +
	"\rVehicleAction\x12\x0e\n" +
	"\x02id\x18\x04 \x01(\x05R\x02id\x12\x10\n" +
	"\x03acc\x18\x01 \x01(\x01R\x03acc\x12%\n" +
	"\flc_target_id\x18\x02 \x01(\x05H\x00R\n" +
	"lcTargetId\x88\x01\x01\x12\x14\n" +
	"\x05angle\x18\x03 \x01(\x01R\x05angleB\x0f\n" +
	"\r_lc_target_id\"\xeb\x02\n" +
	"\x0eVehicleRuntime\x120\n" +
	"\x04base\x18\x01 \x01(\v2\x1c.city.person.v1.PersonMotionR\x04base\x12'\n" +
	"\x02lc\x18\x04 \x01(\v2\x12.city.person.v1.LCH\x00R\x02lc\x88\x01\x01\x12:\n" +
	"\x06action\x18\x05 \x01(\v2\x1d.city.person.v1.VehicleActionH\x01R\x06action\x88\x01\x01\x12)\n" +
	"\x10running_distance\x18\x06 \x01(\x01R\x0frunningDistance\x12(\n" +
	"\x10num_going_astray\x18\a \x01(\x05R\x0enumGoingAstray\x12%\n" +
	"\x0edeparture_time\x18\b \x01(\x01R\rdepartureTime\x12\x10\n" +
	"\x03eta\x18\t \x01(\x01R\x03eta\x12\"\n" +
	"\reta_free_flow\x18\n" +
	" \x01(\x01R\vetaFreeFlowB\x05\n" +
	"\x03_lcB\t\n" +
	"\a_action\"\xc1\x01\n" +
	"\x0fObservedVehicle\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x124\n" +
	"\x06motion\x18\x02 \x01(\v2\x1c.city.person.v1.PersonMotionR\x06motion\x12+\n" +
	"\x11relative_distance\x18\x03 \x01(\x01R\x10relativeDistance\x12;\n" +
	"\brelation\x18\x04 \x01(\x0e2\x1f.city.person.v1.VehicleRelationR\brelation\"\xaf\x01\n" +
	"\fObservedLane\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12 \n" +
	"\vrestriction\x18\x02 \x01(\bR\vrestriction\x12;\n" +
	"\vlight_state\x18\x03 \x01(\x0e2\x1a.city.person.v1.LightStateR\n" +
	"lightState\x120\n" +
	"\x14light_remaining_time\x18\x04 \x01(\x01R\x12lightRemainingTime\"\x9d\x02\n" +
	"\n" +
	"VehicleEnv\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x128\n" +
	"\aruntime\x18\x02 \x01(\v2\x1e.city.person.v1.VehicleRuntimeR\aruntime\x122\n" +
	"\ajourney\x18\x03 \x01(\v2\x18.city.routing.v2.JourneyR\ajourney\x12L\n" +
	"\x11observed_vehicles\x18\x04 \x03(\v2\x1f.city.person.v1.ObservedVehicleR\x10observedVehicles\x12C\n" +
	"\x0eobserved_lanes\x18\x05 \x03(\v2\x1c.city.person.v1.ObservedLaneR\robservedLanes*\xbb\x02\n" +
	"\x0fVehicleRelation\x12 \n" +
	"\x1cVEHICLE_RELATION_UNSPECIFIED\x10\x00\x12\x1a\n" +
	"\x16VEHICLE_RELATION_AHEAD\x10\x01\x12\x1b\n" +
	"\x17VEHICLE_RELATION_BEHIND\x10\x02\x12!\n" +
	"\x1dVEHICLE_RELATION_SHADOW_AHEAD\x10\x03\x12\"\n" +
	"\x1eVEHICLE_RELATION_SHADOW_BEHIND\x10\x04\x12\x1f\n" +
	"\x1bVEHICLE_RELATION_LEFT_AHEAD\x10\x05\x12 \n" +
	"\x1cVEHICLE_RELATION_RIGHT_AHEAD\x10\x06\x12 \n" +
	"\x1cVEHICLE_RELATION_LEFT_BEHIND\x10\a\x12!\n" +
	"\x1dVEHICLE_RELATION_RIGHT_BEHIND\x10\b*m\n" +
	"\n" +
	"LightState\x12\x1b\n" +
	"\x17LIGHT_STATE_UNSPECIFIED\x10\x00\x12\x13\n" +
	"\x0fLIGHT_STATE_RED\x10\x01\x12\x15\n" +
	"\x11LIGHT_STATE_GREEN\x10\x02\x12\x16\n" +
	"\x12LIGHT_STATE_YELLOW\x10\x03B\xb5\x01\n" +
	"\x12com.city.person.v1B\fVehicleProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1b\x06proto3"

var (
	file_city_person_v1_vehicle_proto_rawDescOnce sync.Once
	file_city_person_v1_vehicle_proto_rawDescData []byte
)

func file_city_person_v1_vehicle_proto_rawDescGZIP() []byte {
	file_city_person_v1_vehicle_proto_rawDescOnce.Do(func() {
		file_city_person_v1_vehicle_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_person_v1_vehicle_proto_rawDesc), len(file_city_person_v1_vehicle_proto_rawDesc)))
	})
	return file_city_person_v1_vehicle_proto_rawDescData
}

var file_city_person_v1_vehicle_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_city_person_v1_vehicle_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_city_person_v1_vehicle_proto_goTypes = []any{
	(VehicleRelation)(0),    // 0: city.person.v1.VehicleRelation
	(LightState)(0),         // 1: city.person.v1.LightState
	(*LC)(nil),              // 2: city.person.v1.LC
	(*VehicleAction)(nil),   // 3: city.person.v1.VehicleAction
	(*VehicleRuntime)(nil),  // 4: city.person.v1.VehicleRuntime
	(*ObservedVehicle)(nil), // 5: city.person.v1.ObservedVehicle
	(*ObservedLane)(nil),    // 6: city.person.v1.ObservedLane
	(*VehicleEnv)(nil),      // 7: city.person.v1.VehicleEnv
	(*PersonMotion)(nil),    // 8: city.person.v1.PersonMotion
	(*v2.Journey)(nil),      // 9: city.routing.v2.Journey
}
var file_city_person_v1_vehicle_proto_depIdxs = []int32{
	8,  // 0: city.person.v1.VehicleRuntime.base:type_name -> city.person.v1.PersonMotion
	2,  // 1: city.person.v1.VehicleRuntime.lc:type_name -> city.person.v1.LC
	3,  // 2: city.person.v1.VehicleRuntime.action:type_name -> city.person.v1.VehicleAction
	8,  // 3: city.person.v1.ObservedVehicle.motion:type_name -> city.person.v1.PersonMotion
	0,  // 4: city.person.v1.ObservedVehicle.relation:type_name -> city.person.v1.VehicleRelation
	1,  // 5: city.person.v1.ObservedLane.light_state:type_name -> city.person.v1.LightState
	4,  // 6: city.person.v1.VehicleEnv.runtime:type_name -> city.person.v1.VehicleRuntime
	9,  // 7: city.person.v1.VehicleEnv.journey:type_name -> city.routing.v2.Journey
	5,  // 8: city.person.v1.VehicleEnv.observed_vehicles:type_name -> city.person.v1.ObservedVehicle
	6,  // 9: city.person.v1.VehicleEnv.observed_lanes:type_name -> city.person.v1.ObservedLane
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_city_person_v1_vehicle_proto_init() }
func file_city_person_v1_vehicle_proto_init() {
	if File_city_person_v1_vehicle_proto != nil {
		return
	}
	file_city_person_v1_motion_proto_init()
	file_city_person_v1_vehicle_proto_msgTypes[1].OneofWrappers = []any{}
	file_city_person_v1_vehicle_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_person_v1_vehicle_proto_rawDesc), len(file_city_person_v1_vehicle_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_person_v1_vehicle_proto_goTypes,
		DependencyIndexes: file_city_person_v1_vehicle_proto_depIdxs,
		EnumInfos:         file_city_person_v1_vehicle_proto_enumTypes,
		MessageInfos:      file_city_person_v1_vehicle_proto_msgTypes,
	}.Build()
	File_city_person_v1_vehicle_proto = out.File
	file_city_person_v1_vehicle_proto_goTypes = nil
	file_city_person_v1_vehicle_proto_depIdxs = nil
}
