from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class SetStatusRequest(_message.Message):
    __slots__ = ['id', 'status']
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    id: int
    status: bool

    def __init__(self, id: _Optional[int]=..., status: bool=...) -> None:
        ...

class SetStatusResponse(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetPowerRequest(_message.Message):
    __slots__ = ['id']
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int

    def __init__(self, id: _Optional[int]=...) -> None:
        ...

class GetPowerResponse(_message.Message):
    __slots__ = ['power']
    POWER_FIELD_NUMBER: _ClassVar[int]
    power: float

    def __init__(self, power: _Optional[float]=...) -> None:
        ...

class GetPowerStatusRequest(_message.Message):
    __slots__ = ['flag']
    FLAG_FIELD_NUMBER: _ClassVar[int]
    flag: int

    def __init__(self, flag: _Optional[int]=...) -> None:
        ...

class GetPowerStatusResponse(_message.Message):
    __slots__ = ['power_status']

    class PowerStatusEntry(_message.Message):
        __slots__ = ['key', 'value']
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: int
        value: float

        def __init__(self, key: _Optional[int]=..., value: _Optional[float]=...) -> None:
            ...
    POWER_STATUS_FIELD_NUMBER: _ClassVar[int]
    power_status: _containers.ScalarMap[int, float]

    def __init__(self, power_status: _Optional[_Mapping[int, float]]=...) -> None:
        ...

class GetNoPowerAOIRequest(_message.Message):
    __slots__ = ['flag']
    FLAG_FIELD_NUMBER: _ClassVar[int]
    flag: int

    def __init__(self, flag: _Optional[int]=...) -> None:
        ...

class GetNoPowerAOIResponse(_message.Message):
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

class GetEdgeStatusRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class GetEdgeStatusResponse(_message.Message):
    __slots__ = ['reason1', 'reason2', 'reason3']
    REASON1_FIELD_NUMBER: _ClassVar[int]
    REASON2_FIELD_NUMBER: _ClassVar[int]
    REASON3_FIELD_NUMBER: _ClassVar[int]
    reason1: _containers.RepeatedScalarFieldContainer[str]
    reason2: _containers.RepeatedScalarFieldContainer[str]
    reason3: _containers.RepeatedScalarFieldContainer[str]

    def __init__(self, reason1: _Optional[_Iterable[str]]=..., reason2: _Optional[_Iterable[str]]=..., reason3: _Optional[_Iterable[str]]=...) -> None:
        ...