"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.config.v1 import config_pb2 as city_dot_config_dot_v1_dot_config__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n city/water/input/v1/config.proto\x12\x13city.water.input.v1\x1a\x1bcity/config/v1/config.proto"u\n\x05Mongo\x12\x10\n\x03uri\x18\x01 \x01(\tR\x03uri\x12+\n\x03map\x18\x02 \x01(\x0b2\x19.city.config.v1.MongoPathR\x03map\x12-\n\x04rain\x18\x03 \x01(\x0b2\x19.city.config.v1.MongoPathR\x04rain"U\n\x0bControlStep\x12\x14\n\x05start\x18\x01 \x01(\x05R\x05start\x12\x14\n\x05total\x18\x02 \x01(\x05R\x05total\x12\x1a\n\x08interval\x18\x03 \x01(\x01R\x08interval"?\n\x07Control\x124\n\x04step\x18\x01 \x01(\x0b2 .city.water.input.v1.ControlStepR\x04step"~\n\x0cOutputSwitch\x12\x12\n\x04road\x18\x01 \x01(\x08R\x04road\x12\x1a\n\x08drainage\x18\x02 \x01(\x08R\x08drainage\x12\x16\n\x06supply\x18\x03 \x01(\x08R\x06supply\x12\x10\n\x03aoi\x18\x04 \x01(\x08R\x03aoi\x12\x14\n\x05event\x18\x05 \x01(\x08R\x05event"y\n\x06Output\x124\n\x06target\x18\x01 \x01(\x0b2\x1c.city.config.v1.OutputTargetR\x06target\x129\n\x06switch\x18\x02 \x01(\x0b2!.city.water.input.v1.OutputSwitchR\x06switch"\xa7\x01\n\x06Config\x120\n\x05mongo\x18\x01 \x01(\x0b2\x1a.city.water.input.v1.MongoR\x05mongo\x126\n\x07control\x18\x02 \x01(\x0b2\x1c.city.water.input.v1.ControlR\x07control\x123\n\x06output\x18\x03 \x01(\x0b2\x1b.city.water.input.v1.OutputR\x06outputB\xcf\x01\n\x17com.city.water.input.v1B\x0bConfigProtoP\x01Z8git.fiblab.net/sim/protos/go/city/water/input/v1;inputv1\xa2\x02\x03CWI\xaa\x02\x13City.Water.Input.V1\xca\x02\x13City\\Water\\Input\\V1\xe2\x02\x1fCity\\Water\\Input\\V1\\GPBMetadata\xea\x02\x16City::Water::Input::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.water.input.v1.config_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x17com.city.water.input.v1B\x0bConfigProtoP\x01Z8git.fiblab.net/sim/protos/go/city/water/input/v1;inputv1\xa2\x02\x03CWI\xaa\x02\x13City.Water.Input.V1\xca\x02\x13City\\Water\\Input\\V1\xe2\x02\x1fCity\\Water\\Input\\V1\\GPBMetadata\xea\x02\x16City::Water::Input::V1'
    _globals['_MONGO']._serialized_start = 86
    _globals['_MONGO']._serialized_end = 203
    _globals['_CONTROLSTEP']._serialized_start = 205
    _globals['_CONTROLSTEP']._serialized_end = 290
    _globals['_CONTROL']._serialized_start = 292
    _globals['_CONTROL']._serialized_end = 355
    _globals['_OUTPUTSWITCH']._serialized_start = 357
    _globals['_OUTPUTSWITCH']._serialized_end = 483
    _globals['_OUTPUT']._serialized_start = 485
    _globals['_OUTPUT']._serialized_end = 606
    _globals['_CONFIG']._serialized_start = 609
    _globals['_CONFIG']._serialized_end = 776