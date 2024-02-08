// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/map/v2/aoi_service.proto (package city.map.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetAoiRequest, GetAoiResponse } from "./aoi_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.map.v2.AoiService
 */
export const AoiService = {
  typeName: "city.map.v2.AoiService",
  methods: {
    /**
     * 获取AOI信息
     * Get AOI information
     *
     * @generated from rpc city.map.v2.AoiService.GetAoi
     */
    getAoi: {
      name: "GetAoi",
      I: GetAoiRequest,
      O: GetAoiResponse,
      kind: MethodKind.Unary,
    },
  }
};

