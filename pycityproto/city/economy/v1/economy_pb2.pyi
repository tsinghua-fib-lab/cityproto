from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Person(_message.Message):
    __slots__ = ['id', 'money', 'org_id']
    ID_FIELD_NUMBER: _ClassVar[int]
    MONEY_FIELD_NUMBER: _ClassVar[int]
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    money: float
    org_id: int

    def __init__(self, id: _Optional[int]=..., money: _Optional[float]=..., org_id: _Optional[int]=...) -> None:
        ...

class Employee(_message.Message):
    __slots__ = ['person_id', 'salary']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    SALARY_FIELD_NUMBER: _ClassVar[int]
    person_id: int
    salary: float

    def __init__(self, person_id: _Optional[int]=..., salary: _Optional[float]=...) -> None:
        ...

class Job(_message.Message):
    __slots__ = ['name', 'employee_count', 'salary']
    NAME_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEE_COUNT_FIELD_NUMBER: _ClassVar[int]
    SALARY_FIELD_NUMBER: _ClassVar[int]
    name: str
    employee_count: int
    salary: float

    def __init__(self, name: _Optional[str]=..., employee_count: _Optional[int]=..., salary: _Optional[float]=...) -> None:
        ...

class Goods(_message.Message):
    __slots__ = ['type', 'name', 'count', 'price']
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    type: str
    name: str
    count: int
    price: float

    def __init__(self, type: _Optional[str]=..., name: _Optional[str]=..., count: _Optional[int]=..., price: _Optional[float]=...) -> None:
        ...

class Org(_message.Message):
    __slots__ = ['id', 'poi_id', 'employees', 'jobs', 'money', 'goods', 'functions']
    ID_FIELD_NUMBER: _ClassVar[int]
    POI_ID_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEES_FIELD_NUMBER: _ClassVar[int]
    JOBS_FIELD_NUMBER: _ClassVar[int]
    MONEY_FIELD_NUMBER: _ClassVar[int]
    GOODS_FIELD_NUMBER: _ClassVar[int]
    FUNCTIONS_FIELD_NUMBER: _ClassVar[int]
    id: int
    poi_id: int
    employees: _containers.RepeatedCompositeFieldContainer[Employee]
    jobs: _containers.RepeatedCompositeFieldContainer[Job]
    money: float
    goods: _containers.RepeatedCompositeFieldContainer[Goods]
    functions: _containers.RepeatedScalarFieldContainer[str]

    def __init__(self, id: _Optional[int]=..., poi_id: _Optional[int]=..., employees: _Optional[_Iterable[_Union[Employee, _Mapping]]]=..., jobs: _Optional[_Iterable[_Union[Job, _Mapping]]]=..., money: _Optional[float]=..., goods: _Optional[_Iterable[_Union[Goods, _Mapping]]]=..., functions: _Optional[_Iterable[str]]=...) -> None:
        ...

class Economy(_message.Message):
    __slots__ = ['persons', 'orgs']
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    ORGS_FIELD_NUMBER: _ClassVar[int]
    persons: _containers.RepeatedCompositeFieldContainer[Person]
    orgs: _containers.RepeatedCompositeFieldContainer[Org]

    def __init__(self, persons: _Optional[_Iterable[_Union[Person, _Mapping]]]=..., orgs: _Optional[_Iterable[_Union[Org, _Mapping]]]=...) -> None:
        ...