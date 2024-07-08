from city.event.v2 import event_pb2 as _event_pb2
from city.person.v1 import motion_pb2 as _motion_pb2
from city.person.v1 import person_pb2 as _person_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class PersonRuntime(_message.Message):
    __slots__ = ['base', 'motion', 'events']
    BASE_FIELD_NUMBER: _ClassVar[int]
    MOTION_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    base: _person_pb2.Person
    motion: _motion_pb2.PersonMotion
    events: _containers.RepeatedCompositeFieldContainer[_event_pb2.Event]

    def __init__(self, base: _Optional[_Union[_person_pb2.Person, _Mapping]]=..., motion: _Optional[_Union[_motion_pb2.PersonMotion, _Mapping]]=..., events: _Optional[_Iterable[_Union[_event_pb2.Event, _Mapping]]]=...) -> None:
        ...