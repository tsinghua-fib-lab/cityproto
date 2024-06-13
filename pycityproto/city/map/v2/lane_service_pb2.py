"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.map.v2 import lane_state_pb2 as city_dot_map_dot_v2_dot_lane__state__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ecity/map/v2/lane_service.proto\x12\x0bcity.map.v2\x1a\x15city/geo/v2/geo.proto\x1a\x1ccity/map/v2/lane_state.proto"B\n\x12SetLaneMaxVRequest\x12\x17\n\x07lane_id\x18\x01 \x01(\x05R\x06laneId\x12\x13\n\x05max_v\x18\x02 \x01(\x01R\x04maxV"\x15\n\x13SetLaneMaxVResponse"V\n\x19SetLaneRestrictionRequest\x12\x17\n\x07lane_id\x18\x01 \x01(\x05R\x06laneId\x12 \n\x0brestriction\x18\x02 \x01(\x08R\x0brestriction"\x1c\n\x1aSetLaneRestrictionResponse"R\n\x0eGetLaneRequest\x12\x19\n\x08lane_ids\x18\x01 \x03(\x05R\x07laneIds\x12%\n\x0eexclude_person\x18\x02 \x01(\x08R\rexcludePerson"A\n\x0fGetLaneResponse\x12.\n\x06states\x18\x01 \x03(\x0b2\x16.city.map.v2.LaneStateR\x06states"r\n\x1bGetLaneByLongLatBBoxRequest\x12,\n\x04bbox\x18\x01 \x01(\x0b2\x18.city.geo.v2.LongLatBBoxR\x04bbox\x12%\n\x0eexclude_person\x18\x02 \x01(\x08R\rexcludePerson"N\n\x1cGetLaneByLongLatBBoxResponse\x12.\n\x06states\x18\x01 \x03(\x0b2\x16.city.map.v2.LaneStateR\x06states2\xf9\x02\n\x0bLaneService\x12P\n\x0bSetLaneMaxV\x12\x1f.city.map.v2.SetLaneMaxVRequest\x1a .city.map.v2.SetLaneMaxVResponse\x12e\n\x12SetLaneRestriction\x12&.city.map.v2.SetLaneRestrictionRequest\x1a\'.city.map.v2.SetLaneRestrictionResponse\x12D\n\x07GetLane\x12\x1b.city.map.v2.GetLaneRequest\x1a\x1c.city.map.v2.GetLaneResponse\x12k\n\x14GetLaneByLongLatBBox\x12(.city.map.v2.GetLaneByLongLatBBoxRequest\x1a).city.map.v2.GetLaneByLongLatBBoxResponseB\xa1\x01\n\x0fcom.city.map.v2B\x10LaneServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.lane_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\x10LaneServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_SETLANEMAXVREQUEST']._serialized_start = 100
    _globals['_SETLANEMAXVREQUEST']._serialized_end = 166
    _globals['_SETLANEMAXVRESPONSE']._serialized_start = 168
    _globals['_SETLANEMAXVRESPONSE']._serialized_end = 189
    _globals['_SETLANERESTRICTIONREQUEST']._serialized_start = 191
    _globals['_SETLANERESTRICTIONREQUEST']._serialized_end = 277
    _globals['_SETLANERESTRICTIONRESPONSE']._serialized_start = 279
    _globals['_SETLANERESTRICTIONRESPONSE']._serialized_end = 307
    _globals['_GETLANEREQUEST']._serialized_start = 309
    _globals['_GETLANEREQUEST']._serialized_end = 391
    _globals['_GETLANERESPONSE']._serialized_start = 393
    _globals['_GETLANERESPONSE']._serialized_end = 458
    _globals['_GETLANEBYLONGLATBBOXREQUEST']._serialized_start = 460
    _globals['_GETLANEBYLONGLATBBOXREQUEST']._serialized_end = 574
    _globals['_GETLANEBYLONGLATBBOXRESPONSE']._serialized_start = 576
    _globals['_GETLANEBYLONGLATBBOXRESPONSE']._serialized_end = 654
    _globals['_LANESERVICE']._serialized_start = 657
    _globals['_LANESERVICE']._serialized_end = 1034