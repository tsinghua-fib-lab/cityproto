// @generated by protoc-gen-es v1.10.0
// @generated from file city/economy/v2/economy.proto (package city.economy.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * Firm represents a company entity in the economic system
 * Firm 代表经济系统中的公司实体
 *
 * @generated from message city.economy.v2.Firm
 */
export const Firm = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.Firm",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "employees", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 3, name: "price", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 4, name: "inventory", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "demand", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 6, name: "sales", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 7, name: "currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * NBS (National Bureau of Statistics) represents the statistical authority
 * NBS (国家统计局) 代表负责经济数据统计的权威机构
 *
 * @generated from message city.economy.v2.NBS
 */
export const NBS = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.NBS",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "citizen_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 3, name: "nominal_gdp", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 4, name: "real_gdp", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 5, name: "unemployment", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 6, name: "wages", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 7, name: "prices", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 8, name: "working_hours", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 9, name: "depression", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 10, name: "consumption_currency", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 11, name: "income_currency", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 12, name: "locus_control", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 2 /* ScalarType.FLOAT */} },
    { no: 13, name: "citizen_agent_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 14, name: "currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * Government represents a government institution
 * Government 代表政府机构
 *
 * @generated from message city.economy.v2.Government
 */
export const Government = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.Government",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "citizen_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 3, name: "bracket_cutoffs", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 4, name: "bracket_rates", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 5, name: "currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * Bank represents a banking institution
 * Bank 代表银行机构
 *
 * @generated from message city.economy.v2.Bank
 */
export const Bank = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.Bank",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "citizen_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 3, name: "interest_rate", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 4, name: "currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * Agent represents an individual economic agent (e.g., a resident)
 * Agent 代表经济系统中的个体代理（如居民个人）
 *
 * @generated from message city.economy.v2.Agent
 */
export const Agent = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.Agent",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 3, name: "firm_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 4, name: "skill", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 5, name: "consumption", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
    { no: 6, name: "income", kind: "scalar", T: 2 /* ScalarType.FLOAT */, opt: true },
  ],
);

/**
 * EconomyEntities represents all entities in the economic system
 * EconomyEntities 代表经济系统中的所有实体
 *
 * @generated from message city.economy.v2.EconomyEntities
 */
export const EconomyEntities = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.EconomyEntities",
  () => [
    { no: 1, name: "firms", kind: "message", T: Firm, repeated: true },
    { no: 2, name: "nbs", kind: "message", T: NBS, repeated: true },
    { no: 3, name: "governments", kind: "message", T: Government, repeated: true },
    { no: 4, name: "banks", kind: "message", T: Bank, repeated: true },
    { no: 5, name: "agents", kind: "message", T: Agent, repeated: true },
  ],
);

