// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: city/routing/v2/routing_service.proto

package routingv2

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

// 获取导航路线请求
// Request for getting routing path
type GetRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 导航类型
	// routing type
	Type RouteType `protobuf:"varint,1,opt,name=type,proto3,enum=city.routing.v2.RouteType" json:"type,omitempty" yaml:"type" bson:"type" db:"type"`
	// 起点，约定：包含LanePosition或AoiPosition中的一种
	// Starting point, convention: as LanePosition or AoiPosition
	Start *v2.Position `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty" bson:"start" db:"start" yaml:"start"`
	// 终点，约定：包含LanePosition或AoiPosition中的一种
	// Ending point, convention: as LanePosition or AoiPosition
	End *v2.Position `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty" yaml:"end" bson:"end" db:"end"`
	// 发送导航请求的时间（目前仅在行车导航中使用）
	// The time to send routing request (currently only used in driving routing)
	Time float64 `protobuf:"fixed64,5,opt,name=time,proto3" json:"time,omitempty" yaml:"time" bson:"time" db:"time"`
}

func (x *GetRouteRequest) Reset() {
	*x = GetRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_routing_v2_routing_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteRequest) ProtoMessage() {}

func (x *GetRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteRequest.ProtoReflect.Descriptor instead.
func (*GetRouteRequest) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRouteRequest) GetType() RouteType {
	if x != nil {
		return x.Type
	}
	return RouteType_ROUTE_TYPE_UNSPECIFIED
}

func (x *GetRouteRequest) GetStart() *v2.Position {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GetRouteRequest) GetEnd() *v2.Position {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *GetRouteRequest) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

// 获取导航路线响应
// Response of getting routing path
type GetRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Journeys []*Journey `protobuf:"bytes,1,rep,name=journeys,proto3" json:"journeys,omitempty" yaml:"journeys" bson:"journeys" db:"journeys"`
}

func (x *GetRouteResponse) Reset() {
	*x = GetRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_routing_v2_routing_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteResponse) ProtoMessage() {}

func (x *GetRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteResponse.ProtoReflect.Descriptor instead.
func (*GetRouteResponse) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRouteResponse) GetJourneys() []*Journey {
	if x != nil {
		return x.Journeys
	}
	return nil
}

// 设置行车导航道路通行成本请求
// Request for setting driving routing travelling cost
type SetDrivingCostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 道路通行成本
	// travelling cost
	Costs []*Cost `protobuf:"bytes,1,rep,name=costs,proto3" json:"costs,omitempty" yaml:"costs" bson:"costs" db:"costs"`
}

func (x *SetDrivingCostsRequest) Reset() {
	*x = SetDrivingCostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_routing_v2_routing_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDrivingCostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDrivingCostsRequest) ProtoMessage() {}

func (x *SetDrivingCostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDrivingCostsRequest.ProtoReflect.Descriptor instead.
func (*SetDrivingCostsRequest) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetDrivingCostsRequest) GetCosts() []*Cost {
	if x != nil {
		return x.Costs
	}
	return nil
}

// 设置行车导航道路通行成本响应
// Response of setting driving routing travelling cost
type SetDrivingCostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetDrivingCostsResponse) Reset() {
	*x = SetDrivingCostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_routing_v2_routing_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDrivingCostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDrivingCostsResponse) ProtoMessage() {}

func (x *SetDrivingCostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDrivingCostsResponse.ProtoReflect.Descriptor instead.
func (*SetDrivingCostsResponse) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_service_proto_rawDescGZIP(), []int{3}
}

// 获取行车导航道路通行成本请求
// Request for getting driving routing travelling cost
type GetDrivingCostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 道路通行成本（按照给定的id和time进行查询）
	// travelling cost (query via the given ID and time)
	Costs []*Cost `protobuf:"bytes,1,rep,name=costs,proto3" json:"costs,omitempty" yaml:"costs" bson:"costs" db:"costs"`
}

func (x *GetDrivingCostsRequest) Reset() {
	*x = GetDrivingCostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_routing_v2_routing_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrivingCostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrivingCostsRequest) ProtoMessage() {}

func (x *GetDrivingCostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrivingCostsRequest.ProtoReflect.Descriptor instead.
func (*GetDrivingCostsRequest) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetDrivingCostsRequest) GetCosts() []*Cost {
	if x != nil {
		return x.Costs
	}
	return nil
}

// 获取行车导航道路通行成本响应
// Response of getting driving routing travelling cost
type GetDrivingCostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 道路通行成本（补全cost后的结果）
	// travelling cost (results after completing the cost)
	Costs []*Cost `protobuf:"bytes,1,rep,name=costs,proto3" json:"costs,omitempty" db:"costs" yaml:"costs" bson:"costs"`
}

func (x *GetDrivingCostsResponse) Reset() {
	*x = GetDrivingCostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_routing_v2_routing_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrivingCostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrivingCostsResponse) ProtoMessage() {}

func (x *GetDrivingCostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_routing_v2_routing_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrivingCostsResponse.ProtoReflect.Descriptor instead.
func (*GetDrivingCostsResponse) Descriptor() ([]byte, []int) {
	return file_city_routing_v2_routing_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetDrivingCostsResponse) GetCosts() []*Cost {
	if x != nil {
		return x.Costs
	}
	return nil
}

var File_city_routing_v2_routing_service_proto protoreflect.FileDescriptor

var file_city_routing_v2_routing_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x15, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67,
	0x65, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x69, 0x74,
	0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08,
	0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32,
	0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x52, 0x08, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x73, 0x22, 0x45, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05,
	0x63, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x65, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x46, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x63, 0x6f,
	0x73, 0x74, 0x73, 0x32, 0xad, 0x02, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x20, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x44, 0x72,
	0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x73,
	0x12, 0x27, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0xc0, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x13, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e,
	0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32,
	0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x52, 0x58,
	0xaa, 0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x32, 0xca, 0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_routing_v2_routing_service_proto_rawDescOnce sync.Once
	file_city_routing_v2_routing_service_proto_rawDescData = file_city_routing_v2_routing_service_proto_rawDesc
)

func file_city_routing_v2_routing_service_proto_rawDescGZIP() []byte {
	file_city_routing_v2_routing_service_proto_rawDescOnce.Do(func() {
		file_city_routing_v2_routing_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_routing_v2_routing_service_proto_rawDescData)
	})
	return file_city_routing_v2_routing_service_proto_rawDescData
}

var file_city_routing_v2_routing_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_city_routing_v2_routing_service_proto_goTypes = []interface{}{
	(*GetRouteRequest)(nil),         // 0: city.routing.v2.GetRouteRequest
	(*GetRouteResponse)(nil),        // 1: city.routing.v2.GetRouteResponse
	(*SetDrivingCostsRequest)(nil),  // 2: city.routing.v2.SetDrivingCostsRequest
	(*SetDrivingCostsResponse)(nil), // 3: city.routing.v2.SetDrivingCostsResponse
	(*GetDrivingCostsRequest)(nil),  // 4: city.routing.v2.GetDrivingCostsRequest
	(*GetDrivingCostsResponse)(nil), // 5: city.routing.v2.GetDrivingCostsResponse
	(RouteType)(0),                  // 6: city.routing.v2.RouteType
	(*v2.Position)(nil),             // 7: city.geo.v2.Position
	(*Journey)(nil),                 // 8: city.routing.v2.Journey
	(*Cost)(nil),                    // 9: city.routing.v2.Cost
}
var file_city_routing_v2_routing_service_proto_depIdxs = []int32{
	6,  // 0: city.routing.v2.GetRouteRequest.type:type_name -> city.routing.v2.RouteType
	7,  // 1: city.routing.v2.GetRouteRequest.start:type_name -> city.geo.v2.Position
	7,  // 2: city.routing.v2.GetRouteRequest.end:type_name -> city.geo.v2.Position
	8,  // 3: city.routing.v2.GetRouteResponse.journeys:type_name -> city.routing.v2.Journey
	9,  // 4: city.routing.v2.SetDrivingCostsRequest.costs:type_name -> city.routing.v2.Cost
	9,  // 5: city.routing.v2.GetDrivingCostsRequest.costs:type_name -> city.routing.v2.Cost
	9,  // 6: city.routing.v2.GetDrivingCostsResponse.costs:type_name -> city.routing.v2.Cost
	0,  // 7: city.routing.v2.RoutingService.GetRoute:input_type -> city.routing.v2.GetRouteRequest
	2,  // 8: city.routing.v2.RoutingService.SetDrivingCosts:input_type -> city.routing.v2.SetDrivingCostsRequest
	4,  // 9: city.routing.v2.RoutingService.GetDrivingCosts:input_type -> city.routing.v2.GetDrivingCostsRequest
	1,  // 10: city.routing.v2.RoutingService.GetRoute:output_type -> city.routing.v2.GetRouteResponse
	3,  // 11: city.routing.v2.RoutingService.SetDrivingCosts:output_type -> city.routing.v2.SetDrivingCostsResponse
	5,  // 12: city.routing.v2.RoutingService.GetDrivingCosts:output_type -> city.routing.v2.GetDrivingCostsResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_city_routing_v2_routing_service_proto_init() }
func file_city_routing_v2_routing_service_proto_init() {
	if File_city_routing_v2_routing_service_proto != nil {
		return
	}
	file_city_routing_v2_cost_proto_init()
	file_city_routing_v2_routing_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_routing_v2_routing_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRouteRequest); i {
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
		file_city_routing_v2_routing_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRouteResponse); i {
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
		file_city_routing_v2_routing_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDrivingCostsRequest); i {
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
		file_city_routing_v2_routing_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDrivingCostsResponse); i {
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
		file_city_routing_v2_routing_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrivingCostsRequest); i {
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
		file_city_routing_v2_routing_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrivingCostsResponse); i {
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
			RawDescriptor: file_city_routing_v2_routing_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_routing_v2_routing_service_proto_goTypes,
		DependencyIndexes: file_city_routing_v2_routing_service_proto_depIdxs,
		MessageInfos:      file_city_routing_v2_routing_service_proto_msgTypes,
	}.Build()
	File_city_routing_v2_routing_service_proto = out.File
	file_city_routing_v2_routing_service_proto_rawDesc = nil
	file_city_routing_v2_routing_service_proto_goTypes = nil
	file_city_routing_v2_routing_service_proto_depIdxs = nil
}
