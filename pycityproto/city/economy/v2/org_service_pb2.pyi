from city.economy.v2 import economy_pb2 as _economy_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class AddOrgRequest(_message.Message):
    __slots__ = ['org']
    ORG_FIELD_NUMBER: _ClassVar[int]
    org: _economy_pb2.Org

    def __init__(self, org: _Optional[_Union[_economy_pb2.Org, _Mapping]]=...) -> None:
        ...

class AddOrgResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RemoveOrgRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class RemoveOrgResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetOrgRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetOrgResponse(_message.Message):
    __slots__ = ['org']
    ORG_FIELD_NUMBER: _ClassVar[int]
    org: _economy_pb2.Org

    def __init__(self, org: _Optional[_Union[_economy_pb2.Org, _Mapping]]=...) -> None:
        ...

class UpdateOrgRequest(_message.Message):
    __slots__ = ['org']
    ORG_FIELD_NUMBER: _ClassVar[int]
    org: _economy_pb2.Org

    def __init__(self, org: _Optional[_Union[_economy_pb2.Org, _Mapping]]=...) -> None:
        ...

class UpdateOrgResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class AddAgentRequest(_message.Message):
    __slots__ = ['agent']
    AGENT_FIELD_NUMBER: _ClassVar[int]
    agent: _economy_pb2.Agent

    def __init__(self, agent: _Optional[_Union[_economy_pb2.Agent, _Mapping]]=...) -> None:
        ...

class AddAgentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RemoveAgentRequest(_message.Message):
    __slots__ = ['agent_id']
    AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    agent_id: int

    def __init__(self, agent_id: _Optional[int]=...) -> None:
        ...

class RemoveAgentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetNominalGDPRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetNominalGDPResponse(_message.Message):
    __slots__ = ['nominal_gdp']
    NOMINAL_GDP_FIELD_NUMBER: _ClassVar[int]
    nominal_gdp: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, nominal_gdp: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetNominalGDPRequest(_message.Message):
    __slots__ = ['org_id', 'nominal_gdp']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    NOMINAL_GDP_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    nominal_gdp: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., nominal_gdp: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetNominalGDPResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetRealGDPRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetRealGDPResponse(_message.Message):
    __slots__ = ['real_gdp']
    REAL_GDP_FIELD_NUMBER: _ClassVar[int]
    real_gdp: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, real_gdp: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetRealGDPRequest(_message.Message):
    __slots__ = ['org_id', 'real_gdp']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    REAL_GDP_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    real_gdp: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., real_gdp: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetRealGDPResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetUnemploymentRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetUnemploymentResponse(_message.Message):
    __slots__ = ['unemployment']
    UNEMPLOYMENT_FIELD_NUMBER: _ClassVar[int]
    unemployment: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, unemployment: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetUnemploymentRequest(_message.Message):
    __slots__ = ['org_id', 'unemployment']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    UNEMPLOYMENT_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    unemployment: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., unemployment: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetUnemploymentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetWagesRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetWagesResponse(_message.Message):
    __slots__ = ['wages']
    WAGES_FIELD_NUMBER: _ClassVar[int]
    wages: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, wages: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetWagesRequest(_message.Message):
    __slots__ = ['org_id', 'wages']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    WAGES_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    wages: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., wages: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetWagesResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetPricesRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetPricesResponse(_message.Message):
    __slots__ = ['prices']
    PRICES_FIELD_NUMBER: _ClassVar[int]
    prices: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, prices: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetPricesRequest(_message.Message):
    __slots__ = ['org_id', 'prices']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    PRICES_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    prices: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., prices: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetPricesResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetInventoryRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetInventoryResponse(_message.Message):
    __slots__ = ['inventory']
    INVENTORY_FIELD_NUMBER: _ClassVar[int]
    inventory: int

    def __init__(self, inventory: _Optional[int]=...) -> None:
        ...

class SetInventoryRequest(_message.Message):
    __slots__ = ['org_id', 'inventory']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    INVENTORY_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    inventory: int

    def __init__(self, org_id: _Optional[int]=..., inventory: _Optional[int]=...) -> None:
        ...

class SetInventoryResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetPriceRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetPriceResponse(_message.Message):
    __slots__ = ['price']
    PRICE_FIELD_NUMBER: _ClassVar[int]
    price: float

    def __init__(self, price: _Optional[float]=...) -> None:
        ...

class SetPriceRequest(_message.Message):
    __slots__ = ['org_id', 'price']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    price: float

    def __init__(self, org_id: _Optional[int]=..., price: _Optional[float]=...) -> None:
        ...

class SetPriceResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetCurrencyRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetCurrencyResponse(_message.Message):
    __slots__ = ['currency']
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    currency: float

    def __init__(self, currency: _Optional[float]=...) -> None:
        ...

class SetCurrencyRequest(_message.Message):
    __slots__ = ['org_id', 'currency']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    currency: float

    def __init__(self, org_id: _Optional[int]=..., currency: _Optional[float]=...) -> None:
        ...

class SetCurrencyResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetInterestRateRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetInterestRateResponse(_message.Message):
    __slots__ = ['interest_rate']
    INTEREST_RATE_FIELD_NUMBER: _ClassVar[int]
    interest_rate: float

    def __init__(self, interest_rate: _Optional[float]=...) -> None:
        ...

class SetInterestRateRequest(_message.Message):
    __slots__ = ['org_id', 'interest_rate']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    INTEREST_RATE_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    interest_rate: float

    def __init__(self, org_id: _Optional[int]=..., interest_rate: _Optional[float]=...) -> None:
        ...

class SetInterestRateResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetBracketCutoffsRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetBracketCutoffsResponse(_message.Message):
    __slots__ = ['bracket_cutoffs']
    BRACKET_CUTOFFS_FIELD_NUMBER: _ClassVar[int]
    bracket_cutoffs: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, bracket_cutoffs: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetBracketCutoffsRequest(_message.Message):
    __slots__ = ['org_id', 'bracket_cutoffs']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    BRACKET_CUTOFFS_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    bracket_cutoffs: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., bracket_cutoffs: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetBracketCutoffsResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetBracketRatesRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetBracketRatesResponse(_message.Message):
    __slots__ = ['bracket_rates']
    BRACKET_RATES_FIELD_NUMBER: _ClassVar[int]
    bracket_rates: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, bracket_rates: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetBracketRatesRequest(_message.Message):
    __slots__ = ['org_id', 'bracket_rates']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    BRACKET_RATES_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    bracket_rates: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., bracket_rates: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetBracketRatesResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class CalculateTaxesDueRequest(_message.Message):
    __slots__ = ['government_id', 'agent_ids', 'incomes', 'enable_redistribution']
    GOVERNMENT_ID_FIELD_NUMBER: _ClassVar[int]
    AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    INCOMES_FIELD_NUMBER: _ClassVar[int]
    ENABLE_REDISTRIBUTION_FIELD_NUMBER: _ClassVar[int]
    government_id: int
    agent_ids: _containers.RepeatedScalarFieldContainer[int]
    incomes: _containers.RepeatedScalarFieldContainer[float]
    enable_redistribution: bool

    def __init__(self, government_id: _Optional[int]=..., agent_ids: _Optional[_Iterable[int]]=..., incomes: _Optional[_Iterable[float]]=..., enable_redistribution: bool=...) -> None:
        ...

class CalculateTaxesDueResponse(_message.Message):
    __slots__ = ['taxes_due', 'updated_incomes']
    TAXES_DUE_FIELD_NUMBER: _ClassVar[int]
    UPDATED_INCOMES_FIELD_NUMBER: _ClassVar[int]
    taxes_due: float
    updated_incomes: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, taxes_due: _Optional[float]=..., updated_incomes: _Optional[_Iterable[float]]=...) -> None:
        ...

class CalculateConsumptionRequest(_message.Message):
    __slots__ = ['firm_ids', 'agent_id', 'demands', 'consumption_accumulation']
    FIRM_IDS_FIELD_NUMBER: _ClassVar[int]
    AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    DEMANDS_FIELD_NUMBER: _ClassVar[int]
    CONSUMPTION_ACCUMULATION_FIELD_NUMBER: _ClassVar[int]
    firm_ids: _containers.RepeatedScalarFieldContainer[int]
    agent_id: int
    demands: _containers.RepeatedScalarFieldContainer[int]
    consumption_accumulation: bool

    def __init__(self, firm_ids: _Optional[_Iterable[int]]=..., agent_id: _Optional[int]=..., demands: _Optional[_Iterable[int]]=..., consumption_accumulation: bool=...) -> None:
        ...

class CalculateConsumptionResponse(_message.Message):
    __slots__ = ['actual_consumption', 'success']
    ACTUAL_CONSUMPTION_FIELD_NUMBER: _ClassVar[int]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    actual_consumption: float
    success: bool

    def __init__(self, actual_consumption: _Optional[float]=..., success: bool=...) -> None:
        ...

class CalculateInterestRequest(_message.Message):
    __slots__ = ['bank_id', 'agent_ids']
    BANK_ID_FIELD_NUMBER: _ClassVar[int]
    AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    bank_id: int
    agent_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, bank_id: _Optional[int]=..., agent_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class CalculateInterestResponse(_message.Message):
    __slots__ = ['total_interest', 'updated_currencies']
    TOTAL_INTEREST_FIELD_NUMBER: _ClassVar[int]
    UPDATED_CURRENCIES_FIELD_NUMBER: _ClassVar[int]
    total_interest: float
    updated_currencies: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, total_interest: _Optional[float]=..., updated_currencies: _Optional[_Iterable[float]]=...) -> None:
        ...

class SaveEconomyEntitiesRequest(_message.Message):
    __slots__ = ['file_path']
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    file_path: str

    def __init__(self, file_path: _Optional[str]=...) -> None:
        ...

class SaveEconomyEntitiesResponse(_message.Message):
    __slots__ = ['org_ids', 'agent_ids']
    ORG_IDS_FIELD_NUMBER: _ClassVar[int]
    AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    org_ids: _containers.RepeatedScalarFieldContainer[int]
    agent_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, org_ids: _Optional[_Iterable[int]]=..., agent_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class LoadEconomyEntitiesRequest(_message.Message):
    __slots__ = ['file_path']
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    file_path: str

    def __init__(self, file_path: _Optional[str]=...) -> None:
        ...

class LoadEconomyEntitiesResponse(_message.Message):
    __slots__ = ['org_ids', 'agent_ids']
    ORG_IDS_FIELD_NUMBER: _ClassVar[int]
    AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    org_ids: _containers.RepeatedScalarFieldContainer[int]
    agent_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, org_ids: _Optional[_Iterable[int]]=..., agent_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetConsumptionCurrencyRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetConsumptionCurrencyResponse(_message.Message):
    __slots__ = ['consumption_currency']
    CONSUMPTION_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    consumption_currency: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, consumption_currency: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetConsumptionCurrencyRequest(_message.Message):
    __slots__ = ['org_id', 'consumption_currency']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    CONSUMPTION_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    consumption_currency: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., consumption_currency: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetConsumptionCurrencyResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetConsumptionPropensityRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetConsumptionPropensityResponse(_message.Message):
    __slots__ = ['consumption_propensity']
    CONSUMPTION_PROPENSITY_FIELD_NUMBER: _ClassVar[int]
    consumption_propensity: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, consumption_propensity: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetConsumptionPropensityRequest(_message.Message):
    __slots__ = ['org_id', 'consumption_propensity']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    CONSUMPTION_PROPENSITY_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    consumption_propensity: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., consumption_propensity: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetConsumptionPropensityResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetIncomeCurrencyRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetIncomeCurrencyResponse(_message.Message):
    __slots__ = ['income_currency']
    INCOME_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    income_currency: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, income_currency: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetIncomeCurrencyRequest(_message.Message):
    __slots__ = ['org_id', 'income_currency']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    INCOME_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    income_currency: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., income_currency: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetIncomeCurrencyResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetDepressionRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetDepressionResponse(_message.Message):
    __slots__ = ['depression']
    DEPRESSION_FIELD_NUMBER: _ClassVar[int]
    depression: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, depression: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetDepressionRequest(_message.Message):
    __slots__ = ['org_id', 'depression']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    DEPRESSION_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    depression: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., depression: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetDepressionResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetLocusControlRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetLocusControlResponse(_message.Message):
    __slots__ = ['locus_control']
    LOCUS_CONTROL_FIELD_NUMBER: _ClassVar[int]
    locus_control: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, locus_control: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetLocusControlRequest(_message.Message):
    __slots__ = ['org_id', 'locus_control']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    LOCUS_CONTROL_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    locus_control: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., locus_control: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetLocusControlResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetWorkingHoursRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetWorkingHoursResponse(_message.Message):
    __slots__ = ['working_hours']
    WORKING_HOURS_FIELD_NUMBER: _ClassVar[int]
    working_hours: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, working_hours: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetWorkingHoursRequest(_message.Message):
    __slots__ = ['org_id', 'working_hours']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    WORKING_HOURS_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    working_hours: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., working_hours: _Optional[_Iterable[float]]=...) -> None:
        ...

class SetWorkingHoursResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetOrgEntityIdsRequest(_message.Message):
    __slots__ = ['type']
    TYPE_FIELD_NUMBER: _ClassVar[int]
    type: _economy_pb2.OrgType

    def __init__(self, type: _Optional[_Union[_economy_pb2.OrgType, str]]=...) -> None:
        ...

class GetOrgEntityIdsResponse(_message.Message):
    __slots__ = ['org_ids']
    ORG_IDS_FIELD_NUMBER: _ClassVar[int]
    org_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, org_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class AddInventoryRequest(_message.Message):
    __slots__ = ['org_id', 'delta_inventory']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_INVENTORY_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    delta_inventory: int

    def __init__(self, org_id: _Optional[int]=..., delta_inventory: _Optional[int]=...) -> None:
        ...

class AddInventoryResponse(_message.Message):
    __slots__ = ['inventory']
    INVENTORY_FIELD_NUMBER: _ClassVar[int]
    inventory: int

    def __init__(self, inventory: _Optional[int]=...) -> None:
        ...

class AddPriceRequest(_message.Message):
    __slots__ = ['org_id', 'delta_price']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_PRICE_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    delta_price: float

    def __init__(self, org_id: _Optional[int]=..., delta_price: _Optional[float]=...) -> None:
        ...

class AddPriceResponse(_message.Message):
    __slots__ = ['price']
    PRICE_FIELD_NUMBER: _ClassVar[int]
    price: float

    def __init__(self, price: _Optional[float]=...) -> None:
        ...

class AddCurrencyRequest(_message.Message):
    __slots__ = ['org_id', 'delta_currency']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    delta_currency: float

    def __init__(self, org_id: _Optional[int]=..., delta_currency: _Optional[float]=...) -> None:
        ...

class AddCurrencyResponse(_message.Message):
    __slots__ = ['currency']
    CURRENCY_FIELD_NUMBER: _ClassVar[int]
    currency: float

    def __init__(self, currency: _Optional[float]=...) -> None:
        ...

class AddInterestRateRequest(_message.Message):
    __slots__ = ['org_id', 'delta_interest_rate']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_INTEREST_RATE_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    delta_interest_rate: float

    def __init__(self, org_id: _Optional[int]=..., delta_interest_rate: _Optional[float]=...) -> None:
        ...

class AddInterestRateResponse(_message.Message):
    __slots__ = ['interest_rate']
    INTEREST_RATE_FIELD_NUMBER: _ClassVar[int]
    interest_rate: float

    def __init__(self, interest_rate: _Optional[float]=...) -> None:
        ...

class GetEmployeesRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetEmployeesResponse(_message.Message):
    __slots__ = ['employee_ids']
    EMPLOYEE_IDS_FIELD_NUMBER: _ClassVar[int]
    employee_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, employee_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class SetEmployeesRequest(_message.Message):
    __slots__ = ['org_id', 'employee_ids']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEE_IDS_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    employee_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, org_id: _Optional[int]=..., employee_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class SetEmployeesResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class AddEmployeeRequest(_message.Message):
    __slots__ = ['org_id', 'employee_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEE_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    employee_id: int

    def __init__(self, org_id: _Optional[int]=..., employee_id: _Optional[int]=...) -> None:
        ...

class AddEmployeeResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RemoveEmployeeRequest(_message.Message):
    __slots__ = ['org_id', 'employee_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEE_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    employee_id: int

    def __init__(self, org_id: _Optional[int]=..., employee_id: _Optional[int]=...) -> None:
        ...

class RemoveEmployeeResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetCitizensRequest(_message.Message):
    __slots__ = ['org_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int

    def __init__(self, org_id: _Optional[int]=...) -> None:
        ...

class GetCitizensResponse(_message.Message):
    __slots__ = ['citizen_ids']
    CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    citizen_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, citizen_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class SetCitizensRequest(_message.Message):
    __slots__ = ['org_id', 'citizen_ids']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    citizen_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, org_id: _Optional[int]=..., citizen_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class SetCitizensResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class AddCitizenRequest(_message.Message):
    __slots__ = ['org_id', 'citizen_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    CITIZEN_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    citizen_id: int

    def __init__(self, org_id: _Optional[int]=..., citizen_id: _Optional[int]=...) -> None:
        ...

class AddCitizenResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RemoveCitizenRequest(_message.Message):
    __slots__ = ['org_id', 'citizen_id']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    CITIZEN_ID_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    citizen_id: int

    def __init__(self, org_id: _Optional[int]=..., citizen_id: _Optional[int]=...) -> None:
        ...

class RemoveCitizenResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetAgentRequest(_message.Message):
    __slots__ = ['agent_id']
    AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    agent_id: int

    def __init__(self, agent_id: _Optional[int]=...) -> None:
        ...

class GetAgentResponse(_message.Message):
    __slots__ = ['agent']
    AGENT_FIELD_NUMBER: _ClassVar[int]
    agent: _economy_pb2.Agent

    def __init__(self, agent: _Optional[_Union[_economy_pb2.Agent, _Mapping]]=...) -> None:
        ...

class UpdateAgentRequest(_message.Message):
    __slots__ = ['agent']
    AGENT_FIELD_NUMBER: _ClassVar[int]
    agent: _economy_pb2.Agent

    def __init__(self, agent: _Optional[_Union[_economy_pb2.Agent, _Mapping]]=...) -> None:
        ...

class UpdateAgentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class BatchGetRequest(_message.Message):
    __slots__ = ['ids', 'type']
    IDS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    ids: _containers.RepeatedScalarFieldContainer[int]
    type: str

    def __init__(self, ids: _Optional[_Iterable[int]]=..., type: _Optional[str]=...) -> None:
        ...

class BatchGetResponse(_message.Message):
    __slots__ = ['orgs', 'agents']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Org]
    agents: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Agent]

    def __init__(self, orgs: _Optional[_Iterable[_Union[_economy_pb2.Org, _Mapping]]]=..., agents: _Optional[_Iterable[_Union[_economy_pb2.Agent, _Mapping]]]=...) -> None:
        ...

class BatchUpdateRequest(_message.Message):
    __slots__ = ['orgs', 'agents']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Org]
    agents: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Agent]

    def __init__(self, orgs: _Optional[_Iterable[_Union[_economy_pb2.Org, _Mapping]]]=..., agents: _Optional[_Iterable[_Union[_economy_pb2.Agent, _Mapping]]]=...) -> None:
        ...

class BatchUpdateResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class DeltaUpdateOrgRequest(_message.Message):
    __slots__ = ['org_id', 'delta_inventory', 'delta_price', 'delta_currency', 'delta_interest_rate', 'add_employees', 'remove_employees']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_INVENTORY_FIELD_NUMBER: _ClassVar[int]
    DELTA_PRICE_FIELD_NUMBER: _ClassVar[int]
    DELTA_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    DELTA_INTEREST_RATE_FIELD_NUMBER: _ClassVar[int]
    ADD_EMPLOYEES_FIELD_NUMBER: _ClassVar[int]
    REMOVE_EMPLOYEES_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    delta_inventory: float
    delta_price: float
    delta_currency: float
    delta_interest_rate: float
    add_employees: _containers.RepeatedScalarFieldContainer[int]
    remove_employees: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, org_id: _Optional[int]=..., delta_inventory: _Optional[float]=..., delta_price: _Optional[float]=..., delta_currency: _Optional[float]=..., delta_interest_rate: _Optional[float]=..., add_employees: _Optional[_Iterable[int]]=..., remove_employees: _Optional[_Iterable[int]]=...) -> None:
        ...

class DeltaUpdateOrgResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class DeltaUpdateAgentRequest(_message.Message):
    __slots__ = ['agent_id', 'delta_currency', 'delta_skill', 'delta_consumption', 'delta_income']
    AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    DELTA_SKILL_FIELD_NUMBER: _ClassVar[int]
    DELTA_CONSUMPTION_FIELD_NUMBER: _ClassVar[int]
    DELTA_INCOME_FIELD_NUMBER: _ClassVar[int]
    agent_id: int
    delta_currency: float
    delta_skill: float
    delta_consumption: float
    delta_income: float

    def __init__(self, agent_id: _Optional[int]=..., delta_currency: _Optional[float]=..., delta_skill: _Optional[float]=..., delta_consumption: _Optional[float]=..., delta_income: _Optional[float]=...) -> None:
        ...

class DeltaUpdateAgentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class BatchDeltaUpdateRequest(_message.Message):
    __slots__ = ['orgs', 'agents']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[DeltaUpdateOrgRequest]
    agents: _containers.RepeatedCompositeFieldContainer[DeltaUpdateAgentRequest]

    def __init__(self, orgs: _Optional[_Iterable[_Union[DeltaUpdateOrgRequest, _Mapping]]]=..., agents: _Optional[_Iterable[_Union[DeltaUpdateAgentRequest, _Mapping]]]=...) -> None:
        ...

class BatchDeltaUpdateResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class CalculateRealGDPRequest(_message.Message):
    __slots__ = ['nbs_agent_id']
    NBS_AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    nbs_agent_id: int

    def __init__(self, nbs_agent_id: _Optional[int]=...) -> None:
        ...

class CalculateRealGDPResponse(_message.Message):
    __slots__ = ['real_gdp']
    REAL_GDP_FIELD_NUMBER: _ClassVar[int]
    real_gdp: float

    def __init__(self, real_gdp: _Optional[float]=...) -> None:
        ...