"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.elec.input.v1 import config_pb2 as city_dot_elec_dot_input_dot_v1_dot_config__pb2
from .....city.elec.input.v1 import input_pb2 as city_dot_elec_dot_input_dot_v1_dot_input__pb2
from .....city.map.v2 import map_pb2 as city_dot_map_dot_v2_dot_map__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&city/elec/input/v1/input_service.proto\x12\x12city.elec.input.v1\x1a\x1fcity/elec/input/v1/config.proto\x1a\x1ecity/elec/input/v1/input.proto\x1a\x15city/map/v2/map.proto"\r\n\x0bInitRequest"\xc3\x01\n\x0cInitResponse\x12\x18\n\x07address\x18\x02 \x01(\tR\x07address\x125\n\x07control\x18\x03 \x01(\x0b2\x1b.city.elec.input.v1.ControlR\x07control\x12>\n\nfacilities\x18\x01 \x01(\x0b2\x1e.city.elec.input.v1.FacilitiesR\nfacilities\x12"\n\x03map\x18\x04 \x01(\x0b2\x10.city.map.v2.MapR\x03map2[\n\x0cInputService\x12K\n\x04Init\x12\x1f.city.elec.input.v1.InitRequest\x1a .city.elec.input.v1.InitResponse"\x00B\xcf\x01\n\x16com.city.elec.input.v1B\x11InputServiceProtoP\x01Z7git.fiblab.net/sim/protos/go/city/elec/input/v1;inputv1\xa2\x02\x03CEI\xaa\x02\x12City.Elec.Input.V1\xca\x02\x12City\\Elec\\Input\\V1\xe2\x02\x1eCity\\Elec\\Input\\V1\\GPBMetadata\xea\x02\x15City::Elec::Input::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.elec.input.v1.input_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x16com.city.elec.input.v1B\x11InputServiceProtoP\x01Z7git.fiblab.net/sim/protos/go/city/elec/input/v1;inputv1\xa2\x02\x03CEI\xaa\x02\x12City.Elec.Input.V1\xca\x02\x12City\\Elec\\Input\\V1\xe2\x02\x1eCity\\Elec\\Input\\V1\\GPBMetadata\xea\x02\x15City::Elec::Input::V1'
    _globals['_INITREQUEST']._serialized_start = 150
    _globals['_INITREQUEST']._serialized_end = 163
    _globals['_INITRESPONSE']._serialized_start = 166
    _globals['_INITRESPONSE']._serialized_end = 361
    _globals['_INPUTSERVICE']._serialized_start = 363
    _globals['_INPUTSERVICE']._serialized_end = 454