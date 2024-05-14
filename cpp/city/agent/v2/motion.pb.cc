// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/agent/v2/motion.proto

#include "city/agent/v2/motion.pb.h"

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
namespace agent {
namespace v2 {
PROTOBUF_CONSTEXPR AgentMotion::AgentMotion(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.activity_)*/{&::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized{}}
  , /*decltype(_impl_.position_)*/nullptr
  , /*decltype(_impl_.id_)*/0
  , /*decltype(_impl_.status_)*/0
  , /*decltype(_impl_.v_)*/0
  , /*decltype(_impl_.direction_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct AgentMotionDefaultTypeInternal {
  PROTOBUF_CONSTEXPR AgentMotionDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~AgentMotionDefaultTypeInternal() {}
  union {
    AgentMotion _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 AgentMotionDefaultTypeInternal _AgentMotion_default_instance_;
}  // namespace v2
}  // namespace agent
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fagent_2fv2_2fmotion_2eproto[1];
static const ::_pb::EnumDescriptor* file_level_enum_descriptors_city_2fagent_2fv2_2fmotion_2eproto[1];
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fagent_2fv2_2fmotion_2eproto = nullptr;

const uint32_t TableStruct_city_2fagent_2fv2_2fmotion_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::agent::v2::AgentMotion, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::agent::v2::AgentMotion, _impl_.id_),
  PROTOBUF_FIELD_OFFSET(::city::agent::v2::AgentMotion, _impl_.status_),
  PROTOBUF_FIELD_OFFSET(::city::agent::v2::AgentMotion, _impl_.position_),
  PROTOBUF_FIELD_OFFSET(::city::agent::v2::AgentMotion, _impl_.v_),
  PROTOBUF_FIELD_OFFSET(::city::agent::v2::AgentMotion, _impl_.direction_),
  PROTOBUF_FIELD_OFFSET(::city::agent::v2::AgentMotion, _impl_.activity_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::city::agent::v2::AgentMotion)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::agent::v2::_AgentMotion_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fagent_2fv2_2fmotion_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\032city/agent/v2/motion.proto\022\rcity.agent"
  ".v2\032\025city/geo/v2/geo.proto\"\313\001\n\013AgentMoti"
  "on\022\016\n\002id\030\001 \001(\005R\002id\022-\n\006status\030\002 \001(\0162\025.cit"
  "y.agent.v2.StatusR\006status\0221\n\010position\030\003 "
  "\001(\0132\025.city.geo.v2.PositionR\010position\022\014\n\001"
  "v\030\004 \001(\001R\001v\022\034\n\tdirection\030\005 \001(\001R\tdirection"
  "\022\032\n\010activity\030\006 \001(\tR\010activity:\002\030\001*\231\001\n\006Sta"
  "tus\022\026\n\022STATUS_UNSPECIFIED\020\000\022\020\n\014STATUS_SL"
  "EEP\020\001\022\022\n\016STATUS_DRIVING\020\002\022\022\n\016STATUS_WALK"
  "ING\020\003\022\020\n\014STATUS_CROWD\020\004\022\024\n\020STATUS_PASSEN"
  "GER\020\005\022\025\n\021STATUS_WAIT_ROUTE\020\006B\252\001\n\021com.cit"
  "y.agent.v2B\013MotionProtoP\001Z2git.fiblab.ne"
  "t/sim/protos/go/city/agent/v2;agentv2\242\002\003"
  "CAX\252\002\rCity.Agent.V2\312\002\rCity\\Agent\\V2\342\002\031Ci"
  "ty\\Agent\\V2\\GPBMetadata\352\002\017City::Agent::V"
  "2b\006proto3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_city_2fagent_2fv2_2fmotion_2eproto_deps[1] = {
  &::descriptor_table_city_2fgeo_2fv2_2fgeo_2eproto,
};
static ::_pbi::once_flag descriptor_table_city_2fagent_2fv2_2fmotion_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fagent_2fv2_2fmotion_2eproto = {
    false, false, 609, descriptor_table_protodef_city_2fagent_2fv2_2fmotion_2eproto,
    "city/agent/v2/motion.proto",
    &descriptor_table_city_2fagent_2fv2_2fmotion_2eproto_once, descriptor_table_city_2fagent_2fv2_2fmotion_2eproto_deps, 1, 1,
    schemas, file_default_instances, TableStruct_city_2fagent_2fv2_2fmotion_2eproto::offsets,
    file_level_metadata_city_2fagent_2fv2_2fmotion_2eproto, file_level_enum_descriptors_city_2fagent_2fv2_2fmotion_2eproto,
    file_level_service_descriptors_city_2fagent_2fv2_2fmotion_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fagent_2fv2_2fmotion_2eproto_getter() {
  return &descriptor_table_city_2fagent_2fv2_2fmotion_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fagent_2fv2_2fmotion_2eproto(&descriptor_table_city_2fagent_2fv2_2fmotion_2eproto);
namespace city {
namespace agent {
namespace v2 {
const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* Status_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_city_2fagent_2fv2_2fmotion_2eproto);
  return file_level_enum_descriptors_city_2fagent_2fv2_2fmotion_2eproto[0];
}
bool Status_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
    case 2:
    case 3:
    case 4:
    case 5:
    case 6:
      return true;
    default:
      return false;
  }
}


// ===================================================================

class AgentMotion::_Internal {
 public:
  static const ::city::geo::v2::Position& position(const AgentMotion* msg);
};

const ::city::geo::v2::Position&
AgentMotion::_Internal::position(const AgentMotion* msg) {
  return *msg->_impl_.position_;
}
void AgentMotion::clear_position() {
  if (GetArenaForAllocation() == nullptr && _impl_.position_ != nullptr) {
    delete _impl_.position_;
  }
  _impl_.position_ = nullptr;
}
AgentMotion::AgentMotion(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.agent.v2.AgentMotion)
}
AgentMotion::AgentMotion(const AgentMotion& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  AgentMotion* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.activity_){}
    , decltype(_impl_.position_){nullptr}
    , decltype(_impl_.id_){}
    , decltype(_impl_.status_){}
    , decltype(_impl_.v_){}
    , decltype(_impl_.direction_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.activity_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.activity_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_activity().empty()) {
    _this->_impl_.activity_.Set(from._internal_activity(), 
      _this->GetArenaForAllocation());
  }
  if (from._internal_has_position()) {
    _this->_impl_.position_ = new ::city::geo::v2::Position(*from._impl_.position_);
  }
  ::memcpy(&_impl_.id_, &from._impl_.id_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.direction_) -
    reinterpret_cast<char*>(&_impl_.id_)) + sizeof(_impl_.direction_));
  // @@protoc_insertion_point(copy_constructor:city.agent.v2.AgentMotion)
}

inline void AgentMotion::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.activity_){}
    , decltype(_impl_.position_){nullptr}
    , decltype(_impl_.id_){0}
    , decltype(_impl_.status_){0}
    , decltype(_impl_.v_){0}
    , decltype(_impl_.direction_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
  _impl_.activity_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.activity_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

AgentMotion::~AgentMotion() {
  // @@protoc_insertion_point(destructor:city.agent.v2.AgentMotion)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void AgentMotion::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.activity_.Destroy();
  if (this != internal_default_instance()) delete _impl_.position_;
}

void AgentMotion::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void AgentMotion::Clear() {
// @@protoc_insertion_point(message_clear_start:city.agent.v2.AgentMotion)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.activity_.ClearToEmpty();
  if (GetArenaForAllocation() == nullptr && _impl_.position_ != nullptr) {
    delete _impl_.position_;
  }
  _impl_.position_ = nullptr;
  ::memset(&_impl_.id_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&_impl_.direction_) -
      reinterpret_cast<char*>(&_impl_.id_)) + sizeof(_impl_.direction_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* AgentMotion::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // int32 id = 1 [json_name = "id"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 8)) {
          _impl_.id_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // .city.agent.v2.Status status = 2 [json_name = "status"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 16)) {
          uint64_t val = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
          _internal_set_status(static_cast<::city::agent::v2::Status>(val));
        } else
          goto handle_unusual;
        continue;
      // .city.geo.v2.Position position = 3 [json_name = "position"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 26)) {
          ptr = ctx->ParseMessage(_internal_mutable_position(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // double v = 4 [json_name = "v"];
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 33)) {
          _impl_.v_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // double direction = 5 [json_name = "direction"];
      case 5:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 41)) {
          _impl_.direction_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // string activity = 6 [json_name = "activity"];
      case 6:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 50)) {
          auto str = _internal_mutable_activity();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "city.agent.v2.AgentMotion.activity"));
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

uint8_t* AgentMotion::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.agent.v2.AgentMotion)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 id = 1 [json_name = "id"];
  if (this->_internal_id() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(1, this->_internal_id(), target);
  }

  // .city.agent.v2.Status status = 2 [json_name = "status"];
  if (this->_internal_status() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteEnumToArray(
      2, this->_internal_status(), target);
  }

  // .city.geo.v2.Position position = 3 [json_name = "position"];
  if (this->_internal_has_position()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(3, _Internal::position(this),
        _Internal::position(this).GetCachedSize(), target, stream);
  }

  // double v = 4 [json_name = "v"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_v = this->_internal_v();
  uint64_t raw_v;
  memcpy(&raw_v, &tmp_v, sizeof(tmp_v));
  if (raw_v != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(4, this->_internal_v(), target);
  }

  // double direction = 5 [json_name = "direction"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_direction = this->_internal_direction();
  uint64_t raw_direction;
  memcpy(&raw_direction, &tmp_direction, sizeof(tmp_direction));
  if (raw_direction != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(5, this->_internal_direction(), target);
  }

  // string activity = 6 [json_name = "activity"];
  if (!this->_internal_activity().empty()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_activity().data(), static_cast<int>(this->_internal_activity().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "city.agent.v2.AgentMotion.activity");
    target = stream->WriteStringMaybeAliased(
        6, this->_internal_activity(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.agent.v2.AgentMotion)
  return target;
}

size_t AgentMotion::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.agent.v2.AgentMotion)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string activity = 6 [json_name = "activity"];
  if (!this->_internal_activity().empty()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_activity());
  }

  // .city.geo.v2.Position position = 3 [json_name = "position"];
  if (this->_internal_has_position()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.position_);
  }

  // int32 id = 1 [json_name = "id"];
  if (this->_internal_id() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_id());
  }

  // .city.agent.v2.Status status = 2 [json_name = "status"];
  if (this->_internal_status() != 0) {
    total_size += 1 +
      ::_pbi::WireFormatLite::EnumSize(this->_internal_status());
  }

  // double v = 4 [json_name = "v"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_v = this->_internal_v();
  uint64_t raw_v;
  memcpy(&raw_v, &tmp_v, sizeof(tmp_v));
  if (raw_v != 0) {
    total_size += 1 + 8;
  }

  // double direction = 5 [json_name = "direction"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_direction = this->_internal_direction();
  uint64_t raw_direction;
  memcpy(&raw_direction, &tmp_direction, sizeof(tmp_direction));
  if (raw_direction != 0) {
    total_size += 1 + 8;
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData AgentMotion::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    AgentMotion::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*AgentMotion::GetClassData() const { return &_class_data_; }


void AgentMotion::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<AgentMotion*>(&to_msg);
  auto& from = static_cast<const AgentMotion&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.agent.v2.AgentMotion)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_activity().empty()) {
    _this->_internal_set_activity(from._internal_activity());
  }
  if (from._internal_has_position()) {
    _this->_internal_mutable_position()->::city::geo::v2::Position::MergeFrom(
        from._internal_position());
  }
  if (from._internal_id() != 0) {
    _this->_internal_set_id(from._internal_id());
  }
  if (from._internal_status() != 0) {
    _this->_internal_set_status(from._internal_status());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_v = from._internal_v();
  uint64_t raw_v;
  memcpy(&raw_v, &tmp_v, sizeof(tmp_v));
  if (raw_v != 0) {
    _this->_internal_set_v(from._internal_v());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_direction = from._internal_direction();
  uint64_t raw_direction;
  memcpy(&raw_direction, &tmp_direction, sizeof(tmp_direction));
  if (raw_direction != 0) {
    _this->_internal_set_direction(from._internal_direction());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void AgentMotion::CopyFrom(const AgentMotion& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.agent.v2.AgentMotion)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool AgentMotion::IsInitialized() const {
  return true;
}

void AgentMotion::InternalSwap(AgentMotion* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::InternalSwap(
      &_impl_.activity_, lhs_arena,
      &other->_impl_.activity_, rhs_arena
  );
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(AgentMotion, _impl_.direction_)
      + sizeof(AgentMotion::_impl_.direction_)
      - PROTOBUF_FIELD_OFFSET(AgentMotion, _impl_.position_)>(
          reinterpret_cast<char*>(&_impl_.position_),
          reinterpret_cast<char*>(&other->_impl_.position_));
}

::PROTOBUF_NAMESPACE_ID::Metadata AgentMotion::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fagent_2fv2_2fmotion_2eproto_getter, &descriptor_table_city_2fagent_2fv2_2fmotion_2eproto_once,
      file_level_metadata_city_2fagent_2fv2_2fmotion_2eproto[0]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v2
}  // namespace agent
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::agent::v2::AgentMotion*
Arena::CreateMaybeMessage< ::city::agent::v2::AgentMotion >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::agent::v2::AgentMotion >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
