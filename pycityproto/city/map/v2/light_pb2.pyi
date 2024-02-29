from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class LightState(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    LIGHT_STATE_UNSPECIFIED: _ClassVar[LightState]
    LIGHT_STATE_RED: _ClassVar[LightState]
    LIGHT_STATE_GREEN: _ClassVar[LightState]
    LIGHT_STATE_YELLOW: _ClassVar[LightState]
LIGHT_STATE_UNSPECIFIED: LightState
LIGHT_STATE_RED: LightState
LIGHT_STATE_GREEN: LightState
LIGHT_STATE_YELLOW: LightState

class Phase(_message.Message):
    __slots__ = ['duration', 'states']
    DURATION_FIELD_NUMBER: _ClassVar[int]
    STATES_FIELD_NUMBER: _ClassVar[int]
    duration: float
    states: _containers.RepeatedScalarFieldContainer[LightState]

    def __init__(self, duration: _Optional[float]=..., states: _Optional[_Iterable[_Union[LightState, str]]]=...) -> None:
        ...

class AvailablePhase(_message.Message):
    __slots__ = ['states']
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedScalarFieldContainer[LightState]

    def __init__(self, states: _Optional[_Iterable[_Union[LightState, str]]]=...) -> None:
        ...

class TrafficLight(_message.Message):
    __slots__ = ['junction_id', 'phases']
    JUNCTION_ID_FIELD_NUMBER: _ClassVar[int]
    PHASES_FIELD_NUMBER: _ClassVar[int]
    junction_id: int
    phases: _containers.RepeatedCompositeFieldContainer[Phase]

    def __init__(self, junction_id: _Optional[int]=..., phases: _Optional[_Iterable[_Union[Phase, _Mapping]]]=...) -> None:
        ...