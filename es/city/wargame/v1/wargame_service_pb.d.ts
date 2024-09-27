// @generated by protoc-gen-es v1.10.0
// @generated from file city/wargame/v1/wargame_service.proto (package city.wargame.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Camp, Casualty, HitHistory, Point, RecoAlgoType, RecoPoint } from "./wargame_pb.js";

/**
 * @generated from message city.wargame.v1.PickPointsRequest
 */
export declare class PickPointsRequest extends Message<PickPointsRequest> {
  /**
   * @generated from field: city.wargame.v1.Camp camp = 1;
   */
  camp: Camp;

  /**
   * @generated from field: repeated city.wargame.v1.Point points = 2;
   */
  points: Point[];

  constructor(data?: PartialMessage<PickPointsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.PickPointsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PickPointsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PickPointsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PickPointsRequest;

  static equals(a: PickPointsRequest | PlainMessage<PickPointsRequest> | undefined, b: PickPointsRequest | PlainMessage<PickPointsRequest> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.PickPointsResponse
 */
export declare class PickPointsResponse extends Message<PickPointsResponse> {
  constructor(data?: PartialMessage<PickPointsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.PickPointsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PickPointsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PickPointsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PickPointsResponse;

  static equals(a: PickPointsResponse | PlainMessage<PickPointsResponse> | undefined, b: PickPointsResponse | PlainMessage<PickPointsResponse> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetPickPointsRequest
 */
export declare class GetPickPointsRequest extends Message<GetPickPointsRequest> {
  /**
   * @generated from field: city.wargame.v1.Camp camp = 1;
   */
  camp: Camp;

  constructor(data?: PartialMessage<GetPickPointsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetPickPointsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPickPointsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPickPointsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPickPointsRequest;

  static equals(a: GetPickPointsRequest | PlainMessage<GetPickPointsRequest> | undefined, b: GetPickPointsRequest | PlainMessage<GetPickPointsRequest> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetPickPointsResponse
 */
export declare class GetPickPointsResponse extends Message<GetPickPointsResponse> {
  /**
   * @generated from field: repeated city.wargame.v1.Point points = 1;
   */
  points: Point[];

  constructor(data?: PartialMessage<GetPickPointsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetPickPointsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPickPointsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPickPointsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPickPointsResponse;

  static equals(a: GetPickPointsResponse | PlainMessage<GetPickPointsResponse> | undefined, b: GetPickPointsResponse | PlainMessage<GetPickPointsResponse> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.SetScoreWeightRequest
 */
export declare class SetScoreWeightRequest extends Message<SetScoreWeightRequest> {
  /**
   * @generated from field: double money = 1;
   */
  money: number;

  /**
   * @generated from field: double population_loss = 2;
   */
  populationLoss: number;

  /**
   * @generated from field: double elec_power = 3;
   */
  elecPower: number;

  /**
   * @generated from field: double elec_distory_1 = 4;
   */
  elecDistory1: number;

  /**
   * @generated from field: double water_distory_1 = 5;
   */
  waterDistory1: number;

  /**
   * @generated from field: double bs_distory_1 = 6;
   */
  bsDistory1: number;

  /**
   * @generated from field: double traffic_distory_1 = 7;
   */
  trafficDistory1: number;

  /**
   * @generated from field: double elec_distory_2 = 8;
   */
  elecDistory2: number;

  /**
   * @generated from field: double water_distory_2 = 9;
   */
  waterDistory2: number;

  /**
   * @generated from field: double bs_distory_2 = 10;
   */
  bsDistory2: number;

  /**
   * @generated from field: double traffic_distory_2 = 11;
   */
  trafficDistory2: number;

  /**
   * @generated from field: double elec_distory_3 = 12;
   */
  elecDistory3: number;

  /**
   * @generated from field: double water_distory_3 = 13;
   */
  waterDistory3: number;

  /**
   * @generated from field: double bs_distory_3 = 14;
   */
  bsDistory3: number;

  /**
   * @generated from field: double traffic_distory_3 = 15;
   */
  trafficDistory3: number;

  /**
   * @generated from field: double defense_success = 16;
   */
  defenseSuccess: number;

  constructor(data?: PartialMessage<SetScoreWeightRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.SetScoreWeightRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetScoreWeightRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetScoreWeightRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetScoreWeightRequest;

  static equals(a: SetScoreWeightRequest | PlainMessage<SetScoreWeightRequest> | undefined, b: SetScoreWeightRequest | PlainMessage<SetScoreWeightRequest> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.SetScoreWeightResponse
 */
export declare class SetScoreWeightResponse extends Message<SetScoreWeightResponse> {
  constructor(data?: PartialMessage<SetScoreWeightResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.SetScoreWeightResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetScoreWeightResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetScoreWeightResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetScoreWeightResponse;

  static equals(a: SetScoreWeightResponse | PlainMessage<SetScoreWeightResponse> | undefined, b: SetScoreWeightResponse | PlainMessage<SetScoreWeightResponse> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GiveDefenseOrderRequest
 */
export declare class GiveDefenseOrderRequest extends Message<GiveDefenseOrderRequest> {
  /**
   * @generated from field: double weight_radius = 1;
   */
  weightRadius: number;

  /**
   * @generated from field: double weight_550 = 2;
   */
  weight550: number;

  /**
   * @generated from field: double weight_220 = 3;
   */
  weight220: number;

  /**
   * @generated from field: double weight_110 = 4;
   */
  weight110: number;

  /**
   * @generated from field: double prob_threshold = 5;
   */
  probThreshold: number;

  constructor(data?: PartialMessage<GiveDefenseOrderRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GiveDefenseOrderRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GiveDefenseOrderRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GiveDefenseOrderRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GiveDefenseOrderRequest;

  static equals(a: GiveDefenseOrderRequest | PlainMessage<GiveDefenseOrderRequest> | undefined, b: GiveDefenseOrderRequest | PlainMessage<GiveDefenseOrderRequest> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GiveDefenseOrderResponse
 */
export declare class GiveDefenseOrderResponse extends Message<GiveDefenseOrderResponse> {
  constructor(data?: PartialMessage<GiveDefenseOrderResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GiveDefenseOrderResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GiveDefenseOrderResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GiveDefenseOrderResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GiveDefenseOrderResponse;

  static equals(a: GiveDefenseOrderResponse | PlainMessage<GiveDefenseOrderResponse> | undefined, b: GiveDefenseOrderResponse | PlainMessage<GiveDefenseOrderResponse> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetHitHistoryRequest
 */
export declare class GetHitHistoryRequest extends Message<GetHitHistoryRequest> {
  constructor(data?: PartialMessage<GetHitHistoryRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetHitHistoryRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetHitHistoryRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetHitHistoryRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetHitHistoryRequest;

  static equals(a: GetHitHistoryRequest | PlainMessage<GetHitHistoryRequest> | undefined, b: GetHitHistoryRequest | PlainMessage<GetHitHistoryRequest> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetHitHistoryResponse
 */
export declare class GetHitHistoryResponse extends Message<GetHitHistoryResponse> {
  /**
   * @generated from field: repeated city.wargame.v1.HitHistory histories = 1;
   */
  histories: HitHistory[];

  constructor(data?: PartialMessage<GetHitHistoryResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetHitHistoryResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetHitHistoryResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetHitHistoryResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetHitHistoryResponse;

  static equals(a: GetHitHistoryResponse | PlainMessage<GetHitHistoryResponse> | undefined, b: GetHitHistoryResponse | PlainMessage<GetHitHistoryResponse> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetRecoPointsRequest
 */
export declare class GetRecoPointsRequest extends Message<GetRecoPointsRequest> {
  /**
   * @generated from field: city.wargame.v1.Camp camp = 1;
   */
  camp: Camp;

  /**
   * @generated from field: city.wargame.v1.RecoAlgoType type = 2;
   */
  type: RecoAlgoType;

  constructor(data?: PartialMessage<GetRecoPointsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetRecoPointsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRecoPointsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRecoPointsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRecoPointsRequest;

  static equals(a: GetRecoPointsRequest | PlainMessage<GetRecoPointsRequest> | undefined, b: GetRecoPointsRequest | PlainMessage<GetRecoPointsRequest> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetRecoPointsResponse
 */
export declare class GetRecoPointsResponse extends Message<GetRecoPointsResponse> {
  /**
   * @generated from field: repeated city.wargame.v1.RecoPoint points = 1;
   */
  points: RecoPoint[];

  constructor(data?: PartialMessage<GetRecoPointsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetRecoPointsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRecoPointsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRecoPointsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRecoPointsResponse;

  static equals(a: GetRecoPointsResponse | PlainMessage<GetRecoPointsResponse> | undefined, b: GetRecoPointsResponse | PlainMessage<GetRecoPointsResponse> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetStepRequest
 */
export declare class GetStepRequest extends Message<GetStepRequest> {
  constructor(data?: PartialMessage<GetStepRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetStepRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetStepRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetStepRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetStepRequest;

  static equals(a: GetStepRequest | PlainMessage<GetStepRequest> | undefined, b: GetStepRequest | PlainMessage<GetStepRequest> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetStepResponse
 */
export declare class GetStepResponse extends Message<GetStepResponse> {
  /**
   * @generated from field: int32 step = 1;
   */
  step: number;

  /**
   * @generated from field: bool red_pick_ready = 2;
   */
  redPickReady: boolean;

  /**
   * @generated from field: bool blue_pick_ready = 3;
   */
  bluePickReady: boolean;

  /**
   * @generated from field: int32 round = 4;
   */
  round: number;

  /**
   * @generated from field: bool is_game_started = 5;
   */
  isGameStarted: boolean;

  constructor(data?: PartialMessage<GetStepResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetStepResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetStepResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetStepResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetStepResponse;

  static equals(a: GetStepResponse | PlainMessage<GetStepResponse> | undefined, b: GetStepResponse | PlainMessage<GetStepResponse> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetCasualtiesRequest
 */
export declare class GetCasualtiesRequest extends Message<GetCasualtiesRequest> {
  constructor(data?: PartialMessage<GetCasualtiesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetCasualtiesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCasualtiesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCasualtiesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCasualtiesRequest;

  static equals(a: GetCasualtiesRequest | PlainMessage<GetCasualtiesRequest> | undefined, b: GetCasualtiesRequest | PlainMessage<GetCasualtiesRequest> | undefined): boolean;
}

/**
 * @generated from message city.wargame.v1.GetCasualtiesResponse
 */
export declare class GetCasualtiesResponse extends Message<GetCasualtiesResponse> {
  /**
   * @generated from field: repeated city.wargame.v1.Casualty casualties = 1;
   */
  casualties: Casualty[];

  constructor(data?: PartialMessage<GetCasualtiesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.wargame.v1.GetCasualtiesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCasualtiesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCasualtiesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCasualtiesResponse;

  static equals(a: GetCasualtiesResponse | PlainMessage<GetCasualtiesResponse> | undefined, b: GetCasualtiesResponse | PlainMessage<GetCasualtiesResponse> | undefined): boolean;
}

