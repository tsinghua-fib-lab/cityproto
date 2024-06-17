// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/clock/v1/clock_service.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2fclock_2fv1_2fclock_5fservice_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2fclock_2fv1_2fclock_5fservice_2eproto

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
#include <google/protobuf/generated_message_bases.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_city_2fclock_2fv1_2fclock_5fservice_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2fclock_2fv1_2fclock_5fservice_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2fclock_2fv1_2fclock_5fservice_2eproto;
namespace city {
namespace clock {
namespace v1 {
class NowRequest;
struct NowRequestDefaultTypeInternal;
extern NowRequestDefaultTypeInternal _NowRequest_default_instance_;
class NowResponse;
struct NowResponseDefaultTypeInternal;
extern NowResponseDefaultTypeInternal _NowResponse_default_instance_;
}  // namespace v1
}  // namespace clock
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::clock::v1::NowRequest* Arena::CreateMaybeMessage<::city::clock::v1::NowRequest>(Arena*);
template<> ::city::clock::v1::NowResponse* Arena::CreateMaybeMessage<::city::clock::v1::NowResponse>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace clock {
namespace v1 {

// ===================================================================

class NowRequest final :
    public ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase /* @@protoc_insertion_point(class_definition:city.clock.v1.NowRequest) */ {
 public:
  inline NowRequest() : NowRequest(nullptr) {}
  explicit PROTOBUF_CONSTEXPR NowRequest(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  NowRequest(const NowRequest& from);
  NowRequest(NowRequest&& from) noexcept
    : NowRequest() {
    *this = ::std::move(from);
  }

  inline NowRequest& operator=(const NowRequest& from) {
    CopyFrom(from);
    return *this;
  }
  inline NowRequest& operator=(NowRequest&& from) noexcept {
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
  static const NowRequest& default_instance() {
    return *internal_default_instance();
  }
  static inline const NowRequest* internal_default_instance() {
    return reinterpret_cast<const NowRequest*>(
               &_NowRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(NowRequest& a, NowRequest& b) {
    a.Swap(&b);
  }
  inline void Swap(NowRequest* other) {
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
  void UnsafeArenaSwap(NowRequest* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  NowRequest* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<NowRequest>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyFrom;
  inline void CopyFrom(const NowRequest& from) {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyImpl(*this, from);
  }
  using ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeFrom;
  void MergeFrom(const NowRequest& from) {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeImpl(*this, from);
  }
  public:

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.clock.v1.NowRequest";
  }
  protected:
  explicit NowRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // @@protoc_insertion_point(class_scope:city.clock.v1.NowRequest)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
  };
  friend struct ::TableStruct_city_2fclock_2fv1_2fclock_5fservice_2eproto;
};
// -------------------------------------------------------------------

class NowResponse final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.clock.v1.NowResponse) */ {
 public:
  inline NowResponse() : NowResponse(nullptr) {}
  ~NowResponse() override;
  explicit PROTOBUF_CONSTEXPR NowResponse(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  NowResponse(const NowResponse& from);
  NowResponse(NowResponse&& from) noexcept
    : NowResponse() {
    *this = ::std::move(from);
  }

  inline NowResponse& operator=(const NowResponse& from) {
    CopyFrom(from);
    return *this;
  }
  inline NowResponse& operator=(NowResponse&& from) noexcept {
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
  static const NowResponse& default_instance() {
    return *internal_default_instance();
  }
  static inline const NowResponse* internal_default_instance() {
    return reinterpret_cast<const NowResponse*>(
               &_NowResponse_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(NowResponse& a, NowResponse& b) {
    a.Swap(&b);
  }
  inline void Swap(NowResponse* other) {
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
  void UnsafeArenaSwap(NowResponse* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  NowResponse* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<NowResponse>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const NowResponse& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const NowResponse& from) {
    NowResponse::MergeImpl(*this, from);
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
  void InternalSwap(NowResponse* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.clock.v1.NowResponse";
  }
  protected:
  explicit NowResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kTFieldNumber = 1,
  };
  // double t = 1 [json_name = "t"];
  void clear_t();
  double t() const;
  void set_t(double value);
  private:
  double _internal_t() const;
  void _internal_set_t(double value);
  public:

  // @@protoc_insertion_point(class_scope:city.clock.v1.NowResponse)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    double t_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fclock_2fv1_2fclock_5fservice_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// NowRequest

// -------------------------------------------------------------------

// NowResponse

// double t = 1 [json_name = "t"];
inline void NowResponse::clear_t() {
  _impl_.t_ = 0;
}
inline double NowResponse::_internal_t() const {
  return _impl_.t_;
}
inline double NowResponse::t() const {
  // @@protoc_insertion_point(field_get:city.clock.v1.NowResponse.t)
  return _internal_t();
}
inline void NowResponse::_internal_set_t(double value) {
  
  _impl_.t_ = value;
}
inline void NowResponse::set_t(double value) {
  _internal_set_t(value);
  // @@protoc_insertion_point(field_set:city.clock.v1.NowResponse.t)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace clock
}  // namespace city

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2fclock_2fv1_2fclock_5fservice_2eproto