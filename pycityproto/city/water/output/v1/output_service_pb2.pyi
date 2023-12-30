from city.event.v1 import event_pb2 as _event_pb2
from city.water.output.v1 import output_pb2 as _output_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class OutputRequest(_message.Message):
    __slots__ = ['step', 'roads', 'detailed_roads', 'drainage_nodes', 'drainage_links', 'supply_nodes', 'supply_links', 'aois', 'drainage_metric', 'events', 'drainage_metrics', 'supply_metrics']
    STEP_FIELD_NUMBER: _ClassVar[int]
    ROADS_FIELD_NUMBER: _ClassVar[int]
    DETAILED_ROADS_FIELD_NUMBER: _ClassVar[int]
    DRAINAGE_NODES_FIELD_NUMBER: _ClassVar[int]
    DRAINAGE_LINKS_FIELD_NUMBER: _ClassVar[int]
    SUPPLY_NODES_FIELD_NUMBER: _ClassVar[int]
    SUPPLY_LINKS_FIELD_NUMBER: _ClassVar[int]
    AOIS_FIELD_NUMBER: _ClassVar[int]
    DRAINAGE_METRIC_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    DRAINAGE_METRICS_FIELD_NUMBER: _ClassVar[int]
    SUPPLY_METRICS_FIELD_NUMBER: _ClassVar[int]
    step: int
    roads: _containers.RepeatedCompositeFieldContainer[_output_pb2.Road]
    detailed_roads: _containers.RepeatedCompositeFieldContainer[_output_pb2.DetailedRoad]
    drainage_nodes: _containers.RepeatedCompositeFieldContainer[_output_pb2.Node]
    drainage_links: _containers.RepeatedCompositeFieldContainer[_output_pb2.Link]
    supply_nodes: _containers.RepeatedCompositeFieldContainer[_output_pb2.Node]
    supply_links: _containers.RepeatedCompositeFieldContainer[_output_pb2.Link]
    aois: _containers.RepeatedCompositeFieldContainer[_output_pb2.Aoi]
    drainage_metric: float
    events: _event_pb2.Events
    drainage_metrics: _output_pb2.DrainageMetrics
    supply_metrics: _output_pb2.SupplyMetrics

    def __init__(self, step: _Optional[int]=..., roads: _Optional[_Iterable[_Union[_output_pb2.Road, _Mapping]]]=..., detailed_roads: _Optional[_Iterable[_Union[_output_pb2.DetailedRoad, _Mapping]]]=..., drainage_nodes: _Optional[_Iterable[_Union[_output_pb2.Node, _Mapping]]]=..., drainage_links: _Optional[_Iterable[_Union[_output_pb2.Link, _Mapping]]]=..., supply_nodes: _Optional[_Iterable[_Union[_output_pb2.Node, _Mapping]]]=..., supply_links: _Optional[_Iterable[_Union[_output_pb2.Link, _Mapping]]]=..., aois: _Optional[_Iterable[_Union[_output_pb2.Aoi, _Mapping]]]=..., drainage_metric: _Optional[float]=..., events: _Optional[_Union[_event_pb2.Events, _Mapping]]=..., drainage_metrics: _Optional[_Union[_output_pb2.DrainageMetrics, _Mapping]]=..., supply_metrics: _Optional[_Union[_output_pb2.SupplyMetrics, _Mapping]]=...) -> None:
        ...

class OutputResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...