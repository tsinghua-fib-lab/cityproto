// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/trip/v2/trip.proto

package tripv2

import (
	v2 "git.fiblab.net/sim/protos/go/city/geo/v2"
	v21 "git.fiblab.net/sim/protos/go/city/routing/v2"
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

// 出行偏好
// Preferred trip travel mode
type TripMode int32

const (
	// 未指定出行方式
	// unspecified
	TripMode_TRIP_MODE_UNSPECIFIED TripMode = 0
	// 仅步行
	// walking only
	TripMode_TRIP_MODE_WALK_ONLY TripMode = 1
	// 仅开车
	// driving only
	TripMode_TRIP_MODE_DRIVE_ONLY TripMode = 2
	// 乘坐公交车（含站点间步行换乘）
	// taking bus with walking to transit bus line between stations
	TripMode_TRIP_MODE_BUS_WALK TripMode = 4
	// 当有自行车时选择骑自行车，否则步行
	// Riding bikes if avaible, otherwise walking
	TripMode_TRIP_MODE_BIKE_WALK TripMode = 5
)

// Enum value maps for TripMode.
var (
	TripMode_name = map[int32]string{
		0: "TRIP_MODE_UNSPECIFIED",
		1: "TRIP_MODE_WALK_ONLY",
		2: "TRIP_MODE_DRIVE_ONLY",
		4: "TRIP_MODE_BUS_WALK",
		5: "TRIP_MODE_BIKE_WALK",
	}
	TripMode_value = map[string]int32{
		"TRIP_MODE_UNSPECIFIED": 0,
		"TRIP_MODE_WALK_ONLY":   1,
		"TRIP_MODE_DRIVE_ONLY":  2,
		"TRIP_MODE_BUS_WALK":    4,
		"TRIP_MODE_BIKE_WALK":   5,
	}
)

func (x TripMode) Enum() *TripMode {
	p := new(TripMode)
	*p = x
	return p
}

func (x TripMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TripMode) Descriptor() protoreflect.EnumDescriptor {
	return file_city_trip_v2_trip_proto_enumTypes[0].Descriptor()
}

func (TripMode) Type() protoreflect.EnumType {
	return &file_city_trip_v2_trip_proto_enumTypes[0]
}

func (x TripMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TripMode.Descriptor instead.
func (TripMode) EnumDescriptor() ([]byte, []int) {
	return file_city_trip_v2_trip_proto_rawDescGZIP(), []int{0}
}

// 出行
// Trip
type Trip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 出行方式
	// trip mode
	Mode TripMode `protobuf:"varint,1,opt,name=mode,proto3,enum=city.trip.v2.TripMode" json:"mode,omitempty" bson:"mode" db:"mode" yaml:"mode"`
	// 目的地
	// destination
	End *v2.Position `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty" yaml:"end" bson:"end" db:"end"`
	// 期望的出发时间（单位: 秒）
	// Expected departure time (in seconds)
	DepartureTime *float64 `protobuf:"fixed64,3,opt,name=departure_time,json=departureTime,proto3,oneof" json:"departure_time,omitempty" yaml:"departure_time" bson:"departure_time" db:"departure_time"`
	// 期望的等待时间（单位：秒），如果departure_time为空则wait_time默认为0
	// The expected waiting time (in seconds), if departure_time is empty, wait_time defaults to 0
	WaitTime *float64 `protobuf:"fixed64,4,opt,name=wait_time,json=waitTime,proto3,oneof" json:"wait_time,omitempty" yaml:"wait_time" bson:"wait_time" db:"wait_time"`
	// 期望的到达时间（单位: 秒）
	// Expected arrival time (in seconds)
	ArrivalTime *float64 `protobuf:"fixed64,5,opt,name=arrival_time,json=arrivalTime,proto3,oneof" json:"arrival_time,omitempty" yaml:"arrival_time" bson:"arrival_time" db:"arrival_time"`
	// 本次出行目的地的活动名
	// The activity name of the destination for this trip
	Activity *string `protobuf:"bytes,6,opt,name=activity,proto3,oneof" json:"activity,omitempty" db:"activity" yaml:"activity" bson:"activity"`
	// 本次出行对应的可视化模型（覆盖Person Attribute中的默认模型）
	// The visual model corresponding to this trip (overriding the default model in Person Attribute)
	Model *string `protobuf:"bytes,8,opt,name=model,proto3,oneof" json:"model,omitempty" yaml:"model" bson:"model" db:"model"`
	// 预计算的导航结果
	// Pre calculated routing results
	Routes []*v21.Journey `protobuf:"bytes,7,rep,name=routes,proto3" json:"routes,omitempty" yaml:"routes" bson:"routes" db:"routes"`
}

func (x *Trip) Reset() {
	*x = Trip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_trip_v2_trip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trip) ProtoMessage() {}

func (x *Trip) ProtoReflect() protoreflect.Message {
	mi := &file_city_trip_v2_trip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trip.ProtoReflect.Descriptor instead.
func (*Trip) Descriptor() ([]byte, []int) {
	return file_city_trip_v2_trip_proto_rawDescGZIP(), []int{0}
}

func (x *Trip) GetMode() TripMode {
	if x != nil {
		return x.Mode
	}
	return TripMode_TRIP_MODE_UNSPECIFIED
}

func (x *Trip) GetEnd() *v2.Position {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Trip) GetDepartureTime() float64 {
	if x != nil && x.DepartureTime != nil {
		return *x.DepartureTime
	}
	return 0
}

func (x *Trip) GetWaitTime() float64 {
	if x != nil && x.WaitTime != nil {
		return *x.WaitTime
	}
	return 0
}

func (x *Trip) GetArrivalTime() float64 {
	if x != nil && x.ArrivalTime != nil {
		return *x.ArrivalTime
	}
	return 0
}

func (x *Trip) GetActivity() string {
	if x != nil && x.Activity != nil {
		return *x.Activity
	}
	return ""
}

func (x *Trip) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *Trip) GetRoutes() []*v21.Journey {
	if x != nil {
		return x.Routes
	}
	return nil
}

// 时刻表
// Schedule
// 关于出发时间的说明如下：
// The explanation about the departure time is as follows:
//  1. Schedule的开始时刻是 departure_time 或者 参考时刻+wait_time，
//  1. The start time of the Schedule is either departure_time or reference time+wait_time,
//     参考时刻定义为上一Schedule的结束时刻(即它最后一个Trip的结束时刻)，
//     The reference time is defined as the end time of the previous Schedule (i.e. the end time of its last Trip),
//     或者当它为第一个Schedule时定义为Person更新Schedule后的首次Update
//     Alternatively, when it is the first Schedule, it can be defined as the first Update time after Person updates the Schedule
//     时刻(当有准确时间要求时建议直接指定departure_time)
//     (it is recommended to specify departuretime directly when there is an accurate time requirement)
//  2. Trip的开始时刻是 departure_time 或者 参考时刻+wait_time，参考
//  2. The start time of the Trip is either departure_time or reference time+wait_time,
//     时刻定义为上一Trip的结束时刻，或者当它为第一个Trip时定义为所属的
//     The reference time is defined as the end time of the previous Trip, or when it is the first Trip,
//     Schedule的开始时刻
//     it is defined as the start time of the Schedule to which it belongs
//  3. Person的实际运行时刻取决于Trip的开始时刻，例如它的首次运行是第一
//  3. The actual running time of a Person depends on the start time of the Trip,
//     个Schedule中第一个Trip的开始时刻
//     for example, its first run is the start time of the first Trip in the first Schedule
//
// FAQ
// Q1: 同时指定Schedule和第一个Trip的departure_time会怎样？
// Q1: What would happen if both the Schedule and the departuretime of the first Trip were specified simultaneously?
// A1: 参照(2)，只看Trip的departure_time
// A1: Referring to (2), only depend on the departuretime of Trip
// Q2: Schedule和第一个Trip同时指定wait_time=10会怎样？
// Q2: What would happen if both the Schedule and the first Trip were specified with wait_time=10 at the same time?
// A2: 参照(2)，等待时间为10+10=20
// A2: Referring to (2), the waiting time is 10+10=20
type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 出行列表
	// List of trips
	Trips []*Trip `protobuf:"bytes,1,rep,name=trips,proto3" json:"trips,omitempty" yaml:"trips" bson:"trips" db:"trips"`
	// trips的执行次数，0表示无限循环，大于0表示执行几次
	// The number of times trips are executed, where 0 represents infinite loops and greater than 0 represents how many times they are executed
	LoopCount int32 `protobuf:"varint,2,opt,name=loop_count,json=loopCount,proto3" json:"loop_count,omitempty" yaml:"loop_count" bson:"loop_count" db:"loop_count"`
	// 期望的出发时间（单位: 秒）
	// Expected departure time (in seconds)
	DepartureTime *float64 `protobuf:"fixed64,3,opt,name=departure_time,json=departureTime,proto3,oneof" json:"departure_time,omitempty" yaml:"departure_time" bson:"departure_time" db:"departure_time"`
	// 期望的等待时间（单位：秒），如果departure_time为空则wait_time默认为0
	// Expected waiting time (in seconds), if departure_time is empty, wait_time defaults to 0
	WaitTime *float64 `protobuf:"fixed64,4,opt,name=wait_time,json=waitTime,proto3,oneof" json:"wait_time,omitempty" yaml:"wait_time" bson:"wait_time" db:"wait_time"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_trip_v2_trip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_city_trip_v2_trip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_city_trip_v2_trip_proto_rawDescGZIP(), []int{1}
}

func (x *Schedule) GetTrips() []*Trip {
	if x != nil {
		return x.Trips
	}
	return nil
}

func (x *Schedule) GetLoopCount() int32 {
	if x != nil {
		return x.LoopCount
	}
	return 0
}

func (x *Schedule) GetDepartureTime() float64 {
	if x != nil && x.DepartureTime != nil {
		return *x.DepartureTime
	}
	return 0
}

func (x *Schedule) GetWaitTime() float64 {
	if x != nil && x.WaitTime != nil {
		return *x.WaitTime
	}
	return 0
}

var File_city_trip_v2_trip_proto protoreflect.FileDescriptor

var file_city_trip_v2_trip_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x74,
	0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x69, 0x74, 0x79, 0x2e,
	0x74, 0x72, 0x69, 0x70, 0x2e, 0x76, 0x32, 0x1a, 0x15, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65,
	0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x63, 0x69, 0x74, 0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x03,
	0x0a, 0x04, 0x54, 0x72, 0x69, 0x70, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x70,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x0e, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x77, 0x61, 0x69, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x08, 0x77, 0x61,
	0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x61, 0x72, 0x72,
	0x69, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x02, 0x52, 0x0b, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xc2, 0x01, 0x0a, 0x08, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x70,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x52, 0x05, 0x74, 0x72, 0x69, 0x70, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x0e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x77, 0x61,
	0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52,
	0x08, 0x77, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0x89, 0x01,
	0x0a, 0x08, 0x54, 0x72, 0x69, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x52,
	0x49, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x52, 0x49, 0x50, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x5f, 0x57, 0x41, 0x4c, 0x4b, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x18,
	0x0a, 0x14, 0x54, 0x52, 0x49, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x52, 0x49, 0x56,
	0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x52, 0x49, 0x50,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x4c, 0x4b, 0x10, 0x04,
	0x12, 0x17, 0x0a, 0x13, 0x54, 0x52, 0x49, 0x50, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x49,
	0x4b, 0x45, 0x5f, 0x57, 0x41, 0x4c, 0x4b, 0x10, 0x05, 0x42, 0xa1, 0x01, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x76, 0x32, 0x42, 0x09,
	0x54, 0x72, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74,
	0x72, 0x69, 0x70, 0x2f, 0x76, 0x32, 0x3b, 0x74, 0x72, 0x69, 0x70, 0x76, 0x32, 0xa2, 0x02, 0x03,
	0x43, 0x54, 0x58, 0xaa, 0x02, 0x0c, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x2e,
	0x56, 0x32, 0xca, 0x02, 0x0c, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x54, 0x72, 0x69, 0x70, 0x5c, 0x56,
	0x32, 0xe2, 0x02, 0x18, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x54, 0x72, 0x69, 0x70, 0x5c, 0x56, 0x32,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x43,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x54, 0x72, 0x69, 0x70, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_trip_v2_trip_proto_rawDescOnce sync.Once
	file_city_trip_v2_trip_proto_rawDescData = file_city_trip_v2_trip_proto_rawDesc
)

func file_city_trip_v2_trip_proto_rawDescGZIP() []byte {
	file_city_trip_v2_trip_proto_rawDescOnce.Do(func() {
		file_city_trip_v2_trip_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_trip_v2_trip_proto_rawDescData)
	})
	return file_city_trip_v2_trip_proto_rawDescData
}

var file_city_trip_v2_trip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_city_trip_v2_trip_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_city_trip_v2_trip_proto_goTypes = []any{
	(TripMode)(0),       // 0: city.trip.v2.TripMode
	(*Trip)(nil),        // 1: city.trip.v2.Trip
	(*Schedule)(nil),    // 2: city.trip.v2.Schedule
	(*v2.Position)(nil), // 3: city.geo.v2.Position
	(*v21.Journey)(nil), // 4: city.routing.v2.Journey
}
var file_city_trip_v2_trip_proto_depIdxs = []int32{
	0, // 0: city.trip.v2.Trip.mode:type_name -> city.trip.v2.TripMode
	3, // 1: city.trip.v2.Trip.end:type_name -> city.geo.v2.Position
	4, // 2: city.trip.v2.Trip.routes:type_name -> city.routing.v2.Journey
	1, // 3: city.trip.v2.Schedule.trips:type_name -> city.trip.v2.Trip
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_city_trip_v2_trip_proto_init() }
func file_city_trip_v2_trip_proto_init() {
	if File_city_trip_v2_trip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_trip_v2_trip_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Trip); i {
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
		file_city_trip_v2_trip_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Schedule); i {
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
	file_city_trip_v2_trip_proto_msgTypes[0].OneofWrappers = []any{}
	file_city_trip_v2_trip_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_trip_v2_trip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_trip_v2_trip_proto_goTypes,
		DependencyIndexes: file_city_trip_v2_trip_proto_depIdxs,
		EnumInfos:         file_city_trip_v2_trip_proto_enumTypes,
		MessageInfos:      file_city_trip_v2_trip_proto_msgTypes,
	}.Build()
	File_city_trip_v2_trip_proto = out.File
	file_city_trip_v2_trip_proto_rawDesc = nil
	file_city_trip_v2_trip_proto_goTypes = nil
	file_city_trip_v2_trip_proto_depIdxs = nil
}
