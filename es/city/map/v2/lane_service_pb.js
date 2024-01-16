// @generated by protoc-gen-es v1.6.0
// @generated from file city/map/v2/lane_service.proto (package city.map.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { AgentMotion } from "../../agent/v2/motion_pb.js";

/**
 * 设置Lane的最大速度（限速）请求
 *
 * @generated from message city.map.v2.SetLaneMaxVRequest
 */
export const SetLaneMaxVRequest = proto3.makeMessageType(
  "city.map.v2.SetLaneMaxVRequest",
  () => [
    { no: 1, name: "lane_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "max_v", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 设置Lane的最大速度（限速）响应
 *
 * @generated from message city.map.v2.SetLaneMaxVResponse
 */
export const SetLaneMaxVResponse = proto3.makeMessageType(
  "city.map.v2.SetLaneMaxVResponse",
  [],
);

/**
 * 获取Lane的信息请求
 *
 * @generated from message city.map.v2.GetLaneRequest
 */
export const GetLaneRequest = proto3.makeMessageType(
  "city.map.v2.GetLaneRequest",
  () => [
    { no: 1, name: "lane_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ],
);

/**
 * 获取Lane的信息响应
 *
 * @generated from message city.map.v2.GetLaneResponse
 */
export const GetLaneResponse = proto3.makeMessageType(
  "city.map.v2.GetLaneResponse",
  () => [
    { no: 1, name: "states", kind: "message", T: LaneState, repeated: true },
  ],
);

/**
 * Lane状态
 *
 * @generated from message city.map.v2.LaneState
 */
export const LaneState = proto3.makeMessageType(
  "city.map.v2.LaneState",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "persons", kind: "message", T: AgentMotion, repeated: true },
    { no: 3, name: "avg_v", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "restriction", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

