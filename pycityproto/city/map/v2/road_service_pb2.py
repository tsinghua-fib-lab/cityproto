"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.event.v1 import event_pb2 as city_dot_event_dot_v1_dot_event__pb2
from ....city.map.v2 import lane_state_pb2 as city_dot_map_dot_v2_dot_lane__state__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ecity/map/v2/road_service.proto\x12\x0bcity.map.v2\x1a\x19city/event/v1/event.proto\x1a\x1ccity/map/v2/lane_state.proto"u\n\x0eGetRoadRequest\x12\x19\n\x08road_ids\x18\x01 \x03(\x05R\x07roadIds\x12!\n\x0cexclude_lane\x18\x02 \x01(\x08R\x0bexcludeLane\x12%\n\x0eexclude_person\x18\x03 \x01(\x08R\rexcludePerson"A\n\x0fGetRoadResponse\x12.\n\x06states\x18\x01 \x03(\x0b2\x16.city.map.v2.RoadStateR\x06states" \n\x1eGetRoadGlobalStatisticsRequest"\\\n\x1fGetRoadGlobalStatisticsResponse\x129\n\x19avg_road_congestion_index\x18\x01 \x01(\x01R\x16avgRoadCongestionIndex"\x14\n\x12GetRuinInfoRequest"2\n\x08RuinInfo\x12\x10\n\x03num\x18\x01 \x01(\x05R\x03num\x12\x14\n\x05ratio\x18\x02 \x01(\x01R\x05ratio"\x94\x01\n\x13GetRuinInfoResponse\x12\'\n\x03one\x18\x01 \x01(\x0b2\x15.city.map.v2.RuinInfoR\x03one\x12\'\n\x03two\x18\x02 \x01(\x0b2\x15.city.map.v2.RuinInfoR\x03two\x12+\n\x05three\x18\x03 \x01(\x0b2\x15.city.map.v2.RuinInfoR\x05three"\x12\n\x10GetEventsRequest"B\n\x11GetEventsResponse\x12-\n\x06events\x18\x01 \x01(\x0b2\x15.city.event.v1.EventsR\x06events"\xf8\x03\n\tRoadState\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12$\n\x0ein_vehicle_cnt\x18\x07 \x01(\x05R\x0cinVehicleCnt\x12&\n\x0fout_vehicle_cnt\x18\x08 \x01(\x05R\routVehicleCnt\x12\x1f\n\x0bvehicle_cnt\x18\t \x01(\x05R\nvehicleCnt\x12+\n\x12cum_in_vehicle_cnt\x18\n \x01(\x05R\x0fcumInVehicleCnt\x12-\n\x13cum_out_vehicle_cnt\x18\x0b \x01(\x05R\x10cumOutVehicleCnt\x12\x13\n\x05avg_v\x18\x04 \x01(\x01R\x04avgV\x12&\n\x0favg_travel_time\x18\x0c \x01(\x01R\ravgTravelTime\x12,\n\x05level\x18\x02 \x01(\x0e2\x16.city.map.v2.RoadLevelR\x05level\x12)\n\x10congestion_index\x18\r \x01(\x01R\x0fcongestionIndex\x127\n\x06reason\x18\x03 \x01(\x0e2\x1f.city.map.v2.InterruptionReasonR\x06reason\x12,\n\x05lanes\x18\x05 \x03(\x0b2\x16.city.map.v2.LaneStateR\x05lanes\x12\x13\n\x05max_v\x18\x06 \x01(\x01R\x04maxV*\xc3\x01\n\tRoadLevel\x12\x1a\n\x16ROAD_LEVEL_UNSPECIFIED\x10\x00\x12\x14\n\x10ROAD_LEVEL_CLEAR\x10\x01\x12\x19\n\x15ROAD_LEVEL_LIGHT_LOAD\x10\x02\x12\x1a\n\x16ROAD_LEVEL_MEDIUM_LOAD\x10\x03\x12\x19\n\x15ROAD_LEVEL_HEAVY_LOAD\x10\x04\x12\x17\n\x13ROAD_LEVEL_OVERLOAD\x10\x05\x12\x19\n\x15ROAD_LEVEL_RESTRICTED\x10\x06*\x9e\x01\n\x12InterruptionReason\x12#\n\x1fINTERRUPTION_REASON_UNSPECIFIED\x10\x00\x12\x1e\n\x1aINTERRUPTION_REASON_RUINED\x10\x01\x12\x1f\n\x1bINTERRUPTION_REASON_CASCADE\x10\x02\x12"\n\x1eINTERRUPTION_REASON_CONGESTION\x10\x032\xe7\x02\n\x0bRoadService\x12D\n\x07GetRoad\x12\x1b.city.map.v2.GetRoadRequest\x1a\x1c.city.map.v2.GetRoadResponse\x12t\n\x17GetRoadGlobalStatistics\x12+.city.map.v2.GetRoadGlobalStatisticsRequest\x1a,.city.map.v2.GetRoadGlobalStatisticsResponse\x12P\n\x0bGetRuinInfo\x12\x1f.city.map.v2.GetRuinInfoRequest\x1a .city.map.v2.GetRuinInfoResponse\x12J\n\tGetEvents\x12\x1d.city.map.v2.GetEventsRequest\x1a\x1e.city.map.v2.GetEventsResponseB\xa4\x01\n\x0fcom.city.map.v2B\x10RoadServiceProtoP\x01Z1git.fiblab.net/sim/protos/v2/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.road_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\x10RoadServiceProtoP\x01Z1git.fiblab.net/sim/protos/v2/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_ROADLEVEL']._serialized_start = 1239
    _globals['_ROADLEVEL']._serialized_end = 1434
    _globals['_INTERRUPTIONREASON']._serialized_start = 1437
    _globals['_INTERRUPTIONREASON']._serialized_end = 1595
    _globals['_GETROADREQUEST']._serialized_start = 104
    _globals['_GETROADREQUEST']._serialized_end = 221
    _globals['_GETROADRESPONSE']._serialized_start = 223
    _globals['_GETROADRESPONSE']._serialized_end = 288
    _globals['_GETROADGLOBALSTATISTICSREQUEST']._serialized_start = 290
    _globals['_GETROADGLOBALSTATISTICSREQUEST']._serialized_end = 322
    _globals['_GETROADGLOBALSTATISTICSRESPONSE']._serialized_start = 324
    _globals['_GETROADGLOBALSTATISTICSRESPONSE']._serialized_end = 416
    _globals['_GETRUININFOREQUEST']._serialized_start = 418
    _globals['_GETRUININFOREQUEST']._serialized_end = 438
    _globals['_RUININFO']._serialized_start = 440
    _globals['_RUININFO']._serialized_end = 490
    _globals['_GETRUININFORESPONSE']._serialized_start = 493
    _globals['_GETRUININFORESPONSE']._serialized_end = 641
    _globals['_GETEVENTSREQUEST']._serialized_start = 643
    _globals['_GETEVENTSREQUEST']._serialized_end = 661
    _globals['_GETEVENTSRESPONSE']._serialized_start = 663
    _globals['_GETEVENTSRESPONSE']._serialized_end = 729
    _globals['_ROADSTATE']._serialized_start = 732
    _globals['_ROADSTATE']._serialized_end = 1236
    _globals['_ROADSERVICE']._serialized_start = 1598
    _globals['_ROADSERVICE']._serialized_end = 1957