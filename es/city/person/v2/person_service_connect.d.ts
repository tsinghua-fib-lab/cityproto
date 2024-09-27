// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/person/v2/person_service.proto (package city.person.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddPersonRequest, AddPersonResponse, FetchControlledVehicleEnvsRequest, FetchControlledVehicleEnvsResponse, GetAllVehiclesRequest, GetAllVehiclesResponse, GetPersonByLongLatBBoxRequest, GetPersonByLongLatBBoxResponse, GetPersonRequest, GetPersonResponse, GetPersonsRequest, GetPersonsResponse, ResetPersonPositionRequest, ResetPersonPositionResponse, SetControlledVehicleActionsRequest, SetControlledVehicleActionsResponse, SetControlledVehicleIDsRequest, SetControlledVehicleIDsResponse, SetScheduleRequest, SetScheduleResponse } from "./person_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.person.v2.PersonService
 */
export declare const PersonService: {
  readonly typeName: "city.person.v2.PersonService",
  readonly methods: {
    /**
     * 获取person信息
     * Get person information
     *
     * @generated from rpc city.person.v2.PersonService.GetPerson
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
     * @generated from rpc city.person.v2.PersonService.AddPerson
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
     * @generated from rpc city.person.v2.PersonService.SetSchedule
     */
    readonly setSchedule: {
      readonly name: "SetSchedule",
      readonly I: typeof SetScheduleRequest,
      readonly O: typeof SetScheduleResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 获取多个person信息
     * Get information of multiple persons
     *
     * @generated from rpc city.person.v2.PersonService.GetPersons
     */
    readonly getPersons: {
      readonly name: "GetPersons",
      readonly I: typeof GetPersonsRequest,
      readonly O: typeof GetPersonsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 获取特定区域内的person
     * Get persons in a specific region
     *
     * @generated from rpc city.person.v2.PersonService.GetPersonByLongLatBBox
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
     * @generated from rpc city.person.v2.PersonService.GetAllVehicles
     */
    readonly getAllVehicles: {
      readonly name: "GetAllVehicles",
      readonly I: typeof GetAllVehiclesRequest,
      readonly O: typeof GetAllVehiclesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 重置人的位置（将停止当前正在进行的出行，转为sleep状态）
     * Reset person's position (stop the current trip and switch to sleep status)
     *
     * @generated from rpc city.person.v2.PersonService.ResetPersonPosition
     */
    readonly resetPersonPosition: {
      readonly name: "ResetPersonPosition",
      readonly I: typeof ResetPersonPositionRequest,
      readonly O: typeof ResetPersonPositionResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 设置由外部控制行为的vehicle
     * Set vehicle controlled by external behavior
     *
     * @generated from rpc city.person.v2.PersonService.SetControlledVehicleIDs
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
     * @generated from rpc city.person.v2.PersonService.FetchControlledVehicleEnvs
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
     * @generated from rpc city.person.v2.PersonService.SetControlledVehicleActions
     */
    readonly setControlledVehicleActions: {
      readonly name: "SetControlledVehicleActions",
      readonly I: typeof SetControlledVehicleActionsRequest,
      readonly O: typeof SetControlledVehicleActionsResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

