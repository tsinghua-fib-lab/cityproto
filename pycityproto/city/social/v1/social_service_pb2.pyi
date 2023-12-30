from city.social.v1 import message_pb2 as _message_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class SendRequest(_message.Message):
    __slots__ = ['messages']
    MESSAGES_FIELD_NUMBER: _ClassVar[int]
    messages: _containers.RepeatedCompositeFieldContainer[_message_pb2.Message]

    def __init__(self, messages: _Optional[_Iterable[_Union[_message_pb2.Message, _Mapping]]]=...) -> None:
        ...

class SendResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class ReceiveRequest(_message.Message):
    __slots__ = ['id']
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int

    def __init__(self, id: _Optional[int]=...) -> None:
        ...

class ReceiveResponse(_message.Message):
    __slots__ = ['messages']
    MESSAGES_FIELD_NUMBER: _ClassVar[int]
    messages: _containers.RepeatedCompositeFieldContainer[_message_pb2.Message]

    def __init__(self, messages: _Optional[_Iterable[_Union[_message_pb2.Message, _Mapping]]]=...) -> None:
        ...