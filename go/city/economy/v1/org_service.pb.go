// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: city/economy/v1/org_service.proto

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

// 批量查询组织的经济情况请求
type GetOrgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待查询的组织的ID列表（为空时查询所有组织）
	OrgIds []int32 `protobuf:"varint,1,rep,packed,name=org_ids,json=orgIds,proto3" json:"org_ids,omitempty" yaml:"org_ids" bson:"org_ids" db:"org_ids"`
}

func (x *GetOrgRequest) Reset() {
	*x = GetOrgRequest{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrgRequest) ProtoMessage() {}

func (x *GetOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrgRequest.ProtoReflect.Descriptor instead.
func (*GetOrgRequest) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrgRequest) GetOrgIds() []int32 {
	if x != nil {
		return x.OrgIds
	}
	return nil
}

// 批量查询组织的经济情况响应
type GetOrgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组织的经济情况
	Orgs []*Org `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty" yaml:"orgs" bson:"orgs" db:"orgs"`
}

func (x *GetOrgResponse) Reset() {
	*x = GetOrgResponse{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrgResponse) ProtoMessage() {}

func (x *GetOrgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrgResponse.ProtoReflect.Descriptor instead.
func (*GetOrgResponse) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrgResponse) GetOrgs() []*Org {
	if x != nil {
		return x.Orgs
	}
	return nil
}

// 批量修改组织的资金请求
type UpdateOrgMoneyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待修改的组织资金变动
	Items []*UpdateOrgMoneyRequestItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" bson:"items" db:"items" yaml:"items"`
}

func (x *UpdateOrgMoneyRequest) Reset() {
	*x = UpdateOrgMoneyRequest{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgMoneyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgMoneyRequest) ProtoMessage() {}

func (x *UpdateOrgMoneyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgMoneyRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrgMoneyRequest) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateOrgMoneyRequest) GetItems() []*UpdateOrgMoneyRequestItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// 组织资金变动
type UpdateOrgMoneyRequestItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待修改的组织
	OrgId int32 `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" yaml:"org_id" bson:"org_id" db:"org_id"`
	// 正数表示增加，负数表示减少
	Money float64 `protobuf:"fixed64,2,opt,name=money,proto3" json:"money,omitempty" yaml:"money" bson:"money" db:"money"`
}

func (x *UpdateOrgMoneyRequestItem) Reset() {
	*x = UpdateOrgMoneyRequestItem{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgMoneyRequestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgMoneyRequestItem) ProtoMessage() {}

func (x *UpdateOrgMoneyRequestItem) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgMoneyRequestItem.ProtoReflect.Descriptor instead.
func (*UpdateOrgMoneyRequestItem) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOrgMoneyRequestItem) GetOrgId() int32 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *UpdateOrgMoneyRequestItem) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

// 批量修改组织的资金响应
type UpdateOrgMoneyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 修改后的组织的经济情况
	Orgs []*Org `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty" db:"orgs" yaml:"orgs" bson:"orgs"`
}

func (x *UpdateOrgMoneyResponse) Reset() {
	*x = UpdateOrgMoneyResponse{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgMoneyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgMoneyResponse) ProtoMessage() {}

func (x *UpdateOrgMoneyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgMoneyResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrgMoneyResponse) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOrgMoneyResponse) GetOrgs() []*Org {
	if x != nil {
		return x.Orgs
	}
	return nil
}

// 批量修改组织的货物请求
type UpdateOrgGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待修改的组织货物变动
	Items []*UpdateOrgGoodsRequestItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" yaml:"items" bson:"items" db:"items"`
}

func (x *UpdateOrgGoodsRequest) Reset() {
	*x = UpdateOrgGoodsRequest{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgGoodsRequest) ProtoMessage() {}

func (x *UpdateOrgGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgGoodsRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrgGoodsRequest) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateOrgGoodsRequest) GetItems() []*UpdateOrgGoodsRequestItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// 组织货物变动
type UpdateOrgGoodsRequestItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待修改的组织
	OrgId int32 `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" yaml:"org_id" bson:"org_id" db:"org_id"`
	// 货物变动
	// 按照(type, name)相等来判断是否为同一种货物
	// 货物数量为增量，正数表示增加，负数表示减少
	// price如果未设定则沿用原来的价格，否则使用新的价格
	Goods []*Goods `protobuf:"bytes,2,rep,name=goods,proto3" json:"goods,omitempty" yaml:"goods" bson:"goods" db:"goods"`
}

func (x *UpdateOrgGoodsRequestItem) Reset() {
	*x = UpdateOrgGoodsRequestItem{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgGoodsRequestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgGoodsRequestItem) ProtoMessage() {}

func (x *UpdateOrgGoodsRequestItem) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgGoodsRequestItem.ProtoReflect.Descriptor instead.
func (*UpdateOrgGoodsRequestItem) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateOrgGoodsRequestItem) GetOrgId() int32 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *UpdateOrgGoodsRequestItem) GetGoods() []*Goods {
	if x != nil {
		return x.Goods
	}
	return nil
}

// 批量修改组织的货物响应
type UpdateOrgGoodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 修改后的组织的经济情况
	Orgs []*Org `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty" bson:"orgs" db:"orgs" yaml:"orgs"`
}

func (x *UpdateOrgGoodsResponse) Reset() {
	*x = UpdateOrgGoodsResponse{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgGoodsResponse) ProtoMessage() {}

func (x *UpdateOrgGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgGoodsResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrgGoodsResponse) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateOrgGoodsResponse) GetOrgs() []*Org {
	if x != nil {
		return x.Orgs
	}
	return nil
}

// 批量修改组织的员工请求
type UpdateOrgEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待修改的组织员工变动
	Items []*UpdateOrgEmployeeRequestItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" yaml:"items" bson:"items" db:"items"`
}

func (x *UpdateOrgEmployeeRequest) Reset() {
	*x = UpdateOrgEmployeeRequest{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgEmployeeRequest) ProtoMessage() {}

func (x *UpdateOrgEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgEmployeeRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrgEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateOrgEmployeeRequest) GetItems() []*UpdateOrgEmployeeRequestItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// 组织员工变动
type UpdateOrgEmployeeRequestItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待修改的组织
	OrgId int32 `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" yaml:"org_id" bson:"org_id" db:"org_id"`
	// 新增的员工
	Adds []*Employee `protobuf:"bytes,2,rep,name=adds,proto3" json:"adds,omitempty" yaml:"adds" bson:"adds" db:"adds"`
	// 删除的员工
	Dels []int32 `protobuf:"varint,3,rep,packed,name=dels,proto3" json:"dels,omitempty" yaml:"dels" bson:"dels" db:"dels"`
	// 修改薪水的员工
	Updates []*Employee `protobuf:"bytes,4,rep,name=updates,proto3" json:"updates,omitempty" yaml:"updates" bson:"updates" db:"updates"`
}

func (x *UpdateOrgEmployeeRequestItem) Reset() {
	*x = UpdateOrgEmployeeRequestItem{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgEmployeeRequestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgEmployeeRequestItem) ProtoMessage() {}

func (x *UpdateOrgEmployeeRequestItem) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgEmployeeRequestItem.ProtoReflect.Descriptor instead.
func (*UpdateOrgEmployeeRequestItem) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateOrgEmployeeRequestItem) GetOrgId() int32 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *UpdateOrgEmployeeRequestItem) GetAdds() []*Employee {
	if x != nil {
		return x.Adds
	}
	return nil
}

func (x *UpdateOrgEmployeeRequestItem) GetDels() []int32 {
	if x != nil {
		return x.Dels
	}
	return nil
}

func (x *UpdateOrgEmployeeRequestItem) GetUpdates() []*Employee {
	if x != nil {
		return x.Updates
	}
	return nil
}

// 批量修改组织的员工响应
type UpdateOrgEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 修改后的组织的经济情况
	Orgs []*Org `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty" yaml:"orgs" bson:"orgs" db:"orgs"`
}

func (x *UpdateOrgEmployeeResponse) Reset() {
	*x = UpdateOrgEmployeeResponse{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgEmployeeResponse) ProtoMessage() {}

func (x *UpdateOrgEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgEmployeeResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrgEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateOrgEmployeeResponse) GetOrgs() []*Org {
	if x != nil {
		return x.Orgs
	}
	return nil
}

// 批量修改组织的岗位请求
type UpdateOrgJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待修改的组织岗位变动
	Items []*UpdateOrgJobRequestItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" yaml:"items" bson:"items" db:"items"`
}

func (x *UpdateOrgJobRequest) Reset() {
	*x = UpdateOrgJobRequest{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgJobRequest) ProtoMessage() {}

func (x *UpdateOrgJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrgJobRequest) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateOrgJobRequest) GetItems() []*UpdateOrgJobRequestItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// 组织岗位变动
type UpdateOrgJobRequestItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 待修改的组织
	OrgId int32 `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" yaml:"org_id" bson:"org_id" db:"org_id"`
	// 岗位变动
	// 按照name相等来判断是否为同一种岗位
	// 岗位数量为增量，正数表示增加，负数表示减少
	// salary如果未设定则沿用原来的薪水，否则使用新的薪水
	Jobs []*Job `protobuf:"bytes,2,rep,name=jobs,proto3" json:"jobs,omitempty" yaml:"jobs" bson:"jobs" db:"jobs"`
}

func (x *UpdateOrgJobRequestItem) Reset() {
	*x = UpdateOrgJobRequestItem{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgJobRequestItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgJobRequestItem) ProtoMessage() {}

func (x *UpdateOrgJobRequestItem) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgJobRequestItem.ProtoReflect.Descriptor instead.
func (*UpdateOrgJobRequestItem) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateOrgJobRequestItem) GetOrgId() int32 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *UpdateOrgJobRequestItem) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

// 批量修改组织的岗位响应
type UpdateOrgJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 修改后的组织的经济情况
	Orgs []*Org `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty" yaml:"orgs" bson:"orgs" db:"orgs"`
}

func (x *UpdateOrgJobResponse) Reset() {
	*x = UpdateOrgJobResponse{}
	mi := &file_city_economy_v1_org_service_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrgJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrgJobResponse) ProtoMessage() {}

func (x *UpdateOrgJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_economy_v1_org_service_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrgJobResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrgJobResponse) Descriptor() ([]byte, []int) {
	return file_city_economy_v1_org_service_proto_rawDescGZIP(), []int{13}
}

func (x *UpdateOrgJobResponse) GetOrgs() []*Org {
	if x != nil {
		return x.Orgs
	}
	return nil
}

var File_city_economy_v1_org_service_proto protoreflect.FileDescriptor

var file_city_economy_v1_org_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x73, 0x22, 0x3a, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x67, 0x52, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x22, 0x59, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x40, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x48, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x67, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x42,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6f, 0x72, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x04, 0x6f, 0x72,
	0x67, 0x73, 0x22, 0x59, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x60, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x12, 0x2c, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22,
	0x42, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6f, 0x72, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x04, 0x6f,
	0x72, 0x67, 0x73, 0x22, 0x5f, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x43, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0xad, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x67, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x04, 0x61, 0x64, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x64, 0x65, 0x6c, 0x73, 0x12,
	0x33, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x67, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x22, 0x55, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x5a, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x0a,
	0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x40,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x52, 0x04, 0x6f, 0x72, 0x67, 0x73,
	0x32, 0xf0, 0x03, 0x0a, 0x0a, 0x4f, 0x72, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x12, 0x1e, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x26,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x67, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x63, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x67, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x67, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x67, 0x4a, 0x6f, 0x62, 0x12, 0x24, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0xbf, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x4f, 0x72, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x2e, 0x66, 0x69, 0x62, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73,
	0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x69, 0x74, 0x79, 0x2f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x45, 0x58, 0xaa,
	0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0f, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x69, 0x74, 0x79, 0x5c, 0x45, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x11, 0x43, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_economy_v1_org_service_proto_rawDescOnce sync.Once
	file_city_economy_v1_org_service_proto_rawDescData = file_city_economy_v1_org_service_proto_rawDesc
)

func file_city_economy_v1_org_service_proto_rawDescGZIP() []byte {
	file_city_economy_v1_org_service_proto_rawDescOnce.Do(func() {
		file_city_economy_v1_org_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_economy_v1_org_service_proto_rawDescData)
	})
	return file_city_economy_v1_org_service_proto_rawDescData
}

var file_city_economy_v1_org_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_city_economy_v1_org_service_proto_goTypes = []any{
	(*GetOrgRequest)(nil),                // 0: city.economy.v1.GetOrgRequest
	(*GetOrgResponse)(nil),               // 1: city.economy.v1.GetOrgResponse
	(*UpdateOrgMoneyRequest)(nil),        // 2: city.economy.v1.UpdateOrgMoneyRequest
	(*UpdateOrgMoneyRequestItem)(nil),    // 3: city.economy.v1.UpdateOrgMoneyRequestItem
	(*UpdateOrgMoneyResponse)(nil),       // 4: city.economy.v1.UpdateOrgMoneyResponse
	(*UpdateOrgGoodsRequest)(nil),        // 5: city.economy.v1.UpdateOrgGoodsRequest
	(*UpdateOrgGoodsRequestItem)(nil),    // 6: city.economy.v1.UpdateOrgGoodsRequestItem
	(*UpdateOrgGoodsResponse)(nil),       // 7: city.economy.v1.UpdateOrgGoodsResponse
	(*UpdateOrgEmployeeRequest)(nil),     // 8: city.economy.v1.UpdateOrgEmployeeRequest
	(*UpdateOrgEmployeeRequestItem)(nil), // 9: city.economy.v1.UpdateOrgEmployeeRequestItem
	(*UpdateOrgEmployeeResponse)(nil),    // 10: city.economy.v1.UpdateOrgEmployeeResponse
	(*UpdateOrgJobRequest)(nil),          // 11: city.economy.v1.UpdateOrgJobRequest
	(*UpdateOrgJobRequestItem)(nil),      // 12: city.economy.v1.UpdateOrgJobRequestItem
	(*UpdateOrgJobResponse)(nil),         // 13: city.economy.v1.UpdateOrgJobResponse
	(*Org)(nil),                          // 14: city.economy.v1.Org
	(*Goods)(nil),                        // 15: city.economy.v1.Goods
	(*Employee)(nil),                     // 16: city.economy.v1.Employee
	(*Job)(nil),                          // 17: city.economy.v1.Job
}
var file_city_economy_v1_org_service_proto_depIdxs = []int32{
	14, // 0: city.economy.v1.GetOrgResponse.orgs:type_name -> city.economy.v1.Org
	3,  // 1: city.economy.v1.UpdateOrgMoneyRequest.items:type_name -> city.economy.v1.UpdateOrgMoneyRequestItem
	14, // 2: city.economy.v1.UpdateOrgMoneyResponse.orgs:type_name -> city.economy.v1.Org
	6,  // 3: city.economy.v1.UpdateOrgGoodsRequest.items:type_name -> city.economy.v1.UpdateOrgGoodsRequestItem
	15, // 4: city.economy.v1.UpdateOrgGoodsRequestItem.goods:type_name -> city.economy.v1.Goods
	14, // 5: city.economy.v1.UpdateOrgGoodsResponse.orgs:type_name -> city.economy.v1.Org
	9,  // 6: city.economy.v1.UpdateOrgEmployeeRequest.items:type_name -> city.economy.v1.UpdateOrgEmployeeRequestItem
	16, // 7: city.economy.v1.UpdateOrgEmployeeRequestItem.adds:type_name -> city.economy.v1.Employee
	16, // 8: city.economy.v1.UpdateOrgEmployeeRequestItem.updates:type_name -> city.economy.v1.Employee
	14, // 9: city.economy.v1.UpdateOrgEmployeeResponse.orgs:type_name -> city.economy.v1.Org
	12, // 10: city.economy.v1.UpdateOrgJobRequest.items:type_name -> city.economy.v1.UpdateOrgJobRequestItem
	17, // 11: city.economy.v1.UpdateOrgJobRequestItem.jobs:type_name -> city.economy.v1.Job
	14, // 12: city.economy.v1.UpdateOrgJobResponse.orgs:type_name -> city.economy.v1.Org
	0,  // 13: city.economy.v1.OrgService.GetOrg:input_type -> city.economy.v1.GetOrgRequest
	2,  // 14: city.economy.v1.OrgService.UpdateOrgMoney:input_type -> city.economy.v1.UpdateOrgMoneyRequest
	5,  // 15: city.economy.v1.OrgService.UpdateOrgGoods:input_type -> city.economy.v1.UpdateOrgGoodsRequest
	8,  // 16: city.economy.v1.OrgService.UpdateOrgEmployee:input_type -> city.economy.v1.UpdateOrgEmployeeRequest
	11, // 17: city.economy.v1.OrgService.UpdateOrgJob:input_type -> city.economy.v1.UpdateOrgJobRequest
	1,  // 18: city.economy.v1.OrgService.GetOrg:output_type -> city.economy.v1.GetOrgResponse
	4,  // 19: city.economy.v1.OrgService.UpdateOrgMoney:output_type -> city.economy.v1.UpdateOrgMoneyResponse
	7,  // 20: city.economy.v1.OrgService.UpdateOrgGoods:output_type -> city.economy.v1.UpdateOrgGoodsResponse
	10, // 21: city.economy.v1.OrgService.UpdateOrgEmployee:output_type -> city.economy.v1.UpdateOrgEmployeeResponse
	13, // 22: city.economy.v1.OrgService.UpdateOrgJob:output_type -> city.economy.v1.UpdateOrgJobResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_city_economy_v1_org_service_proto_init() }
func file_city_economy_v1_org_service_proto_init() {
	if File_city_economy_v1_org_service_proto != nil {
		return
	}
	file_city_economy_v1_economy_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_economy_v1_org_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_economy_v1_org_service_proto_goTypes,
		DependencyIndexes: file_city_economy_v1_org_service_proto_depIdxs,
		MessageInfos:      file_city_economy_v1_org_service_proto_msgTypes,
	}.Build()
	File_city_economy_v1_org_service_proto = out.File
	file_city_economy_v1_org_service_proto_rawDesc = nil
	file_city_economy_v1_org_service_proto_goTypes = nil
	file_city_economy_v1_org_service_proto_depIdxs = nil
}
