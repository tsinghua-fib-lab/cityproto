from city.comm.interaction.gateway.v1 import gateway_pb2 as _gateway_pb2
from city.event.v1 import event_pb2 as _event_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class SetGatewayPowerStatusRequest(_message.Message):
    __slots__ = ['id', 'status']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: bool

    def __init__(self, id: _Optional[int]=..., status: bool=...) -> None:
        ...

class SetGatewayPowerStatusResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class SetGatewayRuinStatusRequest(_message.Message):
    __slots__ = ['id', 'status']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: bool

    def __init__(self, id: _Optional[int]=..., status: bool=...) -> None:
        ...

class SetGatewayRuinStatusResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetAllStatusRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetAllStatusResponse(_message.Message):
    __slots__ = ['stations']
    STATIONS_FIELD_NUMBER: _ClassVar[int]
    stations: _containers.RepeatedCompositeFieldContainer[_gateway_pb2.Station]

    def __init__(self, stations: _Optional[_Iterable[_Union[_gateway_pb2.Station, _Mapping]]]=...) -> None:
        ...

class GetRuinInfoRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RuinInfo(_message.Message):
    __slots__ = ['num', 'ratio']
    NUM_FIELD_NUMBER: _ClassVar[int]
    RATIO_FIELD_NUMBER: _ClassVar[int]
    num: int
    ratio: float

    def __init__(self, num: _Optional[int]=..., ratio: _Optional[float]=...) -> None:
        ...

class GetRuinInfoResponse(_message.Message):
    __slots__ = ['one', 'two', 'three']
    ONE_FIELD_NUMBER: _ClassVar[int]
    TWO_FIELD_NUMBER: _ClassVar[int]
    THREE_FIELD_NUMBER: _ClassVar[int]
    one: RuinInfo
    two: RuinInfo
    three: RuinInfo

    def __init__(self, one: _Optional[_Union[RuinInfo, _Mapping]]=..., two: _Optional[_Union[RuinInfo, _Mapping]]=..., three: _Optional[_Union[RuinInfo, _Mapping]]=...) -> None:
        ...

class GetEventsRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetEventsResponse(_message.Message):
    __slots__ = ['events']
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _event_pb2.Events

    def __init__(self, events: _Optional[_Union[_event_pb2.Events, _Mapping]]=...) -> None:
        ...