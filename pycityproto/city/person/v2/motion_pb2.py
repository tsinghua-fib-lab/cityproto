"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bcity/person/v2/motion.proto\x12\x0ecity.person.v2\x1a\x15city/geo/v2/geo.proto"\xd7\x01\n\x0cPersonMotion\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12.\n\x06status\x18\x02 \x01(\x0e2\x16.city.person.v2.StatusR\x06status\x121\n\x08position\x18\x03 \x01(\x0b2\x15.city.geo.v2.PositionR\x08position\x12\x0c\n\x01v\x18\x04 \x01(\x01R\x01v\x12\x1c\n\tdirection\x18\x05 \x01(\x01R\tdirection\x12\x1a\n\x08activity\x18\x06 \x01(\tR\x08activity\x12\x0c\n\x01l\x18\x07 \x01(\x01R\x01l*\xdd\x01\n\x06Status\x12\x16\n\x12STATUS_UNSPECIFIED\x10\x00\x12\x10\n\x0cSTATUS_SLEEP\x10\x01\x12\x12\n\x0eSTATUS_DRIVING\x10\x02\x12\x12\n\x0eSTATUS_WALKING\x10\x03\x12\x10\n\x0cSTATUS_CROWD\x10\x04\x12\x14\n\x10STATUS_PASSENGER\x10\x05\x12\x15\n\x11STATUS_WAIT_ROUTE\x10\x06\x12\x13\n\x0fSTATUS_WAIT_BUS\x10\x07\x12\x17\n\x13STATUS_RAIL_TRANSIT\x10\x08\x12\x14\n\x10STATUS_WAIT_TAXI\x10\tB\xb4\x01\n\x12com.city.person.v2B\x0bMotionProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v2.motion_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v2B\x0bMotionProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2'
    _globals['_STATUS']._serialized_start = 289
    _globals['_STATUS']._serialized_end = 510
    _globals['_PERSONMOTION']._serialized_start = 71
    _globals['_PERSONMOTION']._serialized_end = 286