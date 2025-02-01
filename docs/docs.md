# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [city/clock/v1/clock_service.proto](#city_clock_v1_clock_service-proto)
    - [NowRequest](#city-clock-v1-NowRequest)
    - [NowResponse](#city-clock-v1-NowResponse)
  
    - [ClockService](#city-clock-v1-ClockService)
  
- [city/geo/v2/geo.proto](#city_geo_v2_geo-proto)
    - [AoiPosition](#city-geo-v2-AoiPosition)
    - [LanePosition](#city-geo-v2-LanePosition)
    - [LongLatBBox](#city-geo-v2-LongLatBBox)
    - [LongLatPosition](#city-geo-v2-LongLatPosition)
    - [Position](#city-geo-v2-Position)
    - [XYPosition](#city-geo-v2-XYPosition)
  
- [city/comm/input/v1/comm.proto](#city_comm_input_v1_comm-proto)
    - [CommDemand](#city-comm-input-v1-CommDemand)
    - [CommDemands](#city-comm-input-v1-CommDemands)
    - [Node](#city-comm-input-v1-Node)
    - [Nodes](#city-comm-input-v1-Nodes)
    - [Pump](#city-comm-input-v1-Pump)
    - [RepairStation](#city-comm-input-v1-RepairStation)
  
    - [BaseStationType](#city-comm-input-v1-BaseStationType)
    - [NodeType](#city-comm-input-v1-NodeType)
  
- [city/comm/interaction/aoi/v1/aoi_service.proto](#city_comm_interaction_aoi_v1_aoi_service-proto)
    - [GetBadAoiIDRequest](#city-comm-interaction-aoi-v1-GetBadAoiIDRequest)
    - [GetBadAoiIDResponse](#city-comm-interaction-aoi-v1-GetBadAoiIDResponse)
  
    - [AoiService](#city-comm-interaction-aoi-v1-AoiService)
  
- [city/comm/interaction/demand/v1/demand_service.proto](#city_comm_interaction_demand_v1_demand_service-proto)
    - [SetDemandStatusRequest](#city-comm-interaction-demand-v1-SetDemandStatusRequest)
    - [SetDemandStatusResponse](#city-comm-interaction-demand-v1-SetDemandStatusResponse)
  
    - [DemandService](#city-comm-interaction-demand-v1-DemandService)
  
- [city/comm/interaction/gateway/v1/gateway.proto](#city_comm_interaction_gateway_v1_gateway-proto)
    - [Station](#city-comm-interaction-gateway-v1-Station)
  
    - [Reason](#city-comm-interaction-gateway-v1-Reason)
  
- [city/event/v1/event.proto](#city_event_v1_event-proto)
    - [Event](#city-event-v1-Event)
    - [Events](#city-event-v1-Events)
  
    - [EventType](#city-event-v1-EventType)
  
- [city/comm/interaction/gateway/v1/gateway_service.proto](#city_comm_interaction_gateway_v1_gateway_service-proto)
    - [GetAllStatusRequest](#city-comm-interaction-gateway-v1-GetAllStatusRequest)
    - [GetAllStatusResponse](#city-comm-interaction-gateway-v1-GetAllStatusResponse)
    - [GetEventsRequest](#city-comm-interaction-gateway-v1-GetEventsRequest)
    - [GetEventsResponse](#city-comm-interaction-gateway-v1-GetEventsResponse)
    - [GetRuinInfoRequest](#city-comm-interaction-gateway-v1-GetRuinInfoRequest)
    - [GetRuinInfoResponse](#city-comm-interaction-gateway-v1-GetRuinInfoResponse)
    - [RuinInfo](#city-comm-interaction-gateway-v1-RuinInfo)
    - [SetGatewayPowerStatusRequest](#city-comm-interaction-gateway-v1-SetGatewayPowerStatusRequest)
    - [SetGatewayPowerStatusResponse](#city-comm-interaction-gateway-v1-SetGatewayPowerStatusResponse)
    - [SetGatewayRuinStatusRequest](#city-comm-interaction-gateway-v1-SetGatewayRuinStatusRequest)
    - [SetGatewayRuinStatusResponse](#city-comm-interaction-gateway-v1-SetGatewayRuinStatusResponse)
  
    - [GatewayService](#city-comm-interaction-gateway-v1-GatewayService)
  
- [city/comm/output/v1/output.proto](#city_comm_output_v1_output-proto)
    - [Aoi](#city-comm-output-v1-Aoi)
    - [BaseStation](#city-comm-output-v1-BaseStation)
    - [Node](#city-comm-output-v1-Node)
    - [Person](#city-comm-output-v1-Person)
    - [Signal](#city-comm-output-v1-Signal)
    - [Statistics](#city-comm-output-v1-Statistics)
  
    - [NodeStatus](#city-comm-output-v1-NodeStatus)
    - [PersonConnectStatus](#city-comm-output-v1-PersonConnectStatus)
    - [PersonDemandStatus](#city-comm-output-v1-PersonDemandStatus)
  
- [city/comm/output/v1/output_service.proto](#city_comm_output_v1_output_service-proto)
    - [OutputRequest](#city-comm-output-v1-OutputRequest)
    - [OutputResponse](#city-comm-output-v1-OutputResponse)
  
    - [OutputService](#city-comm-output-v1-OutputService)
  
- [city/config/v1/config.proto](#city_config_v1_config-proto)
    - [MongoPath](#city-config-v1-MongoPath)
    - [OutputTarget](#city-config-v1-OutputTarget)
  
- [city/economy/v1/economy.proto](#city_economy_v1_economy-proto)
    - [Economy](#city-economy-v1-Economy)
    - [Employee](#city-economy-v1-Employee)
    - [Goods](#city-economy-v1-Goods)
    - [Job](#city-economy-v1-Job)
    - [Org](#city-economy-v1-Org)
    - [Person](#city-economy-v1-Person)
  
- [city/economy/v1/org_service.proto](#city_economy_v1_org_service-proto)
    - [GetOrgRequest](#city-economy-v1-GetOrgRequest)
    - [GetOrgResponse](#city-economy-v1-GetOrgResponse)
    - [UpdateOrgEmployeeRequest](#city-economy-v1-UpdateOrgEmployeeRequest)
    - [UpdateOrgEmployeeRequestItem](#city-economy-v1-UpdateOrgEmployeeRequestItem)
    - [UpdateOrgEmployeeResponse](#city-economy-v1-UpdateOrgEmployeeResponse)
    - [UpdateOrgGoodsRequest](#city-economy-v1-UpdateOrgGoodsRequest)
    - [UpdateOrgGoodsRequestItem](#city-economy-v1-UpdateOrgGoodsRequestItem)
    - [UpdateOrgGoodsResponse](#city-economy-v1-UpdateOrgGoodsResponse)
    - [UpdateOrgJobRequest](#city-economy-v1-UpdateOrgJobRequest)
    - [UpdateOrgJobRequestItem](#city-economy-v1-UpdateOrgJobRequestItem)
    - [UpdateOrgJobResponse](#city-economy-v1-UpdateOrgJobResponse)
    - [UpdateOrgMoneyRequest](#city-economy-v1-UpdateOrgMoneyRequest)
    - [UpdateOrgMoneyRequestItem](#city-economy-v1-UpdateOrgMoneyRequestItem)
    - [UpdateOrgMoneyResponse](#city-economy-v1-UpdateOrgMoneyResponse)
  
    - [OrgService](#city-economy-v1-OrgService)
  
- [city/economy/v1/person_service.proto](#city_economy_v1_person_service-proto)
    - [GetPersonRequest](#city-economy-v1-GetPersonRequest)
    - [GetPersonResponse](#city-economy-v1-GetPersonResponse)
    - [UpdatePersonMoneyRequest](#city-economy-v1-UpdatePersonMoneyRequest)
    - [UpdatePersonMoneyRequestItem](#city-economy-v1-UpdatePersonMoneyRequestItem)
    - [UpdatePersonMoneyResponse](#city-economy-v1-UpdatePersonMoneyResponse)
  
    - [PersonService](#city-economy-v1-PersonService)
  
- [city/economy/v2/economy.proto](#city_economy_v2_economy-proto)
    - [Agent](#city-economy-v2-Agent)
    - [EconomyEntities](#city-economy-v2-EconomyEntities)
    - [Org](#city-economy-v2-Org)
  
    - [OrgType](#city-economy-v2-OrgType)
  
- [city/economy/v2/org_service.proto](#city_economy_v2_org_service-proto)
    - [AddAgentRequest](#city-economy-v2-AddAgentRequest)
    - [AddAgentResponse](#city-economy-v2-AddAgentResponse)
    - [AddCitizenRequest](#city-economy-v2-AddCitizenRequest)
    - [AddCitizenResponse](#city-economy-v2-AddCitizenResponse)
    - [AddCurrencyRequest](#city-economy-v2-AddCurrencyRequest)
    - [AddCurrencyResponse](#city-economy-v2-AddCurrencyResponse)
    - [AddEmployeeRequest](#city-economy-v2-AddEmployeeRequest)
    - [AddEmployeeResponse](#city-economy-v2-AddEmployeeResponse)
    - [AddInterestRateRequest](#city-economy-v2-AddInterestRateRequest)
    - [AddInterestRateResponse](#city-economy-v2-AddInterestRateResponse)
    - [AddInventoryRequest](#city-economy-v2-AddInventoryRequest)
    - [AddInventoryResponse](#city-economy-v2-AddInventoryResponse)
    - [AddOrgRequest](#city-economy-v2-AddOrgRequest)
    - [AddOrgResponse](#city-economy-v2-AddOrgResponse)
    - [AddPriceRequest](#city-economy-v2-AddPriceRequest)
    - [AddPriceResponse](#city-economy-v2-AddPriceResponse)
    - [BatchDeltaUpdateRequest](#city-economy-v2-BatchDeltaUpdateRequest)
    - [BatchDeltaUpdateResponse](#city-economy-v2-BatchDeltaUpdateResponse)
    - [BatchGetRequest](#city-economy-v2-BatchGetRequest)
    - [BatchGetResponse](#city-economy-v2-BatchGetResponse)
    - [BatchUpdateRequest](#city-economy-v2-BatchUpdateRequest)
    - [BatchUpdateResponse](#city-economy-v2-BatchUpdateResponse)
    - [CalculateConsumptionRequest](#city-economy-v2-CalculateConsumptionRequest)
    - [CalculateConsumptionResponse](#city-economy-v2-CalculateConsumptionResponse)
    - [CalculateInterestRequest](#city-economy-v2-CalculateInterestRequest)
    - [CalculateInterestResponse](#city-economy-v2-CalculateInterestResponse)
    - [CalculateRealGDPRequest](#city-economy-v2-CalculateRealGDPRequest)
    - [CalculateRealGDPResponse](#city-economy-v2-CalculateRealGDPResponse)
    - [CalculateTaxesDueRequest](#city-economy-v2-CalculateTaxesDueRequest)
    - [CalculateTaxesDueResponse](#city-economy-v2-CalculateTaxesDueResponse)
    - [DeltaUpdateAgentRequest](#city-economy-v2-DeltaUpdateAgentRequest)
    - [DeltaUpdateAgentResponse](#city-economy-v2-DeltaUpdateAgentResponse)
    - [DeltaUpdateOrgRequest](#city-economy-v2-DeltaUpdateOrgRequest)
    - [DeltaUpdateOrgResponse](#city-economy-v2-DeltaUpdateOrgResponse)
    - [GetAgentRequest](#city-economy-v2-GetAgentRequest)
    - [GetAgentResponse](#city-economy-v2-GetAgentResponse)
    - [GetBracketCutoffsRequest](#city-economy-v2-GetBracketCutoffsRequest)
    - [GetBracketCutoffsResponse](#city-economy-v2-GetBracketCutoffsResponse)
    - [GetBracketRatesRequest](#city-economy-v2-GetBracketRatesRequest)
    - [GetBracketRatesResponse](#city-economy-v2-GetBracketRatesResponse)
    - [GetCitizensRequest](#city-economy-v2-GetCitizensRequest)
    - [GetCitizensResponse](#city-economy-v2-GetCitizensResponse)
    - [GetConsumptionCurrencyRequest](#city-economy-v2-GetConsumptionCurrencyRequest)
    - [GetConsumptionCurrencyResponse](#city-economy-v2-GetConsumptionCurrencyResponse)
    - [GetConsumptionPropensityRequest](#city-economy-v2-GetConsumptionPropensityRequest)
    - [GetConsumptionPropensityResponse](#city-economy-v2-GetConsumptionPropensityResponse)
    - [GetCurrencyRequest](#city-economy-v2-GetCurrencyRequest)
    - [GetCurrencyResponse](#city-economy-v2-GetCurrencyResponse)
    - [GetDepressionRequest](#city-economy-v2-GetDepressionRequest)
    - [GetDepressionResponse](#city-economy-v2-GetDepressionResponse)
    - [GetEmployeesRequest](#city-economy-v2-GetEmployeesRequest)
    - [GetEmployeesResponse](#city-economy-v2-GetEmployeesResponse)
    - [GetIncomeCurrencyRequest](#city-economy-v2-GetIncomeCurrencyRequest)
    - [GetIncomeCurrencyResponse](#city-economy-v2-GetIncomeCurrencyResponse)
    - [GetInterestRateRequest](#city-economy-v2-GetInterestRateRequest)
    - [GetInterestRateResponse](#city-economy-v2-GetInterestRateResponse)
    - [GetInventoryRequest](#city-economy-v2-GetInventoryRequest)
    - [GetInventoryResponse](#city-economy-v2-GetInventoryResponse)
    - [GetLocusControlRequest](#city-economy-v2-GetLocusControlRequest)
    - [GetLocusControlResponse](#city-economy-v2-GetLocusControlResponse)
    - [GetNominalGDPRequest](#city-economy-v2-GetNominalGDPRequest)
    - [GetNominalGDPResponse](#city-economy-v2-GetNominalGDPResponse)
    - [GetOrgEntityIdsRequest](#city-economy-v2-GetOrgEntityIdsRequest)
    - [GetOrgEntityIdsResponse](#city-economy-v2-GetOrgEntityIdsResponse)
    - [GetOrgRequest](#city-economy-v2-GetOrgRequest)
    - [GetOrgResponse](#city-economy-v2-GetOrgResponse)
    - [GetPriceRequest](#city-economy-v2-GetPriceRequest)
    - [GetPriceResponse](#city-economy-v2-GetPriceResponse)
    - [GetPricesRequest](#city-economy-v2-GetPricesRequest)
    - [GetPricesResponse](#city-economy-v2-GetPricesResponse)
    - [GetRealGDPRequest](#city-economy-v2-GetRealGDPRequest)
    - [GetRealGDPResponse](#city-economy-v2-GetRealGDPResponse)
    - [GetUnemploymentRequest](#city-economy-v2-GetUnemploymentRequest)
    - [GetUnemploymentResponse](#city-economy-v2-GetUnemploymentResponse)
    - [GetWagesRequest](#city-economy-v2-GetWagesRequest)
    - [GetWagesResponse](#city-economy-v2-GetWagesResponse)
    - [GetWorkingHoursRequest](#city-economy-v2-GetWorkingHoursRequest)
    - [GetWorkingHoursResponse](#city-economy-v2-GetWorkingHoursResponse)
    - [LoadEconomyEntitiesRequest](#city-economy-v2-LoadEconomyEntitiesRequest)
    - [LoadEconomyEntitiesResponse](#city-economy-v2-LoadEconomyEntitiesResponse)
    - [RemoveAgentRequest](#city-economy-v2-RemoveAgentRequest)
    - [RemoveAgentResponse](#city-economy-v2-RemoveAgentResponse)
    - [RemoveCitizenRequest](#city-economy-v2-RemoveCitizenRequest)
    - [RemoveCitizenResponse](#city-economy-v2-RemoveCitizenResponse)
    - [RemoveEmployeeRequest](#city-economy-v2-RemoveEmployeeRequest)
    - [RemoveEmployeeResponse](#city-economy-v2-RemoveEmployeeResponse)
    - [RemoveOrgRequest](#city-economy-v2-RemoveOrgRequest)
    - [RemoveOrgResponse](#city-economy-v2-RemoveOrgResponse)
    - [SaveEconomyEntitiesRequest](#city-economy-v2-SaveEconomyEntitiesRequest)
    - [SaveEconomyEntitiesResponse](#city-economy-v2-SaveEconomyEntitiesResponse)
    - [SetBracketCutoffsRequest](#city-economy-v2-SetBracketCutoffsRequest)
    - [SetBracketCutoffsResponse](#city-economy-v2-SetBracketCutoffsResponse)
    - [SetBracketRatesRequest](#city-economy-v2-SetBracketRatesRequest)
    - [SetBracketRatesResponse](#city-economy-v2-SetBracketRatesResponse)
    - [SetCitizensRequest](#city-economy-v2-SetCitizensRequest)
    - [SetCitizensResponse](#city-economy-v2-SetCitizensResponse)
    - [SetConsumptionCurrencyRequest](#city-economy-v2-SetConsumptionCurrencyRequest)
    - [SetConsumptionCurrencyResponse](#city-economy-v2-SetConsumptionCurrencyResponse)
    - [SetConsumptionPropensityRequest](#city-economy-v2-SetConsumptionPropensityRequest)
    - [SetConsumptionPropensityResponse](#city-economy-v2-SetConsumptionPropensityResponse)
    - [SetCurrencyRequest](#city-economy-v2-SetCurrencyRequest)
    - [SetCurrencyResponse](#city-economy-v2-SetCurrencyResponse)
    - [SetDepressionRequest](#city-economy-v2-SetDepressionRequest)
    - [SetDepressionResponse](#city-economy-v2-SetDepressionResponse)
    - [SetEmployeesRequest](#city-economy-v2-SetEmployeesRequest)
    - [SetEmployeesResponse](#city-economy-v2-SetEmployeesResponse)
    - [SetIncomeCurrencyRequest](#city-economy-v2-SetIncomeCurrencyRequest)
    - [SetIncomeCurrencyResponse](#city-economy-v2-SetIncomeCurrencyResponse)
    - [SetInterestRateRequest](#city-economy-v2-SetInterestRateRequest)
    - [SetInterestRateResponse](#city-economy-v2-SetInterestRateResponse)
    - [SetInventoryRequest](#city-economy-v2-SetInventoryRequest)
    - [SetInventoryResponse](#city-economy-v2-SetInventoryResponse)
    - [SetLocusControlRequest](#city-economy-v2-SetLocusControlRequest)
    - [SetLocusControlResponse](#city-economy-v2-SetLocusControlResponse)
    - [SetNominalGDPRequest](#city-economy-v2-SetNominalGDPRequest)
    - [SetNominalGDPResponse](#city-economy-v2-SetNominalGDPResponse)
    - [SetPriceRequest](#city-economy-v2-SetPriceRequest)
    - [SetPriceResponse](#city-economy-v2-SetPriceResponse)
    - [SetPricesRequest](#city-economy-v2-SetPricesRequest)
    - [SetPricesResponse](#city-economy-v2-SetPricesResponse)
    - [SetRealGDPRequest](#city-economy-v2-SetRealGDPRequest)
    - [SetRealGDPResponse](#city-economy-v2-SetRealGDPResponse)
    - [SetUnemploymentRequest](#city-economy-v2-SetUnemploymentRequest)
    - [SetUnemploymentResponse](#city-economy-v2-SetUnemploymentResponse)
    - [SetWagesRequest](#city-economy-v2-SetWagesRequest)
    - [SetWagesResponse](#city-economy-v2-SetWagesResponse)
    - [SetWorkingHoursRequest](#city-economy-v2-SetWorkingHoursRequest)
    - [SetWorkingHoursResponse](#city-economy-v2-SetWorkingHoursResponse)
    - [UpdateAgentRequest](#city-economy-v2-UpdateAgentRequest)
    - [UpdateAgentResponse](#city-economy-v2-UpdateAgentResponse)
    - [UpdateOrgRequest](#city-economy-v2-UpdateOrgRequest)
    - [UpdateOrgResponse](#city-economy-v2-UpdateOrgResponse)
  
    - [OrgService](#city-economy-v2-OrgService)
  
- [city/elec/input/v1/config.proto](#city_elec_input_v1_config-proto)
    - [Config](#city-elec-input-v1-Config)
    - [Control](#city-elec-input-v1-Control)
    - [ControlStep](#city-elec-input-v1-ControlStep)
    - [Mongo](#city-elec-input-v1-Mongo)
    - [Output](#city-elec-input-v1-Output)
    - [OutputSwitch](#city-elec-input-v1-OutputSwitch)
  
- [city/elec/input/v1/input.proto](#city_elec_input_v1_input-proto)
    - [Facilities](#city-elec-input-v1-Facilities)
    - [Facility](#city-elec-input-v1-Facility)
    - [RepairStation](#city-elec-input-v1-RepairStation)
  
    - [FacilityType](#city-elec-input-v1-FacilityType)
  
- [city/map/v2/light.proto](#city_map_v2_light-proto)
    - [AvailablePhase](#city-map-v2-AvailablePhase)
    - [Phase](#city-map-v2-Phase)
    - [TrafficLight](#city-map-v2-TrafficLight)
  
    - [LightState](#city-map-v2-LightState)
  
- [city/map/v2/map.proto](#city_map_v2_map-proto)
    - [Aoi](#city-map-v2-Aoi)
    - [Header](#city-map-v2-Header)
    - [HeuristicTAZCost](#city-map-v2-HeuristicTAZCost)
    - [Junction](#city-map-v2-Junction)
    - [JunctionLaneGroup](#city-map-v2-JunctionLaneGroup)
    - [Lane](#city-map-v2-Lane)
    - [LaneConnection](#city-map-v2-LaneConnection)
    - [LaneOverlap](#city-map-v2-LaneOverlap)
    - [Map](#city-map-v2-Map)
    - [NextRoadLane](#city-map-v2-NextRoadLane)
    - [NextRoadLanePlan](#city-map-v2-NextRoadLanePlan)
    - [Poi](#city-map-v2-Poi)
    - [Polyline](#city-map-v2-Polyline)
    - [PublicTransportSubline](#city-map-v2-PublicTransportSubline)
    - [Road](#city-map-v2-Road)
    - [RoadIds](#city-map-v2-RoadIds)
    - [SublineSchedules](#city-map-v2-SublineSchedules)
  
    - [AoiType](#city-map-v2-AoiType)
    - [LandUseType](#city-map-v2-LandUseType)
    - [LaneConnectionType](#city-map-v2-LaneConnectionType)
    - [LaneTurn](#city-map-v2-LaneTurn)
    - [LaneType](#city-map-v2-LaneType)
    - [SublineType](#city-map-v2-SublineType)
  
- [city/elec/input/v1/input_service.proto](#city_elec_input_v1_input_service-proto)
    - [InitRequest](#city-elec-input-v1-InitRequest)
    - [InitResponse](#city-elec-input-v1-InitResponse)
  
    - [InputService](#city-elec-input-v1-InputService)
  
- [city/elec/interaction/v1/elec_service.proto](#city_elec_interaction_v1_elec_service-proto)
    - [GetEdgeStatusRequest](#city-elec-interaction-v1-GetEdgeStatusRequest)
    - [GetEdgeStatusResponse](#city-elec-interaction-v1-GetEdgeStatusResponse)
    - [GetNoPowerAOIRequest](#city-elec-interaction-v1-GetNoPowerAOIRequest)
    - [GetNoPowerAOIResponse](#city-elec-interaction-v1-GetNoPowerAOIResponse)
    - [GetPowerRequest](#city-elec-interaction-v1-GetPowerRequest)
    - [GetPowerResponse](#city-elec-interaction-v1-GetPowerResponse)
    - [GetPowerStatusRequest](#city-elec-interaction-v1-GetPowerStatusRequest)
    - [GetPowerStatusResponse](#city-elec-interaction-v1-GetPowerStatusResponse)
    - [GetPowerStatusResponse.PowerStatusEntry](#city-elec-interaction-v1-GetPowerStatusResponse-PowerStatusEntry)
    - [GetRuinInfoRequest](#city-elec-interaction-v1-GetRuinInfoRequest)
    - [GetRuinInfoResponse](#city-elec-interaction-v1-GetRuinInfoResponse)
    - [RuinInfo](#city-elec-interaction-v1-RuinInfo)
    - [SetStatusRequest](#city-elec-interaction-v1-SetStatusRequest)
    - [SetStatusResponse](#city-elec-interaction-v1-SetStatusResponse)
  
    - [ElecService](#city-elec-interaction-v1-ElecService)
  
- [city/elec/output/v1/output.proto](#city_elec_output_v1_output-proto)
    - [Aoi](#city-elec-output-v1-Aoi)
    - [AveragePower](#city-elec-output-v1-AveragePower)
  
- [city/elec/output/v1/output_service.proto](#city_elec_output_v1_output_service-proto)
    - [OutputRequest](#city-elec-output-v1-OutputRequest)
    - [OutputResponse](#city-elec-output-v1-OutputResponse)
  
    - [OutputService](#city-elec-output-v1-OutputService)
  
- [city/event/v1/event_service.proto](#city_event_v1_event_service-proto)
    - [PublishRequest](#city-event-v1-PublishRequest)
    - [PublishResponse](#city-event-v1-PublishResponse)
    - [PullRequest](#city-event-v1-PullRequest)
    - [PullResponse](#city-event-v1-PullResponse)
  
    - [EventService](#city-event-v1-EventService)
  
- [city/event/v2/event.proto](#city_event_v2_event-proto)
    - [Entity](#city-event-v2-Entity)
    - [Event](#city-event-v2-Event)
  
    - [EntityType](#city-event-v2-EntityType)
  
- [city/event/v2/event_service.proto](#city_event_v2_event_service-proto)
    - [GetEventsByTopicRequest](#city-event-v2-GetEventsByTopicRequest)
    - [GetEventsByTopicResponse](#city-event-v2-GetEventsByTopicResponse)
    - [ResolveEventsRequest](#city-event-v2-ResolveEventsRequest)
    - [ResolveEventsResponse](#city-event-v2-ResolveEventsResponse)
  
    - [EventService](#city-event-v2-EventService)
  
- [city/person/v2/motion.proto](#city_person_v2_motion-proto)
    - [PersonMotion](#city-person-v2-PersonMotion)
  
    - [Status](#city-person-v2-Status)
  
- [city/map/v2/aoi_service.proto](#city_map_v2_aoi_service-proto)
    - [AoiState](#city-map-v2-AoiState)
    - [GetAoiRequest](#city-map-v2-GetAoiRequest)
    - [GetAoiResponse](#city-map-v2-GetAoiResponse)
  
    - [AoiService](#city-map-v2-AoiService)
  
- [city/map/v2/lane_state.proto](#city_map_v2_lane_state-proto)
    - [LaneState](#city-map-v2-LaneState)
  
- [city/map/v2/lane_service.proto](#city_map_v2_lane_service-proto)
    - [GetLaneByLongLatBBoxRequest](#city-map-v2-GetLaneByLongLatBBoxRequest)
    - [GetLaneByLongLatBBoxResponse](#city-map-v2-GetLaneByLongLatBBoxResponse)
    - [GetLaneRequest](#city-map-v2-GetLaneRequest)
    - [GetLaneResponse](#city-map-v2-GetLaneResponse)
    - [SetLaneMaxVRequest](#city-map-v2-SetLaneMaxVRequest)
    - [SetLaneMaxVResponse](#city-map-v2-SetLaneMaxVResponse)
    - [SetLaneRestrictionRequest](#city-map-v2-SetLaneRestrictionRequest)
    - [SetLaneRestrictionResponse](#city-map-v2-SetLaneRestrictionResponse)
  
    - [LaneService](#city-map-v2-LaneService)
  
- [city/map/v2/road_service.proto](#city_map_v2_road_service-proto)
    - [GetEventsRequest](#city-map-v2-GetEventsRequest)
    - [GetEventsResponse](#city-map-v2-GetEventsResponse)
    - [GetRoadRequest](#city-map-v2-GetRoadRequest)
    - [GetRoadResponse](#city-map-v2-GetRoadResponse)
    - [GetRuinInfoRequest](#city-map-v2-GetRuinInfoRequest)
    - [GetRuinInfoResponse](#city-map-v2-GetRuinInfoResponse)
    - [RoadState](#city-map-v2-RoadState)
    - [RuinInfo](#city-map-v2-RuinInfo)
  
    - [InterruptionReason](#city-map-v2-InterruptionReason)
    - [RoadLevel](#city-map-v2-RoadLevel)
  
    - [RoadService](#city-map-v2-RoadService)
  
- [city/map/v2/traffic_light_service.proto](#city_map_v2_traffic_light_service-proto)
    - [GetTrafficLightRequest](#city-map-v2-GetTrafficLightRequest)
    - [GetTrafficLightResponse](#city-map-v2-GetTrafficLightResponse)
    - [SetTrafficLightPhaseRequest](#city-map-v2-SetTrafficLightPhaseRequest)
    - [SetTrafficLightPhaseResponse](#city-map-v2-SetTrafficLightPhaseResponse)
    - [SetTrafficLightRequest](#city-map-v2-SetTrafficLightRequest)
    - [SetTrafficLightResponse](#city-map-v2-SetTrafficLightResponse)
    - [SetTrafficLightStatusRequest](#city-map-v2-SetTrafficLightStatusRequest)
    - [SetTrafficLightStatusResponse](#city-map-v2-SetTrafficLightStatusResponse)
  
    - [TrafficLightService](#city-map-v2-TrafficLightService)
  
- [city/pause/v1/pause_service.proto](#city_pause_v1_pause_service-proto)
    - [PauseRequest](#city-pause-v1-PauseRequest)
    - [PauseResponse](#city-pause-v1-PauseResponse)
    - [ResumeRequest](#city-pause-v1-ResumeRequest)
    - [ResumeResponse](#city-pause-v1-ResumeResponse)
  
    - [PauseService](#city-pause-v1-PauseService)
  
- [city/person/v1/motion.proto](#city_person_v1_motion-proto)
    - [PersonMotion](#city-person-v1-PersonMotion)
  
    - [Status](#city-person-v1-Status)
  
- [city/routing/v2/routing.proto](#city_routing_v2_routing-proto)
    - [BusJourneyBody](#city-routing-v2-BusJourneyBody)
    - [DrivingJourneyBody](#city-routing-v2-DrivingJourneyBody)
    - [Journey](#city-routing-v2-Journey)
    - [RoadStatus](#city-routing-v2-RoadStatus)
    - [RoadStatuses](#city-routing-v2-RoadStatuses)
    - [TransferSegment](#city-routing-v2-TransferSegment)
    - [WalkingJourneyBody](#city-routing-v2-WalkingJourneyBody)
    - [WalkingRouteSegment](#city-routing-v2-WalkingRouteSegment)
  
    - [JourneyType](#city-routing-v2-JourneyType)
    - [MovingDirection](#city-routing-v2-MovingDirection)
    - [RouteType](#city-routing-v2-RouteType)
  
- [city/trip/v2/trip.proto](#city_trip_v2_trip-proto)
    - [Schedule](#city-trip-v2-Schedule)
    - [Trip](#city-trip-v2-Trip)
    - [TripStop](#city-trip-v2-TripStop)
  
    - [TripMode](#city-trip-v2-TripMode)
  
- [city/person/v1/person.proto](#city_person_v1_person-proto)
    - [BikeAttribute](#city-person-v1-BikeAttribute)
    - [BusAttribute](#city-person-v1-BusAttribute)
    - [PedestrianAttribute](#city-person-v1-PedestrianAttribute)
    - [Person](#city-person-v1-Person)
    - [Person.LabelsEntry](#city-person-v1-Person-LabelsEntry)
    - [PersonAttribute](#city-person-v1-PersonAttribute)
    - [PersonProfile](#city-person-v1-PersonProfile)
    - [Persons](#city-person-v1-Persons)
    - [VehicleAttribute](#city-person-v1-VehicleAttribute)
  
    - [BusType](#city-person-v1-BusType)
    - [Consumption](#city-person-v1-Consumption)
    - [Education](#city-person-v1-Education)
    - [Gender](#city-person-v1-Gender)
  
- [city/person/v1/person_runtime.proto](#city_person_v1_person_runtime-proto)
    - [PersonRuntime](#city-person-v1-PersonRuntime)
  
- [city/person/v1/vehicle.proto](#city_person_v1_vehicle-proto)
    - [LC](#city-person-v1-LC)
    - [ObservedLane](#city-person-v1-ObservedLane)
    - [ObservedVehicle](#city-person-v1-ObservedVehicle)
    - [VehicleAction](#city-person-v1-VehicleAction)
    - [VehicleEnv](#city-person-v1-VehicleEnv)
    - [VehicleRuntime](#city-person-v1-VehicleRuntime)
  
    - [LightState](#city-person-v1-LightState)
    - [VehicleRelation](#city-person-v1-VehicleRelation)
  
- [city/person/v1/person_service.proto](#city_person_v1_person_service-proto)
    - [AddPersonRequest](#city-person-v1-AddPersonRequest)
    - [AddPersonResponse](#city-person-v1-AddPersonResponse)
    - [FetchControlledVehicleEnvsRequest](#city-person-v1-FetchControlledVehicleEnvsRequest)
    - [FetchControlledVehicleEnvsResponse](#city-person-v1-FetchControlledVehicleEnvsResponse)
    - [GetAllVehiclesRequest](#city-person-v1-GetAllVehiclesRequest)
    - [GetAllVehiclesResponse](#city-person-v1-GetAllVehiclesResponse)
    - [GetPersonByLongLatBBoxRequest](#city-person-v1-GetPersonByLongLatBBoxRequest)
    - [GetPersonByLongLatBBoxResponse](#city-person-v1-GetPersonByLongLatBBoxResponse)
    - [GetPersonRequest](#city-person-v1-GetPersonRequest)
    - [GetPersonResponse](#city-person-v1-GetPersonResponse)
    - [GetPersonsRequest](#city-person-v1-GetPersonsRequest)
    - [GetPersonsResponse](#city-person-v1-GetPersonsResponse)
    - [ResetPersonPositionRequest](#city-person-v1-ResetPersonPositionRequest)
    - [ResetPersonPositionResponse](#city-person-v1-ResetPersonPositionResponse)
    - [SetControlledVehicleActionsRequest](#city-person-v1-SetControlledVehicleActionsRequest)
    - [SetControlledVehicleActionsResponse](#city-person-v1-SetControlledVehicleActionsResponse)
    - [SetControlledVehicleIDsRequest](#city-person-v1-SetControlledVehicleIDsRequest)
    - [SetControlledVehicleIDsResponse](#city-person-v1-SetControlledVehicleIDsResponse)
    - [SetScheduleRequest](#city-person-v1-SetScheduleRequest)
    - [SetScheduleResponse](#city-person-v1-SetScheduleResponse)
  
    - [PersonService](#city-person-v1-PersonService)
  
- [city/person/v2/carbon.proto](#city_person_v2_carbon-proto)
    - [VehicleCarbon](#city-person-v2-VehicleCarbon)
  
- [city/person/v2/person.proto](#city_person_v2_person-proto)
    - [BikeAttribute](#city-person-v2-BikeAttribute)
    - [BusAttribute](#city-person-v2-BusAttribute)
    - [EmissionAttribute](#city-person-v2-EmissionAttribute)
    - [PedestrianAttribute](#city-person-v2-PedestrianAttribute)
    - [Person](#city-person-v2-Person)
    - [Person.LabelsEntry](#city-person-v2-Person-LabelsEntry)
    - [PersonAttribute](#city-person-v2-PersonAttribute)
    - [PersonProfile](#city-person-v2-PersonProfile)
    - [Persons](#city-person-v2-Persons)
    - [VehicleAttribute](#city-person-v2-VehicleAttribute)
    - [VehicleEngineEfficiency](#city-person-v2-VehicleEngineEfficiency)
  
    - [BusType](#city-person-v2-BusType)
    - [Consumption](#city-person-v2-Consumption)
    - [Education](#city-person-v2-Education)
    - [Gender](#city-person-v2-Gender)
    - [PersonType](#city-person-v2-PersonType)
    - [VehicleEngineType](#city-person-v2-VehicleEngineType)
  
- [city/person/v2/person_runtime.proto](#city_person_v2_person_runtime-proto)
    - [PersonRuntime](#city-person-v2-PersonRuntime)
  
- [city/person/v2/vehicle.proto](#city_person_v2_vehicle-proto)
    - [LC](#city-person-v2-LC)
    - [ObservedLane](#city-person-v2-ObservedLane)
    - [ObservedVehicle](#city-person-v2-ObservedVehicle)
    - [VehicleAction](#city-person-v2-VehicleAction)
    - [VehicleEnv](#city-person-v2-VehicleEnv)
    - [VehicleRouteAction](#city-person-v2-VehicleRouteAction)
    - [VehicleRuntime](#city-person-v2-VehicleRuntime)
  
    - [LightState](#city-person-v2-LightState)
    - [VehicleRelation](#city-person-v2-VehicleRelation)
  
- [city/person/v2/person_service.proto](#city_person_v2_person_service-proto)
    - [AddPersonRequest](#city-person-v2-AddPersonRequest)
    - [AddPersonResponse](#city-person-v2-AddPersonResponse)
    - [FetchControlledVehicleEnvsRequest](#city-person-v2-FetchControlledVehicleEnvsRequest)
    - [FetchControlledVehicleEnvsResponse](#city-person-v2-FetchControlledVehicleEnvsResponse)
    - [GetAllVehiclesRequest](#city-person-v2-GetAllVehiclesRequest)
    - [GetAllVehiclesResponse](#city-person-v2-GetAllVehiclesResponse)
    - [GetPersonByLongLatBBoxRequest](#city-person-v2-GetPersonByLongLatBBoxRequest)
    - [GetPersonByLongLatBBoxResponse](#city-person-v2-GetPersonByLongLatBBoxResponse)
    - [GetPersonRequest](#city-person-v2-GetPersonRequest)
    - [GetPersonResponse](#city-person-v2-GetPersonResponse)
    - [GetPersonsRequest](#city-person-v2-GetPersonsRequest)
    - [GetPersonsResponse](#city-person-v2-GetPersonsResponse)
    - [ResetPersonPositionRequest](#city-person-v2-ResetPersonPositionRequest)
    - [ResetPersonPositionResponse](#city-person-v2-ResetPersonPositionResponse)
    - [SetControlledVehicleActionsRequest](#city-person-v2-SetControlledVehicleActionsRequest)
    - [SetControlledVehicleActionsResponse](#city-person-v2-SetControlledVehicleActionsResponse)
    - [SetControlledVehicleIDsRequest](#city-person-v2-SetControlledVehicleIDsRequest)
    - [SetControlledVehicleIDsResponse](#city-person-v2-SetControlledVehicleIDsResponse)
    - [SetScheduleRequest](#city-person-v2-SetScheduleRequest)
    - [SetScheduleResponse](#city-person-v2-SetScheduleResponse)
  
    - [PersonService](#city-person-v2-PersonService)
  
- [city/ping/v1/ping_service.proto](#city_ping_v1_ping_service-proto)
    - [PingRequest](#city-ping-v1-PingRequest)
    - [PingResponse](#city-ping-v1-PingResponse)
  
    - [PingService](#city-ping-v1-PingService)
  
- [city/routing/v2/cost.proto](#city_routing_v2_cost-proto)
    - [Cost](#city-routing-v2-Cost)
  
- [city/routing/v2/routing_service.proto](#city_routing_v2_routing_service-proto)
    - [GetDrivingCostsRequest](#city-routing-v2-GetDrivingCostsRequest)
    - [GetDrivingCostsResponse](#city-routing-v2-GetDrivingCostsResponse)
    - [GetRouteRequest](#city-routing-v2-GetRouteRequest)
    - [GetRouteResponse](#city-routing-v2-GetRouteResponse)
    - [SetDrivingCostsRequest](#city-routing-v2-SetDrivingCostsRequest)
    - [SetDrivingCostsResponse](#city-routing-v2-SetDrivingCostsResponse)
  
    - [RoutingService](#city-routing-v2-RoutingService)
  
- [city/social/v1/message.proto](#city_social_v1_message-proto)
    - [Message](#city-social-v1-Message)
  
- [city/social/v1/social_service.proto](#city_social_v1_social_service-proto)
    - [ReceiveRequest](#city-social-v1-ReceiveRequest)
    - [ReceiveResponse](#city-social-v1-ReceiveResponse)
    - [SendRequest](#city-social-v1-SendRequest)
    - [SendResponse](#city-social-v1-SendResponse)
  
    - [SocialService](#city-social-v1-SocialService)
  
- [city/streetview/v1/streetview.proto](#city_streetview_v1_streetview-proto)
    - [StreetView](#city-streetview-v1-StreetView)
    - [StreetViewImage](#city-streetview-v1-StreetViewImage)
    - [StreetViews](#city-streetview-v1-StreetViews)
  
- [city/sync/v2/sync_service.proto](#city_sync_v2_sync_service-proto)
    - [EnterStepSyncRequest](#city-sync-v2-EnterStepSyncRequest)
    - [EnterStepSyncResponse](#city-sync-v2-EnterStepSyncResponse)
    - [ExitStepSyncRequest](#city-sync-v2-ExitStepSyncRequest)
    - [ExitStepSyncResponse](#city-sync-v2-ExitStepSyncResponse)
    - [GetURLRequest](#city-sync-v2-GetURLRequest)
    - [GetURLResponse](#city-sync-v2-GetURLResponse)
    - [SetURLRequest](#city-sync-v2-SetURLRequest)
    - [SetURLResponse](#city-sync-v2-SetURLResponse)
  
    - [SyncService](#city-sync-v2-SyncService)
  
- [city/water/input/v1/config.proto](#city_water_input_v1_config-proto)
    - [Config](#city-water-input-v1-Config)
    - [Control](#city-water-input-v1-Control)
    - [ControlStep](#city-water-input-v1-ControlStep)
    - [Mongo](#city-water-input-v1-Mongo)
    - [Output](#city-water-input-v1-Output)
    - [OutputSwitch](#city-water-input-v1-OutputSwitch)
  
- [city/water/input/v1/water.proto](#city_water_input_v1_water-proto)
    - [Rain](#city-water-input-v1-Rain)
    - [RainPeriod](#city-water-input-v1-RainPeriod)
  
- [city/water/input/v1/input_service.proto](#city_water_input_v1_input_service-proto)
    - [InitRequest](#city-water-input-v1-InitRequest)
    - [InitResponse](#city-water-input-v1-InitResponse)
  
    - [InputService](#city-water-input-v1-InputService)
  
- [city/water/interaction/v1/water_service.proto](#city_water_interaction_v1_water_service-proto)
    - [GetNoWaterAOIRequest](#city-water-interaction-v1-GetNoWaterAOIRequest)
    - [GetNoWaterAOIResponse](#city-water-interaction-v1-GetNoWaterAOIResponse)
    - [GetPumpStatusRequest](#city-water-interaction-v1-GetPumpStatusRequest)
    - [GetPumpStatusResponse](#city-water-interaction-v1-GetPumpStatusResponse)
    - [GetPumpStatusResponse.PumpStatusEntry](#city-water-interaction-v1-GetPumpStatusResponse-PumpStatusEntry)
    - [GetRuinInfoRequest](#city-water-interaction-v1-GetRuinInfoRequest)
    - [GetRuinInfoResponse](#city-water-interaction-v1-GetRuinInfoResponse)
    - [RuinInfo](#city-water-interaction-v1-RuinInfo)
    - [SetPumpNetworkStatusRequest](#city-water-interaction-v1-SetPumpNetworkStatusRequest)
    - [SetPumpNetworkStatusResponse](#city-water-interaction-v1-SetPumpNetworkStatusResponse)
    - [SetPumpPowerStatusRequest](#city-water-interaction-v1-SetPumpPowerStatusRequest)
    - [SetPumpPowerStatusResponse](#city-water-interaction-v1-SetPumpPowerStatusResponse)
    - [SetPumpStatusRequest](#city-water-interaction-v1-SetPumpStatusRequest)
    - [SetPumpStatusResponse](#city-water-interaction-v1-SetPumpStatusResponse)
  
    - [WaterFacilityType](#city-water-interaction-v1-WaterFacilityType)
  
    - [WaterService](#city-water-interaction-v1-WaterService)
  
- [city/water/output/v1/output.proto](#city_water_output_v1_output-proto)
    - [Aoi](#city-water-output-v1-Aoi)
    - [DetailedRoad](#city-water-output-v1-DetailedRoad)
    - [DrainageBasicInfo](#city-water-output-v1-DrainageBasicInfo)
    - [DrainageMetrics](#city-water-output-v1-DrainageMetrics)
    - [FailureStatistics](#city-water-output-v1-FailureStatistics)
    - [Link](#city-water-output-v1-Link)
    - [Node](#city-water-output-v1-Node)
    - [Road](#city-water-output-v1-Road)
    - [RoadFlood](#city-water-output-v1-RoadFlood)
    - [SupplyBasicInfo](#city-water-output-v1-SupplyBasicInfo)
    - [SupplyDemandStatistics](#city-water-output-v1-SupplyDemandStatistics)
    - [SupplyMetrics](#city-water-output-v1-SupplyMetrics)
  
    - [LinkType](#city-water-output-v1-LinkType)
  
- [city/water/output/v1/output_service.proto](#city_water_output_v1_output_service-proto)
    - [OutputRequest](#city-water-output-v1-OutputRequest)
    - [OutputResponse](#city-water-output-v1-OutputResponse)
  
    - [OutputService](#city-water-output-v1-OutputService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="city_clock_v1_clock_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/clock/v1/clock_service.proto



<a name="city-clock-v1-NowRequest"></a>

### NowRequest
获取当前的模拟时间请求
request of getting current simulation clock






<a name="city-clock-v1-NowResponse"></a>

### NowResponse
获取当前的模拟时间响应
response of getting current simulation clock


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| day | [int32](#int32) | optional | 当前模拟的天数 current simulation day |
| t | [double](#double) |  | 当前的模拟时间，单位为秒 current simulation clock, in seconds |





 

 

 


<a name="city-clock-v1-ClockService"></a>

### ClockService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Now | [NowRequest](#city-clock-v1-NowRequest) | [NowResponse](#city-clock-v1-NowResponse) | 获取当前的模拟时间 get current simulation clock |

 



<a name="city_geo_v2_geo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/geo/v2/geo.proto



<a name="city-geo-v2-AoiPosition"></a>

### AoiPosition
地图坐标（AOI）
Map coordinates (AOI)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aoi_id | [int32](#int32) |  | AOI ID |
| poi_id | [int32](#int32) | optional | POI ID，需要是aoi_id的子poi，否则该值无效 POI ID, needs to be a sub-poi of aoi_id, otherwise the value is invalid |






<a name="city-geo-v2-LanePosition"></a>

### LanePosition
地图坐标（车道&#43;距离s）
Map coordinates (lane ID &#43; distance s)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lane_id | [int32](#int32) |  | 车道id Lane ID |
| s | [double](#double) |  | s是车道上的点到车道起点的距离 s is the distance from the point on the lane to the starting point of the lane |






<a name="city-geo-v2-LongLatBBox"></a>

### LongLatBBox
经纬度矩形区域
latitude and longitude rectangular area


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min_longitude | [double](#double) |  | 最小经度 minimum longitude |
| min_latitude | [double](#double) |  | 最小纬度 minimum latitude |
| max_longitude | [double](#double) |  | 最大经度 maximu longitude |
| max_latitude | [double](#double) |  | 最大纬度 minimum longitude |






<a name="city-geo-v2-LongLatPosition"></a>

### LongLatPosition
WGS84经纬度坐标
WGS84 longitute and latitude coordinates


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| longitude | [double](#double) |  | 经度 longitude |
| latitude | [double](#double) |  | 纬度 latitude |
| z | [double](#double) | optional | 高程（单位：米） elevation (unit: meters) |






<a name="city-geo-v2-Position"></a>

### Position
坐标，如果多种坐标同时存在，两两之间必须满足映射关系，同时逻辑坐标是必须提供的
Coordinates, if multiple coordinates exist at the same time, the mapping relationship between them must be satisfied, and logical coordinates must be provided.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lane_position | [LanePosition](#city-geo-v2-LanePosition) | optional | 地图坐标AOI（必须提供其中之一） Map coordinates AOI (one of these must be provided) |
| aoi_position | [AoiPosition](#city-geo-v2-AoiPosition) | optional | 地图坐标Lane&#43;S（必须提供其中之一） Map coordinates Lane&#43;S (one of these must be provided) |
| longlat_position | [LongLatPosition](#city-geo-v2-LongLatPosition) | optional | WGS84经纬度坐标 WGS84 longitute and latitude coordinates |
| xy_position | [XYPosition](#city-geo-v2-XYPosition) | optional | XY坐标 XY coordinates |






<a name="city-geo-v2-XYPosition"></a>

### XYPosition
XY坐标
XY coordinates


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| x | [double](#double) |  | x坐标，单位米，对应经度 x coordinate, in meters, corresponding to longitude |
| y | [double](#double) |  | y坐标，单位米，对应纬度 y coordinate, in meters, corresponding to latitude |
| z | [double](#double) | optional | z坐标，单位米，对应高程 z coordinate, in meters, corresponding to elevation |





 

 

 

 



<a name="city_comm_input_v1_comm-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/comm/input/v1/comm.proto



<a name="city-comm-input-v1-CommDemand"></a>

### CommDemand
终端通信需求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| demands | [double](#double) | repeated |  |






<a name="city-comm-input-v1-CommDemands"></a>

### CommDemands



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| comm_demands | [CommDemand](#city-comm-input-v1-CommDemand) | repeated |  |






<a name="city-comm-input-v1-Node"></a>

### Node



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| type | [NodeType](#city-comm-input-v1-NodeType) |  |  |
| parent_id | [int32](#int32) |  | 父节点 |
| children_ids | [int32](#int32) | repeated | 子节点 |
| position | [city.geo.v2.Position](#city-geo-v2-Position) | optional | 节点经纬度位置 |
| aoi_id | [int32](#int32) | optional | 节点所在aoi |
| freq_range_id | [int32](#int32) | optional | 基站频段id |
| base_station_type | [BaseStationType](#city-comm-input-v1-BaseStationType) | optional | 室内外基站类型 |






<a name="city-comm-input-v1-Nodes"></a>

### Nodes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [Node](#city-comm-input-v1-Node) | repeated |  |
| repair_stations | [RepairStation](#city-comm-input-v1-RepairStation) | repeated |  |
| pumps | [Pump](#city-comm-input-v1-Pump) | repeated |  |






<a name="city-comm-input-v1-Pump"></a>

### Pump
水泵


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| position | [city.geo.v2.Position](#city-geo-v2-Position) |  |  |






<a name="city-comm-input-v1-RepairStation"></a>

### RepairStation
抢修站


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| aoi_id | [int32](#int32) |  |  |
| position | [city.geo.v2.Position](#city-geo-v2-Position) |  |  |





 


<a name="city-comm-input-v1-BaseStationType"></a>

### BaseStationType


| Name | Number | Description |
| ---- | ------ | ----------- |
| BASE_STATION_TYPE_UNSPECIFIED | 0 |  |
| BASE_STATION_TYPE_INDOOR | 1 |  |
| BASE_STATION_TYPE_OUTDOOR | 2 |  |



<a name="city-comm-input-v1-NodeType"></a>

### NodeType
本文件描述通信部分拓扑结构
三种节点类型

| Name | Number | Description |
| ---- | ------ | ----------- |
| NODE_TYPE_UNSPECIFIED | 0 |  |
| NODE_TYPE_INTERNET | 1 |  |
| NODE_TYPE_GATEWAY | 2 |  |
| NODE_TYPE_BASE_STATION | 3 |  |


 

 

 



<a name="city_comm_interaction_aoi_v1_aoi_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/comm/interaction/aoi/v1/aoi_service.proto



<a name="city-comm-interaction-aoi-v1-GetBadAoiIDRequest"></a>

### GetBadAoiIDRequest







<a name="city-comm-interaction-aoi-v1-GetBadAoiIDResponse"></a>

### GetBadAoiIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [int32](#int32) | repeated |  |





 

 

 


<a name="city-comm-interaction-aoi-v1-AoiService"></a>

### AoiService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetBadAoiID | [GetBadAoiIDRequest](#city-comm-interaction-aoi-v1-GetBadAoiIDRequest) | [GetBadAoiIDResponse](#city-comm-interaction-aoi-v1-GetBadAoiIDResponse) |  |

 



<a name="city_comm_interaction_demand_v1_demand_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/comm/interaction/demand/v1/demand_service.proto



<a name="city-comm-interaction-demand-v1-SetDemandStatusRequest"></a>

### SetDemandStatusRequest
设置用户通信需求激增
用户通信需求激增公式：
result=demand*multiple_times*exp(-power_times*(current_time-start_time))
demand: 用户正常通信需求
current_time: 当前时间, start_time: 开始激增时间


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| multiple_times | [double](#double) |  |  |
| power_times | [double](#double) |  |  |






<a name="city-comm-interaction-demand-v1-SetDemandStatusResponse"></a>

### SetDemandStatusResponse






 

 

 


<a name="city-comm-interaction-demand-v1-DemandService"></a>

### DemandService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetDemandStatus | [SetDemandStatusRequest](#city-comm-interaction-demand-v1-SetDemandStatusRequest) | [SetDemandStatusResponse](#city-comm-interaction-demand-v1-SetDemandStatusResponse) |  |

 



<a name="city_comm_interaction_gateway_v1_gateway-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/comm/interaction/gateway/v1/gateway.proto



<a name="city-comm-interaction-gateway-v1-Station"></a>

### Station



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| status | [bool](#bool) |  |  |
| reason | [Reason](#city-comm-interaction-gateway-v1-Reason) |  |  |





 


<a name="city-comm-interaction-gateway-v1-Reason"></a>

### Reason


| Name | Number | Description |
| ---- | ------ | ----------- |
| REASON_UNSPECIFIED | 0 |  |
| REASON_RUIN | 1 |  |
| REASON_CASCADE | 2 |  |


 

 

 



<a name="city_event_v1_event-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/event/v1/event.proto



<a name="city-event-v1-Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EventType](#city-event-v1-EventType) |  |  |
| level | [int32](#int32) |  |  |
| step | [int32](#int32) |  |  |






<a name="city-event-v1-Events"></a>

### Events



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [Event](#city-event-v1-Event) | repeated |  |





 


<a name="city-event-v1-EventType"></a>

### EventType


| Name | Number | Description |
| ---- | ------ | ----------- |
| EVENT_TYPE_UNSPECIFIED | 0 |  |
| EVENT_TYPE_HEAVY_RAIN | 1 | 特大暴雨 |
| EVENT_TYPE_MILITARY_STRIKE | 2 | 军事打击 |
| EVENT_TYPE_TRAFFIC_CONGESTION | 3 | 交通拥堵 |
| EVENT_TYPE_TRAFFIC_LANE_RESTRICTION | 4 | 道路限行 |
| EVENT_TYPE_TRAFFIC_BAD_LIGHT | 5 | 信控失效 |
| EVENT_TYPE_ELEC_RUINED_TRANSFORMER | 6 | 变压器损坏（被摧毁） |
| EVENT_TYPE_WATER_RUINED_PUMP | 7 | 水泵损坏（被摧毁） |
| EVENT_TYPE_WATER_STOPPED_PUMP | 8 | 水泵停电（变压器停电影响） |
| EVENT_TYPE_COMM_RUINED_BASE_STATION | 9 | 基站损坏 |
| EVENT_TYPE_COMM_STOPPED_BASE_STATION | 10 | 基站停电 |
| EVENT_TYPE_COMM_OVERLOAD_BASE_STATION | 11 | 基站过载 |


 

 

 



<a name="city_comm_interaction_gateway_v1_gateway_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/comm/interaction/gateway/v1/gateway_service.proto



<a name="city-comm-interaction-gateway-v1-GetAllStatusRequest"></a>

### GetAllStatusRequest







<a name="city-comm-interaction-gateway-v1-GetAllStatusResponse"></a>

### GetAllStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stations | [Station](#city-comm-interaction-gateway-v1-Station) | repeated |  |






<a name="city-comm-interaction-gateway-v1-GetEventsRequest"></a>

### GetEventsRequest







<a name="city-comm-interaction-gateway-v1-GetEventsResponse"></a>

### GetEventsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [city.event.v1.Events](#city-event-v1-Events) |  |  |






<a name="city-comm-interaction-gateway-v1-GetRuinInfoRequest"></a>

### GetRuinInfoRequest







<a name="city-comm-interaction-gateway-v1-GetRuinInfoResponse"></a>

### GetRuinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| one | [RuinInfo](#city-comm-interaction-gateway-v1-RuinInfo) |  | 三级级损伤信息 |
| two | [RuinInfo](#city-comm-interaction-gateway-v1-RuinInfo) |  |  |
| three | [RuinInfo](#city-comm-interaction-gateway-v1-RuinInfo) |  |  |






<a name="city-comm-interaction-gateway-v1-RuinInfo"></a>

### RuinInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num | [int32](#int32) |  | 损坏数量 |
| ratio | [double](#double) |  | 损坏占比 |






<a name="city-comm-interaction-gateway-v1-SetGatewayPowerStatusRequest"></a>

### SetGatewayPowerStatusRequest
断电或恢复状态，True为修复，False为断电


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| status | [bool](#bool) |  |  |






<a name="city-comm-interaction-gateway-v1-SetGatewayPowerStatusResponse"></a>

### SetGatewayPowerStatusResponse







<a name="city-comm-interaction-gateway-v1-SetGatewayRuinStatusRequest"></a>

### SetGatewayRuinStatusRequest
摧毁或恢复状态，True为修复，False为摧毁


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| status | [bool](#bool) |  |  |






<a name="city-comm-interaction-gateway-v1-SetGatewayRuinStatusResponse"></a>

### SetGatewayRuinStatusResponse






 

 

 


<a name="city-comm-interaction-gateway-v1-GatewayService"></a>

### GatewayService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetGatewayPowerStatus | [SetGatewayPowerStatusRequest](#city-comm-interaction-gateway-v1-SetGatewayPowerStatusRequest) | [SetGatewayPowerStatusResponse](#city-comm-interaction-gateway-v1-SetGatewayPowerStatusResponse) |  |
| SetGatewayRuinStatus | [SetGatewayRuinStatusRequest](#city-comm-interaction-gateway-v1-SetGatewayRuinStatusRequest) | [SetGatewayRuinStatusResponse](#city-comm-interaction-gateway-v1-SetGatewayRuinStatusResponse) |  |
| GetAllStatus | [GetAllStatusRequest](#city-comm-interaction-gateway-v1-GetAllStatusRequest) | [GetAllStatusResponse](#city-comm-interaction-gateway-v1-GetAllStatusResponse) |  |
| GetRuinInfo | [GetRuinInfoRequest](#city-comm-interaction-gateway-v1-GetRuinInfoRequest) | [GetRuinInfoResponse](#city-comm-interaction-gateway-v1-GetRuinInfoResponse) |  |
| GetEvents | [GetEventsRequest](#city-comm-interaction-gateway-v1-GetEventsRequest) | [GetEventsResponse](#city-comm-interaction-gateway-v1-GetEventsResponse) |  |

 



<a name="city_comm_output_v1_output-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/comm/output/v1/output.proto



<a name="city-comm-output-v1-Aoi"></a>

### Aoi



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| demand_flow | [double](#double) |  | 通信需求总量（单位：Bytes） |
| actual_flow | [double](#double) |  | 通信真实总量（单位：Bytes） |
| outage_num | [int32](#int32) |  | 通信中断人数 |
| satisfied_num | [int32](#int32) |  | 通信满足人数 |
| unsatisfied_num | [int32](#int32) |  | 通信不满足人数 |
| active_user_num | [int32](#int32) |  | 有通信需求的人数 |






<a name="city-comm-output-v1-BaseStation"></a>

### BaseStation
基站状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| demand_flow | [double](#double) |  | 基站需求流量（单位：Bytes） |
| actual_flow | [double](#double) |  | 基站真实流量（单位：Bytes） |
| num_agents | [int32](#int32) |  | 基站用户数 |
| overload | [bool](#bool) |  | 基站是否过载 |
| unsatisfied_num | [int32](#int32) |  | 不满足通信需求用户数 |
| satisfied_num | [int32](#int32) |  | 满足通信需求用户数 |
| outage_num | [int32](#int32) |  | 通信中断用户数 |
| active_num | [int32](#int32) |  | 有通信需求人数 |
| transmit_power | [double](#double) |  | 基站发射功率 |






<a name="city-comm-output-v1-Node"></a>

### Node
设备状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| status | [NodeStatus](#city-comm-output-v1-NodeStatus) |  |  |
| battery_remaining_time | [double](#double) | optional | 电池剩余可用时间（单位：秒）（仅当gateway状态为“电池供电”时启用） |






<a name="city-comm-output-v1-Person"></a>

### Person
人（可见的，交通模拟或室内模拟正在计算位置变动的）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| demand_rate | [double](#double) |  | 通信需求速率（单位：Bytes/s） |
| actual_rate | [double](#double) |  | 通信真实速率（单位：Bytes/s） |
| connect_status | [PersonConnectStatus](#city-comm-output-v1-PersonConnectStatus) |  |  |
| demand_status | [PersonDemandStatus](#city-comm-output-v1-PersonDemandStatus) |  |  |
| strength | [double](#double) |  | 信号强度（单位：dBm） |
| base_station_id | [int32](#int32) |  | 连接基站 |
| freq_range_ids | [int32](#int32) | repeated | 信道分配 |
| received_power | [double](#double) |  | 接收发射功率（单位：dBm） |






<a name="city-comm-output-v1-Signal"></a>

### Signal
信号情况


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num_rows | [int32](#int32) |  |  |
| num_columns | [int32](#int32) |  |  |
| strength | [double](#double) | repeated | 信号强度（单位：dBm） |
| base_station_id | [int32](#int32) | repeated | 基站ID |
| freq_range_id | [int32](#int32) | repeated | 基站频段 |






<a name="city-comm-output-v1-Statistics"></a>

### Statistics



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num_satisfied_agents | [int32](#int32) |  | 满足通信需求用户数 |
| num_unsatisfied_agents | [int32](#int32) |  | 未能满足通信需求的用户数 |
| num_outage_agents | [int32](#int32) |  | 通信中断的用户数 |
| num_active_agents | [int32](#int32) |  | 有通信需求的用户数 |
| demand_flow | [double](#double) |  | 需求总流量（单位：Bytes） |
| actual_flow | [double](#double) |  | 真实总流量（单位：Bytes） |
| num_base_station | [int32](#int32) |  | 总基站数 |
| num_ok_base_station | [int32](#int32) |  | 正常基站数 |
| num_ruined_base_station | [int32](#int32) |  | 损坏基站数 |
| num_stopped_base_station | [int32](#int32) |  | 断电基站数 |
| num_overloaded_base_station | [int32](#int32) |  | 过载基站数 |
| num_gateway | [int32](#int32) |  | 总网关数 |
| num_ok_gateway | [int32](#int32) |  | 正常网关数 |
| num_ruined_gateway | [int32](#int32) |  | 损坏网关数 |
| num_stopped_gateway | [int32](#int32) |  | 断电网关数 |
| num_overloaded_gateway | [int32](#int32) |  | 过载网关数 |
| num_battery_gateway | [int32](#int32) |  | 直流电网关数 |
| power_consumption | [double](#double) |  | 基站耗电量(单位：kW·h) |





 


<a name="city-comm-output-v1-NodeStatus"></a>

### NodeStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| NODE_STATUS_UNSPECIFIED | 0 |  |
| NODE_STATUS_OK | 1 | 正常供电 |
| NODE_STATUS_BATTERY | 2 | 电池供电 |
| NODE_STATUS_FAILURE | 3 | 停电 |
| NODE_STATUS_RUINED | 4 | 损坏 |



<a name="city-comm-output-v1-PersonConnectStatus"></a>

### PersonConnectStatus
用户连接基站状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| PERSON_CONNECT_STATUS_UNSPECIFIED | 0 |  |
| PERSON_CONNECT_STATUS_OK | 1 | 可以连接到基站 |
| PERSON_CONNECT_STATUS_OUTAGE | 2 | 无法连接到基站 |



<a name="city-comm-output-v1-PersonDemandStatus"></a>

### PersonDemandStatus
用户需求满足状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| PERSON_DEMAND_STATUS_UNSPECIFIED | 0 |  |
| PERSON_DEMAND_STATUS_SATISFIED | 1 | 需求被满足 |
| PERSON_DEMAND_STATUS_UNSATISFIED | 2 | 需求不满足 |
| PERSON_DEMAND_STATUS_NO | 3 | 没有需求 |


 

 

 



<a name="city_comm_output_v1_output_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/comm/output/v1/output_service.proto



<a name="city-comm-output-v1-OutputRequest"></a>

### OutputRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| step | [int32](#int32) |  |  |
| nodes | [Node](#city-comm-output-v1-Node) | repeated | 设备状态 |
| base_stations | [BaseStation](#city-comm-output-v1-BaseStation) | repeated |  |
| signal_heatmap | [Signal](#city-comm-output-v1-Signal) |  | 五环区域信号强度热力图（500m网格形式呈现） |
| small_signal_heatmap | [Signal](#city-comm-output-v1-Signal) |  | 国贸区域信号强度热力图（50m网格形式呈现） |
| persons | [Person](#city-comm-output-v1-Person) | repeated | TODO(张钧): 基站和人的连接怎么表示？ 人相关的数据 |
| aois | [Aoi](#city-comm-output-v1-Aoi) | repeated | AOI相关的数据 |
| events | [city.event.v1.Events](#city-event-v1-Events) |  |  |
| statistics | [Statistics](#city-comm-output-v1-Statistics) |  | 统计值 |






<a name="city-comm-output-v1-OutputResponse"></a>

### OutputResponse






 

 

 


<a name="city-comm-output-v1-OutputService"></a>

### OutputService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Output | [OutputRequest](#city-comm-output-v1-OutputRequest) | [OutputResponse](#city-comm-output-v1-OutputResponse) |  |

 



<a name="city_config_v1_config-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/config/v1/config.proto



<a name="city-config-v1-MongoPath"></a>

### MongoPath
MongoDB配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| db | [string](#string) |  | 数据库名 |
| col | [string](#string) |  | 集合名 |






<a name="city-config-v1-OutputTarget"></a>

### OutputTarget
输出目标PostgreSQL


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sql | [string](#string) |  |  |





 

 

 

 



<a name="city_economy_v1_economy-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/economy/v1/economy.proto



<a name="city-economy-v1-Economy"></a>

### Economy
组织列表，对应MongoDB中的集合


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons | [Person](#city-economy-v1-Person) | repeated | 人 |
| orgs | [Org](#city-economy-v1-Org) | repeated | 组织列表 |






<a name="city-economy-v1-Employee"></a>

### Employee
员工


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | 员工ID |
| salary | [double](#double) |  | 薪水 |






<a name="city-economy-v1-Goods"></a>

### Goods
货物


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | 货物类型 |
| name | [string](#string) |  | 货物名称 |
| count | [int32](#int32) |  | 货物数量 |
| price | [double](#double) | optional | 货物价格（允许暂未定价） |






<a name="city-economy-v1-Job"></a>

### Job
岗位


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 岗位名称 |
| employee_count | [int32](#int32) |  | 岗位所需员工数量 |
| salary | [double](#double) | optional | 岗位薪水 |






<a name="city-economy-v1-Org"></a>

### Org
组织，具有员工、货物、资金等属性


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 组织ID |
| poi_id | [int32](#int32) |  | 组织所在的POI ID |
| employees | [Employee](#city-economy-v1-Employee) | repeated | 员工列表（初始化时由Orgs列表提供，单向绑定到person上） |
| jobs | [Job](#city-economy-v1-Job) | repeated | 岗位列表 |
| money | [double](#double) |  | 资金 |
| goods | [Goods](#city-economy-v1-Goods) | repeated | 货物 |
| functions | [string](#string) | repeated | 功能列表 buy: 购买货物 apply: 申请岗位 ... |






<a name="city-economy-v1-Person"></a>

### Person
个人（与Person一一对应，
没有绑定到city.economy.v1.Person的Person将无法参与经济模拟）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 与person id一致 |
| money | [double](#double) |  | 起始资金 |
| org_id | [int32](#int32) | optional | 所在组织ID（初始化时不提供，由组织中的员工列表设定） |





 

 

 

 



<a name="city_economy_v1_org_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/economy/v1/org_service.proto



<a name="city-economy-v1-GetOrgRequest"></a>

### GetOrgRequest
批量查询组织的经济情况请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_ids | [int32](#int32) | repeated | 待查询的组织的ID列表（为空时查询所有组织） |






<a name="city-economy-v1-GetOrgResponse"></a>

### GetOrgResponse
批量查询组织的经济情况响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orgs | [Org](#city-economy-v1-Org) | repeated | 组织的经济情况 |






<a name="city-economy-v1-UpdateOrgEmployeeRequest"></a>

### UpdateOrgEmployeeRequest
批量修改组织的员工请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [UpdateOrgEmployeeRequestItem](#city-economy-v1-UpdateOrgEmployeeRequestItem) | repeated | 待修改的组织员工变动 |






<a name="city-economy-v1-UpdateOrgEmployeeRequestItem"></a>

### UpdateOrgEmployeeRequestItem
组织员工变动


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  | 待修改的组织 |
| adds | [Employee](#city-economy-v1-Employee) | repeated | 新增的员工 |
| dels | [int32](#int32) | repeated | 删除的员工 |
| updates | [Employee](#city-economy-v1-Employee) | repeated | 修改薪水的员工 |






<a name="city-economy-v1-UpdateOrgEmployeeResponse"></a>

### UpdateOrgEmployeeResponse
批量修改组织的员工响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orgs | [Org](#city-economy-v1-Org) | repeated | 修改后的组织的经济情况 |






<a name="city-economy-v1-UpdateOrgGoodsRequest"></a>

### UpdateOrgGoodsRequest
批量修改组织的货物请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [UpdateOrgGoodsRequestItem](#city-economy-v1-UpdateOrgGoodsRequestItem) | repeated | 待修改的组织货物变动 |






<a name="city-economy-v1-UpdateOrgGoodsRequestItem"></a>

### UpdateOrgGoodsRequestItem
组织货物变动


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  | 待修改的组织 |
| goods | [Goods](#city-economy-v1-Goods) | repeated | 货物变动 按照(type, name)相等来判断是否为同一种货物 货物数量为增量，正数表示增加，负数表示减少 price如果未设定则沿用原来的价格，否则使用新的价格 |






<a name="city-economy-v1-UpdateOrgGoodsResponse"></a>

### UpdateOrgGoodsResponse
批量修改组织的货物响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orgs | [Org](#city-economy-v1-Org) | repeated | 修改后的组织的经济情况 |






<a name="city-economy-v1-UpdateOrgJobRequest"></a>

### UpdateOrgJobRequest
批量修改组织的岗位请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [UpdateOrgJobRequestItem](#city-economy-v1-UpdateOrgJobRequestItem) | repeated | 待修改的组织岗位变动 |






<a name="city-economy-v1-UpdateOrgJobRequestItem"></a>

### UpdateOrgJobRequestItem
组织岗位变动


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  | 待修改的组织 |
| jobs | [Job](#city-economy-v1-Job) | repeated | 岗位变动 按照name相等来判断是否为同一种岗位 岗位数量为增量，正数表示增加，负数表示减少 salary如果未设定则沿用原来的薪水，否则使用新的薪水 |






<a name="city-economy-v1-UpdateOrgJobResponse"></a>

### UpdateOrgJobResponse
批量修改组织的岗位响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orgs | [Org](#city-economy-v1-Org) | repeated | 修改后的组织的经济情况 |






<a name="city-economy-v1-UpdateOrgMoneyRequest"></a>

### UpdateOrgMoneyRequest
批量修改组织的资金请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [UpdateOrgMoneyRequestItem](#city-economy-v1-UpdateOrgMoneyRequestItem) | repeated | 待修改的组织资金变动 |






<a name="city-economy-v1-UpdateOrgMoneyRequestItem"></a>

### UpdateOrgMoneyRequestItem
组织资金变动


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  | 待修改的组织 |
| money | [double](#double) |  | 正数表示增加，负数表示减少 |






<a name="city-economy-v1-UpdateOrgMoneyResponse"></a>

### UpdateOrgMoneyResponse
批量修改组织的资金响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orgs | [Org](#city-economy-v1-Org) | repeated | 修改后的组织的经济情况 |





 

 

 


<a name="city-economy-v1-OrgService"></a>

### OrgService
组织经济情况接口

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetOrg | [GetOrgRequest](#city-economy-v1-GetOrgRequest) | [GetOrgResponse](#city-economy-v1-GetOrgResponse) | 批量查询组织的经济情况（员工、岗位、资金、货物） |
| UpdateOrgMoney | [UpdateOrgMoneyRequest](#city-economy-v1-UpdateOrgMoneyRequest) | [UpdateOrgMoneyResponse](#city-economy-v1-UpdateOrgMoneyResponse) | 批量修改组织的资金 |
| UpdateOrgGoods | [UpdateOrgGoodsRequest](#city-economy-v1-UpdateOrgGoodsRequest) | [UpdateOrgGoodsResponse](#city-economy-v1-UpdateOrgGoodsResponse) | 批量修改组织的货物 |
| UpdateOrgEmployee | [UpdateOrgEmployeeRequest](#city-economy-v1-UpdateOrgEmployeeRequest) | [UpdateOrgEmployeeResponse](#city-economy-v1-UpdateOrgEmployeeResponse) | 批量修改组织的员工 |
| UpdateOrgJob | [UpdateOrgJobRequest](#city-economy-v1-UpdateOrgJobRequest) | [UpdateOrgJobResponse](#city-economy-v1-UpdateOrgJobResponse) | 批量修改组织的岗位 |

 



<a name="city_economy_v1_person_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/economy/v1/person_service.proto



<a name="city-economy-v1-GetPersonRequest"></a>

### GetPersonRequest
批量查询人的经济情况请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_ids | [int32](#int32) | repeated | 待查询的人的ID列表（为空时查询所有人） |






<a name="city-economy-v1-GetPersonResponse"></a>

### GetPersonResponse
批量查询组织的经济情况响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons | [Person](#city-economy-v1-Person) | repeated | 人的经济情况 |






<a name="city-economy-v1-UpdatePersonMoneyRequest"></a>

### UpdatePersonMoneyRequest
批量修改人的资金请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [UpdatePersonMoneyRequestItem](#city-economy-v1-UpdatePersonMoneyRequestItem) | repeated | 待修改的人员资金变动 |






<a name="city-economy-v1-UpdatePersonMoneyRequestItem"></a>

### UpdatePersonMoneyRequestItem
人员资金变动


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | 待修改的人员id |
| money | [double](#double) |  | 资金变动（正数表示增加，负数表示减少） |






<a name="city-economy-v1-UpdatePersonMoneyResponse"></a>

### UpdatePersonMoneyResponse
批量修改人的资金响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons | [Person](#city-economy-v1-Person) | repeated | 修改后的人的经济情况 |





 

 

 


<a name="city-economy-v1-PersonService"></a>

### PersonService
个人经济情况接口

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPerson | [GetPersonRequest](#city-economy-v1-GetPersonRequest) | [GetPersonResponse](#city-economy-v1-GetPersonResponse) | 批量查询人的经济情况（资金、雇佣关系） |
| UpdatePersonMoney | [UpdatePersonMoneyRequest](#city-economy-v1-UpdatePersonMoneyRequest) | [UpdatePersonMoneyResponse](#city-economy-v1-UpdatePersonMoneyResponse) | 批量修改人的资金 |

 



<a name="city_economy_v2_economy-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/economy/v2/economy.proto



<a name="city-economy-v2-Agent"></a>

### Agent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | person ID |
| currency | [float](#float) | optional | currency |
| firm_id | [int32](#int32) | optional | 所属企业ID |
| skill | [float](#float) | optional | 技能水平 |
| consumption | [float](#float) | optional | 消费 |
| income | [float](#float) | optional | 收入 |






<a name="city-economy-v2-EconomyEntities"></a>

### EconomyEntities



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orgs | [Org](#city-economy-v2-Org) | repeated |  |
| agents | [Agent](#city-economy-v2-Agent) | repeated |  |






<a name="city-economy-v2-Org"></a>

### Org
组织
Organization


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 组织ID organization id |
| type | [OrgType](#city-economy-v2-OrgType) |  | 组织类别 organization type |
| nominal_gdp | [float](#float) | repeated | NBS |
| real_gdp | [float](#float) | repeated |  |
| unemployment | [float](#float) | repeated |  |
| wages | [float](#float) | repeated |  |
| prices | [float](#float) | repeated |  |
| inventory | [int32](#int32) | optional | Firm |
| price | [float](#float) | optional |  |
| currency | [float](#float) | optional | Firm &amp; Bank &amp; Government |
| interest_rate | [float](#float) | optional | Bank |
| bracket_cutoffs | [float](#float) | repeated | Government |
| bracket_rates | [float](#float) | repeated |  |
| consumption_currency | [float](#float) | repeated | NBS |
| consumption_propensity | [float](#float) | repeated |  |
| income_currency | [float](#float) | repeated |  |
| depression | [float](#float) | repeated |  |
| locus_control | [float](#float) | repeated |  |
| working_hours | [float](#float) | repeated |  |
| employees | [int32](#int32) | repeated | Firm: 企业的雇员列表 employees list for firm |
| citizens | [int32](#int32) | repeated | NBS &amp; Government: 公民列表 citizens list for NBS and government |
| demand | [int32](#int32) | optional | Firm: 总需求量 total demand for firm |
| sales | [int32](#int32) | optional | Firm: 总销售量 total sales for firm |





 


<a name="city-economy-v2-OrgType"></a>

### OrgType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ORG_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| ORG_TYPE_NBS | 1 | 国家统计局 NBS |
| ORG_TYPE_FIRM | 2 | 公司 firm |
| ORG_TYPE_BANK | 3 | 银行 bank |
| ORG_TYPE_GOVERNMENT | 4 | 政府 government |


 

 

 



<a name="city_economy_v2_org_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/economy/v2/org_service.proto



<a name="city-economy-v2-AddAgentRequest"></a>

### AddAgentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent | [Agent](#city-economy-v2-Agent) |  |  |






<a name="city-economy-v2-AddAgentResponse"></a>

### AddAgentResponse







<a name="city-economy-v2-AddCitizenRequest"></a>

### AddCitizenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| citizen_id | [int32](#int32) |  |  |






<a name="city-economy-v2-AddCitizenResponse"></a>

### AddCitizenResponse







<a name="city-economy-v2-AddCurrencyRequest"></a>

### AddCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| delta_currency | [float](#float) |  |  |






<a name="city-economy-v2-AddCurrencyResponse"></a>

### AddCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currency | [float](#float) |  |  |






<a name="city-economy-v2-AddEmployeeRequest"></a>

### AddEmployeeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| employee_id | [int32](#int32) |  |  |






<a name="city-economy-v2-AddEmployeeResponse"></a>

### AddEmployeeResponse







<a name="city-economy-v2-AddInterestRateRequest"></a>

### AddInterestRateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| delta_interest_rate | [float](#float) |  |  |






<a name="city-economy-v2-AddInterestRateResponse"></a>

### AddInterestRateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interest_rate | [float](#float) |  |  |






<a name="city-economy-v2-AddInventoryRequest"></a>

### AddInventoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| delta_inventory | [int32](#int32) |  |  |






<a name="city-economy-v2-AddInventoryResponse"></a>

### AddInventoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inventory | [int32](#int32) |  |  |






<a name="city-economy-v2-AddOrgRequest"></a>

### AddOrgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [Org](#city-economy-v2-Org) |  |  |






<a name="city-economy-v2-AddOrgResponse"></a>

### AddOrgResponse







<a name="city-economy-v2-AddPriceRequest"></a>

### AddPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| delta_price | [float](#float) |  |  |






<a name="city-economy-v2-AddPriceResponse"></a>

### AddPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [float](#float) |  |  |






<a name="city-economy-v2-BatchDeltaUpdateRequest"></a>

### BatchDeltaUpdateRequest
批量增量更新请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orgs | [DeltaUpdateOrgRequest](#city-economy-v2-DeltaUpdateOrgRequest) | repeated |  |
| agents | [DeltaUpdateAgentRequest](#city-economy-v2-DeltaUpdateAgentRequest) | repeated |  |






<a name="city-economy-v2-BatchDeltaUpdateResponse"></a>

### BatchDeltaUpdateResponse







<a name="city-economy-v2-BatchGetRequest"></a>

### BatchGetRequest
批量获取和更新


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [int32](#int32) | repeated | 要获取的 ID 列表 |
| type | [string](#string) |  | &#34;org&#34; 或 &#34;agent&#34; |






<a name="city-economy-v2-BatchGetResponse"></a>

### BatchGetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orgs | [Org](#city-economy-v2-Org) | repeated |  |
| agents | [Agent](#city-economy-v2-Agent) | repeated |  |






<a name="city-economy-v2-BatchUpdateRequest"></a>

### BatchUpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orgs | [Org](#city-economy-v2-Org) | repeated |  |
| agents | [Agent](#city-economy-v2-Agent) | repeated |  |






<a name="city-economy-v2-BatchUpdateResponse"></a>

### BatchUpdateResponse







<a name="city-economy-v2-CalculateConsumptionRequest"></a>

### CalculateConsumptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| firm_ids | [int32](#int32) | repeated | 多个公司ID |
| agent_id | [int32](#int32) |  | 代理ID |
| demands | [int32](#int32) | repeated | 对应每个公司的需求量 |
| consumption_accumulation | [bool](#bool) | optional | 是否累加消费，默认为true |






<a name="city-economy-v2-CalculateConsumptionResponse"></a>

### CalculateConsumptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actual_consumption | [float](#float) |  | 代理实际消费的钱 |
| success | [bool](#bool) |  | 消费是否成功 |






<a name="city-economy-v2-CalculateInterestRequest"></a>

### CalculateInterestRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bank_id | [int32](#int32) |  |  |
| agent_ids | [int32](#int32) | repeated | id of agents who has currency stored in the bank |






<a name="city-economy-v2-CalculateInterestResponse"></a>

### CalculateInterestResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_interest | [float](#float) |  | total interest that the agents got from the bank |
| updated_currencies | [float](#float) | repeated | currencies with interest of each agents, corresponds one-to-one with `agent_ids` by index |






<a name="city-economy-v2-CalculateRealGDPRequest"></a>

### CalculateRealGDPRequest
计算实际GDP的请求和响应消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nbs_agent_id | [int32](#int32) |  | NBS Agent的ID |






<a name="city-economy-v2-CalculateRealGDPResponse"></a>

### CalculateRealGDPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| real_gdp | [float](#float) |  | 计算得到的实际GDP |






<a name="city-economy-v2-CalculateTaxesDueRequest"></a>

### CalculateTaxesDueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| government_id | [int32](#int32) |  |  |
| agent_ids | [int32](#int32) | repeated | id of agents who needs to pay tax |
| incomes | [float](#float) | repeated | income of agents who needs to pay tax, corresponds one-to-one with `agent_ids` by index |
| enable_redistribution | [bool](#bool) |  | Whether redistribute the taxes to all agents or not |






<a name="city-economy-v2-CalculateTaxesDueResponse"></a>

### CalculateTaxesDueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taxes_due | [float](#float) |  | total taxes from agents |
| updated_incomes | [float](#float) | repeated | after-tax income of agents, corresponds one-to-one with `agent_ids` by index |






<a name="city-economy-v2-DeltaUpdateAgentRequest"></a>

### DeltaUpdateAgentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent_id | [int32](#int32) |  |  |
| delta_currency | [float](#float) | optional |  |
| delta_skill | [float](#float) | optional |  |
| delta_consumption | [float](#float) | optional |  |
| delta_income | [float](#float) | optional |  |






<a name="city-economy-v2-DeltaUpdateAgentResponse"></a>

### DeltaUpdateAgentResponse







<a name="city-economy-v2-DeltaUpdateOrgRequest"></a>

### DeltaUpdateOrgRequest
增量更新请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| delta_inventory | [float](#float) | optional |  |
| delta_price | [float](#float) | optional |  |
| delta_currency | [float](#float) | optional |  |
| delta_interest_rate | [float](#float) | optional |  |
| add_employees | [int32](#int32) | repeated |  |
| remove_employees | [int32](#int32) | repeated |  |






<a name="city-economy-v2-DeltaUpdateOrgResponse"></a>

### DeltaUpdateOrgResponse







<a name="city-economy-v2-GetAgentRequest"></a>

### GetAgentRequest
Agent 相关消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetAgentResponse"></a>

### GetAgentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent | [Agent](#city-economy-v2-Agent) |  |  |






<a name="city-economy-v2-GetBracketCutoffsRequest"></a>

### GetBracketCutoffsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetBracketCutoffsResponse"></a>

### GetBracketCutoffsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bracket_cutoffs | [float](#float) | repeated |  |






<a name="city-economy-v2-GetBracketRatesRequest"></a>

### GetBracketRatesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetBracketRatesResponse"></a>

### GetBracketRatesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bracket_rates | [float](#float) | repeated |  |






<a name="city-economy-v2-GetCitizensRequest"></a>

### GetCitizensRequest
Citizens 相关消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetCitizensResponse"></a>

### GetCitizensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| citizen_ids | [int32](#int32) | repeated |  |






<a name="city-economy-v2-GetConsumptionCurrencyRequest"></a>

### GetConsumptionCurrencyRequest
Consumption Currency


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetConsumptionCurrencyResponse"></a>

### GetConsumptionCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| consumption_currency | [float](#float) | repeated |  |






<a name="city-economy-v2-GetConsumptionPropensityRequest"></a>

### GetConsumptionPropensityRequest
Consumption Propensity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetConsumptionPropensityResponse"></a>

### GetConsumptionPropensityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| consumption_propensity | [float](#float) | repeated |  |






<a name="city-economy-v2-GetCurrencyRequest"></a>

### GetCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetCurrencyResponse"></a>

### GetCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currency | [float](#float) |  |  |






<a name="city-economy-v2-GetDepressionRequest"></a>

### GetDepressionRequest
Depression


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetDepressionResponse"></a>

### GetDepressionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| depression | [float](#float) | repeated |  |






<a name="city-economy-v2-GetEmployeesRequest"></a>

### GetEmployeesRequest
Employees 相关消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetEmployeesResponse"></a>

### GetEmployeesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| employee_ids | [int32](#int32) | repeated |  |






<a name="city-economy-v2-GetIncomeCurrencyRequest"></a>

### GetIncomeCurrencyRequest
Income Currency


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetIncomeCurrencyResponse"></a>

### GetIncomeCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| income_currency | [float](#float) | repeated |  |






<a name="city-economy-v2-GetInterestRateRequest"></a>

### GetInterestRateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetInterestRateResponse"></a>

### GetInterestRateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interest_rate | [float](#float) |  |  |






<a name="city-economy-v2-GetInventoryRequest"></a>

### GetInventoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetInventoryResponse"></a>

### GetInventoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inventory | [int32](#int32) |  |  |






<a name="city-economy-v2-GetLocusControlRequest"></a>

### GetLocusControlRequest
Locus of Control


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetLocusControlResponse"></a>

### GetLocusControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locus_control | [float](#float) | repeated |  |






<a name="city-economy-v2-GetNominalGDPRequest"></a>

### GetNominalGDPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetNominalGDPResponse"></a>

### GetNominalGDPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nominal_gdp | [float](#float) | repeated |  |






<a name="city-economy-v2-GetOrgEntityIdsRequest"></a>

### GetOrgEntityIdsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [OrgType](#city-economy-v2-OrgType) |  |  |






<a name="city-economy-v2-GetOrgEntityIdsResponse"></a>

### GetOrgEntityIdsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_ids | [int32](#int32) | repeated |  |






<a name="city-economy-v2-GetOrgRequest"></a>

### GetOrgRequest
Org 相关消息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetOrgResponse"></a>

### GetOrgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [Org](#city-economy-v2-Org) |  |  |






<a name="city-economy-v2-GetPriceRequest"></a>

### GetPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetPriceResponse"></a>

### GetPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [float](#float) |  |  |






<a name="city-economy-v2-GetPricesRequest"></a>

### GetPricesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetPricesResponse"></a>

### GetPricesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prices | [float](#float) | repeated |  |






<a name="city-economy-v2-GetRealGDPRequest"></a>

### GetRealGDPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetRealGDPResponse"></a>

### GetRealGDPResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| real_gdp | [float](#float) | repeated |  |






<a name="city-economy-v2-GetUnemploymentRequest"></a>

### GetUnemploymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetUnemploymentResponse"></a>

### GetUnemploymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unemployment | [float](#float) | repeated |  |






<a name="city-economy-v2-GetWagesRequest"></a>

### GetWagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetWagesResponse"></a>

### GetWagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wages | [float](#float) | repeated |  |






<a name="city-economy-v2-GetWorkingHoursRequest"></a>

### GetWorkingHoursRequest
Working Hours


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-GetWorkingHoursResponse"></a>

### GetWorkingHoursResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| working_hours | [float](#float) | repeated |  |






<a name="city-economy-v2-LoadEconomyEntitiesRequest"></a>

### LoadEconomyEntitiesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_path | [string](#string) |  |  |






<a name="city-economy-v2-LoadEconomyEntitiesResponse"></a>

### LoadEconomyEntitiesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_ids | [int32](#int32) | repeated | 组织ID列表 |
| agent_ids | [int32](#int32) | repeated | Agent ID列表 |






<a name="city-economy-v2-RemoveAgentRequest"></a>

### RemoveAgentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent_id | [int32](#int32) |  |  |






<a name="city-economy-v2-RemoveAgentResponse"></a>

### RemoveAgentResponse







<a name="city-economy-v2-RemoveCitizenRequest"></a>

### RemoveCitizenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| citizen_id | [int32](#int32) |  |  |






<a name="city-economy-v2-RemoveCitizenResponse"></a>

### RemoveCitizenResponse







<a name="city-economy-v2-RemoveEmployeeRequest"></a>

### RemoveEmployeeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| employee_id | [int32](#int32) |  |  |






<a name="city-economy-v2-RemoveEmployeeResponse"></a>

### RemoveEmployeeResponse







<a name="city-economy-v2-RemoveOrgRequest"></a>

### RemoveOrgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |






<a name="city-economy-v2-RemoveOrgResponse"></a>

### RemoveOrgResponse







<a name="city-economy-v2-SaveEconomyEntitiesRequest"></a>

### SaveEconomyEntitiesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_path | [string](#string) |  |  |






<a name="city-economy-v2-SaveEconomyEntitiesResponse"></a>

### SaveEconomyEntitiesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_ids | [int32](#int32) | repeated | 组织ID列表 |
| agent_ids | [int32](#int32) | repeated | Agent ID列表 |






<a name="city-economy-v2-SetBracketCutoffsRequest"></a>

### SetBracketCutoffsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| bracket_cutoffs | [float](#float) | repeated |  |






<a name="city-economy-v2-SetBracketCutoffsResponse"></a>

### SetBracketCutoffsResponse







<a name="city-economy-v2-SetBracketRatesRequest"></a>

### SetBracketRatesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| bracket_rates | [float](#float) | repeated |  |






<a name="city-economy-v2-SetBracketRatesResponse"></a>

### SetBracketRatesResponse







<a name="city-economy-v2-SetCitizensRequest"></a>

### SetCitizensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| citizen_ids | [int32](#int32) | repeated |  |






<a name="city-economy-v2-SetCitizensResponse"></a>

### SetCitizensResponse







<a name="city-economy-v2-SetConsumptionCurrencyRequest"></a>

### SetConsumptionCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| consumption_currency | [float](#float) | repeated |  |






<a name="city-economy-v2-SetConsumptionCurrencyResponse"></a>

### SetConsumptionCurrencyResponse







<a name="city-economy-v2-SetConsumptionPropensityRequest"></a>

### SetConsumptionPropensityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| consumption_propensity | [float](#float) | repeated |  |






<a name="city-economy-v2-SetConsumptionPropensityResponse"></a>

### SetConsumptionPropensityResponse







<a name="city-economy-v2-SetCurrencyRequest"></a>

### SetCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| currency | [float](#float) |  |  |






<a name="city-economy-v2-SetCurrencyResponse"></a>

### SetCurrencyResponse







<a name="city-economy-v2-SetDepressionRequest"></a>

### SetDepressionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| depression | [float](#float) | repeated |  |






<a name="city-economy-v2-SetDepressionResponse"></a>

### SetDepressionResponse







<a name="city-economy-v2-SetEmployeesRequest"></a>

### SetEmployeesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| employee_ids | [int32](#int32) | repeated |  |






<a name="city-economy-v2-SetEmployeesResponse"></a>

### SetEmployeesResponse







<a name="city-economy-v2-SetIncomeCurrencyRequest"></a>

### SetIncomeCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| income_currency | [float](#float) | repeated |  |






<a name="city-economy-v2-SetIncomeCurrencyResponse"></a>

### SetIncomeCurrencyResponse







<a name="city-economy-v2-SetInterestRateRequest"></a>

### SetInterestRateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| interest_rate | [float](#float) |  |  |






<a name="city-economy-v2-SetInterestRateResponse"></a>

### SetInterestRateResponse







<a name="city-economy-v2-SetInventoryRequest"></a>

### SetInventoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| inventory | [int32](#int32) |  |  |






<a name="city-economy-v2-SetInventoryResponse"></a>

### SetInventoryResponse







<a name="city-economy-v2-SetLocusControlRequest"></a>

### SetLocusControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| locus_control | [float](#float) | repeated |  |






<a name="city-economy-v2-SetLocusControlResponse"></a>

### SetLocusControlResponse







<a name="city-economy-v2-SetNominalGDPRequest"></a>

### SetNominalGDPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| nominal_gdp | [float](#float) | repeated |  |






<a name="city-economy-v2-SetNominalGDPResponse"></a>

### SetNominalGDPResponse







<a name="city-economy-v2-SetPriceRequest"></a>

### SetPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| price | [float](#float) |  |  |






<a name="city-economy-v2-SetPriceResponse"></a>

### SetPriceResponse







<a name="city-economy-v2-SetPricesRequest"></a>

### SetPricesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| prices | [float](#float) | repeated |  |






<a name="city-economy-v2-SetPricesResponse"></a>

### SetPricesResponse







<a name="city-economy-v2-SetRealGDPRequest"></a>

### SetRealGDPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| real_gdp | [float](#float) | repeated |  |






<a name="city-economy-v2-SetRealGDPResponse"></a>

### SetRealGDPResponse







<a name="city-economy-v2-SetUnemploymentRequest"></a>

### SetUnemploymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| unemployment | [float](#float) | repeated |  |






<a name="city-economy-v2-SetUnemploymentResponse"></a>

### SetUnemploymentResponse







<a name="city-economy-v2-SetWagesRequest"></a>

### SetWagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| wages | [float](#float) | repeated |  |






<a name="city-economy-v2-SetWagesResponse"></a>

### SetWagesResponse







<a name="city-economy-v2-SetWorkingHoursRequest"></a>

### SetWorkingHoursRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org_id | [int32](#int32) |  |  |
| working_hours | [float](#float) | repeated |  |






<a name="city-economy-v2-SetWorkingHoursResponse"></a>

### SetWorkingHoursResponse







<a name="city-economy-v2-UpdateAgentRequest"></a>

### UpdateAgentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent | [Agent](#city-economy-v2-Agent) |  |  |






<a name="city-economy-v2-UpdateAgentResponse"></a>

### UpdateAgentResponse







<a name="city-economy-v2-UpdateOrgRequest"></a>

### UpdateOrgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| org | [Org](#city-economy-v2-Org) |  |  |






<a name="city-economy-v2-UpdateOrgResponse"></a>

### UpdateOrgResponse






 

 

 


<a name="city-economy-v2-OrgService"></a>

### OrgService
组织经济情况接口

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddOrg | [AddOrgRequest](#city-economy-v2-AddOrgRequest) | [AddOrgResponse](#city-economy-v2-AddOrgResponse) | 添加组织 add org |
| RemoveOrg | [RemoveOrgRequest](#city-economy-v2-RemoveOrgRequest) | [RemoveOrgResponse](#city-economy-v2-RemoveOrgResponse) | 移除组织 remove org |
| GetOrg | [GetOrgRequest](#city-economy-v2-GetOrgRequest) | [GetOrgResponse](#city-economy-v2-GetOrgResponse) | 获取组织 get org |
| UpdateOrg | [UpdateOrgRequest](#city-economy-v2-UpdateOrgRequest) | [UpdateOrgResponse](#city-economy-v2-UpdateOrgResponse) | 更新组织 update org |
| AddAgent | [AddAgentRequest](#city-economy-v2-AddAgentRequest) | [AddAgentResponse](#city-economy-v2-AddAgentResponse) | 添加Agent add agent |
| RemoveAgent | [RemoveAgentRequest](#city-economy-v2-RemoveAgentRequest) | [RemoveAgentResponse](#city-economy-v2-RemoveAgentResponse) | 移除Agent remove agent |
| GetNominalGDP | [GetNominalGDPRequest](#city-economy-v2-GetNominalGDPRequest) | [GetNominalGDPResponse](#city-economy-v2-GetNominalGDPResponse) | Nominal GDP |
| SetNominalGDP | [SetNominalGDPRequest](#city-economy-v2-SetNominalGDPRequest) | [SetNominalGDPResponse](#city-economy-v2-SetNominalGDPResponse) |  |
| GetRealGDP | [GetRealGDPRequest](#city-economy-v2-GetRealGDPRequest) | [GetRealGDPResponse](#city-economy-v2-GetRealGDPResponse) | Real GDP |
| SetRealGDP | [SetRealGDPRequest](#city-economy-v2-SetRealGDPRequest) | [SetRealGDPResponse](#city-economy-v2-SetRealGDPResponse) |  |
| GetUnemployment | [GetUnemploymentRequest](#city-economy-v2-GetUnemploymentRequest) | [GetUnemploymentResponse](#city-economy-v2-GetUnemploymentResponse) | Unemployment |
| SetUnemployment | [SetUnemploymentRequest](#city-economy-v2-SetUnemploymentRequest) | [SetUnemploymentResponse](#city-economy-v2-SetUnemploymentResponse) |  |
| GetWages | [GetWagesRequest](#city-economy-v2-GetWagesRequest) | [GetWagesResponse](#city-economy-v2-GetWagesResponse) | Wages |
| SetWages | [SetWagesRequest](#city-economy-v2-SetWagesRequest) | [SetWagesResponse](#city-economy-v2-SetWagesResponse) |  |
| GetPrices | [GetPricesRequest](#city-economy-v2-GetPricesRequest) | [GetPricesResponse](#city-economy-v2-GetPricesResponse) | Prices |
| SetPrices | [SetPricesRequest](#city-economy-v2-SetPricesRequest) | [SetPricesResponse](#city-economy-v2-SetPricesResponse) |  |
| GetInventory | [GetInventoryRequest](#city-economy-v2-GetInventoryRequest) | [GetInventoryResponse](#city-economy-v2-GetInventoryResponse) | Inventory |
| SetInventory | [SetInventoryRequest](#city-economy-v2-SetInventoryRequest) | [SetInventoryResponse](#city-economy-v2-SetInventoryResponse) |  |
| AddInventory | [AddInventoryRequest](#city-economy-v2-AddInventoryRequest) | [AddInventoryResponse](#city-economy-v2-AddInventoryResponse) |  |
| GetPrice | [GetPriceRequest](#city-economy-v2-GetPriceRequest) | [GetPriceResponse](#city-economy-v2-GetPriceResponse) | Price |
| SetPrice | [SetPriceRequest](#city-economy-v2-SetPriceRequest) | [SetPriceResponse](#city-economy-v2-SetPriceResponse) |  |
| AddPrice | [AddPriceRequest](#city-economy-v2-AddPriceRequest) | [AddPriceResponse](#city-economy-v2-AddPriceResponse) |  |
| GetCurrency | [GetCurrencyRequest](#city-economy-v2-GetCurrencyRequest) | [GetCurrencyResponse](#city-economy-v2-GetCurrencyResponse) | Currency |
| SetCurrency | [SetCurrencyRequest](#city-economy-v2-SetCurrencyRequest) | [SetCurrencyResponse](#city-economy-v2-SetCurrencyResponse) |  |
| AddCurrency | [AddCurrencyRequest](#city-economy-v2-AddCurrencyRequest) | [AddCurrencyResponse](#city-economy-v2-AddCurrencyResponse) |  |
| GetInterestRate | [GetInterestRateRequest](#city-economy-v2-GetInterestRateRequest) | [GetInterestRateResponse](#city-economy-v2-GetInterestRateResponse) | Interest Rate |
| SetInterestRate | [SetInterestRateRequest](#city-economy-v2-SetInterestRateRequest) | [SetInterestRateResponse](#city-economy-v2-SetInterestRateResponse) |  |
| AddInterestRate | [AddInterestRateRequest](#city-economy-v2-AddInterestRateRequest) | [AddInterestRateResponse](#city-economy-v2-AddInterestRateResponse) |  |
| GetBracketCutoffs | [GetBracketCutoffsRequest](#city-economy-v2-GetBracketCutoffsRequest) | [GetBracketCutoffsResponse](#city-economy-v2-GetBracketCutoffsResponse) | Bracket Cutoffs |
| SetBracketCutoffs | [SetBracketCutoffsRequest](#city-economy-v2-SetBracketCutoffsRequest) | [SetBracketCutoffsResponse](#city-economy-v2-SetBracketCutoffsResponse) |  |
| GetBracketRates | [GetBracketRatesRequest](#city-economy-v2-GetBracketRatesRequest) | [GetBracketRatesResponse](#city-economy-v2-GetBracketRatesResponse) | Bracket Rates |
| SetBracketRates | [SetBracketRatesRequest](#city-economy-v2-SetBracketRatesRequest) | [SetBracketRatesResponse](#city-economy-v2-SetBracketRatesResponse) |  |
| CalculateTaxesDue | [CalculateTaxesDueRequest](#city-economy-v2-CalculateTaxesDueRequest) | [CalculateTaxesDueResponse](#city-economy-v2-CalculateTaxesDueResponse) | Taxes Due |
| CalculateConsumption | [CalculateConsumptionRequest](#city-economy-v2-CalculateConsumptionRequest) | [CalculateConsumptionResponse](#city-economy-v2-CalculateConsumptionResponse) | Consumption |
| CalculateInterest | [CalculateInterestRequest](#city-economy-v2-CalculateInterestRequest) | [CalculateInterestResponse](#city-economy-v2-CalculateInterestResponse) | Consumption |
| SaveEconomyEntities | [SaveEconomyEntitiesRequest](#city-economy-v2-SaveEconomyEntitiesRequest) | [SaveEconomyEntitiesResponse](#city-economy-v2-SaveEconomyEntitiesResponse) | Save |
| LoadEconomyEntities | [LoadEconomyEntitiesRequest](#city-economy-v2-LoadEconomyEntitiesRequest) | [LoadEconomyEntitiesResponse](#city-economy-v2-LoadEconomyEntitiesResponse) | Load |
| GetConsumptionCurrency | [GetConsumptionCurrencyRequest](#city-economy-v2-GetConsumptionCurrencyRequest) | [GetConsumptionCurrencyResponse](#city-economy-v2-GetConsumptionCurrencyResponse) | Consumption Currency |
| SetConsumptionCurrency | [SetConsumptionCurrencyRequest](#city-economy-v2-SetConsumptionCurrencyRequest) | [SetConsumptionCurrencyResponse](#city-economy-v2-SetConsumptionCurrencyResponse) |  |
| GetConsumptionPropensity | [GetConsumptionPropensityRequest](#city-economy-v2-GetConsumptionPropensityRequest) | [GetConsumptionPropensityResponse](#city-economy-v2-GetConsumptionPropensityResponse) | Consumption Propensity |
| SetConsumptionPropensity | [SetConsumptionPropensityRequest](#city-economy-v2-SetConsumptionPropensityRequest) | [SetConsumptionPropensityResponse](#city-economy-v2-SetConsumptionPropensityResponse) |  |
| GetIncomeCurrency | [GetIncomeCurrencyRequest](#city-economy-v2-GetIncomeCurrencyRequest) | [GetIncomeCurrencyResponse](#city-economy-v2-GetIncomeCurrencyResponse) | Income Currency |
| SetIncomeCurrency | [SetIncomeCurrencyRequest](#city-economy-v2-SetIncomeCurrencyRequest) | [SetIncomeCurrencyResponse](#city-economy-v2-SetIncomeCurrencyResponse) |  |
| GetDepression | [GetDepressionRequest](#city-economy-v2-GetDepressionRequest) | [GetDepressionResponse](#city-economy-v2-GetDepressionResponse) | Depression |
| SetDepression | [SetDepressionRequest](#city-economy-v2-SetDepressionRequest) | [SetDepressionResponse](#city-economy-v2-SetDepressionResponse) |  |
| GetLocusControl | [GetLocusControlRequest](#city-economy-v2-GetLocusControlRequest) | [GetLocusControlResponse](#city-economy-v2-GetLocusControlResponse) | Locus of Control |
| SetLocusControl | [SetLocusControlRequest](#city-economy-v2-SetLocusControlRequest) | [SetLocusControlResponse](#city-economy-v2-SetLocusControlResponse) |  |
| GetWorkingHours | [GetWorkingHoursRequest](#city-economy-v2-GetWorkingHoursRequest) | [GetWorkingHoursResponse](#city-economy-v2-GetWorkingHoursResponse) | Working Hours |
| SetWorkingHours | [SetWorkingHoursRequest](#city-economy-v2-SetWorkingHoursRequest) | [SetWorkingHoursResponse](#city-economy-v2-SetWorkingHoursResponse) |  |
| GetOrgEntityIds | [GetOrgEntityIdsRequest](#city-economy-v2-GetOrgEntityIdsRequest) | [GetOrgEntityIdsResponse](#city-economy-v2-GetOrgEntityIdsResponse) | Org Entity Ids |
| GetEmployees | [GetEmployeesRequest](#city-economy-v2-GetEmployeesRequest) | [GetEmployeesResponse](#city-economy-v2-GetEmployeesResponse) | Employees 相关接口 |
| SetEmployees | [SetEmployeesRequest](#city-economy-v2-SetEmployeesRequest) | [SetEmployeesResponse](#city-economy-v2-SetEmployeesResponse) |  |
| AddEmployee | [AddEmployeeRequest](#city-economy-v2-AddEmployeeRequest) | [AddEmployeeResponse](#city-economy-v2-AddEmployeeResponse) |  |
| RemoveEmployee | [RemoveEmployeeRequest](#city-economy-v2-RemoveEmployeeRequest) | [RemoveEmployeeResponse](#city-economy-v2-RemoveEmployeeResponse) |  |
| GetCitizens | [GetCitizensRequest](#city-economy-v2-GetCitizensRequest) | [GetCitizensResponse](#city-economy-v2-GetCitizensResponse) | Citizens 相关接口 |
| SetCitizens | [SetCitizensRequest](#city-economy-v2-SetCitizensRequest) | [SetCitizensResponse](#city-economy-v2-SetCitizensResponse) |  |
| AddCitizen | [AddCitizenRequest](#city-economy-v2-AddCitizenRequest) | [AddCitizenResponse](#city-economy-v2-AddCitizenResponse) |  |
| RemoveCitizen | [RemoveCitizenRequest](#city-economy-v2-RemoveCitizenRequest) | [RemoveCitizenResponse](#city-economy-v2-RemoveCitizenResponse) |  |
| GetAgent | [GetAgentRequest](#city-economy-v2-GetAgentRequest) | [GetAgentResponse](#city-economy-v2-GetAgentResponse) | Agent 相关接口 |
| UpdateAgent | [UpdateAgentRequest](#city-economy-v2-UpdateAgentRequest) | [UpdateAgentResponse](#city-economy-v2-UpdateAgentResponse) |  |
| BatchGet | [BatchGetRequest](#city-economy-v2-BatchGetRequest) | [BatchGetResponse](#city-economy-v2-BatchGetResponse) | 批量获取 |
| BatchUpdate | [BatchUpdateRequest](#city-economy-v2-BatchUpdateRequest) | [BatchUpdateResponse](#city-economy-v2-BatchUpdateResponse) | 批量更新 |
| DeltaUpdateOrg | [DeltaUpdateOrgRequest](#city-economy-v2-DeltaUpdateOrgRequest) | [DeltaUpdateOrgResponse](#city-economy-v2-DeltaUpdateOrgResponse) | 增量更新 |
| DeltaUpdateAgent | [DeltaUpdateAgentRequest](#city-economy-v2-DeltaUpdateAgentRequest) | [DeltaUpdateAgentResponse](#city-economy-v2-DeltaUpdateAgentResponse) |  |
| BatchDeltaUpdate | [BatchDeltaUpdateRequest](#city-economy-v2-BatchDeltaUpdateRequest) | [BatchDeltaUpdateResponse](#city-economy-v2-BatchDeltaUpdateResponse) | 批量增量更新 |
| CalculateRealGDP | [CalculateRealGDPRequest](#city-economy-v2-CalculateRealGDPRequest) | [CalculateRealGDPResponse](#city-economy-v2-CalculateRealGDPResponse) | 计算实际GDP |

 



<a name="city_elec_input_v1_config-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/elec/input/v1/config.proto



<a name="city-elec-input-v1-Config"></a>

### Config



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mongo | [Mongo](#city-elec-input-v1-Mongo) |  |  |
| control | [Control](#city-elec-input-v1-Control) |  |  |
| output | [Output](#city-elec-input-v1-Output) |  |  |






<a name="city-elec-input-v1-Control"></a>

### Control



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| step | [ControlStep](#city-elec-input-v1-ControlStep) |  |  |






<a name="city-elec-input-v1-ControlStep"></a>

### ControlStep



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [int32](#int32) |  |  |
| total | [int32](#int32) |  |  |






<a name="city-elec-input-v1-Mongo"></a>

### Mongo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  |  |
| map | [city.config.v1.MongoPath](#city-config-v1-MongoPath) |  |  |
| facilities | [city.config.v1.MongoPath](#city-config-v1-MongoPath) |  |  |






<a name="city-elec-input-v1-Output"></a>

### Output



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target | [city.config.v1.OutputTarget](#city-config-v1-OutputTarget) |  |  |
| switch | [OutputSwitch](#city-elec-input-v1-OutputSwitch) |  |  |






<a name="city-elec-input-v1-OutputSwitch"></a>

### OutputSwitch
是否输出各类数据


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [bool](#bool) |  | 电网节点状态 |
| aoi | [bool](#bool) |  |  |
| event | [bool](#bool) |  |  |





 

 

 

 



<a name="city_elec_input_v1_input-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/elec/input/v1/input.proto



<a name="city-elec-input-v1-Facilities"></a>

### Facilities
设施集合，对应于mongodb一个collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| facilities | [Facility](#city-elec-input-v1-Facility) | repeated |  |
| repair_stations | [RepairStation](#city-elec-input-v1-RepairStation) | repeated |  |






<a name="city-elec-input-v1-Facility"></a>

### Facility



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| position | [city.geo.v2.LongLatPosition](#city-geo-v2-LongLatPosition) |  |  |
| type | [FacilityType](#city-elec-input-v1-FacilityType) |  |  |
| relation | [int32](#int32) | repeated | 当前节点的邻居节点的id |
| foreign_id | [int32](#int32) | optional | 在其它关联的网络中如水网使用时，可使用外部id 对于负载，该值表示其在对应模拟中的id |
| aoi_id | [int32](#int32) | optional | 对于电力设施，该值表示所在aoi id |
| num_transformer | [int32](#int32) | optional | 对于10kv变压器组，该值表示变压器组中变压器的数量 |






<a name="city-elec-input-v1-RepairStation"></a>

### RepairStation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aoi_id | [int32](#int32) |  |  |
| position | [city.geo.v2.LongLatPosition](#city-geo-v2-LongLatPosition) |  |  |





 


<a name="city-elec-input-v1-FacilityType"></a>

### FacilityType


| Name | Number | Description |
| ---- | ------ | ----------- |
| FACILITY_TYPE_UNSPECIFIED | 0 | 电网相关的基础设施分类 |
| FACILITY_TYPE_POWER_STATION | 1 | 发电站 |
| FACILITY_TYPE_TRANSFORMER_500 | 2 | 不同电压的变压器 |
| FACILITY_TYPE_TRANSFORMER_220 | 3 |  |
| FACILITY_TYPE_TRANSFORMER_110 | 4 |  |
| FACILITY_TYPE_TRANSFORMER_10 | 5 |  |
| FACILITY_TYPE_BASE_STATION | 6 | 通信基站 |
| FACILITY_TYPE_GATEWAY | 7 | 网关 |
| FACILITY_TYPE_DRAINAGE_PUMP | 8 | 排水水泵 |
| FACILITY_TYPE_TRAFFIC_LIGHT | 9 | 交通灯 |
| FACILITY_TYPE_AOI | 10 | AOI |
| FACILITY_TYPE_SUPPLY_PUMP | 11 | 供水水泵 |


 

 

 



<a name="city_map_v2_light-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/map/v2/light.proto



<a name="city-map-v2-AvailablePhase"></a>

### AvailablePhase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| states | [LightState](#city-map-v2-LightState) | repeated | 描述最大压力信控的可行相位，由每个lane的灯控情况组成，lane与Junction.lane_ids一一对应 Describes the feasible phase for max pressure algorithm, consisting of the lighting control situation for each lane in the junction, nd the lane corresponds one-to-one with junction.lane_ids |






<a name="city-map-v2-Phase"></a>

### Phase
交通灯相位
traffic light phase


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| duration | [double](#double) |  | 相位持续时间，单位秒 Phase duration in seconds |
| states | [LightState](#city-map-v2-LightState) | repeated | 描述该相位下每个lane的灯控情况，lane与Junction.lane_ids一一对应 The lighting control situation of each lane in this phase, and the lane corresponds one-to-one with junction.lane_ids |






<a name="city-map-v2-TrafficLight"></a>

### TrafficLight
交通灯
traffic light


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| junction_id | [int32](#int32) |  | 所在路口id ID of the junction where the traffic light is at |
| phases | [Phase](#city-map-v2-Phase) | repeated | 相位循环的一个循环周期 One cycle of phase cycling |





 


<a name="city-map-v2-LightState"></a>

### LightState
交通灯的状态
traffic light state

| Name | Number | Description |
| ---- | ------ | ----------- |
| LIGHT_STATE_UNSPECIFIED | 0 | 未指定 unspecified |
| LIGHT_STATE_RED | 1 | 红灯 red light |
| LIGHT_STATE_GREEN | 2 | 绿灯 green light |
| LIGHT_STATE_YELLOW | 3 | 黄灯 yellow light |


 

 

 



<a name="city_map_v2_map-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/map/v2/map.proto



<a name="city-map-v2-Aoi"></a>

### Aoi
Aoi，用于描述地图上的区域
Aoi, describing a region on the map


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Aoi ID（从5_0000_0000开始） Aoi ID (starting from 5_0000_0000) |
| name | [string](#string) |  | Aoi名字 Aoi name |
| type | [AoiType](#city-map-v2-AoiType) |  | Aoi类型 Aoi type |
| driving_positions | [city.geo.v2.LanePosition](#city-geo-v2-LanePosition) | repeated | Aoi与行车路网的连接点 Connection point between Aoi and driving lanes |
| walking_positions | [city.geo.v2.LanePosition](#city-geo-v2-LanePosition) | repeated | Aoi与步行路网的连接点 Connection point between Aoi and walking lanes |
| positions | [city.geo.v2.XYPosition](#city-geo-v2-XYPosition) | repeated | Aoi原始位置（如果只有一个值，则为Aoi所在的点，否则为Aoi多边形的边界） Aoi original position (if there is only one value, it is the point where Aoi is located, otherwise it is the boundary of the Aoi polygon) |
| driving_gates | [city.geo.v2.XYPosition](#city-geo-v2-XYPosition) | repeated | Aoi与行车路网连接时在自身边界上的连接点, 与driving_positions按索引一一对应 The connection point on its own boundary when Aoi is connected to the driving lanes corresponds one-to-one with driving_positions by index. |
| walking_gates | [city.geo.v2.XYPosition](#city-geo-v2-XYPosition) | repeated | Aoi与步行路网连接时在自身边界上的连接点, 与walking_positions按索引一一对应 The connection point on its own boundary when Aoi is connected to the walking lanes corresponds one-to-one with walking_positions by index. |
| area | [double](#double) | optional | Aoi面积, 若是Poi则无此字段 Aoi area, if it is Poi, there is no such field |
| land_use | [LandUseType](#city-map-v2-LandUseType) | optional | **Deprecated.** 土地利用分类，若是Poi则无此字段 Land use type, if it is Poi, there is no such field

弃用 deprecated |
| urban_land_use | [string](#string) | optional | 城市建设用地分类，参照执行标准GB 50137-2011（https://www.planning.org.cn/law/uploads/2013/1383993139.pdf） Urban Land use type, refer to the national standard GB 50137-2011 |
| poi_ids | [int32](#int32) | repeated | Aoi包含的Poi Pois contained in Aoi |






<a name="city-map-v2-Header"></a>

### Header
地图元信息
Map meta information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 地图名称 Map name |
| date | [string](#string) |  | 地图创建时间 Map creation time |
| north | [double](#double) |  | 最大纬度对应的y坐标 y coordinate corresponding to the maximum latitude |
| south | [double](#double) |  | 最小纬度对应的y坐标 y coordinate corresponding to the minimum latitude |
| east | [double](#double) |  | 最大经度对应的x坐标 x coordinate corresponding to the maximum longitude |
| west | [double](#double) |  | 最小经度对应的x坐标 x coordinate corresponding to the minimum longitude |
| projection | [string](#string) |  | PROJ.4 投影字符串，用以支持xy坐标到其他坐标系的转换 PROJ.4 projection string to support the conversion of xy coordinates to other coordinate systems |
| taz_x_step | [double](#double) | optional | 在x方向划分TAZ的步长 Step size of the TAZ in the x-direction |
| taz_y_step | [double](#double) | optional | 在y方向划分TAZ的步长 Step size of the TAZ in the y-direction |






<a name="city-map-v2-HeuristicTAZCost"></a>

### HeuristicTAZCost
交通分析区，用于预计算公共交通权重
// TAZ, used to precalculate public transport costs
message TransportationAnalysisZone{
  int32 x_id = 1;
  int32 y_id = 2;
  double x_minimum = 3;
  double x_maximum = 4;
  double y_minimum = 5;
  double y_maximum = 6;
}
预计算公共交通权重
Precalculate public transport costs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taz_x_id | [int32](#int32) |  |  |
| taz_y_id | [int32](#int32) |  |  |
| aoi_id | [int32](#int32) |  |  |
| cost | [double](#double) |  |  |






<a name="city-map-v2-Junction"></a>

### Junction
Junction，用于描述路口
Junction, describing road intersections


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 路口id（从3_0000_0000开始） Junction id (starting from 3_0000_0000) |
| lane_ids | [int32](#int32) | repeated | 属于该路口Junction的所有车道/人行道等lane All driving/walking lanes belonging to this junction. |
| driving_lane_groups | [JunctionLaneGroup](#city-map-v2-JunctionLaneGroup) | repeated | 属于该路口Junction的所有行车车道组 All driving lane groups belonging to this junction |
| phases | [AvailablePhase](#city-map-v2-AvailablePhase) | repeated | 所有可用信号灯相位 All available phases for max pressure algorithm |
| fixed_program | [TrafficLight](#city-map-v2-TrafficLight) | optional | 默认固定相位信号灯 Default fixed phases traffic light |






<a name="city-map-v2-JunctionLaneGroup"></a>

### JunctionLaneGroup
车道组，用于描述路口内的车道组合
Lane group, describing the combination of lanes within an intersection
具有相同入口道路和出口道路的车道组成一个车道组
Lanes with the same entrance and exit roads form a lane group
车道组是信控处理、路口通行的基本单元
Lane group is the basic unit for signal control and traffic in the junction.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| in_road_id | [int32](#int32) |  | 该车道组的入口道路 The entrance road to this lane group |
| in_angle | [double](#double) |  | 该车道组的入口角度（弧度制） The entrance angle of this lane group (in radians) |
| out_road_id | [int32](#int32) |  | 该车道组的出口道路 The exit road for this lane group |
| out_angle | [double](#double) |  | 该车道组的出口角度（弧度制） The exit angle of this lane group (in radians) |
| lane_ids | [int32](#int32) | repeated | 该车道组包含的车道 Lanes in the group |
| turn | [LaneTurn](#city-map-v2-LaneTurn) |  | 该车道组的转向属性 The turn type of this lane group |






<a name="city-map-v2-Lane"></a>

### Lane
Lane，用于描述道路上的车道、人行道等
Lane, used to describe lanes, sidewalks, etc. on the road


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 车道id（从0开始） lane id (starts from 0) |
| type | [LaneType](#city-map-v2-LaneType) |  | 车道类型 lane type |
| turn | [LaneTurn](#city-map-v2-LaneTurn) |  | 车道转向 lane turn type |
| max_speed | [double](#double) |  | 限速 (m/s) max speed (m/s) |
| length | [double](#double) |  | 中心线长度（单位：米） centerline length (in meters) |
| width | [double](#double) |  | 车道宽度（单位：米） lane width (in meters) |
| center_line | [Polyline](#city-map-v2-Polyline) |  | 车道中心线（车辆/行车轨迹线） Lane center line (vehicle/driving line) |
| left_border_line | [Polyline](#city-map-v2-Polyline) |  | **Deprecated.** 车道左边界线 Lane left boundary line

弃用 deprecated |
| right_border_line | [Polyline](#city-map-v2-Polyline) |  | **Deprecated.** 车道右边界线 Lane right boundary line

弃用 deprecated |
| predecessors | [LaneConnection](#city-map-v2-LaneConnection) | repeated | Lanes can drive / walk from 对于Junction内的车道至多1个前驱 For lanes within junction, there is at most 1 predecessor 对于LANE_TYPE_DRIVING，连接类型必为LANE_CONNECTION_TYPE_TAIL For LANE_TYPE_DRIVING, the connection type must be LANE_CONNECTION_TYPE_TAIL 对于LANE_TYPE_WALKING连接类型，两种都有可能 For LANE_TYPE_WALKING, both connection types are possible |
| successors | [LaneConnection](#city-map-v2-LaneConnection) | repeated | Lanes can drive / walk to 对于Junction内的车道至多1个后继 For lanes within junction, there is at most 1 successor 对于LANE_TYPE_DRIVING，连接类型必为LANE_CONNECTION_TYPE_HEAD For LANE_TYPE_DRIVING, the connection type must be LANE_CONNECTION_TYPE_HEAD 对于LANE_TYPE_WALKING连接类型，两种都有可能 For LANE_TYPE_WALKING, both connection types are possible |
| left_lane_ids | [int32](#int32) | repeated | 左侧相邻车道（按从近到远排列） Adjacent lanes on the left (arranged from nearest to far) |
| right_lane_ids | [int32](#int32) | repeated | 右侧相邻车道（按从近到远排列） Adjacent lanes on the right (arranged from nearest to far) |
| parent_id | [int32](#int32) |  | 所属的道路road id或路口junction id The road id or junction id it belongs to |
| overlaps | [LaneOverlap](#city-map-v2-LaneOverlap) | repeated | 车道之间的冲突点（仅在Junction内有效），保证按照self_s从小到大排序 Conflict points between lanes (valid only within Junction), guaranteed to be sorted from small to large according to self_s |
| aoi_ids | [int32](#int32) | repeated | 连接到该车道的所有AOI All AOIs connected to this lane |






<a name="city-map-v2-LaneConnection"></a>

### LaneConnection
车道连接信息
Lane connection information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 所连接的车道Lane的ID ID of the connected lane |
| type | [LaneConnectionType](#city-map-v2-LaneConnectionType) |  | 连接类型 Connection type |






<a name="city-map-v2-LaneOverlap"></a>

### LaneOverlap
两个lane的冲突区域
Conflict area between two lanes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| self | [city.geo.v2.LanePosition](#city-geo-v2-LanePosition) |  | 冲突点在本车道上的坐标 Coordinates of the conflict point on this lane |
| other | [city.geo.v2.LanePosition](#city-geo-v2-LanePosition) |  | 冲突点在冲突车道上的坐标 Coordinates of the conflict point on the conflicted lane |
| self_first | [bool](#bool) |  | 本车道是否有优先通行权 Whether this lane has priority |






<a name="city-map-v2-Map"></a>

### Map
地图，对应一个地图pb文件或一个地图mongodb collection
Map, corresponding to a map pb file or a map MongoDB collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Header](#city-map-v2-Header) |  |  |
| lanes | [Lane](#city-map-v2-Lane) | repeated |  |
| roads | [Road](#city-map-v2-Road) | repeated |  |
| junctions | [Junction](#city-map-v2-Junction) | repeated |  |
| aois | [Aoi](#city-map-v2-Aoi) | repeated |  |
| pois | [Poi](#city-map-v2-Poi) | repeated |  |
| sublines | [PublicTransportSubline](#city-map-v2-PublicTransportSubline) | repeated |  |






<a name="city-map-v2-NextRoadLane"></a>

### NextRoadLane



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| road_id | [int32](#int32) |  | 下一条路的id ID of the next road |
| lane_id_a | [int32](#int32) |  | 我们假定能去往对应道路的车道id范围是连续的，用[a,b]表示 We assume that the range of lane IDs of the next road is continuous, represented by [a, b] |
| lane_id_b | [int32](#int32) |  |  |






<a name="city-map-v2-NextRoadLanePlan"></a>

### NextRoadLanePlan



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| next_road_lanes | [NextRoadLane](#city-map-v2-NextRoadLane) | repeated | 记录去往目标next_road的可行lane集合 set of feasible lanes for going to the next_road |






<a name="city-map-v2-Poi"></a>

### Poi
Poi，用于描述地图上的兴趣点
Poi, describing points of interest on the map


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Poi id(从7_0000_0000开始) Poi id (starting from 7_0000_0000) |
| name | [string](#string) |  | Poi名称 Poi name |
| category | [string](#string) |  | Poi分类编码 Poi category code |
| position | [city.geo.v2.XYPosition](#city-geo-v2-XYPosition) |  | Poi原始位置 Poi original position |
| aoi_id | [int32](#int32) |  | Poi所属的Aoi Aoi to which the Poi belongs |
| capacity | [int32](#int32) | optional | Poi的容量（能同时容纳的人数），若无则表示无人数限制 The capacity of Poi (the number of people it can accommodate at the same time), if none, it means there is no limit on the number of people |
| functions | [string](#string) | repeated | Poi所能提供的功能 The functions the Poi can offer |






<a name="city-map-v2-Polyline"></a>

### Polyline
折线，用于定义车道等的形状
Polyline, used to define the shape of lanes, etc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [city.geo.v2.XYPosition](#city-geo-v2-XYPosition) | repeated | 折线上的点 Points of the polyline |






<a name="city-map-v2-PublicTransportSubline"></a>

### PublicTransportSubline
公共交通支线 描述依附于行车路网的公共交通线路
Public transport sub-lines, describe public transport routes attached to the road network


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 支线ID Subline ID |
| name | [string](#string) |  | 支线名字 Subline name |
| aoi_ids | [int32](#int32) | repeated | 该条支线沿途的所有车站AOI，按支线前进顺序排列 All stations along this subline, in the order of advancement of the subline.&#34; |
| station_connection_road_ids | [RoadIds](#city-map-v2-RoadIds) | repeated | 依次连接两个相邻车站的road ids Road IDs between two adjacent stations. |
| type | [SublineType](#city-map-v2-SublineType) |  | 支线类型 Type of subline |
| parent_name | [string](#string) |  | 所属线路名称 Name of the belonging line |
| schedules | [SublineSchedules](#city-map-v2-SublineSchedules) |  | 发车时刻表 departure schedule |
| taz_costs | [HeuristicTAZCost](#city-map-v2-HeuristicTAZCost) | repeated | 预计算公共交通权重 Precalculate public transport costs |






<a name="city-map-v2-Road"></a>

### Road
Road，用于描述道路
Road, describing roads


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 道路id（从2_0000_0000开始） Road ID (starting from 2_0000_0000) |
| name | [string](#string) |  | 道路名字 road name |
| lane_ids | [int32](#int32) | repeated | 属于该道路Road的所有车道/人行道等lane All lanes/sidewalks belonging to the road lane_id是按从最左侧车道到最右侧车道(从前进方向来看)的顺序给出的 lane_ids are given in order from the leftmost lane to the rightmost lane (viewed from the forward direction) |
| next_road_lane_plans | [NextRoadLanePlan](#city-map-v2-NextRoadLanePlan) | repeated | 对于包含动态车道的道路，需要通过这一项来指定所有的候选方案 For roads containing dynamic lanes, this is required to specify all candidates |






<a name="city-map-v2-RoadIds"></a>

### RoadIds



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| road_ids | [int32](#int32) | repeated |  |






<a name="city-map-v2-SublineSchedules"></a>

### SublineSchedules
发车时刻表
Subline departure schedule


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| departure_times | [double](#double) | repeated | 始发站发车时间 Departure time from the departure station |
| offset_times | [double](#double) | repeated | 到达沿途站点预计时间 Estimated time to arrive at stops along the way |





 


<a name="city-map-v2-AoiType"></a>

### AoiType
Aoi类型 Aoi Type

| Name | Number | Description |
| ---- | ------ | ----------- |
| AOI_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| AOI_TYPE_BUS_STATION | 1 | 公交站点 bus station |
| AOI_TYPE_OTHER | 2 | 其他 other |



<a name="city-map-v2-LandUseType"></a>

### LandUseType
土地利用类型，参照国标GB/T 21010—2007
Land use type, refer to the national standard GB/T 21010-2007
http://www.gscloud.cn/static/cases/%E3%80%8A%E5%9C%9F%E5%9C%B0%E5%88%A9%E7%94%A8%E7%8E%B0%E7%8A%B6%E5%88%86%E7%B1%BB%E3%80%8B%E5%9B%BD%E5%AE%B6%E6%A0%87%E5%87%86gb_t21010-2007(1).pdf

| Name | Number | Description |
| ---- | ------ | ----------- |
| LAND_USE_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| LAND_USE_TYPE_COMMERCIAL | 5 | 商服用地 commercial land |
| LAND_USE_TYPE_INDUSTRIAL | 6 | 工矿仓储用地 Industrial and storage land |
| LAND_USE_TYPE_RESIDENTIAL | 7 | 住宅用地 residential land |
| LAND_USE_TYPE_PUBLIC | 8 | 公共管理与公共服务用地 Public management and public service land |
| LAND_USE_TYPE_TRANSPORTATION | 10 | 交通运输用地 transportation land |
| LAND_USE_TYPE_OTHER | 12 | 其他土地 other land |



<a name="city-map-v2-LaneConnectionType"></a>

### LaneConnectionType
车道连接类型
Lane connection type

| Name | Number | Description |
| ---- | ------ | ----------- |
| LANE_CONNECTION_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| LANE_CONNECTION_TYPE_HEAD | 1 | 连接处为道路开头 The connection is at the lane head |
| LANE_CONNECTION_TYPE_TAIL | 2 | 连接处为道路结尾 The connection is at the lane tail |



<a name="city-map-v2-LaneTurn"></a>

### LaneTurn
车道转向
lane turn type

| Name | Number | Description |
| ---- | ------ | ----------- |
| LANE_TURN_UNSPECIFIED | 0 | 未指定 unspecified |
| LANE_TURN_STRAIGHT | 1 | 直行 go straight |
| LANE_TURN_LEFT | 2 | 左转 turn left |
| LANE_TURN_RIGHT | 3 | 右转 turn right |
| LANE_TURN_AROUND | 4 | 掉头 turn around |



<a name="city-map-v2-LaneType"></a>

### LaneType
车道类型
Lane type

| Name | Number | Description |
| ---- | ------ | ----------- |
| LANE_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| LANE_TYPE_DRIVING | 1 | 行车 driving |
| LANE_TYPE_WALKING | 2 | 步行 walking |
| LANE_TYPE_RAIL_TRANSIT | 3 | 轨道交通 rail transit |



<a name="city-map-v2-SublineType"></a>

### SublineType
支线类型
Type of subline

| Name | Number | Description |
| ---- | ------ | ----------- |
| SUBLINE_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| SUBLINE_TYPE_BUS | 1 | 公交类型支线 The subline is a bus line |
| SUBLINE_TYPE_SUBWAY | 2 | 地铁类型支线 The subline is a subway line |


 

 

 



<a name="city_elec_input_v1_input_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/elec/input/v1/input_service.proto



<a name="city-elec-input-v1-InitRequest"></a>

### InitRequest







<a name="city-elec-input-v1-InitResponse"></a>

### InitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | 模拟器gRPC监听地址 |
| control | [Control](#city-elec-input-v1-Control) |  |  |
| facilities | [Facilities](#city-elec-input-v1-Facilities) |  |  |
| map | [city.map.v2.Map](#city-map-v2-Map) |  |  |





 

 

 


<a name="city-elec-input-v1-InputService"></a>

### InputService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Init | [InitRequest](#city-elec-input-v1-InitRequest) | [InitResponse](#city-elec-input-v1-InitResponse) |  |

 



<a name="city_elec_interaction_v1_elec_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/elec/interaction/v1/elec_service.proto



<a name="city-elec-interaction-v1-GetEdgeStatusRequest"></a>

### GetEdgeStatusRequest







<a name="city-elec-interaction-v1-GetEdgeStatusResponse"></a>

### GetEdgeStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reason1 | [string](#string) | repeated |  |
| reason2 | [string](#string) | repeated |  |
| reason3 | [string](#string) | repeated |  |






<a name="city-elec-interaction-v1-GetNoPowerAOIRequest"></a>

### GetNoPowerAOIRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flag | [int32](#int32) |  |  |






<a name="city-elec-interaction-v1-GetNoPowerAOIResponse"></a>

### GetNoPowerAOIResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aoi | [int32](#int32) | repeated |  |






<a name="city-elec-interaction-v1-GetPowerRequest"></a>

### GetPowerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 设备id |






<a name="city-elec-interaction-v1-GetPowerResponse"></a>

### GetPowerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| power | [double](#double) |  | 功率 |






<a name="city-elec-interaction-v1-GetPowerStatusRequest"></a>

### GetPowerStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flag | [int32](#int32) |  |  |






<a name="city-elec-interaction-v1-GetPowerStatusResponse"></a>

### GetPowerStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| power_status | [GetPowerStatusResponse.PowerStatusEntry](#city-elec-interaction-v1-GetPowerStatusResponse-PowerStatusEntry) | repeated |  |






<a name="city-elec-interaction-v1-GetPowerStatusResponse-PowerStatusEntry"></a>

### GetPowerStatusResponse.PowerStatusEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [double](#double) |  |  |






<a name="city-elec-interaction-v1-GetRuinInfoRequest"></a>

### GetRuinInfoRequest







<a name="city-elec-interaction-v1-GetRuinInfoResponse"></a>

### GetRuinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| one | [RuinInfo](#city-elec-interaction-v1-RuinInfo) |  | 三级损伤信息 |
| two | [RuinInfo](#city-elec-interaction-v1-RuinInfo) |  |  |
| three | [RuinInfo](#city-elec-interaction-v1-RuinInfo) |  |  |






<a name="city-elec-interaction-v1-RuinInfo"></a>

### RuinInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num | [int32](#int32) |  | 损坏数量 |
| ratio | [double](#double) |  | 损坏占比 |






<a name="city-elec-interaction-v1-SetStatusRequest"></a>

### SetStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 设施id |
| status | [bool](#bool) |  | True 表示恢复，False表示摧毁 |






<a name="city-elec-interaction-v1-SetStatusResponse"></a>

### SetStatusResponse






 

 

 


<a name="city-elec-interaction-v1-ElecService"></a>

### ElecService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetStatus | [SetStatusRequest](#city-elec-interaction-v1-SetStatusRequest) | [SetStatusResponse](#city-elec-interaction-v1-SetStatusResponse) |  |
| GetPower | [GetPowerRequest](#city-elec-interaction-v1-GetPowerRequest) | [GetPowerResponse](#city-elec-interaction-v1-GetPowerResponse) |  |
| GetPowerStatus | [GetPowerStatusRequest](#city-elec-interaction-v1-GetPowerStatusRequest) | [GetPowerStatusResponse](#city-elec-interaction-v1-GetPowerStatusResponse) |  |
| GetNoPowerAOI | [GetNoPowerAOIRequest](#city-elec-interaction-v1-GetNoPowerAOIRequest) | [GetNoPowerAOIResponse](#city-elec-interaction-v1-GetNoPowerAOIResponse) |  |
| GetRuinInfo | [GetRuinInfoRequest](#city-elec-interaction-v1-GetRuinInfoRequest) | [GetRuinInfoResponse](#city-elec-interaction-v1-GetRuinInfoResponse) |  |
| GetEdgeStatus | [GetEdgeStatusRequest](#city-elec-interaction-v1-GetEdgeStatusRequest) | [GetEdgeStatusResponse](#city-elec-interaction-v1-GetEdgeStatusResponse) |  |

 



<a name="city_elec_output_v1_output-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/elec/output/v1/output.proto



<a name="city-elec-output-v1-Aoi"></a>

### Aoi



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| unsatisfied_demand_ratio | [double](#double) |  | 用电需求不满足比例（没电100%/有电0） |
| unsatisfied_demand_num | [int32](#int32) |  | 用电需求不满足人数 （没电就是aoi里的人都不满足） |
| demand | [double](#double) |  | 该aoi当前时刻的用电需求,单位为KW |
| supply | [double](#double) |  | 电力系统向该aoi供应的电力,单位为KW |






<a name="city-elec-output-v1-AveragePower"></a>

### AveragePower



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transformer_500 | [double](#double) |  | 当前各类变压器的平均承受功率，单位为MW |
| transformer_220 | [double](#double) |  |  |
| transformer_110 | [double](#double) |  |  |
| transformer_10 | [double](#double) |  |  |





 

 

 

 



<a name="city_elec_output_v1_output_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/elec/output/v1/output_service.proto



<a name="city-elec-output-v1-OutputRequest"></a>

### OutputRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| step | [int32](#int32) |  |  |
| ruined_ids | [int32](#int32) | repeated | 被破坏的节点ID |
| stopped_ids | [int32](#int32) | repeated | 由于其他节点被破坏导致停止工作的节点ID |
| aois | [Aoi](#city-elec-output-v1-Aoi) | repeated | AOI相关的数据 |
| events | [city.event.v1.Events](#city-event-v1-Events) |  | 事件数据 |
| resident_unsatisfied_ratio | [double](#double) |  | 居民用电需求不满足比例 |
| resident_demand | [double](#double) |  | 居民总用电需求,MW |
| aoi_unsatisfied_ratio | [double](#double) |  | aoi用电需求不满足比例 |
| aoi_unsatisfied_num | [int32](#int32) |  | 不满足用电的aoi数量,个数 |
| aoi_demand | [double](#double) |  | aoi总用电需求,MW |
| average_powers | [AveragePower](#city-elec-output-v1-AveragePower) |  | 各类变压器当前的平均承受功率，单位为MW |






<a name="city-elec-output-v1-OutputResponse"></a>

### OutputResponse






 

 

 


<a name="city-elec-output-v1-OutputService"></a>

### OutputService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Output | [OutputRequest](#city-elec-output-v1-OutputRequest) | [OutputResponse](#city-elec-output-v1-OutputResponse) |  |

 



<a name="city_event_v1_event_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/event/v1/event_service.proto



<a name="city-event-v1-PublishRequest"></a>

### PublishRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#city-event-v1-Event) |  |  |






<a name="city-event-v1-PublishResponse"></a>

### PublishResponse







<a name="city-event-v1-PullRequest"></a>

### PullRequest







<a name="city-event-v1-PullResponse"></a>

### PullResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [Event](#city-event-v1-Event) | repeated |  |





 

 

 


<a name="city-event-v1-EventService"></a>

### EventService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Publish | [PublishRequest](#city-event-v1-PublishRequest) | [PublishResponse](#city-event-v1-PublishResponse) | 发布事件 |
| Pull | [PullRequest](#city-event-v1-PullRequest) | [PullResponse](#city-event-v1-PullResponse) | 从事件中心拉取事件 |

 



<a name="city_event_v2_event-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/event/v2/event.proto



<a name="city-event-v2-Entity"></a>

### Entity
主语


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [EntityType](#city-event-v2-EntityType) |  | 实体类型 |
| id | [int32](#int32) |  | 实体ID |






<a name="city-event-v2-Event"></a>

### Event
模拟器中的事件
包含主题、主语（Who）、谓词（内容）、地点（Where）、时间（When）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| topic | [string](#string) |  | 主题 |
| id | [int32](#int32) | optional | ID |
| subject | [Entity](#city-event-v2-Entity) |  | 主语 |
| content | [string](#string) |  | 谓词 |
| position | [city.geo.v2.Position](#city-geo-v2-Position) |  | 地点 |
| t | [double](#double) |  | 时间 |





 


<a name="city-event-v2-EntityType"></a>

### EntityType
实体类型

| Name | Number | Description |
| ---- | ------ | ----------- |
| ENTITY_TYPE_UNSPECIFIED | 0 | 未指定 |
| ENTITY_TYPE_LANE | 1 | Lane |
| ENTITY_TYPE_ROAD | 2 | Road |
| ENTITY_TYPE_JUNCTION | 3 | Junction |
| ENTITY_TYPE_AOI | 4 | AOI |
| ENTITY_TYPE_POI | 5 | POI |
| ENTITY_TYPE_PERSON | 6 | 人 |
| ENTITY_TYPE_ORG | 7 | 组织 |


 

 

 



<a name="city_event_v2_event_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/event/v2/event_service.proto



<a name="city-event-v2-GetEventsByTopicRequest"></a>

### GetEventsByTopicRequest
按照topic查询事件请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| topic | [string](#string) |  | 主题 |






<a name="city-event-v2-GetEventsByTopicResponse"></a>

### GetEventsByTopicResponse
按照topic查询事件响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [Event](#city-event-v2-Event) | repeated | 事件列表 |






<a name="city-event-v2-ResolveEventsRequest"></a>

### ResolveEventsRequest
确认事件已被处理请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [Event](#city-event-v2-Event) | repeated | 事件列表 |






<a name="city-event-v2-ResolveEventsResponse"></a>

### ResolveEventsResponse
确认事件已被处理响应





 

 

 


<a name="city-event-v2-EventService"></a>

### EventService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetEventsByTopic | [GetEventsByTopicRequest](#city-event-v2-GetEventsByTopicRequest) | [GetEventsByTopicResponse](#city-event-v2-GetEventsByTopicResponse) | 按照topic查询事件 |
| ResolveEvents | [ResolveEventsRequest](#city-event-v2-ResolveEventsRequest) | [ResolveEventsResponse](#city-event-v2-ResolveEventsResponse) | 确认事件已被处理 |

 



<a name="city_person_v2_motion-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v2/motion.proto



<a name="city-person-v2-PersonMotion"></a>

### PersonMotion
Person（人）的运动状态
Person&#39;s motion state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | ID |
| status | [Status](#city-person-v2-Status) |  | 状态 status |
| position | [city.geo.v2.Position](#city-geo-v2-Position) |  | 位置（包含逻辑位置、XY位置、经纬度位置） Position (including logical position, XY position, longitude and latitude position) |
| v | [double](#double) |  | speed |
| direction | [double](#double) |  | 方向角（atan2计算得到的弧度） Direction angle (radians calculated by atan2) |
| activity | [string](#string) |  | 活动描述 activity descriptions |
| l | [double](#double) |  | 长度 length |





 


<a name="city-person-v2-Status"></a>

### Status
Person（人）的运行时状态
Person&#39;s runtime state

| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | 未指定 unspecified |
| STATUS_SLEEP | 1 | 没有移动行为 no mobility behaviors |
| STATUS_DRIVING | 2 | 开车 driving |
| STATUS_WALKING | 3 | 步行 walking |
| STATUS_CROWD | 4 | 室内行人 indoor pedestrian |
| STATUS_PASSENGER | 5 | 乘客 vehicle passenger |
| STATUS_WAIT_ROUTE | 6 | 等待路径规划 wait for path routing |
| STATUS_WAIT_BUS | 7 | 等待公交车 wait for bus coming |
| STATUS_RAIL_TRANSIT | 8 | 操控轨道交通 operating rail transit |
| STATUS_WAIT_TAXI | 9 | 等待出租车 wait for taxi |


 

 

 



<a name="city_map_v2_aoi_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/map/v2/aoi_service.proto



<a name="city-map-v2-AoiState"></a>

### AoiState
AOI状态
AOI state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | AOI ID |
| persons | [city.person.v2.PersonMotion](#city-person-v2-PersonMotion) | repeated | AOI内的人 Persons in AOI |






<a name="city-map-v2-GetAoiRequest"></a>

### GetAoiRequest
获取AOI信息请求
Request for getting AOI information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aoi_ids | [int32](#int32) | repeated | 指定AOI ID列表，如果为空，则返回所有AOI信息 List of targeted AOI IDs, if empty, returns all information of AOIs |






<a name="city-map-v2-GetAoiResponse"></a>

### GetAoiResponse
获取AOI信息响应
Response for getting AOI information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| states | [AoiState](#city-map-v2-AoiState) | repeated | AOI信息列表 Lis of AOIs information |





 

 

 


<a name="city-map-v2-AoiService"></a>

### AoiService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAoi | [GetAoiRequest](#city-map-v2-GetAoiRequest) | [GetAoiResponse](#city-map-v2-GetAoiResponse) | 获取AOI信息 Get AOI information |

 



<a name="city_map_v2_lane_state-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/map/v2/lane_state.proto



<a name="city-map-v2-LaneState"></a>

### LaneState
Lane状态
Lane state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Lane ID |
| persons | [city.person.v2.PersonMotion](#city-person-v2-PersonMotion) | repeated | Lane上的人/车 person/vehicle on the lane |
| avg_v | [double](#double) |  | 平均速度（m/s） average speed (m/s) |
| restriction | [bool](#bool) |  | 是否限行 whether restricted |
| light_state | [LightState](#city-map-v2-LightState) |  | 交通灯状态 traffic light state |





 

 

 

 



<a name="city_map_v2_lane_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/map/v2/lane_service.proto



<a name="city-map-v2-GetLaneByLongLatBBoxRequest"></a>

### GetLaneByLongLatBBoxRequest
获取特定区域内的Lane的信息请求
Request for getting lane information in a specific region


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bbox | [city.geo.v2.LongLatBBox](#city-geo-v2-LongLatBBox) |  | 经纬度范围 latitude and longitude bounding box |
| exclude_person | [bool](#bool) |  | 是否要排除车道上的人的信息 Whether to exclude information of person on the lane |






<a name="city-map-v2-GetLaneByLongLatBBoxResponse"></a>

### GetLaneByLongLatBBoxResponse
获取特定区域内的Lane的信息响应
Response of getting lane information in a specific region


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| states | [LaneState](#city-map-v2-LaneState) | repeated | Lane的信息 Lane information |






<a name="city-map-v2-GetLaneRequest"></a>

### GetLaneRequest
获取Lane的信息请求
Request for getting lane information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lane_ids | [int32](#int32) | repeated | 指定的Lane id列表，如果为空，则返回所有Lane的信息 List of targeted lane IDs, if empty, returns all information of lanes |
| exclude_person | [bool](#bool) |  | 是否要排除车道上的人的信息 Whether to exclude information of person on the lane |






<a name="city-map-v2-GetLaneResponse"></a>

### GetLaneResponse
获取Lane的信息响应
Response of getting lane information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| states | [LaneState](#city-map-v2-LaneState) | repeated | Lane的信息 Lane information |






<a name="city-map-v2-SetLaneMaxVRequest"></a>

### SetLaneMaxVRequest
设置Lane的最大速度（限速）请求
Request for setting lane&#39;s maximum speed (speed limit)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lane_id | [int32](#int32) |  | Lane id |
| max_v | [double](#double) |  | 最大速度（限速），单位：m/s Maximum speed (speed limit), unit: m/s |






<a name="city-map-v2-SetLaneMaxVResponse"></a>

### SetLaneMaxVResponse
设置Lane的最大速度（限速）响应
Response of setting lane&#39;s maximum speed (speed limit)






<a name="city-map-v2-SetLaneRestrictionRequest"></a>

### SetLaneRestrictionRequest
设置Lane限行请求
Request for setting lane&#39;s traffic restriction


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lane_id | [int32](#int32) |  | Lane id |
| restriction | [bool](#bool) |  | 限行 Traffic restriction |






<a name="city-map-v2-SetLaneRestrictionResponse"></a>

### SetLaneRestrictionResponse
设置Lane限行响应
Response of setting lane&#39;s traffic restriction





 

 

 


<a name="city-map-v2-LaneService"></a>

### LaneService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetLaneMaxV | [SetLaneMaxVRequest](#city-map-v2-SetLaneMaxVRequest) | [SetLaneMaxVResponse](#city-map-v2-SetLaneMaxVResponse) | 设置Lane的最大速度（限速） Set Lane&#39;s maximum speed (speed limit) |
| SetLaneRestriction | [SetLaneRestrictionRequest](#city-map-v2-SetLaneRestrictionRequest) | [SetLaneRestrictionResponse](#city-map-v2-SetLaneRestrictionResponse) | 设置Lane限行 Set Lane&#39;s traffic restriction |
| GetLane | [GetLaneRequest](#city-map-v2-GetLaneRequest) | [GetLaneResponse](#city-map-v2-GetLaneResponse) | 获取Lane的信息 Get Lane information |
| GetLaneByLongLatBBox | [GetLaneByLongLatBBoxRequest](#city-map-v2-GetLaneByLongLatBBoxRequest) | [GetLaneByLongLatBBoxResponse](#city-map-v2-GetLaneByLongLatBBoxResponse) | 获取特定区域内的Lane的信息 Get Lane information in a specific region |

 



<a name="city_map_v2_road_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/map/v2/road_service.proto



<a name="city-map-v2-GetEventsRequest"></a>

### GetEventsRequest







<a name="city-map-v2-GetEventsResponse"></a>

### GetEventsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [city.event.v1.Events](#city-event-v1-Events) |  |  |






<a name="city-map-v2-GetRoadRequest"></a>

### GetRoadRequest
查询道路信息请求
Request for getting road information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| road_ids | [int32](#int32) | repeated | 指定查询的道路ID列表，为空代表查询所有道路 List of targeted road IDs. If empty, it means querying all roads. |
| exclude_lane | [bool](#bool) |  | 是否要排除车道信息 Whether to exclude lane information |
| exclude_person | [bool](#bool) |  | 是否要排除车道上的人的信息（仅在包含车道信息时有效） Whether to exclude information about person in the lane (only valid when lane information is included) |






<a name="city-map-v2-GetRoadResponse"></a>

### GetRoadResponse
查询道路信息响应
Response of getting road information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| states | [RoadState](#city-map-v2-RoadState) | repeated | 道路信息列表 List of road information |






<a name="city-map-v2-GetRuinInfoRequest"></a>

### GetRuinInfoRequest







<a name="city-map-v2-GetRuinInfoResponse"></a>

### GetRuinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| one | [RuinInfo](#city-map-v2-RuinInfo) |  | 三级损伤信息 Three-level ruin information |
| two | [RuinInfo](#city-map-v2-RuinInfo) |  |  |
| three | [RuinInfo](#city-map-v2-RuinInfo) |  |  |






<a name="city-map-v2-RoadState"></a>

### RoadState
道路状态
road state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 道路ID road ID |
| avg_v | [double](#double) |  | 道路平均速度（m/s） road average speed (m/s) |
| level | [RoadLevel](#city-map-v2-RoadLevel) |  | 道路拥堵情况 road congestion level |
| reason | [InterruptionReason](#city-map-v2-InterruptionReason) |  | 道路中断原因 road interruption reason |
| lanes | [LaneState](#city-map-v2-LaneState) | repeated | 车道情况 lane state |






<a name="city-map-v2-RuinInfo"></a>

### RuinInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num | [int32](#int32) |  | 损坏数量。Ruined number |
| ratio | [double](#double) |  | 损坏占比。Ruined ratio |





 


<a name="city-map-v2-InterruptionReason"></a>

### InterruptionReason
道路中断原因
road interruption reason

| Name | Number | Description |
| ---- | ------ | ----------- |
| INTERRUPTION_REASON_UNSPECIFIED | 0 |  |
| INTERRUPTION_REASON_RUINED | 1 |  |
| INTERRUPTION_REASON_CASCADE | 2 |  |
| INTERRUPTION_REASON_CONGESTION | 3 |  |



<a name="city-map-v2-RoadLevel"></a>

### RoadLevel
道路拥堵情况
road congestion level

| Name | Number | Description |
| ---- | ------ | ----------- |
| ROAD_LEVEL_UNSPECIFIED | 0 | 未指定 unspecified |
| ROAD_LEVEL_CLEAR | 1 | 畅通 clear |
| ROAD_LEVEL_LIGHT_LOAD | 2 | 轻度拥堵 light load |
| ROAD_LEVEL_MEDIUM_LOAD | 3 | 中度拥堵 medium load |
| ROAD_LEVEL_HEAVY_LOAD | 4 | 重度拥堵 heavy load |
| ROAD_LEVEL_OVERLOAD | 5 | 极端拥堵 overload |
| ROAD_LEVEL_RESTRICTED | 6 | 限行 restricted |


 

 


<a name="city-map-v2-RoadService"></a>

### RoadService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetRoad | [GetRoadRequest](#city-map-v2-GetRoadRequest) | [GetRoadResponse](#city-map-v2-GetRoadResponse) | 查询道路信息 Get road information |
| GetRuinInfo | [GetRuinInfoRequest](#city-map-v2-GetRuinInfoRequest) | [GetRuinInfoResponse](#city-map-v2-GetRuinInfoResponse) |  |
| GetEvents | [GetEventsRequest](#city-map-v2-GetEventsRequest) | [GetEventsResponse](#city-map-v2-GetEventsResponse) |  |

 



<a name="city_map_v2_traffic_light_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/map/v2/traffic_light_service.proto



<a name="city-map-v2-GetTrafficLightRequest"></a>

### GetTrafficLightRequest
获取路口的红绿灯信息请求
Reqeust for getting traffic light information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| junction_id | [int32](#int32) |  | 信号等相关的接口精确到junction The interfaces related to signals are precise to junction |






<a name="city-map-v2-GetTrafficLightResponse"></a>

### GetTrafficLightResponse
获取路口的红绿灯信息响应
Response of getting traffic light information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| traffic_light | [TrafficLight](#city-map-v2-TrafficLight) |  | 当前路口处的红绿灯 The traffic light at the junction |
| phase_index | [int32](#int32) |  | 表示当前路口处的红绿灯处于哪一个相位 Which phase the traffic light is currently in |
| time_remaining | [double](#double) |  | 当前相位的剩余时间 The remaining time of the current phase |






<a name="city-map-v2-SetTrafficLightPhaseRequest"></a>

### SetTrafficLightPhaseRequest
设置路口的红绿灯相位请求
Request for setting traffic light phase


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| junction_id | [int32](#int32) |  | 需要改变相位的路口编号 The target junction ID |
| phase_index | [int32](#int32) |  | 指定当前路口红绿灯的相位 Specify the traffic light phase |
| time_remaining | [double](#double) |  | 当前相位的剩余时间 The remaining time of the current phase |






<a name="city-map-v2-SetTrafficLightPhaseResponse"></a>

### SetTrafficLightPhaseResponse
设置路口的红绿灯相位响应
Response of setting traffic light phase






<a name="city-map-v2-SetTrafficLightRequest"></a>

### SetTrafficLightRequest
设置路口的红绿灯信息请求
Request for setting traffic light information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| traffic_light | [TrafficLight](#city-map-v2-TrafficLight) |  | 需要改变的红绿灯（含路口编号） The target traffic light (including junction ID) |
| phase_index | [int32](#int32) |  | 指定当前路口处的红绿灯的相位 Specify the phase of the traffic light |
| time_remaining | [double](#double) |  | 当前相位的剩余时间 The remaining time of the current phase |






<a name="city-map-v2-SetTrafficLightResponse"></a>

### SetTrafficLightResponse
设置路口的红绿灯信息响应
Response of setting traffic light information






<a name="city-map-v2-SetTrafficLightStatusRequest"></a>

### SetTrafficLightStatusRequest
设置路口的红绿灯状态请求
Request for setting traffic light status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| junction_id | [int32](#int32) |  | 需要改变状态的路口编号 The target junction ID |
| ok | [bool](#bool) |  | 当前路口红绿灯状态，true为通，false为断 The current traffic light status at the junction, true is on, false is off |






<a name="city-map-v2-SetTrafficLightStatusResponse"></a>

### SetTrafficLightStatusResponse
设置路口的红绿灯状态响应
Response of setting traffic light status





 

 

 


<a name="city-map-v2-TrafficLightService"></a>

### TrafficLightService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTrafficLight | [GetTrafficLightRequest](#city-map-v2-GetTrafficLightRequest) | [GetTrafficLightResponse](#city-map-v2-GetTrafficLightResponse) | 获取路口的红绿灯信息 Get traffic light information |
| SetTrafficLight | [SetTrafficLightRequest](#city-map-v2-SetTrafficLightRequest) | [SetTrafficLightResponse](#city-map-v2-SetTrafficLightResponse) | 设置路口的红绿灯信息 Set traffic light information |
| SetTrafficLightPhase | [SetTrafficLightPhaseRequest](#city-map-v2-SetTrafficLightPhaseRequest) | [SetTrafficLightPhaseResponse](#city-map-v2-SetTrafficLightPhaseResponse) | 设置路口的红绿灯相位 Set traffic light phase |
| SetTrafficLightStatus | [SetTrafficLightStatusRequest](#city-map-v2-SetTrafficLightStatusRequest) | [SetTrafficLightStatusResponse](#city-map-v2-SetTrafficLightStatusResponse) | 设置路口的红绿灯状态 Set traffic light status |

 



<a name="city_pause_v1_pause_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/pause/v1/pause_service.proto



<a name="city-pause-v1-PauseRequest"></a>

### PauseRequest
暂停程序请求
Pause program request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional | 组件名，用于暂停某一特定组件的运行，不设置时暂停整个程序 Component name, used to pause a specific component, pause the entire program when not set |






<a name="city-pause-v1-PauseResponse"></a>

### PauseResponse
暂停程序响应
Pause program response






<a name="city-pause-v1-ResumeRequest"></a>

### ResumeRequest
恢复程序请求
Resume program request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional | 组件名，用于恢复某一特定组件的运行，不设置时恢复整个程序 Component name, used to resume a specific component, resume the entire program when not set |






<a name="city-pause-v1-ResumeResponse"></a>

### ResumeResponse
恢复程序响应
Resume program response





 

 

 


<a name="city-pause-v1-PauseService"></a>

### PauseService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Pause | [PauseRequest](#city-pause-v1-PauseRequest) | [PauseResponse](#city-pause-v1-PauseResponse) | 暂停程序 Pause the program |
| Resume | [ResumeRequest](#city-pause-v1-ResumeRequest) | [ResumeResponse](#city-pause-v1-ResumeResponse) | 恢复程序 Resume the program |

 



<a name="city_person_v1_motion-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v1/motion.proto



<a name="city-person-v1-PersonMotion"></a>

### PersonMotion
Person（人）的运动状态
Person&#39;s motion state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | ID |
| status | [Status](#city-person-v1-Status) |  | 状态 status |
| position | [city.geo.v2.Position](#city-geo-v2-Position) |  | 位置（包含逻辑位置、XY位置、经纬度位置） Position (including logical position, XY position, longitude and latitude position) |
| v | [double](#double) |  | speed |
| direction | [double](#double) |  | 方向角（atan2计算得到的弧度） Direction angle (radians calculated by atan2) |
| activity | [string](#string) |  | 活动描述 activity descriptions |
| l | [double](#double) |  | 长度 length |





 


<a name="city-person-v1-Status"></a>

### Status
Person（人）的运行时状态
Person&#39;s runtime state

| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | 未指定 unspecified |
| STATUS_SLEEP | 1 | 没有移动行为 no mobility behaviors |
| STATUS_DRIVING | 2 | 开车 driving |
| STATUS_WALKING | 3 | 步行 walking |
| STATUS_CROWD | 4 | 室内行人 indoor pedestrian |
| STATUS_PASSENGER | 5 | 乘客 vehicle passenger |
| STATUS_WAIT_ROUTE | 6 | 等待路径规划 wait for path routing |
| STATUS_WAIT_BUS | 7 | 等待公交车 wait for bus coming |
| STATUS_RAIL_TRANSIT | 8 | 操控轨道交通 operating rail transit |


 

 

 



<a name="city_routing_v2_routing-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/routing/v2/routing.proto



<a name="city-routing-v2-BusJourneyBody"></a>

### BusJourneyBody



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transfers | [TransferSegment](#city-routing-v2-TransferSegment) | repeated |  |
| eta | [double](#double) |  | 从起点到终点预计的时间(estimation time of arrival) estimation time of arrival |






<a name="city-routing-v2-DrivingJourneyBody"></a>

### DrivingJourneyBody
驾车出行方式的路径规划结果
Routing results for driving journey


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| road_ids | [int32](#int32) | repeated | 从起点到终点的道路序列 Road sequence from origin to destination 采用道路序列的原因是主动变道对车道级的导航需要频繁修改 The reason for using road sequences is that active lane changing requires frequent modifications to lane level navigation 优先使用road_ids，如果road_ids为空，则使用route（也可以直接忽略route） Prioritize using road_ids. If road_ids is empty, use route (or simply ignore route) |
| eta | [double](#double) |  | 从起点到终点预计的时间(estimation time of arrival) estimation time of arrival |






<a name="city-routing-v2-Journey"></a>

### Journey
路径规划结果的一部分，含且仅含采用一种交通出行方式的完整出行序列
Part of the routing results, including a complete travel sequence using exactly one traveling mode


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [JourneyType](#city-routing-v2-JourneyType) |  | 出行方式 journey traveling mode |
| driving | [DrivingJourneyBody](#city-routing-v2-DrivingJourneyBody) | optional | 驾车 Routing results for driving journey |
| walking | [WalkingJourneyBody](#city-routing-v2-WalkingJourneyBody) | optional | 步行 Routing results of walking journey |
| by_bus | [BusJourneyBody](#city-routing-v2-BusJourneyBody) | optional | 公交 Routing results of bus journey |






<a name="city-routing-v2-RoadStatus"></a>

### RoadStatus
预计算路况信息
Pre calculate road condition information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 车道ID Lane ID |
| speed | [double](#double) | repeated | 车道在各个时间片（每个5min）的速度 The speed of the lane at each time slot (5 minutes each) |






<a name="city-routing-v2-RoadStatuses"></a>

### RoadStatuses
预计算道路路况信息集合，对应一个预计算道路况信息pb文件或一个预计算路况信息mongodb collection
Pre calculated road condition information set, corresponding to a pre calculated road condition information PB file or a pre calculated road condition information mongodb collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| road_statuses | [RoadStatus](#city-routing-v2-RoadStatus) | repeated |  |






<a name="city-routing-v2-TransferSegment"></a>

### TransferSegment
message BusJourneyBody {
  int32 line_id = 1;
  int32 start_station_id = 2;
  int32 end_station_id = 3;
}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subline_id | [int32](#int32) |  |  |
| start_station_id | [int32](#int32) |  |  |
| end_station_id | [int32](#int32) |  |  |






<a name="city-routing-v2-WalkingJourneyBody"></a>

### WalkingJourneyBody
步行出行方式的路径规划结果
Routing results of walking journey


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| route | [WalkingRouteSegment](#city-routing-v2-WalkingRouteSegment) | repeated | 从起点到终点的（Lane&#43;方向）序列 The (Lane&#43;direction) sequence from the origin to destination |
| eta | [double](#double) |  | 从起点到终点预计的时间(estimation time of arrival) estimation time of arrival |






<a name="city-routing-v2-WalkingRouteSegment"></a>

### WalkingRouteSegment
步行出行方式的路径规划结果中的一段
A segment in the routing results of walking journey


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lane_id | [int32](#int32) |  | Lane ID |
| moving_direction | [MovingDirection](#city-routing-v2-MovingDirection) |  | 移动方向 moving direction |





 


<a name="city-routing-v2-JourneyType"></a>

### JourneyType
移动方式
travelling mode
Journey用以描述采用一种特定交通方式从一点出发到达另一点的路径。
Journey is used to describe the path from one point to another using one specific travelling mode
一般来说，多个Journey是一个Trip的“实现”。
Generally, multiple Journeys are used to &#34;implement&#34; a Trip
例如：Trip(从清华乘地铁到天安门):
For example: Trip (taking the subway from Tsinghua to Tiananmen Square):
Journey(步行到地铁站)-&gt;Journey(地铁)-&gt;Journey(步行到天安门)
Journey (walking to subway station) -&gt; Journey (subway) -&gt; Journey (walking to Tiananmen Square)

| Name | Number | Description |
| ---- | ------ | ----------- |
| JOURNEY_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| JOURNEY_TYPE_DRIVING | 1 | 驾车 driving |
| JOURNEY_TYPE_WALKING | 2 | 步行 walking |
| JOURNEY_TYPE_BY_BUS | 3 | 公交 taking bus |
| JOURNEY_TYPE_BY_TAXI | 4 | 出租车 taking taxi |



<a name="city-routing-v2-MovingDirection"></a>

### MovingDirection
步行移动方向
Walking direction
行人前进的方向与Lane的正方向（s增大的方向）的关系
The relationship between the direction of pedestrian movement and the positive direction of Lane (the direction where s increases)

| Name | Number | Description |
| ---- | ------ | ----------- |
| MOVING_DIRECTION_UNSPECIFIED | 0 | 未指定 unspecified |
| MOVING_DIRECTION_FORWARD | 1 | 与正方向同向 In the same direction as the positive lane direction |
| MOVING_DIRECTION_BACKWARD | 2 | 与正方向反向 In the opposite direction as the positive lane direction |



<a name="city-routing-v2-RouteType"></a>

### RouteType
导航请求类型
routing type

| Name | Number | Description |
| ---- | ------ | ----------- |
| ROUTE_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| ROUTE_TYPE_DRIVING | 1 | 驾车 driving |
| ROUTE_TYPE_WALKING | 2 | 步行 walking |
| ROUTE_TYPE_BUS | 3 | 公交 by bus |
| ROUTE_TYPE_SUBWAY | 4 | 地铁 by subway |
| ROUTE_TYPE_BUS_SUBWAY | 5 | 地铁&#43;公交，包含两者的换乘 both bus and subway are available, including multimodal transfers |
| ROUTE_TYPE_TAXI | 6 | 出租车 by taxi |


 

 

 



<a name="city_trip_v2_trip-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/trip/v2/trip.proto



<a name="city-trip-v2-Schedule"></a>

### Schedule
时刻表
Schedule
关于出发时间的说明如下：
The explanation about the departure time is as follows:
1. Schedule的开始时刻是 departure_time 或者 参考时刻&#43;wait_time，
1. The start time of the Schedule is either departure_time or reference time&#43;wait_time,
   参考时刻定义为上一Schedule的结束时刻(即它最后一个Trip的结束时刻)，
   The reference time is defined as the end time of the previous Schedule (i.e. the end time of its last Trip),
   或者当它为第一个Schedule时定义为Person更新Schedule后的首次Update
   Alternatively, when it is the first Schedule, it can be defined as the first Update time after Person updates the Schedule
   时刻(当有准确时间要求时建议直接指定departure_time)
   (it is recommended to specify departuretime directly when there is an accurate time requirement)
2. Trip的开始时刻是 departure_time 或者 参考时刻&#43;wait_time，参考
2. The start time of the Trip is either departure_time or reference time&#43;wait_time,
   时刻定义为上一Trip的结束时刻，或者当它为第一个Trip时定义为所属的
   The reference time is defined as the end time of the previous Trip, or when it is the first Trip,
   Schedule的开始时刻
   it is defined as the start time of the Schedule to which it belongs
3. Person的实际运行时刻取决于Trip的开始时刻，例如它的首次运行是第一
3. The actual running time of a Person depends on the start time of the Trip,
   个Schedule中第一个Trip的开始时刻
   for example, its first run is the start time of the first Trip in the first Schedule
FAQ
Q1: 同时指定Schedule和第一个Trip的departure_time会怎样？
Q1: What would happen if both the Schedule and the departuretime of the first Trip were specified simultaneously?
A1: 参照(2)，只看Trip的departure_time
A1: Referring to (2), only depend on the departuretime of Trip
Q2: Schedule和第一个Trip同时指定wait_time=10会怎样？
Q2: What would happen if both the Schedule and the first Trip were specified with wait_time=10 at the same time?
A2: 参照(2)，等待时间为10&#43;10=20
A2: Referring to (2), the waiting time is 10&#43;10=20


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trips | [Trip](#city-trip-v2-Trip) | repeated | 出行列表 List of trips |
| loop_count | [int32](#int32) |  | trips的执行次数，0表示无限循环，大于0表示执行几次 The number of times trips are executed, where 0 represents infinite loops and greater than 0 represents how many times they are executed |
| departure_time | [double](#double) | optional | 期望的出发时间（单位: 秒） Expected departure time (in seconds) |
| wait_time | [double](#double) | optional | 期望的等待时间（单位：秒），如果departure_time为空则wait_time默认为0 Expected waiting time (in seconds), if departure_time is empty, wait_time defaults to 0 |






<a name="city-trip-v2-Trip"></a>

### Trip
出行
Trip


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [TripMode](#city-trip-v2-TripMode) |  | 出行方式 trip mode |
| end | [city.geo.v2.Position](#city-geo-v2-Position) |  | 目的地，如果目的地是AOI且指定了XYPosition，则以XYPosition为室内步行的终点 destination, if the destination is AOI and XYPosition is specified, XYPosition is the end point of indoor walking |
| departure_time | [double](#double) | optional | 期望的出发时间（单位: 秒） Expected departure time (in seconds) |
| wait_time | [double](#double) | optional | 期望的等待时间（单位：秒），如果departure_time为空则wait_time默认为0 The expected waiting time (in seconds), if departure_time is empty, wait_time defaults to 0 |
| arrival_time | [double](#double) | optional | 期望的到达时间（单位: 秒） Expected arrival time (in seconds) |
| activity | [string](#string) | optional | 本次出行目的地的活动名 The activity name of the destination for this trip |
| model | [string](#string) | optional | 本次出行对应的可视化模型（覆盖Person Attribute中的默认模型） The visual model corresponding to this trip (overriding the default model in Person Attribute) |
| routes | [city.routing.v2.Journey](#city-routing-v2-Journey) | repeated | 预计算的导航结果 Pre calculated routing results |
| trip_stops | [TripStop](#city-trip-v2-TripStop) | repeated | 本次出行的所有停靠点 stop points of this trip |






<a name="city-trip-v2-TripStop"></a>

### TripStop
停靠点
stop points of person


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aoi_position | [city.geo.v2.AoiPosition](#city-geo-v2-AoiPosition) | optional | 停车点AOI坐标（可选） Parking position coordinates AOI (optional) |
| lane_position | [city.geo.v2.LanePosition](#city-geo-v2-LanePosition) |  | 停车点Lane&#43;S坐标（必须提供） Parking position coordinates Lane&#43;S (must be provided) |
| duration | [double](#double) |  | 停车持续时间 Parking duration time (s) |





 


<a name="city-trip-v2-TripMode"></a>

### TripMode
出行偏好
Preferred trip travel mode

| Name | Number | Description |
| ---- | ------ | ----------- |
| TRIP_MODE_UNSPECIFIED | 0 | 未指定出行方式 unspecified |
| TRIP_MODE_WALK_ONLY | 1 | 仅步行 walking only |
| TRIP_MODE_DRIVE_ONLY | 2 | 仅开车 driving only |
| TRIP_MODE_BUS_WALK | 4 | 乘坐公交车（含站点间步行换乘） taking bus with walking to transit bus line between stations |
| TRIP_MODE_BIKE_WALK | 5 | 当有自行车时选择骑自行车，否则步行 Riding bikes if avaible, otherwise walking |
| TRIP_MODE_TAXI | 6 | 乘出租车 taking a taxi |
| TRIP_MODE_SUBWAY_WALK | 7 | 乘坐地铁（含站点间步行换乘） taking subway with walking to transit bus line between stations |
| TRIP_MODE_BUS_SUBWAY_WALK | 8 | 乘坐公交车（含站点间步行换乘） taking bus and subway with walking to transit bus line between stations |


 

 

 



<a name="city_person_v1_person-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v1/person.proto



<a name="city-person-v1-BikeAttribute"></a>

### BikeAttribute
自行车附加属性
Bike additional attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| speed | [double](#double) |  | 单位: m/s speed: m/s |
| model | [string](#string) | optional | 自行车模型标签 Bike model tag |






<a name="city-person-v1-BusAttribute"></a>

### BusAttribute
公交车附加属性
Bus additional attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subline_id | [int32](#int32) |  | 公交线路ID bus line ID |
| capacity | [int32](#int32) |  | 公交车容量 bus capacity |
| model | [string](#string) | optional | 公交车模型标签 bus model tag |
| type | [BusType](#city-person-v1-BusType) |  | 公交车类型 type of bus |






<a name="city-person-v1-PedestrianAttribute"></a>

### PedestrianAttribute
行人附加属性
Pedestrian additional attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| speed | [double](#double) |  | 单位: m/s speed: m/s |
| model | [string](#string) | optional | 行人模型标签 Pedestrian model tag |






<a name="city-person-v1-Person"></a>

### Person
智能体
agent


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 智能体ID agent ID |
| attribute | [PersonAttribute](#city-person-v1-PersonAttribute) |  | 参数 attribute |
| home | [city.geo.v2.Position](#city-geo-v2-Position) |  | 初始位置 initial position |
| schedules | [city.trip.v2.Schedule](#city-trip-v2-Schedule) | repeated | 初始日程 initial schedules |
| vehicle_attribute | [VehicleAttribute](#city-person-v1-VehicleAttribute) | optional | 车辆附加属性 vehicle addtional attribute |
| bus_attribute | [BusAttribute](#city-person-v1-BusAttribute) | optional | 公交车附加属性 bus additional attribute |
| pedestrian_attribute | [PedestrianAttribute](#city-person-v1-PedestrianAttribute) | optional | 行人附加属性 pedestrian additional attribute |
| bike_attribute | [BikeAttribute](#city-person-v1-BikeAttribute) | optional | 自行车附加属性 bike addition attribute |
| labels | [Person.LabelsEntry](#city-person-v1-Person-LabelsEntry) | repeated | [可空] 额外的标签（例如：抢修车类型-&gt;电网） [can be empty] additional tags (e.g. repair vehicle type -&gt; power grid) |
| profile | [PersonProfile](#city-person-v1-PersonProfile) | optional | [可空] 智能体简介 [can be empty] agent profile |
| work | [city.geo.v2.Position](#city-geo-v2-Position) | optional | 工作地位置 work position |
| output_when_sleep | [bool](#bool) | optional | 是否在SLEEP状态下也输出可视化（仅限车辆） Whether to output visualization in the SLEEP state (vehicles only) |






<a name="city-person-v1-Person-LabelsEntry"></a>

### Person.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="city-person-v1-PersonAttribute"></a>

### PersonAttribute
智能体属性（通用）
Agent properties (general)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| length | [double](#double) |  | 单位: m，长度 length: m |
| width | [double](#double) |  | 单位: m，宽度 width: m |
| max_speed | [double](#double) |  | 单位: m/s max speed: m/s |
| max_acceleration | [double](#double) |  | 单位: m/s^2, 最大加速度（正值） max accelaration: m/s^2 (positive value) |
| max_braking_acceleration | [double](#double) |  | 单位: m/s^2, 最大减速度（负值） max deceleration: m/s^2 (negative value) |
| usual_acceleration | [double](#double) |  | 单位: m/s^2, 一般加速度（正值），要求小于最大加速度 usual acceleration: m/s^2 (positive value), required to be less than the max acceleration |
| usual_braking_acceleration | [double](#double) |  | 单位: m/s^2, 一般减速度（负值），要求大于最大减速度 usual deceleration: m/s^2 (negative value), required to be greater than the max deceleration |






<a name="city-person-v1-PersonProfile"></a>

### PersonProfile
智能体简介
agent profile


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| age | [int32](#int32) |  | 年龄 age |
| education | [Education](#city-person-v1-Education) |  | 教育水平 education level |
| gender | [Gender](#city-person-v1-Gender) |  | 性别 gender |
| consumption | [Consumption](#city-person-v1-Consumption) |  | 消费水平 consumption level |
| house_id | [int32](#int32) |  | 房屋ID 区分不同家庭 House ID, identify which family this person belongs to |






<a name="city-person-v1-Persons"></a>

### Persons
智能体集合，对应一个智能体pb文件或一个智能体mongodb collection
Agent collection, corresponding to an agent pb file or an agent mongodb collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons | [Person](#city-person-v1-Person) | repeated |  |






<a name="city-person-v1-VehicleAttribute"></a>

### VehicleAttribute
车辆附加属性
Vehicle additional attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lane_change_length | [double](#double) |  | 单位: m, 完成变道所需路程 Distance required to complete lane change: m |
| min_gap | [double](#double) |  | 单位：米，本车距离前车的最小距离 The minimum distance between the vehicle and the vehicle in front: m |
| model | [string](#string) | optional | 车辆模型标签 Vehicle model tag |





 


<a name="city-person-v1-BusType"></a>

### BusType
公交车
Type of Bus

| Name | Number | Description |
| ---- | ------ | ----------- |
| BUS_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| BUS_TYPE_BUS | 1 | 公交类型 The bus is a trolleybus, BRT, eta. |
| BUS_TYPE_SUBWAY | 2 | 地铁类型 The bus is a subway |



<a name="city-person-v1-Consumption"></a>

### Consumption
智能体消费水平
agent consumption level

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONSUMPTION_UNSPECIFIED | 0 | 未指定 unspecified |
| CONSUMPTION_LOW | 1 | 低 low |
| CONSUMPTION_RELATIVELY_LOW | 2 | 较低 relatively low |
| CONSUMPTION_MEDIUM | 3 | 中等 medium |
| CONSUMPTION_RELATIVELY_HIGH | 4 | 较高 relatively high |
| CONSUMPTION_HIGH | 5 | 高 high |



<a name="city-person-v1-Education"></a>

### Education
智能体教育等级
Agent education level

| Name | Number | Description |
| ---- | ------ | ----------- |
| EDUCATION_UNSPECIFIED | 0 | 未指定 unspecified |
| EDUCATION_DOCTOR | 1 | 博士 doctor |
| EDUCATION_MASTER | 2 | 硕士 master |
| EDUCATION_BACHELOR | 3 | 本科 bachelor |
| EDUCATION_HIGH_SCHOOL | 4 | 高中 high school |
| EDUCATION_JUNIOR_HIGH_SCHOOL | 5 | 初中 junior high school |
| EDUCATION_PRIMARY_SCHOOL | 6 | 小学 primary school |
| EDUCATION_COLLEGE | 7 | 大专 college |



<a name="city-person-v1-Gender"></a>

### Gender
智能体性别
agent gender

| Name | Number | Description |
| ---- | ------ | ----------- |
| GENDER_UNSPECIFIED | 0 | 未指定 unspecified |
| GENDER_MALE | 1 | 男性 male |
| GENDER_FEMALE | 2 | 女性 female |


 

 

 



<a name="city_person_v1_person_runtime-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v1/person_runtime.proto



<a name="city-person-v1-PersonRuntime"></a>

### PersonRuntime
智能体运行时信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base | [Person](#city-person-v1-Person) | optional | person信息 person information |
| motion | [PersonMotion](#city-person-v1-PersonMotion) |  | person运动信息 person motion information |
| events | [city.event.v2.Event](#city-event-v2-Event) | repeated | 事件信息 event information |





 

 

 

 



<a name="city_person_v1_vehicle-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v1/vehicle.proto



<a name="city-person-v1-LC"></a>

### LC
变道相关的信息
lane change related information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shadow_lane_id | [int32](#int32) |  | 影子车道ID（变道前的车道） shadow lane id (lane before lane change) |
| shadow_s | [double](#double) |  | 投影到影子车道的坐标 s coordinate projected to shadow lane |
| angle | [double](#double) |  | 变道过程车头相对于前进方向的偏转角（弧度，总是为正，0代表不转向） deviation angle of the vehicle head relative to the forward direction during lane change (radians, always positive, 0 means no steering) |
| completed_ratio | [double](#double) |  | 已完成的变道比例 completed ratio of lane change |






<a name="city-person-v1-ObservedLane"></a>

### ObservedLane
观测到的车道
observed lane


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Lane ID |
| restriction | [bool](#bool) |  | 是否限行 whether restricted |
| light_state | [LightState](#city-person-v1-LightState) |  | 交通灯状态 traffic light state |
| light_remaining_time | [double](#double) |  | 交通灯剩余时间 remaining time of traffic light |






<a name="city-person-v1-ObservedVehicle"></a>

### ObservedVehicle
观测到的车辆
observed vehicles


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 车辆编号 vehicle id |
| motion | [PersonMotion](#city-person-v1-PersonMotion) |  | 当前的车辆运行时信息 current vehicle runtime information |
| relative_distance | [double](#double) |  | 相对距离 relative distance |
| relation | [VehicleRelation](#city-person-v1-VehicleRelation) |  | 关系枚举 relation enumeration |






<a name="city-person-v1-VehicleAction"></a>

### VehicleAction
车辆控制信息
vehicle control information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 车辆编号 vehicle id |
| acc | [double](#double) |  | 本轮更新中设定的加速度 acceleration set in this step |
| lc_target_id | [int32](#int32) | optional | 变道目标（可选，不设置代表不变道或保持变道状态） lane change target (optional, not set means no lane change) |
| angle | [double](#double) |  | 变道过程的转向角度 steering angle during lane change |






<a name="city-person-v1-VehicleEnv"></a>

### VehicleEnv



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 车辆编号 vehicle id |
| runtime | [VehicleRuntime](#city-person-v1-VehicleRuntime) |  | 当前的车辆运行时信息 current vehicle runtime information |
| journey | [city.routing.v2.Journey](#city-routing-v2-Journey) |  | 当前的路径规划结果 journey: current routing result |
| observed_vehicles | [ObservedVehicle](#city-person-v1-ObservedVehicle) | repeated | 观测到的车辆 observed vehicles |
| observed_lanes | [ObservedLane](#city-person-v1-ObservedLane) | repeated | 观测到的车道状态 observed lane states |






<a name="city-person-v1-VehicleRuntime"></a>

### VehicleRuntime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base | [PersonMotion](#city-person-v1-PersonMotion) |  | 基本运行时信息 basic runtime information |
| lc | [LC](#city-person-v1-LC) | optional | 变道信息 lane change information |
| action | [VehicleAction](#city-person-v1-VehicleAction) | optional | 本轮车辆行为（获取车辆环境信息时不返回） vehicle action in the step (not returned when getting vehicle environment information) |
| running_distance | [double](#double) |  | 走过的里程 running distance |
| num_going_astray | [int32](#int32) |  | 走错路次数 number of going astray |
| departure_time | [double](#double) |  | 出发时刻 departure time |
| eta | [double](#double) |  | 预计到达时刻（导航返回的eta&#43;出发时刻） estimated arrival time (eta returned by routing &#43; departure time) |
| eta_free_flow | [double](#double) |  | 自由流下的预计到达时刻 estimated arrival time under free flow |





 


<a name="city-person-v1-LightState"></a>

### LightState
交通灯的状态
traffic light state

| Name | Number | Description |
| ---- | ------ | ----------- |
| LIGHT_STATE_UNSPECIFIED | 0 | 未指定 unspecified |
| LIGHT_STATE_RED | 1 | 红灯 red light |
| LIGHT_STATE_GREEN | 2 | 绿灯 green light |
| LIGHT_STATE_YELLOW | 3 | 黄灯 yellow light |



<a name="city-person-v1-VehicleRelation"></a>

### VehicleRelation


| Name | Number | Description |
| ---- | ------ | ----------- |
| VEHICLE_RELATION_UNSPECIFIED | 0 | 未指定 unspecified |
| VEHICLE_RELATION_AHEAD | 1 | 当前车道前车 vehicle ahead in the current lane |
| VEHICLE_RELATION_BEHIND | 2 | 当前车道后车 vehicle behind in the current lane |
| VEHICLE_RELATION_SHADOW_AHEAD | 3 | 影子车道前车 vehicle ahead in the shadow lane |
| VEHICLE_RELATION_SHADOW_BEHIND | 4 | 影子车道后车 vehicle behind in the shadow lane |
| VEHICLE_RELATION_LEFT_AHEAD | 5 | 当前车道左侧车道前车 vehicle ahead in the left lane |
| VEHICLE_RELATION_RIGHT_AHEAD | 6 | 当前车道右侧车道前车 vehicle ahead in the right lane |
| VEHICLE_RELATION_LEFT_BEHIND | 7 | 当前车道左侧车道后车 vehicle behind in the left lane |
| VEHICLE_RELATION_RIGHT_BEHIND | 8 | 当前车道右侧车道后车 vehicle behind in the right lane |


 

 

 



<a name="city_person_v1_person_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v1/person_service.proto



<a name="city-person-v1-AddPersonRequest"></a>

### AddPersonRequest
新增person请求
Request for adding a new person


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person | [Person](#city-person-v1-Person) |  | 约定：person中不设置id Convention: personid is not set here |






<a name="city-person-v1-AddPersonResponse"></a>

### AddPersonResponse
新增person响应
Response of adding a new person


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | 新增的person分配得到的ID The ID assigned to the newly added person |






<a name="city-person-v1-FetchControlledVehicleEnvsRequest"></a>

### FetchControlledVehicleEnvsRequest
获取由外部控制行为的vehicle信息请求
Request for getting information of vehicle controlled by external behavior






<a name="city-person-v1-FetchControlledVehicleEnvsResponse"></a>

### FetchControlledVehicleEnvsResponse
获取由外部控制行为的vehicle信息响应
Response of getting information of vehicle controlled by external behavior


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicle_envs | [VehicleEnv](#city-person-v1-VehicleEnv) | repeated | 由外部控制行为的vehicle信息 Information of vehicle controlled by external behavior |






<a name="city-person-v1-GetAllVehiclesRequest"></a>

### GetAllVehiclesRequest
获取所有车辆请求
Request for getting all vehicles






<a name="city-person-v1-GetAllVehiclesResponse"></a>

### GetAllVehiclesResponse
获取所有车辆响应
Response of getting all vehicles


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicles | [VehicleRuntime](#city-person-v1-VehicleRuntime) | repeated | 所有车辆的信息 Information of all vehicles |






<a name="city-person-v1-GetPersonByLongLatBBoxRequest"></a>

### GetPersonByLongLatBBoxRequest
获取特定区域内的person请求
Request for getting persons in region


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bbox | [city.geo.v2.LongLatBBox](#city-geo-v2-LongLatBBox) |  | 经纬度范围 longitude and latitude bounding box |
| exclude_statuses | [Status](#city-person-v1-Status) | repeated | 过滤人的状态（状态为列表内的值的人不返回） Filter person&#39;s status (person whose status is in the list will not be returned) |
| return_base | [bool](#bool) |  | 设置是否返回base信息 Set whether to return base information |






<a name="city-person-v1-GetPersonByLongLatBBoxResponse"></a>

### GetPersonByLongLatBBoxResponse
获取特定区域内的person响应
Response of getting persons in region


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons | [PersonRuntime](#city-person-v1-PersonRuntime) | repeated | 区域内的person的信息 Information of persons in the region |






<a name="city-person-v1-GetPersonRequest"></a>

### GetPersonRequest
获取person信息请求
Request for getting person information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | person id |






<a name="city-person-v1-GetPersonResponse"></a>

### GetPersonResponse
获取person信息响应
Response of getting person information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person | [PersonRuntime](#city-person-v1-PersonRuntime) |  |  |






<a name="city-person-v1-GetPersonsRequest"></a>

### GetPersonsRequest
获取多个person信息请求
Request for getting information of multiple persons


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_ids | [int32](#int32) | repeated | person id列表，为空则返回所有person List of person ids, return all persons if empty |
| exclude_statuses | [Status](#city-person-v1-Status) | repeated | 过滤人的状态（状态为列表内的值的人不返回），即使包含在person_ids中 Filter person&#39;s status (person whose status is in the list will not be returned), even if included in person_ids |
| return_base | [bool](#bool) |  | 设置是否返回base信息 Set whether to return base information |






<a name="city-person-v1-GetPersonsResponse"></a>

### GetPersonsResponse
获取多个person信息响应
Response of getting information of multiple persons


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons | [PersonRuntime](#city-person-v1-PersonRuntime) | repeated | person信息 person information |






<a name="city-person-v1-ResetPersonPositionRequest"></a>

### ResetPersonPositionRequest
重置人的位置请求
Request for resetting person&#39;s position


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | person id |
| position | [city.geo.v2.Position](#city-geo-v2-Position) |  | 重置位置 reset position |






<a name="city-person-v1-ResetPersonPositionResponse"></a>

### ResetPersonPositionResponse
重置人的位置响应
Response of resetting person&#39;s position






<a name="city-person-v1-SetControlledVehicleActionsRequest"></a>

### SetControlledVehicleActionsRequest
设置由外部控制行为的vehicle的行为请求
Request for setting behavior of vehicle controlled by external behavior


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicle_actions | [VehicleAction](#city-person-v1-VehicleAction) | repeated | 由外部控制行为的vehicle的行为 Behavior of vehicle controlled by external behavior |






<a name="city-person-v1-SetControlledVehicleActionsResponse"></a>

### SetControlledVehicleActionsResponse
设置由外部控制行为的vehicle的行为响应
Response of setting behavior of vehicle controlled by external behavior






<a name="city-person-v1-SetControlledVehicleIDsRequest"></a>

### SetControlledVehicleIDsRequest
设置由外部控制行为的vehicle请求（下一个step生效）
Request for setting vehicle controlled by external behavior


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicle_ids | [int32](#int32) | repeated | 由外部控制行为的vehicle id列表 List of vehicle ids controlled by external behavior |






<a name="city-person-v1-SetControlledVehicleIDsResponse"></a>

### SetControlledVehicleIDsResponse
设置由外部控制行为的vehicle响应
Response of setting vehicle controlled by external behavior






<a name="city-person-v1-SetScheduleRequest"></a>

### SetScheduleRequest
修改person的schedule请求
Request for setting person schedule


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | person id |
| schedules | [city.trip.v2.Schedule](#city-trip-v2-Schedule) | repeated | 新的schedule（覆盖原有的schedule） New schedule (overwrites the original schedule) |






<a name="city-person-v1-SetScheduleResponse"></a>

### SetScheduleResponse
修改person的schedule响应
Response of setting person schedule





 

 

 


<a name="city-person-v1-PersonService"></a>

### PersonService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPerson | [GetPersonRequest](#city-person-v1-GetPersonRequest) | [GetPersonResponse](#city-person-v1-GetPersonResponse) | 获取person信息 Get person information |
| AddPerson | [AddPersonRequest](#city-person-v1-AddPersonRequest) | [AddPersonResponse](#city-person-v1-AddPersonResponse) | 新增person 传入person初始位置、目的地表、属性 返回personid Add a new person. Input person&#39;s initial location, destination table, and attributes, return personid |
| SetSchedule | [SetScheduleRequest](#city-person-v1-SetScheduleRequest) | [SetScheduleResponse](#city-person-v1-SetScheduleResponse) | 修改person的schedule 传入personid、目的地表 Set person&#39;s schedule. Input personid and destination table |
| GetPersons | [GetPersonsRequest](#city-person-v1-GetPersonsRequest) | [GetPersonsResponse](#city-person-v1-GetPersonsResponse) | 获取多个person信息 Get information of multiple persons |
| GetPersonByLongLatBBox | [GetPersonByLongLatBBoxRequest](#city-person-v1-GetPersonByLongLatBBoxRequest) | [GetPersonByLongLatBBoxResponse](#city-person-v1-GetPersonByLongLatBBoxResponse) | 获取特定区域内的person Get persons in a specific region |
| GetAllVehicles | [GetAllVehiclesRequest](#city-person-v1-GetAllVehiclesRequest) | [GetAllVehiclesResponse](#city-person-v1-GetAllVehiclesResponse) | 获取所有车辆 Get all vehicles |
| ResetPersonPosition | [ResetPersonPositionRequest](#city-person-v1-ResetPersonPositionRequest) | [ResetPersonPositionResponse](#city-person-v1-ResetPersonPositionResponse) | 重置人的位置（将停止当前正在进行的出行，转为sleep状态） Reset person&#39;s position (stop the current trip and switch to sleep status) |
| SetControlledVehicleIDs | [SetControlledVehicleIDsRequest](#city-person-v1-SetControlledVehicleIDsRequest) | [SetControlledVehicleIDsResponse](#city-person-v1-SetControlledVehicleIDsResponse) | 设置由外部控制行为的vehicle Set vehicle controlled by external behavior |
| FetchControlledVehicleEnvs | [FetchControlledVehicleEnvsRequest](#city-person-v1-FetchControlledVehicleEnvsRequest) | [FetchControlledVehicleEnvsResponse](#city-person-v1-FetchControlledVehicleEnvsResponse) | 获取由外部控制行为的vehicle信息 Get information of vehicle controlled by external behavior |
| SetControlledVehicleActions | [SetControlledVehicleActionsRequest](#city-person-v1-SetControlledVehicleActionsRequest) | [SetControlledVehicleActionsResponse](#city-person-v1-SetControlledVehicleActionsResponse) | 设置由外部控制行为的vehicle的行为 Set behavior of vehicle controlled by external behavior |

 



<a name="city_person_v2_carbon-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v2/carbon.proto



<a name="city-person-v2-VehicleCarbon"></a>

### VehicleCarbon
车辆瞬时碳排放信息
Vehicle instantaneous carbon emission information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | ID |
| ds | [double](#double) |  | delta distance (m) |
| v | [double](#double) |  | vehicle speed (m/s) |
| a | [double](#double) |  | vehicle acceleration (m/s^2) |
| u_acc | [double](#double) |  | energy for acceleration (J) |
| u_roll | [double](#double) |  | energy for rolling resistance (J) |
| u_aero | [double](#double) |  | energy for air resistance (J) |
| c_d | [double](#double) |  | C_D: drag coefficient |





 

 

 

 



<a name="city_person_v2_person-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v2/person.proto



<a name="city-person-v2-BikeAttribute"></a>

### BikeAttribute
自行车附加属性
Bike additional attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| speed | [double](#double) |  | 单位: m/s speed: m/s |
| model | [string](#string) | optional | 自行车模型标签 Bike model tag |






<a name="city-person-v2-BusAttribute"></a>

### BusAttribute
公交车附加属性
Bus additional attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subline_id | [int32](#int32) |  | 公交线路ID bus line ID |
| capacity | [int32](#int32) |  | 公交车容量 bus capacity |
| type | [BusType](#city-person-v2-BusType) |  | 公交车类型 type of bus |






<a name="city-person-v2-EmissionAttribute"></a>

### EmissionAttribute
车辆碳排附加属性
Carbon emission additional attributes for Vehicles


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| weight | [double](#double) |  | 单位: kg，车重 vehicle weight: kg |
| type | [VehicleEngineType](#city-person-v2-VehicleEngineType) |  | 车辆引擎类型 vehicle engine type |
| coefficient_drag | [double](#double) |  | 汽车空气阻力系数 Drag coefficient of the vehicle |
| lambda_s | [double](#double) |  | 路面摩擦系数 Pavement friction coefficient |
| frontal_area | [double](#double) |  | 单位: m^2，迎风面积 Frontal area: m^2 |
| fuel_efficiency | [VehicleEngineEfficiency](#city-person-v2-VehicleEngineEfficiency) | optional | 燃油引擎车辆效率 Fuel vehicle efficiency |
| electric_efficiency | [VehicleEngineEfficiency](#city-person-v2-VehicleEngineEfficiency) | optional | 电动引擎车辆效率 Fuel vehicle efficiency |






<a name="city-person-v2-PedestrianAttribute"></a>

### PedestrianAttribute
行人附加属性
Pedestrian additional attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| speed | [double](#double) |  | 单位: m/s speed: m/s |
| model | [string](#string) | optional | 行人模型标签 Pedestrian model tag |






<a name="city-person-v2-Person"></a>

### Person
智能体
agent


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 智能体ID agent ID |
| attribute | [PersonAttribute](#city-person-v2-PersonAttribute) |  | 参数 attribute |
| home | [city.geo.v2.Position](#city-geo-v2-Position) |  | 初始位置 initial position |
| schedules | [city.trip.v2.Schedule](#city-trip-v2-Schedule) | repeated | 初始日程 initial schedules |
| vehicle_attribute | [VehicleAttribute](#city-person-v2-VehicleAttribute) | optional | 车辆附加属性 vehicle addtional attribute |
| bus_attribute | [BusAttribute](#city-person-v2-BusAttribute) | optional | 公交车附加属性 bus additional attribute |
| pedestrian_attribute | [PedestrianAttribute](#city-person-v2-PedestrianAttribute) | optional | 行人附加属性 pedestrian additional attribute |
| bike_attribute | [BikeAttribute](#city-person-v2-BikeAttribute) | optional | 自行车附加属性 bike addition attribute |
| labels | [Person.LabelsEntry](#city-person-v2-Person-LabelsEntry) | repeated | [可空] 额外的标签（例如：抢修车类型-&gt;电网） [can be empty] additional tags (e.g. repair vehicle type -&gt; power grid) |
| profile | [PersonProfile](#city-person-v2-PersonProfile) | optional | [可空] 智能体简介 [can be empty] agent profile |
| work | [city.geo.v2.Position](#city-geo-v2-Position) | optional | 工作地位置 work position |
| output_when_sleep | [bool](#bool) | optional | 是否在SLEEP状态下也输出可视化（仅限车辆） Whether to output visualization in the SLEEP state (vehicles only) |
| type | [PersonType](#city-person-v2-PersonType) |  | 智能体类型 agent type |






<a name="city-person-v2-Person-LabelsEntry"></a>

### Person.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="city-person-v2-PersonAttribute"></a>

### PersonAttribute
智能体属性（通用）
Agent properties (general)






<a name="city-person-v2-PersonProfile"></a>

### PersonProfile
智能体简介
agent profile


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| age | [int32](#int32) |  | 年龄 age |
| education | [Education](#city-person-v2-Education) |  | 教育水平 education level |
| gender | [Gender](#city-person-v2-Gender) |  | 性别 gender |
| consumption | [Consumption](#city-person-v2-Consumption) |  | 消费水平 consumption level |
| house_id | [int32](#int32) |  | 房屋ID 区分不同家庭 House ID, identify which family this person belongs to |






<a name="city-person-v2-Persons"></a>

### Persons
智能体集合，对应一个智能体pb文件或一个智能体mongodb collection
Agent collection, corresponding to an agent pb file or an agent mongodb collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons | [Person](#city-person-v2-Person) | repeated |  |






<a name="city-person-v2-VehicleAttribute"></a>

### VehicleAttribute
车辆附加属性
Vehicle additional attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| length | [double](#double) |  | 单位: m，长度 length: m |
| width | [double](#double) |  | 单位: m，宽度 width: m |
| max_speed | [double](#double) |  | 单位: m/s max speed: m/s |
| max_acceleration | [double](#double) |  | 单位: m/s^2, 最大加速度（正值） max accelaration: m/s^2 (positive value) |
| max_braking_acceleration | [double](#double) |  | 单位: m/s^2, 最大减速度（负值） max deceleration: m/s^2 (negative value) |
| usual_acceleration | [double](#double) |  | 单位: m/s^2, 一般加速度（正值），要求小于最大加速度 usual acceleration: m/s^2 (positive value), required to be less than the max acceleration |
| usual_braking_acceleration | [double](#double) |  | 单位: m/s^2, 一般减速度（负值），要求大于最大减速度 usual deceleration: m/s^2 (negative value), required to be greater than the max deceleration |
| lane_change_length | [double](#double) |  | 单位: m, 完成变道所需路程 Distance required to complete lane change: m |
| min_gap | [double](#double) |  | 单位：米，本车距离前车的最小距离 The minimum distance between the vehicle and the vehicle in front: m |
| headway | [double](#double) |  | 安全车头时距 Safe time headway |
| model | [string](#string) | optional | 车辆模型标签 Vehicle model tag |
| lane_max_speed_recognition_deviation | [double](#double) |  | 本车对车道限速认知的偏差百分比 The deviation of the vehicle&#39;s recognition of lane max speed |
| emission_attribute | [EmissionAttribute](#city-person-v2-EmissionAttribute) |  | 碳排属性 Carbon emission attribute |






<a name="city-person-v2-VehicleEngineEfficiency"></a>

### VehicleEngineEfficiency



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| energy_conversion_efficiency | [double](#double) |  | 能量转换效率：车辆消耗的能量 / 燃料的能量 the energy conversion efficiency: E_{vehicle consumed} / E_{fuel or electricity} |
| c_ef | [double](#double) |  | 消耗能量(MJ) 折合到CO2排放(g)的系数 the conversion factor from consumed energy (MJ) to CO2 emissions (g) |





 


<a name="city-person-v2-BusType"></a>

### BusType
公交车
Type of Bus

| Name | Number | Description |
| ---- | ------ | ----------- |
| BUS_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| BUS_TYPE_BUS | 1 | 公交类型 The bus is a trolleybus, BRT, eta. |
| BUS_TYPE_SUBWAY | 2 | 地铁类型 The bus is a subway |



<a name="city-person-v2-Consumption"></a>

### Consumption
智能体消费水平
agent consumption level

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONSUMPTION_UNSPECIFIED | 0 | 未指定 unspecified |
| CONSUMPTION_LOW | 1 | 低 low |
| CONSUMPTION_RELATIVELY_LOW | 2 | 较低 relatively low |
| CONSUMPTION_MEDIUM | 3 | 中等 medium |
| CONSUMPTION_RELATIVELY_HIGH | 4 | 较高 relatively high |
| CONSUMPTION_HIGH | 5 | 高 high |



<a name="city-person-v2-Education"></a>

### Education
智能体教育等级
Agent education level

| Name | Number | Description |
| ---- | ------ | ----------- |
| EDUCATION_UNSPECIFIED | 0 | 未指定 unspecified |
| EDUCATION_DOCTOR | 1 | 博士 doctor |
| EDUCATION_MASTER | 2 | 硕士 master |
| EDUCATION_BACHELOR | 3 | 本科 bachelor |
| EDUCATION_HIGH_SCHOOL | 4 | 高中 high school |
| EDUCATION_JUNIOR_HIGH_SCHOOL | 5 | 初中 junior high school |
| EDUCATION_PRIMARY_SCHOOL | 6 | 小学 primary school |
| EDUCATION_COLLEGE | 7 | 大专 college |



<a name="city-person-v2-Gender"></a>

### Gender
智能体性别
agent gender

| Name | Number | Description |
| ---- | ------ | ----------- |
| GENDER_UNSPECIFIED | 0 | 未指定 unspecified |
| GENDER_MALE | 1 | 男性 male |
| GENDER_FEMALE | 2 | 女性 female |



<a name="city-person-v2-PersonType"></a>

### PersonType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PERSON_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| PERSON_TYPE_TAXI | 1 | 出租车 taxi |
| PERSON_TYPE_NORMAL | 2 | 常规智能体 normal person |



<a name="city-person-v2-VehicleEngineType"></a>

### VehicleEngineType
车辆引擎类型
vehicle type

| Name | Number | Description |
| ---- | ------ | ----------- |
| VEHICLE_ENGINE_TYPE_UNSPECIFIED | 0 | 未指定 unspecified |
| VEHICLE_ENGINE_TYPE_FUEL | 1 | 油车 gasoline vehicle |
| VEHICLE_ENGINE_TYPE_ELECTRIC | 2 | 电车 electric vehicle |
| VEHICLE_ENGINE_TYPE_HYBRID | 3 | 混合动力汽车 hybrid vehicle |


 

 

 



<a name="city_person_v2_person_runtime-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v2/person_runtime.proto



<a name="city-person-v2-PersonRuntime"></a>

### PersonRuntime
智能体运行时信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base | [Person](#city-person-v2-Person) | optional | person信息 person information |
| motion | [PersonMotion](#city-person-v2-PersonMotion) |  | person运动信息 person motion information |
| events | [city.event.v2.Event](#city-event-v2-Event) | repeated | 事件信息 event information |





 

 

 

 



<a name="city_person_v2_vehicle-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v2/vehicle.proto



<a name="city-person-v2-LC"></a>

### LC
变道相关的信息
lane change related information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shadow_lane_id | [int32](#int32) |  | 影子车道ID（变道前的车道） shadow lane id (lane before lane change) |
| shadow_s | [double](#double) |  | 投影到影子车道的坐标 s coordinate projected to shadow lane |
| angle | [double](#double) |  | 变道过程车头相对于前进方向的偏转角（弧度，总是为正，0代表不转向） deviation angle of the vehicle head relative to the forward direction during lane change (radians, always positive, 0 means no steering) |
| completed_ratio | [double](#double) |  | 已完成的变道比例 completed ratio of lane change |






<a name="city-person-v2-ObservedLane"></a>

### ObservedLane
观测到的车道
observed lane


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Lane ID |
| restriction | [bool](#bool) |  | 是否限行 whether restricted |
| light_state | [LightState](#city-person-v2-LightState) |  | 交通灯状态 traffic light state |
| light_remaining_time | [double](#double) |  | 交通灯剩余时间 remaining time of traffic light |






<a name="city-person-v2-ObservedVehicle"></a>

### ObservedVehicle
观测到的车辆
observed vehicles


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 车辆编号 vehicle id |
| motion | [PersonMotion](#city-person-v2-PersonMotion) |  | 当前的车辆运行时信息 current vehicle runtime information |
| relative_distance | [double](#double) |  | 相对距离 relative distance |
| relation | [VehicleRelation](#city-person-v2-VehicleRelation) |  | 关系枚举 relation enumeration |






<a name="city-person-v2-VehicleAction"></a>

### VehicleAction
车辆控制信息
vehicle control information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 车辆编号 vehicle id |
| acc | [double](#double) |  | 本轮更新中设定的加速度 acceleration set in this step |
| lc_target_id | [int32](#int32) | optional | 变道目标（可选，不设置代表不变道或保持变道状态） lane change target (optional, not set means no lane change) |
| angle | [double](#double) |  | 变道过程的转向角度 steering angle during lane change |






<a name="city-person-v2-VehicleEnv"></a>

### VehicleEnv



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 车辆编号 vehicle id |
| runtime | [VehicleRuntime](#city-person-v2-VehicleRuntime) |  | 当前的车辆运行时信息 current vehicle runtime information |
| journey | [city.routing.v2.Journey](#city-routing-v2-Journey) |  | 当前的路径规划结果 journey: current routing result |
| observed_vehicles | [ObservedVehicle](#city-person-v2-ObservedVehicle) | repeated | 观测到的车辆 observed vehicles |
| observed_lanes | [ObservedLane](#city-person-v2-ObservedLane) | repeated | 观测到的车道状态 observed lane states |






<a name="city-person-v2-VehicleRouteAction"></a>

### VehicleRouteAction
修改车辆路由信息
vehicle routing information modification


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 车辆编号 vehicle id |
| journey | [city.routing.v2.Journey](#city-routing-v2-Journey) |  | 新的路径规划结果 new routing result |






<a name="city-person-v2-VehicleRuntime"></a>

### VehicleRuntime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base | [PersonMotion](#city-person-v2-PersonMotion) |  | 基本运行时信息 basic runtime information |
| lc | [LC](#city-person-v2-LC) | optional | 变道信息 lane change information |
| action | [VehicleAction](#city-person-v2-VehicleAction) | optional | 本轮车辆行为（获取车辆环境信息时不返回） vehicle action in the step (not returned when getting vehicle environment information) |
| running_distance | [double](#double) |  | 走过的里程 running distance |
| num_going_astray | [int32](#int32) |  | 走错路次数 number of going astray |
| departure_time | [double](#double) |  | 出发时刻 departure time |
| eta | [double](#double) |  | 预计到达时刻（导航返回的eta&#43;出发时刻） estimated arrival time (eta returned by routing &#43; departure time) |
| eta_free_flow | [double](#double) |  | 自由流下的预计到达时刻 estimated arrival time under free flow |
| carbon | [VehicleCarbon](#city-person-v2-VehicleCarbon) | optional | 碳排放信息 carbon emission information |





 


<a name="city-person-v2-LightState"></a>

### LightState
交通灯的状态
traffic light state

| Name | Number | Description |
| ---- | ------ | ----------- |
| LIGHT_STATE_UNSPECIFIED | 0 | 未指定 unspecified |
| LIGHT_STATE_RED | 1 | 红灯 red light |
| LIGHT_STATE_GREEN | 2 | 绿灯 green light |
| LIGHT_STATE_YELLOW | 3 | 黄灯 yellow light |



<a name="city-person-v2-VehicleRelation"></a>

### VehicleRelation


| Name | Number | Description |
| ---- | ------ | ----------- |
| VEHICLE_RELATION_UNSPECIFIED | 0 | 未指定 unspecified |
| VEHICLE_RELATION_AHEAD | 1 | 当前车道前车 vehicle ahead in the current lane |
| VEHICLE_RELATION_BEHIND | 2 | 当前车道后车 vehicle behind in the current lane |
| VEHICLE_RELATION_SHADOW_AHEAD | 3 | 影子车道前车 vehicle ahead in the shadow lane |
| VEHICLE_RELATION_SHADOW_BEHIND | 4 | 影子车道后车 vehicle behind in the shadow lane |
| VEHICLE_RELATION_LEFT_AHEAD | 5 | 当前车道左侧车道前车 vehicle ahead in the left lane |
| VEHICLE_RELATION_RIGHT_AHEAD | 6 | 当前车道右侧车道前车 vehicle ahead in the right lane |
| VEHICLE_RELATION_LEFT_BEHIND | 7 | 当前车道左侧车道后车 vehicle behind in the left lane |
| VEHICLE_RELATION_RIGHT_BEHIND | 8 | 当前车道右侧车道后车 vehicle behind in the right lane |


 

 

 



<a name="city_person_v2_person_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/person/v2/person_service.proto



<a name="city-person-v2-AddPersonRequest"></a>

### AddPersonRequest
新增person请求
Request for adding a new person


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person | [Person](#city-person-v2-Person) |  | 约定：person中不设置id Convention: personid is not set here |






<a name="city-person-v2-AddPersonResponse"></a>

### AddPersonResponse
新增person响应
Response of adding a new person


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | 新增的person分配得到的ID The ID assigned to the newly added person |






<a name="city-person-v2-FetchControlledVehicleEnvsRequest"></a>

### FetchControlledVehicleEnvsRequest
获取由外部控制行为的vehicle信息请求
Request for getting information of vehicle controlled by external behavior






<a name="city-person-v2-FetchControlledVehicleEnvsResponse"></a>

### FetchControlledVehicleEnvsResponse
获取由外部控制行为的vehicle信息响应
Response of getting information of vehicle controlled by external behavior


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicle_envs | [VehicleEnv](#city-person-v2-VehicleEnv) | repeated | 由外部控制行为的vehicle信息 Information of vehicle controlled by external behavior |
| route_vehicle_envs | [VehicleEnv](#city-person-v2-VehicleEnv) | repeated | 由外部控制车辆路由的vehicle信息 Information of vehicle controlled by external behavior (control is triggered after entering a new road) |






<a name="city-person-v2-GetAllVehiclesRequest"></a>

### GetAllVehiclesRequest
获取所有车辆请求
Request for getting all vehicles






<a name="city-person-v2-GetAllVehiclesResponse"></a>

### GetAllVehiclesResponse
获取所有车辆响应
Response of getting all vehicles


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicles | [VehicleRuntime](#city-person-v2-VehicleRuntime) | repeated | 所有车辆的信息 Information of all vehicles |






<a name="city-person-v2-GetPersonByLongLatBBoxRequest"></a>

### GetPersonByLongLatBBoxRequest
获取特定区域内的person请求
Request for getting persons in region


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bbox | [city.geo.v2.LongLatBBox](#city-geo-v2-LongLatBBox) |  | 经纬度范围 longitude and latitude bounding box |
| exclude_statuses | [Status](#city-person-v2-Status) | repeated | 过滤人的状态（状态为列表内的值的人不返回） Filter person&#39;s status (person whose status is in the list will not be returned) |
| return_base | [bool](#bool) |  | 设置是否返回base信息 Set whether to return base information |






<a name="city-person-v2-GetPersonByLongLatBBoxResponse"></a>

### GetPersonByLongLatBBoxResponse
获取特定区域内的person响应
Response of getting persons in region


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons | [PersonRuntime](#city-person-v2-PersonRuntime) | repeated | 区域内的person的信息 Information of persons in the region |






<a name="city-person-v2-GetPersonRequest"></a>

### GetPersonRequest
获取person信息请求
Request for getting person information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | person id |






<a name="city-person-v2-GetPersonResponse"></a>

### GetPersonResponse
获取person信息响应
Response of getting person information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person | [PersonRuntime](#city-person-v2-PersonRuntime) |  |  |






<a name="city-person-v2-GetPersonsRequest"></a>

### GetPersonsRequest
获取多个person信息请求
Request for getting information of multiple persons


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_ids | [int32](#int32) | repeated | person id列表，为空则返回所有person List of person ids, return all persons if empty |
| exclude_statuses | [Status](#city-person-v2-Status) | repeated | 过滤人的状态（状态为列表内的值的人不返回），即使包含在person_ids中 Filter person&#39;s status (person whose status is in the list will not be returned), even if included in person_ids |
| return_base | [bool](#bool) |  | 设置是否返回base信息 Set whether to return base information |






<a name="city-person-v2-GetPersonsResponse"></a>

### GetPersonsResponse
获取多个person信息响应
Response of getting information of multiple persons


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons | [PersonRuntime](#city-person-v2-PersonRuntime) | repeated | person信息 person information |






<a name="city-person-v2-ResetPersonPositionRequest"></a>

### ResetPersonPositionRequest
重置人的位置请求
Request for resetting person&#39;s position


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | person id |
| position | [city.geo.v2.Position](#city-geo-v2-Position) |  | 重置位置 reset position |






<a name="city-person-v2-ResetPersonPositionResponse"></a>

### ResetPersonPositionResponse
重置人的位置响应
Response of resetting person&#39;s position






<a name="city-person-v2-SetControlledVehicleActionsRequest"></a>

### SetControlledVehicleActionsRequest
设置由外部控制行为的vehicle的行为请求
Request for setting behavior of vehicle controlled by external behavior


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicle_actions | [VehicleAction](#city-person-v2-VehicleAction) | repeated | 由外部控制行为的vehicle的行为 Behavior of vehicle controlled by external behavior |
| vehicle_journeys | [VehicleRouteAction](#city-person-v2-VehicleRouteAction) | repeated | 由外部控制车辆路由的vehicle的新路由 New route of vehicle controlled by external behavior (control is triggered after entering a new road) |






<a name="city-person-v2-SetControlledVehicleActionsResponse"></a>

### SetControlledVehicleActionsResponse
设置由外部控制行为的vehicle的行为响应
Response of setting behavior of vehicle controlled by external behavior






<a name="city-person-v2-SetControlledVehicleIDsRequest"></a>

### SetControlledVehicleIDsRequest
设置由外部控制行为的vehicle请求（下一个step生效）
Request for setting vehicle controlled by external behavior


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicle_ids | [int32](#int32) | repeated | 由外部控制行为的vehicle id列表 List of vehicle ids controlled by external behavior |
| route_vehicle_ids | [int32](#int32) | repeated | 由外部控制车辆路由的vehicle id列表（在进入新的road后触发控制） List of vehicle ids controlled by external behavior (control is triggered after entering a new road) |






<a name="city-person-v2-SetControlledVehicleIDsResponse"></a>

### SetControlledVehicleIDsResponse
设置由外部控制行为的vehicle响应
Response of setting vehicle controlled by external behavior






<a name="city-person-v2-SetScheduleRequest"></a>

### SetScheduleRequest
修改person的schedule请求
Request for setting person schedule


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person_id | [int32](#int32) |  | person id |
| schedules | [city.trip.v2.Schedule](#city-trip-v2-Schedule) | repeated | 新的schedule（覆盖原有的schedule） New schedule (overwrites the original schedule) |






<a name="city-person-v2-SetScheduleResponse"></a>

### SetScheduleResponse
修改person的schedule响应
Response of setting person schedule





 

 

 


<a name="city-person-v2-PersonService"></a>

### PersonService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPerson | [GetPersonRequest](#city-person-v2-GetPersonRequest) | [GetPersonResponse](#city-person-v2-GetPersonResponse) | 获取person信息 Get person information |
| AddPerson | [AddPersonRequest](#city-person-v2-AddPersonRequest) | [AddPersonResponse](#city-person-v2-AddPersonResponse) | 新增person 传入person初始位置、目的地表、属性 返回personid Add a new person. Input person&#39;s initial location, destination table, and attributes, return personid |
| SetSchedule | [SetScheduleRequest](#city-person-v2-SetScheduleRequest) | [SetScheduleResponse](#city-person-v2-SetScheduleResponse) | 修改person的schedule 传入personid、目的地表 Set person&#39;s schedule. Input personid and destination table |
| GetPersons | [GetPersonsRequest](#city-person-v2-GetPersonsRequest) | [GetPersonsResponse](#city-person-v2-GetPersonsResponse) | 获取多个person信息 Get information of multiple persons |
| GetPersonByLongLatBBox | [GetPersonByLongLatBBoxRequest](#city-person-v2-GetPersonByLongLatBBoxRequest) | [GetPersonByLongLatBBoxResponse](#city-person-v2-GetPersonByLongLatBBoxResponse) | 获取特定区域内的person Get persons in a specific region |
| GetAllVehicles | [GetAllVehiclesRequest](#city-person-v2-GetAllVehiclesRequest) | [GetAllVehiclesResponse](#city-person-v2-GetAllVehiclesResponse) | 获取所有车辆 Get all vehicles |
| ResetPersonPosition | [ResetPersonPositionRequest](#city-person-v2-ResetPersonPositionRequest) | [ResetPersonPositionResponse](#city-person-v2-ResetPersonPositionResponse) | 重置人的位置（将停止当前正在进行的出行，转为sleep状态） Reset person&#39;s position (stop the current trip and switch to sleep status) |
| SetControlledVehicleIDs | [SetControlledVehicleIDsRequest](#city-person-v2-SetControlledVehicleIDsRequest) | [SetControlledVehicleIDsResponse](#city-person-v2-SetControlledVehicleIDsResponse) | 设置由外部控制行为的vehicle Set vehicle controlled by external behavior |
| FetchControlledVehicleEnvs | [FetchControlledVehicleEnvsRequest](#city-person-v2-FetchControlledVehicleEnvsRequest) | [FetchControlledVehicleEnvsResponse](#city-person-v2-FetchControlledVehicleEnvsResponse) | 获取由外部控制行为的vehicle信息 Get information of vehicle controlled by external behavior |
| SetControlledVehicleActions | [SetControlledVehicleActionsRequest](#city-person-v2-SetControlledVehicleActionsRequest) | [SetControlledVehicleActionsResponse](#city-person-v2-SetControlledVehicleActionsResponse) | 设置由外部控制行为的vehicle的行为 Set behavior of vehicle controlled by external behavior |

 



<a name="city_ping_v1_ping_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/ping/v1/ping_service.proto



<a name="city-ping-v1-PingRequest"></a>

### PingRequest
连接测试请求






<a name="city-ping-v1-PingResponse"></a>

### PingResponse
连接测试响应





 

 

 


<a name="city-ping-v1-PingService"></a>

### PingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Ping | [PingRequest](#city-ping-v1-PingRequest) | [PingResponse](#city-ping-v1-PingResponse) | 连接测试 |

 



<a name="city_routing_v2_cost-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/routing/v2/cost.proto



<a name="city-routing-v2-Cost"></a>

### Cost
路径成本设置
Route cost settings


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 目标拓扑元素（只支持道路Road） Target topology element (only supports roads) |
| cost | [double](#double) |  | 路径成本（单位：秒） Path cost (in seconds) |
| time | [double](#double) | optional | 设置的时间（单位：秒） Set time (in seconds) 即设置几点几分的道路通行成本为cost That is, set the cost as the value at what time 为空表示设置全天通行成本均为cost If empty, it means that the all-day cost is set to the value. |





 

 

 

 



<a name="city_routing_v2_routing_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/routing/v2/routing_service.proto



<a name="city-routing-v2-GetDrivingCostsRequest"></a>

### GetDrivingCostsRequest
获取行车导航道路通行成本请求
Request for getting driving routing travelling cost


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| costs | [Cost](#city-routing-v2-Cost) | repeated | 道路通行成本（按照给定的id和time进行查询） travelling cost (query via the given ID and time) |






<a name="city-routing-v2-GetDrivingCostsResponse"></a>

### GetDrivingCostsResponse
获取行车导航道路通行成本响应
Response of getting driving routing travelling cost


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| costs | [Cost](#city-routing-v2-Cost) | repeated | 道路通行成本（补全cost后的结果） travelling cost (results after completing the cost) |






<a name="city-routing-v2-GetRouteRequest"></a>

### GetRouteRequest
获取导航路线请求
Request for getting routing path


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [RouteType](#city-routing-v2-RouteType) |  | 导航类型 routing type |
| start | [city.geo.v2.Position](#city-geo-v2-Position) |  | 起点，约定：包含LanePosition或AoiPosition中的一种 Starting point, convention: as LanePosition or AoiPosition |
| end | [city.geo.v2.Position](#city-geo-v2-Position) |  | 终点，约定：包含LanePosition或AoiPosition中的一种 Ending point, convention: as LanePosition or AoiPosition |
| time | [double](#double) |  | 发送导航请求的时间（目前仅在行车导航中使用） The time to send routing request (currently only used in driving routing) |






<a name="city-routing-v2-GetRouteResponse"></a>

### GetRouteResponse
获取导航路线响应
Response of getting routing path


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| journeys | [Journey](#city-routing-v2-Journey) | repeated |  |






<a name="city-routing-v2-SetDrivingCostsRequest"></a>

### SetDrivingCostsRequest
设置行车导航道路通行成本请求
Request for setting driving routing travelling cost


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| costs | [Cost](#city-routing-v2-Cost) | repeated | 道路通行成本 travelling cost |






<a name="city-routing-v2-SetDrivingCostsResponse"></a>

### SetDrivingCostsResponse
设置行车导航道路通行成本响应
Response of setting driving routing travelling cost





 

 

 


<a name="city-routing-v2-RoutingService"></a>

### RoutingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetRoute | [GetRouteRequest](#city-routing-v2-GetRouteRequest) | [GetRouteResponse](#city-routing-v2-GetRouteResponse) | 获取导航路线 Get routing path |
| SetDrivingCosts | [SetDrivingCostsRequest](#city-routing-v2-SetDrivingCostsRequest) | [SetDrivingCostsResponse](#city-routing-v2-SetDrivingCostsResponse) | 设置行车导航道路通行成本 Set traveling cost of driving routing |
| GetDrivingCosts | [GetDrivingCostsRequest](#city-routing-v2-GetDrivingCostsRequest) | [GetDrivingCostsResponse](#city-routing-v2-GetDrivingCostsResponse) | 获取行车导航道路通行成本 Get traveling cost of driving routing |

 



<a name="city_social_v1_message-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/social/v1/message.proto



<a name="city-social-v1-Message"></a>

### Message
消息
message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| from | [int32](#int32) |  | 消息发送者ID message sender ID |
| to | [int32](#int32) |  | 消息接收者ID message receiver ID |
| message | [string](#string) |  | 消息内容 message content |
| t | [double](#double) | optional | 消息发出时间（秒），如无则为当前模拟时间 Message sending time (in seconds), if none, it is the current simulation time 对于接收到的消息，该字段总是存在 For received messages, this field always exists |





 

 

 

 



<a name="city_social_v1_social_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/social/v1/social_service.proto



<a name="city-social-v1-ReceiveRequest"></a>

### ReceiveRequest
接收消息请求
Request for receiving message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 消息接收者ID（即为自身ID） Message receiver ID (i.e. self.ID) |






<a name="city-social-v1-ReceiveResponse"></a>

### ReceiveResponse
接收消息响应
Response of receiving messages


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messages | [Message](#city-social-v1-Message) | repeated | 接收到的消息 Received messages |






<a name="city-social-v1-SendRequest"></a>

### SendRequest
发送消息请求
Request for sendding message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messages | [Message](#city-social-v1-Message) | repeated | 待发送的消息 Messages to send |






<a name="city-social-v1-SendResponse"></a>

### SendResponse
发送消息响应
Response of sendding message





 

 

 


<a name="city-social-v1-SocialService"></a>

### SocialService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Send | [SendRequest](#city-social-v1-SendRequest) | [SendResponse](#city-social-v1-SendResponse) | 发送消息 Send message |
| Receive | [ReceiveRequest](#city-social-v1-ReceiveRequest) | [ReceiveResponse](#city-social-v1-ReceiveResponse) | 接收消息，并清空该用户的消息队列 Receive messages and clear the user&#39;s message queue |

 



<a name="city_streetview_v1_streetview-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/streetview/v1/streetview.proto



<a name="city-streetview-v1-StreetView"></a>

### StreetView
街景图片元数据


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lnglat | [city.geo.v2.LongLatPosition](#city-geo-v2-LongLatPosition) |  | WGS84经纬度位置 |
| images | [StreetViewImage](#city-streetview-v1-StreetViewImage) | repeated | 该位置的不同朝向街景图列表 |






<a name="city-streetview-v1-StreetViewImage"></a>

### StreetViewImage
街景图片描述


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| heading | [double](#double) |  | 朝向，单位度，0-360，0为正北，90为正东，180为正南，270为正西 |
| object | [string](#string) |  | 对象存储的object key |






<a name="city-streetview-v1-StreetViews"></a>

### StreetViews
导出Message，对应一个MongoDB Collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| street_views | [StreetView](#city-streetview-v1-StreetView) | repeated |  |





 

 

 

 



<a name="city_sync_v2_sync_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/sync/v2/sync_service.proto



<a name="city-sync-v2-EnterStepSyncRequest"></a>

### EnterStepSyncRequest
进入同步状态请求
Enter step sync request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 组件名，需要在同步器启动参数列表中 |






<a name="city-sync-v2-EnterStepSyncResponse"></a>

### EnterStepSyncResponse
进入同步状态响应
Enter step sync response






<a name="city-sync-v2-ExitStepSyncRequest"></a>

### ExitStepSyncRequest
退出同步状态请求
Exit step sync request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 组件名，需要在同步器启动参数列表中 |
| close | [bool](#bool) |  | 是否退出服务 |






<a name="city-sync-v2-ExitStepSyncResponse"></a>

### ExitStepSyncResponse
退出同步状态响应
Exit step sync response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| close | [bool](#bool) |  | 服务是否关闭 |






<a name="city-sync-v2-GetURLRequest"></a>

### GetURLRequest
获取程序URL请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 组件名，需要在同步器启动参数列表中 |






<a name="city-sync-v2-GetURLResponse"></a>

### GetURLResponse
获取程序URL响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | 程序URL |






<a name="city-sync-v2-SetURLRequest"></a>

### SetURLRequest
注册程序URL请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 组件名，需要在同步器启动参数列表中 |
| url | [string](#string) |  | 程序URL |






<a name="city-sync-v2-SetURLResponse"></a>

### SetURLResponse
注册程序URL响应





 

 

 


<a name="city-sync-v2-SyncService"></a>

### SyncService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetURL | [SetURLRequest](#city-sync-v2-SetURLRequest) | [SetURLResponse](#city-sync-v2-SetURLResponse) | 注册程序URL |
| GetURL | [GetURLRequest](#city-sync-v2-GetURLRequest) | [GetURLResponse](#city-sync-v2-GetURLResponse) | 获取程序URL |
| EnterStepSync | [EnterStepSyncRequest](#city-sync-v2-EnterStepSyncRequest) | [EnterStepSyncResponse](#city-sync-v2-EnterStepSyncResponse) | 程序完成本步所有操作，进入同步状态。 要求：进入同步状态的程序不再向其他程序发送消息，直到下一步开始。 |
| ExitStepSync | [ExitStepSyncRequest](#city-sync-v2-ExitStepSyncRequest) | [ExitStepSyncResponse](#city-sync-v2-ExitStepSyncResponse) | 程序完成同步阶段（无通信的安全区域）中必要的处理，如为prepare阶段加锁，可以进入准备阶段（恢复通信）。 |

 



<a name="city_water_input_v1_config-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/water/input/v1/config.proto



<a name="city-water-input-v1-Config"></a>

### Config



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mongo | [Mongo](#city-water-input-v1-Mongo) |  |  |
| control | [Control](#city-water-input-v1-Control) |  |  |
| output | [Output](#city-water-input-v1-Output) |  |  |






<a name="city-water-input-v1-Control"></a>

### Control



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| step | [ControlStep](#city-water-input-v1-ControlStep) |  |  |






<a name="city-water-input-v1-ControlStep"></a>

### ControlStep



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [int32](#int32) |  |  |
| total | [int32](#int32) |  |  |
| interval | [double](#double) |  |  |






<a name="city-water-input-v1-Mongo"></a>

### Mongo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  |  |
| map | [city.config.v1.MongoPath](#city-config-v1-MongoPath) |  |  |
| rain | [city.config.v1.MongoPath](#city-config-v1-MongoPath) |  |  |






<a name="city-water-input-v1-Output"></a>

### Output



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target | [city.config.v1.OutputTarget](#city-config-v1-OutputTarget) |  | 统一的输出目标 |
| switch | [OutputSwitch](#city-water-input-v1-OutputSwitch) |  |  |






<a name="city-water-input-v1-OutputSwitch"></a>

### OutputSwitch
是否输出各类数据


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| road | [bool](#bool) |  |  |
| drainage | [bool](#bool) |  |  |
| supply | [bool](#bool) |  |  |
| aoi | [bool](#bool) |  |  |
| event | [bool](#bool) |  |  |





 

 

 

 



<a name="city_water_input_v1_water-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/water/input/v1/water.proto



<a name="city-water-input-v1-Rain"></a>

### Rain
全天降雨情况，在数据库中体现为一条数据


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rains | [RainPeriod](#city-water-input-v1-RainPeriod) | repeated |  |






<a name="city-water-input-v1-RainPeriod"></a>

### RainPeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [int32](#int32) |  | 起始时间点，单位为秒，但必须整小时 |
| rainfall | [double](#double) |  | 降雨量：单位mm |





 

 

 

 



<a name="city_water_input_v1_input_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/water/input/v1/input_service.proto



<a name="city-water-input-v1-InitRequest"></a>

### InitRequest







<a name="city-water-input-v1-InitResponse"></a>

### InitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | 模拟器gRPC监听地址 |
| control | [Control](#city-water-input-v1-Control) |  |  |
| rain | [Rain](#city-water-input-v1-Rain) |  |  |
| map | [city.map.v2.Map](#city-map-v2-Map) |  | 仅包括header与roads |





 

 

 


<a name="city-water-input-v1-InputService"></a>

### InputService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Init | [InitRequest](#city-water-input-v1-InitRequest) | [InitResponse](#city-water-input-v1-InitResponse) |  |

 



<a name="city_water_interaction_v1_water_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/water/interaction/v1/water_service.proto



<a name="city-water-interaction-v1-GetNoWaterAOIRequest"></a>

### GetNoWaterAOIRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flag | [int32](#int32) |  |  |






<a name="city-water-interaction-v1-GetNoWaterAOIResponse"></a>

### GetNoWaterAOIResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aoi | [int32](#int32) | repeated |  |






<a name="city-water-interaction-v1-GetPumpStatusRequest"></a>

### GetPumpStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flag | [int32](#int32) |  |  |






<a name="city-water-interaction-v1-GetPumpStatusResponse"></a>

### GetPumpStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pump_status | [GetPumpStatusResponse.PumpStatusEntry](#city-water-interaction-v1-GetPumpStatusResponse-PumpStatusEntry) | repeated |  |






<a name="city-water-interaction-v1-GetPumpStatusResponse-PumpStatusEntry"></a>

### GetPumpStatusResponse.PumpStatusEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [int32](#int32) |  |  |






<a name="city-water-interaction-v1-GetRuinInfoRequest"></a>

### GetRuinInfoRequest







<a name="city-water-interaction-v1-GetRuinInfoResponse"></a>

### GetRuinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| one | [RuinInfo](#city-water-interaction-v1-RuinInfo) |  | 三级损伤信息 |
| two | [RuinInfo](#city-water-interaction-v1-RuinInfo) |  |  |
| three | [RuinInfo](#city-water-interaction-v1-RuinInfo) |  |  |






<a name="city-water-interaction-v1-RuinInfo"></a>

### RuinInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num | [int32](#int32) |  | 损坏数量 |
| ratio | [double](#double) |  | 损坏占比 |






<a name="city-water-interaction-v1-SetPumpNetworkStatusRequest"></a>

### SetPumpNetworkStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 水泵id |
| status | [bool](#bool) |  | True表示恢复，False表示摧毁 |
| type | [WaterFacilityType](#city-water-interaction-v1-WaterFacilityType) |  | 供水水泵还是排水水泵 |






<a name="city-water-interaction-v1-SetPumpNetworkStatusResponse"></a>

### SetPumpNetworkStatusResponse







<a name="city-water-interaction-v1-SetPumpPowerStatusRequest"></a>

### SetPumpPowerStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 水泵id |
| status | [bool](#bool) |  | True表示恢复，False表示摧毁 |
| type | [WaterFacilityType](#city-water-interaction-v1-WaterFacilityType) |  | 供水水泵还是排水水泵 |






<a name="city-water-interaction-v1-SetPumpPowerStatusResponse"></a>

### SetPumpPowerStatusResponse







<a name="city-water-interaction-v1-SetPumpStatusRequest"></a>

### SetPumpStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 水泵id |
| status | [bool](#bool) |  | True表示恢复，False表示摧毁 |
| type | [WaterFacilityType](#city-water-interaction-v1-WaterFacilityType) |  | 供水水泵还是排水水泵 |






<a name="city-water-interaction-v1-SetPumpStatusResponse"></a>

### SetPumpStatusResponse






 


<a name="city-water-interaction-v1-WaterFacilityType"></a>

### WaterFacilityType


| Name | Number | Description |
| ---- | ------ | ----------- |
| WATER_FACILITY_TYPE_UNSPECIFIED | 0 |  |
| WATER_FACILITY_TYPE_SUPPLY | 1 | 供水设施 |
| WATER_FACILITY_TYPE_DRAINAGE | 2 | 排水设施 |


 

 


<a name="city-water-interaction-v1-WaterService"></a>

### WaterService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetPumpPowerStatus | [SetPumpPowerStatusRequest](#city-water-interaction-v1-SetPumpPowerStatusRequest) | [SetPumpPowerStatusResponse](#city-water-interaction-v1-SetPumpPowerStatusResponse) |  |
| SetPumpNetworkStatus | [SetPumpNetworkStatusRequest](#city-water-interaction-v1-SetPumpNetworkStatusRequest) | [SetPumpNetworkStatusResponse](#city-water-interaction-v1-SetPumpNetworkStatusResponse) |  |
| SetPumpStatus | [SetPumpStatusRequest](#city-water-interaction-v1-SetPumpStatusRequest) | [SetPumpStatusResponse](#city-water-interaction-v1-SetPumpStatusResponse) |  |
| GetPumpStatus | [GetPumpStatusRequest](#city-water-interaction-v1-GetPumpStatusRequest) | [GetPumpStatusResponse](#city-water-interaction-v1-GetPumpStatusResponse) |  |
| GetNoWaterAOI | [GetNoWaterAOIRequest](#city-water-interaction-v1-GetNoWaterAOIRequest) | [GetNoWaterAOIResponse](#city-water-interaction-v1-GetNoWaterAOIResponse) |  |
| GetRuinInfo | [GetRuinInfoRequest](#city-water-interaction-v1-GetRuinInfoRequest) | [GetRuinInfoResponse](#city-water-interaction-v1-GetRuinInfoResponse) |  |

 



<a name="city_water_output_v1_output-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/water/output/v1/output.proto



<a name="city-water-output-v1-Aoi"></a>

### Aoi



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| unsatisfied_num | [int32](#int32) |  | AOI用水需求不满足人数 |
| unsatisfied_ratio | [double](#double) |  | AOI用水需求不满足比例 |
| demand | [double](#double) |  | AOI用水需求量 m3/s |
| supply | [double](#double) |  | AOI供水量 m3/s |






<a name="city-water-output-v1-DetailedRoad"></a>

### DetailedRoad



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| depths | [RoadFlood](#city-water-output-v1-RoadFlood) | repeated |  |






<a name="city-water-output-v1-DrainageBasicInfo"></a>

### DrainageBasicInfo
排水基础信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| average_power | [double](#double) |  | 排水泵平均功率 kW |
| undrained_volume | [double](#double) |  | 待排水量 m3 |
| drained_volume | [double](#double) |  | 已排水量 m3 |
| average_flow | [double](#double) |  | 平均流量 m3/s |
| flooded_volume | [double](#double) |  | 积水水量 m3 |






<a name="city-water-output-v1-DrainageMetrics"></a>

### DrainageMetrics
排水指标


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| drainage_basic_info | [DrainageBasicInfo](#city-water-output-v1-DrainageBasicInfo) |  |  |
| load_ratio | [double](#double) |  | 负载 |
| failure_statistics | [FailureStatistics](#city-water-output-v1-FailureStatistics) |  |  |






<a name="city-water-output-v1-FailureStatistics"></a>

### FailureStatistics
水泵损坏情况


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| failure_num | [int32](#int32) |  |  |
| normal_num | [int32](#int32) |  |  |
| failure_ratio | [double](#double) |  |  |






<a name="city-water-output-v1-Link"></a>

### Link
边的状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | type为LINK_TYPE_PUMP的id将以&#34;Pump_&#34;前缀 |
| type | [LinkType](#city-water-output-v1-LinkType) |  |  |
| flow | [double](#double) |  | 流量 排水网，单位：m3/s 供水网，单位：L/s |
| ok | [bool](#bool) |  |  |






<a name="city-water-output-v1-Node"></a>

### Node
节点状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Node分两种, junction和outfall, outfall的id将以&#34;_out&#34;后缀 但输出不关心NodeType, 不需用type字段显示记录 |
| head | [double](#double) |  | 水头，单位：米 |






<a name="city-water-output-v1-Road"></a>

### Road
宏观道路水深


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| depth | [double](#double) |  |  |






<a name="city-water-output-v1-RoadFlood"></a>

### RoadFlood
微观水深点


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| position | [city.geo.v2.LongLatPosition](#city-geo-v2-LongLatPosition) |  |  |
| depth | [double](#double) |  |  |






<a name="city-water-output-v1-SupplyBasicInfo"></a>

### SupplyBasicInfo
供水基础信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| average_power | [double](#double) |  | 供水泵平均功率 kW |
| average_flow | [double](#double) |  | 平均流量 m3/s |






<a name="city-water-output-v1-SupplyDemandStatistics"></a>

### SupplyDemandStatistics
供水需求及满足情况


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| persons_demand | [double](#double) |  | 居民总用水需求 m3/s |
| unsatisfied_persons | [int32](#int32) |  | 居民需求不满足人数 |
| unsatisfied_persons_ratio | [double](#double) |  | 居民需求不满足比例 |
| aois_demand | [double](#double) |  | AOI总用水需求 m3/s |
| unsatisfied_aois | [int32](#int32) |  | AOI需求不满足个数 |
| unsatisfied_aois_ratio | [double](#double) |  | AOI需求不满足比例 |






<a name="city-water-output-v1-SupplyMetrics"></a>

### SupplyMetrics
供水指标


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| supply_basic_info | [SupplyBasicInfo](#city-water-output-v1-SupplyBasicInfo) |  |  |
| supply_demand_statistics | [SupplyDemandStatistics](#city-water-output-v1-SupplyDemandStatistics) |  |  |
| load_ratio | [double](#double) |  | 负载 |
| failure_statistics | [FailureStatistics](#city-water-output-v1-FailureStatistics) |  |  |





 


<a name="city-water-output-v1-LinkType"></a>

### LinkType


| Name | Number | Description |
| ---- | ------ | ----------- |
| LINK_TYPE_UNSPECIFIED | 0 |  |
| LINK_TYPE_PIPE | 1 |  |
| LINK_TYPE_PUMP | 2 |  |


 

 

 



<a name="city_water_output_v1_output_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## city/water/output/v1/output_service.proto



<a name="city-water-output-v1-OutputRequest"></a>

### OutputRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| step | [int32](#int32) |  |  |
| roads | [Road](#city-water-output-v1-Road) | repeated | 宏观道路水深 |
| detailed_roads | [DetailedRoad](#city-water-output-v1-DetailedRoad) | repeated | 微观道路点位水深 |
| drainage_nodes | [Node](#city-water-output-v1-Node) | repeated | 排水节点 |
| drainage_links | [Link](#city-water-output-v1-Link) | repeated | 排水的边 |
| supply_nodes | [Node](#city-water-output-v1-Node) | repeated | 供水节点 |
| supply_links | [Link](#city-water-output-v1-Link) | repeated | 供水的边 |
| aois | [Aoi](#city-water-output-v1-Aoi) | repeated | AOI粒度的供水指标 |
| drainage_metric | [double](#double) |  | 排水负载指标 |
| events | [city.event.v1.Events](#city-event-v1-Events) |  | 水网模拟的各种事件 |
| drainage_metrics | [DrainageMetrics](#city-water-output-v1-DrainageMetrics) |  | 排水网指标 |
| supply_metrics | [SupplyMetrics](#city-water-output-v1-SupplyMetrics) |  | 供水网指标 |






<a name="city-water-output-v1-OutputResponse"></a>

### OutputResponse






 

 

 


<a name="city-water-output-v1-OutputService"></a>

### OutputService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Output | [OutputRequest](#city-water-output-v1-OutputRequest) | [OutputResponse](#city-water-output-v1-OutputResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

