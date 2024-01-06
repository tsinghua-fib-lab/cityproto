"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.agent.v2 import motion_pb2 as city_dot_agent_dot_v2_dot_motion__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dcity/map/v2/aoi_service.proto\x12\x0bcity.map.v2\x1a\x1acity/agent/v2/motion.proto"P\n\x08AoiState\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x124\n\x07persons\x18\x02 \x03(\x0b2\x1a.city.agent.v2.AgentMotionR\x07persons"(\n\rGetAoiRequest\x12\x17\n\x07aoi_ids\x18\x01 \x03(\x05R\x06aoiIds"?\n\x0eGetAoiResponse\x12-\n\x06states\x18\x01 \x03(\x0b2\x15.city.map.v2.AoiStateR\x06states2O\n\nAoiService\x12A\n\x06GetAoi\x12\x1a.city.map.v2.GetAoiRequest\x1a\x1b.city.map.v2.GetAoiResponseB\xa0\x01\n\x0fcom.city.map.v2B\x0fAoiServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.aoi_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\x0fAoiServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_AOISTATE']._serialized_start = 74
    _globals['_AOISTATE']._serialized_end = 154
    _globals['_GETAOIREQUEST']._serialized_start = 156
    _globals['_GETAOIREQUEST']._serialized_end = 196
    _globals['_GETAOIRESPONSE']._serialized_start = 198
    _globals['_GETAOIRESPONSE']._serialized_end = 261
    _globals['_AOISERVICE']._serialized_start = 263
    _globals['_AOISERVICE']._serialized_end = 342