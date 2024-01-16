"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.ping.v1 import ping_service_pb2 as city_dot_ping_dot_v1_dot_ping__service__pb2

class PingServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Ping = channel.unary_unary('/city.ping.v1.PingService/Ping', request_serializer=city_dot_ping_dot_v1_dot_ping__service__pb2.PingRequest.SerializeToString, response_deserializer=city_dot_ping_dot_v1_dot_ping__service__pb2.PingResponse.FromString)

class PingServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Ping(self, request, context):
        """连接测试
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_PingServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'Ping': grpc.unary_unary_rpc_method_handler(servicer.Ping, request_deserializer=city_dot_ping_dot_v1_dot_ping__service__pb2.PingRequest.FromString, response_serializer=city_dot_ping_dot_v1_dot_ping__service__pb2.PingResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.ping.v1.PingService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class PingService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Ping(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.ping.v1.PingService/Ping', city_dot_ping_dot_v1_dot_ping__service__pb2.PingRequest.SerializeToString, city_dot_ping_dot_v1_dot_ping__service__pb2.PingResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)