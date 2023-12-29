"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from .....city.elec.interaction.v1 import elec_service_pb2 as city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2

class ElecServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SetStatus = channel.unary_unary('/city.elec.interaction.v1.ElecService/SetStatus', request_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.SetStatusRequest.SerializeToString, response_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.SetStatusResponse.FromString)
        self.GetPower = channel.unary_unary('/city.elec.interaction.v1.ElecService/GetPower', request_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerRequest.SerializeToString, response_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerResponse.FromString)
        self.GetPowerStatus = channel.unary_unary('/city.elec.interaction.v1.ElecService/GetPowerStatus', request_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerStatusRequest.SerializeToString, response_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerStatusResponse.FromString)
        self.GetNoPowerAOI = channel.unary_unary('/city.elec.interaction.v1.ElecService/GetNoPowerAOI', request_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetNoPowerAOIRequest.SerializeToString, response_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetNoPowerAOIResponse.FromString)
        self.GetRuinInfo = channel.unary_unary('/city.elec.interaction.v1.ElecService/GetRuinInfo', request_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetRuinInfoRequest.SerializeToString, response_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetRuinInfoResponse.FromString)
        self.GetEdgeStatus = channel.unary_unary('/city.elec.interaction.v1.ElecService/GetEdgeStatus', request_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetEdgeStatusRequest.SerializeToString, response_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetEdgeStatusResponse.FromString)

class ElecServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SetStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPower(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPowerStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetNoPowerAOI(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRuinInfo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetEdgeStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_ElecServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'SetStatus': grpc.unary_unary_rpc_method_handler(servicer.SetStatus, request_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.SetStatusRequest.FromString, response_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.SetStatusResponse.SerializeToString), 'GetPower': grpc.unary_unary_rpc_method_handler(servicer.GetPower, request_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerRequest.FromString, response_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerResponse.SerializeToString), 'GetPowerStatus': grpc.unary_unary_rpc_method_handler(servicer.GetPowerStatus, request_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerStatusRequest.FromString, response_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerStatusResponse.SerializeToString), 'GetNoPowerAOI': grpc.unary_unary_rpc_method_handler(servicer.GetNoPowerAOI, request_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetNoPowerAOIRequest.FromString, response_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetNoPowerAOIResponse.SerializeToString), 'GetRuinInfo': grpc.unary_unary_rpc_method_handler(servicer.GetRuinInfo, request_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetRuinInfoRequest.FromString, response_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetRuinInfoResponse.SerializeToString), 'GetEdgeStatus': grpc.unary_unary_rpc_method_handler(servicer.GetEdgeStatus, request_deserializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetEdgeStatusRequest.FromString, response_serializer=city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetEdgeStatusResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.elec.interaction.v1.ElecService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class ElecService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SetStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.elec.interaction.v1.ElecService/SetStatus', city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.SetStatusRequest.SerializeToString, city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.SetStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPower(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.elec.interaction.v1.ElecService/GetPower', city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerRequest.SerializeToString, city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPowerStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.elec.interaction.v1.ElecService/GetPowerStatus', city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerStatusRequest.SerializeToString, city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetPowerStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetNoPowerAOI(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.elec.interaction.v1.ElecService/GetNoPowerAOI', city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetNoPowerAOIRequest.SerializeToString, city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetNoPowerAOIResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetRuinInfo(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.elec.interaction.v1.ElecService/GetRuinInfo', city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetRuinInfoRequest.SerializeToString, city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetRuinInfoResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetEdgeStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.elec.interaction.v1.ElecService/GetEdgeStatus', city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetEdgeStatusRequest.SerializeToString, city_dot_elec_dot_interaction_dot_v1_dot_elec__service__pb2.GetEdgeStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)