from city.map.v2 import light_pb2 as _light_pb2
from city.person.v1 import motion_pb2 as _motion_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

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
    light_state: _light_pb2.LightState

    def __init__(self, id: _Optional[int]=..., persons: _Optional[_Iterable[_Union[_motion_pb2.PersonMotion, _Mapping]]]=..., avg_v: _Optional[float]=..., restriction: bool=..., light_state: _Optional[_Union[_light_pb2.LightState, str]]=...) -> None:
        ...