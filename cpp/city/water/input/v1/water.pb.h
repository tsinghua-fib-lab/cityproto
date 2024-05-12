// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/water/input/v1/water.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2fwater_2finput_2fv1_2fwater_2eproto_2epb_2eh
#define GOOGLE_PROTOBUF_INCLUDED_city_2fwater_2finput_2fv1_2fwater_2eproto_2epb_2eh

#include <limits>
#include <string>
#include <type_traits>

#include "google/protobuf/port_def.inc"
#if PROTOBUF_VERSION < 4024000
#error "This file was generated by a newer version of protoc which is"
#error "incompatible with your Protocol Buffer headers. Please update"
#error "your headers."
#endif  // PROTOBUF_VERSION

#if 4024004 < PROTOBUF_MIN_PROTOC_VERSION
#error "This file was generated by an older version of protoc which is"
#error "incompatible with your Protocol Buffer headers. Please"
#error "regenerate this file with a newer version of protoc."
#endif  // PROTOBUF_MIN_PROTOC_VERSION
#include "google/protobuf/port_undef.inc"
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/arena.h"
#include "google/protobuf/arenastring.h"
#include "google/protobuf/generated_message_tctable_decl.h"
#include "google/protobuf/generated_message_util.h"
#include "google/protobuf/metadata_lite.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/message.h"
#include "google/protobuf/repeated_field.h"  // IWYU pragma: export
#include "google/protobuf/extension_set.h"  // IWYU pragma: export
#include "google/protobuf/unknown_field_set.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"

#define PROTOBUF_INTERNAL_EXPORT_city_2fwater_2finput_2fv1_2fwater_2eproto

namespace google {
namespace protobuf {
namespace internal {
class AnyMetadata;
}  // namespace internal
}  // namespace protobuf
}  // namespace google

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2fwater_2finput_2fv1_2fwater_2eproto {
  static const ::uint32_t offsets[];
};
extern const ::google::protobuf::internal::DescriptorTable
    descriptor_table_city_2fwater_2finput_2fv1_2fwater_2eproto;
namespace city {
namespace water {
namespace input {
namespace v1 {
class Rain;
struct RainDefaultTypeInternal;
extern RainDefaultTypeInternal _Rain_default_instance_;
class RainPeriod;
struct RainPeriodDefaultTypeInternal;
extern RainPeriodDefaultTypeInternal _RainPeriod_default_instance_;
}  // namespace v1
}  // namespace input
}  // namespace water
}  // namespace city
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

namespace city {
namespace water {
namespace input {
namespace v1 {

// ===================================================================


// -------------------------------------------------------------------

class RainPeriod final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:city.water.input.v1.RainPeriod) */ {
 public:
  inline RainPeriod() : RainPeriod(nullptr) {}
  ~RainPeriod() override;
  template<typename = void>
  explicit PROTOBUF_CONSTEXPR RainPeriod(::google::protobuf::internal::ConstantInitialized);

  RainPeriod(const RainPeriod& from);
  RainPeriod(RainPeriod&& from) noexcept
    : RainPeriod() {
    *this = ::std::move(from);
  }

  inline RainPeriod& operator=(const RainPeriod& from) {
    CopyFrom(from);
    return *this;
  }
  inline RainPeriod& operator=(RainPeriod&& from) noexcept {
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

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields() {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const RainPeriod& default_instance() {
    return *internal_default_instance();
  }
  static inline const RainPeriod* internal_default_instance() {
    return reinterpret_cast<const RainPeriod*>(
               &_RainPeriod_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(RainPeriod& a, RainPeriod& b) {
    a.Swap(&b);
  }
  inline void Swap(RainPeriod* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(RainPeriod* other) {
    if (other == this) return;
    ABSL_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  RainPeriod* New(::google::protobuf::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<RainPeriod>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const RainPeriod& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom( const RainPeriod& from) {
    RainPeriod::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  ::size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::google::protobuf::internal::ParseContext* ctx) final;
  ::uint8_t* _InternalSerialize(
      ::uint8_t* target, ::google::protobuf::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(RainPeriod* other);

  private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() {
    return "city.water.input.v1.RainPeriod";
  }
  protected:
  explicit RainPeriod(::google::protobuf::Arena* arena);
  public:

  static const ClassData _class_data_;
  const ::google::protobuf::Message::ClassData*GetClassData() const final;

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kRainfallFieldNumber = 2,
    kStartFieldNumber = 1,
  };
  // double rainfall = 2 [json_name = "rainfall"];
  void clear_rainfall() ;
  double rainfall() const;
  void set_rainfall(double value);

  private:
  double _internal_rainfall() const;
  void _internal_set_rainfall(double value);

  public:
  // int32 start = 1 [json_name = "start"];
  void clear_start() ;
  ::int32_t start() const;
  void set_start(::int32_t value);

  private:
  ::int32_t _internal_start() const;
  void _internal_set_start(::int32_t value);

  public:
  // @@protoc_insertion_point(class_scope:city.water.input.v1.RainPeriod)
 private:
  class _Internal;

  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<1, 2, 0, 0, 2> _table_;
  template <typename T> friend class ::google::protobuf::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    double rainfall_;
    ::int32_t start_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fwater_2finput_2fv1_2fwater_2eproto;
};// -------------------------------------------------------------------

class Rain final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:city.water.input.v1.Rain) */ {
 public:
  inline Rain() : Rain(nullptr) {}
  ~Rain() override;
  template<typename = void>
  explicit PROTOBUF_CONSTEXPR Rain(::google::protobuf::internal::ConstantInitialized);

  Rain(const Rain& from);
  Rain(Rain&& from) noexcept
    : Rain() {
    *this = ::std::move(from);
  }

  inline Rain& operator=(const Rain& from) {
    CopyFrom(from);
    return *this;
  }
  inline Rain& operator=(Rain&& from) noexcept {
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

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields() {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const Rain& default_instance() {
    return *internal_default_instance();
  }
  static inline const Rain* internal_default_instance() {
    return reinterpret_cast<const Rain*>(
               &_Rain_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(Rain& a, Rain& b) {
    a.Swap(&b);
  }
  inline void Swap(Rain* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(Rain* other) {
    if (other == this) return;
    ABSL_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Rain* New(::google::protobuf::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<Rain>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const Rain& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom( const Rain& from) {
    Rain::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  ::size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::google::protobuf::internal::ParseContext* ctx) final;
  ::uint8_t* _InternalSerialize(
      ::uint8_t* target, ::google::protobuf::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(Rain* other);

  private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() {
    return "city.water.input.v1.Rain";
  }
  protected:
  explicit Rain(::google::protobuf::Arena* arena);
  public:

  static const ClassData _class_data_;
  const ::google::protobuf::Message::ClassData*GetClassData() const final;

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kRainsFieldNumber = 1,
  };
  // repeated .city.water.input.v1.RainPeriod rains = 1 [json_name = "rains"];
  int rains_size() const;
  private:
  int _internal_rains_size() const;

  public:
  void clear_rains() ;
  ::city::water::input::v1::RainPeriod* mutable_rains(int index);
  ::google::protobuf::RepeatedPtrField< ::city::water::input::v1::RainPeriod >*
      mutable_rains();
  private:
  const ::google::protobuf::RepeatedPtrField<::city::water::input::v1::RainPeriod>& _internal_rains() const;
  ::google::protobuf::RepeatedPtrField<::city::water::input::v1::RainPeriod>* _internal_mutable_rains();
  public:
  const ::city::water::input::v1::RainPeriod& rains(int index) const;
  ::city::water::input::v1::RainPeriod* add_rains();
  const ::google::protobuf::RepeatedPtrField< ::city::water::input::v1::RainPeriod >&
      rains() const;
  // @@protoc_insertion_point(class_scope:city.water.input.v1.Rain)
 private:
  class _Internal;

  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<0, 1, 1, 0, 2> _table_;
  template <typename T> friend class ::google::protobuf::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::google::protobuf::RepeatedPtrField< ::city::water::input::v1::RainPeriod > rains_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fwater_2finput_2fv1_2fwater_2eproto;
};

// ===================================================================




// ===================================================================


#ifdef __GNUC__
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// RainPeriod

// int32 start = 1 [json_name = "start"];
inline void RainPeriod::clear_start() {
  _impl_.start_ = 0;
}
inline ::int32_t RainPeriod::start() const {
  // @@protoc_insertion_point(field_get:city.water.input.v1.RainPeriod.start)
  return _internal_start();
}
inline void RainPeriod::set_start(::int32_t value) {
  _internal_set_start(value);
  // @@protoc_insertion_point(field_set:city.water.input.v1.RainPeriod.start)
}
inline ::int32_t RainPeriod::_internal_start() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.start_;
}
inline void RainPeriod::_internal_set_start(::int32_t value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.start_ = value;
}

// double rainfall = 2 [json_name = "rainfall"];
inline void RainPeriod::clear_rainfall() {
  _impl_.rainfall_ = 0;
}
inline double RainPeriod::rainfall() const {
  // @@protoc_insertion_point(field_get:city.water.input.v1.RainPeriod.rainfall)
  return _internal_rainfall();
}
inline void RainPeriod::set_rainfall(double value) {
  _internal_set_rainfall(value);
  // @@protoc_insertion_point(field_set:city.water.input.v1.RainPeriod.rainfall)
}
inline double RainPeriod::_internal_rainfall() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.rainfall_;
}
inline void RainPeriod::_internal_set_rainfall(double value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.rainfall_ = value;
}

// -------------------------------------------------------------------

// Rain

// repeated .city.water.input.v1.RainPeriod rains = 1 [json_name = "rains"];
inline int Rain::_internal_rains_size() const {
  return _internal_rains().size();
}
inline int Rain::rains_size() const {
  return _internal_rains_size();
}
inline void Rain::clear_rains() {
  _internal_mutable_rains()->Clear();
}
inline ::city::water::input::v1::RainPeriod* Rain::mutable_rains(int index) {
  // @@protoc_insertion_point(field_mutable:city.water.input.v1.Rain.rains)
  return _internal_mutable_rains()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::city::water::input::v1::RainPeriod >*
Rain::mutable_rains() {
  // @@protoc_insertion_point(field_mutable_list:city.water.input.v1.Rain.rains)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_rains();
}
inline const ::city::water::input::v1::RainPeriod& Rain::rains(int index) const {
  // @@protoc_insertion_point(field_get:city.water.input.v1.Rain.rains)
    return _internal_rains().Get(index);
}
inline ::city::water::input::v1::RainPeriod* Rain::add_rains() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::city::water::input::v1::RainPeriod* _add = _internal_mutable_rains()->Add();
  // @@protoc_insertion_point(field_add:city.water.input.v1.Rain.rains)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::city::water::input::v1::RainPeriod >&
Rain::rains() const {
  // @@protoc_insertion_point(field_list:city.water.input.v1.Rain.rains)
  return _internal_rains();
}
inline const ::google::protobuf::RepeatedPtrField<::city::water::input::v1::RainPeriod>&
Rain::_internal_rains() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.rains_;
}
inline ::google::protobuf::RepeatedPtrField<::city::water::input::v1::RainPeriod>*
Rain::_internal_mutable_rains() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.rains_;
}

#ifdef __GNUC__
#pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace input
}  // namespace water
}  // namespace city


// @@protoc_insertion_point(global_scope)

#include "google/protobuf/port_undef.inc"

#endif  // GOOGLE_PROTOBUF_INCLUDED_city_2fwater_2finput_2fv1_2fwater_2eproto_2epb_2eh
