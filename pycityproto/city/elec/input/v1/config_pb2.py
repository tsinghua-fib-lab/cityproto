"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.config.v1 import config_pb2 as city_dot_config_dot_v1_dot_config__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fcity/elec/input/v1/config.proto\x12\x12city.elec.input.v1\x1a\x1bcity/config/v1/config.proto"\x81\x01\n\x05Mongo\x12\x10\n\x03uri\x18\x01 \x01(\tR\x03uri\x12+\n\x03map\x18\x02 \x01(\x0b2\x19.city.config.v1.MongoPathR\x03map\x129\n\nfacilities\x18\x03 \x01(\x0b2\x19.city.config.v1.MongoPathR\nfacilities"9\n\x0bControlStep\x12\x14\n\x05start\x18\x01 \x01(\x05R\x05start\x12\x14\n\x05total\x18\x02 \x01(\x05R\x05total">\n\x07Control\x123\n\x04step\x18\x01 \x01(\x0b2\x1f.city.elec.input.v1.ControlStepR\x04step"J\n\x0cOutputSwitch\x12\x12\n\x04node\x18\x01 \x01(\x08R\x04node\x12\x10\n\x03aoi\x18\x02 \x01(\x08R\x03aoi\x12\x14\n\x05event\x18\x03 \x01(\x08R\x05event"x\n\x06Output\x124\n\x06target\x18\x01 \x01(\x0b2\x1c.city.config.v1.OutputTargetR\x06target\x128\n\x06switch\x18\x02 \x01(\x0b2 .city.elec.input.v1.OutputSwitchR\x06switch"\xa4\x01\n\x06Config\x12/\n\x05mongo\x18\x01 \x01(\x0b2\x19.city.elec.input.v1.MongoR\x05mongo\x125\n\x07control\x18\x02 \x01(\x0b2\x1b.city.elec.input.v1.ControlR\x07control\x122\n\x06output\x18\x03 \x01(\x0b2\x1a.city.elec.input.v1.OutputR\x06outputB\xc9\x01\n\x16com.city.elec.input.v1B\x0bConfigProtoP\x01Z7git.fiblab.net/sim/protos/go/city/elec/input/v1;inputv1\xa2\x02\x03CEI\xaa\x02\x12City.Elec.Input.V1\xca\x02\x12City\\Elec\\Input\\V1\xe2\x02\x1eCity\\Elec\\Input\\V1\\GPBMetadata\xea\x02\x15City::Elec::Input::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.elec.input.v1.config_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x16com.city.elec.input.v1B\x0bConfigProtoP\x01Z7git.fiblab.net/sim/protos/go/city/elec/input/v1;inputv1\xa2\x02\x03CEI\xaa\x02\x12City.Elec.Input.V1\xca\x02\x12City\\Elec\\Input\\V1\xe2\x02\x1eCity\\Elec\\Input\\V1\\GPBMetadata\xea\x02\x15City::Elec::Input::V1'
    _globals['_MONGO']._serialized_start = 85
    _globals['_MONGO']._serialized_end = 214
    _globals['_CONTROLSTEP']._serialized_start = 216
    _globals['_CONTROLSTEP']._serialized_end = 273
    _globals['_CONTROL']._serialized_start = 275
    _globals['_CONTROL']._serialized_end = 337
    _globals['_OUTPUTSWITCH']._serialized_start = 339
    _globals['_OUTPUTSWITCH']._serialized_end = 413
    _globals['_OUTPUT']._serialized_start = 415
    _globals['_OUTPUT']._serialized_end = 535
    _globals['_CONFIG']._serialized_start = 538
    _globals['_CONFIG']._serialized_end = 702