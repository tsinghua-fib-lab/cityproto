// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/map/v2/lane_service.proto (package city.map.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetLaneRequest, GetLaneResponse, SetLaneMaxVRequest, SetLaneMaxVResponse } from "./lane_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.map.v2.LaneService
 */
export declare const LaneService: {
  readonly typeName: "city.map.v2.LaneService",
  readonly methods: {
    /**
     * 设置Lane的最大速度（限速）
     *
     * @generated from rpc city.map.v2.LaneService.SetLaneMaxV
     */
    readonly setLaneMaxV: {
      readonly name: "SetLaneMaxV",
      readonly I: typeof SetLaneMaxVRequest,
      readonly O: typeof SetLaneMaxVResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 获取Lane的信息
     *
     * @generated from rpc city.map.v2.LaneService.GetLane
     */
    readonly getLane: {
      readonly name: "GetLane",
      readonly I: typeof GetLaneRequest,
      readonly O: typeof GetLaneResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};
