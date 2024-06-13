"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.map.v2 import lane_service_pb2 as city_dot_map_dot_v2_dot_lane__service__pb2

class LaneServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SetLaneMaxV = channel.unary_unary('/city.map.v2.LaneService/SetLaneMaxV', request_serializer=city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneMaxVRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneMaxVResponse.FromString)
        self.SetLaneRestriction = channel.unary_unary('/city.map.v2.LaneService/SetLaneRestriction', request_serializer=city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneRestrictionRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneRestrictionResponse.FromString)
        self.GetLane = channel.unary_unary('/city.map.v2.LaneService/GetLane', request_serializer=city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneResponse.FromString)
        self.GetLaneByLongLatBBox = channel.unary_unary('/city.map.v2.LaneService/GetLaneByLongLatBBox', request_serializer=city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneByLongLatBBoxRequest.SerializeToString, response_deserializer=city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneByLongLatBBoxResponse.FromString)

class LaneServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SetLaneMaxV(self, request, context):
        """设置Lane的最大速度（限速）
        Set Lane's maximum speed (speed limit)
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetLaneRestriction(self, request, context):
        """设置Lane限行
        Set Lane's traffic restriction
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetLane(self, request, context):
        """获取Lane的信息
        Get Lane information
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetLaneByLongLatBBox(self, request, context):
        """获取特定区域内的Lane的信息
        Get Lane information in a specific region
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_LaneServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'SetLaneMaxV': grpc.unary_unary_rpc_method_handler(servicer.SetLaneMaxV, request_deserializer=city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneMaxVRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneMaxVResponse.SerializeToString), 'SetLaneRestriction': grpc.unary_unary_rpc_method_handler(servicer.SetLaneRestriction, request_deserializer=city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneRestrictionRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneRestrictionResponse.SerializeToString), 'GetLane': grpc.unary_unary_rpc_method_handler(servicer.GetLane, request_deserializer=city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneResponse.SerializeToString), 'GetLaneByLongLatBBox': grpc.unary_unary_rpc_method_handler(servicer.GetLaneByLongLatBBox, request_deserializer=city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneByLongLatBBoxRequest.FromString, response_serializer=city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneByLongLatBBoxResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.map.v2.LaneService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class LaneService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SetLaneMaxV(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.LaneService/SetLaneMaxV', city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneMaxVRequest.SerializeToString, city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneMaxVResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetLaneRestriction(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.LaneService/SetLaneRestriction', city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneRestrictionRequest.SerializeToString, city_dot_map_dot_v2_dot_lane__service__pb2.SetLaneRestrictionResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetLane(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.LaneService/GetLane', city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneRequest.SerializeToString, city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetLaneByLongLatBBox(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.map.v2.LaneService/GetLaneByLongLatBBox', city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneByLongLatBBoxRequest.SerializeToString, city_dot_map_dot_v2_dot_lane__service__pb2.GetLaneByLongLatBBoxResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)