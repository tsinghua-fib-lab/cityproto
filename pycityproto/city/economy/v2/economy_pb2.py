"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dcity/economy/v2/economy.proto\x12\x0fcity.economy.v2"\xdf\x03\n\x03Org\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12,\n\x04type\x18\x02 \x01(\x0e2\x18.city.economy.v2.OrgTypeR\x04type\x12\x1f\n\x0bnominal_gdp\x18\x03 \x03(\x02R\nnominalGdp\x12\x19\n\x08real_gdp\x18\x04 \x03(\x02R\x07realGdp\x12"\n\x0cunemployment\x18\x05 \x03(\x02R\x0cunemployment\x12\x14\n\x05wages\x18\x06 \x03(\x02R\x05wages\x12\x16\n\x06prices\x18\x07 \x03(\x02R\x06prices\x12!\n\tinventory\x18\x08 \x01(\x05H\x00R\tinventory\x88\x01\x01\x12\x19\n\x05price\x18\t \x01(\x02H\x01R\x05price\x88\x01\x01\x12\x1f\n\x08currency\x18\n \x01(\x02H\x02R\x08currency\x88\x01\x01\x12(\n\rinterest_rate\x18\x0b \x01(\x02H\x03R\x0cinterestRate\x88\x01\x01\x12\'\n\x0fbracket_cutoffs\x18\x0c \x03(\x02R\x0ebracketCutoffs\x12#\n\rbracket_rates\x18\r \x03(\x02R\x0cbracketRatesB\x0c\n\n_inventoryB\x08\n\x06_priceB\x0b\n\t_currencyB\x10\n\x0e_interest_rate"E\n\x05Agent\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x1f\n\x08currency\x18\x02 \x01(\x02H\x00R\x08currency\x88\x01\x01B\x0b\n\t_currency*t\n\x07OrgType\x12\x18\n\x14ORG_TYPE_UNSPECIFIED\x10\x00\x12\x10\n\x0cORG_TYPE_NBS\x10\x01\x12\x11\n\rORG_TYPE_FIRM\x10\x02\x12\x11\n\rORG_TYPE_BANK\x10\x03\x12\x17\n\x13ORG_TYPE_GOVERNMENT\x10\x04B\xbc\x01\n\x13com.city.economy.v2B\x0cEconomyProtoP\x01Z9git.fiblab.net/sim/protos/v2/go/city/economy/v2;economyv2\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V2\xca\x02\x0fCity\\Economy\\V2\xe2\x02\x1bCity\\Economy\\V2\\GPBMetadata\xea\x02\x11City::Economy::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.economy.v2.economy_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.economy.v2B\x0cEconomyProtoP\x01Z9git.fiblab.net/sim/protos/v2/go/city/economy/v2;economyv2\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V2\xca\x02\x0fCity\\Economy\\V2\xe2\x02\x1bCity\\Economy\\V2\\GPBMetadata\xea\x02\x11City::Economy::V2'
    _globals['_ORGTYPE']._serialized_start = 603
    _globals['_ORGTYPE']._serialized_end = 719
    _globals['_ORG']._serialized_start = 51
    _globals['_ORG']._serialized_end = 530
    _globals['_AGENT']._serialized_start = 532
    _globals['_AGENT']._serialized_end = 601