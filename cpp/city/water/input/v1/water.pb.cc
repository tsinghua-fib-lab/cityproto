// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/water/input/v1/water.proto

#include "city/water/input/v1/water.pb.h"

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
namespace water {
namespace input {
namespace v1 {
PROTOBUF_CONSTEXPR RainPeriod::RainPeriod(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.rainfall_)*/0
  , /*decltype(_impl_.start_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct RainPeriodDefaultTypeInternal {
  PROTOBUF_CONSTEXPR RainPeriodDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~RainPeriodDefaultTypeInternal() {}
  union {
    RainPeriod _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 RainPeriodDefaultTypeInternal _RainPeriod_default_instance_;
PROTOBUF_CONSTEXPR Rain::Rain(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.rains_)*/{}
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct RainDefaultTypeInternal {
  PROTOBUF_CONSTEXPR RainDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~RainDefaultTypeInternal() {}
  union {
    Rain _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 RainDefaultTypeInternal _Rain_default_instance_;
}  // namespace v1
}  // namespace input
}  // namespace water
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fwater_2finput_2fv1_2fwater_2eproto[2];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_city_2fwater_2finput_2fv1_2fwater_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fwater_2finput_2fv1_2fwater_2eproto = nullptr;

const uint32_t TableStruct_city_2fwater_2finput_2fv1_2fwater_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::water::input::v1::RainPeriod, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::water::input::v1::RainPeriod, _impl_.start_),
  PROTOBUF_FIELD_OFFSET(::city::water::input::v1::RainPeriod, _impl_.rainfall_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::water::input::v1::Rain, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::water::input::v1::Rain, _impl_.rains_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::city::water::input::v1::RainPeriod)},
  { 8, -1, -1, sizeof(::city::water::input::v1::Rain)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::water::input::v1::_RainPeriod_default_instance_._instance,
  &::city::water::input::v1::_Rain_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fwater_2finput_2fv1_2fwater_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\037city/water/input/v1/water.proto\022\023city."
  "water.input.v1\">\n\nRainPeriod\022\024\n\005start\030\001 "
  "\001(\005R\005start\022\032\n\010rainfall\030\002 \001(\001R\010rainfall\"="
  "\n\004Rain\0225\n\005rains\030\001 \003(\0132\037.city.water.input"
  ".v1.RainPeriodR\005rainsB\316\001\n\027com.city.water"
  ".input.v1B\nWaterProtoP\001Z8git.fiblab.net/"
  "sim/protos/go/city/water/input/v1;inputv"
  "1\242\002\003CWI\252\002\023City.Water.Input.V1\312\002\023City\\Wat"
  "er\\Input\\V1\342\002\037City\\Water\\Input\\V1\\GPBMet"
  "adata\352\002\026City::Water::Input::V1b\006proto3"
  ;
static ::_pbi::once_flag descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto = {
    false, false, 398, descriptor_table_protodef_city_2fwater_2finput_2fv1_2fwater_2eproto,
    "city/water/input/v1/water.proto",
    &descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto_once, nullptr, 0, 2,
    schemas, file_default_instances, TableStruct_city_2fwater_2finput_2fv1_2fwater_2eproto::offsets,
    file_level_metadata_city_2fwater_2finput_2fv1_2fwater_2eproto, file_level_enum_descriptors_city_2fwater_2finput_2fv1_2fwater_2eproto,
    file_level_service_descriptors_city_2fwater_2finput_2fv1_2fwater_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto_getter() {
  return &descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fwater_2finput_2fv1_2fwater_2eproto(&descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto);
namespace city {
namespace water {
namespace input {
namespace v1 {

// ===================================================================

class RainPeriod::_Internal {
 public:
};

RainPeriod::RainPeriod(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.water.input.v1.RainPeriod)
}
RainPeriod::RainPeriod(const RainPeriod& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  RainPeriod* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.rainfall_){}
    , decltype(_impl_.start_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::memcpy(&_impl_.rainfall_, &from._impl_.rainfall_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.start_) -
    reinterpret_cast<char*>(&_impl_.rainfall_)) + sizeof(_impl_.start_));
  // @@protoc_insertion_point(copy_constructor:city.water.input.v1.RainPeriod)
}

inline void RainPeriod::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.rainfall_){0}
    , decltype(_impl_.start_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

RainPeriod::~RainPeriod() {
  // @@protoc_insertion_point(destructor:city.water.input.v1.RainPeriod)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void RainPeriod::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
}

void RainPeriod::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void RainPeriod::Clear() {
// @@protoc_insertion_point(message_clear_start:city.water.input.v1.RainPeriod)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  ::memset(&_impl_.rainfall_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&_impl_.start_) -
      reinterpret_cast<char*>(&_impl_.rainfall_)) + sizeof(_impl_.start_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* RainPeriod::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // int32 start = 1 [json_name = "start"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 8)) {
          _impl_.start_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // double rainfall = 2 [json_name = "rainfall"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 17)) {
          _impl_.rainfall_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
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

uint8_t* RainPeriod::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.water.input.v1.RainPeriod)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 start = 1 [json_name = "start"];
  if (this->_internal_start() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(1, this->_internal_start(), target);
  }

  // double rainfall = 2 [json_name = "rainfall"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_rainfall = this->_internal_rainfall();
  uint64_t raw_rainfall;
  memcpy(&raw_rainfall, &tmp_rainfall, sizeof(tmp_rainfall));
  if (raw_rainfall != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(2, this->_internal_rainfall(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.water.input.v1.RainPeriod)
  return target;
}

size_t RainPeriod::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.water.input.v1.RainPeriod)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // double rainfall = 2 [json_name = "rainfall"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_rainfall = this->_internal_rainfall();
  uint64_t raw_rainfall;
  memcpy(&raw_rainfall, &tmp_rainfall, sizeof(tmp_rainfall));
  if (raw_rainfall != 0) {
    total_size += 1 + 8;
  }

  // int32 start = 1 [json_name = "start"];
  if (this->_internal_start() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_start());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData RainPeriod::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    RainPeriod::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*RainPeriod::GetClassData() const { return &_class_data_; }


void RainPeriod::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<RainPeriod*>(&to_msg);
  auto& from = static_cast<const RainPeriod&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.water.input.v1.RainPeriod)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_rainfall = from._internal_rainfall();
  uint64_t raw_rainfall;
  memcpy(&raw_rainfall, &tmp_rainfall, sizeof(tmp_rainfall));
  if (raw_rainfall != 0) {
    _this->_internal_set_rainfall(from._internal_rainfall());
  }
  if (from._internal_start() != 0) {
    _this->_internal_set_start(from._internal_start());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void RainPeriod::CopyFrom(const RainPeriod& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.water.input.v1.RainPeriod)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool RainPeriod::IsInitialized() const {
  return true;
}

void RainPeriod::InternalSwap(RainPeriod* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(RainPeriod, _impl_.start_)
      + sizeof(RainPeriod::_impl_.start_)
      - PROTOBUF_FIELD_OFFSET(RainPeriod, _impl_.rainfall_)>(
          reinterpret_cast<char*>(&_impl_.rainfall_),
          reinterpret_cast<char*>(&other->_impl_.rainfall_));
}

::PROTOBUF_NAMESPACE_ID::Metadata RainPeriod::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto_getter, &descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto_once,
      file_level_metadata_city_2fwater_2finput_2fv1_2fwater_2eproto[0]);
}

// ===================================================================

class Rain::_Internal {
 public:
};

Rain::Rain(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.water.input.v1.Rain)
}
Rain::Rain(const Rain& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  Rain* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.rains_){from._impl_.rains_}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:city.water.input.v1.Rain)
}

inline void Rain::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.rains_){arena}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

Rain::~Rain() {
  // @@protoc_insertion_point(destructor:city.water.input.v1.Rain)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void Rain::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.rains_.~RepeatedPtrField();
}

void Rain::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void Rain::Clear() {
// @@protoc_insertion_point(message_clear_start:city.water.input.v1.Rain)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.rains_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Rain::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // repeated .city.water.input.v1.RainPeriod rains = 1 [json_name = "rains"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_rains(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<10>(ptr));
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

uint8_t* Rain::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.water.input.v1.Rain)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // repeated .city.water.input.v1.RainPeriod rains = 1 [json_name = "rains"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_rains_size()); i < n; i++) {
    const auto& repfield = this->_internal_rains(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(1, repfield, repfield.GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.water.input.v1.Rain)
  return target;
}

size_t Rain::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.water.input.v1.Rain)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.water.input.v1.RainPeriod rains = 1 [json_name = "rains"];
  total_size += 1UL * this->_internal_rains_size();
  for (const auto& msg : this->_impl_.rains_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData Rain::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    Rain::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*Rain::GetClassData() const { return &_class_data_; }


void Rain::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<Rain*>(&to_msg);
  auto& from = static_cast<const Rain&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.water.input.v1.Rain)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.rains_.MergeFrom(from._impl_.rains_);
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void Rain::CopyFrom(const Rain& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.water.input.v1.Rain)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Rain::IsInitialized() const {
  return true;
}

void Rain::InternalSwap(Rain* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.rains_.InternalSwap(&other->_impl_.rains_);
}

::PROTOBUF_NAMESPACE_ID::Metadata Rain::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto_getter, &descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto_once,
      file_level_metadata_city_2fwater_2finput_2fv1_2fwater_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace input
}  // namespace water
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::water::input::v1::RainPeriod*
Arena::CreateMaybeMessage< ::city::water::input::v1::RainPeriod >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::water::input::v1::RainPeriod >(arena);
}
template<> PROTOBUF_NOINLINE ::city::water::input::v1::Rain*
Arena::CreateMaybeMessage< ::city::water::input::v1::Rain >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::water::input::v1::Rain >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>