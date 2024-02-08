// @generated by protoc-gen-es v1.6.0
// @generated from file city/geo/v2/geo.proto (package city.geo.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * WGS84经纬度坐标
 * WGS84 longitute and latitude coordinates
 *
 * @generated from message city.geo.v2.LongLatPosition
 */
export declare class LongLatPosition extends Message<LongLatPosition> {
  /**
   * 经度
   * longitude
   *
   * @generated from field: double longitude = 1;
   */
  longitude: number;

  /**
   * 纬度
   * latitude
   *
   * @generated from field: double latitude = 2;
   */
  latitude: number;

  constructor(data?: PartialMessage<LongLatPosition>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.geo.v2.LongLatPosition";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LongLatPosition;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LongLatPosition;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LongLatPosition;

  static equals(a: LongLatPosition | PlainMessage<LongLatPosition> | undefined, b: LongLatPosition | PlainMessage<LongLatPosition> | undefined): boolean;
}

/**
 * XY坐标
 * XY coordinates
 *
 * @generated from message city.geo.v2.XYPosition
 */
export declare class XYPosition extends Message<XYPosition> {
  /**
   * x坐标，单位米，对应经度
   * x coordinate, in meters, corresponding to longitude
   *
   * @generated from field: double x = 1;
   */
  x: number;

  /**
   * y坐标，单位米，对应纬度
   * y coordinate, in meters, corresponding to latitude
   *
   * @generated from field: double y = 2;
   */
  y: number;

  constructor(data?: PartialMessage<XYPosition>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.geo.v2.XYPosition";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): XYPosition;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): XYPosition;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): XYPosition;

  static equals(a: XYPosition | PlainMessage<XYPosition> | undefined, b: XYPosition | PlainMessage<XYPosition> | undefined): boolean;
}

/**
 * 地图坐标（车道+距离s）
 * Map coordinates (lane ID + distance s)
 *
 * @generated from message city.geo.v2.LanePosition
 */
export declare class LanePosition extends Message<LanePosition> {
  /**
   * 车道id
   * Lane ID
   *
   * @generated from field: int32 lane_id = 1;
   */
  laneId: number;

  /**
   * s是车道上的点到车道起点的距离
   * s is the distance from the point on the lane to the starting point of the lane
   *
   * @generated from field: double s = 2;
   */
  s: number;

  constructor(data?: PartialMessage<LanePosition>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.geo.v2.LanePosition";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LanePosition;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LanePosition;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LanePosition;

  static equals(a: LanePosition | PlainMessage<LanePosition> | undefined, b: LanePosition | PlainMessage<LanePosition> | undefined): boolean;
}

/**
 * 地图坐标（AOI）
 * Map coordinates (AOI)
 *
 * @generated from message city.geo.v2.AoiPosition
 */
export declare class AoiPosition extends Message<AoiPosition> {
  /**
   * AOI ID
   *
   * @generated from field: int32 aoi_id = 1;
   */
  aoiId: number;

  /**
   * POI ID，需要是aoi_id的子poi，否则该值无效
   * POI ID, needs to be a sub-poi of aoi_id, otherwise the value is invalid
   *
   * @generated from field: optional int32 poi_id = 2;
   */
  poiId?: number;

  constructor(data?: PartialMessage<AoiPosition>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.geo.v2.AoiPosition";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AoiPosition;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AoiPosition;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AoiPosition;

  static equals(a: AoiPosition | PlainMessage<AoiPosition> | undefined, b: AoiPosition | PlainMessage<AoiPosition> | undefined): boolean;
}

/**
 * 坐标，如果多种坐标同时存在，两两之间必须满足映射关系，同时逻辑坐标是必须提供的
 * Coordinates, if multiple coordinates exist at the same time, the mapping relationship between them must be satisfied, and logical coordinates must be provided.
 *
 * @generated from message city.geo.v2.Position
 */
export declare class Position extends Message<Position> {
  /**
   * 地图坐标AOI（必须提供其中之一）
   * Map coordinates AOI (one of these must be provided)
   *
   * @generated from field: optional city.geo.v2.LanePosition lane_position = 1;
   */
  lanePosition?: LanePosition;

  /**
   * 地图坐标Lane+S（必须提供其中之一）
   * Map coordinates Lane+S (one of these must be provided)
   *
   * @generated from field: optional city.geo.v2.AoiPosition aoi_position = 2;
   */
  aoiPosition?: AoiPosition;

  /**
   * WGS84经纬度坐标
   * WGS84 longitute and latitude coordinates
   *
   * @generated from field: optional city.geo.v2.LongLatPosition longlat_position = 3;
   */
  longlatPosition?: LongLatPosition;

  /**
   * XY坐标
   * XY coordinates
   *
   * @generated from field: optional city.geo.v2.XYPosition xy_position = 4;
   */
  xyPosition?: XYPosition;

  constructor(data?: PartialMessage<Position>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.geo.v2.Position";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Position;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Position;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Position;

  static equals(a: Position | PlainMessage<Position> | undefined, b: Position | PlainMessage<Position> | undefined): boolean;
}

/**
 * 经纬度矩形区域
 * latitude and longitude rectangular area
 *
 * @generated from message city.geo.v2.LongLatBBox
 */
export declare class LongLatBBox extends Message<LongLatBBox> {
  /**
   * 最小经度
   * minimum longitude
   *
   * @generated from field: double min_longitude = 1;
   */
  minLongitude: number;

  /**
   * 最小纬度
   * minimum latitude
   *
   * @generated from field: double min_latitude = 2;
   */
  minLatitude: number;

  /**
   * 最大经度
   * maximu longitude
   *
   * @generated from field: double max_longitude = 3;
   */
  maxLongitude: number;

  /**
   * 最大纬度
   * minimum longitude
   *
   * @generated from field: double max_latitude = 4;
   */
  maxLatitude: number;

  constructor(data?: PartialMessage<LongLatBBox>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.geo.v2.LongLatBBox";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LongLatBBox;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LongLatBBox;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LongLatBBox;

  static equals(a: LongLatBBox | PlainMessage<LongLatBBox> | undefined, b: LongLatBBox | PlainMessage<LongLatBBox> | undefined): boolean;
}

