"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.event.v2 import event_service_pb2 as city_dot_event_dot_v2_dot_event__service__pb2

class EventServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetEventsByTopic = channel.unary_unary('/city.event.v2.EventService/GetEventsByTopic', request_serializer=city_dot_event_dot_v2_dot_event__service__pb2.GetEventsByTopicRequest.SerializeToString, response_deserializer=city_dot_event_dot_v2_dot_event__service__pb2.GetEventsByTopicResponse.FromString)
        self.ResolveEvents = channel.unary_unary('/city.event.v2.EventService/ResolveEvents', request_serializer=city_dot_event_dot_v2_dot_event__service__pb2.ResolveEventsRequest.SerializeToString, response_deserializer=city_dot_event_dot_v2_dot_event__service__pb2.ResolveEventsResponse.FromString)

class EventServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetEventsByTopic(self, request, context):
        """按照topic查询事件
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ResolveEvents(self, request, context):
        """确认事件已被处理
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_EventServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetEventsByTopic': grpc.unary_unary_rpc_method_handler(servicer.GetEventsByTopic, request_deserializer=city_dot_event_dot_v2_dot_event__service__pb2.GetEventsByTopicRequest.FromString, response_serializer=city_dot_event_dot_v2_dot_event__service__pb2.GetEventsByTopicResponse.SerializeToString), 'ResolveEvents': grpc.unary_unary_rpc_method_handler(servicer.ResolveEvents, request_deserializer=city_dot_event_dot_v2_dot_event__service__pb2.ResolveEventsRequest.FromString, response_serializer=city_dot_event_dot_v2_dot_event__service__pb2.ResolveEventsResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.event.v2.EventService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class EventService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetEventsByTopic(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.event.v2.EventService/GetEventsByTopic', city_dot_event_dot_v2_dot_event__service__pb2.GetEventsByTopicRequest.SerializeToString, city_dot_event_dot_v2_dot_event__service__pb2.GetEventsByTopicResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ResolveEvents(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.event.v2.EventService/ResolveEvents', city_dot_event_dot_v2_dot_event__service__pb2.ResolveEventsRequest.SerializeToString, city_dot_event_dot_v2_dot_event__service__pb2.ResolveEventsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)