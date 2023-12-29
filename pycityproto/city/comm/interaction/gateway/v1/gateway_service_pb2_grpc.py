"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ......city.comm.interaction.gateway.v1 import gateway_service_pb2 as city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2

class GatewayServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SetGatewayPowerStatus = channel.unary_unary('/city.comm.interaction.gateway.v1.GatewayService/SetGatewayPowerStatus', request_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayPowerStatusRequest.SerializeToString, response_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayPowerStatusResponse.FromString)
        self.SetGatewayRuinStatus = channel.unary_unary('/city.comm.interaction.gateway.v1.GatewayService/SetGatewayRuinStatus', request_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayRuinStatusRequest.SerializeToString, response_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayRuinStatusResponse.FromString)
        self.GetAllStatus = channel.unary_unary('/city.comm.interaction.gateway.v1.GatewayService/GetAllStatus', request_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetAllStatusRequest.SerializeToString, response_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetAllStatusResponse.FromString)
        self.GetRuinInfo = channel.unary_unary('/city.comm.interaction.gateway.v1.GatewayService/GetRuinInfo', request_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetRuinInfoRequest.SerializeToString, response_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetRuinInfoResponse.FromString)
        self.GetEvents = channel.unary_unary('/city.comm.interaction.gateway.v1.GatewayService/GetEvents', request_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetEventsRequest.SerializeToString, response_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetEventsResponse.FromString)

class GatewayServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SetGatewayPowerStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetGatewayRuinStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRuinInfo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetEvents(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_GatewayServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'SetGatewayPowerStatus': grpc.unary_unary_rpc_method_handler(servicer.SetGatewayPowerStatus, request_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayPowerStatusRequest.FromString, response_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayPowerStatusResponse.SerializeToString), 'SetGatewayRuinStatus': grpc.unary_unary_rpc_method_handler(servicer.SetGatewayRuinStatus, request_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayRuinStatusRequest.FromString, response_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayRuinStatusResponse.SerializeToString), 'GetAllStatus': grpc.unary_unary_rpc_method_handler(servicer.GetAllStatus, request_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetAllStatusRequest.FromString, response_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetAllStatusResponse.SerializeToString), 'GetRuinInfo': grpc.unary_unary_rpc_method_handler(servicer.GetRuinInfo, request_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetRuinInfoRequest.FromString, response_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetRuinInfoResponse.SerializeToString), 'GetEvents': grpc.unary_unary_rpc_method_handler(servicer.GetEvents, request_deserializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetEventsRequest.FromString, response_serializer=city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetEventsResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.comm.interaction.gateway.v1.GatewayService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class GatewayService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SetGatewayPowerStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.comm.interaction.gateway.v1.GatewayService/SetGatewayPowerStatus', city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayPowerStatusRequest.SerializeToString, city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayPowerStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetGatewayRuinStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.comm.interaction.gateway.v1.GatewayService/SetGatewayRuinStatus', city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayRuinStatusRequest.SerializeToString, city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.SetGatewayRuinStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAllStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.comm.interaction.gateway.v1.GatewayService/GetAllStatus', city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetAllStatusRequest.SerializeToString, city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetAllStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetRuinInfo(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.comm.interaction.gateway.v1.GatewayService/GetRuinInfo', city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetRuinInfoRequest.SerializeToString, city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetRuinInfoResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetEvents(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.comm.interaction.gateway.v1.GatewayService/GetEvents', city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetEventsRequest.SerializeToString, city_dot_comm_dot_interaction_dot_gateway_dot_v1_dot_gateway__service__pb2.GetEventsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)