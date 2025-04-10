// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/person/v2/carbon.proto

#include "city/person/v2/carbon.pb.h"

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
namespace v2 {
PROTOBUF_CONSTEXPR VehicleCarbon::VehicleCarbon(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.ds_)*/0
  , /*decltype(_impl_.v_)*/0
  , /*decltype(_impl_.a_)*/0
  , /*decltype(_impl_.u_acc_)*/0
  , /*decltype(_impl_.u_roll_)*/0
  , /*decltype(_impl_.u_aero_)*/0
  , /*decltype(_impl_.c_d_)*/0
  , /*decltype(_impl_.id_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct VehicleCarbonDefaultTypeInternal {
  PROTOBUF_CONSTEXPR VehicleCarbonDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~VehicleCarbonDefaultTypeInternal() {}
  union {
    VehicleCarbon _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 VehicleCarbonDefaultTypeInternal _VehicleCarbon_default_instance_;
PROTOBUF_CONSTEXPR EmissionStatistics::EmissionStatistics(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.u_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct EmissionStatisticsDefaultTypeInternal {
  PROTOBUF_CONSTEXPR EmissionStatisticsDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~EmissionStatisticsDefaultTypeInternal() {}
  union {
    EmissionStatistics _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 EmissionStatisticsDefaultTypeInternal _EmissionStatistics_default_instance_;
}  // namespace v2
}  // namespace person
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fperson_2fv2_2fcarbon_2eproto[2];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_city_2fperson_2fv2_2fcarbon_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fperson_2fv2_2fcarbon_2eproto = nullptr;

const uint32_t TableStruct_city_2fperson_2fv2_2fcarbon_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::person::v2::VehicleCarbon, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::person::v2::VehicleCarbon, _impl_.id_),
  PROTOBUF_FIELD_OFFSET(::city::person::v2::VehicleCarbon, _impl_.ds_),
  PROTOBUF_FIELD_OFFSET(::city::person::v2::VehicleCarbon, _impl_.v_),
  PROTOBUF_FIELD_OFFSET(::city::person::v2::VehicleCarbon, _impl_.a_),
  PROTOBUF_FIELD_OFFSET(::city::person::v2::VehicleCarbon, _impl_.u_acc_),
  PROTOBUF_FIELD_OFFSET(::city::person::v2::VehicleCarbon, _impl_.u_roll_),
  PROTOBUF_FIELD_OFFSET(::city::person::v2::VehicleCarbon, _impl_.u_aero_),
  PROTOBUF_FIELD_OFFSET(::city::person::v2::VehicleCarbon, _impl_.c_d_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::person::v2::EmissionStatistics, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::person::v2::EmissionStatistics, _impl_.u_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::city::person::v2::VehicleCarbon)},
  { 14, -1, -1, sizeof(::city::person::v2::EmissionStatistics)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::person::v2::_VehicleCarbon_default_instance_._instance,
  &::city::person::v2::_EmissionStatistics_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fperson_2fv2_2fcarbon_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\033city/person/v2/carbon.proto\022\016city.pers"
  "on.v2\"\237\001\n\rVehicleCarbon\022\016\n\002id\030\001 \001(\005R\002id\022"
  "\016\n\002ds\030\002 \001(\001R\002ds\022\014\n\001v\030\003 \001(\001R\001v\022\014\n\001a\030\004 \001(\001"
  "R\001a\022\023\n\005u_acc\030\005 \001(\001R\004uAcc\022\025\n\006u_roll\030\006 \001(\001"
  "R\005uRoll\022\025\n\006u_aero\030\007 \001(\001R\005uAero\022\017\n\003c_d\030\010 "
  "\001(\001R\002cD\"\"\n\022EmissionStatistics\022\014\n\001u\030\001 \001(\001"
  "R\001uB\264\001\n\022com.city.person.v2B\013CarbonProtoP"
  "\001Z7git.fiblab.net/sim/protos/v2/go/city/"
  "person/v2;personv2\242\002\003CPX\252\002\016City.Person.V"
  "2\312\002\016City\\Person\\V2\342\002\032City\\Person\\V2\\GPBM"
  "etadata\352\002\020City::Person::V2b\006proto3"
  ;
static ::_pbi::once_flag descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto = {
    false, false, 434, descriptor_table_protodef_city_2fperson_2fv2_2fcarbon_2eproto,
    "city/person/v2/carbon.proto",
    &descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto_once, nullptr, 0, 2,
    schemas, file_default_instances, TableStruct_city_2fperson_2fv2_2fcarbon_2eproto::offsets,
    file_level_metadata_city_2fperson_2fv2_2fcarbon_2eproto, file_level_enum_descriptors_city_2fperson_2fv2_2fcarbon_2eproto,
    file_level_service_descriptors_city_2fperson_2fv2_2fcarbon_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto_getter() {
  return &descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fperson_2fv2_2fcarbon_2eproto(&descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto);
namespace city {
namespace person {
namespace v2 {

// ===================================================================

class VehicleCarbon::_Internal {
 public:
};

VehicleCarbon::VehicleCarbon(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.person.v2.VehicleCarbon)
}
VehicleCarbon::VehicleCarbon(const VehicleCarbon& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  VehicleCarbon* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.ds_){}
    , decltype(_impl_.v_){}
    , decltype(_impl_.a_){}
    , decltype(_impl_.u_acc_){}
    , decltype(_impl_.u_roll_){}
    , decltype(_impl_.u_aero_){}
    , decltype(_impl_.c_d_){}
    , decltype(_impl_.id_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::memcpy(&_impl_.ds_, &from._impl_.ds_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.id_) -
    reinterpret_cast<char*>(&_impl_.ds_)) + sizeof(_impl_.id_));
  // @@protoc_insertion_point(copy_constructor:city.person.v2.VehicleCarbon)
}

inline void VehicleCarbon::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.ds_){0}
    , decltype(_impl_.v_){0}
    , decltype(_impl_.a_){0}
    , decltype(_impl_.u_acc_){0}
    , decltype(_impl_.u_roll_){0}
    , decltype(_impl_.u_aero_){0}
    , decltype(_impl_.c_d_){0}
    , decltype(_impl_.id_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

VehicleCarbon::~VehicleCarbon() {
  // @@protoc_insertion_point(destructor:city.person.v2.VehicleCarbon)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void VehicleCarbon::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
}

void VehicleCarbon::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void VehicleCarbon::Clear() {
// @@protoc_insertion_point(message_clear_start:city.person.v2.VehicleCarbon)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  ::memset(&_impl_.ds_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&_impl_.id_) -
      reinterpret_cast<char*>(&_impl_.ds_)) + sizeof(_impl_.id_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* VehicleCarbon::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
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
      // double ds = 2 [json_name = "ds"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 17)) {
          _impl_.ds_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // double v = 3 [json_name = "v"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 25)) {
          _impl_.v_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // double a = 4 [json_name = "a"];
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 33)) {
          _impl_.a_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // double u_acc = 5 [json_name = "uAcc"];
      case 5:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 41)) {
          _impl_.u_acc_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // double u_roll = 6 [json_name = "uRoll"];
      case 6:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 49)) {
          _impl_.u_roll_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // double u_aero = 7 [json_name = "uAero"];
      case 7:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 57)) {
          _impl_.u_aero_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // double c_d = 8 [json_name = "cD"];
      case 8:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 65)) {
          _impl_.c_d_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
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

uint8_t* VehicleCarbon::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.person.v2.VehicleCarbon)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 id = 1 [json_name = "id"];
  if (this->_internal_id() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(1, this->_internal_id(), target);
  }

  // double ds = 2 [json_name = "ds"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_ds = this->_internal_ds();
  uint64_t raw_ds;
  memcpy(&raw_ds, &tmp_ds, sizeof(tmp_ds));
  if (raw_ds != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(2, this->_internal_ds(), target);
  }

  // double v = 3 [json_name = "v"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_v = this->_internal_v();
  uint64_t raw_v;
  memcpy(&raw_v, &tmp_v, sizeof(tmp_v));
  if (raw_v != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(3, this->_internal_v(), target);
  }

  // double a = 4 [json_name = "a"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_a = this->_internal_a();
  uint64_t raw_a;
  memcpy(&raw_a, &tmp_a, sizeof(tmp_a));
  if (raw_a != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(4, this->_internal_a(), target);
  }

  // double u_acc = 5 [json_name = "uAcc"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u_acc = this->_internal_u_acc();
  uint64_t raw_u_acc;
  memcpy(&raw_u_acc, &tmp_u_acc, sizeof(tmp_u_acc));
  if (raw_u_acc != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(5, this->_internal_u_acc(), target);
  }

  // double u_roll = 6 [json_name = "uRoll"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u_roll = this->_internal_u_roll();
  uint64_t raw_u_roll;
  memcpy(&raw_u_roll, &tmp_u_roll, sizeof(tmp_u_roll));
  if (raw_u_roll != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(6, this->_internal_u_roll(), target);
  }

  // double u_aero = 7 [json_name = "uAero"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u_aero = this->_internal_u_aero();
  uint64_t raw_u_aero;
  memcpy(&raw_u_aero, &tmp_u_aero, sizeof(tmp_u_aero));
  if (raw_u_aero != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(7, this->_internal_u_aero(), target);
  }

  // double c_d = 8 [json_name = "cD"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_c_d = this->_internal_c_d();
  uint64_t raw_c_d;
  memcpy(&raw_c_d, &tmp_c_d, sizeof(tmp_c_d));
  if (raw_c_d != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(8, this->_internal_c_d(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.person.v2.VehicleCarbon)
  return target;
}

size_t VehicleCarbon::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.person.v2.VehicleCarbon)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // double ds = 2 [json_name = "ds"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_ds = this->_internal_ds();
  uint64_t raw_ds;
  memcpy(&raw_ds, &tmp_ds, sizeof(tmp_ds));
  if (raw_ds != 0) {
    total_size += 1 + 8;
  }

  // double v = 3 [json_name = "v"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_v = this->_internal_v();
  uint64_t raw_v;
  memcpy(&raw_v, &tmp_v, sizeof(tmp_v));
  if (raw_v != 0) {
    total_size += 1 + 8;
  }

  // double a = 4 [json_name = "a"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_a = this->_internal_a();
  uint64_t raw_a;
  memcpy(&raw_a, &tmp_a, sizeof(tmp_a));
  if (raw_a != 0) {
    total_size += 1 + 8;
  }

  // double u_acc = 5 [json_name = "uAcc"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u_acc = this->_internal_u_acc();
  uint64_t raw_u_acc;
  memcpy(&raw_u_acc, &tmp_u_acc, sizeof(tmp_u_acc));
  if (raw_u_acc != 0) {
    total_size += 1 + 8;
  }

  // double u_roll = 6 [json_name = "uRoll"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u_roll = this->_internal_u_roll();
  uint64_t raw_u_roll;
  memcpy(&raw_u_roll, &tmp_u_roll, sizeof(tmp_u_roll));
  if (raw_u_roll != 0) {
    total_size += 1 + 8;
  }

  // double u_aero = 7 [json_name = "uAero"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u_aero = this->_internal_u_aero();
  uint64_t raw_u_aero;
  memcpy(&raw_u_aero, &tmp_u_aero, sizeof(tmp_u_aero));
  if (raw_u_aero != 0) {
    total_size += 1 + 8;
  }

  // double c_d = 8 [json_name = "cD"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_c_d = this->_internal_c_d();
  uint64_t raw_c_d;
  memcpy(&raw_c_d, &tmp_c_d, sizeof(tmp_c_d));
  if (raw_c_d != 0) {
    total_size += 1 + 8;
  }

  // int32 id = 1 [json_name = "id"];
  if (this->_internal_id() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_id());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData VehicleCarbon::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    VehicleCarbon::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*VehicleCarbon::GetClassData() const { return &_class_data_; }


void VehicleCarbon::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<VehicleCarbon*>(&to_msg);
  auto& from = static_cast<const VehicleCarbon&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.person.v2.VehicleCarbon)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_ds = from._internal_ds();
  uint64_t raw_ds;
  memcpy(&raw_ds, &tmp_ds, sizeof(tmp_ds));
  if (raw_ds != 0) {
    _this->_internal_set_ds(from._internal_ds());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_v = from._internal_v();
  uint64_t raw_v;
  memcpy(&raw_v, &tmp_v, sizeof(tmp_v));
  if (raw_v != 0) {
    _this->_internal_set_v(from._internal_v());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_a = from._internal_a();
  uint64_t raw_a;
  memcpy(&raw_a, &tmp_a, sizeof(tmp_a));
  if (raw_a != 0) {
    _this->_internal_set_a(from._internal_a());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u_acc = from._internal_u_acc();
  uint64_t raw_u_acc;
  memcpy(&raw_u_acc, &tmp_u_acc, sizeof(tmp_u_acc));
  if (raw_u_acc != 0) {
    _this->_internal_set_u_acc(from._internal_u_acc());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u_roll = from._internal_u_roll();
  uint64_t raw_u_roll;
  memcpy(&raw_u_roll, &tmp_u_roll, sizeof(tmp_u_roll));
  if (raw_u_roll != 0) {
    _this->_internal_set_u_roll(from._internal_u_roll());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u_aero = from._internal_u_aero();
  uint64_t raw_u_aero;
  memcpy(&raw_u_aero, &tmp_u_aero, sizeof(tmp_u_aero));
  if (raw_u_aero != 0) {
    _this->_internal_set_u_aero(from._internal_u_aero());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_c_d = from._internal_c_d();
  uint64_t raw_c_d;
  memcpy(&raw_c_d, &tmp_c_d, sizeof(tmp_c_d));
  if (raw_c_d != 0) {
    _this->_internal_set_c_d(from._internal_c_d());
  }
  if (from._internal_id() != 0) {
    _this->_internal_set_id(from._internal_id());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void VehicleCarbon::CopyFrom(const VehicleCarbon& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.person.v2.VehicleCarbon)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool VehicleCarbon::IsInitialized() const {
  return true;
}

void VehicleCarbon::InternalSwap(VehicleCarbon* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(VehicleCarbon, _impl_.id_)
      + sizeof(VehicleCarbon::_impl_.id_)
      - PROTOBUF_FIELD_OFFSET(VehicleCarbon, _impl_.ds_)>(
          reinterpret_cast<char*>(&_impl_.ds_),
          reinterpret_cast<char*>(&other->_impl_.ds_));
}

::PROTOBUF_NAMESPACE_ID::Metadata VehicleCarbon::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto_getter, &descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto_once,
      file_level_metadata_city_2fperson_2fv2_2fcarbon_2eproto[0]);
}

// ===================================================================

class EmissionStatistics::_Internal {
 public:
};

EmissionStatistics::EmissionStatistics(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.person.v2.EmissionStatistics)
}
EmissionStatistics::EmissionStatistics(const EmissionStatistics& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  EmissionStatistics* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.u_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _this->_impl_.u_ = from._impl_.u_;
  // @@protoc_insertion_point(copy_constructor:city.person.v2.EmissionStatistics)
}

inline void EmissionStatistics::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.u_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

EmissionStatistics::~EmissionStatistics() {
  // @@protoc_insertion_point(destructor:city.person.v2.EmissionStatistics)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void EmissionStatistics::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
}

void EmissionStatistics::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void EmissionStatistics::Clear() {
// @@protoc_insertion_point(message_clear_start:city.person.v2.EmissionStatistics)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.u_ = 0;
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* EmissionStatistics::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // double u = 1 [json_name = "u"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 9)) {
          _impl_.u_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
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

uint8_t* EmissionStatistics::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.person.v2.EmissionStatistics)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // double u = 1 [json_name = "u"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u = this->_internal_u();
  uint64_t raw_u;
  memcpy(&raw_u, &tmp_u, sizeof(tmp_u));
  if (raw_u != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(1, this->_internal_u(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.person.v2.EmissionStatistics)
  return target;
}

size_t EmissionStatistics::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.person.v2.EmissionStatistics)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // double u = 1 [json_name = "u"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u = this->_internal_u();
  uint64_t raw_u;
  memcpy(&raw_u, &tmp_u, sizeof(tmp_u));
  if (raw_u != 0) {
    total_size += 1 + 8;
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData EmissionStatistics::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    EmissionStatistics::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*EmissionStatistics::GetClassData() const { return &_class_data_; }


void EmissionStatistics::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<EmissionStatistics*>(&to_msg);
  auto& from = static_cast<const EmissionStatistics&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.person.v2.EmissionStatistics)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_u = from._internal_u();
  uint64_t raw_u;
  memcpy(&raw_u, &tmp_u, sizeof(tmp_u));
  if (raw_u != 0) {
    _this->_internal_set_u(from._internal_u());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void EmissionStatistics::CopyFrom(const EmissionStatistics& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.person.v2.EmissionStatistics)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool EmissionStatistics::IsInitialized() const {
  return true;
}

void EmissionStatistics::InternalSwap(EmissionStatistics* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_.u_, other->_impl_.u_);
}

::PROTOBUF_NAMESPACE_ID::Metadata EmissionStatistics::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto_getter, &descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto_once,
      file_level_metadata_city_2fperson_2fv2_2fcarbon_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v2
}  // namespace person
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::person::v2::VehicleCarbon*
Arena::CreateMaybeMessage< ::city::person::v2::VehicleCarbon >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::person::v2::VehicleCarbon >(arena);
}
template<> PROTOBUF_NOINLINE ::city::person::v2::EmissionStatistics*
Arena::CreateMaybeMessage< ::city::person::v2::EmissionStatistics >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::person::v2::EmissionStatistics >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
