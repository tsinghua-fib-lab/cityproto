from city.person.v2 import motion_pb2 as _motion_pb2
from city.routing.v2 import routing_pb2 as _routing_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class PedestrianEnv(_message.Message):
    __slots__ = ['id', 'motion', 'journey', 'is_current_lane_no_entry', 'is_next_lane_no_entry']
    ID_FIELD_NUMBER: _ClassVar[int]
    MOTION_FIELD_NUMBER: _ClassVar[int]
    JOURNEY_FIELD_NUMBER: _ClassVar[int]
    IS_CURRENT_LANE_NO_ENTRY_FIELD_NUMBER: _ClassVar[int]
    IS_NEXT_LANE_NO_ENTRY_FIELD_NUMBER: _ClassVar[int]
    id: int
    motion: _motion_pb2.PersonMotion
    journey: _routing_pb2.Journey
    is_current_lane_no_entry: bool
    is_next_lane_no_entry: bool

    def __init__(self, id: _Optional[int]=..., motion: _Optional[_Union[_motion_pb2.PersonMotion, _Mapping]]=..., journey: _Optional[_Union[_routing_pb2.Journey, _Mapping]]=..., is_current_lane_no_entry: bool=..., is_next_lane_no_entry: bool=...) -> None:
        ...

class PedestrianAction(_message.Message):
    __slots__ = ['id', 'vx', 'vy']
    ID_FIELD_NUMBER: _ClassVar[int]
    VX_FIELD_NUMBER: _ClassVar[int]
    VY_FIELD_NUMBER: _ClassVar[int]
    id: int
    vx: float
    vy: float

    def __init__(self, id: _Optional[int]=..., vx: _Optional[float]=..., vy: _Optional[float]=...) -> None:
        ...