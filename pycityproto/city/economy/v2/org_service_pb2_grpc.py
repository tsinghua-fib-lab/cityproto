"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.economy.v2 import org_service_pb2 as city_dot_economy_dot_v2_dot_org__service__pb2

class OrgServiceStub(object):
    """OrgService provides management interfaces for organizations and agents in the economic system
    All basic operations support batch processing by default

    OrgService 提供了经济系统中组织和代理的管理接口
    所有基础操作默认支持批量处理
    @section Firm Operations
    @section 企业相关操作
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.AddFirm = channel.unary_unary('/city.economy.v2.OrgService/AddFirm', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddFirmRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddFirmResponse.FromString)
        self.RemoveFirm = channel.unary_unary('/city.economy.v2.OrgService/RemoveFirm', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveFirmRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveFirmResponse.FromString)
        self.GetFirm = channel.unary_unary('/city.economy.v2.OrgService/GetFirm', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetFirmRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetFirmResponse.FromString)
        self.UpdateFirm = channel.unary_unary('/city.economy.v2.OrgService/UpdateFirm', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateFirmRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateFirmResponse.FromString)
        self.ListFirms = channel.unary_unary('/city.economy.v2.OrgService/ListFirms', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListFirmsRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListFirmsResponse.FromString)
        self.DeltaUpdateFirm = channel.unary_unary('/city.economy.v2.OrgService/DeltaUpdateFirm', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateFirmRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateFirmResponse.FromString)
        self.AddNBS = channel.unary_unary('/city.economy.v2.OrgService/AddNBS', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddNBSRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddNBSResponse.FromString)
        self.RemoveNBS = channel.unary_unary('/city.economy.v2.OrgService/RemoveNBS', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveNBSRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveNBSResponse.FromString)
        self.GetNBS = channel.unary_unary('/city.economy.v2.OrgService/GetNBS', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNBSRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNBSResponse.FromString)
        self.UpdateNBS = channel.unary_unary('/city.economy.v2.OrgService/UpdateNBS', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateNBSRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateNBSResponse.FromString)
        self.ListNBS = channel.unary_unary('/city.economy.v2.OrgService/ListNBS', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListNBSRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListNBSResponse.FromString)
        self.DeltaUpdateNBS = channel.unary_unary('/city.economy.v2.OrgService/DeltaUpdateNBS', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateNBSRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateNBSResponse.FromString)
        self.AddGovernment = channel.unary_unary('/city.economy.v2.OrgService/AddGovernment', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddGovernmentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddGovernmentResponse.FromString)
        self.RemoveGovernment = channel.unary_unary('/city.economy.v2.OrgService/RemoveGovernment', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveGovernmentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveGovernmentResponse.FromString)
        self.GetGovernment = channel.unary_unary('/city.economy.v2.OrgService/GetGovernment', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetGovernmentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetGovernmentResponse.FromString)
        self.UpdateGovernment = channel.unary_unary('/city.economy.v2.OrgService/UpdateGovernment', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateGovernmentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateGovernmentResponse.FromString)
        self.ListGovernments = channel.unary_unary('/city.economy.v2.OrgService/ListGovernments', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListGovernmentsRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListGovernmentsResponse.FromString)
        self.DeltaUpdateGovernment = channel.unary_unary('/city.economy.v2.OrgService/DeltaUpdateGovernment', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateGovernmentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateGovernmentResponse.FromString)
        self.AddBank = channel.unary_unary('/city.economy.v2.OrgService/AddBank', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddBankRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddBankResponse.FromString)
        self.RemoveBank = channel.unary_unary('/city.economy.v2.OrgService/RemoveBank', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveBankRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveBankResponse.FromString)
        self.GetBank = channel.unary_unary('/city.economy.v2.OrgService/GetBank', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBankRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBankResponse.FromString)
        self.UpdateBank = channel.unary_unary('/city.economy.v2.OrgService/UpdateBank', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateBankRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateBankResponse.FromString)
        self.ListBanks = channel.unary_unary('/city.economy.v2.OrgService/ListBanks', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListBanksRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListBanksResponse.FromString)
        self.DeltaUpdateBank = channel.unary_unary('/city.economy.v2.OrgService/DeltaUpdateBank', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateBankRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateBankResponse.FromString)
        self.AddAgent = channel.unary_unary('/city.economy.v2.OrgService/AddAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentResponse.FromString)
        self.RemoveAgent = channel.unary_unary('/city.economy.v2.OrgService/RemoveAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentResponse.FromString)
        self.GetAgent = channel.unary_unary('/city.economy.v2.OrgService/GetAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentResponse.FromString)
        self.UpdateAgent = channel.unary_unary('/city.economy.v2.OrgService/UpdateAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentResponse.FromString)
        self.ListAgents = channel.unary_unary('/city.economy.v2.OrgService/ListAgents', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListAgentsRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListAgentsResponse.FromString)
        self.DeltaUpdateAgent = channel.unary_unary('/city.economy.v2.OrgService/DeltaUpdateAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentResponse.FromString)
        self.CalculateTaxesDue = channel.unary_unary('/city.economy.v2.OrgService/CalculateTaxesDue', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueResponse.FromString)
        self.CalculateConsumption = channel.unary_unary('/city.economy.v2.OrgService/CalculateConsumption', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionResponse.FromString)
        self.CalculateInterest = channel.unary_unary('/city.economy.v2.OrgService/CalculateInterest', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestResponse.FromString)
        self.CalculateRealGDP = channel.unary_unary('/city.economy.v2.OrgService/CalculateRealGDP', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPResponse.FromString)
        self.SaveEconomyEntities = channel.unary_unary('/city.economy.v2.OrgService/SaveEconomyEntities', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesResponse.FromString)
        self.LoadEconomyEntities = channel.unary_unary('/city.economy.v2.OrgService/LoadEconomyEntities', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesResponse.FromString)

class OrgServiceServicer(object):
    """OrgService provides management interfaces for organizations and agents in the economic system
    All basic operations support batch processing by default

    OrgService 提供了经济系统中组织和代理的管理接口
    所有基础操作默认支持批量处理
    @section Firm Operations
    @section 企业相关操作
    """

    def AddFirm(self, request, context):
        """Creates one or more firm entities
        Returns the list of created firm IDs

        创建一个或多个企业实体
        返回创建的企业ID列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveFirm(self, request, context):
        """Deletes one or more firms by their IDs
        Also cleans up related employee associations

        根据ID删除一个或多个企业实体
        同时清理相关的员工关联关系
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetFirm(self, request, context):
        """Retrieves detailed information for one or more firms
        Returns a list of firm entities

        获取一个或多个企业的详细信息
        返回企业实体信息列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateFirm(self, request, context):
        """Updates complete information for one or more firms

        更新一个或多个企业的完整信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListFirms(self, request, context):
        """Lists all firms in the system

        获取系统中所有企业的列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeltaUpdateFirm(self, request, context):
        """Performs incremental updates on one or more firms
        Allows updating specific fields like price, inventory, demand etc.

        对一个或多个企业进行增量更新
        可以更新价格、库存、需求等具体字段
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddNBS(self, request, context):
        """@section NBS Operations
        @section 国家统计局相关操作

        Creates a new NBS entity

        创建新的国家统计局实体
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveNBS(self, request, context):
        """Deletes an NBS entity by ID

        根据ID删除国家统计局实体
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetNBS(self, request, context):
        """Retrieves detailed information for an NBS entity

        获取国家统计局的详细信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateNBS(self, request, context):
        """Updates complete information for an NBS entity

        更新国家统计局的完整信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListNBS(self, request, context):
        """Lists all NBS entities in the system

        获取系统中所有国家统计局的列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeltaUpdateNBS(self, request, context):
        """Performs incremental updates on an NBS entity
        Allows updating specific statistics and citizen relationships

        对国家统计局进行增量更新
        可以更新具体统计数据和公民关系
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddGovernment(self, request, context):
        """@section Government Operations
        @section 政府相关操作

        Creates a new government entity

        创建新的政府实体
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveGovernment(self, request, context):
        """Deletes a government entity by ID

        根据ID删除政府实体
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetGovernment(self, request, context):
        """Retrieves detailed information for a government entity

        获取政府实体的详细信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateGovernment(self, request, context):
        """Updates complete information for a government entity

        更新政府实体的完整信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListGovernments(self, request, context):
        """Lists all government entities in the system

        获取系统中所有政府实体的列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeltaUpdateGovernment(self, request, context):
        """Performs incremental updates on a government entity
        Allows updating tax brackets and citizen relationships

        对政府实体进行增量更新
        可以更新税收档位和公民关系
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddBank(self, request, context):
        """@section Bank Operations
        @section 银行相关操作

        Creates a new bank entity

        创建新的银行实体
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveBank(self, request, context):
        """Deletes a bank entity by ID

        根据ID删除银行实体
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetBank(self, request, context):
        """Retrieves detailed information for a bank entity

        获取银行实体的详细信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateBank(self, request, context):
        """Updates complete information for a bank entity

        更新银行实体的完整信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListBanks(self, request, context):
        """Lists all bank entities in the system

        获取系统中所有银行实体的列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeltaUpdateBank(self, request, context):
        """Performs incremental updates on a bank entity
        Allows updating interest rates and customer relationships

        对银行实体进行增量更新
        可以更新利率和客户关系
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddAgent(self, request, context):
        """@section Agent Operations
        @section 经济主体相关操作

        Creates one or more agent entities
        Returns the list of created agent IDs

        创建一个或多个经济主体
        返回创建的经济主体ID列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveAgent(self, request, context):
        """Deletes one or more agents by their IDs

        根据ID删除一个或多个经济主体
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAgent(self, request, context):
        """Retrieves detailed information for one or more agents

        获取一个或多个经济主体的详细信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateAgent(self, request, context):
        """Updates complete information for one or more agents

        更新一个或多个经济主体的完整信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListAgents(self, request, context):
        """Lists all agent entities in the system

        获取系统中所有经济主体的列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeltaUpdateAgent(self, request, context):
        """Performs incremental updates on one or more agents
        Allows updating specific attributes like currency, skills etc.

        对一个或多个经济主体进行增量更新
        可以更新货币、技能等具体属性
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateTaxesDue(self, request, context):
        """@section Economic Calculations
        @section 经济计算相关操作

        Calculates taxes due for specified agents
        Supports income redistribution if enabled

        计算指定经济主体的应缴税额
        支持开启收入再分配功能
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateConsumption(self, request, context):
        """Calculates actual consumption based on supply and demand

        基于供给和需求计算实际消费量
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateInterest(self, request, context):
        """Calculates interest for specified agents

        计算指定经济主体的利息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateRealGDP(self, request, context):
        """Calculates real GDP adjusted for inflation

        计算经通货膨胀调整后的实际GDP
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SaveEconomyEntities(self, request, context):
        """@section System State Operations
        @section 系统状态相关操作

        Saves the current state of all economic entities to a file

        将当前所有经济实体的状态保存到文件
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def LoadEconomyEntities(self, request, context):
        """Loads economic entities state from a file

        从文件中加载经济实体的状态
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_OrgServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'AddFirm': grpc.unary_unary_rpc_method_handler(servicer.AddFirm, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddFirmRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddFirmResponse.SerializeToString), 'RemoveFirm': grpc.unary_unary_rpc_method_handler(servicer.RemoveFirm, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveFirmRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveFirmResponse.SerializeToString), 'GetFirm': grpc.unary_unary_rpc_method_handler(servicer.GetFirm, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetFirmRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetFirmResponse.SerializeToString), 'UpdateFirm': grpc.unary_unary_rpc_method_handler(servicer.UpdateFirm, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateFirmRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateFirmResponse.SerializeToString), 'ListFirms': grpc.unary_unary_rpc_method_handler(servicer.ListFirms, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListFirmsRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListFirmsResponse.SerializeToString), 'DeltaUpdateFirm': grpc.unary_unary_rpc_method_handler(servicer.DeltaUpdateFirm, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateFirmRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateFirmResponse.SerializeToString), 'AddNBS': grpc.unary_unary_rpc_method_handler(servicer.AddNBS, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddNBSRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddNBSResponse.SerializeToString), 'RemoveNBS': grpc.unary_unary_rpc_method_handler(servicer.RemoveNBS, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveNBSRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveNBSResponse.SerializeToString), 'GetNBS': grpc.unary_unary_rpc_method_handler(servicer.GetNBS, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNBSRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNBSResponse.SerializeToString), 'UpdateNBS': grpc.unary_unary_rpc_method_handler(servicer.UpdateNBS, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateNBSRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateNBSResponse.SerializeToString), 'ListNBS': grpc.unary_unary_rpc_method_handler(servicer.ListNBS, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListNBSRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListNBSResponse.SerializeToString), 'DeltaUpdateNBS': grpc.unary_unary_rpc_method_handler(servicer.DeltaUpdateNBS, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateNBSRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateNBSResponse.SerializeToString), 'AddGovernment': grpc.unary_unary_rpc_method_handler(servicer.AddGovernment, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddGovernmentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddGovernmentResponse.SerializeToString), 'RemoveGovernment': grpc.unary_unary_rpc_method_handler(servicer.RemoveGovernment, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveGovernmentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveGovernmentResponse.SerializeToString), 'GetGovernment': grpc.unary_unary_rpc_method_handler(servicer.GetGovernment, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetGovernmentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetGovernmentResponse.SerializeToString), 'UpdateGovernment': grpc.unary_unary_rpc_method_handler(servicer.UpdateGovernment, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateGovernmentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateGovernmentResponse.SerializeToString), 'ListGovernments': grpc.unary_unary_rpc_method_handler(servicer.ListGovernments, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListGovernmentsRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListGovernmentsResponse.SerializeToString), 'DeltaUpdateGovernment': grpc.unary_unary_rpc_method_handler(servicer.DeltaUpdateGovernment, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateGovernmentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateGovernmentResponse.SerializeToString), 'AddBank': grpc.unary_unary_rpc_method_handler(servicer.AddBank, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddBankRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddBankResponse.SerializeToString), 'RemoveBank': grpc.unary_unary_rpc_method_handler(servicer.RemoveBank, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveBankRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveBankResponse.SerializeToString), 'GetBank': grpc.unary_unary_rpc_method_handler(servicer.GetBank, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBankRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBankResponse.SerializeToString), 'UpdateBank': grpc.unary_unary_rpc_method_handler(servicer.UpdateBank, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateBankRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateBankResponse.SerializeToString), 'ListBanks': grpc.unary_unary_rpc_method_handler(servicer.ListBanks, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListBanksRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListBanksResponse.SerializeToString), 'DeltaUpdateBank': grpc.unary_unary_rpc_method_handler(servicer.DeltaUpdateBank, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateBankRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateBankResponse.SerializeToString), 'AddAgent': grpc.unary_unary_rpc_method_handler(servicer.AddAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentResponse.SerializeToString), 'RemoveAgent': grpc.unary_unary_rpc_method_handler(servicer.RemoveAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentResponse.SerializeToString), 'GetAgent': grpc.unary_unary_rpc_method_handler(servicer.GetAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentResponse.SerializeToString), 'UpdateAgent': grpc.unary_unary_rpc_method_handler(servicer.UpdateAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentResponse.SerializeToString), 'ListAgents': grpc.unary_unary_rpc_method_handler(servicer.ListAgents, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListAgentsRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.ListAgentsResponse.SerializeToString), 'DeltaUpdateAgent': grpc.unary_unary_rpc_method_handler(servicer.DeltaUpdateAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentResponse.SerializeToString), 'CalculateTaxesDue': grpc.unary_unary_rpc_method_handler(servicer.CalculateTaxesDue, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueResponse.SerializeToString), 'CalculateConsumption': grpc.unary_unary_rpc_method_handler(servicer.CalculateConsumption, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionResponse.SerializeToString), 'CalculateInterest': grpc.unary_unary_rpc_method_handler(servicer.CalculateInterest, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestResponse.SerializeToString), 'CalculateRealGDP': grpc.unary_unary_rpc_method_handler(servicer.CalculateRealGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPResponse.SerializeToString), 'SaveEconomyEntities': grpc.unary_unary_rpc_method_handler(servicer.SaveEconomyEntities, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesResponse.SerializeToString), 'LoadEconomyEntities': grpc.unary_unary_rpc_method_handler(servicer.LoadEconomyEntities, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.economy.v2.OrgService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class OrgService(object):
    """OrgService provides management interfaces for organizations and agents in the economic system
    All basic operations support batch processing by default

    OrgService 提供了经济系统中组织和代理的管理接口
    所有基础操作默认支持批量处理
    @section Firm Operations
    @section 企业相关操作
    """

    @staticmethod
    def AddFirm(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddFirm', city_dot_economy_dot_v2_dot_org__service__pb2.AddFirmRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddFirmResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveFirm(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/RemoveFirm', city_dot_economy_dot_v2_dot_org__service__pb2.RemoveFirmRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.RemoveFirmResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetFirm(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetFirm', city_dot_economy_dot_v2_dot_org__service__pb2.GetFirmRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetFirmResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateFirm(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/UpdateFirm', city_dot_economy_dot_v2_dot_org__service__pb2.UpdateFirmRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.UpdateFirmResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListFirms(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/ListFirms', city_dot_economy_dot_v2_dot_org__service__pb2.ListFirmsRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.ListFirmsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeltaUpdateFirm(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/DeltaUpdateFirm', city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateFirmRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateFirmResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddNBS(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddNBS', city_dot_economy_dot_v2_dot_org__service__pb2.AddNBSRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddNBSResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveNBS(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/RemoveNBS', city_dot_economy_dot_v2_dot_org__service__pb2.RemoveNBSRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.RemoveNBSResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetNBS(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetNBS', city_dot_economy_dot_v2_dot_org__service__pb2.GetNBSRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetNBSResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateNBS(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/UpdateNBS', city_dot_economy_dot_v2_dot_org__service__pb2.UpdateNBSRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.UpdateNBSResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListNBS(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/ListNBS', city_dot_economy_dot_v2_dot_org__service__pb2.ListNBSRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.ListNBSResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeltaUpdateNBS(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/DeltaUpdateNBS', city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateNBSRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateNBSResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddGovernment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddGovernment', city_dot_economy_dot_v2_dot_org__service__pb2.AddGovernmentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddGovernmentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveGovernment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/RemoveGovernment', city_dot_economy_dot_v2_dot_org__service__pb2.RemoveGovernmentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.RemoveGovernmentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetGovernment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetGovernment', city_dot_economy_dot_v2_dot_org__service__pb2.GetGovernmentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetGovernmentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateGovernment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/UpdateGovernment', city_dot_economy_dot_v2_dot_org__service__pb2.UpdateGovernmentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.UpdateGovernmentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListGovernments(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/ListGovernments', city_dot_economy_dot_v2_dot_org__service__pb2.ListGovernmentsRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.ListGovernmentsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeltaUpdateGovernment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/DeltaUpdateGovernment', city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateGovernmentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateGovernmentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddBank(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddBank', city_dot_economy_dot_v2_dot_org__service__pb2.AddBankRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddBankResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveBank(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/RemoveBank', city_dot_economy_dot_v2_dot_org__service__pb2.RemoveBankRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.RemoveBankResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetBank(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetBank', city_dot_economy_dot_v2_dot_org__service__pb2.GetBankRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetBankResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateBank(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/UpdateBank', city_dot_economy_dot_v2_dot_org__service__pb2.UpdateBankRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.UpdateBankResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListBanks(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/ListBanks', city_dot_economy_dot_v2_dot_org__service__pb2.ListBanksRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.ListBanksResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeltaUpdateBank(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/DeltaUpdateBank', city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateBankRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateBankResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
    def ListAgents(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/ListAgents', city_dot_economy_dot_v2_dot_org__service__pb2.ListAgentsRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.ListAgentsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeltaUpdateAgent(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/DeltaUpdateAgent', city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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