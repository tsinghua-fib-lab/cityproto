// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/map/v2/traffic_light_service.proto

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

// 获取路口的红绿灯信息请求
// Reqeust for getting traffic light information
type GetTrafficLightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 信号等相关的接口精确到junction
	// The interfaces related to signals are precise to junction
	JunctionId int32 `protobuf:"varint,1,opt,name=junction_id,json=junctionId,proto3" json:"junction_id,omitempty" yaml:"junction_id" bson:"junction_id" db:"junction_id"`
}

func (x *GetTrafficLightRequest) Reset() {
	*x = GetTrafficLightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrafficLightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrafficLightRequest) ProtoMessage() {}

func (x *GetTrafficLightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrafficLightRequest.ProtoReflect.Descriptor instead.
func (*GetTrafficLightRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_traffic_light_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTrafficLightRequest) GetJunctionId() int32 {
	if x != nil {
		return x.JunctionId
	}
	return 0
}

// 获取路口的红绿灯信息响应
// Response of getting traffic light information
type GetTrafficLightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前路口处的红绿灯
	// The traffic light at the junction
	TrafficLight *TrafficLight `protobuf:"bytes,1,opt,name=traffic_light,json=trafficLight,proto3" json:"traffic_light,omitempty" yaml:"traffic_light" bson:"traffic_light" db:"traffic_light"`
	// 表示当前路口处的红绿灯处于哪一个相位
	// Which phase the traffic light is currently in
	PhaseIndex int32 `protobuf:"varint,2,opt,name=phase_index,json=phaseIndex,proto3" json:"phase_index,omitempty" yaml:"phase_index" bson:"phase_index" db:"phase_index"`
	// 当前相位的剩余时间
	// The remaining time of the current phase
	TimeRemaining float64 `protobuf:"fixed64,3,opt,name=time_remaining,json=timeRemaining,proto3" json:"time_remaining,omitempty" yaml:"time_remaining" bson:"time_remaining" db:"time_remaining"`
}

func (x *GetTrafficLightResponse) Reset() {
	*x = GetTrafficLightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrafficLightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrafficLightResponse) ProtoMessage() {}

func (x *GetTrafficLightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrafficLightResponse.ProtoReflect.Descriptor instead.
func (*GetTrafficLightResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_traffic_light_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTrafficLightResponse) GetTrafficLight() *TrafficLight {
	if x != nil {
		return x.TrafficLight
	}
	return nil
}

func (x *GetTrafficLightResponse) GetPhaseIndex() int32 {
	if x != nil {
		return x.PhaseIndex
	}
	return 0
}

func (x *GetTrafficLightResponse) GetTimeRemaining() float64 {
	if x != nil {
		return x.TimeRemaining
	}
	return 0
}

// 设置路口的红绿灯信息请求
// Request for setting traffic light information
type SetTrafficLightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要改变的红绿灯（含路口编号）
	// The target traffic light (including junction ID)
	TrafficLight *TrafficLight `protobuf:"bytes,1,opt,name=traffic_light,json=trafficLight,proto3" json:"traffic_light,omitempty" yaml:"traffic_light" bson:"traffic_light" db:"traffic_light"`
	// 指定当前路口处的红绿灯的相位
	// Specify the phase of the traffic light
	PhaseIndex int32 `protobuf:"varint,2,opt,name=phase_index,json=phaseIndex,proto3" json:"phase_index,omitempty" yaml:"phase_index" bson:"phase_index" db:"phase_index"`
	// 当前相位的剩余时间
	// The remaining time of the current phase
	TimeRemaining float64 `protobuf:"fixed64,3,opt,name=time_remaining,json=timeRemaining,proto3" json:"time_remaining,omitempty" bson:"time_remaining" db:"time_remaining" yaml:"time_remaining"`
}

func (x *SetTrafficLightRequest) Reset() {
	*x = SetTrafficLightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTrafficLightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTrafficLightRequest) ProtoMessage() {}

func (x *SetTrafficLightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTrafficLightRequest.ProtoReflect.Descriptor instead.
func (*SetTrafficLightRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_traffic_light_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetTrafficLightRequest) GetTrafficLight() *TrafficLight {
	if x != nil {
		return x.TrafficLight
	}
	return nil
}

func (x *SetTrafficLightRequest) GetPhaseIndex() int32 {
	if x != nil {
		return x.PhaseIndex
	}
	return 0
}

func (x *SetTrafficLightRequest) GetTimeRemaining() float64 {
	if x != nil {
		return x.TimeRemaining
	}
	return 0
}

// 设置路口的红绿灯信息响应
// Response of setting traffic light information
type SetTrafficLightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetTrafficLightResponse) Reset() {
	*x = SetTrafficLightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTrafficLightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTrafficLightResponse) ProtoMessage() {}

func (x *SetTrafficLightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTrafficLightResponse.ProtoReflect.Descriptor instead.
func (*SetTrafficLightResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_traffic_light_service_proto_rawDescGZIP(), []int{3}
}

// 设置路口的红绿灯相位请求
// Request for setting traffic light phase
type SetTrafficLightPhaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要改变相位的路口编号
	// The target junction ID
	JunctionId int32 `protobuf:"varint,1,opt,name=junction_id,json=junctionId,proto3" json:"junction_id,omitempty" yaml:"junction_id" bson:"junction_id" db:"junction_id"`
	// 指定当前路口红绿灯的相位
	// Specify the traffic light phase
	PhaseIndex int32 `protobuf:"varint,2,opt,name=phase_index,json=phaseIndex,proto3" json:"phase_index,omitempty" yaml:"phase_index" bson:"phase_index" db:"phase_index"`
	// 当前相位的剩余时间
	// The remaining time of the current phase
	TimeRemaining float64 `protobuf:"fixed64,3,opt,name=time_remaining,json=timeRemaining,proto3" json:"time_remaining,omitempty" yaml:"time_remaining" bson:"time_remaining" db:"time_remaining"`
}

func (x *SetTrafficLightPhaseRequest) Reset() {
	*x = SetTrafficLightPhaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTrafficLightPhaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTrafficLightPhaseRequest) ProtoMessage() {}

func (x *SetTrafficLightPhaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTrafficLightPhaseRequest.ProtoReflect.Descriptor instead.
func (*SetTrafficLightPhaseRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_traffic_light_service_proto_rawDescGZIP(), []int{4}
}

func (x *SetTrafficLightPhaseRequest) GetJunctionId() int32 {
	if x != nil {
		return x.JunctionId
	}
	return 0
}

func (x *SetTrafficLightPhaseRequest) GetPhaseIndex() int32 {
	if x != nil {
		return x.PhaseIndex
	}
	return 0
}

func (x *SetTrafficLightPhaseRequest) GetTimeRemaining() float64 {
	if x != nil {
		return x.TimeRemaining
	}
	return 0
}

// 设置路口的红绿灯相位响应
// Response of setting traffic light phase
type SetTrafficLightPhaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetTrafficLightPhaseResponse) Reset() {
	*x = SetTrafficLightPhaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTrafficLightPhaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTrafficLightPhaseResponse) ProtoMessage() {}

func (x *SetTrafficLightPhaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTrafficLightPhaseResponse.ProtoReflect.Descriptor instead.
func (*SetTrafficLightPhaseResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_traffic_light_service_proto_rawDescGZIP(), []int{5}
}

// 设置路口的红绿灯状态请求
// Request for setting traffic light status
type SetTrafficLightStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要改变状态的路口编号
	// The target junction ID
	JunctionId int32 `protobuf:"varint,1,opt,name=junction_id,json=junctionId,proto3" json:"junction_id,omitempty" yaml:"junction_id" bson:"junction_id" db:"junction_id"`
	// 当前路口红绿灯状态，true为通，false为断
	// The current traffic light status at the junction, true is on, false is off
	Ok bool `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty" yaml:"ok" bson:"ok" db:"ok"`
}

func (x *SetTrafficLightStatusRequest) Reset() {
	*x = SetTrafficLightStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTrafficLightStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTrafficLightStatusRequest) ProtoMessage() {}

func (x *SetTrafficLightStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTrafficLightStatusRequest.ProtoReflect.Descriptor instead.
func (*SetTrafficLightStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_traffic_light_service_proto_rawDescGZIP(), []int{6}
}

func (x *SetTrafficLightStatusRequest) GetJunctionId() int32 {
	if x != nil {
		return x.JunctionId
	}
	return 0
}

func (x *SetTrafficLightStatusRequest) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

// 设置路口的红绿灯状态响应
// Response of setting traffic light status
type SetTrafficLightStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetTrafficLightStatusResponse) Reset() {
	*x = SetTrafficLightStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTrafficLightStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTrafficLightStatusResponse) ProtoMessage() {}

func (x *SetTrafficLightStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_traffic_light_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTrafficLightStatusResponse.ProtoReflect.Descriptor instead.
func (*SetTrafficLightStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_traffic_light_service_proto_rawDescGZIP(), []int{7}
}

var File_city_map_v2_traffic_light_service_proto protoreflect.FileDescriptor

var file_city_map_v2_traffic_light_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70,
	0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6a, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xa1, 0x01, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x68, 0x61,
	0x73, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0xa0,
	0x01, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x68, 0x61,
	0x73, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x70, 0x68, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x86, 0x01, 0x0a,
	0x1b, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x68, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x1c, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x0a, 0x1c, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6a, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x1f, 0x0a, 0x1d, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xae, 0x03, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x23, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d,
	0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x4c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a,
	0x0f, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x23, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x14, 0x53,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x68,
	0x61, 0x73, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76,
	0x32, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68,
	0x74, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x68, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x29, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa9, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x42, 0x18, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69,
	0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f,
	0x76, 0x32, 0x3b, 0x6d, 0x61, 0x70, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02,
	0x0b, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0b, 0x43,
	0x69, 0x74, 0x79, 0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x17, 0x43, 0x69, 0x74,
	0x79, 0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x4d, 0x61, 0x70,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_map_v2_traffic_light_service_proto_rawDescOnce sync.Once
	file_city_map_v2_traffic_light_service_proto_rawDescData = file_city_map_v2_traffic_light_service_proto_rawDesc
)

func file_city_map_v2_traffic_light_service_proto_rawDescGZIP() []byte {
	file_city_map_v2_traffic_light_service_proto_rawDescOnce.Do(func() {
		file_city_map_v2_traffic_light_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_map_v2_traffic_light_service_proto_rawDescData)
	})
	return file_city_map_v2_traffic_light_service_proto_rawDescData
}

var file_city_map_v2_traffic_light_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_city_map_v2_traffic_light_service_proto_goTypes = []any{
	(*GetTrafficLightRequest)(nil),        // 0: city.map.v2.GetTrafficLightRequest
	(*GetTrafficLightResponse)(nil),       // 1: city.map.v2.GetTrafficLightResponse
	(*SetTrafficLightRequest)(nil),        // 2: city.map.v2.SetTrafficLightRequest
	(*SetTrafficLightResponse)(nil),       // 3: city.map.v2.SetTrafficLightResponse
	(*SetTrafficLightPhaseRequest)(nil),   // 4: city.map.v2.SetTrafficLightPhaseRequest
	(*SetTrafficLightPhaseResponse)(nil),  // 5: city.map.v2.SetTrafficLightPhaseResponse
	(*SetTrafficLightStatusRequest)(nil),  // 6: city.map.v2.SetTrafficLightStatusRequest
	(*SetTrafficLightStatusResponse)(nil), // 7: city.map.v2.SetTrafficLightStatusResponse
	(*TrafficLight)(nil),                  // 8: city.map.v2.TrafficLight
}
var file_city_map_v2_traffic_light_service_proto_depIdxs = []int32{
	8, // 0: city.map.v2.GetTrafficLightResponse.traffic_light:type_name -> city.map.v2.TrafficLight
	8, // 1: city.map.v2.SetTrafficLightRequest.traffic_light:type_name -> city.map.v2.TrafficLight
	0, // 2: city.map.v2.TrafficLightService.GetTrafficLight:input_type -> city.map.v2.GetTrafficLightRequest
	2, // 3: city.map.v2.TrafficLightService.SetTrafficLight:input_type -> city.map.v2.SetTrafficLightRequest
	4, // 4: city.map.v2.TrafficLightService.SetTrafficLightPhase:input_type -> city.map.v2.SetTrafficLightPhaseRequest
	6, // 5: city.map.v2.TrafficLightService.SetTrafficLightStatus:input_type -> city.map.v2.SetTrafficLightStatusRequest
	1, // 6: city.map.v2.TrafficLightService.GetTrafficLight:output_type -> city.map.v2.GetTrafficLightResponse
	3, // 7: city.map.v2.TrafficLightService.SetTrafficLight:output_type -> city.map.v2.SetTrafficLightResponse
	5, // 8: city.map.v2.TrafficLightService.SetTrafficLightPhase:output_type -> city.map.v2.SetTrafficLightPhaseResponse
	7, // 9: city.map.v2.TrafficLightService.SetTrafficLightStatus:output_type -> city.map.v2.SetTrafficLightStatusResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_city_map_v2_traffic_light_service_proto_init() }
func file_city_map_v2_traffic_light_service_proto_init() {
	if File_city_map_v2_traffic_light_service_proto != nil {
		return
	}
	file_city_map_v2_light_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_map_v2_traffic_light_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetTrafficLightRequest); i {
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
		file_city_map_v2_traffic_light_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetTrafficLightResponse); i {
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
		file_city_map_v2_traffic_light_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SetTrafficLightRequest); i {
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
		file_city_map_v2_traffic_light_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SetTrafficLightResponse); i {
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
		file_city_map_v2_traffic_light_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SetTrafficLightPhaseRequest); i {
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
		file_city_map_v2_traffic_light_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SetTrafficLightPhaseResponse); i {
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
		file_city_map_v2_traffic_light_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SetTrafficLightStatusRequest); i {
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
		file_city_map_v2_traffic_light_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SetTrafficLightStatusResponse); i {
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
			RawDescriptor: file_city_map_v2_traffic_light_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_map_v2_traffic_light_service_proto_goTypes,
		DependencyIndexes: file_city_map_v2_traffic_light_service_proto_depIdxs,
		MessageInfos:      file_city_map_v2_traffic_light_service_proto_msgTypes,
	}.Build()
	File_city_map_v2_traffic_light_service_proto = out.File
	file_city_map_v2_traffic_light_service_proto_rawDesc = nil
	file_city_map_v2_traffic_light_service_proto_goTypes = nil
	file_city_map_v2_traffic_light_service_proto_depIdxs = nil
}
