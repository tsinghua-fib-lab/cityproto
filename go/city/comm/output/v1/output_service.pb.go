// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/comm/output/v1/output_service.proto

package outputv1

import (
	v1 "git.fiblab.net/sim/protos/go/v2/city/event/v1"
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

type OutputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step int32 `protobuf:"varint,1,opt,name=step,proto3" json:"step,omitempty"`
	// 设备状态
	Nodes        []*Node        `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
	BaseStations []*BaseStation `protobuf:"bytes,3,rep,name=base_stations,json=baseStations,proto3" json:"base_stations,omitempty"`
	// 五环区域信号强度热力图（500m网格形式呈现）
	SignalHeatmap *Signal `protobuf:"bytes,4,opt,name=signal_heatmap,json=signalHeatmap,proto3" json:"signal_heatmap,omitempty"`
	// 国贸区域信号强度热力图（50m网格形式呈现）
	SmallSignalHeatmap *Signal `protobuf:"bytes,9,opt,name=small_signal_heatmap,json=smallSignalHeatmap,proto3" json:"small_signal_heatmap,omitempty"`
	// TODO(张钧): 基站和人的连接怎么表示？
	// 人相关的数据
	Persons []*Person `protobuf:"bytes,5,rep,name=persons,proto3" json:"persons,omitempty"`
	// AOI相关的数据
	Aois   []*Aoi     `protobuf:"bytes,6,rep,name=aois,proto3" json:"aois,omitempty"`
	Events *v1.Events `protobuf:"bytes,7,opt,name=events,proto3" json:"events,omitempty"`
	// 统计值
	Statistics *Statistics `protobuf:"bytes,8,opt,name=statistics,proto3" json:"statistics,omitempty"`
}

func (x *OutputRequest) Reset() {
	*x = OutputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_output_v1_output_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputRequest) ProtoMessage() {}

func (x *OutputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_output_v1_output_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputRequest.ProtoReflect.Descriptor instead.
func (*OutputRequest) Descriptor() ([]byte, []int) {
	return file_city_comm_output_v1_output_service_proto_rawDescGZIP(), []int{0}
}

func (x *OutputRequest) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *OutputRequest) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *OutputRequest) GetBaseStations() []*BaseStation {
	if x != nil {
		return x.BaseStations
	}
	return nil
}

func (x *OutputRequest) GetSignalHeatmap() *Signal {
	if x != nil {
		return x.SignalHeatmap
	}
	return nil
}

func (x *OutputRequest) GetSmallSignalHeatmap() *Signal {
	if x != nil {
		return x.SmallSignalHeatmap
	}
	return nil
}

func (x *OutputRequest) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

func (x *OutputRequest) GetAois() []*Aoi {
	if x != nil {
		return x.Aois
	}
	return nil
}

func (x *OutputRequest) GetEvents() *v1.Events {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *OutputRequest) GetStatistics() *Statistics {
	if x != nil {
		return x.Statistics
	}
	return nil
}

type OutputResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OutputResponse) Reset() {
	*x = OutputResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_comm_output_v1_output_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputResponse) ProtoMessage() {}

func (x *OutputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_output_v1_output_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputResponse.ProtoReflect.Descriptor instead.
func (*OutputResponse) Descriptor() ([]byte, []int) {
	return file_city_comm_output_v1_output_service_proto_rawDescGZIP(), []int{1}
}

var File_city_comm_output_v1_output_service_proto protoreflect.FileDescriptor

var file_city_comm_output_v1_output_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x20, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x04, 0x0a,
	0x0d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x12, 0x2f, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x62, 0x61,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x68, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52,
	0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x12, 0x4d,
	0x0a, 0x14, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x68,
	0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x12, 0x73, 0x6d, 0x61, 0x6c, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x12, 0x35, 0x0a,
	0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x61, 0x6f, 0x69, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6f, 0x69, 0x52, 0x04, 0x61, 0x6f,
	0x69, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x62, 0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x22, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xda, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x2e,
	0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x69, 0x74, 0x79,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x4f, 0xaa, 0x02,
	0x13, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x5c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x69, 0x74,
	0x79, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x5c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x43,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x3a, 0x3a, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_comm_output_v1_output_service_proto_rawDescOnce sync.Once
	file_city_comm_output_v1_output_service_proto_rawDescData = file_city_comm_output_v1_output_service_proto_rawDesc
)

func file_city_comm_output_v1_output_service_proto_rawDescGZIP() []byte {
	file_city_comm_output_v1_output_service_proto_rawDescOnce.Do(func() {
		file_city_comm_output_v1_output_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_comm_output_v1_output_service_proto_rawDescData)
	})
	return file_city_comm_output_v1_output_service_proto_rawDescData
}

var file_city_comm_output_v1_output_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_comm_output_v1_output_service_proto_goTypes = []any{
	(*OutputRequest)(nil),  // 0: city.comm.output.v1.OutputRequest
	(*OutputResponse)(nil), // 1: city.comm.output.v1.OutputResponse
	(*Node)(nil),           // 2: city.comm.output.v1.Node
	(*BaseStation)(nil),    // 3: city.comm.output.v1.BaseStation
	(*Signal)(nil),         // 4: city.comm.output.v1.Signal
	(*Person)(nil),         // 5: city.comm.output.v1.Person
	(*Aoi)(nil),            // 6: city.comm.output.v1.Aoi
	(*v1.Events)(nil),      // 7: city.event.v1.Events
	(*Statistics)(nil),     // 8: city.comm.output.v1.Statistics
}
var file_city_comm_output_v1_output_service_proto_depIdxs = []int32{
	2, // 0: city.comm.output.v1.OutputRequest.nodes:type_name -> city.comm.output.v1.Node
	3, // 1: city.comm.output.v1.OutputRequest.base_stations:type_name -> city.comm.output.v1.BaseStation
	4, // 2: city.comm.output.v1.OutputRequest.signal_heatmap:type_name -> city.comm.output.v1.Signal
	4, // 3: city.comm.output.v1.OutputRequest.small_signal_heatmap:type_name -> city.comm.output.v1.Signal
	5, // 4: city.comm.output.v1.OutputRequest.persons:type_name -> city.comm.output.v1.Person
	6, // 5: city.comm.output.v1.OutputRequest.aois:type_name -> city.comm.output.v1.Aoi
	7, // 6: city.comm.output.v1.OutputRequest.events:type_name -> city.event.v1.Events
	8, // 7: city.comm.output.v1.OutputRequest.statistics:type_name -> city.comm.output.v1.Statistics
	0, // 8: city.comm.output.v1.OutputService.Output:input_type -> city.comm.output.v1.OutputRequest
	1, // 9: city.comm.output.v1.OutputService.Output:output_type -> city.comm.output.v1.OutputResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_city_comm_output_v1_output_service_proto_init() }
func file_city_comm_output_v1_output_service_proto_init() {
	if File_city_comm_output_v1_output_service_proto != nil {
		return
	}
	file_city_comm_output_v1_output_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_comm_output_v1_output_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OutputRequest); i {
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
		file_city_comm_output_v1_output_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OutputResponse); i {
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
			RawDescriptor: file_city_comm_output_v1_output_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_comm_output_v1_output_service_proto_goTypes,
		DependencyIndexes: file_city_comm_output_v1_output_service_proto_depIdxs,
		MessageInfos:      file_city_comm_output_v1_output_service_proto_msgTypes,
	}.Build()
	File_city_comm_output_v1_output_service_proto = out.File
	file_city_comm_output_v1_output_service_proto_rawDesc = nil
	file_city_comm_output_v1_output_service_proto_goTypes = nil
	file_city_comm_output_v1_output_service_proto_depIdxs = nil
}
