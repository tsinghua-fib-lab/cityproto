"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ......city.comm.interaction.gateway.v1 import gateway_pb2 as city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__pb2
from ......city.event.v1 import event_pb2 as city_dot_event_dot_v1_dot_event__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n6city/comm/interaction/gateway/v1/gateway_service.proto\x12 city.comm.interaction.gateway.v1\x1a.city/comm/interaction/gateway/v1/gateway.proto\x1a\x19city/event/v1/event.proto"F\n\x1cSetGatewayPowerStatusRequest\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n\x06status\x18\x02 \x01(\x08R\x06status"\x1f\n\x1dSetGatewayPowerStatusResponse"E\n\x1bSetGatewayRuinStatusRequest\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n\x06status\x18\x02 \x01(\x08R\x06status"\x1e\n\x1cSetGatewayRuinStatusResponse"\x15\n\x13GetAllStatusRequest"]\n\x14GetAllStatusResponse\x12E\n\x08stations\x18\x01 \x03(\x0b2).city.comm.interaction.gateway.v1.StationR\x08stations"\x14\n\x12GetRuinInfoRequest"2\n\x08RuinInfo\x12\x10\n\x03num\x18\x01 \x01(\x05R\x03num\x12\x14\n\x05ratio\x18\x02 \x01(\x01R\x05ratio"\xd3\x01\n\x13GetRuinInfoResponse\x12<\n\x03one\x18\x01 \x01(\x0b2*.city.comm.interaction.gateway.v1.RuinInfoR\x03one\x12<\n\x03two\x18\x02 \x01(\x0b2*.city.comm.interaction.gateway.v1.RuinInfoR\x03two\x12@\n\x05three\x18\x03 \x01(\x0b2*.city.comm.interaction.gateway.v1.RuinInfoR\x05three"\x12\n\x10GetEventsRequest"B\n\x11GetEventsResponse\x12-\n\x06events\x18\x01 \x01(\x0b2\x15.city.event.v1.EventsR\x06events2\xb4\x05\n\x0eGatewayService\x12\x98\x01\n\x15SetGatewayPowerStatus\x12>.city.comm.interaction.gateway.v1.SetGatewayPowerStatusRequest\x1a?.city.comm.interaction.gateway.v1.SetGatewayPowerStatusResponse\x12\x95\x01\n\x14SetGatewayRuinStatus\x12=.city.comm.interaction.gateway.v1.SetGatewayRuinStatusRequest\x1a>.city.comm.interaction.gateway.v1.SetGatewayRuinStatusResponse\x12}\n\x0cGetAllStatus\x125.city.comm.interaction.gateway.v1.GetAllStatusRequest\x1a6.city.comm.interaction.gateway.v1.GetAllStatusResponse\x12z\n\x0bGetRuinInfo\x124.city.comm.interaction.gateway.v1.GetRuinInfoRequest\x1a5.city.comm.interaction.gateway.v1.GetRuinInfoResponse\x12t\n\tGetEvents\x122.city.comm.interaction.gateway.v1.GetEventsRequest\x1a3.city.comm.interaction.gateway.v1.GetEventsResponseB\xa9\x02\n$com.city.comm.interaction.gateway.v1B\x13GatewayServiceProtoP\x01ZGgit.fiblab.net/sim/protos/go/city/comm/interaction/gateway/v1;gatewayv1\xa2\x02\x04CCIG\xaa\x02 City.Comm.Interaction.Gateway.V1\xca\x02 City\\Comm\\Interaction\\Gateway\\V1\xe2\x02,City\\Comm\\Interaction\\Gateway\\V1\\GPBMetadata\xea\x02$City::Comm::Interaction::Gateway::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.comm.interaction.gateway.v1.gateway_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n$com.city.comm.interaction.gateway.v1B\x13GatewayServiceProtoP\x01ZGgit.fiblab.net/sim/protos/go/city/comm/interaction/gateway/v1;gatewayv1\xa2\x02\x04CCIG\xaa\x02 City.Comm.Interaction.Gateway.V1\xca\x02 City\\Comm\\Interaction\\Gateway\\V1\xe2\x02,City\\Comm\\Interaction\\Gateway\\V1\\GPBMetadata\xea\x02$City::Comm::Interaction::Gateway::V1'
    _globals['_SETGATEWAYPOWERSTATUSREQUEST']._serialized_start = 167
    _globals['_SETGATEWAYPOWERSTATUSREQUEST']._serialized_end = 237
    _globals['_SETGATEWAYPOWERSTATUSRESPONSE']._serialized_start = 239
    _globals['_SETGATEWAYPOWERSTATUSRESPONSE']._serialized_end = 270
    _globals['_SETGATEWAYRUINSTATUSREQUEST']._serialized_start = 272
    _globals['_SETGATEWAYRUINSTATUSREQUEST']._serialized_end = 341
    _globals['_SETGATEWAYRUINSTATUSRESPONSE']._serialized_start = 343
    _globals['_SETGATEWAYRUINSTATUSRESPONSE']._serialized_end = 373
    _globals['_GETALLSTATUSREQUEST']._serialized_start = 375
    _globals['_GETALLSTATUSREQUEST']._serialized_end = 396
    _globals['_GETALLSTATUSRESPONSE']._serialized_start = 398
    _globals['_GETALLSTATUSRESPONSE']._serialized_end = 491
    _globals['_GETRUININFOREQUEST']._serialized_start = 493
    _globals['_GETRUININFOREQUEST']._serialized_end = 513
    _globals['_RUININFO']._serialized_start = 515
    _globals['_RUININFO']._serialized_end = 565
    _globals['_GETRUININFORESPONSE']._serialized_start = 568
    _globals['_GETRUININFORESPONSE']._serialized_end = 779
    _globals['_GETEVENTSREQUEST']._serialized_start = 781
    _globals['_GETEVENTSREQUEST']._serialized_end = 799
    _globals['_GETEVENTSRESPONSE']._serialized_start = 801
    _globals['_GETEVENTSRESPONSE']._serialized_end = 867
    _globals['_GATEWAYSERVICE']._serialized_start = 870
    _globals['_GATEWAYSERVICE']._serialized_end = 1562