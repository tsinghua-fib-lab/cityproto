// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/ping/v1/ping_service.proto (package city.ping.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { PingRequest, PingResponse } from "./ping_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.ping.v1.PingService
 */
export const PingService = {
  typeName: "city.ping.v1.PingService",
  methods: {
    /**
     * 连接测试
     *
     * @generated from rpc city.ping.v1.PingService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
    },
  }
};

