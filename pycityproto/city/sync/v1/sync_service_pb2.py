"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fcity/sync/v1/sync_service.proto\x12\x0ccity.sync.v1"5\n\rSetURLRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x10\n\x03url\x18\x02 \x01(\tR\x03url"\x10\n\x0eSetURLResponse"#\n\rGetURLRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name""\n\x0eGetURLResponse\x12\x10\n\x03url\x18\x01 \x01(\tR\x03url"5\n\x0bStepRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n\x04step\x18\x02 \x01(\x05R\x04step"$\n\x0cStepResponse\x12\x14\n\x05close\x18\x01 \x01(\x08R\x05close2\xd6\x01\n\x0bSyncService\x12C\n\x06SetURL\x12\x1b.city.sync.v1.SetURLRequest\x1a\x1c.city.sync.v1.SetURLResponse\x12C\n\x06GetURL\x12\x1b.city.sync.v1.GetURLRequest\x1a\x1c.city.sync.v1.GetURLResponse\x12=\n\x04Step\x12\x19.city.sync.v1.StepRequest\x1a\x1a.city.sync.v1.StepResponseB\xa8\x01\n\x10com.city.sync.v1B\x10SyncServiceProtoP\x01Z0git.fiblab.net/sim/protos/go/city/sync/v1;syncv1\xa2\x02\x03CSX\xaa\x02\x0cCity.Sync.V1\xca\x02\x0cCity\\Sync\\V1\xe2\x02\x18City\\Sync\\V1\\GPBMetadata\xea\x02\x0eCity::Sync::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.sync.v1.sync_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x10com.city.sync.v1B\x10SyncServiceProtoP\x01Z0git.fiblab.net/sim/protos/go/city/sync/v1;syncv1\xa2\x02\x03CSX\xaa\x02\x0cCity.Sync.V1\xca\x02\x0cCity\\Sync\\V1\xe2\x02\x18City\\Sync\\V1\\GPBMetadata\xea\x02\x0eCity::Sync::V1'
    _globals['_SETURLREQUEST']._serialized_start = 49
    _globals['_SETURLREQUEST']._serialized_end = 102
    _globals['_SETURLRESPONSE']._serialized_start = 104
    _globals['_SETURLRESPONSE']._serialized_end = 120
    _globals['_GETURLREQUEST']._serialized_start = 122
    _globals['_GETURLREQUEST']._serialized_end = 157
    _globals['_GETURLRESPONSE']._serialized_start = 159
    _globals['_GETURLRESPONSE']._serialized_end = 193
    _globals['_STEPREQUEST']._serialized_start = 195
    _globals['_STEPREQUEST']._serialized_end = 248
    _globals['_STEPRESPONSE']._serialized_start = 250
    _globals['_STEPRESPONSE']._serialized_end = 286
    _globals['_SYNCSERVICE']._serialized_start = 289
    _globals['_SYNCSERVICE']._serialized_end = 503