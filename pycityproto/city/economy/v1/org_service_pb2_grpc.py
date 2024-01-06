"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.economy.v1 import org_service_pb2 as city_dot_economy_dot_v1_dot_org__service__pb2

class OrgServiceStub(object):
    """组织经济情况接口
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetOrg = channel.unary_unary('/city.economy.v1.OrgService/GetOrg', request_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.GetOrgRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.GetOrgResponse.FromString)
        self.UpdateOrgMoney = channel.unary_unary('/city.economy.v1.OrgService/UpdateOrgMoney', request_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgMoneyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgMoneyResponse.FromString)
        self.UpdateOrgGoods = channel.unary_unary('/city.economy.v1.OrgService/UpdateOrgGoods', request_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgGoodsRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgGoodsResponse.FromString)
        self.UpdateOrgEmployee = channel.unary_unary('/city.economy.v1.OrgService/UpdateOrgEmployee', request_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgEmployeeRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgEmployeeResponse.FromString)
        self.UpdateOrgJob = channel.unary_unary('/city.economy.v1.OrgService/UpdateOrgJob', request_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgJobRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgJobResponse.FromString)

class OrgServiceServicer(object):
    """组织经济情况接口
    """

    def GetOrg(self, request, context):
        """批量查询组织的经济情况（员工、岗位、资金、货物）
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateOrgMoney(self, request, context):
        """批量修改组织的资金
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateOrgGoods(self, request, context):
        """批量修改组织的货物
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateOrgEmployee(self, request, context):
        """批量修改组织的员工
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateOrgJob(self, request, context):
        """批量修改组织的岗位
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_OrgServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetOrg': grpc.unary_unary_rpc_method_handler(servicer.GetOrg, request_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.GetOrgRequest.FromString, response_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.GetOrgResponse.SerializeToString), 'UpdateOrgMoney': grpc.unary_unary_rpc_method_handler(servicer.UpdateOrgMoney, request_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgMoneyRequest.FromString, response_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgMoneyResponse.SerializeToString), 'UpdateOrgGoods': grpc.unary_unary_rpc_method_handler(servicer.UpdateOrgGoods, request_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgGoodsRequest.FromString, response_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgGoodsResponse.SerializeToString), 'UpdateOrgEmployee': grpc.unary_unary_rpc_method_handler(servicer.UpdateOrgEmployee, request_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgEmployeeRequest.FromString, response_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgEmployeeResponse.SerializeToString), 'UpdateOrgJob': grpc.unary_unary_rpc_method_handler(servicer.UpdateOrgJob, request_deserializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgJobRequest.FromString, response_serializer=city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgJobResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.economy.v1.OrgService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class OrgService(object):
    """组织经济情况接口
    """

    @staticmethod
    def GetOrg(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v1.OrgService/GetOrg', city_dot_economy_dot_v1_dot_org__service__pb2.GetOrgRequest.SerializeToString, city_dot_economy_dot_v1_dot_org__service__pb2.GetOrgResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateOrgMoney(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v1.OrgService/UpdateOrgMoney', city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgMoneyRequest.SerializeToString, city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgMoneyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateOrgGoods(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v1.OrgService/UpdateOrgGoods', city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgGoodsRequest.SerializeToString, city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgGoodsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateOrgEmployee(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v1.OrgService/UpdateOrgEmployee', city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgEmployeeRequest.SerializeToString, city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgEmployeeResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateOrgJob(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v1.OrgService/UpdateOrgJob', city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgJobRequest.SerializeToString, city_dot_economy_dot_v1_dot_org__service__pb2.UpdateOrgJobResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)