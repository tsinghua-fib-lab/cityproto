// @generated by protoc-gen-es v1.6.0
// @generated from file city/person/v1/person_service.proto (package city.person.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { PersonRuntime } from "./person_runtime_pb.js";
import type { Person } from "./person_pb.js";
import type { Schedule } from "../../trip/v2/trip_pb.js";
import type { Status } from "./motion_pb.js";
import type { LongLatBBox, Position } from "../../geo/v2/geo_pb.js";
import type { VehicleAction, VehicleEnv, VehicleRuntime } from "./vehicle_pb.js";

/**
 * 获取person信息请求
 * Request for getting person information
 *
 * @generated from message city.person.v1.GetPersonRequest
 */
export declare class GetPersonRequest extends Message<GetPersonRequest> {
  /**
   * person id
   *
   * @generated from field: int32 person_id = 1;
   */
  personId: number;

  constructor(data?: PartialMessage<GetPersonRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.GetPersonRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPersonRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPersonRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPersonRequest;

  static equals(a: GetPersonRequest | PlainMessage<GetPersonRequest> | undefined, b: GetPersonRequest | PlainMessage<GetPersonRequest> | undefined): boolean;
}

/**
 * 获取person信息响应
 * Response of getting person information
 *
 * @generated from message city.person.v1.GetPersonResponse
 */
export declare class GetPersonResponse extends Message<GetPersonResponse> {
  /**
   * @generated from field: city.person.v1.PersonRuntime person = 1;
   */
  person?: PersonRuntime;

  constructor(data?: PartialMessage<GetPersonResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.GetPersonResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPersonResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPersonResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPersonResponse;

  static equals(a: GetPersonResponse | PlainMessage<GetPersonResponse> | undefined, b: GetPersonResponse | PlainMessage<GetPersonResponse> | undefined): boolean;
}

/**
 * 新增person请求
 * Request for adding a new person
 *
 * @generated from message city.person.v1.AddPersonRequest
 */
export declare class AddPersonRequest extends Message<AddPersonRequest> {
  /**
   * 约定：person中不设置id
   * Convention: personid is not set here
   *
   * @generated from field: city.person.v1.Person person = 1;
   */
  person?: Person;

  constructor(data?: PartialMessage<AddPersonRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.AddPersonRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddPersonRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddPersonRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddPersonRequest;

  static equals(a: AddPersonRequest | PlainMessage<AddPersonRequest> | undefined, b: AddPersonRequest | PlainMessage<AddPersonRequest> | undefined): boolean;
}

/**
 * 新增person响应
 * Response of adding a new person
 *
 * @generated from message city.person.v1.AddPersonResponse
 */
export declare class AddPersonResponse extends Message<AddPersonResponse> {
  /**
   * 新增的person分配得到的ID
   * The ID assigned to the newly added person
   *
   * @generated from field: int32 person_id = 1;
   */
  personId: number;

  constructor(data?: PartialMessage<AddPersonResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.AddPersonResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddPersonResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddPersonResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddPersonResponse;

  static equals(a: AddPersonResponse | PlainMessage<AddPersonResponse> | undefined, b: AddPersonResponse | PlainMessage<AddPersonResponse> | undefined): boolean;
}

/**
 * 修改person的schedule请求
 * Request for setting person schedule
 *
 * @generated from message city.person.v1.SetScheduleRequest
 */
export declare class SetScheduleRequest extends Message<SetScheduleRequest> {
  /**
   * person id
   *
   * @generated from field: int32 person_id = 1;
   */
  personId: number;

  /**
   * 新的schedule（覆盖原有的schedule）
   * New schedule (overwrites the original schedule)
   *
   * @generated from field: repeated city.trip.v2.Schedule schedules = 2;
   */
  schedules: Schedule[];

  constructor(data?: PartialMessage<SetScheduleRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.SetScheduleRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetScheduleRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetScheduleRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetScheduleRequest;

  static equals(a: SetScheduleRequest | PlainMessage<SetScheduleRequest> | undefined, b: SetScheduleRequest | PlainMessage<SetScheduleRequest> | undefined): boolean;
}

/**
 * 修改person的schedule响应
 * Response of setting person schedule
 *
 * @generated from message city.person.v1.SetScheduleResponse
 */
export declare class SetScheduleResponse extends Message<SetScheduleResponse> {
  constructor(data?: PartialMessage<SetScheduleResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.SetScheduleResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetScheduleResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetScheduleResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetScheduleResponse;

  static equals(a: SetScheduleResponse | PlainMessage<SetScheduleResponse> | undefined, b: SetScheduleResponse | PlainMessage<SetScheduleResponse> | undefined): boolean;
}

/**
 * 获取多个person信息请求
 * Request for getting information of multiple persons
 *
 * @generated from message city.person.v1.GetPersonsRequest
 */
export declare class GetPersonsRequest extends Message<GetPersonsRequest> {
  /**
   * person id列表，为空则返回所有person
   * List of person ids, return all persons if empty
   *
   * @generated from field: repeated int32 person_ids = 1;
   */
  personIds: number[];

  /**
   * 过滤人的状态（状态为列表内的值的人不返回），即使包含在person_ids中
   * Filter person's status (person whose status is in the list will not be returned), even if included in person_ids
   *
   * @generated from field: repeated city.person.v1.Status exclude_statuses = 2;
   */
  excludeStatuses: Status[];

  /**
   * 设置是否返回base信息
   * Set whether to return base information
   *
   * @generated from field: bool return_base = 3;
   */
  returnBase: boolean;

  constructor(data?: PartialMessage<GetPersonsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.GetPersonsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPersonsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPersonsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPersonsRequest;

  static equals(a: GetPersonsRequest | PlainMessage<GetPersonsRequest> | undefined, b: GetPersonsRequest | PlainMessage<GetPersonsRequest> | undefined): boolean;
}

/**
 * 获取多个person信息响应
 * Response of getting information of multiple persons
 *
 * @generated from message city.person.v1.GetPersonsResponse
 */
export declare class GetPersonsResponse extends Message<GetPersonsResponse> {
  /**
   * person信息
   * person information
   *
   * @generated from field: repeated city.person.v1.PersonRuntime persons = 1;
   */
  persons: PersonRuntime[];

  constructor(data?: PartialMessage<GetPersonsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.GetPersonsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPersonsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPersonsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPersonsResponse;

  static equals(a: GetPersonsResponse | PlainMessage<GetPersonsResponse> | undefined, b: GetPersonsResponse | PlainMessage<GetPersonsResponse> | undefined): boolean;
}

/**
 * 获取特定区域内的person请求
 * Request for getting persons in region
 *
 * @generated from message city.person.v1.GetPersonByLongLatBBoxRequest
 */
export declare class GetPersonByLongLatBBoxRequest extends Message<GetPersonByLongLatBBoxRequest> {
  /**
   * 经纬度范围
   * longitude and latitude bounding box
   *
   * @generated from field: city.geo.v2.LongLatBBox bbox = 1;
   */
  bbox?: LongLatBBox;

  /**
   * 过滤人的状态（状态为列表内的值的人不返回）
   * Filter person's status (person whose status is in the list will not be returned)
   *
   * @generated from field: repeated city.person.v1.Status exclude_statuses = 2;
   */
  excludeStatuses: Status[];

  /**
   * 设置是否返回base信息
   * Set whether to return base information
   *
   * @generated from field: bool return_base = 3;
   */
  returnBase: boolean;

  constructor(data?: PartialMessage<GetPersonByLongLatBBoxRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.GetPersonByLongLatBBoxRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPersonByLongLatBBoxRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPersonByLongLatBBoxRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPersonByLongLatBBoxRequest;

  static equals(a: GetPersonByLongLatBBoxRequest | PlainMessage<GetPersonByLongLatBBoxRequest> | undefined, b: GetPersonByLongLatBBoxRequest | PlainMessage<GetPersonByLongLatBBoxRequest> | undefined): boolean;
}

/**
 * 获取特定区域内的person响应
 * Response of getting persons in region
 *
 * @generated from message city.person.v1.GetPersonByLongLatBBoxResponse
 */
export declare class GetPersonByLongLatBBoxResponse extends Message<GetPersonByLongLatBBoxResponse> {
  /**
   * 区域内的person的信息
   * Information of persons in the region
   *
   * @generated from field: repeated city.person.v1.PersonRuntime persons = 1;
   */
  persons: PersonRuntime[];

  constructor(data?: PartialMessage<GetPersonByLongLatBBoxResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.GetPersonByLongLatBBoxResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPersonByLongLatBBoxResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPersonByLongLatBBoxResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPersonByLongLatBBoxResponse;

  static equals(a: GetPersonByLongLatBBoxResponse | PlainMessage<GetPersonByLongLatBBoxResponse> | undefined, b: GetPersonByLongLatBBoxResponse | PlainMessage<GetPersonByLongLatBBoxResponse> | undefined): boolean;
}

/**
 * 获取所有车辆请求
 * Request for getting all vehicles
 *
 * @generated from message city.person.v1.GetAllVehiclesRequest
 */
export declare class GetAllVehiclesRequest extends Message<GetAllVehiclesRequest> {
  constructor(data?: PartialMessage<GetAllVehiclesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.GetAllVehiclesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAllVehiclesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAllVehiclesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAllVehiclesRequest;

  static equals(a: GetAllVehiclesRequest | PlainMessage<GetAllVehiclesRequest> | undefined, b: GetAllVehiclesRequest | PlainMessage<GetAllVehiclesRequest> | undefined): boolean;
}

/**
 * 获取所有车辆响应
 * Response of getting all vehicles
 *
 * @generated from message city.person.v1.GetAllVehiclesResponse
 */
export declare class GetAllVehiclesResponse extends Message<GetAllVehiclesResponse> {
  /**
   * 所有车辆的信息
   * Information of all vehicles
   *
   * @generated from field: repeated city.person.v1.VehicleRuntime vehicles = 1;
   */
  vehicles: VehicleRuntime[];

  constructor(data?: PartialMessage<GetAllVehiclesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.GetAllVehiclesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAllVehiclesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAllVehiclesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAllVehiclesResponse;

  static equals(a: GetAllVehiclesResponse | PlainMessage<GetAllVehiclesResponse> | undefined, b: GetAllVehiclesResponse | PlainMessage<GetAllVehiclesResponse> | undefined): boolean;
}

/**
 * 重置人的位置请求
 * Request for resetting person's position
 *
 * @generated from message city.person.v1.ResetPersonPositionRequest
 */
export declare class ResetPersonPositionRequest extends Message<ResetPersonPositionRequest> {
  /**
   * person id
   *
   * @generated from field: int32 person_id = 1;
   */
  personId: number;

  /**
   * 重置位置
   * reset position
   *
   * @generated from field: city.geo.v2.Position position = 2;
   */
  position?: Position;

  constructor(data?: PartialMessage<ResetPersonPositionRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.ResetPersonPositionRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResetPersonPositionRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResetPersonPositionRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResetPersonPositionRequest;

  static equals(a: ResetPersonPositionRequest | PlainMessage<ResetPersonPositionRequest> | undefined, b: ResetPersonPositionRequest | PlainMessage<ResetPersonPositionRequest> | undefined): boolean;
}

/**
 * 重置人的位置响应
 * Response of resetting person's position
 *
 * @generated from message city.person.v1.ResetPersonPositionResponse
 */
export declare class ResetPersonPositionResponse extends Message<ResetPersonPositionResponse> {
  constructor(data?: PartialMessage<ResetPersonPositionResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.ResetPersonPositionResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResetPersonPositionResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResetPersonPositionResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResetPersonPositionResponse;

  static equals(a: ResetPersonPositionResponse | PlainMessage<ResetPersonPositionResponse> | undefined, b: ResetPersonPositionResponse | PlainMessage<ResetPersonPositionResponse> | undefined): boolean;
}

/**
 * 设置由外部控制行为的vehicle请求（下一个step生效）
 * Request for setting vehicle controlled by external behavior
 *
 * @generated from message city.person.v1.SetControlledVehicleIDsRequest
 */
export declare class SetControlledVehicleIDsRequest extends Message<SetControlledVehicleIDsRequest> {
  /**
   * 由外部控制行为的vehicle id列表
   * List of vehicle ids controlled by external behavior
   *
   * @generated from field: repeated int32 vehicle_ids = 1;
   */
  vehicleIds: number[];

  constructor(data?: PartialMessage<SetControlledVehicleIDsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.SetControlledVehicleIDsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetControlledVehicleIDsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetControlledVehicleIDsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetControlledVehicleIDsRequest;

  static equals(a: SetControlledVehicleIDsRequest | PlainMessage<SetControlledVehicleIDsRequest> | undefined, b: SetControlledVehicleIDsRequest | PlainMessage<SetControlledVehicleIDsRequest> | undefined): boolean;
}

/**
 * 设置由外部控制行为的vehicle响应
 * Response of setting vehicle controlled by external behavior
 *
 * @generated from message city.person.v1.SetControlledVehicleIDsResponse
 */
export declare class SetControlledVehicleIDsResponse extends Message<SetControlledVehicleIDsResponse> {
  constructor(data?: PartialMessage<SetControlledVehicleIDsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.SetControlledVehicleIDsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetControlledVehicleIDsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetControlledVehicleIDsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetControlledVehicleIDsResponse;

  static equals(a: SetControlledVehicleIDsResponse | PlainMessage<SetControlledVehicleIDsResponse> | undefined, b: SetControlledVehicleIDsResponse | PlainMessage<SetControlledVehicleIDsResponse> | undefined): boolean;
}

/**
 * 获取由外部控制行为的vehicle信息请求
 * Request for getting information of vehicle controlled by external behavior
 *
 * @generated from message city.person.v1.FetchControlledVehicleEnvsRequest
 */
export declare class FetchControlledVehicleEnvsRequest extends Message<FetchControlledVehicleEnvsRequest> {
  constructor(data?: PartialMessage<FetchControlledVehicleEnvsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.FetchControlledVehicleEnvsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FetchControlledVehicleEnvsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FetchControlledVehicleEnvsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FetchControlledVehicleEnvsRequest;

  static equals(a: FetchControlledVehicleEnvsRequest | PlainMessage<FetchControlledVehicleEnvsRequest> | undefined, b: FetchControlledVehicleEnvsRequest | PlainMessage<FetchControlledVehicleEnvsRequest> | undefined): boolean;
}

/**
 * 获取由外部控制行为的vehicle信息响应
 * Response of getting information of vehicle controlled by external behavior
 *
 * @generated from message city.person.v1.FetchControlledVehicleEnvsResponse
 */
export declare class FetchControlledVehicleEnvsResponse extends Message<FetchControlledVehicleEnvsResponse> {
  /**
   * 由外部控制行为的vehicle信息
   * Information of vehicle controlled by external behavior
   *
   * @generated from field: repeated city.person.v1.VehicleEnv vehicle_envs = 1;
   */
  vehicleEnvs: VehicleEnv[];

  constructor(data?: PartialMessage<FetchControlledVehicleEnvsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.FetchControlledVehicleEnvsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FetchControlledVehicleEnvsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FetchControlledVehicleEnvsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FetchControlledVehicleEnvsResponse;

  static equals(a: FetchControlledVehicleEnvsResponse | PlainMessage<FetchControlledVehicleEnvsResponse> | undefined, b: FetchControlledVehicleEnvsResponse | PlainMessage<FetchControlledVehicleEnvsResponse> | undefined): boolean;
}

/**
 * 设置由外部控制行为的vehicle的行为请求
 * Request for setting behavior of vehicle controlled by external behavior
 *
 * @generated from message city.person.v1.SetControlledVehicleActionsRequest
 */
export declare class SetControlledVehicleActionsRequest extends Message<SetControlledVehicleActionsRequest> {
  /**
   * 由外部控制行为的vehicle的行为
   * Behavior of vehicle controlled by external behavior
   *
   * @generated from field: repeated city.person.v1.VehicleAction vehicle_actions = 1;
   */
  vehicleActions: VehicleAction[];

  constructor(data?: PartialMessage<SetControlledVehicleActionsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.SetControlledVehicleActionsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetControlledVehicleActionsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetControlledVehicleActionsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetControlledVehicleActionsRequest;

  static equals(a: SetControlledVehicleActionsRequest | PlainMessage<SetControlledVehicleActionsRequest> | undefined, b: SetControlledVehicleActionsRequest | PlainMessage<SetControlledVehicleActionsRequest> | undefined): boolean;
}

/**
 * 设置由外部控制行为的vehicle的行为响应
 * Response of setting behavior of vehicle controlled by external behavior
 *
 * @generated from message city.person.v1.SetControlledVehicleActionsResponse
 */
export declare class SetControlledVehicleActionsResponse extends Message<SetControlledVehicleActionsResponse> {
  constructor(data?: PartialMessage<SetControlledVehicleActionsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.SetControlledVehicleActionsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetControlledVehicleActionsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetControlledVehicleActionsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetControlledVehicleActionsResponse;

  static equals(a: SetControlledVehicleActionsResponse | PlainMessage<SetControlledVehicleActionsResponse> | undefined, b: SetControlledVehicleActionsResponse | PlainMessage<SetControlledVehicleActionsResponse> | undefined): boolean;
}

