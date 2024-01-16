// @generated by protoc-gen-es v1.6.0
// @generated from file city/elec/input/v1/config.proto (package city.elec.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { MongoPath, OutputTarget } from "../../../config/v1/config_pb.js";

/**
 * @generated from message city.elec.input.v1.Mongo
 */
export const Mongo = proto3.makeMessageType(
  "city.elec.input.v1.Mongo",
  () => [
    { no: 1, name: "uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "map", kind: "message", T: MongoPath },
    { no: 3, name: "facilities", kind: "message", T: MongoPath },
  ],
);

/**
 * @generated from message city.elec.input.v1.ControlStep
 */
export const ControlStep = proto3.makeMessageType(
  "city.elec.input.v1.ControlStep",
  () => [
    { no: 1, name: "start", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "total", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.elec.input.v1.Control
 */
export const Control = proto3.makeMessageType(
  "city.elec.input.v1.Control",
  () => [
    { no: 1, name: "step", kind: "message", T: ControlStep },
  ],
);

/**
 * 是否输出各类数据
 *
 * @generated from message city.elec.input.v1.OutputSwitch
 */
export const OutputSwitch = proto3.makeMessageType(
  "city.elec.input.v1.OutputSwitch",
  () => [
    { no: 1, name: "node", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "aoi", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "event", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message city.elec.input.v1.Output
 */
export const Output = proto3.makeMessageType(
  "city.elec.input.v1.Output",
  () => [
    { no: 1, name: "target", kind: "message", T: OutputTarget },
    { no: 2, name: "switch", kind: "message", T: OutputSwitch },
  ],
);

/**
 * @generated from message city.elec.input.v1.Config
 */
export const Config = proto3.makeMessageType(
  "city.elec.input.v1.Config",
  () => [
    { no: 1, name: "mongo", kind: "message", T: Mongo },
    { no: 2, name: "control", kind: "message", T: Control },
    { no: 3, name: "output", kind: "message", T: Output },
  ],
);

