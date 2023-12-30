from city.elec.output.v1 import output_pb2 as _output_pb2
from city.event.v1 import event_pb2 as _event_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class OutputRequest(_message.Message):
    __slots__ = ['step', 'ruined_ids', 'stopped_ids', 'aois', 'events', 'resident_unsatisfied_ratio', 'resident_demand', 'aoi_unsatisfied_ratio', 'aoi_unsatisfied_num', 'aoi_demand', 'average_powers']
    STEP_FIELD_NUMBER: _ClassVar[int]
    RUINED_IDS_FIELD_NUMBER: _ClassVar[int]
    STOPPED_IDS_FIELD_NUMBER: _ClassVar[int]
    AOIS_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    RESIDENT_UNSATISFIED_RATIO_FIELD_NUMBER: _ClassVar[int]
    RESIDENT_DEMAND_FIELD_NUMBER: _ClassVar[int]
    AOI_UNSATISFIED_RATIO_FIELD_NUMBER: _ClassVar[int]
    AOI_UNSATISFIED_NUM_FIELD_NUMBER: _ClassVar[int]
    AOI_DEMAND_FIELD_NUMBER: _ClassVar[int]
    AVERAGE_POWERS_FIELD_NUMBER: _ClassVar[int]
    step: int
    ruined_ids: _containers.RepeatedScalarFieldContainer[int]
    stopped_ids: _containers.RepeatedScalarFieldContainer[int]
    aois: _containers.RepeatedCompositeFieldContainer[_output_pb2.Aoi]
    events: _event_pb2.Events
    resident_unsatisfied_ratio: float
    resident_demand: float
    aoi_unsatisfied_ratio: float
    aoi_unsatisfied_num: int
    aoi_demand: float
    average_powers: _output_pb2.AveragePower

    def __init__(self, step: _Optional[int]=..., ruined_ids: _Optional[_Iterable[int]]=..., stopped_ids: _Optional[_Iterable[int]]=..., aois: _Optional[_Iterable[_Union[_output_pb2.Aoi, _Mapping]]]=..., events: _Optional[_Union[_event_pb2.Events, _Mapping]]=..., resident_unsatisfied_ratio: _Optional[float]=..., resident_demand: _Optional[float]=..., aoi_unsatisfied_ratio: _Optional[float]=..., aoi_unsatisfied_num: _Optional[int]=..., aoi_demand: _Optional[float]=..., average_powers: _Optional[_Union[_output_pb2.AveragePower, _Mapping]]=...) -> None:
        ...

class OutputResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...