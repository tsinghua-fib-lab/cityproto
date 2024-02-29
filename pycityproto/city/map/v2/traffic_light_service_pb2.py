"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.map.v2 import light_pb2 as city_dot_map_dot_v2_dot_light__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'city/map/v2/traffic_light_service.proto\x12\x0bcity.map.v2\x1a\x17city/map/v2/light.proto"9\n\x16GetTrafficLightRequest\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId"\xa1\x01\n\x17GetTrafficLightResponse\x12>\n\rtraffic_light\x18\x01 \x01(\x0b2\x19.city.map.v2.TrafficLightR\x0ctrafficLight\x12\x1f\n\x0bphase_index\x18\x02 \x01(\x05R\nphaseIndex\x12%\n\x0etime_remaining\x18\x03 \x01(\x01R\rtimeRemaining"\xa0\x01\n\x16SetTrafficLightRequest\x12>\n\rtraffic_light\x18\x01 \x01(\x0b2\x19.city.map.v2.TrafficLightR\x0ctrafficLight\x12\x1f\n\x0bphase_index\x18\x02 \x01(\x05R\nphaseIndex\x12%\n\x0etime_remaining\x18\x03 \x01(\x01R\rtimeRemaining"\x19\n\x17SetTrafficLightResponse"\x86\x01\n\x1bSetTrafficLightPhaseRequest\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId\x12\x1f\n\x0bphase_index\x18\x02 \x01(\x05R\nphaseIndex\x12%\n\x0etime_remaining\x18\x03 \x01(\x01R\rtimeRemaining"\x1e\n\x1cSetTrafficLightPhaseResponse"O\n\x1cSetTrafficLightStatusRequest\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId\x12\x0e\n\x02ok\x18\x02 \x01(\x08R\x02ok"\x1f\n\x1dSetTrafficLightStatusResponse2\xae\x03\n\x13TrafficLightService\x12\\\n\x0fGetTrafficLight\x12#.city.map.v2.GetTrafficLightRequest\x1a$.city.map.v2.GetTrafficLightResponse\x12\\\n\x0fSetTrafficLight\x12#.city.map.v2.SetTrafficLightRequest\x1a$.city.map.v2.SetTrafficLightResponse\x12k\n\x14SetTrafficLightPhase\x12(.city.map.v2.SetTrafficLightPhaseRequest\x1a).city.map.v2.SetTrafficLightPhaseResponse\x12n\n\x15SetTrafficLightStatus\x12).city.map.v2.SetTrafficLightStatusRequest\x1a*.city.map.v2.SetTrafficLightStatusResponseB\xa9\x01\n\x0fcom.city.map.v2B\x18TrafficLightServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.traffic_light_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\x18TrafficLightServiceProtoP\x01Z.git.fiblab.net/sim/protos/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_GETTRAFFICLIGHTREQUEST']._serialized_start = 81
    _globals['_GETTRAFFICLIGHTREQUEST']._serialized_end = 138
    _globals['_GETTRAFFICLIGHTRESPONSE']._serialized_start = 141
    _globals['_GETTRAFFICLIGHTRESPONSE']._serialized_end = 302
    _globals['_SETTRAFFICLIGHTREQUEST']._serialized_start = 305
    _globals['_SETTRAFFICLIGHTREQUEST']._serialized_end = 465
    _globals['_SETTRAFFICLIGHTRESPONSE']._serialized_start = 467
    _globals['_SETTRAFFICLIGHTRESPONSE']._serialized_end = 492
    _globals['_SETTRAFFICLIGHTPHASEREQUEST']._serialized_start = 495
    _globals['_SETTRAFFICLIGHTPHASEREQUEST']._serialized_end = 629
    _globals['_SETTRAFFICLIGHTPHASERESPONSE']._serialized_start = 631
    _globals['_SETTRAFFICLIGHTPHASERESPONSE']._serialized_end = 661
    _globals['_SETTRAFFICLIGHTSTATUSREQUEST']._serialized_start = 663
    _globals['_SETTRAFFICLIGHTSTATUSREQUEST']._serialized_end = 742
    _globals['_SETTRAFFICLIGHTSTATUSRESPONSE']._serialized_start = 744
    _globals['_SETTRAFFICLIGHTSTATUSRESPONSE']._serialized_end = 775
    _globals['_TRAFFICLIGHTSERVICE']._serialized_start = 778
    _globals['_TRAFFICLIGHTSERVICE']._serialized_end = 1208