// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/comm/interaction/aoi/v1/aoi_service.proto (package city.comm.interaction.aoi.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetBadAoiIDRequest, GetBadAoiIDResponse } from "./aoi_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.comm.interaction.aoi.v1.AoiService
 */
export declare const AoiService: {
  readonly typeName: "city.comm.interaction.aoi.v1.AoiService",
  readonly methods: {
    /**
     * @generated from rpc city.comm.interaction.aoi.v1.AoiService.GetBadAoiID
     */
    readonly getBadAoiID: {
      readonly name: "GetBadAoiID",
      readonly I: typeof GetBadAoiIDRequest,
      readonly O: typeof GetBadAoiIDResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

