from city.traffic.interaction.lane.v1 import lane_pb2 as _lane_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class SetMaxVRequest(_message.Message):
    __slots__ = ['lane_id', 'max_v']
    LANE_ID_FIELD_NUMBER: _ClassVar[int]
    MAX_V_FIELD_NUMBER: _ClassVar[int]
    lane_id: int
    max_v: float

    def __init__(self, lane_id: _Optional[int]=..., max_v: _Optional[float]=...) -> None:
        ...

class SetMaxVResponse(_message.Message):
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
    states: _containers.RepeatedCompositeFieldContainer[_lane_pb2.State]

    def __init__(self, states: _Optional[_Iterable[_Union[_lane_pb2.State, _Mapping]]]=...) -> None:
        ...