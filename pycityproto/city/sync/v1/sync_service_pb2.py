"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fcity/sync/v1/sync_service.proto\x12\x0ccity.sync.v1"\x10\n\x0eGetEtcdRequest"%\n\x0fGetEtcdResponse\x12\x12\n\x04port\x18\x01 \x01(\x05R\x04port"5\n\x0bStepRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n\x04step\x18\x02 \x01(\x05R\x04step"$\n\x0cStepResponse\x12\x14\n\x05close\x18\x01 \x01(\x08R\x05close2\x94\x01\n\x0bSyncService\x12F\n\x07GetEtcd\x12\x1c.city.sync.v1.GetEtcdRequest\x1a\x1d.city.sync.v1.GetEtcdResponse\x12=\n\x04Step\x12\x19.city.sync.v1.StepRequest\x1a\x1a.city.sync.v1.StepResponseB\xa8\x01\n\x10com.city.sync.v1B\x10SyncServiceProtoP\x01Z0git.fiblab.net/sim/protos/go/city/sync/v1;syncv1\xa2\x02\x03CSX\xaa\x02\x0cCity.Sync.V1\xca\x02\x0cCity\\Sync\\V1\xe2\x02\x18City\\Sync\\V1\\GPBMetadata\xea\x02\x0eCity::Sync::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.sync.v1.sync_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x10com.city.sync.v1B\x10SyncServiceProtoP\x01Z0git.fiblab.net/sim/protos/go/city/sync/v1;syncv1\xa2\x02\x03CSX\xaa\x02\x0cCity.Sync.V1\xca\x02\x0cCity\\Sync\\V1\xe2\x02\x18City\\Sync\\V1\\GPBMetadata\xea\x02\x0eCity::Sync::V1'
    _globals['_GETETCDREQUEST']._serialized_start = 49
    _globals['_GETETCDREQUEST']._serialized_end = 65
    _globals['_GETETCDRESPONSE']._serialized_start = 67
    _globals['_GETETCDRESPONSE']._serialized_end = 104
    _globals['_STEPREQUEST']._serialized_start = 106
    _globals['_STEPREQUEST']._serialized_end = 159
    _globals['_STEPRESPONSE']._serialized_start = 161
    _globals['_STEPRESPONSE']._serialized_end = 197
    _globals['_SYNCSERVICE']._serialized_start = 200
    _globals['_SYNCSERVICE']._serialized_end = 348