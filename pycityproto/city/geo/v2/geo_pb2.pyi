from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class LongLatPosition(_message.Message):
    __slots__ = ['longitude', 'latitude']
    LONGITUDE_FIELD_NUMBER: _ClassVar[int]
    LATITUDE_FIELD_NUMBER: _ClassVar[int]
    longitude: float
    latitude: float

    def __init__(self, longitude: _Optional[float]=..., latitude: _Optional[float]=...) -> None:
        ...

class XYPosition(_message.Message):
    __slots__ = ['x', 'y']
    X_FIELD_NUMBER: _ClassVar[int]
    Y_FIELD_NUMBER: _ClassVar[int]
    x: float
    y: float

    def __init__(self, x: _Optional[float]=..., y: _Optional[float]=...) -> None:
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

class LongLatRectArea(_message.Message):
    __slots__ = ['ne', 'sw']
    NE_FIELD_NUMBER: _ClassVar[int]
    SW_FIELD_NUMBER: _ClassVar[int]
    ne: LongLatPosition
    sw: LongLatPosition

    def __init__(self, ne: _Optional[_Union[LongLatPosition, _Mapping]]=..., sw: _Optional[_Union[LongLatPosition, _Mapping]]=...) -> None:
        ...