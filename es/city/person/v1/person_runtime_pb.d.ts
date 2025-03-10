// @generated by protoc-gen-es v1.10.0
// @generated from file city/person/v1/person_runtime.proto (package city.person.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Person } from "./person_pb.js";
import type { PersonMotion } from "./motion_pb.js";
import type { Event } from "../../event/v2/event_pb.js";

/**
 * 智能体运行时信息
 *
 * @generated from message city.person.v1.PersonRuntime
 */
export declare class PersonRuntime extends Message<PersonRuntime> {
  /**
   * person信息
   * person information
   *
   * @generated from field: optional city.person.v1.Person base = 1;
   */
  base?: Person;

  /**
   * person运动信息
   * person motion information
   *
   * @generated from field: city.person.v1.PersonMotion motion = 2;
   */
  motion?: PersonMotion;

  /**
   * 事件信息
   * event information
   *
   * @generated from field: repeated city.event.v2.Event events = 3;
   */
  events: Event[];

  constructor(data?: PartialMessage<PersonRuntime>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "city.person.v1.PersonRuntime";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PersonRuntime;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PersonRuntime;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PersonRuntime;

  static equals(a: PersonRuntime | PlainMessage<PersonRuntime> | undefined, b: PersonRuntime | PlainMessage<PersonRuntime> | undefined): boolean;
}

