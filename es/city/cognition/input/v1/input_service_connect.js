// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/cognition/input/v1/input_service.proto (package city.cognition.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { InitRequest, InitResponse } from "./input_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.cognition.input.v1.InputService
 */
export const InputService = {
  typeName: "city.cognition.input.v1.InputService",
  methods: {
    /**
     * @generated from rpc city.cognition.input.v1.InputService.Init
     */
    init: {
      name: "Init",
      I: InitRequest,
      O: InitResponse,
      kind: MethodKind.Unary,
    },
  }
};

