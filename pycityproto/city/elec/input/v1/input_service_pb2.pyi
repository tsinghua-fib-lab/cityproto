from city.elec.input.v1 import config_pb2 as _config_pb2
from city.elec.input.v1 import input_pb2 as _input_pb2
from city.map.v2 import map_pb2 as _map_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class InitRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class InitResponse(_message.Message):
    __slots__ = ['address', 'control', 'facilities', 'map']
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    CONTROL_FIELD_NUMBER: _ClassVar[int]
    FACILITIES_FIELD_NUMBER: _ClassVar[int]
    MAP_FIELD_NUMBER: _ClassVar[int]
    address: str
    control: _config_pb2.Control
    facilities: _input_pb2.Facilities
    map: _map_pb2.Map

    def __init__(self, address: _Optional[str]=..., control: _Optional[_Union[_config_pb2.Control, _Mapping]]=..., facilities: _Optional[_Union[_input_pb2.Facilities, _Mapping]]=..., map: _Optional[_Union[_map_pb2.Map, _Mapping]]=...) -> None:
        ...