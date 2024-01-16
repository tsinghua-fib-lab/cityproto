// @generated by protoc-gen-es v1.6.0
// @generated from file city/person/v1/person_service.proto (package city.person.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Person } from "./person_pb.js";
import { PersonMotion } from "./motion_pb.js";
import { Schedule } from "../../trip/v2/trip_pb.js";
import { LongLatBBox } from "../../geo/v2/geo_pb.js";

/**
 * 获取person信息请求
 *
 * @generated from message city.person.v1.GetPersonRequest
 */
export const GetPersonRequest = proto3.makeMessageType(
  "city.person.v1.GetPersonRequest",
  () => [
    { no: 1, name: "person_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * 获取person信息响应
 *
 * @generated from message city.person.v1.GetPersonResponse
 */
export const GetPersonResponse = proto3.makeMessageType(
  "city.person.v1.GetPersonResponse",
  () => [
    { no: 1, name: "base", kind: "message", T: Person },
    { no: 2, name: "motion", kind: "message", T: PersonMotion },
  ],
);

/**
 * 新增person请求
 *
 * @generated from message city.person.v1.AddPersonRequest
 */
export const AddPersonRequest = proto3.makeMessageType(
  "city.person.v1.AddPersonRequest",
  () => [
    { no: 1, name: "person", kind: "message", T: Person },
  ],
);

/**
 * 新增person响应
 *
 * @generated from message city.person.v1.AddPersonResponse
 */
export const AddPersonResponse = proto3.makeMessageType(
  "city.person.v1.AddPersonResponse",
  () => [
    { no: 1, name: "person_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * 修改person的schedule请求
 *
 * @generated from message city.person.v1.SetScheduleRequest
 */
export const SetScheduleRequest = proto3.makeMessageType(
  "city.person.v1.SetScheduleRequest",
  () => [
    { no: 1, name: "person_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "schedules", kind: "message", T: Schedule, repeated: true },
  ],
);

/**
 * 修改person的schedule响应
 *
 * @generated from message city.person.v1.SetScheduleResponse
 */
export const SetScheduleResponse = proto3.makeMessageType(
  "city.person.v1.SetScheduleResponse",
  [],
);

/**
 * 获取特定区域内的person请求
 *
 * @generated from message city.person.v1.GetPersonByLongLatBBoxRequest
 */
export const GetPersonByLongLatBBoxRequest = proto3.makeMessageType(
  "city.person.v1.GetPersonByLongLatBBoxRequest",
  () => [
    { no: 1, name: "bound", kind: "message", T: LongLatBBox },
  ],
);

/**
 * 获取特定区域内的person响应
 *
 * @generated from message city.person.v1.GetPersonByLongLatBBoxResponse
 */
export const GetPersonByLongLatBBoxResponse = proto3.makeMessageType(
  "city.person.v1.GetPersonByLongLatBBoxResponse",
  () => [
    { no: 1, name: "motions", kind: "message", T: PersonMotion, repeated: true },
  ],
);

