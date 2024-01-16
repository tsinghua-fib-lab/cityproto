// @generated by protoc-gen-es v1.6.0
// @generated from file city/routing/v2/routing_service.proto (package city.routing.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Journey, RouteType } from "./routing_pb.js";
import { Position } from "../../geo/v2/geo_pb.js";
import { Cost } from "./cost_pb.js";

/**
 * 获取导航路线请求
 *
 * @generated from message city.routing.v2.GetRouteRequest
 */
export const GetRouteRequest = proto3.makeMessageType(
  "city.routing.v2.GetRouteRequest",
  () => [
    { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(RouteType) },
    { no: 2, name: "start", kind: "message", T: Position },
    { no: 3, name: "end", kind: "message", T: Position },
    { no: 4, name: "agent_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 获取导航路线响应
 *
 * @generated from message city.routing.v2.GetRouteResponse
 */
export const GetRouteResponse = proto3.makeMessageType(
  "city.routing.v2.GetRouteResponse",
  () => [
    { no: 1, name: "journeys", kind: "message", T: Journey, repeated: true },
  ],
);

/**
 * 设置行车导航道路通行成本请求
 *
 * @generated from message city.routing.v2.SetDrivingCostsRequest
 */
export const SetDrivingCostsRequest = proto3.makeMessageType(
  "city.routing.v2.SetDrivingCostsRequest",
  () => [
    { no: 1, name: "costs", kind: "message", T: Cost, repeated: true },
  ],
);

/**
 * 设置行车导航道路通行成本响应
 *
 * @generated from message city.routing.v2.SetDrivingCostsResponse
 */
export const SetDrivingCostsResponse = proto3.makeMessageType(
  "city.routing.v2.SetDrivingCostsResponse",
  [],
);

/**
 * 获取行车导航道路通行成本请求
 *
 * @generated from message city.routing.v2.GetDrivingCostsRequest
 */
export const GetDrivingCostsRequest = proto3.makeMessageType(
  "city.routing.v2.GetDrivingCostsRequest",
  () => [
    { no: 1, name: "costs", kind: "message", T: Cost, repeated: true },
  ],
);

/**
 * 获取行车导航道路通行成本响应
 *
 * @generated from message city.routing.v2.GetDrivingCostsResponse
 */
export const GetDrivingCostsResponse = proto3.makeMessageType(
  "city.routing.v2.GetDrivingCostsResponse",
  () => [
    { no: 1, name: "costs", kind: "message", T: Cost, repeated: true },
  ],
);

