// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/person/v1/person_runtime.proto

#include "city/person/v1/person_runtime.pb.h"

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
namespace person {
namespace v1 {
PROTOBUF_CONSTEXPR PersonRuntime::PersonRuntime(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_._has_bits_)*/{}
  , /*decltype(_impl_._cached_size_)*/{}
  , /*decltype(_impl_.events_)*/{}
  , /*decltype(_impl_.base_)*/nullptr
  , /*decltype(_impl_.motion_)*/nullptr} {}
struct PersonRuntimeDefaultTypeInternal {
  PROTOBUF_CONSTEXPR PersonRuntimeDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~PersonRuntimeDefaultTypeInternal() {}
  union {
    PersonRuntime _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 PersonRuntimeDefaultTypeInternal _PersonRuntime_default_instance_;
}  // namespace v1
}  // namespace person
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fperson_2fv1_2fperson_5fruntime_2eproto[1];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_city_2fperson_2fv1_2fperson_5fruntime_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fperson_2fv1_2fperson_5fruntime_2eproto = nullptr;

const uint32_t TableStruct_city_2fperson_2fv1_2fperson_5fruntime_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  PROTOBUF_FIELD_OFFSET(::city::person::v1::PersonRuntime, _impl_._has_bits_),
  PROTOBUF_FIELD_OFFSET(::city::person::v1::PersonRuntime, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::person::v1::PersonRuntime, _impl_.base_),
  PROTOBUF_FIELD_OFFSET(::city::person::v1::PersonRuntime, _impl_.motion_),
  PROTOBUF_FIELD_OFFSET(::city::person::v1::PersonRuntime, _impl_.events_),
  0,
  ~0u,
  ~0u,
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, 9, -1, sizeof(::city::person::v1::PersonRuntime)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::person::v1::_PersonRuntime_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fperson_2fv1_2fperson_5fruntime_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n#city/person/v1/person_runtime.proto\022\016c"
  "ity.person.v1\032\031city/event/v2/event.proto"
  "\032\033city/person/v1/motion.proto\032\033city/pers"
  "on/v1/person.proto\"\255\001\n\rPersonRuntime\022/\n\004"
  "base\030\001 \001(\0132\026.city.person.v1.PersonH\000R\004ba"
  "se\210\001\001\0224\n\006motion\030\002 \001(\0132\034.city.person.v1.P"
  "ersonMotionR\006motion\022,\n\006events\030\003 \003(\0132\024.ci"
  "ty.event.v2.EventR\006eventsB\007\n\005_baseB\273\001\n\022c"
  "om.city.person.v1B\022PersonRuntimeProtoP\001Z"
  "7git.fiblab.net/sim/protos/v2/go/city/pe"
  "rson/v1;personv1\242\002\003CPX\252\002\016City.Person.V1\312"
  "\002\016City\\Person\\V1\342\002\032City\\Person\\V1\\GPBMet"
  "adata\352\002\020City::Person::V1b\006proto3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto_deps[3] = {
  &::descriptor_table_city_2fevent_2fv2_2fevent_2eproto,
  &::descriptor_table_city_2fperson_2fv1_2fmotion_2eproto,
  &::descriptor_table_city_2fperson_2fv1_2fperson_2eproto,
};
static ::_pbi::once_flag descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto = {
    false, false, 512, descriptor_table_protodef_city_2fperson_2fv1_2fperson_5fruntime_2eproto,
    "city/person/v1/person_runtime.proto",
    &descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto_once, descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto_deps, 3, 1,
    schemas, file_default_instances, TableStruct_city_2fperson_2fv1_2fperson_5fruntime_2eproto::offsets,
    file_level_metadata_city_2fperson_2fv1_2fperson_5fruntime_2eproto, file_level_enum_descriptors_city_2fperson_2fv1_2fperson_5fruntime_2eproto,
    file_level_service_descriptors_city_2fperson_2fv1_2fperson_5fruntime_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto_getter() {
  return &descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fperson_2fv1_2fperson_5fruntime_2eproto(&descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto);
namespace city {
namespace person {
namespace v1 {

// ===================================================================

class PersonRuntime::_Internal {
 public:
  using HasBits = decltype(std::declval<PersonRuntime>()._impl_._has_bits_);
  static const ::city::person::v1::Person& base(const PersonRuntime* msg);
  static void set_has_base(HasBits* has_bits) {
    (*has_bits)[0] |= 1u;
  }
  static const ::city::person::v1::PersonMotion& motion(const PersonRuntime* msg);
};

const ::city::person::v1::Person&
PersonRuntime::_Internal::base(const PersonRuntime* msg) {
  return *msg->_impl_.base_;
}
const ::city::person::v1::PersonMotion&
PersonRuntime::_Internal::motion(const PersonRuntime* msg) {
  return *msg->_impl_.motion_;
}
void PersonRuntime::clear_base() {
  if (_impl_.base_ != nullptr) _impl_.base_->Clear();
  _impl_._has_bits_[0] &= ~0x00000001u;
}
void PersonRuntime::clear_motion() {
  if (GetArenaForAllocation() == nullptr && _impl_.motion_ != nullptr) {
    delete _impl_.motion_;
  }
  _impl_.motion_ = nullptr;
}
void PersonRuntime::clear_events() {
  _impl_.events_.Clear();
}
PersonRuntime::PersonRuntime(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.person.v1.PersonRuntime)
}
PersonRuntime::PersonRuntime(const PersonRuntime& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  PersonRuntime* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){from._impl_._has_bits_}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.events_){from._impl_.events_}
    , decltype(_impl_.base_){nullptr}
    , decltype(_impl_.motion_){nullptr}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  if (from._internal_has_base()) {
    _this->_impl_.base_ = new ::city::person::v1::Person(*from._impl_.base_);
  }
  if (from._internal_has_motion()) {
    _this->_impl_.motion_ = new ::city::person::v1::PersonMotion(*from._impl_.motion_);
  }
  // @@protoc_insertion_point(copy_constructor:city.person.v1.PersonRuntime)
}

inline void PersonRuntime::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.events_){arena}
    , decltype(_impl_.base_){nullptr}
    , decltype(_impl_.motion_){nullptr}
  };
}

PersonRuntime::~PersonRuntime() {
  // @@protoc_insertion_point(destructor:city.person.v1.PersonRuntime)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void PersonRuntime::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.events_.~RepeatedPtrField();
  if (this != internal_default_instance()) delete _impl_.base_;
  if (this != internal_default_instance()) delete _impl_.motion_;
}

void PersonRuntime::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void PersonRuntime::Clear() {
// @@protoc_insertion_point(message_clear_start:city.person.v1.PersonRuntime)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.events_.Clear();
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000001u) {
    GOOGLE_DCHECK(_impl_.base_ != nullptr);
    _impl_.base_->Clear();
  }
  if (GetArenaForAllocation() == nullptr && _impl_.motion_ != nullptr) {
    delete _impl_.motion_;
  }
  _impl_.motion_ = nullptr;
  _impl_._has_bits_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* PersonRuntime::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  _Internal::HasBits has_bits{};
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // optional .city.person.v1.Person base = 1 [json_name = "base"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          ptr = ctx->ParseMessage(_internal_mutable_base(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // .city.person.v1.PersonMotion motion = 2 [json_name = "motion"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 18)) {
          ptr = ctx->ParseMessage(_internal_mutable_motion(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // repeated .city.event.v2.Event events = 3 [json_name = "events"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 26)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_events(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<26>(ptr));
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

uint8_t* PersonRuntime::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.person.v1.PersonRuntime)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // optional .city.person.v1.Person base = 1 [json_name = "base"];
  if (_internal_has_base()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(1, _Internal::base(this),
        _Internal::base(this).GetCachedSize(), target, stream);
  }

  // .city.person.v1.PersonMotion motion = 2 [json_name = "motion"];
  if (this->_internal_has_motion()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(2, _Internal::motion(this),
        _Internal::motion(this).GetCachedSize(), target, stream);
  }

  // repeated .city.event.v2.Event events = 3 [json_name = "events"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_events_size()); i < n; i++) {
    const auto& repfield = this->_internal_events(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(3, repfield, repfield.GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.person.v1.PersonRuntime)
  return target;
}

size_t PersonRuntime::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.person.v1.PersonRuntime)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.event.v2.Event events = 3 [json_name = "events"];
  total_size += 1UL * this->_internal_events_size();
  for (const auto& msg : this->_impl_.events_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  // optional .city.person.v1.Person base = 1 [json_name = "base"];
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000001u) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.base_);
  }

  // .city.person.v1.PersonMotion motion = 2 [json_name = "motion"];
  if (this->_internal_has_motion()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.motion_);
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData PersonRuntime::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    PersonRuntime::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*PersonRuntime::GetClassData() const { return &_class_data_; }


void PersonRuntime::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<PersonRuntime*>(&to_msg);
  auto& from = static_cast<const PersonRuntime&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.person.v1.PersonRuntime)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.events_.MergeFrom(from._impl_.events_);
  if (from._internal_has_base()) {
    _this->_internal_mutable_base()->::city::person::v1::Person::MergeFrom(
        from._internal_base());
  }
  if (from._internal_has_motion()) {
    _this->_internal_mutable_motion()->::city::person::v1::PersonMotion::MergeFrom(
        from._internal_motion());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void PersonRuntime::CopyFrom(const PersonRuntime& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.person.v1.PersonRuntime)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool PersonRuntime::IsInitialized() const {
  return true;
}

void PersonRuntime::InternalSwap(PersonRuntime* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_._has_bits_[0], other->_impl_._has_bits_[0]);
  _impl_.events_.InternalSwap(&other->_impl_.events_);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(PersonRuntime, _impl_.motion_)
      + sizeof(PersonRuntime::_impl_.motion_)
      - PROTOBUF_FIELD_OFFSET(PersonRuntime, _impl_.base_)>(
          reinterpret_cast<char*>(&_impl_.base_),
          reinterpret_cast<char*>(&other->_impl_.base_));
}

::PROTOBUF_NAMESPACE_ID::Metadata PersonRuntime::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto_getter, &descriptor_table_city_2fperson_2fv1_2fperson_5fruntime_2eproto_once,
      file_level_metadata_city_2fperson_2fv1_2fperson_5fruntime_2eproto[0]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace person
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::person::v1::PersonRuntime*
Arena::CreateMaybeMessage< ::city::person::v1::PersonRuntime >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::person::v1::PersonRuntime >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
