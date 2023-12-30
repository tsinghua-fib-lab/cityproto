from city.map.v2 import map_pb2 as _map_pb2
from city.water.input.v1 import config_pb2 as _config_pb2
from city.water.input.v1 import water_pb2 as _water_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class InitRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class InitResponse(_message.Message):
    __slots__ = ['address', 'control', 'rain', 'map']
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    CONTROL_FIELD_NUMBER: _ClassVar[int]
    RAIN_FIELD_NUMBER: _ClassVar[int]
    MAP_FIELD_NUMBER: _ClassVar[int]
    address: str
    control: _config_pb2.Control
    rain: _water_pb2.Rain
    map: _map_pb2.Map

    def __init__(self, address: _Optional[str]=..., control: _Optional[_Union[_config_pb2.Control, _Mapping]]=..., rain: _Optional[_Union[_water_pb2.Rain, _Mapping]]=..., map: _Optional[_Union[_map_pb2.Map, _Mapping]]=...) -> None:
        ...