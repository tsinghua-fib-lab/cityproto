// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/cognition/input/v1/input.proto

package inputv1

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

// 社交网络连接关系
type Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source int32 `protobuf:"varint,1,opt,name=source,proto3" json:"source,omitempty"`
	Target int32 `protobuf:"varint,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *Edge) Reset() {
	*x = Edge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_input_v1_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edge) ProtoMessage() {}

func (x *Edge) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_input_v1_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edge.ProtoReflect.Descriptor instead.
func (*Edge) Descriptor() ([]byte, []int) {
	return file_city_cognition_input_v1_input_proto_rawDescGZIP(), []int{0}
}

func (x *Edge) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *Edge) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

type Edges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Edges []*Edge `protobuf:"bytes,1,rep,name=edges,proto3" json:"edges,omitempty"`
}

func (x *Edges) Reset() {
	*x = Edges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_cognition_input_v1_input_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edges) ProtoMessage() {}

func (x *Edges) ProtoReflect() protoreflect.Message {
	mi := &file_city_cognition_input_v1_input_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edges.ProtoReflect.Descriptor instead.
func (*Edges) Descriptor() ([]byte, []int) {
	return file_city_cognition_input_v1_input_proto_rawDescGZIP(), []int{1}
}

func (x *Edges) GetEdges() []*Edge {
	if x != nil {
		return x.Edges
	}
	return nil
}

var File_city_cognition_input_v1_input_proto protoreflect.FileDescriptor

var file_city_cognition_input_v1_input_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x36,
	0x0a, 0x04, 0x45, 0x64, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x3c, 0x0a, 0x05, 0x45, 0x64, 0x67, 0x65, 0x73, 0x12,
	0x33, 0x0a, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65,
	0x64, 0x67, 0x65, 0x73, 0x42, 0xe6, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e,
	0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x43, 0x49, 0xaa, 0x02, 0x17, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x17, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x43, 0x69, 0x74,
	0x79, 0x5c, 0x43, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1a, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x3a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_cognition_input_v1_input_proto_rawDescOnce sync.Once
	file_city_cognition_input_v1_input_proto_rawDescData = file_city_cognition_input_v1_input_proto_rawDesc
)

func file_city_cognition_input_v1_input_proto_rawDescGZIP() []byte {
	file_city_cognition_input_v1_input_proto_rawDescOnce.Do(func() {
		file_city_cognition_input_v1_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_cognition_input_v1_input_proto_rawDescData)
	})
	return file_city_cognition_input_v1_input_proto_rawDescData
}

var file_city_cognition_input_v1_input_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_cognition_input_v1_input_proto_goTypes = []any{
	(*Edge)(nil),  // 0: city.cognition.input.v1.Edge
	(*Edges)(nil), // 1: city.cognition.input.v1.Edges
}
var file_city_cognition_input_v1_input_proto_depIdxs = []int32{
	0, // 0: city.cognition.input.v1.Edges.edges:type_name -> city.cognition.input.v1.Edge
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_city_cognition_input_v1_input_proto_init() }
func file_city_cognition_input_v1_input_proto_init() {
	if File_city_cognition_input_v1_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_cognition_input_v1_input_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Edge); i {
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
		file_city_cognition_input_v1_input_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Edges); i {
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
			RawDescriptor: file_city_cognition_input_v1_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_cognition_input_v1_input_proto_goTypes,
		DependencyIndexes: file_city_cognition_input_v1_input_proto_depIdxs,
		MessageInfos:      file_city_cognition_input_v1_input_proto_msgTypes,
	}.Build()
	File_city_cognition_input_v1_input_proto = out.File
	file_city_cognition_input_v1_input_proto_rawDesc = nil
	file_city_cognition_input_v1_input_proto_goTypes = nil
	file_city_cognition_input_v1_input_proto_depIdxs = nil
}
