// @generated by protoc-gen-es v1.10.0
// @generated from file city/comm/output/v1/output.proto (package city.comm.output.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum city.comm.output.v1.NodeStatus
 */
export const NodeStatus = /*@__PURE__*/ proto3.makeEnum(
  "city.comm.output.v1.NodeStatus",
  [
    {no: 0, name: "NODE_STATUS_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "NODE_STATUS_OK", localName: "OK"},
    {no: 2, name: "NODE_STATUS_BATTERY", localName: "BATTERY"},
    {no: 3, name: "NODE_STATUS_FAILURE", localName: "FAILURE"},
    {no: 4, name: "NODE_STATUS_RUINED", localName: "RUINED"},
  ],
);

/**
 * 用户连接基站状态
 *
 * @generated from enum city.comm.output.v1.PersonConnectStatus
 */
export const PersonConnectStatus = /*@__PURE__*/ proto3.makeEnum(
  "city.comm.output.v1.PersonConnectStatus",
  [
    {no: 0, name: "PERSON_CONNECT_STATUS_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "PERSON_CONNECT_STATUS_OK", localName: "OK"},
    {no: 2, name: "PERSON_CONNECT_STATUS_OUTAGE", localName: "OUTAGE"},
  ],
);

/**
 * 用户需求满足状态
 *
 * @generated from enum city.comm.output.v1.PersonDemandStatus
 */
export const PersonDemandStatus = /*@__PURE__*/ proto3.makeEnum(
  "city.comm.output.v1.PersonDemandStatus",
  [
    {no: 0, name: "PERSON_DEMAND_STATUS_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "PERSON_DEMAND_STATUS_SATISFIED", localName: "SATISFIED"},
    {no: 2, name: "PERSON_DEMAND_STATUS_UNSATISFIED", localName: "UNSATISFIED"},
    {no: 3, name: "PERSON_DEMAND_STATUS_NO", localName: "NO"},
  ],
);

/**
 * 设备状态
 *
 * @generated from message city.comm.output.v1.Node
 */
export const Node = /*@__PURE__*/ proto3.makeMessageType(
  "city.comm.output.v1.Node",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(NodeStatus) },
    { no: 3, name: "battery_remaining_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
  ],
);

/**
 * 基站状态
 *
 * @generated from message city.comm.output.v1.BaseStation
 */
export const BaseStation = /*@__PURE__*/ proto3.makeMessageType(
  "city.comm.output.v1.BaseStation",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "demand_flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "actual_flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "num_agents", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "overload", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "unsatisfied_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "satisfied_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "outage_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 9, name: "active_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "transmit_power", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 信号情况
 *
 * @generated from message city.comm.output.v1.Signal
 */
export const Signal = /*@__PURE__*/ proto3.makeMessageType(
  "city.comm.output.v1.Signal",
  () => [
    { no: 1, name: "num_rows", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "num_columns", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "strength", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, repeated: true },
    { no: 4, name: "base_station_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 5, name: "freq_range_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ],
);

/**
 * 人（可见的，交通模拟或室内模拟正在计算位置变动的）
 *
 * @generated from message city.comm.output.v1.Person
 */
export const Person = /*@__PURE__*/ proto3.makeMessageType(
  "city.comm.output.v1.Person",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "demand_rate", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "actual_rate", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "connect_status", kind: "enum", T: proto3.getEnumType(PersonConnectStatus) },
    { no: 9, name: "demand_status", kind: "enum", T: proto3.getEnumType(PersonDemandStatus) },
    { no: 5, name: "strength", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 6, name: "base_station_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "freq_range_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 8, name: "received_power", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * @generated from message city.comm.output.v1.Aoi
 */
export const Aoi = /*@__PURE__*/ proto3.makeMessageType(
  "city.comm.output.v1.Aoi",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "demand_flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "actual_flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "outage_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "satisfied_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "unsatisfied_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "active_user_num", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.comm.output.v1.Statistics
 */
export const Statistics = /*@__PURE__*/ proto3.makeMessageType(
  "city.comm.output.v1.Statistics",
  () => [
    { no: 1, name: "num_satisfied_agents", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "num_unsatisfied_agents", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "num_outage_agents", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "num_active_agents", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "demand_flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 6, name: "actual_flow", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 7, name: "num_base_station", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "num_ok_base_station", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 9, name: "num_ruined_base_station", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "num_stopped_base_station", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 11, name: "num_overloaded_base_station", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 12, name: "num_gateway", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 13, name: "num_ok_gateway", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 14, name: "num_ruined_gateway", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 15, name: "num_stopped_gateway", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 16, name: "num_overloaded_gateway", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 17, name: "num_battery_gateway", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 18, name: "power_consumption", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

