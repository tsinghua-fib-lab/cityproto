// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/cognition/output/v1/output_service.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_city_2fcognition_2foutput_2fv1_2foutput_5fservice_2eproto_2epb_2eh
#define GOOGLE_PROTOBUF_INCLUDED_city_2fcognition_2foutput_2fv1_2foutput_5fservice_2eproto_2epb_2eh

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
#include "google/protobuf/generated_message_bases.h"
#include "google/protobuf/generated_message_tctable_decl.h"
#include "google/protobuf/generated_message_util.h"
#include "google/protobuf/metadata_lite.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/message.h"
#include "google/protobuf/repeated_field.h"  // IWYU pragma: export
#include "google/protobuf/extension_set.h"  // IWYU pragma: export
#include "google/protobuf/unknown_field_set.h"
#include "city/cognition/output/v1/output.pb.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"

#define PROTOBUF_INTERNAL_EXPORT_city_2fcognition_2foutput_2fv1_2foutput_5fservice_2eproto

namespace google {
namespace protobuf {
namespace internal {
class AnyMetadata;
}  // namespace internal
}  // namespace protobuf
}  // namespace google

// Internal implementation detail -- do not use these members.
struct TableStruct_city_2fcognition_2foutput_2fv1_2foutput_5fservice_2eproto {
  static const ::uint32_t offsets[];
};
extern const ::google::protobuf::internal::DescriptorTable
    descriptor_table_city_2fcognition_2foutput_2fv1_2foutput_5fservice_2eproto;
namespace city {
namespace cognition {
namespace output {
namespace v1 {
class OutputRequest;
struct OutputRequestDefaultTypeInternal;
extern OutputRequestDefaultTypeInternal _OutputRequest_default_instance_;
class OutputResponse;
struct OutputResponseDefaultTypeInternal;
extern OutputResponseDefaultTypeInternal _OutputResponse_default_instance_;
}  // namespace v1
}  // namespace output
}  // namespace cognition
}  // namespace city
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

namespace city {
namespace cognition {
namespace output {
namespace v1 {

// ===================================================================


// -------------------------------------------------------------------

class OutputRequest final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:city.cognition.output.v1.OutputRequest) */ {
 public:
  inline OutputRequest() : OutputRequest(nullptr) {}
  ~OutputRequest() override;
  template<typename = void>
  explicit PROTOBUF_CONSTEXPR OutputRequest(::google::protobuf::internal::ConstantInitialized);

  OutputRequest(const OutputRequest& from);
  OutputRequest(OutputRequest&& from) noexcept
    : OutputRequest() {
    *this = ::std::move(from);
  }

  inline OutputRequest& operator=(const OutputRequest& from) {
    CopyFrom(from);
    return *this;
  }
  inline OutputRequest& operator=(OutputRequest&& from) noexcept {
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
  static const OutputRequest& default_instance() {
    return *internal_default_instance();
  }
  static inline const OutputRequest* internal_default_instance() {
    return reinterpret_cast<const OutputRequest*>(
               &_OutputRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(OutputRequest& a, OutputRequest& b) {
    a.Swap(&b);
  }
  inline void Swap(OutputRequest* other) {
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
  void UnsafeArenaSwap(OutputRequest* other) {
    if (other == this) return;
    ABSL_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  OutputRequest* New(::google::protobuf::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<OutputRequest>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const OutputRequest& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom( const OutputRequest& from) {
    OutputRequest::MergeImpl(*this, from);
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
  void InternalSwap(OutputRequest* other);

  private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() {
    return "city.cognition.output.v1.OutputRequest";
  }
  protected:
  explicit OutputRequest(::google::protobuf::Arena* arena);
  public:

  static const ClassData _class_data_;
  const ::google::protobuf::Message::ClassData*GetClassData() const final;

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kNodesFieldNumber = 2,
    kInfluencesFieldNumber = 3,
    kContentsFieldNumber = 6,
    kGroupInfluencesFieldNumber = 7,
    kGroupsFieldNumber = 8,
    kHeatmapFieldNumber = 4,
    kStatFieldNumber = 5,
    kStepFieldNumber = 1,
  };
  // repeated .city.cognition.output.v1.Node nodes = 2 [json_name = "nodes"];
  int nodes_size() const;
  private:
  int _internal_nodes_size() const;

  public:
  void clear_nodes() ;
  ::city::cognition::output::v1::Node* mutable_nodes(int index);
  ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Node >*
      mutable_nodes();
  private:
  const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Node>& _internal_nodes() const;
  ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Node>* _internal_mutable_nodes();
  public:
  const ::city::cognition::output::v1::Node& nodes(int index) const;
  ::city::cognition::output::v1::Node* add_nodes();
  const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Node >&
      nodes() const;
  // repeated .city.cognition.output.v1.Influence influences = 3 [json_name = "influences"];
  int influences_size() const;
  private:
  int _internal_influences_size() const;

  public:
  void clear_influences() ;
  ::city::cognition::output::v1::Influence* mutable_influences(int index);
  ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence >*
      mutable_influences();
  private:
  const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Influence>& _internal_influences() const;
  ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Influence>* _internal_mutable_influences();
  public:
  const ::city::cognition::output::v1::Influence& influences(int index) const;
  ::city::cognition::output::v1::Influence* add_influences();
  const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence >&
      influences() const;
  // repeated .city.cognition.output.v1.Content contents = 6 [json_name = "contents"];
  int contents_size() const;
  private:
  int _internal_contents_size() const;

  public:
  void clear_contents() ;
  ::city::cognition::output::v1::Content* mutable_contents(int index);
  ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Content >*
      mutable_contents();
  private:
  const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Content>& _internal_contents() const;
  ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Content>* _internal_mutable_contents();
  public:
  const ::city::cognition::output::v1::Content& contents(int index) const;
  ::city::cognition::output::v1::Content* add_contents();
  const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Content >&
      contents() const;
  // repeated .city.cognition.output.v1.Influence group_influences = 7 [json_name = "groupInfluences"];
  int group_influences_size() const;
  private:
  int _internal_group_influences_size() const;

  public:
  void clear_group_influences() ;
  ::city::cognition::output::v1::Influence* mutable_group_influences(int index);
  ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence >*
      mutable_group_influences();
  private:
  const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Influence>& _internal_group_influences() const;
  ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Influence>* _internal_mutable_group_influences();
  public:
  const ::city::cognition::output::v1::Influence& group_influences(int index) const;
  ::city::cognition::output::v1::Influence* add_group_influences();
  const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence >&
      group_influences() const;
  // repeated .city.cognition.output.v1.Group groups = 8 [json_name = "groups"];
  int groups_size() const;
  private:
  int _internal_groups_size() const;

  public:
  void clear_groups() ;
  ::city::cognition::output::v1::Group* mutable_groups(int index);
  ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Group >*
      mutable_groups();
  private:
  const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Group>& _internal_groups() const;
  ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Group>* _internal_mutable_groups();
  public:
  const ::city::cognition::output::v1::Group& groups(int index) const;
  ::city::cognition::output::v1::Group* add_groups();
  const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Group >&
      groups() const;
  // .city.cognition.output.v1.Heatmap heatmap = 4 [json_name = "heatmap"];
  bool has_heatmap() const;
  void clear_heatmap() ;
  const ::city::cognition::output::v1::Heatmap& heatmap() const;
  PROTOBUF_NODISCARD ::city::cognition::output::v1::Heatmap* release_heatmap();
  ::city::cognition::output::v1::Heatmap* mutable_heatmap();
  void set_allocated_heatmap(::city::cognition::output::v1::Heatmap* value);
  void unsafe_arena_set_allocated_heatmap(::city::cognition::output::v1::Heatmap* value);
  ::city::cognition::output::v1::Heatmap* unsafe_arena_release_heatmap();

  private:
  const ::city::cognition::output::v1::Heatmap& _internal_heatmap() const;
  ::city::cognition::output::v1::Heatmap* _internal_mutable_heatmap();

  public:
  // .city.cognition.output.v1.Stat stat = 5 [json_name = "stat"];
  bool has_stat() const;
  void clear_stat() ;
  const ::city::cognition::output::v1::Stat& stat() const;
  PROTOBUF_NODISCARD ::city::cognition::output::v1::Stat* release_stat();
  ::city::cognition::output::v1::Stat* mutable_stat();
  void set_allocated_stat(::city::cognition::output::v1::Stat* value);
  void unsafe_arena_set_allocated_stat(::city::cognition::output::v1::Stat* value);
  ::city::cognition::output::v1::Stat* unsafe_arena_release_stat();

  private:
  const ::city::cognition::output::v1::Stat& _internal_stat() const;
  ::city::cognition::output::v1::Stat* _internal_mutable_stat();

  public:
  // int32 step = 1 [json_name = "step"];
  void clear_step() ;
  ::int32_t step() const;
  void set_step(::int32_t value);

  private:
  ::int32_t _internal_step() const;
  void _internal_set_step(::int32_t value);

  public:
  // @@protoc_insertion_point(class_scope:city.cognition.output.v1.OutputRequest)
 private:
  class _Internal;

  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<3, 8, 7, 0, 2> _table_;
  template <typename T> friend class ::google::protobuf::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::google::protobuf::internal::HasBits<1> _has_bits_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Node > nodes_;
    ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence > influences_;
    ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Content > contents_;
    ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence > group_influences_;
    ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Group > groups_;
    ::city::cognition::output::v1::Heatmap* heatmap_;
    ::city::cognition::output::v1::Stat* stat_;
    ::int32_t step_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_city_2fcognition_2foutput_2fv1_2foutput_5fservice_2eproto;
};// -------------------------------------------------------------------

class OutputResponse final :
    public ::google::protobuf::internal::ZeroFieldsBase /* @@protoc_insertion_point(class_definition:city.cognition.output.v1.OutputResponse) */ {
 public:
  inline OutputResponse() : OutputResponse(nullptr) {}
  template<typename = void>
  explicit PROTOBUF_CONSTEXPR OutputResponse(::google::protobuf::internal::ConstantInitialized);

  OutputResponse(const OutputResponse& from);
  OutputResponse(OutputResponse&& from) noexcept
    : OutputResponse() {
    *this = ::std::move(from);
  }

  inline OutputResponse& operator=(const OutputResponse& from) {
    CopyFrom(from);
    return *this;
  }
  inline OutputResponse& operator=(OutputResponse&& from) noexcept {
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
  static const OutputResponse& default_instance() {
    return *internal_default_instance();
  }
  static inline const OutputResponse* internal_default_instance() {
    return reinterpret_cast<const OutputResponse*>(
               &_OutputResponse_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(OutputResponse& a, OutputResponse& b) {
    a.Swap(&b);
  }
  inline void Swap(OutputResponse* other) {
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
  void UnsafeArenaSwap(OutputResponse* other) {
    if (other == this) return;
    ABSL_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  OutputResponse* New(::google::protobuf::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<OutputResponse>(arena);
  }
  using ::google::protobuf::internal::ZeroFieldsBase::CopyFrom;
  inline void CopyFrom(const OutputResponse& from) {
    ::google::protobuf::internal::ZeroFieldsBase::CopyImpl(*this, from);
  }
  using ::google::protobuf::internal::ZeroFieldsBase::MergeFrom;
  void MergeFrom(const OutputResponse& from) {
    ::google::protobuf::internal::ZeroFieldsBase::MergeImpl(*this, from);
  }
  public:

  private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() {
    return "city.cognition.output.v1.OutputResponse";
  }
  protected:
  explicit OutputResponse(::google::protobuf::Arena* arena);
  public:

  static const ClassData _class_data_;
  const ::google::protobuf::Message::ClassData*GetClassData() const final;

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // @@protoc_insertion_point(class_scope:city.cognition.output.v1.OutputResponse)
 private:
  class _Internal;

  template <typename T> friend class ::google::protobuf::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  friend struct ::TableStruct_city_2fcognition_2foutput_2fv1_2foutput_5fservice_2eproto;
};

// ===================================================================




// ===================================================================


#ifdef __GNUC__
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// OutputRequest

// int32 step = 1 [json_name = "step"];
inline void OutputRequest::clear_step() {
  _impl_.step_ = 0;
}
inline ::int32_t OutputRequest::step() const {
  // @@protoc_insertion_point(field_get:city.cognition.output.v1.OutputRequest.step)
  return _internal_step();
}
inline void OutputRequest::set_step(::int32_t value) {
  _internal_set_step(value);
  // @@protoc_insertion_point(field_set:city.cognition.output.v1.OutputRequest.step)
}
inline ::int32_t OutputRequest::_internal_step() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.step_;
}
inline void OutputRequest::_internal_set_step(::int32_t value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.step_ = value;
}

// repeated .city.cognition.output.v1.Node nodes = 2 [json_name = "nodes"];
inline int OutputRequest::_internal_nodes_size() const {
  return _internal_nodes().size();
}
inline int OutputRequest::nodes_size() const {
  return _internal_nodes_size();
}
inline ::city::cognition::output::v1::Node* OutputRequest::mutable_nodes(int index) {
  // @@protoc_insertion_point(field_mutable:city.cognition.output.v1.OutputRequest.nodes)
  return _internal_mutable_nodes()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Node >*
OutputRequest::mutable_nodes() {
  // @@protoc_insertion_point(field_mutable_list:city.cognition.output.v1.OutputRequest.nodes)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_nodes();
}
inline const ::city::cognition::output::v1::Node& OutputRequest::nodes(int index) const {
  // @@protoc_insertion_point(field_get:city.cognition.output.v1.OutputRequest.nodes)
    return _internal_nodes().Get(index);
}
inline ::city::cognition::output::v1::Node* OutputRequest::add_nodes() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::city::cognition::output::v1::Node* _add = _internal_mutable_nodes()->Add();
  // @@protoc_insertion_point(field_add:city.cognition.output.v1.OutputRequest.nodes)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Node >&
OutputRequest::nodes() const {
  // @@protoc_insertion_point(field_list:city.cognition.output.v1.OutputRequest.nodes)
  return _internal_nodes();
}
inline const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Node>&
OutputRequest::_internal_nodes() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.nodes_;
}
inline ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Node>*
OutputRequest::_internal_mutable_nodes() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.nodes_;
}

// repeated .city.cognition.output.v1.Influence influences = 3 [json_name = "influences"];
inline int OutputRequest::_internal_influences_size() const {
  return _internal_influences().size();
}
inline int OutputRequest::influences_size() const {
  return _internal_influences_size();
}
inline ::city::cognition::output::v1::Influence* OutputRequest::mutable_influences(int index) {
  // @@protoc_insertion_point(field_mutable:city.cognition.output.v1.OutputRequest.influences)
  return _internal_mutable_influences()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence >*
OutputRequest::mutable_influences() {
  // @@protoc_insertion_point(field_mutable_list:city.cognition.output.v1.OutputRequest.influences)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_influences();
}
inline const ::city::cognition::output::v1::Influence& OutputRequest::influences(int index) const {
  // @@protoc_insertion_point(field_get:city.cognition.output.v1.OutputRequest.influences)
    return _internal_influences().Get(index);
}
inline ::city::cognition::output::v1::Influence* OutputRequest::add_influences() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::city::cognition::output::v1::Influence* _add = _internal_mutable_influences()->Add();
  // @@protoc_insertion_point(field_add:city.cognition.output.v1.OutputRequest.influences)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence >&
OutputRequest::influences() const {
  // @@protoc_insertion_point(field_list:city.cognition.output.v1.OutputRequest.influences)
  return _internal_influences();
}
inline const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Influence>&
OutputRequest::_internal_influences() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.influences_;
}
inline ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Influence>*
OutputRequest::_internal_mutable_influences() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.influences_;
}

// .city.cognition.output.v1.Heatmap heatmap = 4 [json_name = "heatmap"];
inline bool OutputRequest::has_heatmap() const {
  bool value = (_impl_._has_bits_[0] & 0x00000001u) != 0;
  PROTOBUF_ASSUME(!value || _impl_.heatmap_ != nullptr);
  return value;
}
inline const ::city::cognition::output::v1::Heatmap& OutputRequest::_internal_heatmap() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  const ::city::cognition::output::v1::Heatmap* p = _impl_.heatmap_;
  return p != nullptr ? *p : reinterpret_cast<const ::city::cognition::output::v1::Heatmap&>(::city::cognition::output::v1::_Heatmap_default_instance_);
}
inline const ::city::cognition::output::v1::Heatmap& OutputRequest::heatmap() const {
  // @@protoc_insertion_point(field_get:city.cognition.output.v1.OutputRequest.heatmap)
  return _internal_heatmap();
}
inline void OutputRequest::unsafe_arena_set_allocated_heatmap(::city::cognition::output::v1::Heatmap* value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::google::protobuf::MessageLite*>(_impl_.heatmap_);
  }
  _impl_.heatmap_ = reinterpret_cast<::city::cognition::output::v1::Heatmap*>(value);
  if (value != nullptr) {
    _impl_._has_bits_[0] |= 0x00000001u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000001u;
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:city.cognition.output.v1.OutputRequest.heatmap)
}
inline ::city::cognition::output::v1::Heatmap* OutputRequest::release_heatmap() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);

  _impl_._has_bits_[0] &= ~0x00000001u;
  ::city::cognition::output::v1::Heatmap* released = _impl_.heatmap_;
  _impl_.heatmap_ = nullptr;
#ifdef PROTOBUF_FORCE_COPY_IN_RELEASE
  auto* old = reinterpret_cast<::google::protobuf::MessageLite*>(released);
  released = ::google::protobuf::internal::DuplicateIfNonNull(released);
  if (GetArenaForAllocation() == nullptr) {
    delete old;
  }
#else   // PROTOBUF_FORCE_COPY_IN_RELEASE
  if (GetArenaForAllocation() != nullptr) {
    released = ::google::protobuf::internal::DuplicateIfNonNull(released);
  }
#endif  // !PROTOBUF_FORCE_COPY_IN_RELEASE
  return released;
}
inline ::city::cognition::output::v1::Heatmap* OutputRequest::unsafe_arena_release_heatmap() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  // @@protoc_insertion_point(field_release:city.cognition.output.v1.OutputRequest.heatmap)

  _impl_._has_bits_[0] &= ~0x00000001u;
  ::city::cognition::output::v1::Heatmap* temp = _impl_.heatmap_;
  _impl_.heatmap_ = nullptr;
  return temp;
}
inline ::city::cognition::output::v1::Heatmap* OutputRequest::_internal_mutable_heatmap() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  _impl_._has_bits_[0] |= 0x00000001u;
  if (_impl_.heatmap_ == nullptr) {
    auto* p = CreateMaybeMessage<::city::cognition::output::v1::Heatmap>(GetArenaForAllocation());
    _impl_.heatmap_ = reinterpret_cast<::city::cognition::output::v1::Heatmap*>(p);
  }
  return _impl_.heatmap_;
}
inline ::city::cognition::output::v1::Heatmap* OutputRequest::mutable_heatmap() {
  ::city::cognition::output::v1::Heatmap* _msg = _internal_mutable_heatmap();
  // @@protoc_insertion_point(field_mutable:city.cognition.output.v1.OutputRequest.heatmap)
  return _msg;
}
inline void OutputRequest::set_allocated_heatmap(::city::cognition::output::v1::Heatmap* value) {
  ::google::protobuf::Arena* message_arena = GetArenaForAllocation();
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  if (message_arena == nullptr) {
    delete reinterpret_cast<::google::protobuf::MessageLite*>(_impl_.heatmap_);
  }

  if (value != nullptr) {
    ::google::protobuf::Arena* submessage_arena =
        ::google::protobuf::Arena::InternalGetOwningArena(reinterpret_cast<::google::protobuf::MessageLite*>(value));
    if (message_arena != submessage_arena) {
      value = ::google::protobuf::internal::GetOwnedMessage(message_arena, value, submessage_arena);
    }
    _impl_._has_bits_[0] |= 0x00000001u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000001u;
  }

  _impl_.heatmap_ = reinterpret_cast<::city::cognition::output::v1::Heatmap*>(value);
  // @@protoc_insertion_point(field_set_allocated:city.cognition.output.v1.OutputRequest.heatmap)
}

// .city.cognition.output.v1.Stat stat = 5 [json_name = "stat"];
inline bool OutputRequest::has_stat() const {
  bool value = (_impl_._has_bits_[0] & 0x00000002u) != 0;
  PROTOBUF_ASSUME(!value || _impl_.stat_ != nullptr);
  return value;
}
inline const ::city::cognition::output::v1::Stat& OutputRequest::_internal_stat() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  const ::city::cognition::output::v1::Stat* p = _impl_.stat_;
  return p != nullptr ? *p : reinterpret_cast<const ::city::cognition::output::v1::Stat&>(::city::cognition::output::v1::_Stat_default_instance_);
}
inline const ::city::cognition::output::v1::Stat& OutputRequest::stat() const {
  // @@protoc_insertion_point(field_get:city.cognition.output.v1.OutputRequest.stat)
  return _internal_stat();
}
inline void OutputRequest::unsafe_arena_set_allocated_stat(::city::cognition::output::v1::Stat* value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::google::protobuf::MessageLite*>(_impl_.stat_);
  }
  _impl_.stat_ = reinterpret_cast<::city::cognition::output::v1::Stat*>(value);
  if (value != nullptr) {
    _impl_._has_bits_[0] |= 0x00000002u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000002u;
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:city.cognition.output.v1.OutputRequest.stat)
}
inline ::city::cognition::output::v1::Stat* OutputRequest::release_stat() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);

  _impl_._has_bits_[0] &= ~0x00000002u;
  ::city::cognition::output::v1::Stat* released = _impl_.stat_;
  _impl_.stat_ = nullptr;
#ifdef PROTOBUF_FORCE_COPY_IN_RELEASE
  auto* old = reinterpret_cast<::google::protobuf::MessageLite*>(released);
  released = ::google::protobuf::internal::DuplicateIfNonNull(released);
  if (GetArenaForAllocation() == nullptr) {
    delete old;
  }
#else   // PROTOBUF_FORCE_COPY_IN_RELEASE
  if (GetArenaForAllocation() != nullptr) {
    released = ::google::protobuf::internal::DuplicateIfNonNull(released);
  }
#endif  // !PROTOBUF_FORCE_COPY_IN_RELEASE
  return released;
}
inline ::city::cognition::output::v1::Stat* OutputRequest::unsafe_arena_release_stat() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  // @@protoc_insertion_point(field_release:city.cognition.output.v1.OutputRequest.stat)

  _impl_._has_bits_[0] &= ~0x00000002u;
  ::city::cognition::output::v1::Stat* temp = _impl_.stat_;
  _impl_.stat_ = nullptr;
  return temp;
}
inline ::city::cognition::output::v1::Stat* OutputRequest::_internal_mutable_stat() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  _impl_._has_bits_[0] |= 0x00000002u;
  if (_impl_.stat_ == nullptr) {
    auto* p = CreateMaybeMessage<::city::cognition::output::v1::Stat>(GetArenaForAllocation());
    _impl_.stat_ = reinterpret_cast<::city::cognition::output::v1::Stat*>(p);
  }
  return _impl_.stat_;
}
inline ::city::cognition::output::v1::Stat* OutputRequest::mutable_stat() {
  ::city::cognition::output::v1::Stat* _msg = _internal_mutable_stat();
  // @@protoc_insertion_point(field_mutable:city.cognition.output.v1.OutputRequest.stat)
  return _msg;
}
inline void OutputRequest::set_allocated_stat(::city::cognition::output::v1::Stat* value) {
  ::google::protobuf::Arena* message_arena = GetArenaForAllocation();
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  if (message_arena == nullptr) {
    delete reinterpret_cast<::google::protobuf::MessageLite*>(_impl_.stat_);
  }

  if (value != nullptr) {
    ::google::protobuf::Arena* submessage_arena =
        ::google::protobuf::Arena::InternalGetOwningArena(reinterpret_cast<::google::protobuf::MessageLite*>(value));
    if (message_arena != submessage_arena) {
      value = ::google::protobuf::internal::GetOwnedMessage(message_arena, value, submessage_arena);
    }
    _impl_._has_bits_[0] |= 0x00000002u;
  } else {
    _impl_._has_bits_[0] &= ~0x00000002u;
  }

  _impl_.stat_ = reinterpret_cast<::city::cognition::output::v1::Stat*>(value);
  // @@protoc_insertion_point(field_set_allocated:city.cognition.output.v1.OutputRequest.stat)
}

// repeated .city.cognition.output.v1.Content contents = 6 [json_name = "contents"];
inline int OutputRequest::_internal_contents_size() const {
  return _internal_contents().size();
}
inline int OutputRequest::contents_size() const {
  return _internal_contents_size();
}
inline ::city::cognition::output::v1::Content* OutputRequest::mutable_contents(int index) {
  // @@protoc_insertion_point(field_mutable:city.cognition.output.v1.OutputRequest.contents)
  return _internal_mutable_contents()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Content >*
OutputRequest::mutable_contents() {
  // @@protoc_insertion_point(field_mutable_list:city.cognition.output.v1.OutputRequest.contents)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_contents();
}
inline const ::city::cognition::output::v1::Content& OutputRequest::contents(int index) const {
  // @@protoc_insertion_point(field_get:city.cognition.output.v1.OutputRequest.contents)
    return _internal_contents().Get(index);
}
inline ::city::cognition::output::v1::Content* OutputRequest::add_contents() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::city::cognition::output::v1::Content* _add = _internal_mutable_contents()->Add();
  // @@protoc_insertion_point(field_add:city.cognition.output.v1.OutputRequest.contents)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Content >&
OutputRequest::contents() const {
  // @@protoc_insertion_point(field_list:city.cognition.output.v1.OutputRequest.contents)
  return _internal_contents();
}
inline const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Content>&
OutputRequest::_internal_contents() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.contents_;
}
inline ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Content>*
OutputRequest::_internal_mutable_contents() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.contents_;
}

// repeated .city.cognition.output.v1.Influence group_influences = 7 [json_name = "groupInfluences"];
inline int OutputRequest::_internal_group_influences_size() const {
  return _internal_group_influences().size();
}
inline int OutputRequest::group_influences_size() const {
  return _internal_group_influences_size();
}
inline ::city::cognition::output::v1::Influence* OutputRequest::mutable_group_influences(int index) {
  // @@protoc_insertion_point(field_mutable:city.cognition.output.v1.OutputRequest.group_influences)
  return _internal_mutable_group_influences()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence >*
OutputRequest::mutable_group_influences() {
  // @@protoc_insertion_point(field_mutable_list:city.cognition.output.v1.OutputRequest.group_influences)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_group_influences();
}
inline const ::city::cognition::output::v1::Influence& OutputRequest::group_influences(int index) const {
  // @@protoc_insertion_point(field_get:city.cognition.output.v1.OutputRequest.group_influences)
    return _internal_group_influences().Get(index);
}
inline ::city::cognition::output::v1::Influence* OutputRequest::add_group_influences() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::city::cognition::output::v1::Influence* _add = _internal_mutable_group_influences()->Add();
  // @@protoc_insertion_point(field_add:city.cognition.output.v1.OutputRequest.group_influences)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Influence >&
OutputRequest::group_influences() const {
  // @@protoc_insertion_point(field_list:city.cognition.output.v1.OutputRequest.group_influences)
  return _internal_group_influences();
}
inline const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Influence>&
OutputRequest::_internal_group_influences() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.group_influences_;
}
inline ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Influence>*
OutputRequest::_internal_mutable_group_influences() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.group_influences_;
}

// repeated .city.cognition.output.v1.Group groups = 8 [json_name = "groups"];
inline int OutputRequest::_internal_groups_size() const {
  return _internal_groups().size();
}
inline int OutputRequest::groups_size() const {
  return _internal_groups_size();
}
inline ::city::cognition::output::v1::Group* OutputRequest::mutable_groups(int index) {
  // @@protoc_insertion_point(field_mutable:city.cognition.output.v1.OutputRequest.groups)
  return _internal_mutable_groups()->Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Group >*
OutputRequest::mutable_groups() {
  // @@protoc_insertion_point(field_mutable_list:city.cognition.output.v1.OutputRequest.groups)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  return _internal_mutable_groups();
}
inline const ::city::cognition::output::v1::Group& OutputRequest::groups(int index) const {
  // @@protoc_insertion_point(field_get:city.cognition.output.v1.OutputRequest.groups)
    return _internal_groups().Get(index);
}
inline ::city::cognition::output::v1::Group* OutputRequest::add_groups() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::city::cognition::output::v1::Group* _add = _internal_mutable_groups()->Add();
  // @@protoc_insertion_point(field_add:city.cognition.output.v1.OutputRequest.groups)
  return _add;
}
inline const ::google::protobuf::RepeatedPtrField< ::city::cognition::output::v1::Group >&
OutputRequest::groups() const {
  // @@protoc_insertion_point(field_list:city.cognition.output.v1.OutputRequest.groups)
  return _internal_groups();
}
inline const ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Group>&
OutputRequest::_internal_groups() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.groups_;
}
inline ::google::protobuf::RepeatedPtrField<::city::cognition::output::v1::Group>*
OutputRequest::_internal_mutable_groups() {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return &_impl_.groups_;
}

// -------------------------------------------------------------------

// OutputResponse

#ifdef __GNUC__
#pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace output
}  // namespace cognition
}  // namespace city


// @@protoc_insertion_point(global_scope)

#include "google/protobuf/port_undef.inc"

#endif  // GOOGLE_PROTOBUF_INCLUDED_city_2fcognition_2foutput_2fv1_2foutput_5fservice_2eproto_2epb_2eh
