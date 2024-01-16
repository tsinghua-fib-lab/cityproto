// @generated by protoc-gen-es v1.6.0
// @generated from file city/event/v1/event_service.proto (package city.event.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Event } from "./event_pb.js";

/**
 * @generated from message city.event.v1.PublishRequest
 */
export declare class PublishRequest extends Message<PublishRequest> {
  /**
   * @generated from field: city.event.v1.Event event = 1;
   */
  event?: Event;

  constructor(data?: PartialMessage<PublishRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.event.v1.PublishRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PublishRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PublishRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PublishRequest;

  static equals(a: PublishRequest | PlainMessage<PublishRequest> | undefined, b: PublishRequest | PlainMessage<PublishRequest> | undefined): boolean;
}

/**
 * @generated from message city.event.v1.PublishResponse
 */
export declare class PublishResponse extends Message<PublishResponse> {
  constructor(data?: PartialMessage<PublishResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.event.v1.PublishResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PublishResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PublishResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PublishResponse;

  static equals(a: PublishResponse | PlainMessage<PublishResponse> | undefined, b: PublishResponse | PlainMessage<PublishResponse> | undefined): boolean;
}

/**
 * @generated from message city.event.v1.PullRequest
 */
export declare class PullRequest extends Message<PullRequest> {
  constructor(data?: PartialMessage<PullRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.event.v1.PullRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PullRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PullRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PullRequest;

  static equals(a: PullRequest | PlainMessage<PullRequest> | undefined, b: PullRequest | PlainMessage<PullRequest> | undefined): boolean;
}

/**
 * @generated from message city.event.v1.PullResponse
 */
export declare class PullResponse extends Message<PullResponse> {
  /**
   * @generated from field: repeated city.event.v1.Event events = 1;
   */
  events: Event[];

  constructor(data?: PartialMessage<PullResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.event.v1.PullResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PullResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PullResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PullResponse;

  static equals(a: PullResponse | PlainMessage<PullResponse> | undefined, b: PullResponse | PlainMessage<PullResponse> | undefined): boolean;
}

