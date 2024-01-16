// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: city/agent/v2/agent_service.proto

package agentv2

import (
	v21 "git.fiblab.net/sim/protos/go/city/geo/v2"
	v2 "git.fiblab.net/sim/protos/go/city/trip/v2"
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

// 获取agent信息请求
type GetAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// agent id
	AgentId int32 `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty" bson:"agent_id" db:"agent_id" yaml:"agent_id"`
}

func (x *GetAgentRequest) Reset() {
	*x = GetAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentRequest) ProtoMessage() {}

func (x *GetAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentRequest.ProtoReflect.Descriptor instead.
func (*GetAgentRequest) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAgentRequest) GetAgentId() int32 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

// 获取agent信息响应
type GetAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// agent信息
	Base *Agent `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty" yaml:"base" bson:"base" db:"base"`
	// agent运动信息
	Motion *AgentMotion `protobuf:"bytes,2,opt,name=motion,proto3" json:"motion,omitempty" yaml:"motion" bson:"motion" db:"motion"`
}

func (x *GetAgentResponse) Reset() {
	*x = GetAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentResponse) ProtoMessage() {}

func (x *GetAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentResponse.ProtoReflect.Descriptor instead.
func (*GetAgentResponse) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAgentResponse) GetBase() *Agent {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *GetAgentResponse) GetMotion() *AgentMotion {
	if x != nil {
		return x.Motion
	}
	return nil
}

// 新增agent请求
type AddAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 约定：agent中不设置id
	Agent *Agent `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty" yaml:"agent" bson:"agent" db:"agent"`
}

func (x *AddAgentRequest) Reset() {
	*x = AddAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAgentRequest) ProtoMessage() {}

func (x *AddAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAgentRequest.ProtoReflect.Descriptor instead.
func (*AddAgentRequest) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddAgentRequest) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

// 新增agent响应
type AddAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 新增的agent分配得到的ID
	AgentId int32 `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty" yaml:"agent_id" bson:"agent_id" db:"agent_id"`
}

func (x *AddAgentResponse) Reset() {
	*x = AddAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAgentResponse) ProtoMessage() {}

func (x *AddAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAgentResponse.ProtoReflect.Descriptor instead.
func (*AddAgentResponse) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddAgentResponse) GetAgentId() int32 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

// 修改agent的schedule请求
type SetScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// agent id
	AgentId int32 `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty" yaml:"agent_id" bson:"agent_id" db:"agent_id"`
	// 新的schedule（覆盖原有的schedule）
	Schedules []*v2.Schedule `protobuf:"bytes,2,rep,name=schedules,proto3" json:"schedules,omitempty" yaml:"schedules" bson:"schedules" db:"schedules"`
}

func (x *SetScheduleRequest) Reset() {
	*x = SetScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScheduleRequest) ProtoMessage() {}

func (x *SetScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScheduleRequest.ProtoReflect.Descriptor instead.
func (*SetScheduleRequest) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_service_proto_rawDescGZIP(), []int{4}
}

func (x *SetScheduleRequest) GetAgentId() int32 {
	if x != nil {
		return x.AgentId
	}
	return 0
}

func (x *SetScheduleRequest) GetSchedules() []*v2.Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

// 修改agent的schedule响应
type SetScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetScheduleResponse) Reset() {
	*x = SetScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScheduleResponse) ProtoMessage() {}

func (x *SetScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScheduleResponse.ProtoReflect.Descriptor instead.
func (*SetScheduleResponse) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_service_proto_rawDescGZIP(), []int{5}
}

// 获取特定区域内的agent请求
type GetAgentsByLongLatAreaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 经纬度范围
	Area *v21.LongLatBBox `protobuf:"bytes,1,opt,name=area,proto3" json:"area,omitempty" yaml:"area" bson:"area" db:"area"`
}

func (x *GetAgentsByLongLatAreaRequest) Reset() {
	*x = GetAgentsByLongLatAreaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentsByLongLatAreaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentsByLongLatAreaRequest) ProtoMessage() {}

func (x *GetAgentsByLongLatAreaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentsByLongLatAreaRequest.ProtoReflect.Descriptor instead.
func (*GetAgentsByLongLatAreaRequest) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetAgentsByLongLatAreaRequest) GetArea() *v21.LongLatBBox {
	if x != nil {
		return x.Area
	}
	return nil
}

// 获取特定区域内的agent响应
type GetAgentsByLongLatAreaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前模拟步数
	Step int32 `protobuf:"varint,1,opt,name=step,proto3" json:"step,omitempty" yaml:"step" bson:"step" db:"step"`
	// 区域内的agent的运动信息
	Motions []*AgentMotion `protobuf:"bytes,2,rep,name=motions,proto3" json:"motions,omitempty" yaml:"motions" bson:"motions" db:"motions"`
}

func (x *GetAgentsByLongLatAreaResponse) Reset() {
	*x = GetAgentsByLongLatAreaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_agent_v2_agent_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentsByLongLatAreaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentsByLongLatAreaResponse) ProtoMessage() {}

func (x *GetAgentsByLongLatAreaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_agent_v2_agent_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentsByLongLatAreaResponse.ProtoReflect.Descriptor instead.
func (*GetAgentsByLongLatAreaResponse) Descriptor() ([]byte, []int) {
	return file_city_agent_v2_agent_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetAgentsByLongLatAreaResponse) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *GetAgentsByLongLatAreaResponse) GetMotions() []*AgentMotion {
	if x != nil {
		return x.Motions
	}
	return nil
}

var File_city_agent_v2_agent_service_proto protoreflect.FileDescriptor

var file_city_agent_v2_agent_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x32, 0x1a, 0x19, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x69, 0x74, 0x79, 0x2f,
	0x67, 0x65, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x74,
	0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x0f, 0x41, 0x64, 0x64,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x15,
	0x0a, 0x13, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x79, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x04,
	0x61, 0x72, 0x65, 0x61, 0x22, 0x6a, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x79, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0x89, 0x03, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03,
	0x88, 0x02, 0x01, 0x12, 0x50, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12,
	0x1e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x64, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x64, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x03, 0x88, 0x02, 0x01, 0x12, 0x59, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01,
	0x12, 0x7a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4c,
	0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x41, 0x72, 0x65, 0x61, 0x12, 0x2c, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x41, 0x72, 0x65,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x79, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x42, 0xb0, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x32, 0x42, 0x11, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62,
	0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x32, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x41,
	0x58, 0xaa, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x56,
	0x32, 0xca, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x5c, 0x56,
	0x32, 0xe2, 0x02, 0x19, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x5c, 0x56,
	0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f,
	0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_agent_v2_agent_service_proto_rawDescOnce sync.Once
	file_city_agent_v2_agent_service_proto_rawDescData = file_city_agent_v2_agent_service_proto_rawDesc
)

func file_city_agent_v2_agent_service_proto_rawDescGZIP() []byte {
	file_city_agent_v2_agent_service_proto_rawDescOnce.Do(func() {
		file_city_agent_v2_agent_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_agent_v2_agent_service_proto_rawDescData)
	})
	return file_city_agent_v2_agent_service_proto_rawDescData
}

var file_city_agent_v2_agent_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_city_agent_v2_agent_service_proto_goTypes = []interface{}{
	(*GetAgentRequest)(nil),                // 0: city.agent.v2.GetAgentRequest
	(*GetAgentResponse)(nil),               // 1: city.agent.v2.GetAgentResponse
	(*AddAgentRequest)(nil),                // 2: city.agent.v2.AddAgentRequest
	(*AddAgentResponse)(nil),               // 3: city.agent.v2.AddAgentResponse
	(*SetScheduleRequest)(nil),             // 4: city.agent.v2.SetScheduleRequest
	(*SetScheduleResponse)(nil),            // 5: city.agent.v2.SetScheduleResponse
	(*GetAgentsByLongLatAreaRequest)(nil),  // 6: city.agent.v2.GetAgentsByLongLatAreaRequest
	(*GetAgentsByLongLatAreaResponse)(nil), // 7: city.agent.v2.GetAgentsByLongLatAreaResponse
	(*Agent)(nil),                          // 8: city.agent.v2.Agent
	(*AgentMotion)(nil),                    // 9: city.agent.v2.AgentMotion
	(*v2.Schedule)(nil),                    // 10: city.trip.v2.Schedule
	(*v21.LongLatBBox)(nil),                // 11: city.geo.v2.LongLatBBox
}
var file_city_agent_v2_agent_service_proto_depIdxs = []int32{
	8,  // 0: city.agent.v2.GetAgentResponse.base:type_name -> city.agent.v2.Agent
	9,  // 1: city.agent.v2.GetAgentResponse.motion:type_name -> city.agent.v2.AgentMotion
	8,  // 2: city.agent.v2.AddAgentRequest.agent:type_name -> city.agent.v2.Agent
	10, // 3: city.agent.v2.SetScheduleRequest.schedules:type_name -> city.trip.v2.Schedule
	11, // 4: city.agent.v2.GetAgentsByLongLatAreaRequest.area:type_name -> city.geo.v2.LongLatBBox
	9,  // 5: city.agent.v2.GetAgentsByLongLatAreaResponse.motions:type_name -> city.agent.v2.AgentMotion
	0,  // 6: city.agent.v2.AgentService.GetAgent:input_type -> city.agent.v2.GetAgentRequest
	2,  // 7: city.agent.v2.AgentService.AddAgent:input_type -> city.agent.v2.AddAgentRequest
	4,  // 8: city.agent.v2.AgentService.SetSchedule:input_type -> city.agent.v2.SetScheduleRequest
	6,  // 9: city.agent.v2.AgentService.GetAgentsByLongLatArea:input_type -> city.agent.v2.GetAgentsByLongLatAreaRequest
	1,  // 10: city.agent.v2.AgentService.GetAgent:output_type -> city.agent.v2.GetAgentResponse
	3,  // 11: city.agent.v2.AgentService.AddAgent:output_type -> city.agent.v2.AddAgentResponse
	5,  // 12: city.agent.v2.AgentService.SetSchedule:output_type -> city.agent.v2.SetScheduleResponse
	7,  // 13: city.agent.v2.AgentService.GetAgentsByLongLatArea:output_type -> city.agent.v2.GetAgentsByLongLatAreaResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_city_agent_v2_agent_service_proto_init() }
func file_city_agent_v2_agent_service_proto_init() {
	if File_city_agent_v2_agent_service_proto != nil {
		return
	}
	file_city_agent_v2_agent_proto_init()
	file_city_agent_v2_motion_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_agent_v2_agent_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentRequest); i {
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
		file_city_agent_v2_agent_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentResponse); i {
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
		file_city_agent_v2_agent_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAgentRequest); i {
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
		file_city_agent_v2_agent_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAgentResponse); i {
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
		file_city_agent_v2_agent_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScheduleRequest); i {
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
		file_city_agent_v2_agent_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScheduleResponse); i {
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
		file_city_agent_v2_agent_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentsByLongLatAreaRequest); i {
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
		file_city_agent_v2_agent_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentsByLongLatAreaResponse); i {
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
			RawDescriptor: file_city_agent_v2_agent_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_agent_v2_agent_service_proto_goTypes,
		DependencyIndexes: file_city_agent_v2_agent_service_proto_depIdxs,
		MessageInfos:      file_city_agent_v2_agent_service_proto_msgTypes,
	}.Build()
	File_city_agent_v2_agent_service_proto = out.File
	file_city_agent_v2_agent_service_proto_rawDesc = nil
	file_city_agent_v2_agent_service_proto_goTypes = nil
	file_city_agent_v2_agent_service_proto_depIdxs = nil
}
