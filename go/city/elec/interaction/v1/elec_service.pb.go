// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: city/elec/interaction/v1/elec_service.proto

package interactionv1

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

type SetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 设施id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// True 表示恢复，False表示摧毁
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetStatusRequest) Reset() {
	*x = SetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusRequest) ProtoMessage() {}

func (x *SetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusRequest.ProtoReflect.Descriptor instead.
func (*SetStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetStatusRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetStatusRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type SetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetStatusResponse) Reset() {
	*x = SetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusResponse) ProtoMessage() {}

func (x *SetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusResponse.ProtoReflect.Descriptor instead.
func (*SetStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{1}
}

type GetPowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 设备id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPowerRequest) Reset() {
	*x = GetPowerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerRequest) ProtoMessage() {}

func (x *GetPowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerRequest.ProtoReflect.Descriptor instead.
func (*GetPowerRequest) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPowerRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 功率
	Power float64 `protobuf:"fixed64,1,opt,name=power,proto3" json:"power,omitempty"`
}

func (x *GetPowerResponse) Reset() {
	*x = GetPowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerResponse) ProtoMessage() {}

func (x *GetPowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerResponse.ProtoReflect.Descriptor instead.
func (*GetPowerResponse) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPowerResponse) GetPower() float64 {
	if x != nil {
		return x.Power
	}
	return 0
}

type GetPowerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag int32 `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *GetPowerStatusRequest) Reset() {
	*x = GetPowerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPowerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerStatusRequest) ProtoMessage() {}

func (x *GetPowerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerStatusRequest.ProtoReflect.Descriptor instead.
func (*GetPowerStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetPowerStatusRequest) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

type GetPowerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PowerStatus map[int32]float64 `protobuf:"bytes,1,rep,name=power_status,json=powerStatus,proto3" json:"power_status,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *GetPowerStatusResponse) Reset() {
	*x = GetPowerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPowerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPowerStatusResponse) ProtoMessage() {}

func (x *GetPowerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPowerStatusResponse.ProtoReflect.Descriptor instead.
func (*GetPowerStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetPowerStatusResponse) GetPowerStatus() map[int32]float64 {
	if x != nil {
		return x.PowerStatus
	}
	return nil
}

type GetNoPowerAOIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag int32 `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *GetNoPowerAOIRequest) Reset() {
	*x = GetNoPowerAOIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoPowerAOIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoPowerAOIRequest) ProtoMessage() {}

func (x *GetNoPowerAOIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoPowerAOIRequest.ProtoReflect.Descriptor instead.
func (*GetNoPowerAOIRequest) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetNoPowerAOIRequest) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

type GetNoPowerAOIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aoi []int32 `protobuf:"varint,1,rep,packed,name=aoi,proto3" json:"aoi,omitempty"`
}

func (x *GetNoPowerAOIResponse) Reset() {
	*x = GetNoPowerAOIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoPowerAOIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoPowerAOIResponse) ProtoMessage() {}

func (x *GetNoPowerAOIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoPowerAOIResponse.ProtoReflect.Descriptor instead.
func (*GetNoPowerAOIResponse) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetNoPowerAOIResponse) GetAoi() []int32 {
	if x != nil {
		return x.Aoi
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
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoRequest) ProtoMessage() {}

func (x *GetRuinInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{8}
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
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuinInfo) ProtoMessage() {}

func (x *RuinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{9}
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

	// 三级损伤信息
	One   *RuinInfo `protobuf:"bytes,1,opt,name=one,proto3" json:"one,omitempty"`
	Two   *RuinInfo `protobuf:"bytes,2,opt,name=two,proto3" json:"two,omitempty"`
	Three *RuinInfo `protobuf:"bytes,3,opt,name=three,proto3" json:"three,omitempty"`
}

func (x *GetRuinInfoResponse) Reset() {
	*x = GetRuinInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoResponse) ProtoMessage() {}

func (x *GetRuinInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{10}
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

type GetEdgeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEdgeStatusRequest) Reset() {
	*x = GetEdgeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEdgeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEdgeStatusRequest) ProtoMessage() {}

func (x *GetEdgeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEdgeStatusRequest.ProtoReflect.Descriptor instead.
func (*GetEdgeStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{11}
}

type GetEdgeStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason1 []string `protobuf:"bytes,1,rep,name=reason1,proto3" json:"reason1,omitempty"`
	Reason2 []string `protobuf:"bytes,2,rep,name=reason2,proto3" json:"reason2,omitempty"`
	Reason3 []string `protobuf:"bytes,3,rep,name=reason3,proto3" json:"reason3,omitempty"`
}

func (x *GetEdgeStatusResponse) Reset() {
	*x = GetEdgeStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEdgeStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEdgeStatusResponse) ProtoMessage() {}

func (x *GetEdgeStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_interaction_v1_elec_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEdgeStatusResponse.ProtoReflect.Descriptor instead.
func (*GetEdgeStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP(), []int{12}
}

func (x *GetEdgeStatusResponse) GetReason1() []string {
	if x != nil {
		return x.Reason1
	}
	return nil
}

func (x *GetEdgeStatusResponse) GetReason2() []string {
	if x != nil {
		return x.Reason2
	}
	return nil
}

func (x *GetEdgeStatusResponse) GetReason3() []string {
	if x != nil {
		return x.Reason3
	}
	return nil
}

var File_city_elec_interaction_v1_elec_service_proto protoreflect.FileDescriptor

var file_city_elec_interaction_v1_elec_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x3a, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x22, 0xbe, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a,
	0x0c, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x41, 0x4f, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x22,
	0x29, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x4f, 0x49,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6f, 0x69, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x61, 0x6f, 0x69, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x32, 0x0a, 0x08, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x75, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x03,
	0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x6f,
	0x6e, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x03, 0x74, 0x77, 0x6f, 0x12, 0x38, 0x0a, 0x05, 0x74, 0x68, 0x72, 0x65,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65,
	0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x68, 0x72,
	0x65, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x45, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x31, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x31, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x32, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x33, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x33, 0x32, 0x9b, 0x05, 0x0a, 0x0b, 0x45, 0x6c, 0x65, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x64, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x70, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x4f, 0x49,
	0x12, 0x2e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x4f, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x4f, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x75, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x69,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x45, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x64, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x64, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0xf8, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65,
	0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x42, 0x10, 0x45, 0x6c, 0x65, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62,
	0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x45, 0x49, 0xaa,
	0x02, 0x18, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x43, 0x69, 0x74,
	0x79, 0x5c, 0x45, 0x6c, 0x65, 0x63, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x6c, 0x65,
	0x63, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x43,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x6c, 0x65, 0x63, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_city_elec_interaction_v1_elec_service_proto_rawDescOnce sync.Once
	file_city_elec_interaction_v1_elec_service_proto_rawDescData = file_city_elec_interaction_v1_elec_service_proto_rawDesc
)

func file_city_elec_interaction_v1_elec_service_proto_rawDescGZIP() []byte {
	file_city_elec_interaction_v1_elec_service_proto_rawDescOnce.Do(func() {
		file_city_elec_interaction_v1_elec_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_elec_interaction_v1_elec_service_proto_rawDescData)
	})
	return file_city_elec_interaction_v1_elec_service_proto_rawDescData
}

var file_city_elec_interaction_v1_elec_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_city_elec_interaction_v1_elec_service_proto_goTypes = []interface{}{
	(*SetStatusRequest)(nil),       // 0: city.elec.interaction.v1.SetStatusRequest
	(*SetStatusResponse)(nil),      // 1: city.elec.interaction.v1.SetStatusResponse
	(*GetPowerRequest)(nil),        // 2: city.elec.interaction.v1.GetPowerRequest
	(*GetPowerResponse)(nil),       // 3: city.elec.interaction.v1.GetPowerResponse
	(*GetPowerStatusRequest)(nil),  // 4: city.elec.interaction.v1.GetPowerStatusRequest
	(*GetPowerStatusResponse)(nil), // 5: city.elec.interaction.v1.GetPowerStatusResponse
	(*GetNoPowerAOIRequest)(nil),   // 6: city.elec.interaction.v1.GetNoPowerAOIRequest
	(*GetNoPowerAOIResponse)(nil),  // 7: city.elec.interaction.v1.GetNoPowerAOIResponse
	(*GetRuinInfoRequest)(nil),     // 8: city.elec.interaction.v1.GetRuinInfoRequest
	(*RuinInfo)(nil),               // 9: city.elec.interaction.v1.RuinInfo
	(*GetRuinInfoResponse)(nil),    // 10: city.elec.interaction.v1.GetRuinInfoResponse
	(*GetEdgeStatusRequest)(nil),   // 11: city.elec.interaction.v1.GetEdgeStatusRequest
	(*GetEdgeStatusResponse)(nil),  // 12: city.elec.interaction.v1.GetEdgeStatusResponse
	nil,                            // 13: city.elec.interaction.v1.GetPowerStatusResponse.PowerStatusEntry
}
var file_city_elec_interaction_v1_elec_service_proto_depIdxs = []int32{
	13, // 0: city.elec.interaction.v1.GetPowerStatusResponse.power_status:type_name -> city.elec.interaction.v1.GetPowerStatusResponse.PowerStatusEntry
	9,  // 1: city.elec.interaction.v1.GetRuinInfoResponse.one:type_name -> city.elec.interaction.v1.RuinInfo
	9,  // 2: city.elec.interaction.v1.GetRuinInfoResponse.two:type_name -> city.elec.interaction.v1.RuinInfo
	9,  // 3: city.elec.interaction.v1.GetRuinInfoResponse.three:type_name -> city.elec.interaction.v1.RuinInfo
	0,  // 4: city.elec.interaction.v1.ElecService.SetStatus:input_type -> city.elec.interaction.v1.SetStatusRequest
	2,  // 5: city.elec.interaction.v1.ElecService.GetPower:input_type -> city.elec.interaction.v1.GetPowerRequest
	4,  // 6: city.elec.interaction.v1.ElecService.GetPowerStatus:input_type -> city.elec.interaction.v1.GetPowerStatusRequest
	6,  // 7: city.elec.interaction.v1.ElecService.GetNoPowerAOI:input_type -> city.elec.interaction.v1.GetNoPowerAOIRequest
	8,  // 8: city.elec.interaction.v1.ElecService.GetRuinInfo:input_type -> city.elec.interaction.v1.GetRuinInfoRequest
	11, // 9: city.elec.interaction.v1.ElecService.GetEdgeStatus:input_type -> city.elec.interaction.v1.GetEdgeStatusRequest
	1,  // 10: city.elec.interaction.v1.ElecService.SetStatus:output_type -> city.elec.interaction.v1.SetStatusResponse
	3,  // 11: city.elec.interaction.v1.ElecService.GetPower:output_type -> city.elec.interaction.v1.GetPowerResponse
	5,  // 12: city.elec.interaction.v1.ElecService.GetPowerStatus:output_type -> city.elec.interaction.v1.GetPowerStatusResponse
	7,  // 13: city.elec.interaction.v1.ElecService.GetNoPowerAOI:output_type -> city.elec.interaction.v1.GetNoPowerAOIResponse
	10, // 14: city.elec.interaction.v1.ElecService.GetRuinInfo:output_type -> city.elec.interaction.v1.GetRuinInfoResponse
	12, // 15: city.elec.interaction.v1.ElecService.GetEdgeStatus:output_type -> city.elec.interaction.v1.GetEdgeStatusResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_city_elec_interaction_v1_elec_service_proto_init() }
func file_city_elec_interaction_v1_elec_service_proto_init() {
	if File_city_elec_interaction_v1_elec_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStatusRequest); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStatusResponse); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPowerRequest); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPowerResponse); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPowerStatusRequest); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPowerStatusResponse); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoPowerAOIRequest); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoPowerAOIResponse); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuinInfoRequest); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuinInfo); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuinInfoResponse); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEdgeStatusRequest); i {
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
		file_city_elec_interaction_v1_elec_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEdgeStatusResponse); i {
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
			RawDescriptor: file_city_elec_interaction_v1_elec_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_elec_interaction_v1_elec_service_proto_goTypes,
		DependencyIndexes: file_city_elec_interaction_v1_elec_service_proto_depIdxs,
		MessageInfos:      file_city_elec_interaction_v1_elec_service_proto_msgTypes,
	}.Build()
	File_city_elec_interaction_v1_elec_service_proto = out.File
	file_city_elec_interaction_v1_elec_service_proto_rawDesc = nil
	file_city_elec_interaction_v1_elec_service_proto_goTypes = nil
	file_city_elec_interaction_v1_elec_service_proto_depIdxs = nil
}
