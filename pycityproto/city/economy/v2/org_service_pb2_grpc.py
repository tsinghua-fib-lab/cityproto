"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.economy.v2 import org_service_pb2 as city_dot_economy_dot_v2_dot_org__service__pb2

class OrgServiceStub(object):
    """OrgService 提供了经济系统中组织和代理的管理接口
    包括基本的CRUD操作、批量操作、增量更新和各种计算功能
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.AddOrg = channel.unary_unary('/city.economy.v2.OrgService/AddOrg', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgResponse.FromString)
        self.RemoveOrg = channel.unary_unary('/city.economy.v2.OrgService/RemoveOrg', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgResponse.FromString)
        self.GetOrg = channel.unary_unary('/city.economy.v2.OrgService/GetOrg', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgResponse.FromString)
        self.UpdateOrg = channel.unary_unary('/city.economy.v2.OrgService/UpdateOrg', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateOrgRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateOrgResponse.FromString)
        self.AddAgent = channel.unary_unary('/city.economy.v2.OrgService/AddAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentResponse.FromString)
        self.RemoveAgent = channel.unary_unary('/city.economy.v2.OrgService/RemoveAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentResponse.FromString)
        self.GetAgent = channel.unary_unary('/city.economy.v2.OrgService/GetAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentResponse.FromString)
        self.UpdateAgent = channel.unary_unary('/city.economy.v2.OrgService/UpdateAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentResponse.FromString)
        self.BatchGet = channel.unary_unary('/city.economy.v2.OrgService/BatchGet', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetResponse.FromString)
        self.BatchUpdate = channel.unary_unary('/city.economy.v2.OrgService/BatchUpdate', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateResponse.FromString)
        self.BatchSet = channel.unary_unary('/city.economy.v2.OrgService/BatchSet', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchSetRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchSetResponse.FromString)
        self.DeltaUpdateOrg = channel.unary_unary('/city.economy.v2.OrgService/DeltaUpdateOrg', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgResponse.FromString)
        self.DeltaUpdateAgent = channel.unary_unary('/city.economy.v2.OrgService/DeltaUpdateAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentResponse.FromString)
        self.BatchDeltaUpdate = channel.unary_unary('/city.economy.v2.OrgService/BatchDeltaUpdate', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateResponse.FromString)
        self.CalculateTaxesDue = channel.unary_unary('/city.economy.v2.OrgService/CalculateTaxesDue', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueResponse.FromString)
        self.CalculateConsumption = channel.unary_unary('/city.economy.v2.OrgService/CalculateConsumption', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionResponse.FromString)
        self.CalculateInterest = channel.unary_unary('/city.economy.v2.OrgService/CalculateInterest', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestResponse.FromString)
        self.CalculateRealGDP = channel.unary_unary('/city.economy.v2.OrgService/CalculateRealGDP', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPResponse.FromString)
        self.SaveEconomyEntities = channel.unary_unary('/city.economy.v2.OrgService/SaveEconomyEntities', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesResponse.FromString)
        self.LoadEconomyEntities = channel.unary_unary('/city.economy.v2.OrgService/LoadEconomyEntities', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesResponse.FromString)

class OrgServiceServicer(object):
    """OrgService 提供了经济系统中组织和代理的管理接口
    包括基本的CRUD操作、批量操作、增量更新和各种计算功能
    """

    def AddOrg(self, request, context):
        """AddOrg 添加一个新的组织到系统中
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveOrg(self, request, context):
        """RemoveOrg 从系统中移除指定的组织
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOrg(self, request, context):
        """GetOrg 获取指定组织的完整信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateOrg(self, request, context):
        """UpdateOrg 更新指定组织的信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddAgent(self, request, context):
        """AddAgent 添加一个新的代理到系统中
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveAgent(self, request, context):
        """RemoveAgent 从系统中移除指定的代理
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAgent(self, request, context):
        """GetAgent 获取指定代理的完整信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateAgent(self, request, context):
        """UpdateAgent 更新指定代理的信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchGet(self, request, context):
        """BatchGet 批量获取多个组织或代理的信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchUpdate(self, request, context):
        """BatchUpdate 批量更新多个组织或代理的信息，只更新请求中指定的字段
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchSet(self, request, context):
        """BatchSet 批量设置多个组织或代理的信息，完全替换所有字段
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeltaUpdateOrg(self, request, context):
        """DeltaUpdateOrg 对组织进行增量更新
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeltaUpdateAgent(self, request, context):
        """DeltaUpdateAgent 对代理进行增量更新
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchDeltaUpdate(self, request, context):
        """BatchDeltaUpdate 批量进行增量更新
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateTaxesDue(self, request, context):
        """CalculateTaxesDue 计算应缴税额并可选择进行再分配
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateConsumption(self, request, context):
        """CalculateConsumption 计算代理的消费情况
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateInterest(self, request, context):
        """CalculateInterest 计算银行利息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateRealGDP(self, request, context):
        """CalculateRealGDP 计算实际GDP
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SaveEconomyEntities(self, request, context):
        """SaveEconomyEntities 保存经济系统的当前状态
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def LoadEconomyEntities(self, request, context):
        """LoadEconomyEntities 加载经济系统的状态
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_OrgServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'AddOrg': grpc.unary_unary_rpc_method_handler(servicer.AddOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgResponse.SerializeToString), 'RemoveOrg': grpc.unary_unary_rpc_method_handler(servicer.RemoveOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgResponse.SerializeToString), 'GetOrg': grpc.unary_unary_rpc_method_handler(servicer.GetOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgResponse.SerializeToString), 'UpdateOrg': grpc.unary_unary_rpc_method_handler(servicer.UpdateOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateOrgResponse.SerializeToString), 'AddAgent': grpc.unary_unary_rpc_method_handler(servicer.AddAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentResponse.SerializeToString), 'RemoveAgent': grpc.unary_unary_rpc_method_handler(servicer.RemoveAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentResponse.SerializeToString), 'GetAgent': grpc.unary_unary_rpc_method_handler(servicer.GetAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentResponse.SerializeToString), 'UpdateAgent': grpc.unary_unary_rpc_method_handler(servicer.UpdateAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentResponse.SerializeToString), 'BatchGet': grpc.unary_unary_rpc_method_handler(servicer.BatchGet, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetResponse.SerializeToString), 'BatchUpdate': grpc.unary_unary_rpc_method_handler(servicer.BatchUpdate, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateResponse.SerializeToString), 'BatchSet': grpc.unary_unary_rpc_method_handler(servicer.BatchSet, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchSetRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchSetResponse.SerializeToString), 'DeltaUpdateOrg': grpc.unary_unary_rpc_method_handler(servicer.DeltaUpdateOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgResponse.SerializeToString), 'DeltaUpdateAgent': grpc.unary_unary_rpc_method_handler(servicer.DeltaUpdateAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentResponse.SerializeToString), 'BatchDeltaUpdate': grpc.unary_unary_rpc_method_handler(servicer.BatchDeltaUpdate, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateResponse.SerializeToString), 'CalculateTaxesDue': grpc.unary_unary_rpc_method_handler(servicer.CalculateTaxesDue, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueResponse.SerializeToString), 'CalculateConsumption': grpc.unary_unary_rpc_method_handler(servicer.CalculateConsumption, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionResponse.SerializeToString), 'CalculateInterest': grpc.unary_unary_rpc_method_handler(servicer.CalculateInterest, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestResponse.SerializeToString), 'CalculateRealGDP': grpc.unary_unary_rpc_method_handler(servicer.CalculateRealGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPResponse.SerializeToString), 'SaveEconomyEntities': grpc.unary_unary_rpc_method_handler(servicer.SaveEconomyEntities, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesResponse.SerializeToString), 'LoadEconomyEntities': grpc.unary_unary_rpc_method_handler(servicer.LoadEconomyEntities, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.economy.v2.OrgService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class OrgService(object):
    """OrgService 提供了经济系统中组织和代理的管理接口
    包括基本的CRUD操作、批量操作、增量更新和各种计算功能
    """

    @staticmethod
    def AddOrg(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddOrg', city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveOrg(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/RemoveOrg', city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetOrg(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetOrg', city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateOrg(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/UpdateOrg', city_dot_economy_dot_v2_dot_org__service__pb2.UpdateOrgRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.UpdateOrgResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddAgent(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddAgent', city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveAgent(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/RemoveAgent', city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAgent(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetAgent', city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateAgent(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/UpdateAgent', city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchGet(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/BatchGet', city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchUpdate(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/BatchUpdate', city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchSet(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/BatchSet', city_dot_economy_dot_v2_dot_org__service__pb2.BatchSetRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.BatchSetResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeltaUpdateOrg(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/DeltaUpdateOrg', city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeltaUpdateAgent(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/DeltaUpdateAgent', city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchDeltaUpdate(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/BatchDeltaUpdate', city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CalculateTaxesDue(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/CalculateTaxesDue', city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CalculateConsumption(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/CalculateConsumption', city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CalculateInterest(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/CalculateInterest', city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CalculateRealGDP(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/CalculateRealGDP', city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SaveEconomyEntities(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SaveEconomyEntities', city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def LoadEconomyEntities(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/LoadEconomyEntities', city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)