// @generated by protoc-gen-es v1.10.0
// @generated from file city/sync/v2/sync_service.proto (package city.sync.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * 注册程序URL请求
 *
 * @generated from message city.sync.v2.SetURLRequest
 */
export declare class SetURLRequest extends Message<SetURLRequest> {
  /**
   * 组件名，需要在同步器启动参数列表中
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * 程序URL
   *
   * @generated from field: string url = 2;
   */
  url: string;

  constructor(data?: PartialMessage<SetURLRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.sync.v2.SetURLRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetURLRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetURLRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetURLRequest;

  static equals(a: SetURLRequest | PlainMessage<SetURLRequest> | undefined, b: SetURLRequest | PlainMessage<SetURLRequest> | undefined): boolean;
}

/**
 * 注册程序URL响应
 *
 * @generated from message city.sync.v2.SetURLResponse
 */
export declare class SetURLResponse extends Message<SetURLResponse> {
  constructor(data?: PartialMessage<SetURLResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.sync.v2.SetURLResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetURLResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetURLResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetURLResponse;

  static equals(a: SetURLResponse | PlainMessage<SetURLResponse> | undefined, b: SetURLResponse | PlainMessage<SetURLResponse> | undefined): boolean;
}

/**
 * 获取程序URL请求
 *
 * @generated from message city.sync.v2.GetURLRequest
 */
export declare class GetURLRequest extends Message<GetURLRequest> {
  /**
   * 组件名，需要在同步器启动参数列表中
   *
   * @generated from field: string name = 1;
   */
  name: string;

  constructor(data?: PartialMessage<GetURLRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.sync.v2.GetURLRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetURLRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetURLRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetURLRequest;

  static equals(a: GetURLRequest | PlainMessage<GetURLRequest> | undefined, b: GetURLRequest | PlainMessage<GetURLRequest> | undefined): boolean;
}

/**
 * 获取程序URL响应
 *
 * @generated from message city.sync.v2.GetURLResponse
 */
export declare class GetURLResponse extends Message<GetURLResponse> {
  /**
   * 程序URL
   *
   * @generated from field: string url = 1;
   */
  url: string;

  constructor(data?: PartialMessage<GetURLResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.sync.v2.GetURLResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetURLResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetURLResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetURLResponse;

  static equals(a: GetURLResponse | PlainMessage<GetURLResponse> | undefined, b: GetURLResponse | PlainMessage<GetURLResponse> | undefined): boolean;
}

/**
 * 进入同步状态请求
 * Enter step sync request
 *
 * @generated from message city.sync.v2.EnterStepSyncRequest
 */
export declare class EnterStepSyncRequest extends Message<EnterStepSyncRequest> {
  /**
   * 组件名，需要在同步器启动参数列表中
   *
   * @generated from field: string name = 1;
   */
  name: string;

  constructor(data?: PartialMessage<EnterStepSyncRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.sync.v2.EnterStepSyncRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EnterStepSyncRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EnterStepSyncRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EnterStepSyncRequest;

  static equals(a: EnterStepSyncRequest | PlainMessage<EnterStepSyncRequest> | undefined, b: EnterStepSyncRequest | PlainMessage<EnterStepSyncRequest> | undefined): boolean;
}

/**
 * 进入同步状态响应
 * Enter step sync response
 *
 * @generated from message city.sync.v2.EnterStepSyncResponse
 */
export declare class EnterStepSyncResponse extends Message<EnterStepSyncResponse> {
  constructor(data?: PartialMessage<EnterStepSyncResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.sync.v2.EnterStepSyncResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EnterStepSyncResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EnterStepSyncResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EnterStepSyncResponse;

  static equals(a: EnterStepSyncResponse | PlainMessage<EnterStepSyncResponse> | undefined, b: EnterStepSyncResponse | PlainMessage<EnterStepSyncResponse> | undefined): boolean;
}

/**
 * 退出同步状态请求
 * Exit step sync request
 *
 * @generated from message city.sync.v2.ExitStepSyncRequest
 */
export declare class ExitStepSyncRequest extends Message<ExitStepSyncRequest> {
  /**
   * 组件名，需要在同步器启动参数列表中
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * 是否退出服务
   *
   * @generated from field: bool close = 2;
   */
  close: boolean;

  constructor(data?: PartialMessage<ExitStepSyncRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.sync.v2.ExitStepSyncRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExitStepSyncRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExitStepSyncRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExitStepSyncRequest;

  static equals(a: ExitStepSyncRequest | PlainMessage<ExitStepSyncRequest> | undefined, b: ExitStepSyncRequest | PlainMessage<ExitStepSyncRequest> | undefined): boolean;
}

/**
 * 退出同步状态响应
 * Exit step sync response
 *
 * @generated from message city.sync.v2.ExitStepSyncResponse
 */
export declare class ExitStepSyncResponse extends Message<ExitStepSyncResponse> {
  /**
   * 服务是否关闭
   *
   * @generated from field: bool close = 1;
   */
  close: boolean;

  constructor(data?: PartialMessage<ExitStepSyncResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.sync.v2.ExitStepSyncResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExitStepSyncResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExitStepSyncResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExitStepSyncResponse;

  static equals(a: ExitStepSyncResponse | PlainMessage<ExitStepSyncResponse> | undefined, b: ExitStepSyncResponse | PlainMessage<ExitStepSyncResponse> | undefined): boolean;
}

