// @generated by protoc-gen-es v1.6.0
// @generated from file city/comm/output/v1/output_service.proto (package city.comm.output.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Aoi, BaseStation, Node, Person, Signal, Statistics } from "./output_pb.js";
import { Events } from "../../../event/v1/event_pb.js";

/**
 * @generated from message city.comm.output.v1.OutputRequest
 */
export const OutputRequest = proto3.makeMessageType(
  "city.comm.output.v1.OutputRequest",
  () => [
    { no: 1, name: "step", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "nodes", kind: "message", T: Node, repeated: true },
    { no: 3, name: "base_stations", kind: "message", T: BaseStation, repeated: true },
    { no: 4, name: "signal_heatmap", kind: "message", T: Signal },
    { no: 9, name: "small_signal_heatmap", kind: "message", T: Signal },
    { no: 5, name: "persons", kind: "message", T: Person, repeated: true },
    { no: 6, name: "aois", kind: "message", T: Aoi, repeated: true },
    { no: 7, name: "events", kind: "message", T: Events },
    { no: 8, name: "statistics", kind: "message", T: Statistics },
  ],
);

/**
 * @generated from message city.comm.output.v1.OutputResponse
 */
export const OutputResponse = proto3.makeMessageType(
  "city.comm.output.v1.OutputResponse",
  [],
);

