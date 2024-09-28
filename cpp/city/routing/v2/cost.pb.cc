// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/routing/v2/cost.proto

#include "city/routing/v2/cost.pb.h"

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
namespace routing {
namespace v2 {
PROTOBUF_CONSTEXPR Cost::Cost(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_._has_bits_)*/{}
  , /*decltype(_impl_._cached_size_)*/{}
  , /*decltype(_impl_.cost_)*/0
  , /*decltype(_impl_.time_)*/0
  , /*decltype(_impl_.id_)*/0} {}
struct CostDefaultTypeInternal {
  PROTOBUF_CONSTEXPR CostDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~CostDefaultTypeInternal() {}
  union {
    Cost _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 CostDefaultTypeInternal _Cost_default_instance_;
}  // namespace v2
}  // namespace routing
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2frouting_2fv2_2fcost_2eproto[1];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_city_2frouting_2fv2_2fcost_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2frouting_2fv2_2fcost_2eproto = nullptr;

const uint32_t TableStruct_city_2frouting_2fv2_2fcost_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  PROTOBUF_FIELD_OFFSET(::city::routing::v2::Cost, _impl_._has_bits_),
  PROTOBUF_FIELD_OFFSET(::city::routing::v2::Cost, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::routing::v2::Cost, _impl_.id_),
  PROTOBUF_FIELD_OFFSET(::city::routing::v2::Cost, _impl_.cost_),
  PROTOBUF_FIELD_OFFSET(::city::routing::v2::Cost, _impl_.time_),
  ~0u,
  ~0u,
  0,
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, 9, -1, sizeof(::city::routing::v2::Cost)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::routing::v2::_Cost_default_instance_._instance,
};

const char descriptor_table_protodef_city_2frouting_2fv2_2fcost_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\032city/routing/v2/cost.proto\022\017city.routi"
  "ng.v2\"L\n\004Cost\022\016\n\002id\030\001 \001(\005R\002id\022\022\n\004cost\030\002 "
  "\001(\001R\004cost\022\027\n\004time\030\003 \001(\001H\000R\004time\210\001\001B\007\n\005_t"
  "imeB\271\001\n\023com.city.routing.v2B\tCostProtoP\001"
  "Z9git.fiblab.net/sim/protos/go/v2/city/r"
  "outing/v2;routingv2\242\002\003CRX\252\002\017City.Routing"
  ".V2\312\002\017City\\Routing\\V2\342\002\033City\\Routing\\V2\\"
  "GPBMetadata\352\002\021City::Routing::V2b\006proto3"
  ;
static ::_pbi::once_flag descriptor_table_city_2frouting_2fv2_2fcost_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2frouting_2fv2_2fcost_2eproto = {
    false, false, 319, descriptor_table_protodef_city_2frouting_2fv2_2fcost_2eproto,
    "city/routing/v2/cost.proto",
    &descriptor_table_city_2frouting_2fv2_2fcost_2eproto_once, nullptr, 0, 1,
    schemas, file_default_instances, TableStruct_city_2frouting_2fv2_2fcost_2eproto::offsets,
    file_level_metadata_city_2frouting_2fv2_2fcost_2eproto, file_level_enum_descriptors_city_2frouting_2fv2_2fcost_2eproto,
    file_level_service_descriptors_city_2frouting_2fv2_2fcost_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2frouting_2fv2_2fcost_2eproto_getter() {
  return &descriptor_table_city_2frouting_2fv2_2fcost_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2frouting_2fv2_2fcost_2eproto(&descriptor_table_city_2frouting_2fv2_2fcost_2eproto);
namespace city {
namespace routing {
namespace v2 {

// ===================================================================

class Cost::_Internal {
 public:
  using HasBits = decltype(std::declval<Cost>()._impl_._has_bits_);
  static void set_has_time(HasBits* has_bits) {
    (*has_bits)[0] |= 1u;
  }
};

Cost::Cost(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.routing.v2.Cost)
}
Cost::Cost(const Cost& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  Cost* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){from._impl_._has_bits_}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.cost_){}
    , decltype(_impl_.time_){}
    , decltype(_impl_.id_){}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::memcpy(&_impl_.cost_, &from._impl_.cost_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.id_) -
    reinterpret_cast<char*>(&_impl_.cost_)) + sizeof(_impl_.id_));
  // @@protoc_insertion_point(copy_constructor:city.routing.v2.Cost)
}

inline void Cost::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.cost_){0}
    , decltype(_impl_.time_){0}
    , decltype(_impl_.id_){0}
  };
}

Cost::~Cost() {
  // @@protoc_insertion_point(destructor:city.routing.v2.Cost)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void Cost::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
}

void Cost::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void Cost::Clear() {
// @@protoc_insertion_point(message_clear_start:city.routing.v2.Cost)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.cost_ = 0;
  _impl_.time_ = 0;
  _impl_.id_ = 0;
  _impl_._has_bits_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Cost::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  _Internal::HasBits has_bits{};
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
      // double cost = 2 [json_name = "cost"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 17)) {
          _impl_.cost_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // optional double time = 3 [json_name = "time"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 25)) {
          _Internal::set_has_time(&has_bits);
          _impl_.time_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
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

uint8_t* Cost::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.routing.v2.Cost)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 id = 1 [json_name = "id"];
  if (this->_internal_id() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(1, this->_internal_id(), target);
  }

  // double cost = 2 [json_name = "cost"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_cost = this->_internal_cost();
  uint64_t raw_cost;
  memcpy(&raw_cost, &tmp_cost, sizeof(tmp_cost));
  if (raw_cost != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(2, this->_internal_cost(), target);
  }

  // optional double time = 3 [json_name = "time"];
  if (_internal_has_time()) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(3, this->_internal_time(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.routing.v2.Cost)
  return target;
}

size_t Cost::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.routing.v2.Cost)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // double cost = 2 [json_name = "cost"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_cost = this->_internal_cost();
  uint64_t raw_cost;
  memcpy(&raw_cost, &tmp_cost, sizeof(tmp_cost));
  if (raw_cost != 0) {
    total_size += 1 + 8;
  }

  // optional double time = 3 [json_name = "time"];
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000001u) {
    total_size += 1 + 8;
  }

  // int32 id = 1 [json_name = "id"];
  if (this->_internal_id() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_id());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData Cost::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    Cost::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*Cost::GetClassData() const { return &_class_data_; }


void Cost::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<Cost*>(&to_msg);
  auto& from = static_cast<const Cost&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.routing.v2.Cost)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_cost = from._internal_cost();
  uint64_t raw_cost;
  memcpy(&raw_cost, &tmp_cost, sizeof(tmp_cost));
  if (raw_cost != 0) {
    _this->_internal_set_cost(from._internal_cost());
  }
  if (from._internal_has_time()) {
    _this->_internal_set_time(from._internal_time());
  }
  if (from._internal_id() != 0) {
    _this->_internal_set_id(from._internal_id());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void Cost::CopyFrom(const Cost& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.routing.v2.Cost)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Cost::IsInitialized() const {
  return true;
}

void Cost::InternalSwap(Cost* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_._has_bits_[0], other->_impl_._has_bits_[0]);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(Cost, _impl_.id_)
      + sizeof(Cost::_impl_.id_)
      - PROTOBUF_FIELD_OFFSET(Cost, _impl_.cost_)>(
          reinterpret_cast<char*>(&_impl_.cost_),
          reinterpret_cast<char*>(&other->_impl_.cost_));
}

::PROTOBUF_NAMESPACE_ID::Metadata Cost::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2frouting_2fv2_2fcost_2eproto_getter, &descriptor_table_city_2frouting_2fv2_2fcost_2eproto_once,
      file_level_metadata_city_2frouting_2fv2_2fcost_2eproto[0]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v2
}  // namespace routing
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::routing::v2::Cost*
Arena::CreateMaybeMessage< ::city::routing::v2::Cost >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::routing::v2::Cost >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
