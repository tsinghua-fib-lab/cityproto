from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Edge(_message.Message):
    __slots__ = ['source', 'target']
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    TARGET_FIELD_NUMBER: _ClassVar[int]
    source: int
    target: int

    def __init__(self, source: _Optional[int]=..., target: _Optional[int]=...) -> None:
        ...

class Edges(_message.Message):
    __slots__ = ['edges']
    EDGES_FIELD_NUMBER: _ClassVar[int]
    edges: _containers.RepeatedCompositeFieldContainer[Edge]

    def __init__(self, edges: _Optional[_Iterable[_Union[Edge, _Mapping]]]=...) -> None:
        ...