// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/event/v2/event_service.proto

#include "city/event/v2/event_service.pb.h"

#include <algorithm>
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/extension_set.h"
#include "google/protobuf/wire_format_lite.h"
#include "google/protobuf/descriptor.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/reflection_ops.h"
#include "google/protobuf/wire_format.h"
#include "google/protobuf/generated_message_tctable_impl.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"
PROTOBUF_PRAGMA_INIT_SEG
namespace _pb = ::google::protobuf;
namespace _pbi = ::google::protobuf::internal;
namespace _fl = ::google::protobuf::internal::field_layout;
namespace city {
namespace event {
namespace v2 {
        template <typename>
PROTOBUF_CONSTEXPR GetEventsByTopicRequest::GetEventsByTopicRequest(::_pbi::ConstantInitialized)
    : _impl_{
      /*decltype(_impl_.topic_)*/ {
          &::_pbi::fixed_address_empty_string,
          ::_pbi::ConstantInitialized{},
      },
      /*decltype(_impl_._cached_size_)*/ {},
    } {}
struct GetEventsByTopicRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR GetEventsByTopicRequestDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~GetEventsByTopicRequestDefaultTypeInternal() {}
  union {
    GetEventsByTopicRequest _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 GetEventsByTopicRequestDefaultTypeInternal _GetEventsByTopicRequest_default_instance_;
        template <typename>
PROTOBUF_CONSTEXPR GetEventsByTopicResponse::GetEventsByTopicResponse(::_pbi::ConstantInitialized)
    : _impl_{
      /*decltype(_impl_.events_)*/ {},
      /*decltype(_impl_._cached_size_)*/ {},
    } {}
struct GetEventsByTopicResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR GetEventsByTopicResponseDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~GetEventsByTopicResponseDefaultTypeInternal() {}
  union {
    GetEventsByTopicResponse _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 GetEventsByTopicResponseDefaultTypeInternal _GetEventsByTopicResponse_default_instance_;
        template <typename>
PROTOBUF_CONSTEXPR ResolveEventsRequest::ResolveEventsRequest(::_pbi::ConstantInitialized)
    : _impl_{
      /*decltype(_impl_.events_)*/ {},
      /*decltype(_impl_._cached_size_)*/ {},
    } {}
struct ResolveEventsRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR ResolveEventsRequestDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~ResolveEventsRequestDefaultTypeInternal() {}
  union {
    ResolveEventsRequest _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 ResolveEventsRequestDefaultTypeInternal _ResolveEventsRequest_default_instance_;
      template <typename>
PROTOBUF_CONSTEXPR ResolveEventsResponse::ResolveEventsResponse(::_pbi::ConstantInitialized) {}
struct ResolveEventsResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR ResolveEventsResponseDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~ResolveEventsResponseDefaultTypeInternal() {}
  union {
    ResolveEventsResponse _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 ResolveEventsResponseDefaultTypeInternal _ResolveEventsResponse_default_instance_;
}  // namespace v2
}  // namespace event
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fevent_2fv2_2fevent_5fservice_2eproto[4];
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_city_2fevent_2fv2_2fevent_5fservice_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_city_2fevent_2fv2_2fevent_5fservice_2eproto = nullptr;
const ::uint32_t TableStruct_city_2fevent_2fv2_2fevent_5fservice_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(
    protodesc_cold) = {
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::city::event::v2::GetEventsByTopicRequest, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::city::event::v2::GetEventsByTopicRequest, _impl_.topic_),
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::city::event::v2::GetEventsByTopicResponse, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::city::event::v2::GetEventsByTopicResponse, _impl_.events_),
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::city::event::v2::ResolveEventsRequest, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::city::event::v2::ResolveEventsRequest, _impl_.events_),
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::city::event::v2::ResolveEventsResponse, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
};

static const ::_pbi::MigrationSchema
    schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
        {0, -1, -1, sizeof(::city::event::v2::GetEventsByTopicRequest)},
        {9, -1, -1, sizeof(::city::event::v2::GetEventsByTopicResponse)},
        {18, -1, -1, sizeof(::city::event::v2::ResolveEventsRequest)},
        {27, -1, -1, sizeof(::city::event::v2::ResolveEventsResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
    &::city::event::v2::_GetEventsByTopicRequest_default_instance_._instance,
    &::city::event::v2::_GetEventsByTopicResponse_default_instance_._instance,
    &::city::event::v2::_ResolveEventsRequest_default_instance_._instance,
    &::city::event::v2::_ResolveEventsResponse_default_instance_._instance,
};
const char descriptor_table_protodef_city_2fevent_2fv2_2fevent_5fservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n!city/event/v2/event_service.proto\022\rcit"
    "y.event.v2\032\031city/event/v2/event.proto\"/\n"
    "\027GetEventsByTopicRequest\022\024\n\005topic\030\001 \001(\tR"
    "\005topic\"H\n\030GetEventsByTopicResponse\022,\n\006ev"
    "ents\030\001 \003(\0132\024.city.event.v2.EventR\006events"
    "\"D\n\024ResolveEventsRequest\022,\n\006events\030\001 \003(\013"
    "2\024.city.event.v2.EventR\006events\"\027\n\025Resolv"
    "eEventsResponse2\317\001\n\014EventService\022c\n\020GetE"
    "ventsByTopic\022&.city.event.v2.GetEventsBy"
    "TopicRequest\032\'.city.event.v2.GetEventsBy"
    "TopicResponse\022Z\n\rResolveEvents\022#.city.ev"
    "ent.v2.ResolveEventsRequest\032$.city.event"
    ".v2.ResolveEventsResponseB\260\001\n\021com.city.e"
    "vent.v2B\021EventServiceProtoP\001Z2git.fiblab"
    ".net/sim/protos/go/city/event/v2;eventv2"
    "\242\002\003CEX\252\002\rCity.Event.V2\312\002\rCity\\Event\\V2\342\002"
    "\031City\\Event\\V2\\GPBMetadata\352\002\017City::Event"
    "::V2b\006proto3"
};
static const ::_pbi::DescriptorTable* const descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_deps[1] =
    {
        &::descriptor_table_city_2fevent_2fv2_2fevent_2eproto,
};
static ::absl::once_flag descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto = {
    false,
    false,
    692,
    descriptor_table_protodef_city_2fevent_2fv2_2fevent_5fservice_2eproto,
    "city/event/v2/event_service.proto",
    &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_once,
    descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_deps,
    1,
    4,
    schemas,
    file_default_instances,
    TableStruct_city_2fevent_2fv2_2fevent_5fservice_2eproto::offsets,
    file_level_metadata_city_2fevent_2fv2_2fevent_5fservice_2eproto,
    file_level_enum_descriptors_city_2fevent_2fv2_2fevent_5fservice_2eproto,
    file_level_service_descriptors_city_2fevent_2fv2_2fevent_5fservice_2eproto,
};

// This function exists to be marked as weak.
// It can significantly speed up compilation by breaking up LLVM's SCC
// in the .pb.cc translation units. Large translation units see a
// reduction of more than 35% of walltime for optimized builds. Without
// the weak attribute all the messages in the file, including all the
// vtables and everything they use become part of the same SCC through
// a cycle like:
// GetMetadata -> descriptor table -> default instances ->
//   vtables -> GetMetadata
// By adding a weak function here we break the connection from the
// individual vtables back into the descriptor table.
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_getter() {
  return &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fevent_2fv2_2fevent_5fservice_2eproto(&descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto);
namespace city {
namespace event {
namespace v2 {
// ===================================================================

class GetEventsByTopicRequest::_Internal {
 public:
};

GetEventsByTopicRequest::GetEventsByTopicRequest(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:city.event.v2.GetEventsByTopicRequest)
}
GetEventsByTopicRequest::GetEventsByTopicRequest(const GetEventsByTopicRequest& from) : ::google::protobuf::Message() {
  GetEventsByTopicRequest* const _this = this;
  (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.topic_){},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);
  _impl_.topic_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.topic_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_topic().empty()) {
    _this->_impl_.topic_.Set(from._internal_topic(), _this->GetArenaForAllocation());
  }

  // @@protoc_insertion_point(copy_constructor:city.event.v2.GetEventsByTopicRequest)
}
inline void GetEventsByTopicRequest::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.topic_){},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _impl_.topic_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.topic_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}
GetEventsByTopicRequest::~GetEventsByTopicRequest() {
  // @@protoc_insertion_point(destructor:city.event.v2.GetEventsByTopicRequest)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void GetEventsByTopicRequest::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.topic_.Destroy();
}
void GetEventsByTopicRequest::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

PROTOBUF_NOINLINE void GetEventsByTopicRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:city.event.v2.GetEventsByTopicRequest)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.topic_.ClearToEmpty();
  _internal_metadata_.Clear<::google::protobuf::UnknownFieldSet>();
}

const char* GetEventsByTopicRequest::_InternalParse(
    const char* ptr, ::_pbi::ParseContext* ctx) {
  ptr = ::_pbi::TcParser::ParseLoop(this, ptr, ctx, &_table_.header);
  return ptr;
}


PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1
const ::_pbi::TcParseTable<0, 1, 0, 51, 2> GetEventsByTopicRequest::_table_ = {
  {
    0,  // no _has_bits_
    0, // no _extensions_
    1, 0,  // max_field_number, fast_idx_mask
    offsetof(decltype(_table_), field_lookup_table),
    4294967294,  // skipmap
    offsetof(decltype(_table_), field_entries),
    1,  // num_field_entries
    0,  // num_aux_entries
    offsetof(decltype(_table_), field_names),  // no aux_entries
    &_GetEventsByTopicRequest_default_instance_._instance,
    ::_pbi::TcParser::GenericFallback,  // fallback
  }, {{
    // string topic = 1 [json_name = "topic"];
    {::_pbi::TcParser::FastUS1,
     {10, 63, 0, PROTOBUF_FIELD_OFFSET(GetEventsByTopicRequest, _impl_.topic_)}},
  }}, {{
    65535, 65535
  }}, {{
    // string topic = 1 [json_name = "topic"];
    {PROTOBUF_FIELD_OFFSET(GetEventsByTopicRequest, _impl_.topic_), 0, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kUtf8String | ::_fl::kRepAString)},
  }},
  // no aux_entries
  {{
    "\45\5\0\0\0\0\0\0"
    "city.event.v2.GetEventsByTopicRequest"
    "topic"
  }},
};

::uint8_t* GetEventsByTopicRequest::_InternalSerialize(
    ::uint8_t* target,
    ::google::protobuf::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.event.v2.GetEventsByTopicRequest)
  ::uint32_t cached_has_bits = 0;
  (void)cached_has_bits;

  // string topic = 1 [json_name = "topic"];
  if (!this->_internal_topic().empty()) {
    const std::string& _s = this->_internal_topic();
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::google::protobuf::internal::WireFormatLite::SERIALIZE, "city.event.v2.GetEventsByTopicRequest.topic");
    target = stream->WriteStringMaybeAliased(1, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target =
        ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
            _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.event.v2.GetEventsByTopicRequest)
  return target;
}

::size_t GetEventsByTopicRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.event.v2.GetEventsByTopicRequest)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string topic = 1 [json_name = "topic"];
  if (!this->_internal_topic().empty()) {
    total_size += 1 + ::google::protobuf::internal::WireFormatLite::StringSize(
                                    this->_internal_topic());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::google::protobuf::Message::ClassData GetEventsByTopicRequest::_class_data_ = {
    ::google::protobuf::Message::CopyWithSourceCheck,
    GetEventsByTopicRequest::MergeImpl
};
const ::google::protobuf::Message::ClassData*GetEventsByTopicRequest::GetClassData() const { return &_class_data_; }


void GetEventsByTopicRequest::MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg) {
  auto* const _this = static_cast<GetEventsByTopicRequest*>(&to_msg);
  auto& from = static_cast<const GetEventsByTopicRequest&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.event.v2.GetEventsByTopicRequest)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_topic().empty()) {
    _this->_internal_set_topic(from._internal_topic());
  }
  _this->_internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(from._internal_metadata_);
}

void GetEventsByTopicRequest::CopyFrom(const GetEventsByTopicRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.event.v2.GetEventsByTopicRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

PROTOBUF_NOINLINE bool GetEventsByTopicRequest::IsInitialized() const {
  return true;
}

void GetEventsByTopicRequest::InternalSwap(GetEventsByTopicRequest* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.topic_, lhs_arena,
                                       &other->_impl_.topic_, rhs_arena);
}

::google::protobuf::Metadata GetEventsByTopicRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_getter, &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_once,
      file_level_metadata_city_2fevent_2fv2_2fevent_5fservice_2eproto[0]);
}
// ===================================================================

class GetEventsByTopicResponse::_Internal {
 public:
};

void GetEventsByTopicResponse::clear_events() {
  _internal_mutable_events()->Clear();
}
GetEventsByTopicResponse::GetEventsByTopicResponse(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:city.event.v2.GetEventsByTopicResponse)
}
GetEventsByTopicResponse::GetEventsByTopicResponse(const GetEventsByTopicResponse& from) : ::google::protobuf::Message() {
  GetEventsByTopicResponse* const _this = this;
  (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.events_){from._impl_.events_},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);

  // @@protoc_insertion_point(copy_constructor:city.event.v2.GetEventsByTopicResponse)
}
inline void GetEventsByTopicResponse::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.events_){arena},
      /*decltype(_impl_._cached_size_)*/ {},
  };
}
GetEventsByTopicResponse::~GetEventsByTopicResponse() {
  // @@protoc_insertion_point(destructor:city.event.v2.GetEventsByTopicResponse)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void GetEventsByTopicResponse::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.events_.~RepeatedPtrField();
}
void GetEventsByTopicResponse::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

PROTOBUF_NOINLINE void GetEventsByTopicResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:city.event.v2.GetEventsByTopicResponse)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _internal_mutable_events()->Clear();
  _internal_metadata_.Clear<::google::protobuf::UnknownFieldSet>();
}

const char* GetEventsByTopicResponse::_InternalParse(
    const char* ptr, ::_pbi::ParseContext* ctx) {
  ptr = ::_pbi::TcParser::ParseLoop(this, ptr, ctx, &_table_.header);
  return ptr;
}


PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1
const ::_pbi::TcParseTable<0, 1, 1, 0, 2> GetEventsByTopicResponse::_table_ = {
  {
    0,  // no _has_bits_
    0, // no _extensions_
    1, 0,  // max_field_number, fast_idx_mask
    offsetof(decltype(_table_), field_lookup_table),
    4294967294,  // skipmap
    offsetof(decltype(_table_), field_entries),
    1,  // num_field_entries
    1,  // num_aux_entries
    offsetof(decltype(_table_), aux_entries),
    &_GetEventsByTopicResponse_default_instance_._instance,
    ::_pbi::TcParser::GenericFallback,  // fallback
  }, {{
    // repeated .city.event.v2.Event events = 1 [json_name = "events"];
    {::_pbi::TcParser::FastMtR1,
     {10, 63, 0, PROTOBUF_FIELD_OFFSET(GetEventsByTopicResponse, _impl_.events_)}},
  }}, {{
    65535, 65535
  }}, {{
    // repeated .city.event.v2.Event events = 1 [json_name = "events"];
    {PROTOBUF_FIELD_OFFSET(GetEventsByTopicResponse, _impl_.events_), 0, 0,
    (0 | ::_fl::kFcRepeated | ::_fl::kMessage | ::_fl::kTvTable)},
  }}, {{
    {::_pbi::TcParser::GetTable<::city::event::v2::Event>()},
  }}, {{
  }},
};

::uint8_t* GetEventsByTopicResponse::_InternalSerialize(
    ::uint8_t* target,
    ::google::protobuf::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.event.v2.GetEventsByTopicResponse)
  ::uint32_t cached_has_bits = 0;
  (void)cached_has_bits;

  // repeated .city.event.v2.Event events = 1 [json_name = "events"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_events_size()); i < n; i++) {
    const auto& repfield = this->_internal_events().Get(i);
    target = ::google::protobuf::internal::WireFormatLite::
        InternalWriteMessage(1, repfield, repfield.GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target =
        ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
            _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.event.v2.GetEventsByTopicResponse)
  return target;
}

::size_t GetEventsByTopicResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.event.v2.GetEventsByTopicResponse)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.event.v2.Event events = 1 [json_name = "events"];
  total_size += 1UL * this->_internal_events_size();
  for (const auto& msg : this->_internal_events()) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSize(msg);
  }
  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::google::protobuf::Message::ClassData GetEventsByTopicResponse::_class_data_ = {
    ::google::protobuf::Message::CopyWithSourceCheck,
    GetEventsByTopicResponse::MergeImpl
};
const ::google::protobuf::Message::ClassData*GetEventsByTopicResponse::GetClassData() const { return &_class_data_; }


void GetEventsByTopicResponse::MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg) {
  auto* const _this = static_cast<GetEventsByTopicResponse*>(&to_msg);
  auto& from = static_cast<const GetEventsByTopicResponse&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.event.v2.GetEventsByTopicResponse)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_internal_mutable_events()->MergeFrom(from._internal_events());
  _this->_internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(from._internal_metadata_);
}

void GetEventsByTopicResponse::CopyFrom(const GetEventsByTopicResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.event.v2.GetEventsByTopicResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

PROTOBUF_NOINLINE bool GetEventsByTopicResponse::IsInitialized() const {
  return true;
}

void GetEventsByTopicResponse::InternalSwap(GetEventsByTopicResponse* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.events_.InternalSwap(&other->_impl_.events_);
}

::google::protobuf::Metadata GetEventsByTopicResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_getter, &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_once,
      file_level_metadata_city_2fevent_2fv2_2fevent_5fservice_2eproto[1]);
}
// ===================================================================

class ResolveEventsRequest::_Internal {
 public:
};

void ResolveEventsRequest::clear_events() {
  _internal_mutable_events()->Clear();
}
ResolveEventsRequest::ResolveEventsRequest(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:city.event.v2.ResolveEventsRequest)
}
ResolveEventsRequest::ResolveEventsRequest(const ResolveEventsRequest& from) : ::google::protobuf::Message() {
  ResolveEventsRequest* const _this = this;
  (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.events_){from._impl_.events_},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);

  // @@protoc_insertion_point(copy_constructor:city.event.v2.ResolveEventsRequest)
}
inline void ResolveEventsRequest::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.events_){arena},
      /*decltype(_impl_._cached_size_)*/ {},
  };
}
ResolveEventsRequest::~ResolveEventsRequest() {
  // @@protoc_insertion_point(destructor:city.event.v2.ResolveEventsRequest)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void ResolveEventsRequest::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.events_.~RepeatedPtrField();
}
void ResolveEventsRequest::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

PROTOBUF_NOINLINE void ResolveEventsRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:city.event.v2.ResolveEventsRequest)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _internal_mutable_events()->Clear();
  _internal_metadata_.Clear<::google::protobuf::UnknownFieldSet>();
}

const char* ResolveEventsRequest::_InternalParse(
    const char* ptr, ::_pbi::ParseContext* ctx) {
  ptr = ::_pbi::TcParser::ParseLoop(this, ptr, ctx, &_table_.header);
  return ptr;
}


PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1
const ::_pbi::TcParseTable<0, 1, 1, 0, 2> ResolveEventsRequest::_table_ = {
  {
    0,  // no _has_bits_
    0, // no _extensions_
    1, 0,  // max_field_number, fast_idx_mask
    offsetof(decltype(_table_), field_lookup_table),
    4294967294,  // skipmap
    offsetof(decltype(_table_), field_entries),
    1,  // num_field_entries
    1,  // num_aux_entries
    offsetof(decltype(_table_), aux_entries),
    &_ResolveEventsRequest_default_instance_._instance,
    ::_pbi::TcParser::GenericFallback,  // fallback
  }, {{
    // repeated .city.event.v2.Event events = 1 [json_name = "events"];
    {::_pbi::TcParser::FastMtR1,
     {10, 63, 0, PROTOBUF_FIELD_OFFSET(ResolveEventsRequest, _impl_.events_)}},
  }}, {{
    65535, 65535
  }}, {{
    // repeated .city.event.v2.Event events = 1 [json_name = "events"];
    {PROTOBUF_FIELD_OFFSET(ResolveEventsRequest, _impl_.events_), 0, 0,
    (0 | ::_fl::kFcRepeated | ::_fl::kMessage | ::_fl::kTvTable)},
  }}, {{
    {::_pbi::TcParser::GetTable<::city::event::v2::Event>()},
  }}, {{
  }},
};

::uint8_t* ResolveEventsRequest::_InternalSerialize(
    ::uint8_t* target,
    ::google::protobuf::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.event.v2.ResolveEventsRequest)
  ::uint32_t cached_has_bits = 0;
  (void)cached_has_bits;

  // repeated .city.event.v2.Event events = 1 [json_name = "events"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_events_size()); i < n; i++) {
    const auto& repfield = this->_internal_events().Get(i);
    target = ::google::protobuf::internal::WireFormatLite::
        InternalWriteMessage(1, repfield, repfield.GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target =
        ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
            _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.event.v2.ResolveEventsRequest)
  return target;
}

::size_t ResolveEventsRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.event.v2.ResolveEventsRequest)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.event.v2.Event events = 1 [json_name = "events"];
  total_size += 1UL * this->_internal_events_size();
  for (const auto& msg : this->_internal_events()) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSize(msg);
  }
  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::google::protobuf::Message::ClassData ResolveEventsRequest::_class_data_ = {
    ::google::protobuf::Message::CopyWithSourceCheck,
    ResolveEventsRequest::MergeImpl
};
const ::google::protobuf::Message::ClassData*ResolveEventsRequest::GetClassData() const { return &_class_data_; }


void ResolveEventsRequest::MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg) {
  auto* const _this = static_cast<ResolveEventsRequest*>(&to_msg);
  auto& from = static_cast<const ResolveEventsRequest&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.event.v2.ResolveEventsRequest)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_internal_mutable_events()->MergeFrom(from._internal_events());
  _this->_internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(from._internal_metadata_);
}

void ResolveEventsRequest::CopyFrom(const ResolveEventsRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.event.v2.ResolveEventsRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

PROTOBUF_NOINLINE bool ResolveEventsRequest::IsInitialized() const {
  return true;
}

void ResolveEventsRequest::InternalSwap(ResolveEventsRequest* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.events_.InternalSwap(&other->_impl_.events_);
}

::google::protobuf::Metadata ResolveEventsRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_getter, &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_once,
      file_level_metadata_city_2fevent_2fv2_2fevent_5fservice_2eproto[2]);
}
// ===================================================================

class ResolveEventsResponse::_Internal {
 public:
};

ResolveEventsResponse::ResolveEventsResponse(::google::protobuf::Arena* arena)
    : ::google::protobuf::internal::ZeroFieldsBase(arena) {
  // @@protoc_insertion_point(arena_constructor:city.event.v2.ResolveEventsResponse)
}
ResolveEventsResponse::ResolveEventsResponse(const ResolveEventsResponse& from) : ::google::protobuf::internal::ZeroFieldsBase() {
  ResolveEventsResponse* const _this = this;
  (void)_this;
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);

  // @@protoc_insertion_point(copy_constructor:city.event.v2.ResolveEventsResponse)
}




const ::google::protobuf::Message::ClassData ResolveEventsResponse::_class_data_ = {
    ::google::protobuf::internal::ZeroFieldsBase::CopyImpl,
    ::google::protobuf::internal::ZeroFieldsBase::MergeImpl,
};
const ::google::protobuf::Message::ClassData*ResolveEventsResponse::GetClassData() const { return &_class_data_; }







::google::protobuf::Metadata ResolveEventsResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_getter, &descriptor_table_city_2fevent_2fv2_2fevent_5fservice_2eproto_once,
      file_level_metadata_city_2fevent_2fv2_2fevent_5fservice_2eproto[3]);
}
// @@protoc_insertion_point(namespace_scope)
}  // namespace v2
}  // namespace event
}  // namespace city
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"
