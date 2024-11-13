// @generated by protoc-gen-es v1.10.0
// @generated from file city/person/v1/vehicle.proto (package city.person.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { PersonMotion } from "./motion_pb.js";
import type { Journey } from "../../routing/v2/routing_pb.js";

/**
 * @generated from enum city.person.v1.VehicleRelation
 */
export declare enum VehicleRelation {
  /**
   * 未指定
   * unspecified
   *
   * @generated from enum value: VEHICLE_RELATION_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 当前车道前车
   * vehicle ahead in the current lane
   *
   * @generated from enum value: VEHICLE_RELATION_AHEAD = 1;
   */
  AHEAD = 1,

  /**
   * 当前车道后车
   * vehicle behind in the current lane
   *
   * @generated from enum value: VEHICLE_RELATION_BEHIND = 2;
   */
  BEHIND = 2,

  /**
   * 影子车道前车
   * vehicle ahead in the shadow lane
   *
   * @generated from enum value: VEHICLE_RELATION_SHADOW_AHEAD = 3;
   */
  SHADOW_AHEAD = 3,

  /**
   * 影子车道后车
   * vehicle behind in the shadow lane
   *
   * @generated from enum value: VEHICLE_RELATION_SHADOW_BEHIND = 4;
   */
  SHADOW_BEHIND = 4,

  /**
   * 当前车道左侧车道前车
   * vehicle ahead in the left lane
   *
   * @generated from enum value: VEHICLE_RELATION_LEFT_AHEAD = 5;
   */
  LEFT_AHEAD = 5,

  /**
   * 当前车道右侧车道前车
   * vehicle ahead in the right lane
   *
   * @generated from enum value: VEHICLE_RELATION_RIGHT_AHEAD = 6;
   */
  RIGHT_AHEAD = 6,

  /**
   * 当前车道左侧车道后车
   * vehicle behind in the left lane
   *
   * @generated from enum value: VEHICLE_RELATION_LEFT_BEHIND = 7;
   */
  LEFT_BEHIND = 7,

  /**
   * 当前车道右侧车道后车
   * vehicle behind in the right lane
   *
   * @generated from enum value: VEHICLE_RELATION_RIGHT_BEHIND = 8;
   */
  RIGHT_BEHIND = 8,
}

/**
 * 交通灯的状态
 * traffic light state
 *
 * @generated from enum city.person.v1.LightState
 */
export declare enum LightState {
  /**
   * 未指定
   * unspecified
   *
   * @generated from enum value: LIGHT_STATE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 红灯
   * red light
   *
   * @generated from enum value: LIGHT_STATE_RED = 1;
   */
  RED = 1,

  /**
   * 绿灯
   * green light
   *
   * @generated from enum value: LIGHT_STATE_GREEN = 2;
   */
  GREEN = 2,

  /**
   * 黄灯
   * yellow light
   *
   * @generated from enum value: LIGHT_STATE_YELLOW = 3;
   */
  YELLOW = 3,
}

/**
 * 变道相关的信息
 * lane change related information
 *
 * @generated from message city.person.v1.LC
 */
export declare class LC extends Message<LC> {
  /**
   * 影子车道ID（变道前的车道）
   * shadow lane id (lane before lane change)
   *
   * @generated from field: int32 shadow_lane_id = 1;
   */
  shadowLaneId: number;

  /**
   * 投影到影子车道的坐标
   * s coordinate projected to shadow lane
   *
   * @generated from field: double shadow_s = 2;
   */
  shadowS: number;

  /**
   * 变道过程车头相对于前进方向的偏转角（弧度，总是为正，0代表不转向）
   * deviation angle of the vehicle head relative to the forward direction during lane change (radians, always positive, 0 means no steering)
   *
   * @generated from field: double angle = 3;
   */
  angle: number;

  /**
   * 已完成的变道比例
   * completed ratio of lane change
   *
   * @generated from field: double completed_ratio = 4;
   */
  completedRatio: number;

  constructor(data?: PartialMessage<LC>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.LC";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LC;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LC;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LC;

  static equals(a: LC | PlainMessage<LC> | undefined, b: LC | PlainMessage<LC> | undefined): boolean;
}

/**
 * 车辆控制信息
 * vehicle control information
 *
 * @generated from message city.person.v1.VehicleAction
 */
export declare class VehicleAction extends Message<VehicleAction> {
  /**
   * 车辆编号
   * vehicle id
   *
   * @generated from field: int32 id = 4;
   */
  id: number;

  /**
   * 本轮更新中设定的加速度
   * acceleration set in this step
   *
   * @generated from field: double acc = 1;
   */
  acc: number;

  /**
   * 变道目标（可选，不设置代表不变道或保持变道状态）
   * lane change target (optional, not set means no lane change)
   *
   * @generated from field: optional int32 lc_target_id = 2;
   */
  lcTargetId?: number;

  /**
   * 变道过程的转向角度
   * steering angle during lane change
   *
   * @generated from field: double angle = 3;
   */
  angle: number;

  constructor(data?: PartialMessage<VehicleAction>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.VehicleAction";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VehicleAction;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VehicleAction;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VehicleAction;

  static equals(a: VehicleAction | PlainMessage<VehicleAction> | undefined, b: VehicleAction | PlainMessage<VehicleAction> | undefined): boolean;
}

/**
 * @generated from message city.person.v1.VehicleRuntime
 */
export declare class VehicleRuntime extends Message<VehicleRuntime> {
  /**
   * 基本运行时信息
   * basic runtime information
   *
   * @generated from field: city.person.v1.PersonMotion base = 1;
   */
  base?: PersonMotion;

  /**
   * 变道信息
   * lane change information
   *
   * @generated from field: optional city.person.v1.LC lc = 4;
   */
  lc?: LC;

  /**
   * 本轮车辆行为（获取车辆环境信息时不返回）
   * vehicle action in the step (not returned when getting vehicle environment information)
   *
   * @generated from field: optional city.person.v1.VehicleAction action = 5;
   */
  action?: VehicleAction;

  /**
   * 走过的里程
   * running distance
   *
   * @generated from field: double running_distance = 6;
   */
  runningDistance: number;

  /**
   * 走错路次数
   * number of going astray
   *
   * @generated from field: int32 num_going_astray = 7;
   */
  numGoingAstray: number;

  /**
   * 出发时刻
   * departure time
   *
   * @generated from field: double departure_time = 8;
   */
  departureTime: number;

  /**
   * 预计到达时刻（导航返回的eta+出发时刻）
   * estimated arrival time (eta returned by routing + departure time)
   *
   * @generated from field: double eta = 9;
   */
  eta: number;

  /**
   * 自由流下的预计到达时刻
   * estimated arrival time under free flow
   *
   * @generated from field: double eta_free_flow = 10;
   */
  etaFreeFlow: number;

  constructor(data?: PartialMessage<VehicleRuntime>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.VehicleRuntime";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VehicleRuntime;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VehicleRuntime;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VehicleRuntime;

  static equals(a: VehicleRuntime | PlainMessage<VehicleRuntime> | undefined, b: VehicleRuntime | PlainMessage<VehicleRuntime> | undefined): boolean;
}

/**
 * 观测到的车辆
 * observed vehicles
 *
 * @generated from message city.person.v1.ObservedVehicle
 */
export declare class ObservedVehicle extends Message<ObservedVehicle> {
  /**
   * 车辆编号
   * vehicle id
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * 当前的车辆运行时信息
   * current vehicle runtime information
   *
   * @generated from field: city.person.v1.PersonMotion motion = 2;
   */
  motion?: PersonMotion;

  /**
   * 相对距离
   * relative distance
   *
   * @generated from field: double relative_distance = 3;
   */
  relativeDistance: number;

  /**
   * 关系枚举
   * relation enumeration
   *
   * @generated from field: city.person.v1.VehicleRelation relation = 4;
   */
  relation: VehicleRelation;

  constructor(data?: PartialMessage<ObservedVehicle>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.ObservedVehicle";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ObservedVehicle;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ObservedVehicle;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ObservedVehicle;

  static equals(a: ObservedVehicle | PlainMessage<ObservedVehicle> | undefined, b: ObservedVehicle | PlainMessage<ObservedVehicle> | undefined): boolean;
}

/**
 * 观测到的车道
 * observed lane
 *
 * @generated from message city.person.v1.ObservedLane
 */
export declare class ObservedLane extends Message<ObservedLane> {
  /**
   * Lane ID
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * 是否限行
   * whether restricted
   *
   * @generated from field: bool restriction = 2;
   */
  restriction: boolean;

  /**
   * 交通灯状态
   * traffic light state
   *
   * @generated from field: city.person.v1.LightState light_state = 3;
   */
  lightState: LightState;

  /**
   * 交通灯剩余时间
   * remaining time of traffic light
   *
   * @generated from field: double light_remaining_time = 4;
   */
  lightRemainingTime: number;

  constructor(data?: PartialMessage<ObservedLane>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.ObservedLane";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ObservedLane;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ObservedLane;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ObservedLane;

  static equals(a: ObservedLane | PlainMessage<ObservedLane> | undefined, b: ObservedLane | PlainMessage<ObservedLane> | undefined): boolean;
}

/**
 * @generated from message city.person.v1.VehicleEnv
 */
export declare class VehicleEnv extends Message<VehicleEnv> {
  /**
   * 车辆编号
   * vehicle id
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * 当前的车辆运行时信息
   * current vehicle runtime information
   *
   * @generated from field: city.person.v1.VehicleRuntime runtime = 2;
   */
  runtime?: VehicleRuntime;

  /**
   * 当前的路径规划结果
   * journey: current routing result
   *
   * @generated from field: city.routing.v2.Journey journey = 3;
   */
  journey?: Journey;

  /**
   * 观测到的车辆
   * observed vehicles
   *
   * @generated from field: repeated city.person.v1.ObservedVehicle observed_vehicles = 4;
   */
  observedVehicles: ObservedVehicle[];

  /**
   * 观测到的车道状态
   * observed lane states
   *
   * @generated from field: repeated city.person.v1.ObservedLane observed_lanes = 5;
   */
  observedLanes: ObservedLane[];

  constructor(data?: PartialMessage<VehicleEnv>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.VehicleEnv";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VehicleEnv;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VehicleEnv;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VehicleEnv;

  static equals(a: VehicleEnv | PlainMessage<VehicleEnv> | undefined, b: VehicleEnv | PlainMessage<VehicleEnv> | undefined): boolean;
}

