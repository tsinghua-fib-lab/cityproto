// @generated by protoc-gen-es v1.6.0
// @generated from file city/cognition/input/v1/config.proto (package city.cognition.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { MongoPath, OutputTarget } from "../../../config/v1/config_pb.js";

/**
 * @generated from message city.cognition.input.v1.Mongo
 */
export declare class Mongo extends Message<Mongo> {
  /**
   * @generated from field: string uri = 1;
   */
  uri: string;

  /**
   * @generated from field: city.config.v1.MongoPath map = 2;
   */
  map?: MongoPath;

  /**
   * @generated from field: city.config.v1.MongoPath agents = 3;
   */
  agents?: MongoPath;

  /**
   * @generated from field: city.config.v1.MongoPath edges = 4;
   */
  edges?: MongoPath;

  constructor(data?: PartialMessage<Mongo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.input.v1.Mongo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Mongo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Mongo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Mongo;

  static equals(a: Mongo | PlainMessage<Mongo> | undefined, b: Mongo | PlainMessage<Mongo> | undefined): boolean;
}

/**
 * @generated from message city.cognition.input.v1.ControlStep
 */
export declare class ControlStep extends Message<ControlStep> {
  /**
   * @generated from field: int32 start = 1;
   */
  start: number;

  /**
   * @generated from field: int32 total = 2;
   */
  total: number;

  /**
   * @generated from field: float interval = 3;
   */
  interval: number;

  constructor(data?: PartialMessage<ControlStep>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.input.v1.ControlStep";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ControlStep;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ControlStep;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ControlStep;

  static equals(a: ControlStep | PlainMessage<ControlStep> | undefined, b: ControlStep | PlainMessage<ControlStep> | undefined): boolean;
}

/**
 * @generated from message city.cognition.input.v1.Control
 */
export declare class Control extends Message<Control> {
  /**
   * @generated from field: city.cognition.input.v1.ControlStep step = 1;
   */
  step?: ControlStep;

  constructor(data?: PartialMessage<Control>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.input.v1.Control";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Control;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Control;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Control;

  static equals(a: Control | PlainMessage<Control> | undefined, b: Control | PlainMessage<Control> | undefined): boolean;
}

/**
 * 是否输出各类数据
 *
 * @generated from message city.cognition.input.v1.OutputSwitch
 */
export declare class OutputSwitch extends Message<OutputSwitch> {
  /**
   * @generated from field: bool common = 1;
   */
  common: boolean;

  constructor(data?: PartialMessage<OutputSwitch>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.input.v1.OutputSwitch";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OutputSwitch;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OutputSwitch;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OutputSwitch;

  static equals(a: OutputSwitch | PlainMessage<OutputSwitch> | undefined, b: OutputSwitch | PlainMessage<OutputSwitch> | undefined): boolean;
}

/**
 * @generated from message city.cognition.input.v1.Output
 */
export declare class Output extends Message<Output> {
  /**
   * @generated from field: city.config.v1.OutputTarget target = 1;
   */
  target?: OutputTarget;

  /**
   * @generated from field: city.cognition.input.v1.OutputSwitch switch = 2;
   */
  switch?: OutputSwitch;

  constructor(data?: PartialMessage<Output>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.input.v1.Output";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Output;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Output;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Output;

  static equals(a: Output | PlainMessage<Output> | undefined, b: Output | PlainMessage<Output> | undefined): boolean;
}

/**
 * @generated from message city.cognition.input.v1.Config
 */
export declare class Config extends Message<Config> {
  /**
   * @generated from field: city.cognition.input.v1.Mongo mongo = 1;
   */
  mongo?: Mongo;

  /**
   * @generated from field: city.cognition.input.v1.Control control = 2;
   */
  control?: Control;

  /**
   * @generated from field: city.cognition.input.v1.Output output = 3;
   */
  output?: Output;

  constructor(data?: PartialMessage<Config>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.cognition.input.v1.Config";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Config;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Config;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Config;

  static equals(a: Config | PlainMessage<Config> | undefined, b: Config | PlainMessage<Config> | undefined): boolean;
}

