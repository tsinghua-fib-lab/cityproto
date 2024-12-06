// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/agentcomm/v1/agent_service.proto (package city.agentcomm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CommunicateRequest, CommunicateResponse } from "./agent_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.agentcomm.v1.AgentService
 */
export declare const AgentService: {
  readonly typeName: "city.agentcomm.v1.AgentService",
  readonly methods: {
    /**
     * @generated from rpc city.agentcomm.v1.AgentService.Communicate
     */
    readonly communicate: {
      readonly name: "Communicate",
      readonly I: typeof CommunicateRequest,
      readonly O: typeof CommunicateResponse,
      readonly kind: MethodKind.BiDiStreaming,
    },
  }
};

