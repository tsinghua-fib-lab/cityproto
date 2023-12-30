"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ......city.traffic.interaction.lane.v1 import lane_service_pb2 as city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2

class LaneServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SetMaxV = channel.unary_unary('/city.traffic.interaction.lane.v1.LaneService/SetMaxV', request_serializer=city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.SetMaxVRequest.SerializeToString, response_deserializer=city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.SetMaxVResponse.FromString)
        self.GetLane = channel.unary_unary('/city.traffic.interaction.lane.v1.LaneService/GetLane', request_serializer=city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.GetLaneRequest.SerializeToString, response_deserializer=city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.GetLaneResponse.FromString)

class LaneServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SetMaxV(self, request, context):
        """设置Lane的最大速度（限速）
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetLane(self, request, context):
        """获取Lane的信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_LaneServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'SetMaxV': grpc.unary_unary_rpc_method_handler(servicer.SetMaxV, request_deserializer=city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.SetMaxVRequest.FromString, response_serializer=city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.SetMaxVResponse.SerializeToString), 'GetLane': grpc.unary_unary_rpc_method_handler(servicer.GetLane, request_deserializer=city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.GetLaneRequest.FromString, response_serializer=city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.GetLaneResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.traffic.interaction.lane.v1.LaneService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class LaneService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SetMaxV(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.traffic.interaction.lane.v1.LaneService/SetMaxV', city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.SetMaxVRequest.SerializeToString, city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.SetMaxVResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetLane(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.traffic.interaction.lane.v1.LaneService/GetLane', city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.GetLaneRequest.SerializeToString, city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__service__pb2.GetLaneResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)