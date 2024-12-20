// @generated by protoc-gen-connect-es v1.5.0
// @generated from file city/economy/v2/org_service.proto (package city.economy.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddAgentRequest, AddAgentResponse, AddOrgRequest, AddOrgResponse, CalculateConsumptionRequest, CalculateConsumptionResponse, CalculateInterestRequest, CalculateInterestResponse, CalculateTaxesDueRequest, CalculateTaxesDueResponse, GetBracketCutoffsRequest, GetBracketCutoffsResponse, GetBracketRatesRequest, GetBracketRatesResponse, GetCurrencyRequest, GetCurrencyResponse, GetInterestRateRequest, GetInterestRateResponse, GetInventoryRequest, GetInventoryResponse, GetNominalGDPRequest, GetNominalGDPResponse, GetPriceRequest, GetPriceResponse, GetPricesRequest, GetPricesResponse, GetRealGDPRequest, GetRealGDPResponse, GetUnemploymentRequest, GetUnemploymentResponse, GetWagesRequest, GetWagesResponse, LoadEconomyEntitiesRequest, LoadEconomyEntitiesResponse, RemoveAgentRequest, RemoveAgentResponse, RemoveOrgRequest, RemoveOrgResponse, SaveEconomyEntitiesRequest, SaveEconomyEntitiesResponse, SetBracketCutoffsRequest, SetBracketCutoffsResponse, SetBracketRatesRequest, SetBracketRatesResponse, SetCurrencyRequest, SetCurrencyResponse, SetInterestRateRequest, SetInterestRateResponse, SetInventoryRequest, SetInventoryResponse, SetNominalGDPRequest, SetNominalGDPResponse, SetPriceRequest, SetPriceResponse, SetPricesRequest, SetPricesResponse, SetRealGDPRequest, SetRealGDPResponse, SetUnemploymentRequest, SetUnemploymentResponse, SetWagesRequest, SetWagesResponse } from "./org_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * 组织经济情况接口
 *
 * @generated from service city.economy.v2.OrgService
 */
export const OrgService = {
  typeName: "city.economy.v2.OrgService",
  methods: {
    /**
     * 添加组织
     * add org
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
     * 移除组织
     * remove org
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
     * 添加Agent
     * add agent
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
     * 移除Agent
     * remove agent
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
     * Nominal GDP
     *
     * @generated from rpc city.economy.v2.OrgService.GetNominalGDP
     */
    getNominalGDP: {
      name: "GetNominalGDP",
      I: GetNominalGDPRequest,
      O: GetNominalGDPResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetNominalGDP
     */
    setNominalGDP: {
      name: "SetNominalGDP",
      I: SetNominalGDPRequest,
      O: SetNominalGDPResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Real GDP
     *
     * @generated from rpc city.economy.v2.OrgService.GetRealGDP
     */
    getRealGDP: {
      name: "GetRealGDP",
      I: GetRealGDPRequest,
      O: GetRealGDPResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetRealGDP
     */
    setRealGDP: {
      name: "SetRealGDP",
      I: SetRealGDPRequest,
      O: SetRealGDPResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Unemployment
     *
     * @generated from rpc city.economy.v2.OrgService.GetUnemployment
     */
    getUnemployment: {
      name: "GetUnemployment",
      I: GetUnemploymentRequest,
      O: GetUnemploymentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetUnemployment
     */
    setUnemployment: {
      name: "SetUnemployment",
      I: SetUnemploymentRequest,
      O: SetUnemploymentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Wages
     *
     * @generated from rpc city.economy.v2.OrgService.GetWages
     */
    getWages: {
      name: "GetWages",
      I: GetWagesRequest,
      O: GetWagesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetWages
     */
    setWages: {
      name: "SetWages",
      I: SetWagesRequest,
      O: SetWagesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Prices
     *
     * @generated from rpc city.economy.v2.OrgService.GetPrices
     */
    getPrices: {
      name: "GetPrices",
      I: GetPricesRequest,
      O: GetPricesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetPrices
     */
    setPrices: {
      name: "SetPrices",
      I: SetPricesRequest,
      O: SetPricesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Inventory
     *
     * @generated from rpc city.economy.v2.OrgService.GetInventory
     */
    getInventory: {
      name: "GetInventory",
      I: GetInventoryRequest,
      O: GetInventoryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetInventory
     */
    setInventory: {
      name: "SetInventory",
      I: SetInventoryRequest,
      O: SetInventoryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Price
     *
     * @generated from rpc city.economy.v2.OrgService.GetPrice
     */
    getPrice: {
      name: "GetPrice",
      I: GetPriceRequest,
      O: GetPriceResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetPrice
     */
    setPrice: {
      name: "SetPrice",
      I: SetPriceRequest,
      O: SetPriceResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Currency
     *
     * @generated from rpc city.economy.v2.OrgService.GetCurrency
     */
    getCurrency: {
      name: "GetCurrency",
      I: GetCurrencyRequest,
      O: GetCurrencyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetCurrency
     */
    setCurrency: {
      name: "SetCurrency",
      I: SetCurrencyRequest,
      O: SetCurrencyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Interest Rate
     *
     * @generated from rpc city.economy.v2.OrgService.GetInterestRate
     */
    getInterestRate: {
      name: "GetInterestRate",
      I: GetInterestRateRequest,
      O: GetInterestRateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetInterestRate
     */
    setInterestRate: {
      name: "SetInterestRate",
      I: SetInterestRateRequest,
      O: SetInterestRateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Bracket Cutoffs
     *
     * @generated from rpc city.economy.v2.OrgService.GetBracketCutoffs
     */
    getBracketCutoffs: {
      name: "GetBracketCutoffs",
      I: GetBracketCutoffsRequest,
      O: GetBracketCutoffsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetBracketCutoffs
     */
    setBracketCutoffs: {
      name: "SetBracketCutoffs",
      I: SetBracketCutoffsRequest,
      O: SetBracketCutoffsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Bracket Rates
     *
     * @generated from rpc city.economy.v2.OrgService.GetBracketRates
     */
    getBracketRates: {
      name: "GetBracketRates",
      I: GetBracketRatesRequest,
      O: GetBracketRatesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetBracketRates
     */
    setBracketRates: {
      name: "SetBracketRates",
      I: SetBracketRatesRequest,
      O: SetBracketRatesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Taxes Due
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
     * Consumption
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
     * Consumption
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
     * Save
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
     * Load
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

