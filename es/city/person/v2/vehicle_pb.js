// @generated by protoc-gen-es v1.10.0
// @generated from file city/person/v2/vehicle.proto (package city.person.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Journey } from "../../routing/v2/routing_pb.js";
import { PersonMotion } from "./motion_pb.js";

/**
 * @generated from enum city.person.v2.VehicleRelation
 */
export const VehicleRelation = /*@__PURE__*/ proto3.makeEnum(
  "city.person.v2.VehicleRelation",
  [
    {no: 0, name: "VEHICLE_RELATION_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "VEHICLE_RELATION_AHEAD", localName: "AHEAD"},
    {no: 2, name: "VEHICLE_RELATION_BEHIND", localName: "BEHIND"},
    {no: 3, name: "VEHICLE_RELATION_SHADOW_AHEAD", localName: "SHADOW_AHEAD"},
    {no: 4, name: "VEHICLE_RELATION_SHADOW_BEHIND", localName: "SHADOW_BEHIND"},
    {no: 5, name: "VEHICLE_RELATION_LEFT_AHEAD", localName: "LEFT_AHEAD"},
    {no: 6, name: "VEHICLE_RELATION_RIGHT_AHEAD", localName: "RIGHT_AHEAD"},
    {no: 7, name: "VEHICLE_RELATION_LEFT_BEHIND", localName: "LEFT_BEHIND"},
    {no: 8, name: "VEHICLE_RELATION_RIGHT_BEHIND", localName: "RIGHT_BEHIND"},
  ],
);

/**
 * 交通灯的状态
 * traffic light state
 *
 * @generated from enum city.person.v2.LightState
 */
export const LightState = /*@__PURE__*/ proto3.makeEnum(
  "city.person.v2.LightState",
  [
    {no: 0, name: "LIGHT_STATE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "LIGHT_STATE_RED", localName: "RED"},
    {no: 2, name: "LIGHT_STATE_GREEN", localName: "GREEN"},
    {no: 3, name: "LIGHT_STATE_YELLOW", localName: "YELLOW"},
  ],
);

/**
 * 变道相关的信息
 * lane change related information
 *
 * @generated from message city.person.v2.LC
 */
export const LC = /*@__PURE__*/ proto3.makeMessageType(
  "city.person.v2.LC",
  () => [
    { no: 1, name: "shadow_lane_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "shadow_s", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "angle", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "completed_ratio", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 车辆控制信息
 * vehicle control information
 *
 * @generated from message city.person.v2.VehicleAction
 */
export const VehicleAction = /*@__PURE__*/ proto3.makeMessageType(
  "city.person.v2.VehicleAction",
  () => [
    { no: 4, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 1, name: "acc", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "lc_target_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 3, name: "angle", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 修改车辆路由信息
 * vehicle routing information modification
 *
 * @generated from message city.person.v2.VehicleRouteAction
 */
export const VehicleRouteAction = /*@__PURE__*/ proto3.makeMessageType(
  "city.person.v2.VehicleRouteAction",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "journey", kind: "message", T: Journey },
  ],
);

/**
 * @generated from message city.person.v2.VehicleRuntime
 */
export const VehicleRuntime = /*@__PURE__*/ proto3.makeMessageType(
  "city.person.v2.VehicleRuntime",
  () => [
    { no: 1, name: "base", kind: "message", T: PersonMotion },
    { no: 4, name: "lc", kind: "message", T: LC, opt: true },
    { no: 5, name: "action", kind: "message", T: VehicleAction, opt: true },
    { no: 6, name: "running_distance", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 7, name: "num_going_astray", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "departure_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 9, name: "eta", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 10, name: "eta_free_flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 观测到的车辆
 * observed vehicles
 *
 * @generated from message city.person.v2.ObservedVehicle
 */
export const ObservedVehicle = /*@__PURE__*/ proto3.makeMessageType(
  "city.person.v2.ObservedVehicle",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "motion", kind: "message", T: PersonMotion },
    { no: 3, name: "relative_distance", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "relation", kind: "enum", T: proto3.getEnumType(VehicleRelation) },
  ],
);

/**
 * 观测到的车道
 * observed lane
 *
 * @generated from message city.person.v2.ObservedLane
 */
export const ObservedLane = /*@__PURE__*/ proto3.makeMessageType(
  "city.person.v2.ObservedLane",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "restriction", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "light_state", kind: "enum", T: proto3.getEnumType(LightState) },
    { no: 4, name: "light_remaining_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * @generated from message city.person.v2.VehicleEnv
 */
export const VehicleEnv = /*@__PURE__*/ proto3.makeMessageType(
  "city.person.v2.VehicleEnv",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "runtime", kind: "message", T: VehicleRuntime },
    { no: 3, name: "journey", kind: "message", T: Journey },
    { no: 4, name: "observed_vehicles", kind: "message", T: ObservedVehicle, repeated: true },
    { no: 5, name: "observed_lanes", kind: "message", T: ObservedLane, repeated: true },
  ],
);

