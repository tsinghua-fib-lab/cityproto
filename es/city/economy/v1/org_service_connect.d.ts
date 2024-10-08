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
export declare const OrgService: {
  readonly typeName: "city.economy.v1.OrgService",
  readonly methods: {
    /**
     * 批量查询组织的经济情况（员工、岗位、资金、货物）
     *
     * @generated from rpc city.economy.v1.OrgService.GetOrg
     */
    readonly getOrg: {
      readonly name: "GetOrg",
      readonly I: typeof GetOrgRequest,
      readonly O: typeof GetOrgResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 批量修改组织的资金
     *
     * @generated from rpc city.economy.v1.OrgService.UpdateOrgMoney
     */
    readonly updateOrgMoney: {
      readonly name: "UpdateOrgMoney",
      readonly I: typeof UpdateOrgMoneyRequest,
      readonly O: typeof UpdateOrgMoneyResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 批量修改组织的货物
     *
     * @generated from rpc city.economy.v1.OrgService.UpdateOrgGoods
     */
    readonly updateOrgGoods: {
      readonly name: "UpdateOrgGoods",
      readonly I: typeof UpdateOrgGoodsRequest,
      readonly O: typeof UpdateOrgGoodsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 批量修改组织的员工
     *
     * @generated from rpc city.economy.v1.OrgService.UpdateOrgEmployee
     */
    readonly updateOrgEmployee: {
      readonly name: "UpdateOrgEmployee",
      readonly I: typeof UpdateOrgEmployeeRequest,
      readonly O: typeof UpdateOrgEmployeeResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 批量修改组织的岗位
     *
     * @generated from rpc city.economy.v1.OrgService.UpdateOrgJob
     */
    readonly updateOrgJob: {
      readonly name: "UpdateOrgJob",
      readonly I: typeof UpdateOrgJobRequest,
      readonly O: typeof UpdateOrgJobResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

