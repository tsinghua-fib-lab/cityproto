from city.economy.v2 import economy_pb2 as _economy_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class AddFirmRequest(_message.Message):
    __slots__ = ['firms']
    FIRMS_FIELD_NUMBER: _ClassVar[int]
    firms: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Firm]

    def __init__(self, firms: _Optional[_Iterable[_Union[_economy_pb2.Firm, _Mapping]]]=...) -> None:
        ...

class AddFirmResponse(_message.Message):
    __slots__ = ['firm_ids']
    FIRM_IDS_FIELD_NUMBER: _ClassVar[int]
    firm_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, firm_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class RemoveFirmRequest(_message.Message):
    __slots__ = ['firm_ids']
    FIRM_IDS_FIELD_NUMBER: _ClassVar[int]
    firm_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, firm_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class RemoveFirmResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetFirmRequest(_message.Message):
    __slots__ = ['firm_ids']
    FIRM_IDS_FIELD_NUMBER: _ClassVar[int]
    firm_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, firm_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetFirmResponse(_message.Message):
    __slots__ = ['firms']
    FIRMS_FIELD_NUMBER: _ClassVar[int]
    firms: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Firm]

    def __init__(self, firms: _Optional[_Iterable[_Union[_economy_pb2.Firm, _Mapping]]]=...) -> None:
        ...

class UpdateFirmRequest(_message.Message):
    __slots__ = ['firms']
    FIRMS_FIELD_NUMBER: _ClassVar[int]
    firms: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Firm]

    def __init__(self, firms: _Optional[_Iterable[_Union[_economy_pb2.Firm, _Mapping]]]=...) -> None:
        ...

class UpdateFirmResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListFirmsRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListFirmsResponse(_message.Message):
    __slots__ = ['firms']
    FIRMS_FIELD_NUMBER: _ClassVar[int]
    firms: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Firm]

    def __init__(self, firms: _Optional[_Iterable[_Union[_economy_pb2.Firm, _Mapping]]]=...) -> None:
        ...

class DeltaUpdateFirmRequest(_message.Message):
    __slots__ = ['updates']
    UPDATES_FIELD_NUMBER: _ClassVar[int]
    updates: _containers.RepeatedCompositeFieldContainer[FirmDeltaUpdate]

    def __init__(self, updates: _Optional[_Iterable[_Union[FirmDeltaUpdate, _Mapping]]]=...) -> None:
        ...

class FirmDeltaUpdate(_message.Message):
    __slots__ = ['firm_id', 'delta_price', 'delta_inventory', 'delta_demand', 'delta_sales', 'delta_currency', 'add_employees', 'remove_employees']
    FIRM_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_PRICE_FIELD_NUMBER: _ClassVar[int]
    DELTA_INVENTORY_FIELD_NUMBER: _ClassVar[int]
    DELTA_DEMAND_FIELD_NUMBER: _ClassVar[int]
    DELTA_SALES_FIELD_NUMBER: _ClassVar[int]
    DELTA_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    ADD_EMPLOYEES_FIELD_NUMBER: _ClassVar[int]
    REMOVE_EMPLOYEES_FIELD_NUMBER: _ClassVar[int]
    firm_id: int
    delta_price: float
    delta_inventory: int
    delta_demand: float
    delta_sales: float
    delta_currency: float
    add_employees: _containers.RepeatedScalarFieldContainer[int]
    remove_employees: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, firm_id: _Optional[int]=..., delta_price: _Optional[float]=..., delta_inventory: _Optional[int]=..., delta_demand: _Optional[float]=..., delta_sales: _Optional[float]=..., delta_currency: _Optional[float]=..., add_employees: _Optional[_Iterable[int]]=..., remove_employees: _Optional[_Iterable[int]]=...) -> None:
        ...

class DeltaUpdateFirmResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class AddNBSRequest(_message.Message):
    __slots__ = ['nbs']
    NBS_FIELD_NUMBER: _ClassVar[int]
    nbs: _economy_pb2.NBS

    def __init__(self, nbs: _Optional[_Union[_economy_pb2.NBS, _Mapping]]=...) -> None:
        ...

class AddNBSResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RemoveNBSRequest(_message.Message):
    __slots__ = ['nbs_id']
    NBS_ID_FIELD_NUMBER: _ClassVar[int]
    nbs_id: int

    def __init__(self, nbs_id: _Optional[int]=...) -> None:
        ...

class RemoveNBSResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetNBSRequest(_message.Message):
    __slots__ = ['nbs_id']
    NBS_ID_FIELD_NUMBER: _ClassVar[int]
    nbs_id: int

    def __init__(self, nbs_id: _Optional[int]=...) -> None:
        ...

class GetNBSResponse(_message.Message):
    __slots__ = ['nbs']
    NBS_FIELD_NUMBER: _ClassVar[int]
    nbs: _economy_pb2.NBS

    def __init__(self, nbs: _Optional[_Union[_economy_pb2.NBS, _Mapping]]=...) -> None:
        ...

class UpdateNBSRequest(_message.Message):
    __slots__ = ['nbs']
    NBS_FIELD_NUMBER: _ClassVar[int]
    nbs: _economy_pb2.NBS

    def __init__(self, nbs: _Optional[_Union[_economy_pb2.NBS, _Mapping]]=...) -> None:
        ...

class UpdateNBSResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListNBSRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListNBSResponse(_message.Message):
    __slots__ = ['nbs_list']
    NBS_LIST_FIELD_NUMBER: _ClassVar[int]
    nbs_list: _containers.RepeatedCompositeFieldContainer[_economy_pb2.NBS]

    def __init__(self, nbs_list: _Optional[_Iterable[_Union[_economy_pb2.NBS, _Mapping]]]=...) -> None:
        ...

class DeltaUpdateNBSRequest(_message.Message):
    __slots__ = ['nbs_id', 'delta_nominal_gdp', 'delta_real_gdp', 'delta_unemployment', 'delta_wages', 'delta_prices', 'delta_working_hours', 'delta_depression', 'delta_consumption_currency', 'delta_income_currency', 'delta_locus_control', 'delta_currency', 'add_citizen_ids', 'remove_citizen_ids', 'add_citizen_agent_ids', 'remove_citizen_agent_ids']

    class DeltaNominalGdpEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DeltaRealGdpEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DeltaUnemploymentEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DeltaWagesEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DeltaPricesEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DeltaWorkingHoursEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DeltaDepressionEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DeltaConsumptionCurrencyEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DeltaIncomeCurrencyEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...

    class DeltaLocusControlEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: float

        def __init__(self, key: _Optional[str]=..., value: _Optional[float]=...) -> None:
            ...
    NBS_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_NOMINAL_GDP_FIELD_NUMBER: _ClassVar[int]
    DELTA_REAL_GDP_FIELD_NUMBER: _ClassVar[int]
    DELTA_UNEMPLOYMENT_FIELD_NUMBER: _ClassVar[int]
    DELTA_WAGES_FIELD_NUMBER: _ClassVar[int]
    DELTA_PRICES_FIELD_NUMBER: _ClassVar[int]
    DELTA_WORKING_HOURS_FIELD_NUMBER: _ClassVar[int]
    DELTA_DEPRESSION_FIELD_NUMBER: _ClassVar[int]
    DELTA_CONSUMPTION_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    DELTA_INCOME_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    DELTA_LOCUS_CONTROL_FIELD_NUMBER: _ClassVar[int]
    DELTA_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    ADD_CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    REMOVE_CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    ADD_CITIZEN_AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    REMOVE_CITIZEN_AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    nbs_id: int
    delta_nominal_gdp: _containers.ScalarMap[str, float]
    delta_real_gdp: _containers.ScalarMap[str, float]
    delta_unemployment: _containers.ScalarMap[str, float]
    delta_wages: _containers.ScalarMap[str, float]
    delta_prices: _containers.ScalarMap[str, float]
    delta_working_hours: _containers.ScalarMap[str, float]
    delta_depression: _containers.ScalarMap[str, float]
    delta_consumption_currency: _containers.ScalarMap[str, float]
    delta_income_currency: _containers.ScalarMap[str, float]
    delta_locus_control: _containers.ScalarMap[str, float]
    delta_currency: float
    add_citizen_ids: _containers.RepeatedScalarFieldContainer[int]
    remove_citizen_ids: _containers.RepeatedScalarFieldContainer[int]
    add_citizen_agent_ids: _containers.RepeatedScalarFieldContainer[int]
    remove_citizen_agent_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, nbs_id: _Optional[int]=..., delta_nominal_gdp: _Optional[_Mapping[str, float]]=..., delta_real_gdp: _Optional[_Mapping[str, float]]=..., delta_unemployment: _Optional[_Mapping[str, float]]=..., delta_wages: _Optional[_Mapping[str, float]]=..., delta_prices: _Optional[_Mapping[str, float]]=..., delta_working_hours: _Optional[_Mapping[str, float]]=..., delta_depression: _Optional[_Mapping[str, float]]=..., delta_consumption_currency: _Optional[_Mapping[str, float]]=..., delta_income_currency: _Optional[_Mapping[str, float]]=..., delta_locus_control: _Optional[_Mapping[str, float]]=..., delta_currency: _Optional[float]=..., add_citizen_ids: _Optional[_Iterable[int]]=..., remove_citizen_ids: _Optional[_Iterable[int]]=..., add_citizen_agent_ids: _Optional[_Iterable[int]]=..., remove_citizen_agent_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class DeltaUpdateNBSResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class AddGovernmentRequest(_message.Message):
    __slots__ = ['government']
    GOVERNMENT_FIELD_NUMBER: _ClassVar[int]
    government: _economy_pb2.Government

    def __init__(self, government: _Optional[_Union[_economy_pb2.Government, _Mapping]]=...) -> None:
        ...

class AddGovernmentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RemoveGovernmentRequest(_message.Message):
    __slots__ = ['government_id']
    GOVERNMENT_ID_FIELD_NUMBER: _ClassVar[int]
    government_id: int

    def __init__(self, government_id: _Optional[int]=...) -> None:
        ...

class RemoveGovernmentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetGovernmentRequest(_message.Message):
    __slots__ = ['government_id']
    GOVERNMENT_ID_FIELD_NUMBER: _ClassVar[int]
    government_id: int

    def __init__(self, government_id: _Optional[int]=...) -> None:
        ...

class GetGovernmentResponse(_message.Message):
    __slots__ = ['government']
    GOVERNMENT_FIELD_NUMBER: _ClassVar[int]
    government: _economy_pb2.Government

    def __init__(self, government: _Optional[_Union[_economy_pb2.Government, _Mapping]]=...) -> None:
        ...

class UpdateGovernmentRequest(_message.Message):
    __slots__ = ['government']
    GOVERNMENT_FIELD_NUMBER: _ClassVar[int]
    government: _economy_pb2.Government

    def __init__(self, government: _Optional[_Union[_economy_pb2.Government, _Mapping]]=...) -> None:
        ...

class UpdateGovernmentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListGovernmentsRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListGovernmentsResponse(_message.Message):
    __slots__ = ['governments']
    GOVERNMENTS_FIELD_NUMBER: _ClassVar[int]
    governments: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Government]

    def __init__(self, governments: _Optional[_Iterable[_Union[_economy_pb2.Government, _Mapping]]]=...) -> None:
        ...

class DeltaUpdateGovernmentRequest(_message.Message):
    __slots__ = ['government_id', 'delta_bracket_cutoffs', 'delta_bracket_rates', 'delta_currency', 'add_citizen_ids', 'remove_citizen_ids']
    GOVERNMENT_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_BRACKET_CUTOFFS_FIELD_NUMBER: _ClassVar[int]
    DELTA_BRACKET_RATES_FIELD_NUMBER: _ClassVar[int]
    DELTA_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    ADD_CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    REMOVE_CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    government_id: int
    delta_bracket_cutoffs: _containers.RepeatedScalarFieldContainer[float]
    delta_bracket_rates: _containers.RepeatedScalarFieldContainer[float]
    delta_currency: float
    add_citizen_ids: _containers.RepeatedScalarFieldContainer[int]
    remove_citizen_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, government_id: _Optional[int]=..., delta_bracket_cutoffs: _Optional[_Iterable[float]]=..., delta_bracket_rates: _Optional[_Iterable[float]]=..., delta_currency: _Optional[float]=..., add_citizen_ids: _Optional[_Iterable[int]]=..., remove_citizen_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class DeltaUpdateGovernmentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class AddBankRequest(_message.Message):
    __slots__ = ['bank']
    BANK_FIELD_NUMBER: _ClassVar[int]
    bank: _economy_pb2.Bank

    def __init__(self, bank: _Optional[_Union[_economy_pb2.Bank, _Mapping]]=...) -> None:
        ...

class AddBankResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RemoveBankRequest(_message.Message):
    __slots__ = ['bank_id']
    BANK_ID_FIELD_NUMBER: _ClassVar[int]
    bank_id: int

    def __init__(self, bank_id: _Optional[int]=...) -> None:
        ...

class RemoveBankResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetBankRequest(_message.Message):
    __slots__ = ['bank_id']
    BANK_ID_FIELD_NUMBER: _ClassVar[int]
    bank_id: int

    def __init__(self, bank_id: _Optional[int]=...) -> None:
        ...

class GetBankResponse(_message.Message):
    __slots__ = ['bank']
    BANK_FIELD_NUMBER: _ClassVar[int]
    bank: _economy_pb2.Bank

    def __init__(self, bank: _Optional[_Union[_economy_pb2.Bank, _Mapping]]=...) -> None:
        ...

class UpdateBankRequest(_message.Message):
    __slots__ = ['bank']
    BANK_FIELD_NUMBER: _ClassVar[int]
    bank: _economy_pb2.Bank

    def __init__(self, bank: _Optional[_Union[_economy_pb2.Bank, _Mapping]]=...) -> None:
        ...

class UpdateBankResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListBanksRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListBanksResponse(_message.Message):
    __slots__ = ['banks']
    BANKS_FIELD_NUMBER: _ClassVar[int]
    banks: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Bank]

    def __init__(self, banks: _Optional[_Iterable[_Union[_economy_pb2.Bank, _Mapping]]]=...) -> None:
        ...

class DeltaUpdateBankRequest(_message.Message):
    __slots__ = ['bank_id', 'delta_interest_rate', 'delta_currency', 'add_citizen_ids', 'remove_citizen_ids']
    BANK_ID_FIELD_NUMBER: _ClassVar[int]
    DELTA_INTEREST_RATE_FIELD_NUMBER: _ClassVar[int]
    DELTA_CURRENCY_FIELD_NUMBER: _ClassVar[int]
    ADD_CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    REMOVE_CITIZEN_IDS_FIELD_NUMBER: _ClassVar[int]
    bank_id: int
    delta_interest_rate: float
    delta_currency: float
    add_citizen_ids: _containers.RepeatedScalarFieldContainer[int]
    remove_citizen_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, bank_id: _Optional[int]=..., delta_interest_rate: _Optional[float]=..., delta_currency: _Optional[float]=..., add_citizen_ids: _Optional[_Iterable[int]]=..., remove_citizen_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class DeltaUpdateBankResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class AddAgentRequest(_message.Message):
    __slots__ = ['agents']
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    agents: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Agent]

    def __init__(self, agents: _Optional[_Iterable[_Union[_economy_pb2.Agent, _Mapping]]]=...) -> None:
        ...

class AddAgentResponse(_message.Message):
    __slots__ = ['agent_ids']
    AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    agent_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, agent_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class RemoveAgentRequest(_message.Message):
    __slots__ = ['agent_ids']
    AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    agent_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, agent_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class RemoveAgentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetAgentRequest(_message.Message):
    __slots__ = ['agent_ids']
    AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    agent_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, agent_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetAgentResponse(_message.Message):
    __slots__ = ['agents']
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    agents: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Agent]

    def __init__(self, agents: _Optional[_Iterable[_Union[_economy_pb2.Agent, _Mapping]]]=...) -> None:
        ...

class UpdateAgentRequest(_message.Message):
    __slots__ = ['agents']
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    agents: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Agent]

    def __init__(self, agents: _Optional[_Iterable[_Union[_economy_pb2.Agent, _Mapping]]]=...) -> None:
        ...

class UpdateAgentResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListAgentsRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ListAgentsResponse(_message.Message):
    __slots__ = ['agents']
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    agents: _containers.RepeatedCompositeFieldContainer[_economy_pb2.Agent]

    def __init__(self, agents: _Optional[_Iterable[_Union[_economy_pb2.Agent, _Mapping]]]=...) -> None:
        ...

class DeltaUpdateAgentRequest(_message.Message):
    __slots__ = ['updates']
    UPDATES_FIELD_NUMBER: _ClassVar[int]
    updates: _containers.RepeatedCompositeFieldContainer[AgentDeltaUpdate]

    def __init__(self, updates: _Optional[_Iterable[_Union[AgentDeltaUpdate, _Mapping]]]=...) -> None:
        ...

class AgentDeltaUpdate(_message.Message):
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
    __slots__ = ['nbs_id']
    NBS_ID_FIELD_NUMBER: _ClassVar[int]
    nbs_id: int

    def __init__(self, nbs_id: _Optional[int]=...) -> None:
        ...

class CalculateRealGDPResponse(_message.Message):
    __slots__ = ['real_gdp']
    REAL_GDP_FIELD_NUMBER: _ClassVar[int]
    real_gdp: float

    def __init__(self, real_gdp: _Optional[float]=...) -> None:
        ...

class SaveEconomyEntitiesRequest(_message.Message):
    __slots__ = ['file_path']
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    file_path: str

    def __init__(self, file_path: _Optional[str]=...) -> None:
        ...

class SaveEconomyEntitiesResponse(_message.Message):
    __slots__ = ['firm_ids', 'nbs_ids', 'government_ids', 'bank_ids', 'agent_ids']
    FIRM_IDS_FIELD_NUMBER: _ClassVar[int]
    NBS_IDS_FIELD_NUMBER: _ClassVar[int]
    GOVERNMENT_IDS_FIELD_NUMBER: _ClassVar[int]
    BANK_IDS_FIELD_NUMBER: _ClassVar[int]
    AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    firm_ids: _containers.RepeatedScalarFieldContainer[int]
    nbs_ids: _containers.RepeatedScalarFieldContainer[int]
    government_ids: _containers.RepeatedScalarFieldContainer[int]
    bank_ids: _containers.RepeatedScalarFieldContainer[int]
    agent_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, firm_ids: _Optional[_Iterable[int]]=..., nbs_ids: _Optional[_Iterable[int]]=..., government_ids: _Optional[_Iterable[int]]=..., bank_ids: _Optional[_Iterable[int]]=..., agent_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class LoadEconomyEntitiesRequest(_message.Message):
    __slots__ = ['file_path']
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    file_path: str

    def __init__(self, file_path: _Optional[str]=...) -> None:
        ...

class LoadEconomyEntitiesResponse(_message.Message):
    __slots__ = ['firm_ids', 'nbs_ids', 'government_ids', 'bank_ids', 'agent_ids']
    FIRM_IDS_FIELD_NUMBER: _ClassVar[int]
    NBS_IDS_FIELD_NUMBER: _ClassVar[int]
    GOVERNMENT_IDS_FIELD_NUMBER: _ClassVar[int]
    BANK_IDS_FIELD_NUMBER: _ClassVar[int]
    AGENT_IDS_FIELD_NUMBER: _ClassVar[int]
    firm_ids: _containers.RepeatedScalarFieldContainer[int]
    nbs_ids: _containers.RepeatedScalarFieldContainer[int]
    government_ids: _containers.RepeatedScalarFieldContainer[int]
    bank_ids: _containers.RepeatedScalarFieldContainer[int]
    agent_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, firm_ids: _Optional[_Iterable[int]]=..., nbs_ids: _Optional[_Iterable[int]]=..., government_ids: _Optional[_Iterable[int]]=..., bank_ids: _Optional[_Iterable[int]]=..., agent_ids: _Optional[_Iterable[int]]=...) -> None:
        ...