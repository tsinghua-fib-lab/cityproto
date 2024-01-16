"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.routing.v2 import cost_pb2 as city_dot_routing_dot_v2_dot_cost__pb2
from ....city.routing.v2 import routing_pb2 as city_dot_routing_dot_v2_dot_routing__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%city/routing/v2/routing_service.proto\x12\x0fcity.routing.v2\x1a\x15city/geo/v2/geo.proto\x1a\x1acity/routing/v2/cost.proto\x1a\x1dcity/routing/v2/routing.proto"\xab\x01\n\x0fGetRouteRequest\x12.\n\x04type\x18\x01 \x01(\x0e2\x1a.city.routing.v2.RouteTypeR\x04type\x12+\n\x05start\x18\x02 \x01(\x0b2\x15.city.geo.v2.PositionR\x05start\x12\'\n\x03end\x18\x03 \x01(\x0b2\x15.city.geo.v2.PositionR\x03end\x12\x12\n\x04time\x18\x05 \x01(\x01R\x04time"H\n\x10GetRouteResponse\x124\n\x08journeys\x18\x01 \x03(\x0b2\x18.city.routing.v2.JourneyR\x08journeys"E\n\x16SetDrivingCostsRequest\x12+\n\x05costs\x18\x01 \x03(\x0b2\x15.city.routing.v2.CostR\x05costs"\x19\n\x17SetDrivingCostsResponse"E\n\x16GetDrivingCostsRequest\x12+\n\x05costs\x18\x01 \x03(\x0b2\x15.city.routing.v2.CostR\x05costs"F\n\x17GetDrivingCostsResponse\x12+\n\x05costs\x18\x01 \x03(\x0b2\x15.city.routing.v2.CostR\x05costs2\xad\x02\n\x0eRoutingService\x12O\n\x08GetRoute\x12 .city.routing.v2.GetRouteRequest\x1a!.city.routing.v2.GetRouteResponse\x12d\n\x0fSetDrivingCosts\x12\'.city.routing.v2.SetDrivingCostsRequest\x1a(.city.routing.v2.SetDrivingCostsResponse\x12d\n\x0fGetDrivingCosts\x12\'.city.routing.v2.GetDrivingCostsRequest\x1a(.city.routing.v2.GetDrivingCostsResponseB\xc0\x01\n\x13com.city.routing.v2B\x13RoutingServiceProtoP\x01Z6git.fiblab.net/sim/protos/go/city/routing/v2;routingv2\xa2\x02\x03CRX\xaa\x02\x0fCity.Routing.V2\xca\x02\x0fCity\\Routing\\V2\xe2\x02\x1bCity\\Routing\\V2\\GPBMetadata\xea\x02\x11City::Routing::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.routing.v2.routing_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.routing.v2B\x13RoutingServiceProtoP\x01Z6git.fiblab.net/sim/protos/go/city/routing/v2;routingv2\xa2\x02\x03CRX\xaa\x02\x0fCity.Routing.V2\xca\x02\x0fCity\\Routing\\V2\xe2\x02\x1bCity\\Routing\\V2\\GPBMetadata\xea\x02\x11City::Routing::V2'
    _globals['_GETROUTEREQUEST']._serialized_start = 141
    _globals['_GETROUTEREQUEST']._serialized_end = 312
    _globals['_GETROUTERESPONSE']._serialized_start = 314
    _globals['_GETROUTERESPONSE']._serialized_end = 386
    _globals['_SETDRIVINGCOSTSREQUEST']._serialized_start = 388
    _globals['_SETDRIVINGCOSTSREQUEST']._serialized_end = 457
    _globals['_SETDRIVINGCOSTSRESPONSE']._serialized_start = 459
    _globals['_SETDRIVINGCOSTSRESPONSE']._serialized_end = 484
    _globals['_GETDRIVINGCOSTSREQUEST']._serialized_start = 486
    _globals['_GETDRIVINGCOSTSREQUEST']._serialized_end = 555
    _globals['_GETDRIVINGCOSTSRESPONSE']._serialized_start = 557
    _globals['_GETDRIVINGCOSTSRESPONSE']._serialized_end = 627
    _globals['_ROUTINGSERVICE']._serialized_start = 630
    _globals['_ROUTINGSERVICE']._serialized_end = 931