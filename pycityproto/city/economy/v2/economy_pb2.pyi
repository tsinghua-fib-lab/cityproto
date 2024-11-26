from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class OrgType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    ORG_TYPE_UNSPECIFIED: _ClassVar[OrgType]
    ORG_TYPE_NBS: _ClassVar[OrgType]
    ORG_TYPE_FIRM: _ClassVar[OrgType]
    ORG_TYPE_BANK: _ClassVar[OrgType]
    ORG_TYPE_GOVERNMENT: _ClassVar[OrgType]
ORG_TYPE_UNSPECIFIED: OrgType
ORG_TYPE_NBS: OrgType
ORG_TYPE_FIRM: OrgType
ORG_TYPE_BANK: OrgType
ORG_TYPE_GOVERNMENT: OrgType

class Org(_message.Message):
    __slots__ = ['id', 'type', 'nominal_gdp', 'real_gdp', 'unemployment', 'wages', 'prices', 'inventory', 'price', 'currency', 'interest_rate', 'bracket_cutoffs', 'bracket_rates']
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NOMINAL_GDP_FIELD_NUMBER: _ClassVar[int]
    REAL_GDP_FIELD_NUMBER: _ClassVar[int]
    UNEMPLOYMENT_FIELD_NUMBER: _ClassVar[int]
    WAGES_FIELD_NUMBER: _ClassVar[int]
    PRICES_FIELD_NUMBER: _ClassVar[int]
    INVENTORY_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    INTEREST_RATE_FIELD_NUMBER: _ClassVar[int]
    BRACKET_CUTOFFS_FIELD_NUMBER: _ClassVar[int]
    BRACKET_RATES_FIELD_NUMBER: _ClassVar[int]
    id: int
    type: OrgType
    nominal_gdp: _containers.RepeatedScalarFieldContainer[float]
    real_gdp: _containers.RepeatedScalarFieldContainer[float]
    unemployment: _containers.RepeatedScalarFieldContainer[float]
    wages: _containers.RepeatedScalarFieldContainer[float]
    prices: _containers.RepeatedScalarFieldContainer[float]
    inventory: int
    price: float
    currency: float
    interest_rate: float
    bracket_cutoffs: _containers.RepeatedScalarFieldContainer[float]
    bracket_rates: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, id: _Optional[int]=..., type: _Optional[_Union[OrgType, str]]=..., nominal_gdp: _Optional[_Iterable[float]]=..., real_gdp: _Optional[_Iterable[float]]=..., unemployment: _Optional[_Iterable[float]]=..., wages: _Optional[_Iterable[float]]=..., prices: _Optional[_Iterable[float]]=..., inventory: _Optional[int]=..., price: _Optional[float]=..., currency: _Optional[float]=..., interest_rate: _Optional[float]=..., bracket_cutoffs: _Optional[_Iterable[float]]=..., bracket_rates: _Optional[_Iterable[float]]=...) -> None:
        ...

class Agent(_message.Message):
    __slots__ = ['id', 'currency']
    ID_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    id: int
    currency: float

    def __init__(self, id: _Optional[int]=..., currency: _Optional[float]=...) -> None:
        ...