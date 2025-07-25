from city.geo.v2 import geo_pb2 as _geo_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Status(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    STATUS_UNSPECIFIED: _ClassVar[Status]
    STATUS_SLEEP: _ClassVar[Status]
    STATUS_DRIVING: _ClassVar[Status]
    STATUS_WALKING: _ClassVar[Status]
    STATUS_CROWD: _ClassVar[Status]
    STATUS_PASSENGER: _ClassVar[Status]
    STATUS_WAIT_ROUTE: _ClassVar[Status]
    STATUS_WAIT_BUS: _ClassVar[Status]
    STATUS_RAIL_TRANSIT: _ClassVar[Status]
    STATUS_WAIT_TAXI: _ClassVar[Status]
STATUS_UNSPECIFIED: Status
STATUS_SLEEP: Status
STATUS_DRIVING: Status
STATUS_WALKING: Status
STATUS_CROWD: Status
STATUS_PASSENGER: Status
STATUS_WAIT_ROUTE: Status
STATUS_WAIT_BUS: Status
STATUS_RAIL_TRANSIT: Status
STATUS_WAIT_TAXI: Status

class PersonMotion(_message.Message):
    __slots__ = ['id', 'status', 'position', 'v', 'direction', 'activity', 'l', 'a', 'is_queuing_on_lane', 'queuing_time_on_cur_lane', 'num_passengers']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    V_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    ACTIVITY_FIELD_NUMBER: _ClassVar[int]
    L_FIELD_NUMBER: _ClassVar[int]
    A_FIELD_NUMBER: _ClassVar[int]
    IS_QUEUING_ON_LANE_FIELD_NUMBER: _ClassVar[int]
    QUEUING_TIME_ON_CUR_LANE_FIELD_NUMBER: _ClassVar[int]
    NUM_PASSENGERS_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: Status
    position: _geo_pb2.Position
    v: float
    direction: float
    activity: str
    l: float
    a: float
    is_queuing_on_lane: bool
    queuing_time_on_cur_lane: float
    num_passengers: int

    def __init__(self, id: _Optional[int]=..., status: _Optional[_Union[Status, str]]=..., position: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., v: _Optional[float]=..., direction: _Optional[float]=..., activity: _Optional[str]=..., l: _Optional[float]=..., a: _Optional[float]=..., is_queuing_on_lane: bool=..., queuing_time_on_cur_lane: _Optional[float]=..., num_passengers: _Optional[int]=...) -> None:
        ...