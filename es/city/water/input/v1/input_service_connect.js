// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/water/input/v1/input_service.proto (package city.water.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { InitRequest, InitResponse } from "./input_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.water.input.v1.InputService
 */
export const InputService = {
  typeName: "city.water.input.v1.InputService",
  methods: {
    /**
     * @generated from rpc city.water.input.v1.InputService.Init
     */
    init: {
      name: "Init",
      I: InitRequest,
      O: InitResponse,
      kind: MethodKind.Unary,
    },
  }
};

