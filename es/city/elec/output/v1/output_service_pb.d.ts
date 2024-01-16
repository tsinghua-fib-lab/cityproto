// @generated by protoc-gen-es v1.6.0
// @generated from file city/elec/output/v1/output_service.proto (package city.elec.output.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Aoi, AveragePower } from "./output_pb.js";
import type { Events } from "../../../event/v1/event_pb.js";

/**
 * @generated from message city.elec.output.v1.OutputRequest
 */
export declare class OutputRequest extends Message<OutputRequest> {
  /**
   * @generated from field: int32 step = 1;
   */
  step: number;

  /**
   * 被破坏的节点ID
   *
   * @generated from field: repeated int32 ruined_ids = 2;
   */
  ruinedIds: number[];

  /**
   * 由于其他节点被破坏导致停止工作的节点ID
   *
   * @generated from field: repeated int32 stopped_ids = 3;
   */
  stoppedIds: number[];

  /**
   * AOI相关的数据
   *
   * @generated from field: repeated city.elec.output.v1.Aoi aois = 4;
   */
  aois: Aoi[];

  /**
   * 事件数据
   *
   * @generated from field: city.event.v1.Events events = 5;
   */
  events?: Events;

  /**
   * 居民用电需求不满足比例
   *
   * @generated from field: double resident_unsatisfied_ratio = 6;
   */
  residentUnsatisfiedRatio: number;

  /**
   * 居民总用电需求,MW
   *
   * @generated from field: double resident_demand = 7;
   */
  residentDemand: number;

  /**
   * aoi用电需求不满足比例
   *
   * @generated from field: double aoi_unsatisfied_ratio = 8;
   */
  aoiUnsatisfiedRatio: number;

  /**
   * 不满足用电的aoi数量,个数
   *
   * @generated from field: int32 aoi_unsatisfied_num = 9;
   */
  aoiUnsatisfiedNum: number;

  /**
   * aoi总用电需求,MW
   *
   * @generated from field: double aoi_demand = 10;
   */
  aoiDemand: number;

  /**
   * 各类变压器当前的平均承受功率，单位为MW
   *
   * @generated from field: city.elec.output.v1.AveragePower average_powers = 11;
   */
  averagePowers?: AveragePower;

  constructor(data?: PartialMessage<OutputRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.elec.output.v1.OutputRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OutputRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OutputRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OutputRequest;

  static equals(a: OutputRequest | PlainMessage<OutputRequest> | undefined, b: OutputRequest | PlainMessage<OutputRequest> | undefined): boolean;
}

/**
 * @generated from message city.elec.output.v1.OutputResponse
 */
export declare class OutputResponse extends Message<OutputResponse> {
  constructor(data?: PartialMessage<OutputResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.elec.output.v1.OutputResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OutputResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OutputResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OutputResponse;

  static equals(a: OutputResponse | PlainMessage<OutputResponse> | undefined, b: OutputResponse | PlainMessage<OutputResponse> | undefined): boolean;
}

