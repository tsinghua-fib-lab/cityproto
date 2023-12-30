from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class Message(_message.Message):
    __slots__ = ['to', 'message', 't']
    FROM_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    T_FIELD_NUMBER: _ClassVar[int]
    to: int
    message: str
    t: float

    def __init__(self, to: _Optional[int]=..., message: _Optional[str]=..., t: _Optional[float]=..., **kwargs) -> None:
        ...