// @generated by protoc-gen-es v1.10.0
// @generated from file city/map/v2/road_service.proto (package city.map.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Events } from "../../event/v1/event_pb.js";
import { LaneState } from "./lane_state_pb.js";

/**
 * 道路拥堵情况
 * road congestion level
 *
 * @generated from enum city.map.v2.RoadLevel
 */
export const RoadLevel = /*@__PURE__*/ proto3.makeEnum(
  "city.map.v2.RoadLevel",
  [
    {no: 0, name: "ROAD_LEVEL_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ROAD_LEVEL_CLEAR", localName: "CLEAR"},
    {no: 2, name: "ROAD_LEVEL_LIGHT_LOAD", localName: "LIGHT_LOAD"},
    {no: 3, name: "ROAD_LEVEL_MEDIUM_LOAD", localName: "MEDIUM_LOAD"},
    {no: 4, name: "ROAD_LEVEL_HEAVY_LOAD", localName: "HEAVY_LOAD"},
    {no: 5, name: "ROAD_LEVEL_OVERLOAD", localName: "OVERLOAD"},
    {no: 6, name: "ROAD_LEVEL_RESTRICTED", localName: "RESTRICTED"},
  ],
);

/**
 * 道路中断原因
 * road interruption reason
 *
 * @generated from enum city.map.v2.InterruptionReason
 */
export const InterruptionReason = /*@__PURE__*/ proto3.makeEnum(
  "city.map.v2.InterruptionReason",
  [
    {no: 0, name: "INTERRUPTION_REASON_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "INTERRUPTION_REASON_RUINED", localName: "RUINED"},
    {no: 2, name: "INTERRUPTION_REASON_CASCADE", localName: "CASCADE"},
    {no: 3, name: "INTERRUPTION_REASON_CONGESTION", localName: "CONGESTION"},
  ],
);

/**
 * 查询道路信息请求
 * Request for getting road information
 *
 * @generated from message city.map.v2.GetRoadRequest
 */
export const GetRoadRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetRoadRequest",
  () => [
    { no: 1, name: "road_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 2, name: "exclude_lane", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "exclude_person", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * 查询道路信息响应
 * Response of getting road information
 *
 * @generated from message city.map.v2.GetRoadResponse
 */
export const GetRoadResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetRoadResponse",
  () => [
    { no: 1, name: "states", kind: "message", T: RoadState, repeated: true },
  ],
);

/**
 * @generated from message city.map.v2.GetRuinInfoRequest
 */
export const GetRuinInfoRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetRuinInfoRequest",
  [],
);

/**
 * @generated from message city.map.v2.RuinInfo
 */
export const RuinInfo = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.RuinInfo",
  () => [
    { no: 1, name: "num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "ratio", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * @generated from message city.map.v2.GetRuinInfoResponse
 */
export const GetRuinInfoResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetRuinInfoResponse",
  () => [
    { no: 1, name: "one", kind: "message", T: RuinInfo },
    { no: 2, name: "two", kind: "message", T: RuinInfo },
    { no: 3, name: "three", kind: "message", T: RuinInfo },
  ],
);

/**
 * @generated from message city.map.v2.GetEventsRequest
 */
export const GetEventsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetEventsRequest",
  [],
);

/**
 * @generated from message city.map.v2.GetEventsResponse
 */
export const GetEventsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetEventsResponse",
  () => [
    { no: 1, name: "events", kind: "message", T: Events },
  ],
);

/**
 * 道路状态
 * road state
 *
 * @generated from message city.map.v2.RoadState
 */
export const RoadState = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.RoadState",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "avg_v", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "level", kind: "enum", T: proto3.getEnumType(RoadLevel) },
    { no: 3, name: "reason", kind: "enum", T: proto3.getEnumType(InterruptionReason) },
    { no: 5, name: "lanes", kind: "message", T: LaneState, repeated: true },
  ],
);

