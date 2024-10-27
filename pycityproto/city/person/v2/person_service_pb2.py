"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.person.v2 import motion_pb2 as city_dot_person_dot_v2_dot_motion__pb2
from ....city.person.v2 import person_pb2 as city_dot_person_dot_v2_dot_person__pb2
from ....city.person.v2 import person_runtime_pb2 as city_dot_person_dot_v2_dot_person__runtime__pb2
from ....city.person.v2 import vehicle_pb2 as city_dot_person_dot_v2_dot_vehicle__pb2
from ....city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#city/person/v2/person_service.proto\x12\x0ecity.person.v2\x1a\x15city/geo/v2/geo.proto\x1a\x1bcity/person/v2/motion.proto\x1a\x1bcity/person/v2/person.proto\x1a#city/person/v2/person_runtime.proto\x1a\x1ccity/person/v2/vehicle.proto\x1a\x17city/trip/v2/trip.proto"/\n\x10GetPersonRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId"J\n\x11GetPersonResponse\x125\n\x06person\x18\x01 \x01(\x0b2\x1d.city.person.v2.PersonRuntimeR\x06person"B\n\x10AddPersonRequest\x12.\n\x06person\x18\x01 \x01(\x0b2\x16.city.person.v2.PersonR\x06person"0\n\x11AddPersonResponse\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId"g\n\x12SetScheduleRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId\x124\n\tschedules\x18\x02 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules"\x15\n\x13SetScheduleResponse"\x96\x01\n\x11GetPersonsRequest\x12\x1d\n\nperson_ids\x18\x01 \x03(\x05R\tpersonIds\x12A\n\x10exclude_statuses\x18\x02 \x03(\x0e2\x16.city.person.v2.StatusR\x0fexcludeStatuses\x12\x1f\n\x0breturn_base\x18\x03 \x01(\x08R\nreturnBase"M\n\x12GetPersonsResponse\x127\n\x07persons\x18\x01 \x03(\x0b2\x1d.city.person.v2.PersonRuntimeR\x07persons"\xb1\x01\n\x1dGetPersonByLongLatBBoxRequest\x12,\n\x04bbox\x18\x01 \x01(\x0b2\x18.city.geo.v2.LongLatBBoxR\x04bbox\x12A\n\x10exclude_statuses\x18\x02 \x03(\x0e2\x16.city.person.v2.StatusR\x0fexcludeStatuses\x12\x1f\n\x0breturn_base\x18\x03 \x01(\x08R\nreturnBase"Y\n\x1eGetPersonByLongLatBBoxResponse\x127\n\x07persons\x18\x01 \x03(\x0b2\x1d.city.person.v2.PersonRuntimeR\x07persons"\x17\n\x15GetAllVehiclesRequest"T\n\x16GetAllVehiclesResponse\x12:\n\x08vehicles\x18\x01 \x03(\x0b2\x1e.city.person.v2.VehicleRuntimeR\x08vehicles"l\n\x1aResetPersonPositionRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId\x121\n\x08position\x18\x02 \x01(\x0b2\x15.city.geo.v2.PositionR\x08position"\x1d\n\x1bResetPersonPositionResponse"m\n\x1eSetControlledVehicleIDsRequest\x12\x1f\n\x0bvehicle_ids\x18\x01 \x03(\x05R\nvehicleIds\x12*\n\x11route_vehicle_ids\x18\x02 \x03(\x05R\x0frouteVehicleIds"!\n\x1fSetControlledVehicleIDsResponse"#\n!FetchControlledVehicleEnvsRequest"\xad\x01\n"FetchControlledVehicleEnvsResponse\x12=\n\x0cvehicle_envs\x18\x01 \x03(\x0b2\x1a.city.person.v2.VehicleEnvR\x0bvehicleEnvs\x12H\n\x12route_vehicle_envs\x18\x02 \x03(\x0b2\x1a.city.person.v2.VehicleEnvR\x10routeVehicleEnvs"\xbb\x01\n"SetControlledVehicleActionsRequest\x12F\n\x0fvehicle_actions\x18\x01 \x03(\x0b2\x1d.city.person.v2.VehicleActionR\x0evehicleActions\x12M\n\x10vehicle_journeys\x18\x02 \x03(\x0b2".city.person.v2.VehicleRouteActionR\x0fvehicleJourneys"%\n#SetControlledVehicleActionsResponse2\xb5\x08\n\rPersonService\x12P\n\tGetPerson\x12 .city.person.v2.GetPersonRequest\x1a!.city.person.v2.GetPersonResponse\x12P\n\tAddPerson\x12 .city.person.v2.AddPersonRequest\x1a!.city.person.v2.AddPersonResponse\x12V\n\x0bSetSchedule\x12".city.person.v2.SetScheduleRequest\x1a#.city.person.v2.SetScheduleResponse\x12S\n\nGetPersons\x12!.city.person.v2.GetPersonsRequest\x1a".city.person.v2.GetPersonsResponse\x12w\n\x16GetPersonByLongLatBBox\x12-.city.person.v2.GetPersonByLongLatBBoxRequest\x1a..city.person.v2.GetPersonByLongLatBBoxResponse\x12_\n\x0eGetAllVehicles\x12%.city.person.v2.GetAllVehiclesRequest\x1a&.city.person.v2.GetAllVehiclesResponse\x12n\n\x13ResetPersonPosition\x12*.city.person.v2.ResetPersonPositionRequest\x1a+.city.person.v2.ResetPersonPositionResponse\x12z\n\x17SetControlledVehicleIDs\x12..city.person.v2.SetControlledVehicleIDsRequest\x1a/.city.person.v2.SetControlledVehicleIDsResponse\x12\x83\x01\n\x1aFetchControlledVehicleEnvs\x121.city.person.v2.FetchControlledVehicleEnvsRequest\x1a2.city.person.v2.FetchControlledVehicleEnvsResponse\x12\x86\x01\n\x1bSetControlledVehicleActions\x122.city.person.v2.SetControlledVehicleActionsRequest\x1a3.city.person.v2.SetControlledVehicleActionsResponseB\xbb\x01\n\x12com.city.person.v2B\x12PersonServiceProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v2.person_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v2B\x12PersonServiceProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2'
    _globals['_GETPERSONREQUEST']._serialized_start = 228
    _globals['_GETPERSONREQUEST']._serialized_end = 275
    _globals['_GETPERSONRESPONSE']._serialized_start = 277
    _globals['_GETPERSONRESPONSE']._serialized_end = 351
    _globals['_ADDPERSONREQUEST']._serialized_start = 353
    _globals['_ADDPERSONREQUEST']._serialized_end = 419
    _globals['_ADDPERSONRESPONSE']._serialized_start = 421
    _globals['_ADDPERSONRESPONSE']._serialized_end = 469
    _globals['_SETSCHEDULEREQUEST']._serialized_start = 471
    _globals['_SETSCHEDULEREQUEST']._serialized_end = 574
    _globals['_SETSCHEDULERESPONSE']._serialized_start = 576
    _globals['_SETSCHEDULERESPONSE']._serialized_end = 597
    _globals['_GETPERSONSREQUEST']._serialized_start = 600
    _globals['_GETPERSONSREQUEST']._serialized_end = 750
    _globals['_GETPERSONSRESPONSE']._serialized_start = 752
    _globals['_GETPERSONSRESPONSE']._serialized_end = 829
    _globals['_GETPERSONBYLONGLATBBOXREQUEST']._serialized_start = 832
    _globals['_GETPERSONBYLONGLATBBOXREQUEST']._serialized_end = 1009
    _globals['_GETPERSONBYLONGLATBBOXRESPONSE']._serialized_start = 1011
    _globals['_GETPERSONBYLONGLATBBOXRESPONSE']._serialized_end = 1100
    _globals['_GETALLVEHICLESREQUEST']._serialized_start = 1102
    _globals['_GETALLVEHICLESREQUEST']._serialized_end = 1125
    _globals['_GETALLVEHICLESRESPONSE']._serialized_start = 1127
    _globals['_GETALLVEHICLESRESPONSE']._serialized_end = 1211
    _globals['_RESETPERSONPOSITIONREQUEST']._serialized_start = 1213
    _globals['_RESETPERSONPOSITIONREQUEST']._serialized_end = 1321
    _globals['_RESETPERSONPOSITIONRESPONSE']._serialized_start = 1323
    _globals['_RESETPERSONPOSITIONRESPONSE']._serialized_end = 1352
    _globals['_SETCONTROLLEDVEHICLEIDSREQUEST']._serialized_start = 1354
    _globals['_SETCONTROLLEDVEHICLEIDSREQUEST']._serialized_end = 1463
    _globals['_SETCONTROLLEDVEHICLEIDSRESPONSE']._serialized_start = 1465
    _globals['_SETCONTROLLEDVEHICLEIDSRESPONSE']._serialized_end = 1498
    _globals['_FETCHCONTROLLEDVEHICLEENVSREQUEST']._serialized_start = 1500
    _globals['_FETCHCONTROLLEDVEHICLEENVSREQUEST']._serialized_end = 1535
    _globals['_FETCHCONTROLLEDVEHICLEENVSRESPONSE']._serialized_start = 1538
    _globals['_FETCHCONTROLLEDVEHICLEENVSRESPONSE']._serialized_end = 1711
    _globals['_SETCONTROLLEDVEHICLEACTIONSREQUEST']._serialized_start = 1714
    _globals['_SETCONTROLLEDVEHICLEACTIONSREQUEST']._serialized_end = 1901
    _globals['_SETCONTROLLEDVEHICLEACTIONSRESPONSE']._serialized_start = 1903
    _globals['_SETCONTROLLEDVEHICLEACTIONSRESPONSE']._serialized_end = 1940
    _globals['_PERSONSERVICE']._serialized_start = 1943
    _globals['_PERSONSERVICE']._serialized_end = 3020