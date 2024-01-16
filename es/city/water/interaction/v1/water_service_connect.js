// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/water/interaction/v1/water_service.proto (package city.water.interaction.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetNoWaterAOIRequest, GetNoWaterAOIResponse, GetPumpStatusRequest, GetPumpStatusResponse, GetRuinInfoRequest, GetRuinInfoResponse, SetPumpNetworkStatusRequest, SetPumpNetworkStatusResponse, SetPumpPowerStatusRequest, SetPumpPowerStatusResponse, SetPumpStatusRequest, SetPumpStatusResponse } from "./water_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.water.interaction.v1.WaterService
 */
export const WaterService = {
  typeName: "city.water.interaction.v1.WaterService",
  methods: {
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.SetPumpPowerStatus
     */
    setPumpPowerStatus: {
      name: "SetPumpPowerStatus",
      I: SetPumpPowerStatusRequest,
      O: SetPumpPowerStatusResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.SetPumpNetworkStatus
     */
    setPumpNetworkStatus: {
      name: "SetPumpNetworkStatus",
      I: SetPumpNetworkStatusRequest,
      O: SetPumpNetworkStatusResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.SetPumpStatus
     */
    setPumpStatus: {
      name: "SetPumpStatus",
      I: SetPumpStatusRequest,
      O: SetPumpStatusResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.GetPumpStatus
     */
    getPumpStatus: {
      name: "GetPumpStatus",
      I: GetPumpStatusRequest,
      O: GetPumpStatusResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.GetNoWaterAOI
     */
    getNoWaterAOI: {
      name: "GetNoWaterAOI",
      I: GetNoWaterAOIRequest,
      O: GetNoWaterAOIResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.water.interaction.v1.WaterService.GetRuinInfo
     */
    getRuinInfo: {
      name: "GetRuinInfo",
      I: GetRuinInfoRequest,
      O: GetRuinInfoResponse,
      kind: MethodKind.Unary,
    },
  }
};
