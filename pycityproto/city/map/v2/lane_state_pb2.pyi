from city.map.v2 import light_pb2 as _light_pb2
from city.person.v2 import motion_pb2 as _motion_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class LaneState(_message.Message):
    __slots__ = ['id', 'persons', 'avg_v', 'restriction', 'light_state', 'in_vehicle_cnt', 'out_vehicle_cnt', 'vehicle_cnt', 'total_queuing_vehicle_cnt', 'total_queuing_time', 'avg_queuing_time']
    ID_FIELD_NUMBER: _ClassVar[int]
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    AVG_V_FIELD_NUMBER: _ClassVar[int]
    RESTRICTION_FIELD_NUMBER: _ClassVar[int]
    LIGHT_STATE_FIELD_NUMBER: _ClassVar[int]
    IN_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    OUT_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    TOTAL_QUEUING_VEHICLE_CNT_FIELD_NUMBER: _ClassVar[int]
    TOTAL_QUEUING_TIME_FIELD_NUMBER: _ClassVar[int]
    AVG_QUEUING_TIME_FIELD_NUMBER: _ClassVar[int]
    id: int
    persons: _containers.RepeatedCompositeFieldContainer[_motion_pb2.PersonMotion]
    avg_v: float
    restriction: bool
    light_state: _light_pb2.LightState
    in_vehicle_cnt: int
    out_vehicle_cnt: int
    vehicle_cnt: int
    total_queuing_vehicle_cnt: int
    total_queuing_time: float
    avg_queuing_time: float

    def __init__(self, id: _Optional[int]=..., persons: _Optional[_Iterable[_Union[_motion_pb2.PersonMotion, _Mapping]]]=..., avg_v: _Optional[float]=..., restriction: bool=..., light_state: _Optional[_Union[_light_pb2.LightState, str]]=..., in_vehicle_cnt: _Optional[int]=..., out_vehicle_cnt: _Optional[int]=..., vehicle_cnt: _Optional[int]=..., total_queuing_vehicle_cnt: _Optional[int]=..., total_queuing_time: _Optional[float]=..., avg_queuing_time: _Optional[float]=...) -> None:
        ...