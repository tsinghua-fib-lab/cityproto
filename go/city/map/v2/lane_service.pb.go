// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: city/map/v2/lane_service.proto

package mapv2

import (
	v2 "git.fiblab.net/sim/protos/go/city/geo/v2"
	v1 "git.fiblab.net/sim/protos/go/city/person/v1"
	v21 "git.fiblab.net/sim/protos/go/city/traffic_light/v2"
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

// 设置Lane的最大速度（限速）请求
type SetLaneMaxVRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lane id
	LaneId int32 `protobuf:"varint,1,opt,name=lane_id,json=laneId,proto3" json:"lane_id,omitempty" db:"lane_id" yaml:"lane_id" bson:"lane_id"`
	// 最大速度（限速），单位：m/s
	MaxV float64 `protobuf:"fixed64,2,opt,name=max_v,json=maxV,proto3" json:"max_v,omitempty" yaml:"max_v" bson:"max_v" db:"max_v"`
}

func (x *SetLaneMaxVRequest) Reset() {
	*x = SetLaneMaxVRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_lane_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLaneMaxVRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLaneMaxVRequest) ProtoMessage() {}

func (x *SetLaneMaxVRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_lane_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLaneMaxVRequest.ProtoReflect.Descriptor instead.
func (*SetLaneMaxVRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_lane_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetLaneMaxVRequest) GetLaneId() int32 {
	if x != nil {
		return x.LaneId
	}
	return 0
}

func (x *SetLaneMaxVRequest) GetMaxV() float64 {
	if x != nil {
		return x.MaxV
	}
	return 0
}

// 设置Lane的最大速度（限速）响应
type SetLaneMaxVResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetLaneMaxVResponse) Reset() {
	*x = SetLaneMaxVResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_lane_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLaneMaxVResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLaneMaxVResponse) ProtoMessage() {}

func (x *SetLaneMaxVResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_lane_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLaneMaxVResponse.ProtoReflect.Descriptor instead.
func (*SetLaneMaxVResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_lane_service_proto_rawDescGZIP(), []int{1}
}

// 获取Lane的信息请求
type GetLaneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 指定的Lane id列表，如果为空，则返回所有Lane的信息
	LaneIds []int32 `protobuf:"varint,1,rep,packed,name=lane_ids,json=laneIds,proto3" json:"lane_ids,omitempty" yaml:"lane_ids" bson:"lane_ids" db:"lane_ids"`
	// 是否要排除车道上的人的信息
	ExcludePerson bool `protobuf:"varint,2,opt,name=exclude_person,json=excludePerson,proto3" json:"exclude_person,omitempty" bson:"exclude_person" db:"exclude_person" yaml:"exclude_person"`
}

func (x *GetLaneRequest) Reset() {
	*x = GetLaneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_lane_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLaneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLaneRequest) ProtoMessage() {}

func (x *GetLaneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_lane_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLaneRequest.ProtoReflect.Descriptor instead.
func (*GetLaneRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_lane_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLaneRequest) GetLaneIds() []int32 {
	if x != nil {
		return x.LaneIds
	}
	return nil
}

func (x *GetLaneRequest) GetExcludePerson() bool {
	if x != nil {
		return x.ExcludePerson
	}
	return false
}

// 获取Lane的信息响应
type GetLaneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lane的信息
	States []*LaneState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty" yaml:"states" bson:"states" db:"states"`
}

func (x *GetLaneResponse) Reset() {
	*x = GetLaneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_lane_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLaneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLaneResponse) ProtoMessage() {}

func (x *GetLaneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_lane_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLaneResponse.ProtoReflect.Descriptor instead.
func (*GetLaneResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_lane_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetLaneResponse) GetStates() []*LaneState {
	if x != nil {
		return x.States
	}
	return nil
}

// 获取特定区域内的Lane的信息请求
type GetLaneByLongLatBBoxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 经纬度范围
	Bound *v2.LongLatBBox `protobuf:"bytes,1,opt,name=bound,proto3" json:"bound,omitempty" db:"bound" yaml:"bound" bson:"bound"`
	// 是否要排除车道上的人的信息
	ExcludePerson bool `protobuf:"varint,2,opt,name=exclude_person,json=excludePerson,proto3" json:"exclude_person,omitempty" yaml:"exclude_person" bson:"exclude_person" db:"exclude_person"`
}

func (x *GetLaneByLongLatBBoxRequest) Reset() {
	*x = GetLaneByLongLatBBoxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_lane_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLaneByLongLatBBoxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLaneByLongLatBBoxRequest) ProtoMessage() {}

func (x *GetLaneByLongLatBBoxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_lane_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLaneByLongLatBBoxRequest.ProtoReflect.Descriptor instead.
func (*GetLaneByLongLatBBoxRequest) Descriptor() ([]byte, []int) {
	return file_city_map_v2_lane_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetLaneByLongLatBBoxRequest) GetBound() *v2.LongLatBBox {
	if x != nil {
		return x.Bound
	}
	return nil
}

func (x *GetLaneByLongLatBBoxRequest) GetExcludePerson() bool {
	if x != nil {
		return x.ExcludePerson
	}
	return false
}

// 获取特定区域内的Lane的信息响应
type GetLaneByLongLatBBoxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lane的信息
	States []*LaneState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty" yaml:"states" bson:"states" db:"states"`
}

func (x *GetLaneByLongLatBBoxResponse) Reset() {
	*x = GetLaneByLongLatBBoxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_lane_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLaneByLongLatBBoxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLaneByLongLatBBoxResponse) ProtoMessage() {}

func (x *GetLaneByLongLatBBoxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_lane_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLaneByLongLatBBoxResponse.ProtoReflect.Descriptor instead.
func (*GetLaneByLongLatBBoxResponse) Descriptor() ([]byte, []int) {
	return file_city_map_v2_lane_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetLaneByLongLatBBoxResponse) GetStates() []*LaneState {
	if x != nil {
		return x.States
	}
	return nil
}

// Lane状态
type LaneState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lane ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id" bson:"id" db:"id"`
	// Lane上的人/车
	Persons []*v1.PersonMotion `protobuf:"bytes,2,rep,name=persons,proto3" json:"persons,omitempty" yaml:"persons" bson:"persons" db:"persons"`
	// 平均速度（m/s）
	AvgV float64 `protobuf:"fixed64,3,opt,name=avg_v,json=avgV,proto3" json:"avg_v,omitempty" db:"avg_v" yaml:"avg_v" bson:"avg_v"`
	// 是否限行
	Restriction bool `protobuf:"varint,4,opt,name=restriction,proto3" json:"restriction,omitempty" bson:"restriction" db:"restriction" yaml:"restriction"`
	// 交通灯状态
	LightState v21.LightState `protobuf:"varint,5,opt,name=light_state,json=lightState,proto3,enum=city.traffic_light.v2.LightState" json:"light_state,omitempty" yaml:"light_state" bson:"light_state" db:"light_state"`
}

func (x *LaneState) Reset() {
	*x = LaneState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_map_v2_lane_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaneState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneState) ProtoMessage() {}

func (x *LaneState) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_lane_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneState.ProtoReflect.Descriptor instead.
func (*LaneState) Descriptor() ([]byte, []int) {
	return file_city_map_v2_lane_service_proto_rawDescGZIP(), []int{6}
}

func (x *LaneState) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LaneState) GetPersons() []*v1.PersonMotion {
	if x != nil {
		return x.Persons
	}
	return nil
}

func (x *LaneState) GetAvgV() float64 {
	if x != nil {
		return x.AvgV
	}
	return 0
}

func (x *LaneState) GetRestriction() bool {
	if x != nil {
		return x.Restriction
	}
	return false
}

func (x *LaneState) GetLightState() v21.LightState {
	if x != nil {
		return x.LightState
	}
	return v21.LightState(0)
}

var File_city_map_v2_lane_service_proto protoreflect.FileDescriptor

var file_city_map_v2_lane_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x61,
	0x6e, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x1a, 0x15, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x12,
	0x53, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x4d, 0x61, 0x78, 0x56, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x6d,
	0x61, 0x78, 0x5f, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x61, 0x78, 0x56,
	0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x4d, 0x61, 0x78, 0x56, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x6e,
	0x65, 0x49, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x61, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x74,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x42, 0x79, 0x4c, 0x6f, 0x6e, 0x67, 0x4c,
	0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x05, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4c,
	0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x05, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x42,
	0x79, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x22, 0xce, 0x01, 0x0a, 0x09, 0x4c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x76,
	0x67, 0x5f, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x61, 0x76, 0x67, 0x56, 0x12,
	0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x42, 0x0a, 0x0b, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x32, 0x92, 0x02, 0x0a, 0x0b, 0x4c, 0x61, 0x6e, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65,
	0x4d, 0x61, 0x78, 0x56, 0x12, 0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x4d, 0x61, 0x78, 0x56, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x4d, 0x61, 0x78, 0x56, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x6e, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x42, 0x79, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61,
	0x74, 0x42, 0x42, 0x6f, 0x78, 0x12, 0x28, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70,
	0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x42, 0x79, 0x4c, 0x6f, 0x6e,
	0x67, 0x4c, 0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x6e, 0x65, 0x42, 0x79, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x42, 0x42,
	0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa1, 0x01, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x42, 0x10,
	0x4c, 0x61, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e,
	0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x3b, 0x6d, 0x61, 0x70,
	0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x2e,
	0x4d, 0x61, 0x70, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x4d, 0x61,
	0x70, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x17, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x4d, 0x61, 0x70, 0x5c,
	0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0d, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_map_v2_lane_service_proto_rawDescOnce sync.Once
	file_city_map_v2_lane_service_proto_rawDescData = file_city_map_v2_lane_service_proto_rawDesc
)

func file_city_map_v2_lane_service_proto_rawDescGZIP() []byte {
	file_city_map_v2_lane_service_proto_rawDescOnce.Do(func() {
		file_city_map_v2_lane_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_map_v2_lane_service_proto_rawDescData)
	})
	return file_city_map_v2_lane_service_proto_rawDescData
}

var file_city_map_v2_lane_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_city_map_v2_lane_service_proto_goTypes = []interface{}{
	(*SetLaneMaxVRequest)(nil),           // 0: city.map.v2.SetLaneMaxVRequest
	(*SetLaneMaxVResponse)(nil),          // 1: city.map.v2.SetLaneMaxVResponse
	(*GetLaneRequest)(nil),               // 2: city.map.v2.GetLaneRequest
	(*GetLaneResponse)(nil),              // 3: city.map.v2.GetLaneResponse
	(*GetLaneByLongLatBBoxRequest)(nil),  // 4: city.map.v2.GetLaneByLongLatBBoxRequest
	(*GetLaneByLongLatBBoxResponse)(nil), // 5: city.map.v2.GetLaneByLongLatBBoxResponse
	(*LaneState)(nil),                    // 6: city.map.v2.LaneState
	(*v2.LongLatBBox)(nil),               // 7: city.geo.v2.LongLatBBox
	(*v1.PersonMotion)(nil),              // 8: city.person.v1.PersonMotion
	(v21.LightState)(0),                  // 9: city.traffic_light.v2.LightState
}
var file_city_map_v2_lane_service_proto_depIdxs = []int32{
	6, // 0: city.map.v2.GetLaneResponse.states:type_name -> city.map.v2.LaneState
	7, // 1: city.map.v2.GetLaneByLongLatBBoxRequest.bound:type_name -> city.geo.v2.LongLatBBox
	6, // 2: city.map.v2.GetLaneByLongLatBBoxResponse.states:type_name -> city.map.v2.LaneState
	8, // 3: city.map.v2.LaneState.persons:type_name -> city.person.v1.PersonMotion
	9, // 4: city.map.v2.LaneState.light_state:type_name -> city.traffic_light.v2.LightState
	0, // 5: city.map.v2.LaneService.SetLaneMaxV:input_type -> city.map.v2.SetLaneMaxVRequest
	2, // 6: city.map.v2.LaneService.GetLane:input_type -> city.map.v2.GetLaneRequest
	4, // 7: city.map.v2.LaneService.GetLaneByLongLatBBox:input_type -> city.map.v2.GetLaneByLongLatBBoxRequest
	1, // 8: city.map.v2.LaneService.SetLaneMaxV:output_type -> city.map.v2.SetLaneMaxVResponse
	3, // 9: city.map.v2.LaneService.GetLane:output_type -> city.map.v2.GetLaneResponse
	5, // 10: city.map.v2.LaneService.GetLaneByLongLatBBox:output_type -> city.map.v2.GetLaneByLongLatBBoxResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_city_map_v2_lane_service_proto_init() }
func file_city_map_v2_lane_service_proto_init() {
	if File_city_map_v2_lane_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_map_v2_lane_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLaneMaxVRequest); i {
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
		file_city_map_v2_lane_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLaneMaxVResponse); i {
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
		file_city_map_v2_lane_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLaneRequest); i {
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
		file_city_map_v2_lane_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLaneResponse); i {
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
		file_city_map_v2_lane_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLaneByLongLatBBoxRequest); i {
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
		file_city_map_v2_lane_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLaneByLongLatBBoxResponse); i {
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
		file_city_map_v2_lane_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaneState); i {
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
			RawDescriptor: file_city_map_v2_lane_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_map_v2_lane_service_proto_goTypes,
		DependencyIndexes: file_city_map_v2_lane_service_proto_depIdxs,
		MessageInfos:      file_city_map_v2_lane_service_proto_msgTypes,
	}.Build()
	File_city_map_v2_lane_service_proto = out.File
	file_city_map_v2_lane_service_proto_rawDesc = nil
	file_city_map_v2_lane_service_proto_goTypes = nil
	file_city_map_v2_lane_service_proto_depIdxs = nil
}
