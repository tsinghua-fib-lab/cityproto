"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%city/agentcomm/v1/agent_service.proto\x12\x11city.agentcomm.v1"x\n\x12CommunicateRequest\x12&\n\x0fsource_agent_id\x18\x01 \x01(\tR\rsourceAgentId\x12&\n\x0ftarget_agent_id\x18\x02 \x01(\tR\rtargetAgentId\x12\x12\n\x04data\x18\x03 \x01(\tR\x04data"/\n\x13CommunicateResponse\x12\x18\n\x07message\x18\x01 \x01(\tR\x07message2p\n\x0cAgentService\x12`\n\x0bCommunicate\x12%.city.agentcomm.v1.CommunicateRequest\x1a&.city.agentcomm.v1.CommunicateResponse(\x010\x01B\xcf\x01\n\x15com.city.agentcomm.v1B\x11AgentServiceProtoP\x01Z=git.fiblab.net/sim/protos/v2/go/city/agentcomm/v1;agentcommv1\xa2\x02\x03CAX\xaa\x02\x11City.Agentcomm.V1\xca\x02\x11City\\Agentcomm\\V1\xe2\x02\x1dCity\\Agentcomm\\V1\\GPBMetadata\xea\x02\x13City::Agentcomm::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.agentcomm.v1.agent_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x15com.city.agentcomm.v1B\x11AgentServiceProtoP\x01Z=git.fiblab.net/sim/protos/v2/go/city/agentcomm/v1;agentcommv1\xa2\x02\x03CAX\xaa\x02\x11City.Agentcomm.V1\xca\x02\x11City\\Agentcomm\\V1\xe2\x02\x1dCity\\Agentcomm\\V1\\GPBMetadata\xea\x02\x13City::Agentcomm::V1'
    _globals['_COMMUNICATEREQUEST']._serialized_start = 60
    _globals['_COMMUNICATEREQUEST']._serialized_end = 180
    _globals['_COMMUNICATERESPONSE']._serialized_start = 182
    _globals['_COMMUNICATERESPONSE']._serialized_end = 229
    _globals['_AGENTSERVICE']._serialized_start = 231
    _globals['_AGENTSERVICE']._serialized_end = 343