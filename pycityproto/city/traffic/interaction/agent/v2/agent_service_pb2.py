"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ......city.agent.v2 import agent_pb2 as city_dot_agent_dot_v2_dot_agent__pb2
from ......city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ......city.traffic.interaction.agent.v2 import agent_pb2 as city_dot_traffic_dot_interaction_dot_agent_dot_v2_dot_agent__pb2
from ......city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n5city/traffic/interaction/agent/v2/agent_service.proto\x12!city.traffic.interaction.agent.v2\x1a\x19city/agent/v2/agent.proto\x1a\x15city/geo/v2/geo.proto\x1a-city/traffic/interaction/agent/v2/agent.proto\x1a\x17city/trip/v2/trip.proto",\n\x0fGetAgentRequest\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId"\x87\x01\n\x10GetAgentResponse\x12(\n\x04base\x18\x01 \x01(\x0b2\x14.city.agent.v2.AgentR\x04base\x12I\n\x07runtime\x18\x02 \x01(\x0b2/.city.traffic.interaction.agent.v2.AgentRuntimeR\x07runtime"=\n\x0fAddAgentRequest\x12*\n\x05agent\x18\x01 \x01(\x0b2\x14.city.agent.v2.AgentR\x05agent"-\n\x10AddAgentResponse\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId"e\n\x12SetScheduleRequest\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId\x124\n\tschedules\x18\x02 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules"\x15\n\x13SetScheduleResponse"Q\n\x1dGetAgentsByLongLatAreaRequest\x120\n\x04area\x18\x01 \x01(\x0b2\x1c.city.geo.v2.LongLatRectAreaR\x04area"}\n\x1eGetAgentsByLongLatAreaResponse\x12\x12\n\x04step\x18\x01 \x01(\x05R\x04step\x12G\n\x06agents\x18\x02 \x03(\x0b2/.city.traffic.interaction.agent.v2.AgentRuntimeR\x06agents2\x96\x04\n\x0cAgentService\x12s\n\x08GetAgent\x122.city.traffic.interaction.agent.v2.GetAgentRequest\x1a3.city.traffic.interaction.agent.v2.GetAgentResponse\x12s\n\x08AddAgent\x122.city.traffic.interaction.agent.v2.AddAgentRequest\x1a3.city.traffic.interaction.agent.v2.AddAgentResponse\x12|\n\x0bSetSchedule\x125.city.traffic.interaction.agent.v2.SetScheduleRequest\x1a6.city.traffic.interaction.agent.v2.SetScheduleResponse\x12\x9d\x01\n\x16GetAgentsByLongLatArea\x12@.city.traffic.interaction.agent.v2.GetAgentsByLongLatAreaRequest\x1aA.city.traffic.interaction.agent.v2.GetAgentsByLongLatAreaResponseB\xab\x02\n%com.city.traffic.interaction.agent.v2B\x11AgentServiceProtoP\x01ZFgit.fiblab.net/sim/protos/go/city/traffic/interaction/agent/v2;agentv2\xa2\x02\x04CTIA\xaa\x02!City.Traffic.Interaction.Agent.V2\xca\x02!City\\Traffic\\Interaction\\Agent\\V2\xe2\x02-City\\Traffic\\Interaction\\Agent\\V2\\GPBMetadata\xea\x02%City::Traffic::Interaction::Agent::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.traffic.interaction.agent.v2.agent_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n%com.city.traffic.interaction.agent.v2B\x11AgentServiceProtoP\x01ZFgit.fiblab.net/sim/protos/go/city/traffic/interaction/agent/v2;agentv2\xa2\x02\x04CTIA\xaa\x02!City.Traffic.Interaction.Agent.V2\xca\x02!City\\Traffic\\Interaction\\Agent\\V2\xe2\x02-City\\Traffic\\Interaction\\Agent\\V2\\GPBMetadata\xea\x02%City::Traffic::Interaction::Agent::V2'
    _globals['_GETAGENTREQUEST']._serialized_start = 214
    _globals['_GETAGENTREQUEST']._serialized_end = 258
    _globals['_GETAGENTRESPONSE']._serialized_start = 261
    _globals['_GETAGENTRESPONSE']._serialized_end = 396
    _globals['_ADDAGENTREQUEST']._serialized_start = 398
    _globals['_ADDAGENTREQUEST']._serialized_end = 459
    _globals['_ADDAGENTRESPONSE']._serialized_start = 461
    _globals['_ADDAGENTRESPONSE']._serialized_end = 506
    _globals['_SETSCHEDULEREQUEST']._serialized_start = 508
    _globals['_SETSCHEDULEREQUEST']._serialized_end = 609
    _globals['_SETSCHEDULERESPONSE']._serialized_start = 611
    _globals['_SETSCHEDULERESPONSE']._serialized_end = 632
    _globals['_GETAGENTSBYLONGLATAREAREQUEST']._serialized_start = 634
    _globals['_GETAGENTSBYLONGLATAREAREQUEST']._serialized_end = 715
    _globals['_GETAGENTSBYLONGLATAREARESPONSE']._serialized_start = 717
    _globals['_GETAGENTSBYLONGLATAREARESPONSE']._serialized_end = 842
    _globals['_AGENTSERVICE']._serialized_start = 845
    _globals['_AGENTSERVICE']._serialized_end = 1379