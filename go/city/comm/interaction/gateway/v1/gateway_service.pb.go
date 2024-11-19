// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: city/comm/interaction/gateway/v1/gateway_service.proto

package gatewayv1

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

// 断电或恢复状态，True为修复，False为断电
type SetGatewayPowerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status bool  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetGatewayPowerStatusRequest) Reset() {
	*x = SetGatewayPowerStatusRequest{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetGatewayPowerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGatewayPowerStatusRequest) ProtoMessage() {}

func (x *SetGatewayPowerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGatewayPowerStatusRequest.ProtoReflect.Descriptor instead.
func (*SetGatewayPowerStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetGatewayPowerStatusRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetGatewayPowerStatusRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type SetGatewayPowerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetGatewayPowerStatusResponse) Reset() {
	*x = SetGatewayPowerStatusResponse{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetGatewayPowerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGatewayPowerStatusResponse) ProtoMessage() {}

func (x *SetGatewayPowerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGatewayPowerStatusResponse.ProtoReflect.Descriptor instead.
func (*SetGatewayPowerStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{1}
}

// 摧毁或恢复状态，True为修复，False为摧毁
type SetGatewayRuinStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status bool  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetGatewayRuinStatusRequest) Reset() {
	*x = SetGatewayRuinStatusRequest{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetGatewayRuinStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGatewayRuinStatusRequest) ProtoMessage() {}

func (x *SetGatewayRuinStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGatewayRuinStatusRequest.ProtoReflect.Descriptor instead.
func (*SetGatewayRuinStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetGatewayRuinStatusRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetGatewayRuinStatusRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type SetGatewayRuinStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetGatewayRuinStatusResponse) Reset() {
	*x = SetGatewayRuinStatusResponse{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetGatewayRuinStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGatewayRuinStatusResponse) ProtoMessage() {}

func (x *SetGatewayRuinStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGatewayRuinStatusResponse.ProtoReflect.Descriptor instead.
func (*SetGatewayRuinStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{3}
}

type GetAllStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllStatusRequest) Reset() {
	*x = GetAllStatusRequest{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStatusRequest) ProtoMessage() {}

func (x *GetAllStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStatusRequest.ProtoReflect.Descriptor instead.
func (*GetAllStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{4}
}

type GetAllStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stations []*Station `protobuf:"bytes,1,rep,name=stations,proto3" json:"stations,omitempty"`
}

func (x *GetAllStatusResponse) Reset() {
	*x = GetAllStatusResponse{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStatusResponse) ProtoMessage() {}

func (x *GetAllStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStatusResponse.ProtoReflect.Descriptor instead.
func (*GetAllStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllStatusResponse) GetStations() []*Station {
	if x != nil {
		return x.Stations
	}
	return nil
}

type GetRuinInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRuinInfoRequest) Reset() {
	*x = GetRuinInfoRequest{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoRequest) ProtoMessage() {}

func (x *GetRuinInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[6]
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
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{6}
}

type RuinInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num   int32   `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`      // 损坏数量
	Ratio float64 `protobuf:"fixed64,2,opt,name=ratio,proto3" json:"ratio,omitempty"` // 损坏占比
}

func (x *RuinInfo) Reset() {
	*x = RuinInfo{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuinInfo) ProtoMessage() {}

func (x *RuinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[7]
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
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{7}
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 三级级损伤信息
	One   *RuinInfo `protobuf:"bytes,1,opt,name=one,proto3" json:"one,omitempty"`
	Two   *RuinInfo `protobuf:"bytes,2,opt,name=two,proto3" json:"two,omitempty"`
	Three *RuinInfo `protobuf:"bytes,3,opt,name=three,proto3" json:"three,omitempty"`
}

func (x *GetRuinInfoResponse) Reset() {
	*x = GetRuinInfoResponse{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoResponse) ProtoMessage() {}

func (x *GetRuinInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[8]
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
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{8}
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[9]
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
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{9}
}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events *v1.Events `protobuf:"bytes,1,opt,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes[10]
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
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP(), []int{10}
}

func (x *GetEventsResponse) GetEvents() *v1.Events {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_city_comm_interaction_gateway_v1_gateway_service_proto protoreflect.FileDescriptor

var file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x69, 0x74, 0x79,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x1c, 0x53, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1f, 0x0a,
	0x1d, 0x53, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45,
	0x0a, 0x1b, 0x53, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x75, 0x69, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x53, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x52, 0x75, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x32, 0x0a, 0x08, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x75, 0x69,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x03, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x74,
	0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x74, 0x77, 0x6f, 0x12, 0x40, 0x0a, 0x05, 0x74, 0x68, 0x72,
	0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x68, 0x72, 0x65, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0xb4, 0x05, 0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3f, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x95, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x52, 0x75, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x75, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x75, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x32, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xac, 0x02, 0x0a, 0x24, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x42, 0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x2e,
	0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x43, 0x43, 0x49, 0x47, 0xaa, 0x02, 0x20,
	0x43, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x20, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2c, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x24, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x3a,
	0x3a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescOnce sync.Once
	file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescData = file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDesc
)

func file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescGZIP() []byte {
	file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescOnce.Do(func() {
		file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescData)
	})
	return file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDescData
}

var file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_city_comm_interaction_gateway_v1_gateway_service_proto_goTypes = []any{
	(*SetGatewayPowerStatusRequest)(nil),  // 0: city.comm.interaction.gateway.v1.SetGatewayPowerStatusRequest
	(*SetGatewayPowerStatusResponse)(nil), // 1: city.comm.interaction.gateway.v1.SetGatewayPowerStatusResponse
	(*SetGatewayRuinStatusRequest)(nil),   // 2: city.comm.interaction.gateway.v1.SetGatewayRuinStatusRequest
	(*SetGatewayRuinStatusResponse)(nil),  // 3: city.comm.interaction.gateway.v1.SetGatewayRuinStatusResponse
	(*GetAllStatusRequest)(nil),           // 4: city.comm.interaction.gateway.v1.GetAllStatusRequest
	(*GetAllStatusResponse)(nil),          // 5: city.comm.interaction.gateway.v1.GetAllStatusResponse
	(*GetRuinInfoRequest)(nil),            // 6: city.comm.interaction.gateway.v1.GetRuinInfoRequest
	(*RuinInfo)(nil),                      // 7: city.comm.interaction.gateway.v1.RuinInfo
	(*GetRuinInfoResponse)(nil),           // 8: city.comm.interaction.gateway.v1.GetRuinInfoResponse
	(*GetEventsRequest)(nil),              // 9: city.comm.interaction.gateway.v1.GetEventsRequest
	(*GetEventsResponse)(nil),             // 10: city.comm.interaction.gateway.v1.GetEventsResponse
	(*Station)(nil),                       // 11: city.comm.interaction.gateway.v1.Station
	(*v1.Events)(nil),                     // 12: city.event.v1.Events
}
var file_city_comm_interaction_gateway_v1_gateway_service_proto_depIdxs = []int32{
	11, // 0: city.comm.interaction.gateway.v1.GetAllStatusResponse.stations:type_name -> city.comm.interaction.gateway.v1.Station
	7,  // 1: city.comm.interaction.gateway.v1.GetRuinInfoResponse.one:type_name -> city.comm.interaction.gateway.v1.RuinInfo
	7,  // 2: city.comm.interaction.gateway.v1.GetRuinInfoResponse.two:type_name -> city.comm.interaction.gateway.v1.RuinInfo
	7,  // 3: city.comm.interaction.gateway.v1.GetRuinInfoResponse.three:type_name -> city.comm.interaction.gateway.v1.RuinInfo
	12, // 4: city.comm.interaction.gateway.v1.GetEventsResponse.events:type_name -> city.event.v1.Events
	0,  // 5: city.comm.interaction.gateway.v1.GatewayService.SetGatewayPowerStatus:input_type -> city.comm.interaction.gateway.v1.SetGatewayPowerStatusRequest
	2,  // 6: city.comm.interaction.gateway.v1.GatewayService.SetGatewayRuinStatus:input_type -> city.comm.interaction.gateway.v1.SetGatewayRuinStatusRequest
	4,  // 7: city.comm.interaction.gateway.v1.GatewayService.GetAllStatus:input_type -> city.comm.interaction.gateway.v1.GetAllStatusRequest
	6,  // 8: city.comm.interaction.gateway.v1.GatewayService.GetRuinInfo:input_type -> city.comm.interaction.gateway.v1.GetRuinInfoRequest
	9,  // 9: city.comm.interaction.gateway.v1.GatewayService.GetEvents:input_type -> city.comm.interaction.gateway.v1.GetEventsRequest
	1,  // 10: city.comm.interaction.gateway.v1.GatewayService.SetGatewayPowerStatus:output_type -> city.comm.interaction.gateway.v1.SetGatewayPowerStatusResponse
	3,  // 11: city.comm.interaction.gateway.v1.GatewayService.SetGatewayRuinStatus:output_type -> city.comm.interaction.gateway.v1.SetGatewayRuinStatusResponse
	5,  // 12: city.comm.interaction.gateway.v1.GatewayService.GetAllStatus:output_type -> city.comm.interaction.gateway.v1.GetAllStatusResponse
	8,  // 13: city.comm.interaction.gateway.v1.GatewayService.GetRuinInfo:output_type -> city.comm.interaction.gateway.v1.GetRuinInfoResponse
	10, // 14: city.comm.interaction.gateway.v1.GatewayService.GetEvents:output_type -> city.comm.interaction.gateway.v1.GetEventsResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_city_comm_interaction_gateway_v1_gateway_service_proto_init() }
func file_city_comm_interaction_gateway_v1_gateway_service_proto_init() {
	if File_city_comm_interaction_gateway_v1_gateway_service_proto != nil {
		return
	}
	file_city_comm_interaction_gateway_v1_gateway_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_comm_interaction_gateway_v1_gateway_service_proto_goTypes,
		DependencyIndexes: file_city_comm_interaction_gateway_v1_gateway_service_proto_depIdxs,
		MessageInfos:      file_city_comm_interaction_gateway_v1_gateway_service_proto_msgTypes,
	}.Build()
	File_city_comm_interaction_gateway_v1_gateway_service_proto = out.File
	file_city_comm_interaction_gateway_v1_gateway_service_proto_rawDesc = nil
	file_city_comm_interaction_gateway_v1_gateway_service_proto_goTypes = nil
	file_city_comm_interaction_gateway_v1_gateway_service_proto_depIdxs = nil
}
