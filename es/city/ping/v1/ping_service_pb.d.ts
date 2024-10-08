// @generated by protoc-gen-es v1.10.0
// @generated from file city/ping/v1/ping_service.proto (package city.ping.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * 连接测试请求
 *
 * @generated from message city.ping.v1.PingRequest
 */
export declare class PingRequest extends Message<PingRequest> {
  constructor(data?: PartialMessage<PingRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.ping.v1.PingRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PingRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PingRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PingRequest;

  static equals(a: PingRequest | PlainMessage<PingRequest> | undefined, b: PingRequest | PlainMessage<PingRequest> | undefined): boolean;
}

/**
 * 连接测试响应
 *
 * @generated from message city.ping.v1.PingResponse
 */
export declare class PingResponse extends Message<PingResponse> {
  constructor(data?: PartialMessage<PingResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.ping.v1.PingResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PingResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PingResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PingResponse;

  static equals(a: PingResponse | PlainMessage<PingResponse> | undefined, b: PingResponse | PlainMessage<PingResponse> | undefined): boolean;
}

