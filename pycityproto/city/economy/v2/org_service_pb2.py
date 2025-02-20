"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.economy.v2 import economy_pb2 as city_dot_economy_dot_v2_dot_economy__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!city/economy/v2/org_service.proto\x12\x0fcity.economy.v2\x1a\x1dcity/economy/v2/economy.proto"7\n\rAddOrgRequest\x12&\n\x03org\x18\x01 \x01(\x0b2\x14.city.economy.v2.OrgR\x03org"\x10\n\x0eAddOrgResponse")\n\x10RemoveOrgRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"\x13\n\x11RemoveOrgResponse"&\n\rGetOrgRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"8\n\x0eGetOrgResponse\x12&\n\x03org\x18\x01 \x01(\x0b2\x14.city.economy.v2.OrgR\x03org":\n\x10UpdateOrgRequest\x12&\n\x03org\x18\x01 \x01(\x0b2\x14.city.economy.v2.OrgR\x03org"\x13\n\x11UpdateOrgResponse"?\n\x0fAddAgentRequest\x12,\n\x05agent\x18\x01 \x01(\x0b2\x16.city.economy.v2.AgentR\x05agent"\x12\n\x10AddAgentResponse"/\n\x12RemoveAgentRequest\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId"\x15\n\x13RemoveAgentResponse",\n\x0fGetAgentRequest\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId"@\n\x10GetAgentResponse\x12,\n\x05agent\x18\x01 \x01(\x0b2\x16.city.economy.v2.AgentR\x05agent"B\n\x12UpdateAgentRequest\x12,\n\x05agent\x18\x01 \x01(\x0b2\x16.city.economy.v2.AgentR\x05agent"\x15\n\x13UpdateAgentResponse"\xab\x01\n\x18CalculateTaxesDueRequest\x12#\n\rgovernment_id\x18\x01 \x01(\x05R\x0cgovernmentId\x12\x1b\n\tagent_ids\x18\x02 \x03(\x05R\x08agentIds\x12\x18\n\x07incomes\x18\x03 \x03(\x02R\x07incomes\x123\n\x15enable_redistribution\x18\x04 \x01(\x08R\x14enableRedistribution"a\n\x19CalculateTaxesDueResponse\x12\x1b\n\ttaxes_due\x18\x01 \x01(\x02R\x08taxesDue\x12\'\n\x0fupdated_incomes\x18\x02 \x03(\x02R\x0eupdatedIncomes"\xca\x01\n\x1bCalculateConsumptionRequest\x12\x19\n\x08firm_ids\x18\x01 \x03(\x05R\x07firmIds\x12\x19\n\x08agent_id\x18\x02 \x01(\x05R\x07agentId\x12\x18\n\x07demands\x18\x03 \x03(\x05R\x07demands\x12>\n\x18consumption_accumulation\x18\x04 \x01(\x08H\x00R\x17consumptionAccumulation\x88\x01\x01B\x1b\n\x19_consumption_accumulation"g\n\x1cCalculateConsumptionResponse\x12-\n\x12actual_consumption\x18\x01 \x01(\x02R\x11actualConsumption\x12\x18\n\x07success\x18\x02 \x01(\x08R\x07success"P\n\x18CalculateInterestRequest\x12\x17\n\x07bank_id\x18\x01 \x01(\x05R\x06bankId\x12\x1b\n\tagent_ids\x18\x02 \x03(\x05R\x08agentIds"q\n\x19CalculateInterestResponse\x12%\n\x0etotal_interest\x18\x01 \x01(\x02R\rtotalInterest\x12-\n\x12updated_currencies\x18\x02 \x03(\x02R\x11updatedCurrencies";\n\x17CalculateRealGDPRequest\x12 \n\x0cnbs_agent_id\x18\x01 \x01(\x05R\nnbsAgentId"5\n\x18CalculateRealGDPResponse\x12\x19\n\x08real_gdp\x18\x01 \x01(\x02R\x07realGdp"7\n\x0fBatchGetRequest\x12\x10\n\x03ids\x18\x01 \x03(\x05R\x03ids\x12\x12\n\x04type\x18\x02 \x01(\tR\x04type"l\n\x10BatchGetResponse\x12(\n\x04orgs\x18\x01 \x03(\x0b2\x14.city.economy.v2.OrgR\x04orgs\x12.\n\x06agents\x18\x02 \x03(\x0b2\x16.city.economy.v2.AgentR\x06agents"n\n\x12BatchUpdateRequest\x12(\n\x04orgs\x18\x01 \x03(\x0b2\x14.city.economy.v2.OrgR\x04orgs\x12.\n\x06agents\x18\x02 \x03(\x0b2\x16.city.economy.v2.AgentR\x06agents"\x15\n\x13BatchUpdateResponse"k\n\x0fBatchSetRequest\x12(\n\x04orgs\x18\x01 \x03(\x0b2\x14.city.economy.v2.OrgR\x04orgs\x12.\n\x06agents\x18\x02 \x03(\x0b2\x16.city.economy.v2.AgentR\x06agents"\x12\n\x10BatchSetResponse"\xa5\t\n\x15DeltaUpdateOrgRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12*\n\x11delta_nominal_gdp\x18\x02 \x03(\x02R\x0fdeltaNominalGdp\x12$\n\x0edelta_real_gdp\x18\x03 \x03(\x02R\x0cdeltaRealGdp\x12-\n\x12delta_unemployment\x18\x04 \x03(\x02R\x11deltaUnemployment\x12\x1f\n\x0bdelta_wages\x18\x05 \x03(\x02R\ndeltaWages\x12!\n\x0cdelta_prices\x18\x06 \x03(\x02R\x0bdeltaPrices\x12,\n\x0fdelta_inventory\x18\x07 \x01(\x05H\x00R\x0edeltaInventory\x88\x01\x01\x12$\n\x0bdelta_price\x18\x08 \x01(\x02H\x01R\ndeltaPrice\x88\x01\x01\x12*\n\x0edelta_currency\x18\t \x01(\x02H\x02R\rdeltaCurrency\x88\x01\x01\x123\n\x13delta_interest_rate\x18\n \x01(\x02H\x03R\x11deltaInterestRate\x88\x01\x01\x122\n\x15delta_bracket_cutoffs\x18\x0b \x03(\x02R\x13deltaBracketCutoffs\x12.\n\x13delta_bracket_rates\x18\x0c \x03(\x02R\x11deltaBracketRates\x12&\n\x0cdelta_demand\x18\r \x01(\x05H\x04R\x0bdeltaDemand\x88\x01\x01\x12$\n\x0bdelta_sales\x18\x0e \x01(\x05H\x05R\ndeltaSales\x88\x01\x01\x12#\n\radd_employees\x18\x0f \x03(\x05R\x0caddEmployees\x12)\n\x10remove_employees\x18\x10 \x03(\x05R\x0fremoveEmployees\x12!\n\x0cadd_citizens\x18\x11 \x03(\x05R\x0baddCitizens\x12\'\n\x0fremove_citizens\x18\x12 \x03(\x05R\x0eremoveCitizens\x12<\n\x1adelta_consumption_currency\x18\x13 \x03(\x02R\x18deltaConsumptionCurrency\x12@\n\x1cdelta_consumption_propensity\x18\x14 \x03(\x02R\x1adeltaConsumptionPropensity\x122\n\x15delta_income_currency\x18\x15 \x03(\x02R\x13deltaIncomeCurrency\x12)\n\x10delta_depression\x18\x16 \x03(\x02R\x0fdeltaDepression\x12.\n\x13delta_locus_control\x18\x17 \x03(\x02R\x11deltaLocusControl\x12.\n\x13delta_working_hours\x18\x18 \x03(\x02R\x11deltaWorkingHoursB\x12\n\x10_delta_inventoryB\x0e\n\x0c_delta_priceB\x11\n\x0f_delta_currencyB\x16\n\x14_delta_interest_rateB\x0f\n\r_delta_demandB\x0e\n\x0c_delta_sales"\x18\n\x16DeltaUpdateOrgResponse"\xdf\x02\n\x17DeltaUpdateAgentRequest\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId\x12*\n\x0edelta_currency\x18\x02 \x01(\x02H\x00R\rdeltaCurrency\x88\x01\x01\x12#\n\x0bnew_firm_id\x18\x03 \x01(\x05H\x01R\tnewFirmId\x88\x01\x01\x12$\n\x0bdelta_skill\x18\x04 \x01(\x02H\x02R\ndeltaSkill\x88\x01\x01\x120\n\x11delta_consumption\x18\x05 \x01(\x02H\x03R\x10deltaConsumption\x88\x01\x01\x12&\n\x0cdelta_income\x18\x06 \x01(\x02H\x04R\x0bdeltaIncome\x88\x01\x01B\x11\n\x0f_delta_currencyB\x0e\n\x0c_new_firm_idB\x0e\n\x0c_delta_skillB\x14\n\x12_delta_consumptionB\x0f\n\r_delta_income"\x1a\n\x18DeltaUpdateAgentResponse"\x97\x01\n\x17BatchDeltaUpdateRequest\x12:\n\x04orgs\x18\x01 \x03(\x0b2&.city.economy.v2.DeltaUpdateOrgRequestR\x04orgs\x12@\n\x06agents\x18\x02 \x03(\x0b2(.city.economy.v2.DeltaUpdateAgentRequestR\x06agents"\x1a\n\x18BatchDeltaUpdateResponse"9\n\x1aSaveEconomyEntitiesRequest\x12\x1b\n\tfile_path\x18\x01 \x01(\tR\x08filePath"S\n\x1bSaveEconomyEntitiesResponse\x12\x17\n\x07org_ids\x18\x01 \x03(\x05R\x06orgIds\x12\x1b\n\tagent_ids\x18\x02 \x03(\x05R\x08agentIds"9\n\x1aLoadEconomyEntitiesRequest\x12\x1b\n\tfile_path\x18\x01 \x01(\tR\x08filePath"S\n\x1bLoadEconomyEntitiesResponse\x12\x17\n\x07org_ids\x18\x01 \x03(\x05R\x06orgIds\x12\x1b\n\tagent_ids\x18\x02 \x03(\x05R\x08agentIds2\x83\x0f\n\nOrgService\x12I\n\x06AddOrg\x12\x1e.city.economy.v2.AddOrgRequest\x1a\x1f.city.economy.v2.AddOrgResponse\x12R\n\tRemoveOrg\x12!.city.economy.v2.RemoveOrgRequest\x1a".city.economy.v2.RemoveOrgResponse\x12I\n\x06GetOrg\x12\x1e.city.economy.v2.GetOrgRequest\x1a\x1f.city.economy.v2.GetOrgResponse\x12R\n\tUpdateOrg\x12!.city.economy.v2.UpdateOrgRequest\x1a".city.economy.v2.UpdateOrgResponse\x12O\n\x08AddAgent\x12 .city.economy.v2.AddAgentRequest\x1a!.city.economy.v2.AddAgentResponse\x12X\n\x0bRemoveAgent\x12#.city.economy.v2.RemoveAgentRequest\x1a$.city.economy.v2.RemoveAgentResponse\x12O\n\x08GetAgent\x12 .city.economy.v2.GetAgentRequest\x1a!.city.economy.v2.GetAgentResponse\x12X\n\x0bUpdateAgent\x12#.city.economy.v2.UpdateAgentRequest\x1a$.city.economy.v2.UpdateAgentResponse\x12Q\n\x08BatchGet\x12 .city.economy.v2.BatchGetRequest\x1a!.city.economy.v2.BatchGetResponse"\x00\x12Z\n\x0bBatchUpdate\x12#.city.economy.v2.BatchUpdateRequest\x1a$.city.economy.v2.BatchUpdateResponse"\x00\x12Q\n\x08BatchSet\x12 .city.economy.v2.BatchSetRequest\x1a!.city.economy.v2.BatchSetResponse"\x00\x12c\n\x0eDeltaUpdateOrg\x12&.city.economy.v2.DeltaUpdateOrgRequest\x1a\'.city.economy.v2.DeltaUpdateOrgResponse"\x00\x12i\n\x10DeltaUpdateAgent\x12(.city.economy.v2.DeltaUpdateAgentRequest\x1a).city.economy.v2.DeltaUpdateAgentResponse"\x00\x12i\n\x10BatchDeltaUpdate\x12(.city.economy.v2.BatchDeltaUpdateRequest\x1a).city.economy.v2.BatchDeltaUpdateResponse"\x00\x12l\n\x11CalculateTaxesDue\x12).city.economy.v2.CalculateTaxesDueRequest\x1a*.city.economy.v2.CalculateTaxesDueResponse"\x00\x12u\n\x14CalculateConsumption\x12,.city.economy.v2.CalculateConsumptionRequest\x1a-.city.economy.v2.CalculateConsumptionResponse"\x00\x12l\n\x11CalculateInterest\x12).city.economy.v2.CalculateInterestRequest\x1a*.city.economy.v2.CalculateInterestResponse"\x00\x12i\n\x10CalculateRealGDP\x12(.city.economy.v2.CalculateRealGDPRequest\x1a).city.economy.v2.CalculateRealGDPResponse"\x00\x12r\n\x13SaveEconomyEntities\x12+.city.economy.v2.SaveEconomyEntitiesRequest\x1a,.city.economy.v2.SaveEconomyEntitiesResponse"\x00\x12r\n\x13LoadEconomyEntities\x12+.city.economy.v2.LoadEconomyEntitiesRequest\x1a,.city.economy.v2.LoadEconomyEntitiesResponse"\x00B\xbf\x01\n\x13com.city.economy.v2B\x0fOrgServiceProtoP\x01Z9git.fiblab.net/sim/protos/v2/go/city/economy/v2;economyv2\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V2\xca\x02\x0fCity\\Economy\\V2\xe2\x02\x1bCity\\Economy\\V2\\GPBMetadata\xea\x02\x11City::Economy::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.economy.v2.org_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.economy.v2B\x0fOrgServiceProtoP\x01Z9git.fiblab.net/sim/protos/v2/go/city/economy/v2;economyv2\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V2\xca\x02\x0fCity\\Economy\\V2\xe2\x02\x1bCity\\Economy\\V2\\GPBMetadata\xea\x02\x11City::Economy::V2'
    _globals['_ADDORGREQUEST']._serialized_start = 85
    _globals['_ADDORGREQUEST']._serialized_end = 140
    _globals['_ADDORGRESPONSE']._serialized_start = 142
    _globals['_ADDORGRESPONSE']._serialized_end = 158
    _globals['_REMOVEORGREQUEST']._serialized_start = 160
    _globals['_REMOVEORGREQUEST']._serialized_end = 201
    _globals['_REMOVEORGRESPONSE']._serialized_start = 203
    _globals['_REMOVEORGRESPONSE']._serialized_end = 222
    _globals['_GETORGREQUEST']._serialized_start = 224
    _globals['_GETORGREQUEST']._serialized_end = 262
    _globals['_GETORGRESPONSE']._serialized_start = 264
    _globals['_GETORGRESPONSE']._serialized_end = 320
    _globals['_UPDATEORGREQUEST']._serialized_start = 322
    _globals['_UPDATEORGREQUEST']._serialized_end = 380
    _globals['_UPDATEORGRESPONSE']._serialized_start = 382
    _globals['_UPDATEORGRESPONSE']._serialized_end = 401
    _globals['_ADDAGENTREQUEST']._serialized_start = 403
    _globals['_ADDAGENTREQUEST']._serialized_end = 466
    _globals['_ADDAGENTRESPONSE']._serialized_start = 468
    _globals['_ADDAGENTRESPONSE']._serialized_end = 486
    _globals['_REMOVEAGENTREQUEST']._serialized_start = 488
    _globals['_REMOVEAGENTREQUEST']._serialized_end = 535
    _globals['_REMOVEAGENTRESPONSE']._serialized_start = 537
    _globals['_REMOVEAGENTRESPONSE']._serialized_end = 558
    _globals['_GETAGENTREQUEST']._serialized_start = 560
    _globals['_GETAGENTREQUEST']._serialized_end = 604
    _globals['_GETAGENTRESPONSE']._serialized_start = 606
    _globals['_GETAGENTRESPONSE']._serialized_end = 670
    _globals['_UPDATEAGENTREQUEST']._serialized_start = 672
    _globals['_UPDATEAGENTREQUEST']._serialized_end = 738
    _globals['_UPDATEAGENTRESPONSE']._serialized_start = 740
    _globals['_UPDATEAGENTRESPONSE']._serialized_end = 761
    _globals['_CALCULATETAXESDUEREQUEST']._serialized_start = 764
    _globals['_CALCULATETAXESDUEREQUEST']._serialized_end = 935
    _globals['_CALCULATETAXESDUERESPONSE']._serialized_start = 937
    _globals['_CALCULATETAXESDUERESPONSE']._serialized_end = 1034
    _globals['_CALCULATECONSUMPTIONREQUEST']._serialized_start = 1037
    _globals['_CALCULATECONSUMPTIONREQUEST']._serialized_end = 1239
    _globals['_CALCULATECONSUMPTIONRESPONSE']._serialized_start = 1241
    _globals['_CALCULATECONSUMPTIONRESPONSE']._serialized_end = 1344
    _globals['_CALCULATEINTERESTREQUEST']._serialized_start = 1346
    _globals['_CALCULATEINTERESTREQUEST']._serialized_end = 1426
    _globals['_CALCULATEINTERESTRESPONSE']._serialized_start = 1428
    _globals['_CALCULATEINTERESTRESPONSE']._serialized_end = 1541
    _globals['_CALCULATEREALGDPREQUEST']._serialized_start = 1543
    _globals['_CALCULATEREALGDPREQUEST']._serialized_end = 1602
    _globals['_CALCULATEREALGDPRESPONSE']._serialized_start = 1604
    _globals['_CALCULATEREALGDPRESPONSE']._serialized_end = 1657
    _globals['_BATCHGETREQUEST']._serialized_start = 1659
    _globals['_BATCHGETREQUEST']._serialized_end = 1714
    _globals['_BATCHGETRESPONSE']._serialized_start = 1716
    _globals['_BATCHGETRESPONSE']._serialized_end = 1824
    _globals['_BATCHUPDATEREQUEST']._serialized_start = 1826
    _globals['_BATCHUPDATEREQUEST']._serialized_end = 1936
    _globals['_BATCHUPDATERESPONSE']._serialized_start = 1938
    _globals['_BATCHUPDATERESPONSE']._serialized_end = 1959
    _globals['_BATCHSETREQUEST']._serialized_start = 1961
    _globals['_BATCHSETREQUEST']._serialized_end = 2068
    _globals['_BATCHSETRESPONSE']._serialized_start = 2070
    _globals['_BATCHSETRESPONSE']._serialized_end = 2088
    _globals['_DELTAUPDATEORGREQUEST']._serialized_start = 2091
    _globals['_DELTAUPDATEORGREQUEST']._serialized_end = 3280
    _globals['_DELTAUPDATEORGRESPONSE']._serialized_start = 3282
    _globals['_DELTAUPDATEORGRESPONSE']._serialized_end = 3306
    _globals['_DELTAUPDATEAGENTREQUEST']._serialized_start = 3309
    _globals['_DELTAUPDATEAGENTREQUEST']._serialized_end = 3660
    _globals['_DELTAUPDATEAGENTRESPONSE']._serialized_start = 3662
    _globals['_DELTAUPDATEAGENTRESPONSE']._serialized_end = 3688
    _globals['_BATCHDELTAUPDATEREQUEST']._serialized_start = 3691
    _globals['_BATCHDELTAUPDATEREQUEST']._serialized_end = 3842
    _globals['_BATCHDELTAUPDATERESPONSE']._serialized_start = 3844
    _globals['_BATCHDELTAUPDATERESPONSE']._serialized_end = 3870
    _globals['_SAVEECONOMYENTITIESREQUEST']._serialized_start = 3872
    _globals['_SAVEECONOMYENTITIESREQUEST']._serialized_end = 3929
    _globals['_SAVEECONOMYENTITIESRESPONSE']._serialized_start = 3931
    _globals['_SAVEECONOMYENTITIESRESPONSE']._serialized_end = 4014
    _globals['_LOADECONOMYENTITIESREQUEST']._serialized_start = 4016
    _globals['_LOADECONOMYENTITIESREQUEST']._serialized_end = 4073
    _globals['_LOADECONOMYENTITIESRESPONSE']._serialized_start = 4075
    _globals['_LOADECONOMYENTITIESRESPONSE']._serialized_end = 4158
    _globals['_ORGSERVICE']._serialized_start = 4161
    _globals['_ORGSERVICE']._serialized_end = 6084