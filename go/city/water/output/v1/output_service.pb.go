// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: city/water/output/v1/output_service.proto

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
	// 宏观道路水深
	Roads []*Road `protobuf:"bytes,2,rep,name=roads,proto3" json:"roads,omitempty"`
	// 微观道路点位水深
	DetailedRoads []*DetailedRoad `protobuf:"bytes,3,rep,name=detailed_roads,json=detailedRoads,proto3" json:"detailed_roads,omitempty"`
	// 排水节点
	DrainageNodes []*Node `protobuf:"bytes,4,rep,name=drainage_nodes,json=drainageNodes,proto3" json:"drainage_nodes,omitempty"`
	// 排水的边
	DrainageLinks []*Link `protobuf:"bytes,5,rep,name=drainage_links,json=drainageLinks,proto3" json:"drainage_links,omitempty"`
	// 供水节点
	SupplyNodes []*Node `protobuf:"bytes,6,rep,name=supply_nodes,json=supplyNodes,proto3" json:"supply_nodes,omitempty"`
	// 供水的边
	SupplyLinks []*Link `protobuf:"bytes,7,rep,name=supply_links,json=supplyLinks,proto3" json:"supply_links,omitempty"`
	// AOI粒度的供水指标
	Aois []*Aoi `protobuf:"bytes,8,rep,name=aois,proto3" json:"aois,omitempty"`
	// 排水负载指标
	DrainageMetric float64 `protobuf:"fixed64,9,opt,name=drainage_metric,json=drainageMetric,proto3" json:"drainage_metric,omitempty"`
	// 水网模拟的各种事件
	Events *v1.Events `protobuf:"bytes,10,opt,name=events,proto3" json:"events,omitempty"`
	// 排水网指标
	DrainageMetrics *DrainageMetrics `protobuf:"bytes,11,opt,name=drainage_metrics,json=drainageMetrics,proto3" json:"drainage_metrics,omitempty"`
	// 供水网指标
	SupplyMetrics *SupplyMetrics `protobuf:"bytes,12,opt,name=supply_metrics,json=supplyMetrics,proto3" json:"supply_metrics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutputRequest) Reset() {
	*x = OutputRequest{}
	mi := &file_city_water_output_v1_output_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputRequest) ProtoMessage() {}

func (x *OutputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_output_v1_output_service_proto_msgTypes[0]
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
	return file_city_water_output_v1_output_service_proto_rawDescGZIP(), []int{0}
}

func (x *OutputRequest) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *OutputRequest) GetRoads() []*Road {
	if x != nil {
		return x.Roads
	}
	return nil
}

func (x *OutputRequest) GetDetailedRoads() []*DetailedRoad {
	if x != nil {
		return x.DetailedRoads
	}
	return nil
}

func (x *OutputRequest) GetDrainageNodes() []*Node {
	if x != nil {
		return x.DrainageNodes
	}
	return nil
}

func (x *OutputRequest) GetDrainageLinks() []*Link {
	if x != nil {
		return x.DrainageLinks
	}
	return nil
}

func (x *OutputRequest) GetSupplyNodes() []*Node {
	if x != nil {
		return x.SupplyNodes
	}
	return nil
}

func (x *OutputRequest) GetSupplyLinks() []*Link {
	if x != nil {
		return x.SupplyLinks
	}
	return nil
}

func (x *OutputRequest) GetAois() []*Aoi {
	if x != nil {
		return x.Aois
	}
	return nil
}

func (x *OutputRequest) GetDrainageMetric() float64 {
	if x != nil {
		return x.DrainageMetric
	}
	return 0
}

func (x *OutputRequest) GetEvents() *v1.Events {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *OutputRequest) GetDrainageMetrics() *DrainageMetrics {
	if x != nil {
		return x.DrainageMetrics
	}
	return nil
}

func (x *OutputRequest) GetSupplyMetrics() *SupplyMetrics {
	if x != nil {
		return x.SupplyMetrics
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
	mi := &file_city_water_output_v1_output_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputResponse) ProtoMessage() {}

func (x *OutputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_water_output_v1_output_service_proto_msgTypes[1]
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
	return file_city_water_output_v1_output_service_proto_rawDescGZIP(), []int{1}
}

var File_city_water_output_v1_output_service_proto protoreflect.FileDescriptor

var file_city_water_output_v1_output_service_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2f, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x69,
	0x74, 0x79, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc9, 0x05, 0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x30, 0x0a, 0x05, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65,
	0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x61, 0x64,
	0x52, 0x05, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x49, 0x0a, 0x0e, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52,
	0x6f, 0x61, 0x64, 0x52, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x6f, 0x61,
	0x64, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0d, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x61, 0x67,
	0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0d, 0x64, 0x72, 0x61, 0x69, 0x6e,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x61, 0x6f, 0x69, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65,
	0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6f, 0x69, 0x52,
	0x04, 0x61, 0x6f, 0x69, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x61, 0x67,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e,
	0x64, 0x72, 0x61, 0x69, 0x6e, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x2d,
	0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x50, 0x0a,
	0x10, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x77,
	0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x72, 0x61, 0x69, 0x6e, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x0f,
	0x64, 0x72, 0x61, 0x69, 0x6e, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x4a, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x77,
	0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x0d, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x64, 0x0a,
	0x0d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x23, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0xe0, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x76, 0x31,
	0x42, 0x12, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c,
	0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x77, 0x61, 0x74,
	0x65, 0x72, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x57, 0x4f, 0xaa, 0x02, 0x14, 0x43, 0x69,
	0x74, 0x79, 0x2e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x14, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x57, 0x61, 0x74, 0x65, 0x72, 0x5c,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x20, 0x43, 0x69, 0x74, 0x79,
	0x5c, 0x57, 0x61, 0x74, 0x65, 0x72, 0x5c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x43,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x57, 0x61, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_city_water_output_v1_output_service_proto_rawDescOnce sync.Once
	file_city_water_output_v1_output_service_proto_rawDescData []byte
)

func file_city_water_output_v1_output_service_proto_rawDescGZIP() []byte {
	file_city_water_output_v1_output_service_proto_rawDescOnce.Do(func() {
		file_city_water_output_v1_output_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_water_output_v1_output_service_proto_rawDesc), len(file_city_water_output_v1_output_service_proto_rawDesc)))
	})
	return file_city_water_output_v1_output_service_proto_rawDescData
}

var file_city_water_output_v1_output_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_water_output_v1_output_service_proto_goTypes = []any{
	(*OutputRequest)(nil),   // 0: city.water.output.v1.OutputRequest
	(*OutputResponse)(nil),  // 1: city.water.output.v1.OutputResponse
	(*Road)(nil),            // 2: city.water.output.v1.Road
	(*DetailedRoad)(nil),    // 3: city.water.output.v1.DetailedRoad
	(*Node)(nil),            // 4: city.water.output.v1.Node
	(*Link)(nil),            // 5: city.water.output.v1.Link
	(*Aoi)(nil),             // 6: city.water.output.v1.Aoi
	(*v1.Events)(nil),       // 7: city.event.v1.Events
	(*DrainageMetrics)(nil), // 8: city.water.output.v1.DrainageMetrics
	(*SupplyMetrics)(nil),   // 9: city.water.output.v1.SupplyMetrics
}
var file_city_water_output_v1_output_service_proto_depIdxs = []int32{
	2,  // 0: city.water.output.v1.OutputRequest.roads:type_name -> city.water.output.v1.Road
	3,  // 1: city.water.output.v1.OutputRequest.detailed_roads:type_name -> city.water.output.v1.DetailedRoad
	4,  // 2: city.water.output.v1.OutputRequest.drainage_nodes:type_name -> city.water.output.v1.Node
	5,  // 3: city.water.output.v1.OutputRequest.drainage_links:type_name -> city.water.output.v1.Link
	4,  // 4: city.water.output.v1.OutputRequest.supply_nodes:type_name -> city.water.output.v1.Node
	5,  // 5: city.water.output.v1.OutputRequest.supply_links:type_name -> city.water.output.v1.Link
	6,  // 6: city.water.output.v1.OutputRequest.aois:type_name -> city.water.output.v1.Aoi
	7,  // 7: city.water.output.v1.OutputRequest.events:type_name -> city.event.v1.Events
	8,  // 8: city.water.output.v1.OutputRequest.drainage_metrics:type_name -> city.water.output.v1.DrainageMetrics
	9,  // 9: city.water.output.v1.OutputRequest.supply_metrics:type_name -> city.water.output.v1.SupplyMetrics
	0,  // 10: city.water.output.v1.OutputService.Output:input_type -> city.water.output.v1.OutputRequest
	1,  // 11: city.water.output.v1.OutputService.Output:output_type -> city.water.output.v1.OutputResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_city_water_output_v1_output_service_proto_init() }
func file_city_water_output_v1_output_service_proto_init() {
	if File_city_water_output_v1_output_service_proto != nil {
		return
	}
	file_city_water_output_v1_output_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_water_output_v1_output_service_proto_rawDesc), len(file_city_water_output_v1_output_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_water_output_v1_output_service_proto_goTypes,
		DependencyIndexes: file_city_water_output_v1_output_service_proto_depIdxs,
		MessageInfos:      file_city_water_output_v1_output_service_proto_msgTypes,
	}.Build()
	File_city_water_output_v1_output_service_proto = out.File
	file_city_water_output_v1_output_service_proto_goTypes = nil
	file_city_water_output_v1_output_service_proto_depIdxs = nil
}
