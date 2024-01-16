// @generated by protoc-gen-connect-es v1.2.1
// @generated from file city/wargame/v1/wargame_service.proto (package city.wargame.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetCasualtiesRequest, GetCasualtiesResponse, GetHitHistoryRequest, GetHitHistoryResponse, GetPickPointsRequest, GetPickPointsResponse, GetRecoPointsRequest, GetRecoPointsResponse, GetStepRequest, GetStepResponse, GiveDefenseOrderRequest, GiveDefenseOrderResponse, PickPointsRequest, PickPointsResponse, SetScoreWeightRequest, SetScoreWeightResponse } from "./wargame_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service city.wargame.v1.WarGameService
 */
export const WarGameService = {
  typeName: "city.wargame.v1.WarGameService",
  methods: {
    /**
     * 地图选点
     *
     * @generated from rpc city.wargame.v1.WarGameService.PickPoints
     */
    pickPoints: {
      name: "PickPoints",
      I: PickPointsRequest,
      O: PickPointsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 获取当前轮选点
     *
     * @generated from rpc city.wargame.v1.WarGameService.GetPickPoints
     */
    getPickPoints: {
      name: "GetPickPoints",
      I: GetPickPointsRequest,
      O: GetPickPointsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 指定防守方条令
     *
     * @generated from rpc city.wargame.v1.WarGameService.GiveDefenseOrder
     */
    giveDefenseOrder: {
      name: "GiveDefenseOrder",
      I: GiveDefenseOrderRequest,
      O: GiveDefenseOrderResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 分值权重设定
     *
     * @generated from rpc city.wargame.v1.WarGameService.SetScoreWeight
     */
    setScoreWeight: {
      name: "SetScoreWeight",
      I: SetScoreWeightRequest,
      O: SetScoreWeightResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 打击历史获取
     *
     * @generated from rpc city.wargame.v1.WarGameService.GetHitHistory
     */
    getHitHistory: {
      name: "GetHitHistory",
      I: GetHitHistoryRequest,
      O: GetHitHistoryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 推荐选点获取
     *
     * @generated from rpc city.wargame.v1.WarGameService.GetRecoPoints
     */
    getRecoPoints: {
      name: "GetRecoPoints",
      I: GetRecoPointsRequest,
      O: GetRecoPointsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 当前步与状态获取
     *
     * @generated from rpc city.wargame.v1.WarGameService.GetStep
     */
    getStep: {
      name: "GetStep",
      I: GetStepRequest,
      O: GetStepResponse,
      kind: MethodKind.Unary,
    },
    /**
     * 人口损伤人数和aoi的id获取
     *
     * @generated from rpc city.wargame.v1.WarGameService.GetCasualties
     */
    getCasualties: {
      name: "GetCasualties",
      I: GetCasualtiesRequest,
      O: GetCasualtiesResponse,
      kind: MethodKind.Unary,
    },
  }
};
