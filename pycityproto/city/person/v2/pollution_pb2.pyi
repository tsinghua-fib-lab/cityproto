from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class PollutionStatistics(_message.Message):
    __slots__ = ['co2', 'pm', 'voc', 'nox']
    CO2_FIELD_NUMBER: _ClassVar[int]
    PM_FIELD_NUMBER: _ClassVar[int]
    VOC_FIELD_NUMBER: _ClassVar[int]
    NOX_FIELD_NUMBER: _ClassVar[int]
    co2: float
    pm: float
    voc: float
    nox: float

    def __init__(self, co2: _Optional[float]=..., pm: _Optional[float]=..., voc: _Optional[float]=..., nox: _Optional[float]=...) -> None:
        ...