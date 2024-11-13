// @generated by protoc-gen-es v1.10.0
// @generated from file city/person/v1/motion.proto (package city.person.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Position } from "../../geo/v2/geo_pb.js";

/**
 * Person（人）的运行时状态
 * Person's runtime state
 *
 * @generated from enum city.person.v1.Status
 */
export declare enum Status {
  /**
   * 未指定
   * unspecified
   *
   * @generated from enum value: STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 没有移动行为
   * no mobility behaviors
   *
   * @generated from enum value: STATUS_SLEEP = 1;
   */
  SLEEP = 1,

  /**
   * 开车
   * driving
   *
   * @generated from enum value: STATUS_DRIVING = 2;
   */
  DRIVING = 2,

  /**
   * 步行
   * walking
   *
   * @generated from enum value: STATUS_WALKING = 3;
   */
  WALKING = 3,

  /**
   * 室内行人
   * indoor pedestrian
   *
   * @generated from enum value: STATUS_CROWD = 4;
   */
  CROWD = 4,

  /**
   * 乘客
   * vehicle passenger
   *
   * @generated from enum value: STATUS_PASSENGER = 5;
   */
  PASSENGER = 5,

  /**
   * 等待路径规划
   * wait for path routing
   *
   * @generated from enum value: STATUS_WAIT_ROUTE = 6;
   */
  WAIT_ROUTE = 6,

  /**
   * 等待公交车
   * wait for bus coming
   *
   * @generated from enum value: STATUS_WAIT_BUS = 7;
   */
  WAIT_BUS = 7,

  /**
   * 操控轨道交通
   * operating rail transit
   *
   * @generated from enum value: STATUS_RAIL_TRANSIT = 8;
   */
  RAIL_TRANSIT = 8,
}

/**
 * Person（人）的运动状态
 * Person's motion state
 *
 * @generated from message city.person.v1.PersonMotion
 */
export declare class PersonMotion extends Message<PersonMotion> {
  /**
   * ID
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * 状态
   * status
   *
   * @generated from field: city.person.v1.Status status = 2;
   */
  status: Status;

  /**
   * 位置（包含逻辑位置、XY位置、经纬度位置）
   * Position (including logical position, XY position, longitude and latitude position)
   *
   * @generated from field: city.geo.v2.Position position = 3;
   */
  position?: Position;

  /**
   * speed
   *
   * @generated from field: double v = 4;
   */
  v: number;

  /**
   * 方向角（atan2计算得到的弧度）
   * Direction angle (radians calculated by atan2)
   *
   * @generated from field: double direction = 5;
   */
  direction: number;

  /**
   * 活动描述
   * activity descriptions
   *
   * @generated from field: string activity = 6;
   */
  activity: string;

  /**
   * 长度
   * length
   *
   * @generated from field: double l = 7;
   */
  l: number;

  constructor(data?: PartialMessage<PersonMotion>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.PersonMotion";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PersonMotion;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PersonMotion;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PersonMotion;

  static equals(a: PersonMotion | PlainMessage<PersonMotion> | undefined, b: PersonMotion | PlainMessage<PersonMotion> | undefined): boolean;
}

