// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/map/v2/lane_service.proto (package city.map.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetLaneByLongLatBBoxRequest, GetLaneByLongLatBBoxResponse, GetLaneRequest, GetLaneResponse, SetLaneMaxVRequest, SetLaneMaxVResponse } from "./lane_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.map.v2.LaneService
 */
export const LaneService = {
  typeName: "city.map.v2.LaneService",
  methods: {
    /**
     * 设置Lane的最大速度（限速）
     * Set Lane's maximum speed (speed limit)
     *
     * @generated from rpc city.map.v2.LaneService.SetLaneMaxV
     */
    setLaneMaxV: {
      name: "SetLaneMaxV",
      I: SetLaneMaxVRequest,
      O: SetLaneMaxVResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取Lane的信息
     * Get Lane information
     *
     * @generated from rpc city.map.v2.LaneService.GetLane
     */
    getLane: {
      name: "GetLane",
      I: GetLaneRequest,
      O: GetLaneResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取特定区域内的Lane的信息
     * Get Lane information in a specific region
     *
     * @generated from rpc city.map.v2.LaneService.GetLaneByLongLatBBox
     */
    getLaneByLongLatBBox: {
      name: "GetLaneByLongLatBBox",
      I: GetLaneByLongLatBBoxRequest,
      O: GetLaneByLongLatBBoxResponse,
      kind: MethodKind.Unary,
    },
  }
};

