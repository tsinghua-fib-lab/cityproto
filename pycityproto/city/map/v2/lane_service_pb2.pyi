from city.geo.v2 import geo_pb2 as _geo_pb2
from city.map.v2 import lane_state_pb2 as _lane_state_pb2
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

class SetLaneRestrictionRequest(_message.Message):
    __slots__ = ['lane_id', 'restriction']
    LANE_ID_FIELD_NUMBER: _ClassVar[int]
    RESTRICTION_FIELD_NUMBER: _ClassVar[int]
    lane_id: int
    restriction: bool

    def __init__(self, lane_id: _Optional[int]=..., restriction: bool=...) -> None:
        ...

class SetLaneRestrictionResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetLaneRequest(_message.Message):
    __slots__ = ['lane_ids', 'exclude_person']
    LANE_IDS_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_PERSON_FIELD_NUMBER: _ClassVar[int]
    lane_ids: _containers.RepeatedScalarFieldContainer[int]
    exclude_person: bool

    def __init__(self, lane_ids: _Optional[_Iterable[int]]=..., exclude_person: bool=...) -> None:
        ...

class GetLaneResponse(_message.Message):
    __slots__ = ['states']
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedCompositeFieldContainer[_lane_state_pb2.LaneState]

    def __init__(self, states: _Optional[_Iterable[_Union[_lane_state_pb2.LaneState, _Mapping]]]=...) -> None:
        ...

class GetLaneByLongLatBBoxRequest(_message.Message):
    __slots__ = ['bbox', 'exclude_person']
    BBOX_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_PERSON_FIELD_NUMBER: _ClassVar[int]
    bbox: _geo_pb2.LongLatBBox
    exclude_person: bool

    def __init__(self, bbox: _Optional[_Union[_geo_pb2.LongLatBBox, _Mapping]]=..., exclude_person: bool=...) -> None:
        ...

class GetLaneByLongLatBBoxResponse(_message.Message):
    __slots__ = ['states']
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedCompositeFieldContainer[_lane_state_pb2.LaneState]

    def __init__(self, states: _Optional[_Iterable[_Union[_lane_state_pb2.LaneState, _Mapping]]]=...) -> None:
        ...