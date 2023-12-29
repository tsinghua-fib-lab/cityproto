from city.traffic.interaction.aoi.v1 import aoi_pb2 as _aoi_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetAoiRequest(_message.Message):
    __slots__ = ('aoi_ids',)
    AOI_IDS_FIELD_NUMBER: _ClassVar[int]
    aoi_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, aoi_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetAoiResponse(_message.Message):
    __slots__ = ('states',)
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedCompositeFieldContainer[_aoi_pb2.State]

    def __init__(self, states: _Optional[_Iterable[_Union[_aoi_pb2.State, _Mapping]]]=...) -> None:
        ...