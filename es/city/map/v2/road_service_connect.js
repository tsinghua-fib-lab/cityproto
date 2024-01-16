// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/map/v2/road_service.proto (package city.map.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetEventsRequest, GetEventsResponse, GetRoadByLongLatBBoxRequest, GetRoadByLongLatBBoxResponse, GetRoadRequest, GetRoadResponse, GetRuinInfoRequest, GetRuinInfoResponse } from "./road_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.map.v2.RoadService
 */
export const RoadService = {
  typeName: "city.map.v2.RoadService",
  methods: {
    /**
     * 查询道路信息
     *
     * @generated from rpc city.map.v2.RoadService.GetRoad
     */
    getRoad: {
      name: "GetRoad",
      I: GetRoadRequest,
      O: GetRoadResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 查询特定区域内的道路信息
     *
     * @generated from rpc city.map.v2.RoadService.GetRoadByLongLatBBox
     */
    getRoadByLongLatBBox: {
      name: "GetRoadByLongLatBBox",
      I: GetRoadByLongLatBBoxRequest,
      O: GetRoadByLongLatBBoxResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.map.v2.RoadService.GetRuinInfo
     */
    getRuinInfo: {
      name: "GetRuinInfo",
      I: GetRuinInfoRequest,
      O: GetRuinInfoResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.map.v2.RoadService.GetEvents
     */
    getEvents: {
      name: "GetEvents",
      I: GetEventsRequest,
      O: GetEventsResponse,
      kind: MethodKind.Unary,
    },
  }
};

