"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.person.v1 import person_service_pb2 as city_dot_person_dot_v1_dot_person__service__pb2

class PersonServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetPerson = channel.unary_unary('/city.person.v1.PersonService/GetPerson', request_serializer=city_dot_person_dot_v1_dot_person__service__pb2.GetPersonRequest.SerializeToString, response_deserializer=city_dot_person_dot_v1_dot_person__service__pb2.GetPersonResponse.FromString)
        self.AddPerson = channel.unary_unary('/city.person.v1.PersonService/AddPerson', request_serializer=city_dot_person_dot_v1_dot_person__service__pb2.AddPersonRequest.SerializeToString, response_deserializer=city_dot_person_dot_v1_dot_person__service__pb2.AddPersonResponse.FromString)
        self.SetSchedule = channel.unary_unary('/city.person.v1.PersonService/SetSchedule', request_serializer=city_dot_person_dot_v1_dot_person__service__pb2.SetScheduleRequest.SerializeToString, response_deserializer=city_dot_person_dot_v1_dot_person__service__pb2.SetScheduleResponse.FromString)
        self.GetPersonByLongLatBBox = channel.unary_unary('/city.person.v1.PersonService/GetPersonByLongLatBBox', request_serializer=city_dot_person_dot_v1_dot_person__service__pb2.GetPersonByLongLatBBoxRequest.SerializeToString, response_deserializer=city_dot_person_dot_v1_dot_person__service__pb2.GetPersonByLongLatBBoxResponse.FromString)

class PersonServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetPerson(self, request, context):
        """获取person信息
        Get person information
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddPerson(self, request, context):
        """新增person 传入person初始位置、目的地表、属性 返回personid
        Add a new person. Input person's initial location, destination table, and attributes, return personid
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetSchedule(self, request, context):
        """修改person的schedule 传入personid、目的地表
        Set person's schedule. Input personid and destination table
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPersonByLongLatBBox(self, request, context):
        """获取特定区域内的person
        Get persons in a specific region
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_PersonServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetPerson': grpc.unary_unary_rpc_method_handler(servicer.GetPerson, request_deserializer=city_dot_person_dot_v1_dot_person__service__pb2.GetPersonRequest.FromString, response_serializer=city_dot_person_dot_v1_dot_person__service__pb2.GetPersonResponse.SerializeToString), 'AddPerson': grpc.unary_unary_rpc_method_handler(servicer.AddPerson, request_deserializer=city_dot_person_dot_v1_dot_person__service__pb2.AddPersonRequest.FromString, response_serializer=city_dot_person_dot_v1_dot_person__service__pb2.AddPersonResponse.SerializeToString), 'SetSchedule': grpc.unary_unary_rpc_method_handler(servicer.SetSchedule, request_deserializer=city_dot_person_dot_v1_dot_person__service__pb2.SetScheduleRequest.FromString, response_serializer=city_dot_person_dot_v1_dot_person__service__pb2.SetScheduleResponse.SerializeToString), 'GetPersonByLongLatBBox': grpc.unary_unary_rpc_method_handler(servicer.GetPersonByLongLatBBox, request_deserializer=city_dot_person_dot_v1_dot_person__service__pb2.GetPersonByLongLatBBoxRequest.FromString, response_serializer=city_dot_person_dot_v1_dot_person__service__pb2.GetPersonByLongLatBBoxResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.person.v1.PersonService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class PersonService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetPerson(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v1.PersonService/GetPerson', city_dot_person_dot_v1_dot_person__service__pb2.GetPersonRequest.SerializeToString, city_dot_person_dot_v1_dot_person__service__pb2.GetPersonResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddPerson(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v1.PersonService/AddPerson', city_dot_person_dot_v1_dot_person__service__pb2.AddPersonRequest.SerializeToString, city_dot_person_dot_v1_dot_person__service__pb2.AddPersonResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetSchedule(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v1.PersonService/SetSchedule', city_dot_person_dot_v1_dot_person__service__pb2.SetScheduleRequest.SerializeToString, city_dot_person_dot_v1_dot_person__service__pb2.SetScheduleResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPersonByLongLatBBox(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v1.PersonService/GetPersonByLongLatBBox', city_dot_person_dot_v1_dot_person__service__pb2.GetPersonByLongLatBBoxRequest.SerializeToString, city_dot_person_dot_v1_dot_person__service__pb2.GetPersonByLongLatBBoxResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)