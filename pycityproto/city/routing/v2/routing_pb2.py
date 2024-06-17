"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dcity/routing/v2/routing.proto\x12\x0fcity.routing.v2"A\n\x12DrivingJourneyBody\x12\x19\n\x08road_ids\x18\x02 \x03(\x05R\x07roadIds\x12\x10\n\x03eta\x18\x03 \x01(\x01R\x03eta"{\n\x13WalkingRouteSegment\x12\x17\n\x07lane_id\x18\x01 \x01(\x05R\x06laneId\x12K\n\x10moving_direction\x18\x02 \x01(\x0e2 .city.routing.v2.MovingDirectionR\x0fmovingDirection"b\n\x12WalkingJourneyBody\x12:\n\x05route\x18\x01 \x03(\x0b2$.city.routing.v2.WalkingRouteSegmentR\x05route\x12\x10\n\x03eta\x18\x02 \x01(\x01R\x03eta"\x80\x01\n\x0fTransferSegment\x12\x1d\n\nsubline_id\x18\x01 \x01(\x05R\tsublineId\x12(\n\x10start_station_id\x18\x02 \x01(\x05R\x0estartStationId\x12$\n\x0eend_station_id\x18\x03 \x01(\x05R\x0cendStationId"b\n\x0eBusJourneyBody\x12>\n\ttransfers\x18\x01 \x03(\x0b2 .city.routing.v2.TransferSegmentR\ttransfers\x12\x10\n\x03eta\x18\x02 \x01(\x01R\x03eta"\xa3\x02\n\x07Journey\x120\n\x04type\x18\x01 \x01(\x0e2\x1c.city.routing.v2.JourneyTypeR\x04type\x12B\n\x07driving\x18\x02 \x01(\x0b2#.city.routing.v2.DrivingJourneyBodyH\x00R\x07driving\x88\x01\x01\x12B\n\x07walking\x18\x03 \x01(\x0b2#.city.routing.v2.WalkingJourneyBodyH\x01R\x07walking\x88\x01\x01\x12;\n\x06by_bus\x18\x04 \x01(\x0b2\x1f.city.routing.v2.BusJourneyBodyH\x02R\x05byBus\x88\x01\x01B\n\n\x08_drivingB\n\n\x08_walkingB\t\n\x07_by_bus"2\n\nRoadStatus\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n\x05speed\x18\x02 \x03(\x01R\x05speed"P\n\x0cRoadStatuses\x12@\n\rroad_statuses\x18\x01 \x03(\x0b2\x1b.city.routing.v2.RoadStatusR\x0croadStatuses*n\n\tRouteType\x12\x1a\n\x16ROUTE_TYPE_UNSPECIFIED\x10\x00\x12\x16\n\x12ROUTE_TYPE_DRIVING\x10\x01\x12\x16\n\x12ROUTE_TYPE_WALKING\x10\x02\x12\x15\n\x11ROUTE_TYPE_BY_BUS\x10\x03*x\n\x0bJourneyType\x12\x1c\n\x18JOURNEY_TYPE_UNSPECIFIED\x10\x00\x12\x18\n\x14JOURNEY_TYPE_DRIVING\x10\x01\x12\x18\n\x14JOURNEY_TYPE_WALKING\x10\x02\x12\x17\n\x13JOURNEY_TYPE_BY_BUS\x10\x03*p\n\x0fMovingDirection\x12 \n\x1cMOVING_DIRECTION_UNSPECIFIED\x10\x00\x12\x1c\n\x18MOVING_DIRECTION_FORWARD\x10\x01\x12\x1d\n\x19MOVING_DIRECTION_BACKWARD\x10\x02B\xb9\x01\n\x13com.city.routing.v2B\x0cRoutingProtoP\x01Z6git.fiblab.net/sim/protos/go/city/routing/v2;routingv2\xa2\x02\x03CRX\xaa\x02\x0fCity.Routing.V2\xca\x02\x0fCity\\Routing\\V2\xe2\x02\x1bCity\\Routing\\V2\\GPBMetadata\xea\x02\x11City::Routing::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.routing.v2.routing_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.routing.v2B\x0cRoutingProtoP\x01Z6git.fiblab.net/sim/protos/go/city/routing/v2;routingv2\xa2\x02\x03CRX\xaa\x02\x0fCity.Routing.V2\xca\x02\x0fCity\\Routing\\V2\xe2\x02\x1bCity\\Routing\\V2\\GPBMetadata\xea\x02\x11City::Routing::V2'
    _globals['_ROUTETYPE']._serialized_start = 1001
    _globals['_ROUTETYPE']._serialized_end = 1111
    _globals['_JOURNEYTYPE']._serialized_start = 1113
    _globals['_JOURNEYTYPE']._serialized_end = 1233
    _globals['_MOVINGDIRECTION']._serialized_start = 1235
    _globals['_MOVINGDIRECTION']._serialized_end = 1347
    _globals['_DRIVINGJOURNEYBODY']._serialized_start = 50
    _globals['_DRIVINGJOURNEYBODY']._serialized_end = 115
    _globals['_WALKINGROUTESEGMENT']._serialized_start = 117
    _globals['_WALKINGROUTESEGMENT']._serialized_end = 240
    _globals['_WALKINGJOURNEYBODY']._serialized_start = 242
    _globals['_WALKINGJOURNEYBODY']._serialized_end = 340
    _globals['_TRANSFERSEGMENT']._serialized_start = 343
    _globals['_TRANSFERSEGMENT']._serialized_end = 471
    _globals['_BUSJOURNEYBODY']._serialized_start = 473
    _globals['_BUSJOURNEYBODY']._serialized_end = 571
    _globals['_JOURNEY']._serialized_start = 574
    _globals['_JOURNEY']._serialized_end = 865
    _globals['_ROADSTATUS']._serialized_start = 867
    _globals['_ROADSTATUS']._serialized_end = 917
    _globals['_ROADSTATUSES']._serialized_start = 919
    _globals['_ROADSTATUSES']._serialized_end = 999