"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n4city/comm/interaction/demand/v1/demand_service.proto\x12\x1fcity.comm.interaction.demand.v1"`\n\x16SetDemandStatusRequest\x12%\n\x0emultiple_times\x18\x01 \x01(\x01R\rmultipleTimes\x12\x1f\n\x0bpower_times\x18\x02 \x01(\x01R\npowerTimes"\x19\n\x17SetDemandStatusResponse2\x96\x01\n\rDemandService\x12\x84\x01\n\x0fSetDemandStatus\x127.city.comm.interaction.demand.v1.SetDemandStatusRequest\x1a8.city.comm.interaction.demand.v1.SetDemandStatusResponseB\xa1\x02\n#com.city.comm.interaction.demand.v1B\x12DemandServiceProtoP\x01ZEgit.fiblab.net/sim/protos/go/city/comm/interaction/demand/v1;demandv1\xa2\x02\x04CCID\xaa\x02\x1fCity.Comm.Interaction.Demand.V1\xca\x02\x1fCity\\Comm\\Interaction\\Demand\\V1\xe2\x02+City\\Comm\\Interaction\\Demand\\V1\\GPBMetadata\xea\x02#City::Comm::Interaction::Demand::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.comm.interaction.demand.v1.demand_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n#com.city.comm.interaction.demand.v1B\x12DemandServiceProtoP\x01ZEgit.fiblab.net/sim/protos/go/city/comm/interaction/demand/v1;demandv1\xa2\x02\x04CCID\xaa\x02\x1fCity.Comm.Interaction.Demand.V1\xca\x02\x1fCity\\Comm\\Interaction\\Demand\\V1\xe2\x02+City\\Comm\\Interaction\\Demand\\V1\\GPBMetadata\xea\x02#City::Comm::Interaction::Demand::V1'
    _globals['_SETDEMANDSTATUSREQUEST']._serialized_start = 89
    _globals['_SETDEMANDSTATUSREQUEST']._serialized_end = 185
    _globals['_SETDEMANDSTATUSRESPONSE']._serialized_start = 187
    _globals['_SETDEMANDSTATUSRESPONSE']._serialized_end = 212
    _globals['_DEMANDSERVICE']._serialized_start = 215
    _globals['_DEMANDSERVICE']._serialized_end = 365