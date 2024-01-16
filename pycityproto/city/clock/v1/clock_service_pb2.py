"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!city/clock/v1/clock_service.proto\x12\rcity.clock.v1"\x0c\n\nNowRequest"\x1b\n\x0bNowResponse\x12\x0c\n\x01t\x18\x01 \x01(\x01R\x01t2L\n\x0cClockService\x12<\n\x03Now\x12\x19.city.clock.v1.NowRequest\x1a\x1a.city.clock.v1.NowResponseB\xb0\x01\n\x11com.city.clock.v1B\x11ClockServiceProtoP\x01Z2git.fiblab.net/sim/protos/go/city/clock/v1;clockv1\xa2\x02\x03CCX\xaa\x02\rCity.Clock.V1\xca\x02\rCity\\Clock\\V1\xe2\x02\x19City\\Clock\\V1\\GPBMetadata\xea\x02\x0fCity::Clock::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.clock.v1.clock_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x11com.city.clock.v1B\x11ClockServiceProtoP\x01Z2git.fiblab.net/sim/protos/go/city/clock/v1;clockv1\xa2\x02\x03CCX\xaa\x02\rCity.Clock.V1\xca\x02\rCity\\Clock\\V1\xe2\x02\x19City\\Clock\\V1\\GPBMetadata\xea\x02\x0fCity::Clock::V1'
    _globals['_NOWREQUEST']._serialized_start = 52
    _globals['_NOWREQUEST']._serialized_end = 64
    _globals['_NOWRESPONSE']._serialized_start = 66
    _globals['_NOWRESPONSE']._serialized_end = 93
    _globals['_CLOCKSERVICE']._serialized_start = 95
    _globals['_CLOCKSERVICE']._serialized_end = 171