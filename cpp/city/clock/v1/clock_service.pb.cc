// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/clock/v1/clock_service.proto

#include "city/clock/v1/clock_service.pb.h"

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
namespace clock {
namespace v1 {
PROTOBUF_CONSTEXPR NowRequest::NowRequest(
    ::_pbi::ConstantInitialized) {}
struct NowRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR NowRequestDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~NowRequestDefaultTypeInternal() {}
  union {
    NowRequest _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 NowRequestDefaultTypeInternal _NowRequest_default_instance_;
PROTOBUF_CONSTEXPR NowResponse::NowResponse(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_._has_bits_)*/{}
  , /*decltype(_impl_._cached_size_)*/{}
  , /*decltype(_impl_.t_)*/0
  , /*decltype(_impl_.day_)*/0} {}
struct NowResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR NowResponseDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~NowResponseDefaultTypeInternal() {}
  union {
    NowResponse _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 NowResponseDefaultTypeInternal _NowResponse_default_instance_;
}  // namespace v1
}  // namespace clock
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fclock_2fv1_2fclock_5fservice_2eproto[2];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_city_2fclock_2fv1_2fclock_5fservice_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fclock_2fv1_2fclock_5fservice_2eproto = nullptr;

const uint32_t TableStruct_city_2fclock_2fv1_2fclock_5fservice_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::clock::v1::NowRequest, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::clock::v1::NowResponse, _impl_._has_bits_),
  PROTOBUF_FIELD_OFFSET(::city::clock::v1::NowResponse, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::clock::v1::NowResponse, _impl_.day_),
  PROTOBUF_FIELD_OFFSET(::city::clock::v1::NowResponse, _impl_.t_),
  0,
  ~0u,
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::city::clock::v1::NowRequest)},
  { 6, 14, -1, sizeof(::city::clock::v1::NowResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::clock::v1::_NowRequest_default_instance_._instance,
  &::city::clock::v1::_NowResponse_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fclock_2fv1_2fclock_5fservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n!city/clock/v1/clock_service.proto\022\rcit"
  "y.clock.v1\"\014\n\nNowRequest\":\n\013NowResponse\022"
  "\025\n\003day\030\002 \001(\005H\000R\003day\210\001\001\022\014\n\001t\030\001 \001(\001R\001tB\006\n\004"
  "_day2L\n\014ClockService\022<\n\003Now\022\031.city.clock"
  ".v1.NowRequest\032\032.city.clock.v1.NowRespon"
  "seB\263\001\n\021com.city.clock.v1B\021ClockServicePr"
  "otoP\001Z5git.fiblab.net/sim/protos/v2/go/c"
  "ity/clock/v1;clockv1\242\002\003CCX\252\002\rCity.Clock."
  "V1\312\002\rCity\\Clock\\V1\342\002\031City\\Clock\\V1\\GPBMe"
  "tadata\352\002\017City::Clock::V1b\006proto3"
  ;
static ::_pbi::once_flag descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto = {
    false, false, 392, descriptor_table_protodef_city_2fclock_2fv1_2fclock_5fservice_2eproto,
    "city/clock/v1/clock_service.proto",
    &descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto_once, nullptr, 0, 2,
    schemas, file_default_instances, TableStruct_city_2fclock_2fv1_2fclock_5fservice_2eproto::offsets,
    file_level_metadata_city_2fclock_2fv1_2fclock_5fservice_2eproto, file_level_enum_descriptors_city_2fclock_2fv1_2fclock_5fservice_2eproto,
    file_level_service_descriptors_city_2fclock_2fv1_2fclock_5fservice_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto_getter() {
  return &descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fclock_2fv1_2fclock_5fservice_2eproto(&descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto);
namespace city {
namespace clock {
namespace v1 {

// ===================================================================

class NowRequest::_Internal {
 public:
};

NowRequest::NowRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase(arena, is_message_owned) {
  // @@protoc_insertion_point(arena_constructor:city.clock.v1.NowRequest)
}
NowRequest::NowRequest(const NowRequest& from)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase() {
  NowRequest* const _this = this; (void)_this;
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:city.clock.v1.NowRequest)
}





const ::PROTOBUF_NAMESPACE_ID::Message::ClassData NowRequest::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyImpl,
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeImpl,
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*NowRequest::GetClassData() const { return &_class_data_; }







::PROTOBUF_NAMESPACE_ID::Metadata NowRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto_getter, &descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto_once,
      file_level_metadata_city_2fclock_2fv1_2fclock_5fservice_2eproto[0]);
}

// ===================================================================

class NowResponse::_Internal {
 public:
  using HasBits = decltype(std::declval<NowResponse>()._impl_._has_bits_);
  static void set_has_day(HasBits* has_bits) {
    (*has_bits)[0] |= 1u;
  }
};

NowResponse::NowResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.clock.v1.NowResponse)
}
NowResponse::NowResponse(const NowResponse& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  NowResponse* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){from._impl_._has_bits_}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.t_){}
    , decltype(_impl_.day_){}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::memcpy(&_impl_.t_, &from._impl_.t_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.day_) -
    reinterpret_cast<char*>(&_impl_.t_)) + sizeof(_impl_.day_));
  // @@protoc_insertion_point(copy_constructor:city.clock.v1.NowResponse)
}

inline void NowResponse::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.t_){0}
    , decltype(_impl_.day_){0}
  };
}

NowResponse::~NowResponse() {
  // @@protoc_insertion_point(destructor:city.clock.v1.NowResponse)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void NowResponse::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
}

void NowResponse::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void NowResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:city.clock.v1.NowResponse)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.t_ = 0;
  _impl_.day_ = 0;
  _impl_._has_bits_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* NowResponse::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  _Internal::HasBits has_bits{};
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // double t = 1 [json_name = "t"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 9)) {
          _impl_.t_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // optional int32 day = 2 [json_name = "day"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 16)) {
          _Internal::set_has_day(&has_bits);
          _impl_.day_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
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
  _impl_._has_bits_.Or(has_bits);
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

uint8_t* NowResponse::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.clock.v1.NowResponse)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // double t = 1 [json_name = "t"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_t = this->_internal_t();
  uint64_t raw_t;
  memcpy(&raw_t, &tmp_t, sizeof(tmp_t));
  if (raw_t != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(1, this->_internal_t(), target);
  }

  // optional int32 day = 2 [json_name = "day"];
  if (_internal_has_day()) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(2, this->_internal_day(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.clock.v1.NowResponse)
  return target;
}

size_t NowResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.clock.v1.NowResponse)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // double t = 1 [json_name = "t"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_t = this->_internal_t();
  uint64_t raw_t;
  memcpy(&raw_t, &tmp_t, sizeof(tmp_t));
  if (raw_t != 0) {
    total_size += 1 + 8;
  }

  // optional int32 day = 2 [json_name = "day"];
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000001u) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_day());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData NowResponse::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    NowResponse::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*NowResponse::GetClassData() const { return &_class_data_; }


void NowResponse::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<NowResponse*>(&to_msg);
  auto& from = static_cast<const NowResponse&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.clock.v1.NowResponse)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_t = from._internal_t();
  uint64_t raw_t;
  memcpy(&raw_t, &tmp_t, sizeof(tmp_t));
  if (raw_t != 0) {
    _this->_internal_set_t(from._internal_t());
  }
  if (from._internal_has_day()) {
    _this->_internal_set_day(from._internal_day());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void NowResponse::CopyFrom(const NowResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.clock.v1.NowResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool NowResponse::IsInitialized() const {
  return true;
}

void NowResponse::InternalSwap(NowResponse* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_._has_bits_[0], other->_impl_._has_bits_[0]);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(NowResponse, _impl_.day_)
      + sizeof(NowResponse::_impl_.day_)
      - PROTOBUF_FIELD_OFFSET(NowResponse, _impl_.t_)>(
          reinterpret_cast<char*>(&_impl_.t_),
          reinterpret_cast<char*>(&other->_impl_.t_));
}

::PROTOBUF_NAMESPACE_ID::Metadata NowResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto_getter, &descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto_once,
      file_level_metadata_city_2fclock_2fv1_2fclock_5fservice_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace clock
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::clock::v1::NowRequest*
Arena::CreateMaybeMessage< ::city::clock::v1::NowRequest >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::clock::v1::NowRequest >(arena);
}
template<> PROTOBUF_NOINLINE ::city::clock::v1::NowResponse*
Arena::CreateMaybeMessage< ::city::clock::v1::NowResponse >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::clock::v1::NowResponse >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
