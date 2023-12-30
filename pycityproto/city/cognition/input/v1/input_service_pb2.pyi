from city.agent.v2 import agent_pb2 as _agent_pb2
from city.cognition.input.v1 import config_pb2 as _config_pb2
from city.cognition.input.v1 import input_pb2 as _input_pb2
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
    __slots__ = ['address', 'control', 'map', 'persons', 'edges']
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    CONTROL_FIELD_NUMBER: _ClassVar[int]
    MAP_FIELD_NUMBER: _ClassVar[int]
    PERSONS_FIELD_NUMBER: _ClassVar[int]
    EDGES_FIELD_NUMBER: _ClassVar[int]
    address: str
    control: _config_pb2.Control
    map: _map_pb2.Map
    persons: _agent_pb2.Agents
    edges: _input_pb2.Edges

    def __init__(self, address: _Optional[str]=..., control: _Optional[_Union[_config_pb2.Control, _Mapping]]=..., map: _Optional[_Union[_map_pb2.Map, _Mapping]]=..., persons: _Optional[_Union[_agent_pb2.Agents, _Mapping]]=..., edges: _Optional[_Union[_input_pb2.Edges, _Mapping]]=...) -> None:
        ...