from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class GetEtcdRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetEtcdResponse(_message.Message):
    __slots__ = ['port']
    PORT_FIELD_NUMBER: _ClassVar[int]
    port: int

    def __init__(self, port: _Optional[int]=...) -> None:
        ...

class StepRequest(_message.Message):
    __slots__ = ['name', 'step']
    NAME_FIELD_NUMBER: _ClassVar[int]
    STEP_FIELD_NUMBER: _ClassVar[int]
    name: str
    step: int

    def __init__(self, name: _Optional[str]=..., step: _Optional[int]=...) -> None:
        ...

class StepResponse(_message.Message):
    __slots__ = ['close']
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    close: bool

    def __init__(self, close: bool=...) -> None:
        ...