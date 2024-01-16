from city.geo.v2 import geo_pb2 as _geo_pb2
from city.routing.v2 import cost_pb2 as _cost_pb2
from city.routing.v2 import routing_pb2 as _routing_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetRouteRequest(_message.Message):
    __slots__ = ['type', 'start', 'end', 'time']
    TYPE_FIELD_NUMBER: _ClassVar[int]
    START_FIELD_NUMBER: _ClassVar[int]
    END_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    type: _routing_pb2.RouteType
    start: _geo_pb2.Position
    end: _geo_pb2.Position
    time: float

    def __init__(self, type: _Optional[_Union[_routing_pb2.RouteType, str]]=..., start: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., end: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., time: _Optional[float]=...) -> None:
        ...

class GetRouteResponse(_message.Message):
    __slots__ = ['journeys']
    JOURNEYS_FIELD_NUMBER: _ClassVar[int]
    journeys: _containers.RepeatedCompositeFieldContainer[_routing_pb2.Journey]

    def __init__(self, journeys: _Optional[_Iterable[_Union[_routing_pb2.Journey, _Mapping]]]=...) -> None:
        ...

class SetDrivingCostsRequest(_message.Message):
    __slots__ = ['costs']
    COSTS_FIELD_NUMBER: _ClassVar[int]
    costs: _containers.RepeatedCompositeFieldContainer[_cost_pb2.Cost]

    def __init__(self, costs: _Optional[_Iterable[_Union[_cost_pb2.Cost, _Mapping]]]=...) -> None:
        ...

class SetDrivingCostsResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetDrivingCostsRequest(_message.Message):
    __slots__ = ['costs']
    COSTS_FIELD_NUMBER: _ClassVar[int]
    costs: _containers.RepeatedCompositeFieldContainer[_cost_pb2.Cost]

    def __init__(self, costs: _Optional[_Iterable[_Union[_cost_pb2.Cost, _Mapping]]]=...) -> None:
        ...

class GetDrivingCostsResponse(_message.Message):
    __slots__ = ['costs']
    COSTS_FIELD_NUMBER: _ClassVar[int]
    costs: _containers.RepeatedCompositeFieldContainer[_cost_pb2.Cost]

    def __init__(self, costs: _Optional[_Iterable[_Union[_cost_pb2.Cost, _Mapping]]]=...) -> None:
        ...