// @generated by protoc-gen-es v1.6.0
// @generated from file city/config/v1/config.proto (package city.config.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * MongoDB配置
 *
 * @generated from message city.config.v1.MongoPath
 */
export const MongoPath = proto3.makeMessageType(
  "city.config.v1.MongoPath",
  () => [
    { no: 1, name: "db", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "col", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * 输出目标PostgreSQL
 *
 * @generated from message city.config.v1.OutputTarget
 */
export const OutputTarget = proto3.makeMessageType(
  "city.config.v1.OutputTarget",
  () => [
    { no: 1, name: "sql", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

