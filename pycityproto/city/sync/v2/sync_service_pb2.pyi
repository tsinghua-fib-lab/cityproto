from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class SetURLRequest(_message.Message):
    __slots__ = ['name', 'url']
    NAME_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    name: str
    url: str

    def __init__(self, name: _Optional[str]=..., url: _Optional[str]=...) -> None:
        ...

class SetURLResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetURLRequest(_message.Message):
    __slots__ = ['name']
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str

    def __init__(self, name: _Optional[str]=...) -> None:
        ...

class GetURLResponse(_message.Message):
    __slots__ = ['url']
    URL_FIELD_NUMBER: _ClassVar[int]
    url: str

    def __init__(self, url: _Optional[str]=...) -> None:
        ...

class EnterStepSyncRequest(_message.Message):
    __slots__ = ['name']
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str

    def __init__(self, name: _Optional[str]=...) -> None:
        ...

class EnterStepSyncResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ExitStepSyncRequest(_message.Message):
    __slots__ = ['name', 'close']
    NAME_FIELD_NUMBER: _ClassVar[int]
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    name: str
    close: bool

    def __init__(self, name: _Optional[str]=..., close: bool=...) -> None:
        ...

class ExitStepSyncResponse(_message.Message):
    __slots__ = ['close']
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    close: bool

    def __init__(self, close: bool=...) -> None:
        ...