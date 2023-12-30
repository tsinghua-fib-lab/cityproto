from city.geo.v2 import geo_pb2 as _geo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class RedWeaponType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    RED_WEAPON_TYPE_UNSPECIFIED: _ClassVar[RedWeaponType]
    RED_WEAPON_TYPE_AGM_158: _ClassVar[RedWeaponType]
    RED_WEAPON_TYPE_AGM_183: _ClassVar[RedWeaponType]
    RED_WEAPON_TYPE_AGM_88: _ClassVar[RedWeaponType]
    RED_WEAPON_TYPE_UGM_109: _ClassVar[RedWeaponType]
    RED_WEAPON_TYPE_AGM_142: _ClassVar[RedWeaponType]

class BlueWeaponType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    BLUE_WEAPON_TYPE_UNSPECIFIED: _ClassVar[BlueWeaponType]
    BLUE_WEAPON_TYPE_HQ_9: _ClassVar[BlueWeaponType]
    BLUE_WEAPON_TYPE_HQ_12: _ClassVar[BlueWeaponType]
    BLUE_WEAPON_TYPE_HQ_16: _ClassVar[BlueWeaponType]
    BLUE_WEAPON_TYPE_HQ_22: _ClassVar[BlueWeaponType]

class Camp(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    CAMP_UNSPECIFIED: _ClassVar[Camp]
    CAMP_RED: _ClassVar[Camp]
    CAMP_BLUE: _ClassVar[Camp]

class RecoType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    RECO_TYPE_UNSPECIFIED: _ClassVar[RecoType]
    RECO_TYPE_MISSILE: _ClassVar[RecoType]
    RECO_TYPE_ELEC: _ClassVar[RecoType]

class RecoAlgoType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    RECO_ALGO_TYPE_UNSPECIFIED: _ClassVar[RecoAlgoType]
    RECO_ALGO_TYPE_MAX_DESTORY: _ClassVar[RecoAlgoType]
    RECO_ALGO_TYPE_RANDOM: _ClassVar[RecoAlgoType]
    RECO_ALGO_TYPE_MIN_DESTORY: _ClassVar[RecoAlgoType]
RED_WEAPON_TYPE_UNSPECIFIED: RedWeaponType
RED_WEAPON_TYPE_AGM_158: RedWeaponType
RED_WEAPON_TYPE_AGM_183: RedWeaponType
RED_WEAPON_TYPE_AGM_88: RedWeaponType
RED_WEAPON_TYPE_UGM_109: RedWeaponType
RED_WEAPON_TYPE_AGM_142: RedWeaponType
BLUE_WEAPON_TYPE_UNSPECIFIED: BlueWeaponType
BLUE_WEAPON_TYPE_HQ_9: BlueWeaponType
BLUE_WEAPON_TYPE_HQ_12: BlueWeaponType
BLUE_WEAPON_TYPE_HQ_16: BlueWeaponType
BLUE_WEAPON_TYPE_HQ_22: BlueWeaponType
CAMP_UNSPECIFIED: Camp
CAMP_RED: Camp
CAMP_BLUE: Camp
RECO_TYPE_UNSPECIFIED: RecoType
RECO_TYPE_MISSILE: RecoType
RECO_TYPE_ELEC: RecoType
RECO_ALGO_TYPE_UNSPECIFIED: RecoAlgoType
RECO_ALGO_TYPE_MAX_DESTORY: RecoAlgoType
RECO_ALGO_TYPE_RANDOM: RecoAlgoType
RECO_ALGO_TYPE_MIN_DESTORY: RecoAlgoType

class RecoPoint(_message.Message):
    __slots__ = ['pos', 'type']
    POS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    pos: _geo_pb2.LongLatPosition
    type: RecoType

    def __init__(self, pos: _Optional[_Union[_geo_pb2.LongLatPosition, _Mapping]]=..., type: _Optional[_Union[RecoType, str]]=...) -> None:
        ...

class Weapon(_message.Message):
    __slots__ = ['red_type', 'blue_type', 'num']
    RED_TYPE_FIELD_NUMBER: _ClassVar[int]
    BLUE_TYPE_FIELD_NUMBER: _ClassVar[int]
    NUM_FIELD_NUMBER: _ClassVar[int]
    red_type: RedWeaponType
    blue_type: BlueWeaponType
    num: int

    def __init__(self, red_type: _Optional[_Union[RedWeaponType, str]]=..., blue_type: _Optional[_Union[BlueWeaponType, str]]=..., num: _Optional[int]=...) -> None:
        ...

class Point(_message.Message):
    __slots__ = ['weapon', 'pos', 'radius']
    WEAPON_FIELD_NUMBER: _ClassVar[int]
    POS_FIELD_NUMBER: _ClassVar[int]
    RADIUS_FIELD_NUMBER: _ClassVar[int]
    weapon: Weapon
    pos: _geo_pb2.LongLatPosition
    radius: float

    def __init__(self, weapon: _Optional[_Union[Weapon, _Mapping]]=..., pos: _Optional[_Union[_geo_pb2.LongLatPosition, _Mapping]]=..., radius: _Optional[float]=...) -> None:
        ...

class Hit(_message.Message):
    __slots__ = ['attack', 'defenses', 'result', 'arrive']
    ATTACK_FIELD_NUMBER: _ClassVar[int]
    DEFENSES_FIELD_NUMBER: _ClassVar[int]
    RESULT_FIELD_NUMBER: _ClassVar[int]
    ARRIVE_FIELD_NUMBER: _ClassVar[int]
    attack: Point
    defenses: _containers.RepeatedCompositeFieldContainer[Point]
    result: bool
    arrive: int

    def __init__(self, attack: _Optional[_Union[Point, _Mapping]]=..., defenses: _Optional[_Iterable[_Union[Point, _Mapping]]]=..., result: bool=..., arrive: _Optional[int]=...) -> None:
        ...

class HitHistory(_message.Message):
    __slots__ = ['hits']
    HITS_FIELD_NUMBER: _ClassVar[int]
    hits: _containers.RepeatedCompositeFieldContainer[Hit]

    def __init__(self, hits: _Optional[_Iterable[_Union[Hit, _Mapping]]]=...) -> None:
        ...

class Casualty(_message.Message):
    __slots__ = ['aoi_id', 'num']
    AOI_ID_FIELD_NUMBER: _ClassVar[int]
    NUM_FIELD_NUMBER: _ClassVar[int]
    aoi_id: int
    num: int

    def __init__(self, aoi_id: _Optional[int]=..., num: _Optional[int]=...) -> None:
        ...