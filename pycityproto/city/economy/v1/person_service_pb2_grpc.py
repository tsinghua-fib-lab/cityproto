"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.economy.v1 import person_service_pb2 as city_dot_economy_dot_v1_dot_person__service__pb2

class PersonServiceStub(object):
    """个人经济情况接口
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetPerson = channel.unary_unary('/city.economy.v1.PersonService/GetPerson', request_serializer=city_dot_economy_dot_v1_dot_person__service__pb2.GetPersonRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v1_dot_person__service__pb2.GetPersonResponse.FromString)
        self.UpdatePersonMoney = channel.unary_unary('/city.economy.v1.PersonService/UpdatePersonMoney', request_serializer=city_dot_economy_dot_v1_dot_person__service__pb2.UpdatePersonMoneyRequest.SerializeToString, response_deserializer=city_dot_economy_dot_v1_dot_person__service__pb2.UpdatePersonMoneyResponse.FromString)

class PersonServiceServicer(object):
    """个人经济情况接口
    """

    def GetPerson(self, request, context):
        """批量查询人的经济情况（资金、雇佣关系）
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdatePersonMoney(self, request, context):
        """批量修改人的资金
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_PersonServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetPerson': grpc.unary_unary_rpc_method_handler(servicer.GetPerson, request_deserializer=city_dot_economy_dot_v1_dot_person__service__pb2.GetPersonRequest.FromString, response_serializer=city_dot_economy_dot_v1_dot_person__service__pb2.GetPersonResponse.SerializeToString), 'UpdatePersonMoney': grpc.unary_unary_rpc_method_handler(servicer.UpdatePersonMoney, request_deserializer=city_dot_economy_dot_v1_dot_person__service__pb2.UpdatePersonMoneyRequest.FromString, response_serializer=city_dot_economy_dot_v1_dot_person__service__pb2.UpdatePersonMoneyResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.economy.v1.PersonService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class PersonService(object):
    """个人经济情况接口
    """

    @staticmethod
    def GetPerson(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v1.PersonService/GetPerson', city_dot_economy_dot_v1_dot_person__service__pb2.GetPersonRequest.SerializeToString, city_dot_economy_dot_v1_dot_person__service__pb2.GetPersonResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdatePersonMoney(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.economy.v1.PersonService/UpdatePersonMoney', city_dot_economy_dot_v1_dot_person__service__pb2.UpdatePersonMoneyRequest.SerializeToString, city_dot_economy_dot_v1_dot_person__service__pb2.UpdatePersonMoneyResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)