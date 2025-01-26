// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
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

var file_city_person_v1_vehicle_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x69, 0x74,
	0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x02, 0x4c,
	0x43, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x5f, 0x6c, 0x61, 0x6e, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x4c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x5f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x53, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x69,
	0x6f, 0x22, 0x7f, 0x0a, 0x0d, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x61, 0x63, 0x63, 0x12, 0x25, 0x0a, 0x0c, 0x6c, 0x63, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x63,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x6e, 0x67, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x61, 0x6e, 0x67, 0x6c,
	0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x63, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x22, 0xeb, 0x02, 0x0a, 0x0e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x02, 0x6c, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x43, 0x48, 0x00, 0x52, 0x02, 0x6c, 0x63, 0x88, 0x01, 0x01,
	0x12, 0x3a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x01, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x10,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x44,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x5f, 0x67,
	0x6f, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x73, 0x74, 0x72, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x41, 0x73, 0x74, 0x72, 0x61,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x74, 0x61, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x74,
	0x61, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x65, 0x74, 0x61, 0x46, 0x72, 0x65, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x6c, 0x63, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xc1, 0x01, 0x0a, 0x0f, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x44,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xaf, 0x01, 0x0a, 0x0c, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x4c, 0x61, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x12, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x9d, 0x02, 0x0a, 0x0a, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x45, 0x6e, 0x76, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x32, 0x0a, 0x07, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x52, 0x07, 0x6a, 0x6f, 0x75, 0x72,
	0x6e, 0x65, 0x79, 0x12, 0x4c, 0x0a, 0x11, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x10, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x12, 0x43, 0x0a, 0x0e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x6c, 0x61,
	0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x65, 0x52, 0x0d, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x4c, 0x61, 0x6e, 0x65, 0x73, 0x2a, 0xbb, 0x02, 0x0a, 0x0f, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x56, 0x45,
	0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x48, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x56, 0x45, 0x48, 0x49,
	0x43, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x45, 0x48,
	0x49, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45,
	0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x48, 0x41, 0x44, 0x4f, 0x57,
	0x5f, 0x41, 0x48, 0x45, 0x41, 0x44, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x56, 0x45, 0x48, 0x49,
	0x43, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x48, 0x41,
	0x44, 0x4f, 0x57, 0x5f, 0x42, 0x45, 0x48, 0x49, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b,
	0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4c, 0x45, 0x46, 0x54, 0x5f, 0x41, 0x48, 0x45, 0x41, 0x44, 0x10, 0x05, 0x12, 0x20, 0x0a,
	0x1c, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x41, 0x48, 0x45, 0x41, 0x44, 0x10, 0x06, 0x12,
	0x20, 0x0a, 0x1c, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x5f, 0x42, 0x45, 0x48, 0x49, 0x4e, 0x44, 0x10,
	0x07, 0x12, 0x21, 0x0a, 0x1d, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x4c,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x42, 0x45, 0x48, 0x49,
	0x4e, 0x44, 0x10, 0x08, 0x2a, 0x6d, 0x0a, 0x0a, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x4c,
	0x49, 0x47, 0x48, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x59, 0x45, 0x4c, 0x4c, 0x4f,
	0x57, 0x10, 0x03, 0x42, 0xb5, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x2e,
	0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x50, 0x58, 0xaa, 0x02, 0x0e, 0x43, 0x69, 0x74, 0x79,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x43, 0x69, 0x74,
	0x79, 0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x43, 0x69,
	0x74, 0x79, 0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x69, 0x74, 0x79, 0x3a,
	0x3a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

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
