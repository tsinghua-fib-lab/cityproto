from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class MongoPath(_message.Message):
    __slots__ = ['db', 'col']
    DB_FIELD_NUMBER: _ClassVar[int]
    COL_FIELD_NUMBER: _ClassVar[int]
    db: str
    col: str

    def __init__(self, db: _Optional[str]=..., col: _Optional[str]=...) -> None:
        ...

class OutputTarget(_message.Message):
    __slots__ = ['sql']
    SQL_FIELD_NUMBER: _ClassVar[int]
    sql: str

    def __init__(self, sql: _Optional[str]=...) -> None:
        ...