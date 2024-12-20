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
export declare const OrgService: {
  readonly typeName: "city.economy.v2.OrgService",
  readonly methods: {
    /**
     * 添加组织
     * add org
     *
     * @generated from rpc city.economy.v2.OrgService.AddOrg
     */
    readonly addOrg: {
      readonly name: "AddOrg",
      readonly I: typeof AddOrgRequest,
      readonly O: typeof AddOrgResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 移除组织
     * remove org
     *
     * @generated from rpc city.economy.v2.OrgService.RemoveOrg
     */
    readonly removeOrg: {
      readonly name: "RemoveOrg",
      readonly I: typeof RemoveOrgRequest,
      readonly O: typeof RemoveOrgResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 添加Agent
     * add agent
     *
     * @generated from rpc city.economy.v2.OrgService.AddAgent
     */
    readonly addAgent: {
      readonly name: "AddAgent",
      readonly I: typeof AddAgentRequest,
      readonly O: typeof AddAgentResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * 移除Agent
     * remove agent
     *
     * @generated from rpc city.economy.v2.OrgService.RemoveAgent
     */
    readonly removeAgent: {
      readonly name: "RemoveAgent",
      readonly I: typeof RemoveAgentRequest,
      readonly O: typeof RemoveAgentResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Nominal GDP
     *
     * @generated from rpc city.economy.v2.OrgService.GetNominalGDP
     */
    readonly getNominalGDP: {
      readonly name: "GetNominalGDP",
      readonly I: typeof GetNominalGDPRequest,
      readonly O: typeof GetNominalGDPResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetNominalGDP
     */
    readonly setNominalGDP: {
      readonly name: "SetNominalGDP",
      readonly I: typeof SetNominalGDPRequest,
      readonly O: typeof SetNominalGDPResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Real GDP
     *
     * @generated from rpc city.economy.v2.OrgService.GetRealGDP
     */
    readonly getRealGDP: {
      readonly name: "GetRealGDP",
      readonly I: typeof GetRealGDPRequest,
      readonly O: typeof GetRealGDPResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetRealGDP
     */
    readonly setRealGDP: {
      readonly name: "SetRealGDP",
      readonly I: typeof SetRealGDPRequest,
      readonly O: typeof SetRealGDPResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Unemployment
     *
     * @generated from rpc city.economy.v2.OrgService.GetUnemployment
     */
    readonly getUnemployment: {
      readonly name: "GetUnemployment",
      readonly I: typeof GetUnemploymentRequest,
      readonly O: typeof GetUnemploymentResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetUnemployment
     */
    readonly setUnemployment: {
      readonly name: "SetUnemployment",
      readonly I: typeof SetUnemploymentRequest,
      readonly O: typeof SetUnemploymentResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Wages
     *
     * @generated from rpc city.economy.v2.OrgService.GetWages
     */
    readonly getWages: {
      readonly name: "GetWages",
      readonly I: typeof GetWagesRequest,
      readonly O: typeof GetWagesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetWages
     */
    readonly setWages: {
      readonly name: "SetWages",
      readonly I: typeof SetWagesRequest,
      readonly O: typeof SetWagesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Prices
     *
     * @generated from rpc city.economy.v2.OrgService.GetPrices
     */
    readonly getPrices: {
      readonly name: "GetPrices",
      readonly I: typeof GetPricesRequest,
      readonly O: typeof GetPricesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetPrices
     */
    readonly setPrices: {
      readonly name: "SetPrices",
      readonly I: typeof SetPricesRequest,
      readonly O: typeof SetPricesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Inventory
     *
     * @generated from rpc city.economy.v2.OrgService.GetInventory
     */
    readonly getInventory: {
      readonly name: "GetInventory",
      readonly I: typeof GetInventoryRequest,
      readonly O: typeof GetInventoryResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetInventory
     */
    readonly setInventory: {
      readonly name: "SetInventory",
      readonly I: typeof SetInventoryRequest,
      readonly O: typeof SetInventoryResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Price
     *
     * @generated from rpc city.economy.v2.OrgService.GetPrice
     */
    readonly getPrice: {
      readonly name: "GetPrice",
      readonly I: typeof GetPriceRequest,
      readonly O: typeof GetPriceResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetPrice
     */
    readonly setPrice: {
      readonly name: "SetPrice",
      readonly I: typeof SetPriceRequest,
      readonly O: typeof SetPriceResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Currency
     *
     * @generated from rpc city.economy.v2.OrgService.GetCurrency
     */
    readonly getCurrency: {
      readonly name: "GetCurrency",
      readonly I: typeof GetCurrencyRequest,
      readonly O: typeof GetCurrencyResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetCurrency
     */
    readonly setCurrency: {
      readonly name: "SetCurrency",
      readonly I: typeof SetCurrencyRequest,
      readonly O: typeof SetCurrencyResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Interest Rate
     *
     * @generated from rpc city.economy.v2.OrgService.GetInterestRate
     */
    readonly getInterestRate: {
      readonly name: "GetInterestRate",
      readonly I: typeof GetInterestRateRequest,
      readonly O: typeof GetInterestRateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetInterestRate
     */
    readonly setInterestRate: {
      readonly name: "SetInterestRate",
      readonly I: typeof SetInterestRateRequest,
      readonly O: typeof SetInterestRateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Bracket Cutoffs
     *
     * @generated from rpc city.economy.v2.OrgService.GetBracketCutoffs
     */
    readonly getBracketCutoffs: {
      readonly name: "GetBracketCutoffs",
      readonly I: typeof GetBracketCutoffsRequest,
      readonly O: typeof GetBracketCutoffsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetBracketCutoffs
     */
    readonly setBracketCutoffs: {
      readonly name: "SetBracketCutoffs",
      readonly I: typeof SetBracketCutoffsRequest,
      readonly O: typeof SetBracketCutoffsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Bracket Rates
     *
     * @generated from rpc city.economy.v2.OrgService.GetBracketRates
     */
    readonly getBracketRates: {
      readonly name: "GetBracketRates",
      readonly I: typeof GetBracketRatesRequest,
      readonly O: typeof GetBracketRatesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc city.economy.v2.OrgService.SetBracketRates
     */
    readonly setBracketRates: {
      readonly name: "SetBracketRates",
      readonly I: typeof SetBracketRatesRequest,
      readonly O: typeof SetBracketRatesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Taxes Due
     *
     * @generated from rpc city.economy.v2.OrgService.CalculateTaxesDue
     */
    readonly calculateTaxesDue: {
      readonly name: "CalculateTaxesDue",
      readonly I: typeof CalculateTaxesDueRequest,
      readonly O: typeof CalculateTaxesDueResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Consumption
     *
     * @generated from rpc city.economy.v2.OrgService.CalculateConsumption
     */
    readonly calculateConsumption: {
      readonly name: "CalculateConsumption",
      readonly I: typeof CalculateConsumptionRequest,
      readonly O: typeof CalculateConsumptionResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Consumption
     *
     * @generated from rpc city.economy.v2.OrgService.CalculateInterest
     */
    readonly calculateInterest: {
      readonly name: "CalculateInterest",
      readonly I: typeof CalculateInterestRequest,
      readonly O: typeof CalculateInterestResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Save
     *
     * @generated from rpc city.economy.v2.OrgService.SaveEconomyEntities
     */
    readonly saveEconomyEntities: {
      readonly name: "SaveEconomyEntities",
      readonly I: typeof SaveEconomyEntitiesRequest,
      readonly O: typeof SaveEconomyEntitiesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Load
     *
     * @generated from rpc city.economy.v2.OrgService.LoadEconomyEntities
     */
    readonly loadEconomyEntities: {
      readonly name: "LoadEconomyEntities",
      readonly I: typeof LoadEconomyEntitiesRequest,
      readonly O: typeof LoadEconomyEntitiesResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

