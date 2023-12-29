// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: city/elec/input/v1/input_service.proto

package inputv1

import (
	v2 "git.fiblab.net/sim/protos/go/city/map/v2"
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

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_input_v1_input_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_input_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_input_service_proto_rawDescGZIP(), []int{0}
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 模拟器gRPC监听地址
	Address    string      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Control    *Control    `protobuf:"bytes,3,opt,name=control,proto3" json:"control,omitempty"`
	Facilities *Facilities `protobuf:"bytes,1,opt,name=facilities,proto3" json:"facilities,omitempty"`
	Map        *v2.Map     `protobuf:"bytes,4,opt,name=map,proto3" json:"map,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_elec_input_v1_input_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_elec_input_v1_input_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_city_elec_input_v1_input_service_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InitResponse) GetControl() *Control {
	if x != nil {
		return x.Control
	}
	return nil
}

func (x *InitResponse) GetFacilities() *Facilities {
	if x != nil {
		return x.Facilities
	}
	return nil
}

func (x *InitResponse) GetMap() *v2.Map {
	if x != nil {
		return x.Map
	}
	return nil
}

var File_city_elec_input_v1_input_service_proto protoreflect.FileDescriptor

var file_city_elec_input_v1_input_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65,
	0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x63, 0x69,
	0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x61, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xc3, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x35,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x3e, 0x0a, 0x0a, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0a, 0x66, 0x61, 0x63, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x32,
	0x2e, 0x4d, 0x61, 0x70, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x32, 0x5b, 0x0a, 0x0c, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x49, 0x6e, 0x69,
	0x74, 0x12, 0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xcf, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x11, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c,
	0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x2f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x45, 0x49, 0xaa, 0x02, 0x12, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6c, 0x65,
	0x63, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x43, 0x69, 0x74,
	0x79, 0x5c, 0x45, 0x6c, 0x65, 0x63, 0x5c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1e, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x6c, 0x65, 0x63, 0x5c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x15, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x6c, 0x65, 0x63, 0x3a, 0x3a, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_elec_input_v1_input_service_proto_rawDescOnce sync.Once
	file_city_elec_input_v1_input_service_proto_rawDescData = file_city_elec_input_v1_input_service_proto_rawDesc
)

func file_city_elec_input_v1_input_service_proto_rawDescGZIP() []byte {
	file_city_elec_input_v1_input_service_proto_rawDescOnce.Do(func() {
		file_city_elec_input_v1_input_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_elec_input_v1_input_service_proto_rawDescData)
	})
	return file_city_elec_input_v1_input_service_proto_rawDescData
}

var file_city_elec_input_v1_input_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_elec_input_v1_input_service_proto_goTypes = []interface{}{
	(*InitRequest)(nil),  // 0: city.elec.input.v1.InitRequest
	(*InitResponse)(nil), // 1: city.elec.input.v1.InitResponse
	(*Control)(nil),      // 2: city.elec.input.v1.Control
	(*Facilities)(nil),   // 3: city.elec.input.v1.Facilities
	(*v2.Map)(nil),       // 4: city.map.v2.Map
}
var file_city_elec_input_v1_input_service_proto_depIdxs = []int32{
	2, // 0: city.elec.input.v1.InitResponse.control:type_name -> city.elec.input.v1.Control
	3, // 1: city.elec.input.v1.InitResponse.facilities:type_name -> city.elec.input.v1.Facilities
	4, // 2: city.elec.input.v1.InitResponse.map:type_name -> city.map.v2.Map
	0, // 3: city.elec.input.v1.InputService.Init:input_type -> city.elec.input.v1.InitRequest
	1, // 4: city.elec.input.v1.InputService.Init:output_type -> city.elec.input.v1.InitResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_city_elec_input_v1_input_service_proto_init() }
func file_city_elec_input_v1_input_service_proto_init() {
	if File_city_elec_input_v1_input_service_proto != nil {
		return
	}
	file_city_elec_input_v1_config_proto_init()
	file_city_elec_input_v1_input_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_elec_input_v1_input_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_city_elec_input_v1_input_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
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
			RawDescriptor: file_city_elec_input_v1_input_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_elec_input_v1_input_service_proto_goTypes,
		DependencyIndexes: file_city_elec_input_v1_input_service_proto_depIdxs,
		MessageInfos:      file_city_elec_input_v1_input_service_proto_msgTypes,
	}.Build()
	File_city_elec_input_v1_input_service_proto = out.File
	file_city_elec_input_v1_input_service_proto_rawDesc = nil
	file_city_elec_input_v1_input_service_proto_goTypes = nil
	file_city_elec_input_v1_input_service_proto_depIdxs = nil
}
