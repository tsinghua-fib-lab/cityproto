from city.geo.v2 import geo_pb2 as _geo_pb2
from city.trip.v2 import trip_pb2 as _trip_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Education(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    EDUCATION_UNSPECIFIED: _ClassVar[Education]
    EDUCATION_DOCTOR: _ClassVar[Education]
    EDUCATION_MASTER: _ClassVar[Education]
    EDUCATION_BACHELOR: _ClassVar[Education]
    EDUCATION_HIGH_SCHOOL: _ClassVar[Education]
    EDUCATION_JUNIOR_HIGH_SCHOOL: _ClassVar[Education]
    EDUCATION_PRIMARY_SCHOOL: _ClassVar[Education]
    EDUCATION_COLLEGE: _ClassVar[Education]

class Gender(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    GENDER_UNSPECIFIED: _ClassVar[Gender]
    GENDER_MALE: _ClassVar[Gender]
    GENDER_FEMALE: _ClassVar[Gender]

class Consumption(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    CONSUMPTION_UNSPECIFIED: _ClassVar[Consumption]
    CONSUMPTION_LOW: _ClassVar[Consumption]
    CONSUMPTION_RELATIVELY_LOW: _ClassVar[Consumption]
    CONSUMPTION_MEDIUM: _ClassVar[Consumption]
    CONSUMPTION_RELATIVELY_HIGH: _ClassVar[Consumption]
    CONSUMPTION_HIGH: _ClassVar[Consumption]
EDUCATION_UNSPECIFIED: Education
EDUCATION_DOCTOR: Education
EDUCATION_MASTER: Education
EDUCATION_BACHELOR: Education
EDUCATION_HIGH_SCHOOL: Education
EDUCATION_JUNIOR_HIGH_SCHOOL: Education
EDUCATION_PRIMARY_SCHOOL: Education
EDUCATION_COLLEGE: Education
GENDER_UNSPECIFIED: Gender
GENDER_MALE: Gender
GENDER_FEMALE: Gender
CONSUMPTION_UNSPECIFIED: Consumption
CONSUMPTION_LOW: Consumption
CONSUMPTION_RELATIVELY_LOW: Consumption
CONSUMPTION_MEDIUM: Consumption
CONSUMPTION_RELATIVELY_HIGH: Consumption
CONSUMPTION_HIGH: Consumption

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
    __slots__ = ['lane_change_length', 'min_gap', 'model']
    LANE_CHANGE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    MIN_GAP_FIELD_NUMBER: _ClassVar[int]
    MODEL_FIELD_NUMBER: _ClassVar[int]
    lane_change_length: float
    min_gap: float
    model: str

    def __init__(self, lane_change_length: _Optional[float]=..., min_gap: _Optional[float]=..., model: _Optional[str]=...) -> None:
        ...

class BusAttribute(_message.Message):
    __slots__ = ['line_id', 'capacity', 'model']
    LINE_ID_FIELD_NUMBER: _ClassVar[int]
    CAPACITY_FIELD_NUMBER: _ClassVar[int]
    MODEL_FIELD_NUMBER: _ClassVar[int]
    line_id: int
    capacity: int
    model: str

    def __init__(self, line_id: _Optional[int]=..., capacity: _Optional[int]=..., model: _Optional[str]=...) -> None:
        ...

class PedestrianAttribute(_message.Message):
    __slots__ = ['speed', 'model']
    SPEED_FIELD_NUMBER: _ClassVar[int]
    MODEL_FIELD_NUMBER: _ClassVar[int]
    speed: float
    model: str

    def __init__(self, speed: _Optional[float]=..., model: _Optional[str]=...) -> None:
        ...

class BikeAttribute(_message.Message):
    __slots__ = ['speed', 'model']
    SPEED_FIELD_NUMBER: _ClassVar[int]
    MODEL_FIELD_NUMBER: _ClassVar[int]
    speed: float
    model: str

    def __init__(self, speed: _Optional[float]=..., model: _Optional[str]=...) -> None:
        ...

class PersonProfile(_message.Message):
    __slots__ = ['age', 'education', 'gender', 'consumption']
    AGE_FIELD_NUMBER: _ClassVar[int]
    EDUCATION_FIELD_NUMBER: _ClassVar[int]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    CONSUMPTION_FIELD_NUMBER: _ClassVar[int]
    age: int
    education: Education
    gender: Gender
    consumption: Consumption

    def __init__(self, age: _Optional[int]=..., education: _Optional[_Union[Education, str]]=..., gender: _Optional[_Union[Gender, str]]=..., consumption: _Optional[_Union[Consumption, str]]=...) -> None:
        ...

class Person(_message.Message):
    __slots__ = ['id', 'attribute', 'home', 'schedules', 'vehicle_attribute', 'bus_attribute', 'pedestrian_attribute', 'bike_attribute', 'labels', 'profile', 'work']

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
    PEDESTRIAN_ATTRIBUTE_FIELD_NUMBER: _ClassVar[int]
    BIKE_ATTRIBUTE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    PROFILE_FIELD_NUMBER: _ClassVar[int]
    WORK_FIELD_NUMBER: _ClassVar[int]
    id: int
    attribute: PersonAttribute
    home: _geo_pb2.Position
    schedules: _containers.RepeatedCompositeFieldContainer[_trip_pb2.Schedule]
    vehicle_attribute: VehicleAttribute
    bus_attribute: BusAttribute
    pedestrian_attribute: PedestrianAttribute
    bike_attribute: BikeAttribute
    labels: _containers.ScalarMap[str, str]
    profile: PersonProfile
    work: _geo_pb2.Position

    def __init__(self, id: _Optional[int]=..., attribute: _Optional[_Union[PersonAttribute, _Mapping]]=..., home: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., schedules: _Optional[_Iterable[_Union[_trip_pb2.Schedule, _Mapping]]]=..., vehicle_attribute: _Optional[_Union[VehicleAttribute, _Mapping]]=..., bus_attribute: _Optional[_Union[BusAttribute, _Mapping]]=..., pedestrian_attribute: _Optional[_Union[PedestrianAttribute, _Mapping]]=..., bike_attribute: _Optional[_Union[BikeAttribute, _Mapping]]=..., labels: _Optional[_Mapping[str, str]]=..., profile: _Optional[_Union[PersonProfile, _Mapping]]=..., work: _Optional[_Union[_geo_pb2.Position, _Mapping]]=...) -> None:
        ...

class Persons(_message.Message):
    __slots__ = ['persons']
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    persons: _containers.RepeatedCompositeFieldContainer[Person]

    def __init__(self, persons: _Optional[_Iterable[_Union[Person, _Mapping]]]=...) -> None:
        ...