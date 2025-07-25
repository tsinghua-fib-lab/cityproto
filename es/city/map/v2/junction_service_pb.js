// @generated by protoc-gen-es v1.10.0
// @generated from file city/map/v2/junction_service.proto (package city.map.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { LaneState } from "./lane_state_pb.js";

/**
 * 查询路口信息请求
 * Request for getting junction information
 *
 * @generated from message city.map.v2.GetJunctionRequest
 */
export const GetJunctionRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetJunctionRequest",
  () => [
    { no: 1, name: "junction_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 2, name: "exclude_lane", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "exclude_person", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * 查询路口信息响应
 * Response of getting junction information
 *
 * @generated from message city.map.v2.GetJunctionResponse
 */
export const GetJunctionResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.GetJunctionResponse",
  () => [
    { no: 1, name: "states", kind: "message", T: JunctionState, repeated: true },
  ],
);

/**
 * 路口状态
 * junction state
 *
 * @generated from message city.map.v2.JunctionState
 */
export const JunctionState = /*@__PURE__*/ proto3.makeMessageType(
  "city.map.v2.JunctionState",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "in_vehicle_cnt", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "out_vehicle_cnt", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "vehicle_cnt", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "cum_in_vehicle_cnt", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "cum_out_vehicle_cnt", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "lanes", kind: "message", T: LaneState, repeated: true },
    { no: 8, name: "predecessor_driving_lanes", kind: "message", T: LaneState, repeated: true },
    { no: 9, name: "total_queuing_vehicle_cnt", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "total_queuing_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 11, name: "avg_queuing_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 12, name: "max_queuing_vehicle_cnt", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 13, name: "has_traffic_light", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

