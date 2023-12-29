from city.traffic.interaction.lane.v1 import lane_pb2 as _lane_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class RoadLevel(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ROAD_LEVEL_UNSPECIFIED: _ClassVar[RoadLevel]
    ROAD_LEVEL_CLEAR: _ClassVar[RoadLevel]
    ROAD_LEVEL_LIGHT_LOAD: _ClassVar[RoadLevel]
    ROAD_LEVEL_MEDIUM_LOAD: _ClassVar[RoadLevel]
    ROAD_LEVEL_HEAVY_LOAD: _ClassVar[RoadLevel]
    ROAD_LEVEL_OVERLOAD: _ClassVar[RoadLevel]
    ROAD_LEVEL_RESTRICTED: _ClassVar[RoadLevel]

class InterruptionReason(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
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

class State(_message.Message):
    __slots__ = ('id', 'avg_v', 'level', 'reason', 'lanes')
    ID_FIELD_NUMBER: _ClassVar[int]
    AVG_V_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    REASON_FIELD_NUMBER: _ClassVar[int]
    LANES_FIELD_NUMBER: _ClassVar[int]
    id: int
    avg_v: float
    level: RoadLevel
    reason: InterruptionReason
    lanes: _containers.RepeatedCompositeFieldContainer[_lane_pb2.State]

    def __init__(self, id: _Optional[int]=..., avg_v: _Optional[float]=..., level: _Optional[_Union[RoadLevel, str]]=..., reason: _Optional[_Union[InterruptionReason, str]]=..., lanes: _Optional[_Iterable[_Union[_lane_pb2.State, _Mapping]]]=...) -> None:
        ...