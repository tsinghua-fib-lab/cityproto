// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: city/cognition/output/v1/output.proto

package outputv1

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

// 社交网络节点状态
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status float64 `protobuf:"fixed64,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_output_v1_output_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_output_v1_output_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_city_cognition_output_v1_output_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetStatus() float64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 一次认知影响的过程
type Influence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId int32   `protobuf:"varint,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	TargetId int32   `protobuf:"varint,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Strength float64 `protobuf:"fixed64,3,opt,name=strength,proto3" json:"strength,omitempty"`
}

func (x *Influence) Reset() {
	*x = Influence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_output_v1_output_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Influence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Influence) ProtoMessage() {}

func (x *Influence) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_output_v1_output_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Influence.ProtoReflect.Descriptor instead.
func (*Influence) Descriptor() ([]byte, []int) {
	return file_city_cognition_output_v1_output_proto_rawDescGZIP(), []int{1}
}

func (x *Influence) GetSourceId() int32 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *Influence) GetTargetId() int32 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *Influence) GetStrength() float64 {
	if x != nil {
		return x.Strength
	}
	return 0
}

// 热力图数据
type Heatmap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumRows    int32     `protobuf:"varint,1,opt,name=num_rows,json=numRows,proto3" json:"num_rows,omitempty"`
	NumColumns int32     `protobuf:"varint,2,opt,name=num_columns,json=numColumns,proto3" json:"num_columns,omitempty"`
	Strength   []float64 `protobuf:"fixed64,3,rep,packed,name=strength,proto3" json:"strength,omitempty"`
}

func (x *Heatmap) Reset() {
	*x = Heatmap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_output_v1_output_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heatmap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heatmap) ProtoMessage() {}

func (x *Heatmap) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_output_v1_output_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heatmap.ProtoReflect.Descriptor instead.
func (*Heatmap) Descriptor() ([]byte, []int) {
	return file_city_cognition_output_v1_output_proto_rawDescGZIP(), []int{2}
}

func (x *Heatmap) GetNumRows() int32 {
	if x != nil {
		return x.NumRows
	}
	return 0
}

func (x *Heatmap) GetNumColumns() int32 {
	if x != nil {
		return x.NumColumns
	}
	return 0
}

func (x *Heatmap) GetStrength() []float64 {
	if x != nil {
		return x.Strength
	}
	return nil
}

// 统计数据
type Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrowdCnts   []int32   `protobuf:"varint,1,rep,packed,name=crowd_cnts,json=crowdCnts,proto3" json:"crowd_cnts,omitempty"`
	CrowdRatios []float64 `protobuf:"fixed64,2,rep,packed,name=crowd_ratios,json=crowdRatios,proto3" json:"crowd_ratios,omitempty"`
	KeyNodes    []int32   `protobuf:"varint,3,rep,packed,name=key_nodes,json=keyNodes,proto3" json:"key_nodes,omitempty"`
}

func (x *Stat) Reset() {
	*x = Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_output_v1_output_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat) ProtoMessage() {}

func (x *Stat) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_output_v1_output_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat.ProtoReflect.Descriptor instead.
func (*Stat) Descriptor() ([]byte, []int) {
	return file_city_cognition_output_v1_output_proto_rawDescGZIP(), []int{3}
}

func (x *Stat) GetCrowdCnts() []int32 {
	if x != nil {
		return x.CrowdCnts
	}
	return nil
}

func (x *Stat) GetCrowdRatios() []float64 {
	if x != nil {
		return x.CrowdRatios
	}
	return nil
}

func (x *Stat) GetKeyNodes() []int32 {
	if x != nil {
		return x.KeyNodes
	}
	return nil
}

// 用户发布内容
type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_output_v1_output_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_output_v1_output_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_city_cognition_output_v1_output_proto_rawDescGZIP(), []int{4}
}

func (x *Content) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Content) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// 社交网络节点静态属性
type NodeMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Lng   float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat   float64 `protobuf:"fixed64,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Level int32   `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *NodeMeta) Reset() {
	*x = NodeMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_output_v1_output_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMeta) ProtoMessage() {}

func (x *NodeMeta) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_output_v1_output_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMeta.ProtoReflect.Descriptor instead.
func (*NodeMeta) Descriptor() ([]byte, []int) {
	return file_city_cognition_output_v1_output_proto_rawDescGZIP(), []int{5}
}

func (x *NodeMeta) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeMeta) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *NodeMeta) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *NodeMeta) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type NodesMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*NodeMeta `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *NodesMeta) Reset() {
	*x = NodesMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_output_v1_output_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodesMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodesMeta) ProtoMessage() {}

func (x *NodesMeta) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_output_v1_output_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodesMeta.ProtoReflect.Descriptor instead.
func (*NodesMeta) Descriptor() ([]byte, []int) {
	return file_city_cognition_output_v1_output_proto_rawDescGZIP(), []int{6}
}

func (x *NodesMeta) GetNodes() []*NodeMeta {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// 群体信息
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Size    int32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Changes []int32 `protobuf:"varint,3,rep,packed,name=changes,proto3" json:"changes,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_output_v1_output_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_output_v1_output_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_city_cognition_output_v1_output_proto_rawDescGZIP(), []int{7}
}

func (x *Group) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Group) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Group) GetChanges() []int32 {
	if x != nil {
		return x.Changes
	}
	return nil
}

var File_city_cognition_output_v1_output_proto protoreflect.FileDescriptor

var file_city_cognition_output_v1_output_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76,
	0x31, 0x22, 0x2e, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x61, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x22, 0x61, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x12,
	0x19, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75,
	0x6d, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x01, 0x52, 0x08, 0x73,
	0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x65, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x5f, 0x63, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x43, 0x6e, 0x74, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x2d,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x54, 0x0a,
	0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0x45, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x38, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x42, 0xee, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63,
	0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65,
	0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x4f, 0xaa, 0x02, 0x18, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x43,
	0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x18, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x67, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24,
	0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_cognition_output_v1_output_proto_rawDescOnce sync.Once
	file_city_cognition_output_v1_output_proto_rawDescData = file_city_cognition_output_v1_output_proto_rawDesc
)

func file_city_cognition_output_v1_output_proto_rawDescGZIP() []byte {
	file_city_cognition_output_v1_output_proto_rawDescOnce.Do(func() {
		file_city_cognition_output_v1_output_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_cognition_output_v1_output_proto_rawDescData)
	})
	return file_city_cognition_output_v1_output_proto_rawDescData
}

var file_city_cognition_output_v1_output_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_city_cognition_output_v1_output_proto_goTypes = []interface{}{
	(*Node)(nil),      // 0: city.cognition.output.v1.Node
	(*Influence)(nil), // 1: city.cognition.output.v1.Influence
	(*Heatmap)(nil),   // 2: city.cognition.output.v1.Heatmap
	(*Stat)(nil),      // 3: city.cognition.output.v1.Stat
	(*Content)(nil),   // 4: city.cognition.output.v1.Content
	(*NodeMeta)(nil),  // 5: city.cognition.output.v1.NodeMeta
	(*NodesMeta)(nil), // 6: city.cognition.output.v1.NodesMeta
	(*Group)(nil),     // 7: city.cognition.output.v1.Group
}
var file_city_cognition_output_v1_output_proto_depIdxs = []int32{
	5, // 0: city.cognition.output.v1.NodesMeta.nodes:type_name -> city.cognition.output.v1.NodeMeta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_city_cognition_output_v1_output_proto_init() }
func file_city_cognition_output_v1_output_proto_init() {
	if File_city_cognition_output_v1_output_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_cognition_output_v1_output_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_city_cognition_output_v1_output_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Influence); i {
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
		file_city_cognition_output_v1_output_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heatmap); i {
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
		file_city_cognition_output_v1_output_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stat); i {
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
		file_city_cognition_output_v1_output_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_city_cognition_output_v1_output_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMeta); i {
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
		file_city_cognition_output_v1_output_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodesMeta); i {
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
		file_city_cognition_output_v1_output_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
			RawDescriptor: file_city_cognition_output_v1_output_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_cognition_output_v1_output_proto_goTypes,
		DependencyIndexes: file_city_cognition_output_v1_output_proto_depIdxs,
		MessageInfos:      file_city_cognition_output_v1_output_proto_msgTypes,
	}.Build()
	File_city_cognition_output_v1_output_proto = out.File
	file_city_cognition_output_v1_output_proto_rawDesc = nil
	file_city_cognition_output_v1_output_proto_goTypes = nil
	file_city_cognition_output_v1_output_proto_depIdxs = nil
}
