from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class EventType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    EVENT_TYPE_UNSPECIFIED: _ClassVar[EventType]
    EVENT_TYPE_HEAVY_RAIN: _ClassVar[EventType]
    EVENT_TYPE_MILITARY_STRIKE: _ClassVar[EventType]
    EVENT_TYPE_TRAFFIC_CONGESTION: _ClassVar[EventType]
    EVENT_TYPE_TRAFFIC_LANE_RESTRICTION: _ClassVar[EventType]
    EVENT_TYPE_TRAFFIC_BAD_LIGHT: _ClassVar[EventType]
    EVENT_TYPE_ELEC_RUINED_TRANSFORMER: _ClassVar[EventType]
    EVENT_TYPE_WATER_RUINED_PUMP: _ClassVar[EventType]
    EVENT_TYPE_WATER_STOPPED_PUMP: _ClassVar[EventType]
    EVENT_TYPE_COMM_RUINED_BASE_STATION: _ClassVar[EventType]
    EVENT_TYPE_COMM_STOPPED_BASE_STATION: _ClassVar[EventType]
    EVENT_TYPE_COMM_OVERLOAD_BASE_STATION: _ClassVar[EventType]
EVENT_TYPE_UNSPECIFIED: EventType
EVENT_TYPE_HEAVY_RAIN: EventType
EVENT_TYPE_MILITARY_STRIKE: EventType
EVENT_TYPE_TRAFFIC_CONGESTION: EventType
EVENT_TYPE_TRAFFIC_LANE_RESTRICTION: EventType
EVENT_TYPE_TRAFFIC_BAD_LIGHT: EventType
EVENT_TYPE_ELEC_RUINED_TRANSFORMER: EventType
EVENT_TYPE_WATER_RUINED_PUMP: EventType
EVENT_TYPE_WATER_STOPPED_PUMP: EventType
EVENT_TYPE_COMM_RUINED_BASE_STATION: EventType
EVENT_TYPE_COMM_STOPPED_BASE_STATION: EventType
EVENT_TYPE_COMM_OVERLOAD_BASE_STATION: EventType

class Event(_message.Message):
    __slots__ = ['type', 'level', 'step']
    TYPE_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    STEP_FIELD_NUMBER: _ClassVar[int]
    type: EventType
    level: int
    step: int

    def __init__(self, type: _Optional[_Union[EventType, str]]=..., level: _Optional[int]=..., step: _Optional[int]=...) -> None:
        ...

class Events(_message.Message):
    __slots__ = ['events']
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _containers.RepeatedCompositeFieldContainer[Event]

    def __init__(self, events: _Optional[_Iterable[_Union[Event, _Mapping]]]=...) -> None:
        ...