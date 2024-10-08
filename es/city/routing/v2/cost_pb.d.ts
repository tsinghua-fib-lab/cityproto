// @generated by protoc-gen-es v1.10.0
// @generated from file city/routing/v2/cost.proto (package city.routing.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * 路径成本设置
 * Route cost settings
 *
 * @generated from message city.routing.v2.Cost
 */
export declare class Cost extends Message<Cost> {
  /**
   * 目标拓扑元素（只支持道路Road）
   * Target topology element (only supports roads)
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * 路径成本（单位：秒）
   * Path cost (in seconds)
   *
   * @generated from field: double cost = 2;
   */
  cost: number;

  /**
   * 设置的时间（单位：秒）
   * Set time (in seconds)
   * 即设置几点几分的道路通行成本为cost
   * That is, set the cost as the value at what time
   * 为空表示设置全天通行成本均为cost
   * If empty, it means that the all-day cost is set to the value.
   *
   * @generated from field: optional double time = 3;
   */
  time?: number;

  constructor(data?: PartialMessage<Cost>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.routing.v2.Cost";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Cost;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Cost;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Cost;

  static equals(a: Cost | PlainMessage<Cost> | undefined, b: Cost | PlainMessage<Cost> | undefined): boolean;
}

