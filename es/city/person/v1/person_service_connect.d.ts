// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/person/v1/person_service.proto (package city.person.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddPersonRequest, AddPersonResponse, FetchControlledVehicleEnvsRequest, FetchControlledVehicleEnvsResponse, GetAllVehiclesRequest, GetAllVehiclesResponse, GetPersonByLongLatBBoxRequest, GetPersonByLongLatBBoxResponse, GetPersonRequest, GetPersonResponse, SetControlledVehicleActionsRequest, SetControlledVehicleActionsResponse, SetControlledVehicleIDsRequest, SetControlledVehicleIDsResponse, SetScheduleRequest, SetScheduleResponse } from "./person_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.person.v1.PersonService
 */
export declare const PersonService: {
  readonly typeName: "city.person.v1.PersonService",
  readonly methods: {
    /**
     * 获取person信息
     * Get person information
     *
     * @generated from rpc city.person.v1.PersonService.GetPerson
     */
    readonly getPerson: {
      readonly name: "GetPerson",
      readonly I: typeof GetPersonRequest,
      readonly O: typeof GetPersonResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 新增person 传入person初始位置、目的地表、属性 返回personid
     * Add a new person. Input person's initial location, destination table, and attributes, return personid
     *
     * @generated from rpc city.person.v1.PersonService.AddPerson
     */
    readonly addPerson: {
      readonly name: "AddPerson",
      readonly I: typeof AddPersonRequest,
      readonly O: typeof AddPersonResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 修改person的schedule 传入personid、目的地表
     * Set person's schedule. Input personid and destination table
     *
     * @generated from rpc city.person.v1.PersonService.SetSchedule
     */
    readonly setSchedule: {
      readonly name: "SetSchedule",
      readonly I: typeof SetScheduleRequest,
      readonly O: typeof SetScheduleResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 获取特定区域内的person
     * Get persons in a specific region
     *
     * @generated from rpc city.person.v1.PersonService.GetPersonByLongLatBBox
     */
    readonly getPersonByLongLatBBox: {
      readonly name: "GetPersonByLongLatBBox",
      readonly I: typeof GetPersonByLongLatBBoxRequest,
      readonly O: typeof GetPersonByLongLatBBoxResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 获取所有车辆
     * Get all vehicles
     *
     * @generated from rpc city.person.v1.PersonService.GetAllVehicles
     */
    readonly getAllVehicles: {
      readonly name: "GetAllVehicles",
      readonly I: typeof GetAllVehiclesRequest,
      readonly O: typeof GetAllVehiclesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 设置由外部控制行为的vehicle
     * Set vehicle controlled by external behavior
     *
     * @generated from rpc city.person.v1.PersonService.SetControlledVehicleIDs
     */
    readonly setControlledVehicleIDs: {
      readonly name: "SetControlledVehicleIDs",
      readonly I: typeof SetControlledVehicleIDsRequest,
      readonly O: typeof SetControlledVehicleIDsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 获取由外部控制行为的vehicle信息
     * Get information of vehicle controlled by external behavior
     *
     * @generated from rpc city.person.v1.PersonService.FetchControlledVehicleEnvs
     */
    readonly fetchControlledVehicleEnvs: {
      readonly name: "FetchControlledVehicleEnvs",
      readonly I: typeof FetchControlledVehicleEnvsRequest,
      readonly O: typeof FetchControlledVehicleEnvsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 设置由外部控制行为的vehicle的行为
     * Set behavior of vehicle controlled by external behavior
     *
     * @generated from rpc city.person.v1.PersonService.SetControlledVehicleActions
     */
    readonly setControlledVehicleActions: {
      readonly name: "SetControlledVehicleActions",
      readonly I: typeof SetControlledVehicleActionsRequest,
      readonly O: typeof SetControlledVehicleActionsResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

