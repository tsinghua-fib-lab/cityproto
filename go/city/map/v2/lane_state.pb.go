// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/map/v2/lane_state.proto

package mapv2

import (
	v2 "git.fiblab.net/sim/protos/v2/go/city/person/v2"
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

// Lane状态
// Lane state
type LaneState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Lane ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// Lane上的人/车
	// person/vehicle on the lane
	Persons []*v2.PersonMotion `protobuf:"bytes,2,rep,name=persons,proto3" json:"persons,omitempty" bson:"persons" db:"persons" yaml:"persons"`
	// 平均速度（m/s）
	// average speed (m/s)
	AvgV float64 `protobuf:"fixed64,3,opt,name=avg_v,json=avgV,proto3" json:"avg_v,omitempty" bson:"avg_v" db:"avg_v" yaml:"avg_v"`
	// 是否限行
	// whether restricted
	Restriction bool `protobuf:"varint,4,opt,name=restriction,proto3" json:"restriction,omitempty" bson:"restriction" db:"restriction" yaml:"restriction"`
	// 交通灯状态
	// traffic light state
	LightState LightState `protobuf:"varint,5,opt,name=light_state,json=lightState,proto3,enum=city.map.v2.LightState" json:"light_state,omitempty" bson:"light_state" db:"light_state" yaml:"light_state"`
	// 当前进入车道的车辆数
	// current entering vehicle count
	InVehicleCnt int32 `protobuf:"varint,6,opt,name=in_vehicle_cnt,json=inVehicleCnt,proto3" json:"in_vehicle_cnt,omitempty" bson:"in_vehicle_cnt" db:"in_vehicle_cnt" yaml:"in_vehicle_cnt"`
	// 当前离开车道的车辆数
	// current leaving vehicle count
	OutVehicleCnt int32 `protobuf:"varint,7,opt,name=out_vehicle_cnt,json=outVehicleCnt,proto3" json:"out_vehicle_cnt,omitempty" bson:"out_vehicle_cnt" db:"out_vehicle_cnt" yaml:"out_vehicle_cnt"`
	// 总车数
	// total vehicle count
	VehicleCnt int32 `protobuf:"varint,8,opt,name=vehicle_cnt,json=vehicleCnt,proto3" json:"vehicle_cnt,omitempty" bson:"vehicle_cnt" db:"vehicle_cnt" yaml:"vehicle_cnt"`
	// 排队数量
	// queueing vehicle count
	TotalQueuingVehicleCnt int32 `protobuf:"varint,9,opt,name=total_queuing_vehicle_cnt,json=totalQueuingVehicleCnt,proto3" json:"total_queuing_vehicle_cnt,omitempty" bson:"total_queuing_vehicle_cnt" db:"total_queuing_vehicle_cnt" yaml:"total_queuing_vehicle_cnt"`
	// 排队时间
	// queueing time
	TotalQueuingTime float64 `protobuf:"fixed64,10,opt,name=total_queuing_time,json=totalQueuingTime,proto3" json:"total_queuing_time,omitempty" bson:"total_queuing_time" db:"total_queuing_time" yaml:"total_queuing_time"`
	// 平均排队时间
	// average queueing time
	AvgQueuingTime float64 `protobuf:"fixed64,11,opt,name=avg_queuing_time,json=avgQueuingTime,proto3" json:"avg_queuing_time,omitempty" bson:"avg_queuing_time" db:"avg_queuing_time" yaml:"avg_queuing_time"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LaneState) Reset() {
	*x = LaneState{}
	mi := &file_city_map_v2_lane_state_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LaneState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneState) ProtoMessage() {}

func (x *LaneState) ProtoReflect() protoreflect.Message {
	mi := &file_city_map_v2_lane_state_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneState.ProtoReflect.Descriptor instead.
func (*LaneState) Descriptor() ([]byte, []int) {
	return file_city_map_v2_lane_state_proto_rawDescGZIP(), []int{0}
}

func (x *LaneState) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LaneState) GetPersons() []*v2.PersonMotion {
	if x != nil {
		return x.Persons
	}
	return nil
}

func (x *LaneState) GetAvgV() float64 {
	if x != nil {
		return x.AvgV
	}
	return 0
}

func (x *LaneState) GetRestriction() bool {
	if x != nil {
		return x.Restriction
	}
	return false
}

func (x *LaneState) GetLightState() LightState {
	if x != nil {
		return x.LightState
	}
	return LightState_LIGHT_STATE_UNSPECIFIED
}

func (x *LaneState) GetInVehicleCnt() int32 {
	if x != nil {
		return x.InVehicleCnt
	}
	return 0
}

func (x *LaneState) GetOutVehicleCnt() int32 {
	if x != nil {
		return x.OutVehicleCnt
	}
	return 0
}

func (x *LaneState) GetVehicleCnt() int32 {
	if x != nil {
		return x.VehicleCnt
	}
	return 0
}

func (x *LaneState) GetTotalQueuingVehicleCnt() int32 {
	if x != nil {
		return x.TotalQueuingVehicleCnt
	}
	return 0
}

func (x *LaneState) GetTotalQueuingTime() float64 {
	if x != nil {
		return x.TotalQueuingTime
	}
	return 0
}

func (x *LaneState) GetAvgQueuingTime() float64 {
	if x != nil {
		return x.AvgQueuingTime
	}
	return 0
}

var File_city_map_v2_lane_state_proto protoreflect.FileDescriptor

const file_city_map_v2_lane_state_proto_rawDesc = "" +
	"\n" +
	"\x1ccity/map/v2/lane_state.proto\x12\vcity.map.v2\x1a\x17city/map/v2/light.proto\x1a\x1bcity/person/v2/motion.proto\"\xc6\x03\n" +
	"\tLaneState\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x126\n" +
	"\apersons\x18\x02 \x03(\v2\x1c.city.person.v2.PersonMotionR\apersons\x12\x13\n" +
	"\x05avg_v\x18\x03 \x01(\x01R\x04avgV\x12 \n" +
	"\vrestriction\x18\x04 \x01(\bR\vrestriction\x128\n" +
	"\vlight_state\x18\x05 \x01(\x0e2\x17.city.map.v2.LightStateR\n" +
	"lightState\x12$\n" +
	"\x0ein_vehicle_cnt\x18\x06 \x01(\x05R\finVehicleCnt\x12&\n" +
	"\x0fout_vehicle_cnt\x18\a \x01(\x05R\routVehicleCnt\x12\x1f\n" +
	"\vvehicle_cnt\x18\b \x01(\x05R\n" +
	"vehicleCnt\x129\n" +
	"\x19total_queuing_vehicle_cnt\x18\t \x01(\x05R\x16totalQueuingVehicleCnt\x12,\n" +
	"\x12total_queuing_time\x18\n" +
	" \x01(\x01R\x10totalQueuingTime\x12(\n" +
	"\x10avg_queuing_time\x18\v \x01(\x01R\x0eavgQueuingTimeB\xa2\x01\n" +
	"\x0fcom.city.map.v2B\x0eLaneStateProtoP\x01Z1git.fiblab.net/sim/protos/v2/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\vCity.Map.V2\xca\x02\vCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3"

var (
	file_city_map_v2_lane_state_proto_rawDescOnce sync.Once
	file_city_map_v2_lane_state_proto_rawDescData []byte
)

func file_city_map_v2_lane_state_proto_rawDescGZIP() []byte {
	file_city_map_v2_lane_state_proto_rawDescOnce.Do(func() {
		file_city_map_v2_lane_state_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_map_v2_lane_state_proto_rawDesc), len(file_city_map_v2_lane_state_proto_rawDesc)))
	})
	return file_city_map_v2_lane_state_proto_rawDescData
}

var file_city_map_v2_lane_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_city_map_v2_lane_state_proto_goTypes = []any{
	(*LaneState)(nil),       // 0: city.map.v2.LaneState
	(*v2.PersonMotion)(nil), // 1: city.person.v2.PersonMotion
	(LightState)(0),         // 2: city.map.v2.LightState
}
var file_city_map_v2_lane_state_proto_depIdxs = []int32{
	1, // 0: city.map.v2.LaneState.persons:type_name -> city.person.v2.PersonMotion
	2, // 1: city.map.v2.LaneState.light_state:type_name -> city.map.v2.LightState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_city_map_v2_lane_state_proto_init() }
func file_city_map_v2_lane_state_proto_init() {
	if File_city_map_v2_lane_state_proto != nil {
		return
	}
	file_city_map_v2_light_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_map_v2_lane_state_proto_rawDesc), len(file_city_map_v2_lane_state_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_map_v2_lane_state_proto_goTypes,
		DependencyIndexes: file_city_map_v2_lane_state_proto_depIdxs,
		MessageInfos:      file_city_map_v2_lane_state_proto_msgTypes,
	}.Build()
	File_city_map_v2_lane_state_proto = out.File
	file_city_map_v2_lane_state_proto_goTypes = nil
	file_city_map_v2_lane_state_proto_depIdxs = nil
}
