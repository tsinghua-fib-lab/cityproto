"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.agent.v2 import agent_service_pb2 as city_dot_agent_dot_v2_dot_agent__service__pb2

class AgentServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetAgent = channel.unary_unary('/city.agent.v2.AgentService/GetAgent', request_serializer=city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentRequest.SerializeToString, response_deserializer=city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentResponse.FromString)
        self.AddAgent = channel.unary_unary('/city.agent.v2.AgentService/AddAgent', request_serializer=city_dot_agent_dot_v2_dot_agent__service__pb2.AddAgentRequest.SerializeToString, response_deserializer=city_dot_agent_dot_v2_dot_agent__service__pb2.AddAgentResponse.FromString)
        self.SetSchedule = channel.unary_unary('/city.agent.v2.AgentService/SetSchedule', request_serializer=city_dot_agent_dot_v2_dot_agent__service__pb2.SetScheduleRequest.SerializeToString, response_deserializer=city_dot_agent_dot_v2_dot_agent__service__pb2.SetScheduleResponse.FromString)
        self.GetAgentsByLongLatArea = channel.unary_unary('/city.agent.v2.AgentService/GetAgentsByLongLatArea', request_serializer=city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentsByLongLatAreaRequest.SerializeToString, response_deserializer=city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentsByLongLatAreaResponse.FromString)

class AgentServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetAgent(self, request, context):
        """获取agent信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddAgent(self, request, context):
        """新增agent 传入agent初始位置、目的地表、属性 返回agentid
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetSchedule(self, request, context):
        """修改agent的schedule 传入agentid、目的地表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAgentsByLongLatArea(self, request, context):
        """获取特定区域内的agent
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_AgentServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetAgent': grpc.unary_unary_rpc_method_handler(servicer.GetAgent, request_deserializer=city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentRequest.FromString, response_serializer=city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentResponse.SerializeToString), 'AddAgent': grpc.unary_unary_rpc_method_handler(servicer.AddAgent, request_deserializer=city_dot_agent_dot_v2_dot_agent__service__pb2.AddAgentRequest.FromString, response_serializer=city_dot_agent_dot_v2_dot_agent__service__pb2.AddAgentResponse.SerializeToString), 'SetSchedule': grpc.unary_unary_rpc_method_handler(servicer.SetSchedule, request_deserializer=city_dot_agent_dot_v2_dot_agent__service__pb2.SetScheduleRequest.FromString, response_serializer=city_dot_agent_dot_v2_dot_agent__service__pb2.SetScheduleResponse.SerializeToString), 'GetAgentsByLongLatArea': grpc.unary_unary_rpc_method_handler(servicer.GetAgentsByLongLatArea, request_deserializer=city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentsByLongLatAreaRequest.FromString, response_serializer=city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentsByLongLatAreaResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.agent.v2.AgentService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class AgentService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetAgent(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.agent.v2.AgentService/GetAgent', city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentRequest.SerializeToString, city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddAgent(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.agent.v2.AgentService/AddAgent', city_dot_agent_dot_v2_dot_agent__service__pb2.AddAgentRequest.SerializeToString, city_dot_agent_dot_v2_dot_agent__service__pb2.AddAgentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetSchedule(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.agent.v2.AgentService/SetSchedule', city_dot_agent_dot_v2_dot_agent__service__pb2.SetScheduleRequest.SerializeToString, city_dot_agent_dot_v2_dot_agent__service__pb2.SetScheduleResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAgentsByLongLatArea(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.agent.v2.AgentService/GetAgentsByLongLatArea', city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentsByLongLatAreaRequest.SerializeToString, city_dot_agent_dot_v2_dot_agent__service__pb2.GetAgentsByLongLatAreaResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)