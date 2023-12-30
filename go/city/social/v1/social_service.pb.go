// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: city/social/v1/social_service.proto

package socialv1

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

// 发送消息请求
type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待发送的消息
	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_social_v1_social_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_social_v1_social_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_city_social_v1_social_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

// 发送消息响应
type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_social_v1_social_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_social_v1_social_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_city_social_v1_social_service_proto_rawDescGZIP(), []int{1}
}

// 接收消息请求
type ReceiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息接收者ID（即为自身ID）
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReceiveRequest) Reset() {
	*x = ReceiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_social_v1_social_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveRequest) ProtoMessage() {}

func (x *ReceiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_social_v1_social_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveRequest.ProtoReflect.Descriptor instead.
func (*ReceiveRequest) Descriptor() ([]byte, []int) {
	return file_city_social_v1_social_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReceiveRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 接收消息响应
type ReceiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 接收到的消息
	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ReceiveResponse) Reset() {
	*x = ReceiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_social_v1_social_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveResponse) ProtoMessage() {}

func (x *ReceiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_social_v1_social_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveResponse.ProtoReflect.Descriptor instead.
func (*ReceiveResponse) Descriptor() ([]byte, []int) {
	return file_city_social_v1_social_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiveResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_city_social_v1_social_service_proto protoreflect.FileDescriptor

var file_city_social_v1_social_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x0f, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x32, 0x9e, 0x01, 0x0a, 0x0d, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x12, 0x1e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xb8, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x53, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f,
	0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69,
	0x74, 0x79, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0e, 0x43, 0x69,
	0x74, 0x79, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x43,
	0x69, 0x74, 0x79, 0x5c, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a,
	0x43, 0x69, 0x74, 0x79, 0x5c, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x69, 0x74,
	0x79, 0x3a, 0x3a, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_social_v1_social_service_proto_rawDescOnce sync.Once
	file_city_social_v1_social_service_proto_rawDescData = file_city_social_v1_social_service_proto_rawDesc
)

func file_city_social_v1_social_service_proto_rawDescGZIP() []byte {
	file_city_social_v1_social_service_proto_rawDescOnce.Do(func() {
		file_city_social_v1_social_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_social_v1_social_service_proto_rawDescData)
	})
	return file_city_social_v1_social_service_proto_rawDescData
}

var file_city_social_v1_social_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_city_social_v1_social_service_proto_goTypes = []interface{}{
	(*SendRequest)(nil),     // 0: city.social.v1.SendRequest
	(*SendResponse)(nil),    // 1: city.social.v1.SendResponse
	(*ReceiveRequest)(nil),  // 2: city.social.v1.ReceiveRequest
	(*ReceiveResponse)(nil), // 3: city.social.v1.ReceiveResponse
	(*Message)(nil),         // 4: city.social.v1.Message
}
var file_city_social_v1_social_service_proto_depIdxs = []int32{
	4, // 0: city.social.v1.SendRequest.messages:type_name -> city.social.v1.Message
	4, // 1: city.social.v1.ReceiveResponse.messages:type_name -> city.social.v1.Message
	0, // 2: city.social.v1.SocialService.Send:input_type -> city.social.v1.SendRequest
	2, // 3: city.social.v1.SocialService.Receive:input_type -> city.social.v1.ReceiveRequest
	1, // 4: city.social.v1.SocialService.Send:output_type -> city.social.v1.SendResponse
	3, // 5: city.social.v1.SocialService.Receive:output_type -> city.social.v1.ReceiveResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_city_social_v1_social_service_proto_init() }
func file_city_social_v1_social_service_proto_init() {
	if File_city_social_v1_social_service_proto != nil {
		return
	}
	file_city_social_v1_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_social_v1_social_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_city_social_v1_social_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
		file_city_social_v1_social_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveRequest); i {
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
		file_city_social_v1_social_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveResponse); i {
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
			RawDescriptor: file_city_social_v1_social_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_social_v1_social_service_proto_goTypes,
		DependencyIndexes: file_city_social_v1_social_service_proto_depIdxs,
		MessageInfos:      file_city_social_v1_social_service_proto_msgTypes,
	}.Build()
	File_city_social_v1_social_service_proto = out.File
	file_city_social_v1_social_service_proto_rawDesc = nil
	file_city_social_v1_social_service_proto_goTypes = nil
	file_city_social_v1_social_service_proto_depIdxs = nil
}
