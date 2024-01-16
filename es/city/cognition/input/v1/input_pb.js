// @generated by protoc-gen-es v1.6.0
// @generated from file city/cognition/input/v1/input.proto (package city.cognition.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * 社交网络连接关系
 *
 * @generated from message city.cognition.input.v1.Edge
 */
export const Edge = proto3.makeMessageType(
  "city.cognition.input.v1.Edge",
  () => [
    { no: 1, name: "source", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "target", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.cognition.input.v1.Edges
 */
export const Edges = proto3.makeMessageType(
  "city.cognition.input.v1.Edges",
  () => [
    { no: 1, name: "edges", kind: "message", T: Edge, repeated: true },
  ],
);
