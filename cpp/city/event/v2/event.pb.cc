// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/event/v2/event.proto

#include "city/event/v2/event.pb.h"

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
namespace event {
namespace v2 {
PROTOBUF_CONSTEXPR Entity::Entity(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.type_)*/0
  , /*decltype(_impl_.id_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct EntityDefaultTypeInternal {
  PROTOBUF_CONSTEXPR EntityDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~EntityDefaultTypeInternal() {}
  union {
    Entity _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 EntityDefaultTypeInternal _Entity_default_instance_;
PROTOBUF_CONSTEXPR Event::Event(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_._has_bits_)*/{}
  , /*decltype(_impl_._cached_size_)*/{}
  , /*decltype(_impl_.topic_)*/{&::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized{}}
  , /*decltype(_impl_.content_)*/{&::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized{}}
  , /*decltype(_impl_.subject_)*/nullptr
  , /*decltype(_impl_.position_)*/nullptr
  , /*decltype(_impl_.t_)*/0
  , /*decltype(_impl_.id_)*/0} {}
struct EventDefaultTypeInternal {
  PROTOBUF_CONSTEXPR EventDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~EventDefaultTypeInternal() {}
  union {
    Event _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 EventDefaultTypeInternal _Event_default_instance_;
}  // namespace v2
}  // namespace event
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fevent_2fv2_2fevent_2eproto[2];
static const ::_pb::EnumDescriptor* file_level_enum_descriptors_city_2fevent_2fv2_2fevent_2eproto[1];
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fevent_2fv2_2fevent_2eproto = nullptr;

const uint32_t TableStruct_city_2fevent_2fv2_2fevent_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Entity, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Entity, _impl_.type_),
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Entity, _impl_.id_),
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Event, _impl_._has_bits_),
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Event, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Event, _impl_.topic_),
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Event, _impl_.id_),
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Event, _impl_.subject_),
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Event, _impl_.content_),
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Event, _impl_.position_),
  PROTOBUF_FIELD_OFFSET(::city::event::v2::Event, _impl_.t_),
  ~0u,
  0,
  ~0u,
  ~0u,
  ~0u,
  ~0u,
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::city::event::v2::Entity)},
  { 8, 20, -1, sizeof(::city::event::v2::Event)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::event::v2::_Entity_default_instance_._instance,
  &::city::event::v2::_Event_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fevent_2fv2_2fevent_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\031city/event/v2/event.proto\022\rcity.event."
  "v2\032\025city/geo/v2/geo.proto\"G\n\006Entity\022-\n\004t"
  "ype\030\001 \001(\0162\031.city.event.v2.EntityTypeR\004ty"
  "pe\022\016\n\002id\030\002 \001(\005R\002id\"\305\001\n\005Event\022\024\n\005topic\030\001 "
  "\001(\tR\005topic\022\023\n\002id\030\002 \001(\005H\000R\002id\210\001\001\022/\n\007subje"
  "ct\030\003 \001(\0132\025.city.event.v2.EntityR\007subject"
  "\022\030\n\007content\030\004 \001(\tR\007content\0221\n\010position\030\005"
  " \001(\0132\025.city.geo.v2.PositionR\010position\022\014\n"
  "\001t\030\006 \001(\001R\001tB\005\n\003_id*\306\001\n\nEntityType\022\033\n\027ENT"
  "ITY_TYPE_UNSPECIFIED\020\000\022\024\n\020ENTITY_TYPE_LA"
  "NE\020\001\022\024\n\020ENTITY_TYPE_ROAD\020\002\022\030\n\024ENTITY_TYP"
  "E_JUNCTION\020\003\022\023\n\017ENTITY_TYPE_AOI\020\004\022\023\n\017ENT"
  "ITY_TYPE_POI\020\005\022\026\n\022ENTITY_TYPE_PERSON\020\006\022\023"
  "\n\017ENTITY_TYPE_ORG\020\007B\254\001\n\021com.city.event.v"
  "2B\nEventProtoP\001Z5git.fiblab.net/sim/prot"
  "os/v2/go/city/event/v2;eventv2\242\002\003CEX\252\002\rC"
  "ity.Event.V2\312\002\rCity\\Event\\V2\342\002\031City\\Even"
  "t\\V2\\GPBMetadata\352\002\017City::Event::V2b\006prot"
  "o3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_city_2fevent_2fv2_2fevent_2eproto_deps[1] = {
  &::descriptor_table_city_2fgeo_2fv2_2fgeo_2eproto,
};
static ::_pbi::once_flag descriptor_table_city_2fevent_2fv2_2fevent_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fevent_2fv2_2fevent_2eproto = {
    false, false, 722, descriptor_table_protodef_city_2fevent_2fv2_2fevent_2eproto,
    "city/event/v2/event.proto",
    &descriptor_table_city_2fevent_2fv2_2fevent_2eproto_once, descriptor_table_city_2fevent_2fv2_2fevent_2eproto_deps, 1, 2,
    schemas, file_default_instances, TableStruct_city_2fevent_2fv2_2fevent_2eproto::offsets,
    file_level_metadata_city_2fevent_2fv2_2fevent_2eproto, file_level_enum_descriptors_city_2fevent_2fv2_2fevent_2eproto,
    file_level_service_descriptors_city_2fevent_2fv2_2fevent_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fevent_2fv2_2fevent_2eproto_getter() {
  return &descriptor_table_city_2fevent_2fv2_2fevent_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fevent_2fv2_2fevent_2eproto(&descriptor_table_city_2fevent_2fv2_2fevent_2eproto);
namespace city {
namespace event {
namespace v2 {
const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* EntityType_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_city_2fevent_2fv2_2fevent_2eproto);
  return file_level_enum_descriptors_city_2fevent_2fv2_2fevent_2eproto[0];
}
bool EntityType_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
    case 2:
    case 3:
    case 4:
    case 5:
    case 6:
    case 7:
      return true;
    default:
      return false;
  }
}


// ===================================================================

class Entity::_Internal {
 public:
};

Entity::Entity(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.event.v2.Entity)
}
Entity::Entity(const Entity& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  Entity* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.type_){}
    , decltype(_impl_.id_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::memcpy(&_impl_.type_, &from._impl_.type_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.id_) -
    reinterpret_cast<char*>(&_impl_.type_)) + sizeof(_impl_.id_));
  // @@protoc_insertion_point(copy_constructor:city.event.v2.Entity)
}

inline void Entity::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.type_){0}
    , decltype(_impl_.id_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

Entity::~Entity() {
  // @@protoc_insertion_point(destructor:city.event.v2.Entity)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void Entity::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
}

void Entity::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void Entity::Clear() {
// @@protoc_insertion_point(message_clear_start:city.event.v2.Entity)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  ::memset(&_impl_.type_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&_impl_.id_) -
      reinterpret_cast<char*>(&_impl_.type_)) + sizeof(_impl_.id_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Entity::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // .city.event.v2.EntityType type = 1 [json_name = "type"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 8)) {
          uint64_t val = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
          _internal_set_type(static_cast<::city::event::v2::EntityType>(val));
        } else
          goto handle_unusual;
        continue;
      // int32 id = 2 [json_name = "id"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 16)) {
          _impl_.id_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
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

uint8_t* Entity::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.event.v2.Entity)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // .city.event.v2.EntityType type = 1 [json_name = "type"];
  if (this->_internal_type() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteEnumToArray(
      1, this->_internal_type(), target);
  }

  // int32 id = 2 [json_name = "id"];
  if (this->_internal_id() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(2, this->_internal_id(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.event.v2.Entity)
  return target;
}

size_t Entity::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.event.v2.Entity)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // .city.event.v2.EntityType type = 1 [json_name = "type"];
  if (this->_internal_type() != 0) {
    total_size += 1 +
      ::_pbi::WireFormatLite::EnumSize(this->_internal_type());
  }

  // int32 id = 2 [json_name = "id"];
  if (this->_internal_id() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_id());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData Entity::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    Entity::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*Entity::GetClassData() const { return &_class_data_; }


void Entity::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<Entity*>(&to_msg);
  auto& from = static_cast<const Entity&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.event.v2.Entity)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (from._internal_type() != 0) {
    _this->_internal_set_type(from._internal_type());
  }
  if (from._internal_id() != 0) {
    _this->_internal_set_id(from._internal_id());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void Entity::CopyFrom(const Entity& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.event.v2.Entity)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Entity::IsInitialized() const {
  return true;
}

void Entity::InternalSwap(Entity* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(Entity, _impl_.id_)
      + sizeof(Entity::_impl_.id_)
      - PROTOBUF_FIELD_OFFSET(Entity, _impl_.type_)>(
          reinterpret_cast<char*>(&_impl_.type_),
          reinterpret_cast<char*>(&other->_impl_.type_));
}

::PROTOBUF_NAMESPACE_ID::Metadata Entity::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fevent_2fv2_2fevent_2eproto_getter, &descriptor_table_city_2fevent_2fv2_2fevent_2eproto_once,
      file_level_metadata_city_2fevent_2fv2_2fevent_2eproto[0]);
}

// ===================================================================

class Event::_Internal {
 public:
  using HasBits = decltype(std::declval<Event>()._impl_._has_bits_);
  static void set_has_id(HasBits* has_bits) {
    (*has_bits)[0] |= 1u;
  }
  static const ::city::event::v2::Entity& subject(const Event* msg);
  static const ::city::geo::v2::Position& position(const Event* msg);
};

const ::city::event::v2::Entity&
Event::_Internal::subject(const Event* msg) {
  return *msg->_impl_.subject_;
}
const ::city::geo::v2::Position&
Event::_Internal::position(const Event* msg) {
  return *msg->_impl_.position_;
}
void Event::clear_position() {
  if (GetArenaForAllocation() == nullptr && _impl_.position_ != nullptr) {
    delete _impl_.position_;
  }
  _impl_.position_ = nullptr;
}
Event::Event(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.event.v2.Event)
}
Event::Event(const Event& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  Event* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){from._impl_._has_bits_}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.topic_){}
    , decltype(_impl_.content_){}
    , decltype(_impl_.subject_){nullptr}
    , decltype(_impl_.position_){nullptr}
    , decltype(_impl_.t_){}
    , decltype(_impl_.id_){}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.topic_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.topic_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_topic().empty()) {
    _this->_impl_.topic_.Set(from._internal_topic(), 
      _this->GetArenaForAllocation());
  }
  _impl_.content_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.content_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_content().empty()) {
    _this->_impl_.content_.Set(from._internal_content(), 
      _this->GetArenaForAllocation());
  }
  if (from._internal_has_subject()) {
    _this->_impl_.subject_ = new ::city::event::v2::Entity(*from._impl_.subject_);
  }
  if (from._internal_has_position()) {
    _this->_impl_.position_ = new ::city::geo::v2::Position(*from._impl_.position_);
  }
  ::memcpy(&_impl_.t_, &from._impl_.t_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.id_) -
    reinterpret_cast<char*>(&_impl_.t_)) + sizeof(_impl_.id_));
  // @@protoc_insertion_point(copy_constructor:city.event.v2.Event)
}

inline void Event::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.topic_){}
    , decltype(_impl_.content_){}
    , decltype(_impl_.subject_){nullptr}
    , decltype(_impl_.position_){nullptr}
    , decltype(_impl_.t_){0}
    , decltype(_impl_.id_){0}
  };
  _impl_.topic_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.topic_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  _impl_.content_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.content_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

Event::~Event() {
  // @@protoc_insertion_point(destructor:city.event.v2.Event)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void Event::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.topic_.Destroy();
  _impl_.content_.Destroy();
  if (this != internal_default_instance()) delete _impl_.subject_;
  if (this != internal_default_instance()) delete _impl_.position_;
}

void Event::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void Event::Clear() {
// @@protoc_insertion_point(message_clear_start:city.event.v2.Event)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.topic_.ClearToEmpty();
  _impl_.content_.ClearToEmpty();
  if (GetArenaForAllocation() == nullptr && _impl_.subject_ != nullptr) {
    delete _impl_.subject_;
  }
  _impl_.subject_ = nullptr;
  if (GetArenaForAllocation() == nullptr && _impl_.position_ != nullptr) {
    delete _impl_.position_;
  }
  _impl_.position_ = nullptr;
  _impl_.t_ = 0;
  _impl_.id_ = 0;
  _impl_._has_bits_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Event::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  _Internal::HasBits has_bits{};
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // string topic = 1 [json_name = "topic"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          auto str = _internal_mutable_topic();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "city.event.v2.Event.topic"));
        } else
          goto handle_unusual;
        continue;
      // optional int32 id = 2 [json_name = "id"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 16)) {
          _Internal::set_has_id(&has_bits);
          _impl_.id_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // .city.event.v2.Entity subject = 3 [json_name = "subject"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 26)) {
          ptr = ctx->ParseMessage(_internal_mutable_subject(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // string content = 4 [json_name = "content"];
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 34)) {
          auto str = _internal_mutable_content();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "city.event.v2.Event.content"));
        } else
          goto handle_unusual;
        continue;
      // .city.geo.v2.Position position = 5 [json_name = "position"];
      case 5:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 42)) {
          ptr = ctx->ParseMessage(_internal_mutable_position(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // double t = 6 [json_name = "t"];
      case 6:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 49)) {
          _impl_.t_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
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
  _impl_._has_bits_.Or(has_bits);
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

uint8_t* Event::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.event.v2.Event)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // string topic = 1 [json_name = "topic"];
  if (!this->_internal_topic().empty()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_topic().data(), static_cast<int>(this->_internal_topic().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "city.event.v2.Event.topic");
    target = stream->WriteStringMaybeAliased(
        1, this->_internal_topic(), target);
  }

  // optional int32 id = 2 [json_name = "id"];
  if (_internal_has_id()) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(2, this->_internal_id(), target);
  }

  // .city.event.v2.Entity subject = 3 [json_name = "subject"];
  if (this->_internal_has_subject()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(3, _Internal::subject(this),
        _Internal::subject(this).GetCachedSize(), target, stream);
  }

  // string content = 4 [json_name = "content"];
  if (!this->_internal_content().empty()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_content().data(), static_cast<int>(this->_internal_content().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "city.event.v2.Event.content");
    target = stream->WriteStringMaybeAliased(
        4, this->_internal_content(), target);
  }

  // .city.geo.v2.Position position = 5 [json_name = "position"];
  if (this->_internal_has_position()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(5, _Internal::position(this),
        _Internal::position(this).GetCachedSize(), target, stream);
  }

  // double t = 6 [json_name = "t"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_t = this->_internal_t();
  uint64_t raw_t;
  memcpy(&raw_t, &tmp_t, sizeof(tmp_t));
  if (raw_t != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(6, this->_internal_t(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.event.v2.Event)
  return target;
}

size_t Event::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.event.v2.Event)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string topic = 1 [json_name = "topic"];
  if (!this->_internal_topic().empty()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_topic());
  }

  // string content = 4 [json_name = "content"];
  if (!this->_internal_content().empty()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_content());
  }

  // .city.event.v2.Entity subject = 3 [json_name = "subject"];
  if (this->_internal_has_subject()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.subject_);
  }

  // .city.geo.v2.Position position = 5 [json_name = "position"];
  if (this->_internal_has_position()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.position_);
  }

  // double t = 6 [json_name = "t"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_t = this->_internal_t();
  uint64_t raw_t;
  memcpy(&raw_t, &tmp_t, sizeof(tmp_t));
  if (raw_t != 0) {
    total_size += 1 + 8;
  }

  // optional int32 id = 2 [json_name = "id"];
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000001u) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_id());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData Event::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    Event::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*Event::GetClassData() const { return &_class_data_; }


void Event::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<Event*>(&to_msg);
  auto& from = static_cast<const Event&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.event.v2.Event)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_topic().empty()) {
    _this->_internal_set_topic(from._internal_topic());
  }
  if (!from._internal_content().empty()) {
    _this->_internal_set_content(from._internal_content());
  }
  if (from._internal_has_subject()) {
    _this->_internal_mutable_subject()->::city::event::v2::Entity::MergeFrom(
        from._internal_subject());
  }
  if (from._internal_has_position()) {
    _this->_internal_mutable_position()->::city::geo::v2::Position::MergeFrom(
        from._internal_position());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_t = from._internal_t();
  uint64_t raw_t;
  memcpy(&raw_t, &tmp_t, sizeof(tmp_t));
  if (raw_t != 0) {
    _this->_internal_set_t(from._internal_t());
  }
  if (from._internal_has_id()) {
    _this->_internal_set_id(from._internal_id());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void Event::CopyFrom(const Event& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.event.v2.Event)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Event::IsInitialized() const {
  return true;
}

void Event::InternalSwap(Event* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_._has_bits_[0], other->_impl_._has_bits_[0]);
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::InternalSwap(
      &_impl_.topic_, lhs_arena,
      &other->_impl_.topic_, rhs_arena
  );
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::InternalSwap(
      &_impl_.content_, lhs_arena,
      &other->_impl_.content_, rhs_arena
  );
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(Event, _impl_.id_)
      + sizeof(Event::_impl_.id_)
      - PROTOBUF_FIELD_OFFSET(Event, _impl_.subject_)>(
          reinterpret_cast<char*>(&_impl_.subject_),
          reinterpret_cast<char*>(&other->_impl_.subject_));
}

::PROTOBUF_NAMESPACE_ID::Metadata Event::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fevent_2fv2_2fevent_2eproto_getter, &descriptor_table_city_2fevent_2fv2_2fevent_2eproto_once,
      file_level_metadata_city_2fevent_2fv2_2fevent_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v2
}  // namespace event
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::event::v2::Entity*
Arena::CreateMaybeMessage< ::city::event::v2::Entity >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::event::v2::Entity >(arena);
}
template<> PROTOBUF_NOINLINE ::city::event::v2::Event*
Arena::CreateMaybeMessage< ::city::event::v2::Event >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::event::v2::Event >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
