from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class RouteType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    ROUTE_TYPE_UNSPECIFIED: _ClassVar[RouteType]
    ROUTE_TYPE_DRIVING: _ClassVar[RouteType]
    ROUTE_TYPE_WALKING: _ClassVar[RouteType]
    ROUTE_TYPE_BY_BUS: _ClassVar[RouteType]

class JourneyType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    JOURNEY_TYPE_UNSPECIFIED: _ClassVar[JourneyType]
    JOURNEY_TYPE_DRIVING: _ClassVar[JourneyType]
    JOURNEY_TYPE_WALKING: _ClassVar[JourneyType]
    JOURNEY_TYPE_BY_BUS: _ClassVar[JourneyType]

class MovingDirection(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    MOVING_DIRECTION_UNSPECIFIED: _ClassVar[MovingDirection]
    MOVING_DIRECTION_FORWARD: _ClassVar[MovingDirection]
    MOVING_DIRECTION_BACKWARD: _ClassVar[MovingDirection]
ROUTE_TYPE_UNSPECIFIED: RouteType
ROUTE_TYPE_DRIVING: RouteType
ROUTE_TYPE_WALKING: RouteType
ROUTE_TYPE_BY_BUS: RouteType
JOURNEY_TYPE_UNSPECIFIED: JourneyType
JOURNEY_TYPE_DRIVING: JourneyType
JOURNEY_TYPE_WALKING: JourneyType
JOURNEY_TYPE_BY_BUS: JourneyType
MOVING_DIRECTION_UNSPECIFIED: MovingDirection
MOVING_DIRECTION_FORWARD: MovingDirection
MOVING_DIRECTION_BACKWARD: MovingDirection

class DrivingJourneyBody(_message.Message):
    __slots__ = ['road_ids', 'eta']
    ROAD_IDS_FIELD_NUMBER: _ClassVar[int]
    ETA_FIELD_NUMBER: _ClassVar[int]
    road_ids: _containers.RepeatedScalarFieldContainer[int]
    eta: float

    def __init__(self, road_ids: _Optional[_Iterable[int]]=..., eta: _Optional[float]=...) -> None:
        ...

class WalkingRouteSegment(_message.Message):
    __slots__ = ['lane_id', 'moving_direction']
    LANE_ID_FIELD_NUMBER: _ClassVar[int]
    MOVING_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    lane_id: int
    moving_direction: MovingDirection

    def __init__(self, lane_id: _Optional[int]=..., moving_direction: _Optional[_Union[MovingDirection, str]]=...) -> None:
        ...

class WalkingJourneyBody(_message.Message):
    __slots__ = ['route', 'eta']
    ROUTE_FIELD_NUMBER: _ClassVar[int]
    ETA_FIELD_NUMBER: _ClassVar[int]
    route: _containers.RepeatedCompositeFieldContainer[WalkingRouteSegment]
    eta: float

    def __init__(self, route: _Optional[_Iterable[_Union[WalkingRouteSegment, _Mapping]]]=..., eta: _Optional[float]=...) -> None:
        ...

class TransferSegment(_message.Message):
    __slots__ = ['subline_id', 'start_station_id', 'end_station_id']
    SUBLINE_ID_FIELD_NUMBER: _ClassVar[int]
    START_STATION_ID_FIELD_NUMBER: _ClassVar[int]
    END_STATION_ID_FIELD_NUMBER: _ClassVar[int]
    subline_id: int
    start_station_id: int
    end_station_id: int

    def __init__(self, subline_id: _Optional[int]=..., start_station_id: _Optional[int]=..., end_station_id: _Optional[int]=...) -> None:
        ...

class BusJourneyBody(_message.Message):
    __slots__ = ['transfers', 'eta']
    TRANSFERS_FIELD_NUMBER: _ClassVar[int]
    ETA_FIELD_NUMBER: _ClassVar[int]
    transfers: _containers.RepeatedCompositeFieldContainer[TransferSegment]
    eta: float

    def __init__(self, transfers: _Optional[_Iterable[_Union[TransferSegment, _Mapping]]]=..., eta: _Optional[float]=...) -> None:
        ...

class Journey(_message.Message):
    __slots__ = ['type', 'driving', 'walking', 'by_bus']
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DRIVING_FIELD_NUMBER: _ClassVar[int]
    WALKING_FIELD_NUMBER: _ClassVar[int]
    BY_BUS_FIELD_NUMBER: _ClassVar[int]
    type: JourneyType
    driving: DrivingJourneyBody
    walking: WalkingJourneyBody
    by_bus: BusJourneyBody

    def __init__(self, type: _Optional[_Union[JourneyType, str]]=..., driving: _Optional[_Union[DrivingJourneyBody, _Mapping]]=..., walking: _Optional[_Union[WalkingJourneyBody, _Mapping]]=..., by_bus: _Optional[_Union[BusJourneyBody, _Mapping]]=...) -> None:
        ...

class RoadStatus(_message.Message):
    __slots__ = ['id', 'speed']
    ID_FIELD_NUMBER: _ClassVar[int]
    SPEED_FIELD_NUMBER: _ClassVar[int]
    id: int
    speed: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, id: _Optional[int]=..., speed: _Optional[_Iterable[float]]=...) -> None:
        ...

class RoadStatuses(_message.Message):
    __slots__ = ['road_statuses']
    ROAD_STATUSES_FIELD_NUMBER: _ClassVar[int]
    road_statuses: _containers.RepeatedCompositeFieldContainer[RoadStatus]

    def __init__(self, road_statuses: _Optional[_Iterable[_Union[RoadStatus, _Mapping]]]=...) -> None:
        ...