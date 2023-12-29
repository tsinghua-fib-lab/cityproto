"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from .....city.water.input.v1 import input_service_pb2 as city_dot_water_dot_input_dot_v1_dot_input__service__pb2

class InputServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Init = channel.unary_unary('/city.water.input.v1.InputService/Init', request_serializer=city_dot_water_dot_input_dot_v1_dot_input__service__pb2.InitRequest.SerializeToString, response_deserializer=city_dot_water_dot_input_dot_v1_dot_input__service__pb2.InitResponse.FromString)

class InputServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Init(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_InputServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'Init': grpc.unary_unary_rpc_method_handler(servicer.Init, request_deserializer=city_dot_water_dot_input_dot_v1_dot_input__service__pb2.InitRequest.FromString, response_serializer=city_dot_water_dot_input_dot_v1_dot_input__service__pb2.InitResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.water.input.v1.InputService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class InputService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Init(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.water.input.v1.InputService/Init', city_dot_water_dot_input_dot_v1_dot_input__service__pb2.InitRequest.SerializeToString, city_dot_water_dot_input_dot_v1_dot_input__service__pb2.InitResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)