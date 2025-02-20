// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/economy/v2/org_service.proto (package city.economy.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddAgentRequest, AddAgentResponse, AddOrgRequest, AddOrgResponse, BatchDeltaUpdateRequest, BatchDeltaUpdateResponse, BatchGetRequest, BatchGetResponse, BatchSetRequest, BatchSetResponse, BatchUpdateRequest, BatchUpdateResponse, CalculateConsumptionRequest, CalculateConsumptionResponse, CalculateInterestRequest, CalculateInterestResponse, CalculateRealGDPRequest, CalculateRealGDPResponse, CalculateTaxesDueRequest, CalculateTaxesDueResponse, DeltaUpdateAgentRequest, DeltaUpdateAgentResponse, DeltaUpdateOrgRequest, DeltaUpdateOrgResponse, GetAgentRequest, GetAgentResponse, GetOrgRequest, GetOrgResponse, LoadEconomyEntitiesRequest, LoadEconomyEntitiesResponse, RemoveAgentRequest, RemoveAgentResponse, RemoveOrgRequest, RemoveOrgResponse, SaveEconomyEntitiesRequest, SaveEconomyEntitiesResponse, UpdateAgentRequest, UpdateAgentResponse, UpdateOrgRequest, UpdateOrgResponse } from "./org_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * OrgService 提供了经济系统中组织和代理的管理接口
 * 包括基本的CRUD操作、批量操作、增量更新和各种计算功能
 *
 * @generated from service city.economy.v2.OrgService
 */
export const OrgService = {
  typeName: "city.economy.v2.OrgService",
  methods: {
    /**
     * AddOrg 添加一个新的组织到系统中
     *
     * @generated from rpc city.economy.v2.OrgService.AddOrg
     */
    addOrg: {
      name: "AddOrg",
      I: AddOrgRequest,
      O: AddOrgResponse,
      kind: MethodKind.Unary,
    },
    /**
     * RemoveOrg 从系统中移除指定的组织
     *
     * @generated from rpc city.economy.v2.OrgService.RemoveOrg
     */
    removeOrg: {
      name: "RemoveOrg",
      I: RemoveOrgRequest,
      O: RemoveOrgResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetOrg 获取指定组织的完整信息
     *
     * @generated from rpc city.economy.v2.OrgService.GetOrg
     */
    getOrg: {
      name: "GetOrg",
      I: GetOrgRequest,
      O: GetOrgResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateOrg 更新指定组织的信息
     *
     * @generated from rpc city.economy.v2.OrgService.UpdateOrg
     */
    updateOrg: {
      name: "UpdateOrg",
      I: UpdateOrgRequest,
      O: UpdateOrgResponse,
      kind: MethodKind.Unary,
    },
    /**
     * AddAgent 添加一个新的代理到系统中
     *
     * @generated from rpc city.economy.v2.OrgService.AddAgent
     */
    addAgent: {
      name: "AddAgent",
      I: AddAgentRequest,
      O: AddAgentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * RemoveAgent 从系统中移除指定的代理
     *
     * @generated from rpc city.economy.v2.OrgService.RemoveAgent
     */
    removeAgent: {
      name: "RemoveAgent",
      I: RemoveAgentRequest,
      O: RemoveAgentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetAgent 获取指定代理的完整信息
     *
     * @generated from rpc city.economy.v2.OrgService.GetAgent
     */
    getAgent: {
      name: "GetAgent",
      I: GetAgentRequest,
      O: GetAgentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateAgent 更新指定代理的信息
     *
     * @generated from rpc city.economy.v2.OrgService.UpdateAgent
     */
    updateAgent: {
      name: "UpdateAgent",
      I: UpdateAgentRequest,
      O: UpdateAgentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * BatchGet 批量获取多个组织或代理的信息
     *
     * @generated from rpc city.economy.v2.OrgService.BatchGet
     */
    batchGet: {
      name: "BatchGet",
      I: BatchGetRequest,
      O: BatchGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * BatchUpdate 批量更新多个组织或代理的信息，只更新请求中指定的字段
     *
     * @generated from rpc city.economy.v2.OrgService.BatchUpdate
     */
    batchUpdate: {
      name: "BatchUpdate",
      I: BatchUpdateRequest,
      O: BatchUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * BatchSet 批量设置多个组织或代理的信息，完全替换所有字段
     *
     * @generated from rpc city.economy.v2.OrgService.BatchSet
     */
    batchSet: {
      name: "BatchSet",
      I: BatchSetRequest,
      O: BatchSetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeltaUpdateOrg 对组织进行增量更新
     *
     * @generated from rpc city.economy.v2.OrgService.DeltaUpdateOrg
     */
    deltaUpdateOrg: {
      name: "DeltaUpdateOrg",
      I: DeltaUpdateOrgRequest,
      O: DeltaUpdateOrgResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeltaUpdateAgent 对代理进行增量更新
     *
     * @generated from rpc city.economy.v2.OrgService.DeltaUpdateAgent
     */
    deltaUpdateAgent: {
      name: "DeltaUpdateAgent",
      I: DeltaUpdateAgentRequest,
      O: DeltaUpdateAgentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * BatchDeltaUpdate 批量进行增量更新
     *
     * @generated from rpc city.economy.v2.OrgService.BatchDeltaUpdate
     */
    batchDeltaUpdate: {
      name: "BatchDeltaUpdate",
      I: BatchDeltaUpdateRequest,
      O: BatchDeltaUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CalculateTaxesDue 计算应缴税额并可选择进行再分配
     *
     * @generated from rpc city.economy.v2.OrgService.CalculateTaxesDue
     */
    calculateTaxesDue: {
      name: "CalculateTaxesDue",
      I: CalculateTaxesDueRequest,
      O: CalculateTaxesDueResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CalculateConsumption 计算代理的消费情况
     *
     * @generated from rpc city.economy.v2.OrgService.CalculateConsumption
     */
    calculateConsumption: {
      name: "CalculateConsumption",
      I: CalculateConsumptionRequest,
      O: CalculateConsumptionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CalculateInterest 计算银行利息
     *
     * @generated from rpc city.economy.v2.OrgService.CalculateInterest
     */
    calculateInterest: {
      name: "CalculateInterest",
      I: CalculateInterestRequest,
      O: CalculateInterestResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CalculateRealGDP 计算实际GDP
     *
     * @generated from rpc city.economy.v2.OrgService.CalculateRealGDP
     */
    calculateRealGDP: {
      name: "CalculateRealGDP",
      I: CalculateRealGDPRequest,
      O: CalculateRealGDPResponse,
      kind: MethodKind.Unary,
    },
    /**
     * SaveEconomyEntities 保存经济系统的当前状态
     *
     * @generated from rpc city.economy.v2.OrgService.SaveEconomyEntities
     */
    saveEconomyEntities: {
      name: "SaveEconomyEntities",
      I: SaveEconomyEntitiesRequest,
      O: SaveEconomyEntitiesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * LoadEconomyEntities 加载经济系统的状态
     *
     * @generated from rpc city.economy.v2.OrgService.LoadEconomyEntities
     */
    loadEconomyEntities: {
      name: "LoadEconomyEntities",
      I: LoadEconomyEntitiesRequest,
      O: LoadEconomyEntitiesResponse,
      kind: MethodKind.Unary,
    },
  }
};

