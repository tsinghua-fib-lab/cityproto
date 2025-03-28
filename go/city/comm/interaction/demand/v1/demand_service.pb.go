// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/comm/interaction/demand/v1/demand_service.proto

package demandv1

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

// 设置用户通信需求激增
// 用户通信需求激增公式：
// result=demand*multiple_times*exp(-power_times*(current_time-start_time))
// demand: 用户正常通信需求
// current_time: 当前时间, start_time: 开始激增时间
type SetDemandStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MultipleTimes float64                `protobuf:"fixed64,1,opt,name=multiple_times,json=multipleTimes,proto3" json:"multiple_times,omitempty"`
	PowerTimes    float64                `protobuf:"fixed64,2,opt,name=power_times,json=powerTimes,proto3" json:"power_times,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetDemandStatusRequest) Reset() {
	*x = SetDemandStatusRequest{}
	mi := &file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetDemandStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDemandStatusRequest) ProtoMessage() {}

func (x *SetDemandStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[0]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetDemandStatusResponse) Reset() {
	*x = SetDemandStatusResponse{}
	mi := &file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetDemandStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDemandStatusResponse) ProtoMessage() {}

func (x *SetDemandStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_interaction_demand_v1_demand_service_proto_msgTypes[1]
	if x != nil {
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

const file_city_comm_interaction_demand_v1_demand_service_proto_rawDesc = "" +
	"\n" +
	"4city/comm/interaction/demand/v1/demand_service.proto\x12\x1fcity.comm.interaction.demand.v1\"`\n" +
	"\x16SetDemandStatusRequest\x12%\n" +
	"\x0emultiple_times\x18\x01 \x01(\x01R\rmultipleTimes\x12\x1f\n" +
	"\vpower_times\x18\x02 \x01(\x01R\n" +
	"powerTimes\"\x19\n" +
	"\x17SetDemandStatusResponse2\x96\x01\n" +
	"\rDemandService\x12\x84\x01\n" +
	"\x0fSetDemandStatus\x127.city.comm.interaction.demand.v1.SetDemandStatusRequest\x1a8.city.comm.interaction.demand.v1.SetDemandStatusResponseB\xa4\x02\n" +
	"#com.city.comm.interaction.demand.v1B\x12DemandServiceProtoP\x01ZHgit.fiblab.net/sim/protos/v2/go/city/comm/interaction/demand/v1;demandv1\xa2\x02\x04CCID\xaa\x02\x1fCity.Comm.Interaction.Demand.V1\xca\x02\x1fCity\\Comm\\Interaction\\Demand\\V1\xe2\x02+City\\Comm\\Interaction\\Demand\\V1\\GPBMetadata\xea\x02#City::Comm::Interaction::Demand::V1b\x06proto3"

var (
	file_city_comm_interaction_demand_v1_demand_service_proto_rawDescOnce sync.Once
	file_city_comm_interaction_demand_v1_demand_service_proto_rawDescData []byte
)

func file_city_comm_interaction_demand_v1_demand_service_proto_rawDescGZIP() []byte {
	file_city_comm_interaction_demand_v1_demand_service_proto_rawDescOnce.Do(func() {
		file_city_comm_interaction_demand_v1_demand_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_comm_interaction_demand_v1_demand_service_proto_rawDesc), len(file_city_comm_interaction_demand_v1_demand_service_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_comm_interaction_demand_v1_demand_service_proto_rawDesc), len(file_city_comm_interaction_demand_v1_demand_service_proto_rawDesc)),
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
	file_city_comm_interaction_demand_v1_demand_service_proto_goTypes = nil
	file_city_comm_interaction_demand_v1_demand_service_proto_depIdxs = nil
}
