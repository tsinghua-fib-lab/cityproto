from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class BaseRuntime(_message.Message):
    __slots__ = ['x', 'y', 'v', 'direction']
    X_FIELD_NUMBER: _ClassVar[int]
    Y_FIELD_NUMBER: _ClassVar[int]
    V_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    x: float
    y: float
    v: float
    direction: float

    def __init__(self, x: _Optional[float]=..., y: _Optional[float]=..., v: _Optional[float]=..., direction: _Optional[float]=...) -> None:
        ...

class BaseRuntimeOnRoad(_message.Message):
    __slots__ = ['lane_id', 's']
    LANE_ID_FIELD_NUMBER: _ClassVar[int]
    S_FIELD_NUMBER: _ClassVar[int]
    lane_id: int
    s: float

    def __init__(self, lane_id: _Optional[int]=..., s: _Optional[float]=...) -> None:
        ...