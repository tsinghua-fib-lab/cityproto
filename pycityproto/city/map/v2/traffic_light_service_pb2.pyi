from city.map.v2 import light_pb2 as _light_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetTrafficLightRequest(_message.Message):
    __slots__ = ['junction_id']
    JUNCTION_ID_FIELD_NUMBER: _ClassVar[int]
    junction_id: int

    def __init__(self, junction_id: _Optional[int]=...) -> None:
        ...

class GetTrafficLightResponse(_message.Message):
    __slots__ = ['traffic_light', 'phase_index', 'time_remaining']
    TRAFFIC_LIGHT_FIELD_NUMBER: _ClassVar[int]
    PHASE_INDEX_FIELD_NUMBER: _ClassVar[int]
    TIME_REMAINING_FIELD_NUMBER: _ClassVar[int]
    traffic_light: _light_pb2.TrafficLight
    phase_index: int
    time_remaining: float

    def __init__(self, traffic_light: _Optional[_Union[_light_pb2.TrafficLight, _Mapping]]=..., phase_index: _Optional[int]=..., time_remaining: _Optional[float]=...) -> None:
        ...

class SetTrafficLightRequest(_message.Message):
    __slots__ = ['traffic_light', 'phase_index', 'time_remaining']
    TRAFFIC_LIGHT_FIELD_NUMBER: _ClassVar[int]
    PHASE_INDEX_FIELD_NUMBER: _ClassVar[int]
    TIME_REMAINING_FIELD_NUMBER: _ClassVar[int]
    traffic_light: _light_pb2.TrafficLight
    phase_index: int
    time_remaining: float

    def __init__(self, traffic_light: _Optional[_Union[_light_pb2.TrafficLight, _Mapping]]=..., phase_index: _Optional[int]=..., time_remaining: _Optional[float]=...) -> None:
        ...

class SetTrafficLightResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class SetTrafficLightPhaseRequest(_message.Message):
    __slots__ = ['junction_id', 'phase_index', 'time_remaining']
    JUNCTION_ID_FIELD_NUMBER: _ClassVar[int]
    PHASE_INDEX_FIELD_NUMBER: _ClassVar[int]
    TIME_REMAINING_FIELD_NUMBER: _ClassVar[int]
    junction_id: int
    phase_index: int
    time_remaining: float

    def __init__(self, junction_id: _Optional[int]=..., phase_index: _Optional[int]=..., time_remaining: _Optional[float]=...) -> None:
        ...

class SetTrafficLightPhaseResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class SetTrafficLightStatusRequest(_message.Message):
    __slots__ = ['junction_id', 'ok']
    JUNCTION_ID_FIELD_NUMBER: _ClassVar[int]
    OK_FIELD_NUMBER: _ClassVar[int]
    junction_id: int
    ok: bool

    def __init__(self, junction_id: _Optional[int]=..., ok: bool=...) -> None:
        ...

class SetTrafficLightStatusResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...