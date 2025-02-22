from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Firm(_message.Message):
    __slots__ = ['id', 'employees', 'price', 'inventory', 'demand', 'sales', 'currency']
    ID_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEES_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    INVENTORY_FIELD_NUMBER: _ClassVar[int]
    DEMAND_FIELD_NUMBER: _ClassVar[int]
    SALES_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    id: int
    employees: _containers.RepeatedScalarFieldContainer[int]
    price: float
    inventory: int
    demand: float
    sales: float
    currency: float

    def __init__(self, id: _Optional[int]=..., employees: _Optional[_Iterable[int]]=..., price: _Optional[float]=..., inventory: _Optional[int]=..., demand: _Optional[float]=..., sales: _Optional[float]=..., currency: _Optional[float]=...) -> None:
        ...

class NBS(_message.Message):
    __slots__ = ['id', 'citizen_ids', 'nominal_gdp', 'real_gdp', 'unemployment', 'wages', 'prices', 'working_hours', 'depression', 'consumption_currency', 'income_currency', 'locus_control', 'citizen_agent_ids', 'currency']

    class NominalGdpEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class RealGdpEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class UnemploymentEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class WagesEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class PricesEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class WorkingHoursEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DepressionEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class ConsumptionCurrencyEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class IncomeCurrencyEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class LocusControlEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...
    ID_FIELD_NUMBER: _ClassVar[int]
    CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    NOMINAL_GDP_FIELD_NUMBER: _ClassVar[int]
    REAL_GDP_FIELD_NUMBER: _ClassVar[int]
    UNEMPLOYMENT_FIELD_NUMBER: _ClassVar[int]
    WAGES_FIELD_NUMBER: _ClassVar[int]
    PRICES_FIELD_NUMBER: _ClassVar[int]
    WORKING_HOURS_FIELD_NUMBER: _ClassVar[int]
    DEPRESSION_FIELD_NUMBER: _ClassVar[int]
    CONSUMPTION_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    INCOME_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    LOCUS_CONTROL_FIELD_NUMBER: _ClassVar[int]
    CITIZEN_AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    id: int
    citizen_ids: _containers.RepeatedScalarFieldContainer[int]
    nominal_gdp: _containers.ScalarMap[str, float]
    real_gdp: _containers.ScalarMap[str, float]
    unemployment: _containers.ScalarMap[str, float]
    wages: _containers.ScalarMap[str, float]
    prices: _containers.ScalarMap[str, float]
    working_hours: _containers.ScalarMap[str, float]
    depression: _containers.ScalarMap[str, float]
    consumption_currency: _containers.ScalarMap[str, float]
    income_currency: _containers.ScalarMap[str, float]
    locus_control: _containers.ScalarMap[str, float]
    citizen_agent_ids: _containers.RepeatedScalarFieldContainer[int]
    currency: float

    def __init__(self, id: _Optional[int]=..., citizen_ids: _Optional[_Iterable[int]]=..., nominal_gdp: _Optional[_Mapping[str, float]]=..., real_gdp: _Optional[_Mapping[str, float]]=..., unemployment: _Optional[_Mapping[str, float]]=..., wages: _Optional[_Mapping[str, float]]=..., prices: _Optional[_Mapping[str, float]]=..., working_hours: _Optional[_Mapping[str, float]]=..., depression: _Optional[_Mapping[str, float]]=..., consumption_currency: _Optional[_Mapping[str, float]]=..., income_currency: _Optional[_Mapping[str, float]]=..., locus_control: _Optional[_Mapping[str, float]]=..., citizen_agent_ids: _Optional[_Iterable[int]]=..., currency: _Optional[float]=...) -> None:
        ...

class Government(_message.Message):
    __slots__ = ['id', 'citizen_ids', 'bracket_cutoffs', 'bracket_rates', 'currency']
    ID_FIELD_NUMBER: _ClassVar[int]
    CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    BRACKET_CUTOFFS_FIELD_NUMBER: _ClassVar[int]
    BRACKET_RATES_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    id: int
    citizen_ids: _containers.RepeatedScalarFieldContainer[int]
    bracket_cutoffs: _containers.RepeatedScalarFieldContainer[float]
    bracket_rates: _containers.RepeatedScalarFieldContainer[float]
    currency: float

    def __init__(self, id: _Optional[int]=..., citizen_ids: _Optional[_Iterable[int]]=..., bracket_cutoffs: _Optional[_Iterable[float]]=..., bracket_rates: _Optional[_Iterable[float]]=..., currency: _Optional[float]=...) -> None:
        ...

class Bank(_message.Message):
    __slots__ = ['id', 'citizen_ids', 'interest_rate', 'currency']
    ID_FIELD_NUMBER: _ClassVar[int]
    CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    INTEREST_RATE_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    id: int
    citizen_ids: _containers.RepeatedScalarFieldContainer[int]
    interest_rate: float
    currency: float

    def __init__(self, id: _Optional[int]=..., citizen_ids: _Optional[_Iterable[int]]=..., interest_rate: _Optional[float]=..., currency: _Optional[float]=...) -> None:
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
    __slots__ = ['firms', 'nbs', 'governments', 'banks', 'agents']
    FIRMS_FIELD_NUMBER: _ClassVar[int]
    NBS_FIELD_NUMBER: _ClassVar[int]
    GOVERNMENTS_FIELD_NUMBER: _ClassVar[int]
    BANKS_FIELD_NUMBER: _ClassVar[int]
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    firms: _containers.RepeatedCompositeFieldContainer[Firm]
    nbs: _containers.RepeatedCompositeFieldContainer[NBS]
    governments: _containers.RepeatedCompositeFieldContainer[Government]
    banks: _containers.RepeatedCompositeFieldContainer[Bank]
    agents: _containers.RepeatedCompositeFieldContainer[Agent]

    def __init__(self, firms: _Optional[_Iterable[_Union[Firm, _Mapping]]]=..., nbs: _Optional[_Iterable[_Union[NBS, _Mapping]]]=..., governments: _Optional[_Iterable[_Union[Government, _Mapping]]]=..., banks: _Optional[_Iterable[_Union[Bank, _Mapping]]]=..., agents: _Optional[_Iterable[_Union[Agent, _Mapping]]]=...) -> None:
        ...