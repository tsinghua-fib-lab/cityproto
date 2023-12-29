from city.traffic.interaction.agent.v2 import agent_pb2 as _agent_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class State(_message.Message):
    __slots__ = ('id', 'persons', 'avg_v', 'restriction')
    ID_FIELD_NUMBER: _ClassVar[int]
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    AVG_V_FIELD_NUMBER: _ClassVar[int]
    RESTRICTION_FIELD_NUMBER: _ClassVar[int]
    id: int
    persons: _containers.RepeatedCompositeFieldContainer[_agent_pb2.AgentRuntime]
    avg_v: float
    restriction: bool

    def __init__(self, id: _Optional[int]=..., persons: _Optional[_Iterable[_Union[_agent_pb2.AgentRuntime, _Mapping]]]=..., avg_v: _Optional[float]=..., restriction: bool=...) -> None:
        ...