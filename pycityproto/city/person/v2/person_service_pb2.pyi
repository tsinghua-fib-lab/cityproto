from city.geo.v2 import geo_pb2 as _geo_pb2
from city.person.v2 import motion_pb2 as _motion_pb2
from city.person.v2 import pedestrian_pb2 as _pedestrian_pb2
from city.person.v2 import person_pb2 as _person_pb2
from city.person.v2 import person_runtime_pb2 as _person_runtime_pb2
from city.person.v2 import taxi_pb2 as _taxi_pb2
from city.person.v2 import vehicle_pb2 as _vehicle_pb2
from city.trip.v2 import trip_pb2 as _trip_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetPersonRequest(_message.Message):
    __slots__ = ['person_id']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    person_id: int

    def __init__(self, person_id: _Optional[int]=...) -> None:
        ...

class GetPersonResponse(_message.Message):
    __slots__ = ['person']
    PERSON_FIELD_NUMBER: _ClassVar[int]
    person: _person_runtime_pb2.PersonRuntime

    def __init__(self, person: _Optional[_Union[_person_runtime_pb2.PersonRuntime, _Mapping]]=...) -> None:
        ...

class AddPersonRequest(_message.Message):
    __slots__ = ['person']
    PERSON_FIELD_NUMBER: _ClassVar[int]
    person: _person_pb2.Person

    def __init__(self, person: _Optional[_Union[_person_pb2.Person, _Mapping]]=...) -> None:
        ...

class AddPersonResponse(_message.Message):
    __slots__ = ['person_id']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    person_id: int

    def __init__(self, person_id: _Optional[int]=...) -> None:
        ...

class SetScheduleRequest(_message.Message):
    __slots__ = ['person_id', 'schedules']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    SCHEDULES_FIELD_NUMBER: _ClassVar[int]
    person_id: int
    schedules: _containers.RepeatedCompositeFieldContainer[_trip_pb2.Schedule]

    def __init__(self, person_id: _Optional[int]=..., schedules: _Optional[_Iterable[_Union[_trip_pb2.Schedule, _Mapping]]]=...) -> None:
        ...

class SetScheduleResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetPersonsRequest(_message.Message):
    __slots__ = ['person_ids', 'exclude_statuses', 'return_base']
    PERSON_IDS_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_STATUSES_FIELD_NUMBER: _ClassVar[int]
    RETURN_BASE_FIELD_NUMBER: _ClassVar[int]
    person_ids: _containers.RepeatedScalarFieldContainer[int]
    exclude_statuses: _containers.RepeatedScalarFieldContainer[_motion_pb2.Status]
    return_base: bool

    def __init__(self, person_ids: _Optional[_Iterable[int]]=..., exclude_statuses: _Optional[_Iterable[_Union[_motion_pb2.Status, str]]]=..., return_base: bool=...) -> None:
        ...

class GetPersonsResponse(_message.Message):
    __slots__ = ['persons']
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    persons: _containers.RepeatedCompositeFieldContainer[_person_runtime_pb2.PersonRuntime]

    def __init__(self, persons: _Optional[_Iterable[_Union[_person_runtime_pb2.PersonRuntime, _Mapping]]]=...) -> None:
        ...

class GetPersonByLongLatBBoxRequest(_message.Message):
    __slots__ = ['bbox', 'exclude_statuses', 'return_base']
    BBOX_FIELD_NUMBER: _ClassVar[int]
    EXCLUDE_STATUSES_FIELD_NUMBER: _ClassVar[int]
    RETURN_BASE_FIELD_NUMBER: _ClassVar[int]
    bbox: _geo_pb2.LongLatBBox
    exclude_statuses: _containers.RepeatedScalarFieldContainer[_motion_pb2.Status]
    return_base: bool

    def __init__(self, bbox: _Optional[_Union[_geo_pb2.LongLatBBox, _Mapping]]=..., exclude_statuses: _Optional[_Iterable[_Union[_motion_pb2.Status, str]]]=..., return_base: bool=...) -> None:
        ...

class GetPersonByLongLatBBoxResponse(_message.Message):
    __slots__ = ['persons']
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    persons: _containers.RepeatedCompositeFieldContainer[_person_runtime_pb2.PersonRuntime]

    def __init__(self, persons: _Optional[_Iterable[_Union[_person_runtime_pb2.PersonRuntime, _Mapping]]]=...) -> None:
        ...

class GetAllVehiclesRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetAllVehiclesResponse(_message.Message):
    __slots__ = ['vehicles']
    VEHICLES_FIELD_NUMBER: _ClassVar[int]
    vehicles: _containers.RepeatedCompositeFieldContainer[_vehicle_pb2.VehicleRuntime]

    def __init__(self, vehicles: _Optional[_Iterable[_Union[_vehicle_pb2.VehicleRuntime, _Mapping]]]=...) -> None:
        ...

class GetAllPedestriansRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetAllPedestriansResponse(_message.Message):
    __slots__ = ['pedestrians']
    PEDESTRIANS_FIELD_NUMBER: _ClassVar[int]
    pedestrians: _containers.RepeatedCompositeFieldContainer[_person_runtime_pb2.PersonRuntime]

    def __init__(self, pedestrians: _Optional[_Iterable[_Union[_person_runtime_pb2.PersonRuntime, _Mapping]]]=...) -> None:
        ...

class ResetPersonPositionRequest(_message.Message):
    __slots__ = ['person_id', 'position']
    PERSON_ID_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    person_id: int
    position: _geo_pb2.Position

    def __init__(self, person_id: _Optional[int]=..., position: _Optional[_Union[_geo_pb2.Position, _Mapping]]=...) -> None:
        ...

class ResetPersonPositionResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class SetControlledVehicleIDsRequest(_message.Message):
    __slots__ = ['vehicle_ids', 'route_vehicle_ids']
    VEHICLE_IDS_FIELD_NUMBER: _ClassVar[int]
    ROUTE_VEHICLE_IDS_FIELD_NUMBER: _ClassVar[int]
    vehicle_ids: _containers.RepeatedScalarFieldContainer[int]
    route_vehicle_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, vehicle_ids: _Optional[_Iterable[int]]=..., route_vehicle_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class SetControlledVehicleIDsResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class FetchControlledVehicleEnvsRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class FetchControlledVehicleEnvsResponse(_message.Message):
    __slots__ = ['vehicle_envs', 'route_vehicle_envs']
    VEHICLE_ENVS_FIELD_NUMBER: _ClassVar[int]
    ROUTE_VEHICLE_ENVS_FIELD_NUMBER: _ClassVar[int]
    vehicle_envs: _containers.RepeatedCompositeFieldContainer[_vehicle_pb2.VehicleEnv]
    route_vehicle_envs: _containers.RepeatedCompositeFieldContainer[_vehicle_pb2.VehicleEnv]

    def __init__(self, vehicle_envs: _Optional[_Iterable[_Union[_vehicle_pb2.VehicleEnv, _Mapping]]]=..., route_vehicle_envs: _Optional[_Iterable[_Union[_vehicle_pb2.VehicleEnv, _Mapping]]]=...) -> None:
        ...

class SetControlledVehicleActionsRequest(_message.Message):
    __slots__ = ['vehicle_actions', 'vehicle_journeys']
    VEHICLE_ACTIONS_FIELD_NUMBER: _ClassVar[int]
    VEHICLE_JOURNEYS_FIELD_NUMBER: _ClassVar[int]
    vehicle_actions: _containers.RepeatedCompositeFieldContainer[_vehicle_pb2.VehicleAction]
    vehicle_journeys: _containers.RepeatedCompositeFieldContainer[_vehicle_pb2.VehicleRouteAction]

    def __init__(self, vehicle_actions: _Optional[_Iterable[_Union[_vehicle_pb2.VehicleAction, _Mapping]]]=..., vehicle_journeys: _Optional[_Iterable[_Union[_vehicle_pb2.VehicleRouteAction, _Mapping]]]=...) -> None:
        ...

class SetControlledVehicleActionsResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class SetControlledTaxiIDsRequest(_message.Message):
    __slots__ = ['taxi_ids']
    TAXI_IDS_FIELD_NUMBER: _ClassVar[int]
    taxi_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, taxi_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class SetControlledTaxiIDsResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetAllOrdersRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetAllOrdersResponse(_message.Message):
    __slots__ = ['order_infos']
    ORDER_INFOS_FIELD_NUMBER: _ClassVar[int]
    order_infos: _containers.RepeatedCompositeFieldContainer[_taxi_pb2.RequestOrderInfo]

    def __init__(self, order_infos: _Optional[_Iterable[_Union[_taxi_pb2.RequestOrderInfo, _Mapping]]]=...) -> None:
        ...

class SetControlledTaxiToOrdersRequest(_message.Message):
    __slots__ = ['order_plans']
    ORDER_PLANS_FIELD_NUMBER: _ClassVar[int]
    order_plans: _containers.RepeatedCompositeFieldContainer[_taxi_pb2.OrderAllocationPlan]

    def __init__(self, order_plans: _Optional[_Iterable[_Union[_taxi_pb2.OrderAllocationPlan, _Mapping]]]=...) -> None:
        ...

class SetControlledTaxiToOrdersResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class SetControlledPedestriansRequest(_message.Message):
    __slots__ = ['pedestrian_ids']
    PEDESTRIAN_IDS_FIELD_NUMBER: _ClassVar[int]
    pedestrian_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, pedestrian_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class SetControlledPedestriansResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class FetchControlledPedestriansEnvsRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class FetchControlledPedestriansEnvsResponse(_message.Message):
    __slots__ = ['pedestrian_envs']
    PEDESTRIAN_ENVS_FIELD_NUMBER: _ClassVar[int]
    pedestrian_envs: _containers.RepeatedCompositeFieldContainer[_pedestrian_pb2.PedestrianEnv]

    def __init__(self, pedestrian_envs: _Optional[_Iterable[_Union[_pedestrian_pb2.PedestrianEnv, _Mapping]]]=...) -> None:
        ...

class SetControlledPedestriansActionsRequest(_message.Message):
    __slots__ = ['pedestrian_actions']
    PEDESTRIAN_ACTIONS_FIELD_NUMBER: _ClassVar[int]
    pedestrian_actions: _containers.RepeatedCompositeFieldContainer[_pedestrian_pb2.PedestrianAction]

    def __init__(self, pedestrian_actions: _Optional[_Iterable[_Union[_pedestrian_pb2.PedestrianAction, _Mapping]]]=...) -> None:
        ...

class SetControlledPedestriansActionsResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetControlledTaxiOrderAllocationPlanRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetControlledTaxiOrderAllocationPlanResponse(_message.Message):
    __slots__ = ['order_allocations']
    ORDER_ALLOCATIONS_FIELD_NUMBER: _ClassVar[int]
    order_allocations: _containers.RepeatedCompositeFieldContainer[_taxi_pb2.OrderAllocations]

    def __init__(self, order_allocations: _Optional[_Iterable[_Union[_taxi_pb2.OrderAllocations, _Mapping]]]=...) -> None:
        ...

class SetControlledTaxiOrderAllocationPlanRequest(_message.Message):
    __slots__ = ['order_allocations']
    ORDER_ALLOCATIONS_FIELD_NUMBER: _ClassVar[int]
    order_allocations: _containers.RepeatedCompositeFieldContainer[_taxi_pb2.OrderAllocations]

    def __init__(self, order_allocations: _Optional[_Iterable[_Union[_taxi_pb2.OrderAllocations, _Mapping]]]=...) -> None:
        ...

class SetControlledTaxiOrderAllocationPlanResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...