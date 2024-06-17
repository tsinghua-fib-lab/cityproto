// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/comm/interaction/demand/v1/demand_service.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2fcomm_2finteraction_2fdemand_2fv1_2fdemand_5fservice_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2fcomm_2finteraction_2fdemand_2fv1_2fdemand_5fservice_2eproto

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
#define PROTOBUF_INTERNAL_EXPORT_city_2fcomm_2finteraction_2fdemand_2fv1_2fdemand_5fservice_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2fcomm_2finteraction_2fdemand_2fv1_2fdemand_5fservice_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2fcomm_2finteraction_2fdemand_2fv1_2fdemand_5fservice_2eproto;
namespace city {
namespace comm {
namespace interaction {
namespace demand {
namespace v1 {
class SetDemandStatusRequest;
struct SetDemandStatusRequestDefaultTypeInternal;
extern SetDemandStatusRequestDefaultTypeInternal _SetDemandStatusRequest_default_instance_;
class SetDemandStatusResponse;
struct SetDemandStatusResponseDefaultTypeInternal;
extern SetDemandStatusResponseDefaultTypeInternal _SetDemandStatusResponse_default_instance_;
}  // namespace v1
}  // namespace demand
}  // namespace interaction
}  // namespace comm
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::comm::interaction::demand::v1::SetDemandStatusRequest* Arena::CreateMaybeMessage<::city::comm::interaction::demand::v1::SetDemandStatusRequest>(Arena*);
template<> ::city::comm::interaction::demand::v1::SetDemandStatusResponse* Arena::CreateMaybeMessage<::city::comm::interaction::demand::v1::SetDemandStatusResponse>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace comm {
namespace interaction {
namespace demand {
namespace v1 {

// ===================================================================

class SetDemandStatusRequest final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.comm.interaction.demand.v1.SetDemandStatusRequest) */ {
 public:
  inline SetDemandStatusRequest() : SetDemandStatusRequest(nullptr) {}
  ~SetDemandStatusRequest() override;
  explicit PROTOBUF_CONSTEXPR SetDemandStatusRequest(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  SetDemandStatusRequest(const SetDemandStatusRequest& from);
  SetDemandStatusRequest(SetDemandStatusRequest&& from) noexcept
    : SetDemandStatusRequest() {
    *this = ::std::move(from);
  }

  inline SetDemandStatusRequest& operator=(const SetDemandStatusRequest& from) {
    CopyFrom(from);
    return *this;
  }
  inline SetDemandStatusRequest& operator=(SetDemandStatusRequest&& from) noexcept {
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
  static const SetDemandStatusRequest& default_instance() {
    return *internal_default_instance();
  }
  static inline const SetDemandStatusRequest* internal_default_instance() {
    return reinterpret_cast<const SetDemandStatusRequest*>(
               &_SetDemandStatusRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(SetDemandStatusRequest& a, SetDemandStatusRequest& b) {
    a.Swap(&b);
  }
  inline void Swap(SetDemandStatusRequest* other) {
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
  void UnsafeArenaSwap(SetDemandStatusRequest* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  SetDemandStatusRequest* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<SetDemandStatusRequest>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const SetDemandStatusRequest& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const SetDemandStatusRequest& from) {
    SetDemandStatusRequest::MergeImpl(*this, from);
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
  void InternalSwap(SetDemandStatusRequest* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.comm.interaction.demand.v1.SetDemandStatusRequest";
  }
  protected:
  explicit SetDemandStatusRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kMultipleTimesFieldNumber = 1,
    kPowerTimesFieldNumber = 2,
  };
  // double multiple_times = 1 [json_name = "multipleTimes"];
  void clear_multiple_times();
  double multiple_times() const;
  void set_multiple_times(double value);
  private:
  double _internal_multiple_times() const;
  void _internal_set_multiple_times(double value);
  public:

  // double power_times = 2 [json_name = "powerTimes"];
  void clear_power_times();
  double power_times() const;
  void set_power_times(double value);
  private:
  double _internal_power_times() const;
  void _internal_set_power_times(double value);
  public:

  // @@protoc_insertion_point(class_scope:city.comm.interaction.demand.v1.SetDemandStatusRequest)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    double multiple_times_;
    double power_times_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fcomm_2finteraction_2fdemand_2fv1_2fdemand_5fservice_2eproto;
};
// -------------------------------------------------------------------

class SetDemandStatusResponse final :
    public ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase /* @@protoc_insertion_point(class_definition:city.comm.interaction.demand.v1.SetDemandStatusResponse) */ {
 public:
  inline SetDemandStatusResponse() : SetDemandStatusResponse(nullptr) {}
  explicit PROTOBUF_CONSTEXPR SetDemandStatusResponse(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  SetDemandStatusResponse(const SetDemandStatusResponse& from);
  SetDemandStatusResponse(SetDemandStatusResponse&& from) noexcept
    : SetDemandStatusResponse() {
    *this = ::std::move(from);
  }

  inline SetDemandStatusResponse& operator=(const SetDemandStatusResponse& from) {
    CopyFrom(from);
    return *this;
  }
  inline SetDemandStatusResponse& operator=(SetDemandStatusResponse&& from) noexcept {
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
  static const SetDemandStatusResponse& default_instance() {
    return *internal_default_instance();
  }
  static inline const SetDemandStatusResponse* internal_default_instance() {
    return reinterpret_cast<const SetDemandStatusResponse*>(
               &_SetDemandStatusResponse_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(SetDemandStatusResponse& a, SetDemandStatusResponse& b) {
    a.Swap(&b);
  }
  inline void Swap(SetDemandStatusResponse* other) {
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
  void UnsafeArenaSwap(SetDemandStatusResponse* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  SetDemandStatusResponse* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<SetDemandStatusResponse>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyFrom;
  inline void CopyFrom(const SetDemandStatusResponse& from) {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyImpl(*this, from);
  }
  using ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeFrom;
  void MergeFrom(const SetDemandStatusResponse& from) {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeImpl(*this, from);
  }
  public:

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.comm.interaction.demand.v1.SetDemandStatusResponse";
  }
  protected:
  explicit SetDemandStatusResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // @@protoc_insertion_point(class_scope:city.comm.interaction.demand.v1.SetDemandStatusResponse)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
  };
  friend struct ::TableStruct_city_2fcomm_2finteraction_2fdemand_2fv1_2fdemand_5fservice_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// SetDemandStatusRequest

// double multiple_times = 1 [json_name = "multipleTimes"];
inline void SetDemandStatusRequest::clear_multiple_times() {
  _impl_.multiple_times_ = 0;
}
inline double SetDemandStatusRequest::_internal_multiple_times() const {
  return _impl_.multiple_times_;
}
inline double SetDemandStatusRequest::multiple_times() const {
  // @@protoc_insertion_point(field_get:city.comm.interaction.demand.v1.SetDemandStatusRequest.multiple_times)
  return _internal_multiple_times();
}
inline void SetDemandStatusRequest::_internal_set_multiple_times(double value) {
  
  _impl_.multiple_times_ = value;
}
inline void SetDemandStatusRequest::set_multiple_times(double value) {
  _internal_set_multiple_times(value);
  // @@protoc_insertion_point(field_set:city.comm.interaction.demand.v1.SetDemandStatusRequest.multiple_times)
}

// double power_times = 2 [json_name = "powerTimes"];
inline void SetDemandStatusRequest::clear_power_times() {
  _impl_.power_times_ = 0;
}
inline double SetDemandStatusRequest::_internal_power_times() const {
  return _impl_.power_times_;
}
inline double SetDemandStatusRequest::power_times() const {
  // @@protoc_insertion_point(field_get:city.comm.interaction.demand.v1.SetDemandStatusRequest.power_times)
  return _internal_power_times();
}
inline void SetDemandStatusRequest::_internal_set_power_times(double value) {
  
  _impl_.power_times_ = value;
}
inline void SetDemandStatusRequest::set_power_times(double value) {
  _internal_set_power_times(value);
  // @@protoc_insertion_point(field_set:city.comm.interaction.demand.v1.SetDemandStatusRequest.power_times)
}

// -------------------------------------------------------------------

// SetDemandStatusResponse

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace demand
}  // namespace interaction
}  // namespace comm
}  // namespace city

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2fcomm_2finteraction_2fdemand_2fv1_2fdemand_5fservice_2eproto