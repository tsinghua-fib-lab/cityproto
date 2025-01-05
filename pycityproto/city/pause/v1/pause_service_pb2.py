"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!city/pause/v1/pause_service.proto\x12\rcity.pause.v1"0\n\x0cPauseRequest\x12\x17\n\x04name\x18\x01 \x01(\tH\x00R\x04name\x88\x01\x01B\x07\n\x05_name"\x0f\n\rPauseResponse"1\n\rResumeRequest\x12\x17\n\x04name\x18\x01 \x01(\tH\x00R\x04name\x88\x01\x01B\x07\n\x05_name"\x10\n\x0eResumeResponse2\x99\x01\n\x0cPauseService\x12B\n\x05Pause\x12\x1b.city.pause.v1.PauseRequest\x1a\x1c.city.pause.v1.PauseResponse\x12E\n\x06Resume\x12\x1c.city.pause.v1.ResumeRequest\x1a\x1d.city.pause.v1.ResumeResponseB\xb3\x01\n\x11com.city.pause.v1B\x11PauseServiceProtoP\x01Z5git.fiblab.net/sim/protos/v2/go/city/pause/v1;pausev1\xa2\x02\x03CPX\xaa\x02\rCity.Pause.V1\xca\x02\rCity\\Pause\\V1\xe2\x02\x19City\\Pause\\V1\\GPBMetadata\xea\x02\x0fCity::Pause::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.pause.v1.pause_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x11com.city.pause.v1B\x11PauseServiceProtoP\x01Z5git.fiblab.net/sim/protos/v2/go/city/pause/v1;pausev1\xa2\x02\x03CPX\xaa\x02\rCity.Pause.V1\xca\x02\rCity\\Pause\\V1\xe2\x02\x19City\\Pause\\V1\\GPBMetadata\xea\x02\x0fCity::Pause::V1'
    _globals['_PAUSEREQUEST']._serialized_start = 52
    _globals['_PAUSEREQUEST']._serialized_end = 100
    _globals['_PAUSERESPONSE']._serialized_start = 102
    _globals['_PAUSERESPONSE']._serialized_end = 117
    _globals['_RESUMEREQUEST']._serialized_start = 119
    _globals['_RESUMEREQUEST']._serialized_end = 168
    _globals['_RESUMERESPONSE']._serialized_start = 170
    _globals['_RESUMERESPONSE']._serialized_end = 186
    _globals['_PAUSESERVICE']._serialized_start = 189
    _globals['_PAUSESERVICE']._serialized_end = 342