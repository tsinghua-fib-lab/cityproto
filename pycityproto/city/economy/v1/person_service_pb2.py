"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.economy.v1 import economy_pb2 as city_dot_economy_dot_v1_dot_economy__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$city/economy/v1/person_service.proto\x12\x0fcity.economy.v1\x1a\x1dcity/economy/v1/economy.proto"1\n\x10GetPersonRequest\x12\x1d\n\nperson_ids\x18\x01 \x03(\x05R\tpersonIds"F\n\x11GetPersonResponse\x121\n\x07persons\x18\x01 \x03(\x0b2\x17.city.economy.v1.PersonR\x07persons"_\n\x18UpdatePersonMoneyRequest\x12C\n\x05items\x18\x01 \x03(\x0b2-.city.economy.v1.UpdatePersonMoneyRequestItemR\x05items"Q\n\x1cUpdatePersonMoneyRequestItem\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId\x12\x14\n\x05money\x18\x02 \x01(\x01R\x05money"N\n\x19UpdatePersonMoneyResponse\x121\n\x07persons\x18\x01 \x03(\x0b2\x17.city.economy.v1.PersonR\x07persons2\xd3\x01\n\rPersonService\x12T\n\tGetPerson\x12!.city.economy.v1.GetPersonRequest\x1a".city.economy.v1.GetPersonResponse"\x00\x12l\n\x11UpdatePersonMoney\x12).city.economy.v1.UpdatePersonMoneyRequest\x1a*.city.economy.v1.UpdatePersonMoneyResponse"\x00B\xbf\x01\n\x13com.city.economy.v1B\x12PersonServiceProtoP\x01Z6git.fiblab.net/sim/protos/go/city/economy/v1;economyv1\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V1\xca\x02\x0fCity\\Economy\\V1\xe2\x02\x1bCity\\Economy\\V1\\GPBMetadata\xea\x02\x11City::Economy::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.economy.v1.person_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.economy.v1B\x12PersonServiceProtoP\x01Z6git.fiblab.net/sim/protos/go/city/economy/v1;economyv1\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V1\xca\x02\x0fCity\\Economy\\V1\xe2\x02\x1bCity\\Economy\\V1\\GPBMetadata\xea\x02\x11City::Economy::V1'
    _globals['_GETPERSONREQUEST']._serialized_start = 88
    _globals['_GETPERSONREQUEST']._serialized_end = 137
    _globals['_GETPERSONRESPONSE']._serialized_start = 139
    _globals['_GETPERSONRESPONSE']._serialized_end = 209
    _globals['_UPDATEPERSONMONEYREQUEST']._serialized_start = 211
    _globals['_UPDATEPERSONMONEYREQUEST']._serialized_end = 306
    _globals['_UPDATEPERSONMONEYREQUESTITEM']._serialized_start = 308
    _globals['_UPDATEPERSONMONEYREQUESTITEM']._serialized_end = 389
    _globals['_UPDATEPERSONMONEYRESPONSE']._serialized_start = 391
    _globals['_UPDATEPERSONMONEYRESPONSE']._serialized_end = 469
    _globals['_PERSONSERVICE']._serialized_start = 472
    _globals['_PERSONSERVICE']._serialized_end = 683