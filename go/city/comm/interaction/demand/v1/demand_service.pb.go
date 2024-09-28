// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/comm/interaction/demand/v1/demand_service.proto

package demandv1

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

// 设置用户通信需求激增
// 用户通信需求激增公式：
// result=demand*multiple_times*exp(-power_times*(current_time-start_time))
// demand: 用户正常通信需求
// current_time: 当前时间, start_time: 开始激增时间
type SetDemandStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MultipleTimes float64 `protobuf:"fixed64,1,opt,name=multiple_times,json=multipleTimes,proto3" json:"multiple_times,omitempty"`
	PowerTimes    float64 `protobuf:"fixed64,2,opt,name=power_times,json=powerTimes,proto3" json:"power_times,omitempty"`
}

func (x *SetDemandStatusRequest) Reset() {
	*x = SetDemandStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDemandStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDemandStatusRequest) ProtoMessage() {}

func (x *SetDemandStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDemandStatusRequest.ProtoReflect.Descriptor instead.
func (*SetDemandStatusRequest) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_demand_v1_demand_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetDemandStatusRequest) GetMultipleTimes() float64 {
	if x != nil {
		return x.MultipleTimes
	}
	return 0
}

func (x *SetDemandStatusRequest) GetPowerTimes() float64 {
	if x != nil {
		return x.PowerTimes
	}
	return 0
}

type SetDemandStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetDemandStatusResponse) Reset() {
	*x = SetDemandStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDemandStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDemandStatusResponse) ProtoMessage() {}

func (x *SetDemandStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDemandStatusResponse.ProtoReflect.Descriptor instead.
func (*SetDemandStatusResponse) Descriptor() ([]byte, []int) {
	return file_city_comm_interaction_demand_v1_demand_service_proto_rawDescGZIP(), []int{1}
}

var File_city_comm_interaction_demand_v1_demand_service_proto protoreflect.FileDescriptor

var file_city_comm_interaction_demand_v1_demand_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x60, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x44, 0x65,
	0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x77, 0x65,
	0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x96, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x44, 0x65,
	0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa4, 0x02,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74,
	0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74,
	0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x43, 0x43, 0x49, 0x44, 0xaa, 0x02, 0x1f, 0x43,
	0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1f, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x2b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x23, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x3a, 0x3a, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_comm_interaction_demand_v1_demand_service_proto_rawDescOnce sync.Once
	file_city_comm_interaction_demand_v1_demand_service_proto_rawDescData = file_city_comm_interaction_demand_v1_demand_service_proto_rawDesc
)

func file_city_comm_interaction_demand_v1_demand_service_proto_rawDescGZIP() []byte {
	file_city_comm_interaction_demand_v1_demand_service_proto_rawDescOnce.Do(func() {
		file_city_comm_interaction_demand_v1_demand_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_comm_interaction_demand_v1_demand_service_proto_rawDescData)
	})
	return file_city_comm_interaction_demand_v1_demand_service_proto_rawDescData
}

var file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_comm_interaction_demand_v1_demand_service_proto_goTypes = []any{
	(*SetDemandStatusRequest)(nil),  // 0: city.comm.interaction.demand.v1.SetDemandStatusRequest
	(*SetDemandStatusResponse)(nil), // 1: city.comm.interaction.demand.v1.SetDemandStatusResponse
}
var file_city_comm_interaction_demand_v1_demand_service_proto_depIdxs = []int32{
	0, // 0: city.comm.interaction.demand.v1.DemandService.SetDemandStatus:input_type -> city.comm.interaction.demand.v1.SetDemandStatusRequest
	1, // 1: city.comm.interaction.demand.v1.DemandService.SetDemandStatus:output_type -> city.comm.interaction.demand.v1.SetDemandStatusResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_comm_interaction_demand_v1_demand_service_proto_init() }
func file_city_comm_interaction_demand_v1_demand_service_proto_init() {
	if File_city_comm_interaction_demand_v1_demand_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetDemandStatusRequest); i {
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
		file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetDemandStatusResponse); i {
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
			RawDescriptor: file_city_comm_interaction_demand_v1_demand_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_comm_interaction_demand_v1_demand_service_proto_goTypes,
		DependencyIndexes: file_city_comm_interaction_demand_v1_demand_service_proto_depIdxs,
		MessageInfos:      file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes,
	}.Build()
	File_city_comm_interaction_demand_v1_demand_service_proto = out.File
	file_city_comm_interaction_demand_v1_demand_service_proto_rawDesc = nil
	file_city_comm_interaction_demand_v1_demand_service_proto_goTypes = nil
	file_city_comm_interaction_demand_v1_demand_service_proto_depIdxs = nil
}
