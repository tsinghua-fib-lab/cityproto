// @generated by protoc-gen-es v1.10.0
// @generated from file city/event/v1/event.proto (package city.event.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum city.event.v1.EventType
 */
export const EventType = /*@__PURE__*/ proto3.makeEnum(
  "city.event.v1.EventType",
  [
    {no: 0, name: "EVENT_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "EVENT_TYPE_HEAVY_RAIN", localName: "HEAVY_RAIN"},
    {no: 2, name: "EVENT_TYPE_MILITARY_STRIKE", localName: "MILITARY_STRIKE"},
    {no: 3, name: "EVENT_TYPE_TRAFFIC_CONGESTION", localName: "TRAFFIC_CONGESTION"},
    {no: 4, name: "EVENT_TYPE_TRAFFIC_LANE_RESTRICTION", localName: "TRAFFIC_LANE_RESTRICTION"},
    {no: 5, name: "EVENT_TYPE_TRAFFIC_BAD_LIGHT", localName: "TRAFFIC_BAD_LIGHT"},
    {no: 6, name: "EVENT_TYPE_ELEC_RUINED_TRANSFORMER", localName: "ELEC_RUINED_TRANSFORMER"},
    {no: 7, name: "EVENT_TYPE_WATER_RUINED_PUMP", localName: "WATER_RUINED_PUMP"},
    {no: 8, name: "EVENT_TYPE_WATER_STOPPED_PUMP", localName: "WATER_STOPPED_PUMP"},
    {no: 9, name: "EVENT_TYPE_COMM_RUINED_BASE_STATION", localName: "COMM_RUINED_BASE_STATION"},
    {no: 10, name: "EVENT_TYPE_COMM_STOPPED_BASE_STATION", localName: "COMM_STOPPED_BASE_STATION"},
    {no: 11, name: "EVENT_TYPE_COMM_OVERLOAD_BASE_STATION", localName: "COMM_OVERLOAD_BASE_STATION"},
  ],
);

/**
 * @generated from message city.event.v1.Event
 */
export const Event = /*@__PURE__*/ proto3.makeMessageType(
  "city.event.v1.Event",
  () => [
    { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(EventType) },
    { no: 2, name: "level", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "step", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.event.v1.Events
 */
export const Events = /*@__PURE__*/ proto3.makeMessageType(
  "city.event.v1.Events",
  () => [
    { no: 1, name: "events", kind: "message", T: Event, repeated: true },
  ],
);

