"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.person.v1 import motion_pb2 as city_dot_person_dot_v1_dot_motion__pb2
from ....city.person.v1 import person_pb2 as city_dot_person_dot_v1_dot_person__pb2
from ....city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#city/person/v1/person_service.proto\x12\x0ecity.person.v1\x1a\x15city/geo/v2/geo.proto\x1a\x1bcity/person/v1/motion.proto\x1a\x1bcity/person/v1/person.proto\x1a\x17city/trip/v2/trip.proto"/\n\x10GetPersonRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId"u\n\x11GetPersonResponse\x12*\n\x04base\x18\x01 \x01(\x0b2\x16.city.person.v1.PersonR\x04base\x124\n\x06motion\x18\x02 \x01(\x0b2\x1c.city.person.v1.PersonMotionR\x06motion"B\n\x10AddPersonRequest\x12.\n\x06person\x18\x01 \x01(\x0b2\x16.city.person.v1.PersonR\x06person"0\n\x11AddPersonResponse\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId"g\n\x12SetScheduleRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId\x124\n\tschedules\x18\x02 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules"\x15\n\x13SetScheduleResponse"\x90\x01\n\x1dGetPersonByLongLatBBoxRequest\x12,\n\x04bbox\x18\x01 \x01(\x0b2\x18.city.geo.v2.LongLatBBoxR\x04bbox\x12A\n\x10exclude_statuses\x18\x02 \x03(\x0e2\x16.city.person.v1.StatusR\x0fexcludeStatuses"X\n\x1eGetPersonByLongLatBBoxResponse\x126\n\x07motions\x18\x01 \x03(\x0b2\x1c.city.person.v1.PersonMotionR\x07motions2\x84\x03\n\rPersonService\x12P\n\tGetPerson\x12 .city.person.v1.GetPersonRequest\x1a!.city.person.v1.GetPersonResponse\x12P\n\tAddPerson\x12 .city.person.v1.AddPersonRequest\x1a!.city.person.v1.AddPersonResponse\x12V\n\x0bSetSchedule\x12".city.person.v1.SetScheduleRequest\x1a#.city.person.v1.SetScheduleResponse\x12w\n\x16GetPersonByLongLatBBox\x12-.city.person.v1.GetPersonByLongLatBBoxRequest\x1a..city.person.v1.GetPersonByLongLatBBoxResponseB\xb8\x01\n\x12com.city.person.v1B\x12PersonServiceProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v1.person_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v1B\x12PersonServiceProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1'
    _globals['_GETPERSONREQUEST']._serialized_start = 161
    _globals['_GETPERSONREQUEST']._serialized_end = 208
    _globals['_GETPERSONRESPONSE']._serialized_start = 210
    _globals['_GETPERSONRESPONSE']._serialized_end = 327
    _globals['_ADDPERSONREQUEST']._serialized_start = 329
    _globals['_ADDPERSONREQUEST']._serialized_end = 395
    _globals['_ADDPERSONRESPONSE']._serialized_start = 397
    _globals['_ADDPERSONRESPONSE']._serialized_end = 445
    _globals['_SETSCHEDULEREQUEST']._serialized_start = 447
    _globals['_SETSCHEDULEREQUEST']._serialized_end = 550
    _globals['_SETSCHEDULERESPONSE']._serialized_start = 552
    _globals['_SETSCHEDULERESPONSE']._serialized_end = 573
    _globals['_GETPERSONBYLONGLATBBOXREQUEST']._serialized_start = 576
    _globals['_GETPERSONBYLONGLATBBOXREQUEST']._serialized_end = 720
    _globals['_GETPERSONBYLONGLATBBOXRESPONSE']._serialized_start = 722
    _globals['_GETPERSONBYLONGLATBBOXRESPONSE']._serialized_end = 810
    _globals['_PERSONSERVICE']._serialized_start = 813
    _globals['_PERSONSERVICE']._serialized_end = 1201