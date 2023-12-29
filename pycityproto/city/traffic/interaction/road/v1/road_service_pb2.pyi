from city.event.v1 import event_pb2 as _event_pb2
from city.traffic.interaction.road.v1 import road_pb2 as _road_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetRoadRequest(_message.Message):
    __slots__ = ('road_ids',)
    ROAD_IDS_FIELD_NUMBER: _ClassVar[int]
    road_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, road_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetRoadResponse(_message.Message):
    __slots__ = ('states',)
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedCompositeFieldContainer[_road_pb2.State]

    def __init__(self, states: _Optional[_Iterable[_Union[_road_pb2.State, _Mapping]]]=...) -> None:
        ...

class GetRuinInfoRequest(_message.Message):
    __slots__ = ()

    def __init__(self) -> None:
        ...

class RuinInfo(_message.Message):
    __slots__ = ('num', 'ratio')
    NUM_FIELD_NUMBER: _ClassVar[int]
    RATIO_FIELD_NUMBER: _ClassVar[int]
    num: int
    ratio: float

    def __init__(self, num: _Optional[int]=..., ratio: _Optional[float]=...) -> None:
        ...

class GetRuinInfoResponse(_message.Message):
    __slots__ = ('one', 'two', 'three')
    ONE_FIELD_NUMBER: _ClassVar[int]
    TWO_FIELD_NUMBER: _ClassVar[int]
    THREE_FIELD_NUMBER: _ClassVar[int]
    one: RuinInfo
    two: RuinInfo
    three: RuinInfo

    def __init__(self, one: _Optional[_Union[RuinInfo, _Mapping]]=..., two: _Optional[_Union[RuinInfo, _Mapping]]=..., three: _Optional[_Union[RuinInfo, _Mapping]]=...) -> None:
        ...

class GetEventsRequest(_message.Message):
    __slots__ = ()

    def __init__(self) -> None:
        ...

class GetEventsResponse(_message.Message):
    __slots__ = ('events',)
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _event_pb2.Events

    def __init__(self, events: _Optional[_Union[_event_pb2.Events, _Mapping]]=...) -> None:
        ...