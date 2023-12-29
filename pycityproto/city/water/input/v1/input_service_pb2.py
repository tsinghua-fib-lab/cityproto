"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.map.v2 import map_pb2 as city_dot_map_dot_v2_dot_map__pb2
from .....city.water.input.v1 import config_pb2 as city_dot_water_dot_input_dot_v1_dot_config__pb2
from .....city.water.input.v1 import water_pb2 as city_dot_water_dot_input_dot_v1_dot_water__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'city/water/input/v1/input_service.proto\x12\x13city.water.input.v1\x1a\x15city/map/v2/map.proto\x1a city/water/input/v1/config.proto\x1a\x1fcity/water/input/v1/water.proto"\r\n\x0bInitRequest"\xb3\x01\n\x0cInitResponse\x12\x18\n\x07address\x18\x02 \x01(\tR\x07address\x126\n\x07control\x18\x03 \x01(\x0b2\x1c.city.water.input.v1.ControlR\x07control\x12-\n\x04rain\x18\x01 \x01(\x0b2\x19.city.water.input.v1.RainR\x04rain\x12"\n\x03map\x18\x04 \x01(\x0b2\x10.city.map.v2.MapR\x03map2]\n\x0cInputService\x12M\n\x04Init\x12 .city.water.input.v1.InitRequest\x1a!.city.water.input.v1.InitResponse"\x00B\xd5\x01\n\x17com.city.water.input.v1B\x11InputServiceProtoP\x01Z8git.fiblab.net/sim/protos/go/city/water/input/v1;inputv1\xa2\x02\x03CWI\xaa\x02\x13City.Water.Input.V1\xca\x02\x13City\\Water\\Input\\V1\xe2\x02\x1fCity\\Water\\Input\\V1\\GPBMetadata\xea\x02\x16City::Water::Input::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.water.input.v1.input_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x17com.city.water.input.v1B\x11InputServiceProtoP\x01Z8git.fiblab.net/sim/protos/go/city/water/input/v1;inputv1\xa2\x02\x03CWI\xaa\x02\x13City.Water.Input.V1\xca\x02\x13City\\Water\\Input\\V1\xe2\x02\x1fCity\\Water\\Input\\V1\\GPBMetadata\xea\x02\x16City::Water::Input::V1'
    _globals['_INITREQUEST']._serialized_start = 154
    _globals['_INITREQUEST']._serialized_end = 167
    _globals['_INITRESPONSE']._serialized_start = 170
    _globals['_INITRESPONSE']._serialized_end = 349
    _globals['_INPUTSERVICE']._serialized_start = 351
    _globals['_INPUTSERVICE']._serialized_end = 444