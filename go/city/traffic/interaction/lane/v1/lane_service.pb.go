// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: city/traffic/interaction/lane/v1/lane_service.proto

package lanev1

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

type SetMaxVRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaneId int32   `protobuf:"varint,1,opt,name=lane_id,json=laneId,proto3" json:"lane_id,omitempty"`
	MaxV   float64 `protobuf:"fixed64,2,opt,name=max_v,json=maxV,proto3" json:"max_v,omitempty"`
}

func (x *SetMaxVRequest) Reset() {
	*x = SetMaxVRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMaxVRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMaxVRequest) ProtoMessage() {}

func (x *SetMaxVRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMaxVRequest.ProtoReflect.Descriptor instead.
func (*SetMaxVRequest) Descriptor() ([]byte, []int) {
	return file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetMaxVRequest) GetLaneId() int32 {
	if x != nil {
		return x.LaneId
	}
	return 0
}

func (x *SetMaxVRequest) GetMaxV() float64 {
	if x != nil {
		return x.MaxV
	}
	return 0
}

type SetMaxVResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetMaxVResponse) Reset() {
	*x = SetMaxVResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMaxVResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMaxVResponse) ProtoMessage() {}

func (x *SetMaxVResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMaxVResponse.ProtoReflect.Descriptor instead.
func (*SetMaxVResponse) Descriptor() ([]byte, []int) {
	return file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescGZIP(), []int{1}
}

type GetLaneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaneIds []int32 `protobuf:"varint,1,rep,packed,name=lane_ids,json=laneIds,proto3" json:"lane_ids,omitempty"`
}

func (x *GetLaneRequest) Reset() {
	*x = GetLaneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLaneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLaneRequest) ProtoMessage() {}

func (x *GetLaneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLaneRequest.ProtoReflect.Descriptor instead.
func (*GetLaneRequest) Descriptor() ([]byte, []int) {
	return file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLaneRequest) GetLaneIds() []int32 {
	if x != nil {
		return x.LaneIds
	}
	return nil
}

type GetLaneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States []*State `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *GetLaneResponse) Reset() {
	*x = GetLaneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLaneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLaneResponse) ProtoMessage() {}

func (x *GetLaneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLaneResponse.ProtoReflect.Descriptor instead.
func (*GetLaneResponse) Descriptor() ([]byte, []int) {
	return file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetLaneResponse) GetStates() []*State {
	if x != nil {
		return x.States
	}
	return nil
}

var File_city_traffic_interaction_lane_v1_lane_service_proto protoreflect.FileDescriptor

var file_city_traffic_interaction_lane_v1_lane_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2b, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x56, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12,
	0x13, 0x0a, 0x05, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6d, 0x61, 0x78, 0x56, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x56, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x6e,
	0x65, 0x49, 0x64, 0x73, 0x22, 0x52, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x32, 0xed, 0x01, 0x0a, 0x0b, 0x4c, 0x61, 0x6e,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x4d,
	0x61, 0x78, 0x56, 0x12, 0x30, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6c,
	0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x56, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x56,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x6e, 0x65, 0x12, 0x30, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6c,
	0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa3, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x10, 0x4c, 0x61, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61,
	0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x61, 0x6e, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x43, 0x54,
	0x49, 0x4c, 0xaa, 0x02, 0x20, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61,
	0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5c, 0x4c, 0x61, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2c, 0x43, 0x69, 0x74, 0x79, 0x5c,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5c, 0x4c, 0x61, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x4c, 0x61, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescOnce sync.Once
	file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescData = file_city_traffic_interaction_lane_v1_lane_service_proto_rawDesc
)

func file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescGZIP() []byte {
	file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescOnce.Do(func() {
		file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescData)
	})
	return file_city_traffic_interaction_lane_v1_lane_service_proto_rawDescData
}

var file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_city_traffic_interaction_lane_v1_lane_service_proto_goTypes = []interface{}{
	(*SetMaxVRequest)(nil),  // 0: city.traffic.interaction.lane.v1.SetMaxVRequest
	(*SetMaxVResponse)(nil), // 1: city.traffic.interaction.lane.v1.SetMaxVResponse
	(*GetLaneRequest)(nil),  // 2: city.traffic.interaction.lane.v1.GetLaneRequest
	(*GetLaneResponse)(nil), // 3: city.traffic.interaction.lane.v1.GetLaneResponse
	(*State)(nil),           // 4: city.traffic.interaction.lane.v1.State
}
var file_city_traffic_interaction_lane_v1_lane_service_proto_depIdxs = []int32{
	4, // 0: city.traffic.interaction.lane.v1.GetLaneResponse.states:type_name -> city.traffic.interaction.lane.v1.State
	0, // 1: city.traffic.interaction.lane.v1.LaneService.SetMaxV:input_type -> city.traffic.interaction.lane.v1.SetMaxVRequest
	2, // 2: city.traffic.interaction.lane.v1.LaneService.GetLane:input_type -> city.traffic.interaction.lane.v1.GetLaneRequest
	1, // 3: city.traffic.interaction.lane.v1.LaneService.SetMaxV:output_type -> city.traffic.interaction.lane.v1.SetMaxVResponse
	3, // 4: city.traffic.interaction.lane.v1.LaneService.GetLane:output_type -> city.traffic.interaction.lane.v1.GetLaneResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_city_traffic_interaction_lane_v1_lane_service_proto_init() }
func file_city_traffic_interaction_lane_v1_lane_service_proto_init() {
	if File_city_traffic_interaction_lane_v1_lane_service_proto != nil {
		return
	}
	file_city_traffic_interaction_lane_v1_lane_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMaxVRequest); i {
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
		file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMaxVResponse); i {
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
		file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLaneRequest); i {
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
		file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLaneResponse); i {
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
			RawDescriptor: file_city_traffic_interaction_lane_v1_lane_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_traffic_interaction_lane_v1_lane_service_proto_goTypes,
		DependencyIndexes: file_city_traffic_interaction_lane_v1_lane_service_proto_depIdxs,
		MessageInfos:      file_city_traffic_interaction_lane_v1_lane_service_proto_msgTypes,
	}.Build()
	File_city_traffic_interaction_lane_v1_lane_service_proto = out.File
	file_city_traffic_interaction_lane_v1_lane_service_proto_rawDesc = nil
	file_city_traffic_interaction_lane_v1_lane_service_proto_goTypes = nil
	file_city_traffic_interaction_lane_v1_lane_service_proto_depIdxs = nil
}
