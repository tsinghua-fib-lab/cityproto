// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/water/interaction/v1/water_service.proto

package interactionv1

import (
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

type WaterFacilityType int32

const (
	WaterFacilityType_WATER_FACILITY_TYPE_UNSPECIFIED WaterFacilityType = 0
	// 供水设施
	WaterFacilityType_WATER_FACILITY_TYPE_SUPPLY WaterFacilityType = 1
	// 排水设施
	WaterFacilityType_WATER_FACILITY_TYPE_DRAINAGE WaterFacilityType = 2
)

// Enum value maps for WaterFacilityType.
var (
	WaterFacilityType_name = map[int32]string{
		0: "WATER_FACILITY_TYPE_UNSPECIFIED",
		1: "WATER_FACILITY_TYPE_SUPPLY",
		2: "WATER_FACILITY_TYPE_DRAINAGE",
	}
	WaterFacilityType_value = map[string]int32{
		"WATER_FACILITY_TYPE_UNSPECIFIED": 0,
		"WATER_FACILITY_TYPE_SUPPLY":      1,
		"WATER_FACILITY_TYPE_DRAINAGE":    2,
	}
)

func (x WaterFacilityType) Enum() *WaterFacilityType {
	p := new(WaterFacilityType)
	*p = x
	return p
}

func (x WaterFacilityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WaterFacilityType) Descriptor() protoreflect.EnumDescriptor {
	return file_city_water_interaction_v1_water_service_proto_enumTypes[0].Descriptor()
}

func (WaterFacilityType) Type() protoreflect.EnumType {
	return &file_city_water_interaction_v1_water_service_proto_enumTypes[0]
}

func (x WaterFacilityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WaterFacilityType.Descriptor instead.
func (WaterFacilityType) EnumDescriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{0}
}

type SetPumpPowerStatusRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 水泵id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// True表示恢复，False表示摧毁
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// 供水水泵还是排水水泵
	Type          WaterFacilityType `protobuf:"varint,3,opt,name=type,proto3,enum=city.water.interaction.v1.WaterFacilityType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPumpPowerStatusRequest) Reset() {
	*x = SetPumpPowerStatusRequest{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPumpPowerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPumpPowerStatusRequest) ProtoMessage() {}

func (x *SetPumpPowerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPumpPowerStatusRequest.ProtoReflect.Descriptor instead.
func (*SetPumpPowerStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetPumpPowerStatusRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetPumpPowerStatusRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SetPumpPowerStatusRequest) GetType() WaterFacilityType {
	if x != nil {
		return x.Type
	}
	return WaterFacilityType_WATER_FACILITY_TYPE_UNSPECIFIED
}

type SetPumpPowerStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPumpPowerStatusResponse) Reset() {
	*x = SetPumpPowerStatusResponse{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPumpPowerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPumpPowerStatusResponse) ProtoMessage() {}

func (x *SetPumpPowerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPumpPowerStatusResponse.ProtoReflect.Descriptor instead.
func (*SetPumpPowerStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{1}
}

type SetPumpNetworkStatusRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 水泵id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// True表示恢复，False表示摧毁
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// 供水水泵还是排水水泵
	Type          WaterFacilityType `protobuf:"varint,3,opt,name=type,proto3,enum=city.water.interaction.v1.WaterFacilityType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPumpNetworkStatusRequest) Reset() {
	*x = SetPumpNetworkStatusRequest{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPumpNetworkStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPumpNetworkStatusRequest) ProtoMessage() {}

func (x *SetPumpNetworkStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPumpNetworkStatusRequest.ProtoReflect.Descriptor instead.
func (*SetPumpNetworkStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetPumpNetworkStatusRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetPumpNetworkStatusRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SetPumpNetworkStatusRequest) GetType() WaterFacilityType {
	if x != nil {
		return x.Type
	}
	return WaterFacilityType_WATER_FACILITY_TYPE_UNSPECIFIED
}

type SetPumpNetworkStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPumpNetworkStatusResponse) Reset() {
	*x = SetPumpNetworkStatusResponse{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPumpNetworkStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPumpNetworkStatusResponse) ProtoMessage() {}

func (x *SetPumpNetworkStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPumpNetworkStatusResponse.ProtoReflect.Descriptor instead.
func (*SetPumpNetworkStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{3}
}

type SetPumpStatusRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 水泵id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// True表示恢复，False表示摧毁
	Status bool `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// 供水水泵还是排水水泵
	Type          WaterFacilityType `protobuf:"varint,3,opt,name=type,proto3,enum=city.water.interaction.v1.WaterFacilityType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPumpStatusRequest) Reset() {
	*x = SetPumpStatusRequest{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPumpStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPumpStatusRequest) ProtoMessage() {}

func (x *SetPumpStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPumpStatusRequest.ProtoReflect.Descriptor instead.
func (*SetPumpStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{4}
}

func (x *SetPumpStatusRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetPumpStatusRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SetPumpStatusRequest) GetType() WaterFacilityType {
	if x != nil {
		return x.Type
	}
	return WaterFacilityType_WATER_FACILITY_TYPE_UNSPECIFIED
}

type SetPumpStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetPumpStatusResponse) Reset() {
	*x = SetPumpStatusResponse{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetPumpStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPumpStatusResponse) ProtoMessage() {}

func (x *SetPumpStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPumpStatusResponse.ProtoReflect.Descriptor instead.
func (*SetPumpStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{5}
}

type GetPumpStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Flag          int32                  `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPumpStatusRequest) Reset() {
	*x = GetPumpStatusRequest{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPumpStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPumpStatusRequest) ProtoMessage() {}

func (x *GetPumpStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPumpStatusRequest.ProtoReflect.Descriptor instead.
func (*GetPumpStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetPumpStatusRequest) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

type GetPumpStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PumpStatus    map[int32]int32        `protobuf:"bytes,1,rep,name=pump_status,json=pumpStatus,proto3" json:"pump_status,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPumpStatusResponse) Reset() {
	*x = GetPumpStatusResponse{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPumpStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPumpStatusResponse) ProtoMessage() {}

func (x *GetPumpStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPumpStatusResponse.ProtoReflect.Descriptor instead.
func (*GetPumpStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetPumpStatusResponse) GetPumpStatus() map[int32]int32 {
	if x != nil {
		return x.PumpStatus
	}
	return nil
}

type GetNoWaterAOIRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Flag          int32                  `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNoWaterAOIRequest) Reset() {
	*x = GetNoWaterAOIRequest{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNoWaterAOIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoWaterAOIRequest) ProtoMessage() {}

func (x *GetNoWaterAOIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoWaterAOIRequest.ProtoReflect.Descriptor instead.
func (*GetNoWaterAOIRequest) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetNoWaterAOIRequest) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

type GetNoWaterAOIResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Aoi           []int32                `protobuf:"varint,1,rep,packed,name=aoi,proto3" json:"aoi,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNoWaterAOIResponse) Reset() {
	*x = GetNoWaterAOIResponse{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNoWaterAOIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoWaterAOIResponse) ProtoMessage() {}

func (x *GetNoWaterAOIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoWaterAOIResponse.ProtoReflect.Descriptor instead.
func (*GetNoWaterAOIResponse) Descriptor() ([]byte, []int) {
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetNoWaterAOIResponse) GetAoi() []int32 {
	if x != nil {
		return x.Aoi
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
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoRequest) ProtoMessage() {}

func (x *GetRuinInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[10]
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
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{10}
}

type RuinInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Num           int32                  `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`      // 损坏数量
	Ratio         float64                `protobuf:"fixed64,2,opt,name=ratio,proto3" json:"ratio,omitempty"` // 损坏占比
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuinInfo) Reset() {
	*x = RuinInfo{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuinInfo) ProtoMessage() {}

func (x *RuinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[11]
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
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{11}
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
	One           *RuinInfo `protobuf:"bytes,1,opt,name=one,proto3" json:"one,omitempty"`
	Two           *RuinInfo `protobuf:"bytes,2,opt,name=two,proto3" json:"two,omitempty"`
	Three         *RuinInfo `protobuf:"bytes,3,opt,name=three,proto3" json:"three,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRuinInfoResponse) Reset() {
	*x = GetRuinInfoResponse{}
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuinInfoResponse) ProtoMessage() {}

func (x *GetRuinInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_interaction_v1_water_service_proto_msgTypes[12]
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
	return file_city_water_interaction_v1_water_service_proto_rawDescGZIP(), []int{12}
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

var File_city_water_interaction_v1_water_service_proto protoreflect.FileDescriptor

const file_city_water_interaction_v1_water_service_proto_rawDesc = "" +
	"\n" +
	"-city/water/interaction/v1/water_service.proto\x12\x19city.water.interaction.v1\"\x85\x01\n" +
	"\x19SetPumpPowerStatusRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\bR\x06status\x12@\n" +
	"\x04type\x18\x03 \x01(\x0e2,.city.water.interaction.v1.WaterFacilityTypeR\x04type\"\x1c\n" +
	"\x1aSetPumpPowerStatusResponse\"\x87\x01\n" +
	"\x1bSetPumpNetworkStatusRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\bR\x06status\x12@\n" +
	"\x04type\x18\x03 \x01(\x0e2,.city.water.interaction.v1.WaterFacilityTypeR\x04type\"\x1e\n" +
	"\x1cSetPumpNetworkStatusResponse\"\x80\x01\n" +
	"\x14SetPumpStatusRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\bR\x06status\x12@\n" +
	"\x04type\x18\x03 \x01(\x0e2,.city.water.interaction.v1.WaterFacilityTypeR\x04type\"\x17\n" +
	"\x15SetPumpStatusResponse\"*\n" +
	"\x14GetPumpStatusRequest\x12\x12\n" +
	"\x04flag\x18\x01 \x01(\x05R\x04flag\"\xb9\x01\n" +
	"\x15GetPumpStatusResponse\x12a\n" +
	"\vpump_status\x18\x01 \x03(\v2@.city.water.interaction.v1.GetPumpStatusResponse.PumpStatusEntryR\n" +
	"pumpStatus\x1a=\n" +
	"\x0fPumpStatusEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x05R\x05value:\x028\x01\"*\n" +
	"\x14GetNoWaterAOIRequest\x12\x12\n" +
	"\x04flag\x18\x01 \x01(\x05R\x04flag\")\n" +
	"\x15GetNoWaterAOIResponse\x12\x10\n" +
	"\x03aoi\x18\x01 \x03(\x05R\x03aoi\"\x14\n" +
	"\x12GetRuinInfoRequest\"2\n" +
	"\bRuinInfo\x12\x10\n" +
	"\x03num\x18\x01 \x01(\x05R\x03num\x12\x14\n" +
	"\x05ratio\x18\x02 \x01(\x01R\x05ratio\"\xbe\x01\n" +
	"\x13GetRuinInfoResponse\x125\n" +
	"\x03one\x18\x01 \x01(\v2#.city.water.interaction.v1.RuinInfoR\x03one\x125\n" +
	"\x03two\x18\x02 \x01(\v2#.city.water.interaction.v1.RuinInfoR\x03two\x129\n" +
	"\x05three\x18\x03 \x01(\v2#.city.water.interaction.v1.RuinInfoR\x05three*z\n" +
	"\x11WaterFacilityType\x12#\n" +
	"\x1fWATER_FACILITY_TYPE_UNSPECIFIED\x10\x00\x12\x1e\n" +
	"\x1aWATER_FACILITY_TYPE_SUPPLY\x10\x01\x12 \n" +
	"\x1cWATER_FACILITY_TYPE_DRAINAGE\x10\x022\xf2\x05\n" +
	"\fWaterService\x12\x83\x01\n" +
	"\x12SetPumpPowerStatus\x124.city.water.interaction.v1.SetPumpPowerStatusRequest\x1a5.city.water.interaction.v1.SetPumpPowerStatusResponse\"\x00\x12\x89\x01\n" +
	"\x14SetPumpNetworkStatus\x126.city.water.interaction.v1.SetPumpNetworkStatusRequest\x1a7.city.water.interaction.v1.SetPumpNetworkStatusResponse\"\x00\x12t\n" +
	"\rSetPumpStatus\x12/.city.water.interaction.v1.SetPumpStatusRequest\x1a0.city.water.interaction.v1.SetPumpStatusResponse\"\x00\x12t\n" +
	"\rGetPumpStatus\x12/.city.water.interaction.v1.GetPumpStatusRequest\x1a0.city.water.interaction.v1.GetPumpStatusResponse\"\x00\x12t\n" +
	"\rGetNoWaterAOI\x12/.city.water.interaction.v1.GetNoWaterAOIRequest\x1a0.city.water.interaction.v1.GetNoWaterAOIResponse\"\x00\x12n\n" +
	"\vGetRuinInfo\x12-.city.water.interaction.v1.GetRuinInfoRequest\x1a..city.water.interaction.v1.GetRuinInfoResponse\"\x00B\x82\x02\n" +
	"\x1dcom.city.water.interaction.v1B\x11WaterServiceProtoP\x01ZGgit.fiblab.net/sim/protos/v2/go/city/water/interaction/v1;interactionv1\xa2\x02\x03CWI\xaa\x02\x19City.Water.Interaction.V1\xca\x02\x19City\\Water\\Interaction\\V1\xe2\x02%City\\Water\\Interaction\\V1\\GPBMetadata\xea\x02\x1cCity::Water::Interaction::V1b\x06proto3"

var (
	file_city_water_interaction_v1_water_service_proto_rawDescOnce sync.Once
	file_city_water_interaction_v1_water_service_proto_rawDescData []byte
)

func file_city_water_interaction_v1_water_service_proto_rawDescGZIP() []byte {
	file_city_water_interaction_v1_water_service_proto_rawDescOnce.Do(func() {
		file_city_water_interaction_v1_water_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_water_interaction_v1_water_service_proto_rawDesc), len(file_city_water_interaction_v1_water_service_proto_rawDesc)))
	})
	return file_city_water_interaction_v1_water_service_proto_rawDescData
}

var file_city_water_interaction_v1_water_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_city_water_interaction_v1_water_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_city_water_interaction_v1_water_service_proto_goTypes = []any{
	(WaterFacilityType)(0),               // 0: city.water.interaction.v1.WaterFacilityType
	(*SetPumpPowerStatusRequest)(nil),    // 1: city.water.interaction.v1.SetPumpPowerStatusRequest
	(*SetPumpPowerStatusResponse)(nil),   // 2: city.water.interaction.v1.SetPumpPowerStatusResponse
	(*SetPumpNetworkStatusRequest)(nil),  // 3: city.water.interaction.v1.SetPumpNetworkStatusRequest
	(*SetPumpNetworkStatusResponse)(nil), // 4: city.water.interaction.v1.SetPumpNetworkStatusResponse
	(*SetPumpStatusRequest)(nil),         // 5: city.water.interaction.v1.SetPumpStatusRequest
	(*SetPumpStatusResponse)(nil),        // 6: city.water.interaction.v1.SetPumpStatusResponse
	(*GetPumpStatusRequest)(nil),         // 7: city.water.interaction.v1.GetPumpStatusRequest
	(*GetPumpStatusResponse)(nil),        // 8: city.water.interaction.v1.GetPumpStatusResponse
	(*GetNoWaterAOIRequest)(nil),         // 9: city.water.interaction.v1.GetNoWaterAOIRequest
	(*GetNoWaterAOIResponse)(nil),        // 10: city.water.interaction.v1.GetNoWaterAOIResponse
	(*GetRuinInfoRequest)(nil),           // 11: city.water.interaction.v1.GetRuinInfoRequest
	(*RuinInfo)(nil),                     // 12: city.water.interaction.v1.RuinInfo
	(*GetRuinInfoResponse)(nil),          // 13: city.water.interaction.v1.GetRuinInfoResponse
	nil,                                  // 14: city.water.interaction.v1.GetPumpStatusResponse.PumpStatusEntry
}
var file_city_water_interaction_v1_water_service_proto_depIdxs = []int32{
	0,  // 0: city.water.interaction.v1.SetPumpPowerStatusRequest.type:type_name -> city.water.interaction.v1.WaterFacilityType
	0,  // 1: city.water.interaction.v1.SetPumpNetworkStatusRequest.type:type_name -> city.water.interaction.v1.WaterFacilityType
	0,  // 2: city.water.interaction.v1.SetPumpStatusRequest.type:type_name -> city.water.interaction.v1.WaterFacilityType
	14, // 3: city.water.interaction.v1.GetPumpStatusResponse.pump_status:type_name -> city.water.interaction.v1.GetPumpStatusResponse.PumpStatusEntry
	12, // 4: city.water.interaction.v1.GetRuinInfoResponse.one:type_name -> city.water.interaction.v1.RuinInfo
	12, // 5: city.water.interaction.v1.GetRuinInfoResponse.two:type_name -> city.water.interaction.v1.RuinInfo
	12, // 6: city.water.interaction.v1.GetRuinInfoResponse.three:type_name -> city.water.interaction.v1.RuinInfo
	1,  // 7: city.water.interaction.v1.WaterService.SetPumpPowerStatus:input_type -> city.water.interaction.v1.SetPumpPowerStatusRequest
	3,  // 8: city.water.interaction.v1.WaterService.SetPumpNetworkStatus:input_type -> city.water.interaction.v1.SetPumpNetworkStatusRequest
	5,  // 9: city.water.interaction.v1.WaterService.SetPumpStatus:input_type -> city.water.interaction.v1.SetPumpStatusRequest
	7,  // 10: city.water.interaction.v1.WaterService.GetPumpStatus:input_type -> city.water.interaction.v1.GetPumpStatusRequest
	9,  // 11: city.water.interaction.v1.WaterService.GetNoWaterAOI:input_type -> city.water.interaction.v1.GetNoWaterAOIRequest
	11, // 12: city.water.interaction.v1.WaterService.GetRuinInfo:input_type -> city.water.interaction.v1.GetRuinInfoRequest
	2,  // 13: city.water.interaction.v1.WaterService.SetPumpPowerStatus:output_type -> city.water.interaction.v1.SetPumpPowerStatusResponse
	4,  // 14: city.water.interaction.v1.WaterService.SetPumpNetworkStatus:output_type -> city.water.interaction.v1.SetPumpNetworkStatusResponse
	6,  // 15: city.water.interaction.v1.WaterService.SetPumpStatus:output_type -> city.water.interaction.v1.SetPumpStatusResponse
	8,  // 16: city.water.interaction.v1.WaterService.GetPumpStatus:output_type -> city.water.interaction.v1.GetPumpStatusResponse
	10, // 17: city.water.interaction.v1.WaterService.GetNoWaterAOI:output_type -> city.water.interaction.v1.GetNoWaterAOIResponse
	13, // 18: city.water.interaction.v1.WaterService.GetRuinInfo:output_type -> city.water.interaction.v1.GetRuinInfoResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_city_water_interaction_v1_water_service_proto_init() }
func file_city_water_interaction_v1_water_service_proto_init() {
	if File_city_water_interaction_v1_water_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_water_interaction_v1_water_service_proto_rawDesc), len(file_city_water_interaction_v1_water_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_water_interaction_v1_water_service_proto_goTypes,
		DependencyIndexes: file_city_water_interaction_v1_water_service_proto_depIdxs,
		EnumInfos:         file_city_water_interaction_v1_water_service_proto_enumTypes,
		MessageInfos:      file_city_water_interaction_v1_water_service_proto_msgTypes,
	}.Build()
	File_city_water_interaction_v1_water_service_proto = out.File
	file_city_water_interaction_v1_water_service_proto_goTypes = nil
	file_city_water_interaction_v1_water_service_proto_depIdxs = nil
}
