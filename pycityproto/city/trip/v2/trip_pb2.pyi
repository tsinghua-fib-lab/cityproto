from city.geo.v2 import geo_pb2 as _geo_pb2
from city.routing.v2 import routing_pb2 as _routing_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class TripMode(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    TRIP_MODE_UNSPECIFIED: _ClassVar[TripMode]
    TRIP_MODE_WALK_ONLY: _ClassVar[TripMode]
    TRIP_MODE_DRIVE_ONLY: _ClassVar[TripMode]
    TRIP_MODE_BUS_WALK: _ClassVar[TripMode]
    TRIP_MODE_BIKE_WALK: _ClassVar[TripMode]
    TRIP_MODE_TAXI: _ClassVar[TripMode]
    TRIP_MODE_SUBWAY_WALK: _ClassVar[TripMode]
    TRIP_MODE_BUS_SUBWAY_WALK: _ClassVar[TripMode]
TRIP_MODE_UNSPECIFIED: TripMode
TRIP_MODE_WALK_ONLY: TripMode
TRIP_MODE_DRIVE_ONLY: TripMode
TRIP_MODE_BUS_WALK: TripMode
TRIP_MODE_BIKE_WALK: TripMode
TRIP_MODE_TAXI: TripMode
TRIP_MODE_SUBWAY_WALK: TripMode
TRIP_MODE_BUS_SUBWAY_WALK: TripMode

class TripStop(_message.Message):
    __slots__ = ['aoi_position', 'lane_position', 'duration', 'optional_lane_positions']
    AOI_POSITION_FIELD_NUMBER: _ClassVar[int]
    LANE_POSITION_FIELD_NUMBER: _ClassVar[int]
    DURATION_FIELD_NUMBER: _ClassVar[int]
    OPTIONAL_LANE_POSITIONS_FIELD_NUMBER: _ClassVar[int]
    aoi_position: _geo_pb2.AoiPosition
    lane_position: _geo_pb2.LanePosition
    duration: float
    optional_lane_positions: _containers.RepeatedCompositeFieldContainer[_geo_pb2.LanePosition]

    def __init__(self, aoi_position: _Optional[_Union[_geo_pb2.AoiPosition, _Mapping]]=..., lane_position: _Optional[_Union[_geo_pb2.LanePosition, _Mapping]]=..., duration: _Optional[float]=..., optional_lane_positions: _Optional[_Iterable[_Union[_geo_pb2.LanePosition, _Mapping]]]=...) -> None:
        ...

class Trip(_message.Message):
    __slots__ = ['mode', 'end', 'departure_time', 'wait_time', 'arrival_time', 'activity', 'model', 'routes', 'trip_stops']
    MODE_FIELD_NUMBER: _ClassVar[int]
    END_FIELD_NUMBER: _ClassVar[int]
    DEPARTURE_TIME_FIELD_NUMBER: _ClassVar[int]
    WAIT_TIME_FIELD_NUMBER: _ClassVar[int]
    ARRIVAL_TIME_FIELD_NUMBER: _ClassVar[int]
    ACTIVITY_FIELD_NUMBER: _ClassVar[int]
    MODEL_FIELD_NUMBER: _ClassVar[int]
    ROUTES_FIELD_NUMBER: _ClassVar[int]
    TRIP_STOPS_FIELD_NUMBER: _ClassVar[int]
    mode: TripMode
    end: _geo_pb2.Position
    departure_time: float
    wait_time: float
    arrival_time: float
    activity: str
    model: str
    routes: _containers.RepeatedCompositeFieldContainer[_routing_pb2.Journey]
    trip_stops: _containers.RepeatedCompositeFieldContainer[TripStop]

    def __init__(self, mode: _Optional[_Union[TripMode, str]]=..., end: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., departure_time: _Optional[float]=..., wait_time: _Optional[float]=..., arrival_time: _Optional[float]=..., activity: _Optional[str]=..., model: _Optional[str]=..., routes: _Optional[_Iterable[_Union[_routing_pb2.Journey, _Mapping]]]=..., trip_stops: _Optional[_Iterable[_Union[TripStop, _Mapping]]]=...) -> None:
        ...

class Schedule(_message.Message):
    __slots__ = ['trips', 'loop_count', 'departure_time', 'wait_time']
    TRIPS_FIELD_NUMBER: _ClassVar[int]
    LOOP_COUNT_FIELD_NUMBER: _ClassVar[int]
    DEPARTURE_TIME_FIELD_NUMBER: _ClassVar[int]
    WAIT_TIME_FIELD_NUMBER: _ClassVar[int]
    trips: _containers.RepeatedCompositeFieldContainer[Trip]
    loop_count: int
    departure_time: float
    wait_time: float

    def __init__(self, trips: _Optional[_Iterable[_Union[Trip, _Mapping]]]=..., loop_count: _Optional[int]=..., departure_time: _Optional[float]=..., wait_time: _Optional[float]=...) -> None:
        ...