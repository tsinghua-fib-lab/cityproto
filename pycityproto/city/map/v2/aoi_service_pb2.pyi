from city.person.v1 import motion_pb2 as _motion_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class AoiState(_message.Message):
    __slots__ = ['id', 'persons']
    ID_FIELD_NUMBER: _ClassVar[int]
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    id: int
    persons: _containers.RepeatedCompositeFieldContainer[_motion_pb2.PersonMotion]

    def __init__(self, id: _Optional[int]=..., persons: _Optional[_Iterable[_Union[_motion_pb2.PersonMotion, _Mapping]]]=...) -> None:
        ...

class GetAoiRequest(_message.Message):
    __slots__ = ['aoi_ids']
    AOI_IDS_FIELD_NUMBER: _ClassVar[int]
    aoi_ids: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, aoi_ids: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetAoiResponse(_message.Message):
    __slots__ = ['states']
    STATES_FIELD_NUMBER: _ClassVar[int]
    states: _containers.RepeatedCompositeFieldContainer[AoiState]

    def __init__(self, states: _Optional[_Iterable[_Union[AoiState, _Mapping]]]=...) -> None:
        ...