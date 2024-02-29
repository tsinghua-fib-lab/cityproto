"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17city/map/v2/light.proto\x12\x0bcity.map.v2"T\n\x05Phase\x12\x1a\n\x08duration\x18\x01 \x01(\x01R\x08duration\x12/\n\x06states\x18\x02 \x03(\x0e2\x17.city.map.v2.LightStateR\x06states"A\n\x0eAvailablePhase\x12/\n\x06states\x18\x01 \x03(\x0e2\x17.city.map.v2.LightStateR\x06states"[\n\x0cTrafficLight\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId\x12*\n\x06phases\x18\x02 \x03(\x0b2\x12.city.map.v2.PhaseR\x06phases*m\n\nLightState\x12\x1b\n\x17LIGHT_STATE_UNSPECIFIED\x10\x00\x12\x13\n\x0fLIGHT_STATE_RED\x10\x01\x12\x15\n\x11LIGHT_STATE_GREEN\x10\x02\x12\x16\n\x12LIGHT_STATE_YELLOW\x10\x03B\x9b\x01\n\x0fcom.city.map.v2B\nLightProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.light_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\nLightProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_LIGHTSTATE']._serialized_start = 286
    _globals['_LIGHTSTATE']._serialized_end = 395
    _globals['_PHASE']._serialized_start = 40
    _globals['_PHASE']._serialized_end = 124
    _globals['_AVAILABLEPHASE']._serialized_start = 126
    _globals['_AVAILABLEPHASE']._serialized_end = 191
    _globals['_TRAFFICLIGHT']._serialized_start = 193
    _globals['_TRAFFICLIGHT']._serialized_end = 284