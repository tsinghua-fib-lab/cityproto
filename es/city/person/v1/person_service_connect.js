// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/person/v1/person_service.proto (package city.person.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddPersonRequest, AddPersonResponse, GetPersonByLongLatBBoxRequest, GetPersonByLongLatBBoxResponse, GetPersonRequest, GetPersonResponse, SetScheduleRequest, SetScheduleResponse } from "./person_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.person.v1.PersonService
 */
export const PersonService = {
  typeName: "city.person.v1.PersonService",
  methods: {
    /**
     * 获取person信息
     * Get person information
     *
     * @generated from rpc city.person.v1.PersonService.GetPerson
     */
    getPerson: {
      name: "GetPerson",
      I: GetPersonRequest,
      O: GetPersonResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 新增person 传入person初始位置、目的地表、属性 返回personid
     * Add a new person. Input person's initial location, destination table, and attributes, return personid
     *
     * @generated from rpc city.person.v1.PersonService.AddPerson
     */
    addPerson: {
      name: "AddPerson",
      I: AddPersonRequest,
      O: AddPersonResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 修改person的schedule 传入personid、目的地表
     * Set person's schedule. Input personid and destination table
     *
     * @generated from rpc city.person.v1.PersonService.SetSchedule
     */
    setSchedule: {
      name: "SetSchedule",
      I: SetScheduleRequest,
      O: SetScheduleResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取特定区域内的person
     * Get persons in a specific region
     *
     * @generated from rpc city.person.v1.PersonService.GetPersonByLongLatBBox
     */
    getPersonByLongLatBBox: {
      name: "GetPersonByLongLatBBox",
      I: GetPersonByLongLatBBoxRequest,
      O: GetPersonByLongLatBBoxResponse,
      kind: MethodKind.Unary,
    },
  }
};

