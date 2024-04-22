// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: city/person/v1/person_service.proto

package personv1

import (
	v21 "git.fiblab.net/sim/protos/go/city/geo/v2"
	v2 "git.fiblab.net/sim/protos/go/city/trip/v2"
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

// 获取person信息请求
// Request for getting person information
type GetPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// person id
	PersonId int32 `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty" bson:"person_id" db:"person_id" yaml:"person_id"`
}

func (x *GetPersonRequest) Reset() {
	*x = GetPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRequest) ProtoMessage() {}

func (x *GetPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRequest.ProtoReflect.Descriptor instead.
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPersonRequest) GetPersonId() int32 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

// 获取person信息响应
// Response of getting person information
type GetPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// person信息
	// person information
	Base *Person `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty" yaml:"base" bson:"base" db:"base"`
	// person运动信息
	// person motion information
	Motion *PersonMotion `protobuf:"bytes,2,opt,name=motion,proto3" json:"motion,omitempty" db:"motion" yaml:"motion" bson:"motion"`
}

func (x *GetPersonResponse) Reset() {
	*x = GetPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonResponse) ProtoMessage() {}

func (x *GetPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonResponse.ProtoReflect.Descriptor instead.
func (*GetPersonResponse) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPersonResponse) GetBase() *Person {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *GetPersonResponse) GetMotion() *PersonMotion {
	if x != nil {
		return x.Motion
	}
	return nil
}

// 新增person请求
// Request for adding a new person
type AddPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 约定：person中不设置id
	// Convention: personid is not set here
	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty" yaml:"person" bson:"person" db:"person"`
}

func (x *AddPersonRequest) Reset() {
	*x = AddPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPersonRequest) ProtoMessage() {}

func (x *AddPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPersonRequest.ProtoReflect.Descriptor instead.
func (*AddPersonRequest) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddPersonRequest) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

// 新增person响应
// Response of adding a new person
type AddPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 新增的person分配得到的ID
	// The ID assigned to the newly added person
	PersonId int32 `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty" db:"person_id" yaml:"person_id" bson:"person_id"`
}

func (x *AddPersonResponse) Reset() {
	*x = AddPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPersonResponse) ProtoMessage() {}

func (x *AddPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPersonResponse.ProtoReflect.Descriptor instead.
func (*AddPersonResponse) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddPersonResponse) GetPersonId() int32 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

// 修改person的schedule请求
// Request for setting person schedule
type SetScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// person id
	PersonId int32 `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty" yaml:"person_id" bson:"person_id" db:"person_id"`
	// 新的schedule（覆盖原有的schedule）
	// New schedule (overwrites the original schedule)
	Schedules []*v2.Schedule `protobuf:"bytes,2,rep,name=schedules,proto3" json:"schedules,omitempty" yaml:"schedules" bson:"schedules" db:"schedules"`
}

func (x *SetScheduleRequest) Reset() {
	*x = SetScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScheduleRequest) ProtoMessage() {}

func (x *SetScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScheduleRequest.ProtoReflect.Descriptor instead.
func (*SetScheduleRequest) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_service_proto_rawDescGZIP(), []int{4}
}

func (x *SetScheduleRequest) GetPersonId() int32 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *SetScheduleRequest) GetSchedules() []*v2.Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

// 修改person的schedule响应
// Response of setting person schedule
type SetScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetScheduleResponse) Reset() {
	*x = SetScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScheduleResponse) ProtoMessage() {}

func (x *SetScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScheduleResponse.ProtoReflect.Descriptor instead.
func (*SetScheduleResponse) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_service_proto_rawDescGZIP(), []int{5}
}

// 获取特定区域内的person请求
// Request for getting persons in region
type GetPersonByLongLatBBoxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 经纬度范围
	// longitude and latitude bounding box
	Bbox *v21.LongLatBBox `protobuf:"bytes,1,opt,name=bbox,proto3" json:"bbox,omitempty" yaml:"bbox" bson:"bbox" db:"bbox"`
	// 过滤人的状态（状态为列表内的值的人不返回）
	// Filter person's status (person whose status is in the list will not be returned)
	ExcludeStatuses []Status `protobuf:"varint,2,rep,packed,name=exclude_statuses,json=excludeStatuses,proto3,enum=city.person.v1.Status" json:"exclude_statuses,omitempty" yaml:"exclude_statuses" bson:"exclude_statuses" db:"exclude_statuses"`
}

func (x *GetPersonByLongLatBBoxRequest) Reset() {
	*x = GetPersonByLongLatBBoxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonByLongLatBBoxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonByLongLatBBoxRequest) ProtoMessage() {}

func (x *GetPersonByLongLatBBoxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonByLongLatBBoxRequest.ProtoReflect.Descriptor instead.
func (*GetPersonByLongLatBBoxRequest) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetPersonByLongLatBBoxRequest) GetBbox() *v21.LongLatBBox {
	if x != nil {
		return x.Bbox
	}
	return nil
}

func (x *GetPersonByLongLatBBoxRequest) GetExcludeStatuses() []Status {
	if x != nil {
		return x.ExcludeStatuses
	}
	return nil
}

// 获取特定区域内的person响应
// Response of getting persons in region
type GetPersonByLongLatBBoxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 区域内的person的运动信息
	// motion status of persons in the region
	Motions []*PersonMotion `protobuf:"bytes,1,rep,name=motions,proto3" json:"motions,omitempty" db:"motions" yaml:"motions" bson:"motions"`
}

func (x *GetPersonByLongLatBBoxResponse) Reset() {
	*x = GetPersonByLongLatBBoxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_person_v1_person_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonByLongLatBBoxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonByLongLatBBoxResponse) ProtoMessage() {}

func (x *GetPersonByLongLatBBoxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_person_v1_person_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonByLongLatBBoxResponse.ProtoReflect.Descriptor instead.
func (*GetPersonByLongLatBBoxResponse) Descriptor() ([]byte, []int) {
	return file_city_person_v1_person_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetPersonByLongLatBBoxResponse) GetMotions() []*PersonMotion {
	if x != nil {
		return x.Motions
	}
	return nil
}

var File_city_person_v1_person_service_proto protoreflect.FileDescriptor

var file_city_person_v1_person_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x2f,
	0x76, 0x32, 0x2f, 0x67, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x69,
	0x74, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x69, 0x74, 0x79, 0x2f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x69,
	0x70, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x75, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x67, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x90, 0x01,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x4c, 0x6f, 0x6e,
	0x67, 0x4c, 0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x04, 0x62, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x6e, 0x67,
	0x4c, 0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x04, 0x62, 0x62, 0x6f, 0x78, 0x12, 0x41, 0x0a,
	0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x22, 0x58, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x4c,
	0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x84, 0x03, 0x0a, 0x0d, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x56, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x22, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x4c, 0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x42, 0x42,
	0x6f, 0x78, 0x12, 0x2d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x4c,
	0x6f, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x4c, 0x6f,
	0x6e, 0x67, 0x4c, 0x61, 0x74, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xb8, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73,
	0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74,
	0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x50, 0x58, 0xaa, 0x02, 0x0e, 0x43, 0x69, 0x74,
	0x79, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x43, 0x69,
	0x74, 0x79, 0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x43,
	0x69, 0x74, 0x79, 0x5c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x69, 0x74, 0x79,
	0x3a, 0x3a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_person_v1_person_service_proto_rawDescOnce sync.Once
	file_city_person_v1_person_service_proto_rawDescData = file_city_person_v1_person_service_proto_rawDesc
)

func file_city_person_v1_person_service_proto_rawDescGZIP() []byte {
	file_city_person_v1_person_service_proto_rawDescOnce.Do(func() {
		file_city_person_v1_person_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_person_v1_person_service_proto_rawDescData)
	})
	return file_city_person_v1_person_service_proto_rawDescData
}

var file_city_person_v1_person_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_city_person_v1_person_service_proto_goTypes = []interface{}{
	(*GetPersonRequest)(nil),               // 0: city.person.v1.GetPersonRequest
	(*GetPersonResponse)(nil),              // 1: city.person.v1.GetPersonResponse
	(*AddPersonRequest)(nil),               // 2: city.person.v1.AddPersonRequest
	(*AddPersonResponse)(nil),              // 3: city.person.v1.AddPersonResponse
	(*SetScheduleRequest)(nil),             // 4: city.person.v1.SetScheduleRequest
	(*SetScheduleResponse)(nil),            // 5: city.person.v1.SetScheduleResponse
	(*GetPersonByLongLatBBoxRequest)(nil),  // 6: city.person.v1.GetPersonByLongLatBBoxRequest
	(*GetPersonByLongLatBBoxResponse)(nil), // 7: city.person.v1.GetPersonByLongLatBBoxResponse
	(*Person)(nil),                         // 8: city.person.v1.Person
	(*PersonMotion)(nil),                   // 9: city.person.v1.PersonMotion
	(*v2.Schedule)(nil),                    // 10: city.trip.v2.Schedule
	(*v21.LongLatBBox)(nil),                // 11: city.geo.v2.LongLatBBox
	(Status)(0),                            // 12: city.person.v1.Status
}
var file_city_person_v1_person_service_proto_depIdxs = []int32{
	8,  // 0: city.person.v1.GetPersonResponse.base:type_name -> city.person.v1.Person
	9,  // 1: city.person.v1.GetPersonResponse.motion:type_name -> city.person.v1.PersonMotion
	8,  // 2: city.person.v1.AddPersonRequest.person:type_name -> city.person.v1.Person
	10, // 3: city.person.v1.SetScheduleRequest.schedules:type_name -> city.trip.v2.Schedule
	11, // 4: city.person.v1.GetPersonByLongLatBBoxRequest.bbox:type_name -> city.geo.v2.LongLatBBox
	12, // 5: city.person.v1.GetPersonByLongLatBBoxRequest.exclude_statuses:type_name -> city.person.v1.Status
	9,  // 6: city.person.v1.GetPersonByLongLatBBoxResponse.motions:type_name -> city.person.v1.PersonMotion
	0,  // 7: city.person.v1.PersonService.GetPerson:input_type -> city.person.v1.GetPersonRequest
	2,  // 8: city.person.v1.PersonService.AddPerson:input_type -> city.person.v1.AddPersonRequest
	4,  // 9: city.person.v1.PersonService.SetSchedule:input_type -> city.person.v1.SetScheduleRequest
	6,  // 10: city.person.v1.PersonService.GetPersonByLongLatBBox:input_type -> city.person.v1.GetPersonByLongLatBBoxRequest
	1,  // 11: city.person.v1.PersonService.GetPerson:output_type -> city.person.v1.GetPersonResponse
	3,  // 12: city.person.v1.PersonService.AddPerson:output_type -> city.person.v1.AddPersonResponse
	5,  // 13: city.person.v1.PersonService.SetSchedule:output_type -> city.person.v1.SetScheduleResponse
	7,  // 14: city.person.v1.PersonService.GetPersonByLongLatBBox:output_type -> city.person.v1.GetPersonByLongLatBBoxResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_city_person_v1_person_service_proto_init() }
func file_city_person_v1_person_service_proto_init() {
	if File_city_person_v1_person_service_proto != nil {
		return
	}
	file_city_person_v1_motion_proto_init()
	file_city_person_v1_person_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_person_v1_person_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonRequest); i {
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
		file_city_person_v1_person_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonResponse); i {
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
		file_city_person_v1_person_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPersonRequest); i {
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
		file_city_person_v1_person_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPersonResponse); i {
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
		file_city_person_v1_person_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScheduleRequest); i {
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
		file_city_person_v1_person_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScheduleResponse); i {
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
		file_city_person_v1_person_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonByLongLatBBoxRequest); i {
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
		file_city_person_v1_person_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonByLongLatBBoxResponse); i {
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
			RawDescriptor: file_city_person_v1_person_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_person_v1_person_service_proto_goTypes,
		DependencyIndexes: file_city_person_v1_person_service_proto_depIdxs,
		MessageInfos:      file_city_person_v1_person_service_proto_msgTypes,
	}.Build()
	File_city_person_v1_person_service_proto = out.File
	file_city_person_v1_person_service_proto_rawDesc = nil
	file_city_person_v1_person_service_proto_goTypes = nil
	file_city_person_v1_person_service_proto_depIdxs = nil
}
