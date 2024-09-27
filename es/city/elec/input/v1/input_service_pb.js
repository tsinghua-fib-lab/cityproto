// @generated by protoc-gen-es v1.10.0
// @generated from file city/elec/input/v1/input_service.proto (package city.elec.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Control } from "./config_pb.js";
import { Facilities } from "./input_pb.js";
import { Map } from "../../../map/v2/map_pb.js";

/**
 * @generated from message city.elec.input.v1.InitRequest
 */
export const InitRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.elec.input.v1.InitRequest",
  [],
);

/**
 * @generated from message city.elec.input.v1.InitResponse
 */
export const InitResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.elec.input.v1.InitResponse",
  () => [
    { no: 2, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "control", kind: "message", T: Control },
    { no: 1, name: "facilities", kind: "message", T: Facilities },
    { no: 4, name: "map", kind: "message", T: Map },
  ],
);

