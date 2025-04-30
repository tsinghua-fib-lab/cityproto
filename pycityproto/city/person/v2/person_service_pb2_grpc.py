"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from ....city.person.v2 import person_service_pb2 as city_dot_person_dot_v2_dot_person__service__pb2

class PersonServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetPerson = channel.unary_unary('/city.person.v2.PersonService/GetPerson', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonResponse.FromString)
        self.AddPerson = channel.unary_unary('/city.person.v2.PersonService/AddPerson', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.AddPersonRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.AddPersonResponse.FromString)
        self.SetSchedule = channel.unary_unary('/city.person.v2.PersonService/SetSchedule', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetScheduleRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetScheduleResponse.FromString)
        self.GetPersons = channel.unary_unary('/city.person.v2.PersonService/GetPersons', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonsRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonsResponse.FromString)
        self.GetPersonByLongLatBBox = channel.unary_unary('/city.person.v2.PersonService/GetPersonByLongLatBBox', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonByLongLatBBoxRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonByLongLatBBoxResponse.FromString)
        self.GetAllVehicles = channel.unary_unary('/city.person.v2.PersonService/GetAllVehicles', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllVehiclesRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllVehiclesResponse.FromString)
        self.GetAllPedestrians = channel.unary_unary('/city.person.v2.PersonService/GetAllPedestrians', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllPedestriansRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllPedestriansResponse.FromString)
        self.ResetPersonPosition = channel.unary_unary('/city.person.v2.PersonService/ResetPersonPosition', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.ResetPersonPositionRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.ResetPersonPositionResponse.FromString)
        self.SetControlledVehicleIDs = channel.unary_unary('/city.person.v2.PersonService/SetControlledVehicleIDs', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleIDsRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleIDsResponse.FromString)
        self.FetchControlledVehicleEnvs = channel.unary_unary('/city.person.v2.PersonService/FetchControlledVehicleEnvs', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledVehicleEnvsRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledVehicleEnvsResponse.FromString)
        self.SetControlledVehicleActions = channel.unary_unary('/city.person.v2.PersonService/SetControlledVehicleActions', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleActionsRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleActionsResponse.FromString)
        self.SetControlledTaxiIDs = channel.unary_unary('/city.person.v2.PersonService/SetControlledTaxiIDs', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiIDsRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiIDsResponse.FromString)
        self.GetAllOrders = channel.unary_unary('/city.person.v2.PersonService/GetAllOrders', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllOrdersRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllOrdersResponse.FromString)
        self.SetControlledTaxiToOrders = channel.unary_unary('/city.person.v2.PersonService/SetControlledTaxiToOrders', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiToOrdersRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiToOrdersResponse.FromString)
        self.SetControlledPedestrians = channel.unary_unary('/city.person.v2.PersonService/SetControlledPedestrians', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansResponse.FromString)
        self.FetchControlledPedestriansEnvs = channel.unary_unary('/city.person.v2.PersonService/FetchControlledPedestriansEnvs', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledPedestriansEnvsRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledPedestriansEnvsResponse.FromString)
        self.SetControlledPedestriansActions = channel.unary_unary('/city.person.v2.PersonService/SetControlledPedestriansActions', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansActionsRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansActionsResponse.FromString)
        self.GetControlledTaxiOrderAllocationPlan = channel.unary_unary('/city.person.v2.PersonService/GetControlledTaxiOrderAllocationPlan', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetControlledTaxiOrderAllocationPlanRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetControlledTaxiOrderAllocationPlanResponse.FromString)
        self.SetControlledTaxiOrderAllocationPlan = channel.unary_unary('/city.person.v2.PersonService/SetControlledTaxiOrderAllocationPlan', request_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiOrderAllocationPlanRequest.SerializeToString, response_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiOrderAllocationPlanResponse.FromString)

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
        Add a new person. Input person's initial location, destination table, and
        attributes, return personid
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

    def GetPersons(self, request, context):
        """获取多个person信息
        Get information of multiple persons
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

    def GetAllVehicles(self, request, context):
        """获取所有车辆
        Get all vehicles
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllPedestrians(self, request, context):
        """获取所有行人
        Get all pedestrians
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ResetPersonPosition(self, request, context):
        """重置人的位置（将停止当前正在进行的出行，转为sleep状态）
        Reset person's position (stop the current trip and switch to sleep status)
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetControlledVehicleIDs(self, request, context):
        """RL接口
        RL interface

        设置由外部控制行为的vehicle
        Set vehicle controlled by external behavior
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def FetchControlledVehicleEnvs(self, request, context):
        """获取由外部控制行为的vehicle信息
        Get information of vehicle controlled by external behavior
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetControlledVehicleActions(self, request, context):
        """设置由外部控制行为的vehicle的行为
        Set behavior of vehicle controlled by external behavior
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetControlledTaxiIDs(self, request, context):
        """设置由外部控制的taxi
        Set taxi controlled by external behavior
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllOrders(self, request, context):
        """获取所有订单信息
        Get information of all orders
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetControlledTaxiToOrders(self, request, context):
        """设置所有外部控制的出租车接指定的单
        Set all externally controlled taxis to specified orders
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetControlledPedestrians(self, request, context):
        """设置由外部控制的行人
        Set pedestrian controlled by external behavior
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def FetchControlledPedestriansEnvs(self, request, context):
        """获取由外部控制的行人信息
        Get information of pedestrian controlled by external behavior
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetControlledPedestriansActions(self, request, context):
        """设置由外部控制的行人行为
        Set behavior of pedestrian controlled by external behavior
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetControlledTaxiOrderAllocationPlan(self, request, context):
        """获取当前所有受控出租车的订单分配方案
        Get current order allocation plan for all controlled taxis
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetControlledTaxiOrderAllocationPlan(self, request, context):
        """设置当前所有受控出租车的订单分配方案
        Set current order allocation plan for all controlled taxis
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_PersonServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'GetPerson': grpc.unary_unary_rpc_method_handler(servicer.GetPerson, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonResponse.SerializeToString), 'AddPerson': grpc.unary_unary_rpc_method_handler(servicer.AddPerson, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.AddPersonRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.AddPersonResponse.SerializeToString), 'SetSchedule': grpc.unary_unary_rpc_method_handler(servicer.SetSchedule, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetScheduleRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetScheduleResponse.SerializeToString), 'GetPersons': grpc.unary_unary_rpc_method_handler(servicer.GetPersons, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonsRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonsResponse.SerializeToString), 'GetPersonByLongLatBBox': grpc.unary_unary_rpc_method_handler(servicer.GetPersonByLongLatBBox, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonByLongLatBBoxRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetPersonByLongLatBBoxResponse.SerializeToString), 'GetAllVehicles': grpc.unary_unary_rpc_method_handler(servicer.GetAllVehicles, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllVehiclesRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllVehiclesResponse.SerializeToString), 'GetAllPedestrians': grpc.unary_unary_rpc_method_handler(servicer.GetAllPedestrians, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllPedestriansRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllPedestriansResponse.SerializeToString), 'ResetPersonPosition': grpc.unary_unary_rpc_method_handler(servicer.ResetPersonPosition, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.ResetPersonPositionRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.ResetPersonPositionResponse.SerializeToString), 'SetControlledVehicleIDs': grpc.unary_unary_rpc_method_handler(servicer.SetControlledVehicleIDs, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleIDsRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleIDsResponse.SerializeToString), 'FetchControlledVehicleEnvs': grpc.unary_unary_rpc_method_handler(servicer.FetchControlledVehicleEnvs, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledVehicleEnvsRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledVehicleEnvsResponse.SerializeToString), 'SetControlledVehicleActions': grpc.unary_unary_rpc_method_handler(servicer.SetControlledVehicleActions, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleActionsRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleActionsResponse.SerializeToString), 'SetControlledTaxiIDs': grpc.unary_unary_rpc_method_handler(servicer.SetControlledTaxiIDs, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiIDsRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiIDsResponse.SerializeToString), 'GetAllOrders': grpc.unary_unary_rpc_method_handler(servicer.GetAllOrders, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllOrdersRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetAllOrdersResponse.SerializeToString), 'SetControlledTaxiToOrders': grpc.unary_unary_rpc_method_handler(servicer.SetControlledTaxiToOrders, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiToOrdersRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiToOrdersResponse.SerializeToString), 'SetControlledPedestrians': grpc.unary_unary_rpc_method_handler(servicer.SetControlledPedestrians, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansResponse.SerializeToString), 'FetchControlledPedestriansEnvs': grpc.unary_unary_rpc_method_handler(servicer.FetchControlledPedestriansEnvs, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledPedestriansEnvsRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledPedestriansEnvsResponse.SerializeToString), 'SetControlledPedestriansActions': grpc.unary_unary_rpc_method_handler(servicer.SetControlledPedestriansActions, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansActionsRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansActionsResponse.SerializeToString), 'GetControlledTaxiOrderAllocationPlan': grpc.unary_unary_rpc_method_handler(servicer.GetControlledTaxiOrderAllocationPlan, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.GetControlledTaxiOrderAllocationPlanRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.GetControlledTaxiOrderAllocationPlanResponse.SerializeToString), 'SetControlledTaxiOrderAllocationPlan': grpc.unary_unary_rpc_method_handler(servicer.SetControlledTaxiOrderAllocationPlan, request_deserializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiOrderAllocationPlanRequest.FromString, response_serializer=city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiOrderAllocationPlanResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('city.person.v2.PersonService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class PersonService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetPerson(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/GetPerson', city_dot_person_dot_v2_dot_person__service__pb2.GetPersonRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.GetPersonResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddPerson(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/AddPerson', city_dot_person_dot_v2_dot_person__service__pb2.AddPersonRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.AddPersonResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetSchedule(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/SetSchedule', city_dot_person_dot_v2_dot_person__service__pb2.SetScheduleRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.SetScheduleResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPersons(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/GetPersons', city_dot_person_dot_v2_dot_person__service__pb2.GetPersonsRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.GetPersonsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPersonByLongLatBBox(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/GetPersonByLongLatBBox', city_dot_person_dot_v2_dot_person__service__pb2.GetPersonByLongLatBBoxRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.GetPersonByLongLatBBoxResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAllVehicles(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/GetAllVehicles', city_dot_person_dot_v2_dot_person__service__pb2.GetAllVehiclesRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.GetAllVehiclesResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAllPedestrians(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/GetAllPedestrians', city_dot_person_dot_v2_dot_person__service__pb2.GetAllPedestriansRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.GetAllPedestriansResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ResetPersonPosition(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/ResetPersonPosition', city_dot_person_dot_v2_dot_person__service__pb2.ResetPersonPositionRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.ResetPersonPositionResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetControlledVehicleIDs(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/SetControlledVehicleIDs', city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleIDsRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleIDsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def FetchControlledVehicleEnvs(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/FetchControlledVehicleEnvs', city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledVehicleEnvsRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledVehicleEnvsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetControlledVehicleActions(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/SetControlledVehicleActions', city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleActionsRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.SetControlledVehicleActionsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetControlledTaxiIDs(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/SetControlledTaxiIDs', city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiIDsRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiIDsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAllOrders(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/GetAllOrders', city_dot_person_dot_v2_dot_person__service__pb2.GetAllOrdersRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.GetAllOrdersResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetControlledTaxiToOrders(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/SetControlledTaxiToOrders', city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiToOrdersRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiToOrdersResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetControlledPedestrians(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/SetControlledPedestrians', city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def FetchControlledPedestriansEnvs(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/FetchControlledPedestriansEnvs', city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledPedestriansEnvsRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.FetchControlledPedestriansEnvsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetControlledPedestriansActions(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/SetControlledPedestriansActions', city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansActionsRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.SetControlledPedestriansActionsResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetControlledTaxiOrderAllocationPlan(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/GetControlledTaxiOrderAllocationPlan', city_dot_person_dot_v2_dot_person__service__pb2.GetControlledTaxiOrderAllocationPlanRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.GetControlledTaxiOrderAllocationPlanResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetControlledTaxiOrderAllocationPlan(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/city.person.v2.PersonService/SetControlledTaxiOrderAllocationPlan', city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiOrderAllocationPlanRequest.SerializeToString, city_dot_person_dot_v2_dot_person__service__pb2.SetControlledTaxiOrderAllocationPlanResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)