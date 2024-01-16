"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.agent.v2 import agent_pb2 as city_dot_agent_dot_v2_dot_agent__pb2
from ....city.agent.v2 import motion_pb2 as city_dot_agent_dot_v2_dot_motion__pb2
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!city/agent/v2/agent_service.proto\x12\rcity.agent.v2\x1a\x19city/agent/v2/agent.proto\x1a\x1acity/agent/v2/motion.proto\x1a\x15city/geo/v2/geo.proto\x1a\x17city/trip/v2/trip.proto",\n\x0fGetAgentRequest\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId"p\n\x10GetAgentResponse\x12(\n\x04base\x18\x01 \x01(\x0b2\x14.city.agent.v2.AgentR\x04base\x122\n\x06motion\x18\x02 \x01(\x0b2\x1a.city.agent.v2.AgentMotionR\x06motion"=\n\x0fAddAgentRequest\x12*\n\x05agent\x18\x01 \x01(\x0b2\x14.city.agent.v2.AgentR\x05agent"-\n\x10AddAgentResponse\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId"e\n\x12SetScheduleRequest\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId\x124\n\tschedules\x18\x02 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules"\x15\n\x13SetScheduleResponse"Q\n\x1dGetAgentsByLongLatAreaRequest\x120\n\x04area\x18\x01 \x01(\x0b2\x1c.city.geo.v2.LongLatRectAreaR\x04area"j\n\x1eGetAgentsByLongLatAreaResponse\x12\x12\n\x04step\x18\x01 \x01(\x05R\x04step\x124\n\x07motions\x18\x02 \x03(\x0b2\x1a.city.agent.v2.AgentMotionR\x07motions2\xf5\x02\n\x0cAgentService\x12K\n\x08GetAgent\x12\x1e.city.agent.v2.GetAgentRequest\x1a\x1f.city.agent.v2.GetAgentResponse\x12K\n\x08AddAgent\x12\x1e.city.agent.v2.AddAgentRequest\x1a\x1f.city.agent.v2.AddAgentResponse\x12T\n\x0bSetSchedule\x12!.city.agent.v2.SetScheduleRequest\x1a".city.agent.v2.SetScheduleResponse\x12u\n\x16GetAgentsByLongLatArea\x12,.city.agent.v2.GetAgentsByLongLatAreaRequest\x1a-.city.agent.v2.GetAgentsByLongLatAreaResponseB\xb0\x01\n\x11com.city.agent.v2B\x11AgentServiceProtoP\x01Z2git.fiblab.net/sim/protos/go/city/agent/v2;agentv2\xa2\x02\x03CAX\xaa\x02\rCity.Agent.V2\xca\x02\rCity\\Agent\\V2\xe2\x02\x19City\\Agent\\V2\\GPBMetadata\xea\x02\x0fCity::Agent::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.agent.v2.agent_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x11com.city.agent.v2B\x11AgentServiceProtoP\x01Z2git.fiblab.net/sim/protos/go/city/agent/v2;agentv2\xa2\x02\x03CAX\xaa\x02\rCity.Agent.V2\xca\x02\rCity\\Agent\\V2\xe2\x02\x19City\\Agent\\V2\\GPBMetadata\xea\x02\x0fCity::Agent::V2'
    _globals['_GETAGENTREQUEST']._serialized_start = 155
    _globals['_GETAGENTREQUEST']._serialized_end = 199
    _globals['_GETAGENTRESPONSE']._serialized_start = 201
    _globals['_GETAGENTRESPONSE']._serialized_end = 313
    _globals['_ADDAGENTREQUEST']._serialized_start = 315
    _globals['_ADDAGENTREQUEST']._serialized_end = 376
    _globals['_ADDAGENTRESPONSE']._serialized_start = 378
    _globals['_ADDAGENTRESPONSE']._serialized_end = 423
    _globals['_SETSCHEDULEREQUEST']._serialized_start = 425
    _globals['_SETSCHEDULEREQUEST']._serialized_end = 526
    _globals['_SETSCHEDULERESPONSE']._serialized_start = 528
    _globals['_SETSCHEDULERESPONSE']._serialized_end = 549
    _globals['_GETAGENTSBYLONGLATAREAREQUEST']._serialized_start = 551
    _globals['_GETAGENTSBYLONGLATAREAREQUEST']._serialized_end = 632
    _globals['_GETAGENTSBYLONGLATAREARESPONSE']._serialized_start = 634
    _globals['_GETAGENTSBYLONGLATAREARESPONSE']._serialized_end = 740
    _globals['_AGENTSERVICE']._serialized_start = 743
    _globals['_AGENTSERVICE']._serialized_end = 1116