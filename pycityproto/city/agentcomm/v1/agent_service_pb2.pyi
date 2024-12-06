from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class CommunicateRequest(_message.Message):
    __slots__ = ['source_agent_id', 'target_agent_id', 'data']
    SOURCE_AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    TARGET_AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    source_agent_id: str
    target_agent_id: str
    data: str

    def __init__(self, source_agent_id: _Optional[str]=..., target_agent_id: _Optional[str]=..., data: _Optional[str]=...) -> None:
        ...

class CommunicateResponse(_message.Message):
    __slots__ = ['message']
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str

    def __init__(self, message: _Optional[str]=...) -> None:
        ...