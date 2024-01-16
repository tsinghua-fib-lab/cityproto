// @generated by protoc-gen-es v1.6.0
// @generated from file city/comm/output/v1/output_service.proto (package city.comm.output.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Aoi, BaseStation, Node, Person, Signal, Statistics } from "./output_pb.js";
import type { Events } from "../../../event/v1/event_pb.js";

/**
 * @generated from message city.comm.output.v1.OutputRequest
 */
export declare class OutputRequest extends Message<OutputRequest> {
  /**
   * @generated from field: int32 step = 1;
   */
  step: number;

  /**
   * 设备状态
   *
   * @generated from field: repeated city.comm.output.v1.Node nodes = 2;
   */
  nodes: Node[];

  /**
   * @generated from field: repeated city.comm.output.v1.BaseStation base_stations = 3;
   */
  baseStations: BaseStation[];

  /**
   * 五环区域信号强度热力图（500m网格形式呈现）
   *
   * @generated from field: city.comm.output.v1.Signal signal_heatmap = 4;
   */
  signalHeatmap?: Signal;

  /**
   * 国贸区域信号强度热力图（50m网格形式呈现）
   *
   * @generated from field: city.comm.output.v1.Signal small_signal_heatmap = 9;
   */
  smallSignalHeatmap?: Signal;

  /**
   * TODO(张钧): 基站和人的连接怎么表示？
   * 人相关的数据
   *
   * @generated from field: repeated city.comm.output.v1.Person persons = 5;
   */
  persons: Person[];

  /**
   * AOI相关的数据
   *
   * @generated from field: repeated city.comm.output.v1.Aoi aois = 6;
   */
  aois: Aoi[];

  /**
   * @generated from field: city.event.v1.Events events = 7;
   */
  events?: Events;

  /**
   * 统计值
   *
   * @generated from field: city.comm.output.v1.Statistics statistics = 8;
   */
  statistics?: Statistics;

  constructor(data?: PartialMessage<OutputRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.output.v1.OutputRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OutputRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OutputRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OutputRequest;

  static equals(a: OutputRequest | PlainMessage<OutputRequest> | undefined, b: OutputRequest | PlainMessage<OutputRequest> | undefined): boolean;
}

/**
 * @generated from message city.comm.output.v1.OutputResponse
 */
export declare class OutputResponse extends Message<OutputResponse> {
  constructor(data?: PartialMessage<OutputResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.output.v1.OutputResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OutputResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OutputResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OutputResponse;

  static equals(a: OutputResponse | PlainMessage<OutputResponse> | undefined, b: OutputResponse | PlainMessage<OutputResponse> | undefined): boolean;
}

