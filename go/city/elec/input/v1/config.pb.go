// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/elec/input/v1/config.proto

package inputv1

import (
	v1 "git.fiblab.net/sim/protos/v2/go/city/config/v1"
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

type Mongo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri        string        `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty" yaml:"uri" bson:"uri" db:"uri"`
	Map        *v1.MongoPath `protobuf:"bytes,2,opt,name=map,proto3" json:"map,omitempty" yaml:"map" bson:"map" db:"map"`
	Facilities *v1.MongoPath `protobuf:"bytes,3,opt,name=facilities,proto3" json:"facilities,omitempty" yaml:"facilities" bson:"facilities" db:"facilities"`
}

func (x *Mongo) Reset() {
	*x = Mongo{}
	mi := &file_city_elec_input_v1_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mongo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mongo) ProtoMessage() {}

func (x *Mongo) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mongo.ProtoReflect.Descriptor instead.
func (*Mongo) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *Mongo) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Mongo) GetMap() *v1.MongoPath {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *Mongo) GetFacilities() *v1.MongoPath {
	if x != nil {
		return x.Facilities
	}
	return nil
}

type ControlStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty" yaml:"start" bson:"start" db:"start"`
	Total int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty" yaml:"total" bson:"total" db:"total"`
}

func (x *ControlStep) Reset() {
	*x = ControlStep{}
	mi := &file_city_elec_input_v1_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControlStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlStep) ProtoMessage() {}

func (x *ControlStep) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlStep.ProtoReflect.Descriptor instead.
func (*ControlStep) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *ControlStep) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ControlStep) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Control struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step *ControlStep `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty" yaml:"step" bson:"step" db:"step"`
}

func (x *Control) Reset() {
	*x = Control{}
	mi := &file_city_elec_input_v1_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Control) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Control) ProtoMessage() {}

func (x *Control) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Control.ProtoReflect.Descriptor instead.
func (*Control) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_config_proto_rawDescGZIP(), []int{2}
}

func (x *Control) GetStep() *ControlStep {
	if x != nil {
		return x.Step
	}
	return nil
}

// 是否输出各类数据
type OutputSwitch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 电网节点状态
	Node  bool `protobuf:"varint,1,opt,name=node,proto3" json:"node,omitempty" yaml:"node" bson:"node" db:"node"`
	Aoi   bool `protobuf:"varint,2,opt,name=aoi,proto3" json:"aoi,omitempty" db:"aoi" yaml:"aoi" bson:"aoi"`
	Event bool `protobuf:"varint,3,opt,name=event,proto3" json:"event,omitempty" yaml:"event" bson:"event" db:"event"`
}

func (x *OutputSwitch) Reset() {
	*x = OutputSwitch{}
	mi := &file_city_elec_input_v1_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutputSwitch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputSwitch) ProtoMessage() {}

func (x *OutputSwitch) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputSwitch.ProtoReflect.Descriptor instead.
func (*OutputSwitch) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_config_proto_rawDescGZIP(), []int{3}
}

func (x *OutputSwitch) GetNode() bool {
	if x != nil {
		return x.Node
	}
	return false
}

func (x *OutputSwitch) GetAoi() bool {
	if x != nil {
		return x.Aoi
	}
	return false
}

func (x *OutputSwitch) GetEvent() bool {
	if x != nil {
		return x.Event
	}
	return false
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *v1.OutputTarget `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty" yaml:"target" bson:"target" db:"target"`
	Switch *OutputSwitch    `protobuf:"bytes,2,opt,name=switch,proto3" json:"switch,omitempty" yaml:"switch" bson:"switch" db:"switch"`
}

func (x *Output) Reset() {
	*x = Output{}
	mi := &file_city_elec_input_v1_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_config_proto_rawDescGZIP(), []int{4}
}

func (x *Output) GetTarget() *v1.OutputTarget {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Output) GetSwitch() *OutputSwitch {
	if x != nil {
		return x.Switch
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mongo   *Mongo   `protobuf:"bytes,1,opt,name=mongo,proto3" json:"mongo,omitempty" yaml:"mongo" bson:"mongo" db:"mongo"`
	Control *Control `protobuf:"bytes,2,opt,name=control,proto3" json:"control,omitempty" yaml:"control" bson:"control" db:"control"`
	Output  *Output  `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty" bson:"output" db:"output" yaml:"output"`
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_city_elec_input_v1_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_config_proto_rawDescGZIP(), []int{5}
}

func (x *Config) GetMongo() *Mongo {
	if x != nil {
		return x.Mongo
	}
	return nil
}

func (x *Config) GetControl() *Control {
	if x != nil {
		return x.Control
	}
	return nil
}

func (x *Config) GetOutput() *Output {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_city_elec_input_v1_config_proto protoreflect.FileDescriptor

var file_city_elec_input_v1_config_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x2b,
	0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e,
	0x67, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x66,
	0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x52, 0x0a, 0x66, 0x61, 0x63, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x53, 0x74, 0x65, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x3e, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x33, 0x0a, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x22, 0x4a, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6f, 0x69, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x61, 0x6f, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x78, 0x0a,
	0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x38, 0x0a,
	0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52,
	0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0xa4, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x2f, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63,
	0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0xcc,
	0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63,
	0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69,
	0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65,
	0x6c, 0x65, 0x63, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x45, 0x49, 0xaa, 0x02, 0x12, 0x43, 0x69, 0x74,
	0x79, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x12, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x6c, 0x65, 0x63, 0x5c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x6c, 0x65, 0x63,
	0x5c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x6c,
	0x65, 0x63, 0x3a, 0x3a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_elec_input_v1_config_proto_rawDescOnce sync.Once
	file_city_elec_input_v1_config_proto_rawDescData = file_city_elec_input_v1_config_proto_rawDesc
)

func file_city_elec_input_v1_config_proto_rawDescGZIP() []byte {
	file_city_elec_input_v1_config_proto_rawDescOnce.Do(func() {
		file_city_elec_input_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_elec_input_v1_config_proto_rawDescData)
	})
	return file_city_elec_input_v1_config_proto_rawDescData
}

var file_city_elec_input_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_city_elec_input_v1_config_proto_goTypes = []any{
	(*Mongo)(nil),           // 0: city.elec.input.v1.Mongo
	(*ControlStep)(nil),     // 1: city.elec.input.v1.ControlStep
	(*Control)(nil),         // 2: city.elec.input.v1.Control
	(*OutputSwitch)(nil),    // 3: city.elec.input.v1.OutputSwitch
	(*Output)(nil),          // 4: city.elec.input.v1.Output
	(*Config)(nil),          // 5: city.elec.input.v1.Config
	(*v1.MongoPath)(nil),    // 6: city.config.v1.MongoPath
	(*v1.OutputTarget)(nil), // 7: city.config.v1.OutputTarget
}
var file_city_elec_input_v1_config_proto_depIdxs = []int32{
	6, // 0: city.elec.input.v1.Mongo.map:type_name -> city.config.v1.MongoPath
	6, // 1: city.elec.input.v1.Mongo.facilities:type_name -> city.config.v1.MongoPath
	1, // 2: city.elec.input.v1.Control.step:type_name -> city.elec.input.v1.ControlStep
	7, // 3: city.elec.input.v1.Output.target:type_name -> city.config.v1.OutputTarget
	3, // 4: city.elec.input.v1.Output.switch:type_name -> city.elec.input.v1.OutputSwitch
	0, // 5: city.elec.input.v1.Config.mongo:type_name -> city.elec.input.v1.Mongo
	2, // 6: city.elec.input.v1.Config.control:type_name -> city.elec.input.v1.Control
	4, // 7: city.elec.input.v1.Config.output:type_name -> city.elec.input.v1.Output
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_city_elec_input_v1_config_proto_init() }
func file_city_elec_input_v1_config_proto_init() {
	if File_city_elec_input_v1_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_elec_input_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_elec_input_v1_config_proto_goTypes,
		DependencyIndexes: file_city_elec_input_v1_config_proto_depIdxs,
		MessageInfos:      file_city_elec_input_v1_config_proto_msgTypes,
	}.Build()
	File_city_elec_input_v1_config_proto = out.File
	file_city_elec_input_v1_config_proto_rawDesc = nil
	file_city_elec_input_v1_config_proto_goTypes = nil
	file_city_elec_input_v1_config_proto_depIdxs = nil
}
