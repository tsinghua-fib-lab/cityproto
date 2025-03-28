// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: city/geo/v2/geo.proto

package geov2

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

// WGS84经纬度坐标
// WGS84 longitute and latitude coordinates
type LongLatPosition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 经度
	// longitude
	Longitude float64 `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty" bson:"longitude" db:"longitude" yaml:"longitude"`
	// 纬度
	// latitude
	Latitude float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty" bson:"latitude" db:"latitude" yaml:"latitude"`
	// 高程（单位：米）
	// elevation (unit: meters)
	Z             *float64 `protobuf:"fixed64,3,opt,name=z,proto3,oneof" json:"z,omitempty" bson:"z" db:"z" yaml:"z"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LongLatPosition) Reset() {
	*x = LongLatPosition{}
	mi := &file_city_geo_v2_geo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LongLatPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongLatPosition) ProtoMessage() {}

func (x *LongLatPosition) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongLatPosition.ProtoReflect.Descriptor instead.
func (*LongLatPosition) Descriptor() ([]byte, []int) {
	return file_city_geo_v2_geo_proto_rawDescGZIP(), []int{0}
}

func (x *LongLatPosition) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *LongLatPosition) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LongLatPosition) GetZ() float64 {
	if x != nil && x.Z != nil {
		return *x.Z
	}
	return 0
}

// XY坐标
// XY coordinates
type XYPosition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// x坐标，单位米，对应经度
	// x coordinate, in meters, corresponding to longitude
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty" bson:"x" db:"x" yaml:"x"`
	// y坐标，单位米，对应纬度
	// y coordinate, in meters, corresponding to latitude
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty" bson:"y" db:"y" yaml:"y"`
	// z坐标，单位米，对应高程
	// z coordinate, in meters, corresponding to elevation
	Z             *float64 `protobuf:"fixed64,3,opt,name=z,proto3,oneof" json:"z,omitempty" bson:"z" db:"z" yaml:"z"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *XYPosition) Reset() {
	*x = XYPosition{}
	mi := &file_city_geo_v2_geo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *XYPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XYPosition) ProtoMessage() {}

func (x *XYPosition) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XYPosition.ProtoReflect.Descriptor instead.
func (*XYPosition) Descriptor() ([]byte, []int) {
	return file_city_geo_v2_geo_proto_rawDescGZIP(), []int{1}
}

func (x *XYPosition) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *XYPosition) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *XYPosition) GetZ() float64 {
	if x != nil && x.Z != nil {
		return *x.Z
	}
	return 0
}

// 地图坐标（车道+距离s）
// Map coordinates (lane ID + distance s)
type LanePosition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 车道id
	// Lane ID
	LaneId int32 `protobuf:"varint,1,opt,name=lane_id,json=laneId,proto3" json:"lane_id,omitempty" bson:"lane_id" db:"lane_id" yaml:"lane_id"`
	// s是车道上的点到车道起点的距离
	// s is the distance from the point on the lane to the starting point of the lane
	S             float64 `protobuf:"fixed64,2,opt,name=s,proto3" json:"s,omitempty" bson:"s" db:"s" yaml:"s"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LanePosition) Reset() {
	*x = LanePosition{}
	mi := &file_city_geo_v2_geo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LanePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanePosition) ProtoMessage() {}

func (x *LanePosition) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanePosition.ProtoReflect.Descriptor instead.
func (*LanePosition) Descriptor() ([]byte, []int) {
	return file_city_geo_v2_geo_proto_rawDescGZIP(), []int{2}
}

func (x *LanePosition) GetLaneId() int32 {
	if x != nil {
		return x.LaneId
	}
	return 0
}

func (x *LanePosition) GetS() float64 {
	if x != nil {
		return x.S
	}
	return 0
}

// 地图坐标（AOI）
// Map coordinates (AOI)
type AoiPosition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// AOI ID
	AoiId int32 `protobuf:"varint,1,opt,name=aoi_id,json=aoiId,proto3" json:"aoi_id,omitempty" bson:"aoi_id" db:"aoi_id" yaml:"aoi_id"`
	// POI ID，需要是aoi_id的子poi，否则该值无效
	// POI ID, needs to be a sub-poi of aoi_id, otherwise the value is invalid
	PoiId         *int32 `protobuf:"varint,2,opt,name=poi_id,json=poiId,proto3,oneof" json:"poi_id,omitempty" bson:"poi_id" db:"poi_id" yaml:"poi_id"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AoiPosition) Reset() {
	*x = AoiPosition{}
	mi := &file_city_geo_v2_geo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AoiPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AoiPosition) ProtoMessage() {}

func (x *AoiPosition) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AoiPosition.ProtoReflect.Descriptor instead.
func (*AoiPosition) Descriptor() ([]byte, []int) {
	return file_city_geo_v2_geo_proto_rawDescGZIP(), []int{3}
}

func (x *AoiPosition) GetAoiId() int32 {
	if x != nil {
		return x.AoiId
	}
	return 0
}

func (x *AoiPosition) GetPoiId() int32 {
	if x != nil && x.PoiId != nil {
		return *x.PoiId
	}
	return 0
}

// 坐标，如果多种坐标同时存在，两两之间必须满足映射关系，同时逻辑坐标是必须提供的
// Coordinates, if multiple coordinates exist at the same time, the mapping relationship between them must be satisfied, and logical coordinates must be provided.
type Position struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 地图坐标AOI（必须提供其中之一）
	// Map coordinates AOI (one of these must be provided)
	LanePosition *LanePosition `protobuf:"bytes,1,opt,name=lane_position,json=lanePosition,proto3,oneof" json:"lane_position,omitempty" bson:"lane_position" db:"lane_position" yaml:"lane_position"`
	// 地图坐标Lane+S（必须提供其中之一）
	// Map coordinates Lane+S (one of these must be provided)
	AoiPosition *AoiPosition `protobuf:"bytes,2,opt,name=aoi_position,json=aoiPosition,proto3,oneof" json:"aoi_position,omitempty" bson:"aoi_position" db:"aoi_position" yaml:"aoi_position"`
	// WGS84经纬度坐标
	// WGS84 longitute and latitude coordinates
	LonglatPosition *LongLatPosition `protobuf:"bytes,3,opt,name=longlat_position,json=longlatPosition,proto3,oneof" json:"longlat_position,omitempty" bson:"longlat_position" db:"longlat_position" yaml:"longlat_position"`
	// XY坐标
	// XY coordinates
	XyPosition    *XYPosition `protobuf:"bytes,4,opt,name=xy_position,json=xyPosition,proto3,oneof" json:"xy_position,omitempty" bson:"xy_position" db:"xy_position" yaml:"xy_position"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Position) Reset() {
	*x = Position{}
	mi := &file_city_geo_v2_geo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_city_geo_v2_geo_proto_rawDescGZIP(), []int{4}
}

func (x *Position) GetLanePosition() *LanePosition {
	if x != nil {
		return x.LanePosition
	}
	return nil
}

func (x *Position) GetAoiPosition() *AoiPosition {
	if x != nil {
		return x.AoiPosition
	}
	return nil
}

func (x *Position) GetLonglatPosition() *LongLatPosition {
	if x != nil {
		return x.LonglatPosition
	}
	return nil
}

func (x *Position) GetXyPosition() *XYPosition {
	if x != nil {
		return x.XyPosition
	}
	return nil
}

// 经纬度矩形区域
// latitude and longitude rectangular area
type LongLatBBox struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 最小经度
	// minimum longitude
	MinLongitude float64 `protobuf:"fixed64,1,opt,name=min_longitude,json=minLongitude,proto3" json:"min_longitude,omitempty" bson:"min_longitude" db:"min_longitude" yaml:"min_longitude"`
	// 最小纬度
	// minimum latitude
	MinLatitude float64 `protobuf:"fixed64,2,opt,name=min_latitude,json=minLatitude,proto3" json:"min_latitude,omitempty" bson:"min_latitude" db:"min_latitude" yaml:"min_latitude"`
	// 最大经度
	// maximu longitude
	MaxLongitude float64 `protobuf:"fixed64,3,opt,name=max_longitude,json=maxLongitude,proto3" json:"max_longitude,omitempty" bson:"max_longitude" db:"max_longitude" yaml:"max_longitude"`
	// 最大纬度
	// minimum longitude
	MaxLatitude   float64 `protobuf:"fixed64,4,opt,name=max_latitude,json=maxLatitude,proto3" json:"max_latitude,omitempty" bson:"max_latitude" db:"max_latitude" yaml:"max_latitude"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LongLatBBox) Reset() {
	*x = LongLatBBox{}
	mi := &file_city_geo_v2_geo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LongLatBBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongLatBBox) ProtoMessage() {}

func (x *LongLatBBox) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongLatBBox.ProtoReflect.Descriptor instead.
func (*LongLatBBox) Descriptor() ([]byte, []int) {
	return file_city_geo_v2_geo_proto_rawDescGZIP(), []int{5}
}

func (x *LongLatBBox) GetMinLongitude() float64 {
	if x != nil {
		return x.MinLongitude
	}
	return 0
}

func (x *LongLatBBox) GetMinLatitude() float64 {
	if x != nil {
		return x.MinLatitude
	}
	return 0
}

func (x *LongLatBBox) GetMaxLongitude() float64 {
	if x != nil {
		return x.MaxLongitude
	}
	return 0
}

func (x *LongLatBBox) GetMaxLatitude() float64 {
	if x != nil {
		return x.MaxLatitude
	}
	return 0
}

var File_city_geo_v2_geo_proto protoreflect.FileDescriptor

const file_city_geo_v2_geo_proto_rawDesc = "" +
	"\n" +
	"\x15city/geo/v2/geo.proto\x12\vcity.geo.v2\"d\n" +
	"\x0fLongLatPosition\x12\x1c\n" +
	"\tlongitude\x18\x01 \x01(\x01R\tlongitude\x12\x1a\n" +
	"\blatitude\x18\x02 \x01(\x01R\blatitude\x12\x11\n" +
	"\x01z\x18\x03 \x01(\x01H\x00R\x01z\x88\x01\x01B\x04\n" +
	"\x02_z\"A\n" +
	"\n" +
	"XYPosition\x12\f\n" +
	"\x01x\x18\x01 \x01(\x01R\x01x\x12\f\n" +
	"\x01y\x18\x02 \x01(\x01R\x01y\x12\x11\n" +
	"\x01z\x18\x03 \x01(\x01H\x00R\x01z\x88\x01\x01B\x04\n" +
	"\x02_z\"5\n" +
	"\fLanePosition\x12\x17\n" +
	"\alane_id\x18\x01 \x01(\x05R\x06laneId\x12\f\n" +
	"\x01s\x18\x02 \x01(\x01R\x01s\"K\n" +
	"\vAoiPosition\x12\x15\n" +
	"\x06aoi_id\x18\x01 \x01(\x05R\x05aoiId\x12\x1a\n" +
	"\x06poi_id\x18\x02 \x01(\x05H\x00R\x05poiId\x88\x01\x01B\t\n" +
	"\a_poi_id\"\xe6\x02\n" +
	"\bPosition\x12C\n" +
	"\rlane_position\x18\x01 \x01(\v2\x19.city.geo.v2.LanePositionH\x00R\flanePosition\x88\x01\x01\x12@\n" +
	"\faoi_position\x18\x02 \x01(\v2\x18.city.geo.v2.AoiPositionH\x01R\vaoiPosition\x88\x01\x01\x12L\n" +
	"\x10longlat_position\x18\x03 \x01(\v2\x1c.city.geo.v2.LongLatPositionH\x02R\x0flonglatPosition\x88\x01\x01\x12=\n" +
	"\vxy_position\x18\x04 \x01(\v2\x17.city.geo.v2.XYPositionH\x03R\n" +
	"xyPosition\x88\x01\x01B\x10\n" +
	"\x0e_lane_positionB\x0f\n" +
	"\r_aoi_positionB\x13\n" +
	"\x11_longlat_positionB\x0e\n" +
	"\f_xy_position\"\x9d\x01\n" +
	"\vLongLatBBox\x12#\n" +
	"\rmin_longitude\x18\x01 \x01(\x01R\fminLongitude\x12!\n" +
	"\fmin_latitude\x18\x02 \x01(\x01R\vminLatitude\x12#\n" +
	"\rmax_longitude\x18\x03 \x01(\x01R\fmaxLongitude\x12!\n" +
	"\fmax_latitude\x18\x04 \x01(\x01R\vmaxLatitudeB\x9c\x01\n" +
	"\x0fcom.city.geo.v2B\bGeoProtoP\x01Z1git.fiblab.net/sim/protos/v2/go/city/geo/v2;geov2\xa2\x02\x03CGX\xaa\x02\vCity.Geo.V2\xca\x02\vCity\\Geo\\V2\xe2\x02\x17City\\Geo\\V2\\GPBMetadata\xea\x02\rCity::Geo::V2b\x06proto3"

var (
	file_city_geo_v2_geo_proto_rawDescOnce sync.Once
	file_city_geo_v2_geo_proto_rawDescData []byte
)

func file_city_geo_v2_geo_proto_rawDescGZIP() []byte {
	file_city_geo_v2_geo_proto_rawDescOnce.Do(func() {
		file_city_geo_v2_geo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_geo_v2_geo_proto_rawDesc), len(file_city_geo_v2_geo_proto_rawDesc)))
	})
	return file_city_geo_v2_geo_proto_rawDescData
}

var file_city_geo_v2_geo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_city_geo_v2_geo_proto_goTypes = []any{
	(*LongLatPosition)(nil), // 0: city.geo.v2.LongLatPosition
	(*XYPosition)(nil),      // 1: city.geo.v2.XYPosition
	(*LanePosition)(nil),    // 2: city.geo.v2.LanePosition
	(*AoiPosition)(nil),     // 3: city.geo.v2.AoiPosition
	(*Position)(nil),        // 4: city.geo.v2.Position
	(*LongLatBBox)(nil),     // 5: city.geo.v2.LongLatBBox
}
var file_city_geo_v2_geo_proto_depIdxs = []int32{
	2, // 0: city.geo.v2.Position.lane_position:type_name -> city.geo.v2.LanePosition
	3, // 1: city.geo.v2.Position.aoi_position:type_name -> city.geo.v2.AoiPosition
	0, // 2: city.geo.v2.Position.longlat_position:type_name -> city.geo.v2.LongLatPosition
	1, // 3: city.geo.v2.Position.xy_position:type_name -> city.geo.v2.XYPosition
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_city_geo_v2_geo_proto_init() }
func file_city_geo_v2_geo_proto_init() {
	if File_city_geo_v2_geo_proto != nil {
		return
	}
	file_city_geo_v2_geo_proto_msgTypes[0].OneofWrappers = []any{}
	file_city_geo_v2_geo_proto_msgTypes[1].OneofWrappers = []any{}
	file_city_geo_v2_geo_proto_msgTypes[3].OneofWrappers = []any{}
	file_city_geo_v2_geo_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_geo_v2_geo_proto_rawDesc), len(file_city_geo_v2_geo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_geo_v2_geo_proto_goTypes,
		DependencyIndexes: file_city_geo_v2_geo_proto_depIdxs,
		MessageInfos:      file_city_geo_v2_geo_proto_msgTypes,
	}.Build()
	File_city_geo_v2_geo_proto = out.File
	file_city_geo_v2_geo_proto_goTypes = nil
	file_city_geo_v2_geo_proto_depIdxs = nil
}
