// @generated by protoc-gen-es v1.6.0
// @generated from file city/economy/v1/org_service.proto (package city.economy.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Employee, Goods, Job, Org } from "./economy_pb.js";

/**
 * 批量查询组织的经济情况请求
 *
 * @generated from message city.economy.v1.GetOrgRequest
 */
export declare class GetOrgRequest extends Message<GetOrgRequest> {
  /**
   * 待查询的组织的ID列表（为空时查询所有组织）
   *
   * @generated from field: repeated int32 org_ids = 1;
   */
  orgIds: number[];

  constructor(data?: PartialMessage<GetOrgRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.GetOrgRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOrgRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOrgRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOrgRequest;

  static equals(a: GetOrgRequest | PlainMessage<GetOrgRequest> | undefined, b: GetOrgRequest | PlainMessage<GetOrgRequest> | undefined): boolean;
}

/**
 * 批量查询组织的经济情况响应
 *
 * @generated from message city.economy.v1.GetOrgResponse
 */
export declare class GetOrgResponse extends Message<GetOrgResponse> {
  /**
   * 组织的经济情况
   *
   * @generated from field: repeated city.economy.v1.Org orgs = 1;
   */
  orgs: Org[];

  constructor(data?: PartialMessage<GetOrgResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.GetOrgResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOrgResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOrgResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOrgResponse;

  static equals(a: GetOrgResponse | PlainMessage<GetOrgResponse> | undefined, b: GetOrgResponse | PlainMessage<GetOrgResponse> | undefined): boolean;
}

/**
 * 批量修改组织的资金请求
 *
 * @generated from message city.economy.v1.UpdateOrgMoneyRequest
 */
export declare class UpdateOrgMoneyRequest extends Message<UpdateOrgMoneyRequest> {
  /**
   * 待修改的组织资金变动
   *
   * @generated from field: repeated city.economy.v1.UpdateOrgMoneyRequestItem items = 1;
   */
  items: UpdateOrgMoneyRequestItem[];

  constructor(data?: PartialMessage<UpdateOrgMoneyRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgMoneyRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgMoneyRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgMoneyRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgMoneyRequest;

  static equals(a: UpdateOrgMoneyRequest | PlainMessage<UpdateOrgMoneyRequest> | undefined, b: UpdateOrgMoneyRequest | PlainMessage<UpdateOrgMoneyRequest> | undefined): boolean;
}

/**
 * 组织资金变动
 *
 * @generated from message city.economy.v1.UpdateOrgMoneyRequestItem
 */
export declare class UpdateOrgMoneyRequestItem extends Message<UpdateOrgMoneyRequestItem> {
  /**
   * 待修改的组织
   *
   * @generated from field: int32 org_id = 1;
   */
  orgId: number;

  /**
   * 正数表示增加，负数表示减少
   *
   * @generated from field: double money = 2;
   */
  money: number;

  constructor(data?: PartialMessage<UpdateOrgMoneyRequestItem>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgMoneyRequestItem";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgMoneyRequestItem;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgMoneyRequestItem;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgMoneyRequestItem;

  static equals(a: UpdateOrgMoneyRequestItem | PlainMessage<UpdateOrgMoneyRequestItem> | undefined, b: UpdateOrgMoneyRequestItem | PlainMessage<UpdateOrgMoneyRequestItem> | undefined): boolean;
}

/**
 * 批量修改组织的资金响应
 *
 * @generated from message city.economy.v1.UpdateOrgMoneyResponse
 */
export declare class UpdateOrgMoneyResponse extends Message<UpdateOrgMoneyResponse> {
  /**
   * 修改后的组织的经济情况
   *
   * @generated from field: repeated city.economy.v1.Org orgs = 1;
   */
  orgs: Org[];

  constructor(data?: PartialMessage<UpdateOrgMoneyResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgMoneyResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgMoneyResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgMoneyResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgMoneyResponse;

  static equals(a: UpdateOrgMoneyResponse | PlainMessage<UpdateOrgMoneyResponse> | undefined, b: UpdateOrgMoneyResponse | PlainMessage<UpdateOrgMoneyResponse> | undefined): boolean;
}

/**
 * 批量修改组织的货物请求
 *
 * @generated from message city.economy.v1.UpdateOrgGoodsRequest
 */
export declare class UpdateOrgGoodsRequest extends Message<UpdateOrgGoodsRequest> {
  /**
   * 待修改的组织货物变动
   *
   * @generated from field: repeated city.economy.v1.UpdateOrgGoodsRequestItem items = 1;
   */
  items: UpdateOrgGoodsRequestItem[];

  constructor(data?: PartialMessage<UpdateOrgGoodsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgGoodsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgGoodsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgGoodsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgGoodsRequest;

  static equals(a: UpdateOrgGoodsRequest | PlainMessage<UpdateOrgGoodsRequest> | undefined, b: UpdateOrgGoodsRequest | PlainMessage<UpdateOrgGoodsRequest> | undefined): boolean;
}

/**
 * 组织货物变动
 *
 * @generated from message city.economy.v1.UpdateOrgGoodsRequestItem
 */
export declare class UpdateOrgGoodsRequestItem extends Message<UpdateOrgGoodsRequestItem> {
  /**
   * 待修改的组织
   *
   * @generated from field: int32 org_id = 1;
   */
  orgId: number;

  /**
   * 货物变动
   * 按照(type, name)相等来判断是否为同一种货物
   * 货物数量为增量，正数表示增加，负数表示减少
   * price如果未设定则沿用原来的价格，否则使用新的价格
   *
   * @generated from field: repeated city.economy.v1.Goods goods = 2;
   */
  goods: Goods[];

  constructor(data?: PartialMessage<UpdateOrgGoodsRequestItem>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgGoodsRequestItem";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgGoodsRequestItem;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgGoodsRequestItem;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgGoodsRequestItem;

  static equals(a: UpdateOrgGoodsRequestItem | PlainMessage<UpdateOrgGoodsRequestItem> | undefined, b: UpdateOrgGoodsRequestItem | PlainMessage<UpdateOrgGoodsRequestItem> | undefined): boolean;
}

/**
 * 批量修改组织的货物响应
 *
 * @generated from message city.economy.v1.UpdateOrgGoodsResponse
 */
export declare class UpdateOrgGoodsResponse extends Message<UpdateOrgGoodsResponse> {
  /**
   * 修改后的组织的经济情况
   *
   * @generated from field: repeated city.economy.v1.Org orgs = 1;
   */
  orgs: Org[];

  constructor(data?: PartialMessage<UpdateOrgGoodsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgGoodsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgGoodsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgGoodsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgGoodsResponse;

  static equals(a: UpdateOrgGoodsResponse | PlainMessage<UpdateOrgGoodsResponse> | undefined, b: UpdateOrgGoodsResponse | PlainMessage<UpdateOrgGoodsResponse> | undefined): boolean;
}

/**
 * 批量修改组织的员工请求
 *
 * @generated from message city.economy.v1.UpdateOrgEmployeeRequest
 */
export declare class UpdateOrgEmployeeRequest extends Message<UpdateOrgEmployeeRequest> {
  /**
   * 待修改的组织员工变动
   *
   * @generated from field: repeated city.economy.v1.UpdateOrgEmployeeRequestItem items = 1;
   */
  items: UpdateOrgEmployeeRequestItem[];

  constructor(data?: PartialMessage<UpdateOrgEmployeeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgEmployeeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgEmployeeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgEmployeeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgEmployeeRequest;

  static equals(a: UpdateOrgEmployeeRequest | PlainMessage<UpdateOrgEmployeeRequest> | undefined, b: UpdateOrgEmployeeRequest | PlainMessage<UpdateOrgEmployeeRequest> | undefined): boolean;
}

/**
 * 组织员工变动
 *
 * @generated from message city.economy.v1.UpdateOrgEmployeeRequestItem
 */
export declare class UpdateOrgEmployeeRequestItem extends Message<UpdateOrgEmployeeRequestItem> {
  /**
   * 待修改的组织
   *
   * @generated from field: int32 org_id = 1;
   */
  orgId: number;

  /**
   * 新增的员工
   *
   * @generated from field: repeated city.economy.v1.Employee adds = 2;
   */
  adds: Employee[];

  /**
   * 删除的员工
   *
   * @generated from field: repeated int32 dels = 3;
   */
  dels: number[];

  /**
   * 修改薪水的员工
   *
   * @generated from field: repeated city.economy.v1.Employee updates = 4;
   */
  updates: Employee[];

  constructor(data?: PartialMessage<UpdateOrgEmployeeRequestItem>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgEmployeeRequestItem";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgEmployeeRequestItem;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgEmployeeRequestItem;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgEmployeeRequestItem;

  static equals(a: UpdateOrgEmployeeRequestItem | PlainMessage<UpdateOrgEmployeeRequestItem> | undefined, b: UpdateOrgEmployeeRequestItem | PlainMessage<UpdateOrgEmployeeRequestItem> | undefined): boolean;
}

/**
 * 批量修改组织的员工响应
 *
 * @generated from message city.economy.v1.UpdateOrgEmployeeResponse
 */
export declare class UpdateOrgEmployeeResponse extends Message<UpdateOrgEmployeeResponse> {
  /**
   * 修改后的组织的经济情况
   *
   * @generated from field: repeated city.economy.v1.Org orgs = 1;
   */
  orgs: Org[];

  constructor(data?: PartialMessage<UpdateOrgEmployeeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgEmployeeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgEmployeeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgEmployeeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgEmployeeResponse;

  static equals(a: UpdateOrgEmployeeResponse | PlainMessage<UpdateOrgEmployeeResponse> | undefined, b: UpdateOrgEmployeeResponse | PlainMessage<UpdateOrgEmployeeResponse> | undefined): boolean;
}

/**
 * 批量修改组织的岗位请求
 *
 * @generated from message city.economy.v1.UpdateOrgJobRequest
 */
export declare class UpdateOrgJobRequest extends Message<UpdateOrgJobRequest> {
  /**
   * 待修改的组织岗位变动
   *
   * @generated from field: repeated city.economy.v1.UpdateOrgJobRequestItem items = 1;
   */
  items: UpdateOrgJobRequestItem[];

  constructor(data?: PartialMessage<UpdateOrgJobRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgJobRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgJobRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgJobRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgJobRequest;

  static equals(a: UpdateOrgJobRequest | PlainMessage<UpdateOrgJobRequest> | undefined, b: UpdateOrgJobRequest | PlainMessage<UpdateOrgJobRequest> | undefined): boolean;
}

/**
 * 组织岗位变动
 *
 * @generated from message city.economy.v1.UpdateOrgJobRequestItem
 */
export declare class UpdateOrgJobRequestItem extends Message<UpdateOrgJobRequestItem> {
  /**
   * 待修改的组织
   *
   * @generated from field: int32 org_id = 1;
   */
  orgId: number;

  /**
   * 岗位变动
   * 按照name相等来判断是否为同一种岗位
   * 岗位数量为增量，正数表示增加，负数表示减少
   * salary如果未设定则沿用原来的薪水，否则使用新的薪水
   *
   * @generated from field: repeated city.economy.v1.Job jobs = 2;
   */
  jobs: Job[];

  constructor(data?: PartialMessage<UpdateOrgJobRequestItem>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgJobRequestItem";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgJobRequestItem;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgJobRequestItem;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgJobRequestItem;

  static equals(a: UpdateOrgJobRequestItem | PlainMessage<UpdateOrgJobRequestItem> | undefined, b: UpdateOrgJobRequestItem | PlainMessage<UpdateOrgJobRequestItem> | undefined): boolean;
}

/**
 * 批量修改组织的岗位响应
 *
 * @generated from message city.economy.v1.UpdateOrgJobResponse
 */
export declare class UpdateOrgJobResponse extends Message<UpdateOrgJobResponse> {
  /**
   * 修改后的组织的经济情况
   *
   * @generated from field: repeated city.economy.v1.Org orgs = 1;
   */
  orgs: Org[];

  constructor(data?: PartialMessage<UpdateOrgJobResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.economy.v1.UpdateOrgJobResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOrgJobResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOrgJobResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOrgJobResponse;

  static equals(a: UpdateOrgJobResponse | PlainMessage<UpdateOrgJobResponse> | undefined, b: UpdateOrgJobResponse | PlainMessage<UpdateOrgJobResponse> | undefined): boolean;
}

