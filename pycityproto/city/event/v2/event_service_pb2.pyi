from city.event.v2 import event_pb2 as _event_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetEventsByTopicRequest(_message.Message):
    __slots__ = ['topic']
    TOPIC_FIELD_NUMBER: _ClassVar[int]
    topic: str

    def __init__(self, topic: _Optional[str]=...) -> None:
        ...

class GetEventsByTopicResponse(_message.Message):
    __slots__ = ['events']
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _containers.RepeatedCompositeFieldContainer[_event_pb2.Event]

    def __init__(self, events: _Optional[_Iterable[_Union[_event_pb2.Event, _Mapping]]]=...) -> None:
        ...

class ResolveEventsRequest(_message.Message):
    __slots__ = ['events']
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _containers.RepeatedCompositeFieldContainer[_event_pb2.Event]

    def __init__(self, events: _Optional[_Iterable[_Union[_event_pb2.Event, _Mapping]]]=...) -> None:
        ...

class ResolveEventsResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...