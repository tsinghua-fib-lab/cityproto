from city.comm.output.v1 import output_pb2 as _output_pb2
from city.event.v1 import event_pb2 as _event_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class OutputRequest(_message.Message):
    __slots__ = ['step', 'nodes', 'base_stations', 'signal_heatmap', 'small_signal_heatmap', 'persons', 'aois', 'events', 'statistics']
    STEP_FIELD_NUMBER: _ClassVar[int]
    NODES_FIELD_NUMBER: _ClassVar[int]
    BASE_STATIONS_FIELD_NUMBER: _ClassVar[int]
    SIGNAL_HEATMAP_FIELD_NUMBER: _ClassVar[int]
    SMALL_SIGNAL_HEATMAP_FIELD_NUMBER: _ClassVar[int]
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    AOIS_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    STATISTICS_FIELD_NUMBER: _ClassVar[int]
    step: int
    nodes: _containers.RepeatedCompositeFieldContainer[_output_pb2.Node]
    base_stations: _containers.RepeatedCompositeFieldContainer[_output_pb2.BaseStation]
    signal_heatmap: _output_pb2.Signal
    small_signal_heatmap: _output_pb2.Signal
    persons: _containers.RepeatedCompositeFieldContainer[_output_pb2.Person]
    aois: _containers.RepeatedCompositeFieldContainer[_output_pb2.Aoi]
    events: _event_pb2.Events
    statistics: _output_pb2.Statistics

    def __init__(self, step: _Optional[int]=..., nodes: _Optional[_Iterable[_Union[_output_pb2.Node, _Mapping]]]=..., base_stations: _Optional[_Iterable[_Union[_output_pb2.BaseStation, _Mapping]]]=..., signal_heatmap: _Optional[_Union[_output_pb2.Signal, _Mapping]]=..., small_signal_heatmap: _Optional[_Union[_output_pb2.Signal, _Mapping]]=..., persons: _Optional[_Iterable[_Union[_output_pb2.Person, _Mapping]]]=..., aois: _Optional[_Iterable[_Union[_output_pb2.Aoi, _Mapping]]]=..., events: _Optional[_Union[_event_pb2.Events, _Mapping]]=..., statistics: _Optional[_Union[_output_pb2.Statistics, _Mapping]]=...) -> None:
        ...

class OutputResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...