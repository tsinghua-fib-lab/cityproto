// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/map/v2/aoi_service.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2fmap_2fv2_2faoi_5fservice_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2fmap_2fv2_2faoi_5fservice_2eproto

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
#include "city/person/v1/motion.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_city_2fmap_2fv2_2faoi_5fservice_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2fmap_2fv2_2faoi_5fservice_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2fmap_2fv2_2faoi_5fservice_2eproto;
namespace city {
namespace map {
namespace v2 {
class AoiState;
struct AoiStateDefaultTypeInternal;
extern AoiStateDefaultTypeInternal _AoiState_default_instance_;
class GetAoiRequest;
struct GetAoiRequestDefaultTypeInternal;
extern GetAoiRequestDefaultTypeInternal _GetAoiRequest_default_instance_;
class GetAoiResponse;
struct GetAoiResponseDefaultTypeInternal;
extern GetAoiResponseDefaultTypeInternal _GetAoiResponse_default_instance_;
}  // namespace v2
}  // namespace map
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::map::v2::AoiState* Arena::CreateMaybeMessage<::city::map::v2::AoiState>(Arena*);
template<> ::city::map::v2::GetAoiRequest* Arena::CreateMaybeMessage<::city::map::v2::GetAoiRequest>(Arena*);
template<> ::city::map::v2::GetAoiResponse* Arena::CreateMaybeMessage<::city::map::v2::GetAoiResponse>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace map {
namespace v2 {

// ===================================================================

class AoiState final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.map.v2.AoiState) */ {
 public:
  inline AoiState() : AoiState(nullptr) {}
  ~AoiState() override;
  explicit PROTOBUF_CONSTEXPR AoiState(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  AoiState(const AoiState& from);
  AoiState(AoiState&& from) noexcept
    : AoiState() {
    *this = ::std::move(from);
  }

  inline AoiState& operator=(const AoiState& from) {
    CopyFrom(from);
    return *this;
  }
  inline AoiState& operator=(AoiState&& from) noexcept {
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
  static const AoiState& default_instance() {
    return *internal_default_instance();
  }
  static inline const AoiState* internal_default_instance() {
    return reinterpret_cast<const AoiState*>(
               &_AoiState_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(AoiState& a, AoiState& b) {
    a.Swap(&b);
  }
  inline void Swap(AoiState* other) {
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
  void UnsafeArenaSwap(AoiState* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  AoiState* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<AoiState>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const AoiState& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const AoiState& from) {
    AoiState::MergeImpl(*this, from);
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
  void InternalSwap(AoiState* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.map.v2.AoiState";
  }
  protected:
  explicit AoiState(::PROTOBUF_NAMESPACE_ID::Arena* arena,
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
  };
  // repeated .city.person.v1.PersonMotion persons = 2 [json_name = "persons"];
  int persons_size() const;
  private:
  int _internal_persons_size() const;
  public:
  void clear_persons();
  ::city::person::v1::PersonMotion* mutable_persons(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v1::PersonMotion >*
      mutable_persons();
  private:
  const ::city::person::v1::PersonMotion& _internal_persons(int index) const;
  ::city::person::v1::PersonMotion* _internal_add_persons();
  public:
  const ::city::person::v1::PersonMotion& persons(int index) const;
  ::city::person::v1::PersonMotion* add_persons();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v1::PersonMotion >&
      persons() const;

  // int32 id = 1 [json_name = "id"];
  void clear_id();
  int32_t id() const;
  void set_id(int32_t value);
  private:
  int32_t _internal_id() const;
  void _internal_set_id(int32_t value);
  public:

  // @@protoc_insertion_point(class_scope:city.map.v2.AoiState)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v1::PersonMotion > persons_;
    int32_t id_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fmap_2fv2_2faoi_5fservice_2eproto;
};
// -------------------------------------------------------------------

class GetAoiRequest final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.map.v2.GetAoiRequest) */ {
 public:
  inline GetAoiRequest() : GetAoiRequest(nullptr) {}
  ~GetAoiRequest() override;
  explicit PROTOBUF_CONSTEXPR GetAoiRequest(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  GetAoiRequest(const GetAoiRequest& from);
  GetAoiRequest(GetAoiRequest&& from) noexcept
    : GetAoiRequest() {
    *this = ::std::move(from);
  }

  inline GetAoiRequest& operator=(const GetAoiRequest& from) {
    CopyFrom(from);
    return *this;
  }
  inline GetAoiRequest& operator=(GetAoiRequest&& from) noexcept {
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
  static const GetAoiRequest& default_instance() {
    return *internal_default_instance();
  }
  static inline const GetAoiRequest* internal_default_instance() {
    return reinterpret_cast<const GetAoiRequest*>(
               &_GetAoiRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(GetAoiRequest& a, GetAoiRequest& b) {
    a.Swap(&b);
  }
  inline void Swap(GetAoiRequest* other) {
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
  void UnsafeArenaSwap(GetAoiRequest* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  GetAoiRequest* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<GetAoiRequest>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const GetAoiRequest& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const GetAoiRequest& from) {
    GetAoiRequest::MergeImpl(*this, from);
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
  void InternalSwap(GetAoiRequest* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.map.v2.GetAoiRequest";
  }
  protected:
  explicit GetAoiRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kAoiIdsFieldNumber = 1,
  };
  // repeated int32 aoi_ids = 1 [json_name = "aoiIds"];
  int aoi_ids_size() const;
  private:
  int _internal_aoi_ids_size() const;
  public:
  void clear_aoi_ids();
  private:
  int32_t _internal_aoi_ids(int index) const;
  const ::PROTOBUF_NAMESPACE_ID::RepeatedField< int32_t >&
      _internal_aoi_ids() const;
  void _internal_add_aoi_ids(int32_t value);
  ::PROTOBUF_NAMESPACE_ID::RepeatedField< int32_t >*
      _internal_mutable_aoi_ids();
  public:
  int32_t aoi_ids(int index) const;
  void set_aoi_ids(int index, int32_t value);
  void add_aoi_ids(int32_t value);
  const ::PROTOBUF_NAMESPACE_ID::RepeatedField< int32_t >&
      aoi_ids() const;
  ::PROTOBUF_NAMESPACE_ID::RepeatedField< int32_t >*
      mutable_aoi_ids();

  // @@protoc_insertion_point(class_scope:city.map.v2.GetAoiRequest)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::RepeatedField< int32_t > aoi_ids_;
    mutable std::atomic<int> _aoi_ids_cached_byte_size_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fmap_2fv2_2faoi_5fservice_2eproto;
};
// -------------------------------------------------------------------

class GetAoiResponse final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.map.v2.GetAoiResponse) */ {
 public:
  inline GetAoiResponse() : GetAoiResponse(nullptr) {}
  ~GetAoiResponse() override;
  explicit PROTOBUF_CONSTEXPR GetAoiResponse(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  GetAoiResponse(const GetAoiResponse& from);
  GetAoiResponse(GetAoiResponse&& from) noexcept
    : GetAoiResponse() {
    *this = ::std::move(from);
  }

  inline GetAoiResponse& operator=(const GetAoiResponse& from) {
    CopyFrom(from);
    return *this;
  }
  inline GetAoiResponse& operator=(GetAoiResponse&& from) noexcept {
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
  static const GetAoiResponse& default_instance() {
    return *internal_default_instance();
  }
  static inline const GetAoiResponse* internal_default_instance() {
    return reinterpret_cast<const GetAoiResponse*>(
               &_GetAoiResponse_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    2;

  friend void swap(GetAoiResponse& a, GetAoiResponse& b) {
    a.Swap(&b);
  }
  inline void Swap(GetAoiResponse* other) {
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
  void UnsafeArenaSwap(GetAoiResponse* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  GetAoiResponse* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<GetAoiResponse>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const GetAoiResponse& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const GetAoiResponse& from) {
    GetAoiResponse::MergeImpl(*this, from);
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
  void InternalSwap(GetAoiResponse* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.map.v2.GetAoiResponse";
  }
  protected:
  explicit GetAoiResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kStatesFieldNumber = 1,
  };
  // repeated .city.map.v2.AoiState states = 1 [json_name = "states"];
  int states_size() const;
  private:
  int _internal_states_size() const;
  public:
  void clear_states();
  ::city::map::v2::AoiState* mutable_states(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::map::v2::AoiState >*
      mutable_states();
  private:
  const ::city::map::v2::AoiState& _internal_states(int index) const;
  ::city::map::v2::AoiState* _internal_add_states();
  public:
  const ::city::map::v2::AoiState& states(int index) const;
  ::city::map::v2::AoiState* add_states();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::map::v2::AoiState >&
      states() const;

  // @@protoc_insertion_point(class_scope:city.map.v2.GetAoiResponse)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::map::v2::AoiState > states_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fmap_2fv2_2faoi_5fservice_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// AoiState

// int32 id = 1 [json_name = "id"];
inline void AoiState::clear_id() {
  _impl_.id_ = 0;
}
inline int32_t AoiState::_internal_id() const {
  return _impl_.id_;
}
inline int32_t AoiState::id() const {
  // @@protoc_insertion_point(field_get:city.map.v2.AoiState.id)
  return _internal_id();
}
inline void AoiState::_internal_set_id(int32_t value) {
  
  _impl_.id_ = value;
}
inline void AoiState::set_id(int32_t value) {
  _internal_set_id(value);
  // @@protoc_insertion_point(field_set:city.map.v2.AoiState.id)
}

// repeated .city.person.v1.PersonMotion persons = 2 [json_name = "persons"];
inline int AoiState::_internal_persons_size() const {
  return _impl_.persons_.size();
}
inline int AoiState::persons_size() const {
  return _internal_persons_size();
}
inline ::city::person::v1::PersonMotion* AoiState::mutable_persons(int index) {
  // @@protoc_insertion_point(field_mutable:city.map.v2.AoiState.persons)
  return _impl_.persons_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v1::PersonMotion >*
AoiState::mutable_persons() {
  // @@protoc_insertion_point(field_mutable_list:city.map.v2.AoiState.persons)
  return &_impl_.persons_;
}
inline const ::city::person::v1::PersonMotion& AoiState::_internal_persons(int index) const {
  return _impl_.persons_.Get(index);
}
inline const ::city::person::v1::PersonMotion& AoiState::persons(int index) const {
  // @@protoc_insertion_point(field_get:city.map.v2.AoiState.persons)
  return _internal_persons(index);
}
inline ::city::person::v1::PersonMotion* AoiState::_internal_add_persons() {
  return _impl_.persons_.Add();
}
inline ::city::person::v1::PersonMotion* AoiState::add_persons() {
  ::city::person::v1::PersonMotion* _add = _internal_add_persons();
  // @@protoc_insertion_point(field_add:city.map.v2.AoiState.persons)
  return _add;
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::person::v1::PersonMotion >&
AoiState::persons() const {
  // @@protoc_insertion_point(field_list:city.map.v2.AoiState.persons)
  return _impl_.persons_;
}

// -------------------------------------------------------------------

// GetAoiRequest

// repeated int32 aoi_ids = 1 [json_name = "aoiIds"];
inline int GetAoiRequest::_internal_aoi_ids_size() const {
  return _impl_.aoi_ids_.size();
}
inline int GetAoiRequest::aoi_ids_size() const {
  return _internal_aoi_ids_size();
}
inline void GetAoiRequest::clear_aoi_ids() {
  _impl_.aoi_ids_.Clear();
}
inline int32_t GetAoiRequest::_internal_aoi_ids(int index) const {
  return _impl_.aoi_ids_.Get(index);
}
inline int32_t GetAoiRequest::aoi_ids(int index) const {
  // @@protoc_insertion_point(field_get:city.map.v2.GetAoiRequest.aoi_ids)
  return _internal_aoi_ids(index);
}
inline void GetAoiRequest::set_aoi_ids(int index, int32_t value) {
  _impl_.aoi_ids_.Set(index, value);
  // @@protoc_insertion_point(field_set:city.map.v2.GetAoiRequest.aoi_ids)
}
inline void GetAoiRequest::_internal_add_aoi_ids(int32_t value) {
  _impl_.aoi_ids_.Add(value);
}
inline void GetAoiRequest::add_aoi_ids(int32_t value) {
  _internal_add_aoi_ids(value);
  // @@protoc_insertion_point(field_add:city.map.v2.GetAoiRequest.aoi_ids)
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedField< int32_t >&
GetAoiRequest::_internal_aoi_ids() const {
  return _impl_.aoi_ids_;
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedField< int32_t >&
GetAoiRequest::aoi_ids() const {
  // @@protoc_insertion_point(field_list:city.map.v2.GetAoiRequest.aoi_ids)
  return _internal_aoi_ids();
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedField< int32_t >*
GetAoiRequest::_internal_mutable_aoi_ids() {
  return &_impl_.aoi_ids_;
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedField< int32_t >*
GetAoiRequest::mutable_aoi_ids() {
  // @@protoc_insertion_point(field_mutable_list:city.map.v2.GetAoiRequest.aoi_ids)
  return _internal_mutable_aoi_ids();
}

// -------------------------------------------------------------------

// GetAoiResponse

// repeated .city.map.v2.AoiState states = 1 [json_name = "states"];
inline int GetAoiResponse::_internal_states_size() const {
  return _impl_.states_.size();
}
inline int GetAoiResponse::states_size() const {
  return _internal_states_size();
}
inline void GetAoiResponse::clear_states() {
  _impl_.states_.Clear();
}
inline ::city::map::v2::AoiState* GetAoiResponse::mutable_states(int index) {
  // @@protoc_insertion_point(field_mutable:city.map.v2.GetAoiResponse.states)
  return _impl_.states_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::map::v2::AoiState >*
GetAoiResponse::mutable_states() {
  // @@protoc_insertion_point(field_mutable_list:city.map.v2.GetAoiResponse.states)
  return &_impl_.states_;
}
inline const ::city::map::v2::AoiState& GetAoiResponse::_internal_states(int index) const {
  return _impl_.states_.Get(index);
}
inline const ::city::map::v2::AoiState& GetAoiResponse::states(int index) const {
  // @@protoc_insertion_point(field_get:city.map.v2.GetAoiResponse.states)
  return _internal_states(index);
}
inline ::city::map::v2::AoiState* GetAoiResponse::_internal_add_states() {
  return _impl_.states_.Add();
}
inline ::city::map::v2::AoiState* GetAoiResponse::add_states() {
  ::city::map::v2::AoiState* _add = _internal_add_states();
  // @@protoc_insertion_point(field_add:city.map.v2.GetAoiResponse.states)
  return _add;
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::map::v2::AoiState >&
GetAoiResponse::states() const {
  // @@protoc_insertion_point(field_list:city.map.v2.GetAoiResponse.states)
  return _impl_.states_;
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------

// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v2
}  // namespace map
}  // namespace city

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2fmap_2fv2_2faoi_5fservice_2eproto