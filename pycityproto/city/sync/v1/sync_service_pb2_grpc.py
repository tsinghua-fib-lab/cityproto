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
        self.GetEtcd = channel.unary_unary('/city.sync.v1.SyncService/GetEtcd', request_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.GetEtcdRequest.SerializeToString, response_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.GetEtcdResponse.FromString)
        self.Step = channel.unary_unary('/city.sync.v1.SyncService/Step', request_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.StepRequest.SerializeToString, response_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.StepResponse.FromString)

class SyncServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetEtcd(self, request, context):
        """获取内嵌etcd的端口
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
    rpc_method_handlers = {'GetEtcd': grpc.unary_unary_rpc_method_handler(servicer.GetEtcd, request_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.GetEtcdRequest.FromString, response_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.GetEtcdResponse.SerializeToString), 'Step': grpc.unary_unary_rpc_method_handler(servicer.Step, request_deserializer=city_dot_sync_dot_v1_dot_sync__service__pb2.StepRequest.FromString, response_serializer=city_dot_sync_dot_v1_dot_sync__service__pb2.StepResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.sync.v1.SyncService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class SyncService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetEtcd(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.sync.v1.SyncService/GetEtcd', city_dot_sync_dot_v1_dot_sync__service__pb2.GetEtcdRequest.SerializeToString, city_dot_sync_dot_v1_dot_sync__service__pb2.GetEtcdResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Step(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.sync.v1.SyncService/Step', city_dot_sync_dot_v1_dot_sync__service__pb2.StepRequest.SerializeToString, city_dot_sync_dot_v1_dot_sync__service__pb2.StepResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)