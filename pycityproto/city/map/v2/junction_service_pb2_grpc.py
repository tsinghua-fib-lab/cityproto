"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.map.v2 import junction_service_pb2 as city_dot_map_dot_v2_dot_junction__service__pb2

class JunctionServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetJunction = channel.unary_unary('/city.map.v2.JunctionService/GetJunction', request_serializer=city_dot_map_dot_v2_dot_junction__service__pb2.GetJunctionRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_junction__service__pb2.GetJunctionResponse.FromString)

class JunctionServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetJunction(self, request, context):
        """查询路口信息
        Get junction information
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_JunctionServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetJunction': grpc.unary_unary_rpc_method_handler(servicer.GetJunction, request_deserializer=city_dot_map_dot_v2_dot_junction__service__pb2.GetJunctionRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_junction__service__pb2.GetJunctionResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.map.v2.JunctionService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class JunctionService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetJunction(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.JunctionService/GetJunction', city_dot_map_dot_v2_dot_junction__service__pb2.GetJunctionRequest.SerializeToString, city_dot_map_dot_v2_dot_junction__service__pb2.GetJunctionResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)