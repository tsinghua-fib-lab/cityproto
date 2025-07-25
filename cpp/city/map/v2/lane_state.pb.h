// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/map/v2/lane_state.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2fmap_2fv2_2flane_5fstate_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2fmap_2fv2_2flane_5fstate_2eproto

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
#include <google/protobuf/unknown_field_set.h>
#include "city/map/v2/light.pb.h"
#include "city/person/v2/motion.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_city_2fmap_2fv2_2flane_5fstate_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2fmap_2fv2_2flane_5fstate_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2fmap_2fv2_2flane_5fstate_2eproto;
namespace city {
namespace map {
namespace v2 {
class LaneState;
struct LaneStateDefaultTypeInternal;
extern LaneStateDefaultTypeInternal _LaneState_default_instance_;
}  // namespace v2
}  // namespace map
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::map::v2::LaneState* Arena::CreateMaybeMessage<::city::map::v2::LaneState>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace map {
namespace v2 {

// ===================================================================

class LaneState final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.map.v2.LaneState) */ {
 public:
  inline LaneState() : LaneState(nullptr) {}
  ~LaneState() override;
  explicit PROTOBUF_CONSTEXPR LaneState(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  LaneState(const LaneState& from);
  LaneState(LaneState&& from) noexcept
    : LaneState() {
    *this = ::std::move(from);
  }

  inline LaneState& operator=(const LaneState& from) {
    CopyFrom(from);
    return *this;
  }
  inline LaneState& operator=(LaneState&& from) noexcept {
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
  static const LaneState& default_instance() {
    return *internal_default_instance();
  }
  static inline const LaneState* internal_default_instance() {
    return reinterpret_cast<const LaneState*>(
               &_LaneState_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(LaneState& a, LaneState& b) {
    a.Swap(&b);
  }
  inline void Swap(LaneState* other) {
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
  void UnsafeArenaSwap(LaneState* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  LaneState* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<LaneState>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const LaneState& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const LaneState& from) {
    LaneState::MergeImpl(*this, from);
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
  void InternalSwap(LaneState* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.map.v2.LaneState";
  }
  protected:
  explicit LaneState(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kPersonsFieldNumber = 2,
    kIdFieldNumber = 1,
    kRestrictionFieldNumber = 4,
    kAvgVFieldNumber = 3,
    kLightStateFieldNumber = 5,
    kInVehicleCntFieldNumber = 6,
    kOutVehicleCntFieldNumber = 7,
    kVehicleCntFieldNumber = 8,
    kTotalQueuingTimeFieldNumber = 10,
    kAvgQueuingTimeFieldNumber = 11,
    kTotalQueuingVehicleCntFieldNumber = 9,
  };
  // repeated .city.person.v2.PersonMotion persons = 2 [json_name = "persons"];
  int persons_size() const;
  private:
  int _internal_persons_size() const;
  public:
  void clear_persons();
  ::city::person::v2::PersonMotion* mutable_persons(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v2::PersonMotion >*
      mutable_persons();
  private:
  const ::city::person::v2::PersonMotion& _internal_persons(int index) const;
  ::city::person::v2::PersonMotion* _internal_add_persons();
  public:
  const ::city::person::v2::PersonMotion& persons(int index) const;
  ::city::person::v2::PersonMotion* add_persons();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v2::PersonMotion >&
      persons() const;

  // int32 id = 1 [json_name = "id"];
  void clear_id();
  int32_t id() const;
  void set_id(int32_t value);
  private:
  int32_t _internal_id() const;
  void _internal_set_id(int32_t value);
  public:

  // bool restriction = 4 [json_name = "restriction"];
  void clear_restriction();
  bool restriction() const;
  void set_restriction(bool value);
  private:
  bool _internal_restriction() const;
  void _internal_set_restriction(bool value);
  public:

  // double avg_v = 3 [json_name = "avgV"];
  void clear_avg_v();
  double avg_v() const;
  void set_avg_v(double value);
  private:
  double _internal_avg_v() const;
  void _internal_set_avg_v(double value);
  public:

  // .city.map.v2.LightState light_state = 5 [json_name = "lightState"];
  void clear_light_state();
  ::city::map::v2::LightState light_state() const;
  void set_light_state(::city::map::v2::LightState value);
  private:
  ::city::map::v2::LightState _internal_light_state() const;
  void _internal_set_light_state(::city::map::v2::LightState value);
  public:

  // int32 in_vehicle_cnt = 6 [json_name = "inVehicleCnt"];
  void clear_in_vehicle_cnt();
  int32_t in_vehicle_cnt() const;
  void set_in_vehicle_cnt(int32_t value);
  private:
  int32_t _internal_in_vehicle_cnt() const;
  void _internal_set_in_vehicle_cnt(int32_t value);
  public:

  // int32 out_vehicle_cnt = 7 [json_name = "outVehicleCnt"];
  void clear_out_vehicle_cnt();
  int32_t out_vehicle_cnt() const;
  void set_out_vehicle_cnt(int32_t value);
  private:
  int32_t _internal_out_vehicle_cnt() const;
  void _internal_set_out_vehicle_cnt(int32_t value);
  public:

  // int32 vehicle_cnt = 8 [json_name = "vehicleCnt"];
  void clear_vehicle_cnt();
  int32_t vehicle_cnt() const;
  void set_vehicle_cnt(int32_t value);
  private:
  int32_t _internal_vehicle_cnt() const;
  void _internal_set_vehicle_cnt(int32_t value);
  public:

  // double total_queuing_time = 10 [json_name = "totalQueuingTime"];
  void clear_total_queuing_time();
  double total_queuing_time() const;
  void set_total_queuing_time(double value);
  private:
  double _internal_total_queuing_time() const;
  void _internal_set_total_queuing_time(double value);
  public:

  // double avg_queuing_time = 11 [json_name = "avgQueuingTime"];
  void clear_avg_queuing_time();
  double avg_queuing_time() const;
  void set_avg_queuing_time(double value);
  private:
  double _internal_avg_queuing_time() const;
  void _internal_set_avg_queuing_time(double value);
  public:

  // int32 total_queuing_vehicle_cnt = 9 [json_name = "totalQueuingVehicleCnt"];
  void clear_total_queuing_vehicle_cnt();
  int32_t total_queuing_vehicle_cnt() const;
  void set_total_queuing_vehicle_cnt(int32_t value);
  private:
  int32_t _internal_total_queuing_vehicle_cnt() const;
  void _internal_set_total_queuing_vehicle_cnt(int32_t value);
  public:

  // @@protoc_insertion_point(class_scope:city.map.v2.LaneState)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v2::PersonMotion > persons_;
    int32_t id_;
    bool restriction_;
    double avg_v_;
    int light_state_;
    int32_t in_vehicle_cnt_;
    int32_t out_vehicle_cnt_;
    int32_t vehicle_cnt_;
    double total_queuing_time_;
    double avg_queuing_time_;
    int32_t total_queuing_vehicle_cnt_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fmap_2fv2_2flane_5fstate_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// LaneState

// int32 id = 1 [json_name = "id"];
inline void LaneState::clear_id() {
  _impl_.id_ = 0;
}
inline int32_t LaneState::_internal_id() const {
  return _impl_.id_;
}
inline int32_t LaneState::id() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.id)
  return _internal_id();
}
inline void LaneState::_internal_set_id(int32_t value) {
  
  _impl_.id_ = value;
}
inline void LaneState::set_id(int32_t value) {
  _internal_set_id(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.id)
}

// repeated .city.person.v2.PersonMotion persons = 2 [json_name = "persons"];
inline int LaneState::_internal_persons_size() const {
  return _impl_.persons_.size();
}
inline int LaneState::persons_size() const {
  return _internal_persons_size();
}
inline ::city::person::v2::PersonMotion* LaneState::mutable_persons(int index) {
  // @@protoc_insertion_point(field_mutable:city.map.v2.LaneState.persons)
  return _impl_.persons_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v2::PersonMotion >*
LaneState::mutable_persons() {
  // @@protoc_insertion_point(field_mutable_list:city.map.v2.LaneState.persons)
  return &_impl_.persons_;
}
inline const ::city::person::v2::PersonMotion& LaneState::_internal_persons(int index) const {
  return _impl_.persons_.Get(index);
}
inline const ::city::person::v2::PersonMotion& LaneState::persons(int index) const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.persons)
  return _internal_persons(index);
}
inline ::city::person::v2::PersonMotion* LaneState::_internal_add_persons() {
  return _impl_.persons_.Add();
}
inline ::city::person::v2::PersonMotion* LaneState::add_persons() {
  ::city::person::v2::PersonMotion* _add = _internal_add_persons();
  // @@protoc_insertion_point(field_add:city.map.v2.LaneState.persons)
  return _add;
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v2::PersonMotion >&
LaneState::persons() const {
  // @@protoc_insertion_point(field_list:city.map.v2.LaneState.persons)
  return _impl_.persons_;
}

// double avg_v = 3 [json_name = "avgV"];
inline void LaneState::clear_avg_v() {
  _impl_.avg_v_ = 0;
}
inline double LaneState::_internal_avg_v() const {
  return _impl_.avg_v_;
}
inline double LaneState::avg_v() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.avg_v)
  return _internal_avg_v();
}
inline void LaneState::_internal_set_avg_v(double value) {
  
  _impl_.avg_v_ = value;
}
inline void LaneState::set_avg_v(double value) {
  _internal_set_avg_v(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.avg_v)
}

// bool restriction = 4 [json_name = "restriction"];
inline void LaneState::clear_restriction() {
  _impl_.restriction_ = false;
}
inline bool LaneState::_internal_restriction() const {
  return _impl_.restriction_;
}
inline bool LaneState::restriction() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.restriction)
  return _internal_restriction();
}
inline void LaneState::_internal_set_restriction(bool value) {
  
  _impl_.restriction_ = value;
}
inline void LaneState::set_restriction(bool value) {
  _internal_set_restriction(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.restriction)
}

// .city.map.v2.LightState light_state = 5 [json_name = "lightState"];
inline void LaneState::clear_light_state() {
  _impl_.light_state_ = 0;
}
inline ::city::map::v2::LightState LaneState::_internal_light_state() const {
  return static_cast< ::city::map::v2::LightState >(_impl_.light_state_);
}
inline ::city::map::v2::LightState LaneState::light_state() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.light_state)
  return _internal_light_state();
}
inline void LaneState::_internal_set_light_state(::city::map::v2::LightState value) {
  
  _impl_.light_state_ = value;
}
inline void LaneState::set_light_state(::city::map::v2::LightState value) {
  _internal_set_light_state(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.light_state)
}

// int32 in_vehicle_cnt = 6 [json_name = "inVehicleCnt"];
inline void LaneState::clear_in_vehicle_cnt() {
  _impl_.in_vehicle_cnt_ = 0;
}
inline int32_t LaneState::_internal_in_vehicle_cnt() const {
  return _impl_.in_vehicle_cnt_;
}
inline int32_t LaneState::in_vehicle_cnt() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.in_vehicle_cnt)
  return _internal_in_vehicle_cnt();
}
inline void LaneState::_internal_set_in_vehicle_cnt(int32_t value) {
  
  _impl_.in_vehicle_cnt_ = value;
}
inline void LaneState::set_in_vehicle_cnt(int32_t value) {
  _internal_set_in_vehicle_cnt(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.in_vehicle_cnt)
}

// int32 out_vehicle_cnt = 7 [json_name = "outVehicleCnt"];
inline void LaneState::clear_out_vehicle_cnt() {
  _impl_.out_vehicle_cnt_ = 0;
}
inline int32_t LaneState::_internal_out_vehicle_cnt() const {
  return _impl_.out_vehicle_cnt_;
}
inline int32_t LaneState::out_vehicle_cnt() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.out_vehicle_cnt)
  return _internal_out_vehicle_cnt();
}
inline void LaneState::_internal_set_out_vehicle_cnt(int32_t value) {
  
  _impl_.out_vehicle_cnt_ = value;
}
inline void LaneState::set_out_vehicle_cnt(int32_t value) {
  _internal_set_out_vehicle_cnt(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.out_vehicle_cnt)
}

// int32 vehicle_cnt = 8 [json_name = "vehicleCnt"];
inline void LaneState::clear_vehicle_cnt() {
  _impl_.vehicle_cnt_ = 0;
}
inline int32_t LaneState::_internal_vehicle_cnt() const {
  return _impl_.vehicle_cnt_;
}
inline int32_t LaneState::vehicle_cnt() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.vehicle_cnt)
  return _internal_vehicle_cnt();
}
inline void LaneState::_internal_set_vehicle_cnt(int32_t value) {
  
  _impl_.vehicle_cnt_ = value;
}
inline void LaneState::set_vehicle_cnt(int32_t value) {
  _internal_set_vehicle_cnt(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.vehicle_cnt)
}

// int32 total_queuing_vehicle_cnt = 9 [json_name = "totalQueuingVehicleCnt"];
inline void LaneState::clear_total_queuing_vehicle_cnt() {
  _impl_.total_queuing_vehicle_cnt_ = 0;
}
inline int32_t LaneState::_internal_total_queuing_vehicle_cnt() const {
  return _impl_.total_queuing_vehicle_cnt_;
}
inline int32_t LaneState::total_queuing_vehicle_cnt() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.total_queuing_vehicle_cnt)
  return _internal_total_queuing_vehicle_cnt();
}
inline void LaneState::_internal_set_total_queuing_vehicle_cnt(int32_t value) {
  
  _impl_.total_queuing_vehicle_cnt_ = value;
}
inline void LaneState::set_total_queuing_vehicle_cnt(int32_t value) {
  _internal_set_total_queuing_vehicle_cnt(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.total_queuing_vehicle_cnt)
}

// double total_queuing_time = 10 [json_name = "totalQueuingTime"];
inline void LaneState::clear_total_queuing_time() {
  _impl_.total_queuing_time_ = 0;
}
inline double LaneState::_internal_total_queuing_time() const {
  return _impl_.total_queuing_time_;
}
inline double LaneState::total_queuing_time() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.total_queuing_time)
  return _internal_total_queuing_time();
}
inline void LaneState::_internal_set_total_queuing_time(double value) {
  
  _impl_.total_queuing_time_ = value;
}
inline void LaneState::set_total_queuing_time(double value) {
  _internal_set_total_queuing_time(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.total_queuing_time)
}

// double avg_queuing_time = 11 [json_name = "avgQueuingTime"];
inline void LaneState::clear_avg_queuing_time() {
  _impl_.avg_queuing_time_ = 0;
}
inline double LaneState::_internal_avg_queuing_time() const {
  return _impl_.avg_queuing_time_;
}
inline double LaneState::avg_queuing_time() const {
  // @@protoc_insertion_point(field_get:city.map.v2.LaneState.avg_queuing_time)
  return _internal_avg_queuing_time();
}
inline void LaneState::_internal_set_avg_queuing_time(double value) {
  
  _impl_.avg_queuing_time_ = value;
}
inline void LaneState::set_avg_queuing_time(double value) {
  _internal_set_avg_queuing_time(value);
  // @@protoc_insertion_point(field_set:city.map.v2.LaneState.avg_queuing_time)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace v2
}  // namespace map
}  // namespace city

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2fmap_2fv2_2flane_5fstate_2eproto
