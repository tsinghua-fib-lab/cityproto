from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Node(_message.Message):
    __slots__ = ['id', 'status']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: float

    def __init__(self, id: _Optional[int]=..., status: _Optional[float]=...) -> None:
        ...

class Influence(_message.Message):
    __slots__ = ['source_id', 'target_id', 'strength']
    SOURCE_ID_FIELD_NUMBER: _ClassVar[int]
    TARGET_ID_FIELD_NUMBER: _ClassVar[int]
    STRENGTH_FIELD_NUMBER: _ClassVar[int]
    source_id: int
    target_id: int
    strength: float

    def __init__(self, source_id: _Optional[int]=..., target_id: _Optional[int]=..., strength: _Optional[float]=...) -> None:
        ...

class Heatmap(_message.Message):
    __slots__ = ['num_rows', 'num_columns', 'strength']
    NUM_ROWS_FIELD_NUMBER: _ClassVar[int]
    NUM_COLUMNS_FIELD_NUMBER: _ClassVar[int]
    STRENGTH_FIELD_NUMBER: _ClassVar[int]
    num_rows: int
    num_columns: int
    strength: _containers.RepeatedScalarFieldContainer[float]

    def __init__(self, num_rows: _Optional[int]=..., num_columns: _Optional[int]=..., strength: _Optional[_Iterable[float]]=...) -> None:
        ...

class Stat(_message.Message):
    __slots__ = ['crowd_cnts', 'crowd_ratios', 'key_nodes']
    CROWD_CNTS_FIELD_NUMBER: _ClassVar[int]
    CROWD_RATIOS_FIELD_NUMBER: _ClassVar[int]
    KEY_NODES_FIELD_NUMBER: _ClassVar[int]
    crowd_cnts: _containers.RepeatedScalarFieldContainer[int]
    crowd_ratios: _containers.RepeatedScalarFieldContainer[float]
    key_nodes: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, crowd_cnts: _Optional[_Iterable[int]]=..., crowd_ratios: _Optional[_Iterable[float]]=..., key_nodes: _Optional[_Iterable[int]]=...) -> None:
        ...

class Content(_message.Message):
    __slots__ = ['id', 'text']
    ID_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    id: int
    text: str

    def __init__(self, id: _Optional[int]=..., text: _Optional[str]=...) -> None:
        ...

class NodeMeta(_message.Message):
    __slots__ = ['id', 'lng', 'lat', 'level']
    ID_FIELD_NUMBER: _ClassVar[int]
    LNG_FIELD_NUMBER: _ClassVar[int]
    LAT_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    id: int
    lng: float
    lat: float
    level: int

    def __init__(self, id: _Optional[int]=..., lng: _Optional[float]=..., lat: _Optional[float]=..., level: _Optional[int]=...) -> None:
        ...

class NodesMeta(_message.Message):
    __slots__ = ['nodes']
    NODES_FIELD_NUMBER: _ClassVar[int]
    nodes: _containers.RepeatedCompositeFieldContainer[NodeMeta]

    def __init__(self, nodes: _Optional[_Iterable[_Union[NodeMeta, _Mapping]]]=...) -> None:
        ...

class Group(_message.Message):
    __slots__ = ['id', 'size', 'changes']
    ID_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    CHANGES_FIELD_NUMBER: _ClassVar[int]
    id: int
    size: int
    changes: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, id: _Optional[int]=..., size: _Optional[int]=..., changes: _Optional[_Iterable[int]]=...) -> None:
        ...