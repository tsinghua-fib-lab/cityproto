// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/comm/output/v1/output_service.proto

package outputv1

import (
	v1 "git.fiblab.net/sim/protos/v2/go/city/event/v1"
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

type OutputRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Step  int32                  `protobuf:"varint,1,opt,name=step,proto3" json:"step,omitempty"`
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
	Statistics    *Statistics `protobuf:"bytes,8,opt,name=statistics,proto3" json:"statistics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutputRequest) Reset() {
	*x = OutputRequest{}
	mi := &file_city_comm_output_v1_output_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputRequest) ProtoMessage() {}

func (x *OutputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_output_v1_output_service_proto_msgTypes[0]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutputResponse) Reset() {
	*x = OutputResponse{}
	mi := &file_city_comm_output_v1_output_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputResponse) ProtoMessage() {}

func (x *OutputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_comm_output_v1_output_service_proto_msgTypes[1]
	if x != nil {
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

const file_city_comm_output_v1_output_service_proto_rawDesc = "" +
	"\n" +
	"(city/comm/output/v1/output_service.proto\x12\x13city.comm.output.v1\x1a city/comm/output/v1/output.proto\x1a\x19city/event/v1/event.proto\"\x83\x04\n" +
	"\rOutputRequest\x12\x12\n" +
	"\x04step\x18\x01 \x01(\x05R\x04step\x12/\n" +
	"\x05nodes\x18\x02 \x03(\v2\x19.city.comm.output.v1.NodeR\x05nodes\x12E\n" +
	"\rbase_stations\x18\x03 \x03(\v2 .city.comm.output.v1.BaseStationR\fbaseStations\x12B\n" +
	"\x0esignal_heatmap\x18\x04 \x01(\v2\x1b.city.comm.output.v1.SignalR\rsignalHeatmap\x12M\n" +
	"\x14small_signal_heatmap\x18\t \x01(\v2\x1b.city.comm.output.v1.SignalR\x12smallSignalHeatmap\x125\n" +
	"\apersons\x18\x05 \x03(\v2\x1b.city.comm.output.v1.PersonR\apersons\x12,\n" +
	"\x04aois\x18\x06 \x03(\v2\x18.city.comm.output.v1.AoiR\x04aois\x12-\n" +
	"\x06events\x18\a \x01(\v2\x15.city.event.v1.EventsR\x06events\x12?\n" +
	"\n" +
	"statistics\x18\b \x01(\v2\x1f.city.comm.output.v1.StatisticsR\n" +
	"statistics\"\x10\n" +
	"\x0eOutputResponse2b\n" +
	"\rOutputService\x12Q\n" +
	"\x06Output\x12\".city.comm.output.v1.OutputRequest\x1a#.city.comm.output.v1.OutputResponseB\xda\x01\n" +
	"\x17com.city.comm.output.v1B\x12OutputServiceProtoP\x01Z<git.fiblab.net/sim/protos/v2/go/city/comm/output/v1;outputv1\xa2\x02\x03CCO\xaa\x02\x13City.Comm.Output.V1\xca\x02\x13City\\Comm\\Output\\V1\xe2\x02\x1fCity\\Comm\\Output\\V1\\GPBMetadata\xea\x02\x16City::Comm::Output::V1b\x06proto3"

var (
	file_city_comm_output_v1_output_service_proto_rawDescOnce sync.Once
	file_city_comm_output_v1_output_service_proto_rawDescData []byte
)

func file_city_comm_output_v1_output_service_proto_rawDescGZIP() []byte {
	file_city_comm_output_v1_output_service_proto_rawDescOnce.Do(func() {
		file_city_comm_output_v1_output_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_comm_output_v1_output_service_proto_rawDesc), len(file_city_comm_output_v1_output_service_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_comm_output_v1_output_service_proto_rawDesc), len(file_city_comm_output_v1_output_service_proto_rawDesc)),
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
	file_city_comm_output_v1_output_service_proto_goTypes = nil
	file_city_comm_output_v1_output_service_proto_depIdxs = nil
}
