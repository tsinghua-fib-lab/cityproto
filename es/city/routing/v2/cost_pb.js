// @generated by protoc-gen-es v1.10.0
// @generated from file city/routing/v2/cost.proto (package city.routing.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * 路径成本设置
 * Route cost settings
 *
 * @generated from message city.routing.v2.Cost
 */
export const Cost = /*@__PURE__*/ proto3.makeMessageType(
  "city.routing.v2.Cost",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "cost", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
  ],
);

