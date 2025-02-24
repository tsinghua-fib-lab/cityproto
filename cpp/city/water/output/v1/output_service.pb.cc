// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: city/water/output/v1/output_service.proto

#include "city/water/output/v1/output_service.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

PROTOBUF_PRAGMA_INIT_SEG

namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = _pb::internal;

namespace city {
namespace water {
namespace output {
namespace v1 {
PROTOBUF_CONSTEXPR OutputRequest::OutputRequest(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.roads_)*/{}
  , /*decltype(_impl_.detailed_roads_)*/{}
  , /*decltype(_impl_.drainage_nodes_)*/{}
  , /*decltype(_impl_.drainage_links_)*/{}
  , /*decltype(_impl_.supply_nodes_)*/{}
  , /*decltype(_impl_.supply_links_)*/{}
  , /*decltype(_impl_.aois_)*/{}
  , /*decltype(_impl_.events_)*/nullptr
  , /*decltype(_impl_.drainage_metrics_)*/nullptr
  , /*decltype(_impl_.supply_metrics_)*/nullptr
  , /*decltype(_impl_.drainage_metric_)*/0
  , /*decltype(_impl_.step_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct OutputRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR OutputRequestDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~OutputRequestDefaultTypeInternal() {}
  union {
    OutputRequest _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 OutputRequestDefaultTypeInternal _OutputRequest_default_instance_;
PROTOBUF_CONSTEXPR OutputResponse::OutputResponse(
    ::_pbi::ConstantInitialized) {}
struct OutputResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR OutputResponseDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~OutputResponseDefaultTypeInternal() {}
  union {
    OutputResponse _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 OutputResponseDefaultTypeInternal _OutputResponse_default_instance_;
}  // namespace v1
}  // namespace output
}  // namespace water
}  // namespace city
static ::_pb::Metadata file_level_metadata_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto[2];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto = nullptr;

const uint32_t TableStruct_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.step_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.roads_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.detailed_roads_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.drainage_nodes_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.drainage_links_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.supply_nodes_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.supply_links_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.aois_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.drainage_metric_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.events_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.drainage_metrics_),
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputRequest, _impl_.supply_metrics_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::city::water::output::v1::OutputResponse, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::city::water::output::v1::OutputRequest)},
  { 18, -1, -1, sizeof(::city::water::output::v1::OutputResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::city::water::output::v1::_OutputRequest_default_instance_._instance,
  &::city::water::output::v1::_OutputResponse_default_instance_._instance,
};

const char descriptor_table_protodef_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n)city/water/output/v1/output_service.pr"
  "oto\022\024city.water.output.v1\032\031city/event/v1"
  "/event.proto\032!city/water/output/v1/outpu"
  "t.proto\"\311\005\n\rOutputRequest\022\022\n\004step\030\001 \001(\005R"
  "\004step\0220\n\005roads\030\002 \003(\0132\032.city.water.output"
  ".v1.RoadR\005roads\022I\n\016detailed_roads\030\003 \003(\0132"
  "\".city.water.output.v1.DetailedRoadR\rdet"
  "ailedRoads\022A\n\016drainage_nodes\030\004 \003(\0132\032.cit"
  "y.water.output.v1.NodeR\rdrainageNodes\022A\n"
  "\016drainage_links\030\005 \003(\0132\032.city.water.outpu"
  "t.v1.LinkR\rdrainageLinks\022=\n\014supply_nodes"
  "\030\006 \003(\0132\032.city.water.output.v1.NodeR\013supp"
  "lyNodes\022=\n\014supply_links\030\007 \003(\0132\032.city.wat"
  "er.output.v1.LinkR\013supplyLinks\022-\n\004aois\030\010"
  " \003(\0132\031.city.water.output.v1.AoiR\004aois\022\'\n"
  "\017drainage_metric\030\t \001(\001R\016drainageMetric\022-"
  "\n\006events\030\n \001(\0132\025.city.event.v1.EventsR\006e"
  "vents\022P\n\020drainage_metrics\030\013 \001(\0132%.city.w"
  "ater.output.v1.DrainageMetricsR\017drainage"
  "Metrics\022J\n\016supply_metrics\030\014 \001(\0132#.city.w"
  "ater.output.v1.SupplyMetricsR\rsupplyMetr"
  "ics\"\020\n\016OutputResponse2d\n\rOutputService\022S"
  "\n\006Output\022#.city.water.output.v1.OutputRe"
  "quest\032$.city.water.output.v1.OutputRespo"
  "nseB\340\001\n\030com.city.water.output.v1B\022Output"
  "ServiceProtoP\001Z=git.fiblab.net/sim/proto"
  "s/v2/go/city/water/output/v1;outputv1\242\002\003"
  "CWO\252\002\024City.Water.Output.V1\312\002\024City\\Water\\"
  "Output\\V1\342\002 City\\Water\\Output\\V1\\GPBMeta"
  "data\352\002\027City::Water::Output::V1b\006proto3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto_deps[2] = {
  &::descriptor_table_city_2fevent_2fv1_2fevent_2eproto,
  &::descriptor_table_city_2fwater_2foutput_2fv1_2foutput_2eproto,
};
static ::_pbi::once_flag descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto = {
    false, false, 1198, descriptor_table_protodef_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto,
    "city/water/output/v1/output_service.proto",
    &descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto_once, descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto_deps, 2, 2,
    schemas, file_default_instances, TableStruct_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto::offsets,
    file_level_metadata_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto, file_level_enum_descriptors_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto,
    file_level_service_descriptors_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto_getter() {
  return &descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto(&descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto);
namespace city {
namespace water {
namespace output {
namespace v1 {

// ===================================================================

class OutputRequest::_Internal {
 public:
  static const ::city::event::v1::Events& events(const OutputRequest* msg);
  static const ::city::water::output::v1::DrainageMetrics& drainage_metrics(const OutputRequest* msg);
  static const ::city::water::output::v1::SupplyMetrics& supply_metrics(const OutputRequest* msg);
};

const ::city::event::v1::Events&
OutputRequest::_Internal::events(const OutputRequest* msg) {
  return *msg->_impl_.events_;
}
const ::city::water::output::v1::DrainageMetrics&
OutputRequest::_Internal::drainage_metrics(const OutputRequest* msg) {
  return *msg->_impl_.drainage_metrics_;
}
const ::city::water::output::v1::SupplyMetrics&
OutputRequest::_Internal::supply_metrics(const OutputRequest* msg) {
  return *msg->_impl_.supply_metrics_;
}
void OutputRequest::clear_roads() {
  _impl_.roads_.Clear();
}
void OutputRequest::clear_detailed_roads() {
  _impl_.detailed_roads_.Clear();
}
void OutputRequest::clear_drainage_nodes() {
  _impl_.drainage_nodes_.Clear();
}
void OutputRequest::clear_drainage_links() {
  _impl_.drainage_links_.Clear();
}
void OutputRequest::clear_supply_nodes() {
  _impl_.supply_nodes_.Clear();
}
void OutputRequest::clear_supply_links() {
  _impl_.supply_links_.Clear();
}
void OutputRequest::clear_aois() {
  _impl_.aois_.Clear();
}
void OutputRequest::clear_events() {
  if (GetArenaForAllocation() == nullptr && _impl_.events_ != nullptr) {
    delete _impl_.events_;
  }
  _impl_.events_ = nullptr;
}
void OutputRequest::clear_drainage_metrics() {
  if (GetArenaForAllocation() == nullptr && _impl_.drainage_metrics_ != nullptr) {
    delete _impl_.drainage_metrics_;
  }
  _impl_.drainage_metrics_ = nullptr;
}
void OutputRequest::clear_supply_metrics() {
  if (GetArenaForAllocation() == nullptr && _impl_.supply_metrics_ != nullptr) {
    delete _impl_.supply_metrics_;
  }
  _impl_.supply_metrics_ = nullptr;
}
OutputRequest::OutputRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:city.water.output.v1.OutputRequest)
}
OutputRequest::OutputRequest(const OutputRequest& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  OutputRequest* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.roads_){from._impl_.roads_}
    , decltype(_impl_.detailed_roads_){from._impl_.detailed_roads_}
    , decltype(_impl_.drainage_nodes_){from._impl_.drainage_nodes_}
    , decltype(_impl_.drainage_links_){from._impl_.drainage_links_}
    , decltype(_impl_.supply_nodes_){from._impl_.supply_nodes_}
    , decltype(_impl_.supply_links_){from._impl_.supply_links_}
    , decltype(_impl_.aois_){from._impl_.aois_}
    , decltype(_impl_.events_){nullptr}
    , decltype(_impl_.drainage_metrics_){nullptr}
    , decltype(_impl_.supply_metrics_){nullptr}
    , decltype(_impl_.drainage_metric_){}
    , decltype(_impl_.step_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  if (from._internal_has_events()) {
    _this->_impl_.events_ = new ::city::event::v1::Events(*from._impl_.events_);
  }
  if (from._internal_has_drainage_metrics()) {
    _this->_impl_.drainage_metrics_ = new ::city::water::output::v1::DrainageMetrics(*from._impl_.drainage_metrics_);
  }
  if (from._internal_has_supply_metrics()) {
    _this->_impl_.supply_metrics_ = new ::city::water::output::v1::SupplyMetrics(*from._impl_.supply_metrics_);
  }
  ::memcpy(&_impl_.drainage_metric_, &from._impl_.drainage_metric_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.step_) -
    reinterpret_cast<char*>(&_impl_.drainage_metric_)) + sizeof(_impl_.step_));
  // @@protoc_insertion_point(copy_constructor:city.water.output.v1.OutputRequest)
}

inline void OutputRequest::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.roads_){arena}
    , decltype(_impl_.detailed_roads_){arena}
    , decltype(_impl_.drainage_nodes_){arena}
    , decltype(_impl_.drainage_links_){arena}
    , decltype(_impl_.supply_nodes_){arena}
    , decltype(_impl_.supply_links_){arena}
    , decltype(_impl_.aois_){arena}
    , decltype(_impl_.events_){nullptr}
    , decltype(_impl_.drainage_metrics_){nullptr}
    , decltype(_impl_.supply_metrics_){nullptr}
    , decltype(_impl_.drainage_metric_){0}
    , decltype(_impl_.step_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

OutputRequest::~OutputRequest() {
  // @@protoc_insertion_point(destructor:city.water.output.v1.OutputRequest)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void OutputRequest::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.roads_.~RepeatedPtrField();
  _impl_.detailed_roads_.~RepeatedPtrField();
  _impl_.drainage_nodes_.~RepeatedPtrField();
  _impl_.drainage_links_.~RepeatedPtrField();
  _impl_.supply_nodes_.~RepeatedPtrField();
  _impl_.supply_links_.~RepeatedPtrField();
  _impl_.aois_.~RepeatedPtrField();
  if (this != internal_default_instance()) delete _impl_.events_;
  if (this != internal_default_instance()) delete _impl_.drainage_metrics_;
  if (this != internal_default_instance()) delete _impl_.supply_metrics_;
}

void OutputRequest::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void OutputRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:city.water.output.v1.OutputRequest)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.roads_.Clear();
  _impl_.detailed_roads_.Clear();
  _impl_.drainage_nodes_.Clear();
  _impl_.drainage_links_.Clear();
  _impl_.supply_nodes_.Clear();
  _impl_.supply_links_.Clear();
  _impl_.aois_.Clear();
  if (GetArenaForAllocation() == nullptr && _impl_.events_ != nullptr) {
    delete _impl_.events_;
  }
  _impl_.events_ = nullptr;
  if (GetArenaForAllocation() == nullptr && _impl_.drainage_metrics_ != nullptr) {
    delete _impl_.drainage_metrics_;
  }
  _impl_.drainage_metrics_ = nullptr;
  if (GetArenaForAllocation() == nullptr && _impl_.supply_metrics_ != nullptr) {
    delete _impl_.supply_metrics_;
  }
  _impl_.supply_metrics_ = nullptr;
  ::memset(&_impl_.drainage_metric_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&_impl_.step_) -
      reinterpret_cast<char*>(&_impl_.drainage_metric_)) + sizeof(_impl_.step_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* OutputRequest::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // int32 step = 1 [json_name = "step"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 8)) {
          _impl_.step_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // repeated .city.water.output.v1.Road roads = 2 [json_name = "roads"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 18)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_roads(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<18>(ptr));
        } else
          goto handle_unusual;
        continue;
      // repeated .city.water.output.v1.DetailedRoad detailed_roads = 3 [json_name = "detailedRoads"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 26)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_detailed_roads(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<26>(ptr));
        } else
          goto handle_unusual;
        continue;
      // repeated .city.water.output.v1.Node drainage_nodes = 4 [json_name = "drainageNodes"];
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 34)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_drainage_nodes(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<34>(ptr));
        } else
          goto handle_unusual;
        continue;
      // repeated .city.water.output.v1.Link drainage_links = 5 [json_name = "drainageLinks"];
      case 5:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 42)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_drainage_links(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<42>(ptr));
        } else
          goto handle_unusual;
        continue;
      // repeated .city.water.output.v1.Node supply_nodes = 6 [json_name = "supplyNodes"];
      case 6:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 50)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_supply_nodes(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<50>(ptr));
        } else
          goto handle_unusual;
        continue;
      // repeated .city.water.output.v1.Link supply_links = 7 [json_name = "supplyLinks"];
      case 7:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 58)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_supply_links(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<58>(ptr));
        } else
          goto handle_unusual;
        continue;
      // repeated .city.water.output.v1.Aoi aois = 8 [json_name = "aois"];
      case 8:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 66)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_aois(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<66>(ptr));
        } else
          goto handle_unusual;
        continue;
      // double drainage_metric = 9 [json_name = "drainageMetric"];
      case 9:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 73)) {
          _impl_.drainage_metric_ = ::PROTOBUF_NAMESPACE_ID::internal::UnalignedLoad<double>(ptr);
          ptr += sizeof(double);
        } else
          goto handle_unusual;
        continue;
      // .city.event.v1.Events events = 10 [json_name = "events"];
      case 10:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 82)) {
          ptr = ctx->ParseMessage(_internal_mutable_events(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // .city.water.output.v1.DrainageMetrics drainage_metrics = 11 [json_name = "drainageMetrics"];
      case 11:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 90)) {
          ptr = ctx->ParseMessage(_internal_mutable_drainage_metrics(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // .city.water.output.v1.SupplyMetrics supply_metrics = 12 [json_name = "supplyMetrics"];
      case 12:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 98)) {
          ptr = ctx->ParseMessage(_internal_mutable_supply_metrics(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

uint8_t* OutputRequest::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:city.water.output.v1.OutputRequest)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 step = 1 [json_name = "step"];
  if (this->_internal_step() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(1, this->_internal_step(), target);
  }

  // repeated .city.water.output.v1.Road roads = 2 [json_name = "roads"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_roads_size()); i < n; i++) {
    const auto& repfield = this->_internal_roads(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(2, repfield, repfield.GetCachedSize(), target, stream);
  }

  // repeated .city.water.output.v1.DetailedRoad detailed_roads = 3 [json_name = "detailedRoads"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_detailed_roads_size()); i < n; i++) {
    const auto& repfield = this->_internal_detailed_roads(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(3, repfield, repfield.GetCachedSize(), target, stream);
  }

  // repeated .city.water.output.v1.Node drainage_nodes = 4 [json_name = "drainageNodes"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_drainage_nodes_size()); i < n; i++) {
    const auto& repfield = this->_internal_drainage_nodes(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(4, repfield, repfield.GetCachedSize(), target, stream);
  }

  // repeated .city.water.output.v1.Link drainage_links = 5 [json_name = "drainageLinks"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_drainage_links_size()); i < n; i++) {
    const auto& repfield = this->_internal_drainage_links(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(5, repfield, repfield.GetCachedSize(), target, stream);
  }

  // repeated .city.water.output.v1.Node supply_nodes = 6 [json_name = "supplyNodes"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_supply_nodes_size()); i < n; i++) {
    const auto& repfield = this->_internal_supply_nodes(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(6, repfield, repfield.GetCachedSize(), target, stream);
  }

  // repeated .city.water.output.v1.Link supply_links = 7 [json_name = "supplyLinks"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_supply_links_size()); i < n; i++) {
    const auto& repfield = this->_internal_supply_links(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(7, repfield, repfield.GetCachedSize(), target, stream);
  }

  // repeated .city.water.output.v1.Aoi aois = 8 [json_name = "aois"];
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_aois_size()); i < n; i++) {
    const auto& repfield = this->_internal_aois(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(8, repfield, repfield.GetCachedSize(), target, stream);
  }

  // double drainage_metric = 9 [json_name = "drainageMetric"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_drainage_metric = this->_internal_drainage_metric();
  uint64_t raw_drainage_metric;
  memcpy(&raw_drainage_metric, &tmp_drainage_metric, sizeof(tmp_drainage_metric));
  if (raw_drainage_metric != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteDoubleToArray(9, this->_internal_drainage_metric(), target);
  }

  // .city.event.v1.Events events = 10 [json_name = "events"];
  if (this->_internal_has_events()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(10, _Internal::events(this),
        _Internal::events(this).GetCachedSize(), target, stream);
  }

  // .city.water.output.v1.DrainageMetrics drainage_metrics = 11 [json_name = "drainageMetrics"];
  if (this->_internal_has_drainage_metrics()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(11, _Internal::drainage_metrics(this),
        _Internal::drainage_metrics(this).GetCachedSize(), target, stream);
  }

  // .city.water.output.v1.SupplyMetrics supply_metrics = 12 [json_name = "supplyMetrics"];
  if (this->_internal_has_supply_metrics()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(12, _Internal::supply_metrics(this),
        _Internal::supply_metrics(this).GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:city.water.output.v1.OutputRequest)
  return target;
}

size_t OutputRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:city.water.output.v1.OutputRequest)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .city.water.output.v1.Road roads = 2 [json_name = "roads"];
  total_size += 1UL * this->_internal_roads_size();
  for (const auto& msg : this->_impl_.roads_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  // repeated .city.water.output.v1.DetailedRoad detailed_roads = 3 [json_name = "detailedRoads"];
  total_size += 1UL * this->_internal_detailed_roads_size();
  for (const auto& msg : this->_impl_.detailed_roads_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  // repeated .city.water.output.v1.Node drainage_nodes = 4 [json_name = "drainageNodes"];
  total_size += 1UL * this->_internal_drainage_nodes_size();
  for (const auto& msg : this->_impl_.drainage_nodes_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  // repeated .city.water.output.v1.Link drainage_links = 5 [json_name = "drainageLinks"];
  total_size += 1UL * this->_internal_drainage_links_size();
  for (const auto& msg : this->_impl_.drainage_links_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  // repeated .city.water.output.v1.Node supply_nodes = 6 [json_name = "supplyNodes"];
  total_size += 1UL * this->_internal_supply_nodes_size();
  for (const auto& msg : this->_impl_.supply_nodes_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  // repeated .city.water.output.v1.Link supply_links = 7 [json_name = "supplyLinks"];
  total_size += 1UL * this->_internal_supply_links_size();
  for (const auto& msg : this->_impl_.supply_links_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  // repeated .city.water.output.v1.Aoi aois = 8 [json_name = "aois"];
  total_size += 1UL * this->_internal_aois_size();
  for (const auto& msg : this->_impl_.aois_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  // .city.event.v1.Events events = 10 [json_name = "events"];
  if (this->_internal_has_events()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.events_);
  }

  // .city.water.output.v1.DrainageMetrics drainage_metrics = 11 [json_name = "drainageMetrics"];
  if (this->_internal_has_drainage_metrics()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.drainage_metrics_);
  }

  // .city.water.output.v1.SupplyMetrics supply_metrics = 12 [json_name = "supplyMetrics"];
  if (this->_internal_has_supply_metrics()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.supply_metrics_);
  }

  // double drainage_metric = 9 [json_name = "drainageMetric"];
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_drainage_metric = this->_internal_drainage_metric();
  uint64_t raw_drainage_metric;
  memcpy(&raw_drainage_metric, &tmp_drainage_metric, sizeof(tmp_drainage_metric));
  if (raw_drainage_metric != 0) {
    total_size += 1 + 8;
  }

  // int32 step = 1 [json_name = "step"];
  if (this->_internal_step() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(this->_internal_step());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData OutputRequest::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    OutputRequest::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*OutputRequest::GetClassData() const { return &_class_data_; }


void OutputRequest::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<OutputRequest*>(&to_msg);
  auto& from = static_cast<const OutputRequest&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:city.water.output.v1.OutputRequest)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.roads_.MergeFrom(from._impl_.roads_);
  _this->_impl_.detailed_roads_.MergeFrom(from._impl_.detailed_roads_);
  _this->_impl_.drainage_nodes_.MergeFrom(from._impl_.drainage_nodes_);
  _this->_impl_.drainage_links_.MergeFrom(from._impl_.drainage_links_);
  _this->_impl_.supply_nodes_.MergeFrom(from._impl_.supply_nodes_);
  _this->_impl_.supply_links_.MergeFrom(from._impl_.supply_links_);
  _this->_impl_.aois_.MergeFrom(from._impl_.aois_);
  if (from._internal_has_events()) {
    _this->_internal_mutable_events()->::city::event::v1::Events::MergeFrom(
        from._internal_events());
  }
  if (from._internal_has_drainage_metrics()) {
    _this->_internal_mutable_drainage_metrics()->::city::water::output::v1::DrainageMetrics::MergeFrom(
        from._internal_drainage_metrics());
  }
  if (from._internal_has_supply_metrics()) {
    _this->_internal_mutable_supply_metrics()->::city::water::output::v1::SupplyMetrics::MergeFrom(
        from._internal_supply_metrics());
  }
  static_assert(sizeof(uint64_t) == sizeof(double), "Code assumes uint64_t and double are the same size.");
  double tmp_drainage_metric = from._internal_drainage_metric();
  uint64_t raw_drainage_metric;
  memcpy(&raw_drainage_metric, &tmp_drainage_metric, sizeof(tmp_drainage_metric));
  if (raw_drainage_metric != 0) {
    _this->_internal_set_drainage_metric(from._internal_drainage_metric());
  }
  if (from._internal_step() != 0) {
    _this->_internal_set_step(from._internal_step());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void OutputRequest::CopyFrom(const OutputRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:city.water.output.v1.OutputRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool OutputRequest::IsInitialized() const {
  return true;
}

void OutputRequest::InternalSwap(OutputRequest* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.roads_.InternalSwap(&other->_impl_.roads_);
  _impl_.detailed_roads_.InternalSwap(&other->_impl_.detailed_roads_);
  _impl_.drainage_nodes_.InternalSwap(&other->_impl_.drainage_nodes_);
  _impl_.drainage_links_.InternalSwap(&other->_impl_.drainage_links_);
  _impl_.supply_nodes_.InternalSwap(&other->_impl_.supply_nodes_);
  _impl_.supply_links_.InternalSwap(&other->_impl_.supply_links_);
  _impl_.aois_.InternalSwap(&other->_impl_.aois_);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.step_)
      + sizeof(OutputRequest::_impl_.step_)
      - PROTOBUF_FIELD_OFFSET(OutputRequest, _impl_.events_)>(
          reinterpret_cast<char*>(&_impl_.events_),
          reinterpret_cast<char*>(&other->_impl_.events_));
}

::PROTOBUF_NAMESPACE_ID::Metadata OutputRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto_getter, &descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto_once,
      file_level_metadata_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto[0]);
}

// ===================================================================

class OutputResponse::_Internal {
 public:
};

OutputResponse::OutputResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase(arena, is_message_owned) {
  // @@protoc_insertion_point(arena_constructor:city.water.output.v1.OutputResponse)
}
OutputResponse::OutputResponse(const OutputResponse& from)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase() {
  OutputResponse* const _this = this; (void)_this;
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:city.water.output.v1.OutputResponse)
}





const ::PROTOBUF_NAMESPACE_ID::Message::ClassData OutputResponse::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyImpl,
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeImpl,
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*OutputResponse::GetClassData() const { return &_class_data_; }







::PROTOBUF_NAMESPACE_ID::Metadata OutputResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto_getter, &descriptor_table_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto_once,
      file_level_metadata_city_2fwater_2foutput_2fv1_2foutput_5fservice_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace output
}  // namespace water
}  // namespace city
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::city::water::output::v1::OutputRequest*
Arena::CreateMaybeMessage< ::city::water::output::v1::OutputRequest >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::water::output::v1::OutputRequest >(arena);
}
template<> PROTOBUF_NOINLINE ::city::water::output::v1::OutputResponse*
Arena::CreateMaybeMessage< ::city::water::output::v1::OutputResponse >(Arena* arena) {
  return Arena::CreateMessageInternal< ::city::water::output::v1::OutputResponse >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
