// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/cognition/input/v1/input.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2fcognition_2finput_2fv1_2finput_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_city_2fcognition_2finput_2fv1_2finput_2eproto

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
#define PROTOBUF_INTERNAL_EXPORT_city_2fcognition_2finput_2fv1_2finput_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2fcognition_2finput_2fv1_2finput_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_city_2fcognition_2finput_2fv1_2finput_2eproto;
namespace city {
namespace cognition {
namespace input {
namespace v1 {
class Edge;
struct EdgeDefaultTypeInternal;
extern EdgeDefaultTypeInternal _Edge_default_instance_;
class Edges;
struct EdgesDefaultTypeInternal;
extern EdgesDefaultTypeInternal _Edges_default_instance_;
}  // namespace v1
}  // namespace input
}  // namespace cognition
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> ::city::cognition::input::v1::Edge* Arena::CreateMaybeMessage<::city::cognition::input::v1::Edge>(Arena*);
template<> ::city::cognition::input::v1::Edges* Arena::CreateMaybeMessage<::city::cognition::input::v1::Edges>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace city {
namespace cognition {
namespace input {
namespace v1 {

// ===================================================================

class Edge final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.cognition.input.v1.Edge) */ {
 public:
  inline Edge() : Edge(nullptr) {}
  ~Edge() override;
  explicit PROTOBUF_CONSTEXPR Edge(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  Edge(const Edge& from);
  Edge(Edge&& from) noexcept
    : Edge() {
    *this = ::std::move(from);
  }

  inline Edge& operator=(const Edge& from) {
    CopyFrom(from);
    return *this;
  }
  inline Edge& operator=(Edge&& from) noexcept {
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
  static const Edge& default_instance() {
    return *internal_default_instance();
  }
  static inline const Edge* internal_default_instance() {
    return reinterpret_cast<const Edge*>(
               &_Edge_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(Edge& a, Edge& b) {
    a.Swap(&b);
  }
  inline void Swap(Edge* other) {
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
  void UnsafeArenaSwap(Edge* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Edge* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<Edge>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const Edge& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const Edge& from) {
    Edge::MergeImpl(*this, from);
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
  void InternalSwap(Edge* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.cognition.input.v1.Edge";
  }
  protected:
  explicit Edge(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kSourceFieldNumber = 1,
    kTargetFieldNumber = 2,
  };
  // int32 source = 1 [json_name = "source"];
  void clear_source();
  int32_t source() const;
  void set_source(int32_t value);
  private:
  int32_t _internal_source() const;
  void _internal_set_source(int32_t value);
  public:

  // int32 target = 2 [json_name = "target"];
  void clear_target();
  int32_t target() const;
  void set_target(int32_t value);
  private:
  int32_t _internal_target() const;
  void _internal_set_target(int32_t value);
  public:

  // @@protoc_insertion_point(class_scope:city.cognition.input.v1.Edge)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    int32_t source_;
    int32_t target_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fcognition_2finput_2fv1_2finput_2eproto;
};
// -------------------------------------------------------------------

class Edges final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:city.cognition.input.v1.Edges) */ {
 public:
  inline Edges() : Edges(nullptr) {}
  ~Edges() override;
  explicit PROTOBUF_CONSTEXPR Edges(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  Edges(const Edges& from);
  Edges(Edges&& from) noexcept
    : Edges() {
    *this = ::std::move(from);
  }

  inline Edges& operator=(const Edges& from) {
    CopyFrom(from);
    return *this;
  }
  inline Edges& operator=(Edges&& from) noexcept {
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
  static const Edges& default_instance() {
    return *internal_default_instance();
  }
  static inline const Edges* internal_default_instance() {
    return reinterpret_cast<const Edges*>(
               &_Edges_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(Edges& a, Edges& b) {
    a.Swap(&b);
  }
  inline void Swap(Edges* other) {
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
  void UnsafeArenaSwap(Edges* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Edges* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<Edges>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const Edges& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const Edges& from) {
    Edges::MergeImpl(*this, from);
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
  void InternalSwap(Edges* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "city.cognition.input.v1.Edges";
  }
  protected:
  explicit Edges(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kEdgesFieldNumber = 1,
  };
  // repeated .city.cognition.input.v1.Edge edges = 1 [json_name = "edges"];
  int edges_size() const;
  private:
  int _internal_edges_size() const;
  public:
  void clear_edges();
  ::city::cognition::input::v1::Edge* mutable_edges(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::cognition::input::v1::Edge >*
      mutable_edges();
  private:
  const ::city::cognition::input::v1::Edge& _internal_edges(int index) const;
  ::city::cognition::input::v1::Edge* _internal_add_edges();
  public:
  const ::city::cognition::input::v1::Edge& edges(int index) const;
  ::city::cognition::input::v1::Edge* add_edges();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::cognition::input::v1::Edge >&
      edges() const;

  // @@protoc_insertion_point(class_scope:city.cognition.input.v1.Edges)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::cognition::input::v1::Edge > edges_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fcognition_2finput_2fv1_2finput_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// Edge

// int32 source = 1 [json_name = "source"];
inline void Edge::clear_source() {
  _impl_.source_ = 0;
}
inline int32_t Edge::_internal_source() const {
  return _impl_.source_;
}
inline int32_t Edge::source() const {
  // @@protoc_insertion_point(field_get:city.cognition.input.v1.Edge.source)
  return _internal_source();
}
inline void Edge::_internal_set_source(int32_t value) {
  
  _impl_.source_ = value;
}
inline void Edge::set_source(int32_t value) {
  _internal_set_source(value);
  // @@protoc_insertion_point(field_set:city.cognition.input.v1.Edge.source)
}

// int32 target = 2 [json_name = "target"];
inline void Edge::clear_target() {
  _impl_.target_ = 0;
}
inline int32_t Edge::_internal_target() const {
  return _impl_.target_;
}
inline int32_t Edge::target() const {
  // @@protoc_insertion_point(field_get:city.cognition.input.v1.Edge.target)
  return _internal_target();
}
inline void Edge::_internal_set_target(int32_t value) {
  
  _impl_.target_ = value;
}
inline void Edge::set_target(int32_t value) {
  _internal_set_target(value);
  // @@protoc_insertion_point(field_set:city.cognition.input.v1.Edge.target)
}

// -------------------------------------------------------------------

// Edges

// repeated .city.cognition.input.v1.Edge edges = 1 [json_name = "edges"];
inline int Edges::_internal_edges_size() const {
  return _impl_.edges_.size();
}
inline int Edges::edges_size() const {
  return _internal_edges_size();
}
inline void Edges::clear_edges() {
  _impl_.edges_.Clear();
}
inline ::city::cognition::input::v1::Edge* Edges::mutable_edges(int index) {
  // @@protoc_insertion_point(field_mutable:city.cognition.input.v1.Edges.edges)
  return _impl_.edges_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::cognition::input::v1::Edge >*
Edges::mutable_edges() {
  // @@protoc_insertion_point(field_mutable_list:city.cognition.input.v1.Edges.edges)
  return &_impl_.edges_;
}
inline const ::city::cognition::input::v1::Edge& Edges::_internal_edges(int index) const {
  return _impl_.edges_.Get(index);
}
inline const ::city::cognition::input::v1::Edge& Edges::edges(int index) const {
  // @@protoc_insertion_point(field_get:city.cognition.input.v1.Edges.edges)
  return _internal_edges(index);
}
inline ::city::cognition::input::v1::Edge* Edges::_internal_add_edges() {
  return _impl_.edges_.Add();
}
inline ::city::cognition::input::v1::Edge* Edges::add_edges() {
  ::city::cognition::input::v1::Edge* _add = _internal_add_edges();
  // @@protoc_insertion_point(field_add:city.cognition.input.v1.Edges.edges)
  return _add;
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::city::cognition::input::v1::Edge >&
Edges::edges() const {
  // @@protoc_insertion_point(field_list:city.cognition.input.v1.Edges.edges)
  return _impl_.edges_;
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace input
}  // namespace cognition
}  // namespace city

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_city_2fcognition_2finput_2fv1_2finput_2eproto
