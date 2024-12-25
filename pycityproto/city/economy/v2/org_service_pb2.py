"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.economy.v2 import economy_pb2 as city_dot_economy_dot_v2_dot_economy__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!city/economy/v2/org_service.proto\x12\x0fcity.economy.v2\x1a\x1dcity/economy/v2/economy.proto"7\n\rAddOrgRequest\x12&\n\x03org\x18\x01 \x01(\x0b2\x14.city.economy.v2.OrgR\x03org"\x10\n\x0eAddOrgResponse")\n\x10RemoveOrgRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"\x13\n\x11RemoveOrgResponse"?\n\x0fAddAgentRequest\x12,\n\x05agent\x18\x01 \x01(\x0b2\x16.city.economy.v2.AgentR\x05agent"\x12\n\x10AddAgentResponse"/\n\x12RemoveAgentRequest\x12\x19\n\x08agent_id\x18\x01 \x01(\x05R\x07agentId"\x15\n\x13RemoveAgentResponse"-\n\x14GetNominalGDPRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"8\n\x15GetNominalGDPResponse\x12\x1f\n\x0bnominal_gdp\x18\x01 \x03(\x02R\nnominalGdp"N\n\x14SetNominalGDPRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\x1f\n\x0bnominal_gdp\x18\x02 \x03(\x02R\nnominalGdp"\x17\n\x15SetNominalGDPResponse"*\n\x11GetRealGDPRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"/\n\x12GetRealGDPResponse\x12\x19\n\x08real_gdp\x18\x01 \x03(\x02R\x07realGdp"E\n\x11SetRealGDPRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\x19\n\x08real_gdp\x18\x02 \x03(\x02R\x07realGdp"\x14\n\x12SetRealGDPResponse"/\n\x16GetUnemploymentRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"=\n\x17GetUnemploymentResponse\x12"\n\x0cunemployment\x18\x01 \x03(\x02R\x0cunemployment"S\n\x16SetUnemploymentRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12"\n\x0cunemployment\x18\x02 \x03(\x02R\x0cunemployment"\x19\n\x17SetUnemploymentResponse"(\n\x0fGetWagesRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"(\n\x10GetWagesResponse\x12\x14\n\x05wages\x18\x01 \x03(\x02R\x05wages">\n\x0fSetWagesRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\x14\n\x05wages\x18\x02 \x03(\x02R\x05wages"\x12\n\x10SetWagesResponse")\n\x10GetPricesRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"+\n\x11GetPricesResponse\x12\x16\n\x06prices\x18\x01 \x03(\x02R\x06prices"A\n\x10SetPricesRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\x16\n\x06prices\x18\x02 \x03(\x02R\x06prices"\x13\n\x11SetPricesResponse",\n\x13GetInventoryRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"4\n\x14GetInventoryResponse\x12\x1c\n\tinventory\x18\x01 \x01(\x05R\tinventory"J\n\x13SetInventoryRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\x1c\n\tinventory\x18\x02 \x01(\x05R\tinventory"\x16\n\x14SetInventoryResponse"(\n\x0fGetPriceRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"(\n\x10GetPriceResponse\x12\x14\n\x05price\x18\x01 \x01(\x02R\x05price">\n\x0fSetPriceRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\x14\n\x05price\x18\x02 \x01(\x02R\x05price"\x12\n\x10SetPriceResponse"+\n\x12GetCurrencyRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"1\n\x13GetCurrencyResponse\x12\x1a\n\x08currency\x18\x01 \x01(\x02R\x08currency"G\n\x12SetCurrencyRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\x1a\n\x08currency\x18\x02 \x01(\x02R\x08currency"\x15\n\x13SetCurrencyResponse"/\n\x16GetInterestRateRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId">\n\x17GetInterestRateResponse\x12#\n\rinterest_rate\x18\x01 \x01(\x02R\x0cinterestRate"T\n\x16SetInterestRateRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12#\n\rinterest_rate\x18\x02 \x01(\x02R\x0cinterestRate"\x19\n\x17SetInterestRateResponse"1\n\x18GetBracketCutoffsRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"D\n\x19GetBracketCutoffsResponse\x12\'\n\x0fbracket_cutoffs\x18\x01 \x03(\x02R\x0ebracketCutoffs"Z\n\x18SetBracketCutoffsRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\'\n\x0fbracket_cutoffs\x18\x02 \x03(\x02R\x0ebracketCutoffs"\x1b\n\x19SetBracketCutoffsResponse"/\n\x16GetBracketRatesRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId">\n\x17GetBracketRatesResponse\x12#\n\rbracket_rates\x18\x01 \x03(\x02R\x0cbracketRates"T\n\x16SetBracketRatesRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12#\n\rbracket_rates\x18\x02 \x03(\x02R\x0cbracketRates"\x19\n\x17SetBracketRatesResponse"\xab\x01\n\x18CalculateTaxesDueRequest\x12#\n\rgovernment_id\x18\x01 \x01(\x05R\x0cgovernmentId\x12\x1b\n\tagent_ids\x18\x02 \x03(\x05R\x08agentIds\x12\x18\n\x07incomes\x18\x03 \x03(\x02R\x07incomes\x123\n\x15enable_redistribution\x18\x04 \x01(\x08R\x14enableRedistribution"a\n\x19CalculateTaxesDueResponse\x12\x1b\n\ttaxes_due\x18\x01 \x01(\x02R\x08taxesDue\x12\'\n\x0fupdated_incomes\x18\x02 \x03(\x02R\x0eupdatedIncomes"m\n\x1bCalculateConsumptionRequest\x12\x17\n\x07firm_id\x18\x01 \x01(\x05R\x06firmId\x12\x1b\n\tagent_ids\x18\x02 \x03(\x05R\x08agentIds\x12\x18\n\x07demands\x18\x03 \x03(\x05R\x07demands"x\n\x1cCalculateConsumptionResponse\x12)\n\x10remain_inventory\x18\x01 \x01(\x05R\x0fremainInventory\x12-\n\x12updated_currencies\x18\x02 \x03(\x02R\x11updatedCurrencies"P\n\x18CalculateInterestRequest\x12\x17\n\x07bank_id\x18\x01 \x01(\x05R\x06bankId\x12\x1b\n\tagent_ids\x18\x02 \x03(\x05R\x08agentIds"q\n\x19CalculateInterestResponse\x12%\n\x0etotal_interest\x18\x01 \x01(\x02R\rtotalInterest\x12-\n\x12updated_currencies\x18\x02 \x03(\x02R\x11updatedCurrencies"9\n\x1aSaveEconomyEntitiesRequest\x12\x1b\n\tfile_path\x18\x01 \x01(\tR\x08filePath"S\n\x1bSaveEconomyEntitiesResponse\x12\x1b\n\tagent_ids\x18\x01 \x03(\x05R\x08agentIds\x12\x17\n\x07org_ids\x18\x02 \x03(\x05R\x06orgIds"9\n\x1aLoadEconomyEntitiesRequest\x12\x1b\n\tfile_path\x18\x01 \x01(\tR\x08filePath"S\n\x1bLoadEconomyEntitiesResponse\x12\x1b\n\tagent_ids\x18\x01 \x03(\x05R\x08agentIds\x12\x17\n\x07org_ids\x18\x02 \x03(\x05R\x06orgIds"6\n\x1dGetConsumptionCurrencyRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"S\n\x1eGetConsumptionCurrencyResponse\x121\n\x14consumption_currency\x18\x01 \x03(\x02R\x13consumptionCurrency"i\n\x1dSetConsumptionCurrencyRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x121\n\x14consumption_currency\x18\x02 \x03(\x02R\x13consumptionCurrency" \n\x1eSetConsumptionCurrencyResponse"8\n\x1fGetConsumptionPropensityRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"Y\n GetConsumptionPropensityResponse\x125\n\x16consumption_propensity\x18\x01 \x03(\x02R\x15consumptionPropensity"o\n\x1fSetConsumptionPropensityRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x125\n\x16consumption_propensity\x18\x02 \x03(\x02R\x15consumptionPropensity""\n SetConsumptionPropensityResponse"1\n\x18GetIncomeCurrencyRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"D\n\x19GetIncomeCurrencyResponse\x12\'\n\x0fincome_currency\x18\x01 \x03(\x02R\x0eincomeCurrency"Z\n\x18SetIncomeCurrencyRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\'\n\x0fincome_currency\x18\x02 \x03(\x02R\x0eincomeCurrency"\x1b\n\x19SetIncomeCurrencyResponse"-\n\x14GetDepressionRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId"7\n\x15GetDepressionResponse\x12\x1e\n\ndepression\x18\x01 \x03(\x02R\ndepression"M\n\x14SetDepressionRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\x1e\n\ndepression\x18\x02 \x03(\x02R\ndepression"\x17\n\x15SetDepressionResponse"/\n\x16GetLocusControlRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId">\n\x17GetLocusControlResponse\x12#\n\rlocus_control\x18\x01 \x03(\x02R\x0clocusControl"T\n\x16SetLocusControlRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12#\n\rlocus_control\x18\x02 \x03(\x02R\x0clocusControl"\x19\n\x17SetLocusControlResponse"/\n\x16GetWorkingHoursRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId">\n\x17GetWorkingHoursResponse\x12#\n\rworking_hours\x18\x01 \x03(\x02R\x0cworkingHours"T\n\x16SetWorkingHoursRequest\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12#\n\rworking_hours\x18\x02 \x03(\x02R\x0cworkingHours"\x19\n\x17SetWorkingHoursResponse2\xb3!\n\nOrgService\x12I\n\x06AddOrg\x12\x1e.city.economy.v2.AddOrgRequest\x1a\x1f.city.economy.v2.AddOrgResponse\x12R\n\tRemoveOrg\x12!.city.economy.v2.RemoveOrgRequest\x1a".city.economy.v2.RemoveOrgResponse\x12O\n\x08AddAgent\x12 .city.economy.v2.AddAgentRequest\x1a!.city.economy.v2.AddAgentResponse\x12X\n\x0bRemoveAgent\x12#.city.economy.v2.RemoveAgentRequest\x1a$.city.economy.v2.RemoveAgentResponse\x12^\n\rGetNominalGDP\x12%.city.economy.v2.GetNominalGDPRequest\x1a&.city.economy.v2.GetNominalGDPResponse\x12^\n\rSetNominalGDP\x12%.city.economy.v2.SetNominalGDPRequest\x1a&.city.economy.v2.SetNominalGDPResponse\x12U\n\nGetRealGDP\x12".city.economy.v2.GetRealGDPRequest\x1a#.city.economy.v2.GetRealGDPResponse\x12U\n\nSetRealGDP\x12".city.economy.v2.SetRealGDPRequest\x1a#.city.economy.v2.SetRealGDPResponse\x12d\n\x0fGetUnemployment\x12\'.city.economy.v2.GetUnemploymentRequest\x1a(.city.economy.v2.GetUnemploymentResponse\x12d\n\x0fSetUnemployment\x12\'.city.economy.v2.SetUnemploymentRequest\x1a(.city.economy.v2.SetUnemploymentResponse\x12O\n\x08GetWages\x12 .city.economy.v2.GetWagesRequest\x1a!.city.economy.v2.GetWagesResponse\x12O\n\x08SetWages\x12 .city.economy.v2.SetWagesRequest\x1a!.city.economy.v2.SetWagesResponse\x12R\n\tGetPrices\x12!.city.economy.v2.GetPricesRequest\x1a".city.economy.v2.GetPricesResponse\x12R\n\tSetPrices\x12!.city.economy.v2.SetPricesRequest\x1a".city.economy.v2.SetPricesResponse\x12[\n\x0cGetInventory\x12$.city.economy.v2.GetInventoryRequest\x1a%.city.economy.v2.GetInventoryResponse\x12[\n\x0cSetInventory\x12$.city.economy.v2.SetInventoryRequest\x1a%.city.economy.v2.SetInventoryResponse\x12O\n\x08GetPrice\x12 .city.economy.v2.GetPriceRequest\x1a!.city.economy.v2.GetPriceResponse\x12O\n\x08SetPrice\x12 .city.economy.v2.SetPriceRequest\x1a!.city.economy.v2.SetPriceResponse\x12X\n\x0bGetCurrency\x12#.city.economy.v2.GetCurrencyRequest\x1a$.city.economy.v2.GetCurrencyResponse\x12X\n\x0bSetCurrency\x12#.city.economy.v2.SetCurrencyRequest\x1a$.city.economy.v2.SetCurrencyResponse\x12d\n\x0fGetInterestRate\x12\'.city.economy.v2.GetInterestRateRequest\x1a(.city.economy.v2.GetInterestRateResponse\x12d\n\x0fSetInterestRate\x12\'.city.economy.v2.SetInterestRateRequest\x1a(.city.economy.v2.SetInterestRateResponse\x12j\n\x11GetBracketCutoffs\x12).city.economy.v2.GetBracketCutoffsRequest\x1a*.city.economy.v2.GetBracketCutoffsResponse\x12j\n\x11SetBracketCutoffs\x12).city.economy.v2.SetBracketCutoffsRequest\x1a*.city.economy.v2.SetBracketCutoffsResponse\x12d\n\x0fGetBracketRates\x12\'.city.economy.v2.GetBracketRatesRequest\x1a(.city.economy.v2.GetBracketRatesResponse\x12d\n\x0fSetBracketRates\x12\'.city.economy.v2.SetBracketRatesRequest\x1a(.city.economy.v2.SetBracketRatesResponse\x12j\n\x11CalculateTaxesDue\x12).city.economy.v2.CalculateTaxesDueRequest\x1a*.city.economy.v2.CalculateTaxesDueResponse\x12s\n\x14CalculateConsumption\x12,.city.economy.v2.CalculateConsumptionRequest\x1a-.city.economy.v2.CalculateConsumptionResponse\x12j\n\x11CalculateInterest\x12).city.economy.v2.CalculateInterestRequest\x1a*.city.economy.v2.CalculateInterestResponse\x12p\n\x13SaveEconomyEntities\x12+.city.economy.v2.SaveEconomyEntitiesRequest\x1a,.city.economy.v2.SaveEconomyEntitiesResponse\x12p\n\x13LoadEconomyEntities\x12+.city.economy.v2.LoadEconomyEntitiesRequest\x1a,.city.economy.v2.LoadEconomyEntitiesResponse\x12y\n\x16GetConsumptionCurrency\x12..city.economy.v2.GetConsumptionCurrencyRequest\x1a/.city.economy.v2.GetConsumptionCurrencyResponse\x12y\n\x16SetConsumptionCurrency\x12..city.economy.v2.SetConsumptionCurrencyRequest\x1a/.city.economy.v2.SetConsumptionCurrencyResponse\x12\x7f\n\x18GetConsumptionPropensity\x120.city.economy.v2.GetConsumptionPropensityRequest\x1a1.city.economy.v2.GetConsumptionPropensityResponse\x12\x7f\n\x18SetConsumptionPropensity\x120.city.economy.v2.SetConsumptionPropensityRequest\x1a1.city.economy.v2.SetConsumptionPropensityResponse\x12j\n\x11GetIncomeCurrency\x12).city.economy.v2.GetIncomeCurrencyRequest\x1a*.city.economy.v2.GetIncomeCurrencyResponse\x12j\n\x11SetIncomeCurrency\x12).city.economy.v2.SetIncomeCurrencyRequest\x1a*.city.economy.v2.SetIncomeCurrencyResponse\x12^\n\rGetDepression\x12%.city.economy.v2.GetDepressionRequest\x1a&.city.economy.v2.GetDepressionResponse\x12^\n\rSetDepression\x12%.city.economy.v2.SetDepressionRequest\x1a&.city.economy.v2.SetDepressionResponse\x12d\n\x0fGetLocusControl\x12\'.city.economy.v2.GetLocusControlRequest\x1a(.city.economy.v2.GetLocusControlResponse\x12d\n\x0fSetLocusControl\x12\'.city.economy.v2.SetLocusControlRequest\x1a(.city.economy.v2.SetLocusControlResponse\x12d\n\x0fGetWorkingHours\x12\'.city.economy.v2.GetWorkingHoursRequest\x1a(.city.economy.v2.GetWorkingHoursResponse\x12d\n\x0fSetWorkingHours\x12\'.city.economy.v2.SetWorkingHoursRequest\x1a(.city.economy.v2.SetWorkingHoursResponseB\xbf\x01\n\x13com.city.economy.v2B\x0fOrgServiceProtoP\x01Z9git.fiblab.net/sim/protos/v2/go/city/economy/v2;economyv2\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V2\xca\x02\x0fCity\\Economy\\V2\xe2\x02\x1bCity\\Economy\\V2\\GPBMetadata\xea\x02\x11City::Economy::V2b\x06proto3')
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
    _globals['_ADDAGENTREQUEST']._serialized_start = 224
    _globals['_ADDAGENTREQUEST']._serialized_end = 287
    _globals['_ADDAGENTRESPONSE']._serialized_start = 289
    _globals['_ADDAGENTRESPONSE']._serialized_end = 307
    _globals['_REMOVEAGENTREQUEST']._serialized_start = 309
    _globals['_REMOVEAGENTREQUEST']._serialized_end = 356
    _globals['_REMOVEAGENTRESPONSE']._serialized_start = 358
    _globals['_REMOVEAGENTRESPONSE']._serialized_end = 379
    _globals['_GETNOMINALGDPREQUEST']._serialized_start = 381
    _globals['_GETNOMINALGDPREQUEST']._serialized_end = 426
    _globals['_GETNOMINALGDPRESPONSE']._serialized_start = 428
    _globals['_GETNOMINALGDPRESPONSE']._serialized_end = 484
    _globals['_SETNOMINALGDPREQUEST']._serialized_start = 486
    _globals['_SETNOMINALGDPREQUEST']._serialized_end = 564
    _globals['_SETNOMINALGDPRESPONSE']._serialized_start = 566
    _globals['_SETNOMINALGDPRESPONSE']._serialized_end = 589
    _globals['_GETREALGDPREQUEST']._serialized_start = 591
    _globals['_GETREALGDPREQUEST']._serialized_end = 633
    _globals['_GETREALGDPRESPONSE']._serialized_start = 635
    _globals['_GETREALGDPRESPONSE']._serialized_end = 682
    _globals['_SETREALGDPREQUEST']._serialized_start = 684
    _globals['_SETREALGDPREQUEST']._serialized_end = 753
    _globals['_SETREALGDPRESPONSE']._serialized_start = 755
    _globals['_SETREALGDPRESPONSE']._serialized_end = 775
    _globals['_GETUNEMPLOYMENTREQUEST']._serialized_start = 777
    _globals['_GETUNEMPLOYMENTREQUEST']._serialized_end = 824
    _globals['_GETUNEMPLOYMENTRESPONSE']._serialized_start = 826
    _globals['_GETUNEMPLOYMENTRESPONSE']._serialized_end = 887
    _globals['_SETUNEMPLOYMENTREQUEST']._serialized_start = 889
    _globals['_SETUNEMPLOYMENTREQUEST']._serialized_end = 972
    _globals['_SETUNEMPLOYMENTRESPONSE']._serialized_start = 974
    _globals['_SETUNEMPLOYMENTRESPONSE']._serialized_end = 999
    _globals['_GETWAGESREQUEST']._serialized_start = 1001
    _globals['_GETWAGESREQUEST']._serialized_end = 1041
    _globals['_GETWAGESRESPONSE']._serialized_start = 1043
    _globals['_GETWAGESRESPONSE']._serialized_end = 1083
    _globals['_SETWAGESREQUEST']._serialized_start = 1085
    _globals['_SETWAGESREQUEST']._serialized_end = 1147
    _globals['_SETWAGESRESPONSE']._serialized_start = 1149
    _globals['_SETWAGESRESPONSE']._serialized_end = 1167
    _globals['_GETPRICESREQUEST']._serialized_start = 1169
    _globals['_GETPRICESREQUEST']._serialized_end = 1210
    _globals['_GETPRICESRESPONSE']._serialized_start = 1212
    _globals['_GETPRICESRESPONSE']._serialized_end = 1255
    _globals['_SETPRICESREQUEST']._serialized_start = 1257
    _globals['_SETPRICESREQUEST']._serialized_end = 1322
    _globals['_SETPRICESRESPONSE']._serialized_start = 1324
    _globals['_SETPRICESRESPONSE']._serialized_end = 1343
    _globals['_GETINVENTORYREQUEST']._serialized_start = 1345
    _globals['_GETINVENTORYREQUEST']._serialized_end = 1389
    _globals['_GETINVENTORYRESPONSE']._serialized_start = 1391
    _globals['_GETINVENTORYRESPONSE']._serialized_end = 1443
    _globals['_SETINVENTORYREQUEST']._serialized_start = 1445
    _globals['_SETINVENTORYREQUEST']._serialized_end = 1519
    _globals['_SETINVENTORYRESPONSE']._serialized_start = 1521
    _globals['_SETINVENTORYRESPONSE']._serialized_end = 1543
    _globals['_GETPRICEREQUEST']._serialized_start = 1545
    _globals['_GETPRICEREQUEST']._serialized_end = 1585
    _globals['_GETPRICERESPONSE']._serialized_start = 1587
    _globals['_GETPRICERESPONSE']._serialized_end = 1627
    _globals['_SETPRICEREQUEST']._serialized_start = 1629
    _globals['_SETPRICEREQUEST']._serialized_end = 1691
    _globals['_SETPRICERESPONSE']._serialized_start = 1693
    _globals['_SETPRICERESPONSE']._serialized_end = 1711
    _globals['_GETCURRENCYREQUEST']._serialized_start = 1713
    _globals['_GETCURRENCYREQUEST']._serialized_end = 1756
    _globals['_GETCURRENCYRESPONSE']._serialized_start = 1758
    _globals['_GETCURRENCYRESPONSE']._serialized_end = 1807
    _globals['_SETCURRENCYREQUEST']._serialized_start = 1809
    _globals['_SETCURRENCYREQUEST']._serialized_end = 1880
    _globals['_SETCURRENCYRESPONSE']._serialized_start = 1882
    _globals['_SETCURRENCYRESPONSE']._serialized_end = 1903
    _globals['_GETINTERESTRATEREQUEST']._serialized_start = 1905
    _globals['_GETINTERESTRATEREQUEST']._serialized_end = 1952
    _globals['_GETINTERESTRATERESPONSE']._serialized_start = 1954
    _globals['_GETINTERESTRATERESPONSE']._serialized_end = 2016
    _globals['_SETINTERESTRATEREQUEST']._serialized_start = 2018
    _globals['_SETINTERESTRATEREQUEST']._serialized_end = 2102
    _globals['_SETINTERESTRATERESPONSE']._serialized_start = 2104
    _globals['_SETINTERESTRATERESPONSE']._serialized_end = 2129
    _globals['_GETBRACKETCUTOFFSREQUEST']._serialized_start = 2131
    _globals['_GETBRACKETCUTOFFSREQUEST']._serialized_end = 2180
    _globals['_GETBRACKETCUTOFFSRESPONSE']._serialized_start = 2182
    _globals['_GETBRACKETCUTOFFSRESPONSE']._serialized_end = 2250
    _globals['_SETBRACKETCUTOFFSREQUEST']._serialized_start = 2252
    _globals['_SETBRACKETCUTOFFSREQUEST']._serialized_end = 2342
    _globals['_SETBRACKETCUTOFFSRESPONSE']._serialized_start = 2344
    _globals['_SETBRACKETCUTOFFSRESPONSE']._serialized_end = 2371
    _globals['_GETBRACKETRATESREQUEST']._serialized_start = 2373
    _globals['_GETBRACKETRATESREQUEST']._serialized_end = 2420
    _globals['_GETBRACKETRATESRESPONSE']._serialized_start = 2422
    _globals['_GETBRACKETRATESRESPONSE']._serialized_end = 2484
    _globals['_SETBRACKETRATESREQUEST']._serialized_start = 2486
    _globals['_SETBRACKETRATESREQUEST']._serialized_end = 2570
    _globals['_SETBRACKETRATESRESPONSE']._serialized_start = 2572
    _globals['_SETBRACKETRATESRESPONSE']._serialized_end = 2597
    _globals['_CALCULATETAXESDUEREQUEST']._serialized_start = 2600
    _globals['_CALCULATETAXESDUEREQUEST']._serialized_end = 2771
    _globals['_CALCULATETAXESDUERESPONSE']._serialized_start = 2773
    _globals['_CALCULATETAXESDUERESPONSE']._serialized_end = 2870
    _globals['_CALCULATECONSUMPTIONREQUEST']._serialized_start = 2872
    _globals['_CALCULATECONSUMPTIONREQUEST']._serialized_end = 2981
    _globals['_CALCULATECONSUMPTIONRESPONSE']._serialized_start = 2983
    _globals['_CALCULATECONSUMPTIONRESPONSE']._serialized_end = 3103
    _globals['_CALCULATEINTERESTREQUEST']._serialized_start = 3105
    _globals['_CALCULATEINTERESTREQUEST']._serialized_end = 3185
    _globals['_CALCULATEINTERESTRESPONSE']._serialized_start = 3187
    _globals['_CALCULATEINTERESTRESPONSE']._serialized_end = 3300
    _globals['_SAVEECONOMYENTITIESREQUEST']._serialized_start = 3302
    _globals['_SAVEECONOMYENTITIESREQUEST']._serialized_end = 3359
    _globals['_SAVEECONOMYENTITIESRESPONSE']._serialized_start = 3361
    _globals['_SAVEECONOMYENTITIESRESPONSE']._serialized_end = 3444
    _globals['_LOADECONOMYENTITIESREQUEST']._serialized_start = 3446
    _globals['_LOADECONOMYENTITIESREQUEST']._serialized_end = 3503
    _globals['_LOADECONOMYENTITIESRESPONSE']._serialized_start = 3505
    _globals['_LOADECONOMYENTITIESRESPONSE']._serialized_end = 3588
    _globals['_GETCONSUMPTIONCURRENCYREQUEST']._serialized_start = 3590
    _globals['_GETCONSUMPTIONCURRENCYREQUEST']._serialized_end = 3644
    _globals['_GETCONSUMPTIONCURRENCYRESPONSE']._serialized_start = 3646
    _globals['_GETCONSUMPTIONCURRENCYRESPONSE']._serialized_end = 3729
    _globals['_SETCONSUMPTIONCURRENCYREQUEST']._serialized_start = 3731
    _globals['_SETCONSUMPTIONCURRENCYREQUEST']._serialized_end = 3836
    _globals['_SETCONSUMPTIONCURRENCYRESPONSE']._serialized_start = 3838
    _globals['_SETCONSUMPTIONCURRENCYRESPONSE']._serialized_end = 3870
    _globals['_GETCONSUMPTIONPROPENSITYREQUEST']._serialized_start = 3872
    _globals['_GETCONSUMPTIONPROPENSITYREQUEST']._serialized_end = 3928
    _globals['_GETCONSUMPTIONPROPENSITYRESPONSE']._serialized_start = 3930
    _globals['_GETCONSUMPTIONPROPENSITYRESPONSE']._serialized_end = 4019
    _globals['_SETCONSUMPTIONPROPENSITYREQUEST']._serialized_start = 4021
    _globals['_SETCONSUMPTIONPROPENSITYREQUEST']._serialized_end = 4132
    _globals['_SETCONSUMPTIONPROPENSITYRESPONSE']._serialized_start = 4134
    _globals['_SETCONSUMPTIONPROPENSITYRESPONSE']._serialized_end = 4168
    _globals['_GETINCOMECURRENCYREQUEST']._serialized_start = 4170
    _globals['_GETINCOMECURRENCYREQUEST']._serialized_end = 4219
    _globals['_GETINCOMECURRENCYRESPONSE']._serialized_start = 4221
    _globals['_GETINCOMECURRENCYRESPONSE']._serialized_end = 4289
    _globals['_SETINCOMECURRENCYREQUEST']._serialized_start = 4291
    _globals['_SETINCOMECURRENCYREQUEST']._serialized_end = 4381
    _globals['_SETINCOMECURRENCYRESPONSE']._serialized_start = 4383
    _globals['_SETINCOMECURRENCYRESPONSE']._serialized_end = 4410
    _globals['_GETDEPRESSIONREQUEST']._serialized_start = 4412
    _globals['_GETDEPRESSIONREQUEST']._serialized_end = 4457
    _globals['_GETDEPRESSIONRESPONSE']._serialized_start = 4459
    _globals['_GETDEPRESSIONRESPONSE']._serialized_end = 4514
    _globals['_SETDEPRESSIONREQUEST']._serialized_start = 4516
    _globals['_SETDEPRESSIONREQUEST']._serialized_end = 4593
    _globals['_SETDEPRESSIONRESPONSE']._serialized_start = 4595
    _globals['_SETDEPRESSIONRESPONSE']._serialized_end = 4618
    _globals['_GETLOCUSCONTROLREQUEST']._serialized_start = 4620
    _globals['_GETLOCUSCONTROLREQUEST']._serialized_end = 4667
    _globals['_GETLOCUSCONTROLRESPONSE']._serialized_start = 4669
    _globals['_GETLOCUSCONTROLRESPONSE']._serialized_end = 4731
    _globals['_SETLOCUSCONTROLREQUEST']._serialized_start = 4733
    _globals['_SETLOCUSCONTROLREQUEST']._serialized_end = 4817
    _globals['_SETLOCUSCONTROLRESPONSE']._serialized_start = 4819
    _globals['_SETLOCUSCONTROLRESPONSE']._serialized_end = 4844
    _globals['_GETWORKINGHOURSREQUEST']._serialized_start = 4846
    _globals['_GETWORKINGHOURSREQUEST']._serialized_end = 4893
    _globals['_GETWORKINGHOURSRESPONSE']._serialized_start = 4895
    _globals['_GETWORKINGHOURSRESPONSE']._serialized_end = 4957
    _globals['_SETWORKINGHOURSREQUEST']._serialized_start = 4959
    _globals['_SETWORKINGHOURSREQUEST']._serialized_end = 5043
    _globals['_SETWORKINGHOURSRESPONSE']._serialized_start = 5045
    _globals['_SETWORKINGHOURSRESPONSE']._serialized_end = 5070
    _globals['_ORGSERVICE']._serialized_start = 5073
    _globals['_ORGSERVICE']._serialized_end = 9348