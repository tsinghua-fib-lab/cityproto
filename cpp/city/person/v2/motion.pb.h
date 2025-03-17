// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/person/v2/motion.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2fperson_2fv2_2fmotion_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2fperson_2fv2_2fmotion_2eproto

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
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_city_2fperson_2fv2_2fmotion_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2fperson_2fv2_2fmotion_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2fperson_2fv2_2fmotion_2eproto;
namespace city {
namespace person {
namespace v2 {
class PersonMotion;
struct PersonMotionDefaultTypeInternal;
extern PersonMotionDefaultTypeInternal _PersonMotion_default_instance_;
}  // namespace v2
}  // namespace person
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::person::v2::PersonMotion* Arena::CreateMaybeMessage<::city::person::v2::PersonMotion>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace person {
namespace v2 {

enum Status : int {
  STATUS_UNSPECIFIED = 0,
  STATUS_SLEEP = 1,
  STATUS_DRIVING = 2,
  STATUS_WALKING = 3,
  STATUS_CROWD = 4,
  STATUS_PASSENGER = 5,
  STATUS_WAIT_ROUTE = 6,
  STATUS_WAIT_BUS = 7,
  STATUS_RAIL_TRANSIT = 8,
  STATUS_WAIT_TAXI = 9,
  Status_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::min(),
  Status_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::max()
};
bool Status_IsValid(int value);
constexpr Status Status_MIN = STATUS_UNSPECIFIED;
constexpr Status Status_MAX = STATUS_WAIT_TAXI;
constexpr int Status_ARRAYSIZE = Status_MAX + 1;

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* Status_descriptor();
template<typename T>
inline const std::string& Status_Name(T enum_t_value) {
  static_assert(::std::is_same<T, Status>::value ||
    ::std::is_integral<T>::value,
    "Incorrect type passed to function Status_Name.");
  return ::PROTOBUF_NAMESPACE_ID::internal::NameOfEnum(
    Status_descriptor(), enum_t_value);
}
inline bool Status_Parse(
    ::PROTOBUF_NAMESPACE_ID::ConstStringParam name, Status* value) {
  return ::PROTOBUF_NAMESPACE_ID::internal::ParseNamedEnum<Status>(
    Status_descriptor(), name, value);
}
// ===================================================================

class PersonMotion final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.person.v2.PersonMotion) */ {
 public:
  inline PersonMotion() : PersonMotion(nullptr) {}
  ~PersonMotion() override;
  explicit PROTOBUF_CONSTEXPR PersonMotion(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  PersonMotion(const PersonMotion& from);
  PersonMotion(PersonMotion&& from) noexcept
    : PersonMotion() {
    *this = ::std::move(from);
  }

  inline PersonMotion& operator=(const PersonMotion& from) {
    CopyFrom(from);
    return *this;
  }
  inline PersonMotion& operator=(PersonMotion&& from) noexcept {
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
  static const PersonMotion& default_instance() {
    return *internal_default_instance();
  }
  static inline const PersonMotion* internal_default_instance() {
    return reinterpret_cast<const PersonMotion*>(
               &_PersonMotion_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(PersonMotion& a, PersonMotion& b) {
    a.Swap(&b);
  }
  inline void Swap(PersonMotion* other) {
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
  void UnsafeArenaSwap(PersonMotion* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  PersonMotion* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<PersonMotion>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const PersonMotion& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const PersonMotion& from) {
    PersonMotion::MergeImpl(*this, from);
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
  void InternalSwap(PersonMotion* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.person.v2.PersonMotion";
  }
  protected:
  explicit PersonMotion(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kActivityFieldNumber = 6,
    kPositionFieldNumber = 3,
    kIdFieldNumber = 1,
    kStatusFieldNumber = 2,
    kVFieldNumber = 4,
    kDirectionFieldNumber = 5,
    kLFieldNumber = 7,
    kAFieldNumber = 8,
  };
  // string activity = 6 [json_name = "activity"];
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

  // .city.geo.v2.Position position = 3 [json_name = "position"];
  bool has_position() const;
  private:
  bool _internal_has_position() const;
  public:
  void clear_position();
  const ::city::geo::v2::Position& position() const;
  PROTOBUF_NODISCARD ::city::geo::v2::Position* release_position();
  ::city::geo::v2::Position* mutable_position();
  void set_allocated_position(::city::geo::v2::Position* position);
  private:
  const ::city::geo::v2::Position& _internal_position() const;
  ::city::geo::v2::Position* _internal_mutable_position();
  public:
  void unsafe_arena_set_allocated_position(
      ::city::geo::v2::Position* position);
  ::city::geo::v2::Position* unsafe_arena_release_position();

  // int32 id = 1 [json_name = "id"];
  void clear_id();
  int32_t id() const;
  void set_id(int32_t value);
  private:
  int32_t _internal_id() const;
  void _internal_set_id(int32_t value);
  public:

  // .city.person.v2.Status status = 2 [json_name = "status"];
  void clear_status();
  ::city::person::v2::Status status() const;
  void set_status(::city::person::v2::Status value);
  private:
  ::city::person::v2::Status _internal_status() const;
  void _internal_set_status(::city::person::v2::Status value);
  public:

  // double v = 4 [json_name = "v"];
  void clear_v();
  double v() const;
  void set_v(double value);
  private:
  double _internal_v() const;
  void _internal_set_v(double value);
  public:

  // double direction = 5 [json_name = "direction"];
  void clear_direction();
  double direction() const;
  void set_direction(double value);
  private:
  double _internal_direction() const;
  void _internal_set_direction(double value);
  public:

  // double l = 7 [json_name = "l"];
  void clear_l();
  double l() const;
  void set_l(double value);
  private:
  double _internal_l() const;
  void _internal_set_l(double value);
  public:

  // double a = 8 [json_name = "a"];
  void clear_a();
  double a() const;
  void set_a(double value);
  private:
  double _internal_a() const;
  void _internal_set_a(double value);
  public:

  // @@protoc_insertion_point(class_scope:city.person.v2.PersonMotion)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr activity_;
    ::city::geo::v2::Position* position_;
    int32_t id_;
    int status_;
    double v_;
    double direction_;
    double l_;
    double a_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fperson_2fv2_2fmotion_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// PersonMotion

// int32 id = 1 [json_name = "id"];
inline void PersonMotion::clear_id() {
  _impl_.id_ = 0;
}
inline int32_t PersonMotion::_internal_id() const {
  return _impl_.id_;
}
inline int32_t PersonMotion::id() const {
  // @@protoc_insertion_point(field_get:city.person.v2.PersonMotion.id)
  return _internal_id();
}
inline void PersonMotion::_internal_set_id(int32_t value) {
  
  _impl_.id_ = value;
}
inline void PersonMotion::set_id(int32_t value) {
  _internal_set_id(value);
  // @@protoc_insertion_point(field_set:city.person.v2.PersonMotion.id)
}

// .city.person.v2.Status status = 2 [json_name = "status"];
inline void PersonMotion::clear_status() {
  _impl_.status_ = 0;
}
inline ::city::person::v2::Status PersonMotion::_internal_status() const {
  return static_cast< ::city::person::v2::Status >(_impl_.status_);
}
inline ::city::person::v2::Status PersonMotion::status() const {
  // @@protoc_insertion_point(field_get:city.person.v2.PersonMotion.status)
  return _internal_status();
}
inline void PersonMotion::_internal_set_status(::city::person::v2::Status value) {
  
  _impl_.status_ = value;
}
inline void PersonMotion::set_status(::city::person::v2::Status value) {
  _internal_set_status(value);
  // @@protoc_insertion_point(field_set:city.person.v2.PersonMotion.status)
}

// .city.geo.v2.Position position = 3 [json_name = "position"];
inline bool PersonMotion::_internal_has_position() const {
  return this != internal_default_instance() && _impl_.position_ != nullptr;
}
inline bool PersonMotion::has_position() const {
  return _internal_has_position();
}
inline const ::city::geo::v2::Position& PersonMotion::_internal_position() const {
  const ::city::geo::v2::Position* p = _impl_.position_;
  return p != nullptr ? *p : reinterpret_cast<const ::city::geo::v2::Position&>(
      ::city::geo::v2::_Position_default_instance_);
}
inline const ::city::geo::v2::Position& PersonMotion::position() const {
  // @@protoc_insertion_point(field_get:city.person.v2.PersonMotion.position)
  return _internal_position();
}
inline void PersonMotion::unsafe_arena_set_allocated_position(
    ::city::geo::v2::Position* position) {
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.position_);
  }
  _impl_.position_ = position;
  if (position) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:city.person.v2.PersonMotion.position)
}
inline ::city::geo::v2::Position* PersonMotion::release_position() {
  
  ::city::geo::v2::Position* temp = _impl_.position_;
  _impl_.position_ = nullptr;
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
inline ::city::geo::v2::Position* PersonMotion::unsafe_arena_release_position() {
  // @@protoc_insertion_point(field_release:city.person.v2.PersonMotion.position)
  
  ::city::geo::v2::Position* temp = _impl_.position_;
  _impl_.position_ = nullptr;
  return temp;
}
inline ::city::geo::v2::Position* PersonMotion::_internal_mutable_position() {
  
  if (_impl_.position_ == nullptr) {
    auto* p = CreateMaybeMessage<::city::geo::v2::Position>(GetArenaForAllocation());
    _impl_.position_ = p;
  }
  return _impl_.position_;
}
inline ::city::geo::v2::Position* PersonMotion::mutable_position() {
  ::city::geo::v2::Position* _msg = _internal_mutable_position();
  // @@protoc_insertion_point(field_mutable:city.person.v2.PersonMotion.position)
  return _msg;
}
inline void PersonMotion::set_allocated_position(::city::geo::v2::Position* position) {
  ::PROTOBUF_NAMESPACE_ID::Arena* message_arena = GetArenaForAllocation();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.position_);
  }
  if (position) {
    ::PROTOBUF_NAMESPACE_ID::Arena* submessage_arena =
        ::PROTOBUF_NAMESPACE_ID::Arena::InternalGetOwningArena(
                reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(position));
    if (message_arena != submessage_arena) {
      position = ::PROTOBUF_NAMESPACE_ID::internal::GetOwnedMessage(
          message_arena, position, submessage_arena);
    }
    
  } else {
    
  }
  _impl_.position_ = position;
  // @@protoc_insertion_point(field_set_allocated:city.person.v2.PersonMotion.position)
}

// double v = 4 [json_name = "v"];
inline void PersonMotion::clear_v() {
  _impl_.v_ = 0;
}
inline double PersonMotion::_internal_v() const {
  return _impl_.v_;
}
inline double PersonMotion::v() const {
  // @@protoc_insertion_point(field_get:city.person.v2.PersonMotion.v)
  return _internal_v();
}
inline void PersonMotion::_internal_set_v(double value) {
  
  _impl_.v_ = value;
}
inline void PersonMotion::set_v(double value) {
  _internal_set_v(value);
  // @@protoc_insertion_point(field_set:city.person.v2.PersonMotion.v)
}

// double direction = 5 [json_name = "direction"];
inline void PersonMotion::clear_direction() {
  _impl_.direction_ = 0;
}
inline double PersonMotion::_internal_direction() const {
  return _impl_.direction_;
}
inline double PersonMotion::direction() const {
  // @@protoc_insertion_point(field_get:city.person.v2.PersonMotion.direction)
  return _internal_direction();
}
inline void PersonMotion::_internal_set_direction(double value) {
  
  _impl_.direction_ = value;
}
inline void PersonMotion::set_direction(double value) {
  _internal_set_direction(value);
  // @@protoc_insertion_point(field_set:city.person.v2.PersonMotion.direction)
}

// string activity = 6 [json_name = "activity"];
inline void PersonMotion::clear_activity() {
  _impl_.activity_.ClearToEmpty();
}
inline const std::string& PersonMotion::activity() const {
  // @@protoc_insertion_point(field_get:city.person.v2.PersonMotion.activity)
  return _internal_activity();
}
template <typename ArgT0, typename... ArgT>
inline PROTOBUF_ALWAYS_INLINE
void PersonMotion::set_activity(ArgT0&& arg0, ArgT... args) {
 
 _impl_.activity_.Set(static_cast<ArgT0 &&>(arg0), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:city.person.v2.PersonMotion.activity)
}
inline std::string* PersonMotion::mutable_activity() {
  std::string* _s = _internal_mutable_activity();
  // @@protoc_insertion_point(field_mutable:city.person.v2.PersonMotion.activity)
  return _s;
}
inline const std::string& PersonMotion::_internal_activity() const {
  return _impl_.activity_.Get();
}
inline void PersonMotion::_internal_set_activity(const std::string& value) {
  
  _impl_.activity_.Set(value, GetArenaForAllocation());
}
inline std::string* PersonMotion::_internal_mutable_activity() {
  
  return _impl_.activity_.Mutable(GetArenaForAllocation());
}
inline std::string* PersonMotion::release_activity() {
  // @@protoc_insertion_point(field_release:city.person.v2.PersonMotion.activity)
  return _impl_.activity_.Release();
}
inline void PersonMotion::set_allocated_activity(std::string* activity) {
  if (activity != nullptr) {
    
  } else {
    
  }
  _impl_.activity_.SetAllocated(activity, GetArenaForAllocation());
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (_impl_.activity_.IsDefault()) {
    _impl_.activity_.Set("", GetArenaForAllocation());
  }
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:city.person.v2.PersonMotion.activity)
}

// double l = 7 [json_name = "l"];
inline void PersonMotion::clear_l() {
  _impl_.l_ = 0;
}
inline double PersonMotion::_internal_l() const {
  return _impl_.l_;
}
inline double PersonMotion::l() const {
  // @@protoc_insertion_point(field_get:city.person.v2.PersonMotion.l)
  return _internal_l();
}
inline void PersonMotion::_internal_set_l(double value) {
  
  _impl_.l_ = value;
}
inline void PersonMotion::set_l(double value) {
  _internal_set_l(value);
  // @@protoc_insertion_point(field_set:city.person.v2.PersonMotion.l)
}

// double a = 8 [json_name = "a"];
inline void PersonMotion::clear_a() {
  _impl_.a_ = 0;
}
inline double PersonMotion::_internal_a() const {
  return _impl_.a_;
}
inline double PersonMotion::a() const {
  // @@protoc_insertion_point(field_get:city.person.v2.PersonMotion.a)
  return _internal_a();
}
inline void PersonMotion::_internal_set_a(double value) {
  
  _impl_.a_ = value;
}
inline void PersonMotion::set_a(double value) {
  _internal_set_a(value);
  // @@protoc_insertion_point(field_set:city.person.v2.PersonMotion.a)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace v2
}  // namespace person
}  // namespace city

PROTOBUF_NAMESPACE_OPEN

template <> struct is_proto_enum< ::city::person::v2::Status> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::city::person::v2::Status>() {
  return ::city::person::v2::Status_descriptor();
}

PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2fperson_2fv2_2fmotion_2eproto
