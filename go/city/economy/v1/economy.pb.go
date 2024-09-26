// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: city/economy/v1/economy.proto

package economyv1

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

// 个人（与Person一一对应，
// 没有绑定到city.economy.v1.Person的Person将无法参与经济模拟）
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 与person id一致
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id" bson:"id" db:"id"`
	// 起始资金
	Money float64 `protobuf:"fixed64,2,opt,name=money,proto3" json:"money,omitempty" yaml:"money" bson:"money" db:"money"`
	// 所在组织ID（初始化时不提供，由组织中的员工列表设定）
	OrgId *int32 `protobuf:"varint,3,opt,name=org_id,json=orgId,proto3,oneof" json:"org_id,omitempty" db:"org_id" yaml:"org_id" bson:"org_id"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_economy_v1_economy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_economy_proto_msgTypes[0]
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
	return file_city_economy_v1_economy_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *Person) GetOrgId() int32 {
	if x != nil && x.OrgId != nil {
		return *x.OrgId
	}
	return 0
}

// 员工
type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 员工ID
	PersonId int32 `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty" yaml:"person_id" bson:"person_id" db:"person_id"`
	// 薪水
	Salary float64 `protobuf:"fixed64,2,opt,name=salary,proto3" json:"salary,omitempty" yaml:"salary" bson:"salary" db:"salary"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_economy_v1_economy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_economy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_economy_proto_rawDescGZIP(), []int{1}
}

func (x *Employee) GetPersonId() int32 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *Employee) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

// 岗位
type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 岗位名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name" bson:"name" db:"name"`
	// 岗位所需员工数量
	EmployeeCount int32 `protobuf:"varint,2,opt,name=employee_count,json=employeeCount,proto3" json:"employee_count,omitempty" yaml:"employee_count" bson:"employee_count" db:"employee_count"`
	// 岗位薪水
	Salary *float64 `protobuf:"fixed64,3,opt,name=salary,proto3,oneof" json:"salary,omitempty" yaml:"salary" bson:"salary" db:"salary"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_economy_v1_economy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_economy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_economy_proto_rawDescGZIP(), []int{2}
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetEmployeeCount() int32 {
	if x != nil {
		return x.EmployeeCount
	}
	return 0
}

func (x *Job) GetSalary() float64 {
	if x != nil && x.Salary != nil {
		return *x.Salary
	}
	return 0
}

// 货物
type Goods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 货物类型
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" db:"type" yaml:"type" bson:"type"`
	// 货物名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name" bson:"name" db:"name"`
	// 货物数量
	Count int32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty" bson:"count" db:"count" yaml:"count"`
	// 货物价格（允许暂未定价）
	Price *float64 `protobuf:"fixed64,4,opt,name=price,proto3,oneof" json:"price,omitempty" yaml:"price" bson:"price" db:"price"`
}

func (x *Goods) Reset() {
	*x = Goods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_economy_v1_economy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Goods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Goods) ProtoMessage() {}

func (x *Goods) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_economy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Goods.ProtoReflect.Descriptor instead.
func (*Goods) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_economy_proto_rawDescGZIP(), []int{3}
}

func (x *Goods) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Goods) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Goods) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Goods) GetPrice() float64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

// 组织，具有员工、货物、资金等属性
type Org struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组织ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id" db:"id" yaml:"id"`
	// 组织所在的POI ID
	PoiId int32 `protobuf:"varint,2,opt,name=poi_id,json=poiId,proto3" json:"poi_id,omitempty" yaml:"poi_id" bson:"poi_id" db:"poi_id"`
	// 员工列表（初始化时由Orgs列表提供，单向绑定到person上）
	Employees []*Employee `protobuf:"bytes,3,rep,name=employees,proto3" json:"employees,omitempty" db:"employees" yaml:"employees" bson:"employees"`
	// 岗位列表
	Jobs []*Job `protobuf:"bytes,4,rep,name=jobs,proto3" json:"jobs,omitempty" yaml:"jobs" bson:"jobs" db:"jobs"`
	// 资金
	Money float64 `protobuf:"fixed64,5,opt,name=money,proto3" json:"money,omitempty" yaml:"money" bson:"money" db:"money"`
	// 货物
	Goods []*Goods `protobuf:"bytes,6,rep,name=goods,proto3" json:"goods,omitempty" db:"goods" yaml:"goods" bson:"goods"`
	// 功能列表
	// buy: 购买货物
	// apply: 申请岗位
	// ...
	Functions []string `protobuf:"bytes,7,rep,name=functions,proto3" json:"functions,omitempty" yaml:"functions" bson:"functions" db:"functions"`
}

func (x *Org) Reset() {
	*x = Org{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_economy_v1_economy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Org) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Org) ProtoMessage() {}

func (x *Org) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_economy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Org.ProtoReflect.Descriptor instead.
func (*Org) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_economy_proto_rawDescGZIP(), []int{4}
}

func (x *Org) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Org) GetPoiId() int32 {
	if x != nil {
		return x.PoiId
	}
	return 0
}

func (x *Org) GetEmployees() []*Employee {
	if x != nil {
		return x.Employees
	}
	return nil
}

func (x *Org) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *Org) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *Org) GetGoods() []*Goods {
	if x != nil {
		return x.Goods
	}
	return nil
}

func (x *Org) GetFunctions() []string {
	if x != nil {
		return x.Functions
	}
	return nil
}

// 组织列表，对应MongoDB中的集合
type Economy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 人
	Persons []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty" db:"persons" yaml:"persons" bson:"persons"`
	// 组织列表
	Orgs []*Org `protobuf:"bytes,2,rep,name=orgs,proto3" json:"orgs,omitempty" yaml:"orgs" bson:"orgs" db:"orgs"`
}

func (x *Economy) Reset() {
	*x = Economy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_economy_v1_economy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Economy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Economy) ProtoMessage() {}

func (x *Economy) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_economy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Economy.ProtoReflect.Descriptor instead.
func (*Economy) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_economy_proto_rawDescGZIP(), []int{5}
}

func (x *Economy) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

func (x *Economy) GetOrgs() []*Org {
	if x != nil {
		return x.Orgs
	}
	return nil
}

var File_city_economy_v1_economy_proto protoreflect.FileDescriptor

var file_city_economy_v1_economy_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31,
	0x22, 0x55, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x22, 0x68, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x22, 0x6a, 0x0a, 0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0xf1,
	0x01, 0x0a, 0x03, 0x4f, 0x72, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x49, 0x64, 0x12, 0x37, 0x0a,
	0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x09, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x66, 0x0a, 0x07, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x12, 0x31, 0x0a,
	0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x12, 0x28, 0x0a, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x67, 0x52, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x42, 0xb9, 0x01, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e,
	0x65, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2f, 0x76, 0x31,
	0x3b, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x45, 0x58,
	0xaa, 0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_economy_v1_economy_proto_rawDescOnce sync.Once
	file_city_economy_v1_economy_proto_rawDescData = file_city_economy_v1_economy_proto_rawDesc
)

func file_city_economy_v1_economy_proto_rawDescGZIP() []byte {
	file_city_economy_v1_economy_proto_rawDescOnce.Do(func() {
		file_city_economy_v1_economy_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_economy_v1_economy_proto_rawDescData)
	})
	return file_city_economy_v1_economy_proto_rawDescData
}

var file_city_economy_v1_economy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_city_economy_v1_economy_proto_goTypes = []any{
	(*Person)(nil),   // 0: city.economy.v1.Person
	(*Employee)(nil), // 1: city.economy.v1.Employee
	(*Job)(nil),      // 2: city.economy.v1.Job
	(*Goods)(nil),    // 3: city.economy.v1.Goods
	(*Org)(nil),      // 4: city.economy.v1.Org
	(*Economy)(nil),  // 5: city.economy.v1.Economy
}
var file_city_economy_v1_economy_proto_depIdxs = []int32{
	1, // 0: city.economy.v1.Org.employees:type_name -> city.economy.v1.Employee
	2, // 1: city.economy.v1.Org.jobs:type_name -> city.economy.v1.Job
	3, // 2: city.economy.v1.Org.goods:type_name -> city.economy.v1.Goods
	0, // 3: city.economy.v1.Economy.persons:type_name -> city.economy.v1.Person
	4, // 4: city.economy.v1.Economy.orgs:type_name -> city.economy.v1.Org
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_city_economy_v1_economy_proto_init() }
func file_city_economy_v1_economy_proto_init() {
	if File_city_economy_v1_economy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_economy_v1_economy_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_city_economy_v1_economy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Employee); i {
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
		file_city_economy_v1_economy_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Job); i {
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
		file_city_economy_v1_economy_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Goods); i {
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
		file_city_economy_v1_economy_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Org); i {
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
		file_city_economy_v1_economy_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Economy); i {
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
	file_city_economy_v1_economy_proto_msgTypes[0].OneofWrappers = []any{}
	file_city_economy_v1_economy_proto_msgTypes[2].OneofWrappers = []any{}
	file_city_economy_v1_economy_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_economy_v1_economy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_economy_v1_economy_proto_goTypes,
		DependencyIndexes: file_city_economy_v1_economy_proto_depIdxs,
		MessageInfos:      file_city_economy_v1_economy_proto_msgTypes,
	}.Build()
	File_city_economy_v1_economy_proto = out.File
	file_city_economy_v1_economy_proto_rawDesc = nil
	file_city_economy_v1_economy_proto_goTypes = nil
	file_city_economy_v1_economy_proto_depIdxs = nil
}
