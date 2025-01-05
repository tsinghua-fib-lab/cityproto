"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.pause.v1 import pause_service_pb2 as city_dot_pause_dot_v1_dot_pause__service__pb2

class PauseServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Pause = channel.unary_unary('/city.pause.v1.PauseService/Pause', request_serializer=city_dot_pause_dot_v1_dot_pause__service__pb2.PauseRequest.SerializeToString, response_deserializer=city_dot_pause_dot_v1_dot_pause__service__pb2.PauseResponse.FromString)
        self.Resume = channel.unary_unary('/city.pause.v1.PauseService/Resume', request_serializer=city_dot_pause_dot_v1_dot_pause__service__pb2.ResumeRequest.SerializeToString, response_deserializer=city_dot_pause_dot_v1_dot_pause__service__pb2.ResumeResponse.FromString)

class PauseServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Pause(self, request, context):
        """暂停程序
        Pause the program
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Resume(self, request, context):
        """恢复程序
        Resume the program
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_PauseServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'Pause': grpc.unary_unary_rpc_method_handler(servicer.Pause, request_deserializer=city_dot_pause_dot_v1_dot_pause__service__pb2.PauseRequest.FromString, response_serializer=city_dot_pause_dot_v1_dot_pause__service__pb2.PauseResponse.SerializeToString), 'Resume': grpc.unary_unary_rpc_method_handler(servicer.Resume, request_deserializer=city_dot_pause_dot_v1_dot_pause__service__pb2.ResumeRequest.FromString, response_serializer=city_dot_pause_dot_v1_dot_pause__service__pb2.ResumeResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.pause.v1.PauseService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class PauseService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Pause(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.pause.v1.PauseService/Pause', city_dot_pause_dot_v1_dot_pause__service__pb2.PauseRequest.SerializeToString, city_dot_pause_dot_v1_dot_pause__service__pb2.PauseResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Resume(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.pause.v1.PauseService/Resume', city_dot_pause_dot_v1_dot_pause__service__pb2.ResumeRequest.SerializeToString, city_dot_pause_dot_v1_dot_pause__service__pb2.ResumeResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)