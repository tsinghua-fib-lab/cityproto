"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.routing.v2 import routing_pb2 as city_dot_routing_dot_v2_dot_routing__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17city/trip/v2/trip.proto\x12\x0ccity.trip.v2\x1a\x15city/geo/v2/geo.proto\x1a\x1dcity/routing/v2/routing.proto"\x8c\x02\n\x08TripStop\x12@\n\x0caoi_position\x18\x01 \x01(\x0b2\x18.city.geo.v2.AoiPositionH\x00R\x0baoiPosition\x88\x01\x01\x12>\n\rlane_position\x18\x02 \x01(\x0b2\x19.city.geo.v2.LanePositionR\x0clanePosition\x12\x1a\n\x08duration\x18\x03 \x01(\x01R\x08duration\x12Q\n\x17optional_lane_positions\x18\x04 \x03(\x0b2\x19.city.geo.v2.LanePositionR\x15optionalLanePositionsB\x0f\n\r_aoi_position"\xbf\x03\n\x04Trip\x12*\n\x04mode\x18\x01 \x01(\x0e2\x16.city.trip.v2.TripModeR\x04mode\x12\'\n\x03end\x18\x02 \x01(\x0b2\x15.city.geo.v2.PositionR\x03end\x12*\n\x0edeparture_time\x18\x03 \x01(\x01H\x00R\rdepartureTime\x88\x01\x01\x12 \n\twait_time\x18\x04 \x01(\x01H\x01R\x08waitTime\x88\x01\x01\x12&\n\x0carrival_time\x18\x05 \x01(\x01H\x02R\x0barrivalTime\x88\x01\x01\x12\x1f\n\x08activity\x18\x06 \x01(\tH\x03R\x08activity\x88\x01\x01\x12\x19\n\x05model\x18\x08 \x01(\tH\x04R\x05model\x88\x01\x01\x120\n\x06routes\x18\x07 \x03(\x0b2\x18.city.routing.v2.JourneyR\x06routes\x125\n\ntrip_stops\x18\t \x03(\x0b2\x16.city.trip.v2.TripStopR\ttripStopsB\x11\n\x0f_departure_timeB\x0c\n\n_wait_timeB\x0f\n\r_arrival_timeB\x0b\n\t_activityB\x08\n\x06_model"\xc2\x01\n\x08Schedule\x12(\n\x05trips\x18\x01 \x03(\x0b2\x12.city.trip.v2.TripR\x05trips\x12\x1d\n\nloop_count\x18\x02 \x01(\x05R\tloopCount\x12*\n\x0edeparture_time\x18\x03 \x01(\x01H\x00R\rdepartureTime\x88\x01\x01\x12 \n\twait_time\x18\x04 \x01(\x01H\x01R\x08waitTime\x88\x01\x01B\x11\n\x0f_departure_timeB\x0c\n\n_wait_time*\xd7\x01\n\x08TripMode\x12\x19\n\x15TRIP_MODE_UNSPECIFIED\x10\x00\x12\x17\n\x13TRIP_MODE_WALK_ONLY\x10\x01\x12\x18\n\x14TRIP_MODE_DRIVE_ONLY\x10\x02\x12\x16\n\x12TRIP_MODE_BUS_WALK\x10\x04\x12\x17\n\x13TRIP_MODE_BIKE_WALK\x10\x05\x12\x12\n\x0eTRIP_MODE_TAXI\x10\x06\x12\x19\n\x15TRIP_MODE_SUBWAY_WALK\x10\x07\x12\x1d\n\x19TRIP_MODE_BUS_SUBWAY_WALK\x10\x08B\xa4\x01\n\x10com.city.trip.v2B\tTripProtoP\x01Z3git.fiblab.net/sim/protos/v2/go/city/trip/v2;tripv2\xa2\x02\x03CTX\xaa\x02\x0cCity.Trip.V2\xca\x02\x0cCity\\Trip\\V2\xe2\x02\x18City\\Trip\\V2\\GPBMetadata\xea\x02\x0eCity::Trip::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.trip.v2.trip_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x10com.city.trip.v2B\tTripProtoP\x01Z3git.fiblab.net/sim/protos/v2/go/city/trip/v2;tripv2\xa2\x02\x03CTX\xaa\x02\x0cCity.Trip.V2\xca\x02\x0cCity\\Trip\\V2\xe2\x02\x18City\\Trip\\V2\\GPBMetadata\xea\x02\x0eCity::Trip::V2'
    _globals['_TRIPMODE']._serialized_start = 1014
    _globals['_TRIPMODE']._serialized_end = 1229
    _globals['_TRIPSTOP']._serialized_start = 96
    _globals['_TRIPSTOP']._serialized_end = 364
    _globals['_TRIP']._serialized_start = 367
    _globals['_TRIP']._serialized_end = 814
    _globals['_SCHEDULE']._serialized_start = 817
    _globals['_SCHEDULE']._serialized_end = 1011