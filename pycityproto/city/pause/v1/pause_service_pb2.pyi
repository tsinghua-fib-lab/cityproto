from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class PauseRequest(_message.Message):
    __slots__ = ['name']
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str

    def __init__(self, name: _Optional[str]=...) -> None:
        ...

class PauseResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ResumeRequest(_message.Message):
    __slots__ = ['name']
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str

    def __init__(self, name: _Optional[str]=...) -> None:
        ...

class ResumeResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...