from city.geo.v2 import geo_pb2 as _geo_pb2
from city.trip.v2 import trip_pb2 as _trip_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class PersonAttribute(_message.Message):
    __slots__ = ['length', 'width', 'max_speed', 'max_acceleration', 'max_braking_acceleration', 'usual_acceleration', 'usual_braking_acceleration']
    LENGTH_FIELD_NUMBER: _ClassVar[int]
    WIDTH_FIELD_NUMBER: _ClassVar[int]
    MAX_SPEED_FIELD_NUMBER: _ClassVar[int]
    MAX_ACCELERATION_FIELD_NUMBER: _ClassVar[int]
    MAX_BRAKING_ACCELERATION_FIELD_NUMBER: _ClassVar[int]
    USUAL_ACCELERATION_FIELD_NUMBER: _ClassVar[int]
    USUAL_BRAKING_ACCELERATION_FIELD_NUMBER: _ClassVar[int]
    length: float
    width: float
    max_speed: float
    max_acceleration: float
    max_braking_acceleration: float
    usual_acceleration: float
    usual_braking_acceleration: float

    def __init__(self, length: _Optional[float]=..., width: _Optional[float]=..., max_speed: _Optional[float]=..., max_acceleration: _Optional[float]=..., max_braking_acceleration: _Optional[float]=..., usual_acceleration: _Optional[float]=..., usual_braking_acceleration: _Optional[float]=...) -> None:
        ...

class VehicleAttribute(_message.Message):
    __slots__ = ['lane_change_length', 'min_gap']
    LANE_CHANGE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    MIN_GAP_FIELD_NUMBER: _ClassVar[int]
    lane_change_length: float
    min_gap: float

    def __init__(self, lane_change_length: _Optional[float]=..., min_gap: _Optional[float]=...) -> None:
        ...

class BusAttribute(_message.Message):
    __slots__ = ['line_id', 'capacity']
    LINE_ID_FIELD_NUMBER: _ClassVar[int]
    CAPACITY_FIELD_NUMBER: _ClassVar[int]
    line_id: int
    capacity: int

    def __init__(self, line_id: _Optional[int]=..., capacity: _Optional[int]=...) -> None:
        ...

class BikeAttribute(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class Person(_message.Message):
    __slots__ = ['id', 'attribute', 'home', 'schedules', 'vehicle_attribute', 'bus_attribute', 'bike_attribute', 'labels']

    class LabelsEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str

        def __init__(self, key: _Optional[str]=..., value: _Optional[str]=...) -> None:
            ...
    ID_FIELD_NUMBER: _ClassVar[int]
    ATTRIBUTE_FIELD_NUMBER: _ClassVar[int]
    HOME_FIELD_NUMBER: _ClassVar[int]
    SCHEDULES_FIELD_NUMBER: _ClassVar[int]
    VEHICLE_ATTRIBUTE_FIELD_NUMBER: _ClassVar[int]
    BUS_ATTRIBUTE_FIELD_NUMBER: _ClassVar[int]
    BIKE_ATTRIBUTE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    id: int
    attribute: PersonAttribute
    home: _geo_pb2.Position
    schedules: _containers.RepeatedCompositeFieldContainer[_trip_pb2.Schedule]
    vehicle_attribute: VehicleAttribute
    bus_attribute: BusAttribute
    bike_attribute: BikeAttribute
    labels: _containers.ScalarMap[str, str]

    def __init__(self, id: _Optional[int]=..., attribute: _Optional[_Union[PersonAttribute, _Mapping]]=..., home: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., schedules: _Optional[_Iterable[_Union[_trip_pb2.Schedule, _Mapping]]]=..., vehicle_attribute: _Optional[_Union[VehicleAttribute, _Mapping]]=..., bus_attribute: _Optional[_Union[BusAttribute, _Mapping]]=..., bike_attribute: _Optional[_Union[BikeAttribute, _Mapping]]=..., labels: _Optional[_Mapping[str, str]]=...) -> None:
        ...

class Persons(_message.Message):
    __slots__ = ['persons']
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    persons: _containers.RepeatedCompositeFieldContainer[Person]

    def __init__(self, persons: _Optional[_Iterable[_Union[Person, _Mapping]]]=...) -> None:
        ...