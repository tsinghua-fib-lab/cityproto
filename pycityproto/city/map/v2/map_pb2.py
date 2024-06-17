"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.map.v2 import light_pb2 as city_dot_map_dot_v2_dot_light__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15city/map/v2/map.proto\x12\x0bcity.map.v2\x1a\x15city/geo/v2/geo.proto\x1a\x17city/map/v2/light.proto"9\n\x08Polyline\x12-\n\x05nodes\x18\x01 \x03(\x0b2\x17.city.geo.v2.XYPositionR\x05nodes"\x88\x02\n\x06Header\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n\x04date\x18\x02 \x01(\tR\x04date\x12\x14\n\x05north\x18\x03 \x01(\x01R\x05north\x12\x14\n\x05south\x18\x04 \x01(\x01R\x05south\x12\x12\n\x04east\x18\x05 \x01(\x01R\x04east\x12\x12\n\x04west\x18\x06 \x01(\x01R\x04west\x12\x1e\n\nprojection\x18\x07 \x01(\tR\nprojection\x12!\n\ntaz_x_step\x18\x08 \x01(\x01H\x00R\x08tazXStep\x88\x01\x01\x12!\n\ntaz_y_step\x18\t \x01(\x01H\x01R\x08tazYStep\x88\x01\x01B\r\n\x0b_taz_x_stepB\r\n\x0b_taz_y_step"\x8c\x01\n\x0bLaneOverlap\x12-\n\x04self\x18\x01 \x01(\x0b2\x19.city.geo.v2.LanePositionR\x04self\x12/\n\x05other\x18\x02 \x01(\x0b2\x19.city.geo.v2.LanePositionR\x05other\x12\x1d\n\nself_first\x18\x03 \x01(\x08R\tselfFirst"U\n\x0eLaneConnection\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x123\n\x04type\x18\x02 \x01(\x0e2\x1f.city.map.v2.LaneConnectionTypeR\x04type"\xaf\x05\n\x04Lane\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12)\n\x04type\x18\x02 \x01(\x0e2\x15.city.map.v2.LaneTypeR\x04type\x12)\n\x04turn\x18\x03 \x01(\x0e2\x15.city.map.v2.LaneTurnR\x04turn\x12\x1b\n\tmax_speed\x18\x04 \x01(\x01R\x08maxSpeed\x12\x16\n\x06length\x18\x05 \x01(\x01R\x06length\x12\x14\n\x05width\x18\x06 \x01(\x01R\x05width\x126\n\x0bcenter_line\x18\x07 \x01(\x0b2\x15.city.map.v2.PolylineR\ncenterLine\x12C\n\x10left_border_line\x18\x08 \x01(\x0b2\x15.city.map.v2.PolylineB\x02\x18\x01R\x0eleftBorderLine\x12E\n\x11right_border_line\x18\t \x01(\x0b2\x15.city.map.v2.PolylineB\x02\x18\x01R\x0frightBorderLine\x12?\n\x0cpredecessors\x18\n \x03(\x0b2\x1b.city.map.v2.LaneConnectionR\x0cpredecessors\x12;\n\nsuccessors\x18\x0b \x03(\x0b2\x1b.city.map.v2.LaneConnectionR\nsuccessors\x12"\n\rleft_lane_ids\x18\x0c \x03(\x05R\x0bleftLaneIds\x12$\n\x0eright_lane_ids\x18\r \x03(\x05R\x0crightLaneIds\x12\x1b\n\tparent_id\x18\x0e \x01(\x05R\x08parentId\x124\n\x08overlaps\x18\x0f \x03(\x0b2\x18.city.map.v2.LaneOverlapR\x08overlaps\x12\x17\n\x07aoi_ids\x18\x10 \x03(\x05R\x06aoiIds"_\n\x0cNextRoadLane\x12\x17\n\x07road_id\x18\x01 \x01(\x05R\x06roadId\x12\x1a\n\tlane_id_a\x18\x02 \x01(\x05R\x07laneIdA\x12\x1a\n\tlane_id_b\x18\x03 \x01(\x05R\x07laneIdB"U\n\x10NextRoadLanePlan\x12A\n\x0fnext_road_lanes\x18\x01 \x03(\x0b2\x19.city.map.v2.NextRoadLaneR\rnextRoadLanes"\x95\x01\n\x04Road\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n\x04name\x18\x04 \x01(\tR\x04name\x12\x19\n\x08lane_ids\x18\x02 \x03(\x05R\x07laneIds\x12N\n\x14next_road_lane_plans\x18\x03 \x03(\x0b2\x1d.city.map.v2.NextRoadLanePlanR\x11nextRoadLanePlans"\xcf\x01\n\x11JunctionLaneGroup\x12\x1c\n\nin_road_id\x18\x01 \x01(\x05R\x08inRoadId\x12\x19\n\x08in_angle\x18\x02 \x01(\x01R\x07inAngle\x12\x1e\n\x0bout_road_id\x18\x03 \x01(\x05R\toutRoadId\x12\x1b\n\tout_angle\x18\x04 \x01(\x01R\x08outAngle\x12\x19\n\x08lane_ids\x18\x05 \x03(\x05R\x07laneIds\x12)\n\x04turn\x18\x06 \x01(\x0e2\x15.city.map.v2.LaneTurnR\x04turn"\x91\x02\n\x08Junction\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x19\n\x08lane_ids\x18\x02 \x03(\x05R\x07laneIds\x12N\n\x13driving_lane_groups\x18\x03 \x03(\x0b2\x1e.city.map.v2.JunctionLaneGroupR\x11drivingLaneGroups\x123\n\x06phases\x18\x04 \x03(\x0b2\x1b.city.map.v2.AvailablePhaseR\x06phases\x12C\n\rfixed_program\x18\x05 \x01(\x0b2\x19.city.map.v2.TrafficLightH\x00R\x0cfixedProgram\x88\x01\x01B\x10\n\x0e_fixed_program"$\n\x07RoadIds\x12\x19\n\x08road_ids\x18\x01 \x03(\x05R\x07roadIds"^\n\x10SublineSchedules\x12\'\n\x0fdeparture_times\x18\x01 \x03(\x01R\x0edepartureTimes\x12!\n\x0coffset_times\x18\x02 \x03(\x01R\x0boffsetTimes"q\n\x10HeuristicTAZCost\x12\x18\n\x08taz_x_id\x18\x01 \x01(\x05R\x06tazXId\x12\x18\n\x08taz_y_id\x18\x02 \x01(\x05R\x06tazYId\x12\x15\n\x06aoi_id\x18\x03 \x01(\x05R\x05aoiId\x12\x12\n\x04cost\x18\x04 \x01(\x01R\x04cost"\xf2\x02\n\x16PublicTransportSubline\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x17\n\x07aoi_ids\x18\x03 \x03(\x05R\x06aoiIds\x12S\n\x1bstation_connection_road_ids\x18\x04 \x03(\x0b2\x14.city.map.v2.RoadIdsR\x18stationConnectionRoadIds\x12,\n\x04type\x18\x05 \x01(\x0e2\x18.city.map.v2.SublineTypeR\x04type\x12\x1f\n\x0bparent_name\x18\x06 \x01(\tR\nparentName\x12;\n\tschedules\x18\x07 \x01(\x0b2\x1d.city.map.v2.SublineSchedulesR\tschedules\x12:\n\ttaz_costs\x18\x08 \x03(\x0b2\x1d.city.map.v2.HeuristicTAZCostR\x08tazCosts"\xff\x04\n\x03Aoi\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n\x04name\x18\x0b \x01(\tR\x04name\x12,\n\x04type\x18\x02 \x01(\x0e2\x14.city.map.v2.AoiTypeB\x02\x18\x01R\x04type\x12F\n\x11driving_positions\x18\x03 \x03(\x0b2\x19.city.geo.v2.LanePositionR\x10drivingPositions\x12F\n\x11walking_positions\x18\x04 \x03(\x0b2\x19.city.geo.v2.LanePositionR\x10walkingPositions\x125\n\tpositions\x18\x05 \x03(\x0b2\x17.city.geo.v2.XYPositionR\tpositions\x12<\n\rdriving_gates\x18\x06 \x03(\x0b2\x17.city.geo.v2.XYPositionR\x0cdrivingGates\x12<\n\rwalking_gates\x18\x07 \x03(\x0b2\x17.city.geo.v2.XYPositionR\x0cwalkingGates\x12\x17\n\x04area\x18\x08 \x01(\x01H\x00R\x04area\x88\x01\x01\x12<\n\x08land_use\x18\n \x01(\x0e2\x18.city.map.v2.LandUseTypeB\x02\x18\x01H\x01R\x07landUse\x88\x01\x01\x12)\n\x0eurban_land_use\x18\x0c \x01(\tH\x02R\x0curbanLandUse\x88\x01\x01\x12\x17\n\x07poi_ids\x18\t \x03(\x05R\x06poiIds\x12\x1f\n\x0bsubline_ids\x18\r \x03(\x05R\nsublineIdsB\x07\n\x05_areaB\x0b\n\t_land_useB\x11\n\x0f_urban_land_use"\xdd\x01\n\x03Poi\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n\x08category\x18\x03 \x01(\tR\x08category\x123\n\x08position\x18\x04 \x01(\x0b2\x17.city.geo.v2.XYPositionR\x08position\x12\x15\n\x06aoi_id\x18\x05 \x01(\x05R\x05aoiId\x12\x1f\n\x08capacity\x18\x06 \x01(\x05H\x00R\x08capacity\x88\x01\x01\x12\x1c\n\tfunctions\x18\x07 \x03(\tR\tfunctionsB\x0b\n\t_capacity"\xc6\x02\n\x03Map\x12+\n\x06header\x18\x01 \x01(\x0b2\x13.city.map.v2.HeaderR\x06header\x12\'\n\x05lanes\x18\x02 \x03(\x0b2\x11.city.map.v2.LaneR\x05lanes\x12\'\n\x05roads\x18\x03 \x03(\x0b2\x11.city.map.v2.RoadR\x05roads\x123\n\tjunctions\x18\x04 \x03(\x0b2\x15.city.map.v2.JunctionR\tjunctions\x12$\n\x04aois\x18\x05 \x03(\x0b2\x10.city.map.v2.AoiR\x04aois\x12$\n\x04pois\x18\x06 \x03(\x0b2\x10.city.map.v2.PoiR\x04pois\x12?\n\x08sublines\x18\x07 \x03(\x0b2#.city.map.v2.PublicTransportSublineR\x08sublines*S\n\x08LaneType\x12\x19\n\x15LANE_TYPE_UNSPECIFIED\x10\x00\x12\x15\n\x11LANE_TYPE_DRIVING\x10\x01\x12\x15\n\x11LANE_TYPE_WALKING\x10\x02*|\n\x08LaneTurn\x12\x19\n\x15LANE_TURN_UNSPECIFIED\x10\x00\x12\x16\n\x12LANE_TURN_STRAIGHT\x10\x01\x12\x12\n\x0eLANE_TURN_LEFT\x10\x02\x12\x13\n\x0fLANE_TURN_RIGHT\x10\x03\x12\x14\n\x10LANE_TURN_AROUND\x10\x04*x\n\x12LaneConnectionType\x12$\n LANE_CONNECTION_TYPE_UNSPECIFIED\x10\x00\x12\x1d\n\x19LANE_CONNECTION_TYPE_HEAD\x10\x01\x12\x1d\n\x19LANE_CONNECTION_TYPE_TAIL\x10\x02*Q\n\x07AoiType\x12\x18\n\x14AOI_TYPE_UNSPECIFIED\x10\x00\x12\x18\n\x14AOI_TYPE_BUS_STATION\x10\x01\x12\x12\n\x0eAOI_TYPE_OTHER\x10\x02*\xdc\x01\n\x0bLandUseType\x12\x1d\n\x19LAND_USE_TYPE_UNSPECIFIED\x10\x00\x12\x1c\n\x18LAND_USE_TYPE_COMMERCIAL\x10\x05\x12\x1c\n\x18LAND_USE_TYPE_INDUSTRIAL\x10\x06\x12\x1d\n\x19LAND_USE_TYPE_RESIDENTIAL\x10\x07\x12\x18\n\x14LAND_USE_TYPE_PUBLIC\x10\x08\x12 \n\x1cLAND_USE_TYPE_TRANSPORTATION\x10\n\x12\x17\n\x13LAND_USE_TYPE_OTHER\x10\x0c*Z\n\x0bSublineType\x12\x1c\n\x18SUBLINE_TYPE_UNSPECIFIED\x10\x00\x12\x14\n\x10SUBLINE_TYPE_BUS\x10\x01\x12\x17\n\x13SUBLINE_TYPE_SUBWAY\x10\x02B\x99\x01\n\x0fcom.city.map.v2B\x08MapProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.map_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\x08MapProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_LANE'].fields_by_name['left_border_line']._options = None
    _globals['_LANE'].fields_by_name['left_border_line']._serialized_options = b'\x18\x01'
    _globals['_LANE'].fields_by_name['right_border_line']._options = None
    _globals['_LANE'].fields_by_name['right_border_line']._serialized_options = b'\x18\x01'
    _globals['_AOI'].fields_by_name['type']._options = None
    _globals['_AOI'].fields_by_name['type']._serialized_options = b'\x18\x01'
    _globals['_AOI'].fields_by_name['land_use']._options = None
    _globals['_AOI'].fields_by_name['land_use']._serialized_options = b'\x18\x01'
    _globals['_LANETYPE']._serialized_start = 3971
    _globals['_LANETYPE']._serialized_end = 4054
    _globals['_LANETURN']._serialized_start = 4056
    _globals['_LANETURN']._serialized_end = 4180
    _globals['_LANECONNECTIONTYPE']._serialized_start = 4182
    _globals['_LANECONNECTIONTYPE']._serialized_end = 4302
    _globals['_AOITYPE']._serialized_start = 4304
    _globals['_AOITYPE']._serialized_end = 4385
    _globals['_LANDUSETYPE']._serialized_start = 4388
    _globals['_LANDUSETYPE']._serialized_end = 4608
    _globals['_SUBLINETYPE']._serialized_start = 4610
    _globals['_SUBLINETYPE']._serialized_end = 4700
    _globals['_POLYLINE']._serialized_start = 86
    _globals['_POLYLINE']._serialized_end = 143
    _globals['_HEADER']._serialized_start = 146
    _globals['_HEADER']._serialized_end = 410
    _globals['_LANEOVERLAP']._serialized_start = 413
    _globals['_LANEOVERLAP']._serialized_end = 553
    _globals['_LANECONNECTION']._serialized_start = 555
    _globals['_LANECONNECTION']._serialized_end = 640
    _globals['_LANE']._serialized_start = 643
    _globals['_LANE']._serialized_end = 1330
    _globals['_NEXTROADLANE']._serialized_start = 1332
    _globals['_NEXTROADLANE']._serialized_end = 1427
    _globals['_NEXTROADLANEPLAN']._serialized_start = 1429
    _globals['_NEXTROADLANEPLAN']._serialized_end = 1514
    _globals['_ROAD']._serialized_start = 1517
    _globals['_ROAD']._serialized_end = 1666
    _globals['_JUNCTIONLANEGROUP']._serialized_start = 1669
    _globals['_JUNCTIONLANEGROUP']._serialized_end = 1876
    _globals['_JUNCTION']._serialized_start = 1879
    _globals['_JUNCTION']._serialized_end = 2152
    _globals['_ROADIDS']._serialized_start = 2154
    _globals['_ROADIDS']._serialized_end = 2190
    _globals['_SUBLINESCHEDULES']._serialized_start = 2192
    _globals['_SUBLINESCHEDULES']._serialized_end = 2286
    _globals['_HEURISTICTAZCOST']._serialized_start = 2288
    _globals['_HEURISTICTAZCOST']._serialized_end = 2401
    _globals['_PUBLICTRANSPORTSUBLINE']._serialized_start = 2404
    _globals['_PUBLICTRANSPORTSUBLINE']._serialized_end = 2774
    _globals['_AOI']._serialized_start = 2777
    _globals['_AOI']._serialized_end = 3416
    _globals['_POI']._serialized_start = 3419
    _globals['_POI']._serialized_end = 3640
    _globals['_MAP']._serialized_start = 3643
    _globals['_MAP']._serialized_end = 3969