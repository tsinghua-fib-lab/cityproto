"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.map.v2 import traffic_light_service_pb2 as city_dot_map_dot_v2_dot_traffic__light__service__pb2

class TrafficLightServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetTrafficLight = channel.unary_unary('/city.map.v2.TrafficLightService/GetTrafficLight', request_serializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.GetTrafficLightRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.GetTrafficLightResponse.FromString)
        self.SetTrafficLight = channel.unary_unary('/city.map.v2.TrafficLightService/SetTrafficLight', request_serializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightResponse.FromString)
        self.SetTrafficLightPhase = channel.unary_unary('/city.map.v2.TrafficLightService/SetTrafficLightPhase', request_serializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightPhaseRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightPhaseResponse.FromString)
        self.SetTrafficLightStatus = channel.unary_unary('/city.map.v2.TrafficLightService/SetTrafficLightStatus', request_serializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightStatusRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightStatusResponse.FromString)

class TrafficLightServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetTrafficLight(self, request, context):
        """获取路口的红绿灯信息
        Get traffic light information
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetTrafficLight(self, request, context):
        """设置路口的红绿灯信息
        Set traffic light information
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetTrafficLightPhase(self, request, context):
        """设置路口的红绿灯相位
        Set traffic light phase
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetTrafficLightStatus(self, request, context):
        """设置路口的红绿灯状态
        Set traffic light status
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_TrafficLightServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetTrafficLight': grpc.unary_unary_rpc_method_handler(servicer.GetTrafficLight, request_deserializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.GetTrafficLightRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.GetTrafficLightResponse.SerializeToString), 'SetTrafficLight': grpc.unary_unary_rpc_method_handler(servicer.SetTrafficLight, request_deserializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightResponse.SerializeToString), 'SetTrafficLightPhase': grpc.unary_unary_rpc_method_handler(servicer.SetTrafficLightPhase, request_deserializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightPhaseRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightPhaseResponse.SerializeToString), 'SetTrafficLightStatus': grpc.unary_unary_rpc_method_handler(servicer.SetTrafficLightStatus, request_deserializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightStatusRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightStatusResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.map.v2.TrafficLightService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class TrafficLightService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetTrafficLight(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.TrafficLightService/GetTrafficLight', city_dot_map_dot_v2_dot_traffic__light__service__pb2.GetTrafficLightRequest.SerializeToString, city_dot_map_dot_v2_dot_traffic__light__service__pb2.GetTrafficLightResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetTrafficLight(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.TrafficLightService/SetTrafficLight', city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightRequest.SerializeToString, city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetTrafficLightPhase(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.TrafficLightService/SetTrafficLightPhase', city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightPhaseRequest.SerializeToString, city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightPhaseResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetTrafficLightStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.TrafficLightService/SetTrafficLightStatus', city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightStatusRequest.SerializeToString, city_dot_map_dot_v2_dot_traffic__light__service__pb2.SetTrafficLightStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)