// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/elec/input/v1/input_service.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2felec_2finput_2fv1_2finput_5fservice_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2felec_2finput_2fv1_2finput_5fservice_2eproto

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
#include "city/elec/input/v1/config.pb.h"
#include "city/elec/input/v1/input.pb.h"
#include "city/map/v2/map.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_city_2felec_2finput_2fv1_2finput_5fservice_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2felec_2finput_2fv1_2finput_5fservice_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2felec_2finput_2fv1_2finput_5fservice_2eproto;
namespace city {
namespace elec {
namespace input {
namespace v1 {
class InitRequest;
struct InitRequestDefaultTypeInternal;
extern InitRequestDefaultTypeInternal _InitRequest_default_instance_;
class InitResponse;
struct InitResponseDefaultTypeInternal;
extern InitResponseDefaultTypeInternal _InitResponse_default_instance_;
}  // namespace v1
}  // namespace input
}  // namespace elec
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::elec::input::v1::InitRequest* Arena::CreateMaybeMessage<::city::elec::input::v1::InitRequest>(Arena*);
template<> ::city::elec::input::v1::InitResponse* Arena::CreateMaybeMessage<::city::elec::input::v1::InitResponse>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace elec {
namespace input {
namespace v1 {

// ===================================================================

class InitRequest final :
    public ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase /* @@protoc_insertion_point(class_definition:city.elec.input.v1.InitRequest) */ {
 public:
  inline InitRequest() : InitRequest(nullptr) {}
  explicit PROTOBUF_CONSTEXPR InitRequest(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  InitRequest(const InitRequest& from);
  InitRequest(InitRequest&& from) noexcept
    : InitRequest() {
    *this = ::std::move(from);
  }

  inline InitRequest& operator=(const InitRequest& from) {
    CopyFrom(from);
    return *this;
  }
  inline InitRequest& operator=(InitRequest&& from) noexcept {
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
  static const InitRequest& default_instance() {
    return *internal_default_instance();
  }
  static inline const InitRequest* internal_default_instance() {
    return reinterpret_cast<const InitRequest*>(
               &_InitRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(InitRequest& a, InitRequest& b) {
    a.Swap(&b);
  }
  inline void Swap(InitRequest* other) {
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
  void UnsafeArenaSwap(InitRequest* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  InitRequest* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<InitRequest>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyFrom;
  inline void CopyFrom(const InitRequest& from) {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyImpl(*this, from);
  }
  using ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeFrom;
  void MergeFrom(const InitRequest& from) {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeImpl(*this, from);
  }
  public:

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.elec.input.v1.InitRequest";
  }
  protected:
  explicit InitRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // @@protoc_insertion_point(class_scope:city.elec.input.v1.InitRequest)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
  };
  friend struct ::TableStruct_city_2felec_2finput_2fv1_2finput_5fservice_2eproto;
};
// -------------------------------------------------------------------

class InitResponse final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.elec.input.v1.InitResponse) */ {
 public:
  inline InitResponse() : InitResponse(nullptr) {}
  ~InitResponse() override;
  explicit PROTOBUF_CONSTEXPR InitResponse(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  InitResponse(const InitResponse& from);
  InitResponse(InitResponse&& from) noexcept
    : InitResponse() {
    *this = ::std::move(from);
  }

  inline InitResponse& operator=(const InitResponse& from) {
    CopyFrom(from);
    return *this;
  }
  inline InitResponse& operator=(InitResponse&& from) noexcept {
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
  static const InitResponse& default_instance() {
    return *internal_default_instance();
  }
  static inline const InitResponse* internal_default_instance() {
    return reinterpret_cast<const InitResponse*>(
               &_InitResponse_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(InitResponse& a, InitResponse& b) {
    a.Swap(&b);
  }
  inline void Swap(InitResponse* other) {
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
  void UnsafeArenaSwap(InitResponse* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  InitResponse* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<InitResponse>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const InitResponse& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const InitResponse& from) {
    InitResponse::MergeImpl(*this, from);
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
  void InternalSwap(InitResponse* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.elec.input.v1.InitResponse";
  }
  protected:
  explicit InitResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kAddressFieldNumber = 2,
    kFacilitiesFieldNumber = 1,
    kControlFieldNumber = 3,
    kMapFieldNumber = 4,
  };
  // string address = 2 [json_name = "address"];
  void clear_address();
  const std::string& address() const;
  template <typename ArgT0 = const std::string&, typename... ArgT>
  void set_address(ArgT0&& arg0, ArgT... args);
  std::string* mutable_address();
  PROTOBUF_NODISCARD std::string* release_address();
  void set_allocated_address(std::string* address);
  private:
  const std::string& _internal_address() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_address(const std::string& value);
  std::string* _internal_mutable_address();
  public:

  // .city.elec.input.v1.Facilities facilities = 1 [json_name = "facilities"];
  bool has_facilities() const;
  private:
  bool _internal_has_facilities() const;
  public:
  void clear_facilities();
  const ::city::elec::input::v1::Facilities& facilities() const;
  PROTOBUF_NODISCARD ::city::elec::input::v1::Facilities* release_facilities();
  ::city::elec::input::v1::Facilities* mutable_facilities();
  void set_allocated_facilities(::city::elec::input::v1::Facilities* facilities);
  private:
  const ::city::elec::input::v1::Facilities& _internal_facilities() const;
  ::city::elec::input::v1::Facilities* _internal_mutable_facilities();
  public:
  void unsafe_arena_set_allocated_facilities(
      ::city::elec::input::v1::Facilities* facilities);
  ::city::elec::input::v1::Facilities* unsafe_arena_release_facilities();

  // .city.elec.input.v1.Control control = 3 [json_name = "control"];
  bool has_control() const;
  private:
  bool _internal_has_control() const;
  public:
  void clear_control();
  const ::city::elec::input::v1::Control& control() const;
  PROTOBUF_NODISCARD ::city::elec::input::v1::Control* release_control();
  ::city::elec::input::v1::Control* mutable_control();
  void set_allocated_control(::city::elec::input::v1::Control* control);
  private:
  const ::city::elec::input::v1::Control& _internal_control() const;
  ::city::elec::input::v1::Control* _internal_mutable_control();
  public:
  void unsafe_arena_set_allocated_control(
      ::city::elec::input::v1::Control* control);
  ::city::elec::input::v1::Control* unsafe_arena_release_control();

  // .city.map.v2.Map map = 4 [json_name = "map"];
  bool has_map() const;
  private:
  bool _internal_has_map() const;
  public:
  void clear_map();
  const ::city::map::v2::Map& map() const;
  PROTOBUF_NODISCARD ::city::map::v2::Map* release_map();
  ::city::map::v2::Map* mutable_map();
  void set_allocated_map(::city::map::v2::Map* map);
  private:
  const ::city::map::v2::Map& _internal_map() const;
  ::city::map::v2::Map* _internal_mutable_map();
  public:
  void unsafe_arena_set_allocated_map(
      ::city::map::v2::Map* map);
  ::city::map::v2::Map* unsafe_arena_release_map();

  // @@protoc_insertion_point(class_scope:city.elec.input.v1.InitResponse)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr address_;
    ::city::elec::input::v1::Facilities* facilities_;
    ::city::elec::input::v1::Control* control_;
    ::city::map::v2::Map* map_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2felec_2finput_2fv1_2finput_5fservice_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// InitRequest

// -------------------------------------------------------------------

// InitResponse

// string address = 2 [json_name = "address"];
inline void InitResponse::clear_address() {
  _impl_.address_.ClearToEmpty();
}
inline const std::string& InitResponse::address() const {
  // @@protoc_insertion_point(field_get:city.elec.input.v1.InitResponse.address)
  return _internal_address();
}
template <typename ArgT0, typename... ArgT>
inline PROTOBUF_ALWAYS_INLINE
void InitResponse::set_address(ArgT0&& arg0, ArgT... args) {
 
 _impl_.address_.Set(static_cast<ArgT0 &&>(arg0), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:city.elec.input.v1.InitResponse.address)
}
inline std::string* InitResponse::mutable_address() {
  std::string* _s = _internal_mutable_address();
  // @@protoc_insertion_point(field_mutable:city.elec.input.v1.InitResponse.address)
  return _s;
}
inline const std::string& InitResponse::_internal_address() const {
  return _impl_.address_.Get();
}
inline void InitResponse::_internal_set_address(const std::string& value) {
  
  _impl_.address_.Set(value, GetArenaForAllocation());
}
inline std::string* InitResponse::_internal_mutable_address() {
  
  return _impl_.address_.Mutable(GetArenaForAllocation());
}
inline std::string* InitResponse::release_address() {
  // @@protoc_insertion_point(field_release:city.elec.input.v1.InitResponse.address)
  return _impl_.address_.Release();
}
inline void InitResponse::set_allocated_address(std::string* address) {
  if (address != nullptr) {
    
  } else {
    
  }
  _impl_.address_.SetAllocated(address, GetArenaForAllocation());
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (_impl_.address_.IsDefault()) {
    _impl_.address_.Set("", GetArenaForAllocation());
  }
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:city.elec.input.v1.InitResponse.address)
}

// .city.elec.input.v1.Control control = 3 [json_name = "control"];
inline bool InitResponse::_internal_has_control() const {
  return this != internal_default_instance() && _impl_.control_ != nullptr;
}
inline bool InitResponse::has_control() const {
  return _internal_has_control();
}
inline const ::city::elec::input::v1::Control& InitResponse::_internal_control() const {
  const ::city::elec::input::v1::Control* p = _impl_.control_;
  return p != nullptr ? *p : reinterpret_cast<const ::city::elec::input::v1::Control&>(
      ::city::elec::input::v1::_Control_default_instance_);
}
inline const ::city::elec::input::v1::Control& InitResponse::control() const {
  // @@protoc_insertion_point(field_get:city.elec.input.v1.InitResponse.control)
  return _internal_control();
}
inline void InitResponse::unsafe_arena_set_allocated_control(
    ::city::elec::input::v1::Control* control) {
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.control_);
  }
  _impl_.control_ = control;
  if (control) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:city.elec.input.v1.InitResponse.control)
}
inline ::city::elec::input::v1::Control* InitResponse::release_control() {
  
  ::city::elec::input::v1::Control* temp = _impl_.control_;
  _impl_.control_ = nullptr;
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
inline ::city::elec::input::v1::Control* InitResponse::unsafe_arena_release_control() {
  // @@protoc_insertion_point(field_release:city.elec.input.v1.InitResponse.control)
  
  ::city::elec::input::v1::Control* temp = _impl_.control_;
  _impl_.control_ = nullptr;
  return temp;
}
inline ::city::elec::input::v1::Control* InitResponse::_internal_mutable_control() {
  
  if (_impl_.control_ == nullptr) {
    auto* p = CreateMaybeMessage<::city::elec::input::v1::Control>(GetArenaForAllocation());
    _impl_.control_ = p;
  }
  return _impl_.control_;
}
inline ::city::elec::input::v1::Control* InitResponse::mutable_control() {
  ::city::elec::input::v1::Control* _msg = _internal_mutable_control();
  // @@protoc_insertion_point(field_mutable:city.elec.input.v1.InitResponse.control)
  return _msg;
}
inline void InitResponse::set_allocated_control(::city::elec::input::v1::Control* control) {
  ::PROTOBUF_NAMESPACE_ID::Arena* message_arena = GetArenaForAllocation();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.control_);
  }
  if (control) {
    ::PROTOBUF_NAMESPACE_ID::Arena* submessage_arena =
        ::PROTOBUF_NAMESPACE_ID::Arena::InternalGetOwningArena(
                reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(control));
    if (message_arena != submessage_arena) {
      control = ::PROTOBUF_NAMESPACE_ID::internal::GetOwnedMessage(
          message_arena, control, submessage_arena);
    }
    
  } else {
    
  }
  _impl_.control_ = control;
  // @@protoc_insertion_point(field_set_allocated:city.elec.input.v1.InitResponse.control)
}

// .city.elec.input.v1.Facilities facilities = 1 [json_name = "facilities"];
inline bool InitResponse::_internal_has_facilities() const {
  return this != internal_default_instance() && _impl_.facilities_ != nullptr;
}
inline bool InitResponse::has_facilities() const {
  return _internal_has_facilities();
}
inline const ::city::elec::input::v1::Facilities& InitResponse::_internal_facilities() const {
  const ::city::elec::input::v1::Facilities* p = _impl_.facilities_;
  return p != nullptr ? *p : reinterpret_cast<const ::city::elec::input::v1::Facilities&>(
      ::city::elec::input::v1::_Facilities_default_instance_);
}
inline const ::city::elec::input::v1::Facilities& InitResponse::facilities() const {
  // @@protoc_insertion_point(field_get:city.elec.input.v1.InitResponse.facilities)
  return _internal_facilities();
}
inline void InitResponse::unsafe_arena_set_allocated_facilities(
    ::city::elec::input::v1::Facilities* facilities) {
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.facilities_);
  }
  _impl_.facilities_ = facilities;
  if (facilities) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:city.elec.input.v1.InitResponse.facilities)
}
inline ::city::elec::input::v1::Facilities* InitResponse::release_facilities() {
  
  ::city::elec::input::v1::Facilities* temp = _impl_.facilities_;
  _impl_.facilities_ = nullptr;
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
inline ::city::elec::input::v1::Facilities* InitResponse::unsafe_arena_release_facilities() {
  // @@protoc_insertion_point(field_release:city.elec.input.v1.InitResponse.facilities)
  
  ::city::elec::input::v1::Facilities* temp = _impl_.facilities_;
  _impl_.facilities_ = nullptr;
  return temp;
}
inline ::city::elec::input::v1::Facilities* InitResponse::_internal_mutable_facilities() {
  
  if (_impl_.facilities_ == nullptr) {
    auto* p = CreateMaybeMessage<::city::elec::input::v1::Facilities>(GetArenaForAllocation());
    _impl_.facilities_ = p;
  }
  return _impl_.facilities_;
}
inline ::city::elec::input::v1::Facilities* InitResponse::mutable_facilities() {
  ::city::elec::input::v1::Facilities* _msg = _internal_mutable_facilities();
  // @@protoc_insertion_point(field_mutable:city.elec.input.v1.InitResponse.facilities)
  return _msg;
}
inline void InitResponse::set_allocated_facilities(::city::elec::input::v1::Facilities* facilities) {
  ::PROTOBUF_NAMESPACE_ID::Arena* message_arena = GetArenaForAllocation();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.facilities_);
  }
  if (facilities) {
    ::PROTOBUF_NAMESPACE_ID::Arena* submessage_arena =
        ::PROTOBUF_NAMESPACE_ID::Arena::InternalGetOwningArena(
                reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(facilities));
    if (message_arena != submessage_arena) {
      facilities = ::PROTOBUF_NAMESPACE_ID::internal::GetOwnedMessage(
          message_arena, facilities, submessage_arena);
    }
    
  } else {
    
  }
  _impl_.facilities_ = facilities;
  // @@protoc_insertion_point(field_set_allocated:city.elec.input.v1.InitResponse.facilities)
}

// .city.map.v2.Map map = 4 [json_name = "map"];
inline bool InitResponse::_internal_has_map() const {
  return this != internal_default_instance() && _impl_.map_ != nullptr;
}
inline bool InitResponse::has_map() const {
  return _internal_has_map();
}
inline const ::city::map::v2::Map& InitResponse::_internal_map() const {
  const ::city::map::v2::Map* p = _impl_.map_;
  return p != nullptr ? *p : reinterpret_cast<const ::city::map::v2::Map&>(
      ::city::map::v2::_Map_default_instance_);
}
inline const ::city::map::v2::Map& InitResponse::map() const {
  // @@protoc_insertion_point(field_get:city.elec.input.v1.InitResponse.map)
  return _internal_map();
}
inline void InitResponse::unsafe_arena_set_allocated_map(
    ::city::map::v2::Map* map) {
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.map_);
  }
  _impl_.map_ = map;
  if (map) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:city.elec.input.v1.InitResponse.map)
}
inline ::city::map::v2::Map* InitResponse::release_map() {
  
  ::city::map::v2::Map* temp = _impl_.map_;
  _impl_.map_ = nullptr;
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
inline ::city::map::v2::Map* InitResponse::unsafe_arena_release_map() {
  // @@protoc_insertion_point(field_release:city.elec.input.v1.InitResponse.map)
  
  ::city::map::v2::Map* temp = _impl_.map_;
  _impl_.map_ = nullptr;
  return temp;
}
inline ::city::map::v2::Map* InitResponse::_internal_mutable_map() {
  
  if (_impl_.map_ == nullptr) {
    auto* p = CreateMaybeMessage<::city::map::v2::Map>(GetArenaForAllocation());
    _impl_.map_ = p;
  }
  return _impl_.map_;
}
inline ::city::map::v2::Map* InitResponse::mutable_map() {
  ::city::map::v2::Map* _msg = _internal_mutable_map();
  // @@protoc_insertion_point(field_mutable:city.elec.input.v1.InitResponse.map)
  return _msg;
}
inline void InitResponse::set_allocated_map(::city::map::v2::Map* map) {
  ::PROTOBUF_NAMESPACE_ID::Arena* message_arena = GetArenaForAllocation();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.map_);
  }
  if (map) {
    ::PROTOBUF_NAMESPACE_ID::Arena* submessage_arena =
        ::PROTOBUF_NAMESPACE_ID::Arena::InternalGetOwningArena(
                reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(map));
    if (message_arena != submessage_arena) {
      map = ::PROTOBUF_NAMESPACE_ID::internal::GetOwnedMessage(
          message_arena, map, submessage_arena);
    }
    
  } else {
    
  }
  _impl_.map_ = map;
  // @@protoc_insertion_point(field_set_allocated:city.elec.input.v1.InitResponse.map)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace input
}  // namespace elec
}  // namespace city

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2felec_2finput_2fv1_2finput_5fservice_2eproto
