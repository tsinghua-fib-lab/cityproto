// @generated by protoc-gen-es v1.6.0
// @generated from file city/comm/interaction/gateway/v1/gateway_service.proto (package city.comm.interaction.gateway.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Station } from "./gateway_pb.js";
import type { Events } from "../../../../event/v1/event_pb.js";

/**
 * 断电或恢复状态，True为修复，False为断电
 *
 * @generated from message city.comm.interaction.gateway.v1.SetGatewayPowerStatusRequest
 */
export declare class SetGatewayPowerStatusRequest extends Message<SetGatewayPowerStatusRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: bool status = 2;
   */
  status: boolean;

  constructor(data?: PartialMessage<SetGatewayPowerStatusRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.SetGatewayPowerStatusRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetGatewayPowerStatusRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetGatewayPowerStatusRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetGatewayPowerStatusRequest;

  static equals(a: SetGatewayPowerStatusRequest | PlainMessage<SetGatewayPowerStatusRequest> | undefined, b: SetGatewayPowerStatusRequest | PlainMessage<SetGatewayPowerStatusRequest> | undefined): boolean;
}

/**
 * @generated from message city.comm.interaction.gateway.v1.SetGatewayPowerStatusResponse
 */
export declare class SetGatewayPowerStatusResponse extends Message<SetGatewayPowerStatusResponse> {
  constructor(data?: PartialMessage<SetGatewayPowerStatusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.SetGatewayPowerStatusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetGatewayPowerStatusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetGatewayPowerStatusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetGatewayPowerStatusResponse;

  static equals(a: SetGatewayPowerStatusResponse | PlainMessage<SetGatewayPowerStatusResponse> | undefined, b: SetGatewayPowerStatusResponse | PlainMessage<SetGatewayPowerStatusResponse> | undefined): boolean;
}

/**
 * 摧毁或恢复状态，True为修复，False为摧毁
 *
 * @generated from message city.comm.interaction.gateway.v1.SetGatewayRuinStatusRequest
 */
export declare class SetGatewayRuinStatusRequest extends Message<SetGatewayRuinStatusRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: bool status = 2;
   */
  status: boolean;

  constructor(data?: PartialMessage<SetGatewayRuinStatusRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.SetGatewayRuinStatusRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetGatewayRuinStatusRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetGatewayRuinStatusRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetGatewayRuinStatusRequest;

  static equals(a: SetGatewayRuinStatusRequest | PlainMessage<SetGatewayRuinStatusRequest> | undefined, b: SetGatewayRuinStatusRequest | PlainMessage<SetGatewayRuinStatusRequest> | undefined): boolean;
}

/**
 * @generated from message city.comm.interaction.gateway.v1.SetGatewayRuinStatusResponse
 */
export declare class SetGatewayRuinStatusResponse extends Message<SetGatewayRuinStatusResponse> {
  constructor(data?: PartialMessage<SetGatewayRuinStatusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.SetGatewayRuinStatusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetGatewayRuinStatusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetGatewayRuinStatusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetGatewayRuinStatusResponse;

  static equals(a: SetGatewayRuinStatusResponse | PlainMessage<SetGatewayRuinStatusResponse> | undefined, b: SetGatewayRuinStatusResponse | PlainMessage<SetGatewayRuinStatusResponse> | undefined): boolean;
}

/**
 * @generated from message city.comm.interaction.gateway.v1.GetAllStatusRequest
 */
export declare class GetAllStatusRequest extends Message<GetAllStatusRequest> {
  constructor(data?: PartialMessage<GetAllStatusRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.GetAllStatusRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAllStatusRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAllStatusRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAllStatusRequest;

  static equals(a: GetAllStatusRequest | PlainMessage<GetAllStatusRequest> | undefined, b: GetAllStatusRequest | PlainMessage<GetAllStatusRequest> | undefined): boolean;
}

/**
 * @generated from message city.comm.interaction.gateway.v1.GetAllStatusResponse
 */
export declare class GetAllStatusResponse extends Message<GetAllStatusResponse> {
  /**
   * @generated from field: repeated city.comm.interaction.gateway.v1.Station stations = 1;
   */
  stations: Station[];

  constructor(data?: PartialMessage<GetAllStatusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.GetAllStatusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAllStatusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAllStatusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAllStatusResponse;

  static equals(a: GetAllStatusResponse | PlainMessage<GetAllStatusResponse> | undefined, b: GetAllStatusResponse | PlainMessage<GetAllStatusResponse> | undefined): boolean;
}

/**
 * @generated from message city.comm.interaction.gateway.v1.GetRuinInfoRequest
 */
export declare class GetRuinInfoRequest extends Message<GetRuinInfoRequest> {
  constructor(data?: PartialMessage<GetRuinInfoRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.GetRuinInfoRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRuinInfoRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRuinInfoRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRuinInfoRequest;

  static equals(a: GetRuinInfoRequest | PlainMessage<GetRuinInfoRequest> | undefined, b: GetRuinInfoRequest | PlainMessage<GetRuinInfoRequest> | undefined): boolean;
}

/**
 * @generated from message city.comm.interaction.gateway.v1.RuinInfo
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
  static readonly typeName = "city.comm.interaction.gateway.v1.RuinInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RuinInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RuinInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RuinInfo;

  static equals(a: RuinInfo | PlainMessage<RuinInfo> | undefined, b: RuinInfo | PlainMessage<RuinInfo> | undefined): boolean;
}

/**
 * @generated from message city.comm.interaction.gateway.v1.GetRuinInfoResponse
 */
export declare class GetRuinInfoResponse extends Message<GetRuinInfoResponse> {
  /**
   * 三级级损伤信息
   *
   * @generated from field: city.comm.interaction.gateway.v1.RuinInfo one = 1;
   */
  one?: RuinInfo;

  /**
   * @generated from field: city.comm.interaction.gateway.v1.RuinInfo two = 2;
   */
  two?: RuinInfo;

  /**
   * @generated from field: city.comm.interaction.gateway.v1.RuinInfo three = 3;
   */
  three?: RuinInfo;

  constructor(data?: PartialMessage<GetRuinInfoResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.GetRuinInfoResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRuinInfoResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRuinInfoResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRuinInfoResponse;

  static equals(a: GetRuinInfoResponse | PlainMessage<GetRuinInfoResponse> | undefined, b: GetRuinInfoResponse | PlainMessage<GetRuinInfoResponse> | undefined): boolean;
}

/**
 * @generated from message city.comm.interaction.gateway.v1.GetEventsRequest
 */
export declare class GetEventsRequest extends Message<GetEventsRequest> {
  constructor(data?: PartialMessage<GetEventsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.GetEventsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetEventsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetEventsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetEventsRequest;

  static equals(a: GetEventsRequest | PlainMessage<GetEventsRequest> | undefined, b: GetEventsRequest | PlainMessage<GetEventsRequest> | undefined): boolean;
}

/**
 * @generated from message city.comm.interaction.gateway.v1.GetEventsResponse
 */
export declare class GetEventsResponse extends Message<GetEventsResponse> {
  /**
   * @generated from field: city.event.v1.Events events = 1;
   */
  events?: Events;

  constructor(data?: PartialMessage<GetEventsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.comm.interaction.gateway.v1.GetEventsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetEventsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetEventsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetEventsResponse;

  static equals(a: GetEventsResponse | PlainMessage<GetEventsResponse> | undefined, b: GetEventsResponse | PlainMessage<GetEventsResponse> | undefined): boolean;
}
