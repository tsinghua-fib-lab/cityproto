"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.traffic_light.v2 import traffic_light_pb2 as city_dot_traffic__light_dot_v2_dot_traffic__light__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1city/traffic_light/v2/traffic_light_service.proto\x12\x15city.traffic_light.v2\x1a)city/traffic_light/v2/traffic_light.proto"9\n\x16GetTrafficLightRequest\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId"\xab\x01\n\x17GetTrafficLightResponse\x12H\n\rtraffic_light\x18\x01 \x01(\x0b2#.city.traffic_light.v2.TrafficLightR\x0ctrafficLight\x12\x1f\n\x0bphase_index\x18\x02 \x01(\x05R\nphaseIndex\x12%\n\x0etime_remaining\x18\x03 \x01(\x01R\rtimeRemaining"\xaa\x01\n\x16SetTrafficLightRequest\x12H\n\rtraffic_light\x18\x01 \x01(\x0b2#.city.traffic_light.v2.TrafficLightR\x0ctrafficLight\x12\x1f\n\x0bphase_index\x18\x02 \x01(\x05R\nphaseIndex\x12%\n\x0etime_remaining\x18\x03 \x01(\x01R\rtimeRemaining"\x19\n\x17SetTrafficLightResponse"\x86\x01\n\x1bSetTrafficLightPhaseRequest\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId\x12\x1f\n\x0bphase_index\x18\x02 \x01(\x05R\nphaseIndex\x12%\n\x0etime_remaining\x18\x03 \x01(\x01R\rtimeRemaining"\x1e\n\x1cSetTrafficLightPhaseResponse"O\n\x1cSetTrafficLightStatusRequest\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId\x12\x0e\n\x02ok\x18\x02 \x01(\x08R\x02ok"\x1f\n\x1dSetTrafficLightStatusResponse2\xff\x03\n\x13TrafficLightService\x12p\n\x0fGetTrafficLight\x12-.city.traffic_light.v2.GetTrafficLightRequest\x1a..city.traffic_light.v2.GetTrafficLightResponse\x12p\n\x0fSetTrafficLight\x12-.city.traffic_light.v2.SetTrafficLightRequest\x1a..city.traffic_light.v2.SetTrafficLightResponse\x12\x7f\n\x14SetTrafficLightPhase\x122.city.traffic_light.v2.SetTrafficLightPhaseRequest\x1a3.city.traffic_light.v2.SetTrafficLightPhaseResponse\x12\x82\x01\n\x15SetTrafficLightStatus\x123.city.traffic_light.v2.SetTrafficLightStatusRequest\x1a4.city.traffic_light.v2.SetTrafficLightStatusResponseB\xeb\x01\n\x19com.city.traffic_light.v2B\x18TrafficLightServiceProtoP\x01ZBgit.fiblab.net/sim/protos/go/city/traffic_light/v2;traffic_lightv2\xa2\x02\x03CTX\xaa\x02\x14City.TrafficLight.V2\xca\x02\x14City\\TrafficLight\\V2\xe2\x02 City\\TrafficLight\\V2\\GPBMetadata\xea\x02\x16City::TrafficLight::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.traffic_light.v2.traffic_light_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x19com.city.traffic_light.v2B\x18TrafficLightServiceProtoP\x01ZBgit.fiblab.net/sim/protos/go/city/traffic_light/v2;traffic_lightv2\xa2\x02\x03CTX\xaa\x02\x14City.TrafficLight.V2\xca\x02\x14City\\TrafficLight\\V2\xe2\x02 City\\TrafficLight\\V2\\GPBMetadata\xea\x02\x16City::TrafficLight::V2'
    _globals['_GETTRAFFICLIGHTREQUEST']._serialized_start = 119
    _globals['_GETTRAFFICLIGHTREQUEST']._serialized_end = 176
    _globals['_GETTRAFFICLIGHTRESPONSE']._serialized_start = 179
    _globals['_GETTRAFFICLIGHTRESPONSE']._serialized_end = 350
    _globals['_SETTRAFFICLIGHTREQUEST']._serialized_start = 353
    _globals['_SETTRAFFICLIGHTREQUEST']._serialized_end = 523
    _globals['_SETTRAFFICLIGHTRESPONSE']._serialized_start = 525
    _globals['_SETTRAFFICLIGHTRESPONSE']._serialized_end = 550
    _globals['_SETTRAFFICLIGHTPHASEREQUEST']._serialized_start = 553
    _globals['_SETTRAFFICLIGHTPHASEREQUEST']._serialized_end = 687
    _globals['_SETTRAFFICLIGHTPHASERESPONSE']._serialized_start = 689
    _globals['_SETTRAFFICLIGHTPHASERESPONSE']._serialized_end = 719
    _globals['_SETTRAFFICLIGHTSTATUSREQUEST']._serialized_start = 721
    _globals['_SETTRAFFICLIGHTSTATUSREQUEST']._serialized_end = 800
    _globals['_SETTRAFFICLIGHTSTATUSRESPONSE']._serialized_start = 802
    _globals['_SETTRAFFICLIGHTSTATUSRESPONSE']._serialized_end = 833
    _globals['_TRAFFICLIGHTSERVICE']._serialized_start = 836
    _globals['_TRAFFICLIGHTSERVICE']._serialized_end = 1347