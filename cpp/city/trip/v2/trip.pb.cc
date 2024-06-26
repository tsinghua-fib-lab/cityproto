// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/trip/v2/trip.proto

#include "city/trip/v2/trip.pb.h"

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
namespace trip {
namespace v2 {
PROTOBUF_CONSTEXPR Trip::Trip(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_._has_bits_)*/{}
  , /*decltype(_impl_._cached_size_)*/{}
  , /*decltype(_impl_.routes_)*/{}
  , /*decltype(_impl_.activity_)*/{&::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized{}}
  , /*decltype(_impl_.model_)*/{&::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized{}}
  , /*decltype(_impl_.end_)*/nullptr
  , /*decltype(_impl_.departure_time_)*/0
  , /*decltype(_impl_.wait_time_)*/0
  , /*decltype(_impl_.arrival_time_)*/0
  , /*decltype(_impl_.mode_)*/0} {}
struct TripDefaultTypeInternal {
  PROTOBUF_CONSTEXPR TripDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~TripDefaultTypeInternal() {}
  union {
    Trip _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 TripDefaultTypeInternal _Trip_default_instance_;
PROTOBUF_CONSTEXPR Schedule::Schedule(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_._has_bits_)*/{}
  , /*decltype(_impl_._cached_size_)*/{}
  , /*decltype(_impl_.trips_)*/{}
  , /*decltype(_impl_.departure_time_)*/0
  , /*decltype(_impl_.wait_time_)*/0
  , /*decltype(_impl_.loop_count_)*/0} {}
struct ScheduleDefaultTypeInternal {
  PROTOBUF_CONSTEXPR ScheduleDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~ScheduleDefaultTypeInternal() {}
  union {
    Schedule _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 ScheduleDefaultTypeInternal _Schedule_default_instance_;
}  // namespace v2
}  // namespace trip
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2ftrip_2fv2_2ftrip_2eproto[2];
static const ::_pb::EnumDescriptor* file_level_enum_descriptors_city_2ftrip_2fv2_2ftrip_2eproto[1];
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2ftrip_2fv2_2ftrip_2eproto = nullptr;

const uint32_t TableStruct_city_2ftrip_2fv2_2ftrip_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _impl_._has_bits_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _impl_.mode_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _impl_.end_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _impl_.departure_time_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _impl_.wait_time_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _impl_.arrival_time_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _impl_.activity_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _impl_.model_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Trip, _impl_.routes_),
  ~0u,
  ~0u,
  2,
  3,
  4,
  0,
  1,
  ~0u,
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Schedule, _impl_._has_bits_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Schedule, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Schedule, _impl_.trips_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Schedule, _impl_.loop_count_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Schedule, _impl_.departure_time_),
  PROTOBUF_FIELD_OFFSET(::city::trip::v2::Schedule, _impl_.wait_time_),
  ~0u,
  ~0u,
  0,
  1,
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, 14, -1, sizeof(::city::trip::v2::Trip)},
  { 22, 32, -1, sizeof(::city::trip::v2::Schedule)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::trip::v2::_Trip_default_instance_._instance,
  &::city::trip::v2::_Schedule_default_instance_._instance,
};

const char descriptor_table_protodef_city_2ftrip_2fv2_2ftrip_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\027city/trip/v2/trip.proto\022\014city.trip.v2\032"
  "\025city/geo/v2/geo.proto\032\035city/routing/v2/"
  "routing.proto\"\210\003\n\004Trip\022*\n\004mode\030\001 \001(\0162\026.c"
  "ity.trip.v2.TripModeR\004mode\022\'\n\003end\030\002 \001(\0132"
  "\025.city.geo.v2.PositionR\003end\022*\n\016departure"
  "_time\030\003 \001(\001H\000R\rdepartureTime\210\001\001\022 \n\twait_"
  "time\030\004 \001(\001H\001R\010waitTime\210\001\001\022&\n\014arrival_tim"
  "e\030\005 \001(\001H\002R\013arrivalTime\210\001\001\022\037\n\010activity\030\006 "
  "\001(\tH\003R\010activity\210\001\001\022\031\n\005model\030\010 \001(\tH\004R\005mod"
  "el\210\001\001\0220\n\006routes\030\007 \003(\0132\030.city.routing.v2."
  "JourneyR\006routesB\021\n\017_departure_timeB\014\n\n_w"
  "ait_timeB\017\n\r_arrival_timeB\013\n\t_activityB\010"
  "\n\006_model\"\302\001\n\010Schedule\022(\n\005trips\030\001 \003(\0132\022.c"
  "ity.trip.v2.TripR\005trips\022\035\n\nloop_count\030\002 "
  "\001(\005R\tloopCount\022*\n\016departure_time\030\003 \001(\001H\000"
  "R\rdepartureTime\210\001\001\022 \n\twait_time\030\004 \001(\001H\001R"
  "\010waitTime\210\001\001B\021\n\017_departure_timeB\014\n\n_wait"
  "_time*\211\001\n\010TripMode\022\031\n\025TRIP_MODE_UNSPECIF"
  "IED\020\000\022\027\n\023TRIP_MODE_WALK_ONLY\020\001\022\030\n\024TRIP_M"
  "ODE_DRIVE_ONLY\020\002\022\026\n\022TRIP_MODE_BUS_WALK\020\004"
  "\022\027\n\023TRIP_MODE_BIKE_WALK\020\005B\241\001\n\020com.city.t"
  "rip.v2B\tTripProtoP\001Z0git.fiblab.net/sim/"
  "protos/go/city/trip/v2;tripv2\242\002\003CTX\252\002\014Ci"
  "ty.Trip.V2\312\002\014City\\Trip\\V2\342\002\030City\\Trip\\V2"
  "\\GPBMetadata\352\002\016City::Trip::V2b\006proto3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto_deps[2] = {
  &::descriptor_table_city_2fgeo_2fv2_2fgeo_2eproto,
  &::descriptor_table_city_2frouting_2fv2_2frouting_2eproto,
};
static ::_pbi::once_flag descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto = {
    false, false, 997, descriptor_table_protodef_city_2ftrip_2fv2_2ftrip_2eproto,
    "city/trip/v2/trip.proto",
    &descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto_once, descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto_deps, 2, 2,
    schemas, file_default_instances, TableStruct_city_2ftrip_2fv2_2ftrip_2eproto::offsets,
    file_level_metadata_city_2ftrip_2fv2_2ftrip_2eproto, file_level_enum_descriptors_city_2ftrip_2fv2_2ftrip_2eproto,
    file_level_service_descriptors_city_2ftrip_2fv2_2ftrip_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto_getter() {
  return &descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2ftrip_2fv2_2ftrip_2eproto(&descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto);
namespace city {
namespace trip {
namespace v2 {
const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* TripMode_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto);
  return file_level_enum_descriptors_city_2ftrip_2fv2_2ftrip_2eproto[0];
}
bool TripMode_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
    case 2:
    case 4:
    case 5:
      return true;
    default:
      return false;
  }
}


// ===================================================================

class Trip::_Internal {
 public:
  using HasBits = decltype(std::declval<Trip>()._impl_._has_bits_);
  static const ::city::geo::v2::Position& end(const Trip* msg);
  static void set_has_departure_time(HasBits* has_bits) {
    (*has_bits)[0] |= 4u;
  }
  static void set_has_wait_time(HasBits* has_bits) {
    (*has_bits)[0] |= 8u;
  }
  static void set_has_arrival_time(HasBits* has_bits) {
    (*has_bits)[0] |= 16u;
  }
  static void set_has_activity(HasBits* has_bits) {
    (*has_bits)[0] |= 1u;
  }
  static void set_has_model(HasBits* has_bits) {
    (*has_bits)[0] |= 2u;
  }
};

const ::city::geo::v2::Position&
Trip::_Internal::end(const Trip* msg) {
  return *msg->_impl_.end_;
}
void Trip::clear_end() {
  if (GetArenaForAllocation() == nullptr && _impl_.end_ != nullptr) {
    delete _impl_.end_;
  }
  _impl_.end_ = nullptr;
}
void Trip::clear_routes() {
  _impl_.routes_.Clear();
}
Trip::Trip(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.trip.v2.Trip)
}
Trip::Trip(const Trip& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  Trip* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){from._impl_._has_bits_}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.routes_){from._impl_.routes_}
    , decltype(_impl_.activity_){}
    , decltype(_impl_.model_){}
    , decltype(_impl_.end_){nullptr}
    , decltype(_impl_.departure_time_){}
    , decltype(_impl_.wait_time_){}
    , decltype(_impl_.arrival_time_){}
    , decltype(_impl_.mode_){}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.activity_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.activity_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (from._internal_has_activity()) {
    _this->_impl_.activity_.Set(from._internal_activity(), 
      _this->GetArenaForAllocation());
  }
  _impl_.model_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.model_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (from._internal_has_model()) {
    _this->_impl_.model_.Set(from._internal_model(), 
      _this->GetArenaForAllocation());
  }
  if (from._internal_has_end()) {
    _this->_impl_.end_ = new ::city::geo::v2::Position(*from._impl_.end_);
  }
  ::memcpy(&_impl_.departure_time_, &from._impl_.departure_time_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.mode_) -
    reinterpret_cast<char*>(&_impl_.departure_time_)) + sizeof(_impl_.mode_));
  // @@protoc_insertion_point(copy_constructor:city.trip.v2.Trip)
}

inline void Trip::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.routes_){arena}
    , decltype(_impl_.activity_){}
    , decltype(_impl_.model_){}
    , decltype(_impl_.end_){nullptr}
    , decltype(_impl_.departure_time_){0}
    , decltype(_impl_.wait_time_){0}
    , decltype(_impl_.arrival_time_){0}
    , decltype(_impl_.mode_){0}
  };
  _impl_.activity_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.activity_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  _impl_.model_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.model_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

Trip::~Trip() {
  // @@protoc_insertion_point(destructor:city.trip.v2.Trip)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void Trip::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.routes_.~RepeatedPtrField();
  _impl_.activity_.Destroy();
  _impl_.model_.Destroy();
  if (this != internal_default_instance()) delete _impl_.end_;
}

void Trip::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void Trip::Clear() {
// @@protoc_insertion_point(message_clear_start:city.trip.v2.Trip)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.routes_.Clear();
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000003u) {
    if (cached_has_bits & 0x00000001u) {
      _impl_.activity_.ClearNonDefaultToEmpty();
    }
    if (cached_has_bits & 0x00000002u) {
      _impl_.model_.ClearNonDefaultToEmpty();
    }
  }
  if (GetArenaForAllocation() == nullptr && _impl_.end_ != nullptr) {
    delete _impl_.end_;
  }
  _impl_.end_ = nullptr;
  if (cached_has_bits & 0x0000001cu) {
    ::memset(&_impl_.departure_time_, 0, static_cast<size_t>(
        reinterpret_cast<char*>(&_impl_.arrival_time_) -
        reinterpret_cast<char*>(&_impl_.departure_time_)) + sizeof(_impl_.arrival_time_));
  }
  _impl_.mode_ = 0;
  _impl_._has_bits_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Trip::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  _Internal::HasBits has_bits{};
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // .city.trip.v2.TripMode mode = 1 [json_name = "mode"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 8)) {
          uint64_t val = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
          _internal_set_mode(static_cast<::city::trip::v2::TripMode>(val));
        } else
          goto handle_unusual;
        continue;
      // .city.geo.v2.Position end = 2 [json_name = "end"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 18)) {
          ptr = ctx->ParseMessage(_internal_mutable_end(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // optional double departure_time = 3 [json_name = "departureTime"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 25)) {
          _Internal::set_has_departure_time(&has_bits);
          _impl_.departure_time_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // optional double wait_time = 4 [json_name = "waitTime"];
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 33)) {
          _Internal::set_has_wait_time(&has_bits);
          _impl_.wait_time_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // optional double arrival_time = 5 [json_name = "arrivalTime"];
      case 5:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 41)) {
          _Internal::set_has_arrival_time(&has_bits);
          _impl_.arrival_time_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // optional string activity = 6 [json_name = "activity"];
      case 6:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 50)) {
          auto str = _internal_mutable_activity();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "city.trip.v2.Trip.activity"));
        } else
          goto handle_unusual;
        continue;
      // repeated .city.routing.v2.Journey routes = 7 [json_name = "routes"];
      case 7:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 58)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_routes(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<58>(ptr));
        } else
          goto handle_unusual;
        continue;
      // optional string model = 8 [json_name = "model"];
      case 8:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 66)) {
          auto str = _internal_mutable_model();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "city.trip.v2.Trip.model"));
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

uint8_t* Trip::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.trip.v2.Trip)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // .city.trip.v2.TripMode mode = 1 [json_name = "mode"];
  if (this->_internal_mode() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteEnumToArray(
      1, this->_internal_mode(), target);
  }

  // .city.geo.v2.Position end = 2 [json_name = "end"];
  if (this->_internal_has_end()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(2, _Internal::end(this),
        _Internal::end(this).GetCachedSize(), target, stream);
  }

  // optional double departure_time = 3 [json_name = "departureTime"];
  if (_internal_has_departure_time()) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(3, this->_internal_departure_time(), target);
  }

  // optional double wait_time = 4 [json_name = "waitTime"];
  if (_internal_has_wait_time()) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(4, this->_internal_wait_time(), target);
  }

  // optional double arrival_time = 5 [json_name = "arrivalTime"];
  if (_internal_has_arrival_time()) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(5, this->_internal_arrival_time(), target);
  }

  // optional string activity = 6 [json_name = "activity"];
  if (_internal_has_activity()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_activity().data(), static_cast<int>(this->_internal_activity().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "city.trip.v2.Trip.activity");
    target = stream->WriteStringMaybeAliased(
        6, this->_internal_activity(), target);
  }

  // repeated .city.routing.v2.Journey routes = 7 [json_name = "routes"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_routes_size()); i < n; i++) {
    const auto& repfield = this->_internal_routes(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(7, repfield, repfield.GetCachedSize(), target, stream);
  }

  // optional string model = 8 [json_name = "model"];
  if (_internal_has_model()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_model().data(), static_cast<int>(this->_internal_model().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "city.trip.v2.Trip.model");
    target = stream->WriteStringMaybeAliased(
        8, this->_internal_model(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.trip.v2.Trip)
  return target;
}

size_t Trip::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.trip.v2.Trip)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.routing.v2.Journey routes = 7 [json_name = "routes"];
  total_size += 1UL * this->_internal_routes_size();
  for (const auto& msg : this->_impl_.routes_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000003u) {
    // optional string activity = 6 [json_name = "activity"];
    if (cached_has_bits & 0x00000001u) {
      total_size += 1 +
        ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
          this->_internal_activity());
    }

    // optional string model = 8 [json_name = "model"];
    if (cached_has_bits & 0x00000002u) {
      total_size += 1 +
        ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
          this->_internal_model());
    }

  }
  // .city.geo.v2.Position end = 2 [json_name = "end"];
  if (this->_internal_has_end()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.end_);
  }

  if (cached_has_bits & 0x0000001cu) {
    // optional double departure_time = 3 [json_name = "departureTime"];
    if (cached_has_bits & 0x00000004u) {
      total_size += 1 + 8;
    }

    // optional double wait_time = 4 [json_name = "waitTime"];
    if (cached_has_bits & 0x00000008u) {
      total_size += 1 + 8;
    }

    // optional double arrival_time = 5 [json_name = "arrivalTime"];
    if (cached_has_bits & 0x00000010u) {
      total_size += 1 + 8;
    }

  }
  // .city.trip.v2.TripMode mode = 1 [json_name = "mode"];
  if (this->_internal_mode() != 0) {
    total_size += 1 +
      ::_pbi::WireFormatLite::EnumSize(this->_internal_mode());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData Trip::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    Trip::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*Trip::GetClassData() const { return &_class_data_; }


void Trip::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<Trip*>(&to_msg);
  auto& from = static_cast<const Trip&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.trip.v2.Trip)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.routes_.MergeFrom(from._impl_.routes_);
  cached_has_bits = from._impl_._has_bits_[0];
  if (cached_has_bits & 0x00000003u) {
    if (cached_has_bits & 0x00000001u) {
      _this->_internal_set_activity(from._internal_activity());
    }
    if (cached_has_bits & 0x00000002u) {
      _this->_internal_set_model(from._internal_model());
    }
  }
  if (from._internal_has_end()) {
    _this->_internal_mutable_end()->::city::geo::v2::Position::MergeFrom(
        from._internal_end());
  }
  if (cached_has_bits & 0x0000001cu) {
    if (cached_has_bits & 0x00000004u) {
      _this->_impl_.departure_time_ = from._impl_.departure_time_;
    }
    if (cached_has_bits & 0x00000008u) {
      _this->_impl_.wait_time_ = from._impl_.wait_time_;
    }
    if (cached_has_bits & 0x00000010u) {
      _this->_impl_.arrival_time_ = from._impl_.arrival_time_;
    }
    _this->_impl_._has_bits_[0] |= cached_has_bits;
  }
  if (from._internal_mode() != 0) {
    _this->_internal_set_mode(from._internal_mode());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void Trip::CopyFrom(const Trip& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.trip.v2.Trip)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Trip::IsInitialized() const {
  return true;
}

void Trip::InternalSwap(Trip* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_._has_bits_[0], other->_impl_._has_bits_[0]);
  _impl_.routes_.InternalSwap(&other->_impl_.routes_);
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::InternalSwap(
      &_impl_.activity_, lhs_arena,
      &other->_impl_.activity_, rhs_arena
  );
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::InternalSwap(
      &_impl_.model_, lhs_arena,
      &other->_impl_.model_, rhs_arena
  );
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(Trip, _impl_.mode_)
      + sizeof(Trip::_impl_.mode_)
      - PROTOBUF_FIELD_OFFSET(Trip, _impl_.end_)>(
          reinterpret_cast<char*>(&_impl_.end_),
          reinterpret_cast<char*>(&other->_impl_.end_));
}

::PROTOBUF_NAMESPACE_ID::Metadata Trip::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto_getter, &descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto_once,
      file_level_metadata_city_2ftrip_2fv2_2ftrip_2eproto[0]);
}

// ===================================================================

class Schedule::_Internal {
 public:
  using HasBits = decltype(std::declval<Schedule>()._impl_._has_bits_);
  static void set_has_departure_time(HasBits* has_bits) {
    (*has_bits)[0] |= 1u;
  }
  static void set_has_wait_time(HasBits* has_bits) {
    (*has_bits)[0] |= 2u;
  }
};

Schedule::Schedule(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.trip.v2.Schedule)
}
Schedule::Schedule(const Schedule& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  Schedule* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){from._impl_._has_bits_}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.trips_){from._impl_.trips_}
    , decltype(_impl_.departure_time_){}
    , decltype(_impl_.wait_time_){}
    , decltype(_impl_.loop_count_){}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::memcpy(&_impl_.departure_time_, &from._impl_.departure_time_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.loop_count_) -
    reinterpret_cast<char*>(&_impl_.departure_time_)) + sizeof(_impl_.loop_count_));
  // @@protoc_insertion_point(copy_constructor:city.trip.v2.Schedule)
}

inline void Schedule::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){}
    , /*decltype(_impl_._cached_size_)*/{}
    , decltype(_impl_.trips_){arena}
    , decltype(_impl_.departure_time_){0}
    , decltype(_impl_.wait_time_){0}
    , decltype(_impl_.loop_count_){0}
  };
}

Schedule::~Schedule() {
  // @@protoc_insertion_point(destructor:city.trip.v2.Schedule)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void Schedule::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.trips_.~RepeatedPtrField();
}

void Schedule::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void Schedule::Clear() {
// @@protoc_insertion_point(message_clear_start:city.trip.v2.Schedule)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.trips_.Clear();
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000003u) {
    ::memset(&_impl_.departure_time_, 0, static_cast<size_t>(
        reinterpret_cast<char*>(&_impl_.wait_time_) -
        reinterpret_cast<char*>(&_impl_.departure_time_)) + sizeof(_impl_.wait_time_));
  }
  _impl_.loop_count_ = 0;
  _impl_._has_bits_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Schedule::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  _Internal::HasBits has_bits{};
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // repeated .city.trip.v2.Trip trips = 1 [json_name = "trips"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_trips(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<10>(ptr));
        } else
          goto handle_unusual;
        continue;
      // int32 loop_count = 2 [json_name = "loopCount"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 16)) {
          _impl_.loop_count_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // optional double departure_time = 3 [json_name = "departureTime"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 25)) {
          _Internal::set_has_departure_time(&has_bits);
          _impl_.departure_time_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // optional double wait_time = 4 [json_name = "waitTime"];
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 33)) {
          _Internal::set_has_wait_time(&has_bits);
          _impl_.wait_time_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
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

uint8_t* Schedule::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.trip.v2.Schedule)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // repeated .city.trip.v2.Trip trips = 1 [json_name = "trips"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_trips_size()); i < n; i++) {
    const auto& repfield = this->_internal_trips(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(1, repfield, repfield.GetCachedSize(), target, stream);
  }

  // int32 loop_count = 2 [json_name = "loopCount"];
  if (this->_internal_loop_count() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(2, this->_internal_loop_count(), target);
  }

  // optional double departure_time = 3 [json_name = "departureTime"];
  if (_internal_has_departure_time()) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(3, this->_internal_departure_time(), target);
  }

  // optional double wait_time = 4 [json_name = "waitTime"];
  if (_internal_has_wait_time()) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(4, this->_internal_wait_time(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.trip.v2.Schedule)
  return target;
}

size_t Schedule::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.trip.v2.Schedule)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.trip.v2.Trip trips = 1 [json_name = "trips"];
  total_size += 1UL * this->_internal_trips_size();
  for (const auto& msg : this->_impl_.trips_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x00000003u) {
    // optional double departure_time = 3 [json_name = "departureTime"];
    if (cached_has_bits & 0x00000001u) {
      total_size += 1 + 8;
    }

    // optional double wait_time = 4 [json_name = "waitTime"];
    if (cached_has_bits & 0x00000002u) {
      total_size += 1 + 8;
    }

  }
  // int32 loop_count = 2 [json_name = "loopCount"];
  if (this->_internal_loop_count() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_loop_count());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData Schedule::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    Schedule::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*Schedule::GetClassData() const { return &_class_data_; }


void Schedule::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<Schedule*>(&to_msg);
  auto& from = static_cast<const Schedule&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.trip.v2.Schedule)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.trips_.MergeFrom(from._impl_.trips_);
  cached_has_bits = from._impl_._has_bits_[0];
  if (cached_has_bits & 0x00000003u) {
    if (cached_has_bits & 0x00000001u) {
      _this->_impl_.departure_time_ = from._impl_.departure_time_;
    }
    if (cached_has_bits & 0x00000002u) {
      _this->_impl_.wait_time_ = from._impl_.wait_time_;
    }
    _this->_impl_._has_bits_[0] |= cached_has_bits;
  }
  if (from._internal_loop_count() != 0) {
    _this->_internal_set_loop_count(from._internal_loop_count());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void Schedule::CopyFrom(const Schedule& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.trip.v2.Schedule)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Schedule::IsInitialized() const {
  return true;
}

void Schedule::InternalSwap(Schedule* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_._has_bits_[0], other->_impl_._has_bits_[0]);
  _impl_.trips_.InternalSwap(&other->_impl_.trips_);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(Schedule, _impl_.loop_count_)
      + sizeof(Schedule::_impl_.loop_count_)
      - PROTOBUF_FIELD_OFFSET(Schedule, _impl_.departure_time_)>(
          reinterpret_cast<char*>(&_impl_.departure_time_),
          reinterpret_cast<char*>(&other->_impl_.departure_time_));
}

::PROTOBUF_NAMESPACE_ID::Metadata Schedule::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto_getter, &descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto_once,
      file_level_metadata_city_2ftrip_2fv2_2ftrip_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v2
}  // namespace trip
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::trip::v2::Trip*
Arena::CreateMaybeMessage< ::city::trip::v2::Trip >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::trip::v2::Trip >(arena);
}
template<> PROTOBUF_NOINLINE ::city::trip::v2::Schedule*
Arena::CreateMaybeMessage< ::city::trip::v2::Schedule >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::trip::v2::Schedule >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
