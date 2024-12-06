"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.agentcomm.v1 import agent_service_pb2 as city_dot_agentcomm_dot_v1_dot_agent__service__pb2

class AgentServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Communicate = channel.stream_stream('/city.agentcomm.v1.AgentService/Communicate', request_serializer=city_dot_agentcomm_dot_v1_dot_agent__service__pb2.CommunicateRequest.SerializeToString, response_deserializer=city_dot_agentcomm_dot_v1_dot_agent__service__pb2.CommunicateResponse.FromString)

class AgentServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Communicate(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_AgentServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'Communicate': grpc.stream_stream_rpc_method_handler(servicer.Communicate, request_deserializer=city_dot_agentcomm_dot_v1_dot_agent__service__pb2.CommunicateRequest.FromString, response_serializer=city_dot_agentcomm_dot_v1_dot_agent__service__pb2.CommunicateResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.agentcomm.v1.AgentService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class AgentService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Communicate(request_iterator, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/city.agentcomm.v1.AgentService/Communicate', city_dot_agentcomm_dot_v1_dot_agent__service__pb2.CommunicateRequest.SerializeToString, city_dot_agentcomm_dot_v1_dot_agent__service__pb2.CommunicateResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)