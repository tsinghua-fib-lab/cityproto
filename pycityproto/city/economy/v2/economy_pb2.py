"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dcity/economy/v2/economy.proto\x12\x0fcity.economy.v2"\xb2\x01\n\x04Firm\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x1c\n\temployees\x18\x02 \x03(\x05R\temployees\x12\x14\n\x05price\x18\x03 \x01(\x02R\x05price\x12\x1c\n\tinventory\x18\x04 \x01(\x05R\tinventory\x12\x16\n\x06demand\x18\x05 \x01(\x02R\x06demand\x12\x14\n\x05sales\x18\x06 \x01(\x02R\x05sales\x12\x1a\n\x08currency\x18\x07 \x01(\x02R\x08currency"\xd2\x0b\n\x03NBS\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x1f\n\x0bcitizen_ids\x18\x02 \x03(\x05R\ncitizenIds\x12E\n\x0bnominal_gdp\x18\x03 \x03(\x0b2$.city.economy.v2.NBS.NominalGdpEntryR\nnominalGdp\x12<\n\x08real_gdp\x18\x04 \x03(\x0b2!.city.economy.v2.NBS.RealGdpEntryR\x07realGdp\x12J\n\x0cunemployment\x18\x05 \x03(\x0b2&.city.economy.v2.NBS.UnemploymentEntryR\x0cunemployment\x125\n\x05wages\x18\x06 \x03(\x0b2\x1f.city.economy.v2.NBS.WagesEntryR\x05wages\x128\n\x06prices\x18\x07 \x03(\x0b2 .city.economy.v2.NBS.PricesEntryR\x06prices\x12K\n\rworking_hours\x18\x08 \x03(\x0b2&.city.economy.v2.NBS.WorkingHoursEntryR\x0cworkingHours\x12D\n\ndepression\x18\t \x03(\x0b2$.city.economy.v2.NBS.DepressionEntryR\ndepression\x12`\n\x14consumption_currency\x18\n \x03(\x0b2-.city.economy.v2.NBS.ConsumptionCurrencyEntryR\x13consumptionCurrency\x12Q\n\x0fincome_currency\x18\x0b \x03(\x0b2(.city.economy.v2.NBS.IncomeCurrencyEntryR\x0eincomeCurrency\x12K\n\rlocus_control\x18\x0c \x03(\x0b2&.city.economy.v2.NBS.LocusControlEntryR\x0clocusControl\x12*\n\x11citizen_agent_ids\x18\r \x03(\x05R\x0fcitizenAgentIds\x12\x1a\n\x08currency\x18\x0e \x01(\x02R\x08currency\x1a=\n\x0fNominalGdpEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\x1a:\n\x0cRealGdpEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\x1a?\n\x11UnemploymentEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\x1a8\n\nWagesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\x1a9\n\x0bPricesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\x1a?\n\x11WorkingHoursEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\x1a=\n\x0fDepressionEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\x1aF\n\x18ConsumptionCurrencyEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\x1aA\n\x13IncomeCurrencyEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\x1a?\n\x11LocusControlEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01"\xa7\x01\n\nGovernment\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x1f\n\x0bcitizen_ids\x18\x02 \x03(\x05R\ncitizenIds\x12\'\n\x0fbracket_cutoffs\x18\x03 \x03(\x02R\x0ebracketCutoffs\x12#\n\rbracket_rates\x18\x04 \x03(\x02R\x0cbracketRates\x12\x1a\n\x08currency\x18\x05 \x01(\x02R\x08currency"x\n\x04Bank\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x1f\n\x0bcitizen_ids\x18\x02 \x03(\x05R\ncitizenIds\x12#\n\rinterest_rate\x18\x03 \x01(\x02R\x0cinterestRate\x12\x1a\n\x08currency\x18\x04 \x01(\x02R\x08currency"\xf3\x01\n\x05Agent\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x1f\n\x08currency\x18\x02 \x01(\x02H\x00R\x08currency\x88\x01\x01\x12\x1c\n\x07firm_id\x18\x03 \x01(\x05H\x01R\x06firmId\x88\x01\x01\x12\x19\n\x05skill\x18\x04 \x01(\x02H\x02R\x05skill\x88\x01\x01\x12%\n\x0bconsumption\x18\x05 \x01(\x02H\x03R\x0bconsumption\x88\x01\x01\x12\x1b\n\x06income\x18\x06 \x01(\x02H\x04R\x06income\x88\x01\x01B\x0b\n\t_currencyB\n\n\x08_firm_idB\x08\n\x06_skillB\x0e\n\x0c_consumptionB\t\n\x07_income"\x82\x02\n\x0fEconomyEntities\x12+\n\x05firms\x18\x01 \x03(\x0b2\x15.city.economy.v2.FirmR\x05firms\x12&\n\x03nbs\x18\x02 \x03(\x0b2\x14.city.economy.v2.NBSR\x03nbs\x12=\n\x0bgovernments\x18\x03 \x03(\x0b2\x1b.city.economy.v2.GovernmentR\x0bgovernments\x12+\n\x05banks\x18\x04 \x03(\x0b2\x15.city.economy.v2.BankR\x05banks\x12.\n\x06agents\x18\x05 \x03(\x0b2\x16.city.economy.v2.AgentR\x06agentsB\xbc\x01\n\x13com.city.economy.v2B\x0cEconomyProtoP\x01Z9git.fiblab.net/sim/protos/v2/go/city/economy/v2;economyv2\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V2\xca\x02\x0fCity\\Economy\\V2\xe2\x02\x1bCity\\Economy\\V2\\GPBMetadata\xea\x02\x11City::Economy::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.economy.v2.economy_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.economy.v2B\x0cEconomyProtoP\x01Z9git.fiblab.net/sim/protos/v2/go/city/economy/v2;economyv2\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V2\xca\x02\x0fCity\\Economy\\V2\xe2\x02\x1bCity\\Economy\\V2\\GPBMetadata\xea\x02\x11City::Economy::V2'
    _globals['_NBS_NOMINALGDPENTRY']._options = None
    _globals['_NBS_NOMINALGDPENTRY']._serialized_options = b'8\x01'
    _globals['_NBS_REALGDPENTRY']._options = None
    _globals['_NBS_REALGDPENTRY']._serialized_options = b'8\x01'
    _globals['_NBS_UNEMPLOYMENTENTRY']._options = None
    _globals['_NBS_UNEMPLOYMENTENTRY']._serialized_options = b'8\x01'
    _globals['_NBS_WAGESENTRY']._options = None
    _globals['_NBS_WAGESENTRY']._serialized_options = b'8\x01'
    _globals['_NBS_PRICESENTRY']._options = None
    _globals['_NBS_PRICESENTRY']._serialized_options = b'8\x01'
    _globals['_NBS_WORKINGHOURSENTRY']._options = None
    _globals['_NBS_WORKINGHOURSENTRY']._serialized_options = b'8\x01'
    _globals['_NBS_DEPRESSIONENTRY']._options = None
    _globals['_NBS_DEPRESSIONENTRY']._serialized_options = b'8\x01'
    _globals['_NBS_CONSUMPTIONCURRENCYENTRY']._options = None
    _globals['_NBS_CONSUMPTIONCURRENCYENTRY']._serialized_options = b'8\x01'
    _globals['_NBS_INCOMECURRENCYENTRY']._options = None
    _globals['_NBS_INCOMECURRENCYENTRY']._serialized_options = b'8\x01'
    _globals['_NBS_LOCUSCONTROLENTRY']._options = None
    _globals['_NBS_LOCUSCONTROLENTRY']._serialized_options = b'8\x01'
    _globals['_FIRM']._serialized_start = 51
    _globals['_FIRM']._serialized_end = 229
    _globals['_NBS']._serialized_start = 232
    _globals['_NBS']._serialized_end = 1722
    _globals['_NBS_NOMINALGDPENTRY']._serialized_start = 1087
    _globals['_NBS_NOMINALGDPENTRY']._serialized_end = 1148
    _globals['_NBS_REALGDPENTRY']._serialized_start = 1150
    _globals['_NBS_REALGDPENTRY']._serialized_end = 1208
    _globals['_NBS_UNEMPLOYMENTENTRY']._serialized_start = 1210
    _globals['_NBS_UNEMPLOYMENTENTRY']._serialized_end = 1273
    _globals['_NBS_WAGESENTRY']._serialized_start = 1275
    _globals['_NBS_WAGESENTRY']._serialized_end = 1331
    _globals['_NBS_PRICESENTRY']._serialized_start = 1333
    _globals['_NBS_PRICESENTRY']._serialized_end = 1390
    _globals['_NBS_WORKINGHOURSENTRY']._serialized_start = 1392
    _globals['_NBS_WORKINGHOURSENTRY']._serialized_end = 1455
    _globals['_NBS_DEPRESSIONENTRY']._serialized_start = 1457
    _globals['_NBS_DEPRESSIONENTRY']._serialized_end = 1518
    _globals['_NBS_CONSUMPTIONCURRENCYENTRY']._serialized_start = 1520
    _globals['_NBS_CONSUMPTIONCURRENCYENTRY']._serialized_end = 1590
    _globals['_NBS_INCOMECURRENCYENTRY']._serialized_start = 1592
    _globals['_NBS_INCOMECURRENCYENTRY']._serialized_end = 1657
    _globals['_NBS_LOCUSCONTROLENTRY']._serialized_start = 1659
    _globals['_NBS_LOCUSCONTROLENTRY']._serialized_end = 1722
    _globals['_GOVERNMENT']._serialized_start = 1725
    _globals['_GOVERNMENT']._serialized_end = 1892
    _globals['_BANK']._serialized_start = 1894
    _globals['_BANK']._serialized_end = 2014
    _globals['_AGENT']._serialized_start = 2017
    _globals['_AGENT']._serialized_end = 2260
    _globals['_ECONOMYENTITIES']._serialized_start = 2263
    _globals['_ECONOMYENTITIES']._serialized_end = 2521