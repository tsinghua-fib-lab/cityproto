"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ......city.comm.interaction.demand.v1 import demand_service_pb2 as city_dot_comm_dot_interaction_dot_demand_dot_v1_dot_demand__service__pb2

class DemandServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SetDemandStatus = channel.unary_unary('/city.comm.interaction.demand.v1.DemandService/SetDemandStatus', request_serializer=city_dot_comm_dot_interaction_dot_demand_dot_v1_dot_demand__service__pb2.SetDemandStatusRequest.SerializeToString, response_deserializer=city_dot_comm_dot_interaction_dot_demand_dot_v1_dot_demand__service__pb2.SetDemandStatusResponse.FromString)

class DemandServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SetDemandStatus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_DemandServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'SetDemandStatus': grpc.unary_unary_rpc_method_handler(servicer.SetDemandStatus, request_deserializer=city_dot_comm_dot_interaction_dot_demand_dot_v1_dot_demand__service__pb2.SetDemandStatusRequest.FromString, response_serializer=city_dot_comm_dot_interaction_dot_demand_dot_v1_dot_demand__service__pb2.SetDemandStatusResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.comm.interaction.demand.v1.DemandService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class DemandService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SetDemandStatus(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.comm.interaction.demand.v1.DemandService/SetDemandStatus', city_dot_comm_dot_interaction_dot_demand_dot_v1_dot_demand__service__pb2.SetDemandStatusRequest.SerializeToString, city_dot_comm_dot_interaction_dot_demand_dot_v1_dot_demand__service__pb2.SetDemandStatusResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)