// @generated by protoc-gen-es v1.10.0
// @generated from file city/economy/v2/economy.proto (package city.economy.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum city.economy.v2.OrgType
 */
export declare enum OrgType {
  /**
   * 未指定
   * unspecified
   *
   * @generated from enum value: ORG_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 国家统计局
   * NBS
   *
   * @generated from enum value: ORG_TYPE_NBS = 1;
   */
  NBS = 1,

  /**
   * 公司
   * firm
   *
   * @generated from enum value: ORG_TYPE_FIRM = 2;
   */
  FIRM = 2,

  /**
   * 银行
   * bank
   *
   * @generated from enum value: ORG_TYPE_BANK = 3;
   */
  BANK = 3,

  /**
   * 政府
   * government
   *
   * @generated from enum value: ORG_TYPE_GOVERNMENT = 4;
   */
  GOVERNMENT = 4,
}

/**
 * 组织
 * Organization
 *
 * @generated from message city.economy.v2.Org
 */
export declare class Org extends Message<Org> {
  /**
   * 组织ID
   * organization id
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * 组织类别
   * organization type
   *
   * @generated from field: city.economy.v2.OrgType type = 2;
   */
  type: OrgType;

  /**
   * NBS
   *
   *
   * @generated from field: repeated float nominal_gdp = 3;
   */
  nominalGdp: number[];

  /**
   * @generated from field: repeated float real_gdp = 4;
   */
  realGdp: number[];

  /**
   * @generated from field: repeated float unemployment = 5;
   */
  unemployment: number[];

  /**
   * @generated from field: repeated float wages = 6;
   */
  wages: number[];

  /**
   * @generated from field: repeated float prices = 7;
   */
  prices: number[];

  /**
   * Firm
   *
   *
   * @generated from field: optional int32 inventory = 8;
   */
  inventory?: number;

  /**
   * @generated from field: optional float price = 9;
   */
  price?: number;

  /**
   * Firm & Bank & Government
   *
   *
   * @generated from field: optional float currency = 10;
   */
  currency?: number;

  /**
   * Bank
   *
   *
   * @generated from field: optional float interest_rate = 11;
   */
  interestRate?: number;

  /**
   * Government
   *
   *
   * @generated from field: repeated float bracket_cutoffs = 12;
   */
  bracketCutoffs: number[];

  /**
   * @generated from field: repeated float bracket_rates = 13;
   */
  bracketRates: number[];

  constructor(data?: PartialMessage<Org>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v2.Org";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Org;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Org;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Org;

  static equals(a: Org | PlainMessage<Org> | undefined, b: Org | PlainMessage<Org> | undefined): boolean;
}

/**
 * @generated from message city.economy.v2.Agent
 */
export declare class Agent extends Message<Agent> {
  /**
   * person ID
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * currency
   *
   * @generated from field: optional float currency = 2;
   */
  currency?: number;

  constructor(data?: PartialMessage<Agent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v2.Agent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Agent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Agent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Agent;

  static equals(a: Agent | PlainMessage<Agent> | undefined, b: Agent | PlainMessage<Agent> | undefined): boolean;
}

/**
 * @generated from message city.economy.v2.EconomyEntities
 */
export declare class EconomyEntities extends Message<EconomyEntities> {
  /**
   * @generated from field: repeated city.economy.v2.Org orgs = 1;
   */
  orgs: Org[];

  /**
   * @generated from field: repeated city.economy.v2.Agent agents = 2;
   */
  agents: Agent[];

  constructor(data?: PartialMessage<EconomyEntities>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v2.EconomyEntities";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EconomyEntities;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EconomyEntities;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EconomyEntities;

  static equals(a: EconomyEntities | PlainMessage<EconomyEntities> | undefined, b: EconomyEntities | PlainMessage<EconomyEntities> | undefined): boolean;
}
