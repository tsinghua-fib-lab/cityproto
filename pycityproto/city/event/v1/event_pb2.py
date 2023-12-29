"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19city/event/v1/event.proto\x12\rcity.event.v1"_\n\x05Event\x12,\n\x04type\x18\x01 \x01(\x0e2\x18.city.event.v1.EventTypeR\x04type\x12\x14\n\x05level\x18\x02 \x01(\x05R\x05level\x12\x12\n\x04step\x18\x03 \x01(\x05R\x04step"6\n\x06Events\x12,\n\x06events\x18\x01 \x03(\x0b2\x14.city.event.v1.EventR\x06events*\xbb\x03\n\tEventType\x12\x1a\n\x16EVENT_TYPE_UNSPECIFIED\x10\x00\x12\x19\n\x15EVENT_TYPE_HEAVY_RAIN\x10\x01\x12\x1e\n\x1aEVENT_TYPE_MILITARY_STRIKE\x10\x02\x12!\n\x1dEVENT_TYPE_TRAFFIC_CONGESTION\x10\x03\x12\'\n#EVENT_TYPE_TRAFFIC_LANE_RESTRICTION\x10\x04\x12 \n\x1cEVENT_TYPE_TRAFFIC_BAD_LIGHT\x10\x05\x12&\n"EVENT_TYPE_ELEC_RUINED_TRANSFORMER\x10\x06\x12 \n\x1cEVENT_TYPE_WATER_RUINED_PUMP\x10\x07\x12!\n\x1dEVENT_TYPE_WATER_STOPPED_PUMP\x10\x08\x12\'\n#EVENT_TYPE_COMM_RUINED_BASE_STATION\x10\t\x12(\n$EVENT_TYPE_COMM_STOPPED_BASE_STATION\x10\n\x12)\n%EVENT_TYPE_COMM_OVERLOAD_BASE_STATION\x10\x0bB\xa9\x01\n\x11com.city.event.v1B\nEventProtoP\x01Z2git.fiblab.net/sim/protos/go/city/event/v1;eventv1\xa2\x02\x03CEX\xaa\x02\rCity.Event.V1\xca\x02\rCity\\Event\\V1\xe2\x02\x19City\\Event\\V1\\GPBMetadata\xea\x02\x0fCity::Event::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.event.v1.event_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x11com.city.event.v1B\nEventProtoP\x01Z2git.fiblab.net/sim/protos/go/city/event/v1;eventv1\xa2\x02\x03CEX\xaa\x02\rCity.Event.V1\xca\x02\rCity\\Event\\V1\xe2\x02\x19City\\Event\\V1\\GPBMetadata\xea\x02\x0fCity::Event::V1'
    _globals['_EVENTTYPE']._serialized_start = 198
    _globals['_EVENTTYPE']._serialized_end = 641
    _globals['_EVENT']._serialized_start = 44
    _globals['_EVENT']._serialized_end = 139
    _globals['_EVENTS']._serialized_start = 141
    _globals['_EVENTS']._serialized_end = 195