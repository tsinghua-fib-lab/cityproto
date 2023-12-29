"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.event.v1 import event_pb2 as city_dot_event_dot_v1_dot_event__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!city/event/v1/event_service.proto\x12\rcity.event.v1\x1a\x19city/event/v1/event.proto"<\n\x0ePublishRequest\x12*\n\x05event\x18\x01 \x01(\x0b2\x14.city.event.v1.EventR\x05event"\x11\n\x0fPublishResponse"\r\n\x0bPullRequest"<\n\x0cPullResponse\x12,\n\x06events\x18\x01 \x03(\x0b2\x14.city.event.v1.EventR\x06events2\x9d\x01\n\x0cEventService\x12J\n\x07Publish\x12\x1d.city.event.v1.PublishRequest\x1a\x1e.city.event.v1.PublishResponse"\x00\x12A\n\x04Pull\x12\x1a.city.event.v1.PullRequest\x1a\x1b.city.event.v1.PullResponse"\x00B\xb0\x01\n\x11com.city.event.v1B\x11EventServiceProtoP\x01Z2git.fiblab.net/sim/protos/go/city/event/v1;eventv1\xa2\x02\x03CEX\xaa\x02\rCity.Event.V1\xca\x02\rCity\\Event\\V1\xe2\x02\x19City\\Event\\V1\\GPBMetadata\xea\x02\x0fCity::Event::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.event.v1.event_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x11com.city.event.v1B\x11EventServiceProtoP\x01Z2git.fiblab.net/sim/protos/go/city/event/v1;eventv1\xa2\x02\x03CEX\xaa\x02\rCity.Event.V1\xca\x02\rCity\\Event\\V1\xe2\x02\x19City\\Event\\V1\\GPBMetadata\xea\x02\x0fCity::Event::V1'
    _globals['_PUBLISHREQUEST']._serialized_start = 79
    _globals['_PUBLISHREQUEST']._serialized_end = 139
    _globals['_PUBLISHRESPONSE']._serialized_start = 141
    _globals['_PUBLISHRESPONSE']._serialized_end = 158
    _globals['_PULLREQUEST']._serialized_start = 160
    _globals['_PULLREQUEST']._serialized_end = 173
    _globals['_PULLRESPONSE']._serialized_start = 175
    _globals['_PULLRESPONSE']._serialized_end = 235
    _globals['_EVENTSERVICE']._serialized_start = 238
    _globals['_EVENTSERVICE']._serialized_end = 395