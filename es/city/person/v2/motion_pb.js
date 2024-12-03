// @generated by protoc-gen-es v1.10.0
// @generated from file city/person/v2/motion.proto (package city.person.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Position } from "../../geo/v2/geo_pb.js";

/**
 * Person（人）的运行时状态
 * Person's runtime state
 *
 * @generated from enum city.person.v2.Status
 */
export const Status = /*@__PURE__*/ proto3.makeEnum(
  "city.person.v2.Status",
  [
    {no: 0, name: "STATUS_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "STATUS_SLEEP", localName: "SLEEP"},
    {no: 2, name: "STATUS_DRIVING", localName: "DRIVING"},
    {no: 3, name: "STATUS_WALKING", localName: "WALKING"},
    {no: 4, name: "STATUS_CROWD", localName: "CROWD"},
    {no: 5, name: "STATUS_PASSENGER", localName: "PASSENGER"},
    {no: 6, name: "STATUS_WAIT_ROUTE", localName: "WAIT_ROUTE"},
    {no: 7, name: "STATUS_WAIT_BUS", localName: "WAIT_BUS"},
    {no: 8, name: "STATUS_RAIL_TRANSIT", localName: "RAIL_TRANSIT"},
    {no: 9, name: "STATUS_WAIT_TAXI", localName: "WAIT_TAXI"},
  ],
);

/**
 * Person（人）的运动状态
 * Person's motion state
 *
 * @generated from message city.person.v2.PersonMotion
 */
export const PersonMotion = /*@__PURE__*/ proto3.makeMessageType(
  "city.person.v2.PersonMotion",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(Status) },
    { no: 3, name: "position", kind: "message", T: Position },
    { no: 4, name: "v", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 5, name: "direction", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 6, name: "activity", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "l", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

