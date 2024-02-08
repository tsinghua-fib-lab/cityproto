// @generated by protoc-gen-es v1.6.0
// @generated from file city/trip/v2/trip.proto (package city.trip.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Position } from "../../geo/v2/geo_pb.js";
import { Journey } from "../../routing/v2/routing_pb.js";

/**
 * 出行偏好
 * Preferred trip travel mode
 *
 * @generated from enum city.trip.v2.TripMode
 */
export const TripMode = proto3.makeEnum(
  "city.trip.v2.TripMode",
  [
    {no: 0, name: "TRIP_MODE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "TRIP_MODE_WALK_ONLY", localName: "WALK_ONLY"},
    {no: 2, name: "TRIP_MODE_DRIVE_ONLY", localName: "DRIVE_ONLY"},
    {no: 4, name: "TRIP_MODE_BUS_WALK", localName: "BUS_WALK"},
    {no: 5, name: "TRIP_MODE_BIKE_WALK", localName: "BIKE_WALK"},
  ],
);

/**
 * 出行
 * Trip
 *
 * @generated from message city.trip.v2.Trip
 */
export const Trip = proto3.makeMessageType(
  "city.trip.v2.Trip",
  () => [
    { no: 1, name: "mode", kind: "enum", T: proto3.getEnumType(TripMode) },
    { no: 2, name: "end", kind: "message", T: Position },
    { no: 3, name: "departure_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 4, name: "wait_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 5, name: "arrival_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 6, name: "activity", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 7, name: "routes", kind: "message", T: Journey, repeated: true },
  ],
);

/**
 * 时刻表
 * Schedule
 * 关于出发时间的说明如下：
 * The explanation about the departure time is as follows:
 * 1. Schedule的开始时刻是 departure_time 或者 参考时刻+wait_time，
 * 1. The start time of the Schedule is either departure_time or reference time+wait_time,
 *    参考时刻定义为上一Schedule的结束时刻(即它最后一个Trip的结束时刻)，
 *    The reference time is defined as the end time of the previous Schedule (i.e. the end time of its last Trip),
 *    或者当它为第一个Schedule时定义为Person更新Schedule后的首次Update
 *    Alternatively, when it is the first Schedule, it can be defined as the first Update time after Person updates the Schedule
 *    时刻(当有准确时间要求时建议直接指定departure_time)
 *    (it is recommended to specify departuretime directly when there is an accurate time requirement)
 * 2. Trip的开始时刻是 departure_time 或者 参考时刻+wait_time，参考
 * 2. The start time of the Trip is either departure_time or reference time+wait_time,
 *    时刻定义为上一Trip的结束时刻，或者当它为第一个Trip时定义为所属的
 *    The reference time is defined as the end time of the previous Trip, or when it is the first Trip,
 *    Schedule的开始时刻
 *    it is defined as the start time of the Schedule to which it belongs
 * 3. Person的实际运行时刻取决于Trip的开始时刻，例如它的首次运行是第一
 * 3. The actual running time of a Person depends on the start time of the Trip,
 *    个Schedule中第一个Trip的开始时刻
 *    for example, its first run is the start time of the first Trip in the first Schedule
 * FAQ
 * Q1: 同时指定Schedule和第一个Trip的departure_time会怎样？
 * Q1: What would happen if both the Schedule and the departuretime of the first Trip were specified simultaneously?
 * A1: 参照(2)，只看Trip的departure_time
 * A1: Referring to (2), only depend on the departuretime of Trip
 * Q2: Schedule和第一个Trip同时指定wait_time=10会怎样？
 * Q2: What would happen if both the Schedule and the first Trip were specified with wait_time=10 at the same time?
 * A2: 参照(2)，等待时间为10+10=20
 * A2: Referring to (2), the waiting time is 10+10=20
 *
 * @generated from message city.trip.v2.Schedule
 */
export const Schedule = proto3.makeMessageType(
  "city.trip.v2.Schedule",
  () => [
    { no: 1, name: "trips", kind: "message", T: Trip, repeated: true },
    { no: 2, name: "loop_count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "departure_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 4, name: "wait_time", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
  ],
);

