// @generated by protoc-gen-es v1.6.0
// @generated from file city/person/v2/person.proto (package city.person.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Position } from "../../geo/v2/geo_pb.js";
import { Schedule } from "../../trip/v2/trip_pb.js";

/**
 * 车辆引擎类型
 * vehicle type
 *
 * @generated from enum city.person.v2.VehicleEngineType
 */
export const VehicleEngineType = proto3.makeEnum(
  "city.person.v2.VehicleEngineType",
  [
    {no: 0, name: "VEHICLE_ENGINE_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "VEHICLE_ENGINE_TYPE_FUEL", localName: "FUEL"},
    {no: 2, name: "VEHICLE_ENGINE_TYPE_ELECTRIC", localName: "ELECTRIC"},
  ],
);

/**
 * 公交车
 * Type of Bus
 *
 * @generated from enum city.person.v2.BusType
 */
export const BusType = proto3.makeEnum(
  "city.person.v2.BusType",
  [
    {no: 0, name: "BUS_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "BUS_TYPE_BUS", localName: "BUS"},
    {no: 2, name: "BUS_TYPE_SUBWAY", localName: "SUBWAY"},
  ],
);

/**
 * 智能体教育等级
 * Agent education level
 *
 * @generated from enum city.person.v2.Education
 */
export const Education = proto3.makeEnum(
  "city.person.v2.Education",
  [
    {no: 0, name: "EDUCATION_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "EDUCATION_DOCTOR", localName: "DOCTOR"},
    {no: 2, name: "EDUCATION_MASTER", localName: "MASTER"},
    {no: 3, name: "EDUCATION_BACHELOR", localName: "BACHELOR"},
    {no: 4, name: "EDUCATION_HIGH_SCHOOL", localName: "HIGH_SCHOOL"},
    {no: 5, name: "EDUCATION_JUNIOR_HIGH_SCHOOL", localName: "JUNIOR_HIGH_SCHOOL"},
    {no: 6, name: "EDUCATION_PRIMARY_SCHOOL", localName: "PRIMARY_SCHOOL"},
    {no: 7, name: "EDUCATION_COLLEGE", localName: "COLLEGE"},
  ],
);

/**
 * 智能体性别
 * agent gender
 *
 * @generated from enum city.person.v2.Gender
 */
export const Gender = proto3.makeEnum(
  "city.person.v2.Gender",
  [
    {no: 0, name: "GENDER_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "GENDER_MALE", localName: "MALE"},
    {no: 2, name: "GENDER_FEMALE", localName: "FEMALE"},
  ],
);

/**
 * 智能体消费水平
 * agent consumption level
 *
 * @generated from enum city.person.v2.Consumption
 */
export const Consumption = proto3.makeEnum(
  "city.person.v2.Consumption",
  [
    {no: 0, name: "CONSUMPTION_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "CONSUMPTION_LOW", localName: "LOW"},
    {no: 2, name: "CONSUMPTION_RELATIVELY_LOW", localName: "RELATIVELY_LOW"},
    {no: 3, name: "CONSUMPTION_MEDIUM", localName: "MEDIUM"},
    {no: 4, name: "CONSUMPTION_RELATIVELY_HIGH", localName: "RELATIVELY_HIGH"},
    {no: 5, name: "CONSUMPTION_HIGH", localName: "HIGH"},
  ],
);

/**
 * 智能体属性（通用）
 * Agent properties (general)
 *
 * @generated from message city.person.v2.PersonAttribute
 */
export const PersonAttribute = proto3.makeMessageType(
  "city.person.v2.PersonAttribute",
  [],
);

/**
 * @generated from message city.person.v2.VehicleEngineEfficiency
 */
export const VehicleEngineEfficiency = proto3.makeMessageType(
  "city.person.v2.VehicleEngineEfficiency",
  () => [
    { no: 1, name: "energy_conversion_efficiency", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * 车辆碳排附加属性
 * Carbon emission additional attributes for Vehicles
 *
 * @generated from message city.person.v2.EmissionAttribute
 */
export const EmissionAttribute = proto3.makeMessageType(
  "city.person.v2.EmissionAttribute",
  () => [
    { no: 1, name: "weight", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(VehicleEngineType) },
    { no: 3, name: "coefficient_drag", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "lambda_s", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 5, name: "frontal_area", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 6, name: "fuel_efficiency", kind: "message", T: VehicleEngineEfficiency, opt: true },
    { no: 7, name: "electric_efficiency", kind: "message", T: VehicleEngineEfficiency, opt: true },
  ],
);

/**
 * 车辆附加属性
 * Vehicle additional attributes
 *
 * @generated from message city.person.v2.VehicleAttribute
 */
export const VehicleAttribute = proto3.makeMessageType(
  "city.person.v2.VehicleAttribute",
  () => [
    { no: 1, name: "length", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "width", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "max_speed", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "max_acceleration", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 5, name: "max_braking_acceleration", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 6, name: "usual_acceleration", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 7, name: "usual_braking_acceleration", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 8, name: "lane_change_length", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 9, name: "min_gap", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 10, name: "headway", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 11, name: "model", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 12, name: "lane_max_speed_recognition_deviation", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 13, name: "emission_attribute", kind: "message", T: EmissionAttribute },
  ],
);

/**
 * 公交车附加属性
 * Bus additional attributes
 *
 * @generated from message city.person.v2.BusAttribute
 */
export const BusAttribute = proto3.makeMessageType(
  "city.person.v2.BusAttribute",
  () => [
    { no: 1, name: "subline_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "capacity", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "model", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "type", kind: "enum", T: proto3.getEnumType(BusType) },
  ],
);

/**
 * 行人附加属性
 * Pedestrian additional attributes
 *
 * @generated from message city.person.v2.PedestrianAttribute
 */
export const PedestrianAttribute = proto3.makeMessageType(
  "city.person.v2.PedestrianAttribute",
  () => [
    { no: 1, name: "speed", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "model", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * 自行车附加属性
 * Bike additional attributes
 *
 * @generated from message city.person.v2.BikeAttribute
 */
export const BikeAttribute = proto3.makeMessageType(
  "city.person.v2.BikeAttribute",
  () => [
    { no: 1, name: "speed", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "model", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * 智能体简介
 * agent profile
 *
 * @generated from message city.person.v2.PersonProfile
 */
export const PersonProfile = proto3.makeMessageType(
  "city.person.v2.PersonProfile",
  () => [
    { no: 1, name: "age", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "education", kind: "enum", T: proto3.getEnumType(Education) },
    { no: 3, name: "gender", kind: "enum", T: proto3.getEnumType(Gender) },
    { no: 4, name: "consumption", kind: "enum", T: proto3.getEnumType(Consumption) },
    { no: 5, name: "house_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * 智能体
 * agent
 *
 * @generated from message city.person.v2.Person
 */
export const Person = proto3.makeMessageType(
  "city.person.v2.Person",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "attribute", kind: "message", T: PersonAttribute },
    { no: 3, name: "home", kind: "message", T: Position },
    { no: 4, name: "schedules", kind: "message", T: Schedule, repeated: true },
    { no: 7, name: "vehicle_attribute", kind: "message", T: VehicleAttribute, opt: true },
    { no: 8, name: "bus_attribute", kind: "message", T: BusAttribute, opt: true },
    { no: 12, name: "pedestrian_attribute", kind: "message", T: PedestrianAttribute, opt: true },
    { no: 9, name: "bike_attribute", kind: "message", T: BikeAttribute, opt: true },
    { no: 10, name: "labels", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 11, name: "profile", kind: "message", T: PersonProfile, opt: true },
    { no: 13, name: "work", kind: "message", T: Position, opt: true },
    { no: 14, name: "output_when_sleep", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
  ],
);

/**
 * 智能体集合，对应一个智能体pb文件或一个智能体mongodb collection
 * Agent collection, corresponding to an agent pb file or an agent mongodb collection
 *
 * @generated from message city.person.v2.Persons
 */
export const Persons = proto3.makeMessageType(
  "city.person.v2.Persons",
  () => [
    { no: 1, name: "persons", kind: "message", T: Person, repeated: true },
  ],
);

