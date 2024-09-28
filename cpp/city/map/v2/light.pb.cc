// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/map/v2/light.proto

#include "city/map/v2/light.pb.h"

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
namespace map {
namespace v2 {
PROTOBUF_CONSTEXPR Phase::Phase(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.states_)*/{}
  , /*decltype(_impl_._states_cached_byte_size_)*/{0}
  , /*decltype(_impl_.duration_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct PhaseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR PhaseDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~PhaseDefaultTypeInternal() {}
  union {
    Phase _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 PhaseDefaultTypeInternal _Phase_default_instance_;
PROTOBUF_CONSTEXPR AvailablePhase::AvailablePhase(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.states_)*/{}
  , /*decltype(_impl_._states_cached_byte_size_)*/{0}
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct AvailablePhaseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR AvailablePhaseDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~AvailablePhaseDefaultTypeInternal() {}
  union {
    AvailablePhase _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 AvailablePhaseDefaultTypeInternal _AvailablePhase_default_instance_;
PROTOBUF_CONSTEXPR TrafficLight::TrafficLight(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.phases_)*/{}
  , /*decltype(_impl_.junction_id_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct TrafficLightDefaultTypeInternal {
  PROTOBUF_CONSTEXPR TrafficLightDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~TrafficLightDefaultTypeInternal() {}
  union {
    TrafficLight _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 TrafficLightDefaultTypeInternal _TrafficLight_default_instance_;
}  // namespace v2
}  // namespace map
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fmap_2fv2_2flight_2eproto[3];
static const ::_pb::EnumDescriptor* file_level_enum_descriptors_city_2fmap_2fv2_2flight_2eproto[1];
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fmap_2fv2_2flight_2eproto = nullptr;

const uint32_t TableStruct_city_2fmap_2fv2_2flight_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::map::v2::Phase, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::map::v2::Phase, _impl_.duration_),
  PROTOBUF_FIELD_OFFSET(::city::map::v2::Phase, _impl_.states_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::map::v2::AvailablePhase, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::map::v2::AvailablePhase, _impl_.states_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::map::v2::TrafficLight, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::map::v2::TrafficLight, _impl_.junction_id_),
  PROTOBUF_FIELD_OFFSET(::city::map::v2::TrafficLight, _impl_.phases_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::city::map::v2::Phase)},
  { 8, -1, -1, sizeof(::city::map::v2::AvailablePhase)},
  { 15, -1, -1, sizeof(::city::map::v2::TrafficLight)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::map::v2::_Phase_default_instance_._instance,
  &::city::map::v2::_AvailablePhase_default_instance_._instance,
  &::city::map::v2::_TrafficLight_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fmap_2fv2_2flight_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\027city/map/v2/light.proto\022\013city.map.v2\"T"
  "\n\005Phase\022\032\n\010duration\030\001 \001(\001R\010duration\022/\n\006s"
  "tates\030\002 \003(\0162\027.city.map.v2.LightStateR\006st"
  "ates\"A\n\016AvailablePhase\022/\n\006states\030\001 \003(\0162\027"
  ".city.map.v2.LightStateR\006states\"[\n\014Traff"
  "icLight\022\037\n\013junction_id\030\001 \001(\005R\njunctionId"
  "\022*\n\006phases\030\002 \003(\0132\022.city.map.v2.PhaseR\006ph"
  "ases*m\n\nLightState\022\033\n\027LIGHT_STATE_UNSPEC"
  "IFIED\020\000\022\023\n\017LIGHT_STATE_RED\020\001\022\025\n\021LIGHT_ST"
  "ATE_GREEN\020\002\022\026\n\022LIGHT_STATE_YELLOW\020\003B\236\001\n\017"
  "com.city.map.v2B\nLightProtoP\001Z1git.fibla"
  "b.net/sim/protos/v2/go/city/map/v2;mapv2"
  "\242\002\003CMX\252\002\013City.Map.V2\312\002\013City\\Map\\V2\342\002\027Cit"
  "y\\Map\\V2\\GPBMetadata\352\002\rCity::Map::V2b\006pr"
  "oto3"
  ;
static ::_pbi::once_flag descriptor_table_city_2fmap_2fv2_2flight_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fmap_2fv2_2flight_2eproto = {
    false, false, 564, descriptor_table_protodef_city_2fmap_2fv2_2flight_2eproto,
    "city/map/v2/light.proto",
    &descriptor_table_city_2fmap_2fv2_2flight_2eproto_once, nullptr, 0, 3,
    schemas, file_default_instances, TableStruct_city_2fmap_2fv2_2flight_2eproto::offsets,
    file_level_metadata_city_2fmap_2fv2_2flight_2eproto, file_level_enum_descriptors_city_2fmap_2fv2_2flight_2eproto,
    file_level_service_descriptors_city_2fmap_2fv2_2flight_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fmap_2fv2_2flight_2eproto_getter() {
  return &descriptor_table_city_2fmap_2fv2_2flight_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fmap_2fv2_2flight_2eproto(&descriptor_table_city_2fmap_2fv2_2flight_2eproto);
namespace city {
namespace map {
namespace v2 {
const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* LightState_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_city_2fmap_2fv2_2flight_2eproto);
  return file_level_enum_descriptors_city_2fmap_2fv2_2flight_2eproto[0];
}
bool LightState_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
    case 2:
    case 3:
      return true;
    default:
      return false;
  }
}


// ===================================================================

class Phase::_Internal {
 public:
};

Phase::Phase(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.map.v2.Phase)
}
Phase::Phase(const Phase& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  Phase* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.states_){from._impl_.states_}
    , /*decltype(_impl_._states_cached_byte_size_)*/{0}
    , decltype(_impl_.duration_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _this->_impl_.duration_ = from._impl_.duration_;
  // @@protoc_insertion_point(copy_constructor:city.map.v2.Phase)
}

inline void Phase::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.states_){arena}
    , /*decltype(_impl_._states_cached_byte_size_)*/{0}
    , decltype(_impl_.duration_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

Phase::~Phase() {
  // @@protoc_insertion_point(destructor:city.map.v2.Phase)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void Phase::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.states_.~RepeatedField();
}

void Phase::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void Phase::Clear() {
// @@protoc_insertion_point(message_clear_start:city.map.v2.Phase)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.states_.Clear();
  _impl_.duration_ = 0;
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Phase::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // double duration = 1 [json_name = "duration"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 9)) {
          _impl_.duration_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // repeated .city.map.v2.LightState states = 2 [json_name = "states"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 18)) {
          ptr = ::PROTOBUF_NAMESPACE_ID::internal::PackedEnumParser(_internal_mutable_states(), ptr, ctx);
          CHK_(ptr);
        } else if (static_cast<uint8_t>(tag) == 16) {
          uint64_t val = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
          _internal_add_states(static_cast<::city::map::v2::LightState>(val));
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

uint8_t* Phase::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.map.v2.Phase)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // double duration = 1 [json_name = "duration"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_duration = this->_internal_duration();
  uint64_t raw_duration;
  memcpy(&raw_duration, &tmp_duration, sizeof(tmp_duration));
  if (raw_duration != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(1, this->_internal_duration(), target);
  }

  // repeated .city.map.v2.LightState states = 2 [json_name = "states"];
  {
    int byte_size = _impl_._states_cached_byte_size_.load(std::memory_order_relaxed);
    if (byte_size > 0) {
      target = stream->WriteEnumPacked(
          2, _impl_.states_, byte_size, target);
    }
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.map.v2.Phase)
  return target;
}

size_t Phase::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.map.v2.Phase)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.map.v2.LightState states = 2 [json_name = "states"];
  {
    size_t data_size = 0;
    unsigned int count = static_cast<unsigned int>(this->_internal_states_size());for (unsigned int i = 0; i < count; i++) {
      data_size += ::_pbi::WireFormatLite::EnumSize(
        this->_internal_states(static_cast<int>(i)));
    }
    if (data_size > 0) {
      total_size += 1 +
        ::_pbi::WireFormatLite::Int32Size(static_cast<int32_t>(data_size));
    }
    int cached_size = ::_pbi::ToCachedSize(data_size);
    _impl_._states_cached_byte_size_.store(cached_size,
                                    std::memory_order_relaxed);
    total_size += data_size;
  }

  // double duration = 1 [json_name = "duration"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_duration = this->_internal_duration();
  uint64_t raw_duration;
  memcpy(&raw_duration, &tmp_duration, sizeof(tmp_duration));
  if (raw_duration != 0) {
    total_size += 1 + 8;
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData Phase::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    Phase::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*Phase::GetClassData() const { return &_class_data_; }


void Phase::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<Phase*>(&to_msg);
  auto& from = static_cast<const Phase&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.map.v2.Phase)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.states_.MergeFrom(from._impl_.states_);
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_duration = from._internal_duration();
  uint64_t raw_duration;
  memcpy(&raw_duration, &tmp_duration, sizeof(tmp_duration));
  if (raw_duration != 0) {
    _this->_internal_set_duration(from._internal_duration());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void Phase::CopyFrom(const Phase& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.map.v2.Phase)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Phase::IsInitialized() const {
  return true;
}

void Phase::InternalSwap(Phase* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.states_.InternalSwap(&other->_impl_.states_);
  swap(_impl_.duration_, other->_impl_.duration_);
}

::PROTOBUF_NAMESPACE_ID::Metadata Phase::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fmap_2fv2_2flight_2eproto_getter, &descriptor_table_city_2fmap_2fv2_2flight_2eproto_once,
      file_level_metadata_city_2fmap_2fv2_2flight_2eproto[0]);
}

// ===================================================================

class AvailablePhase::_Internal {
 public:
};

AvailablePhase::AvailablePhase(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.map.v2.AvailablePhase)
}
AvailablePhase::AvailablePhase(const AvailablePhase& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  AvailablePhase* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.states_){from._impl_.states_}
    , /*decltype(_impl_._states_cached_byte_size_)*/{0}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:city.map.v2.AvailablePhase)
}

inline void AvailablePhase::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.states_){arena}
    , /*decltype(_impl_._states_cached_byte_size_)*/{0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

AvailablePhase::~AvailablePhase() {
  // @@protoc_insertion_point(destructor:city.map.v2.AvailablePhase)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void AvailablePhase::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.states_.~RepeatedField();
}

void AvailablePhase::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void AvailablePhase::Clear() {
// @@protoc_insertion_point(message_clear_start:city.map.v2.AvailablePhase)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.states_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* AvailablePhase::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // repeated .city.map.v2.LightState states = 1 [json_name = "states"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          ptr = ::PROTOBUF_NAMESPACE_ID::internal::PackedEnumParser(_internal_mutable_states(), ptr, ctx);
          CHK_(ptr);
        } else if (static_cast<uint8_t>(tag) == 8) {
          uint64_t val = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
          _internal_add_states(static_cast<::city::map::v2::LightState>(val));
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

uint8_t* AvailablePhase::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.map.v2.AvailablePhase)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // repeated .city.map.v2.LightState states = 1 [json_name = "states"];
  {
    int byte_size = _impl_._states_cached_byte_size_.load(std::memory_order_relaxed);
    if (byte_size > 0) {
      target = stream->WriteEnumPacked(
          1, _impl_.states_, byte_size, target);
    }
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.map.v2.AvailablePhase)
  return target;
}

size_t AvailablePhase::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.map.v2.AvailablePhase)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.map.v2.LightState states = 1 [json_name = "states"];
  {
    size_t data_size = 0;
    unsigned int count = static_cast<unsigned int>(this->_internal_states_size());for (unsigned int i = 0; i < count; i++) {
      data_size += ::_pbi::WireFormatLite::EnumSize(
        this->_internal_states(static_cast<int>(i)));
    }
    if (data_size > 0) {
      total_size += 1 +
        ::_pbi::WireFormatLite::Int32Size(static_cast<int32_t>(data_size));
    }
    int cached_size = ::_pbi::ToCachedSize(data_size);
    _impl_._states_cached_byte_size_.store(cached_size,
                                    std::memory_order_relaxed);
    total_size += data_size;
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData AvailablePhase::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    AvailablePhase::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*AvailablePhase::GetClassData() const { return &_class_data_; }


void AvailablePhase::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<AvailablePhase*>(&to_msg);
  auto& from = static_cast<const AvailablePhase&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.map.v2.AvailablePhase)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.states_.MergeFrom(from._impl_.states_);
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void AvailablePhase::CopyFrom(const AvailablePhase& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.map.v2.AvailablePhase)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool AvailablePhase::IsInitialized() const {
  return true;
}

void AvailablePhase::InternalSwap(AvailablePhase* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.states_.InternalSwap(&other->_impl_.states_);
}

::PROTOBUF_NAMESPACE_ID::Metadata AvailablePhase::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fmap_2fv2_2flight_2eproto_getter, &descriptor_table_city_2fmap_2fv2_2flight_2eproto_once,
      file_level_metadata_city_2fmap_2fv2_2flight_2eproto[1]);
}

// ===================================================================

class TrafficLight::_Internal {
 public:
};

TrafficLight::TrafficLight(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.map.v2.TrafficLight)
}
TrafficLight::TrafficLight(const TrafficLight& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  TrafficLight* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.phases_){from._impl_.phases_}
    , decltype(_impl_.junction_id_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _this->_impl_.junction_id_ = from._impl_.junction_id_;
  // @@protoc_insertion_point(copy_constructor:city.map.v2.TrafficLight)
}

inline void TrafficLight::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.phases_){arena}
    , decltype(_impl_.junction_id_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

TrafficLight::~TrafficLight() {
  // @@protoc_insertion_point(destructor:city.map.v2.TrafficLight)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void TrafficLight::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.phases_.~RepeatedPtrField();
}

void TrafficLight::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void TrafficLight::Clear() {
// @@protoc_insertion_point(message_clear_start:city.map.v2.TrafficLight)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.phases_.Clear();
  _impl_.junction_id_ = 0;
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* TrafficLight::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // int32 junction_id = 1 [json_name = "junctionId"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 8)) {
          _impl_.junction_id_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // repeated .city.map.v2.Phase phases = 2 [json_name = "phases"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 18)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_phases(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<18>(ptr));
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

uint8_t* TrafficLight::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.map.v2.TrafficLight)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 junction_id = 1 [json_name = "junctionId"];
  if (this->_internal_junction_id() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(1, this->_internal_junction_id(), target);
  }

  // repeated .city.map.v2.Phase phases = 2 [json_name = "phases"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_phases_size()); i < n; i++) {
    const auto& repfield = this->_internal_phases(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(2, repfield, repfield.GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.map.v2.TrafficLight)
  return target;
}

size_t TrafficLight::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.map.v2.TrafficLight)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.map.v2.Phase phases = 2 [json_name = "phases"];
  total_size += 1UL * this->_internal_phases_size();
  for (const auto& msg : this->_impl_.phases_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  // int32 junction_id = 1 [json_name = "junctionId"];
  if (this->_internal_junction_id() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_junction_id());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData TrafficLight::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    TrafficLight::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*TrafficLight::GetClassData() const { return &_class_data_; }


void TrafficLight::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<TrafficLight*>(&to_msg);
  auto& from = static_cast<const TrafficLight&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.map.v2.TrafficLight)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.phases_.MergeFrom(from._impl_.phases_);
  if (from._internal_junction_id() != 0) {
    _this->_internal_set_junction_id(from._internal_junction_id());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void TrafficLight::CopyFrom(const TrafficLight& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.map.v2.TrafficLight)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool TrafficLight::IsInitialized() const {
  return true;
}

void TrafficLight::InternalSwap(TrafficLight* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.phases_.InternalSwap(&other->_impl_.phases_);
  swap(_impl_.junction_id_, other->_impl_.junction_id_);
}

::PROTOBUF_NAMESPACE_ID::Metadata TrafficLight::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fmap_2fv2_2flight_2eproto_getter, &descriptor_table_city_2fmap_2fv2_2flight_2eproto_once,
      file_level_metadata_city_2fmap_2fv2_2flight_2eproto[2]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v2
}  // namespace map
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::map::v2::Phase*
Arena::CreateMaybeMessage< ::city::map::v2::Phase >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::map::v2::Phase >(arena);
}
template<> PROTOBUF_NOINLINE ::city::map::v2::AvailablePhase*
Arena::CreateMaybeMessage< ::city::map::v2::AvailablePhase >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::map::v2::AvailablePhase >(arena);
}
template<> PROTOBUF_NOINLINE ::city::map::v2::TrafficLight*
Arena::CreateMaybeMessage< ::city::map::v2::TrafficLight >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::map::v2::TrafficLight >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
