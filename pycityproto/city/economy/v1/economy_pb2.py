"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dcity/economy/v1/economy.proto\x12\x0fcity.economy.v1"U\n\x06Person\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n\x05money\x18\x02 \x01(\x01R\x05money\x12\x1a\n\x06org_id\x18\x03 \x01(\x05H\x00R\x05orgId\x88\x01\x01B\t\n\x07_org_id"?\n\x08Employee\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId\x12\x16\n\x06salary\x18\x02 \x01(\x01R\x06salary"h\n\x03Job\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12%\n\x0eemployee_count\x18\x02 \x01(\x05R\remployeeCount\x12\x1b\n\x06salary\x18\x03 \x01(\x01H\x00R\x06salary\x88\x01\x01B\t\n\x07_salary"j\n\x05Goods\x12\x12\n\x04type\x18\x01 \x01(\tR\x04type\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n\x05count\x18\x03 \x01(\x05R\x05count\x12\x19\n\x05price\x18\x04 \x01(\x01H\x00R\x05price\x88\x01\x01B\x08\n\x06_price"\xf1\x01\n\x03Org\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x15\n\x06poi_id\x18\x02 \x01(\x05R\x05poiId\x127\n\temployees\x18\x03 \x03(\x0b2\x19.city.economy.v1.EmployeeR\temployees\x12(\n\x04jobs\x18\x04 \x03(\x0b2\x14.city.economy.v1.JobR\x04jobs\x12\x14\n\x05money\x18\x05 \x01(\x01R\x05money\x12,\n\x05goods\x18\x06 \x03(\x0b2\x16.city.economy.v1.GoodsR\x05goods\x12\x1c\n\tfunctions\x18\x07 \x03(\tR\tfunctions"f\n\x07Economy\x121\n\x07persons\x18\x01 \x03(\x0b2\x17.city.economy.v1.PersonR\x07persons\x12(\n\x04orgs\x18\x02 \x03(\x0b2\x14.city.economy.v1.OrgR\x04orgsB\xb9\x01\n\x13com.city.economy.v1B\x0cEconomyProtoP\x01Z6git.fiblab.net/sim/protos/go/city/economy/v1;economyv1\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V1\xca\x02\x0fCity\\Economy\\V1\xe2\x02\x1bCity\\Economy\\V1\\GPBMetadata\xea\x02\x11City::Economy::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.economy.v1.economy_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.economy.v1B\x0cEconomyProtoP\x01Z6git.fiblab.net/sim/protos/go/city/economy/v1;economyv1\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V1\xca\x02\x0fCity\\Economy\\V1\xe2\x02\x1bCity\\Economy\\V1\\GPBMetadata\xea\x02\x11City::Economy::V1'
    _globals['_PERSON']._serialized_start = 50
    _globals['_PERSON']._serialized_end = 135
    _globals['_EMPLOYEE']._serialized_start = 137
    _globals['_EMPLOYEE']._serialized_end = 200
    _globals['_JOB']._serialized_start = 202
    _globals['_JOB']._serialized_end = 306
    _globals['_GOODS']._serialized_start = 308
    _globals['_GOODS']._serialized_end = 414
    _globals['_ORG']._serialized_start = 417
    _globals['_ORG']._serialized_end = 658
    _globals['_ECONOMY']._serialized_start = 660
    _globals['_ECONOMY']._serialized_end = 762