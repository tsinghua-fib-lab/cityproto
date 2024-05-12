// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/comm/output/v1/output_service.proto

#include "city/comm/output/v1/output_service.pb.h"

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
namespace comm {
namespace output {
namespace v1 {
        template <typename>
PROTOBUF_CONSTEXPR OutputRequest::OutputRequest(::_pbi::ConstantInitialized)
    : _impl_{
      /*decltype(_impl_._has_bits_)*/ {},
      /*decltype(_impl_._cached_size_)*/ {},
      /*decltype(_impl_.nodes_)*/ {},
      /*decltype(_impl_.base_stations_)*/ {},
      /*decltype(_impl_.persons_)*/ {},
      /*decltype(_impl_.aois_)*/ {},
      /*decltype(_impl_.signal_heatmap_)*/ nullptr,
      /*decltype(_impl_.events_)*/ nullptr,
      /*decltype(_impl_.statistics_)*/ nullptr,
      /*decltype(_impl_.small_signal_heatmap_)*/ nullptr,
      /*decltype(_impl_.step_)*/ 0,
    } {}
struct OutputRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR OutputRequestDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~OutputRequestDefaultTypeInternal() {}
  union {
    OutputRequest _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 OutputRequestDefaultTypeInternal _OutputRequest_default_instance_;
      template <typename>
PROTOBUF_CONSTEXPR OutputResponse::OutputResponse(::_pbi::ConstantInitialized) {}
struct OutputResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR OutputResponseDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~OutputResponseDefaultTypeInternal() {}
  union {
    OutputResponse _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 OutputResponseDefaultTypeInternal _OutputResponse_default_instance_;
}  // namespace v1
}  // namespace output
}  // namespace comm
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto[2];
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto = nullptr;
const ::uint32_t TableStruct_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(
    protodesc_cold) = {
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_._has_bits_),
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_.step_),
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_.nodes_),
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_.base_stations_),
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_.signal_heatmap_),
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_.small_signal_heatmap_),
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_.persons_),
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_.aois_),
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_.events_),
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputRequest, _impl_.statistics_),
    ~0u,
    ~0u,
    ~0u,
    0,
    3,
    ~0u,
    ~0u,
    1,
    2,
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::city::comm::output::v1::OutputResponse, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
};

static const ::_pbi::MigrationSchema
    schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
        {0, 17, -1, sizeof(::city::comm::output::v1::OutputRequest)},
        {26, -1, -1, sizeof(::city::comm::output::v1::OutputResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
    &::city::comm::output::v1::_OutputRequest_default_instance_._instance,
    &::city::comm::output::v1::_OutputResponse_default_instance_._instance,
};
const char descriptor_table_protodef_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n(city/comm/output/v1/output_service.pro"
    "to\022\023city.comm.output.v1\032 city/comm/outpu"
    "t/v1/output.proto\032\031city/event/v1/event.p"
    "roto\"\203\004\n\rOutputRequest\022\022\n\004step\030\001 \001(\005R\004st"
    "ep\022/\n\005nodes\030\002 \003(\0132\031.city.comm.output.v1."
    "NodeR\005nodes\022E\n\rbase_stations\030\003 \003(\0132 .cit"
    "y.comm.output.v1.BaseStationR\014baseStatio"
    "ns\022B\n\016signal_heatmap\030\004 \001(\0132\033.city.comm.o"
    "utput.v1.SignalR\rsignalHeatmap\022M\n\024small_"
    "signal_heatmap\030\t \001(\0132\033.city.comm.output."
    "v1.SignalR\022smallSignalHeatmap\0225\n\007persons"
    "\030\005 \003(\0132\033.city.comm.output.v1.PersonR\007per"
    "sons\022,\n\004aois\030\006 \003(\0132\030.city.comm.output.v1"
    ".AoiR\004aois\022-\n\006events\030\007 \001(\0132\025.city.event."
    "v1.EventsR\006events\022\?\n\nstatistics\030\010 \001(\0132\037."
    "city.comm.output.v1.StatisticsR\nstatisti"
    "cs\"\020\n\016OutputResponse2b\n\rOutputService\022Q\n"
    "\006Output\022\".city.comm.output.v1.OutputRequ"
    "est\032#.city.comm.output.v1.OutputResponse"
    "B\327\001\n\027com.city.comm.output.v1B\022OutputServ"
    "iceProtoP\001Z9git.fiblab.net/sim/protos/go"
    "/city/comm/output/v1;outputv1\242\002\003CCO\252\002\023Ci"
    "ty.Comm.Output.V1\312\002\023City\\Comm\\Output\\V1\342"
    "\002\037City\\Comm\\Output\\V1\\GPBMetadata\352\002\026City"
    "::Comm::Output::V1b\006proto3"
};
static const ::_pbi::DescriptorTable* const descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto_deps[2] =
    {
        &::descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_2eproto,
        &::descriptor_table_city_2fevent_2fv1_2fevent_2eproto,
};
static ::absl::once_flag descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto = {
    false,
    false,
    986,
    descriptor_table_protodef_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto,
    "city/comm/output/v1/output_service.proto",
    &descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto_once,
    descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto_deps,
    2,
    2,
    schemas,
    file_default_instances,
    TableStruct_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto::offsets,
    file_level_metadata_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto,
    file_level_enum_descriptors_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto,
    file_level_service_descriptors_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto,
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
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto_getter() {
  return &descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto(&descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto);
namespace city {
namespace comm {
namespace output {
namespace v1 {
// ===================================================================

class OutputRequest::_Internal {
 public:
  using HasBits = decltype(std::declval<OutputRequest>()._impl_._has_bits_);
  static constexpr ::int32_t kHasBitsOffset =
    8 * PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_._has_bits_);
  static const ::city::comm::output::v1::Signal& signal_heatmap(const OutputRequest* msg);
  static void set_has_signal_heatmap(HasBits* has_bits) {
    (*has_bits)[0] |= 1u;
  }
  static const ::city::comm::output::v1::Signal& small_signal_heatmap(const OutputRequest* msg);
  static void set_has_small_signal_heatmap(HasBits* has_bits) {
    (*has_bits)[0] |= 8u;
  }
  static const ::city::event::v1::Events& events(const OutputRequest* msg);
  static void set_has_events(HasBits* has_bits) {
    (*has_bits)[0] |= 2u;
  }
  static const ::city::comm::output::v1::Statistics& statistics(const OutputRequest* msg);
  static void set_has_statistics(HasBits* has_bits) {
    (*has_bits)[0] |= 4u;
  }
};

const ::city::comm::output::v1::Signal& OutputRequest::_Internal::signal_heatmap(const OutputRequest* msg) {
  return *msg->_impl_.signal_heatmap_;
}
const ::city::comm::output::v1::Signal& OutputRequest::_Internal::small_signal_heatmap(const OutputRequest* msg) {
  return *msg->_impl_.small_signal_heatmap_;
}
const ::city::event::v1::Events& OutputRequest::_Internal::events(const OutputRequest* msg) {
  return *msg->_impl_.events_;
}
const ::city::comm::output::v1::Statistics& OutputRequest::_Internal::statistics(const OutputRequest* msg) {
  return *msg->_impl_.statistics_;
}
void OutputRequest::clear_nodes() {
  _internal_mutable_nodes()->Clear();
}
void OutputRequest::clear_base_stations() {
  _internal_mutable_base_stations()->Clear();
}
void OutputRequest::clear_signal_heatmap() {
  if (_impl_.signal_heatmap_ != nullptr) _impl_.signal_heatmap_->Clear();
  _impl_._has_bits_[0] &= ~0x00000001u;
}
void OutputRequest::clear_small_signal_heatmap() {
  if (_impl_.small_signal_heatmap_ != nullptr) _impl_.small_signal_heatmap_->Clear();
  _impl_._has_bits_[0] &= ~0x00000008u;
}
void OutputRequest::clear_persons() {
  _internal_mutable_persons()->Clear();
}
void OutputRequest::clear_aois() {
  _internal_mutable_aois()->Clear();
}
void OutputRequest::clear_events() {
  if (_impl_.events_ != nullptr) _impl_.events_->Clear();
  _impl_._has_bits_[0] &= ~0x00000002u;
}
void OutputRequest::clear_statistics() {
  if (_impl_.statistics_ != nullptr) _impl_.statistics_->Clear();
  _impl_._has_bits_[0] &= ~0x00000004u;
}
OutputRequest::OutputRequest(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:city.comm.output.v1.OutputRequest)
}
OutputRequest::OutputRequest(const OutputRequest& from) : ::google::protobuf::Message() {
  OutputRequest* const _this = this;
  (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){from._impl_._has_bits_},
      /*decltype(_impl_._cached_size_)*/ {},
      decltype(_impl_.nodes_){from._impl_.nodes_},
      decltype(_impl_.base_stations_){from._impl_.base_stations_},
      decltype(_impl_.persons_){from._impl_.persons_},
      decltype(_impl_.aois_){from._impl_.aois_},
      decltype(_impl_.signal_heatmap_){nullptr},
      decltype(_impl_.events_){nullptr},
      decltype(_impl_.statistics_){nullptr},
      decltype(_impl_.small_signal_heatmap_){nullptr},
      decltype(_impl_.step_){},
  };
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);
  if ((from._impl_._has_bits_[0] & 0x00000001u) != 0) {
    _this->_impl_.signal_heatmap_ = new ::city::comm::output::v1::Signal(*from._impl_.signal_heatmap_);
  }
  if ((from._impl_._has_bits_[0] & 0x00000002u) != 0) {
    _this->_impl_.events_ = new ::city::event::v1::Events(*from._impl_.events_);
  }
  if ((from._impl_._has_bits_[0] & 0x00000004u) != 0) {
    _this->_impl_.statistics_ = new ::city::comm::output::v1::Statistics(*from._impl_.statistics_);
  }
  if ((from._impl_._has_bits_[0] & 0x00000008u) != 0) {
    _this->_impl_.small_signal_heatmap_ = new ::city::comm::output::v1::Signal(*from._impl_.small_signal_heatmap_);
  }
  _this->_impl_.step_ = from._impl_.step_;

  // @@protoc_insertion_point(copy_constructor:city.comm.output.v1.OutputRequest)
}
inline void OutputRequest::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_._has_bits_){},
      /*decltype(_impl_._cached_size_)*/ {},
      decltype(_impl_.nodes_){arena},
      decltype(_impl_.base_stations_){arena},
      decltype(_impl_.persons_){arena},
      decltype(_impl_.aois_){arena},
      decltype(_impl_.signal_heatmap_){nullptr},
      decltype(_impl_.events_){nullptr},
      decltype(_impl_.statistics_){nullptr},
      decltype(_impl_.small_signal_heatmap_){nullptr},
      decltype(_impl_.step_){0},
  };
}
OutputRequest::~OutputRequest() {
  // @@protoc_insertion_point(destructor:city.comm.output.v1.OutputRequest)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void OutputRequest::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.nodes_.~RepeatedPtrField();
  _impl_.base_stations_.~RepeatedPtrField();
  _impl_.persons_.~RepeatedPtrField();
  _impl_.aois_.~RepeatedPtrField();
  if (this != internal_default_instance()) delete _impl_.signal_heatmap_;
  if (this != internal_default_instance()) delete _impl_.events_;
  if (this != internal_default_instance()) delete _impl_.statistics_;
  if (this != internal_default_instance()) delete _impl_.small_signal_heatmap_;
}
void OutputRequest::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

PROTOBUF_NOINLINE void OutputRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:city.comm.output.v1.OutputRequest)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _internal_mutable_nodes()->Clear();
  _internal_mutable_base_stations()->Clear();
  _internal_mutable_persons()->Clear();
  _internal_mutable_aois()->Clear();
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x0000000fu) {
    if (cached_has_bits & 0x00000001u) {
      ABSL_DCHECK(_impl_.signal_heatmap_ != nullptr);
      _impl_.signal_heatmap_->Clear();
    }
    if (cached_has_bits & 0x00000002u) {
      ABSL_DCHECK(_impl_.events_ != nullptr);
      _impl_.events_->Clear();
    }
    if (cached_has_bits & 0x00000004u) {
      ABSL_DCHECK(_impl_.statistics_ != nullptr);
      _impl_.statistics_->Clear();
    }
    if (cached_has_bits & 0x00000008u) {
      ABSL_DCHECK(_impl_.small_signal_heatmap_ != nullptr);
      _impl_.small_signal_heatmap_->Clear();
    }
  }
  _impl_.step_ = 0;
  _impl_._has_bits_.Clear();
  _internal_metadata_.Clear<::google::protobuf::UnknownFieldSet>();
}

const char* OutputRequest::_InternalParse(
    const char* ptr, ::_pbi::ParseContext* ctx) {
  ptr = ::_pbi::TcParser::ParseLoop(this, ptr, ctx, &_table_.header);
  return ptr;
}


PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1
const ::_pbi::TcParseTable<4, 9, 8, 0, 2> OutputRequest::_table_ = {
  {
    PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_._has_bits_),
    0, // no _extensions_
    9, 120,  // max_field_number, fast_idx_mask
    offsetof(decltype(_table_), field_lookup_table),
    4294966784,  // skipmap
    offsetof(decltype(_table_), field_entries),
    9,  // num_field_entries
    8,  // num_aux_entries
    offsetof(decltype(_table_), aux_entries),
    &_OutputRequest_default_instance_._instance,
    ::_pbi::TcParser::GenericFallback,  // fallback
  }, {{
    {::_pbi::TcParser::MiniParse, {}},
    // int32 step = 1 [json_name = "step"];
    {::_pbi::TcParser::SingularVarintNoZag1<::uint32_t, offsetof(OutputRequest, _impl_.step_), 63>(),
     {8, 63, 0, PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.step_)}},
    // repeated .city.comm.output.v1.Node nodes = 2 [json_name = "nodes"];
    {::_pbi::TcParser::FastMtR1,
     {18, 63, 0, PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.nodes_)}},
    // repeated .city.comm.output.v1.BaseStation base_stations = 3 [json_name = "baseStations"];
    {::_pbi::TcParser::FastMtR1,
     {26, 63, 1, PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.base_stations_)}},
    // .city.comm.output.v1.Signal signal_heatmap = 4 [json_name = "signalHeatmap"];
    {::_pbi::TcParser::FastMtS1,
     {34, 0, 2, PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.signal_heatmap_)}},
    // repeated .city.comm.output.v1.Person persons = 5 [json_name = "persons"];
    {::_pbi::TcParser::FastMtR1,
     {42, 63, 3, PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.persons_)}},
    // repeated .city.comm.output.v1.Aoi aois = 6 [json_name = "aois"];
    {::_pbi::TcParser::FastMtR1,
     {50, 63, 4, PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.aois_)}},
    // .city.event.v1.Events events = 7 [json_name = "events"];
    {::_pbi::TcParser::FastMtS1,
     {58, 1, 5, PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.events_)}},
    // .city.comm.output.v1.Statistics statistics = 8 [json_name = "statistics"];
    {::_pbi::TcParser::FastMtS1,
     {66, 2, 6, PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.statistics_)}},
    // .city.comm.output.v1.Signal small_signal_heatmap = 9 [json_name = "smallSignalHeatmap"];
    {::_pbi::TcParser::FastMtS1,
     {74, 3, 7, PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.small_signal_heatmap_)}},
    {::_pbi::TcParser::MiniParse, {}},
    {::_pbi::TcParser::MiniParse, {}},
    {::_pbi::TcParser::MiniParse, {}},
    {::_pbi::TcParser::MiniParse, {}},
    {::_pbi::TcParser::MiniParse, {}},
    {::_pbi::TcParser::MiniParse, {}},
  }}, {{
    65535, 65535
  }}, {{
    // int32 step = 1 [json_name = "step"];
    {PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.step_), -1, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kInt32)},
    // repeated .city.comm.output.v1.Node nodes = 2 [json_name = "nodes"];
    {PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.nodes_), -1, 0,
    (0 | ::_fl::kFcRepeated | ::_fl::kMessage | ::_fl::kTvTable)},
    // repeated .city.comm.output.v1.BaseStation base_stations = 3 [json_name = "baseStations"];
    {PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.base_stations_), -1, 1,
    (0 | ::_fl::kFcRepeated | ::_fl::kMessage | ::_fl::kTvTable)},
    // .city.comm.output.v1.Signal signal_heatmap = 4 [json_name = "signalHeatmap"];
    {PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.signal_heatmap_), _Internal::kHasBitsOffset + 0, 2,
    (0 | ::_fl::kFcOptional | ::_fl::kMessage | ::_fl::kTvTable)},
    // repeated .city.comm.output.v1.Person persons = 5 [json_name = "persons"];
    {PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.persons_), -1, 3,
    (0 | ::_fl::kFcRepeated | ::_fl::kMessage | ::_fl::kTvTable)},
    // repeated .city.comm.output.v1.Aoi aois = 6 [json_name = "aois"];
    {PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.aois_), -1, 4,
    (0 | ::_fl::kFcRepeated | ::_fl::kMessage | ::_fl::kTvTable)},
    // .city.event.v1.Events events = 7 [json_name = "events"];
    {PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.events_), _Internal::kHasBitsOffset + 1, 5,
    (0 | ::_fl::kFcOptional | ::_fl::kMessage | ::_fl::kTvTable)},
    // .city.comm.output.v1.Statistics statistics = 8 [json_name = "statistics"];
    {PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.statistics_), _Internal::kHasBitsOffset + 2, 6,
    (0 | ::_fl::kFcOptional | ::_fl::kMessage | ::_fl::kTvTable)},
    // .city.comm.output.v1.Signal small_signal_heatmap = 9 [json_name = "smallSignalHeatmap"];
    {PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.small_signal_heatmap_), _Internal::kHasBitsOffset + 3, 7,
    (0 | ::_fl::kFcOptional | ::_fl::kMessage | ::_fl::kTvTable)},
  }}, {{
    {::_pbi::TcParser::GetTable<::city::comm::output::v1::Node>()},
    {::_pbi::TcParser::GetTable<::city::comm::output::v1::BaseStation>()},
    {::_pbi::TcParser::GetTable<::city::comm::output::v1::Signal>()},
    {::_pbi::TcParser::GetTable<::city::comm::output::v1::Person>()},
    {::_pbi::TcParser::GetTable<::city::comm::output::v1::Aoi>()},
    {::_pbi::TcParser::GetTable<::city::event::v1::Events>()},
    {::_pbi::TcParser::GetTable<::city::comm::output::v1::Statistics>()},
    {::_pbi::TcParser::GetTable<::city::comm::output::v1::Signal>()},
  }}, {{
  }},
};

::uint8_t* OutputRequest::_InternalSerialize(
    ::uint8_t* target,
    ::google::protobuf::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.comm.output.v1.OutputRequest)
  ::uint32_t cached_has_bits = 0;
  (void)cached_has_bits;

  // int32 step = 1 [json_name = "step"];
  if (this->_internal_step() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::
        WriteInt32ToArrayWithField<1>(
            stream, this->_internal_step(), target);
  }

  // repeated .city.comm.output.v1.Node nodes = 2 [json_name = "nodes"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_nodes_size()); i < n; i++) {
    const auto& repfield = this->_internal_nodes().Get(i);
    target = ::google::protobuf::internal::WireFormatLite::
        InternalWriteMessage(2, repfield, repfield.GetCachedSize(), target, stream);
  }

  // repeated .city.comm.output.v1.BaseStation base_stations = 3 [json_name = "baseStations"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_base_stations_size()); i < n; i++) {
    const auto& repfield = this->_internal_base_stations().Get(i);
    target = ::google::protobuf::internal::WireFormatLite::
        InternalWriteMessage(3, repfield, repfield.GetCachedSize(), target, stream);
  }

  cached_has_bits = _impl_._has_bits_[0];
  // .city.comm.output.v1.Signal signal_heatmap = 4 [json_name = "signalHeatmap"];
  if (cached_has_bits & 0x00000001u) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessage(4, _Internal::signal_heatmap(this),
        _Internal::signal_heatmap(this).GetCachedSize(), target, stream);
  }

  // repeated .city.comm.output.v1.Person persons = 5 [json_name = "persons"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_persons_size()); i < n; i++) {
    const auto& repfield = this->_internal_persons().Get(i);
    target = ::google::protobuf::internal::WireFormatLite::
        InternalWriteMessage(5, repfield, repfield.GetCachedSize(), target, stream);
  }

  // repeated .city.comm.output.v1.Aoi aois = 6 [json_name = "aois"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_aois_size()); i < n; i++) {
    const auto& repfield = this->_internal_aois().Get(i);
    target = ::google::protobuf::internal::WireFormatLite::
        InternalWriteMessage(6, repfield, repfield.GetCachedSize(), target, stream);
  }

  // .city.event.v1.Events events = 7 [json_name = "events"];
  if (cached_has_bits & 0x00000002u) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessage(7, _Internal::events(this),
        _Internal::events(this).GetCachedSize(), target, stream);
  }

  // .city.comm.output.v1.Statistics statistics = 8 [json_name = "statistics"];
  if (cached_has_bits & 0x00000004u) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessage(8, _Internal::statistics(this),
        _Internal::statistics(this).GetCachedSize(), target, stream);
  }

  // .city.comm.output.v1.Signal small_signal_heatmap = 9 [json_name = "smallSignalHeatmap"];
  if (cached_has_bits & 0x00000008u) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessage(9, _Internal::small_signal_heatmap(this),
        _Internal::small_signal_heatmap(this).GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target =
        ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
            _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.comm.output.v1.OutputRequest)
  return target;
}

::size_t OutputRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.comm.output.v1.OutputRequest)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.comm.output.v1.Node nodes = 2 [json_name = "nodes"];
  total_size += 1UL * this->_internal_nodes_size();
  for (const auto& msg : this->_internal_nodes()) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSize(msg);
  }
  // repeated .city.comm.output.v1.BaseStation base_stations = 3 [json_name = "baseStations"];
  total_size += 1UL * this->_internal_base_stations_size();
  for (const auto& msg : this->_internal_base_stations()) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSize(msg);
  }
  // repeated .city.comm.output.v1.Person persons = 5 [json_name = "persons"];
  total_size += 1UL * this->_internal_persons_size();
  for (const auto& msg : this->_internal_persons()) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSize(msg);
  }
  // repeated .city.comm.output.v1.Aoi aois = 6 [json_name = "aois"];
  total_size += 1UL * this->_internal_aois_size();
  for (const auto& msg : this->_internal_aois()) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSize(msg);
  }
  cached_has_bits = _impl_._has_bits_[0];
  if (cached_has_bits & 0x0000000fu) {
    // .city.comm.output.v1.Signal signal_heatmap = 4 [json_name = "signalHeatmap"];
    if (cached_has_bits & 0x00000001u) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          *_impl_.signal_heatmap_);
    }

    // .city.event.v1.Events events = 7 [json_name = "events"];
    if (cached_has_bits & 0x00000002u) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          *_impl_.events_);
    }

    // .city.comm.output.v1.Statistics statistics = 8 [json_name = "statistics"];
    if (cached_has_bits & 0x00000004u) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          *_impl_.statistics_);
    }

    // .city.comm.output.v1.Signal small_signal_heatmap = 9 [json_name = "smallSignalHeatmap"];
    if (cached_has_bits & 0x00000008u) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          *_impl_.small_signal_heatmap_);
    }

  }
  // int32 step = 1 [json_name = "step"];
  if (this->_internal_step() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(
        this->_internal_step());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::google::protobuf::Message::ClassData OutputRequest::_class_data_ = {
    ::google::protobuf::Message::CopyWithSourceCheck,
    OutputRequest::MergeImpl
};
const ::google::protobuf::Message::ClassData*OutputRequest::GetClassData() const { return &_class_data_; }


void OutputRequest::MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg) {
  auto* const _this = static_cast<OutputRequest*>(&to_msg);
  auto& from = static_cast<const OutputRequest&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.comm.output.v1.OutputRequest)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_internal_mutable_nodes()->MergeFrom(from._internal_nodes());
  _this->_internal_mutable_base_stations()->MergeFrom(from._internal_base_stations());
  _this->_internal_mutable_persons()->MergeFrom(from._internal_persons());
  _this->_internal_mutable_aois()->MergeFrom(from._internal_aois());
  cached_has_bits = from._impl_._has_bits_[0];
  if (cached_has_bits & 0x0000000fu) {
    if (cached_has_bits & 0x00000001u) {
      _this->_internal_mutable_signal_heatmap()->::city::comm::output::v1::Signal::MergeFrom(
          from._internal_signal_heatmap());
    }
    if (cached_has_bits & 0x00000002u) {
      _this->_internal_mutable_events()->::city::event::v1::Events::MergeFrom(
          from._internal_events());
    }
    if (cached_has_bits & 0x00000004u) {
      _this->_internal_mutable_statistics()->::city::comm::output::v1::Statistics::MergeFrom(
          from._internal_statistics());
    }
    if (cached_has_bits & 0x00000008u) {
      _this->_internal_mutable_small_signal_heatmap()->::city::comm::output::v1::Signal::MergeFrom(
          from._internal_small_signal_heatmap());
    }
  }
  if (from._internal_step() != 0) {
    _this->_internal_set_step(from._internal_step());
  }
  _this->_internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(from._internal_metadata_);
}

void OutputRequest::CopyFrom(const OutputRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.comm.output.v1.OutputRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

PROTOBUF_NOINLINE bool OutputRequest::IsInitialized() const {
  return true;
}

void OutputRequest::InternalSwap(OutputRequest* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_._has_bits_[0], other->_impl_._has_bits_[0]);
  _impl_.nodes_.InternalSwap(&other->_impl_.nodes_);
  _impl_.base_stations_.InternalSwap(&other->_impl_.base_stations_);
  _impl_.persons_.InternalSwap(&other->_impl_.persons_);
  _impl_.aois_.InternalSwap(&other->_impl_.aois_);
  ::google::protobuf::internal::memswap<
      PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.step_)
      + sizeof(OutputRequest::_impl_.step_)
      - PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.signal_heatmap_)>(
          reinterpret_cast<char*>(&_impl_.signal_heatmap_),
          reinterpret_cast<char*>(&other->_impl_.signal_heatmap_));
}

::google::protobuf::Metadata OutputRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto_getter, &descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto_once,
      file_level_metadata_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto[0]);
}
// ===================================================================

class OutputResponse::_Internal {
 public:
};

OutputResponse::OutputResponse(::google::protobuf::Arena* arena)
    : ::google::protobuf::internal::ZeroFieldsBase(arena) {
  // @@protoc_insertion_point(arena_constructor:city.comm.output.v1.OutputResponse)
}
OutputResponse::OutputResponse(const OutputResponse& from) : ::google::protobuf::internal::ZeroFieldsBase() {
  OutputResponse* const _this = this;
  (void)_this;
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);

  // @@protoc_insertion_point(copy_constructor:city.comm.output.v1.OutputResponse)
}




const ::google::protobuf::Message::ClassData OutputResponse::_class_data_ = {
    ::google::protobuf::internal::ZeroFieldsBase::CopyImpl,
    ::google::protobuf::internal::ZeroFieldsBase::MergeImpl,
};
const ::google::protobuf::Message::ClassData*OutputResponse::GetClassData() const { return &_class_data_; }







::google::protobuf::Metadata OutputResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto_getter, &descriptor_table_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto_once,
      file_level_metadata_city_2fcomm_2foutput_2fv1_2foutput_5fservice_2eproto[1]);
}
// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace output
}  // namespace comm
}  // namespace city
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"
