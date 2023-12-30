from city.config.v1 import config_pb2 as _config_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Mongo(_message.Message):
    __slots__ = ['uri', 'map', 'facilities']
    URI_FIELD_NUMBER: _ClassVar[int]
    MAP_FIELD_NUMBER: _ClassVar[int]
    FACILITIES_FIELD_NUMBER: _ClassVar[int]
    uri: str
    map: _config_pb2.MongoPath
    facilities: _config_pb2.MongoPath

    def __init__(self, uri: _Optional[str]=..., map: _Optional[_Union[_config_pb2.MongoPath, _Mapping]]=..., facilities: _Optional[_Union[_config_pb2.MongoPath, _Mapping]]=...) -> None:
        ...

class ControlStep(_message.Message):
    __slots__ = ['start', 'total']
    START_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    start: int
    total: int

    def __init__(self, start: _Optional[int]=..., total: _Optional[int]=...) -> None:
        ...

class Control(_message.Message):
    __slots__ = ['step']
    STEP_FIELD_NUMBER: _ClassVar[int]
    step: ControlStep

    def __init__(self, step: _Optional[_Union[ControlStep, _Mapping]]=...) -> None:
        ...

class OutputSwitch(_message.Message):
    __slots__ = ['node', 'aoi', 'event']
    NODE_FIELD_NUMBER: _ClassVar[int]
    AOI_FIELD_NUMBER: _ClassVar[int]
    EVENT_FIELD_NUMBER: _ClassVar[int]
    node: bool
    aoi: bool
    event: bool

    def __init__(self, node: bool=..., aoi: bool=..., event: bool=...) -> None:
        ...

class Output(_message.Message):
    __slots__ = ['target', 'switch']
    TARGET_FIELD_NUMBER: _ClassVar[int]
    SWITCH_FIELD_NUMBER: _ClassVar[int]
    target: _config_pb2.OutputTarget
    switch: OutputSwitch

    def __init__(self, target: _Optional[_Union[_config_pb2.OutputTarget, _Mapping]]=..., switch: _Optional[_Union[OutputSwitch, _Mapping]]=...) -> None:
        ...

class Config(_message.Message):
    __slots__ = ['mongo', 'control', 'output']
    MONGO_FIELD_NUMBER: _ClassVar[int]
    CONTROL_FIELD_NUMBER: _ClassVar[int]
    OUTPUT_FIELD_NUMBER: _ClassVar[int]
    mongo: Mongo
    control: Control
    output: Output

    def __init__(self, mongo: _Optional[_Union[Mongo, _Mapping]]=..., control: _Optional[_Union[Control, _Mapping]]=..., output: _Optional[_Union[Output, _Mapping]]=...) -> None:
        ...