// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/person/v2/carbon.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2fperson_2fv2_2fcarbon_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2fperson_2fv2_2fcarbon_2eproto

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
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_city_2fperson_2fv2_2fcarbon_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2fperson_2fv2_2fcarbon_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2fperson_2fv2_2fcarbon_2eproto;
namespace city {
namespace person {
namespace v2 {
class EmissionStatistics;
struct EmissionStatisticsDefaultTypeInternal;
extern EmissionStatisticsDefaultTypeInternal _EmissionStatistics_default_instance_;
class VehicleCarbon;
struct VehicleCarbonDefaultTypeInternal;
extern VehicleCarbonDefaultTypeInternal _VehicleCarbon_default_instance_;
}  // namespace v2
}  // namespace person
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::person::v2::EmissionStatistics* Arena::CreateMaybeMessage<::city::person::v2::EmissionStatistics>(Arena*);
template<> ::city::person::v2::VehicleCarbon* Arena::CreateMaybeMessage<::city::person::v2::VehicleCarbon>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace person {
namespace v2 {

// ===================================================================

class VehicleCarbon final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.person.v2.VehicleCarbon) */ {
 public:
  inline VehicleCarbon() : VehicleCarbon(nullptr) {}
  ~VehicleCarbon() override;
  explicit PROTOBUF_CONSTEXPR VehicleCarbon(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  VehicleCarbon(const VehicleCarbon& from);
  VehicleCarbon(VehicleCarbon&& from) noexcept
    : VehicleCarbon() {
    *this = ::std::move(from);
  }

  inline VehicleCarbon& operator=(const VehicleCarbon& from) {
    CopyFrom(from);
    return *this;
  }
  inline VehicleCarbon& operator=(VehicleCarbon&& from) noexcept {
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
  static const VehicleCarbon& default_instance() {
    return *internal_default_instance();
  }
  static inline const VehicleCarbon* internal_default_instance() {
    return reinterpret_cast<const VehicleCarbon*>(
               &_VehicleCarbon_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(VehicleCarbon& a, VehicleCarbon& b) {
    a.Swap(&b);
  }
  inline void Swap(VehicleCarbon* other) {
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
  void UnsafeArenaSwap(VehicleCarbon* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  VehicleCarbon* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<VehicleCarbon>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const VehicleCarbon& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const VehicleCarbon& from) {
    VehicleCarbon::MergeImpl(*this, from);
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
  void InternalSwap(VehicleCarbon* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.person.v2.VehicleCarbon";
  }
  protected:
  explicit VehicleCarbon(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kDsFieldNumber = 2,
    kVFieldNumber = 3,
    kAFieldNumber = 4,
    kUAccFieldNumber = 5,
    kURollFieldNumber = 6,
    kUAeroFieldNumber = 7,
    kCDFieldNumber = 8,
    kIdFieldNumber = 1,
  };
  // double ds = 2 [json_name = "ds"];
  void clear_ds();
  double ds() const;
  void set_ds(double value);
  private:
  double _internal_ds() const;
  void _internal_set_ds(double value);
  public:

  // double v = 3 [json_name = "v"];
  void clear_v();
  double v() const;
  void set_v(double value);
  private:
  double _internal_v() const;
  void _internal_set_v(double value);
  public:

  // double a = 4 [json_name = "a"];
  void clear_a();
  double a() const;
  void set_a(double value);
  private:
  double _internal_a() const;
  void _internal_set_a(double value);
  public:

  // double u_acc = 5 [json_name = "uAcc"];
  void clear_u_acc();
  double u_acc() const;
  void set_u_acc(double value);
  private:
  double _internal_u_acc() const;
  void _internal_set_u_acc(double value);
  public:

  // double u_roll = 6 [json_name = "uRoll"];
  void clear_u_roll();
  double u_roll() const;
  void set_u_roll(double value);
  private:
  double _internal_u_roll() const;
  void _internal_set_u_roll(double value);
  public:

  // double u_aero = 7 [json_name = "uAero"];
  void clear_u_aero();
  double u_aero() const;
  void set_u_aero(double value);
  private:
  double _internal_u_aero() const;
  void _internal_set_u_aero(double value);
  public:

  // double c_d = 8 [json_name = "cD"];
  void clear_c_d();
  double c_d() const;
  void set_c_d(double value);
  private:
  double _internal_c_d() const;
  void _internal_set_c_d(double value);
  public:

  // int32 id = 1 [json_name = "id"];
  void clear_id();
  int32_t id() const;
  void set_id(int32_t value);
  private:
  int32_t _internal_id() const;
  void _internal_set_id(int32_t value);
  public:

  // @@protoc_insertion_point(class_scope:city.person.v2.VehicleCarbon)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    double ds_;
    double v_;
    double a_;
    double u_acc_;
    double u_roll_;
    double u_aero_;
    double c_d_;
    int32_t id_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fperson_2fv2_2fcarbon_2eproto;
};
// -------------------------------------------------------------------

class EmissionStatistics final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.person.v2.EmissionStatistics) */ {
 public:
  inline EmissionStatistics() : EmissionStatistics(nullptr) {}
  ~EmissionStatistics() override;
  explicit PROTOBUF_CONSTEXPR EmissionStatistics(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  EmissionStatistics(const EmissionStatistics& from);
  EmissionStatistics(EmissionStatistics&& from) noexcept
    : EmissionStatistics() {
    *this = ::std::move(from);
  }

  inline EmissionStatistics& operator=(const EmissionStatistics& from) {
    CopyFrom(from);
    return *this;
  }
  inline EmissionStatistics& operator=(EmissionStatistics&& from) noexcept {
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
  static const EmissionStatistics& default_instance() {
    return *internal_default_instance();
  }
  static inline const EmissionStatistics* internal_default_instance() {
    return reinterpret_cast<const EmissionStatistics*>(
               &_EmissionStatistics_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(EmissionStatistics& a, EmissionStatistics& b) {
    a.Swap(&b);
  }
  inline void Swap(EmissionStatistics* other) {
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
  void UnsafeArenaSwap(EmissionStatistics* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  EmissionStatistics* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<EmissionStatistics>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const EmissionStatistics& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const EmissionStatistics& from) {
    EmissionStatistics::MergeImpl(*this, from);
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
  void InternalSwap(EmissionStatistics* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.person.v2.EmissionStatistics";
  }
  protected:
  explicit EmissionStatistics(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kUFieldNumber = 1,
  };
  // float u = 1 [json_name = "u"];
  void clear_u();
  float u() const;
  void set_u(float value);
  private:
  float _internal_u() const;
  void _internal_set_u(float value);
  public:

  // @@protoc_insertion_point(class_scope:city.person.v2.EmissionStatistics)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    float u_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fperson_2fv2_2fcarbon_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// VehicleCarbon

// int32 id = 1 [json_name = "id"];
inline void VehicleCarbon::clear_id() {
  _impl_.id_ = 0;
}
inline int32_t VehicleCarbon::_internal_id() const {
  return _impl_.id_;
}
inline int32_t VehicleCarbon::id() const {
  // @@protoc_insertion_point(field_get:city.person.v2.VehicleCarbon.id)
  return _internal_id();
}
inline void VehicleCarbon::_internal_set_id(int32_t value) {
  
  _impl_.id_ = value;
}
inline void VehicleCarbon::set_id(int32_t value) {
  _internal_set_id(value);
  // @@protoc_insertion_point(field_set:city.person.v2.VehicleCarbon.id)
}

// double ds = 2 [json_name = "ds"];
inline void VehicleCarbon::clear_ds() {
  _impl_.ds_ = 0;
}
inline double VehicleCarbon::_internal_ds() const {
  return _impl_.ds_;
}
inline double VehicleCarbon::ds() const {
  // @@protoc_insertion_point(field_get:city.person.v2.VehicleCarbon.ds)
  return _internal_ds();
}
inline void VehicleCarbon::_internal_set_ds(double value) {
  
  _impl_.ds_ = value;
}
inline void VehicleCarbon::set_ds(double value) {
  _internal_set_ds(value);
  // @@protoc_insertion_point(field_set:city.person.v2.VehicleCarbon.ds)
}

// double v = 3 [json_name = "v"];
inline void VehicleCarbon::clear_v() {
  _impl_.v_ = 0;
}
inline double VehicleCarbon::_internal_v() const {
  return _impl_.v_;
}
inline double VehicleCarbon::v() const {
  // @@protoc_insertion_point(field_get:city.person.v2.VehicleCarbon.v)
  return _internal_v();
}
inline void VehicleCarbon::_internal_set_v(double value) {
  
  _impl_.v_ = value;
}
inline void VehicleCarbon::set_v(double value) {
  _internal_set_v(value);
  // @@protoc_insertion_point(field_set:city.person.v2.VehicleCarbon.v)
}

// double a = 4 [json_name = "a"];
inline void VehicleCarbon::clear_a() {
  _impl_.a_ = 0;
}
inline double VehicleCarbon::_internal_a() const {
  return _impl_.a_;
}
inline double VehicleCarbon::a() const {
  // @@protoc_insertion_point(field_get:city.person.v2.VehicleCarbon.a)
  return _internal_a();
}
inline void VehicleCarbon::_internal_set_a(double value) {
  
  _impl_.a_ = value;
}
inline void VehicleCarbon::set_a(double value) {
  _internal_set_a(value);
  // @@protoc_insertion_point(field_set:city.person.v2.VehicleCarbon.a)
}

// double u_acc = 5 [json_name = "uAcc"];
inline void VehicleCarbon::clear_u_acc() {
  _impl_.u_acc_ = 0;
}
inline double VehicleCarbon::_internal_u_acc() const {
  return _impl_.u_acc_;
}
inline double VehicleCarbon::u_acc() const {
  // @@protoc_insertion_point(field_get:city.person.v2.VehicleCarbon.u_acc)
  return _internal_u_acc();
}
inline void VehicleCarbon::_internal_set_u_acc(double value) {
  
  _impl_.u_acc_ = value;
}
inline void VehicleCarbon::set_u_acc(double value) {
  _internal_set_u_acc(value);
  // @@protoc_insertion_point(field_set:city.person.v2.VehicleCarbon.u_acc)
}

// double u_roll = 6 [json_name = "uRoll"];
inline void VehicleCarbon::clear_u_roll() {
  _impl_.u_roll_ = 0;
}
inline double VehicleCarbon::_internal_u_roll() const {
  return _impl_.u_roll_;
}
inline double VehicleCarbon::u_roll() const {
  // @@protoc_insertion_point(field_get:city.person.v2.VehicleCarbon.u_roll)
  return _internal_u_roll();
}
inline void VehicleCarbon::_internal_set_u_roll(double value) {
  
  _impl_.u_roll_ = value;
}
inline void VehicleCarbon::set_u_roll(double value) {
  _internal_set_u_roll(value);
  // @@protoc_insertion_point(field_set:city.person.v2.VehicleCarbon.u_roll)
}

// double u_aero = 7 [json_name = "uAero"];
inline void VehicleCarbon::clear_u_aero() {
  _impl_.u_aero_ = 0;
}
inline double VehicleCarbon::_internal_u_aero() const {
  return _impl_.u_aero_;
}
inline double VehicleCarbon::u_aero() const {
  // @@protoc_insertion_point(field_get:city.person.v2.VehicleCarbon.u_aero)
  return _internal_u_aero();
}
inline void VehicleCarbon::_internal_set_u_aero(double value) {
  
  _impl_.u_aero_ = value;
}
inline void VehicleCarbon::set_u_aero(double value) {
  _internal_set_u_aero(value);
  // @@protoc_insertion_point(field_set:city.person.v2.VehicleCarbon.u_aero)
}

// double c_d = 8 [json_name = "cD"];
inline void VehicleCarbon::clear_c_d() {
  _impl_.c_d_ = 0;
}
inline double VehicleCarbon::_internal_c_d() const {
  return _impl_.c_d_;
}
inline double VehicleCarbon::c_d() const {
  // @@protoc_insertion_point(field_get:city.person.v2.VehicleCarbon.c_d)
  return _internal_c_d();
}
inline void VehicleCarbon::_internal_set_c_d(double value) {
  
  _impl_.c_d_ = value;
}
inline void VehicleCarbon::set_c_d(double value) {
  _internal_set_c_d(value);
  // @@protoc_insertion_point(field_set:city.person.v2.VehicleCarbon.c_d)
}

// -------------------------------------------------------------------

// EmissionStatistics

// float u = 1 [json_name = "u"];
inline void EmissionStatistics::clear_u() {
  _impl_.u_ = 0;
}
inline float EmissionStatistics::_internal_u() const {
  return _impl_.u_;
}
inline float EmissionStatistics::u() const {
  // @@protoc_insertion_point(field_get:city.person.v2.EmissionStatistics.u)
  return _internal_u();
}
inline void EmissionStatistics::_internal_set_u(float value) {
  
  _impl_.u_ = value;
}
inline void EmissionStatistics::set_u(float value) {
  _internal_set_u(value);
  // @@protoc_insertion_point(field_set:city.person.v2.EmissionStatistics.u)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v2
}  // namespace person
}  // namespace city

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2fperson_2fv2_2fcarbon_2eproto
