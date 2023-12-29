"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.agent.v2 import agent_pb2 as city_dot_agent_dot_v2_dot_agent__pb2
from .....city.cognition.input.v1 import config_pb2 as city_dot_cognition_dot_input_dot_v1_dot_config__pb2
from .....city.cognition.input.v1 import input_pb2 as city_dot_cognition_dot_input_dot_v1_dot_input__pb2
from .....city.map.v2 import map_pb2 as city_dot_map_dot_v2_dot_map__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+city/cognition/input/v1/input_service.proto\x12\x17city.cognition.input.v1\x1a\x19city/agent/v2/agent.proto\x1a$city/cognition/input/v1/config.proto\x1a#city/cognition/input/v1/input.proto\x1a\x15city/map/v2/map.proto"\r\n\x0bInitRequest"\xef\x01\n\x0cInitResponse\x12\x18\n\x07address\x18\x01 \x01(\tR\x07address\x12:\n\x07control\x18\x02 \x01(\x0b2 .city.cognition.input.v1.ControlR\x07control\x12"\n\x03map\x18\x03 \x01(\x0b2\x10.city.map.v2.MapR\x03map\x12/\n\x07persons\x18\x04 \x01(\x0b2\x15.city.agent.v2.AgentsR\x07persons\x124\n\x05edges\x18\x05 \x01(\x0b2\x1e.city.cognition.input.v1.EdgesR\x05edges2c\n\x0cInputService\x12S\n\x04Init\x12$.city.cognition.input.v1.InitRequest\x1a%.city.cognition.input.v1.InitResponseB\xed\x01\n\x1bcom.city.cognition.input.v1B\x11InputServiceProtoP\x01Z<git.fiblab.net/sim/protos/go/city/cognition/input/v1;inputv1\xa2\x02\x03CCI\xaa\x02\x17City.Cognition.Input.V1\xca\x02\x17City\\Cognition\\Input\\V1\xe2\x02#City\\Cognition\\Input\\V1\\GPBMetadata\xea\x02\x1aCity::Cognition::Input::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.cognition.input.v1.input_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x1bcom.city.cognition.input.v1B\x11InputServiceProtoP\x01Z<git.fiblab.net/sim/protos/go/city/cognition/input/v1;inputv1\xa2\x02\x03CCI\xaa\x02\x17City.Cognition.Input.V1\xca\x02\x17City\\Cognition\\Input\\V1\xe2\x02#City\\Cognition\\Input\\V1\\GPBMetadata\xea\x02\x1aCity::Cognition::Input::V1'
    _globals['_INITREQUEST']._serialized_start = 197
    _globals['_INITREQUEST']._serialized_end = 210
    _globals['_INITRESPONSE']._serialized_start = 213
    _globals['_INITRESPONSE']._serialized_end = 452
    _globals['_INPUTSERVICE']._serialized_start = 454
    _globals['_INPUTSERVICE']._serialized_end = 553