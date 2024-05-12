// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: city/streetview/v1/streetview.proto

package streetviewv1

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

// 街景图片描述
type StreetViewImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 朝向，单位度，0-360，0为正北，90为正东，180为正南，270为正西
	Heading float64 `protobuf:"fixed64,1,opt,name=heading,proto3" json:"heading,omitempty"`
	// 对象存储的object key
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *StreetViewImage) Reset() {
	*x = StreetViewImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_streetview_v1_streetview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreetViewImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreetViewImage) ProtoMessage() {}

func (x *StreetViewImage) ProtoReflect() protoreflect.Message {
	mi := &file_city_streetview_v1_streetview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreetViewImage.ProtoReflect.Descriptor instead.
func (*StreetViewImage) Descriptor() ([]byte, []int) {
	return file_city_streetview_v1_streetview_proto_rawDescGZIP(), []int{0}
}

func (x *StreetViewImage) GetHeading() float64 {
	if x != nil {
		return x.Heading
	}
	return 0
}

func (x *StreetViewImage) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

// 街景图片元数据
type StreetView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// WGS84经纬度位置
	Lnglat *v2.LongLatPosition `protobuf:"bytes,1,opt,name=lnglat,proto3" json:"lnglat,omitempty"`
	// 该位置的不同朝向街景图列表
	Images []*StreetViewImage `protobuf:"bytes,2,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *StreetView) Reset() {
	*x = StreetView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_streetview_v1_streetview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreetView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreetView) ProtoMessage() {}

func (x *StreetView) ProtoReflect() protoreflect.Message {
	mi := &file_city_streetview_v1_streetview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreetView.ProtoReflect.Descriptor instead.
func (*StreetView) Descriptor() ([]byte, []int) {
	return file_city_streetview_v1_streetview_proto_rawDescGZIP(), []int{1}
}

func (x *StreetView) GetLnglat() *v2.LongLatPosition {
	if x != nil {
		return x.Lnglat
	}
	return nil
}

func (x *StreetView) GetImages() []*StreetViewImage {
	if x != nil {
		return x.Images
	}
	return nil
}

// 导出Message，对应一个MongoDB Collection
type StreetViews struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetViews []*StreetView `protobuf:"bytes,1,rep,name=street_views,json=streetViews,proto3" json:"street_views,omitempty"`
}

func (x *StreetViews) Reset() {
	*x = StreetViews{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_streetview_v1_streetview_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreetViews) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreetViews) ProtoMessage() {}

func (x *StreetViews) ProtoReflect() protoreflect.Message {
	mi := &file_city_streetview_v1_streetview_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreetViews.ProtoReflect.Descriptor instead.
func (*StreetViews) Descriptor() ([]byte, []int) {
	return file_city_streetview_v1_streetview_proto_rawDescGZIP(), []int{2}
}

func (x *StreetViews) GetStreetViews() []*StreetView {
	if x != nil {
		return x.StreetViews
	}
	return nil
}

var File_city_streetview_v1_streetview_proto protoreflect.FileDescriptor

var file_city_streetview_v1_streetview_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x76, 0x69, 0x65,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63, 0x69, 0x74, 0x79, 0x2f,
	0x67, 0x65, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x43, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x7f, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x34, 0x0a, 0x06, 0x6c, 0x6e, 0x67, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76,
	0x32, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x6c, 0x6e, 0x67, 0x6c, 0x61, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x50, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x56, 0x69, 0x65, 0x77, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x73, 0x42, 0xd1, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x76, 0x31, 0x42, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x76, 0x69, 0x65, 0x77, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c,
	0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x76, 0x69, 0x65, 0x77, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x76, 0x69,
	0x65, 0x77, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x69, 0x74,
	0x79, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x12, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x76, 0x69, 0x65,
	0x77, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x53, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x76, 0x69, 0x65, 0x77, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x53, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x76, 0x69, 0x65, 0x77, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_streetview_v1_streetview_proto_rawDescOnce sync.Once
	file_city_streetview_v1_streetview_proto_rawDescData = file_city_streetview_v1_streetview_proto_rawDesc
)

func file_city_streetview_v1_streetview_proto_rawDescGZIP() []byte {
	file_city_streetview_v1_streetview_proto_rawDescOnce.Do(func() {
		file_city_streetview_v1_streetview_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_streetview_v1_streetview_proto_rawDescData)
	})
	return file_city_streetview_v1_streetview_proto_rawDescData
}

var file_city_streetview_v1_streetview_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_city_streetview_v1_streetview_proto_goTypes = []interface{}{
	(*StreetViewImage)(nil),    // 0: city.streetview.v1.StreetViewImage
	(*StreetView)(nil),         // 1: city.streetview.v1.StreetView
	(*StreetViews)(nil),        // 2: city.streetview.v1.StreetViews
	(*v2.LongLatPosition)(nil), // 3: city.geo.v2.LongLatPosition
}
var file_city_streetview_v1_streetview_proto_depIdxs = []int32{
	3, // 0: city.streetview.v1.StreetView.lnglat:type_name -> city.geo.v2.LongLatPosition
	0, // 1: city.streetview.v1.StreetView.images:type_name -> city.streetview.v1.StreetViewImage
	1, // 2: city.streetview.v1.StreetViews.street_views:type_name -> city.streetview.v1.StreetView
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_city_streetview_v1_streetview_proto_init() }
func file_city_streetview_v1_streetview_proto_init() {
	if File_city_streetview_v1_streetview_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_streetview_v1_streetview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreetViewImage); i {
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
		file_city_streetview_v1_streetview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreetView); i {
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
		file_city_streetview_v1_streetview_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreetViews); i {
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
			RawDescriptor: file_city_streetview_v1_streetview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_streetview_v1_streetview_proto_goTypes,
		DependencyIndexes: file_city_streetview_v1_streetview_proto_depIdxs,
		MessageInfos:      file_city_streetview_v1_streetview_proto_msgTypes,
	}.Build()
	File_city_streetview_v1_streetview_proto = out.File
	file_city_streetview_v1_streetview_proto_rawDesc = nil
	file_city_streetview_v1_streetview_proto_goTypes = nil
	file_city_streetview_v1_streetview_proto_depIdxs = nil
}
