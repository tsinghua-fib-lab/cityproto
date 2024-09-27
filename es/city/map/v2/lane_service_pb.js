// @generated by protoc-gen-es v1.10.0
// @generated from file city/map/v2/lane_service.proto (package city.map.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { LaneState } from "./lane_state_pb.js";
import { LongLatBBox } from "../../geo/v2/geo_pb.js";

/**
 * 设置Lane的最大速度（限速）请求
 * Request for setting lane's maximum speed (speed limit)
 *
 * @generated from message city.map.v2.SetLaneMaxVRequest
 */
export const SetLaneMaxVRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.SetLaneMaxVRequest",
  () => [
    { no: 1, name: "lane_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "max_v", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 设置Lane的最大速度（限速）响应
 * Response of setting lane's maximum speed (speed limit)
 *
 * @generated from message city.map.v2.SetLaneMaxVResponse
 */
export const SetLaneMaxVResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.SetLaneMaxVResponse",
  [],
);

/**
 * 设置Lane限行请求
 * Request for setting lane's traffic restriction
 *
 * @generated from message city.map.v2.SetLaneRestrictionRequest
 */
export const SetLaneRestrictionRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.SetLaneRestrictionRequest",
  () => [
    { no: 1, name: "lane_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "restriction", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * 设置Lane限行响应
 * Response of setting lane's traffic restriction
 *
 * @generated from message city.map.v2.SetLaneRestrictionResponse
 */
export const SetLaneRestrictionResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.SetLaneRestrictionResponse",
  [],
);

/**
 * 获取Lane的信息请求
 * Request for getting lane information
 *
 * @generated from message city.map.v2.GetLaneRequest
 */
export const GetLaneRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetLaneRequest",
  () => [
    { no: 1, name: "lane_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 2, name: "exclude_person", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * 获取Lane的信息响应
 * Response of getting lane information
 *
 * @generated from message city.map.v2.GetLaneResponse
 */
export const GetLaneResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetLaneResponse",
  () => [
    { no: 1, name: "states", kind: "message", T: LaneState, repeated: true },
  ],
);

/**
 * 获取特定区域内的Lane的信息请求
 * Request for getting lane information in a specific region
 *
 * @generated from message city.map.v2.GetLaneByLongLatBBoxRequest
 */
export const GetLaneByLongLatBBoxRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetLaneByLongLatBBoxRequest",
  () => [
    { no: 1, name: "bbox", kind: "message", T: LongLatBBox },
    { no: 2, name: "exclude_person", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * 获取特定区域内的Lane的信息响应
 * Response of getting lane information in a specific region
 *
 * @generated from message city.map.v2.GetLaneByLongLatBBoxResponse
 */
export const GetLaneByLongLatBBoxResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetLaneByLongLatBBoxResponse",
  () => [
    { no: 1, name: "states", kind: "message", T: LaneState, repeated: true },
  ],
);

