"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19city/person/v2/taxi.proto\x12\x0ecity.person.v2\x1a\x15city/geo/v2/geo.proto"\xdb\x01\n\x10RequestOrderInfo\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId\x12!\n\x0crequest_time\x18\x02 \x01(\x01R\x0brequestTime\x12\x19\n\x08order_id\x18\x03 \x01(\x05R\x07orderId\x123\n\tdeparture\x18\x04 \x01(\x0b2\x15.city.geo.v2.PositionR\tdeparture\x127\n\x0bdestination\x18\x05 \x01(\x0b2\x15.city.geo.v2.PositionR\x0bdestination"\xde\x01\n\x13OrderAllocationPlan\x12\x1b\n\torder_ids\x18\x01 \x03(\x05R\x08orderIds\x12\x17\n\x07taxi_id\x18\x02 \x01(\x05R\x06taxiId\x126\n\x04type\x18\x03 \x01(\x0e2".city.person.v2.AllocationPlanTypeR\x04type\x12+\n\x12pick_up_person_ids\x18\x04 \x03(\x05R\x0fpickUpPersonIds\x12,\n\x12deliver_person_ids\x18\x05 \x03(\x05R\x10deliverPersonIds*~\n\x12AllocationPlanType\x12$\n ALLOCATION_PLAN_TYPE_UNSPECIFIED\x10\x00\x12 \n\x1cALLOCATION_PLAN_TYPE_PICK_UP\x10\x01\x12 \n\x1cALLOCATION_PLAN_TYPE_DELIVER\x10\x02B\xb2\x01\n\x12com.city.person.v2B\tTaxiProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v2.taxi_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v2B\tTaxiProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2'
    _globals['_ALLOCATIONPLANTYPE']._serialized_start = 515
    _globals['_ALLOCATIONPLANTYPE']._serialized_end = 641
    _globals['_REQUESTORDERINFO']._serialized_start = 69
    _globals['_REQUESTORDERINFO']._serialized_end = 288
    _globals['_ORDERALLOCATIONPLAN']._serialized_start = 291
    _globals['_ORDERALLOCATIONPLAN']._serialized_end = 513