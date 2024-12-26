// @generated by protoc-gen-es v1.10.0
// @generated from file city/economy/v2/org_service.proto (package city.economy.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Agent, Org, OrgType } from "./economy_pb.js";

/**
 * @generated from message city.economy.v2.AddOrgRequest
 */
export const AddOrgRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.AddOrgRequest",
  () => [
    { no: 1, name: "org", kind: "message", T: Org },
  ],
);

/**
 * @generated from message city.economy.v2.AddOrgResponse
 */
export const AddOrgResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.AddOrgResponse",
  [],
);

/**
 * @generated from message city.economy.v2.RemoveOrgRequest
 */
export const RemoveOrgRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.RemoveOrgRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.RemoveOrgResponse
 */
export const RemoveOrgResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.RemoveOrgResponse",
  [],
);

/**
 * @generated from message city.economy.v2.AddAgentRequest
 */
export const AddAgentRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.AddAgentRequest",
  () => [
    { no: 1, name: "agent", kind: "message", T: Agent },
  ],
);

/**
 * @generated from message city.economy.v2.AddAgentResponse
 */
export const AddAgentResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.AddAgentResponse",
  [],
);

/**
 * @generated from message city.economy.v2.RemoveAgentRequest
 */
export const RemoveAgentRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.RemoveAgentRequest",
  () => [
    { no: 1, name: "agent_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.RemoveAgentResponse
 */
export const RemoveAgentResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.RemoveAgentResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetNominalGDPRequest
 */
export const GetNominalGDPRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetNominalGDPRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetNominalGDPResponse
 */
export const GetNominalGDPResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetNominalGDPResponse",
  () => [
    { no: 1, name: "nominal_gdp", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetNominalGDPRequest
 */
export const SetNominalGDPRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetNominalGDPRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "nominal_gdp", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetNominalGDPResponse
 */
export const SetNominalGDPResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetNominalGDPResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetRealGDPRequest
 */
export const GetRealGDPRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetRealGDPRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetRealGDPResponse
 */
export const GetRealGDPResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetRealGDPResponse",
  () => [
    { no: 1, name: "real_gdp", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetRealGDPRequest
 */
export const SetRealGDPRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetRealGDPRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "real_gdp", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetRealGDPResponse
 */
export const SetRealGDPResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetRealGDPResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetUnemploymentRequest
 */
export const GetUnemploymentRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetUnemploymentRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetUnemploymentResponse
 */
export const GetUnemploymentResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetUnemploymentResponse",
  () => [
    { no: 1, name: "unemployment", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetUnemploymentRequest
 */
export const SetUnemploymentRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetUnemploymentRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "unemployment", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetUnemploymentResponse
 */
export const SetUnemploymentResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetUnemploymentResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetWagesRequest
 */
export const GetWagesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetWagesRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetWagesResponse
 */
export const GetWagesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetWagesResponse",
  () => [
    { no: 1, name: "wages", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetWagesRequest
 */
export const SetWagesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetWagesRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "wages", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetWagesResponse
 */
export const SetWagesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetWagesResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetPricesRequest
 */
export const GetPricesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetPricesRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetPricesResponse
 */
export const GetPricesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetPricesResponse",
  () => [
    { no: 1, name: "prices", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetPricesRequest
 */
export const SetPricesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetPricesRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "prices", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetPricesResponse
 */
export const SetPricesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetPricesResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetInventoryRequest
 */
export const GetInventoryRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetInventoryRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetInventoryResponse
 */
export const GetInventoryResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetInventoryResponse",
  () => [
    { no: 1, name: "inventory", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.SetInventoryRequest
 */
export const SetInventoryRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetInventoryRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "inventory", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.SetInventoryResponse
 */
export const SetInventoryResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetInventoryResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetPriceRequest
 */
export const GetPriceRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetPriceRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetPriceResponse
 */
export const GetPriceResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetPriceResponse",
  () => [
    { no: 1, name: "price", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * @generated from message city.economy.v2.SetPriceRequest
 */
export const SetPriceRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetPriceRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "price", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * @generated from message city.economy.v2.SetPriceResponse
 */
export const SetPriceResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetPriceResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetCurrencyRequest
 */
export const GetCurrencyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetCurrencyRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetCurrencyResponse
 */
export const GetCurrencyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetCurrencyResponse",
  () => [
    { no: 1, name: "currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * @generated from message city.economy.v2.SetCurrencyRequest
 */
export const SetCurrencyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetCurrencyRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * @generated from message city.economy.v2.SetCurrencyResponse
 */
export const SetCurrencyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetCurrencyResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetInterestRateRequest
 */
export const GetInterestRateRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetInterestRateRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetInterestRateResponse
 */
export const GetInterestRateResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetInterestRateResponse",
  () => [
    { no: 1, name: "interest_rate", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * @generated from message city.economy.v2.SetInterestRateRequest
 */
export const SetInterestRateRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetInterestRateRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "interest_rate", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * @generated from message city.economy.v2.SetInterestRateResponse
 */
export const SetInterestRateResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetInterestRateResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetBracketCutoffsRequest
 */
export const GetBracketCutoffsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetBracketCutoffsRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetBracketCutoffsResponse
 */
export const GetBracketCutoffsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetBracketCutoffsResponse",
  () => [
    { no: 1, name: "bracket_cutoffs", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetBracketCutoffsRequest
 */
export const SetBracketCutoffsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetBracketCutoffsRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "bracket_cutoffs", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetBracketCutoffsResponse
 */
export const SetBracketCutoffsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetBracketCutoffsResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetBracketRatesRequest
 */
export const GetBracketRatesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetBracketRatesRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetBracketRatesResponse
 */
export const GetBracketRatesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetBracketRatesResponse",
  () => [
    { no: 1, name: "bracket_rates", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetBracketRatesRequest
 */
export const SetBracketRatesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetBracketRatesRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "bracket_rates", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetBracketRatesResponse
 */
export const SetBracketRatesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetBracketRatesResponse",
  [],
);

/**
 * @generated from message city.economy.v2.CalculateTaxesDueRequest
 */
export const CalculateTaxesDueRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.CalculateTaxesDueRequest",
  () => [
    { no: 1, name: "government_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "agent_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 3, name: "incomes", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
    { no: 4, name: "enable_redistribution", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message city.economy.v2.CalculateTaxesDueResponse
 */
export const CalculateTaxesDueResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.CalculateTaxesDueResponse",
  () => [
    { no: 1, name: "taxes_due", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 2, name: "updated_incomes", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.CalculateConsumptionRequest
 */
export const CalculateConsumptionRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.CalculateConsumptionRequest",
  () => [
    { no: 1, name: "firm_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "agent_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 3, name: "demands", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.CalculateConsumptionResponse
 */
export const CalculateConsumptionResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.CalculateConsumptionResponse",
  () => [
    { no: 1, name: "remain_inventory", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "updated_currencies", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.CalculateInterestRequest
 */
export const CalculateInterestRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.CalculateInterestRequest",
  () => [
    { no: 1, name: "bank_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "agent_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.CalculateInterestResponse
 */
export const CalculateInterestResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.CalculateInterestResponse",
  () => [
    { no: 1, name: "total_interest", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 2, name: "updated_currencies", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SaveEconomyEntitiesRequest
 */
export const SaveEconomyEntitiesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SaveEconomyEntitiesRequest",
  () => [
    { no: 1, name: "file_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message city.economy.v2.SaveEconomyEntitiesResponse
 */
export const SaveEconomyEntitiesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SaveEconomyEntitiesResponse",
  () => [
    { no: 1, name: "agent_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 2, name: "org_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.LoadEconomyEntitiesRequest
 */
export const LoadEconomyEntitiesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.LoadEconomyEntitiesRequest",
  () => [
    { no: 1, name: "file_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message city.economy.v2.LoadEconomyEntitiesResponse
 */
export const LoadEconomyEntitiesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.LoadEconomyEntitiesResponse",
  () => [
    { no: 1, name: "agent_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 2, name: "org_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ],
);

/**
 * Consumption Currency
 *
 * @generated from message city.economy.v2.GetConsumptionCurrencyRequest
 */
export const GetConsumptionCurrencyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetConsumptionCurrencyRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetConsumptionCurrencyResponse
 */
export const GetConsumptionCurrencyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetConsumptionCurrencyResponse",
  () => [
    { no: 1, name: "consumption_currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetConsumptionCurrencyRequest
 */
export const SetConsumptionCurrencyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetConsumptionCurrencyRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "consumption_currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetConsumptionCurrencyResponse
 */
export const SetConsumptionCurrencyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetConsumptionCurrencyResponse",
  [],
);

/**
 * Consumption Propensity
 *
 * @generated from message city.economy.v2.GetConsumptionPropensityRequest
 */
export const GetConsumptionPropensityRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetConsumptionPropensityRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetConsumptionPropensityResponse
 */
export const GetConsumptionPropensityResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetConsumptionPropensityResponse",
  () => [
    { no: 1, name: "consumption_propensity", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetConsumptionPropensityRequest
 */
export const SetConsumptionPropensityRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetConsumptionPropensityRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "consumption_propensity", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetConsumptionPropensityResponse
 */
export const SetConsumptionPropensityResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetConsumptionPropensityResponse",
  [],
);

/**
 * Income Currency
 *
 * @generated from message city.economy.v2.GetIncomeCurrencyRequest
 */
export const GetIncomeCurrencyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetIncomeCurrencyRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetIncomeCurrencyResponse
 */
export const GetIncomeCurrencyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetIncomeCurrencyResponse",
  () => [
    { no: 1, name: "income_currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetIncomeCurrencyRequest
 */
export const SetIncomeCurrencyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetIncomeCurrencyRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "income_currency", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetIncomeCurrencyResponse
 */
export const SetIncomeCurrencyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetIncomeCurrencyResponse",
  [],
);

/**
 * Depression
 *
 * @generated from message city.economy.v2.GetDepressionRequest
 */
export const GetDepressionRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetDepressionRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetDepressionResponse
 */
export const GetDepressionResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetDepressionResponse",
  () => [
    { no: 1, name: "depression", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetDepressionRequest
 */
export const SetDepressionRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetDepressionRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "depression", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetDepressionResponse
 */
export const SetDepressionResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetDepressionResponse",
  [],
);

/**
 * Locus of Control
 *
 * @generated from message city.economy.v2.GetLocusControlRequest
 */
export const GetLocusControlRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetLocusControlRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetLocusControlResponse
 */
export const GetLocusControlResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetLocusControlResponse",
  () => [
    { no: 1, name: "locus_control", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetLocusControlRequest
 */
export const SetLocusControlRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetLocusControlRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "locus_control", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetLocusControlResponse
 */
export const SetLocusControlResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetLocusControlResponse",
  [],
);

/**
 * Working Hours
 *
 * @generated from message city.economy.v2.GetWorkingHoursRequest
 */
export const GetWorkingHoursRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetWorkingHoursRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message city.economy.v2.GetWorkingHoursResponse
 */
export const GetWorkingHoursResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetWorkingHoursResponse",
  () => [
    { no: 1, name: "working_hours", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetWorkingHoursRequest
 */
export const SetWorkingHoursRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetWorkingHoursRequest",
  () => [
    { no: 1, name: "org_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "working_hours", kind: "scalar", T: 2 /* ScalarType.FLOAT */, repeated: true },
  ],
);

/**
 * @generated from message city.economy.v2.SetWorkingHoursResponse
 */
export const SetWorkingHoursResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.SetWorkingHoursResponse",
  [],
);

/**
 * @generated from message city.economy.v2.GetOrgEntityIdsRequest
 */
export const GetOrgEntityIdsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetOrgEntityIdsRequest",
  () => [
    { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(OrgType) },
  ],
);

/**
 * @generated from message city.economy.v2.GetOrgEntityIdsResponse
 */
export const GetOrgEntityIdsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "city.economy.v2.GetOrgEntityIdsResponse",
  () => [
    { no: 1, name: "org_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ],
);

