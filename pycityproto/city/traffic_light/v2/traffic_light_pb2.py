"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)city/traffic_light/v2/traffic_light.proto\x12\x15city.traffic_light.v2"^\n\x05Phase\x12\x1a\n\x08duration\x18\x01 \x01(\x01R\x08duration\x129\n\x06states\x18\x02 \x03(\x0e2!.city.traffic_light.v2.LightStateR\x06states"e\n\x0cTrafficLight\x12\x1f\n\x0bjunction_id\x18\x01 \x01(\x05R\njunctionId\x124\n\x06phases\x18\x02 \x03(\x0b2\x1c.city.traffic_light.v2.PhaseR\x06phases"[\n\rTrafficLights\x12J\n\x0etraffic_lights\x18\x01 \x03(\x0b2#.city.traffic_light.v2.TrafficLightR\rtrafficLights*m\n\nLightState\x12\x1b\n\x17LIGHT_STATE_UNSPECIFIED\x10\x00\x12\x13\n\x0fLIGHT_STATE_RED\x10\x01\x12\x15\n\x11LIGHT_STATE_GREEN\x10\x02\x12\x16\n\x12LIGHT_STATE_YELLOW\x10\x03B\xe4\x01\n\x19com.city.traffic_light.v2B\x11TrafficLightProtoP\x01ZBgit.fiblab.net/sim/protos/go/city/traffic_light/v2;traffic_lightv2\xa2\x02\x03CTX\xaa\x02\x14City.TrafficLight.V2\xca\x02\x14City\\TrafficLight\\V2\xe2\x02 City\\TrafficLight\\V2\\GPBMetadata\xea\x02\x16City::TrafficLight::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.traffic_light.v2.traffic_light_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x19com.city.traffic_light.v2B\x11TrafficLightProtoP\x01ZBgit.fiblab.net/sim/protos/go/city/traffic_light/v2;traffic_lightv2\xa2\x02\x03CTX\xaa\x02\x14City.TrafficLight.V2\xca\x02\x14City\\TrafficLight\\V2\xe2\x02 City\\TrafficLight\\V2\\GPBMetadata\xea\x02\x16City::TrafficLight::V2'
    _globals['_LIGHTSTATE']._serialized_start = 360
    _globals['_LIGHTSTATE']._serialized_end = 469
    _globals['_PHASE']._serialized_start = 68
    _globals['_PHASE']._serialized_end = 162
    _globals['_TRAFFICLIGHT']._serialized_start = 164
    _globals['_TRAFFICLIGHT']._serialized_end = 265
    _globals['_TRAFFICLIGHTS']._serialized_start = 267
    _globals['_TRAFFICLIGHTS']._serialized_end = 358