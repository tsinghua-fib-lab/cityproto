from city.geo.v2 import geo_pb2 as _geo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class LinkType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    LINK_TYPE_UNSPECIFIED: _ClassVar[LinkType]
    LINK_TYPE_PIPE: _ClassVar[LinkType]
    LINK_TYPE_PUMP: _ClassVar[LinkType]
LINK_TYPE_UNSPECIFIED: LinkType
LINK_TYPE_PIPE: LinkType
LINK_TYPE_PUMP: LinkType

class Road(_message.Message):
    __slots__ = ['id', 'depth']
    ID_FIELD_NUMBER: _ClassVar[int]
    DEPTH_FIELD_NUMBER: _ClassVar[int]
    id: int
    depth: float

    def __init__(self, id: _Optional[int]=..., depth: _Optional[float]=...) -> None:
        ...

class RoadFlood(_message.Message):
    __slots__ = ['position', 'depth']
    POSITION_FIELD_NUMBER: _ClassVar[int]
    DEPTH_FIELD_NUMBER: _ClassVar[int]
    position: _geo_pb2.LongLatPosition
    depth: float

    def __init__(self, position: _Optional[_Union[_geo_pb2.LongLatPosition, _Mapping]]=..., depth: _Optional[float]=...) -> None:
        ...

class DetailedRoad(_message.Message):
    __slots__ = ['id', 'depths']
    ID_FIELD_NUMBER: _ClassVar[int]
    DEPTHS_FIELD_NUMBER: _ClassVar[int]
    id: int
    depths: _containers.RepeatedCompositeFieldContainer[RoadFlood]

    def __init__(self, id: _Optional[int]=..., depths: _Optional[_Iterable[_Union[RoadFlood, _Mapping]]]=...) -> None:
        ...

class Node(_message.Message):
    __slots__ = ['id', 'head']
    ID_FIELD_NUMBER: _ClassVar[int]
    HEAD_FIELD_NUMBER: _ClassVar[int]
    id: str
    head: float

    def __init__(self, id: _Optional[str]=..., head: _Optional[float]=...) -> None:
        ...

class Link(_message.Message):
    __slots__ = ['id', 'type', 'flow', 'ok']
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    FLOW_FIELD_NUMBER: _ClassVar[int]
    OK_FIELD_NUMBER: _ClassVar[int]
    id: str
    type: LinkType
    flow: float
    ok: bool

    def __init__(self, id: _Optional[str]=..., type: _Optional[_Union[LinkType, str]]=..., flow: _Optional[float]=..., ok: bool=...) -> None:
        ...

class Aoi(_message.Message):
    __slots__ = ['id', 'unsatisfied_num', 'unsatisfied_ratio', 'demand', 'supply']
    ID_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_NUM_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_RATIO_FIELD_NUMBER: _ClassVar[int]
    DEMAND_FIELD_NUMBER: _ClassVar[int]
    SUPPLY_FIELD_NUMBER: _ClassVar[int]
    id: int
    unsatisfied_num: int
    unsatisfied_ratio: float
    demand: float
    supply: float

    def __init__(self, id: _Optional[int]=..., unsatisfied_num: _Optional[int]=..., unsatisfied_ratio: _Optional[float]=..., demand: _Optional[float]=..., supply: _Optional[float]=...) -> None:
        ...

class DrainageBasicInfo(_message.Message):
    __slots__ = ['average_power', 'undrained_volume', 'drained_volume', 'average_flow', 'flooded_volume']
    AVERAGE_POWER_FIELD_NUMBER: _ClassVar[int]
    UNDRAINED_VOLUME_FIELD_NUMBER: _ClassVar[int]
    DRAINED_VOLUME_FIELD_NUMBER: _ClassVar[int]
    AVERAGE_FLOW_FIELD_NUMBER: _ClassVar[int]
    FLOODED_VOLUME_FIELD_NUMBER: _ClassVar[int]
    average_power: float
    undrained_volume: float
    drained_volume: float
    average_flow: float
    flooded_volume: float

    def __init__(self, average_power: _Optional[float]=..., undrained_volume: _Optional[float]=..., drained_volume: _Optional[float]=..., average_flow: _Optional[float]=..., flooded_volume: _Optional[float]=...) -> None:
        ...

class SupplyBasicInfo(_message.Message):
    __slots__ = ['average_power', 'average_flow']
    AVERAGE_POWER_FIELD_NUMBER: _ClassVar[int]
    AVERAGE_FLOW_FIELD_NUMBER: _ClassVar[int]
    average_power: float
    average_flow: float

    def __init__(self, average_power: _Optional[float]=..., average_flow: _Optional[float]=...) -> None:
        ...

class SupplyDemandStatistics(_message.Message):
    __slots__ = ['persons_demand', 'unsatisfied_persons', 'unsatisfied_persons_ratio', 'aois_demand', 'unsatisfied_aois', 'unsatisfied_aois_ratio']
    PERSONS_DEMAND_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_PERSONS_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_PERSONS_RATIO_FIELD_NUMBER: _ClassVar[int]
    AOIS_DEMAND_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_AOIS_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_AOIS_RATIO_FIELD_NUMBER: _ClassVar[int]
    persons_demand: float
    unsatisfied_persons: int
    unsatisfied_persons_ratio: float
    aois_demand: float
    unsatisfied_aois: int
    unsatisfied_aois_ratio: float

    def __init__(self, persons_demand: _Optional[float]=..., unsatisfied_persons: _Optional[int]=..., unsatisfied_persons_ratio: _Optional[float]=..., aois_demand: _Optional[float]=..., unsatisfied_aois: _Optional[int]=..., unsatisfied_aois_ratio: _Optional[float]=...) -> None:
        ...

class FailureStatistics(_message.Message):
    __slots__ = ['failure_num', 'normal_num', 'failure_ratio']
    FAILURE_NUM_FIELD_NUMBER: _ClassVar[int]
    NORMAL_NUM_FIELD_NUMBER: _ClassVar[int]
    FAILURE_RATIO_FIELD_NUMBER: _ClassVar[int]
    failure_num: int
    normal_num: int
    failure_ratio: float

    def __init__(self, failure_num: _Optional[int]=..., normal_num: _Optional[int]=..., failure_ratio: _Optional[float]=...) -> None:
        ...

class DrainageMetrics(_message.Message):
    __slots__ = ['drainage_basic_info', 'load_ratio', 'failure_statistics']
    DRAINAGE_BASIC_INFO_FIELD_NUMBER: _ClassVar[int]
    LOAD_RATIO_FIELD_NUMBER: _ClassVar[int]
    FAILURE_STATISTICS_FIELD_NUMBER: _ClassVar[int]
    drainage_basic_info: DrainageBasicInfo
    load_ratio: float
    failure_statistics: FailureStatistics

    def __init__(self, drainage_basic_info: _Optional[_Union[DrainageBasicInfo, _Mapping]]=..., load_ratio: _Optional[float]=..., failure_statistics: _Optional[_Union[FailureStatistics, _Mapping]]=...) -> None:
        ...

class SupplyMetrics(_message.Message):
    __slots__ = ['supply_basic_info', 'supply_demand_statistics', 'load_ratio', 'failure_statistics']
    SUPPLY_BASIC_INFO_FIELD_NUMBER: _ClassVar[int]
    SUPPLY_DEMAND_STATISTICS_FIELD_NUMBER: _ClassVar[int]
    LOAD_RATIO_FIELD_NUMBER: _ClassVar[int]
    FAILURE_STATISTICS_FIELD_NUMBER: _ClassVar[int]
    supply_basic_info: SupplyBasicInfo
    supply_demand_statistics: SupplyDemandStatistics
    load_ratio: float
    failure_statistics: FailureStatistics

    def __init__(self, supply_basic_info: _Optional[_Union[SupplyBasicInfo, _Mapping]]=..., supply_demand_statistics: _Optional[_Union[SupplyDemandStatistics, _Mapping]]=..., load_ratio: _Optional[float]=..., failure_statistics: _Optional[_Union[FailureStatistics, _Mapping]]=...) -> None:
        ...