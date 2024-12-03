// @generated by protoc-gen-es v1.10.0
// @generated from file city/economy/v2/economy.proto (package city.economy.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum city.economy.v2.OrgType
 */
export const OrgType = /*@__PURE__*/ proto3.makeEnum(
  "city.economy.v2.OrgType",
  [
    {no: 0, name: "ORG_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ORG_TYPE_NBS", localName: "NBS"},
    {no: 2, name: "ORG_TYPE_FIRM", localName: "FIRM"},
    {no: 3, name: "ORG_TYPE_BANK", localName: "BANK"},
    {no: 4, name: "ORG_TYPE_GOVERNMENT", localName: "GOVERNMENT"},
  ],
);

/**
 * 组织
 * Organization
 *
 * @generated from message city.economy.v2.Org
 */
export const Org = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.Org",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(OrgType) },
    { no: 3, name: "nominal_gdp", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 4, name: "real_gdp", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 5, name: "unemployment", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 6, name: "wages", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 7, name: "prices", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 8, name: "inventory", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 9, name: "price", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 10, name: "currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 11, name: "interest_rate", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 12, name: "bracket_cutoffs", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 13, name: "bracket_rates", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.Agent
 */
export const Agent = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.Agent",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
  ],
);

/**
 * @generated from message city.economy.v2.EconomyEntities
 */
export const EconomyEntities = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.EconomyEntities",
  () => [
    { no: 1, name: "orgs", kind: "message", T: Org, repeated: true },
    { no: 2, name: "agents", kind: "message", T: Agent, repeated: true },
  ],
);

