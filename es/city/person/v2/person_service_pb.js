// @generated by protoc-gen-es v1.6.0
// @generated from file city/person/v2/person_service.proto (package city.person.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { PersonRuntime } from "./person_runtime_pb.js";
import { Person } from "./person_pb.js";
import { Schedule } from "../../trip/v2/trip_pb.js";
import { Status } from "./motion_pb.js";
import { LongLatBBox, Position } from "../../geo/v2/geo_pb.js";
import { VehicleAction, VehicleEnv, VehicleRuntime } from "./vehicle_pb.js";

/**
 * 获取person信息请求
 * Request for getting person information
 *
 * @generated from message city.person.v2.GetPersonRequest
 */
export const GetPersonRequest = proto3.makeMessageType(
  "city.person.v2.GetPersonRequest",
  () => [
    { no: 1, name: "person_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * 获取person信息响应
 * Response of getting person information
 *
 * @generated from message city.person.v2.GetPersonResponse
 */
export const GetPersonResponse = proto3.makeMessageType(
  "city.person.v2.GetPersonResponse",
  () => [
    { no: 1, name: "person", kind: "message", T: PersonRuntime },
  ],
);

/**
 * 新增person请求
 * Request for adding a new person
 *
 * @generated from message city.person.v2.AddPersonRequest
 */
export const AddPersonRequest = proto3.makeMessageType(
  "city.person.v2.AddPersonRequest",
  () => [
    { no: 1, name: "person", kind: "message", T: Person },
  ],
);

/**
 * 新增person响应
 * Response of adding a new person
 *
 * @generated from message city.person.v2.AddPersonResponse
 */
export const AddPersonResponse = proto3.makeMessageType(
  "city.person.v2.AddPersonResponse",
  () => [
    { no: 1, name: "person_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * 修改person的schedule请求
 * Request for setting person schedule
 *
 * @generated from message city.person.v2.SetScheduleRequest
 */
export const SetScheduleRequest = proto3.makeMessageType(
  "city.person.v2.SetScheduleRequest",
  () => [
    { no: 1, name: "person_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "schedules", kind: "message", T: Schedule, repeated: true },
  ],
);

/**
 * 修改person的schedule响应
 * Response of setting person schedule
 *
 * @generated from message city.person.v2.SetScheduleResponse
 */
export const SetScheduleResponse = proto3.makeMessageType(
  "city.person.v2.SetScheduleResponse",
  [],
);

/**
 * 获取多个person信息请求
 * Request for getting information of multiple persons
 *
 * @generated from message city.person.v2.GetPersonsRequest
 */
export const GetPersonsRequest = proto3.makeMessageType(
  "city.person.v2.GetPersonsRequest",
  () => [
    { no: 1, name: "person_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 2, name: "exclude_statuses", kind: "enum", T: proto3.getEnumType(Status), repeated: true },
    { no: 3, name: "return_base", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * 获取多个person信息响应
 * Response of getting information of multiple persons
 *
 * @generated from message city.person.v2.GetPersonsResponse
 */
export const GetPersonsResponse = proto3.makeMessageType(
  "city.person.v2.GetPersonsResponse",
  () => [
    { no: 1, name: "persons", kind: "message", T: PersonRuntime, repeated: true },
  ],
);

/**
 * 获取特定区域内的person请求
 * Request for getting persons in region
 *
 * @generated from message city.person.v2.GetPersonByLongLatBBoxRequest
 */
export const GetPersonByLongLatBBoxRequest = proto3.makeMessageType(
  "city.person.v2.GetPersonByLongLatBBoxRequest",
  () => [
    { no: 1, name: "bbox", kind: "message", T: LongLatBBox },
    { no: 2, name: "exclude_statuses", kind: "enum", T: proto3.getEnumType(Status), repeated: true },
    { no: 3, name: "return_base", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * 获取特定区域内的person响应
 * Response of getting persons in region
 *
 * @generated from message city.person.v2.GetPersonByLongLatBBoxResponse
 */
export const GetPersonByLongLatBBoxResponse = proto3.makeMessageType(
  "city.person.v2.GetPersonByLongLatBBoxResponse",
  () => [
    { no: 1, name: "persons", kind: "message", T: PersonRuntime, repeated: true },
  ],
);

/**
 * 获取所有车辆请求
 * Request for getting all vehicles
 *
 * @generated from message city.person.v2.GetAllVehiclesRequest
 */
export const GetAllVehiclesRequest = proto3.makeMessageType(
  "city.person.v2.GetAllVehiclesRequest",
  [],
);

/**
 * 获取所有车辆响应
 * Response of getting all vehicles
 *
 * @generated from message city.person.v2.GetAllVehiclesResponse
 */
export const GetAllVehiclesResponse = proto3.makeMessageType(
  "city.person.v2.GetAllVehiclesResponse",
  () => [
    { no: 1, name: "vehicles", kind: "message", T: VehicleRuntime, repeated: true },
  ],
);

/**
 * 重置人的位置请求
 * Request for resetting person's position
 *
 * @generated from message city.person.v2.ResetPersonPositionRequest
 */
export const ResetPersonPositionRequest = proto3.makeMessageType(
  "city.person.v2.ResetPersonPositionRequest",
  () => [
    { no: 1, name: "person_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "position", kind: "message", T: Position },
  ],
);

/**
 * 重置人的位置响应
 * Response of resetting person's position
 *
 * @generated from message city.person.v2.ResetPersonPositionResponse
 */
export const ResetPersonPositionResponse = proto3.makeMessageType(
  "city.person.v2.ResetPersonPositionResponse",
  [],
);

/**
 * 设置由外部控制行为的vehicle请求（下一个step生效）
 * Request for setting vehicle controlled by external behavior
 *
 * @generated from message city.person.v2.SetControlledVehicleIDsRequest
 */
export const SetControlledVehicleIDsRequest = proto3.makeMessageType(
  "city.person.v2.SetControlledVehicleIDsRequest",
  () => [
    { no: 1, name: "vehicle_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ],
);

/**
 * 设置由外部控制行为的vehicle响应
 * Response of setting vehicle controlled by external behavior
 *
 * @generated from message city.person.v2.SetControlledVehicleIDsResponse
 */
export const SetControlledVehicleIDsResponse = proto3.makeMessageType(
  "city.person.v2.SetControlledVehicleIDsResponse",
  [],
);

/**
 * 获取由外部控制行为的vehicle信息请求
 * Request for getting information of vehicle controlled by external behavior
 *
 * @generated from message city.person.v2.FetchControlledVehicleEnvsRequest
 */
export const FetchControlledVehicleEnvsRequest = proto3.makeMessageType(
  "city.person.v2.FetchControlledVehicleEnvsRequest",
  [],
);

/**
 * 获取由外部控制行为的vehicle信息响应
 * Response of getting information of vehicle controlled by external behavior
 *
 * @generated from message city.person.v2.FetchControlledVehicleEnvsResponse
 */
export const FetchControlledVehicleEnvsResponse = proto3.makeMessageType(
  "city.person.v2.FetchControlledVehicleEnvsResponse",
  () => [
    { no: 1, name: "vehicle_envs", kind: "message", T: VehicleEnv, repeated: true },
  ],
);

/**
 * 设置由外部控制行为的vehicle的行为请求
 * Request for setting behavior of vehicle controlled by external behavior
 *
 * @generated from message city.person.v2.SetControlledVehicleActionsRequest
 */
export const SetControlledVehicleActionsRequest = proto3.makeMessageType(
  "city.person.v2.SetControlledVehicleActionsRequest",
  () => [
    { no: 1, name: "vehicle_actions", kind: "message", T: VehicleAction, repeated: true },
  ],
);

/**
 * 设置由外部控制行为的vehicle的行为响应
 * Response of setting behavior of vehicle controlled by external behavior
 *
 * @generated from message city.person.v2.SetControlledVehicleActionsResponse
 */
export const SetControlledVehicleActionsResponse = proto3.makeMessageType(
  "city.person.v2.SetControlledVehicleActionsResponse",
  [],
);
