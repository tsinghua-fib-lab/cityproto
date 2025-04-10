from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class PedestrianAction(_message.Message):
    __slots__ = ['id', 'vx', 'vy']
    ID_FIELD_NUMBER: _ClassVar[int]
    VX_FIELD_NUMBER: _ClassVar[int]
    VY_FIELD_NUMBER: _ClassVar[int]
    id: int
    vx: float
    vy: float

    def __init__(self, id: _Optional[int]=..., vx: _Optional[float]=..., vy: _Optional[float]=...) -> None:
        ...