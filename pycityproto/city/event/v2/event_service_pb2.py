"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.event.v2 import event_pb2 as city_dot_event_dot_v2_dot_event__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!city/event/v2/event_service.proto\x12\rcity.event.v2\x1a\x19city/event/v2/event.proto"/\n\x17GetEventsByTopicRequest\x12\x14\n\x05topic\x18\x01 \x01(\tR\x05topic"H\n\x18GetEventsByTopicResponse\x12,\n\x06events\x18\x01 \x03(\x0b2\x14.city.event.v2.EventR\x06events"D\n\x14ResolveEventsRequest\x12,\n\x06events\x18\x01 \x03(\x0b2\x14.city.event.v2.EventR\x06events"\x17\n\x15ResolveEventsResponse2\xcf\x01\n\x0cEventService\x12c\n\x10GetEventsByTopic\x12&.city.event.v2.GetEventsByTopicRequest\x1a\'.city.event.v2.GetEventsByTopicResponse\x12Z\n\rResolveEvents\x12#.city.event.v2.ResolveEventsRequest\x1a$.city.event.v2.ResolveEventsResponseB\xb0\x01\n\x11com.city.event.v2B\x11EventServiceProtoP\x01Z2git.fiblab.net/sim/protos/go/city/event/v2;eventv2\xa2\x02\x03CEX\xaa\x02\rCity.Event.V2\xca\x02\rCity\\Event\\V2\xe2\x02\x19City\\Event\\V2\\GPBMetadata\xea\x02\x0fCity::Event::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.event.v2.event_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x11com.city.event.v2B\x11EventServiceProtoP\x01Z2git.fiblab.net/sim/protos/go/city/event/v2;eventv2\xa2\x02\x03CEX\xaa\x02\rCity.Event.V2\xca\x02\rCity\\Event\\V2\xe2\x02\x19City\\Event\\V2\\GPBMetadata\xea\x02\x0fCity::Event::V2'
    _globals['_GETEVENTSBYTOPICREQUEST']._serialized_start = 79
    _globals['_GETEVENTSBYTOPICREQUEST']._serialized_end = 126
    _globals['_GETEVENTSBYTOPICRESPONSE']._serialized_start = 128
    _globals['_GETEVENTSBYTOPICRESPONSE']._serialized_end = 200
    _globals['_RESOLVEEVENTSREQUEST']._serialized_start = 202
    _globals['_RESOLVEEVENTSREQUEST']._serialized_end = 270
    _globals['_RESOLVEEVENTSRESPONSE']._serialized_start = 272
    _globals['_RESOLVEEVENTSRESPONSE']._serialized_end = 295
    _globals['_EVENTSERVICE']._serialized_start = 298
    _globals['_EVENTSERVICE']._serialized_end = 505