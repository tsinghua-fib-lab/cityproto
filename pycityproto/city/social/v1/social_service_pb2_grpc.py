"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.social.v1 import social_service_pb2 as city_dot_social_dot_v1_dot_social__service__pb2

class SocialServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Send = channel.unary_unary('/city.social.v1.SocialService/Send', request_serializer=city_dot_social_dot_v1_dot_social__service__pb2.SendRequest.SerializeToString, response_deserializer=city_dot_social_dot_v1_dot_social__service__pb2.SendResponse.FromString)
        self.Receive = channel.unary_unary('/city.social.v1.SocialService/Receive', request_serializer=city_dot_social_dot_v1_dot_social__service__pb2.ReceiveRequest.SerializeToString, response_deserializer=city_dot_social_dot_v1_dot_social__service__pb2.ReceiveResponse.FromString)

class SocialServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Send(self, request, context):
        """发送消息
        Send message
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Receive(self, request, context):
        """接收消息，并清空该用户的消息队列
        Receive messages and clear the user's message queue
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_SocialServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'Send': grpc.unary_unary_rpc_method_handler(servicer.Send, request_deserializer=city_dot_social_dot_v1_dot_social__service__pb2.SendRequest.FromString, response_serializer=city_dot_social_dot_v1_dot_social__service__pb2.SendResponse.SerializeToString), 'Receive': grpc.unary_unary_rpc_method_handler(servicer.Receive, request_deserializer=city_dot_social_dot_v1_dot_social__service__pb2.ReceiveRequest.FromString, response_serializer=city_dot_social_dot_v1_dot_social__service__pb2.ReceiveResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.social.v1.SocialService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class SocialService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Send(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.social.v1.SocialService/Send', city_dot_social_dot_v1_dot_social__service__pb2.SendRequest.SerializeToString, city_dot_social_dot_v1_dot_social__service__pb2.SendResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Receive(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.social.v1.SocialService/Receive', city_dot_social_dot_v1_dot_social__service__pb2.ReceiveRequest.SerializeToString, city_dot_social_dot_v1_dot_social__service__pb2.ReceiveResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)