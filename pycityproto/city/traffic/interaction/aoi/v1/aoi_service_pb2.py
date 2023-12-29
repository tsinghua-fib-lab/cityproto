"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ......city.traffic.interaction.aoi.v1 import aoi_pb2 as city_dot_traffic_dot_interaction_dot_aoi_dot_v1_dot_aoi__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1city/traffic/interaction/aoi/v1/aoi_service.proto\x12\x1fcity.traffic.interaction.aoi.v1\x1a)city/traffic/interaction/aoi/v1/aoi.proto"(\n\rGetAoiRequest\x12\x17\n\x07aoi_ids\x18\x01 \x03(\x05R\x06aoiIds"P\n\x0eGetAoiResponse\x12>\n\x06states\x18\x01 \x03(\x0b2&.city.traffic.interaction.aoi.v1.StateR\x06states2w\n\nAoiService\x12i\n\x06GetAoi\x12..city.traffic.interaction.aoi.v1.GetAoiRequest\x1a/.city.traffic.interaction.aoi.v1.GetAoiResponseB\x9b\x02\n#com.city.traffic.interaction.aoi.v1B\x0fAoiServiceProtoP\x01ZBgit.fiblab.net/sim/protos/go/city/traffic/interaction/aoi/v1;aoiv1\xa2\x02\x04CTIA\xaa\x02\x1fCity.Traffic.Interaction.Aoi.V1\xca\x02\x1fCity\\Traffic\\Interaction\\Aoi\\V1\xe2\x02+City\\Traffic\\Interaction\\Aoi\\V1\\GPBMetadata\xea\x02#City::Traffic::Interaction::Aoi::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.traffic.interaction.aoi.v1.aoi_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n#com.city.traffic.interaction.aoi.v1B\x0fAoiServiceProtoP\x01ZBgit.fiblab.net/sim/protos/go/city/traffic/interaction/aoi/v1;aoiv1\xa2\x02\x04CTIA\xaa\x02\x1fCity.Traffic.Interaction.Aoi.V1\xca\x02\x1fCity\\Traffic\\Interaction\\Aoi\\V1\xe2\x02+City\\Traffic\\Interaction\\Aoi\\V1\\GPBMetadata\xea\x02#City::Traffic::Interaction::Aoi::V1'
    _globals['_GETAOIREQUEST']._serialized_start = 129
    _globals['_GETAOIREQUEST']._serialized_end = 169
    _globals['_GETAOIRESPONSE']._serialized_start = 171
    _globals['_GETAOIRESPONSE']._serialized_end = 251
    _globals['_AOISERVICE']._serialized_start = 253
    _globals['_AOISERVICE']._serialized_end = 372