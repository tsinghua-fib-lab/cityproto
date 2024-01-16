// @generated by protoc-gen-es v1.6.0
// @generated from file city/comm/input/v1/comm.proto (package city.comm.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Position } from "../../../geo/v2/geo_pb.js";

/**
 * 本文件描述通信部分拓扑结构
 * 三种节点类型
 *
 * @generated from enum city.comm.input.v1.NodeType
 */
export declare enum NodeType {
  /**
   * @generated from enum value: NODE_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: NODE_TYPE_INTERNET = 1;
   */
  INTERNET = 1,

  /**
   * @generated from enum value: NODE_TYPE_GATEWAY = 2;
   */
  GATEWAY = 2,

  /**
   * @generated from enum value: NODE_TYPE_BASE_STATION = 3;
   */
  BASE_STATION = 3,
}

/**
 * @generated from enum city.comm.input.v1.BaseStationType
 */
export declare enum BaseStationType {
  /**
   * @generated from enum value: BASE_STATION_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: BASE_STATION_TYPE_INDOOR = 1;
   */
  INDOOR = 1,

  /**
   * @generated from enum value: BASE_STATION_TYPE_OUTDOOR = 2;
   */
  OUTDOOR = 2,
}

/**
 * @generated from message city.comm.input.v1.Node
 */
export declare class Node extends Message<Node> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: city.comm.input.v1.NodeType type = 2;
   */
  type: NodeType;

  /**
   * 父节点
   *
   * @generated from field: int32 parent_id = 3;
   */
  parentId: number;

  /**
   * 子节点
   *
   * @generated from field: repeated int32 children_ids = 4;
   */
  childrenIds: number[];

  /**
   * 节点经纬度位置
   *
   * @generated from field: optional city.geo.v2.Position position = 5;
   */
  position?: Position;

  /**
   * 节点所在aoi
   *
   * @generated from field: optional int32 aoi_id = 6;
   */
  aoiId?: number;

  /**
   * 基站频段id
   *
   * @generated from field: optional int32 freq_range_id = 7;
   */
  freqRangeId?: number;

  /**
   * 室内外基站类型
   *
   * @generated from field: optional city.comm.input.v1.BaseStationType base_station_type = 8;
   */
  baseStationType?: BaseStationType;

  constructor(data?: PartialMessage<Node>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.input.v1.Node";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Node;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Node;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Node;

  static equals(a: Node | PlainMessage<Node> | undefined, b: Node | PlainMessage<Node> | undefined): boolean;
}

/**
 * 抢修站
 *
 * @generated from message city.comm.input.v1.RepairStation
 */
export declare class RepairStation extends Message<RepairStation> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: int32 aoi_id = 2;
   */
  aoiId: number;

  /**
   * @generated from field: city.geo.v2.Position position = 3;
   */
  position?: Position;

  constructor(data?: PartialMessage<RepairStation>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.input.v1.RepairStation";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RepairStation;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RepairStation;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RepairStation;

  static equals(a: RepairStation | PlainMessage<RepairStation> | undefined, b: RepairStation | PlainMessage<RepairStation> | undefined): boolean;
}

/**
 * 水泵
 *
 * @generated from message city.comm.input.v1.Pump
 */
export declare class Pump extends Message<Pump> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: city.geo.v2.Position position = 2;
   */
  position?: Position;

  constructor(data?: PartialMessage<Pump>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.input.v1.Pump";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pump;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pump;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pump;

  static equals(a: Pump | PlainMessage<Pump> | undefined, b: Pump | PlainMessage<Pump> | undefined): boolean;
}

/**
 * 终端通信需求
 *
 * @generated from message city.comm.input.v1.CommDemand
 */
export declare class CommDemand extends Message<CommDemand> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: repeated double demands = 2;
   */
  demands: number[];

  constructor(data?: PartialMessage<CommDemand>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.input.v1.CommDemand";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CommDemand;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CommDemand;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CommDemand;

  static equals(a: CommDemand | PlainMessage<CommDemand> | undefined, b: CommDemand | PlainMessage<CommDemand> | undefined): boolean;
}

/**
 * @generated from message city.comm.input.v1.Nodes
 */
export declare class Nodes extends Message<Nodes> {
  /**
   * @generated from field: repeated city.comm.input.v1.Node nodes = 1;
   */
  nodes: Node[];

  /**
   * @generated from field: repeated city.comm.input.v1.RepairStation repair_stations = 2;
   */
  repairStations: RepairStation[];

  /**
   * @generated from field: repeated city.comm.input.v1.Pump pumps = 3;
   */
  pumps: Pump[];

  constructor(data?: PartialMessage<Nodes>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.input.v1.Nodes";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Nodes;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Nodes;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Nodes;

  static equals(a: Nodes | PlainMessage<Nodes> | undefined, b: Nodes | PlainMessage<Nodes> | undefined): boolean;
}

/**
 * @generated from message city.comm.input.v1.CommDemands
 */
export declare class CommDemands extends Message<CommDemands> {
  /**
   * @generated from field: repeated city.comm.input.v1.CommDemand comm_demands = 1;
   */
  commDemands: CommDemand[];

  constructor(data?: PartialMessage<CommDemands>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.input.v1.CommDemands";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CommDemands;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CommDemands;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CommDemands;

  static equals(a: CommDemands | PlainMessage<CommDemands> | undefined, b: CommDemands | PlainMessage<CommDemands> | undefined): boolean;
}
