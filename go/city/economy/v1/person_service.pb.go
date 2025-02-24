// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: city/economy/v1/person_service.proto

package economyv1

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

// 批量查询人的经济情况请求
type GetPersonRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 待查询的人的ID列表（为空时查询所有人）
	PersonIds     []int32 `protobuf:"varint,1,rep,packed,name=person_ids,json=personIds,proto3" json:"person_ids,omitempty" bson:"person_ids" db:"person_ids" yaml:"person_ids"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPersonRequest) Reset() {
	*x = GetPersonRequest{}
	mi := &file_city_economy_v1_person_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRequest) ProtoMessage() {}

func (x *GetPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_person_service_proto_msgTypes[0]
	if x != nil {
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
	return file_city_economy_v1_person_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPersonRequest) GetPersonIds() []int32 {
	if x != nil {
		return x.PersonIds
	}
	return nil
}

// 批量查询组织的经济情况响应
type GetPersonResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 人的经济情况
	Persons       []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty" bson:"persons" db:"persons" yaml:"persons"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPersonResponse) Reset() {
	*x = GetPersonResponse{}
	mi := &file_city_economy_v1_person_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonResponse) ProtoMessage() {}

func (x *GetPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_person_service_proto_msgTypes[1]
	if x != nil {
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
	return file_city_economy_v1_person_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPersonResponse) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

// 批量修改人的资金请求
type UpdatePersonMoneyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 待修改的人员资金变动
	Items         []*UpdatePersonMoneyRequestItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" bson:"items" db:"items" yaml:"items"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePersonMoneyRequest) Reset() {
	*x = UpdatePersonMoneyRequest{}
	mi := &file_city_economy_v1_person_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePersonMoneyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersonMoneyRequest) ProtoMessage() {}

func (x *UpdatePersonMoneyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_person_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersonMoneyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePersonMoneyRequest) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_person_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePersonMoneyRequest) GetItems() []*UpdatePersonMoneyRequestItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// 人员资金变动
type UpdatePersonMoneyRequestItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 待修改的人员id
	PersonId int32 `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty" bson:"person_id" db:"person_id" yaml:"person_id"`
	// 资金变动（正数表示增加，负数表示减少）
	Money         float64 `protobuf:"fixed64,2,opt,name=money,proto3" json:"money,omitempty" bson:"money" db:"money" yaml:"money"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePersonMoneyRequestItem) Reset() {
	*x = UpdatePersonMoneyRequestItem{}
	mi := &file_city_economy_v1_person_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePersonMoneyRequestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersonMoneyRequestItem) ProtoMessage() {}

func (x *UpdatePersonMoneyRequestItem) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_person_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersonMoneyRequestItem.ProtoReflect.Descriptor instead.
func (*UpdatePersonMoneyRequestItem) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_person_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePersonMoneyRequestItem) GetPersonId() int32 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *UpdatePersonMoneyRequestItem) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

// 批量修改人的资金响应
type UpdatePersonMoneyResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 修改后的人的经济情况
	Persons       []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty" bson:"persons" db:"persons" yaml:"persons"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePersonMoneyResponse) Reset() {
	*x = UpdatePersonMoneyResponse{}
	mi := &file_city_economy_v1_person_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePersonMoneyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersonMoneyResponse) ProtoMessage() {}

func (x *UpdatePersonMoneyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_person_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersonMoneyResponse.ProtoReflect.Descriptor instead.
func (*UpdatePersonMoneyResponse) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_person_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePersonMoneyResponse) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

var File_city_economy_v1_person_service_proto protoreflect.FileDescriptor

var file_city_economy_v1_person_service_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x22, 0x46, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x22, 0x5f, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x51, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x4e, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x73, 0x32, 0xd3, 0x01, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x29, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xc2, 0x01, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x2e, 0x76, 0x31, 0x42, 0x12, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x2e, 0x66,
	0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f,
	0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x45, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x69, 0x74,
	0x79, 0x2e, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43,
	0x69, 0x74, 0x79, 0x5c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_city_economy_v1_person_service_proto_rawDescOnce sync.Once
	file_city_economy_v1_person_service_proto_rawDescData []byte
)

func file_city_economy_v1_person_service_proto_rawDescGZIP() []byte {
	file_city_economy_v1_person_service_proto_rawDescOnce.Do(func() {
		file_city_economy_v1_person_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_city_economy_v1_person_service_proto_rawDesc), len(file_city_economy_v1_person_service_proto_rawDesc)))
	})
	return file_city_economy_v1_person_service_proto_rawDescData
}

var file_city_economy_v1_person_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_city_economy_v1_person_service_proto_goTypes = []any{
	(*GetPersonRequest)(nil),             // 0: city.economy.v1.GetPersonRequest
	(*GetPersonResponse)(nil),            // 1: city.economy.v1.GetPersonResponse
	(*UpdatePersonMoneyRequest)(nil),     // 2: city.economy.v1.UpdatePersonMoneyRequest
	(*UpdatePersonMoneyRequestItem)(nil), // 3: city.economy.v1.UpdatePersonMoneyRequestItem
	(*UpdatePersonMoneyResponse)(nil),    // 4: city.economy.v1.UpdatePersonMoneyResponse
	(*Person)(nil),                       // 5: city.economy.v1.Person
}
var file_city_economy_v1_person_service_proto_depIdxs = []int32{
	5, // 0: city.economy.v1.GetPersonResponse.persons:type_name -> city.economy.v1.Person
	3, // 1: city.economy.v1.UpdatePersonMoneyRequest.items:type_name -> city.economy.v1.UpdatePersonMoneyRequestItem
	5, // 2: city.economy.v1.UpdatePersonMoneyResponse.persons:type_name -> city.economy.v1.Person
	0, // 3: city.economy.v1.PersonService.GetPerson:input_type -> city.economy.v1.GetPersonRequest
	2, // 4: city.economy.v1.PersonService.UpdatePersonMoney:input_type -> city.economy.v1.UpdatePersonMoneyRequest
	1, // 5: city.economy.v1.PersonService.GetPerson:output_type -> city.economy.v1.GetPersonResponse
	4, // 6: city.economy.v1.PersonService.UpdatePersonMoney:output_type -> city.economy.v1.UpdatePersonMoneyResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_city_economy_v1_person_service_proto_init() }
func file_city_economy_v1_person_service_proto_init() {
	if File_city_economy_v1_person_service_proto != nil {
		return
	}
	file_city_economy_v1_economy_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_city_economy_v1_person_service_proto_rawDesc), len(file_city_economy_v1_person_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_economy_v1_person_service_proto_goTypes,
		DependencyIndexes: file_city_economy_v1_person_service_proto_depIdxs,
		MessageInfos:      file_city_economy_v1_person_service_proto_msgTypes,
	}.Build()
	File_city_economy_v1_person_service_proto = out.File
	file_city_economy_v1_person_service_proto_goTypes = nil
	file_city_economy_v1_person_service_proto_depIdxs = nil
}
