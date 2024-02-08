"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.routing.v2 import routing_service_pb2 as city_dot_routing_dot_v2_dot_routing__service__pb2

class RoutingServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetRoute = channel.unary_unary('/city.routing.v2.RoutingService/GetRoute', request_serializer=city_dot_routing_dot_v2_dot_routing__service__pb2.GetRouteRequest.SerializeToString, response_deserializer=city_dot_routing_dot_v2_dot_routing__service__pb2.GetRouteResponse.FromString)
        self.SetDrivingCosts = channel.unary_unary('/city.routing.v2.RoutingService/SetDrivingCosts', request_serializer=city_dot_routing_dot_v2_dot_routing__service__pb2.SetDrivingCostsRequest.SerializeToString, response_deserializer=city_dot_routing_dot_v2_dot_routing__service__pb2.SetDrivingCostsResponse.FromString)
        self.GetDrivingCosts = channel.unary_unary('/city.routing.v2.RoutingService/GetDrivingCosts', request_serializer=city_dot_routing_dot_v2_dot_routing__service__pb2.GetDrivingCostsRequest.SerializeToString, response_deserializer=city_dot_routing_dot_v2_dot_routing__service__pb2.GetDrivingCostsResponse.FromString)

class RoutingServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetRoute(self, request, context):
        """获取导航路线
        Get routing path
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetDrivingCosts(self, request, context):
        """设置行车导航道路通行成本
        Set traveling cost of driving routing
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetDrivingCosts(self, request, context):
        """获取行车导航道路通行成本
        Get traveling cost of driving routing
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_RoutingServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetRoute': grpc.unary_unary_rpc_method_handler(servicer.GetRoute, request_deserializer=city_dot_routing_dot_v2_dot_routing__service__pb2.GetRouteRequest.FromString, response_serializer=city_dot_routing_dot_v2_dot_routing__service__pb2.GetRouteResponse.SerializeToString), 'SetDrivingCosts': grpc.unary_unary_rpc_method_handler(servicer.SetDrivingCosts, request_deserializer=city_dot_routing_dot_v2_dot_routing__service__pb2.SetDrivingCostsRequest.FromString, response_serializer=city_dot_routing_dot_v2_dot_routing__service__pb2.SetDrivingCostsResponse.SerializeToString), 'GetDrivingCosts': grpc.unary_unary_rpc_method_handler(servicer.GetDrivingCosts, request_deserializer=city_dot_routing_dot_v2_dot_routing__service__pb2.GetDrivingCostsRequest.FromString, response_serializer=city_dot_routing_dot_v2_dot_routing__service__pb2.GetDrivingCostsResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.routing.v2.RoutingService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class RoutingService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetRoute(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.routing.v2.RoutingService/GetRoute', city_dot_routing_dot_v2_dot_routing__service__pb2.GetRouteRequest.SerializeToString, city_dot_routing_dot_v2_dot_routing__service__pb2.GetRouteResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetDrivingCosts(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.routing.v2.RoutingService/SetDrivingCosts', city_dot_routing_dot_v2_dot_routing__service__pb2.SetDrivingCostsRequest.SerializeToString, city_dot_routing_dot_v2_dot_routing__service__pb2.SetDrivingCostsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetDrivingCosts(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.routing.v2.RoutingService/GetDrivingCosts', city_dot_routing_dot_v2_dot_routing__service__pb2.GetDrivingCostsRequest.SerializeToString, city_dot_routing_dot_v2_dot_routing__service__pb2.GetDrivingCostsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)