from city.economy.v1 import economy_pb2 as _economy_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetPersonRequest(_message.Message):
    __slots__ = ['person_ids']
    PERSON_IDS_FIELD_NUMBER: _ClassVar[int]
    person_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, person_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetPersonResponse(_message.Message):
    __slots__ = ['persons']
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    persons: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Person]

    def __init__(self, persons: _Optional[_Iterable[_Union[_economy_pb2.Person, _Mapping]]]=...) -> None:
        ...

class UpdatePersonMoneyRequest(_message.Message):
    __slots__ = ['items']
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    items: _containers.RepeatedCompositeFieldContainer[UpdatePersonMoneyRequestItem]

    def __init__(self, items: _Optional[_Iterable[_Union[UpdatePersonMoneyRequestItem, _Mapping]]]=...) -> None:
        ...

class UpdatePersonMoneyRequestItem(_message.Message):
    __slots__ = ['person_id', 'money']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    MONEY_FIELD_NUMBER: _ClassVar[int]
    person_id: int
    money: float

    def __init__(self, person_id: _Optional[int]=..., money: _Optional[float]=...) -> None:
        ...

class UpdatePersonMoneyResponse(_message.Message):
    __slots__ = ['persons']
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    persons: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Person]

    def __init__(self, persons: _Optional[_Iterable[_Union[_economy_pb2.Person, _Mapping]]]=...) -> None:
        ...