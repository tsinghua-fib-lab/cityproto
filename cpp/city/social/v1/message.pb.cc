// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/social/v1/message.proto

#include "city/social/v1/message.pb.h"

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
namespace social {
namespace v1 {
PROTOBUF_CONSTEXPR Message::Message(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_._has_bits_)*/{}
  , /*decltype(_impl_._cached_size_)*/{}
  , /*decltype(_impl_.message_)*/{&::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized{}}
  , /*decltype(_impl_.from_)*/0
  , /*decltype(_impl_.to_)*/0
  , /*decltype(_impl_.t_)*/0} {}
struct MessageDefaultTypeInternal {
  PROTOBUF_CONSTEXPR MessageDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~MessageDefaultTypeInternal() {}
  union {
    Message _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 MessageDefaultTypeInternal _Message_default_instance_;
}  // namespace v1
}  // namespace social
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fsocial_2fv1_2fmessage_2eproto[1];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_city_2fsocial_2fv1_2fmessage_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fsocial_2fv1_2fmessage_2eproto = nullptr;

const uint32_t TableStruct_city_2fsocial_2fv1_2fmessage_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  PROTOBUF_FIELD_OFFSET(::city::social::v1::Message, _impl_._has_bits_),
  PROTOBUF_FIELD_OFFSET(::city::social::v1::Message, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::social::v1::Message, _impl_.from_),
  PROTOBUF_FIELD_OFFSET(::city::social::v1::Message, _impl_.to_),
  PROTOBUF_FIELD_OFFSET(::city::social::v1::Message, _impl_.message_),
  PROTOBUF_FIELD_OFFSET(::city::social::v1::Message, _impl_.t_),
  ~0u,
  ~0u,
  ~0u,
  0,
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, 10, -1, sizeof(::city::social::v1::Message)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::social::v1::_Message_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fsocial_2fv1_2fmessage_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\034city/social/v1/message.proto\022\016city.soc"
  "ial.v1\"`\n\007Message\022\022\n\004from\030\001 \001(\005R\004from\022\016\n"
  "\002to\030\002 \001(\005R\002to\022\030\n\007message\030\003 \001(\tR\007message\022"
  "\021\n\001t\030\004 \001(\001H\000R\001t\210\001\001B\004\n\002_tB\265\001\n\022com.city.so"
  "cial.v1B\014MessageProtoP\001Z7git.fiblab.net/"
  "sim/protos/v2/go/city/social/v1;socialv1"
  "\242\002\003CSX\252\002\016City.Social.V1\312\002\016City\\Social\\V1"
  "\342\002\032City\\Social\\V1\\GPBMetadata\352\002\020City::So"
  "cial::V1b\006proto3"
  ;
static ::_pbi::once_flag descriptor_table_city_2fsocial_2fv1_2fmessage_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fsocial_2fv1_2fmessage_2eproto = {
    false, false, 336, descriptor_table_protodef_city_2fsocial_2fv1_2fmessage_2eproto,
    "city/social/v1/message.proto",
    &descriptor_table_city_2fsocial_2fv1_2fmessage_2eproto_once, nullptr, 0, 1,
    schemas, file_default_instances, TableStruct_city_2fsocial_2fv1_2fmessage_2eproto::offsets,
    file_level_metadata_city_2fsocial_2fv1_2fmessage_2eproto, file_level_enum_descriptors_city_2fsocial_2fv1_2fmessage_2eproto,
    file_level_service_descriptors_city_2fsocial_2fv1_2fmessage_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fsocial_2fv1_2fmessage_2eproto_getter() {
  return &descriptor_table_city_2fsocial_2fv1_2fmessage_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fsocial_2fv1_2fmessage_2eproto(&descriptor_table_city_2fsocial_2fv1_2fmessage_2eproto);
namespace city {
namespace social {
namespace v1 {

// ===================================================================

class Message::_Internal {
 public:
  using HasBits = decltype(std::declval<Message>()._impl_._has_bits_);
  static void set_has_t(HasBits* has_bits) {
    (*has_bits)[0] |= 1u;
  }
};

Message::Message(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.social.v1.Message)
}
Message::Message(const Message& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  Message* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){from._impl_._has_bits_}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.message_){}
    , decltype(_impl_.from_){}
    , decltype(_impl_.to_){}
    , decltype(_impl_.t_){}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.message_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.message_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_message().empty()) {
    _this->_impl_.message_.Set(from._internal_message(), 
      _this->GetArenaForAllocation());
  }
  ::memcpy(&_impl_.from_, &from._impl_.from_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.t_) -
    reinterpret_cast<char*>(&_impl_.from_)) + sizeof(_impl_.t_));
  // @@protoc_insertion_point(copy_constructor:city.social.v1.Message)
}

inline void Message::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.message_){}
    , decltype(_impl_.from_){0}
    , decltype(_impl_.to_){0}
    , decltype(_impl_.t_){0}
  };
  _impl_.message_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.message_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

Message::~Message() {
  // @@protoc_insertion_point(destructor:city.social.v1.Message)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void Message::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.message_.Destroy();
}

void Message::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void Message::Clear() {
// @@protoc_insertion_point(message_clear_start:city.social.v1.Message)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.message_.ClearToEmpty();
  ::memset(&_impl_.from_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&_impl_.to_) -
      reinterpret_cast<char*>(&_impl_.from_)) + sizeof(_impl_.to_));
  _impl_.t_ = 0;
  _impl_._has_bits_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Message::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  _Internal::HasBits has_bits{};
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // int32 from = 1 [json_name = "from"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 8)) {
          _impl_.from_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // int32 to = 2 [json_name = "to"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 16)) {
          _impl_.to_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // string message = 3 [json_name = "message"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 26)) {
          auto str = _internal_mutable_message();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "city.social.v1.Message.message"));
        } else
          goto handle_unusual;
        continue;
      // optional double t = 4 [json_name = "t"];
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 33)) {
          _Internal::set_has_t(&has_bits);
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

uint8_t* Message::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.social.v1.Message)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 from = 1 [json_name = "from"];
  if (this->_internal_from() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(1, this->_internal_from(), target);
  }

  // int32 to = 2 [json_name = "to"];
  if (this->_internal_to() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(2, this->_internal_to(), target);
  }

  // string message = 3 [json_name = "message"];
  if (!this->_internal_message().empty()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_message().data(), static_cast<int>(this->_internal_message().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "city.social.v1.Message.message");
    target = stream->WriteStringMaybeAliased(
        3, this->_internal_message(), target);
  }

  // optional double t = 4 [json_name = "t"];
  if (_internal_has_t()) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(4, this->_internal_t(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.social.v1.Message)
  return target;
}

size_t Message::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.social.v1.Message)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string message = 3 [json_name = "message"];
  if (!this->_internal_message().empty()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_message());
  }

  // int32 from = 1 [json_name = "from"];
  if (this->_internal_from() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_from());
  }

  // int32 to = 2 [json_name = "to"];
  if (this->_internal_to() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_to());
  }

  // optional double t = 4 [json_name = "t"];
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000001u) {
    total_size += 1 + 8;
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData Message::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    Message::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*Message::GetClassData() const { return &_class_data_; }


void Message::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<Message*>(&to_msg);
  auto& from = static_cast<const Message&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.social.v1.Message)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_message().empty()) {
    _this->_internal_set_message(from._internal_message());
  }
  if (from._internal_from() != 0) {
    _this->_internal_set_from(from._internal_from());
  }
  if (from._internal_to() != 0) {
    _this->_internal_set_to(from._internal_to());
  }
  if (from._internal_has_t()) {
    _this->_internal_set_t(from._internal_t());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void Message::CopyFrom(const Message& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.social.v1.Message)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Message::IsInitialized() const {
  return true;
}

void Message::InternalSwap(Message* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_._has_bits_[0], other->_impl_._has_bits_[0]);
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::InternalSwap(
      &_impl_.message_, lhs_arena,
      &other->_impl_.message_, rhs_arena
  );
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(Message, _impl_.t_)
      + sizeof(Message::_impl_.t_)
      - PROTOBUF_FIELD_OFFSET(Message, _impl_.from_)>(
          reinterpret_cast<char*>(&_impl_.from_),
          reinterpret_cast<char*>(&other->_impl_.from_));
}

::PROTOBUF_NAMESPACE_ID::Metadata Message::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fsocial_2fv1_2fmessage_2eproto_getter, &descriptor_table_city_2fsocial_2fv1_2fmessage_2eproto_once,
      file_level_metadata_city_2fsocial_2fv1_2fmessage_2eproto[0]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace social
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::social::v1::Message*
Arena::CreateMaybeMessage< ::city::social::v1::Message >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::social::v1::Message >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
