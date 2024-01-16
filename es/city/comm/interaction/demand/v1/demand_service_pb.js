// @generated by protoc-gen-es v1.6.0
// @generated from file city/comm/interaction/demand/v1/demand_service.proto (package city.comm.interaction.demand.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * 设置用户通信需求激增
 * 用户通信需求激增公式：
 * result=demand*multiple_times*exp(-power_times*(current_time-start_time))
 * demand: 用户正常通信需求
 * current_time: 当前时间, start_time: 开始激增时间
 *
 * @generated from message city.comm.interaction.demand.v1.SetDemandStatusRequest
 */
export const SetDemandStatusRequest = proto3.makeMessageType(
  "city.comm.interaction.demand.v1.SetDemandStatusRequest",
  () => [
    { no: 1, name: "multiple_times", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "power_times", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ],
);

/**
 * @generated from message city.comm.interaction.demand.v1.SetDemandStatusResponse
 */
export const SetDemandStatusResponse = proto3.makeMessageType(
  "city.comm.interaction.demand.v1.SetDemandStatusResponse",
  [],
);

