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
        self.GetPrice = channel.unary_unary('/city.economy.v2.OrgService/GetPrice', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceResponse.FromString)
        self.SetPrice = channel.unary_unary('/city.economy.v2.OrgService/SetPrice', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceResponse.FromString)
        self.GetCurrency = channel.unary_unary('/city.economy.v2.OrgService/GetCurrency', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyResponse.FromString)
        self.SetCurrency = channel.unary_unary('/city.economy.v2.OrgService/SetCurrency', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyResponse.FromString)
        self.GetInterestRate = channel.unary_unary('/city.economy.v2.OrgService/GetInterestRate', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateResponse.FromString)
        self.SetInterestRate = channel.unary_unary('/city.economy.v2.OrgService/SetInterestRate', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateResponse.FromString)
        self.GetBracketCutoffs = channel.unary_unary('/city.economy.v2.OrgService/GetBracketCutoffs', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsResponse.FromString)
        self.SetBracketCutoffs = channel.unary_unary('/city.economy.v2.OrgService/SetBracketCutoffs', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsResponse.FromString)
        self.GetBracketRates = channel.unary_unary('/city.economy.v2.OrgService/GetBracketRates', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesResponse.FromString)
        self.SetBracketRates = channel.unary_unary('/city.economy.v2.OrgService/SetBracketRates', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesResponse.FromString)
        self.CalculateTaxesDue = channel.unary_unary('/city.economy.v2.OrgService/CalculateTaxesDue', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueResponse.FromString)
        self.CalculateConsumption = channel.unary_unary('/city.economy.v2.OrgService/CalculateConsumption', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionResponse.FromString)
        self.CalculateInterest = channel.unary_unary('/city.economy.v2.OrgService/CalculateInterest', request_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestResponse.FromString)

class OrgServiceServicer(object):
    """组织经济情况接口
    """

    def AddOrg(self, request, context):
        """添加组织
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveOrg(self, request, context):
        """移除组织
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddAgent(self, request, context):
        """添加Agent
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveAgent(self, request, context):
        """移除Agent
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

def add_OrgServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'AddOrg': grpc.unary_unary_rpc_method_handler(servicer.AddOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddOrgResponse.SerializeToString), 'RemoveOrg': grpc.unary_unary_rpc_method_handler(servicer.RemoveOrg, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveOrgResponse.SerializeToString), 'AddAgent': grpc.unary_unary_rpc_method_handler(servicer.AddAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.AddAgentResponse.SerializeToString), 'RemoveAgent': grpc.unary_unary_rpc_method_handler(servicer.RemoveAgent, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.RemoveAgentResponse.SerializeToString), 'GetNominalGDP': grpc.unary_unary_rpc_method_handler(servicer.GetNominalGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNominalGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetNominalGDPResponse.SerializeToString), 'SetNominalGDP': grpc.unary_unary_rpc_method_handler(servicer.SetNominalGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetNominalGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetNominalGDPResponse.SerializeToString), 'GetRealGDP': grpc.unary_unary_rpc_method_handler(servicer.GetRealGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetRealGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetRealGDPResponse.SerializeToString), 'SetRealGDP': grpc.unary_unary_rpc_method_handler(servicer.SetRealGDP, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetRealGDPRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetRealGDPResponse.SerializeToString), 'GetUnemployment': grpc.unary_unary_rpc_method_handler(servicer.GetUnemployment, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetUnemploymentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetUnemploymentResponse.SerializeToString), 'SetUnemployment': grpc.unary_unary_rpc_method_handler(servicer.SetUnemployment, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetUnemploymentRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetUnemploymentResponse.SerializeToString), 'GetWages': grpc.unary_unary_rpc_method_handler(servicer.GetWages, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWagesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetWagesResponse.SerializeToString), 'SetWages': grpc.unary_unary_rpc_method_handler(servicer.SetWages, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWagesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetWagesResponse.SerializeToString), 'GetPrices': grpc.unary_unary_rpc_method_handler(servicer.GetPrices, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPricesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPricesResponse.SerializeToString), 'SetPrices': grpc.unary_unary_rpc_method_handler(servicer.SetPrices, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPricesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPricesResponse.SerializeToString), 'GetInventory': grpc.unary_unary_rpc_method_handler(servicer.GetInventory, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInventoryRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInventoryResponse.SerializeToString), 'SetInventory': grpc.unary_unary_rpc_method_handler(servicer.SetInventory, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInventoryRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInventoryResponse.SerializeToString), 'GetPrice': grpc.unary_unary_rpc_method_handler(servicer.GetPrice, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceResponse.SerializeToString), 'SetPrice': grpc.unary_unary_rpc_method_handler(servicer.SetPrice, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceResponse.SerializeToString), 'GetCurrency': grpc.unary_unary_rpc_method_handler(servicer.GetCurrency, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyResponse.SerializeToString), 'SetCurrency': grpc.unary_unary_rpc_method_handler(servicer.SetCurrency, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyResponse.SerializeToString), 'GetInterestRate': grpc.unary_unary_rpc_method_handler(servicer.GetInterestRate, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateResponse.SerializeToString), 'SetInterestRate': grpc.unary_unary_rpc_method_handler(servicer.SetInterestRate, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateResponse.SerializeToString), 'GetBracketCutoffs': grpc.unary_unary_rpc_method_handler(servicer.GetBracketCutoffs, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketCutoffsResponse.SerializeToString), 'SetBracketCutoffs': grpc.unary_unary_rpc_method_handler(servicer.SetBracketCutoffs, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketCutoffsResponse.SerializeToString), 'GetBracketRates': grpc.unary_unary_rpc_method_handler(servicer.GetBracketRates, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.GetBracketRatesResponse.SerializeToString), 'SetBracketRates': grpc.unary_unary_rpc_method_handler(servicer.SetBracketRates, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.SetBracketRatesResponse.SerializeToString), 'CalculateTaxesDue': grpc.unary_unary_rpc_method_handler(servicer.CalculateTaxesDue, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateTaxesDueResponse.SerializeToString), 'CalculateConsumption': grpc.unary_unary_rpc_method_handler(servicer.CalculateConsumption, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateConsumptionResponse.SerializeToString), 'CalculateInterest': grpc.unary_unary_rpc_method_handler(servicer.CalculateInterest, request_deserializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestRequest.FromString, response_serializer=city_dot_economy_dot_v2_dot_org__service__pb2.CalculateInterestResponse.SerializeToString)}
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
    def GetPrice(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetPrice', city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetPriceResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetPrice(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetPrice', city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetPriceResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCurrency(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetCurrency', city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetCurrencyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetCurrency(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetCurrency', city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetCurrencyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetInterestRate(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/GetInterestRate', city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.GetInterestRateResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetInterestRate(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v2.OrgService/SetInterestRate', city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateRequest.SerializeToString, city_dot_economy_dot_v2_dot_org__service__pb2.SetInterestRateResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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