"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1acity/agent/v2/motion.proto\x12\rcity.agent.v2\x1a\x15city/geo/v2/geo.proto"\xcb\x01\n\x0bAgentMotion\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12-\n\x06status\x18\x02 \x01(\x0e2\x15.city.agent.v2.StatusR\x06status\x121\n\x08position\x18\x03 \x01(\x0b2\x15.city.geo.v2.PositionR\x08position\x12\x0c\n\x01v\x18\x04 \x01(\x01R\x01v\x12\x1c\n\tdirection\x18\x05 \x01(\x01R\tdirection\x12\x1a\n\x08activity\x18\x06 \x01(\tR\x08activity:\x02\x18\x01*\x99\x01\n\x06Status\x12\x16\n\x12STATUS_UNSPECIFIED\x10\x00\x12\x10\n\x0cSTATUS_SLEEP\x10\x01\x12\x12\n\x0eSTATUS_DRIVING\x10\x02\x12\x12\n\x0eSTATUS_WALKING\x10\x03\x12\x10\n\x0cSTATUS_CROWD\x10\x04\x12\x14\n\x10STATUS_PASSENGER\x10\x05\x12\x15\n\x11STATUS_WAIT_ROUTE\x10\x06B\xaa\x01\n\x11com.city.agent.v2B\x0bMotionProtoP\x01Z2git.fiblab.net/sim/protos/go/city/agent/v2;agentv2\xa2\x02\x03CAX\xaa\x02\rCity.Agent.V2\xca\x02\rCity\\Agent\\V2\xe2\x02\x19City\\Agent\\V2\\GPBMetadata\xea\x02\x0fCity::Agent::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.agent.v2.motion_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x11com.city.agent.v2B\x0bMotionProtoP\x01Z2git.fiblab.net/sim/protos/go/city/agent/v2;agentv2\xa2\x02\x03CAX\xaa\x02\rCity.Agent.V2\xca\x02\rCity\\Agent\\V2\xe2\x02\x19City\\Agent\\V2\\GPBMetadata\xea\x02\x0fCity::Agent::V2'
    _globals['_AGENTMOTION']._options = None
    _globals['_AGENTMOTION']._serialized_options = b'\x18\x01'
    _globals['_STATUS']._serialized_start = 275
    _globals['_STATUS']._serialized_end = 428
    _globals['_AGENTMOTION']._serialized_start = 69
    _globals['_AGENTMOTION']._serialized_end = 272