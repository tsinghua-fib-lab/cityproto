"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ......city.traffic_light.v2 import traffic_light_pb2 as city_dot_traffic__light_dot_v2_dot_traffic__light__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nEcity/traffic/interaction/traffic_light/v2/traffic_light_service.proto\x12)city.traffic.interaction.traffic_light.v2\x1a)city/traffic_light/v2/traffic_light.proto"9\n\x16GetTrafficLightRequest\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId"\xab\x01\n\x17GetTrafficLightResponse\x12H\n\rtraffic_light\x18\x01 \x01(\x0b2#.city.traffic_light.v2.TrafficLightR\x0ctrafficLight\x12\x1f\n\x0bphase_index\x18\x02 \x01(\x05R\nphaseIndex\x12%\n\x0etime_remaining\x18\x03 \x01(\x01R\rtimeRemaining"\xaa\x01\n\x16SetTrafficLightRequest\x12H\n\rtraffic_light\x18\x01 \x01(\x0b2#.city.traffic_light.v2.TrafficLightR\x0ctrafficLight\x12\x1f\n\x0bphase_index\x18\x02 \x01(\x05R\nphaseIndex\x12%\n\x0etime_remaining\x18\x03 \x01(\x01R\rtimeRemaining"\x19\n\x17SetTrafficLightResponse"\x86\x01\n\x1bSetTrafficLightPhaseRequest\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId\x12\x1f\n\x0bphase_index\x18\x02 \x01(\x05R\nphaseIndex\x12%\n\x0etime_remaining\x18\x03 \x01(\x01R\rtimeRemaining"\x1e\n\x1cSetTrafficLightPhaseResponse"O\n\x1cSetTrafficLightStatusRequest\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId\x12\x0e\n\x02ok\x18\x02 \x01(\x08R\x02ok"\x1f\n\x1dSetTrafficLightStatusResponse2\xa2\x05\n\x13TrafficLightService\x12\x98\x01\n\x0fGetTrafficLight\x12A.city.traffic.interaction.traffic_light.v2.GetTrafficLightRequest\x1aB.city.traffic.interaction.traffic_light.v2.GetTrafficLightResponse\x12\x98\x01\n\x0fSetTrafficLight\x12A.city.traffic.interaction.traffic_light.v2.SetTrafficLightRequest\x1aB.city.traffic.interaction.traffic_light.v2.SetTrafficLightResponse\x12\xa7\x01\n\x14SetTrafficLightPhase\x12F.city.traffic.interaction.traffic_light.v2.SetTrafficLightPhaseRequest\x1aG.city.traffic.interaction.traffic_light.v2.SetTrafficLightPhaseResponse\x12\xaa\x01\n\x15SetTrafficLightStatus\x12G.city.traffic.interaction.traffic_light.v2.SetTrafficLightStatusRequest\x1aH.city.traffic.interaction.traffic_light.v2.SetTrafficLightStatusResponseB\xe6\x02\n-com.city.traffic.interaction.traffic_light.v2B\x18TrafficLightServiceProtoP\x01ZVgit.fiblab.net/sim/protos/go/city/traffic/interaction/traffic_light/v2;traffic_lightv2\xa2\x02\x04CTIT\xaa\x02(City.Traffic.Interaction.TrafficLight.V2\xca\x02(City\\Traffic\\Interaction\\TrafficLight\\V2\xe2\x024City\\Traffic\\Interaction\\TrafficLight\\V2\\GPBMetadata\xea\x02,City::Traffic::Interaction::TrafficLight::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.traffic.interaction.traffic_light.v2.traffic_light_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n-com.city.traffic.interaction.traffic_light.v2B\x18TrafficLightServiceProtoP\x01ZVgit.fiblab.net/sim/protos/go/city/traffic/interaction/traffic_light/v2;traffic_lightv2\xa2\x02\x04CTIT\xaa\x02(City.Traffic.Interaction.TrafficLight.V2\xca\x02(City\\Traffic\\Interaction\\TrafficLight\\V2\xe2\x024City\\Traffic\\Interaction\\TrafficLight\\V2\\GPBMetadata\xea\x02,City::Traffic::Interaction::TrafficLight::V2'
    _globals['_GETTRAFFICLIGHTREQUEST']._serialized_start = 159
    _globals['_GETTRAFFICLIGHTREQUEST']._serialized_end = 216
    _globals['_GETTRAFFICLIGHTRESPONSE']._serialized_start = 219
    _globals['_GETTRAFFICLIGHTRESPONSE']._serialized_end = 390
    _globals['_SETTRAFFICLIGHTREQUEST']._serialized_start = 393
    _globals['_SETTRAFFICLIGHTREQUEST']._serialized_end = 563
    _globals['_SETTRAFFICLIGHTRESPONSE']._serialized_start = 565
    _globals['_SETTRAFFICLIGHTRESPONSE']._serialized_end = 590
    _globals['_SETTRAFFICLIGHTPHASEREQUEST']._serialized_start = 593
    _globals['_SETTRAFFICLIGHTPHASEREQUEST']._serialized_end = 727
    _globals['_SETTRAFFICLIGHTPHASERESPONSE']._serialized_start = 729
    _globals['_SETTRAFFICLIGHTPHASERESPONSE']._serialized_end = 759
    _globals['_SETTRAFFICLIGHTSTATUSREQUEST']._serialized_start = 761
    _globals['_SETTRAFFICLIGHTSTATUSREQUEST']._serialized_end = 840
    _globals['_SETTRAFFICLIGHTSTATUSRESPONSE']._serialized_start = 842
    _globals['_SETTRAFFICLIGHTSTATUSRESPONSE']._serialized_end = 873
    _globals['_TRAFFICLIGHTSERVICE']._serialized_start = 876
    _globals['_TRAFFICLIGHTSERVICE']._serialized_end = 1550