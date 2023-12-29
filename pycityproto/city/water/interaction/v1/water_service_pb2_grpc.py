"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from .....city.water.interaction.v1 import water_service_pb2 as city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2

class WaterServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SetPumpPowerStatus = channel.unary_unary('/city.water.interaction.v1.WaterService/SetPumpPowerStatus', request_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpPowerStatusRequest.SerializeToString, response_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpPowerStatusResponse.FromString)
        self.SetPumpNetworkStatus = channel.unary_unary('/city.water.interaction.v1.WaterService/SetPumpNetworkStatus', request_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpNetworkStatusRequest.SerializeToString, response_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpNetworkStatusResponse.FromString)
        self.SetPumpStatus = channel.unary_unary('/city.water.interaction.v1.WaterService/SetPumpStatus', request_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpStatusRequest.SerializeToString, response_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpStatusResponse.FromString)
        self.GetPumpStatus = channel.unary_unary('/city.water.interaction.v1.WaterService/GetPumpStatus', request_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetPumpStatusRequest.SerializeToString, response_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetPumpStatusResponse.FromString)
        self.GetNoWaterAOI = channel.unary_unary('/city.water.interaction.v1.WaterService/GetNoWaterAOI', request_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetNoWaterAOIRequest.SerializeToString, response_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetNoWaterAOIResponse.FromString)
        self.GetRuinInfo = channel.unary_unary('/city.water.interaction.v1.WaterService/GetRuinInfo', request_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetRuinInfoRequest.SerializeToString, response_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetRuinInfoResponse.FromString)

class WaterServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SetPumpPowerStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetPumpNetworkStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetPumpStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPumpStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetNoWaterAOI(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRuinInfo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_WaterServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'SetPumpPowerStatus': grpc.unary_unary_rpc_method_handler(servicer.SetPumpPowerStatus, request_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpPowerStatusRequest.FromString, response_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpPowerStatusResponse.SerializeToString), 'SetPumpNetworkStatus': grpc.unary_unary_rpc_method_handler(servicer.SetPumpNetworkStatus, request_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpNetworkStatusRequest.FromString, response_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpNetworkStatusResponse.SerializeToString), 'SetPumpStatus': grpc.unary_unary_rpc_method_handler(servicer.SetPumpStatus, request_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpStatusRequest.FromString, response_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpStatusResponse.SerializeToString), 'GetPumpStatus': grpc.unary_unary_rpc_method_handler(servicer.GetPumpStatus, request_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetPumpStatusRequest.FromString, response_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetPumpStatusResponse.SerializeToString), 'GetNoWaterAOI': grpc.unary_unary_rpc_method_handler(servicer.GetNoWaterAOI, request_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetNoWaterAOIRequest.FromString, response_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetNoWaterAOIResponse.SerializeToString), 'GetRuinInfo': grpc.unary_unary_rpc_method_handler(servicer.GetRuinInfo, request_deserializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetRuinInfoRequest.FromString, response_serializer=city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetRuinInfoResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.water.interaction.v1.WaterService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class WaterService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SetPumpPowerStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.water.interaction.v1.WaterService/SetPumpPowerStatus', city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpPowerStatusRequest.SerializeToString, city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpPowerStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetPumpNetworkStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.water.interaction.v1.WaterService/SetPumpNetworkStatus', city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpNetworkStatusRequest.SerializeToString, city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpNetworkStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetPumpStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.water.interaction.v1.WaterService/SetPumpStatus', city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpStatusRequest.SerializeToString, city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.SetPumpStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPumpStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.water.interaction.v1.WaterService/GetPumpStatus', city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetPumpStatusRequest.SerializeToString, city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetPumpStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetNoWaterAOI(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.water.interaction.v1.WaterService/GetNoWaterAOI', city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetNoWaterAOIRequest.SerializeToString, city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetNoWaterAOIResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetRuinInfo(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.water.interaction.v1.WaterService/GetRuinInfo', city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetRuinInfoRequest.SerializeToString, city_dot_water_dot_interaction_dot_v1_dot_water__service__pb2.GetRuinInfoResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)