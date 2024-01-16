from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class NowRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class NowResponse(_message.Message):
    __slots__ = ['t']
    T_FIELD_NUMBER: _ClassVar[int]
    t: float

    def __init__(self, t: _Optional[float]=...) -> None:
        ...