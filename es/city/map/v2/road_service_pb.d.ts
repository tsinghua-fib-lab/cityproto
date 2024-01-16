// @generated by protoc-gen-es v1.6.0
// @generated from file city/map/v2/road_service.proto (package city.map.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Events } from "../../event/v1/event_pb.js";
import type { LaneState } from "./lane_service_pb.js";

/**
 * 道路拥堵情况
 *
 * @generated from enum city.map.v2.RoadLevel
 */
export declare enum RoadLevel {
  /**
   * 未指定
   *
   * @generated from enum value: ROAD_LEVEL_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 畅通
   *
   * @generated from enum value: ROAD_LEVEL_CLEAR = 1;
   */
  CLEAR = 1,

  /**
   * 轻度拥堵
   *
   * @generated from enum value: ROAD_LEVEL_LIGHT_LOAD = 2;
   */
  LIGHT_LOAD = 2,

  /**
   * 中度拥堵
   *
   * @generated from enum value: ROAD_LEVEL_MEDIUM_LOAD = 3;
   */
  MEDIUM_LOAD = 3,

  /**
   * 重度拥堵
   *
   * @generated from enum value: ROAD_LEVEL_HEAVY_LOAD = 4;
   */
  HEAVY_LOAD = 4,

  /**
   * 极端拥堵
   *
   * @generated from enum value: ROAD_LEVEL_OVERLOAD = 5;
   */
  OVERLOAD = 5,

  /**
   * 限行
   *
   * @generated from enum value: ROAD_LEVEL_RESTRICTED = 6;
   */
  RESTRICTED = 6,
}

/**
 * @generated from enum city.map.v2.InterruptionReason
 */
export declare enum InterruptionReason {
  /**
   * @generated from enum value: INTERRUPTION_REASON_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: INTERRUPTION_REASON_RUINED = 1;
   */
  RUINED = 1,

  /**
   * @generated from enum value: INTERRUPTION_REASON_CASCADE = 2;
   */
  CASCADE = 2,

  /**
   * @generated from enum value: INTERRUPTION_REASON_CONGESTION = 3;
   */
  CONGESTION = 3,
}

/**
 * 查询道路信息请求
 *
 * @generated from message city.map.v2.GetRoadRequest
 */
export declare class GetRoadRequest extends Message<GetRoadRequest> {
  /**
   * 指定查询的道路ID列表，为空代表查询所有道路
   *
   * @generated from field: repeated int32 road_ids = 1;
   */
  roadIds: number[];

  /**
   * 是否要排除车道信息
   *
   * @generated from field: bool exclude_lane = 2;
   */
  excludeLane: boolean;

  /**
   * 是否要排除车道上的人的信息（仅在包含车道信息时有效）
   *
   * @generated from field: bool exclude_person = 3;
   */
  excludePerson: boolean;

  constructor(data?: PartialMessage<GetRoadRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.map.v2.GetRoadRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRoadRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRoadRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRoadRequest;

  static equals(a: GetRoadRequest | PlainMessage<GetRoadRequest> | undefined, b: GetRoadRequest | PlainMessage<GetRoadRequest> | undefined): boolean;
}

/**
 * 查询道路信息响应
 *
 * @generated from message city.map.v2.GetRoadResponse
 */
export declare class GetRoadResponse extends Message<GetRoadResponse> {
  /**
   * 道路信息列表
   *
   * @generated from field: repeated city.map.v2.RoadState states = 1;
   */
  states: RoadState[];

  constructor(data?: PartialMessage<GetRoadResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.map.v2.GetRoadResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRoadResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRoadResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRoadResponse;

  static equals(a: GetRoadResponse | PlainMessage<GetRoadResponse> | undefined, b: GetRoadResponse | PlainMessage<GetRoadResponse> | undefined): boolean;
}

/**
 * @generated from message city.map.v2.GetRuinInfoRequest
 */
export declare class GetRuinInfoRequest extends Message<GetRuinInfoRequest> {
  constructor(data?: PartialMessage<GetRuinInfoRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.map.v2.GetRuinInfoRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRuinInfoRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRuinInfoRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRuinInfoRequest;

  static equals(a: GetRuinInfoRequest | PlainMessage<GetRuinInfoRequest> | undefined, b: GetRuinInfoRequest | PlainMessage<GetRuinInfoRequest> | undefined): boolean;
}

/**
 * @generated from message city.map.v2.RuinInfo
 */
export declare class RuinInfo extends Message<RuinInfo> {
  /**
   * 损坏数量
   *
   * @generated from field: int32 num = 1;
   */
  num: number;

  /**
   * 损坏占比
   *
   * @generated from field: double ratio = 2;
   */
  ratio: number;

  constructor(data?: PartialMessage<RuinInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.map.v2.RuinInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RuinInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RuinInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RuinInfo;

  static equals(a: RuinInfo | PlainMessage<RuinInfo> | undefined, b: RuinInfo | PlainMessage<RuinInfo> | undefined): boolean;
}

/**
 * @generated from message city.map.v2.GetRuinInfoResponse
 */
export declare class GetRuinInfoResponse extends Message<GetRuinInfoResponse> {
  /**
   * 三级级损伤信息
   *
   * @generated from field: city.map.v2.RuinInfo one = 1;
   */
  one?: RuinInfo;

  /**
   * @generated from field: city.map.v2.RuinInfo two = 2;
   */
  two?: RuinInfo;

  /**
   * @generated from field: city.map.v2.RuinInfo three = 3;
   */
  three?: RuinInfo;

  constructor(data?: PartialMessage<GetRuinInfoResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.map.v2.GetRuinInfoResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRuinInfoResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRuinInfoResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRuinInfoResponse;

  static equals(a: GetRuinInfoResponse | PlainMessage<GetRuinInfoResponse> | undefined, b: GetRuinInfoResponse | PlainMessage<GetRuinInfoResponse> | undefined): boolean;
}

/**
 * @generated from message city.map.v2.GetEventsRequest
 */
export declare class GetEventsRequest extends Message<GetEventsRequest> {
  constructor(data?: PartialMessage<GetEventsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.map.v2.GetEventsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetEventsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetEventsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetEventsRequest;

  static equals(a: GetEventsRequest | PlainMessage<GetEventsRequest> | undefined, b: GetEventsRequest | PlainMessage<GetEventsRequest> | undefined): boolean;
}

/**
 * @generated from message city.map.v2.GetEventsResponse
 */
export declare class GetEventsResponse extends Message<GetEventsResponse> {
  /**
   * @generated from field: city.event.v1.Events events = 1;
   */
  events?: Events;

  constructor(data?: PartialMessage<GetEventsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.map.v2.GetEventsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetEventsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetEventsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetEventsResponse;

  static equals(a: GetEventsResponse | PlainMessage<GetEventsResponse> | undefined, b: GetEventsResponse | PlainMessage<GetEventsResponse> | undefined): boolean;
}

/**
 * 道路状态
 *
 * @generated from message city.map.v2.RoadState
 */
export declare class RoadState extends Message<RoadState> {
  /**
   * 道路ID
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * 道路平均速度（m/s）
   *
   * @generated from field: double avg_v = 4;
   */
  avgV: number;

  /**
   * 道路拥堵情况
   *
   * @generated from field: city.map.v2.RoadLevel level = 2;
   */
  level: RoadLevel;

  /**
   * 道路中断原因
   *
   * @generated from field: city.map.v2.InterruptionReason reason = 3;
   */
  reason: InterruptionReason;

  /**
   * 车道情况
   *
   * @generated from field: repeated city.map.v2.LaneState lanes = 5;
   */
  lanes: LaneState[];

  constructor(data?: PartialMessage<RoadState>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.map.v2.RoadState";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RoadState;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RoadState;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RoadState;

  static equals(a: RoadState | PlainMessage<RoadState> | undefined, b: RoadState | PlainMessage<RoadState> | undefined): boolean;
}

