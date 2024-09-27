// @generated by protoc-gen-es v1.10.0
// @generated from file city/elec/input/v1/input.proto (package city.elec.input.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { LongLatPosition } from "../../../geo/v2/geo_pb.js";

/**
 * @generated from enum city.elec.input.v1.FacilityType
 */
export declare enum FacilityType {
  /**
   * 电网相关的基础设施分类
   *
   * @generated from enum value: FACILITY_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 发电站
   *
   * @generated from enum value: FACILITY_TYPE_POWER_STATION = 1;
   */
  POWER_STATION = 1,

  /**
   * 不同电压的变压器
   *
   * @generated from enum value: FACILITY_TYPE_TRANSFORMER_500 = 2;
   */
  TRANSFORMER_500 = 2,

  /**
   * @generated from enum value: FACILITY_TYPE_TRANSFORMER_220 = 3;
   */
  TRANSFORMER_220 = 3,

  /**
   * @generated from enum value: FACILITY_TYPE_TRANSFORMER_110 = 4;
   */
  TRANSFORMER_110 = 4,

  /**
   * @generated from enum value: FACILITY_TYPE_TRANSFORMER_10 = 5;
   */
  TRANSFORMER_10 = 5,

  /**
   * 通信基站
   *
   * @generated from enum value: FACILITY_TYPE_BASE_STATION = 6;
   */
  BASE_STATION = 6,

  /**
   * 网关
   *
   * @generated from enum value: FACILITY_TYPE_GATEWAY = 7;
   */
  GATEWAY = 7,

  /**
   * 排水水泵
   *
   * @generated from enum value: FACILITY_TYPE_DRAINAGE_PUMP = 8;
   */
  DRAINAGE_PUMP = 8,

  /**
   * 交通灯
   *
   * @generated from enum value: FACILITY_TYPE_TRAFFIC_LIGHT = 9;
   */
  TRAFFIC_LIGHT = 9,

  /**
   * AOI
   *
   * @generated from enum value: FACILITY_TYPE_AOI = 10;
   */
  AOI = 10,

  /**
   * 供水水泵
   *
   * @generated from enum value: FACILITY_TYPE_SUPPLY_PUMP = 11;
   */
  SUPPLY_PUMP = 11,
}

/**
 * @generated from message city.elec.input.v1.RepairStation
 */
export declare class RepairStation extends Message<RepairStation> {
  /**
   * @generated from field: int32 aoi_id = 1;
   */
  aoiId: number;

  /**
   * @generated from field: city.geo.v2.LongLatPosition position = 2;
   */
  position?: LongLatPosition;

  constructor(data?: PartialMessage<RepairStation>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.elec.input.v1.RepairStation";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RepairStation;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RepairStation;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RepairStation;

  static equals(a: RepairStation | PlainMessage<RepairStation> | undefined, b: RepairStation | PlainMessage<RepairStation> | undefined): boolean;
}

/**
 * @generated from message city.elec.input.v1.Facility
 */
export declare class Facility extends Message<Facility> {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: city.geo.v2.LongLatPosition position = 2;
   */
  position?: LongLatPosition;

  /**
   * @generated from field: city.elec.input.v1.FacilityType type = 3;
   */
  type: FacilityType;

  /**
   * 当前节点的邻居节点的id
   *
   * @generated from field: repeated int32 relation = 4;
   */
  relation: number[];

  /**
   * 在其它关联的网络中如水网使用时，可使用外部id
   * 对于负载，该值表示其在对应模拟中的id
   *
   * @generated from field: optional int32 foreign_id = 5;
   */
  foreignId?: number;

  /**
   * 对于电力设施，该值表示所在aoi id
   *
   * @generated from field: optional int32 aoi_id = 6;
   */
  aoiId?: number;

  /**
   * 对于10kv变压器组，该值表示变压器组中变压器的数量
   *
   * @generated from field: optional int32 num_transformer = 7;
   */
  numTransformer?: number;

  constructor(data?: PartialMessage<Facility>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.elec.input.v1.Facility";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Facility;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Facility;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Facility;

  static equals(a: Facility | PlainMessage<Facility> | undefined, b: Facility | PlainMessage<Facility> | undefined): boolean;
}

/**
 * 设施集合，对应于mongodb一个collection
 *
 * @generated from message city.elec.input.v1.Facilities
 */
export declare class Facilities extends Message<Facilities> {
  /**
   * @generated from field: repeated city.elec.input.v1.Facility facilities = 1;
   */
  facilities: Facility[];

  /**
   * @generated from field: repeated city.elec.input.v1.RepairStation repair_stations = 2;
   */
  repairStations: RepairStation[];

  constructor(data?: PartialMessage<Facilities>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.elec.input.v1.Facilities";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Facilities;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Facilities;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Facilities;

  static equals(a: Facilities | PlainMessage<Facilities> | undefined, b: Facilities | PlainMessage<Facilities> | undefined): boolean;
}

