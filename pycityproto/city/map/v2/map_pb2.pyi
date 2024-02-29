from city.geo.v2 import geo_pb2 as _geo_pb2
from city.map.v2 import light_pb2 as _light_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class LaneType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    LANE_TYPE_UNSPECIFIED: _ClassVar[LaneType]
    LANE_TYPE_DRIVING: _ClassVar[LaneType]
    LANE_TYPE_WALKING: _ClassVar[LaneType]

class LaneTurn(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    LANE_TURN_UNSPECIFIED: _ClassVar[LaneTurn]
    LANE_TURN_STRAIGHT: _ClassVar[LaneTurn]
    LANE_TURN_LEFT: _ClassVar[LaneTurn]
    LANE_TURN_RIGHT: _ClassVar[LaneTurn]
    LANE_TURN_AROUND: _ClassVar[LaneTurn]

class LaneConnectionType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    LANE_CONNECTION_TYPE_UNSPECIFIED: _ClassVar[LaneConnectionType]
    LANE_CONNECTION_TYPE_HEAD: _ClassVar[LaneConnectionType]
    LANE_CONNECTION_TYPE_TAIL: _ClassVar[LaneConnectionType]

class AoiType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    AOI_TYPE_UNSPECIFIED: _ClassVar[AoiType]
    AOI_TYPE_BUS_STATION: _ClassVar[AoiType]
    AOI_TYPE_OTHER: _ClassVar[AoiType]

class LandUseType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    LAND_USE_TYPE_UNSPECIFIED: _ClassVar[LandUseType]
    LAND_USE_TYPE_COMMERCIAL: _ClassVar[LandUseType]
    LAND_USE_TYPE_INDUSTRIAL: _ClassVar[LandUseType]
    LAND_USE_TYPE_RESIDENTIAL: _ClassVar[LandUseType]
    LAND_USE_TYPE_PUBLIC: _ClassVar[LandUseType]
    LAND_USE_TYPE_TRANSPORTATION: _ClassVar[LandUseType]
    LAND_USE_TYPE_OTHER: _ClassVar[LandUseType]
LANE_TYPE_UNSPECIFIED: LaneType
LANE_TYPE_DRIVING: LaneType
LANE_TYPE_WALKING: LaneType
LANE_TURN_UNSPECIFIED: LaneTurn
LANE_TURN_STRAIGHT: LaneTurn
LANE_TURN_LEFT: LaneTurn
LANE_TURN_RIGHT: LaneTurn
LANE_TURN_AROUND: LaneTurn
LANE_CONNECTION_TYPE_UNSPECIFIED: LaneConnectionType
LANE_CONNECTION_TYPE_HEAD: LaneConnectionType
LANE_CONNECTION_TYPE_TAIL: LaneConnectionType
AOI_TYPE_UNSPECIFIED: AoiType
AOI_TYPE_BUS_STATION: AoiType
AOI_TYPE_OTHER: AoiType
LAND_USE_TYPE_UNSPECIFIED: LandUseType
LAND_USE_TYPE_COMMERCIAL: LandUseType
LAND_USE_TYPE_INDUSTRIAL: LandUseType
LAND_USE_TYPE_RESIDENTIAL: LandUseType
LAND_USE_TYPE_PUBLIC: LandUseType
LAND_USE_TYPE_TRANSPORTATION: LandUseType
LAND_USE_TYPE_OTHER: LandUseType

class Polyline(_message.Message):
    __slots__ = ['nodes']
    NODES_FIELD_NUMBER: _ClassVar[int]
    nodes: _containers.RepeatedCompositeFieldContainer[_geo_pb2.XYPosition]

    def __init__(self, nodes: _Optional[_Iterable[_Union[_geo_pb2.XYPosition, _Mapping]]]=...) -> None:
        ...

class Header(_message.Message):
    __slots__ = ['name', 'date', 'north', 'south', 'east', 'west', 'projection']
    NAME_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    NORTH_FIELD_NUMBER: _ClassVar[int]
    SOUTH_FIELD_NUMBER: _ClassVar[int]
    EAST_FIELD_NUMBER: _ClassVar[int]
    WEST_FIELD_NUMBER: _ClassVar[int]
    PROJECTION_FIELD_NUMBER: _ClassVar[int]
    name: str
    date: str
    north: float
    south: float
    east: float
    west: float
    projection: str

    def __init__(self, name: _Optional[str]=..., date: _Optional[str]=..., north: _Optional[float]=..., south: _Optional[float]=..., east: _Optional[float]=..., west: _Optional[float]=..., projection: _Optional[str]=...) -> None:
        ...

class LaneOverlap(_message.Message):
    __slots__ = ['self', 'other', 'self_first']
    SELF_FIELD_NUMBER: _ClassVar[int]
    OTHER_FIELD_NUMBER: _ClassVar[int]
    SELF_FIRST_FIELD_NUMBER: _ClassVar[int]
    self: _geo_pb2.LanePosition
    other: _geo_pb2.LanePosition
    self_first: bool

    def __init__(self, self_: _Optional[_Union[_geo_pb2.LanePosition, _Mapping]]=..., other: _Optional[_Union[_geo_pb2.LanePosition, _Mapping]]=..., self_first: bool=...) -> None:
        ...

class LaneConnection(_message.Message):
    __slots__ = ['id', 'type']
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    id: int
    type: LaneConnectionType

    def __init__(self, id: _Optional[int]=..., type: _Optional[_Union[LaneConnectionType, str]]=...) -> None:
        ...

class Lane(_message.Message):
    __slots__ = ['id', 'type', 'turn', 'max_speed', 'length', 'width', 'center_line', 'left_border_line', 'right_border_line', 'predecessors', 'successors', 'left_lane_ids', 'right_lane_ids', 'parent_id', 'overlaps', 'aoi_ids']
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    TURN_FIELD_NUMBER: _ClassVar[int]
    MAX_SPEED_FIELD_NUMBER: _ClassVar[int]
    LENGTH_FIELD_NUMBER: _ClassVar[int]
    WIDTH_FIELD_NUMBER: _ClassVar[int]
    CENTER_LINE_FIELD_NUMBER: _ClassVar[int]
    LEFT_BORDER_LINE_FIELD_NUMBER: _ClassVar[int]
    RIGHT_BORDER_LINE_FIELD_NUMBER: _ClassVar[int]
    PREDECESSORS_FIELD_NUMBER: _ClassVar[int]
    SUCCESSORS_FIELD_NUMBER: _ClassVar[int]
    LEFT_LANE_IDS_FIELD_NUMBER: _ClassVar[int]
    RIGHT_LANE_IDS_FIELD_NUMBER: _ClassVar[int]
    PARENT_ID_FIELD_NUMBER: _ClassVar[int]
    OVERLAPS_FIELD_NUMBER: _ClassVar[int]
    AOI_IDS_FIELD_NUMBER: _ClassVar[int]
    id: int
    type: LaneType
    turn: LaneTurn
    max_speed: float
    length: float
    width: float
    center_line: Polyline
    left_border_line: Polyline
    right_border_line: Polyline
    predecessors: _containers.RepeatedCompositeFieldContainer[LaneConnection]
    successors: _containers.RepeatedCompositeFieldContainer[LaneConnection]
    left_lane_ids: _containers.RepeatedScalarFieldContainer[int]
    right_lane_ids: _containers.RepeatedScalarFieldContainer[int]
    parent_id: int
    overlaps: _containers.RepeatedCompositeFieldContainer[LaneOverlap]
    aoi_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, id: _Optional[int]=..., type: _Optional[_Union[LaneType, str]]=..., turn: _Optional[_Union[LaneTurn, str]]=..., max_speed: _Optional[float]=..., length: _Optional[float]=..., width: _Optional[float]=..., center_line: _Optional[_Union[Polyline, _Mapping]]=..., left_border_line: _Optional[_Union[Polyline, _Mapping]]=..., right_border_line: _Optional[_Union[Polyline, _Mapping]]=..., predecessors: _Optional[_Iterable[_Union[LaneConnection, _Mapping]]]=..., successors: _Optional[_Iterable[_Union[LaneConnection, _Mapping]]]=..., left_lane_ids: _Optional[_Iterable[int]]=..., right_lane_ids: _Optional[_Iterable[int]]=..., parent_id: _Optional[int]=..., overlaps: _Optional[_Iterable[_Union[LaneOverlap, _Mapping]]]=..., aoi_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class NextRoadLane(_message.Message):
    __slots__ = ['road_id', 'lane_id_a', 'lane_id_b']
    ROAD_ID_FIELD_NUMBER: _ClassVar[int]
    LANE_ID_A_FIELD_NUMBER: _ClassVar[int]
    LANE_ID_B_FIELD_NUMBER: _ClassVar[int]
    road_id: int
    lane_id_a: int
    lane_id_b: int

    def __init__(self, road_id: _Optional[int]=..., lane_id_a: _Optional[int]=..., lane_id_b: _Optional[int]=...) -> None:
        ...

class NextRoadLanePlan(_message.Message):
    __slots__ = ['next_road_lanes']
    NEXT_ROAD_LANES_FIELD_NUMBER: _ClassVar[int]
    next_road_lanes: _containers.RepeatedCompositeFieldContainer[NextRoadLane]

    def __init__(self, next_road_lanes: _Optional[_Iterable[_Union[NextRoadLane, _Mapping]]]=...) -> None:
        ...

class Road(_message.Message):
    __slots__ = ['id', 'name', 'lane_ids', 'next_road_lane_plans']
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    LANE_IDS_FIELD_NUMBER: _ClassVar[int]
    NEXT_ROAD_LANE_PLANS_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    lane_ids: _containers.RepeatedScalarFieldContainer[int]
    next_road_lane_plans: _containers.RepeatedCompositeFieldContainer[NextRoadLanePlan]

    def __init__(self, id: _Optional[int]=..., name: _Optional[str]=..., lane_ids: _Optional[_Iterable[int]]=..., next_road_lane_plans: _Optional[_Iterable[_Union[NextRoadLanePlan, _Mapping]]]=...) -> None:
        ...

class JunctionLaneGroup(_message.Message):
    __slots__ = ['in_road_id', 'in_angle', 'out_road_id', 'out_angle', 'lane_ids', 'turn']
    IN_ROAD_ID_FIELD_NUMBER: _ClassVar[int]
    IN_ANGLE_FIELD_NUMBER: _ClassVar[int]
    OUT_ROAD_ID_FIELD_NUMBER: _ClassVar[int]
    OUT_ANGLE_FIELD_NUMBER: _ClassVar[int]
    LANE_IDS_FIELD_NUMBER: _ClassVar[int]
    TURN_FIELD_NUMBER: _ClassVar[int]
    in_road_id: int
    in_angle: float
    out_road_id: int
    out_angle: float
    lane_ids: _containers.RepeatedScalarFieldContainer[int]
    turn: LaneTurn

    def __init__(self, in_road_id: _Optional[int]=..., in_angle: _Optional[float]=..., out_road_id: _Optional[int]=..., out_angle: _Optional[float]=..., lane_ids: _Optional[_Iterable[int]]=..., turn: _Optional[_Union[LaneTurn, str]]=...) -> None:
        ...

class Junction(_message.Message):
    __slots__ = ['id', 'lane_ids', 'driving_lane_groups', 'phases', 'fixed_program']
    ID_FIELD_NUMBER: _ClassVar[int]
    LANE_IDS_FIELD_NUMBER: _ClassVar[int]
    DRIVING_LANE_GROUPS_FIELD_NUMBER: _ClassVar[int]
    PHASES_FIELD_NUMBER: _ClassVar[int]
    FIXED_PROGRAM_FIELD_NUMBER: _ClassVar[int]
    id: int
    lane_ids: _containers.RepeatedScalarFieldContainer[int]
    driving_lane_groups: _containers.RepeatedCompositeFieldContainer[JunctionLaneGroup]
    phases: _containers.RepeatedCompositeFieldContainer[_light_pb2.AvailablePhase]
    fixed_program: _light_pb2.TrafficLight

    def __init__(self, id: _Optional[int]=..., lane_ids: _Optional[_Iterable[int]]=..., driving_lane_groups: _Optional[_Iterable[_Union[JunctionLaneGroup, _Mapping]]]=..., phases: _Optional[_Iterable[_Union[_light_pb2.AvailablePhase, _Mapping]]]=..., fixed_program: _Optional[_Union[_light_pb2.TrafficLight, _Mapping]]=...) -> None:
        ...

class Aoi(_message.Message):
    __slots__ = ['id', 'name', 'type', 'driving_positions', 'walking_positions', 'positions', 'driving_gates', 'walking_gates', 'area', 'land_use', 'urban_land_use', 'poi_ids']
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DRIVING_POSITIONS_FIELD_NUMBER: _ClassVar[int]
    WALKING_POSITIONS_FIELD_NUMBER: _ClassVar[int]
    POSITIONS_FIELD_NUMBER: _ClassVar[int]
    DRIVING_GATES_FIELD_NUMBER: _ClassVar[int]
    WALKING_GATES_FIELD_NUMBER: _ClassVar[int]
    AREA_FIELD_NUMBER: _ClassVar[int]
    LAND_USE_FIELD_NUMBER: _ClassVar[int]
    URBAN_LAND_USE_FIELD_NUMBER: _ClassVar[int]
    POI_IDS_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    type: AoiType
    driving_positions: _containers.RepeatedCompositeFieldContainer[_geo_pb2.LanePosition]
    walking_positions: _containers.RepeatedCompositeFieldContainer[_geo_pb2.LanePosition]
    positions: _containers.RepeatedCompositeFieldContainer[_geo_pb2.XYPosition]
    driving_gates: _containers.RepeatedCompositeFieldContainer[_geo_pb2.XYPosition]
    walking_gates: _containers.RepeatedCompositeFieldContainer[_geo_pb2.XYPosition]
    area: float
    land_use: LandUseType
    urban_land_use: str
    poi_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, id: _Optional[int]=..., name: _Optional[str]=..., type: _Optional[_Union[AoiType, str]]=..., driving_positions: _Optional[_Iterable[_Union[_geo_pb2.LanePosition, _Mapping]]]=..., walking_positions: _Optional[_Iterable[_Union[_geo_pb2.LanePosition, _Mapping]]]=..., positions: _Optional[_Iterable[_Union[_geo_pb2.XYPosition, _Mapping]]]=..., driving_gates: _Optional[_Iterable[_Union[_geo_pb2.XYPosition, _Mapping]]]=..., walking_gates: _Optional[_Iterable[_Union[_geo_pb2.XYPosition, _Mapping]]]=..., area: _Optional[float]=..., land_use: _Optional[_Union[LandUseType, str]]=..., urban_land_use: _Optional[str]=..., poi_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class Poi(_message.Message):
    __slots__ = ['id', 'name', 'category', 'position', 'aoi_id', 'capacity', 'functions']
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    AOI_ID_FIELD_NUMBER: _ClassVar[int]
    CAPACITY_FIELD_NUMBER: _ClassVar[int]
    FUNCTIONS_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    category: str
    position: _geo_pb2.XYPosition
    aoi_id: int
    capacity: int
    functions: _containers.RepeatedScalarFieldContainer[str]

    def __init__(self, id: _Optional[int]=..., name: _Optional[str]=..., category: _Optional[str]=..., position: _Optional[_Union[_geo_pb2.XYPosition, _Mapping]]=..., aoi_id: _Optional[int]=..., capacity: _Optional[int]=..., functions: _Optional[_Iterable[str]]=...) -> None:
        ...

class Map(_message.Message):
    __slots__ = ['header', 'lanes', 'roads', 'junctions', 'aois', 'pois']
    HEADER_FIELD_NUMBER: _ClassVar[int]
    LANES_FIELD_NUMBER: _ClassVar[int]
    ROADS_FIELD_NUMBER: _ClassVar[int]
    JUNCTIONS_FIELD_NUMBER: _ClassVar[int]
    AOIS_FIELD_NUMBER: _ClassVar[int]
    POIS_FIELD_NUMBER: _ClassVar[int]
    header: Header
    lanes: _containers.RepeatedCompositeFieldContainer[Lane]
    roads: _containers.RepeatedCompositeFieldContainer[Road]
    junctions: _containers.RepeatedCompositeFieldContainer[Junction]
    aois: _containers.RepeatedCompositeFieldContainer[Aoi]
    pois: _containers.RepeatedCompositeFieldContainer[Poi]

    def __init__(self, header: _Optional[_Union[Header, _Mapping]]=..., lanes: _Optional[_Iterable[_Union[Lane, _Mapping]]]=..., roads: _Optional[_Iterable[_Union[Road, _Mapping]]]=..., junctions: _Optional[_Iterable[_Union[Junction, _Mapping]]]=..., aois: _Optional[_Iterable[_Union[Aoi, _Mapping]]]=..., pois: _Optional[_Iterable[_Union[Poi, _Mapping]]]=...) -> None:
        ...