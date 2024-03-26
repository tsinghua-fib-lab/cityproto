from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class LongLatPosition(_message.Message):
    __slots__ = ['longitude', 'latitude', 'z']
    LONGITUDE_FIELD_NUMBER: _ClassVar[int]
    LATITUDE_FIELD_NUMBER: _ClassVar[int]
    Z_FIELD_NUMBER: _ClassVar[int]
    longitude: float
    latitude: float
    z: float

    def __init__(self, longitude: _Optional[float]=..., latitude: _Optional[float]=..., z: _Optional[float]=...) -> None:
        ...

class XYPosition(_message.Message):
    __slots__ = ['x', 'y', 'z']
    X_FIELD_NUMBER: _ClassVar[int]
    Y_FIELD_NUMBER: _ClassVar[int]
    Z_FIELD_NUMBER: _ClassVar[int]
    x: float
    y: float
    z: float

    def __init__(self, x: _Optional[float]=..., y: _Optional[float]=..., z: _Optional[float]=...) -> None:
        ...

class LanePosition(_message.Message):
    __slots__ = ['lane_id', 's']
    LANE_ID_FIELD_NUMBER: _ClassVar[int]
    S_FIELD_NUMBER: _ClassVar[int]
    lane_id: int
    s: float

    def __init__(self, lane_id: _Optional[int]=..., s: _Optional[float]=...) -> None:
        ...

class AoiPosition(_message.Message):
    __slots__ = ['aoi_id', 'poi_id']
    AOI_ID_FIELD_NUMBER: _ClassVar[int]
    POI_ID_FIELD_NUMBER: _ClassVar[int]
    aoi_id: int
    poi_id: int

    def __init__(self, aoi_id: _Optional[int]=..., poi_id: _Optional[int]=...) -> None:
        ...

class Position(_message.Message):
    __slots__ = ['lane_position', 'aoi_position', 'longlat_position', 'xy_position']
    LANE_POSITION_FIELD_NUMBER: _ClassVar[int]
    AOI_POSITION_FIELD_NUMBER: _ClassVar[int]
    LONGLAT_POSITION_FIELD_NUMBER: _ClassVar[int]
    XY_POSITION_FIELD_NUMBER: _ClassVar[int]
    lane_position: LanePosition
    aoi_position: AoiPosition
    longlat_position: LongLatPosition
    xy_position: XYPosition

    def __init__(self, lane_position: _Optional[_Union[LanePosition, _Mapping]]=..., aoi_position: _Optional[_Union[AoiPosition, _Mapping]]=..., longlat_position: _Optional[_Union[LongLatPosition, _Mapping]]=..., xy_position: _Optional[_Union[XYPosition, _Mapping]]=...) -> None:
        ...

class LongLatBBox(_message.Message):
    __slots__ = ['min_longitude', 'min_latitude', 'max_longitude', 'max_latitude']
    MIN_LONGITUDE_FIELD_NUMBER: _ClassVar[int]
    MIN_LATITUDE_FIELD_NUMBER: _ClassVar[int]
    MAX_LONGITUDE_FIELD_NUMBER: _ClassVar[int]
    MAX_LATITUDE_FIELD_NUMBER: _ClassVar[int]
    min_longitude: float
    min_latitude: float
    max_longitude: float
    max_latitude: float

    def __init__(self, min_longitude: _Optional[float]=..., min_latitude: _Optional[float]=..., max_longitude: _Optional[float]=..., max_latitude: _Optional[float]=...) -> None:
        ...