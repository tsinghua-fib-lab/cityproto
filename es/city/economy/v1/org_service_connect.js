// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/economy/v1/org_service.proto (package city.economy.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetOrgRequest, GetOrgResponse, UpdateOrgEmployeeRequest, UpdateOrgEmployeeResponse, UpdateOrgGoodsRequest, UpdateOrgGoodsResponse, UpdateOrgJobRequest, UpdateOrgJobResponse, UpdateOrgMoneyRequest, UpdateOrgMoneyResponse } from "./org_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * 组织经济情况接口
 *
 * @generated from service city.economy.v1.OrgService
 */
export const OrgService = {
  typeName: "city.economy.v1.OrgService",
  methods: {
    /**
     * 批量查询组织的经济情况（员工、岗位、资金、货物）
     *
     * @generated from rpc city.economy.v1.OrgService.GetOrg
     */
    getOrg: {
      name: "GetOrg",
      I: GetOrgRequest,
      O: GetOrgResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 批量修改组织的资金
     *
     * @generated from rpc city.economy.v1.OrgService.UpdateOrgMoney
     */
    updateOrgMoney: {
      name: "UpdateOrgMoney",
      I: UpdateOrgMoneyRequest,
      O: UpdateOrgMoneyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 批量修改组织的货物
     *
     * @generated from rpc city.economy.v1.OrgService.UpdateOrgGoods
     */
    updateOrgGoods: {
      name: "UpdateOrgGoods",
      I: UpdateOrgGoodsRequest,
      O: UpdateOrgGoodsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 批量修改组织的员工
     *
     * @generated from rpc city.economy.v1.OrgService.UpdateOrgEmployee
     */
    updateOrgEmployee: {
      name: "UpdateOrgEmployee",
      I: UpdateOrgEmployeeRequest,
      O: UpdateOrgEmployeeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 批量修改组织的岗位
     *
     * @generated from rpc city.economy.v1.OrgService.UpdateOrgJob
     */
    updateOrgJob: {
      name: "UpdateOrgJob",
      I: UpdateOrgJobRequest,
      O: UpdateOrgJobResponse,
      kind: MethodKind.Unary,
    },
  }
};

