from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class VehicleCarbon(_message.Message):
    __slots__ = ['id', 'ds', 'v', 'a', 'u_acc', 'u_roll', 'u_aero', 'c_d']
    ID_FIELD_NUMBER: _ClassVar[int]
    DS_FIELD_NUMBER: _ClassVar[int]
    V_FIELD_NUMBER: _ClassVar[int]
    A_FIELD_NUMBER: _ClassVar[int]
    U_ACC_FIELD_NUMBER: _ClassVar[int]
    U_ROLL_FIELD_NUMBER: _ClassVar[int]
    U_AERO_FIELD_NUMBER: _ClassVar[int]
    C_D_FIELD_NUMBER: _ClassVar[int]
    id: int
    ds: float
    v: float
    a: float
    u_acc: float
    u_roll: float
    u_aero: float
    c_d: float

    def __init__(self, id: _Optional[int]=..., ds: _Optional[float]=..., v: _Optional[float]=..., a: _Optional[float]=..., u_acc: _Optional[float]=..., u_roll: _Optional[float]=..., u_aero: _Optional[float]=..., c_d: _Optional[float]=...) -> None:
        ...