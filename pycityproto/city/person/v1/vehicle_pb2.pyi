from city.person.v1 import motion_pb2 as _motion_pb2
from city.routing.v2 import routing_pb2 as _routing_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class VehicleRelation(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    VEHICLE_RELATION_UNSPECIFIED: _ClassVar[VehicleRelation]
    VEHICLE_RELATION_AHEAD: _ClassVar[VehicleRelation]
    VEHICLE_RELATION_BEHIND: _ClassVar[VehicleRelation]
    VEHICLE_RELATION_SHADOW_AHEAD: _ClassVar[VehicleRelation]
    VEHICLE_RELATION_SHADOW_BEHIND: _ClassVar[VehicleRelation]
    VEHICLE_RELATION_LEFT_AHEAD: _ClassVar[VehicleRelation]
    VEHICLE_RELATION_RIGHT_AHEAD: _ClassVar[VehicleRelation]
    VEHICLE_RELATION_LEFT_BEHIND: _ClassVar[VehicleRelation]
    VEHICLE_RELATION_RIGHT_BEHIND: _ClassVar[VehicleRelation]

class LightState(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    LIGHT_STATE_UNSPECIFIED: _ClassVar[LightState]
    LIGHT_STATE_RED: _ClassVar[LightState]
    LIGHT_STATE_GREEN: _ClassVar[LightState]
    LIGHT_STATE_YELLOW: _ClassVar[LightState]
VEHICLE_RELATION_UNSPECIFIED: VehicleRelation
VEHICLE_RELATION_AHEAD: VehicleRelation
VEHICLE_RELATION_BEHIND: VehicleRelation
VEHICLE_RELATION_SHADOW_AHEAD: VehicleRelation
VEHICLE_RELATION_SHADOW_BEHIND: VehicleRelation
VEHICLE_RELATION_LEFT_AHEAD: VehicleRelation
VEHICLE_RELATION_RIGHT_AHEAD: VehicleRelation
VEHICLE_RELATION_LEFT_BEHIND: VehicleRelation
VEHICLE_RELATION_RIGHT_BEHIND: VehicleRelation
LIGHT_STATE_UNSPECIFIED: LightState
LIGHT_STATE_RED: LightState
LIGHT_STATE_GREEN: LightState
LIGHT_STATE_YELLOW: LightState

class LC(_message.Message):
    __slots__ = ['shadow_lane_id', 'shadow_s', 'angle', 'completed_ratio']
    SHADOW_LANE_ID_FIELD_NUMBER: _ClassVar[int]
    SHADOW_S_FIELD_NUMBER: _ClassVar[int]
    ANGLE_FIELD_NUMBER: _ClassVar[int]
    COMPLETED_RATIO_FIELD_NUMBER: _ClassVar[int]
    shadow_lane_id: int
    shadow_s: float
    angle: float
    completed_ratio: float

    def __init__(self, shadow_lane_id: _Optional[int]=..., shadow_s: _Optional[float]=..., angle: _Optional[float]=..., completed_ratio: _Optional[float]=...) -> None:
        ...

class VehicleAction(_message.Message):
    __slots__ = ['id', 'acc', 'lc_target_id', 'angle']
    ID_FIELD_NUMBER: _ClassVar[int]
    ACC_FIELD_NUMBER: _ClassVar[int]
    LC_TARGET_ID_FIELD_NUMBER: _ClassVar[int]
    ANGLE_FIELD_NUMBER: _ClassVar[int]
    id: int
    acc: float
    lc_target_id: int
    angle: float

    def __init__(self, id: _Optional[int]=..., acc: _Optional[float]=..., lc_target_id: _Optional[int]=..., angle: _Optional[float]=...) -> None:
        ...

class VehicleRuntime(_message.Message):
    __slots__ = ['base', 'lc', 'action', 'running_distance', 'num_going_astray', 'departure_time', 'eta', 'eta_free_flow']
    BASE_FIELD_NUMBER: _ClassVar[int]
    LC_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    RUNNING_DISTANCE_FIELD_NUMBER: _ClassVar[int]
    NUM_GOING_ASTRAY_FIELD_NUMBER: _ClassVar[int]
    DEPARTURE_TIME_FIELD_NUMBER: _ClassVar[int]
    ETA_FIELD_NUMBER: _ClassVar[int]
    ETA_FREE_FLOW_FIELD_NUMBER: _ClassVar[int]
    base: _motion_pb2.PersonMotion
    lc: LC
    action: VehicleAction
    running_distance: float
    num_going_astray: int
    departure_time: float
    eta: float
    eta_free_flow: float

    def __init__(self, base: _Optional[_Union[_motion_pb2.PersonMotion, _Mapping]]=..., lc: _Optional[_Union[LC, _Mapping]]=..., action: _Optional[_Union[VehicleAction, _Mapping]]=..., running_distance: _Optional[float]=..., num_going_astray: _Optional[int]=..., departure_time: _Optional[float]=..., eta: _Optional[float]=..., eta_free_flow: _Optional[float]=...) -> None:
        ...

class ObservedVehicle(_message.Message):
    __slots__ = ['id', 'motion', 'relative_distance', 'relation']
    ID_FIELD_NUMBER: _ClassVar[int]
    MOTION_FIELD_NUMBER: _ClassVar[int]
    RELATIVE_DISTANCE_FIELD_NUMBER: _ClassVar[int]
    RELATION_FIELD_NUMBER: _ClassVar[int]
    id: int
    motion: _motion_pb2.PersonMotion
    relative_distance: float
    relation: VehicleRelation

    def __init__(self, id: _Optional[int]=..., motion: _Optional[_Union[_motion_pb2.PersonMotion, _Mapping]]=..., relative_distance: _Optional[float]=..., relation: _Optional[_Union[VehicleRelation, str]]=...) -> None:
        ...

class ObservedLane(_message.Message):
    __slots__ = ['id', 'restriction', 'light_state', 'light_remaining_time']
    ID_FIELD_NUMBER: _ClassVar[int]
    RESTRICTION_FIELD_NUMBER: _ClassVar[int]
    LIGHT_STATE_FIELD_NUMBER: _ClassVar[int]
    LIGHT_REMAINING_TIME_FIELD_NUMBER: _ClassVar[int]
    id: int
    restriction: bool
    light_state: LightState
    light_remaining_time: float

    def __init__(self, id: _Optional[int]=..., restriction: bool=..., light_state: _Optional[_Union[LightState, str]]=..., light_remaining_time: _Optional[float]=...) -> None:
        ...

class VehicleEnv(_message.Message):
    __slots__ = ['id', 'runtime', 'journey', 'observed_vehicles', 'observed_lanes']
    ID_FIELD_NUMBER: _ClassVar[int]
    RUNTIME_FIELD_NUMBER: _ClassVar[int]
    JOURNEY_FIELD_NUMBER: _ClassVar[int]
    OBSERVED_VEHICLES_FIELD_NUMBER: _ClassVar[int]
    OBSERVED_LANES_FIELD_NUMBER: _ClassVar[int]
    id: int
    runtime: VehicleRuntime
    journey: _routing_pb2.Journey
    observed_vehicles: _containers.RepeatedCompositeFieldContainer[ObservedVehicle]
    observed_lanes: _containers.RepeatedCompositeFieldContainer[ObservedLane]

    def __init__(self, id: _Optional[int]=..., runtime: _Optional[_Union[VehicleRuntime, _Mapping]]=..., journey: _Optional[_Union[_routing_pb2.Journey, _Mapping]]=..., observed_vehicles: _Optional[_Iterable[_Union[ObservedVehicle, _Mapping]]]=..., observed_lanes: _Optional[_Iterable[_Union[ObservedLane, _Mapping]]]=...) -> None:
        ...