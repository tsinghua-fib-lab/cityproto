from city.person.v1 import motion_pb2 as _motion_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class SetLaneMaxVRequest(_message.Message):
    __slots__ = ['lane_id', 'max_v']
    LANE_ID_FIELD_NUMBER: _ClassVar[int]
    MAX_V_FIELD_NUMBER: _ClassVar[int]
    lane_id: int
    max_v: float

    def __init__(self, lane_id: _Optional[int]=..., max_v: _Optional[float]=...) -> None:
        ...

class SetLaneMaxVResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetLaneRequest(_message.Message):
    __slots__ = ['lane_ids']
    LANE_IDS_FIELD_NUMBER: _ClassVar[int]
    lane_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, lane_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetLaneResponse(_message.Message):
    __slots__ = ['states']
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedCompositeFieldContainer[LaneState]

    def __init__(self, states: _Optional[_Iterable[_Union[LaneState, _Mapping]]]=...) -> None:
        ...

class LaneState(_message.Message):
    __slots__ = ['id', 'persons', 'avg_v', 'restriction']
    ID_FIELD_NUMBER: _ClassVar[int]
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    AVG_V_FIELD_NUMBER: _ClassVar[int]
    RESTRICTION_FIELD_NUMBER: _ClassVar[int]
    id: int
    persons: _containers.RepeatedCompositeFieldContainer[_motion_pb2.PersonMotion]
    avg_v: float
    restriction: bool

    def __init__(self, id: _Optional[int]=..., persons: _Optional[_Iterable[_Union[_motion_pb2.PersonMotion, _Mapping]]]=..., avg_v: _Optional[float]=..., restriction: bool=...) -> None:
        ...