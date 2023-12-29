"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from .....city.comm.output.v1 import output_service_pb2 as city_dot_comm_dot_output_dot_v1_dot_output__service__pb2

class OutputServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Output = channel.unary_unary('/city.comm.output.v1.OutputService/Output', request_serializer=city_dot_comm_dot_output_dot_v1_dot_output__service__pb2.OutputRequest.SerializeToString, response_deserializer=city_dot_comm_dot_output_dot_v1_dot_output__service__pb2.OutputResponse.FromString)

class OutputServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Output(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_OutputServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'Output': grpc.unary_unary_rpc_method_handler(servicer.Output, request_deserializer=city_dot_comm_dot_output_dot_v1_dot_output__service__pb2.OutputRequest.FromString, response_serializer=city_dot_comm_dot_output_dot_v1_dot_output__service__pb2.OutputResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.comm.output.v1.OutputService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class OutputService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Output(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.comm.output.v1.OutputService/Output', city_dot_comm_dot_output_dot_v1_dot_output__service__pb2.OutputRequest.SerializeToString, city_dot_comm_dot_output_dot_v1_dot_output__service__pb2.OutputResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)