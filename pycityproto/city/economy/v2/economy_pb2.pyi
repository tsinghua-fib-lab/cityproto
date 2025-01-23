from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
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
    __slots__ = ['id', 'type', 'nominal_gdp', 'real_gdp', 'unemployment', 'wages', 'prices', 'inventory', 'price', 'currency', 'interest_rate', 'bracket_cutoffs', 'bracket_rates', 'consumption_currency', 'consumption_propensity', 'income_currency', 'depression', 'locus_control', 'working_hours', 'employees', 'citizens', 'demand', 'sales']
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
    CONSUMPTION_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    CONSUMPTION_PROPENSITY_FIELD_NUMBER: _ClassVar[int]
    INCOME_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    DEPRESSION_FIELD_NUMBER: _ClassVar[int]
    LOCUS_CONTROL_FIELD_NUMBER: _ClassVar[int]
    WORKING_HOURS_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEES_FIELD_NUMBER: _ClassVar[int]
    CITIZENS_FIELD_NUMBER: _ClassVar[int]
    DEMAND_FIELD_NUMBER: _ClassVar[int]
    SALES_FIELD_NUMBER: _ClassVar[int]
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
    consumption_currency: _containers.RepeatedScalarFieldContainer[float]
    consumption_propensity: _containers.RepeatedScalarFieldContainer[float]
    income_currency: _containers.RepeatedScalarFieldContainer[float]
    depression: _containers.RepeatedScalarFieldContainer[float]
    locus_control: _containers.RepeatedScalarFieldContainer[float]
    working_hours: _containers.RepeatedScalarFieldContainer[float]
    employees: _containers.RepeatedScalarFieldContainer[int]
    citizens: _containers.RepeatedScalarFieldContainer[int]
    demand: int
    sales: int

    def __init__(self, id: _Optional[int]=..., type: _Optional[_Union[OrgType, str]]=..., nominal_gdp: _Optional[_Iterable[float]]=..., real_gdp: _Optional[_Iterable[float]]=..., unemployment: _Optional[_Iterable[float]]=..., wages: _Optional[_Iterable[float]]=..., prices: _Optional[_Iterable[float]]=..., inventory: _Optional[int]=..., price: _Optional[float]=..., currency: _Optional[float]=..., interest_rate: _Optional[float]=..., bracket_cutoffs: _Optional[_Iterable[float]]=..., bracket_rates: _Optional[_Iterable[float]]=..., consumption_currency: _Optional[_Iterable[float]]=..., consumption_propensity: _Optional[_Iterable[float]]=..., income_currency: _Optional[_Iterable[float]]=..., depression: _Optional[_Iterable[float]]=..., locus_control: _Optional[_Iterable[float]]=..., working_hours: _Optional[_Iterable[float]]=..., employees: _Optional[_Iterable[int]]=..., citizens: _Optional[_Iterable[int]]=..., demand: _Optional[int]=..., sales: _Optional[int]=...) -> None:
        ...

class Agent(_message.Message):
    __slots__ = ['id', 'currency', 'firm_id', 'skill', 'consumption', 'income']
    ID_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    FIRM_ID_FIELD_NUMBER: _ClassVar[int]
    SKILL_FIELD_NUMBER: _ClassVar[int]
    CONSUMPTION_FIELD_NUMBER: _ClassVar[int]
    INCOME_FIELD_NUMBER: _ClassVar[int]
    id: int
    currency: float
    firm_id: int
    skill: float
    consumption: float
    income: float

    def __init__(self, id: _Optional[int]=..., currency: _Optional[float]=..., firm_id: _Optional[int]=..., skill: _Optional[float]=..., consumption: _Optional[float]=..., income: _Optional[float]=...) -> None:
        ...

class EconomyEntities(_message.Message):
    __slots__ = ['orgs', 'agents']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[Org]
    agents: _containers.RepeatedCompositeFieldContainer[Agent]

    def __init__(self, orgs: _Optional[_Iterable[_Union[Org, _Mapping]]]=..., agents: _Optional[_Iterable[_Union[Agent, _Mapping]]]=...) -> None:
        ...