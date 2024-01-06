"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.economy.v1 import economy_pb2 as city_dot_economy_dot_v1_dot_economy__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!city/economy/v1/org_service.proto\x12\x0fcity.economy.v1\x1a\x1dcity/economy/v1/economy.proto"(\n\rGetOrgRequest\x12\x17\n\x07org_ids\x18\x01 \x03(\x05R\x06orgIds":\n\x0eGetOrgResponse\x12(\n\x04orgs\x18\x01 \x03(\x0b2\x14.city.economy.v1.OrgR\x04orgs"Y\n\x15UpdateOrgMoneyRequest\x12@\n\x05items\x18\x01 \x03(\x0b2*.city.economy.v1.UpdateOrgMoneyRequestItemR\x05items"H\n\x19UpdateOrgMoneyRequestItem\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12\x14\n\x05money\x18\x02 \x01(\x01R\x05money"B\n\x16UpdateOrgMoneyResponse\x12(\n\x04orgs\x18\x01 \x03(\x0b2\x14.city.economy.v1.OrgR\x04orgs"Y\n\x15UpdateOrgGoodsRequest\x12@\n\x05items\x18\x01 \x03(\x0b2*.city.economy.v1.UpdateOrgGoodsRequestItemR\x05items"`\n\x19UpdateOrgGoodsRequestItem\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12,\n\x05goods\x18\x02 \x03(\x0b2\x16.city.economy.v1.GoodsR\x05goods"B\n\x16UpdateOrgGoodsResponse\x12(\n\x04orgs\x18\x01 \x03(\x0b2\x14.city.economy.v1.OrgR\x04orgs"_\n\x18UpdateOrgEmployeeRequest\x12C\n\x05items\x18\x01 \x03(\x0b2-.city.economy.v1.UpdateOrgEmployeeRequestItemR\x05items"\xad\x01\n\x1cUpdateOrgEmployeeRequestItem\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12-\n\x04adds\x18\x02 \x03(\x0b2\x19.city.economy.v1.EmployeeR\x04adds\x12\x12\n\x04dels\x18\x03 \x03(\x05R\x04dels\x123\n\x07updates\x18\x04 \x03(\x0b2\x19.city.economy.v1.EmployeeR\x07updates"E\n\x19UpdateOrgEmployeeResponse\x12(\n\x04orgs\x18\x01 \x03(\x0b2\x14.city.economy.v1.OrgR\x04orgs"U\n\x13UpdateOrgJobRequest\x12>\n\x05items\x18\x01 \x03(\x0b2(.city.economy.v1.UpdateOrgJobRequestItemR\x05items"Z\n\x17UpdateOrgJobRequestItem\x12\x15\n\x06org_id\x18\x01 \x01(\x05R\x05orgId\x12(\n\x04jobs\x18\x02 \x03(\x0b2\x14.city.economy.v1.JobR\x04jobs"@\n\x14UpdateOrgJobResponse\x12(\n\x04orgs\x18\x01 \x03(\x0b2\x14.city.economy.v1.OrgR\x04orgs2\xf0\x03\n\nOrgService\x12K\n\x06GetOrg\x12\x1e.city.economy.v1.GetOrgRequest\x1a\x1f.city.economy.v1.GetOrgResponse"\x00\x12c\n\x0eUpdateOrgMoney\x12&.city.economy.v1.UpdateOrgMoneyRequest\x1a\'.city.economy.v1.UpdateOrgMoneyResponse"\x00\x12c\n\x0eUpdateOrgGoods\x12&.city.economy.v1.UpdateOrgGoodsRequest\x1a\'.city.economy.v1.UpdateOrgGoodsResponse"\x00\x12l\n\x11UpdateOrgEmployee\x12).city.economy.v1.UpdateOrgEmployeeRequest\x1a*.city.economy.v1.UpdateOrgEmployeeResponse"\x00\x12]\n\x0cUpdateOrgJob\x12$.city.economy.v1.UpdateOrgJobRequest\x1a%.city.economy.v1.UpdateOrgJobResponse"\x00B\xbc\x01\n\x13com.city.economy.v1B\x0fOrgServiceProtoP\x01Z6git.fiblab.net/sim/protos/go/city/economy/v1;economyv1\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V1\xca\x02\x0fCity\\Economy\\V1\xe2\x02\x1bCity\\Economy\\V1\\GPBMetadata\xea\x02\x11City::Economy::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.economy.v1.org_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.economy.v1B\x0fOrgServiceProtoP\x01Z6git.fiblab.net/sim/protos/go/city/economy/v1;economyv1\xa2\x02\x03CEX\xaa\x02\x0fCity.Economy.V1\xca\x02\x0fCity\\Economy\\V1\xe2\x02\x1bCity\\Economy\\V1\\GPBMetadata\xea\x02\x11City::Economy::V1'
    _globals['_GETORGREQUEST']._serialized_start = 85
    _globals['_GETORGREQUEST']._serialized_end = 125
    _globals['_GETORGRESPONSE']._serialized_start = 127
    _globals['_GETORGRESPONSE']._serialized_end = 185
    _globals['_UPDATEORGMONEYREQUEST']._serialized_start = 187
    _globals['_UPDATEORGMONEYREQUEST']._serialized_end = 276
    _globals['_UPDATEORGMONEYREQUESTITEM']._serialized_start = 278
    _globals['_UPDATEORGMONEYREQUESTITEM']._serialized_end = 350
    _globals['_UPDATEORGMONEYRESPONSE']._serialized_start = 352
    _globals['_UPDATEORGMONEYRESPONSE']._serialized_end = 418
    _globals['_UPDATEORGGOODSREQUEST']._serialized_start = 420
    _globals['_UPDATEORGGOODSREQUEST']._serialized_end = 509
    _globals['_UPDATEORGGOODSREQUESTITEM']._serialized_start = 511
    _globals['_UPDATEORGGOODSREQUESTITEM']._serialized_end = 607
    _globals['_UPDATEORGGOODSRESPONSE']._serialized_start = 609
    _globals['_UPDATEORGGOODSRESPONSE']._serialized_end = 675
    _globals['_UPDATEORGEMPLOYEEREQUEST']._serialized_start = 677
    _globals['_UPDATEORGEMPLOYEEREQUEST']._serialized_end = 772
    _globals['_UPDATEORGEMPLOYEEREQUESTITEM']._serialized_start = 775
    _globals['_UPDATEORGEMPLOYEEREQUESTITEM']._serialized_end = 948
    _globals['_UPDATEORGEMPLOYEERESPONSE']._serialized_start = 950
    _globals['_UPDATEORGEMPLOYEERESPONSE']._serialized_end = 1019
    _globals['_UPDATEORGJOBREQUEST']._serialized_start = 1021
    _globals['_UPDATEORGJOBREQUEST']._serialized_end = 1106
    _globals['_UPDATEORGJOBREQUESTITEM']._serialized_start = 1108
    _globals['_UPDATEORGJOBREQUESTITEM']._serialized_end = 1198
    _globals['_UPDATEORGJOBRESPONSE']._serialized_start = 1200
    _globals['_UPDATEORGJOBRESPONSE']._serialized_end = 1264
    _globals['_ORGSERVICE']._serialized_start = 1267
    _globals['_ORGSERVICE']._serialized_end = 1763