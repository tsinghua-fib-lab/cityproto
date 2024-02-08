"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.map.v2 import road_service_pb2 as city_dot_map_dot_v2_dot_road__service__pb2

class RoadServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetRoad = channel.unary_unary('/city.map.v2.RoadService/GetRoad', request_serializer=city_dot_map_dot_v2_dot_road__service__pb2.GetRoadRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_road__service__pb2.GetRoadResponse.FromString)
        self.GetRuinInfo = channel.unary_unary('/city.map.v2.RoadService/GetRuinInfo', request_serializer=city_dot_map_dot_v2_dot_road__service__pb2.GetRuinInfoRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_road__service__pb2.GetRuinInfoResponse.FromString)
        self.GetEvents = channel.unary_unary('/city.map.v2.RoadService/GetEvents', request_serializer=city_dot_map_dot_v2_dot_road__service__pb2.GetEventsRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_road__service__pb2.GetEventsResponse.FromString)

class RoadServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetRoad(self, request, context):
        """查询道路信息
        Get road information
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRuinInfo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetEvents(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_RoadServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetRoad': grpc.unary_unary_rpc_method_handler(servicer.GetRoad, request_deserializer=city_dot_map_dot_v2_dot_road__service__pb2.GetRoadRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_road__service__pb2.GetRoadResponse.SerializeToString), 'GetRuinInfo': grpc.unary_unary_rpc_method_handler(servicer.GetRuinInfo, request_deserializer=city_dot_map_dot_v2_dot_road__service__pb2.GetRuinInfoRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_road__service__pb2.GetRuinInfoResponse.SerializeToString), 'GetEvents': grpc.unary_unary_rpc_method_handler(servicer.GetEvents, request_deserializer=city_dot_map_dot_v2_dot_road__service__pb2.GetEventsRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_road__service__pb2.GetEventsResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.map.v2.RoadService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class RoadService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetRoad(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.RoadService/GetRoad', city_dot_map_dot_v2_dot_road__service__pb2.GetRoadRequest.SerializeToString, city_dot_map_dot_v2_dot_road__service__pb2.GetRoadResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetRuinInfo(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.RoadService/GetRuinInfo', city_dot_map_dot_v2_dot_road__service__pb2.GetRuinInfoRequest.SerializeToString, city_dot_map_dot_v2_dot_road__service__pb2.GetRuinInfoResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetEvents(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.RoadService/GetEvents', city_dot_map_dot_v2_dot_road__service__pb2.GetEventsRequest.SerializeToString, city_dot_map_dot_v2_dot_road__service__pb2.GetEventsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)