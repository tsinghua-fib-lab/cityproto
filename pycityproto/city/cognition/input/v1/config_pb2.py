"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.config.v1 import config_pb2 as city_dot_config_dot_v1_dot_config__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$city/cognition/input/v1/config.proto\x12\x17city.cognition.input.v1\x1a\x1bcity/config/v1/config.proto"\xaa\x01\n\x05Mongo\x12\x10\n\x03uri\x18\x01 \x01(\tR\x03uri\x12+\n\x03map\x18\x02 \x01(\x0b2\x19.city.config.v1.MongoPathR\x03map\x121\n\x06agents\x18\x03 \x01(\x0b2\x19.city.config.v1.MongoPathR\x06agents\x12/\n\x05edges\x18\x04 \x01(\x0b2\x19.city.config.v1.MongoPathR\x05edges"U\n\x0bControlStep\x12\x14\n\x05start\x18\x01 \x01(\x05R\x05start\x12\x14\n\x05total\x18\x02 \x01(\x05R\x05total\x12\x1a\n\x08interval\x18\x03 \x01(\x02R\x08interval"C\n\x07Control\x128\n\x04step\x18\x01 \x01(\x0b2$.city.cognition.input.v1.ControlStepR\x04step"&\n\x0cOutputSwitch\x12\x16\n\x06common\x18\x01 \x01(\x08R\x06common"}\n\x06Output\x124\n\x06target\x18\x01 \x01(\x0b2\x1c.city.config.v1.OutputTargetR\x06target\x12=\n\x06switch\x18\x02 \x01(\x0b2%.city.cognition.input.v1.OutputSwitchR\x06switch"\xb3\x01\n\x06Config\x124\n\x05mongo\x18\x01 \x01(\x0b2\x1e.city.cognition.input.v1.MongoR\x05mongo\x12:\n\x07control\x18\x02 \x01(\x0b2 .city.cognition.input.v1.ControlR\x07control\x127\n\x06output\x18\x03 \x01(\x0b2\x1f.city.cognition.input.v1.OutputR\x06outputB\xe7\x01\n\x1bcom.city.cognition.input.v1B\x0bConfigProtoP\x01Z<git.fiblab.net/sim/protos/go/city/cognition/input/v1;inputv1\xa2\x02\x03CCI\xaa\x02\x17City.Cognition.Input.V1\xca\x02\x17City\\Cognition\\Input\\V1\xe2\x02#City\\Cognition\\Input\\V1\\GPBMetadata\xea\x02\x1aCity::Cognition::Input::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.cognition.input.v1.config_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x1bcom.city.cognition.input.v1B\x0bConfigProtoP\x01Z<git.fiblab.net/sim/protos/go/city/cognition/input/v1;inputv1\xa2\x02\x03CCI\xaa\x02\x17City.Cognition.Input.V1\xca\x02\x17City\\Cognition\\Input\\V1\xe2\x02#City\\Cognition\\Input\\V1\\GPBMetadata\xea\x02\x1aCity::Cognition::Input::V1'
    _globals['_MONGO']._serialized_start = 95
    _globals['_MONGO']._serialized_end = 265
    _globals['_CONTROLSTEP']._serialized_start = 267
    _globals['_CONTROLSTEP']._serialized_end = 352
    _globals['_CONTROL']._serialized_start = 354
    _globals['_CONTROL']._serialized_end = 421
    _globals['_OUTPUTSWITCH']._serialized_start = 423
    _globals['_OUTPUTSWITCH']._serialized_end = 461
    _globals['_OUTPUT']._serialized_start = 463
    _globals['_OUTPUT']._serialized_end = 588
    _globals['_CONFIG']._serialized_start = 591
    _globals['_CONFIG']._serialized_end = 770