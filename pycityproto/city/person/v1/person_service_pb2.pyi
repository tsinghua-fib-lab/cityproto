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

class GetPersonsByLongLatAreaRequest(_message.Message):
    __slots__ = ['area']
    AREA_FIELD_NUMBER: _ClassVar[int]
    area: _geo_pb2.LongLatRectArea

    def __init__(self, area: _Optional[_Union[_geo_pb2.LongLatRectArea, _Mapping]]=...) -> None:
        ...

class GetPersonsByLongLatAreaResponse(_message.Message):
    __slots__ = ['step', 'motions']
    STEP_FIELD_NUMBER: _ClassVar[int]
    MOTIONS_FIELD_NUMBER: _ClassVar[int]
    step: int
    motions: _containers.RepeatedCompositeFieldContainer[_motion_pb2.PersonMotion]

    def __init__(self, step: _Optional[int]=..., motions: _Optional[_Iterable[_Union[_motion_pb2.PersonMotion, _Mapping]]]=...) -> None:
        ...