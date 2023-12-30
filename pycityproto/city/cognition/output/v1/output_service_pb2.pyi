from city.cognition.output.v1 import output_pb2 as _output_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class OutputRequest(_message.Message):
    __slots__ = ['step', 'nodes', 'influences', 'heatmap', 'stat', 'contents', 'group_influences', 'groups']
    STEP_FIELD_NUMBER: _ClassVar[int]
    NODES_FIELD_NUMBER: _ClassVar[int]
    INFLUENCES_FIELD_NUMBER: _ClassVar[int]
    HEATMAP_FIELD_NUMBER: _ClassVar[int]
    STAT_FIELD_NUMBER: _ClassVar[int]
    CONTENTS_FIELD_NUMBER: _ClassVar[int]
    GROUP_INFLUENCES_FIELD_NUMBER: _ClassVar[int]
    GROUPS_FIELD_NUMBER: _ClassVar[int]
    step: int
    nodes: _containers.RepeatedCompositeFieldContainer[_output_pb2.Node]
    influences: _containers.RepeatedCompositeFieldContainer[_output_pb2.Influence]
    heatmap: _output_pb2.Heatmap
    stat: _output_pb2.Stat
    contents: _containers.RepeatedCompositeFieldContainer[_output_pb2.Content]
    group_influences: _containers.RepeatedCompositeFieldContainer[_output_pb2.Influence]
    groups: _containers.RepeatedCompositeFieldContainer[_output_pb2.Group]

    def __init__(self, step: _Optional[int]=..., nodes: _Optional[_Iterable[_Union[_output_pb2.Node, _Mapping]]]=..., influences: _Optional[_Iterable[_Union[_output_pb2.Influence, _Mapping]]]=..., heatmap: _Optional[_Union[_output_pb2.Heatmap, _Mapping]]=..., stat: _Optional[_Union[_output_pb2.Stat, _Mapping]]=..., contents: _Optional[_Iterable[_Union[_output_pb2.Content, _Mapping]]]=..., group_influences: _Optional[_Iterable[_Union[_output_pb2.Influence, _Mapping]]]=..., groups: _Optional[_Iterable[_Union[_output_pb2.Group, _Mapping]]]=...) -> None:
        ...

class OutputResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...