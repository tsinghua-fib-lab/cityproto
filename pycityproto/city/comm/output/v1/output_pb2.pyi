from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class NodeStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    NODE_STATUS_UNSPECIFIED: _ClassVar[NodeStatus]
    NODE_STATUS_OK: _ClassVar[NodeStatus]
    NODE_STATUS_BATTERY: _ClassVar[NodeStatus]
    NODE_STATUS_FAILURE: _ClassVar[NodeStatus]
    NODE_STATUS_RUINED: _ClassVar[NodeStatus]

class PersonConnectStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    PERSON_CONNECT_STATUS_UNSPECIFIED: _ClassVar[PersonConnectStatus]
    PERSON_CONNECT_STATUS_OK: _ClassVar[PersonConnectStatus]
    PERSON_CONNECT_STATUS_OUTAGE: _ClassVar[PersonConnectStatus]

class PersonDemandStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    PERSON_DEMAND_STATUS_UNSPECIFIED: _ClassVar[PersonDemandStatus]
    PERSON_DEMAND_STATUS_SATISFIED: _ClassVar[PersonDemandStatus]
    PERSON_DEMAND_STATUS_UNSATISFIED: _ClassVar[PersonDemandStatus]
    PERSON_DEMAND_STATUS_NO: _ClassVar[PersonDemandStatus]
NODE_STATUS_UNSPECIFIED: NodeStatus
NODE_STATUS_OK: NodeStatus
NODE_STATUS_BATTERY: NodeStatus
NODE_STATUS_FAILURE: NodeStatus
NODE_STATUS_RUINED: NodeStatus
PERSON_CONNECT_STATUS_UNSPECIFIED: PersonConnectStatus
PERSON_CONNECT_STATUS_OK: PersonConnectStatus
PERSON_CONNECT_STATUS_OUTAGE: PersonConnectStatus
PERSON_DEMAND_STATUS_UNSPECIFIED: PersonDemandStatus
PERSON_DEMAND_STATUS_SATISFIED: PersonDemandStatus
PERSON_DEMAND_STATUS_UNSATISFIED: PersonDemandStatus
PERSON_DEMAND_STATUS_NO: PersonDemandStatus

class Node(_message.Message):
    __slots__ = ['id', 'status', 'battery_remaining_time']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    BATTERY_REMAINING_TIME_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: NodeStatus
    battery_remaining_time: float

    def __init__(self, id: _Optional[int]=..., status: _Optional[_Union[NodeStatus, str]]=..., battery_remaining_time: _Optional[float]=...) -> None:
        ...

class BaseStation(_message.Message):
    __slots__ = ['id', 'demand_flow', 'actual_flow', 'num_agents', 'overload', 'unsatisfied_num', 'satisfied_num', 'outage_num', 'active_num', 'transmit_power']
    ID_FIELD_NUMBER: _ClassVar[int]
    DEMAND_FLOW_FIELD_NUMBER: _ClassVar[int]
    ACTUAL_FLOW_FIELD_NUMBER: _ClassVar[int]
    NUM_AGENTS_FIELD_NUMBER: _ClassVar[int]
    OVERLOAD_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_NUM_FIELD_NUMBER: _ClassVar[int]
    SATISFIED_NUM_FIELD_NUMBER: _ClassVar[int]
    OUTAGE_NUM_FIELD_NUMBER: _ClassVar[int]
    ACTIVE_NUM_FIELD_NUMBER: _ClassVar[int]
    TRANSMIT_POWER_FIELD_NUMBER: _ClassVar[int]
    id: int
    demand_flow: float
    actual_flow: float
    num_agents: int
    overload: bool
    unsatisfied_num: int
    satisfied_num: int
    outage_num: int
    active_num: int
    transmit_power: float

    def __init__(self, id: _Optional[int]=..., demand_flow: _Optional[float]=..., actual_flow: _Optional[float]=..., num_agents: _Optional[int]=..., overload: bool=..., unsatisfied_num: _Optional[int]=..., satisfied_num: _Optional[int]=..., outage_num: _Optional[int]=..., active_num: _Optional[int]=..., transmit_power: _Optional[float]=...) -> None:
        ...

class Signal(_message.Message):
    __slots__ = ['num_rows', 'num_columns', 'strength', 'base_station_id', 'freq_range_id']
    NUM_ROWS_FIELD_NUMBER: _ClassVar[int]
    NUM_COLUMNS_FIELD_NUMBER: _ClassVar[int]
    STRENGTH_FIELD_NUMBER: _ClassVar[int]
    BASE_STATION_ID_FIELD_NUMBER: _ClassVar[int]
    FREQ_RANGE_ID_FIELD_NUMBER: _ClassVar[int]
    num_rows: int
    num_columns: int
    strength: _containers.RepeatedScalarFieldContainer[float]
    base_station_id: _containers.RepeatedScalarFieldContainer[int]
    freq_range_id: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, num_rows: _Optional[int]=..., num_columns: _Optional[int]=..., strength: _Optional[_Iterable[float]]=..., base_station_id: _Optional[_Iterable[int]]=..., freq_range_id: _Optional[_Iterable[int]]=...) -> None:
        ...

class Person(_message.Message):
    __slots__ = ['id', 'demand_rate', 'actual_rate', 'connect_status', 'demand_status', 'strength', 'base_station_id', 'freq_range_ids', 'received_power']
    ID_FIELD_NUMBER: _ClassVar[int]
    DEMAND_RATE_FIELD_NUMBER: _ClassVar[int]
    ACTUAL_RATE_FIELD_NUMBER: _ClassVar[int]
    CONNECT_STATUS_FIELD_NUMBER: _ClassVar[int]
    DEMAND_STATUS_FIELD_NUMBER: _ClassVar[int]
    STRENGTH_FIELD_NUMBER: _ClassVar[int]
    BASE_STATION_ID_FIELD_NUMBER: _ClassVar[int]
    FREQ_RANGE_IDS_FIELD_NUMBER: _ClassVar[int]
    RECEIVED_POWER_FIELD_NUMBER: _ClassVar[int]
    id: int
    demand_rate: float
    actual_rate: float
    connect_status: PersonConnectStatus
    demand_status: PersonDemandStatus
    strength: float
    base_station_id: int
    freq_range_ids: _containers.RepeatedScalarFieldContainer[int]
    received_power: float

    def __init__(self, id: _Optional[int]=..., demand_rate: _Optional[float]=..., actual_rate: _Optional[float]=..., connect_status: _Optional[_Union[PersonConnectStatus, str]]=..., demand_status: _Optional[_Union[PersonDemandStatus, str]]=..., strength: _Optional[float]=..., base_station_id: _Optional[int]=..., freq_range_ids: _Optional[_Iterable[int]]=..., received_power: _Optional[float]=...) -> None:
        ...

class Aoi(_message.Message):
    __slots__ = ['id', 'demand_flow', 'actual_flow', 'outage_num', 'satisfied_num', 'unsatisfied_num', 'active_user_num']
    ID_FIELD_NUMBER: _ClassVar[int]
    DEMAND_FLOW_FIELD_NUMBER: _ClassVar[int]
    ACTUAL_FLOW_FIELD_NUMBER: _ClassVar[int]
    OUTAGE_NUM_FIELD_NUMBER: _ClassVar[int]
    SATISFIED_NUM_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_NUM_FIELD_NUMBER: _ClassVar[int]
    ACTIVE_USER_NUM_FIELD_NUMBER: _ClassVar[int]
    id: int
    demand_flow: float
    actual_flow: float
    outage_num: int
    satisfied_num: int
    unsatisfied_num: int
    active_user_num: int

    def __init__(self, id: _Optional[int]=..., demand_flow: _Optional[float]=..., actual_flow: _Optional[float]=..., outage_num: _Optional[int]=..., satisfied_num: _Optional[int]=..., unsatisfied_num: _Optional[int]=..., active_user_num: _Optional[int]=...) -> None:
        ...

class Statistics(_message.Message):
    __slots__ = ['num_satisfied_agents', 'num_unsatisfied_agents', 'num_outage_agents', 'num_active_agents', 'demand_flow', 'actual_flow', 'num_base_station', 'num_ok_base_station', 'num_ruined_base_station', 'num_stopped_base_station', 'num_overloaded_base_station', 'num_gateway', 'num_ok_gateway', 'num_ruined_gateway', 'num_stopped_gateway', 'num_overloaded_gateway', 'num_battery_gateway', 'power_consumption']
    NUM_SATISFIED_AGENTS_FIELD_NUMBER: _ClassVar[int]
    NUM_UNSATISFIED_AGENTS_FIELD_NUMBER: _ClassVar[int]
    NUM_OUTAGE_AGENTS_FIELD_NUMBER: _ClassVar[int]
    NUM_ACTIVE_AGENTS_FIELD_NUMBER: _ClassVar[int]
    DEMAND_FLOW_FIELD_NUMBER: _ClassVar[int]
    ACTUAL_FLOW_FIELD_NUMBER: _ClassVar[int]
    NUM_BASE_STATION_FIELD_NUMBER: _ClassVar[int]
    NUM_OK_BASE_STATION_FIELD_NUMBER: _ClassVar[int]
    NUM_RUINED_BASE_STATION_FIELD_NUMBER: _ClassVar[int]
    NUM_STOPPED_BASE_STATION_FIELD_NUMBER: _ClassVar[int]
    NUM_OVERLOADED_BASE_STATION_FIELD_NUMBER: _ClassVar[int]
    NUM_GATEWAY_FIELD_NUMBER: _ClassVar[int]
    NUM_OK_GATEWAY_FIELD_NUMBER: _ClassVar[int]
    NUM_RUINED_GATEWAY_FIELD_NUMBER: _ClassVar[int]
    NUM_STOPPED_GATEWAY_FIELD_NUMBER: _ClassVar[int]
    NUM_OVERLOADED_GATEWAY_FIELD_NUMBER: _ClassVar[int]
    NUM_BATTERY_GATEWAY_FIELD_NUMBER: _ClassVar[int]
    POWER_CONSUMPTION_FIELD_NUMBER: _ClassVar[int]
    num_satisfied_agents: int
    num_unsatisfied_agents: int
    num_outage_agents: int
    num_active_agents: int
    demand_flow: float
    actual_flow: float
    num_base_station: int
    num_ok_base_station: int
    num_ruined_base_station: int
    num_stopped_base_station: int
    num_overloaded_base_station: int
    num_gateway: int
    num_ok_gateway: int
    num_ruined_gateway: int
    num_stopped_gateway: int
    num_overloaded_gateway: int
    num_battery_gateway: int
    power_consumption: float

    def __init__(self, num_satisfied_agents: _Optional[int]=..., num_unsatisfied_agents: _Optional[int]=..., num_outage_agents: _Optional[int]=..., num_active_agents: _Optional[int]=..., demand_flow: _Optional[float]=..., actual_flow: _Optional[float]=..., num_base_station: _Optional[int]=..., num_ok_base_station: _Optional[int]=..., num_ruined_base_station: _Optional[int]=..., num_stopped_base_station: _Optional[int]=..., num_overloaded_base_station: _Optional[int]=..., num_gateway: _Optional[int]=..., num_ok_gateway: _Optional[int]=..., num_ruined_gateway: _Optional[int]=..., num_stopped_gateway: _Optional[int]=..., num_overloaded_gateway: _Optional[int]=..., num_battery_gateway: _Optional[int]=..., power_consumption: _Optional[float]=...) -> None:
        ...