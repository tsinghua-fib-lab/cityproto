from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Reason(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    REASON_UNSPECIFIED: _ClassVar[Reason]
    REASON_RUIN: _ClassVar[Reason]
    REASON_CASCADE: _ClassVar[Reason]
REASON_UNSPECIFIED: Reason
REASON_RUIN: Reason
REASON_CASCADE: Reason

class Station(_message.Message):
    __slots__ = ['id', 'status', 'reason']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    REASON_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: bool
    reason: Reason

    def __init__(self, id: _Optional[int]=..., status: bool=..., reason: _Optional[_Union[Reason, str]]=...) -> None:
        ...