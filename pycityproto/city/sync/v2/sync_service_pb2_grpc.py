"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.sync.v2 import sync_service_pb2 as city_dot_sync_dot_v2_dot_sync__service__pb2

class SyncServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SetURL = channel.unary_unary('/city.sync.v2.SyncService/SetURL', request_serializer=city_dot_sync_dot_v2_dot_sync__service__pb2.SetURLRequest.SerializeToString, response_deserializer=city_dot_sync_dot_v2_dot_sync__service__pb2.SetURLResponse.FromString)
        self.GetURL = channel.unary_unary('/city.sync.v2.SyncService/GetURL', request_serializer=city_dot_sync_dot_v2_dot_sync__service__pb2.GetURLRequest.SerializeToString, response_deserializer=city_dot_sync_dot_v2_dot_sync__service__pb2.GetURLResponse.FromString)
        self.EnterStepSync = channel.unary_unary('/city.sync.v2.SyncService/EnterStepSync', request_serializer=city_dot_sync_dot_v2_dot_sync__service__pb2.EnterStepSyncRequest.SerializeToString, response_deserializer=city_dot_sync_dot_v2_dot_sync__service__pb2.EnterStepSyncResponse.FromString)
        self.ExitStepSync = channel.unary_unary('/city.sync.v2.SyncService/ExitStepSync', request_serializer=city_dot_sync_dot_v2_dot_sync__service__pb2.ExitStepSyncRequest.SerializeToString, response_deserializer=city_dot_sync_dot_v2_dot_sync__service__pb2.ExitStepSyncResponse.FromString)

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

    def EnterStepSync(self, request, context):
        """程序完成本步所有操作，进入同步状态。
        要求：进入同步状态的程序不再向其他程序发送消息，直到下一步开始。
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ExitStepSync(self, request, context):
        """程序完成同步阶段（无通信的安全区域）中必要的处理，如为prepare阶段加锁，可以进入准备阶段（恢复通信）。
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_SyncServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'SetURL': grpc.unary_unary_rpc_method_handler(servicer.SetURL, request_deserializer=city_dot_sync_dot_v2_dot_sync__service__pb2.SetURLRequest.FromString, response_serializer=city_dot_sync_dot_v2_dot_sync__service__pb2.SetURLResponse.SerializeToString), 'GetURL': grpc.unary_unary_rpc_method_handler(servicer.GetURL, request_deserializer=city_dot_sync_dot_v2_dot_sync__service__pb2.GetURLRequest.FromString, response_serializer=city_dot_sync_dot_v2_dot_sync__service__pb2.GetURLResponse.SerializeToString), 'EnterStepSync': grpc.unary_unary_rpc_method_handler(servicer.EnterStepSync, request_deserializer=city_dot_sync_dot_v2_dot_sync__service__pb2.EnterStepSyncRequest.FromString, response_serializer=city_dot_sync_dot_v2_dot_sync__service__pb2.EnterStepSyncResponse.SerializeToString), 'ExitStepSync': grpc.unary_unary_rpc_method_handler(servicer.ExitStepSync, request_deserializer=city_dot_sync_dot_v2_dot_sync__service__pb2.ExitStepSyncRequest.FromString, response_serializer=city_dot_sync_dot_v2_dot_sync__service__pb2.ExitStepSyncResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.sync.v2.SyncService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class SyncService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SetURL(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.sync.v2.SyncService/SetURL', city_dot_sync_dot_v2_dot_sync__service__pb2.SetURLRequest.SerializeToString, city_dot_sync_dot_v2_dot_sync__service__pb2.SetURLResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetURL(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.sync.v2.SyncService/GetURL', city_dot_sync_dot_v2_dot_sync__service__pb2.GetURLRequest.SerializeToString, city_dot_sync_dot_v2_dot_sync__service__pb2.GetURLResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def EnterStepSync(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.sync.v2.SyncService/EnterStepSync', city_dot_sync_dot_v2_dot_sync__service__pb2.EnterStepSyncRequest.SerializeToString, city_dot_sync_dot_v2_dot_sync__service__pb2.EnterStepSyncResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ExitStepSync(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.sync.v2.SyncService/ExitStepSync', city_dot_sync_dot_v2_dot_sync__service__pb2.ExitStepSyncRequest.SerializeToString, city_dot_sync_dot_v2_dot_sync__service__pb2.ExitStepSyncResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)