// @generated by protoc-gen-es v1.10.0
// @generated from file city/water/output/v1/output.proto (package city.water.output.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { LongLatPosition } from "../../../geo/v2/geo_pb.js";

/**
 * @generated from enum city.water.output.v1.LinkType
 */
export const LinkType = /*@__PURE__*/ proto3.makeEnum(
  "city.water.output.v1.LinkType",
  [
    {no: 0, name: "LINK_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "LINK_TYPE_PIPE", localName: "PIPE"},
    {no: 2, name: "LINK_TYPE_PUMP", localName: "PUMP"},
  ],
);

/**
 * 宏观道路水深
 *
 * @generated from message city.water.output.v1.Road
 */
export const Road = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.Road",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "depth", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 微观水深点
 *
 * @generated from message city.water.output.v1.RoadFlood
 */
export const RoadFlood = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.RoadFlood",
  () => [
    { no: 1, name: "position", kind: "message", T: LongLatPosition },
    { no: 2, name: "depth", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * @generated from message city.water.output.v1.DetailedRoad
 */
export const DetailedRoad = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.DetailedRoad",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "depths", kind: "message", T: RoadFlood, repeated: true },
  ],
);

/**
 * 节点状态
 *
 * @generated from message city.water.output.v1.Node
 */
export const Node = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.Node",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "head", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 边的状态
 *
 * @generated from message city.water.output.v1.Link
 */
export const Link = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.Link",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(LinkType) },
    { no: 3, name: "flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "ok", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message city.water.output.v1.Aoi
 */
export const Aoi = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.Aoi",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "unsatisfied_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "unsatisfied_ratio", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "demand", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 5, name: "supply", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 排水基础信息
 *
 * @generated from message city.water.output.v1.DrainageBasicInfo
 */
export const DrainageBasicInfo = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.DrainageBasicInfo",
  () => [
    { no: 1, name: "average_power", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "undrained_volume", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "drained_volume", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "average_flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 5, name: "flooded_volume", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 供水基础信息
 *
 * @generated from message city.water.output.v1.SupplyBasicInfo
 */
export const SupplyBasicInfo = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.SupplyBasicInfo",
  () => [
    { no: 1, name: "average_power", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "average_flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 供水需求及满足情况
 *
 * @generated from message city.water.output.v1.SupplyDemandStatistics
 */
export const SupplyDemandStatistics = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.SupplyDemandStatistics",
  () => [
    { no: 1, name: "persons_demand", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "unsatisfied_persons", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "unsatisfied_persons_ratio", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "aois_demand", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 5, name: "unsatisfied_aois", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "unsatisfied_aois_ratio", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 水泵损坏情况
 *
 * @generated from message city.water.output.v1.FailureStatistics
 */
export const FailureStatistics = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.FailureStatistics",
  () => [
    { no: 1, name: "failure_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "normal_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "failure_ratio", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 排水指标
 *
 * @generated from message city.water.output.v1.DrainageMetrics
 */
export const DrainageMetrics = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.DrainageMetrics",
  () => [
    { no: 1, name: "drainage_basic_info", kind: "message", T: DrainageBasicInfo },
    { no: 2, name: "load_ratio", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "failure_statistics", kind: "message", T: FailureStatistics },
  ],
);

/**
 * 供水指标
 *
 * @generated from message city.water.output.v1.SupplyMetrics
 */
export const SupplyMetrics = /*@__PURE__*/ proto3.makeMessageType(
  "city.water.output.v1.SupplyMetrics",
  () => [
    { no: 1, name: "supply_basic_info", kind: "message", T: SupplyBasicInfo },
    { no: 2, name: "supply_demand_statistics", kind: "message", T: SupplyDemandStatistics },
    { no: 3, name: "load_ratio", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "failure_statistics", kind: "message", T: FailureStatistics },
  ],
);

