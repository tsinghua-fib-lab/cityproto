"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.sync.v1 import sync_service_pb2 as city_dot_sync_dot_v1_dot_sync__service__pb2

class SyncServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SetURL = channel.unary_unary('/city.sync.v1.SyncService/SetURL', request_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.SetURLRequest.SerializeToString, response_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.SetURLResponse.FromString)
        self.GetURL = channel.unary_unary('/city.sync.v1.SyncService/GetURL', request_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.GetURLRequest.SerializeToString, response_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.GetURLResponse.FromString)
        self.Step = channel.unary_unary('/city.sync.v1.SyncService/Step', request_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.StepRequest.SerializeToString, response_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.StepResponse.FromString)

class SyncServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SetURL(self, request, context):
        """注册程序URL
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetURL(self, request, context):
        """获取程序URL
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Step(self, request, context):
        """步进
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_SyncServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'SetURL': grpc.unary_unary_rpc_method_handler(servicer.SetURL, request_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.SetURLRequest.FromString, response_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.SetURLResponse.SerializeToString), 'GetURL': grpc.unary_unary_rpc_method_handler(servicer.GetURL, request_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.GetURLRequest.FromString, response_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.GetURLResponse.SerializeToString), 'Step': grpc.unary_unary_rpc_method_handler(servicer.Step, request_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.StepRequest.FromString, response_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.StepResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.sync.v1.SyncService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class SyncService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SetURL(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.sync.v1.SyncService/SetURL', city_dot_sync_dot_v1_dot_sync__service__pb2.SetURLRequest.SerializeToString, city_dot_sync_dot_v1_dot_sync__service__pb2.SetURLResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetURL(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.sync.v1.SyncService/GetURL', city_dot_sync_dot_v1_dot_sync__service__pb2.GetURLRequest.SerializeToString, city_dot_sync_dot_v1_dot_sync__service__pb2.GetURLResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Step(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.sync.v1.SyncService/Step', city_dot_sync_dot_v1_dot_sync__service__pb2.StepRequest.SerializeToString, city_dot_sync_dot_v1_dot_sync__service__pb2.StepResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)