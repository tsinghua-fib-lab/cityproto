// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/comm/interaction/aoi/v1/aoi_service.proto

#include "city/comm/interaction/aoi/v1/aoi_service.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

PROTOBUF_PRAGMA_INIT_SEG

namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = _pb::internal;

namespace city {
namespace comm {
namespace interaction {
namespace aoi {
namespace v1 {
PROTOBUF_CONSTEXPR GetBadAoiIDRequest::GetBadAoiIDRequest(
    ::_pbi::ConstantInitialized) {}
struct GetBadAoiIDRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR GetBadAoiIDRequestDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~GetBadAoiIDRequestDefaultTypeInternal() {}
  union {
    GetBadAoiIDRequest _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 GetBadAoiIDRequestDefaultTypeInternal _GetBadAoiIDRequest_default_instance_;
PROTOBUF_CONSTEXPR GetBadAoiIDResponse::GetBadAoiIDResponse(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.ids_)*/{}
  , /*decltype(_impl_._ids_cached_byte_size_)*/{0}
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct GetBadAoiIDResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR GetBadAoiIDResponseDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~GetBadAoiIDResponseDefaultTypeInternal() {}
  union {
    GetBadAoiIDResponse _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 GetBadAoiIDResponseDefaultTypeInternal _GetBadAoiIDResponse_default_instance_;
}  // namespace v1
}  // namespace aoi
}  // namespace interaction
}  // namespace comm
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto[2];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto = nullptr;

const uint32_t TableStruct_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::comm::interaction::aoi::v1::GetBadAoiIDResponse, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::comm::interaction::aoi::v1::GetBadAoiIDResponse, _impl_.ids_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::city::comm::interaction::aoi::v1::GetBadAoiIDRequest)},
  { 6, -1, -1, sizeof(::city::comm::interaction::aoi::v1::GetBadAoiIDResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::comm::interaction::aoi::v1::_GetBadAoiIDRequest_default_instance_._instance,
  &::city::comm::interaction::aoi::v1::_GetBadAoiIDResponse_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n.city/comm/interaction/aoi/v1/aoi_servi"
  "ce.proto\022\034city.comm.interaction.aoi.v1\"\024"
  "\n\022GetBadAoiIDRequest\"\'\n\023GetBadAoiIDRespo"
  "nse\022\020\n\003ids\030\001 \003(\005R\003ids2\200\001\n\nAoiService\022r\n\013"
  "GetBadAoiID\0220.city.comm.interaction.aoi."
  "v1.GetBadAoiIDRequest\0321.city.comm.intera"
  "ction.aoi.v1.GetBadAoiIDResponseB\211\002\n com"
  ".city.comm.interaction.aoi.v1B\017AoiServic"
  "eProtoP\001Z\?git.fiblab.net/sim/protos/go/c"
  "ity/comm/interaction/aoi/v1;aoiv1\242\002\004CCIA"
  "\252\002\034City.Comm.Interaction.Aoi.V1\312\002\034City\\C"
  "omm\\Interaction\\Aoi\\V1\342\002(City\\Comm\\Inter"
  "action\\Aoi\\V1\\GPBMetadata\352\002 City::Comm::"
  "Interaction::Aoi::V1b\006proto3"
  ;
static ::_pbi::once_flag descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto = {
    false, false, 548, descriptor_table_protodef_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto,
    "city/comm/interaction/aoi/v1/aoi_service.proto",
    &descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto_once, nullptr, 0, 2,
    schemas, file_default_instances, TableStruct_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto::offsets,
    file_level_metadata_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto, file_level_enum_descriptors_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto,
    file_level_service_descriptors_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto_getter() {
  return &descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto(&descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto);
namespace city {
namespace comm {
namespace interaction {
namespace aoi {
namespace v1 {

// ===================================================================

class GetBadAoiIDRequest::_Internal {
 public:
};

GetBadAoiIDRequest::GetBadAoiIDRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase(arena, is_message_owned) {
  // @@protoc_insertion_point(arena_constructor:city.comm.interaction.aoi.v1.GetBadAoiIDRequest)
}
GetBadAoiIDRequest::GetBadAoiIDRequest(const GetBadAoiIDRequest& from)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase() {
  GetBadAoiIDRequest* const _this = this; (void)_this;
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:city.comm.interaction.aoi.v1.GetBadAoiIDRequest)
}





const ::PROTOBUF_NAMESPACE_ID::Message::ClassData GetBadAoiIDRequest::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyImpl,
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeImpl,
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetBadAoiIDRequest::GetClassData() const { return &_class_data_; }







::PROTOBUF_NAMESPACE_ID::Metadata GetBadAoiIDRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto_getter, &descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto_once,
      file_level_metadata_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto[0]);
}

// ===================================================================

class GetBadAoiIDResponse::_Internal {
 public:
};

GetBadAoiIDResponse::GetBadAoiIDResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.comm.interaction.aoi.v1.GetBadAoiIDResponse)
}
GetBadAoiIDResponse::GetBadAoiIDResponse(const GetBadAoiIDResponse& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  GetBadAoiIDResponse* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.ids_){from._impl_.ids_}
    , /*decltype(_impl_._ids_cached_byte_size_)*/{0}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:city.comm.interaction.aoi.v1.GetBadAoiIDResponse)
}

inline void GetBadAoiIDResponse::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.ids_){arena}
    , /*decltype(_impl_._ids_cached_byte_size_)*/{0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

GetBadAoiIDResponse::~GetBadAoiIDResponse() {
  // @@protoc_insertion_point(destructor:city.comm.interaction.aoi.v1.GetBadAoiIDResponse)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void GetBadAoiIDResponse::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.ids_.~RepeatedField();
}

void GetBadAoiIDResponse::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void GetBadAoiIDResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:city.comm.interaction.aoi.v1.GetBadAoiIDResponse)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.ids_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* GetBadAoiIDResponse::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // repeated int32 ids = 1 [json_name = "ids"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          ptr = ::PROTOBUF_NAMESPACE_ID::internal::PackedInt32Parser(_internal_mutable_ids(), ptr, ctx);
          CHK_(ptr);
        } else if (static_cast<uint8_t>(tag) == 8) {
          _internal_add_ids(::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr));
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

uint8_t* GetBadAoiIDResponse::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.comm.interaction.aoi.v1.GetBadAoiIDResponse)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // repeated int32 ids = 1 [json_name = "ids"];
  {
    int byte_size = _impl_._ids_cached_byte_size_.load(std::memory_order_relaxed);
    if (byte_size > 0) {
      target = stream->WriteInt32Packed(
          1, _internal_ids(), byte_size, target);
    }
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.comm.interaction.aoi.v1.GetBadAoiIDResponse)
  return target;
}

size_t GetBadAoiIDResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.comm.interaction.aoi.v1.GetBadAoiIDResponse)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated int32 ids = 1 [json_name = "ids"];
  {
    size_t data_size = ::_pbi::WireFormatLite::
      Int32Size(this->_impl_.ids_);
    if (data_size > 0) {
      total_size += 1 +
        ::_pbi::WireFormatLite::Int32Size(static_cast<int32_t>(data_size));
    }
    int cached_size = ::_pbi::ToCachedSize(data_size);
    _impl_._ids_cached_byte_size_.store(cached_size,
                                    std::memory_order_relaxed);
    total_size += data_size;
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData GetBadAoiIDResponse::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    GetBadAoiIDResponse::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetBadAoiIDResponse::GetClassData() const { return &_class_data_; }


void GetBadAoiIDResponse::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<GetBadAoiIDResponse*>(&to_msg);
  auto& from = static_cast<const GetBadAoiIDResponse&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.comm.interaction.aoi.v1.GetBadAoiIDResponse)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.ids_.MergeFrom(from._impl_.ids_);
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void GetBadAoiIDResponse::CopyFrom(const GetBadAoiIDResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.comm.interaction.aoi.v1.GetBadAoiIDResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool GetBadAoiIDResponse::IsInitialized() const {
  return true;
}

void GetBadAoiIDResponse::InternalSwap(GetBadAoiIDResponse* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.ids_.InternalSwap(&other->_impl_.ids_);
}

::PROTOBUF_NAMESPACE_ID::Metadata GetBadAoiIDResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto_getter, &descriptor_table_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto_once,
      file_level_metadata_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace aoi
}  // namespace interaction
}  // namespace comm
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest*
Arena::CreateMaybeMessage< ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest >(arena);
}
template<> PROTOBUF_NOINLINE ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse*
Arena::CreateMaybeMessage< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>