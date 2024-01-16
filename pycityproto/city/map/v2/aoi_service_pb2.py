"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.person.v1 import motion_pb2 as city_dot_person_dot_v1_dot_motion__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dcity/map/v2/aoi_service.proto\x12\x0bcity.map.v2\x1a\x1bcity/person/v1/motion.proto"R\n\x08AoiState\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x126\n\x07persons\x18\x02 \x03(\x0b2\x1c.city.person.v1.PersonMotionR\x07persons"(\n\rGetAoiRequest\x12\x17\n\x07aoi_ids\x18\x01 \x03(\x05R\x06aoiIds"?\n\x0eGetAoiResponse\x12-\n\x06states\x18\x01 \x03(\x0b2\x15.city.map.v2.AoiStateR\x06states2O\n\nAoiService\x12A\n\x06GetAoi\x12\x1a.city.map.v2.GetAoiRequest\x1a\x1b.city.map.v2.GetAoiResponseB\xa0\x01\n\x0fcom.city.map.v2B\x0fAoiServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.aoi_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\x0fAoiServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_AOISTATE']._serialized_start = 75
    _globals['_AOISTATE']._serialized_end = 157
    _globals['_GETAOIREQUEST']._serialized_start = 159
    _globals['_GETAOIREQUEST']._serialized_end = 199
    _globals['_GETAOIRESPONSE']._serialized_start = 201
    _globals['_GETAOIRESPONSE']._serialized_end = 264
    _globals['_AOISERVICE']._serialized_start = 266
    _globals['_AOISERVICE']._serialized_end = 345