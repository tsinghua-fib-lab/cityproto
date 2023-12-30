from city.event.v1 import event_pb2 as _event_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class PublishRequest(_message.Message):
    __slots__ = ['event']
    EVENT_FIELD_NUMBER: _ClassVar[int]
    event: _event_pb2.Event

    def __init__(self, event: _Optional[_Union[_event_pb2.Event, _Mapping]]=...) -> None:
        ...

class PublishResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class PullRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class PullResponse(_message.Message):
    __slots__ = ['events']
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _containers.RepeatedCompositeFieldContainer[_event_pb2.Event]

    def __init__(self, events: _Optional[_Iterable[_Union[_event_pb2.Event, _Mapping]]]=...) -> None:
        ...