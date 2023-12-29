"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ......city.traffic.interaction.aoi.v1 import aoi_service_pb2 as city_dot_traffic_dot_interaction_dot_aoi_dot_v1_dot_aoi__service__pb2

class AoiServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetAoi = channel.unary_unary('/city.traffic.interaction.aoi.v1.AoiService/GetAoi', request_serializer=city_dot_traffic_dot_interaction_dot_aoi_dot_v1_dot_aoi__service__pb2.GetAoiRequest.SerializeToString, response_deserializer=city_dot_traffic_dot_interaction_dot_aoi_dot_v1_dot_aoi__service__pb2.GetAoiResponse.FromString)

class AoiServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetAoi(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_AoiServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetAoi': grpc.unary_unary_rpc_method_handler(servicer.GetAoi, request_deserializer=city_dot_traffic_dot_interaction_dot_aoi_dot_v1_dot_aoi__service__pb2.GetAoiRequest.FromString, response_serializer=city_dot_traffic_dot_interaction_dot_aoi_dot_v1_dot_aoi__service__pb2.GetAoiResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.traffic.interaction.aoi.v1.AoiService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class AoiService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetAoi(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.traffic.interaction.aoi.v1.AoiService/GetAoi', city_dot_traffic_dot_interaction_dot_aoi_dot_v1_dot_aoi__service__pb2.GetAoiRequest.SerializeToString, city_dot_traffic_dot_interaction_dot_aoi_dot_v1_dot_aoi__service__pb2.GetAoiResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)