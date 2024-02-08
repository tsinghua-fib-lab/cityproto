"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.clock.v1 import clock_service_pb2 as city_dot_clock_dot_v1_dot_clock__service__pb2

class ClockServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Now = channel.unary_unary('/city.clock.v1.ClockService/Now', request_serializer=city_dot_clock_dot_v1_dot_clock__service__pb2.NowRequest.SerializeToString, response_deserializer=city_dot_clock_dot_v1_dot_clock__service__pb2.NowResponse.FromString)

class ClockServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Now(self, request, context):
        """获取当前的模拟时间
        get current simulation clock
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_ClockServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'Now': grpc.unary_unary_rpc_method_handler(servicer.Now, request_deserializer=city_dot_clock_dot_v1_dot_clock__service__pb2.NowRequest.FromString, response_serializer=city_dot_clock_dot_v1_dot_clock__service__pb2.NowResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.clock.v1.ClockService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class ClockService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Now(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.clock.v1.ClockService/Now', city_dot_clock_dot_v1_dot_clock__service__pb2.NowRequest.SerializeToString, city_dot_clock_dot_v1_dot_clock__service__pb2.NowResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)