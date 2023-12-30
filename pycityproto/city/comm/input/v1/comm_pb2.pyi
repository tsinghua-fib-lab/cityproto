from city.geo.v2 import geo_pb2 as _geo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class NodeType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    NODE_TYPE_UNSPECIFIED: _ClassVar[NodeType]
    NODE_TYPE_INTERNET: _ClassVar[NodeType]
    NODE_TYPE_GATEWAY: _ClassVar[NodeType]
    NODE_TYPE_BASE_STATION: _ClassVar[NodeType]

class BaseStationType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    BASE_STATION_TYPE_UNSPECIFIED: _ClassVar[BaseStationType]
    BASE_STATION_TYPE_INDOOR: _ClassVar[BaseStationType]
    BASE_STATION_TYPE_OUTDOOR: _ClassVar[BaseStationType]
NODE_TYPE_UNSPECIFIED: NodeType
NODE_TYPE_INTERNET: NodeType
NODE_TYPE_GATEWAY: NodeType
NODE_TYPE_BASE_STATION: NodeType
BASE_STATION_TYPE_UNSPECIFIED: BaseStationType
BASE_STATION_TYPE_INDOOR: BaseStationType
BASE_STATION_TYPE_OUTDOOR: BaseStationType

class Node(_message.Message):
    __slots__ = ['id', 'type', 'parent_id', 'children_ids', 'position', 'aoi_id', 'freq_range_id', 'base_station_type']
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    PARENT_ID_FIELD_NUMBER: _ClassVar[int]
    CHILDREN_IDS_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    AOI_ID_FIELD_NUMBER: _ClassVar[int]
    FREQ_RANGE_ID_FIELD_NUMBER: _ClassVar[int]
    BASE_STATION_TYPE_FIELD_NUMBER: _ClassVar[int]
    id: int
    type: NodeType
    parent_id: int
    children_ids: _containers.RepeatedScalarFieldContainer[int]
    position: _geo_pb2.Position
    aoi_id: int
    freq_range_id: int
    base_station_type: BaseStationType

    def __init__(self, id: _Optional[int]=..., type: _Optional[_Union[NodeType, str]]=..., parent_id: _Optional[int]=..., children_ids: _Optional[_Iterable[int]]=..., position: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., aoi_id: _Optional[int]=..., freq_range_id: _Optional[int]=..., base_station_type: _Optional[_Union[BaseStationType, str]]=...) -> None:
        ...

class RepairStation(_message.Message):
    __slots__ = ['id', 'aoi_id', 'position']
    ID_FIELD_NUMBER: _ClassVar[int]
    AOI_ID_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    id: int
    aoi_id: int
    position: _geo_pb2.Position

    def __init__(self, id: _Optional[int]=..., aoi_id: _Optional[int]=..., position: _Optional[_Union[_geo_pb2.Position, _Mapping]]=...) -> None:
        ...

class Pump(_message.Message):
    __slots__ = ['id', 'position']
    ID_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    id: int
    position: _geo_pb2.Position

    def __init__(self, id: _Optional[int]=..., position: _Optional[_Union[_geo_pb2.Position, _Mapping]]=...) -> None:
        ...

class CommDemand(_message.Message):
    __slots__ = ['id', 'demands']
    ID_FIELD_NUMBER: _ClassVar[int]
    DEMANDS_FIELD_NUMBER: _ClassVar[int]
    id: int
    demands: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, id: _Optional[int]=..., demands: _Optional[_Iterable[float]]=...) -> None:
        ...

class Nodes(_message.Message):
    __slots__ = ['nodes', 'repair_stations', 'pumps']
    NODES_FIELD_NUMBER: _ClassVar[int]
    REPAIR_STATIONS_FIELD_NUMBER: _ClassVar[int]
    PUMPS_FIELD_NUMBER: _ClassVar[int]
    nodes: _containers.RepeatedCompositeFieldContainer[Node]
    repair_stations: _containers.RepeatedCompositeFieldContainer[RepairStation]
    pumps: _containers.RepeatedCompositeFieldContainer[Pump]

    def __init__(self, nodes: _Optional[_Iterable[_Union[Node, _Mapping]]]=..., repair_stations: _Optional[_Iterable[_Union[RepairStation, _Mapping]]]=..., pumps: _Optional[_Iterable[_Union[Pump, _Mapping]]]=...) -> None:
        ...

class CommDemands(_message.Message):
    __slots__ = ['comm_demands']
    COMM_DEMANDS_FIELD_NUMBER: _ClassVar[int]
    comm_demands: _containers.RepeatedCompositeFieldContainer[CommDemand]

    def __init__(self, comm_demands: _Optional[_Iterable[_Union[CommDemand, _Mapping]]]=...) -> None:
        ...