// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: city/person/v1/person.proto

package personv1

import (
	v2 "git.fiblab.net/sim/protos/go/city/geo/v2"
	v21 "git.fiblab.net/sim/protos/go/city/trip/v2"
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

// 智能体教育等级
// Agent education level
type Education int32

const (
	// 未指定
	// unspecified
	Education_EDUCATION_UNSPECIFIED Education = 0
	// 博士
	// doctor
	Education_EDUCATION_DOCTOR Education = 1
	// 硕士
	// master
	Education_EDUCATION_MASTER Education = 2
	// 本科
	// bachelor
	Education_EDUCATION_BACHELOR Education = 3
	// 高中
	// high school
	Education_EDUCATION_HIGH_SCHOOL Education = 4
	// 初中
	// junior high school
	Education_EDUCATION_JUNIOR_HIGH_SCHOOL Education = 5
	// 小学
	// primary school
	Education_EDUCATION_PRIMARY_SCHOOL Education = 6
	// 大专
	// college
	Education_EDUCATION_COLLEGE Education = 7
)

// Enum value maps for Education.
var (
	Education_name = map[int32]string{
		0: "EDUCATION_UNSPECIFIED",
		1: "EDUCATION_DOCTOR",
		2: "EDUCATION_MASTER",
		3: "EDUCATION_BACHELOR",
		4: "EDUCATION_HIGH_SCHOOL",
		5: "EDUCATION_JUNIOR_HIGH_SCHOOL",
		6: "EDUCATION_PRIMARY_SCHOOL",
		7: "EDUCATION_COLLEGE",
	}
	Education_value = map[string]int32{
		"EDUCATION_UNSPECIFIED":        0,
		"EDUCATION_DOCTOR":             1,
		"EDUCATION_MASTER":             2,
		"EDUCATION_BACHELOR":           3,
		"EDUCATION_HIGH_SCHOOL":        4,
		"EDUCATION_JUNIOR_HIGH_SCHOOL": 5,
		"EDUCATION_PRIMARY_SCHOOL":     6,
		"EDUCATION_COLLEGE":            7,
	}
)

func (x Education) Enum() *Education {
	p := new(Education)
	*p = x
	return p
}

func (x Education) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Education) Descriptor() protoreflect.EnumDescriptor {
	return file_city_person_v1_person_proto_enumTypes[0].Descriptor()
}

func (Education) Type() protoreflect.EnumType {
	return &file_city_person_v1_person_proto_enumTypes[0]
}

func (x Education) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Education.Descriptor instead.
func (Education) EnumDescriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{0}
}

// 智能体性别
// agent gender
type Gender int32

const (
	// 未指定
	// unspecified
	Gender_GENDER_UNSPECIFIED Gender = 0
	// 男性
	// male
	Gender_GENDER_MALE Gender = 1
	// 女性
	// female
	Gender_GENDER_FEMALE Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_UNSPECIFIED",
		1: "GENDER_MALE",
		2: "GENDER_FEMALE",
	}
	Gender_value = map[string]int32{
		"GENDER_UNSPECIFIED": 0,
		"GENDER_MALE":        1,
		"GENDER_FEMALE":      2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_city_person_v1_person_proto_enumTypes[1].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_city_person_v1_person_proto_enumTypes[1]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{1}
}

// 智能体消费水平
// agent consumption level
type Consumption int32

const (
	// 未指定
	// unspecified
	Consumption_CONSUMPTION_UNSPECIFIED Consumption = 0
	// 低
	// low
	Consumption_CONSUMPTION_LOW Consumption = 1
	// 较低
	// relatively low
	Consumption_CONSUMPTION_RELATIVELY_LOW Consumption = 2
	// 中等
	// medium
	Consumption_CONSUMPTION_MEDIUM Consumption = 3
	// 较高
	// relatively high
	Consumption_CONSUMPTION_RELATIVELY_HIGH Consumption = 4
	// 高
	// high
	Consumption_CONSUMPTION_HIGH Consumption = 5
)

// Enum value maps for Consumption.
var (
	Consumption_name = map[int32]string{
		0: "CONSUMPTION_UNSPECIFIED",
		1: "CONSUMPTION_LOW",
		2: "CONSUMPTION_RELATIVELY_LOW",
		3: "CONSUMPTION_MEDIUM",
		4: "CONSUMPTION_RELATIVELY_HIGH",
		5: "CONSUMPTION_HIGH",
	}
	Consumption_value = map[string]int32{
		"CONSUMPTION_UNSPECIFIED":     0,
		"CONSUMPTION_LOW":             1,
		"CONSUMPTION_RELATIVELY_LOW":  2,
		"CONSUMPTION_MEDIUM":          3,
		"CONSUMPTION_RELATIVELY_HIGH": 4,
		"CONSUMPTION_HIGH":            5,
	}
)

func (x Consumption) Enum() *Consumption {
	p := new(Consumption)
	*p = x
	return p
}

func (x Consumption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Consumption) Descriptor() protoreflect.EnumDescriptor {
	return file_city_person_v1_person_proto_enumTypes[2].Descriptor()
}

func (Consumption) Type() protoreflect.EnumType {
	return &file_city_person_v1_person_proto_enumTypes[2]
}

func (x Consumption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Consumption.Descriptor instead.
func (Consumption) EnumDescriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{2}
}

// 智能体属性（通用）
// Agent properties (general)
type PersonAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 单位: m，长度
	// length: m
	Length float64 `protobuf:"fixed64,2,opt,name=length,proto3" json:"length,omitempty" yaml:"length" bson:"length" db:"length"`
	// 单位: m，宽度
	// width: m
	Width float64 `protobuf:"fixed64,3,opt,name=width,proto3" json:"width,omitempty" bson:"width" db:"width" yaml:"width"`
	// 单位: m/s
	// max speed: m/s
	MaxSpeed float64 `protobuf:"fixed64,4,opt,name=max_speed,json=maxSpeed,proto3" json:"max_speed,omitempty" yaml:"max_speed" bson:"max_speed" db:"max_speed"`
	// 单位: m/s^2, 最大加速度（正值）
	// max accelaration: m/s^2 (positive value)
	MaxAcceleration float64 `protobuf:"fixed64,5,opt,name=max_acceleration,json=maxAcceleration,proto3" json:"max_acceleration,omitempty" yaml:"max_acceleration" bson:"max_acceleration" db:"max_acceleration"`
	// 单位: m/s^2, 最大减速度（负值）
	// max deceleration: m/s^2 (negative value)
	MaxBrakingAcceleration float64 `protobuf:"fixed64,6,opt,name=max_braking_acceleration,json=maxBrakingAcceleration,proto3" json:"max_braking_acceleration,omitempty" yaml:"max_braking_acceleration" bson:"max_braking_acceleration" db:"max_braking_acceleration"`
	// 单位: m/s^2, 一般加速度（正值），要求小于最大加速度
	// usual acceleration: m/s^2 (positive value), required to be less than the max acceleration
	UsualAcceleration float64 `protobuf:"fixed64,7,opt,name=usual_acceleration,json=usualAcceleration,proto3" json:"usual_acceleration,omitempty" yaml:"usual_acceleration" bson:"usual_acceleration" db:"usual_acceleration"`
	// 单位: m/s^2, 一般减速度（负值），要求大于最大减速度
	// usual deceleration: m/s^2 (negative value), required to be greater than the max deceleration
	UsualBrakingAcceleration float64 `protobuf:"fixed64,8,opt,name=usual_braking_acceleration,json=usualBrakingAcceleration,proto3" json:"usual_braking_acceleration,omitempty" yaml:"usual_braking_acceleration" bson:"usual_braking_acceleration" db:"usual_braking_acceleration"`
}

func (x *PersonAttribute) Reset() {
	*x = PersonAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonAttribute) ProtoMessage() {}

func (x *PersonAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonAttribute.ProtoReflect.Descriptor instead.
func (*PersonAttribute) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{0}
}

func (x *PersonAttribute) GetLength() float64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *PersonAttribute) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *PersonAttribute) GetMaxSpeed() float64 {
	if x != nil {
		return x.MaxSpeed
	}
	return 0
}

func (x *PersonAttribute) GetMaxAcceleration() float64 {
	if x != nil {
		return x.MaxAcceleration
	}
	return 0
}

func (x *PersonAttribute) GetMaxBrakingAcceleration() float64 {
	if x != nil {
		return x.MaxBrakingAcceleration
	}
	return 0
}

func (x *PersonAttribute) GetUsualAcceleration() float64 {
	if x != nil {
		return x.UsualAcceleration
	}
	return 0
}

func (x *PersonAttribute) GetUsualBrakingAcceleration() float64 {
	if x != nil {
		return x.UsualBrakingAcceleration
	}
	return 0
}

// 车辆附加属性
// Vehicle additional attributes
type VehicleAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 单位: m, 完成变道所需路程
	// Distance required to complete lane change: m
	LaneChangeLength float64 `protobuf:"fixed64,1,opt,name=lane_change_length,json=laneChangeLength,proto3" json:"lane_change_length,omitempty" yaml:"lane_change_length" bson:"lane_change_length" db:"lane_change_length"`
	// 单位：米，本车距离前车的最小距离
	// The minimum distance between the vehicle and the vehicle in front: m
	MinGap float64 `protobuf:"fixed64,2,opt,name=min_gap,json=minGap,proto3" json:"min_gap,omitempty" db:"min_gap" yaml:"min_gap" bson:"min_gap"`
}

func (x *VehicleAttribute) Reset() {
	*x = VehicleAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleAttribute) ProtoMessage() {}

func (x *VehicleAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleAttribute.ProtoReflect.Descriptor instead.
func (*VehicleAttribute) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{1}
}

func (x *VehicleAttribute) GetLaneChangeLength() float64 {
	if x != nil {
		return x.LaneChangeLength
	}
	return 0
}

func (x *VehicleAttribute) GetMinGap() float64 {
	if x != nil {
		return x.MinGap
	}
	return 0
}

// 公交车附加属性
// Bus additional attributes
type BusAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 公交线路ID
	// bus line ID
	LineId int32 `protobuf:"varint,1,opt,name=line_id,json=lineId,proto3" json:"line_id,omitempty" yaml:"line_id" bson:"line_id" db:"line_id"`
	// 公交车容量
	// bus capacity
	Capacity int32 `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty" db:"capacity" yaml:"capacity" bson:"capacity"`
}

func (x *BusAttribute) Reset() {
	*x = BusAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusAttribute) ProtoMessage() {}

func (x *BusAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusAttribute.ProtoReflect.Descriptor instead.
func (*BusAttribute) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{2}
}

func (x *BusAttribute) GetLineId() int32 {
	if x != nil {
		return x.LineId
	}
	return 0
}

func (x *BusAttribute) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

// 自行车附加属性
// Bike additional attributes
type BikeAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BikeAttribute) Reset() {
	*x = BikeAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BikeAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BikeAttribute) ProtoMessage() {}

func (x *BikeAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BikeAttribute.ProtoReflect.Descriptor instead.
func (*BikeAttribute) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{3}
}

// 智能体简介
// agent profile
type PersonProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 年龄
	// age
	Age int32 `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty" yaml:"age" bson:"age" db:"age"`
	// 教育水平
	// education level
	Education Education `protobuf:"varint,2,opt,name=education,proto3,enum=city.person.v1.Education" json:"education,omitempty" yaml:"education" bson:"education" db:"education"`
	// 性别
	// gender
	Gender Gender `protobuf:"varint,3,opt,name=gender,proto3,enum=city.person.v1.Gender" json:"gender,omitempty" db:"gender" yaml:"gender" bson:"gender"`
	// 消费水平
	// consumption level
	Consumption Consumption `protobuf:"varint,4,opt,name=consumption,proto3,enum=city.person.v1.Consumption" json:"consumption,omitempty" yaml:"consumption" bson:"consumption" db:"consumption"`
}

func (x *PersonProfile) Reset() {
	*x = PersonProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonProfile) ProtoMessage() {}

func (x *PersonProfile) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonProfile.ProtoReflect.Descriptor instead.
func (*PersonProfile) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{4}
}

func (x *PersonProfile) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *PersonProfile) GetEducation() Education {
	if x != nil {
		return x.Education
	}
	return Education_EDUCATION_UNSPECIFIED
}

func (x *PersonProfile) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_GENDER_UNSPECIFIED
}

func (x *PersonProfile) GetConsumption() Consumption {
	if x != nil {
		return x.Consumption
	}
	return Consumption_CONSUMPTION_UNSPECIFIED
}

// 智能体
// agent
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 智能体ID
	// agent ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id" bson:"id" db:"id"`
	// 参数
	// attribute
	Attribute *PersonAttribute `protobuf:"bytes,2,opt,name=attribute,proto3" json:"attribute,omitempty" yaml:"attribute" bson:"attribute" db:"attribute"`
	// 初始位置
	// initial position
	Home *v2.Position `protobuf:"bytes,3,opt,name=home,proto3" json:"home,omitempty" yaml:"home" bson:"home" db:"home"`
	// 初始日程
	// initial schedules
	Schedules []*v21.Schedule `protobuf:"bytes,4,rep,name=schedules,proto3" json:"schedules,omitempty" yaml:"schedules" bson:"schedules" db:"schedules"`
	// 车辆附加属性
	// vehicle addtional attribute
	VehicleAttribute *VehicleAttribute `protobuf:"bytes,7,opt,name=vehicle_attribute,json=vehicleAttribute,proto3,oneof" json:"vehicle_attribute,omitempty" yaml:"vehicle_attribute" bson:"vehicle_attribute" db:"vehicle_attribute"`
	// 公交车附加属性
	// bus additional attribute
	BusAttribute *BusAttribute `protobuf:"bytes,8,opt,name=bus_attribute,json=busAttribute,proto3,oneof" json:"bus_attribute,omitempty" yaml:"bus_attribute" bson:"bus_attribute" db:"bus_attribute"`
	// 自行车附加属性
	// bike addition attribute
	BikeAttribute *BikeAttribute `protobuf:"bytes,9,opt,name=bike_attribute,json=bikeAttribute,proto3,oneof" json:"bike_attribute,omitempty" yaml:"bike_attribute" bson:"bike_attribute" db:"bike_attribute"`
	// [可空] 额外的标签（例如：抢修车类型->电网）
	// [can be empty] additional tags (e.g. repair vehicle type -> power grid)
	Labels map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"labels" bson:"labels" db:"labels"`
	// [可空] 智能体简介
	// [can be empty] agent profile
	Profile *PersonProfile `protobuf:"bytes,11,opt,name=profile,proto3,oneof" json:"profile,omitempty" yaml:"profile" bson:"profile" db:"profile"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{5}
}

func (x *Person) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetAttribute() *PersonAttribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

func (x *Person) GetHome() *v2.Position {
	if x != nil {
		return x.Home
	}
	return nil
}

func (x *Person) GetSchedules() []*v21.Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

func (x *Person) GetVehicleAttribute() *VehicleAttribute {
	if x != nil {
		return x.VehicleAttribute
	}
	return nil
}

func (x *Person) GetBusAttribute() *BusAttribute {
	if x != nil {
		return x.BusAttribute
	}
	return nil
}

func (x *Person) GetBikeAttribute() *BikeAttribute {
	if x != nil {
		return x.BikeAttribute
	}
	return nil
}

func (x *Person) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Person) GetProfile() *PersonProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

// 智能体集合，对应一个智能体pb文件或一个智能体mongodb collection
// Agent collection, corresponding to an agent pb file or an agent mongodb collection
type Persons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty" db:"persons" yaml:"persons" bson:"persons"`
}

func (x *Persons) Reset() {
	*x = Persons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Persons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Persons) ProtoMessage() {}

func (x *Persons) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Persons.ProtoReflect.Descriptor instead.
func (*Persons) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_proto_rawDescGZIP(), []int{6}
}

func (x *Persons) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

var File_city_person_v1_person_proto protoreflect.FileDescriptor

var file_city_person_v1_person_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63,
	0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x2f,
	0x76, 0x32, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x02,
	0x0a, 0x0f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x6d, 0x61, 0x78, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x41, 0x63, 0x63, 0x65, 0x6c,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x6d, 0x61, 0x78, 0x5f, 0x62,
	0x72, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x42, 0x72,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x6c,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x75,
	0x73, 0x75, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3c, 0x0a, 0x1a, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x5f, 0x62, 0x72, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x18, 0x75, 0x73, 0x75, 0x61, 0x6c, 0x42, 0x72, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59,
	0x0a, 0x10, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10,
	0x6c, 0x61, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x47, 0x61, 0x70, 0x22, 0x43, 0x0a, 0x0c, 0x42, 0x75, 0x73,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x0f,
	0x0a, 0x0d, 0x42, 0x69, 0x6b, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22,
	0xc9, 0x01, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9b, 0x05, 0x0a, 0x06,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x68, 0x6f, 0x6d, 0x65,
	0x12, 0x34, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x11, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x48, 0x00, 0x52, 0x10, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x0d, 0x62, 0x75,
	0x73, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x48,
	0x01, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x49, 0x0a, 0x0e, 0x62, 0x69, 0x6b, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6b, 0x65,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x48, 0x02, 0x52, 0x0d, 0x62, 0x69, 0x6b,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x03, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x62, 0x75, 0x73,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x62,
	0x69, 0x6b, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x3b, 0x0a, 0x07, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2a, 0xdc, 0x01, 0x0a, 0x09, 0x45, 0x64, 0x75, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x4f, 0x43,
	0x54, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x45,
	0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x43, 0x48, 0x45, 0x4c, 0x4f,
	0x52, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x10, 0x04, 0x12, 0x20,
	0x0a, 0x1c, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4a, 0x55, 0x4e, 0x49,
	0x4f, 0x52, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x10, 0x05,
	0x12, 0x1c, 0x0a, 0x18, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52,
	0x49, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x10, 0x06, 0x12, 0x15,
	0x0a, 0x11, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4c, 0x4c,
	0x45, 0x47, 0x45, 0x10, 0x07, 0x2a, 0x44, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x12, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45,
	0x52, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44,
	0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x2a, 0xae, 0x01, 0x0a, 0x0b,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x17, 0x43,
	0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x53,
	0x55, 0x4d, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x1e, 0x0a,
	0x1a, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4c,
	0x41, 0x54, 0x49, 0x56, 0x45, 0x4c, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x16, 0x0a,
	0x12, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x44,
	0x49, 0x55, 0x4d, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x56, 0x45, 0x4c, 0x59, 0x5f,
	0x48, 0x49, 0x47, 0x48, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d,
	0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x05, 0x42, 0xb1, 0x01, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x42, 0x0b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e,
	0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x50, 0x58, 0xaa, 0x02,
	0x0e, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1a, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10,
	0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_person_v1_person_proto_rawDescOnce sync.Once
	file_city_person_v1_person_proto_rawDescData = file_city_person_v1_person_proto_rawDesc
)

func file_city_person_v1_person_proto_rawDescGZIP() []byte {
	file_city_person_v1_person_proto_rawDescOnce.Do(func() {
		file_city_person_v1_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_person_v1_person_proto_rawDescData)
	})
	return file_city_person_v1_person_proto_rawDescData
}

var file_city_person_v1_person_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_city_person_v1_person_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_city_person_v1_person_proto_goTypes = []interface{}{
	(Education)(0),           // 0: city.person.v1.Education
	(Gender)(0),              // 1: city.person.v1.Gender
	(Consumption)(0),         // 2: city.person.v1.Consumption
	(*PersonAttribute)(nil),  // 3: city.person.v1.PersonAttribute
	(*VehicleAttribute)(nil), // 4: city.person.v1.VehicleAttribute
	(*BusAttribute)(nil),     // 5: city.person.v1.BusAttribute
	(*BikeAttribute)(nil),    // 6: city.person.v1.BikeAttribute
	(*PersonProfile)(nil),    // 7: city.person.v1.PersonProfile
	(*Person)(nil),           // 8: city.person.v1.Person
	(*Persons)(nil),          // 9: city.person.v1.Persons
	nil,                      // 10: city.person.v1.Person.LabelsEntry
	(*v2.Position)(nil),      // 11: city.geo.v2.Position
	(*v21.Schedule)(nil),     // 12: city.trip.v2.Schedule
}
var file_city_person_v1_person_proto_depIdxs = []int32{
	0,  // 0: city.person.v1.PersonProfile.education:type_name -> city.person.v1.Education
	1,  // 1: city.person.v1.PersonProfile.gender:type_name -> city.person.v1.Gender
	2,  // 2: city.person.v1.PersonProfile.consumption:type_name -> city.person.v1.Consumption
	3,  // 3: city.person.v1.Person.attribute:type_name -> city.person.v1.PersonAttribute
	11, // 4: city.person.v1.Person.home:type_name -> city.geo.v2.Position
	12, // 5: city.person.v1.Person.schedules:type_name -> city.trip.v2.Schedule
	4,  // 6: city.person.v1.Person.vehicle_attribute:type_name -> city.person.v1.VehicleAttribute
	5,  // 7: city.person.v1.Person.bus_attribute:type_name -> city.person.v1.BusAttribute
	6,  // 8: city.person.v1.Person.bike_attribute:type_name -> city.person.v1.BikeAttribute
	10, // 9: city.person.v1.Person.labels:type_name -> city.person.v1.Person.LabelsEntry
	7,  // 10: city.person.v1.Person.profile:type_name -> city.person.v1.PersonProfile
	8,  // 11: city.person.v1.Persons.persons:type_name -> city.person.v1.Person
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_city_person_v1_person_proto_init() }
func file_city_person_v1_person_proto_init() {
	if File_city_person_v1_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_person_v1_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonAttribute); i {
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
		file_city_person_v1_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleAttribute); i {
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
		file_city_person_v1_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusAttribute); i {
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
		file_city_person_v1_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BikeAttribute); i {
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
		file_city_person_v1_person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonProfile); i {
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
		file_city_person_v1_person_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_city_person_v1_person_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Persons); i {
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
	file_city_person_v1_person_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_person_v1_person_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_person_v1_person_proto_goTypes,
		DependencyIndexes: file_city_person_v1_person_proto_depIdxs,
		EnumInfos:         file_city_person_v1_person_proto_enumTypes,
		MessageInfos:      file_city_person_v1_person_proto_msgTypes,
	}.Build()
	File_city_person_v1_person_proto = out.File
	file_city_person_v1_person_proto_rawDesc = nil
	file_city_person_v1_person_proto_goTypes = nil
	file_city_person_v1_person_proto_depIdxs = nil
}
