from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class Cost(_message.Message):
    __slots__ = ['id', 'cost', 'time']
    ID_FIELD_NUMBER: _ClassVar[int]
    COST_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    id: int
    cost: float
    time: float

    def __init__(self, id: _Optional[int]=..., cost: _Optional[float]=..., time: _Optional[float]=...) -> None:
        ...