from city.geo.v2 import geo_pb2 as _geo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class OrderStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    ORDER_STATUS_UNSPECIFIED: _ClassVar[OrderStatus]
    ORDER_STATUS_WAITING: _ClassVar[OrderStatus]
    ORDER_STATUS_PICKING_UP: _ClassVar[OrderStatus]
    ORDER_STATUS_DELIVERING: _ClassVar[OrderStatus]
    ORDER_STATUS_COMPLETED: _ClassVar[OrderStatus]

class AllocationPlanType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    ALLOCATION_PLAN_TYPE_UNSPECIFIED: _ClassVar[AllocationPlanType]
    ALLOCATION_PLAN_TYPE_PICK_UP: _ClassVar[AllocationPlanType]
    ALLOCATION_PLAN_TYPE_DELIVER: _ClassVar[AllocationPlanType]
ORDER_STATUS_UNSPECIFIED: OrderStatus
ORDER_STATUS_WAITING: OrderStatus
ORDER_STATUS_PICKING_UP: OrderStatus
ORDER_STATUS_DELIVERING: OrderStatus
ORDER_STATUS_COMPLETED: OrderStatus
ALLOCATION_PLAN_TYPE_UNSPECIFIED: AllocationPlanType
ALLOCATION_PLAN_TYPE_PICK_UP: AllocationPlanType
ALLOCATION_PLAN_TYPE_DELIVER: AllocationPlanType

class RequestOrderInfo(_message.Message):
    __slots__ = ['person_id', 'request_time', 'order_id', 'departure', 'destination', 'status', 'departure_time']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    REQUEST_TIME_FIELD_NUMBER: _ClassVar[int]
    ORDER_ID_FIELD_NUMBER: _ClassVar[int]
    DEPARTURE_FIELD_NUMBER: _ClassVar[int]
    DESTINATION_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    DEPARTURE_TIME_FIELD_NUMBER: _ClassVar[int]
    person_id: int
    request_time: float
    order_id: int
    departure: _geo_pb2.Position
    destination: _geo_pb2.Position
    status: OrderStatus
    departure_time: float

    def __init__(self, person_id: _Optional[int]=..., request_time: _Optional[float]=..., order_id: _Optional[int]=..., departure: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., destination: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., status: _Optional[_Union[OrderStatus, str]]=..., departure_time: _Optional[float]=...) -> None:
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

class OrderAllocations(_message.Message):
    __slots__ = ['order_allocations', 'taxi_id']
    ORDER_ALLOCATIONS_FIELD_NUMBER: _ClassVar[int]
    TAXI_ID_FIELD_NUMBER: _ClassVar[int]
    order_allocations: _containers.RepeatedCompositeFieldContainer[OrderAllocationPlan]
    taxi_id: int

    def __init__(self, order_allocations: _Optional[_Iterable[_Union[OrderAllocationPlan, _Mapping]]]=..., taxi_id: _Optional[int]=...) -> None:
        ...