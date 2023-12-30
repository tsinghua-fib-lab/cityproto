from city.geo.v2 import geo_pb2 as _geo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class StreetViewImage(_message.Message):
    __slots__ = ['heading', 'object']
    HEADING_FIELD_NUMBER: _ClassVar[int]
    OBJECT_FIELD_NUMBER: _ClassVar[int]
    heading: float
    object: str

    def __init__(self, heading: _Optional[float]=..., object: _Optional[str]=...) -> None:
        ...

class StreetView(_message.Message):
    __slots__ = ['lnglat', 'images']
    LNGLAT_FIELD_NUMBER: _ClassVar[int]
    IMAGES_FIELD_NUMBER: _ClassVar[int]
    lnglat: _geo_pb2.LongLatPosition
    images: _containers.RepeatedCompositeFieldContainer[StreetViewImage]

    def __init__(self, lnglat: _Optional[_Union[_geo_pb2.LongLatPosition, _Mapping]]=..., images: _Optional[_Iterable[_Union[StreetViewImage, _Mapping]]]=...) -> None:
        ...

class StreetViews(_message.Message):
    __slots__ = ['street_views']
    STREET_VIEWS_FIELD_NUMBER: _ClassVar[int]
    street_views: _containers.RepeatedCompositeFieldContainer[StreetView]

    def __init__(self, street_views: _Optional[_Iterable[_Union[StreetView, _Mapping]]]=...) -> None:
        ...