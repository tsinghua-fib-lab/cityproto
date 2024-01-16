// @generated by protoc-gen-es v1.6.0
// @generated from file city/agent/v2/motion.proto (package city.agent.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Position } from "../../geo/v2/geo_pb.js";

/**
 * Agent（人）的运行时状态
 *
 * @generated from enum city.agent.v2.Status
 */
export declare enum Status {
  /**
   * 未指定
   *
   * @generated from enum value: STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 没有移动行为
   *
   * @generated from enum value: STATUS_SLEEP = 1;
   */
  SLEEP = 1,

  /**
   * 开车
   *
   * @generated from enum value: STATUS_DRIVING = 2;
   */
  DRIVING = 2,

  /**
   * 步行
   *
   * @generated from enum value: STATUS_WALKING = 3;
   */
  WALKING = 3,

  /**
   * 室内行人
   *
   * @generated from enum value: STATUS_CROWD = 4;
   */
  CROWD = 4,

  /**
   * 乘客
   *
   * @generated from enum value: STATUS_PASSENGER = 5;
   */
  PASSENGER = 5,

  /**
   * 等待路径规划
   *
   * @generated from enum value: STATUS_WAIT_ROUTE = 6;
   */
  WAIT_ROUTE = 6,
}

/**
 * Agent（人）的运动状态
 *
 * @generated from message city.agent.v2.AgentMotion
 */
export declare class AgentMotion extends Message<AgentMotion> {
  /**
   * ID
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * 状态
   *
   * @generated from field: city.agent.v2.Status status = 2;
   */
  status: Status;

  /**
   * 位置（包含逻辑位置、XY位置、经纬度位置）
   *
   * @generated from field: city.geo.v2.Position position = 3;
   */
  position?: Position;

  /**
   * 速度
   *
   * @generated from field: double v = 4;
   */
  v: number;

  /**
   * 方向角（atan2计算得到的弧度）
   *
   * @generated from field: double direction = 5;
   */
  direction: number;

  /**
   * 活动描述
   *
   * @generated from field: string activity = 6;
   */
  activity: string;

  constructor(data?: PartialMessage<AgentMotion>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.agent.v2.AgentMotion";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AgentMotion;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AgentMotion;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AgentMotion;

  static equals(a: AgentMotion | PlainMessage<AgentMotion> | undefined, b: AgentMotion | PlainMessage<AgentMotion> | undefined): boolean;
}
