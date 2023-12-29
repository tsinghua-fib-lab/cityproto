"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.event.v1 import event_service_pb2 as city_dot_event_dot_v1_dot_event__service__pb2

class EventServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Publish = channel.unary_unary('/city.event.v1.EventService/Publish', request_serializer=city_dot_event_dot_v1_dot_event__service__pb2.PublishRequest.SerializeToString, response_deserializer=city_dot_event_dot_v1_dot_event__service__pb2.PublishResponse.FromString)
        self.Pull = channel.unary_unary('/city.event.v1.EventService/Pull', request_serializer=city_dot_event_dot_v1_dot_event__service__pb2.PullRequest.SerializeToString, response_deserializer=city_dot_event_dot_v1_dot_event__service__pb2.PullResponse.FromString)

class EventServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Publish(self, request, context):
        """发布事件
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Pull(self, request, context):
        """从事件中心拉取事件
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_EventServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'Publish': grpc.unary_unary_rpc_method_handler(servicer.Publish, request_deserializer=city_dot_event_dot_v1_dot_event__service__pb2.PublishRequest.FromString, response_serializer=city_dot_event_dot_v1_dot_event__service__pb2.PublishResponse.SerializeToString), 'Pull': grpc.unary_unary_rpc_method_handler(servicer.Pull, request_deserializer=city_dot_event_dot_v1_dot_event__service__pb2.PullRequest.FromString, response_serializer=city_dot_event_dot_v1_dot_event__service__pb2.PullResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.event.v1.EventService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class EventService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Publish(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.event.v1.EventService/Publish', city_dot_event_dot_v1_dot_event__service__pb2.PublishRequest.SerializeToString, city_dot_event_dot_v1_dot_event__service__pb2.PublishResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Pull(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.event.v1.EventService/Pull', city_dot_event_dot_v1_dot_event__service__pb2.PullRequest.SerializeToString, city_dot_event_dot_v1_dot_event__service__pb2.PullResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)