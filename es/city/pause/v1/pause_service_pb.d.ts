// @generated by protoc-gen-es v1.10.0
// @generated from file city/pause/v1/pause_service.proto (package city.pause.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * 暂停程序请求
 * Pause program request
 *
 * @generated from message city.pause.v1.PauseRequest
 */
export declare class PauseRequest extends Message<PauseRequest> {
  /**
   * 组件名，用于暂停某一特定组件的运行，不设置时暂停整个程序
   * Component name, used to pause a specific component, pause the entire program when not set
   *
   * @generated from field: optional string name = 1;
   */
  name?: string;

  constructor(data?: PartialMessage<PauseRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.pause.v1.PauseRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PauseRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PauseRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PauseRequest;

  static equals(a: PauseRequest | PlainMessage<PauseRequest> | undefined, b: PauseRequest | PlainMessage<PauseRequest> | undefined): boolean;
}

/**
 * 暂停程序响应
 * Pause program response
 *
 * @generated from message city.pause.v1.PauseResponse
 */
export declare class PauseResponse extends Message<PauseResponse> {
  constructor(data?: PartialMessage<PauseResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.pause.v1.PauseResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PauseResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PauseResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PauseResponse;

  static equals(a: PauseResponse | PlainMessage<PauseResponse> | undefined, b: PauseResponse | PlainMessage<PauseResponse> | undefined): boolean;
}

/**
 * 恢复程序请求
 * Resume program request
 *
 * @generated from message city.pause.v1.ResumeRequest
 */
export declare class ResumeRequest extends Message<ResumeRequest> {
  /**
   * 组件名，用于恢复某一特定组件的运行，不设置时恢复整个程序
   * Component name, used to resume a specific component, resume the entire program when not set
   *
   * @generated from field: optional string name = 1;
   */
  name?: string;

  constructor(data?: PartialMessage<ResumeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.pause.v1.ResumeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResumeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResumeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResumeRequest;

  static equals(a: ResumeRequest | PlainMessage<ResumeRequest> | undefined, b: ResumeRequest | PlainMessage<ResumeRequest> | undefined): boolean;
}

/**
 * 恢复程序响应
 * Resume program response
 *
 * @generated from message city.pause.v1.ResumeResponse
 */
export declare class ResumeResponse extends Message<ResumeResponse> {
  constructor(data?: PartialMessage<ResumeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.pause.v1.ResumeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResumeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResumeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResumeResponse;

  static equals(a: ResumeResponse | PlainMessage<ResumeResponse> | undefined, b: ResumeResponse | PlainMessage<ResumeResponse> | undefined): boolean;
}

