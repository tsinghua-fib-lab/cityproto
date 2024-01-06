from city.economy.v1 import economy_pb2 as _economy_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetOrgRequest(_message.Message):
    __slots__ = ['org_ids']
    ORG_IDS_FIELD_NUMBER: _ClassVar[int]
    org_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, org_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetOrgResponse(_message.Message):
    __slots__ = ['orgs']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Org]

    def __init__(self, orgs: _Optional[_Iterable[_Union[_economy_pb2.Org, _Mapping]]]=...) -> None:
        ...

class UpdateOrgMoneyRequest(_message.Message):
    __slots__ = ['items']
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    items: _containers.RepeatedCompositeFieldContainer[UpdateOrgMoneyRequestItem]

    def __init__(self, items: _Optional[_Iterable[_Union[UpdateOrgMoneyRequestItem, _Mapping]]]=...) -> None:
        ...

class UpdateOrgMoneyRequestItem(_message.Message):
    __slots__ = ['org_id', 'money']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    MONEY_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    money: float

    def __init__(self, org_id: _Optional[int]=..., money: _Optional[float]=...) -> None:
        ...

class UpdateOrgMoneyResponse(_message.Message):
    __slots__ = ['orgs']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Org]

    def __init__(self, orgs: _Optional[_Iterable[_Union[_economy_pb2.Org, _Mapping]]]=...) -> None:
        ...

class UpdateOrgGoodsRequest(_message.Message):
    __slots__ = ['items']
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    items: _containers.RepeatedCompositeFieldContainer[UpdateOrgGoodsRequestItem]

    def __init__(self, items: _Optional[_Iterable[_Union[UpdateOrgGoodsRequestItem, _Mapping]]]=...) -> None:
        ...

class UpdateOrgGoodsRequestItem(_message.Message):
    __slots__ = ['org_id', 'goods']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    GOODS_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    goods: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Goods]

    def __init__(self, org_id: _Optional[int]=..., goods: _Optional[_Iterable[_Union[_economy_pb2.Goods, _Mapping]]]=...) -> None:
        ...

class UpdateOrgGoodsResponse(_message.Message):
    __slots__ = ['orgs']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Org]

    def __init__(self, orgs: _Optional[_Iterable[_Union[_economy_pb2.Org, _Mapping]]]=...) -> None:
        ...

class UpdateOrgEmployeeRequest(_message.Message):
    __slots__ = ['items']
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    items: _containers.RepeatedCompositeFieldContainer[UpdateOrgEmployeeRequestItem]

    def __init__(self, items: _Optional[_Iterable[_Union[UpdateOrgEmployeeRequestItem, _Mapping]]]=...) -> None:
        ...

class UpdateOrgEmployeeRequestItem(_message.Message):
    __slots__ = ['org_id', 'adds', 'dels', 'updates']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    ADDS_FIELD_NUMBER: _ClassVar[int]
    DELS_FIELD_NUMBER: _ClassVar[int]
    UPDATES_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    adds: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Employee]
    dels: _containers.RepeatedScalarFieldContainer[int]
    updates: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Employee]

    def __init__(self, org_id: _Optional[int]=..., adds: _Optional[_Iterable[_Union[_economy_pb2.Employee, _Mapping]]]=..., dels: _Optional[_Iterable[int]]=..., updates: _Optional[_Iterable[_Union[_economy_pb2.Employee, _Mapping]]]=...) -> None:
        ...

class UpdateOrgEmployeeResponse(_message.Message):
    __slots__ = ['orgs']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Org]

    def __init__(self, orgs: _Optional[_Iterable[_Union[_economy_pb2.Org, _Mapping]]]=...) -> None:
        ...

class UpdateOrgJobRequest(_message.Message):
    __slots__ = ['items']
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    items: _containers.RepeatedCompositeFieldContainer[UpdateOrgJobRequestItem]

    def __init__(self, items: _Optional[_Iterable[_Union[UpdateOrgJobRequestItem, _Mapping]]]=...) -> None:
        ...

class UpdateOrgJobRequestItem(_message.Message):
    __slots__ = ['org_id', 'jobs']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    JOBS_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    jobs: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Job]

    def __init__(self, org_id: _Optional[int]=..., jobs: _Optional[_Iterable[_Union[_economy_pb2.Job, _Mapping]]]=...) -> None:
        ...

class UpdateOrgJobResponse(_message.Message):
    __slots__ = ['orgs']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Org]

    def __init__(self, orgs: _Optional[_Iterable[_Union[_economy_pb2.Org, _Mapping]]]=...) -> None:
        ...