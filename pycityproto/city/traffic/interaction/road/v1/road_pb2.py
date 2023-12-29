"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ......city.traffic.interaction.lane.v1 import lane_pb2 as city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+city/traffic/interaction/road/v1/road.proto\x12 city.traffic.interaction.road.v1\x1a+city/traffic/interaction/lane/v1/lane.proto"\xfc\x01\n\x05State\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x13\n\x05avg_v\x18\x04 \x01(\x01R\x04avgV\x12A\n\x05level\x18\x02 \x01(\x0e2+.city.traffic.interaction.road.v1.RoadLevelR\x05level\x12L\n\x06reason\x18\x03 \x01(\x0e24.city.traffic.interaction.road.v1.InterruptionReasonR\x06reason\x12=\n\x05lanes\x18\x05 \x03(\x0b2\'.city.traffic.interaction.lane.v1.StateR\x05lanes*\xc3\x01\n\tRoadLevel\x12\x1a\n\x16ROAD_LEVEL_UNSPECIFIED\x10\x00\x12\x14\n\x10ROAD_LEVEL_CLEAR\x10\x01\x12\x19\n\x15ROAD_LEVEL_LIGHT_LOAD\x10\x02\x12\x1a\n\x16ROAD_LEVEL_MEDIUM_LOAD\x10\x03\x12\x19\n\x15ROAD_LEVEL_HEAVY_LOAD\x10\x04\x12\x17\n\x13ROAD_LEVEL_OVERLOAD\x10\x05\x12\x19\n\x15ROAD_LEVEL_RESTRICTED\x10\x06*\x9e\x01\n\x12InterruptionReason\x12#\n\x1fINTERRUPTION_REASON_UNSPECIFIED\x10\x00\x12\x1e\n\x1aINTERRUPTION_REASON_RUINED\x10\x01\x12\x1f\n\x1bINTERRUPTION_REASON_CASCADE\x10\x02\x12"\n\x1eINTERRUPTION_REASON_CONGESTION\x10\x03B\x9c\x02\n$com.city.traffic.interaction.road.v1B\tRoadProtoP\x01ZDgit.fiblab.net/sim/protos/go/city/traffic/interaction/road/v1;roadv1\xa2\x02\x04CTIR\xaa\x02 City.Traffic.Interaction.Road.V1\xca\x02 City\\Traffic\\Interaction\\Road\\V1\xe2\x02,City\\Traffic\\Interaction\\Road\\V1\\GPBMetadata\xea\x02$City::Traffic::Interaction::Road::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.traffic.interaction.road.v1.road_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n$com.city.traffic.interaction.road.v1B\tRoadProtoP\x01ZDgit.fiblab.net/sim/protos/go/city/traffic/interaction/road/v1;roadv1\xa2\x02\x04CTIR\xaa\x02 City.Traffic.Interaction.Road.V1\xca\x02 City\\Traffic\\Interaction\\Road\\V1\xe2\x02,City\\Traffic\\Interaction\\Road\\V1\\GPBMetadata\xea\x02$City::Traffic::Interaction::Road::V1'
    _globals['_ROADLEVEL']._serialized_start = 382
    _globals['_ROADLEVEL']._serialized_end = 577
    _globals['_INTERRUPTIONREASON']._serialized_start = 580
    _globals['_INTERRUPTIONREASON']._serialized_end = 738
    _globals['_STATE']._serialized_start = 127
    _globals['_STATE']._serialized_end = 379