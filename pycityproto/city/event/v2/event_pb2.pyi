from city.geo.v2 import geo_pb2 as _geo_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class EntityType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    ENTITY_TYPE_UNSPECIFIED: _ClassVar[EntityType]
    ENTITY_TYPE_LANE: _ClassVar[EntityType]
    ENTITY_TYPE_ROAD: _ClassVar[EntityType]
    ENTITY_TYPE_JUNCTION: _ClassVar[EntityType]
    ENTITY_TYPE_AOI: _ClassVar[EntityType]
    ENTITY_TYPE_POI: _ClassVar[EntityType]
    ENTITY_TYPE_PERSON: _ClassVar[EntityType]
    ENTITY_TYPE_ORG: _ClassVar[EntityType]
ENTITY_TYPE_UNSPECIFIED: EntityType
ENTITY_TYPE_LANE: EntityType
ENTITY_TYPE_ROAD: EntityType
ENTITY_TYPE_JUNCTION: EntityType
ENTITY_TYPE_AOI: EntityType
ENTITY_TYPE_POI: EntityType
ENTITY_TYPE_PERSON: EntityType
ENTITY_TYPE_ORG: EntityType

class Entity(_message.Message):
    __slots__ = ['type', 'id']
    TYPE_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    type: EntityType
    id: int

    def __init__(self, type: _Optional[_Union[EntityType, str]]=..., id: _Optional[int]=...) -> None:
        ...

class Event(_message.Message):
    __slots__ = ['topic', 'id', 'subject', 'content', 'position', 't']
    TOPIC_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    POSITION_FIELD_NUMBER: _ClassVar[int]
    T_FIELD_NUMBER: _ClassVar[int]
    topic: str
    id: int
    subject: Entity
    content: str
    position: _geo_pb2.Position
    t: float

    def __init__(self, topic: _Optional[str]=..., id: _Optional[int]=..., subject: _Optional[_Union[Entity, _Mapping]]=..., content: _Optional[str]=..., position: _Optional[_Union[_geo_pb2.Position, _Mapping]]=..., t: _Optional[float]=...) -> None:
        ...