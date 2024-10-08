// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/routing/v2/routing_service.proto (package city.routing.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetDrivingCostsRequest, GetDrivingCostsResponse, GetRouteRequest, GetRouteResponse, SetDrivingCostsRequest, SetDrivingCostsResponse } from "./routing_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.routing.v2.RoutingService
 */
export declare const RoutingService: {
  readonly typeName: "city.routing.v2.RoutingService",
  readonly methods: {
    /**
     * 获取导航路线
     * Get routing path
     *
     * @generated from rpc city.routing.v2.RoutingService.GetRoute
     */
    readonly getRoute: {
      readonly name: "GetRoute",
      readonly I: typeof GetRouteRequest,
      readonly O: typeof GetRouteResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 设置行车导航道路通行成本
     * Set traveling cost of driving routing
     *
     * @generated from rpc city.routing.v2.RoutingService.SetDrivingCosts
     */
    readonly setDrivingCosts: {
      readonly name: "SetDrivingCosts",
      readonly I: typeof SetDrivingCostsRequest,
      readonly O: typeof SetDrivingCostsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 获取行车导航道路通行成本
     * Get traveling cost of driving routing
     *
     * @generated from rpc city.routing.v2.RoutingService.GetDrivingCosts
     */
    readonly getDrivingCosts: {
      readonly name: "GetDrivingCosts",
      readonly I: typeof GetDrivingCostsRequest,
      readonly O: typeof GetDrivingCostsResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

