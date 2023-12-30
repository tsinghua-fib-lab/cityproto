from city.wargame.v1 import wargame_pb2 as _wargame_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class PickPointsRequest(_message.Message):
    __slots__ = ['camp', 'points']
    CAMP_FIELD_NUMBER: _ClassVar[int]
    POINTS_FIELD_NUMBER: _ClassVar[int]
    camp: _wargame_pb2.Camp
    points: _containers.RepeatedCompositeFieldContainer[_wargame_pb2.Point]

    def __init__(self, camp: _Optional[_Union[_wargame_pb2.Camp, str]]=..., points: _Optional[_Iterable[_Union[_wargame_pb2.Point, _Mapping]]]=...) -> None:
        ...

class PickPointsResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetPickPointsRequest(_message.Message):
    __slots__ = ['camp']
    CAMP_FIELD_NUMBER: _ClassVar[int]
    camp: _wargame_pb2.Camp

    def __init__(self, camp: _Optional[_Union[_wargame_pb2.Camp, str]]=...) -> None:
        ...

class GetPickPointsResponse(_message.Message):
    __slots__ = ['points']
    POINTS_FIELD_NUMBER: _ClassVar[int]
    points: _containers.RepeatedCompositeFieldContainer[_wargame_pb2.Point]

    def __init__(self, points: _Optional[_Iterable[_Union[_wargame_pb2.Point, _Mapping]]]=...) -> None:
        ...

class SetScoreWeightRequest(_message.Message):
    __slots__ = ['money', 'population_loss', 'elec_power', 'elec_distory_1', 'water_distory_1', 'bs_distory_1', 'traffic_distory_1', 'elec_distory_2', 'water_distory_2', 'bs_distory_2', 'traffic_distory_2', 'elec_distory_3', 'water_distory_3', 'bs_distory_3', 'traffic_distory_3', 'defense_success']
    MONEY_FIELD_NUMBER: _ClassVar[int]
    POPULATION_LOSS_FIELD_NUMBER: _ClassVar[int]
    ELEC_POWER_FIELD_NUMBER: _ClassVar[int]
    ELEC_DISTORY_1_FIELD_NUMBER: _ClassVar[int]
    WATER_DISTORY_1_FIELD_NUMBER: _ClassVar[int]
    BS_DISTORY_1_FIELD_NUMBER: _ClassVar[int]
    TRAFFIC_DISTORY_1_FIELD_NUMBER: _ClassVar[int]
    ELEC_DISTORY_2_FIELD_NUMBER: _ClassVar[int]
    WATER_DISTORY_2_FIELD_NUMBER: _ClassVar[int]
    BS_DISTORY_2_FIELD_NUMBER: _ClassVar[int]
    TRAFFIC_DISTORY_2_FIELD_NUMBER: _ClassVar[int]
    ELEC_DISTORY_3_FIELD_NUMBER: _ClassVar[int]
    WATER_DISTORY_3_FIELD_NUMBER: _ClassVar[int]
    BS_DISTORY_3_FIELD_NUMBER: _ClassVar[int]
    TRAFFIC_DISTORY_3_FIELD_NUMBER: _ClassVar[int]
    DEFENSE_SUCCESS_FIELD_NUMBER: _ClassVar[int]
    money: float
    population_loss: float
    elec_power: float
    elec_distory_1: float
    water_distory_1: float
    bs_distory_1: float
    traffic_distory_1: float
    elec_distory_2: float
    water_distory_2: float
    bs_distory_2: float
    traffic_distory_2: float
    elec_distory_3: float
    water_distory_3: float
    bs_distory_3: float
    traffic_distory_3: float
    defense_success: float

    def __init__(self, money: _Optional[float]=..., population_loss: _Optional[float]=..., elec_power: _Optional[float]=..., elec_distory_1: _Optional[float]=..., water_distory_1: _Optional[float]=..., bs_distory_1: _Optional[float]=..., traffic_distory_1: _Optional[float]=..., elec_distory_2: _Optional[float]=..., water_distory_2: _Optional[float]=..., bs_distory_2: _Optional[float]=..., traffic_distory_2: _Optional[float]=..., elec_distory_3: _Optional[float]=..., water_distory_3: _Optional[float]=..., bs_distory_3: _Optional[float]=..., traffic_distory_3: _Optional[float]=..., defense_success: _Optional[float]=...) -> None:
        ...

class SetScoreWeightResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GiveDefenseOrderRequest(_message.Message):
    __slots__ = ['weight_radius', 'weight_550', 'weight_220', 'weight_110', 'prob_threshold']
    WEIGHT_RADIUS_FIELD_NUMBER: _ClassVar[int]
    WEIGHT_550_FIELD_NUMBER: _ClassVar[int]
    WEIGHT_220_FIELD_NUMBER: _ClassVar[int]
    WEIGHT_110_FIELD_NUMBER: _ClassVar[int]
    PROB_THRESHOLD_FIELD_NUMBER: _ClassVar[int]
    weight_radius: float
    weight_550: float
    weight_220: float
    weight_110: float
    prob_threshold: float

    def __init__(self, weight_radius: _Optional[float]=..., weight_550: _Optional[float]=..., weight_220: _Optional[float]=..., weight_110: _Optional[float]=..., prob_threshold: _Optional[float]=...) -> None:
        ...

class GiveDefenseOrderResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetHitHistoryRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetHitHistoryResponse(_message.Message):
    __slots__ = ['histories']
    HISTORIES_FIELD_NUMBER: _ClassVar[int]
    histories: _containers.RepeatedCompositeFieldContainer[_wargame_pb2.HitHistory]

    def __init__(self, histories: _Optional[_Iterable[_Union[_wargame_pb2.HitHistory, _Mapping]]]=...) -> None:
        ...

class GetRecoPointsRequest(_message.Message):
    __slots__ = ['camp', 'type']
    CAMP_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    camp: _wargame_pb2.Camp
    type: _wargame_pb2.RecoAlgoType

    def __init__(self, camp: _Optional[_Union[_wargame_pb2.Camp, str]]=..., type: _Optional[_Union[_wargame_pb2.RecoAlgoType, str]]=...) -> None:
        ...

class GetRecoPointsResponse(_message.Message):
    __slots__ = ['points']
    POINTS_FIELD_NUMBER: _ClassVar[int]
    points: _containers.RepeatedCompositeFieldContainer[_wargame_pb2.RecoPoint]

    def __init__(self, points: _Optional[_Iterable[_Union[_wargame_pb2.RecoPoint, _Mapping]]]=...) -> None:
        ...

class GetStepRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetStepResponse(_message.Message):
    __slots__ = ['step', 'red_pick_ready', 'blue_pick_ready', 'round', 'is_game_started']
    STEP_FIELD_NUMBER: _ClassVar[int]
    RED_PICK_READY_FIELD_NUMBER: _ClassVar[int]
    BLUE_PICK_READY_FIELD_NUMBER: _ClassVar[int]
    ROUND_FIELD_NUMBER: _ClassVar[int]
    IS_GAME_STARTED_FIELD_NUMBER: _ClassVar[int]
    step: int
    red_pick_ready: bool
    blue_pick_ready: bool
    round: int
    is_game_started: bool

    def __init__(self, step: _Optional[int]=..., red_pick_ready: bool=..., blue_pick_ready: bool=..., round: _Optional[int]=..., is_game_started: bool=...) -> None:
        ...

class GetCasualtiesRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetCasualtiesResponse(_message.Message):
    __slots__ = ['casualties']
    CASUALTIES_FIELD_NUMBER: _ClassVar[int]
    casualties: _containers.RepeatedCompositeFieldContainer[_wargame_pb2.Casualty]

    def __init__(self, casualties: _Optional[_Iterable[_Union[_wargame_pb2.Casualty, _Mapping]]]=...) -> None:
        ...