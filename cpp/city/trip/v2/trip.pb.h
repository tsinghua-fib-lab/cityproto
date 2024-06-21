// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/trip/v2/trip.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2ftrip_2fv2_2ftrip_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2ftrip_2fv2_2ftrip_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3021000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3021012 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
#include "city/geo/v2/geo.pb.h"
#include "city/routing/v2/routing.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_city_2ftrip_2fv2_2ftrip_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2ftrip_2fv2_2ftrip_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2ftrip_2fv2_2ftrip_2eproto;
namespace city {
namespace trip {
namespace v2 {
class Schedule;
struct ScheduleDefaultTypeInternal;
extern ScheduleDefaultTypeInternal _Schedule_default_instance_;
class Trip;
struct TripDefaultTypeInternal;
extern TripDefaultTypeInternal _Trip_default_instance_;
}  // namespace v2
}  // namespace trip
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::trip::v2::Schedule* Arena::CreateMaybeMessage<::city::trip::v2::Schedule>(Arena*);
template<> ::city::trip::v2::Trip* Arena::CreateMaybeMessage<::city::trip::v2::Trip>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace trip {
namespace v2 {

enum TripMode : int {
  TRIP_MODE_UNSPECIFIED = 0,
  TRIP_MODE_WALK_ONLY = 1,
  TRIP_MODE_DRIVE_ONLY = 2,
  TRIP_MODE_BUS_WALK = 4,
  TRIP_MODE_BIKE_WALK = 5,
  TripMode_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::min(),
  TripMode_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::max()
};
bool TripMode_IsValid(int value);
constexpr TripMode TripMode_MIN = TRIP_MODE_UNSPECIFIED;
constexpr TripMode TripMode_MAX = TRIP_MODE_BIKE_WALK;
constexpr int TripMode_ARRAYSIZE = TripMode_MAX + 1;

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* TripMode_descriptor();
template<typename T>
inline const std::string& TripMode_Name(T enum_t_value) {
  static_assert(::std::is_same<T, TripMode>::value ||
    ::std::is_integral<T>::value,
    "Incorrect type passed to function TripMode_Name.");
  return ::PROTOBUF_NAMESPACE_ID::internal::NameOfEnum(
    TripMode_descriptor(), enum_t_value);
}
inline bool TripMode_Parse(
    ::PROTOBUF_NAMESPACE_ID::ConstStringParam name, TripMode* value) {
  return ::PROTOBUF_NAMESPACE_ID::internal::ParseNamedEnum<TripMode>(
    TripMode_descriptor(), name, value);
}
// ===================================================================

class Trip final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.trip.v2.Trip) */ {
 public:
  inline Trip() : Trip(nullptr) {}
  ~Trip() override;
  explicit PROTOBUF_CONSTEXPR Trip(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  Trip(const Trip& from);
  Trip(Trip&& from) noexcept
    : Trip() {
    *this = ::std::move(from);
  }

  inline Trip& operator=(const Trip& from) {
    CopyFrom(from);
    return *this;
  }
  inline Trip& operator=(Trip&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const Trip& default_instance() {
    return *internal_default_instance();
  }
  static inline const Trip* internal_default_instance() {
    return reinterpret_cast<const Trip*>(
               &_Trip_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(Trip& a, Trip& b) {
    a.Swap(&b);
  }
  inline void Swap(Trip* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(Trip* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Trip* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<Trip>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const Trip& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const Trip& from) {
    Trip::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::PROTOBUF_NAMESPACE_ID::Arena* arena, bool is_message_owned);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(Trip* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.trip.v2.Trip";
  }
  protected:
  explicit Trip(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kRoutesFieldNumber = 7,
    kActivityFieldNumber = 6,
    kModelFieldNumber = 8,
    kEndFieldNumber = 2,
    kDepartureTimeFieldNumber = 3,
    kWaitTimeFieldNumber = 4,
    kArrivalTimeFieldNumber = 5,
    kModeFieldNumber = 1,
  };
  // repeated .city.routing.v2.Journey routes = 7 [json_name = "routes"];
  int routes_size() const;
  private:
  int _internal_routes_size() const;
  public:
  void clear_routes();
  ::city::routing::v2::Journey* mutable_routes(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::routing::v2::Journey >*
      mutable_routes();
  private:
  const ::city::routing::v2::Journey& _internal_routes(int index) const;
  ::city::routing::v2::Journey* _internal_add_routes();
  public:
  const ::city::routing::v2::Journey& routes(int index) const;
  ::city::routing::v2::Journey* add_routes();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::routing::v2::Journey >&
      routes() const;

  // optional string activity = 6 [json_name = "activity"];
  bool has_activity() const;
  private:
  bool _internal_has_activity() const;
  public:
  void clear_activity();
  const std::string& activity() const;
  template <typename ArgT0 = const std::string&, typename... ArgT>
  void set_activity(ArgT0&& arg0, ArgT... args);
  std::string* mutable_activity();
  PROTOBUF_NODISCARD std::string* release_activity();
  void set_allocated_activity(std::string* activity);
  private:
  const std::string& _internal_activity() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_activity(const std::string& value);
  std::string* _internal_mutable_activity();
  public:

  // optional string model = 8 [json_name = "model"];
  bool has_model() const;
  private:
  bool _internal_has_model() const;
  public:
  void clear_model();
  const std::string& model() const;
  template <typename ArgT0 = const std::string&, typename... ArgT>
  void set_model(ArgT0&& arg0, ArgT... args);
  std::string* mutable_model();
  PROTOBUF_NODISCARD std::string* release_model();
  void set_allocated_model(std::string* model);
  private:
  const std::string& _internal_model() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_model(const std::string& value);
  std::string* _internal_mutable_model();
  public:

  // .city.geo.v2.Position end = 2 [json_name = "end"];
  bool has_end() const;
  private:
  bool _internal_has_end() const;
  public:
  void clear_end();
  const ::city::geo::v2::Position& end() const;
  PROTOBUF_NODISCARD ::city::geo::v2::Position* release_end();
  ::city::geo::v2::Position* mutable_end();
  void set_allocated_end(::city::geo::v2::Position* end);
  private:
  const ::city::geo::v2::Position& _internal_end() const;
  ::city::geo::v2::Position* _internal_mutable_end();
  public:
  void unsafe_arena_set_allocated_end(
      ::city::geo::v2::Position* end);
  ::city::geo::v2::Position* unsafe_arena_release_end();

  // optional double departure_time = 3 [json_name = "departureTime"];
  bool has_departure_time() const;
  private:
  bool _internal_has_departure_time() const;
  public:
  void clear_departure_time();
  double departure_time() const;
  void set_departure_time(double value);
  private:
  double _internal_departure_time() const;
  void _internal_set_departure_time(double value);
  public:

  // optional double wait_time = 4 [json_name = "waitTime"];
  bool has_wait_time() const;
  private:
  bool _internal_has_wait_time() const;
  public:
  void clear_wait_time();
  double wait_time() const;
  void set_wait_time(double value);
  private:
  double _internal_wait_time() const;
  void _internal_set_wait_time(double value);
  public:

  // optional double arrival_time = 5 [json_name = "arrivalTime"];
  bool has_arrival_time() const;
  private:
  bool _internal_has_arrival_time() const;
  public:
  void clear_arrival_time();
  double arrival_time() const;
  void set_arrival_time(double value);
  private:
  double _internal_arrival_time() const;
  void _internal_set_arrival_time(double value);
  public:

  // .city.trip.v2.TripMode mode = 1 [json_name = "mode"];
  void clear_mode();
  ::city::trip::v2::TripMode mode() const;
  void set_mode(::city::trip::v2::TripMode value);
  private:
  ::city::trip::v2::TripMode _internal_mode() const;
  void _internal_set_mode(::city::trip::v2::TripMode value);
  public:

  // @@protoc_insertion_point(class_scope:city.trip.v2.Trip)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::internal::HasBits<1> _has_bits_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
    ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::routing::v2::Journey > routes_;
    ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr activity_;
    ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr model_;
    ::city::geo::v2::Position* end_;
    double departure_time_;
    double wait_time_;
    double arrival_time_;
    int mode_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2ftrip_2fv2_2ftrip_2eproto;
};
// -------------------------------------------------------------------

class Schedule final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.trip.v2.Schedule) */ {
 public:
  inline Schedule() : Schedule(nullptr) {}
  ~Schedule() override;
  explicit PROTOBUF_CONSTEXPR Schedule(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  Schedule(const Schedule& from);
  Schedule(Schedule&& from) noexcept
    : Schedule() {
    *this = ::std::move(from);
  }

  inline Schedule& operator=(const Schedule& from) {
    CopyFrom(from);
    return *this;
  }
  inline Schedule& operator=(Schedule&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const Schedule& default_instance() {
    return *internal_default_instance();
  }
  static inline const Schedule* internal_default_instance() {
    return reinterpret_cast<const Schedule*>(
               &_Schedule_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(Schedule& a, Schedule& b) {
    a.Swap(&b);
  }
  inline void Swap(Schedule* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(Schedule* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Schedule* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<Schedule>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const Schedule& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const Schedule& from) {
    Schedule::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::PROTOBUF_NAMESPACE_ID::Arena* arena, bool is_message_owned);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(Schedule* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.trip.v2.Schedule";
  }
  protected:
  explicit Schedule(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kTripsFieldNumber = 1,
    kDepartureTimeFieldNumber = 3,
    kWaitTimeFieldNumber = 4,
    kLoopCountFieldNumber = 2,
  };
  // repeated .city.trip.v2.Trip trips = 1 [json_name = "trips"];
  int trips_size() const;
  private:
  int _internal_trips_size() const;
  public:
  void clear_trips();
  ::city::trip::v2::Trip* mutable_trips(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::trip::v2::Trip >*
      mutable_trips();
  private:
  const ::city::trip::v2::Trip& _internal_trips(int index) const;
  ::city::trip::v2::Trip* _internal_add_trips();
  public:
  const ::city::trip::v2::Trip& trips(int index) const;
  ::city::trip::v2::Trip* add_trips();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::trip::v2::Trip >&
      trips() const;

  // optional double departure_time = 3 [json_name = "departureTime"];
  bool has_departure_time() const;
  private:
  bool _internal_has_departure_time() const;
  public:
  void clear_departure_time();
  double departure_time() const;
  void set_departure_time(double value);
  private:
  double _internal_departure_time() const;
  void _internal_set_departure_time(double value);
  public:

  // optional double wait_time = 4 [json_name = "waitTime"];
  bool has_wait_time() const;
  private:
  bool _internal_has_wait_time() const;
  public:
  void clear_wait_time();
  double wait_time() const;
  void set_wait_time(double value);
  private:
  double _internal_wait_time() const;
  void _internal_set_wait_time(double value);
  public:

  // int32 loop_count = 2 [json_name = "loopCount"];
  void clear_loop_count();
  int32_t loop_count() const;
  void set_loop_count(int32_t value);
  private:
  int32_t _internal_loop_count() const;
  void _internal_set_loop_count(int32_t value);
  public:

  // @@protoc_insertion_point(class_scope:city.trip.v2.Schedule)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::internal::HasBits<1> _has_bits_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
    ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::trip::v2::Trip > trips_;
    double departure_time_;
    double wait_time_;
    int32_t loop_count_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2ftrip_2fv2_2ftrip_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// Trip

// .city.trip.v2.TripMode mode = 1 [json_name = "mode"];
inline void Trip::clear_mode() {
  _impl_.mode_ = 0;
}
inline ::city::trip::v2::TripMode Trip::_internal_mode() const {
  return static_cast< ::city::trip::v2::TripMode >(_impl_.mode_);
}
inline ::city::trip::v2::TripMode Trip::mode() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Trip.mode)
  return _internal_mode();
}
inline void Trip::_internal_set_mode(::city::trip::v2::TripMode value) {
  
  _impl_.mode_ = value;
}
inline void Trip::set_mode(::city::trip::v2::TripMode value) {
  _internal_set_mode(value);
  // @@protoc_insertion_point(field_set:city.trip.v2.Trip.mode)
}

// .city.geo.v2.Position end = 2 [json_name = "end"];
inline bool Trip::_internal_has_end() const {
  return this != internal_default_instance() && _impl_.end_ != nullptr;
}
inline bool Trip::has_end() const {
  return _internal_has_end();
}
inline const ::city::geo::v2::Position& Trip::_internal_end() const {
  const ::city::geo::v2::Position* p = _impl_.end_;
  return p != nullptr ? *p : reinterpret_cast<const ::city::geo::v2::Position&>(
      ::city::geo::v2::_Position_default_instance_);
}
inline const ::city::geo::v2::Position& Trip::end() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Trip.end)
  return _internal_end();
}
inline void Trip::unsafe_arena_set_allocated_end(
    ::city::geo::v2::Position* end) {
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.end_);
  }
  _impl_.end_ = end;
  if (end) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:city.trip.v2.Trip.end)
}
inline ::city::geo::v2::Position* Trip::release_end() {
  
  ::city::geo::v2::Position* temp = _impl_.end_;
  _impl_.end_ = nullptr;
#ifdef PROTOBUF_FORCE_COPY_IN_RELEASE
  auto* old =  reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(temp);
  temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
  if (GetArenaForAllocation() == nullptr) { delete old; }
#else  // PROTOBUF_FORCE_COPY_IN_RELEASE
  if (GetArenaForAllocation() != nullptr) {
    temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
  }
#endif  // !PROTOBUF_FORCE_COPY_IN_RELEASE
  return temp;
}
inline ::city::geo::v2::Position* Trip::unsafe_arena_release_end() {
  // @@protoc_insertion_point(field_release:city.trip.v2.Trip.end)
  
  ::city::geo::v2::Position* temp = _impl_.end_;
  _impl_.end_ = nullptr;
  return temp;
}
inline ::city::geo::v2::Position* Trip::_internal_mutable_end() {
  
  if (_impl_.end_ == nullptr) {
    auto* p = CreateMaybeMessage<::city::geo::v2::Position>(GetArenaForAllocation());
    _impl_.end_ = p;
  }
  return _impl_.end_;
}
inline ::city::geo::v2::Position* Trip::mutable_end() {
  ::city::geo::v2::Position* _msg = _internal_mutable_end();
  // @@protoc_insertion_point(field_mutable:city.trip.v2.Trip.end)
  return _msg;
}
inline void Trip::set_allocated_end(::city::geo::v2::Position* end) {
  ::PROTOBUF_NAMESPACE_ID::Arena* message_arena = GetArenaForAllocation();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.end_);
  }
  if (end) {
    ::PROTOBUF_NAMESPACE_ID::Arena* submessage_arena =
        ::PROTOBUF_NAMESPACE_ID::Arena::InternalGetOwningArena(
                reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(end));
    if (message_arena != submessage_arena) {
      end = ::PROTOBUF_NAMESPACE_ID::internal::GetOwnedMessage(
          message_arena, end, submessage_arena);
    }
    
  } else {
    
  }
  _impl_.end_ = end;
  // @@protoc_insertion_point(field_set_allocated:city.trip.v2.Trip.end)
}

// optional double departure_time = 3 [json_name = "departureTime"];
inline bool Trip::_internal_has_departure_time() const {
  bool value = (_impl_._has_bits_[0] & 0x00000004u) != 0;
  return value;
}
inline bool Trip::has_departure_time() const {
  return _internal_has_departure_time();
}
inline void Trip::clear_departure_time() {
  _impl_.departure_time_ = 0;
  _impl_._has_bits_[0] &= ~0x00000004u;
}
inline double Trip::_internal_departure_time() const {
  return _impl_.departure_time_;
}
inline double Trip::departure_time() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Trip.departure_time)
  return _internal_departure_time();
}
inline void Trip::_internal_set_departure_time(double value) {
  _impl_._has_bits_[0] |= 0x00000004u;
  _impl_.departure_time_ = value;
}
inline void Trip::set_departure_time(double value) {
  _internal_set_departure_time(value);
  // @@protoc_insertion_point(field_set:city.trip.v2.Trip.departure_time)
}

// optional double wait_time = 4 [json_name = "waitTime"];
inline bool Trip::_internal_has_wait_time() const {
  bool value = (_impl_._has_bits_[0] & 0x00000008u) != 0;
  return value;
}
inline bool Trip::has_wait_time() const {
  return _internal_has_wait_time();
}
inline void Trip::clear_wait_time() {
  _impl_.wait_time_ = 0;
  _impl_._has_bits_[0] &= ~0x00000008u;
}
inline double Trip::_internal_wait_time() const {
  return _impl_.wait_time_;
}
inline double Trip::wait_time() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Trip.wait_time)
  return _internal_wait_time();
}
inline void Trip::_internal_set_wait_time(double value) {
  _impl_._has_bits_[0] |= 0x00000008u;
  _impl_.wait_time_ = value;
}
inline void Trip::set_wait_time(double value) {
  _internal_set_wait_time(value);
  // @@protoc_insertion_point(field_set:city.trip.v2.Trip.wait_time)
}

// optional double arrival_time = 5 [json_name = "arrivalTime"];
inline bool Trip::_internal_has_arrival_time() const {
  bool value = (_impl_._has_bits_[0] & 0x00000010u) != 0;
  return value;
}
inline bool Trip::has_arrival_time() const {
  return _internal_has_arrival_time();
}
inline void Trip::clear_arrival_time() {
  _impl_.arrival_time_ = 0;
  _impl_._has_bits_[0] &= ~0x00000010u;
}
inline double Trip::_internal_arrival_time() const {
  return _impl_.arrival_time_;
}
inline double Trip::arrival_time() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Trip.arrival_time)
  return _internal_arrival_time();
}
inline void Trip::_internal_set_arrival_time(double value) {
  _impl_._has_bits_[0] |= 0x00000010u;
  _impl_.arrival_time_ = value;
}
inline void Trip::set_arrival_time(double value) {
  _internal_set_arrival_time(value);
  // @@protoc_insertion_point(field_set:city.trip.v2.Trip.arrival_time)
}

// optional string activity = 6 [json_name = "activity"];
inline bool Trip::_internal_has_activity() const {
  bool value = (_impl_._has_bits_[0] & 0x00000001u) != 0;
  return value;
}
inline bool Trip::has_activity() const {
  return _internal_has_activity();
}
inline void Trip::clear_activity() {
  _impl_.activity_.ClearToEmpty();
  _impl_._has_bits_[0] &= ~0x00000001u;
}
inline const std::string& Trip::activity() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Trip.activity)
  return _internal_activity();
}
template <typename ArgT0, typename... ArgT>
inline PROTOBUF_ALWAYS_INLINE
void Trip::set_activity(ArgT0&& arg0, ArgT... args) {
 _impl_._has_bits_[0] |= 0x00000001u;
 _impl_.activity_.Set(static_cast<ArgT0 &&>(arg0), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:city.trip.v2.Trip.activity)
}
inline std::string* Trip::mutable_activity() {
  std::string* _s = _internal_mutable_activity();
  // @@protoc_insertion_point(field_mutable:city.trip.v2.Trip.activity)
  return _s;
}
inline const std::string& Trip::_internal_activity() const {
  return _impl_.activity_.Get();
}
inline void Trip::_internal_set_activity(const std::string& value) {
  _impl_._has_bits_[0] |= 0x00000001u;
  _impl_.activity_.Set(value, GetArenaForAllocation());
}
inline std::string* Trip::_internal_mutable_activity() {
  _impl_._has_bits_[0] |= 0x00000001u;
  return _impl_.activity_.Mutable(GetArenaForAllocation());
}
inline std::string* Trip::release_activity() {
  // @@protoc_insertion_point(field_release:city.trip.v2.Trip.activity)
  if (!_internal_has_activity()) {
    return nullptr;
  }
  _impl_._has_bits_[0] &= ~0x00000001u;
  auto* p = _impl_.activity_.Release();
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (_impl_.activity_.IsDefault()) {
    _impl_.activity_.Set("", GetArenaForAllocation());
  }
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  return p;
}
inline void Trip::set_allocated_activity(std::string* activity) {
  if (activity != nullptr) {
    _impl_._has_bits_[0] |= 0x00000001u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000001u;
  }
  _impl_.activity_.SetAllocated(activity, GetArenaForAllocation());
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (_impl_.activity_.IsDefault()) {
    _impl_.activity_.Set("", GetArenaForAllocation());
  }
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:city.trip.v2.Trip.activity)
}

// optional string model = 8 [json_name = "model"];
inline bool Trip::_internal_has_model() const {
  bool value = (_impl_._has_bits_[0] & 0x00000002u) != 0;
  return value;
}
inline bool Trip::has_model() const {
  return _internal_has_model();
}
inline void Trip::clear_model() {
  _impl_.model_.ClearToEmpty();
  _impl_._has_bits_[0] &= ~0x00000002u;
}
inline const std::string& Trip::model() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Trip.model)
  return _internal_model();
}
template <typename ArgT0, typename... ArgT>
inline PROTOBUF_ALWAYS_INLINE
void Trip::set_model(ArgT0&& arg0, ArgT... args) {
 _impl_._has_bits_[0] |= 0x00000002u;
 _impl_.model_.Set(static_cast<ArgT0 &&>(arg0), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:city.trip.v2.Trip.model)
}
inline std::string* Trip::mutable_model() {
  std::string* _s = _internal_mutable_model();
  // @@protoc_insertion_point(field_mutable:city.trip.v2.Trip.model)
  return _s;
}
inline const std::string& Trip::_internal_model() const {
  return _impl_.model_.Get();
}
inline void Trip::_internal_set_model(const std::string& value) {
  _impl_._has_bits_[0] |= 0x00000002u;
  _impl_.model_.Set(value, GetArenaForAllocation());
}
inline std::string* Trip::_internal_mutable_model() {
  _impl_._has_bits_[0] |= 0x00000002u;
  return _impl_.model_.Mutable(GetArenaForAllocation());
}
inline std::string* Trip::release_model() {
  // @@protoc_insertion_point(field_release:city.trip.v2.Trip.model)
  if (!_internal_has_model()) {
    return nullptr;
  }
  _impl_._has_bits_[0] &= ~0x00000002u;
  auto* p = _impl_.model_.Release();
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (_impl_.model_.IsDefault()) {
    _impl_.model_.Set("", GetArenaForAllocation());
  }
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  return p;
}
inline void Trip::set_allocated_model(std::string* model) {
  if (model != nullptr) {
    _impl_._has_bits_[0] |= 0x00000002u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000002u;
  }
  _impl_.model_.SetAllocated(model, GetArenaForAllocation());
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (_impl_.model_.IsDefault()) {
    _impl_.model_.Set("", GetArenaForAllocation());
  }
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:city.trip.v2.Trip.model)
}

// repeated .city.routing.v2.Journey routes = 7 [json_name = "routes"];
inline int Trip::_internal_routes_size() const {
  return _impl_.routes_.size();
}
inline int Trip::routes_size() const {
  return _internal_routes_size();
}
inline ::city::routing::v2::Journey* Trip::mutable_routes(int index) {
  // @@protoc_insertion_point(field_mutable:city.trip.v2.Trip.routes)
  return _impl_.routes_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::routing::v2::Journey >*
Trip::mutable_routes() {
  // @@protoc_insertion_point(field_mutable_list:city.trip.v2.Trip.routes)
  return &_impl_.routes_;
}
inline const ::city::routing::v2::Journey& Trip::_internal_routes(int index) const {
  return _impl_.routes_.Get(index);
}
inline const ::city::routing::v2::Journey& Trip::routes(int index) const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Trip.routes)
  return _internal_routes(index);
}
inline ::city::routing::v2::Journey* Trip::_internal_add_routes() {
  return _impl_.routes_.Add();
}
inline ::city::routing::v2::Journey* Trip::add_routes() {
  ::city::routing::v2::Journey* _add = _internal_add_routes();
  // @@protoc_insertion_point(field_add:city.trip.v2.Trip.routes)
  return _add;
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::routing::v2::Journey >&
Trip::routes() const {
  // @@protoc_insertion_point(field_list:city.trip.v2.Trip.routes)
  return _impl_.routes_;
}

// -------------------------------------------------------------------

// Schedule

// repeated .city.trip.v2.Trip trips = 1 [json_name = "trips"];
inline int Schedule::_internal_trips_size() const {
  return _impl_.trips_.size();
}
inline int Schedule::trips_size() const {
  return _internal_trips_size();
}
inline void Schedule::clear_trips() {
  _impl_.trips_.Clear();
}
inline ::city::trip::v2::Trip* Schedule::mutable_trips(int index) {
  // @@protoc_insertion_point(field_mutable:city.trip.v2.Schedule.trips)
  return _impl_.trips_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::trip::v2::Trip >*
Schedule::mutable_trips() {
  // @@protoc_insertion_point(field_mutable_list:city.trip.v2.Schedule.trips)
  return &_impl_.trips_;
}
inline const ::city::trip::v2::Trip& Schedule::_internal_trips(int index) const {
  return _impl_.trips_.Get(index);
}
inline const ::city::trip::v2::Trip& Schedule::trips(int index) const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Schedule.trips)
  return _internal_trips(index);
}
inline ::city::trip::v2::Trip* Schedule::_internal_add_trips() {
  return _impl_.trips_.Add();
}
inline ::city::trip::v2::Trip* Schedule::add_trips() {
  ::city::trip::v2::Trip* _add = _internal_add_trips();
  // @@protoc_insertion_point(field_add:city.trip.v2.Schedule.trips)
  return _add;
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::trip::v2::Trip >&
Schedule::trips() const {
  // @@protoc_insertion_point(field_list:city.trip.v2.Schedule.trips)
  return _impl_.trips_;
}

// int32 loop_count = 2 [json_name = "loopCount"];
inline void Schedule::clear_loop_count() {
  _impl_.loop_count_ = 0;
}
inline int32_t Schedule::_internal_loop_count() const {
  return _impl_.loop_count_;
}
inline int32_t Schedule::loop_count() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Schedule.loop_count)
  return _internal_loop_count();
}
inline void Schedule::_internal_set_loop_count(int32_t value) {
  
  _impl_.loop_count_ = value;
}
inline void Schedule::set_loop_count(int32_t value) {
  _internal_set_loop_count(value);
  // @@protoc_insertion_point(field_set:city.trip.v2.Schedule.loop_count)
}

// optional double departure_time = 3 [json_name = "departureTime"];
inline bool Schedule::_internal_has_departure_time() const {
  bool value = (_impl_._has_bits_[0] & 0x00000001u) != 0;
  return value;
}
inline bool Schedule::has_departure_time() const {
  return _internal_has_departure_time();
}
inline void Schedule::clear_departure_time() {
  _impl_.departure_time_ = 0;
  _impl_._has_bits_[0] &= ~0x00000001u;
}
inline double Schedule::_internal_departure_time() const {
  return _impl_.departure_time_;
}
inline double Schedule::departure_time() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Schedule.departure_time)
  return _internal_departure_time();
}
inline void Schedule::_internal_set_departure_time(double value) {
  _impl_._has_bits_[0] |= 0x00000001u;
  _impl_.departure_time_ = value;
}
inline void Schedule::set_departure_time(double value) {
  _internal_set_departure_time(value);
  // @@protoc_insertion_point(field_set:city.trip.v2.Schedule.departure_time)
}

// optional double wait_time = 4 [json_name = "waitTime"];
inline bool Schedule::_internal_has_wait_time() const {
  bool value = (_impl_._has_bits_[0] & 0x00000002u) != 0;
  return value;
}
inline bool Schedule::has_wait_time() const {
  return _internal_has_wait_time();
}
inline void Schedule::clear_wait_time() {
  _impl_.wait_time_ = 0;
  _impl_._has_bits_[0] &= ~0x00000002u;
}
inline double Schedule::_internal_wait_time() const {
  return _impl_.wait_time_;
}
inline double Schedule::wait_time() const {
  // @@protoc_insertion_point(field_get:city.trip.v2.Schedule.wait_time)
  return _internal_wait_time();
}
inline void Schedule::_internal_set_wait_time(double value) {
  _impl_._has_bits_[0] |= 0x00000002u;
  _impl_.wait_time_ = value;
}
inline void Schedule::set_wait_time(double value) {
  _internal_set_wait_time(value);
  // @@protoc_insertion_point(field_set:city.trip.v2.Schedule.wait_time)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v2
}  // namespace trip
}  // namespace city

PROTOBUF_NAMESPACE_OPEN

template <> struct is_proto_enum< ::city::trip::v2::TripMode> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::city::trip::v2::TripMode>() {
  return ::city::trip::v2::TripMode_descriptor();
}

PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2ftrip_2fv2_2ftrip_2eproto
