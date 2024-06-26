// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/routing/v2/cost.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2frouting_2fv2_2fcost_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2frouting_2fv2_2fcost_2eproto

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
#define PROTOBUF_INTERNAL_EXPORT_city_2frouting_2fv2_2fcost_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2frouting_2fv2_2fcost_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2frouting_2fv2_2fcost_2eproto;
namespace city {
namespace routing {
namespace v2 {
class Cost;
struct CostDefaultTypeInternal;
extern CostDefaultTypeInternal _Cost_default_instance_;
}  // namespace v2
}  // namespace routing
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::routing::v2::Cost* Arena::CreateMaybeMessage<::city::routing::v2::Cost>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace routing {
namespace v2 {

// ===================================================================

class Cost final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.routing.v2.Cost) */ {
 public:
  inline Cost() : Cost(nullptr) {}
  ~Cost() override;
  explicit PROTOBUF_CONSTEXPR Cost(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  Cost(const Cost& from);
  Cost(Cost&& from) noexcept
    : Cost() {
    *this = ::std::move(from);
  }

  inline Cost& operator=(const Cost& from) {
    CopyFrom(from);
    return *this;
  }
  inline Cost& operator=(Cost&& from) noexcept {
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
  static const Cost& default_instance() {
    return *internal_default_instance();
  }
  static inline const Cost* internal_default_instance() {
    return reinterpret_cast<const Cost*>(
               &_Cost_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(Cost& a, Cost& b) {
    a.Swap(&b);
  }
  inline void Swap(Cost* other) {
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
  void UnsafeArenaSwap(Cost* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Cost* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<Cost>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const Cost& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const Cost& from) {
    Cost::MergeImpl(*this, from);
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
  void InternalSwap(Cost* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.routing.v2.Cost";
  }
  protected:
  explicit Cost(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kCostFieldNumber = 2,
    kTimeFieldNumber = 3,
    kIdFieldNumber = 1,
  };
  // double cost = 2 [json_name = "cost"];
  void clear_cost();
  double cost() const;
  void set_cost(double value);
  private:
  double _internal_cost() const;
  void _internal_set_cost(double value);
  public:

  // optional double time = 3 [json_name = "time"];
  bool has_time() const;
  private:
  bool _internal_has_time() const;
  public:
  void clear_time();
  double time() const;
  void set_time(double value);
  private:
  double _internal_time() const;
  void _internal_set_time(double value);
  public:

  // int32 id = 1 [json_name = "id"];
  void clear_id();
  int32_t id() const;
  void set_id(int32_t value);
  private:
  int32_t _internal_id() const;
  void _internal_set_id(int32_t value);
  public:

  // @@protoc_insertion_point(class_scope:city.routing.v2.Cost)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::internal::HasBits<1> _has_bits_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
    double cost_;
    double time_;
    int32_t id_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2frouting_2fv2_2fcost_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// Cost

// int32 id = 1 [json_name = "id"];
inline void Cost::clear_id() {
  _impl_.id_ = 0;
}
inline int32_t Cost::_internal_id() const {
  return _impl_.id_;
}
inline int32_t Cost::id() const {
  // @@protoc_insertion_point(field_get:city.routing.v2.Cost.id)
  return _internal_id();
}
inline void Cost::_internal_set_id(int32_t value) {
  
  _impl_.id_ = value;
}
inline void Cost::set_id(int32_t value) {
  _internal_set_id(value);
  // @@protoc_insertion_point(field_set:city.routing.v2.Cost.id)
}

// double cost = 2 [json_name = "cost"];
inline void Cost::clear_cost() {
  _impl_.cost_ = 0;
}
inline double Cost::_internal_cost() const {
  return _impl_.cost_;
}
inline double Cost::cost() const {
  // @@protoc_insertion_point(field_get:city.routing.v2.Cost.cost)
  return _internal_cost();
}
inline void Cost::_internal_set_cost(double value) {
  
  _impl_.cost_ = value;
}
inline void Cost::set_cost(double value) {
  _internal_set_cost(value);
  // @@protoc_insertion_point(field_set:city.routing.v2.Cost.cost)
}

// optional double time = 3 [json_name = "time"];
inline bool Cost::_internal_has_time() const {
  bool value = (_impl_._has_bits_[0] & 0x00000001u) != 0;
  return value;
}
inline bool Cost::has_time() const {
  return _internal_has_time();
}
inline void Cost::clear_time() {
  _impl_.time_ = 0;
  _impl_._has_bits_[0] &= ~0x00000001u;
}
inline double Cost::_internal_time() const {
  return _impl_.time_;
}
inline double Cost::time() const {
  // @@protoc_insertion_point(field_get:city.routing.v2.Cost.time)
  return _internal_time();
}
inline void Cost::_internal_set_time(double value) {
  _impl_._has_bits_[0] |= 0x00000001u;
  _impl_.time_ = value;
}
inline void Cost::set_time(double value) {
  _internal_set_time(value);
  // @@protoc_insertion_point(field_set:city.routing.v2.Cost.time)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace v2
}  // namespace routing
}  // namespace city

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2frouting_2fv2_2fcost_2eproto
