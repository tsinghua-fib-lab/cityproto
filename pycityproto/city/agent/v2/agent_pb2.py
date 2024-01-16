"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19city/agent/v2/agent.proto\x12\rcity.agent.v2\x1a\x15city/geo/v2/geo.proto\x1a\x17city/trip/v2/trip.proto"\xdf\x02\n\x0eAgentAttribute\x12,\n\x04type\x18\x01 \x01(\x0e2\x18.city.agent.v2.AgentTypeR\x04type\x12\x16\n\x06length\x18\x02 \x01(\x01R\x06length\x12\x14\n\x05width\x18\x03 \x01(\x01R\x05width\x12\x1b\n\tmax_speed\x18\x04 \x01(\x01R\x08maxSpeed\x12)\n\x10max_acceleration\x18\x05 \x01(\x01R\x0fmaxAcceleration\x128\n\x18max_braking_acceleration\x18\x06 \x01(\x01R\x16maxBrakingAcceleration\x12-\n\x12usual_acceleration\x18\x07 \x01(\x01R\x11usualAcceleration\x12<\n\x1ausual_braking_acceleration\x18\x08 \x01(\x01R\x18usualBrakingAcceleration:\x02\x18\x01"Y\n\x10VehicleAttribute\x12,\n\x12lane_change_length\x18\x01 \x01(\x01R\x10laneChangeLength\x12\x17\n\x07min_gap\x18\x02 \x01(\x01R\x06minGap"C\n\x0cBusAttribute\x12\x17\n\x07line_id\x18\x01 \x01(\x05R\x06lineId\x12\x1a\n\x08capacity\x18\x02 \x01(\x05R\x08capacity"\x0f\n\rBikeAttribute"\xcd\x04\n\x05Agent\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12;\n\tattribute\x18\x02 \x01(\x0b2\x1d.city.agent.v2.AgentAttributeR\tattribute\x12)\n\x04home\x18\x03 \x01(\x0b2\x15.city.geo.v2.PositionR\x04home\x124\n\tschedules\x18\x04 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules\x12Q\n\x11vehicle_attribute\x18\x07 \x01(\x0b2\x1f.city.agent.v2.VehicleAttributeH\x00R\x10vehicleAttribute\x88\x01\x01\x12E\n\rbus_attribute\x18\x08 \x01(\x0b2\x1b.city.agent.v2.BusAttributeH\x01R\x0cbusAttribute\x88\x01\x01\x12H\n\x0ebike_attribute\x18\t \x01(\x0b2\x1c.city.agent.v2.BikeAttributeH\x02R\rbikeAttribute\x88\x01\x01\x128\n\x06labels\x18\n \x03(\x0b2 .city.agent.v2.Agent.LabelsEntryR\x06labels\x1a9\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x028\x01:\x02\x18\x01B\x14\n\x12_vehicle_attributeB\x10\n\x0e_bus_attributeB\x11\n\x0f_bike_attribute":\n\x06Agents\x12,\n\x06agents\x18\x01 \x03(\x0b2\x14.city.agent.v2.AgentR\x06agents:\x02\x18\x01*n\n\tAgentType\x12\x1a\n\x16AGENT_TYPE_UNSPECIFIED\x10\x00\x12\x15\n\x11AGENT_TYPE_PERSON\x10\x01\x12\x1a\n\x16AGENT_TYPE_PRIVATE_CAR\x10\x02\x12\x12\n\x0eAGENT_TYPE_BUS\x10\x03B\xa9\x01\n\x11com.city.agent.v2B\nAgentProtoP\x01Z2git.fiblab.net/sim/protos/go/city/agent/v2;agentv2\xa2\x02\x03CAX\xaa\x02\rCity.Agent.V2\xca\x02\rCity\\Agent\\V2\xe2\x02\x19City\\Agent\\V2\\GPBMetadata\xea\x02\x0fCity::Agent::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.agent.v2.agent_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x11com.city.agent.v2B\nAgentProtoP\x01Z2git.fiblab.net/sim/protos/go/city/agent/v2;agentv2\xa2\x02\x03CAX\xaa\x02\rCity.Agent.V2\xca\x02\rCity\\Agent\\V2\xe2\x02\x19City\\Agent\\V2\\GPBMetadata\xea\x02\x0fCity::Agent::V2'
    _globals['_AGENTATTRIBUTE']._options = None
    _globals['_AGENTATTRIBUTE']._serialized_options = b'\x18\x01'
    _globals['_AGENT_LABELSENTRY']._options = None
    _globals['_AGENT_LABELSENTRY']._serialized_options = b'8\x01'
    _globals['_AGENT']._options = None
    _globals['_AGENT']._serialized_options = b'\x18\x01'
    _globals['_AGENTS']._options = None
    _globals['_AGENTS']._serialized_options = b'\x18\x01'
    _globals['_AGENTTYPE']._serialized_start = 1275
    _globals['_AGENTTYPE']._serialized_end = 1385
    _globals['_AGENTATTRIBUTE']._serialized_start = 93
    _globals['_AGENTATTRIBUTE']._serialized_end = 444
    _globals['_VEHICLEATTRIBUTE']._serialized_start = 446
    _globals['_VEHICLEATTRIBUTE']._serialized_end = 535
    _globals['_BUSATTRIBUTE']._serialized_start = 537
    _globals['_BUSATTRIBUTE']._serialized_end = 604
    _globals['_BIKEATTRIBUTE']._serialized_start = 606
    _globals['_BIKEATTRIBUTE']._serialized_end = 621
    _globals['_AGENT']._serialized_start = 624
    _globals['_AGENT']._serialized_end = 1213
    _globals['_AGENT_LABELSENTRY']._serialized_start = 1093
    _globals['_AGENT_LABELSENTRY']._serialized_end = 1150
    _globals['_AGENTS']._serialized_start = 1215
    _globals['_AGENTS']._serialized_end = 1273