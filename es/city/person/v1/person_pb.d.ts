// @generated by protoc-gen-es v1.6.0
// @generated from file city/person/v1/person.proto (package city.person.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Position } from "../../geo/v2/geo_pb.js";
import type { Schedule } from "../../trip/v2/trip_pb.js";

/**
 * 智能体教育等级
 *
 * @generated from enum city.person.v1.Education
 */
export declare enum Education {
  /**
   * 未指定
   *
   * @generated from enum value: EDUCATION_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 博士
   *
   * @generated from enum value: EDUCATION_DOCTOR = 1;
   */
  DOCTOR = 1,

  /**
   * 硕士
   *
   * @generated from enum value: EDUCATION_MASTER = 2;
   */
  MASTER = 2,

  /**
   * 本科
   *
   * @generated from enum value: EDUCATION_BACHELOR = 3;
   */
  BACHELOR = 3,

  /**
   * 高中
   *
   * @generated from enum value: EDUCATION_HIGH_SCHOOL = 4;
   */
  HIGH_SCHOOL = 4,

  /**
   * 初中
   *
   * @generated from enum value: EDUCATION_JUNIOR_HIGH_SCHOOL = 5;
   */
  JUNIOR_HIGH_SCHOOL = 5,

  /**
   * 小学
   *
   * @generated from enum value: EDUCATION_PRIMARY_SCHOOL = 6;
   */
  PRIMARY_SCHOOL = 6,

  /**
   * 大专
   *
   * @generated from enum value: EDUCATION_COLLEGE = 7;
   */
  COLLEGE = 7,
}

/**
 * 智能体性别
 *
 * @generated from enum city.person.v1.Gender
 */
export declare enum Gender {
  /**
   * 未指定
   *
   * @generated from enum value: GENDER_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 男性
   *
   * @generated from enum value: GENDER_MALE = 1;
   */
  MALE = 1,

  /**
   * 女性
   *
   * @generated from enum value: GENDER_FEMALE = 2;
   */
  FEMALE = 2,
}

/**
 * 智能体消费水平
 *
 * @generated from enum city.person.v1.Consumption
 */
export declare enum Consumption {
  /**
   * 未指定
   *
   * @generated from enum value: CONSUMPTION_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * 低
   *
   * @generated from enum value: CONSUMPTION_LOW = 1;
   */
  LOW = 1,

  /**
   * 较低
   *
   * @generated from enum value: CONSUMPTION_RELATIVELY_LOW = 2;
   */
  RELATIVELY_LOW = 2,

  /**
   * 中等
   *
   * @generated from enum value: CONSUMPTION_MEDIUM = 3;
   */
  MEDIUM = 3,

  /**
   * 较高
   *
   * @generated from enum value: CONSUMPTION_RELATIVELY_HIGH = 4;
   */
  RELATIVELY_HIGH = 4,

  /**
   * 高
   *
   * @generated from enum value: CONSUMPTION_HIGH = 5;
   */
  HIGH = 5,
}

/**
 * 智能体属性（通用）
 *
 * @generated from message city.person.v1.PersonAttribute
 */
export declare class PersonAttribute extends Message<PersonAttribute> {
  /**
   * 单位: m，长度
   *
   * @generated from field: double length = 2;
   */
  length: number;

  /**
   * 单位: m，宽度
   *
   * @generated from field: double width = 3;
   */
  width: number;

  /**
   * 单位: m/s
   *
   * @generated from field: double max_speed = 4;
   */
  maxSpeed: number;

  /**
   * 单位: m/s^2, 最大加速度（正值）
   *
   * @generated from field: double max_acceleration = 5;
   */
  maxAcceleration: number;

  /**
   * 单位: m/s^2, 最大减速度（负值）
   *
   * @generated from field: double max_braking_acceleration = 6;
   */
  maxBrakingAcceleration: number;

  /**
   * 单位: m/s^2, 一般加速度（正值），要求小于最大加速度
   *
   * @generated from field: double usual_acceleration = 7;
   */
  usualAcceleration: number;

  /**
   * 单位: m/s^2, 一般减速度（负值），要求大于最大减速度
   *
   * @generated from field: double usual_braking_acceleration = 8;
   */
  usualBrakingAcceleration: number;

  constructor(data?: PartialMessage<PersonAttribute>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.PersonAttribute";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PersonAttribute;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PersonAttribute;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PersonAttribute;

  static equals(a: PersonAttribute | PlainMessage<PersonAttribute> | undefined, b: PersonAttribute | PlainMessage<PersonAttribute> | undefined): boolean;
}

/**
 * 车辆附加属性
 *
 * @generated from message city.person.v1.VehicleAttribute
 */
export declare class VehicleAttribute extends Message<VehicleAttribute> {
  /**
   * 单位: m, 完成变道所需路程
   *
   * @generated from field: double lane_change_length = 1;
   */
  laneChangeLength: number;

  /**
   * 单位：米，本车距离前车的最小距离
   *
   * @generated from field: double min_gap = 2;
   */
  minGap: number;

  constructor(data?: PartialMessage<VehicleAttribute>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.VehicleAttribute";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VehicleAttribute;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VehicleAttribute;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VehicleAttribute;

  static equals(a: VehicleAttribute | PlainMessage<VehicleAttribute> | undefined, b: VehicleAttribute | PlainMessage<VehicleAttribute> | undefined): boolean;
}

/**
 * 公交车附加属性
 *
 * @generated from message city.person.v1.BusAttribute
 */
export declare class BusAttribute extends Message<BusAttribute> {
  /**
   * 公交线路ID
   *
   * @generated from field: int32 line_id = 1;
   */
  lineId: number;

  /**
   * 公交车容量
   *
   * @generated from field: int32 capacity = 2;
   */
  capacity: number;

  constructor(data?: PartialMessage<BusAttribute>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.BusAttribute";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BusAttribute;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BusAttribute;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BusAttribute;

  static equals(a: BusAttribute | PlainMessage<BusAttribute> | undefined, b: BusAttribute | PlainMessage<BusAttribute> | undefined): boolean;
}

/**
 * 自行车附加属性
 *
 * @generated from message city.person.v1.BikeAttribute
 */
export declare class BikeAttribute extends Message<BikeAttribute> {
  constructor(data?: PartialMessage<BikeAttribute>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.BikeAttribute";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BikeAttribute;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BikeAttribute;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BikeAttribute;

  static equals(a: BikeAttribute | PlainMessage<BikeAttribute> | undefined, b: BikeAttribute | PlainMessage<BikeAttribute> | undefined): boolean;
}

/**
 * 智能体简介
 *
 * @generated from message city.person.v1.PersonProfile
 */
export declare class PersonProfile extends Message<PersonProfile> {
  /**
   * 年龄
   *
   * @generated from field: int32 age = 1;
   */
  age: number;

  /**
   * 教育水平
   *
   * @generated from field: city.person.v1.Education education = 2;
   */
  education: Education;

  /**
   * 性别
   *
   * @generated from field: city.person.v1.Gender gender = 3;
   */
  gender: Gender;

  /**
   * 消费水平
   *
   * @generated from field: city.person.v1.Consumption consumption = 4;
   */
  consumption: Consumption;

  constructor(data?: PartialMessage<PersonProfile>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.PersonProfile";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PersonProfile;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PersonProfile;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PersonProfile;

  static equals(a: PersonProfile | PlainMessage<PersonProfile> | undefined, b: PersonProfile | PlainMessage<PersonProfile> | undefined): boolean;
}

/**
 * 智能体
 *
 * @generated from message city.person.v1.Person
 */
export declare class Person extends Message<Person> {
  /**
   * 智能体ID
   *
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * 参数
   *
   * @generated from field: city.person.v1.PersonAttribute attribute = 2;
   */
  attribute?: PersonAttribute;

  /**
   * 初始位置
   *
   * @generated from field: city.geo.v2.Position home = 3;
   */
  home?: Position;

  /**
   * 初始日程
   *
   * @generated from field: repeated city.trip.v2.Schedule schedules = 4;
   */
  schedules: Schedule[];

  /**
   * 车辆附加属性
   *
   * @generated from field: optional city.person.v1.VehicleAttribute vehicle_attribute = 7;
   */
  vehicleAttribute?: VehicleAttribute;

  /**
   * 公交车附加属性
   *
   * @generated from field: optional city.person.v1.BusAttribute bus_attribute = 8;
   */
  busAttribute?: BusAttribute;

  /**
   * 自行车附加属性
   *
   * @generated from field: optional city.person.v1.BikeAttribute bike_attribute = 9;
   */
  bikeAttribute?: BikeAttribute;

  /**
   * [可空] 额外的标签（例如：抢修车类型->电网）
   *
   * @generated from field: map<string, string> labels = 10;
   */
  labels: { [key: string]: string };

  /**
   * [可空] 智能体简介
   *
   * @generated from field: optional city.person.v1.PersonProfile profile = 11;
   */
  profile?: PersonProfile;

  constructor(data?: PartialMessage<Person>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.Person";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Person;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Person;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Person;

  static equals(a: Person | PlainMessage<Person> | undefined, b: Person | PlainMessage<Person> | undefined): boolean;
}

/**
 * 智能体集合，对应一个智能体pb文件或一个智能体mongodb collection
 *
 * @generated from message city.person.v1.Persons
 */
export declare class Persons extends Message<Persons> {
  /**
   * @generated from field: repeated city.person.v1.Person persons = 1;
   */
  persons: Person[];

  constructor(data?: PartialMessage<Persons>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.Persons";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Persons;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Persons;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Persons;

  static equals(a: Persons | PlainMessage<Persons> | undefined, b: Persons | PlainMessage<Persons> | undefined): boolean;
}

