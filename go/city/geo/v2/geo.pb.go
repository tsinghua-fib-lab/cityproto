// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: city/geo/v2/geo.proto

package geov2

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

// WGS84经纬度坐标
// WGS84 longitute and latitude coordinates
type LongLatPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 经度
	// longitude
	Longitude float64 `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty" yaml:"longitude" bson:"longitude" db:"longitude"`
	// 纬度
	// latitude
	Latitude float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty" yaml:"latitude" bson:"latitude" db:"latitude"`
}

func (x *LongLatPosition) Reset() {
	*x = LongLatPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_geo_v2_geo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongLatPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongLatPosition) ProtoMessage() {}

func (x *LongLatPosition) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

// XY坐标
// XY coordinates
type XYPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// x坐标，单位米，对应经度
	// x coordinate, in meters, corresponding to longitude
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty" db:"x" yaml:"x" bson:"x"`
	// y坐标，单位米，对应纬度
	// y coordinate, in meters, corresponding to latitude
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty" bson:"y" db:"y" yaml:"y"`
}

func (x *XYPosition) Reset() {
	*x = XYPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_geo_v2_geo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XYPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XYPosition) ProtoMessage() {}

func (x *XYPosition) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

// 地图坐标（车道+距离s）
// Map coordinates (lane ID + distance s)
type LanePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 车道id
	// Lane ID
	LaneId int32 `protobuf:"varint,1,opt,name=lane_id,json=laneId,proto3" json:"lane_id,omitempty" db:"lane_id" yaml:"lane_id" bson:"lane_id"`
	// s是车道上的点到车道起点的距离
	// s is the distance from the point on the lane to the starting point of the lane
	S float64 `protobuf:"fixed64,2,opt,name=s,proto3" json:"s,omitempty" yaml:"s" bson:"s" db:"s"`
}

func (x *LanePosition) Reset() {
	*x = LanePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_geo_v2_geo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanePosition) ProtoMessage() {}

func (x *LanePosition) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AOI ID
	AoiId int32 `protobuf:"varint,1,opt,name=aoi_id,json=aoiId,proto3" json:"aoi_id,omitempty" yaml:"aoi_id" bson:"aoi_id" db:"aoi_id"`
	// POI ID，需要是aoi_id的子poi，否则该值无效
	// POI ID, needs to be a sub-poi of aoi_id, otherwise the value is invalid
	PoiId *int32 `protobuf:"varint,2,opt,name=poi_id,json=poiId,proto3,oneof" json:"poi_id,omitempty" yaml:"poi_id" bson:"poi_id" db:"poi_id"`
}

func (x *AoiPosition) Reset() {
	*x = AoiPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_geo_v2_geo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AoiPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AoiPosition) ProtoMessage() {}

func (x *AoiPosition) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 地图坐标AOI（必须提供其中之一）
	// Map coordinates AOI (one of these must be provided)
	LanePosition *LanePosition `protobuf:"bytes,1,opt,name=lane_position,json=lanePosition,proto3,oneof" json:"lane_position,omitempty" db:"lane_position" yaml:"lane_position" bson:"lane_position"`
	// 地图坐标Lane+S（必须提供其中之一）
	// Map coordinates Lane+S (one of these must be provided)
	AoiPosition *AoiPosition `protobuf:"bytes,2,opt,name=aoi_position,json=aoiPosition,proto3,oneof" json:"aoi_position,omitempty" db:"aoi_position" yaml:"aoi_position" bson:"aoi_position"`
	// WGS84经纬度坐标
	// WGS84 longitute and latitude coordinates
	LonglatPosition *LongLatPosition `protobuf:"bytes,3,opt,name=longlat_position,json=longlatPosition,proto3,oneof" json:"longlat_position,omitempty" yaml:"longlat_position" bson:"longlat_position" db:"longlat_position"`
	// XY坐标
	// XY coordinates
	XyPosition *XYPosition `protobuf:"bytes,4,opt,name=xy_position,json=xyPosition,proto3,oneof" json:"xy_position,omitempty" db:"xy_position" yaml:"xy_position" bson:"xy_position"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_geo_v2_geo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 最小经度
	// minimum longitude
	MinLongitude float64 `protobuf:"fixed64,1,opt,name=min_longitude,json=minLongitude,proto3" json:"min_longitude,omitempty" yaml:"min_longitude" bson:"min_longitude" db:"min_longitude"`
	// 最小纬度
	// minimum latitude
	MinLatitude float64 `protobuf:"fixed64,2,opt,name=min_latitude,json=minLatitude,proto3" json:"min_latitude,omitempty" yaml:"min_latitude" bson:"min_latitude" db:"min_latitude"`
	// 最大经度
	// maximu longitude
	MaxLongitude float64 `protobuf:"fixed64,3,opt,name=max_longitude,json=maxLongitude,proto3" json:"max_longitude,omitempty" yaml:"max_longitude" bson:"max_longitude" db:"max_longitude"`
	// 最大纬度
	// minimum longitude
	MaxLatitude float64 `protobuf:"fixed64,4,opt,name=max_latitude,json=maxLatitude,proto3" json:"max_latitude,omitempty" db:"max_latitude" yaml:"max_latitude" bson:"max_latitude"`
}

func (x *LongLatBBox) Reset() {
	*x = LongLatBBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_geo_v2_geo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongLatBBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongLatBBox) ProtoMessage() {}

func (x *LongLatBBox) ProtoReflect() protoreflect.Message {
	mi := &file_city_geo_v2_geo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_city_geo_v2_geo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65,
	0x6f, 0x2e, 0x76, 0x32, 0x22, 0x4b, 0x0a, 0x0f, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x22, 0x28, 0x0a, 0x0a, 0x58, 0x59, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x22, 0x35, 0x0a, 0x0c, 0x4c,
	0x61, 0x6e, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6c,
	0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x61,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x01, 0x73, 0x22, 0x4b, 0x0a, 0x0b, 0x41, 0x6f, 0x69, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x6f, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x61, 0x6f, 0x69, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x6f, 0x69, 0x5f, 0x69, 0x64, 0x22,
	0xe6, 0x02, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0d,
	0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76,
	0x32, 0x2e, 0x4c, 0x61, 0x6e, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x40, 0x0a, 0x0c, 0x61, 0x6f, 0x69, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67,
	0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x6f, 0x69, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x01, 0x52, 0x0b, 0x61, 0x6f, 0x69, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x4c, 0x0a, 0x10, 0x6c, 0x6f, 0x6e, 0x67, 0x6c, 0x61, 0x74, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x6e, 0x67,
	0x4c, 0x61, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52, 0x0f, 0x6c,
	0x6f, 0x6e, 0x67, 0x6c, 0x61, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x3d, 0x0a, 0x0b, 0x78, 0x79, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65,
	0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x58, 0x59, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x03, 0x52, 0x0a, 0x78, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x6f, 0x69, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x6c, 0x61, 0x74, 0x5f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x78, 0x79, 0x5f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x6e,
	0x67, 0x4c, 0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0c, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x4c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6d, 0x61, 0x78,
	0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x99, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32, 0x42, 0x08, 0x47, 0x65,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69,
	0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x2f,
	0x76, 0x32, 0x3b, 0x67, 0x65, 0x6f, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x47, 0x58, 0xaa, 0x02,
	0x0b, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x6f, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0b, 0x43,
	0x69, 0x74, 0x79, 0x5c, 0x47, 0x65, 0x6f, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x17, 0x43, 0x69, 0x74,
	0x79, 0x5c, 0x47, 0x65, 0x6f, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x47, 0x65, 0x6f,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_geo_v2_geo_proto_rawDescOnce sync.Once
	file_city_geo_v2_geo_proto_rawDescData = file_city_geo_v2_geo_proto_rawDesc
)

func file_city_geo_v2_geo_proto_rawDescGZIP() []byte {
	file_city_geo_v2_geo_proto_rawDescOnce.Do(func() {
		file_city_geo_v2_geo_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_geo_v2_geo_proto_rawDescData)
	})
	return file_city_geo_v2_geo_proto_rawDescData
}

var file_city_geo_v2_geo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_city_geo_v2_geo_proto_goTypes = []interface{}{
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
	if !protoimpl.UnsafeEnabled {
		file_city_geo_v2_geo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongLatPosition); i {
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
		file_city_geo_v2_geo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XYPosition); i {
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
		file_city_geo_v2_geo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanePosition); i {
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
		file_city_geo_v2_geo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AoiPosition); i {
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
		file_city_geo_v2_geo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_city_geo_v2_geo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongLatBBox); i {
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
	file_city_geo_v2_geo_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_city_geo_v2_geo_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_geo_v2_geo_proto_rawDesc,
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
	file_city_geo_v2_geo_proto_rawDesc = nil
	file_city_geo_v2_geo_proto_goTypes = nil
	file_city_geo_v2_geo_proto_depIdxs = nil
}
