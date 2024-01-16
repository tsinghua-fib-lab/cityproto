from city.agent.v2 import agent_pb2 as _agent_pb2
from city.agent.v2 import motion_pb2 as _motion_pb2
from city.geo.v2 import geo_pb2 as _geo_pb2
from city.trip.v2 import trip_pb2 as _trip_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class GetAgentRequest(_message.Message):
    __slots__ = ['agent_id']
    AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    agent_id: int

    def __init__(self, agent_id: _Optional[int]=...) -> None:
        ...

class GetAgentResponse(_message.Message):
    __slots__ = ['base', 'motion']
    BASE_FIELD_NUMBER: _ClassVar[int]
    MOTION_FIELD_NUMBER: _ClassVar[int]
    base: _agent_pb2.Agent
    motion: _motion_pb2.AgentMotion

    def __init__(self, base: _Optional[_Union[_agent_pb2.Agent, _Mapping]]=..., motion: _Optional[_Union[_motion_pb2.AgentMotion, _Mapping]]=...) -> None:
        ...

class AddAgentRequest(_message.Message):
    __slots__ = ['agent']
    AGENT_FIELD_NUMBER: _ClassVar[int]
    agent: _agent_pb2.Agent

    def __init__(self, agent: _Optional[_Union[_agent_pb2.Agent, _Mapping]]=...) -> None:
        ...

class AddAgentResponse(_message.Message):
    __slots__ = ['agent_id']
    AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    agent_id: int

    def __init__(self, agent_id: _Optional[int]=...) -> None:
        ...

class SetScheduleRequest(_message.Message):
    __slots__ = ['agent_id', 'schedules']
    AGENT_ID_FIELD_NUMBER: _ClassVar[int]
    SCHEDULES_FIELD_NUMBER: _ClassVar[int]
    agent_id: int
    schedules: _containers.RepeatedCompositeFieldContainer[_trip_pb2.Schedule]

    def __init__(self, agent_id: _Optional[int]=..., schedules: _Optional[_Iterable[_Union[_trip_pb2.Schedule, _Mapping]]]=...) -> None:
        ...

class SetScheduleResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetAgentsByLongLatAreaRequest(_message.Message):
    __slots__ = ['area']
    AREA_FIELD_NUMBER: _ClassVar[int]
    area: _geo_pb2.LongLatBBox

    def __init__(self, area: _Optional[_Union[_geo_pb2.LongLatBBox, _Mapping]]=...) -> None:
        ...

class GetAgentsByLongLatAreaResponse(_message.Message):
    __slots__ = ['step', 'motions']
    STEP_FIELD_NUMBER: _ClassVar[int]
    MOTIONS_FIELD_NUMBER: _ClassVar[int]
    step: int
    motions: _containers.RepeatedCompositeFieldContainer[_motion_pb2.AgentMotion]

    def __init__(self, step: _Optional[int]=..., motions: _Optional[_Iterable[_Union[_motion_pb2.AgentMotion, _Mapping]]]=...) -> None:
        ...