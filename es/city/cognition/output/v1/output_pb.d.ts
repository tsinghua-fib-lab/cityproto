// @generated by protoc-gen-es v1.6.0
// @generated from file city/cognition/output/v1/output.proto (package city.cognition.output.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * 社交网络节点状态
 *
 * @generated from message city.cognition.output.v1.Node
 */
export declare class Node extends Message<Node> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: double status = 3;
   */
  status: number;

  constructor(data?: PartialMessage<Node>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.output.v1.Node";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Node;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Node;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Node;

  static equals(a: Node | PlainMessage<Node> | undefined, b: Node | PlainMessage<Node> | undefined): boolean;
}

/**
 * 一次认知影响的过程
 *
 * @generated from message city.cognition.output.v1.Influence
 */
export declare class Influence extends Message<Influence> {
  /**
   * @generated from field: int32 source_id = 1;
   */
  sourceId: number;

  /**
   * @generated from field: int32 target_id = 2;
   */
  targetId: number;

  /**
   * @generated from field: double strength = 3;
   */
  strength: number;

  constructor(data?: PartialMessage<Influence>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.output.v1.Influence";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Influence;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Influence;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Influence;

  static equals(a: Influence | PlainMessage<Influence> | undefined, b: Influence | PlainMessage<Influence> | undefined): boolean;
}

/**
 * 热力图数据
 *
 * @generated from message city.cognition.output.v1.Heatmap
 */
export declare class Heatmap extends Message<Heatmap> {
  /**
   * @generated from field: int32 num_rows = 1;
   */
  numRows: number;

  /**
   * @generated from field: int32 num_columns = 2;
   */
  numColumns: number;

  /**
   * @generated from field: repeated double strength = 3;
   */
  strength: number[];

  constructor(data?: PartialMessage<Heatmap>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.output.v1.Heatmap";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Heatmap;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Heatmap;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Heatmap;

  static equals(a: Heatmap | PlainMessage<Heatmap> | undefined, b: Heatmap | PlainMessage<Heatmap> | undefined): boolean;
}

/**
 * 统计数据
 *
 * @generated from message city.cognition.output.v1.Stat
 */
export declare class Stat extends Message<Stat> {
  /**
   * @generated from field: repeated int32 crowd_cnts = 1;
   */
  crowdCnts: number[];

  /**
   * @generated from field: repeated double crowd_ratios = 2;
   */
  crowdRatios: number[];

  /**
   * @generated from field: repeated int32 key_nodes = 3;
   */
  keyNodes: number[];

  constructor(data?: PartialMessage<Stat>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.output.v1.Stat";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Stat;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Stat;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Stat;

  static equals(a: Stat | PlainMessage<Stat> | undefined, b: Stat | PlainMessage<Stat> | undefined): boolean;
}

/**
 * 用户发布内容
 *
 * @generated from message city.cognition.output.v1.Content
 */
export declare class Content extends Message<Content> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: string text = 2;
   */
  text: string;

  constructor(data?: PartialMessage<Content>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.output.v1.Content";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Content;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Content;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Content;

  static equals(a: Content | PlainMessage<Content> | undefined, b: Content | PlainMessage<Content> | undefined): boolean;
}

/**
 * 社交网络节点静态属性
 *
 * @generated from message city.cognition.output.v1.NodeMeta
 */
export declare class NodeMeta extends Message<NodeMeta> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: double lng = 2;
   */
  lng: number;

  /**
   * @generated from field: double lat = 3;
   */
  lat: number;

  /**
   * @generated from field: int32 level = 4;
   */
  level: number;

  constructor(data?: PartialMessage<NodeMeta>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.output.v1.NodeMeta";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NodeMeta;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NodeMeta;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NodeMeta;

  static equals(a: NodeMeta | PlainMessage<NodeMeta> | undefined, b: NodeMeta | PlainMessage<NodeMeta> | undefined): boolean;
}

/**
 * @generated from message city.cognition.output.v1.NodesMeta
 */
export declare class NodesMeta extends Message<NodesMeta> {
  /**
   * @generated from field: repeated city.cognition.output.v1.NodeMeta nodes = 1;
   */
  nodes: NodeMeta[];

  constructor(data?: PartialMessage<NodesMeta>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.output.v1.NodesMeta";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NodesMeta;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NodesMeta;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NodesMeta;

  static equals(a: NodesMeta | PlainMessage<NodesMeta> | undefined, b: NodesMeta | PlainMessage<NodesMeta> | undefined): boolean;
}

/**
 * 群体信息
 *
 * @generated from message city.cognition.output.v1.Group
 */
export declare class Group extends Message<Group> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: int32 size = 2;
   */
  size: number;

  /**
   * @generated from field: repeated int32 changes = 3;
   */
  changes: number[];

  constructor(data?: PartialMessage<Group>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.output.v1.Group";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Group;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Group;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Group;

  static equals(a: Group | PlainMessage<Group> | undefined, b: Group | PlainMessage<Group> | undefined): boolean;
}

