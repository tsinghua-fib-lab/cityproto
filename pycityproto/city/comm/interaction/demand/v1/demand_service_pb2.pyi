from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class SetDemandStatusRequest(_message.Message):
    __slots__ = ['multiple_times', 'power_times']
    MULTIPLE_TIMES_FIELD_NUMBER: _ClassVar[int]
    POWER_TIMES_FIELD_NUMBER: _ClassVar[int]
    multiple_times: float
    power_times: float

    def __init__(self, multiple_times: _Optional[float]=..., power_times: _Optional[float]=...) -> None:
        ...

class SetDemandStatusResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...