from city.geo.v2 import geo_pb2 as _geo_pb2
from city.person.v1 import motion_pb2 as _motion_pb2
from city.person.v1 import person_pb2 as _person_pb2
from city.trip.v2 import trip_pb2 as _trip_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetPersonRequest(_message.Message):
    __slots__ = ['person_id']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    person_id: int

    def __init__(self, person_id: _Optional[int]=...) -> None:
        ...

class GetPersonResponse(_message.Message):
    __slots__ = ['base', 'motion']
    BASE_FIELD_NUMBER: _ClassVar[int]
    MOTION_FIELD_NUMBER: _ClassVar[int]
    base: _person_pb2.Person
    motion: _motion_pb2.PersonMotion

    def __init__(self, base: _Optional[_Union[_person_pb2.Person, _Mapping]]=..., motion: _Optional[_Union[_motion_pb2.PersonMotion, _Mapping]]=...) -> None:
        ...

class AddPersonRequest(_message.Message):
    __slots__ = ['person']
    PERSON_FIELD_NUMBER: _ClassVar[int]
    person: _person_pb2.Person

    def __init__(self, person: _Optional[_Union[_person_pb2.Person, _Mapping]]=...) -> None:
        ...

class AddPersonResponse(_message.Message):
    __slots__ = ['person_id']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    person_id: int

    def __init__(self, person_id: _Optional[int]=...) -> None:
        ...

class SetScheduleRequest(_message.Message):
    __slots__ = ['person_id', 'schedules']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    SCHEDULES_FIELD_NUMBER: _ClassVar[int]
    person_id: int
    schedules: _containers.RepeatedCompositeFieldContainer[_trip_pb2.Schedule]

    def __init__(self, person_id: _Optional[int]=..., schedules: _Optional[_Iterable[_Union[_trip_pb2.Schedule, _Mapping]]]=...) -> None:
        ...

class SetScheduleResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetPersonByLongLatBBoxRequest(_message.Message):
    __slots__ = ['bbox', 'exclude_statuses']
    BBOX_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_STATUSES_FIELD_NUMBER: _ClassVar[int]
    bbox: _geo_pb2.LongLatBBox
    exclude_statuses: _containers.RepeatedScalarFieldContainer[_motion_pb2.Status]

    def __init__(self, bbox: _Optional[_Union[_geo_pb2.LongLatBBox, _Mapping]]=..., exclude_statuses: _Optional[_Iterable[_Union[_motion_pb2.Status, str]]]=...) -> None:
        ...

class GetPersonByLongLatBBoxResponse(_message.Message):
    __slots__ = ['motions']
    MOTIONS_FIELD_NUMBER: _ClassVar[int]
    motions: _containers.RepeatedCompositeFieldContainer[_motion_pb2.PersonMotion]

    def __init__(self, motions: _Optional[_Iterable[_Union[_motion_pb2.PersonMotion, _Mapping]]]=...) -> None:
        ...