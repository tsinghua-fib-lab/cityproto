// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/person/v2/person_service.proto (package city.person.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddPersonRequest, AddPersonResponse, FetchControlledVehicleEnvsRequest, FetchControlledVehicleEnvsResponse, GetAllOrdersRequest, GetAllOrdersResponse, GetAllPedestriansRequest, GetAllPedestriansResponse, GetAllVehiclesRequest, GetAllVehiclesResponse, GetPersonByLongLatBBoxRequest, GetPersonByLongLatBBoxResponse, GetPersonRequest, GetPersonResponse, GetPersonsRequest, GetPersonsResponse, ResetPersonPositionRequest, ResetPersonPositionResponse, SetControlledPedestriansActionsRequest, SetControlledPedestriansActionsResponse, SetControlledPedestriansRequest, SetControlledPedestriansResponse, SetControlledTaxiIDsRequest, SetControlledTaxiIDsResponse, SetControlledTaxiToOrdersRequest, SetControlledTaxiToOrdersResponse, SetControlledVehicleActionsRequest, SetControlledVehicleActionsResponse, SetControlledVehicleIDsRequest, SetControlledVehicleIDsResponse, SetScheduleRequest, SetScheduleResponse } from "./person_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.person.v2.PersonService
 */
export const PersonService = {
  typeName: "city.person.v2.PersonService",
  methods: {
    /**
     * 获取person信息
     * Get person information
     *
     * @generated from rpc city.person.v2.PersonService.GetPerson
     */
    getPerson: {
      name: "GetPerson",
      I: GetPersonRequest,
      O: GetPersonResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 新增person 传入person初始位置、目的地表、属性 返回personid
     * Add a new person. Input person's initial location, destination table, and
     * attributes, return personid
     *
     * @generated from rpc city.person.v2.PersonService.AddPerson
     */
    addPerson: {
      name: "AddPerson",
      I: AddPersonRequest,
      O: AddPersonResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 修改person的schedule 传入personid、目的地表
     * Set person's schedule. Input personid and destination table
     *
     * @generated from rpc city.person.v2.PersonService.SetSchedule
     */
    setSchedule: {
      name: "SetSchedule",
      I: SetScheduleRequest,
      O: SetScheduleResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取多个person信息
     * Get information of multiple persons
     *
     * @generated from rpc city.person.v2.PersonService.GetPersons
     */
    getPersons: {
      name: "GetPersons",
      I: GetPersonsRequest,
      O: GetPersonsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取特定区域内的person
     * Get persons in a specific region
     *
     * @generated from rpc city.person.v2.PersonService.GetPersonByLongLatBBox
     */
    getPersonByLongLatBBox: {
      name: "GetPersonByLongLatBBox",
      I: GetPersonByLongLatBBoxRequest,
      O: GetPersonByLongLatBBoxResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取所有车辆
     * Get all vehicles
     *
     * @generated from rpc city.person.v2.PersonService.GetAllVehicles
     */
    getAllVehicles: {
      name: "GetAllVehicles",
      I: GetAllVehiclesRequest,
      O: GetAllVehiclesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取所有行人
     * Get all pedestrians
     *
     * @generated from rpc city.person.v2.PersonService.GetAllPedestrians
     */
    getAllPedestrians: {
      name: "GetAllPedestrians",
      I: GetAllPedestriansRequest,
      O: GetAllPedestriansResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 重置人的位置（将停止当前正在进行的出行，转为sleep状态）
     * Reset person's position (stop the current trip and switch to sleep status)
     *
     * @generated from rpc city.person.v2.PersonService.ResetPersonPosition
     */
    resetPersonPosition: {
      name: "ResetPersonPosition",
      I: ResetPersonPositionRequest,
      O: ResetPersonPositionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 设置由外部控制行为的vehicle
     * Set vehicle controlled by external behavior
     *
     * @generated from rpc city.person.v2.PersonService.SetControlledVehicleIDs
     */
    setControlledVehicleIDs: {
      name: "SetControlledVehicleIDs",
      I: SetControlledVehicleIDsRequest,
      O: SetControlledVehicleIDsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取由外部控制行为的vehicle信息
     * Get information of vehicle controlled by external behavior
     *
     * @generated from rpc city.person.v2.PersonService.FetchControlledVehicleEnvs
     */
    fetchControlledVehicleEnvs: {
      name: "FetchControlledVehicleEnvs",
      I: FetchControlledVehicleEnvsRequest,
      O: FetchControlledVehicleEnvsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 设置由外部控制行为的vehicle的行为
     * Set behavior of vehicle controlled by external behavior
     *
     * @generated from rpc city.person.v2.PersonService.SetControlledVehicleActions
     */
    setControlledVehicleActions: {
      name: "SetControlledVehicleActions",
      I: SetControlledVehicleActionsRequest,
      O: SetControlledVehicleActionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 设置由外部控制的taxi
     * Set taxi controlled by external behavior
     *
     * @generated from rpc city.person.v2.PersonService.SetControlledTaxiIDs
     */
    setControlledTaxiIDs: {
      name: "SetControlledTaxiIDs",
      I: SetControlledTaxiIDsRequest,
      O: SetControlledTaxiIDsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取所有订单信息
     * Get information of all orders
     *
     * @generated from rpc city.person.v2.PersonService.GetAllOrders
     */
    getAllOrders: {
      name: "GetAllOrders",
      I: GetAllOrdersRequest,
      O: GetAllOrdersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 设置所有外部控制的出租车接指定的单
     * Set all externally controlled taxis to specified orders
     *
     * @generated from rpc city.person.v2.PersonService.SetControlledTaxiToOrders
     */
    setControlledTaxiToOrders: {
      name: "SetControlledTaxiToOrders",
      I: SetControlledTaxiToOrdersRequest,
      O: SetControlledTaxiToOrdersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 设置由外部控制的行人
     * Set pedestrian controlled by external behavior
     *
     * @generated from rpc city.person.v2.PersonService.SetControlledPedestrians
     */
    setControlledPedestrians: {
      name: "SetControlledPedestrians",
      I: SetControlledPedestriansRequest,
      O: SetControlledPedestriansResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 设置由外部控制的行人行为
     * Set behavior of pedestrian controlled by external behavior
     *
     * @generated from rpc city.person.v2.PersonService.SetControlledPedestriansActions
     */
    setControlledPedestriansActions: {
      name: "SetControlledPedestriansActions",
      I: SetControlledPedestriansActionsRequest,
      O: SetControlledPedestriansActionsResponse,
      kind: MethodKind.Unary,
    },
  }
};

