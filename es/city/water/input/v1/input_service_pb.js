// @generated by protoc-gen-es v1.6.0
// @generated from file city/water/input/v1/input_service.proto (package city.water.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Control } from "./config_pb.js";
import { Rain } from "./water_pb.js";
import { Map } from "../../../map/v2/map_pb.js";

/**
 * @generated from message city.water.input.v1.InitRequest
 */
export const InitRequest = proto3.makeMessageType(
  "city.water.input.v1.InitRequest",
  [],
);

/**
 * @generated from message city.water.input.v1.InitResponse
 */
export const InitResponse = proto3.makeMessageType(
  "city.water.input.v1.InitResponse",
  () => [
    { no: 2, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "control", kind: "message", T: Control },
    { no: 1, name: "rain", kind: "message", T: Rain },
    { no: 4, name: "map", kind: "message", T: Map },
  ],
);

