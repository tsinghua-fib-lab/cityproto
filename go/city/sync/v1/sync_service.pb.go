// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/sync/v1/sync_service.proto

package syncv1

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

// 注册程序URL请求
type SetURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组件名，需要在同步器启动参数列表中
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 程序URL
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *SetURLRequest) Reset() {
	*x = SetURLRequest{}
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetURLRequest) ProtoMessage() {}

func (x *SetURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetURLRequest.ProtoReflect.Descriptor instead.
func (*SetURLRequest) Descriptor() ([]byte, []int) {
	return file_city_sync_v1_sync_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetURLRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SetURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// 注册程序URL响应
type SetURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetURLResponse) Reset() {
	*x = SetURLResponse{}
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetURLResponse) ProtoMessage() {}

func (x *SetURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetURLResponse.ProtoReflect.Descriptor instead.
func (*SetURLResponse) Descriptor() ([]byte, []int) {
	return file_city_sync_v1_sync_service_proto_rawDescGZIP(), []int{1}
}

// 获取程序URL请求
type GetURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组件名，需要在同步器启动参数列表中
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetURLRequest) Reset() {
	*x = GetURLRequest{}
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLRequest) ProtoMessage() {}

func (x *GetURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLRequest.ProtoReflect.Descriptor instead.
func (*GetURLRequest) Descriptor() ([]byte, []int) {
	return file_city_sync_v1_sync_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetURLRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 获取程序URL响应
type GetURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 程序URL
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetURLResponse) Reset() {
	*x = GetURLResponse{}
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLResponse) ProtoMessage() {}

func (x *GetURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLResponse.ProtoReflect.Descriptor instead.
func (*GetURLResponse) Descriptor() ([]byte, []int) {
	return file_city_sync_v1_sync_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// 步进请求
type StepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组件名，需要在同步器启动参数列表中
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 步进的步数，当step为-1时表示关闭请求，初始化时需要进行一次step=1的同步
	Step int32 `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *StepRequest) Reset() {
	*x = StepRequest{}
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepRequest) ProtoMessage() {}

func (x *StepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepRequest.ProtoReflect.Descriptor instead.
func (*StepRequest) Descriptor() ([]byte, []int) {
	return file_city_sync_v1_sync_service_proto_rawDescGZIP(), []int{4}
}

func (x *StepRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StepRequest) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

// 步进响应
type StepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务是否关闭
	Close bool `protobuf:"varint,1,opt,name=close,proto3" json:"close,omitempty"`
}

func (x *StepResponse) Reset() {
	*x = StepResponse{}
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepResponse) ProtoMessage() {}

func (x *StepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_sync_v1_sync_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepResponse.ProtoReflect.Descriptor instead.
func (*StepResponse) Descriptor() ([]byte, []int) {
	return file_city_sync_v1_sync_service_proto_rawDescGZIP(), []int{5}
}

func (x *StepResponse) GetClose() bool {
	if x != nil {
		return x.Close
	}
	return false
}

var File_city_sync_v1_sync_service_proto protoreflect.FileDescriptor

var file_city_sync_v1_sync_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x22,
	0x35, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x74, 0x65, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x32, 0xd6,
	0x01, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43,
	0x0a, 0x06, 0x53, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x1b, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70,
	0x12, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa8, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x53, 0x79,
	0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74,
	0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x79, 0x6e, 0x63,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0c, 0x43, 0x69, 0x74, 0x79, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x53,
	0x79, 0x6e, 0x63, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x53, 0x79,
	0x6e, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x53, 0x79, 0x6e, 0x63, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_sync_v1_sync_service_proto_rawDescOnce sync.Once
	file_city_sync_v1_sync_service_proto_rawDescData = file_city_sync_v1_sync_service_proto_rawDesc
)

func file_city_sync_v1_sync_service_proto_rawDescGZIP() []byte {
	file_city_sync_v1_sync_service_proto_rawDescOnce.Do(func() {
		file_city_sync_v1_sync_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_sync_v1_sync_service_proto_rawDescData)
	})
	return file_city_sync_v1_sync_service_proto_rawDescData
}

var file_city_sync_v1_sync_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_city_sync_v1_sync_service_proto_goTypes = []any{
	(*SetURLRequest)(nil),  // 0: city.sync.v1.SetURLRequest
	(*SetURLResponse)(nil), // 1: city.sync.v1.SetURLResponse
	(*GetURLRequest)(nil),  // 2: city.sync.v1.GetURLRequest
	(*GetURLResponse)(nil), // 3: city.sync.v1.GetURLResponse
	(*StepRequest)(nil),    // 4: city.sync.v1.StepRequest
	(*StepResponse)(nil),   // 5: city.sync.v1.StepResponse
}
var file_city_sync_v1_sync_service_proto_depIdxs = []int32{
	0, // 0: city.sync.v1.SyncService.SetURL:input_type -> city.sync.v1.SetURLRequest
	2, // 1: city.sync.v1.SyncService.GetURL:input_type -> city.sync.v1.GetURLRequest
	4, // 2: city.sync.v1.SyncService.Step:input_type -> city.sync.v1.StepRequest
	1, // 3: city.sync.v1.SyncService.SetURL:output_type -> city.sync.v1.SetURLResponse
	3, // 4: city.sync.v1.SyncService.GetURL:output_type -> city.sync.v1.GetURLResponse
	5, // 5: city.sync.v1.SyncService.Step:output_type -> city.sync.v1.StepResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_sync_v1_sync_service_proto_init() }
func file_city_sync_v1_sync_service_proto_init() {
	if File_city_sync_v1_sync_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_sync_v1_sync_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_sync_v1_sync_service_proto_goTypes,
		DependencyIndexes: file_city_sync_v1_sync_service_proto_depIdxs,
		MessageInfos:      file_city_sync_v1_sync_service_proto_msgTypes,
	}.Build()
	File_city_sync_v1_sync_service_proto = out.File
	file_city_sync_v1_sync_service_proto_rawDesc = nil
	file_city_sync_v1_sync_service_proto_goTypes = nil
	file_city_sync_v1_sync_service_proto_depIdxs = nil
}
