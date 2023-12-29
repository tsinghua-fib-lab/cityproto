"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.wargame.v1 import wargame_service_pb2 as city_dot_wargame_dot_v1_dot_wargame__service__pb2

class WarGameServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.PickPoints = channel.unary_unary('/city.wargame.v1.WarGameService/PickPoints', request_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.PickPointsRequest.SerializeToString, response_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.PickPointsResponse.FromString)
        self.GetPickPoints = channel.unary_unary('/city.wargame.v1.WarGameService/GetPickPoints', request_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetPickPointsRequest.SerializeToString, response_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetPickPointsResponse.FromString)
        self.GiveDefenseOrder = channel.unary_unary('/city.wargame.v1.WarGameService/GiveDefenseOrder', request_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GiveDefenseOrderRequest.SerializeToString, response_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GiveDefenseOrderResponse.FromString)
        self.SetScoreWeight = channel.unary_unary('/city.wargame.v1.WarGameService/SetScoreWeight', request_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.SetScoreWeightRequest.SerializeToString, response_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.SetScoreWeightResponse.FromString)
        self.GetHitHistory = channel.unary_unary('/city.wargame.v1.WarGameService/GetHitHistory', request_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetHitHistoryRequest.SerializeToString, response_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetHitHistoryResponse.FromString)
        self.GetRecoPoints = channel.unary_unary('/city.wargame.v1.WarGameService/GetRecoPoints', request_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetRecoPointsRequest.SerializeToString, response_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetRecoPointsResponse.FromString)
        self.GetStep = channel.unary_unary('/city.wargame.v1.WarGameService/GetStep', request_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetStepRequest.SerializeToString, response_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetStepResponse.FromString)
        self.GetCasualties = channel.unary_unary('/city.wargame.v1.WarGameService/GetCasualties', request_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetCasualtiesRequest.SerializeToString, response_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetCasualtiesResponse.FromString)

class WarGameServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def PickPoints(self, request, context):
        """地图选点
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPickPoints(self, request, context):
        """获取当前轮选点
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GiveDefenseOrder(self, request, context):
        """指定防守方条令
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetScoreWeight(self, request, context):
        """分值权重设定
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetHitHistory(self, request, context):
        """打击历史获取
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRecoPoints(self, request, context):
        """推荐选点获取
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetStep(self, request, context):
        """当前步与状态获取
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCasualties(self, request, context):
        """人口损伤人数和aoi的id获取
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_WarGameServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'PickPoints': grpc.unary_unary_rpc_method_handler(servicer.PickPoints, request_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.PickPointsRequest.FromString, response_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.PickPointsResponse.SerializeToString), 'GetPickPoints': grpc.unary_unary_rpc_method_handler(servicer.GetPickPoints, request_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetPickPointsRequest.FromString, response_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetPickPointsResponse.SerializeToString), 'GiveDefenseOrder': grpc.unary_unary_rpc_method_handler(servicer.GiveDefenseOrder, request_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GiveDefenseOrderRequest.FromString, response_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GiveDefenseOrderResponse.SerializeToString), 'SetScoreWeight': grpc.unary_unary_rpc_method_handler(servicer.SetScoreWeight, request_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.SetScoreWeightRequest.FromString, response_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.SetScoreWeightResponse.SerializeToString), 'GetHitHistory': grpc.unary_unary_rpc_method_handler(servicer.GetHitHistory, request_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetHitHistoryRequest.FromString, response_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetHitHistoryResponse.SerializeToString), 'GetRecoPoints': grpc.unary_unary_rpc_method_handler(servicer.GetRecoPoints, request_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetRecoPointsRequest.FromString, response_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetRecoPointsResponse.SerializeToString), 'GetStep': grpc.unary_unary_rpc_method_handler(servicer.GetStep, request_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetStepRequest.FromString, response_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetStepResponse.SerializeToString), 'GetCasualties': grpc.unary_unary_rpc_method_handler(servicer.GetCasualties, request_deserializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetCasualtiesRequest.FromString, response_serializer=city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetCasualtiesResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.wargame.v1.WarGameService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class WarGameService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def PickPoints(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.wargame.v1.WarGameService/PickPoints', city_dot_wargame_dot_v1_dot_wargame__service__pb2.PickPointsRequest.SerializeToString, city_dot_wargame_dot_v1_dot_wargame__service__pb2.PickPointsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPickPoints(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.wargame.v1.WarGameService/GetPickPoints', city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetPickPointsRequest.SerializeToString, city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetPickPointsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GiveDefenseOrder(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.wargame.v1.WarGameService/GiveDefenseOrder', city_dot_wargame_dot_v1_dot_wargame__service__pb2.GiveDefenseOrderRequest.SerializeToString, city_dot_wargame_dot_v1_dot_wargame__service__pb2.GiveDefenseOrderResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetScoreWeight(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.wargame.v1.WarGameService/SetScoreWeight', city_dot_wargame_dot_v1_dot_wargame__service__pb2.SetScoreWeightRequest.SerializeToString, city_dot_wargame_dot_v1_dot_wargame__service__pb2.SetScoreWeightResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetHitHistory(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.wargame.v1.WarGameService/GetHitHistory', city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetHitHistoryRequest.SerializeToString, city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetHitHistoryResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetRecoPoints(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.wargame.v1.WarGameService/GetRecoPoints', city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetRecoPointsRequest.SerializeToString, city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetRecoPointsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetStep(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.wargame.v1.WarGameService/GetStep', city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetStepRequest.SerializeToString, city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetStepResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCasualties(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.wargame.v1.WarGameService/GetCasualties', city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetCasualtiesRequest.SerializeToString, city_dot_wargame_dot_v1_dot_wargame__service__pb2.GetCasualtiesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)