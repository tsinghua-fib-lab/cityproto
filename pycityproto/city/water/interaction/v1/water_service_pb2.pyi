from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class WaterFacilityType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    WATER_FACILITY_TYPE_UNSPECIFIED: _ClassVar[WaterFacilityType]
    WATER_FACILITY_TYPE_SUPPLY: _ClassVar[WaterFacilityType]
    WATER_FACILITY_TYPE_DRAINAGE: _ClassVar[WaterFacilityType]
WATER_FACILITY_TYPE_UNSPECIFIED: WaterFacilityType
WATER_FACILITY_TYPE_SUPPLY: WaterFacilityType
WATER_FACILITY_TYPE_DRAINAGE: WaterFacilityType

class SetPumpPowerStatusRequest(_message.Message):
    __slots__ = ['id', 'status', 'type']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: bool
    type: WaterFacilityType

    def __init__(self, id: _Optional[int]=..., status: bool=..., type: _Optional[_Union[WaterFacilityType, str]]=...) -> None:
        ...

class SetPumpPowerStatusResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class SetPumpNetworkStatusRequest(_message.Message):
    __slots__ = ['id', 'status', 'type']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: bool
    type: WaterFacilityType

    def __init__(self, id: _Optional[int]=..., status: bool=..., type: _Optional[_Union[WaterFacilityType, str]]=...) -> None:
        ...

class SetPumpNetworkStatusResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class SetPumpStatusRequest(_message.Message):
    __slots__ = ['id', 'status', 'type']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: bool
    type: WaterFacilityType

    def __init__(self, id: _Optional[int]=..., status: bool=..., type: _Optional[_Union[WaterFacilityType, str]]=...) -> None:
        ...

class SetPumpStatusResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetPumpStatusRequest(_message.Message):
    __slots__ = ['flag']
    FLAG_FIELD_NUMBER: _ClassVar[int]
    flag: int

    def __init__(self, flag: _Optional[int]=...) -> None:
        ...

class GetPumpStatusResponse(_message.Message):
    __slots__ = ['pump_status']

    class PumpStatusEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: int
        value: int

        def __init__(self, key: _Optional[int]=..., value: _Optional[int]=...) -> None:
            ...
    PUMP_STATUS_FIELD_NUMBER: _ClassVar[int]
    pump_status: _containers.ScalarMap[int, int]

    def __init__(self, pump_status: _Optional[_Mapping[int, int]]=...) -> None:
        ...

class GetNoWaterAOIRequest(_message.Message):
    __slots__ = ['flag']
    FLAG_FIELD_NUMBER: _ClassVar[int]
    flag: int

    def __init__(self, flag: _Optional[int]=...) -> None:
        ...

class GetNoWaterAOIResponse(_message.Message):
    __slots__ = ['aoi']
    AOI_FIELD_NUMBER: _ClassVar[int]
    aoi: _containers.RepeatedScalarFieldContainer[int]

    def __init__(self, aoi: _Optional[_Iterable[int]]=...) -> None:
        ...

class GetRuinInfoRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class RuinInfo(_message.Message):
    __slots__ = ['num', 'ratio']
    NUM_FIELD_NUMBER: _ClassVar[int]
    RATIO_FIELD_NUMBER: _ClassVar[int]
    num: int
    ratio: float

    def __init__(self, num: _Optional[int]=..., ratio: _Optional[float]=...) -> None:
        ...

class GetRuinInfoResponse(_message.Message):
    __slots__ = ['one', 'two', 'three']
    ONE_FIELD_NUMBER: _ClassVar[int]
    TWO_FIELD_NUMBER: _ClassVar[int]
    THREE_FIELD_NUMBER: _ClassVar[int]
    one: RuinInfo
    two: RuinInfo
    three: RuinInfo

    def __init__(self, one: _Optional[_Union[RuinInfo, _Mapping]]=..., two: _Optional[_Union[RuinInfo, _Mapping]]=..., three: _Optional[_Union[RuinInfo, _Mapping]]=...) -> None:
        ...