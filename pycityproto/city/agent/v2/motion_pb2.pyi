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
STATUS_UNSPECIFIED: Status
STATUS_SLEEP: Status
STATUS_DRIVING: Status
STATUS_WALKING: Status
STATUS_CROWD: Status
STATUS_PASSENGER: Status
STATUS_WAIT_ROUTE: Status

class AgentMotion(_message.Message):
    __slots__ = ['id', 'status', 'position', 'v', 'direction', 'activity']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    V_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    ACTIVITY_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: Status
    position: _geo_pb2.Position
    v: float
    direction: float
    activity: str

    def __init__(self, id: _Optional[int]=..., status: _Optional[_Union[Status, str]]=..., position: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., v: _Optional[float]=..., direction: _Optional[float]=..., activity: _Optional[str]=...) -> None:
        ...