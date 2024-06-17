// @generated by protoc-gen-es v1.6.0
// @generated from file city/person/v1/runtime.proto (package city.person.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message city.person.v1.BaseRuntime
 */
export declare class BaseRuntime extends Message<BaseRuntime> {
  /**
   * x坐标
   * x coordinate
   *
   * @generated from field: double x = 1;
   */
  x: number;

  /**
   * y坐标
   * y coordinate
   *
   * @generated from field: double y = 2;
   */
  y: number;

  /**
   * 车速
   * velocity (m/s)
   *
   * @generated from field: double v = 3;
   */
  v: number;

  /**
   * 方向(rad)，逆时针方向增加，0为正东方向
   * direction (rad), increase in counterclockwise direction, 0 is east
   *
   * @generated from field: double direction = 4;
   */
  direction: number;

  /**
   * 长度
   * length
   *
   * @generated from field: double l = 5;
   */
  l: number;

  constructor(data?: PartialMessage<BaseRuntime>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.BaseRuntime";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BaseRuntime;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BaseRuntime;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BaseRuntime;

  static equals(a: BaseRuntime | PlainMessage<BaseRuntime> | undefined, b: BaseRuntime | PlainMessage<BaseRuntime> | undefined): boolean;
}

/**
 * @generated from message city.person.v1.BaseRuntimeOnRoad
 */
export declare class BaseRuntimeOnRoad extends Message<BaseRuntimeOnRoad> {
  /**
   * 车道ID
   * lane ID
   *
   * @generated from field: int32 lane_id = 1;
   */
  laneId: number;

  /**
   * 车道中心线坐标
   * lane center line coordinate
   *
   * @generated from field: double s = 2;
   */
  s: number;

  constructor(data?: PartialMessage<BaseRuntimeOnRoad>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.BaseRuntimeOnRoad";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BaseRuntimeOnRoad;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BaseRuntimeOnRoad;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BaseRuntimeOnRoad;

  static equals(a: BaseRuntimeOnRoad | PlainMessage<BaseRuntimeOnRoad> | undefined, b: BaseRuntimeOnRoad | PlainMessage<BaseRuntimeOnRoad> | undefined): boolean;
}
