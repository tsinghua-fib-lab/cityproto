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

class BatchSetRequest(_message.Message):
    __slots__ = ['orgs', 'agents']
    ORGS_FIELD_NUMBER: _ClassVar[int]
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    orgs: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Org]
    agents: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Agent]

    def __init__(self, orgs: _Optional[_Iterable[_Union[_economy_pb2.Org, _Mapping]]]=..., agents: _Optional[_Iterable[_Union[_economy_pb2.Agent, _Mapping]]]=...) -> None:
        ...

class BatchSetResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class DeltaUpdateOrgRequest(_message.Message):
    __slots__ = ['org_id', 'delta_nominal_gdp', 'delta_real_gdp', 'delta_unemployment', 'delta_wages', 'delta_prices', 'delta_inventory', 'delta_price', 'delta_currency', 'delta_interest_rate', 'delta_bracket_cutoffs', 'delta_bracket_rates', 'delta_demand', 'delta_sales', 'add_employees', 'remove_employees', 'add_citizens', 'remove_citizens', 'delta_consumption_currency', 'delta_consumption_propensity', 'delta_income_currency', 'delta_depression', 'delta_locus_control', 'delta_working_hours']
    ORG_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_NOMINAL_GDP_FIELD_NUMBER: _ClassVar[int]
    DELTA_REAL_GDP_FIELD_NUMBER: _ClassVar[int]
    DELTA_UNEMPLOYMENT_FIELD_NUMBER: _ClassVar[int]
    DELTA_WAGES_FIELD_NUMBER: _ClassVar[int]
    DELTA_PRICES_FIELD_NUMBER: _ClassVar[int]
    DELTA_INVENTORY_FIELD_NUMBER: _ClassVar[int]
    DELTA_PRICE_FIELD_NUMBER: _ClassVar[int]
    DELTA_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    DELTA_INTEREST_RATE_FIELD_NUMBER: _ClassVar[int]
    DELTA_BRACKET_CUTOFFS_FIELD_NUMBER: _ClassVar[int]
    DELTA_BRACKET_RATES_FIELD_NUMBER: _ClassVar[int]
    DELTA_DEMAND_FIELD_NUMBER: _ClassVar[int]
    DELTA_SALES_FIELD_NUMBER: _ClassVar[int]
    ADD_EMPLOYEES_FIELD_NUMBER: _ClassVar[int]
    REMOVE_EMPLOYEES_FIELD_NUMBER: _ClassVar[int]
    ADD_CITIZENS_FIELD_NUMBER: _ClassVar[int]
    REMOVE_CITIZENS_FIELD_NUMBER: _ClassVar[int]
    DELTA_CONSUMPTION_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    DELTA_CONSUMPTION_PROPENSITY_FIELD_NUMBER: _ClassVar[int]
    DELTA_INCOME_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    DELTA_DEPRESSION_FIELD_NUMBER: _ClassVar[int]
    DELTA_LOCUS_CONTROL_FIELD_NUMBER: _ClassVar[int]
    DELTA_WORKING_HOURS_FIELD_NUMBER: _ClassVar[int]
    org_id: int
    delta_nominal_gdp: _containers.RepeatedScalarFieldContainer[float]
    delta_real_gdp: _containers.RepeatedScalarFieldContainer[float]
    delta_unemployment: _containers.RepeatedScalarFieldContainer[float]
    delta_wages: _containers.RepeatedScalarFieldContainer[float]
    delta_prices: _containers.RepeatedScalarFieldContainer[float]
    delta_inventory: int
    delta_price: float
    delta_currency: float
    delta_interest_rate: float
    delta_bracket_cutoffs: _containers.RepeatedScalarFieldContainer[float]
    delta_bracket_rates: _containers.RepeatedScalarFieldContainer[float]
    delta_demand: int
    delta_sales: int
    add_employees: _containers.RepeatedScalarFieldContainer[int]
    remove_employees: _containers.RepeatedScalarFieldContainer[int]
    add_citizens: _containers.RepeatedScalarFieldContainer[int]
    remove_citizens: _containers.RepeatedScalarFieldContainer[int]
    delta_consumption_currency: _containers.RepeatedScalarFieldContainer[float]
    delta_consumption_propensity: _containers.RepeatedScalarFieldContainer[float]
    delta_income_currency: _containers.RepeatedScalarFieldContainer[float]
    delta_depression: _containers.RepeatedScalarFieldContainer[float]
    delta_locus_control: _containers.RepeatedScalarFieldContainer[float]
    delta_working_hours: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, org_id: _Optional[int]=..., delta_nominal_gdp: _Optional[_Iterable[float]]=..., delta_real_gdp: _Optional[_Iterable[float]]=..., delta_unemployment: _Optional[_Iterable[float]]=..., delta_wages: _Optional[_Iterable[float]]=..., delta_prices: _Optional[_Iterable[float]]=..., delta_inventory: _Optional[int]=..., delta_price: _Optional[float]=..., delta_currency: _Optional[float]=..., delta_interest_rate: _Optional[float]=..., delta_bracket_cutoffs: _Optional[_Iterable[float]]=..., delta_bracket_rates: _Optional[_Iterable[float]]=..., delta_demand: _Optional[int]=..., delta_sales: _Optional[int]=..., add_employees: _Optional[_Iterable[int]]=..., remove_employees: _Optional[_Iterable[int]]=..., add_citizens: _Optional[_Iterable[int]]=..., remove_citizens: _Optional[_Iterable[int]]=..., delta_consumption_currency: _Optional[_Iterable[float]]=..., delta_consumption_propensity: _Optional[_Iterable[float]]=..., delta_income_currency: _Optional[_Iterable[float]]=..., delta_depression: _Optional[_Iterable[float]]=..., delta_locus_control: _Optional[_Iterable[float]]=..., delta_working_hours: _Optional[_Iterable[float]]=...) -> None:
        ...

class DeltaUpdateOrgResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class DeltaUpdateAgentRequest(_message.Message):
    __slots__ = ['agent_id', 'delta_currency', 'new_firm_id', 'delta_skill', 'delta_consumption', 'delta_income']
    AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    NEW_FIRM_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_SKILL_FIELD_NUMBER: _ClassVar[int]
    DELTA_CONSUMPTION_FIELD_NUMBER: _ClassVar[int]
    DELTA_INCOME_FIELD_NUMBER: _ClassVar[int]
    agent_id: int
    delta_currency: float
    new_firm_id: int
    delta_skill: float
    delta_consumption: float
    delta_income: float

    def __init__(self, agent_id: _Optional[int]=..., delta_currency: _Optional[float]=..., new_firm_id: _Optional[int]=..., delta_skill: _Optional[float]=..., delta_consumption: _Optional[float]=..., delta_income: _Optional[float]=...) -> None:
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