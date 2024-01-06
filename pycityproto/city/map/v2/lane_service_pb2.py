"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.agent.v2 import motion_pb2 as city_dot_agent_dot_v2_dot_motion__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ecity/map/v2/lane_service.proto\x12\x0bcity.map.v2\x1a\x1acity/agent/v2/motion.proto"B\n\x12SetLaneMaxVRequest\x12\x17\n\x07lane_id\x18\x01 \x01(\x05R\x06laneId\x12\x13\n\x05max_v\x18\x02 \x01(\x01R\x04maxV"\x15\n\x13SetLaneMaxVResponse"+\n\x0eGetLaneRequest\x12\x19\n\x08lane_ids\x18\x01 \x03(\x05R\x07laneIds"A\n\x0fGetLaneResponse\x12.\n\x06states\x18\x01 \x03(\x0b2\x16.city.map.v2.LaneStateR\x06states"\x88\x01\n\tLaneState\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x124\n\x07persons\x18\x02 \x03(\x0b2\x1a.city.agent.v2.AgentMotionR\x07persons\x12\x13\n\x05avg_v\x18\x03 \x01(\x01R\x04avgV\x12 \n\x0brestriction\x18\x04 \x01(\x08R\x0brestriction2\xa5\x01\n\x0bLaneService\x12P\n\x0bSetLaneMaxV\x12\x1f.city.map.v2.SetLaneMaxVRequest\x1a .city.map.v2.SetLaneMaxVResponse\x12D\n\x07GetLane\x12\x1b.city.map.v2.GetLaneRequest\x1a\x1c.city.map.v2.GetLaneResponseB\xa1\x01\n\x0fcom.city.map.v2B\x10LaneServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.lane_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\x10LaneServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_SETLANEMAXVREQUEST']._serialized_start = 75
    _globals['_SETLANEMAXVREQUEST']._serialized_end = 141
    _globals['_SETLANEMAXVRESPONSE']._serialized_start = 143
    _globals['_SETLANEMAXVRESPONSE']._serialized_end = 164
    _globals['_GETLANEREQUEST']._serialized_start = 166
    _globals['_GETLANEREQUEST']._serialized_end = 209
    _globals['_GETLANERESPONSE']._serialized_start = 211
    _globals['_GETLANERESPONSE']._serialized_end = 276
    _globals['_LANESTATE']._serialized_start = 279
    _globals['_LANESTATE']._serialized_end = 415
    _globals['_LANESERVICE']._serialized_start = 418
    _globals['_LANESERVICE']._serialized_end = 583