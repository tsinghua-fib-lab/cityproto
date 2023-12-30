from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class Aoi(_message.Message):
    __slots__ = ['id', 'unsatisfied_demand_ratio', 'unsatisfied_demand_num', 'demand', 'supply']
    ID_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_DEMAND_RATIO_FIELD_NUMBER: _ClassVar[int]
    UNSATISFIED_DEMAND_NUM_FIELD_NUMBER: _ClassVar[int]
    DEMAND_FIELD_NUMBER: _ClassVar[int]
    SUPPLY_FIELD_NUMBER: _ClassVar[int]
    id: int
    unsatisfied_demand_ratio: float
    unsatisfied_demand_num: int
    demand: float
    supply: float

    def __init__(self, id: _Optional[int]=..., unsatisfied_demand_ratio: _Optional[float]=..., unsatisfied_demand_num: _Optional[int]=..., demand: _Optional[float]=..., supply: _Optional[float]=...) -> None:
        ...

class AveragePower(_message.Message):
    __slots__ = ['transformer_500', 'transformer_220', 'transformer_110', 'transformer_10']
    TRANSFORMER_500_FIELD_NUMBER: _ClassVar[int]
    TRANSFORMER_220_FIELD_NUMBER: _ClassVar[int]
    TRANSFORMER_110_FIELD_NUMBER: _ClassVar[int]
    TRANSFORMER_10_FIELD_NUMBER: _ClassVar[int]
    transformer_500: float
    transformer_220: float
    transformer_110: float
    transformer_10: float

    def __init__(self, transformer_500: _Optional[float]=..., transformer_220: _Optional[float]=..., transformer_110: _Optional[float]=..., transformer_10: _Optional[float]=...) -> None:
        ...