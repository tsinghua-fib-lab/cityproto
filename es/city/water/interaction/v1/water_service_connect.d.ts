// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/water/interaction/v1/water_service.proto (package city.water.interaction.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetNoWaterAOIRequest, GetNoWaterAOIResponse, GetPumpStatusRequest, GetPumpStatusResponse, GetRuinInfoRequest, GetRuinInfoResponse, SetPumpNetworkStatusRequest, SetPumpNetworkStatusResponse, SetPumpPowerStatusRequest, SetPumpPowerStatusResponse, SetPumpStatusRequest, SetPumpStatusResponse } from "./water_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.water.interaction.v1.WaterService
 */
export declare const WaterService: {
  readonly typeName: "city.water.interaction.v1.WaterService",
  readonly methods: {
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.SetPumpPowerStatus
     */
    readonly setPumpPowerStatus: {
      readonly name: "SetPumpPowerStatus",
      readonly I: typeof SetPumpPowerStatusRequest,
      readonly O: typeof SetPumpPowerStatusResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.SetPumpNetworkStatus
     */
    readonly setPumpNetworkStatus: {
      readonly name: "SetPumpNetworkStatus",
      readonly I: typeof SetPumpNetworkStatusRequest,
      readonly O: typeof SetPumpNetworkStatusResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.SetPumpStatus
     */
    readonly setPumpStatus: {
      readonly name: "SetPumpStatus",
      readonly I: typeof SetPumpStatusRequest,
      readonly O: typeof SetPumpStatusResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.GetPumpStatus
     */
    readonly getPumpStatus: {
      readonly name: "GetPumpStatus",
      readonly I: typeof GetPumpStatusRequest,
      readonly O: typeof GetPumpStatusResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.GetNoWaterAOI
     */
    readonly getNoWaterAOI: {
      readonly name: "GetNoWaterAOI",
      readonly I: typeof GetNoWaterAOIRequest,
      readonly O: typeof GetNoWaterAOIResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.GetRuinInfo
     */
    readonly getRuinInfo: {
      readonly name: "GetRuinInfo",
      readonly I: typeof GetRuinInfoRequest,
      readonly O: typeof GetRuinInfoResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

