// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/event/v1/event.proto

package eventv1

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

type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED EventType = 0
	// 特大暴雨
	EventType_EVENT_TYPE_HEAVY_RAIN EventType = 1
	// 军事打击
	EventType_EVENT_TYPE_MILITARY_STRIKE EventType = 2
	// 交通拥堵
	EventType_EVENT_TYPE_TRAFFIC_CONGESTION EventType = 3
	// 道路限行
	EventType_EVENT_TYPE_TRAFFIC_LANE_RESTRICTION EventType = 4
	// 信控失效
	EventType_EVENT_TYPE_TRAFFIC_BAD_LIGHT EventType = 5
	// 变压器损坏（被摧毁）
	EventType_EVENT_TYPE_ELEC_RUINED_TRANSFORMER EventType = 6
	// 水泵损坏（被摧毁）
	EventType_EVENT_TYPE_WATER_RUINED_PUMP EventType = 7
	// 水泵停电（变压器停电影响）
	EventType_EVENT_TYPE_WATER_STOPPED_PUMP EventType = 8
	// 基站损坏
	EventType_EVENT_TYPE_COMM_RUINED_BASE_STATION EventType = 9
	// 基站停电
	EventType_EVENT_TYPE_COMM_STOPPED_BASE_STATION EventType = 10
	// 基站过载
	EventType_EVENT_TYPE_COMM_OVERLOAD_BASE_STATION EventType = 11
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0:  "EVENT_TYPE_UNSPECIFIED",
		1:  "EVENT_TYPE_HEAVY_RAIN",
		2:  "EVENT_TYPE_MILITARY_STRIKE",
		3:  "EVENT_TYPE_TRAFFIC_CONGESTION",
		4:  "EVENT_TYPE_TRAFFIC_LANE_RESTRICTION",
		5:  "EVENT_TYPE_TRAFFIC_BAD_LIGHT",
		6:  "EVENT_TYPE_ELEC_RUINED_TRANSFORMER",
		7:  "EVENT_TYPE_WATER_RUINED_PUMP",
		8:  "EVENT_TYPE_WATER_STOPPED_PUMP",
		9:  "EVENT_TYPE_COMM_RUINED_BASE_STATION",
		10: "EVENT_TYPE_COMM_STOPPED_BASE_STATION",
		11: "EVENT_TYPE_COMM_OVERLOAD_BASE_STATION",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED":                0,
		"EVENT_TYPE_HEAVY_RAIN":                 1,
		"EVENT_TYPE_MILITARY_STRIKE":            2,
		"EVENT_TYPE_TRAFFIC_CONGESTION":         3,
		"EVENT_TYPE_TRAFFIC_LANE_RESTRICTION":   4,
		"EVENT_TYPE_TRAFFIC_BAD_LIGHT":          5,
		"EVENT_TYPE_ELEC_RUINED_TRANSFORMER":    6,
		"EVENT_TYPE_WATER_RUINED_PUMP":          7,
		"EVENT_TYPE_WATER_STOPPED_PUMP":         8,
		"EVENT_TYPE_COMM_RUINED_BASE_STATION":   9,
		"EVENT_TYPE_COMM_STOPPED_BASE_STATION":  10,
		"EVENT_TYPE_COMM_OVERLOAD_BASE_STATION": 11,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_city_event_v1_event_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_city_event_v1_event_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_city_event_v1_event_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  EventType `protobuf:"varint,1,opt,name=type,proto3,enum=city.event.v1.EventType" json:"type,omitempty" bson:"type" db:"type" yaml:"type"`
	Level int32     `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty" bson:"level" db:"level" yaml:"level"`
	Step  int32     `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty" bson:"step" db:"step" yaml:"step"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_event_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_city_event_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_EVENT_TYPE_UNSPECIFIED
}

func (x *Event) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Event) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

type Events struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty" bson:"events" db:"events" yaml:"events"`
}

func (x *Events) Reset() {
	*x = Events{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_event_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Events) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events) ProtoMessage() {}

func (x *Events) ProtoReflect() protoreflect.Message {
	mi := &file_city_event_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events.ProtoReflect.Descriptor instead.
func (*Events) Descriptor() ([]byte, []int) {
	return file_city_event_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *Events) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_city_event_v1_event_proto protoreflect.FileDescriptor

var file_city_event_v1_event_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x5f, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x36, 0x0a, 0x06, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2a, 0xbb, 0x03, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x45, 0x41, 0x56,
	0x59, 0x5f, 0x52, 0x41, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49, 0x4c, 0x49, 0x54, 0x41, 0x52, 0x59, 0x5f,
	0x53, 0x54, 0x52, 0x49, 0x4b, 0x45, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x43,
	0x4f, 0x4e, 0x47, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49,
	0x43, 0x5f, 0x4c, 0x41, 0x4e, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x42, 0x41, 0x44, 0x5f, 0x4c,
	0x49, 0x47, 0x48, 0x54, 0x10, 0x05, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4c, 0x45, 0x43, 0x5f, 0x52, 0x55, 0x49, 0x4e, 0x45, 0x44,
	0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x52, 0x10, 0x06, 0x12, 0x20,
	0x0a, 0x1c, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x41, 0x54,
	0x45, 0x52, 0x5f, 0x52, 0x55, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x50, 0x55, 0x4d, 0x50, 0x10, 0x07,
	0x12, 0x21, 0x0a, 0x1d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57,
	0x41, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x5f, 0x50, 0x55, 0x4d,
	0x50, 0x10, 0x08, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x5f, 0x52, 0x55, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x28, 0x0a, 0x24,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x5f,
	0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x4f,
	0x41, 0x44, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x0b, 0x42, 0xac, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61,
	0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x45, 0x58, 0xaa, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x19, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0f, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_event_v1_event_proto_rawDescOnce sync.Once
	file_city_event_v1_event_proto_rawDescData = file_city_event_v1_event_proto_rawDesc
)

func file_city_event_v1_event_proto_rawDescGZIP() []byte {
	file_city_event_v1_event_proto_rawDescOnce.Do(func() {
		file_city_event_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_event_v1_event_proto_rawDescData)
	})
	return file_city_event_v1_event_proto_rawDescData
}

var file_city_event_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_city_event_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_event_v1_event_proto_goTypes = []any{
	(EventType)(0), // 0: city.event.v1.EventType
	(*Event)(nil),  // 1: city.event.v1.Event
	(*Events)(nil), // 2: city.event.v1.Events
}
var file_city_event_v1_event_proto_depIdxs = []int32{
	0, // 0: city.event.v1.Event.type:type_name -> city.event.v1.EventType
	1, // 1: city.event.v1.Events.events:type_name -> city.event.v1.Event
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_city_event_v1_event_proto_init() }
func file_city_event_v1_event_proto_init() {
	if File_city_event_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_event_v1_event_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
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
		file_city_event_v1_event_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Events); i {
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
			RawDescriptor: file_city_event_v1_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_event_v1_event_proto_goTypes,
		DependencyIndexes: file_city_event_v1_event_proto_depIdxs,
		EnumInfos:         file_city_event_v1_event_proto_enumTypes,
		MessageInfos:      file_city_event_v1_event_proto_msgTypes,
	}.Build()
	File_city_event_v1_event_proto = out.File
	file_city_event_v1_event_proto_rawDesc = nil
	file_city_event_v1_event_proto_goTypes = nil
	file_city_event_v1_event_proto_depIdxs = nil
}
