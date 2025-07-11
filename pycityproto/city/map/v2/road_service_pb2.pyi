from city.event.v1 import event_pb2 as _event_pb2
from city.map.v2 import lane_state_pb2 as _lane_state_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class RoadLevel(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    ROAD_LEVEL_UNSPECIFIED: _ClassVar[RoadLevel]
    ROAD_LEVEL_CLEAR: _ClassVar[RoadLevel]
    ROAD_LEVEL_LIGHT_LOAD: _ClassVar[RoadLevel]
    ROAD_LEVEL_MEDIUM_LOAD: _ClassVar[RoadLevel]
    ROAD_LEVEL_HEAVY_LOAD: _ClassVar[RoadLevel]
    ROAD_LEVEL_OVERLOAD: _ClassVar[RoadLevel]
    ROAD_LEVEL_RESTRICTED: _ClassVar[RoadLevel]

class InterruptionReason(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    INTERRUPTION_REASON_UNSPECIFIED: _ClassVar[InterruptionReason]
    INTERRUPTION_REASON_RUINED: _ClassVar[InterruptionReason]
    INTERRUPTION_REASON_CASCADE: _ClassVar[InterruptionReason]
    INTERRUPTION_REASON_CONGESTION: _ClassVar[InterruptionReason]
ROAD_LEVEL_UNSPECIFIED: RoadLevel
ROAD_LEVEL_CLEAR: RoadLevel
ROAD_LEVEL_LIGHT_LOAD: RoadLevel
ROAD_LEVEL_MEDIUM_LOAD: RoadLevel
ROAD_LEVEL_HEAVY_LOAD: RoadLevel
ROAD_LEVEL_OVERLOAD: RoadLevel
ROAD_LEVEL_RESTRICTED: RoadLevel
INTERRUPTION_REASON_UNSPECIFIED: InterruptionReason
INTERRUPTION_REASON_RUINED: InterruptionReason
INTERRUPTION_REASON_CASCADE: InterruptionReason
INTERRUPTION_REASON_CONGESTION: InterruptionReason

class GetRoadRequest(_message.Message):
    __slots__ = ['road_ids', 'exclude_lane', 'exclude_person']
    ROAD_IDS_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_LANE_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_PERSON_FIELD_NUMBER: _ClassVar[int]
    road_ids: _containers.RepeatedScalarFieldContainer[int]
    exclude_lane: bool
    exclude_person: bool

    def __init__(self, road_ids: _Optional[_Iterable[int]]=..., exclude_lane: bool=..., exclude_person: bool=...) -> None:
        ...

class GetRoadResponse(_message.Message):
    __slots__ = ['states']
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedCompositeFieldContainer[RoadState]

    def __init__(self, states: _Optional[_Iterable[_Union[RoadState, _Mapping]]]=...) -> None:
        ...

class GetRoadGlobalStatisticsRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetRoadGlobalStatisticsResponse(_message.Message):
    __slots__ = ['avg_road_congestion_index']
    AVG_ROAD_CONGESTION_INDEX_FIELD_NUMBER: _ClassVar[int]
    avg_road_congestion_index: float

    def __init__(self, avg_road_congestion_index: _Optional[float]=...) -> None:
        ...

class GetRuinInfoRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RuinInfo(_message.Message):
    __slots__ = ['num', 'ratio']
    NUM_FIELD_NUMBER: _ClassVar[int]
    RATIO_FIELD_NUMBER: _ClassVar[int]
    num: int
    ratio: float

    def __init__(self, num: _Optional[int]=..., ratio: _Optional[float]=...) -> None:
        ...

class GetRuinInfoResponse(_message.Message):
    __slots__ = ['one', 'two', 'three']
    ONE_FIELD_NUMBER: _ClassVar[int]
    TWO_FIELD_NUMBER: _ClassVar[int]
    THREE_FIELD_NUMBER: _ClassVar[int]
    one: RuinInfo
    two: RuinInfo
    three: RuinInfo

    def __init__(self, one: _Optional[_Union[RuinInfo, _Mapping]]=..., two: _Optional[_Union[RuinInfo, _Mapping]]=..., three: _Optional[_Union[RuinInfo, _Mapping]]=...) -> None:
        ...

class GetEventsRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetEventsResponse(_message.Message):
    __slots__ = ['events']
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _event_pb2.Events

    def __init__(self, events: _Optional[_Union[_event_pb2.Events, _Mapping]]=...) -> None:
        ...

class RoadState(_message.Message):
    __slots__ = ['id', 'in_vehicle_cnt', 'out_vehicle_cnt', 'vehicle_cnt', 'cum_in_vehicle_cnt', 'cum_out_vehicle_cnt', 'avg_v', 'avg_travel_time', 'level', 'congestion_index', 'reason', 'lanes', 'max_v']
    ID_FIELD_NUMBER: _ClassVar[int]
    IN_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    OUT_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    CUM_IN_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    CUM_OUT_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    AVG_V_FIELD_NUMBER: _ClassVar[int]
    AVG_TRAVEL_TIME_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    CONGESTION_INDEX_FIELD_NUMBER: _ClassVar[int]
    REASON_FIELD_NUMBER: _ClassVar[int]
    LANES_FIELD_NUMBER: _ClassVar[int]
    MAX_V_FIELD_NUMBER: _ClassVar[int]
    id: int
    in_vehicle_cnt: int
    out_vehicle_cnt: int
    vehicle_cnt: int
    cum_in_vehicle_cnt: int
    cum_out_vehicle_cnt: int
    avg_v: float
    avg_travel_time: float
    level: RoadLevel
    congestion_index: float
    reason: InterruptionReason
    lanes: _containers.RepeatedCompositeFieldContainer[_lane_state_pb2.LaneState]
    max_v: float

    def __init__(self, id: _Optional[int]=..., in_vehicle_cnt: _Optional[int]=..., out_vehicle_cnt: _Optional[int]=..., vehicle_cnt: _Optional[int]=..., cum_in_vehicle_cnt: _Optional[int]=..., cum_out_vehicle_cnt: _Optional[int]=..., avg_v: _Optional[float]=..., avg_travel_time: _Optional[float]=..., level: _Optional[_Union[RoadLevel, str]]=..., congestion_index: _Optional[float]=..., reason: _Optional[_Union[InterruptionReason, str]]=..., lanes: _Optional[_Iterable[_Union[_lane_state_pb2.LaneState, _Mapping]]]=..., max_v: _Optional[float]=...) -> None:
        ...