// @generated by protoc-gen-es v1.6.0
// @generated from file city/elec/input/v1/input_service.proto (package city.elec.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Control } from "./config_pb.js";
import type { Facilities } from "./input_pb.js";
import type { Map } from "../../../map/v2/map_pb.js";

/**
 * @generated from message city.elec.input.v1.InitRequest
 */
export declare class InitRequest extends Message<InitRequest> {
  constructor(data?: PartialMessage<InitRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.elec.input.v1.InitRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InitRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InitRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InitRequest;

  static equals(a: InitRequest | PlainMessage<InitRequest> | undefined, b: InitRequest | PlainMessage<InitRequest> | undefined): boolean;
}

/**
 * @generated from message city.elec.input.v1.InitResponse
 */
export declare class InitResponse extends Message<InitResponse> {
  /**
   * 模拟器gRPC监听地址
   *
   * @generated from field: string address = 2;
   */
  address: string;

  /**
   * @generated from field: city.elec.input.v1.Control control = 3;
   */
  control?: Control;

  /**
   * @generated from field: city.elec.input.v1.Facilities facilities = 1;
   */
  facilities?: Facilities;

  /**
   * @generated from field: city.map.v2.Map map = 4;
   */
  map?: Map;

  constructor(data?: PartialMessage<InitResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.elec.input.v1.InitResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InitResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InitResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InitResponse;

  static equals(a: InitResponse | PlainMessage<InitResponse> | undefined, b: InitResponse | PlainMessage<InitResponse> | undefined): boolean;
}

