// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: city/elec/input/v1/input.proto

package inputv1

import (
	v2 "git.fiblab.net/sim/protos/go/city/geo/v2"
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

type FacilityType int32

const (
	// 电网相关的基础设施分类
	FacilityType_FACILITY_TYPE_UNSPECIFIED     FacilityType = 0
	FacilityType_FACILITY_TYPE_POWER_STATION   FacilityType = 1 //发电站
	FacilityType_FACILITY_TYPE_TRANSFORMER_500 FacilityType = 2 //不同电压的变压器
	FacilityType_FACILITY_TYPE_TRANSFORMER_220 FacilityType = 3
	FacilityType_FACILITY_TYPE_TRANSFORMER_110 FacilityType = 4
	FacilityType_FACILITY_TYPE_TRANSFORMER_10  FacilityType = 5
	// 通信基站
	FacilityType_FACILITY_TYPE_BASE_STATION FacilityType = 6
	// 网关
	FacilityType_FACILITY_TYPE_GATEWAY FacilityType = 7
	// 排水水泵
	FacilityType_FACILITY_TYPE_DRAINAGE_PUMP FacilityType = 8
	// 交通灯
	FacilityType_FACILITY_TYPE_TRAFFIC_LIGHT FacilityType = 9
	// AOI
	FacilityType_FACILITY_TYPE_AOI FacilityType = 10
	// 供水水泵
	FacilityType_FACILITY_TYPE_SUPPLY_PUMP FacilityType = 11
)

// Enum value maps for FacilityType.
var (
	FacilityType_name = map[int32]string{
		0:  "FACILITY_TYPE_UNSPECIFIED",
		1:  "FACILITY_TYPE_POWER_STATION",
		2:  "FACILITY_TYPE_TRANSFORMER_500",
		3:  "FACILITY_TYPE_TRANSFORMER_220",
		4:  "FACILITY_TYPE_TRANSFORMER_110",
		5:  "FACILITY_TYPE_TRANSFORMER_10",
		6:  "FACILITY_TYPE_BASE_STATION",
		7:  "FACILITY_TYPE_GATEWAY",
		8:  "FACILITY_TYPE_DRAINAGE_PUMP",
		9:  "FACILITY_TYPE_TRAFFIC_LIGHT",
		10: "FACILITY_TYPE_AOI",
		11: "FACILITY_TYPE_SUPPLY_PUMP",
	}
	FacilityType_value = map[string]int32{
		"FACILITY_TYPE_UNSPECIFIED":     0,
		"FACILITY_TYPE_POWER_STATION":   1,
		"FACILITY_TYPE_TRANSFORMER_500": 2,
		"FACILITY_TYPE_TRANSFORMER_220": 3,
		"FACILITY_TYPE_TRANSFORMER_110": 4,
		"FACILITY_TYPE_TRANSFORMER_10":  5,
		"FACILITY_TYPE_BASE_STATION":    6,
		"FACILITY_TYPE_GATEWAY":         7,
		"FACILITY_TYPE_DRAINAGE_PUMP":   8,
		"FACILITY_TYPE_TRAFFIC_LIGHT":   9,
		"FACILITY_TYPE_AOI":             10,
		"FACILITY_TYPE_SUPPLY_PUMP":     11,
	}
)

func (x FacilityType) Enum() *FacilityType {
	p := new(FacilityType)
	*p = x
	return p
}

func (x FacilityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FacilityType) Descriptor() protoreflect.EnumDescriptor {
	return file_city_elec_input_v1_input_proto_enumTypes[0].Descriptor()
}

func (FacilityType) Type() protoreflect.EnumType {
	return &file_city_elec_input_v1_input_proto_enumTypes[0]
}

func (x FacilityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FacilityType.Descriptor instead.
func (FacilityType) EnumDescriptor() ([]byte, []int) {
	return file_city_elec_input_v1_input_proto_rawDescGZIP(), []int{0}
}

type RepairStation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AoiId    int32               `protobuf:"varint,1,opt,name=aoi_id,json=aoiId,proto3" json:"aoi_id,omitempty" yaml:"aoi_id" bson:"aoi_id" db:"aoi_id"`
	Position *v2.LongLatPosition `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty" yaml:"position" bson:"position" db:"position"`
}

func (x *RepairStation) Reset() {
	*x = RepairStation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_input_v1_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepairStation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepairStation) ProtoMessage() {}

func (x *RepairStation) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepairStation.ProtoReflect.Descriptor instead.
func (*RepairStation) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_input_proto_rawDescGZIP(), []int{0}
}

func (x *RepairStation) GetAoiId() int32 {
	if x != nil {
		return x.AoiId
	}
	return 0
}

func (x *RepairStation) GetPosition() *v2.LongLatPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

type Facility struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" yaml:"id" bson:"id"`
	Position *v2.LongLatPosition `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty" yaml:"position" bson:"position" db:"position"`
	Type     FacilityType        `protobuf:"varint,3,opt,name=type,proto3,enum=city.elec.input.v1.FacilityType" json:"type,omitempty" yaml:"type" bson:"type" db:"type"`
	// 当前节点的邻居节点的id
	Relation []int32 `protobuf:"varint,4,rep,packed,name=relation,proto3" json:"relation,omitempty" yaml:"relation" bson:"relation" db:"relation"`
	// 在其它关联的网络中如水网使用时，可使用外部id
	// 对于负载，该值表示其在对应模拟中的id
	ForeignId *int32 `protobuf:"varint,5,opt,name=foreign_id,json=foreignId,proto3,oneof" json:"foreign_id,omitempty" yaml:"foreign_id" bson:"foreign_id" db:"foreign_id"`
	// 对于电力设施，该值表示所在aoi id
	AoiId *int32 `protobuf:"varint,6,opt,name=aoi_id,json=aoiId,proto3,oneof" json:"aoi_id,omitempty" db:"aoi_id" yaml:"aoi_id" bson:"aoi_id"`
	// 对于10kv变压器组，该值表示变压器组中变压器的数量
	NumTransformer *int32 `protobuf:"varint,7,opt,name=num_transformer,json=numTransformer,proto3,oneof" json:"num_transformer,omitempty" yaml:"num_transformer" bson:"num_transformer" db:"num_transformer"`
}

func (x *Facility) Reset() {
	*x = Facility{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_input_v1_input_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Facility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Facility) ProtoMessage() {}

func (x *Facility) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_input_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Facility.ProtoReflect.Descriptor instead.
func (*Facility) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_input_proto_rawDescGZIP(), []int{1}
}

func (x *Facility) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Facility) GetPosition() *v2.LongLatPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Facility) GetType() FacilityType {
	if x != nil {
		return x.Type
	}
	return FacilityType_FACILITY_TYPE_UNSPECIFIED
}

func (x *Facility) GetRelation() []int32 {
	if x != nil {
		return x.Relation
	}
	return nil
}

func (x *Facility) GetForeignId() int32 {
	if x != nil && x.ForeignId != nil {
		return *x.ForeignId
	}
	return 0
}

func (x *Facility) GetAoiId() int32 {
	if x != nil && x.AoiId != nil {
		return *x.AoiId
	}
	return 0
}

func (x *Facility) GetNumTransformer() int32 {
	if x != nil && x.NumTransformer != nil {
		return *x.NumTransformer
	}
	return 0
}

// 设施集合，对应于mongodb一个collection
type Facilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Facilities     []*Facility      `protobuf:"bytes,1,rep,name=facilities,proto3" json:"facilities,omitempty" yaml:"facilities" bson:"facilities" db:"facilities"`
	RepairStations []*RepairStation `protobuf:"bytes,2,rep,name=repair_stations,json=repairStations,proto3" json:"repair_stations,omitempty" yaml:"repair_stations" bson:"repair_stations" db:"repair_stations"`
}

func (x *Facilities) Reset() {
	*x = Facilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_input_v1_input_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Facilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Facilities) ProtoMessage() {}

func (x *Facilities) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_input_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Facilities.ProtoReflect.Descriptor instead.
func (*Facilities) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_input_proto_rawDescGZIP(), []int{2}
}

func (x *Facilities) GetFacilities() []*Facility {
	if x != nil {
		return x.Facilities
	}
	return nil
}

func (x *Facilities) GetRepairStations() []*RepairStation {
	if x != nil {
		return x.RepairStations
	}
	return nil
}

var File_city_elec_input_v1_input_proto protoreflect.FileDescriptor

var file_city_elec_input_v1_input_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x76,
	0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x0d, 0x52,
	0x65, 0x70, 0x61, 0x69, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x6f, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x6f,
	0x69, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f,
	0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc2, 0x02,
	0x0a, 0x08, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4c,
	0x61, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x66, 0x6f,
	0x72, 0x65, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06, 0x61, 0x6f,
	0x69, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05, 0x61, 0x6f,
	0x69, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x02, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e,
	0x5f, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6f, 0x69, 0x5f, 0x69, 0x64, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x65, 0x72, 0x22, 0x96, 0x01, 0x0a, 0x0a, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65,
	0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x63, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x0a, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x4a, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x70,
	0x61, 0x69, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x8c, 0x03, 0x0a, 0x0c,
	0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19,
	0x46, 0x41, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x46,
	0x41, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x57,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d,
	0x46, 0x41, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x52, 0x5f, 0x35, 0x30, 0x30, 0x10, 0x02, 0x12,
	0x21, 0x0a, 0x1d, 0x46, 0x41, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x52, 0x5f, 0x32, 0x32, 0x30,
	0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x46, 0x41, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x52, 0x5f,
	0x31, 0x31, 0x30, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x46, 0x41, 0x43, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x4f, 0x52, 0x4d,
	0x45, 0x52, 0x5f, 0x31, 0x30, 0x10, 0x05, 0x12, 0x1e, 0x0a, 0x1a, 0x46, 0x41, 0x43, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x41, 0x43, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59,
	0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x41, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x52, 0x41, 0x49, 0x4e, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x55, 0x4d,
	0x50, 0x10, 0x08, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x41, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x4c, 0x49, 0x47,
	0x48, 0x54, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x41, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4f, 0x49, 0x10, 0x0a, 0x12, 0x1d, 0x0a, 0x19, 0x46,
	0x41, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x55, 0x50,
	0x50, 0x4c, 0x59, 0x5f, 0x50, 0x55, 0x4d, 0x50, 0x10, 0x0b, 0x42, 0xc8, 0x01, 0x0a, 0x16, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e,
	0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x45, 0x49, 0xaa, 0x02, 0x12, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x2e, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45,
	0x6c, 0x65, 0x63, 0x5c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43,
	0x69, 0x74, 0x79, 0x5c, 0x45, 0x6c, 0x65, 0x63, 0x5c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15,
	0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x6c, 0x65, 0x63, 0x3a, 0x3a, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_elec_input_v1_input_proto_rawDescOnce sync.Once
	file_city_elec_input_v1_input_proto_rawDescData = file_city_elec_input_v1_input_proto_rawDesc
)

func file_city_elec_input_v1_input_proto_rawDescGZIP() []byte {
	file_city_elec_input_v1_input_proto_rawDescOnce.Do(func() {
		file_city_elec_input_v1_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_elec_input_v1_input_proto_rawDescData)
	})
	return file_city_elec_input_v1_input_proto_rawDescData
}

var file_city_elec_input_v1_input_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_city_elec_input_v1_input_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_city_elec_input_v1_input_proto_goTypes = []interface{}{
	(FacilityType)(0),          // 0: city.elec.input.v1.FacilityType
	(*RepairStation)(nil),      // 1: city.elec.input.v1.RepairStation
	(*Facility)(nil),           // 2: city.elec.input.v1.Facility
	(*Facilities)(nil),         // 3: city.elec.input.v1.Facilities
	(*v2.LongLatPosition)(nil), // 4: city.geo.v2.LongLatPosition
}
var file_city_elec_input_v1_input_proto_depIdxs = []int32{
	4, // 0: city.elec.input.v1.RepairStation.position:type_name -> city.geo.v2.LongLatPosition
	4, // 1: city.elec.input.v1.Facility.position:type_name -> city.geo.v2.LongLatPosition
	0, // 2: city.elec.input.v1.Facility.type:type_name -> city.elec.input.v1.FacilityType
	2, // 3: city.elec.input.v1.Facilities.facilities:type_name -> city.elec.input.v1.Facility
	1, // 4: city.elec.input.v1.Facilities.repair_stations:type_name -> city.elec.input.v1.RepairStation
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_city_elec_input_v1_input_proto_init() }
func file_city_elec_input_v1_input_proto_init() {
	if File_city_elec_input_v1_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_elec_input_v1_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepairStation); i {
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
		file_city_elec_input_v1_input_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Facility); i {
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
		file_city_elec_input_v1_input_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Facilities); i {
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
	file_city_elec_input_v1_input_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_elec_input_v1_input_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_elec_input_v1_input_proto_goTypes,
		DependencyIndexes: file_city_elec_input_v1_input_proto_depIdxs,
		EnumInfos:         file_city_elec_input_v1_input_proto_enumTypes,
		MessageInfos:      file_city_elec_input_v1_input_proto_msgTypes,
	}.Build()
	File_city_elec_input_v1_input_proto = out.File
	file_city_elec_input_v1_input_proto_rawDesc = nil
	file_city_elec_input_v1_input_proto_goTypes = nil
	file_city_elec_input_v1_input_proto_depIdxs = nil
}
