// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/comm/interaction/demand/v1/demand_service.proto (package city.comm.interaction.demand.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { SetDemandStatusRequest, SetDemandStatusResponse } from "./demand_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.comm.interaction.demand.v1.DemandService
 */
export const DemandService = {
  typeName: "city.comm.interaction.demand.v1.DemandService",
  methods: {
    /**
     * @generated from rpc city.comm.interaction.demand.v1.DemandService.SetDemandStatus
     */
    setDemandStatus: {
      name: "SetDemandStatus",
      I: SetDemandStatusRequest,
      O: SetDemandStatusResponse,
      kind: MethodKind.Unary,
    },
  }
};

