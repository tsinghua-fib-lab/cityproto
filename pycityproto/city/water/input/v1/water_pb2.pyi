from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class RainPeriod(_message.Message):
    __slots__ = ['start', 'rainfall']
    START_FIELD_NUMBER: _ClassVar[int]
    RAINFALL_FIELD_NUMBER: _ClassVar[int]
    start: int
    rainfall: float

    def __init__(self, start: _Optional[int]=..., rainfall: _Optional[float]=...) -> None:
        ...

class Rain(_message.Message):
    __slots__ = ['rains']
    RAINS_FIELD_NUMBER: _ClassVar[int]
    rains: _containers.RepeatedCompositeFieldContainer[RainPeriod]

    def __init__(self, rains: _Optional[_Iterable[_Union[RainPeriod, _Mapping]]]=...) -> None:
        ...