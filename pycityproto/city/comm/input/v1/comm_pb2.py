"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dcity/comm/input/v1/comm.proto\x12\x12city.comm.input.v1\x1a\x15city/geo/v2/geo.proto"\x9b\x03\n\x04Node\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x120\n\x04type\x18\x02 \x01(\x0e2\x1c.city.comm.input.v1.NodeTypeR\x04type\x12\x1b\n\tparent_id\x18\x03 \x01(\x05R\x08parentId\x12!\n\x0cchildren_ids\x18\x04 \x03(\x05R\x0bchildrenIds\x126\n\x08position\x18\x05 \x01(\x0b2\x15.city.geo.v2.PositionH\x00R\x08position\x88\x01\x01\x12\x1a\n\x06aoi_id\x18\x06 \x01(\x05H\x01R\x05aoiId\x88\x01\x01\x12\'\n\rfreq_range_id\x18\x07 \x01(\x05H\x02R\x0bfreqRangeId\x88\x01\x01\x12T\n\x11base_station_type\x18\x08 \x01(\x0e2#.city.comm.input.v1.BaseStationTypeH\x03R\x0fbaseStationType\x88\x01\x01B\x0b\n\t_positionB\t\n\x07_aoi_idB\x10\n\x0e_freq_range_idB\x14\n\x12_base_station_type"i\n\rRepairStation\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x15\n\x06aoi_id\x18\x02 \x01(\x05R\x05aoiId\x121\n\x08position\x18\x03 \x01(\x0b2\x15.city.geo.v2.PositionR\x08position"I\n\x04Pump\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x121\n\x08position\x18\x02 \x01(\x0b2\x15.city.geo.v2.PositionR\x08position"6\n\nCommDemand\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x18\n\x07demands\x18\x02 \x03(\x01R\x07demands"\xb3\x01\n\x05Nodes\x12.\n\x05nodes\x18\x01 \x03(\x0b2\x18.city.comm.input.v1.NodeR\x05nodes\x12J\n\x0frepair_stations\x18\x02 \x03(\x0b2!.city.comm.input.v1.RepairStationR\x0erepairStations\x12.\n\x05pumps\x18\x03 \x03(\x0b2\x18.city.comm.input.v1.PumpR\x05pumps"P\n\x0bCommDemands\x12A\n\x0ccomm_demands\x18\x01 \x03(\x0b2\x1e.city.comm.input.v1.CommDemandR\x0bcommDemands*p\n\x08NodeType\x12\x19\n\x15NODE_TYPE_UNSPECIFIED\x10\x00\x12\x16\n\x12NODE_TYPE_INTERNET\x10\x01\x12\x15\n\x11NODE_TYPE_GATEWAY\x10\x02\x12\x1a\n\x16NODE_TYPE_BASE_STATION\x10\x03*q\n\x0fBaseStationType\x12!\n\x1dBASE_STATION_TYPE_UNSPECIFIED\x10\x00\x12\x1c\n\x18BASE_STATION_TYPE_INDOOR\x10\x01\x12\x1d\n\x19BASE_STATION_TYPE_OUTDOOR\x10\x02B\xc7\x01\n\x16com.city.comm.input.v1B\tCommProtoP\x01Z7git.fiblab.net/sim/protos/go/city/comm/input/v1;inputv1\xa2\x02\x03CCI\xaa\x02\x12City.Comm.Input.V1\xca\x02\x12City\\Comm\\Input\\V1\xe2\x02\x1eCity\\Comm\\Input\\V1\\GPBMetadata\xea\x02\x15City::Comm::Input::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.comm.input.v1.comm_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x16com.city.comm.input.v1B\tCommProtoP\x01Z7git.fiblab.net/sim/protos/go/city/comm/input/v1;inputv1\xa2\x02\x03CCI\xaa\x02\x12City.Comm.Input.V1\xca\x02\x12City\\Comm\\Input\\V1\xe2\x02\x1eCity\\Comm\\Input\\V1\\GPBMetadata\xea\x02\x15City::Comm::Input::V1'
    _globals['_NODETYPE']._serialized_start = 992
    _globals['_NODETYPE']._serialized_end = 1104
    _globals['_BASESTATIONTYPE']._serialized_start = 1106
    _globals['_BASESTATIONTYPE']._serialized_end = 1219
    _globals['_NODE']._serialized_start = 77
    _globals['_NODE']._serialized_end = 488
    _globals['_REPAIRSTATION']._serialized_start = 490
    _globals['_REPAIRSTATION']._serialized_end = 595
    _globals['_PUMP']._serialized_start = 597
    _globals['_PUMP']._serialized_end = 670
    _globals['_COMMDEMAND']._serialized_start = 672
    _globals['_COMMDEMAND']._serialized_end = 726
    _globals['_NODES']._serialized_start = 729
    _globals['_NODES']._serialized_end = 908
    _globals['_COMMDEMANDS']._serialized_start = 910
    _globals['_COMMDEMANDS']._serialized_end = 990