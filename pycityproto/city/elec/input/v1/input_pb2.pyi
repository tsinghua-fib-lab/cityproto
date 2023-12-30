from city.geo.v2 import geo_pb2 as _geo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class FacilityType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    FACILITY_TYPE_UNSPECIFIED: _ClassVar[FacilityType]
    FACILITY_TYPE_POWER_STATION: _ClassVar[FacilityType]
    FACILITY_TYPE_TRANSFORMER_500: _ClassVar[FacilityType]
    FACILITY_TYPE_TRANSFORMER_220: _ClassVar[FacilityType]
    FACILITY_TYPE_TRANSFORMER_110: _ClassVar[FacilityType]
    FACILITY_TYPE_TRANSFORMER_10: _ClassVar[FacilityType]
    FACILITY_TYPE_BASE_STATION: _ClassVar[FacilityType]
    FACILITY_TYPE_GATEWAY: _ClassVar[FacilityType]
    FACILITY_TYPE_DRAINAGE_PUMP: _ClassVar[FacilityType]
    FACILITY_TYPE_TRAFFIC_LIGHT: _ClassVar[FacilityType]
    FACILITY_TYPE_AOI: _ClassVar[FacilityType]
    FACILITY_TYPE_SUPPLY_PUMP: _ClassVar[FacilityType]
FACILITY_TYPE_UNSPECIFIED: FacilityType
FACILITY_TYPE_POWER_STATION: FacilityType
FACILITY_TYPE_TRANSFORMER_500: FacilityType
FACILITY_TYPE_TRANSFORMER_220: FacilityType
FACILITY_TYPE_TRANSFORMER_110: FacilityType
FACILITY_TYPE_TRANSFORMER_10: FacilityType
FACILITY_TYPE_BASE_STATION: FacilityType
FACILITY_TYPE_GATEWAY: FacilityType
FACILITY_TYPE_DRAINAGE_PUMP: FacilityType
FACILITY_TYPE_TRAFFIC_LIGHT: FacilityType
FACILITY_TYPE_AOI: FacilityType
FACILITY_TYPE_SUPPLY_PUMP: FacilityType

class RepairStation(_message.Message):
    __slots__ = ['aoi_id', 'position']
    AOI_ID_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    aoi_id: int
    position: _geo_pb2.LongLatPosition

    def __init__(self, aoi_id: _Optional[int]=..., position: _Optional[_Union[_geo_pb2.LongLatPosition, _Mapping]]=...) -> None:
        ...

class Facility(_message.Message):
    __slots__ = ['id', 'position', 'type', 'relation', 'foreign_id', 'aoi_id', 'num_transformer']
    ID_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    RELATION_FIELD_NUMBER: _ClassVar[int]
    FOREIGN_ID_FIELD_NUMBER: _ClassVar[int]
    AOI_ID_FIELD_NUMBER: _ClassVar[int]
    NUM_TRANSFORMER_FIELD_NUMBER: _ClassVar[int]
    id: int
    position: _geo_pb2.LongLatPosition
    type: FacilityType
    relation: _containers.RepeatedScalarFieldContainer[int]
    foreign_id: int
    aoi_id: int
    num_transformer: int

    def __init__(self, id: _Optional[int]=..., position: _Optional[_Union[_geo_pb2.LongLatPosition, _Mapping]]=..., type: _Optional[_Union[FacilityType, str]]=..., relation: _Optional[_Iterable[int]]=..., foreign_id: _Optional[int]=..., aoi_id: _Optional[int]=..., num_transformer: _Optional[int]=...) -> None:
        ...

class Facilities(_message.Message):
    __slots__ = ['facilities', 'repair_stations']
    FACILITIES_FIELD_NUMBER: _ClassVar[int]
    REPAIR_STATIONS_FIELD_NUMBER: _ClassVar[int]
    facilities: _containers.RepeatedCompositeFieldContainer[Facility]
    repair_stations: _containers.RepeatedCompositeFieldContainer[RepairStation]

    def __init__(self, facilities: _Optional[_Iterable[_Union[Facility, _Mapping]]]=..., repair_stations: _Optional[_Iterable[_Union[RepairStation, _Mapping]]]=...) -> None:
        ...