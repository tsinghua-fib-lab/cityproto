from city.map.v2 import lane_state_pb2 as _lane_state_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetJunctionRequest(_message.Message):
    __slots__ = ['junction_ids', 'exclude_lane', 'exclude_person']
    JUNCTION_IDS_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_LANE_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_PERSON_FIELD_NUMBER: _ClassVar[int]
    junction_ids: _containers.RepeatedScalarFieldContainer[int]
    exclude_lane: bool
    exclude_person: bool

    def __init__(self, junction_ids: _Optional[_Iterable[int]]=..., exclude_lane: bool=..., exclude_person: bool=...) -> None:
        ...

class GetJunctionResponse(_message.Message):
    __slots__ = ['states']
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedCompositeFieldContainer[JunctionState]

    def __init__(self, states: _Optional[_Iterable[_Union[JunctionState, _Mapping]]]=...) -> None:
        ...

class JunctionState(_message.Message):
    __slots__ = ['id', 'in_vehicle_cnt', 'out_vehicle_cnt', 'vehicle_cnt', 'cum_in_vehicle_cnt', 'cum_out_vehicle_cnt', 'lanes', 'predecessor_driving_lanes', 'total_queuing_vehicle_cnt', 'total_queuing_time', 'avg_queuing_time', 'max_queuing_vehicle_cnt', 'has_traffic_light']
    ID_FIELD_NUMBER: _ClassVar[int]
    IN_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    OUT_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    CUM_IN_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    CUM_OUT_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    LANES_FIELD_NUMBER: _ClassVar[int]
    PREDECESSOR_DRIVING_LANES_FIELD_NUMBER: _ClassVar[int]
    TOTAL_QUEUING_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    TOTAL_QUEUING_TIME_FIELD_NUMBER: _ClassVar[int]
    AVG_QUEUING_TIME_FIELD_NUMBER: _ClassVar[int]
    MAX_QUEUING_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    HAS_TRAFFIC_LIGHT_FIELD_NUMBER: _ClassVar[int]
    id: int
    in_vehicle_cnt: int
    out_vehicle_cnt: int
    vehicle_cnt: int
    cum_in_vehicle_cnt: int
    cum_out_vehicle_cnt: int
    lanes: _containers.RepeatedCompositeFieldContainer[_lane_state_pb2.LaneState]
    predecessor_driving_lanes: _containers.RepeatedCompositeFieldContainer[_lane_state_pb2.LaneState]
    total_queuing_vehicle_cnt: int
    total_queuing_time: float
    avg_queuing_time: float
    max_queuing_vehicle_cnt: int
    has_traffic_light: bool

    def __init__(self, id: _Optional[int]=..., in_vehicle_cnt: _Optional[int]=..., out_vehicle_cnt: _Optional[int]=..., vehicle_cnt: _Optional[int]=..., cum_in_vehicle_cnt: _Optional[int]=..., cum_out_vehicle_cnt: _Optional[int]=..., lanes: _Optional[_Iterable[_Union[_lane_state_pb2.LaneState, _Mapping]]]=..., predecessor_driving_lanes: _Optional[_Iterable[_Union[_lane_state_pb2.LaneState, _Mapping]]]=..., total_queuing_vehicle_cnt: _Optional[int]=..., total_queuing_time: _Optional[float]=..., avg_queuing_time: _Optional[float]=..., max_queuing_vehicle_cnt: _Optional[int]=..., has_traffic_light: bool=...) -> None:
        ...