"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.person.v1 import motion_pb2 as city_dot_person_dot_v1_dot_motion__pb2
from ....city.person.v1 import person_pb2 as city_dot_person_dot_v1_dot_person__pb2
from ....city.person.v1 import vehicle_pb2 as city_dot_person_dot_v1_dot_vehicle__pb2
from ....city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#city/person/v1/person_service.proto\x12\x0ecity.person.v1\x1a\x15city/geo/v2/geo.proto\x1a\x1bcity/person/v1/motion.proto\x1a\x1bcity/person/v1/person.proto\x1a\x1ccity/person/v1/vehicle.proto\x1a\x17city/trip/v2/trip.proto"/\n\x10GetPersonRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId"u\n\x11GetPersonResponse\x12*\n\x04base\x18\x01 \x01(\x0b2\x16.city.person.v1.PersonR\x04base\x124\n\x06motion\x18\x02 \x01(\x0b2\x1c.city.person.v1.PersonMotionR\x06motion"B\n\x10AddPersonRequest\x12.\n\x06person\x18\x01 \x01(\x0b2\x16.city.person.v1.PersonR\x06person"0\n\x11AddPersonResponse\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId"g\n\x12SetScheduleRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId\x124\n\tschedules\x18\x02 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules"\x15\n\x13SetScheduleResponse"\x90\x01\n\x1dGetPersonByLongLatBBoxRequest\x12,\n\x04bbox\x18\x01 \x01(\x0b2\x18.city.geo.v2.LongLatBBoxR\x04bbox\x12A\n\x10exclude_statuses\x18\x02 \x03(\x0e2\x16.city.person.v1.StatusR\x0fexcludeStatuses"X\n\x1eGetPersonByLongLatBBoxResponse\x126\n\x07motions\x18\x01 \x03(\x0b2\x1c.city.person.v1.PersonMotionR\x07motions"\x17\n\x15GetAllVehiclesRequest"T\n\x16GetAllVehiclesResponse\x12:\n\x08vehicles\x18\x01 \x03(\x0b2\x1e.city.person.v1.VehicleRuntimeR\x08vehicles"A\n\x1eSetControlledVehicleIDsRequest\x12\x1f\n\x0bvehicle_ids\x18\x01 \x03(\x05R\nvehicleIds"!\n\x1fSetControlledVehicleIDsResponse"#\n!FetchControlledVehicleEnvsRequest"c\n"FetchControlledVehicleEnvsResponse\x12=\n\x0cvehicle_envs\x18\x01 \x03(\x0b2\x1a.city.person.v1.VehicleEnvR\x0bvehicleEnvs"l\n"SetControlledVehicleActionsRequest\x12F\n\x0fvehicle_actions\x18\x01 \x03(\x0b2\x1d.city.person.v1.VehicleActionR\x0evehicleActions"%\n#SetControlledVehicleActionsResponse2\xf0\x06\n\rPersonService\x12P\n\tGetPerson\x12 .city.person.v1.GetPersonRequest\x1a!.city.person.v1.GetPersonResponse\x12P\n\tAddPerson\x12 .city.person.v1.AddPersonRequest\x1a!.city.person.v1.AddPersonResponse\x12V\n\x0bSetSchedule\x12".city.person.v1.SetScheduleRequest\x1a#.city.person.v1.SetScheduleResponse\x12w\n\x16GetPersonByLongLatBBox\x12-.city.person.v1.GetPersonByLongLatBBoxRequest\x1a..city.person.v1.GetPersonByLongLatBBoxResponse\x12_\n\x0eGetAllVehicles\x12%.city.person.v1.GetAllVehiclesRequest\x1a&.city.person.v1.GetAllVehiclesResponse\x12z\n\x17SetControlledVehicleIDs\x12..city.person.v1.SetControlledVehicleIDsRequest\x1a/.city.person.v1.SetControlledVehicleIDsResponse\x12\x83\x01\n\x1aFetchControlledVehicleEnvs\x121.city.person.v1.FetchControlledVehicleEnvsRequest\x1a2.city.person.v1.FetchControlledVehicleEnvsResponse\x12\x86\x01\n\x1bSetControlledVehicleActions\x122.city.person.v1.SetControlledVehicleActionsRequest\x1a3.city.person.v1.SetControlledVehicleActionsResponseB\xb8\x01\n\x12com.city.person.v1B\x12PersonServiceProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v1.person_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v1B\x12PersonServiceProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1'
    _globals['_GETPERSONREQUEST']._serialized_start = 191
    _globals['_GETPERSONREQUEST']._serialized_end = 238
    _globals['_GETPERSONRESPONSE']._serialized_start = 240
    _globals['_GETPERSONRESPONSE']._serialized_end = 357
    _globals['_ADDPERSONREQUEST']._serialized_start = 359
    _globals['_ADDPERSONREQUEST']._serialized_end = 425
    _globals['_ADDPERSONRESPONSE']._serialized_start = 427
    _globals['_ADDPERSONRESPONSE']._serialized_end = 475
    _globals['_SETSCHEDULEREQUEST']._serialized_start = 477
    _globals['_SETSCHEDULEREQUEST']._serialized_end = 580
    _globals['_SETSCHEDULERESPONSE']._serialized_start = 582
    _globals['_SETSCHEDULERESPONSE']._serialized_end = 603
    _globals['_GETPERSONBYLONGLATBBOXREQUEST']._serialized_start = 606
    _globals['_GETPERSONBYLONGLATBBOXREQUEST']._serialized_end = 750
    _globals['_GETPERSONBYLONGLATBBOXRESPONSE']._serialized_start = 752
    _globals['_GETPERSONBYLONGLATBBOXRESPONSE']._serialized_end = 840
    _globals['_GETALLVEHICLESREQUEST']._serialized_start = 842
    _globals['_GETALLVEHICLESREQUEST']._serialized_end = 865
    _globals['_GETALLVEHICLESRESPONSE']._serialized_start = 867
    _globals['_GETALLVEHICLESRESPONSE']._serialized_end = 951
    _globals['_SETCONTROLLEDVEHICLEIDSREQUEST']._serialized_start = 953
    _globals['_SETCONTROLLEDVEHICLEIDSREQUEST']._serialized_end = 1018
    _globals['_SETCONTROLLEDVEHICLEIDSRESPONSE']._serialized_start = 1020
    _globals['_SETCONTROLLEDVEHICLEIDSRESPONSE']._serialized_end = 1053
    _globals['_FETCHCONTROLLEDVEHICLEENVSREQUEST']._serialized_start = 1055
    _globals['_FETCHCONTROLLEDVEHICLEENVSREQUEST']._serialized_end = 1090
    _globals['_FETCHCONTROLLEDVEHICLEENVSRESPONSE']._serialized_start = 1092
    _globals['_FETCHCONTROLLEDVEHICLEENVSRESPONSE']._serialized_end = 1191
    _globals['_SETCONTROLLEDVEHICLEACTIONSREQUEST']._serialized_start = 1193
    _globals['_SETCONTROLLEDVEHICLEACTIONSREQUEST']._serialized_end = 1301
    _globals['_SETCONTROLLEDVEHICLEACTIONSRESPONSE']._serialized_start = 1303
    _globals['_SETCONTROLLEDVEHICLEACTIONSRESPONSE']._serialized_end = 1340
    _globals['_PERSONSERVICE']._serialized_start = 1343
    _globals['_PERSONSERVICE']._serialized_end = 2223