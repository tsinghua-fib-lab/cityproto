// @generated by protoc-gen-es v1.10.0
// @generated from file city/streetview/v1/streetview.proto (package city.streetview.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { LongLatPosition } from "../../geo/v2/geo_pb.js";

/**
 * 街景图片描述
 *
 * @generated from message city.streetview.v1.StreetViewImage
 */
export const StreetViewImage = /*@__PURE__*/ proto3.makeMessageType(
  "city.streetview.v1.StreetViewImage",
  () => [
    { no: 1, name: "heading", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "object", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * 街景图片元数据
 *
 * @generated from message city.streetview.v1.StreetView
 */
export const StreetView = /*@__PURE__*/ proto3.makeMessageType(
  "city.streetview.v1.StreetView",
  () => [
    { no: 1, name: "lnglat", kind: "message", T: LongLatPosition },
    { no: 2, name: "images", kind: "message", T: StreetViewImage, repeated: true },
  ],
);

/**
 * 导出Message，对应一个MongoDB Collection
 *
 * @generated from message city.streetview.v1.StreetViews
 */
export const StreetViews = /*@__PURE__*/ proto3.makeMessageType(
  "city.streetview.v1.StreetViews",
  () => [
    { no: 1, name: "street_views", kind: "message", T: StreetView, repeated: true },
  ],
);

