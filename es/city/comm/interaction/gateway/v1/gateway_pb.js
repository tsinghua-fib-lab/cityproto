// @generated by protoc-gen-es v1.10.0
// @generated from file city/comm/interaction/gateway/v1/gateway.proto (package city.comm.interaction.gateway.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum city.comm.interaction.gateway.v1.Reason
 */
export const Reason = /*@__PURE__*/ proto3.makeEnum(
  "city.comm.interaction.gateway.v1.Reason",
  [
    {no: 0, name: "REASON_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "REASON_RUIN", localName: "RUIN"},
    {no: 2, name: "REASON_CASCADE", localName: "CASCADE"},
  ],
);

/**
 * @generated from message city.comm.interaction.gateway.v1.Station
 */
export const Station = /*@__PURE__*/ proto3.makeMessageType(
  "city.comm.interaction.gateway.v1.Station",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "status", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "reason", kind: "enum", T: proto3.getEnumType(Reason) },
  ],
);

