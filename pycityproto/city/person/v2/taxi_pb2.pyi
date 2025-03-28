from city.geo.v2 import geo_pb2 as _geo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class AllocationPlanType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    ALLOCATION_PLAN_TYPE_UNSPECIFIED: _ClassVar[AllocationPlanType]
    ALLOCATION_PLAN_TYPE_PICK_UP: _ClassVar[AllocationPlanType]
    ALLOCATION_PLAN_TYPE_DELIVER: _ClassVar[AllocationPlanType]
ALLOCATION_PLAN_TYPE_UNSPECIFIED: AllocationPlanType
ALLOCATION_PLAN_TYPE_PICK_UP: AllocationPlanType
ALLOCATION_PLAN_TYPE_DELIVER: AllocationPlanType

class RequestOrderInfo(_message.Message):
    __slots__ = ['person_id', 'request_time', 'order_id', 'departure', 'destination']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    REQUEST_TIME_FIELD_NUMBER: _ClassVar[int]
    ORDER_ID_FIELD_NUMBER: _ClassVar[int]
    DEPARTURE_FIELD_NUMBER: _ClassVar[int]
    DESTINATION_FIELD_NUMBER: _ClassVar[int]
    person_id: int
    request_time: float
    order_id: int
    departure: _geo_pb2.Position
    destination: _geo_pb2.Position

    def __init__(self, person_id: _Optional[int]=..., request_time: _Optional[float]=..., order_id: _Optional[int]=..., departure: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., destination: _Optional[_Union[_geo_pb2.Position, _Mapping]]=...) -> None:
        ...

class OrderAllocationPlan(_message.Message):
    __slots__ = ['order_ids', 'taxi_id', 'type', 'pick_up_person_ids', 'deliver_person_ids']
    ORDER_IDS_FIELD_NUMBER: _ClassVar[int]
    TAXI_ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    PICK_UP_PERSON_IDS_FIELD_NUMBER: _ClassVar[int]
    DELIVER_PERSON_IDS_FIELD_NUMBER: _ClassVar[int]
    order_ids: _containers.RepeatedScalarFieldContainer[int]
    taxi_id: int
    type: AllocationPlanType
    pick_up_person_ids: _containers.RepeatedScalarFieldContainer[int]
    deliver_person_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, order_ids: _Optional[_Iterable[int]]=..., taxi_id: _Optional[int]=..., type: _Optional[_Union[AllocationPlanType, str]]=..., pick_up_person_ids: _Optional[_Iterable[int]]=..., deliver_person_ids: _Optional[_Iterable[int]]=...) -> None:
        ...