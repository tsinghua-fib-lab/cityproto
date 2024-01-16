from city.geo.v2 import geo_pb2 as _geo_pb2
from city.person.v1 import motion_pb2 as _motion_pb2
from city.traffic_light.v2 import traffic_light_pb2 as _traffic_light_pb2
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
    states: _containers.RepeatedCompositeFieldContainer[LaneState]

    def __init__(self, states: _Optional[_Iterable[_Union[LaneState, _Mapping]]]=...) -> None:
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
    states: _containers.RepeatedCompositeFieldContainer[LaneState]

    def __init__(self, states: _Optional[_Iterable[_Union[LaneState, _Mapping]]]=...) -> None:
        ...

class LaneState(_message.Message):
    __slots__ = ['id', 'persons', 'avg_v', 'restriction', 'light_state']
    ID_FIELD_NUMBER: _ClassVar[int]
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    AVG_V_FIELD_NUMBER: _ClassVar[int]
    RESTRICTION_FIELD_NUMBER: _ClassVar[int]
    LIGHT_STATE_FIELD_NUMBER: _ClassVar[int]
    id: int
    persons: _containers.RepeatedCompositeFieldContainer[_motion_pb2.PersonMotion]
    avg_v: float
    restriction: bool
    light_state: _traffic_light_pb2.LightState

    def __init__(self, id: _Optional[int]=..., persons: _Optional[_Iterable[_Union[_motion_pb2.PersonMotion, _Mapping]]]=..., avg_v: _Optional[float]=..., restriction: bool=..., light_state: _Optional[_Union[_traffic_light_pb2.LightState, str]]=...) -> None:
        ...