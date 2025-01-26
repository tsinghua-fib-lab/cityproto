"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.economy.v2 import org_service_pb2 as city_dot_economy_dot_v2_dot_org__service__pb2

class OrgServiceStub(object):
    """组织经济情况接口
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
        self.GetNominalGDP = channel.unary_unary('/city.economy.v2.OrgService/GetNominalGDP', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNominalGDPRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNominalGDPResponse.FromString)
        self.SetNominalGDP = channel.unary_unary('/city.economy.v2.OrgService/SetNominalGDP', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetNominalGDPRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetNominalGDPResponse.FromString)
        self.GetRealGDP = channel.unary_unary('/city.economy.v2.OrgService/GetRealGDP', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetRealGDPRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetRealGDPResponse.FromString)
        self.SetRealGDP = channel.unary_unary('/city.economy.v2.OrgService/SetRealGDP', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetRealGDPRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetRealGDPResponse.FromString)
        self.GetUnemployment = channel.unary_unary('/city.economy.v2.OrgService/GetUnemployment', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetUnemploymentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetUnemploymentResponse.FromString)
        self.SetUnemployment = channel.unary_unary('/city.economy.v2.OrgService/SetUnemployment', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetUnemploymentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetUnemploymentResponse.FromString)
        self.GetWages = channel.unary_unary('/city.economy.v2.OrgService/GetWages', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWagesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWagesResponse.FromString)
        self.SetWages = channel.unary_unary('/city.economy.v2.OrgService/SetWages', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWagesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWagesResponse.FromString)
        self.GetPrices = channel.unary_unary('/city.economy.v2.OrgService/GetPrices', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPricesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPricesResponse.FromString)
        self.SetPrices = channel.unary_unary('/city.economy.v2.OrgService/SetPrices', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPricesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPricesResponse.FromString)
        self.GetInventory = channel.unary_unary('/city.economy.v2.OrgService/GetInventory', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInventoryRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInventoryResponse.FromString)
        self.SetInventory = channel.unary_unary('/city.economy.v2.OrgService/SetInventory', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInventoryRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInventoryResponse.FromString)
        self.AddInventory = channel.unary_unary('/city.economy.v2.OrgService/AddInventory', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddInventoryRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddInventoryResponse.FromString)
        self.GetPrice = channel.unary_unary('/city.economy.v2.OrgService/GetPrice', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceResponse.FromString)
        self.SetPrice = channel.unary_unary('/city.economy.v2.OrgService/SetPrice', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceResponse.FromString)
        self.AddPrice = channel.unary_unary('/city.economy.v2.OrgService/AddPrice', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddPriceRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddPriceResponse.FromString)
        self.GetCurrency = channel.unary_unary('/city.economy.v2.OrgService/GetCurrency', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyResponse.FromString)
        self.SetCurrency = channel.unary_unary('/city.economy.v2.OrgService/SetCurrency', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyResponse.FromString)
        self.AddCurrency = channel.unary_unary('/city.economy.v2.OrgService/AddCurrency', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddCurrencyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddCurrencyResponse.FromString)
        self.GetInterestRate = channel.unary_unary('/city.economy.v2.OrgService/GetInterestRate', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateResponse.FromString)
        self.SetInterestRate = channel.unary_unary('/city.economy.v2.OrgService/SetInterestRate', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateResponse.FromString)
        self.AddInterestRate = channel.unary_unary('/city.economy.v2.OrgService/AddInterestRate', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddInterestRateRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddInterestRateResponse.FromString)
        self.GetBracketCutoffs = channel.unary_unary('/city.economy.v2.OrgService/GetBracketCutoffs', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsResponse.FromString)
        self.SetBracketCutoffs = channel.unary_unary('/city.economy.v2.OrgService/SetBracketCutoffs', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsResponse.FromString)
        self.GetBracketRates = channel.unary_unary('/city.economy.v2.OrgService/GetBracketRates', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesResponse.FromString)
        self.SetBracketRates = channel.unary_unary('/city.economy.v2.OrgService/SetBracketRates', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesResponse.FromString)
        self.CalculateTaxesDue = channel.unary_unary('/city.economy.v2.OrgService/CalculateTaxesDue', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueResponse.FromString)
        self.CalculateConsumption = channel.unary_unary('/city.economy.v2.OrgService/CalculateConsumption', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionResponse.FromString)
        self.CalculateInterest = channel.unary_unary('/city.economy.v2.OrgService/CalculateInterest', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestResponse.FromString)
        self.SaveEconomyEntities = channel.unary_unary('/city.economy.v2.OrgService/SaveEconomyEntities', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesResponse.FromString)
        self.LoadEconomyEntities = channel.unary_unary('/city.economy.v2.OrgService/LoadEconomyEntities', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesResponse.FromString)
        self.GetConsumptionCurrency = channel.unary_unary('/city.economy.v2.OrgService/GetConsumptionCurrency', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionCurrencyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionCurrencyResponse.FromString)
        self.SetConsumptionCurrency = channel.unary_unary('/city.economy.v2.OrgService/SetConsumptionCurrency', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionCurrencyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionCurrencyResponse.FromString)
        self.GetConsumptionPropensity = channel.unary_unary('/city.economy.v2.OrgService/GetConsumptionPropensity', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionPropensityRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionPropensityResponse.FromString)
        self.SetConsumptionPropensity = channel.unary_unary('/city.economy.v2.OrgService/SetConsumptionPropensity', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionPropensityRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionPropensityResponse.FromString)
        self.GetIncomeCurrency = channel.unary_unary('/city.economy.v2.OrgService/GetIncomeCurrency', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetIncomeCurrencyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetIncomeCurrencyResponse.FromString)
        self.SetIncomeCurrency = channel.unary_unary('/city.economy.v2.OrgService/SetIncomeCurrency', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetIncomeCurrencyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetIncomeCurrencyResponse.FromString)
        self.GetDepression = channel.unary_unary('/city.economy.v2.OrgService/GetDepression', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetDepressionRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetDepressionResponse.FromString)
        self.SetDepression = channel.unary_unary('/city.economy.v2.OrgService/SetDepression', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetDepressionRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetDepressionResponse.FromString)
        self.GetLocusControl = channel.unary_unary('/city.economy.v2.OrgService/GetLocusControl', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetLocusControlRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetLocusControlResponse.FromString)
        self.SetLocusControl = channel.unary_unary('/city.economy.v2.OrgService/SetLocusControl', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetLocusControlRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetLocusControlResponse.FromString)
        self.GetWorkingHours = channel.unary_unary('/city.economy.v2.OrgService/GetWorkingHours', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWorkingHoursRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWorkingHoursResponse.FromString)
        self.SetWorkingHours = channel.unary_unary('/city.economy.v2.OrgService/SetWorkingHours', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWorkingHoursRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWorkingHoursResponse.FromString)
        self.GetOrgEntityIds = channel.unary_unary('/city.economy.v2.OrgService/GetOrgEntityIds', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgEntityIdsRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgEntityIdsResponse.FromString)
        self.GetEmployees = channel.unary_unary('/city.economy.v2.OrgService/GetEmployees', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetEmployeesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetEmployeesResponse.FromString)
        self.SetEmployees = channel.unary_unary('/city.economy.v2.OrgService/SetEmployees', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetEmployeesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetEmployeesResponse.FromString)
        self.AddEmployee = channel.unary_unary('/city.economy.v2.OrgService/AddEmployee', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddEmployeeRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddEmployeeResponse.FromString)
        self.RemoveEmployee = channel.unary_unary('/city.economy.v2.OrgService/RemoveEmployee', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveEmployeeRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveEmployeeResponse.FromString)
        self.GetCitizens = channel.unary_unary('/city.economy.v2.OrgService/GetCitizens', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCitizensRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCitizensResponse.FromString)
        self.SetCitizens = channel.unary_unary('/city.economy.v2.OrgService/SetCitizens', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCitizensRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCitizensResponse.FromString)
        self.AddCitizen = channel.unary_unary('/city.economy.v2.OrgService/AddCitizen', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddCitizenRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddCitizenResponse.FromString)
        self.RemoveCitizen = channel.unary_unary('/city.economy.v2.OrgService/RemoveCitizen', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveCitizenRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveCitizenResponse.FromString)
        self.GetAgent = channel.unary_unary('/city.economy.v2.OrgService/GetAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentResponse.FromString)
        self.UpdateAgent = channel.unary_unary('/city.economy.v2.OrgService/UpdateAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentResponse.FromString)
        self.BatchGet = channel.unary_unary('/city.economy.v2.OrgService/BatchGet', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetResponse.FromString)
        self.BatchUpdate = channel.unary_unary('/city.economy.v2.OrgService/BatchUpdate', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateResponse.FromString)
        self.DeltaUpdateOrg = channel.unary_unary('/city.economy.v2.OrgService/DeltaUpdateOrg', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgResponse.FromString)
        self.DeltaUpdateAgent = channel.unary_unary('/city.economy.v2.OrgService/DeltaUpdateAgent', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentResponse.FromString)
        self.BatchDeltaUpdate = channel.unary_unary('/city.economy.v2.OrgService/BatchDeltaUpdate', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateResponse.FromString)
        self.CalculateRealGDP = channel.unary_unary('/city.economy.v2.OrgService/CalculateRealGDP', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPResponse.FromString)

class OrgServiceServicer(object):
    """组织经济情况接口
    """

    def AddOrg(self, request, context):
        """添加组织
        add org
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveOrg(self, request, context):
        """移除组织
        remove org
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOrg(self, request, context):
        """获取组织
        get org
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateOrg(self, request, context):
        """更新组织
        update org
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddAgent(self, request, context):
        """添加Agent
        add agent
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveAgent(self, request, context):
        """移除Agent
        remove agent
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetNominalGDP(self, request, context):
        """Nominal GDP
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetNominalGDP(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRealGDP(self, request, context):
        """Real GDP
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetRealGDP(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUnemployment(self, request, context):
        """Unemployment
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetUnemployment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetWages(self, request, context):
        """Wages
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetWages(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPrices(self, request, context):
        """Prices
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetPrices(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetInventory(self, request, context):
        """Inventory
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetInventory(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddInventory(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPrice(self, request, context):
        """Price
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetPrice(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddPrice(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCurrency(self, request, context):
        """Currency
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetCurrency(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddCurrency(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetInterestRate(self, request, context):
        """Interest Rate
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetInterestRate(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddInterestRate(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetBracketCutoffs(self, request, context):
        """Bracket Cutoffs
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetBracketCutoffs(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetBracketRates(self, request, context):
        """Bracket Rates
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetBracketRates(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateTaxesDue(self, request, context):
        """Taxes Due
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateConsumption(self, request, context):
        """Consumption
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateInterest(self, request, context):
        """Consumption
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SaveEconomyEntities(self, request, context):
        """Save
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def LoadEconomyEntities(self, request, context):
        """Load
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetConsumptionCurrency(self, request, context):
        """Consumption Currency
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetConsumptionCurrency(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetConsumptionPropensity(self, request, context):
        """Consumption Propensity
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetConsumptionPropensity(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetIncomeCurrency(self, request, context):
        """Income Currency
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetIncomeCurrency(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetDepression(self, request, context):
        """Depression
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetDepression(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetLocusControl(self, request, context):
        """Locus of Control
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetLocusControl(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetWorkingHours(self, request, context):
        """Working Hours
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetWorkingHours(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOrgEntityIds(self, request, context):
        """Org Entity Ids
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetEmployees(self, request, context):
        """Employees 相关接口
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetEmployees(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddEmployee(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveEmployee(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCitizens(self, request, context):
        """Citizens 相关接口
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetCitizens(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddCitizen(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveCitizen(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAgent(self, request, context):
        """Agent 相关接口
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateAgent(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchGet(self, request, context):
        """批量获取
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchUpdate(self, request, context):
        """批量更新
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeltaUpdateOrg(self, request, context):
        """增量更新
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeltaUpdateAgent(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchDeltaUpdate(self, request, context):
        """批量增量更新
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateRealGDP(self, request, context):
        """计算实际GDP
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_OrgServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'AddOrg': grpc.unary_unary_rpc_method_handler(servicer.AddOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgResponse.SerializeToString), 'RemoveOrg': grpc.unary_unary_rpc_method_handler(servicer.RemoveOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgResponse.SerializeToString), 'GetOrg': grpc.unary_unary_rpc_method_handler(servicer.GetOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgResponse.SerializeToString), 'UpdateOrg': grpc.unary_unary_rpc_method_handler(servicer.UpdateOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateOrgResponse.SerializeToString), 'AddAgent': grpc.unary_unary_rpc_method_handler(servicer.AddAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentResponse.SerializeToString), 'RemoveAgent': grpc.unary_unary_rpc_method_handler(servicer.RemoveAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentResponse.SerializeToString), 'GetNominalGDP': grpc.unary_unary_rpc_method_handler(servicer.GetNominalGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNominalGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNominalGDPResponse.SerializeToString), 'SetNominalGDP': grpc.unary_unary_rpc_method_handler(servicer.SetNominalGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetNominalGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetNominalGDPResponse.SerializeToString), 'GetRealGDP': grpc.unary_unary_rpc_method_handler(servicer.GetRealGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetRealGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetRealGDPResponse.SerializeToString), 'SetRealGDP': grpc.unary_unary_rpc_method_handler(servicer.SetRealGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetRealGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetRealGDPResponse.SerializeToString), 'GetUnemployment': grpc.unary_unary_rpc_method_handler(servicer.GetUnemployment, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetUnemploymentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetUnemploymentResponse.SerializeToString), 'SetUnemployment': grpc.unary_unary_rpc_method_handler(servicer.SetUnemployment, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetUnemploymentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetUnemploymentResponse.SerializeToString), 'GetWages': grpc.unary_unary_rpc_method_handler(servicer.GetWages, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWagesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWagesResponse.SerializeToString), 'SetWages': grpc.unary_unary_rpc_method_handler(servicer.SetWages, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWagesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWagesResponse.SerializeToString), 'GetPrices': grpc.unary_unary_rpc_method_handler(servicer.GetPrices, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPricesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPricesResponse.SerializeToString), 'SetPrices': grpc.unary_unary_rpc_method_handler(servicer.SetPrices, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPricesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPricesResponse.SerializeToString), 'GetInventory': grpc.unary_unary_rpc_method_handler(servicer.GetInventory, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInventoryRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInventoryResponse.SerializeToString), 'SetInventory': grpc.unary_unary_rpc_method_handler(servicer.SetInventory, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInventoryRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInventoryResponse.SerializeToString), 'AddInventory': grpc.unary_unary_rpc_method_handler(servicer.AddInventory, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddInventoryRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddInventoryResponse.SerializeToString), 'GetPrice': grpc.unary_unary_rpc_method_handler(servicer.GetPrice, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceResponse.SerializeToString), 'SetPrice': grpc.unary_unary_rpc_method_handler(servicer.SetPrice, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceResponse.SerializeToString), 'AddPrice': grpc.unary_unary_rpc_method_handler(servicer.AddPrice, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddPriceRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddPriceResponse.SerializeToString), 'GetCurrency': grpc.unary_unary_rpc_method_handler(servicer.GetCurrency, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyResponse.SerializeToString), 'SetCurrency': grpc.unary_unary_rpc_method_handler(servicer.SetCurrency, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyResponse.SerializeToString), 'AddCurrency': grpc.unary_unary_rpc_method_handler(servicer.AddCurrency, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddCurrencyRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddCurrencyResponse.SerializeToString), 'GetInterestRate': grpc.unary_unary_rpc_method_handler(servicer.GetInterestRate, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateResponse.SerializeToString), 'SetInterestRate': grpc.unary_unary_rpc_method_handler(servicer.SetInterestRate, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateResponse.SerializeToString), 'AddInterestRate': grpc.unary_unary_rpc_method_handler(servicer.AddInterestRate, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddInterestRateRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddInterestRateResponse.SerializeToString), 'GetBracketCutoffs': grpc.unary_unary_rpc_method_handler(servicer.GetBracketCutoffs, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsResponse.SerializeToString), 'SetBracketCutoffs': grpc.unary_unary_rpc_method_handler(servicer.SetBracketCutoffs, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsResponse.SerializeToString), 'GetBracketRates': grpc.unary_unary_rpc_method_handler(servicer.GetBracketRates, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesResponse.SerializeToString), 'SetBracketRates': grpc.unary_unary_rpc_method_handler(servicer.SetBracketRates, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesResponse.SerializeToString), 'CalculateTaxesDue': grpc.unary_unary_rpc_method_handler(servicer.CalculateTaxesDue, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueResponse.SerializeToString), 'CalculateConsumption': grpc.unary_unary_rpc_method_handler(servicer.CalculateConsumption, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionResponse.SerializeToString), 'CalculateInterest': grpc.unary_unary_rpc_method_handler(servicer.CalculateInterest, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestResponse.SerializeToString), 'SaveEconomyEntities': grpc.unary_unary_rpc_method_handler(servicer.SaveEconomyEntities, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesResponse.SerializeToString), 'LoadEconomyEntities': grpc.unary_unary_rpc_method_handler(servicer.LoadEconomyEntities, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesResponse.SerializeToString), 'GetConsumptionCurrency': grpc.unary_unary_rpc_method_handler(servicer.GetConsumptionCurrency, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionCurrencyRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionCurrencyResponse.SerializeToString), 'SetConsumptionCurrency': grpc.unary_unary_rpc_method_handler(servicer.SetConsumptionCurrency, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionCurrencyRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionCurrencyResponse.SerializeToString), 'GetConsumptionPropensity': grpc.unary_unary_rpc_method_handler(servicer.GetConsumptionPropensity, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionPropensityRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionPropensityResponse.SerializeToString), 'SetConsumptionPropensity': grpc.unary_unary_rpc_method_handler(servicer.SetConsumptionPropensity, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionPropensityRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionPropensityResponse.SerializeToString), 'GetIncomeCurrency': grpc.unary_unary_rpc_method_handler(servicer.GetIncomeCurrency, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetIncomeCurrencyRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetIncomeCurrencyResponse.SerializeToString), 'SetIncomeCurrency': grpc.unary_unary_rpc_method_handler(servicer.SetIncomeCurrency, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetIncomeCurrencyRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetIncomeCurrencyResponse.SerializeToString), 'GetDepression': grpc.unary_unary_rpc_method_handler(servicer.GetDepression, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetDepressionRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetDepressionResponse.SerializeToString), 'SetDepression': grpc.unary_unary_rpc_method_handler(servicer.SetDepression, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetDepressionRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetDepressionResponse.SerializeToString), 'GetLocusControl': grpc.unary_unary_rpc_method_handler(servicer.GetLocusControl, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetLocusControlRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetLocusControlResponse.SerializeToString), 'SetLocusControl': grpc.unary_unary_rpc_method_handler(servicer.SetLocusControl, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetLocusControlRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetLocusControlResponse.SerializeToString), 'GetWorkingHours': grpc.unary_unary_rpc_method_handler(servicer.GetWorkingHours, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWorkingHoursRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWorkingHoursResponse.SerializeToString), 'SetWorkingHours': grpc.unary_unary_rpc_method_handler(servicer.SetWorkingHours, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWorkingHoursRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWorkingHoursResponse.SerializeToString), 'GetOrgEntityIds': grpc.unary_unary_rpc_method_handler(servicer.GetOrgEntityIds, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgEntityIdsRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgEntityIdsResponse.SerializeToString), 'GetEmployees': grpc.unary_unary_rpc_method_handler(servicer.GetEmployees, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetEmployeesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetEmployeesResponse.SerializeToString), 'SetEmployees': grpc.unary_unary_rpc_method_handler(servicer.SetEmployees, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetEmployeesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetEmployeesResponse.SerializeToString), 'AddEmployee': grpc.unary_unary_rpc_method_handler(servicer.AddEmployee, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddEmployeeRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddEmployeeResponse.SerializeToString), 'RemoveEmployee': grpc.unary_unary_rpc_method_handler(servicer.RemoveEmployee, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveEmployeeRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveEmployeeResponse.SerializeToString), 'GetCitizens': grpc.unary_unary_rpc_method_handler(servicer.GetCitizens, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCitizensRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCitizensResponse.SerializeToString), 'SetCitizens': grpc.unary_unary_rpc_method_handler(servicer.SetCitizens, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCitizensRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCitizensResponse.SerializeToString), 'AddCitizen': grpc.unary_unary_rpc_method_handler(servicer.AddCitizen, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddCitizenRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddCitizenResponse.SerializeToString), 'RemoveCitizen': grpc.unary_unary_rpc_method_handler(servicer.RemoveCitizen, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveCitizenRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveCitizenResponse.SerializeToString), 'GetAgent': grpc.unary_unary_rpc_method_handler(servicer.GetAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetAgentResponse.SerializeToString), 'UpdateAgent': grpc.unary_unary_rpc_method_handler(servicer.UpdateAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.UpdateAgentResponse.SerializeToString), 'BatchGet': grpc.unary_unary_rpc_method_handler(servicer.BatchGet, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchGetResponse.SerializeToString), 'BatchUpdate': grpc.unary_unary_rpc_method_handler(servicer.BatchUpdate, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchUpdateResponse.SerializeToString), 'DeltaUpdateOrg': grpc.unary_unary_rpc_method_handler(servicer.DeltaUpdateOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgResponse.SerializeToString), 'DeltaUpdateAgent': grpc.unary_unary_rpc_method_handler(servicer.DeltaUpdateAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentResponse.SerializeToString), 'BatchDeltaUpdate': grpc.unary_unary_rpc_method_handler(servicer.BatchDeltaUpdate, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateResponse.SerializeToString), 'CalculateRealGDP': grpc.unary_unary_rpc_method_handler(servicer.CalculateRealGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.economy.v2.OrgService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class OrgService(object):
    """组织经济情况接口
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
    def GetNominalGDP(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetNominalGDP', city_dot_economy_dot_v2_dot_org__service__pb2.GetNominalGDPRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetNominalGDPResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetNominalGDP(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetNominalGDP', city_dot_economy_dot_v2_dot_org__service__pb2.SetNominalGDPRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetNominalGDPResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetRealGDP(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetRealGDP', city_dot_economy_dot_v2_dot_org__service__pb2.GetRealGDPRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetRealGDPResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetRealGDP(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetRealGDP', city_dot_economy_dot_v2_dot_org__service__pb2.SetRealGDPRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetRealGDPResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUnemployment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetUnemployment', city_dot_economy_dot_v2_dot_org__service__pb2.GetUnemploymentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetUnemploymentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetUnemployment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetUnemployment', city_dot_economy_dot_v2_dot_org__service__pb2.SetUnemploymentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetUnemploymentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetWages(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetWages', city_dot_economy_dot_v2_dot_org__service__pb2.GetWagesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetWagesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetWages(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetWages', city_dot_economy_dot_v2_dot_org__service__pb2.SetWagesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetWagesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPrices(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetPrices', city_dot_economy_dot_v2_dot_org__service__pb2.GetPricesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetPricesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetPrices(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetPrices', city_dot_economy_dot_v2_dot_org__service__pb2.SetPricesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetPricesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetInventory(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetInventory', city_dot_economy_dot_v2_dot_org__service__pb2.GetInventoryRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetInventoryResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetInventory(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetInventory', city_dot_economy_dot_v2_dot_org__service__pb2.SetInventoryRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetInventoryResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddInventory(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddInventory', city_dot_economy_dot_v2_dot_org__service__pb2.AddInventoryRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddInventoryResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPrice(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetPrice', city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetPrice(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetPrice', city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddPrice(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddPrice', city_dot_economy_dot_v2_dot_org__service__pb2.AddPriceRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddPriceResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCurrency(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetCurrency', city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetCurrency(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetCurrency', city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddCurrency(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddCurrency', city_dot_economy_dot_v2_dot_org__service__pb2.AddCurrencyRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddCurrencyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetInterestRate(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetInterestRate', city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetInterestRate(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetInterestRate', city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddInterestRate(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddInterestRate', city_dot_economy_dot_v2_dot_org__service__pb2.AddInterestRateRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddInterestRateResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetBracketCutoffs(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetBracketCutoffs', city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetBracketCutoffs(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetBracketCutoffs', city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetBracketRates(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetBracketRates', city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetBracketRates(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetBracketRates', city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
    def SaveEconomyEntities(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SaveEconomyEntities', city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SaveEconomyEntitiesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def LoadEconomyEntities(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/LoadEconomyEntities', city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.LoadEconomyEntitiesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetConsumptionCurrency(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetConsumptionCurrency', city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionCurrencyRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionCurrencyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetConsumptionCurrency(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetConsumptionCurrency', city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionCurrencyRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionCurrencyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetConsumptionPropensity(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetConsumptionPropensity', city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionPropensityRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetConsumptionPropensityResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetConsumptionPropensity(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetConsumptionPropensity', city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionPropensityRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetConsumptionPropensityResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetIncomeCurrency(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetIncomeCurrency', city_dot_economy_dot_v2_dot_org__service__pb2.GetIncomeCurrencyRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetIncomeCurrencyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetIncomeCurrency(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetIncomeCurrency', city_dot_economy_dot_v2_dot_org__service__pb2.SetIncomeCurrencyRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetIncomeCurrencyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetDepression(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetDepression', city_dot_economy_dot_v2_dot_org__service__pb2.GetDepressionRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetDepressionResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetDepression(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetDepression', city_dot_economy_dot_v2_dot_org__service__pb2.SetDepressionRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetDepressionResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetLocusControl(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetLocusControl', city_dot_economy_dot_v2_dot_org__service__pb2.GetLocusControlRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetLocusControlResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetLocusControl(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetLocusControl', city_dot_economy_dot_v2_dot_org__service__pb2.SetLocusControlRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetLocusControlResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetWorkingHours(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetWorkingHours', city_dot_economy_dot_v2_dot_org__service__pb2.GetWorkingHoursRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetWorkingHoursResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetWorkingHours(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetWorkingHours', city_dot_economy_dot_v2_dot_org__service__pb2.SetWorkingHoursRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetWorkingHoursResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetOrgEntityIds(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetOrgEntityIds', city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgEntityIdsRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetOrgEntityIdsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetEmployees(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetEmployees', city_dot_economy_dot_v2_dot_org__service__pb2.GetEmployeesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetEmployeesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetEmployees(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetEmployees', city_dot_economy_dot_v2_dot_org__service__pb2.SetEmployeesRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetEmployeesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddEmployee(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddEmployee', city_dot_economy_dot_v2_dot_org__service__pb2.AddEmployeeRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddEmployeeResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveEmployee(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/RemoveEmployee', city_dot_economy_dot_v2_dot_org__service__pb2.RemoveEmployeeRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.RemoveEmployeeResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCitizens(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetCitizens', city_dot_economy_dot_v2_dot_org__service__pb2.GetCitizensRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetCitizensResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetCitizens(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetCitizens', city_dot_economy_dot_v2_dot_org__service__pb2.SetCitizensRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetCitizensResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddCitizen(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/AddCitizen', city_dot_economy_dot_v2_dot_org__service__pb2.AddCitizenRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.AddCitizenResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveCitizen(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/RemoveCitizen', city_dot_economy_dot_v2_dot_org__service__pb2.RemoveCitizenRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.RemoveCitizenResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
    def DeltaUpdateOrg(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/DeltaUpdateOrg', city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateOrgResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeltaUpdateAgent(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/DeltaUpdateAgent', city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.DeltaUpdateAgentResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchDeltaUpdate(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/BatchDeltaUpdate', city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.BatchDeltaUpdateResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CalculateRealGDP(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/CalculateRealGDP', city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.CalculateRealGDPResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)