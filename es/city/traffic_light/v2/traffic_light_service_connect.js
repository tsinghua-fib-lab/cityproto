// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/traffic_light/v2/traffic_light_service.proto (package city.traffic_light.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetTrafficLightRequest, GetTrafficLightResponse, SetTrafficLightPhaseRequest, SetTrafficLightPhaseResponse, SetTrafficLightRequest, SetTrafficLightResponse, SetTrafficLightStatusRequest, SetTrafficLightStatusResponse } from "./traffic_light_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.traffic_light.v2.TrafficLightService
 */
export const TrafficLightService = {
  typeName: "city.traffic_light.v2.TrafficLightService",
  methods: {
    /**
     * 获取路口的红绿灯信息
     * Get traffic light information
     *
     * @generated from rpc city.traffic_light.v2.TrafficLightService.GetTrafficLight
     */
    getTrafficLight: {
      name: "GetTrafficLight",
      I: GetTrafficLightRequest,
      O: GetTrafficLightResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 设置路口的红绿灯信息
     * Set traffic light information
     *
     * @generated from rpc city.traffic_light.v2.TrafficLightService.SetTrafficLight
     */
    setTrafficLight: {
      name: "SetTrafficLight",
      I: SetTrafficLightRequest,
      O: SetTrafficLightResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 设置路口的红绿灯相位
     * Set traffic light phase
     *
     * @generated from rpc city.traffic_light.v2.TrafficLightService.SetTrafficLightPhase
     */
    setTrafficLightPhase: {
      name: "SetTrafficLightPhase",
      I: SetTrafficLightPhaseRequest,
      O: SetTrafficLightPhaseResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 设置路口的红绿灯状态
     * Set traffic light status
     *
     * @generated from rpc city.traffic_light.v2.TrafficLightService.SetTrafficLightStatus
     */
    setTrafficLightStatus: {
      name: "SetTrafficLightStatus",
      I: SetTrafficLightStatusRequest,
      O: SetTrafficLightStatusResponse,
      kind: MethodKind.Unary,
    },
  }
};

