"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19city/event/v2/event.proto\x12\rcity.event.v2\x1a\x15city/geo/v2/geo.proto"G\n\x06Entity\x12-\n\x04type\x18\x01 \x01(\x0e2\x19.city.event.v2.EntityTypeR\x04type\x12\x0e\n\x02id\x18\x02 \x01(\x05R\x02id"\xc5\x01\n\x05Event\x12\x14\n\x05topic\x18\x01 \x01(\tR\x05topic\x12\x13\n\x02id\x18\x02 \x01(\x05H\x00R\x02id\x88\x01\x01\x12/\n\x07subject\x18\x03 \x01(\x0b2\x15.city.event.v2.EntityR\x07subject\x12\x18\n\x07content\x18\x04 \x01(\tR\x07content\x121\n\x08position\x18\x05 \x01(\x0b2\x15.city.geo.v2.PositionR\x08position\x12\x0c\n\x01t\x18\x06 \x01(\x01R\x01tB\x05\n\x03_id*\xc6\x01\n\nEntityType\x12\x1b\n\x17ENTITY_TYPE_UNSPECIFIED\x10\x00\x12\x14\n\x10ENTITY_TYPE_LANE\x10\x01\x12\x14\n\x10ENTITY_TYPE_ROAD\x10\x02\x12\x18\n\x14ENTITY_TYPE_JUNCTION\x10\x03\x12\x13\n\x0fENTITY_TYPE_AOI\x10\x04\x12\x13\n\x0fENTITY_TYPE_POI\x10\x05\x12\x16\n\x12ENTITY_TYPE_PERSON\x10\x06\x12\x13\n\x0fENTITY_TYPE_ORG\x10\x07B\xa9\x01\n\x11com.city.event.v2B\nEventProtoP\x01Z2git.fiblab.net/sim/protos/go/city/event/v2;eventv2\xa2\x02\x03CEX\xaa\x02\rCity.Event.V2\xca\x02\rCity\\Event\\V2\xe2\x02\x19City\\Event\\V2\\GPBMetadata\xea\x02\x0fCity::Event::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.event.v2.event_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x11com.city.event.v2B\nEventProtoP\x01Z2git.fiblab.net/sim/protos/go/city/event/v2;eventv2\xa2\x02\x03CEX\xaa\x02\rCity.Event.V2\xca\x02\rCity\\Event\\V2\xe2\x02\x19City\\Event\\V2\\GPBMetadata\xea\x02\x0fCity::Event::V2'
    _globals['_ENTITYTYPE']._serialized_start = 341
    _globals['_ENTITYTYPE']._serialized_end = 539
    _globals['_ENTITY']._serialized_start = 67
    _globals['_ENTITY']._serialized_end = 138
    _globals['_EVENT']._serialized_start = 141
    _globals['_EVENT']._serialized_end = 338