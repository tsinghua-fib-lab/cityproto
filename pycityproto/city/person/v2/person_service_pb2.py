"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.person.v2 import motion_pb2 as city_dot_person_dot_v2_dot_motion__pb2
from ....city.person.v2 import pedestrian_pb2 as city_dot_person_dot_v2_dot_pedestrian__pb2
from ....city.person.v2 import person_pb2 as city_dot_person_dot_v2_dot_person__pb2
from ....city.person.v2 import person_runtime_pb2 as city_dot_person_dot_v2_dot_person__runtime__pb2
from ....city.person.v2 import taxi_pb2 as city_dot_person_dot_v2_dot_taxi__pb2
from ....city.person.v2 import vehicle_pb2 as city_dot_person_dot_v2_dot_vehicle__pb2
from ....city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#city/person/v2/person_service.proto\x12\x0ecity.person.v2\x1a\x15city/geo/v2/geo.proto\x1a\x1bcity/person/v2/motion.proto\x1a\x1fcity/person/v2/pedestrian.proto\x1a\x1bcity/person/v2/person.proto\x1a#city/person/v2/person_runtime.proto\x1a\x19city/person/v2/taxi.proto\x1a\x1ccity/person/v2/vehicle.proto\x1a\x17city/trip/v2/trip.proto"/\n\x10GetPersonRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId"J\n\x11GetPersonResponse\x125\n\x06person\x18\x01 \x01(\x0b2\x1d.city.person.v2.PersonRuntimeR\x06person"B\n\x10AddPersonRequest\x12.\n\x06person\x18\x01 \x01(\x0b2\x16.city.person.v2.PersonR\x06person"0\n\x11AddPersonResponse\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId"g\n\x12SetScheduleRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId\x124\n\tschedules\x18\x02 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules"\x15\n\x13SetScheduleResponse"\x96\x01\n\x11GetPersonsRequest\x12\x1d\n\nperson_ids\x18\x01 \x03(\x05R\tpersonIds\x12A\n\x10exclude_statuses\x18\x02 \x03(\x0e2\x16.city.person.v2.StatusR\x0fexcludeStatuses\x12\x1f\n\x0breturn_base\x18\x03 \x01(\x08R\nreturnBase"M\n\x12GetPersonsResponse\x127\n\x07persons\x18\x01 \x03(\x0b2\x1d.city.person.v2.PersonRuntimeR\x07persons"\xb1\x01\n\x1dGetPersonByLongLatBBoxRequest\x12,\n\x04bbox\x18\x01 \x01(\x0b2\x18.city.geo.v2.LongLatBBoxR\x04bbox\x12A\n\x10exclude_statuses\x18\x02 \x03(\x0e2\x16.city.person.v2.StatusR\x0fexcludeStatuses\x12\x1f\n\x0breturn_base\x18\x03 \x01(\x08R\nreturnBase"Y\n\x1eGetPersonByLongLatBBoxResponse\x127\n\x07persons\x18\x01 \x03(\x0b2\x1d.city.person.v2.PersonRuntimeR\x07persons"\x17\n\x15GetAllVehiclesRequest"T\n\x16GetAllVehiclesResponse\x12:\n\x08vehicles\x18\x01 \x03(\x0b2\x1e.city.person.v2.VehicleRuntimeR\x08vehicles"\x1a\n\x18GetAllPedestriansRequest"\\\n\x19GetAllPedestriansResponse\x12?\n\x0bpedestrians\x18\x01 \x03(\x0b2\x1d.city.person.v2.PersonRuntimeR\x0bpedestrians"l\n\x1aResetPersonPositionRequest\x12\x1b\n\tperson_id\x18\x01 \x01(\x05R\x08personId\x121\n\x08position\x18\x02 \x01(\x0b2\x15.city.geo.v2.PositionR\x08position"\x1d\n\x1bResetPersonPositionResponse"m\n\x1eSetControlledVehicleIDsRequest\x12\x1f\n\x0bvehicle_ids\x18\x01 \x03(\x05R\nvehicleIds\x12*\n\x11route_vehicle_ids\x18\x02 \x03(\x05R\x0frouteVehicleIds"!\n\x1fSetControlledVehicleIDsResponse"#\n!FetchControlledVehicleEnvsRequest"\xad\x01\n"FetchControlledVehicleEnvsResponse\x12=\n\x0cvehicle_envs\x18\x01 \x03(\x0b2\x1a.city.person.v2.VehicleEnvR\x0bvehicleEnvs\x12H\n\x12route_vehicle_envs\x18\x02 \x03(\x0b2\x1a.city.person.v2.VehicleEnvR\x10routeVehicleEnvs"\xbb\x01\n"SetControlledVehicleActionsRequest\x12F\n\x0fvehicle_actions\x18\x01 \x03(\x0b2\x1d.city.person.v2.VehicleActionR\x0evehicleActions\x12M\n\x10vehicle_journeys\x18\x02 \x03(\x0b2".city.person.v2.VehicleRouteActionR\x0fvehicleJourneys"%\n#SetControlledVehicleActionsResponse"8\n\x1bSetControlledTaxiIDsRequest\x12\x19\n\x08taxi_ids\x18\x01 \x03(\x05R\x07taxiIds"\x1e\n\x1cSetControlledTaxiIDsResponse"\x15\n\x13GetAllOrdersRequest"Y\n\x14GetAllOrdersResponse\x12A\n\x0border_infos\x18\x01 \x03(\x0b2 .city.person.v2.RequestOrderInfoR\norderInfos"h\n SetControlledTaxiToOrdersRequest\x12D\n\x0border_plans\x18\x01 \x03(\x0b2#.city.person.v2.OrderAllocationPlanR\norderPlans"#\n!SetControlledTaxiToOrdersResponse"H\n\x1fSetControlledPedestriansRequest\x12%\n\x0epedestrian_ids\x18\x01 \x03(\x05R\rpedestrianIds""\n SetControlledPedestriansResponse"\'\n%FetchControlledPedestriansEnvsRequest"p\n&FetchControlledPedestriansEnvsResponse\x12F\n\x0fpedestrian_envs\x18\x01 \x03(\x0b2\x1d.city.person.v2.PedestrianEnvR\x0epedestrianEnvs"y\n&SetControlledPedestriansActionsRequest\x12O\n\x12pedestrian_actions\x18\x01 \x03(\x0b2 .city.person.v2.PedestrianActionR\x11pedestrianActions")\n\'SetControlledPedestriansActionsResponse"-\n+GetControlledTaxiOrderAllocationPlanRequest"\xa4\x02\n,GetControlledTaxiOrderAllocationPlanResponse\x12\x89\x01\n\x15order_allocations_map\x18\x01 \x03(\x0b2U.city.person.v2.GetControlledTaxiOrderAllocationPlanResponse.OrderAllocationsMapEntryR\x13orderAllocationsMap\x1ah\n\x18OrderAllocationsMapEntry\x12\x10\n\x03key\x18\x01 \x01(\x05R\x03key\x126\n\x05value\x18\x02 \x01(\x0b2 .city.person.v2.OrderAllocationsR\x05value:\x028\x01"\xa2\x02\n+SetControlledTaxiOrderAllocationPlanRequest\x12\x88\x01\n\x15order_allocations_map\x18\x01 \x03(\x0b2T.city.person.v2.SetControlledTaxiOrderAllocationPlanRequest.OrderAllocationsMapEntryR\x13orderAllocationsMap\x1ah\n\x18OrderAllocationsMapEntry\x12\x10\n\x03key\x18\x01 \x01(\x05R\x03key\x126\n\x05value\x18\x02 \x01(\x0b2 .city.person.v2.OrderAllocationsR\x05value:\x028\x01".\n,SetControlledTaxiOrderAllocationPlanResponse2\xde\x11\n\rPersonService\x12P\n\tGetPerson\x12 .city.person.v2.GetPersonRequest\x1a!.city.person.v2.GetPersonResponse\x12P\n\tAddPerson\x12 .city.person.v2.AddPersonRequest\x1a!.city.person.v2.AddPersonResponse\x12V\n\x0bSetSchedule\x12".city.person.v2.SetScheduleRequest\x1a#.city.person.v2.SetScheduleResponse\x12S\n\nGetPersons\x12!.city.person.v2.GetPersonsRequest\x1a".city.person.v2.GetPersonsResponse\x12w\n\x16GetPersonByLongLatBBox\x12-.city.person.v2.GetPersonByLongLatBBoxRequest\x1a..city.person.v2.GetPersonByLongLatBBoxResponse\x12_\n\x0eGetAllVehicles\x12%.city.person.v2.GetAllVehiclesRequest\x1a&.city.person.v2.GetAllVehiclesResponse\x12h\n\x11GetAllPedestrians\x12(.city.person.v2.GetAllPedestriansRequest\x1a).city.person.v2.GetAllPedestriansResponse\x12n\n\x13ResetPersonPosition\x12*.city.person.v2.ResetPersonPositionRequest\x1a+.city.person.v2.ResetPersonPositionResponse\x12z\n\x17SetControlledVehicleIDs\x12..city.person.v2.SetControlledVehicleIDsRequest\x1a/.city.person.v2.SetControlledVehicleIDsResponse\x12\x83\x01\n\x1aFetchControlledVehicleEnvs\x121.city.person.v2.FetchControlledVehicleEnvsRequest\x1a2.city.person.v2.FetchControlledVehicleEnvsResponse\x12\x86\x01\n\x1bSetControlledVehicleActions\x122.city.person.v2.SetControlledVehicleActionsRequest\x1a3.city.person.v2.SetControlledVehicleActionsResponse\x12q\n\x14SetControlledTaxiIDs\x12+.city.person.v2.SetControlledTaxiIDsRequest\x1a,.city.person.v2.SetControlledTaxiIDsResponse\x12Y\n\x0cGetAllOrders\x12#.city.person.v2.GetAllOrdersRequest\x1a$.city.person.v2.GetAllOrdersResponse\x12\x80\x01\n\x19SetControlledTaxiToOrders\x120.city.person.v2.SetControlledTaxiToOrdersRequest\x1a1.city.person.v2.SetControlledTaxiToOrdersResponse\x12}\n\x18SetControlledPedestrians\x12/.city.person.v2.SetControlledPedestriansRequest\x1a0.city.person.v2.SetControlledPedestriansResponse\x12\x8f\x01\n\x1eFetchControlledPedestriansEnvs\x125.city.person.v2.FetchControlledPedestriansEnvsRequest\x1a6.city.person.v2.FetchControlledPedestriansEnvsResponse\x12\x92\x01\n\x1fSetControlledPedestriansActions\x126.city.person.v2.SetControlledPedestriansActionsRequest\x1a7.city.person.v2.SetControlledPedestriansActionsResponse\x12\xa1\x01\n$GetControlledTaxiOrderAllocationPlan\x12;.city.person.v2.GetControlledTaxiOrderAllocationPlanRequest\x1a<.city.person.v2.GetControlledTaxiOrderAllocationPlanResponse\x12\xa1\x01\n$SetControlledTaxiOrderAllocationPlan\x12;.city.person.v2.SetControlledTaxiOrderAllocationPlanRequest\x1a<.city.person.v2.SetControlledTaxiOrderAllocationPlanResponseB\xbb\x01\n\x12com.city.person.v2B\x12PersonServiceProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v2.person_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v2B\x12PersonServiceProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2'
    _globals['_GETCONTROLLEDTAXIORDERALLOCATIONPLANRESPONSE_ORDERALLOCATIONSMAPENTRY']._options = None
    _globals['_GETCONTROLLEDTAXIORDERALLOCATIONPLANRESPONSE_ORDERALLOCATIONSMAPENTRY']._serialized_options = b'8\x01'
    _globals['_SETCONTROLLEDTAXIORDERALLOCATIONPLANREQUEST_ORDERALLOCATIONSMAPENTRY']._options = None
    _globals['_SETCONTROLLEDTAXIORDERALLOCATIONPLANREQUEST_ORDERALLOCATIONSMAPENTRY']._serialized_options = b'8\x01'
    _globals['_GETPERSONREQUEST']._serialized_start = 288
    _globals['_GETPERSONREQUEST']._serialized_end = 335
    _globals['_GETPERSONRESPONSE']._serialized_start = 337
    _globals['_GETPERSONRESPONSE']._serialized_end = 411
    _globals['_ADDPERSONREQUEST']._serialized_start = 413
    _globals['_ADDPERSONREQUEST']._serialized_end = 479
    _globals['_ADDPERSONRESPONSE']._serialized_start = 481
    _globals['_ADDPERSONRESPONSE']._serialized_end = 529
    _globals['_SETSCHEDULEREQUEST']._serialized_start = 531
    _globals['_SETSCHEDULEREQUEST']._serialized_end = 634
    _globals['_SETSCHEDULERESPONSE']._serialized_start = 636
    _globals['_SETSCHEDULERESPONSE']._serialized_end = 657
    _globals['_GETPERSONSREQUEST']._serialized_start = 660
    _globals['_GETPERSONSREQUEST']._serialized_end = 810
    _globals['_GETPERSONSRESPONSE']._serialized_start = 812
    _globals['_GETPERSONSRESPONSE']._serialized_end = 889
    _globals['_GETPERSONBYLONGLATBBOXREQUEST']._serialized_start = 892
    _globals['_GETPERSONBYLONGLATBBOXREQUEST']._serialized_end = 1069
    _globals['_GETPERSONBYLONGLATBBOXRESPONSE']._serialized_start = 1071
    _globals['_GETPERSONBYLONGLATBBOXRESPONSE']._serialized_end = 1160
    _globals['_GETALLVEHICLESREQUEST']._serialized_start = 1162
    _globals['_GETALLVEHICLESREQUEST']._serialized_end = 1185
    _globals['_GETALLVEHICLESRESPONSE']._serialized_start = 1187
    _globals['_GETALLVEHICLESRESPONSE']._serialized_end = 1271
    _globals['_GETALLPEDESTRIANSREQUEST']._serialized_start = 1273
    _globals['_GETALLPEDESTRIANSREQUEST']._serialized_end = 1299
    _globals['_GETALLPEDESTRIANSRESPONSE']._serialized_start = 1301
    _globals['_GETALLPEDESTRIANSRESPONSE']._serialized_end = 1393
    _globals['_RESETPERSONPOSITIONREQUEST']._serialized_start = 1395
    _globals['_RESETPERSONPOSITIONREQUEST']._serialized_end = 1503
    _globals['_RESETPERSONPOSITIONRESPONSE']._serialized_start = 1505
    _globals['_RESETPERSONPOSITIONRESPONSE']._serialized_end = 1534
    _globals['_SETCONTROLLEDVEHICLEIDSREQUEST']._serialized_start = 1536
    _globals['_SETCONTROLLEDVEHICLEIDSREQUEST']._serialized_end = 1645
    _globals['_SETCONTROLLEDVEHICLEIDSRESPONSE']._serialized_start = 1647
    _globals['_SETCONTROLLEDVEHICLEIDSRESPONSE']._serialized_end = 1680
    _globals['_FETCHCONTROLLEDVEHICLEENVSREQUEST']._serialized_start = 1682
    _globals['_FETCHCONTROLLEDVEHICLEENVSREQUEST']._serialized_end = 1717
    _globals['_FETCHCONTROLLEDVEHICLEENVSRESPONSE']._serialized_start = 1720
    _globals['_FETCHCONTROLLEDVEHICLEENVSRESPONSE']._serialized_end = 1893
    _globals['_SETCONTROLLEDVEHICLEACTIONSREQUEST']._serialized_start = 1896
    _globals['_SETCONTROLLEDVEHICLEACTIONSREQUEST']._serialized_end = 2083
    _globals['_SETCONTROLLEDVEHICLEACTIONSRESPONSE']._serialized_start = 2085
    _globals['_SETCONTROLLEDVEHICLEACTIONSRESPONSE']._serialized_end = 2122
    _globals['_SETCONTROLLEDTAXIIDSREQUEST']._serialized_start = 2124
    _globals['_SETCONTROLLEDTAXIIDSREQUEST']._serialized_end = 2180
    _globals['_SETCONTROLLEDTAXIIDSRESPONSE']._serialized_start = 2182
    _globals['_SETCONTROLLEDTAXIIDSRESPONSE']._serialized_end = 2212
    _globals['_GETALLORDERSREQUEST']._serialized_start = 2214
    _globals['_GETALLORDERSREQUEST']._serialized_end = 2235
    _globals['_GETALLORDERSRESPONSE']._serialized_start = 2237
    _globals['_GETALLORDERSRESPONSE']._serialized_end = 2326
    _globals['_SETCONTROLLEDTAXITOORDERSREQUEST']._serialized_start = 2328
    _globals['_SETCONTROLLEDTAXITOORDERSREQUEST']._serialized_end = 2432
    _globals['_SETCONTROLLEDTAXITOORDERSRESPONSE']._serialized_start = 2434
    _globals['_SETCONTROLLEDTAXITOORDERSRESPONSE']._serialized_end = 2469
    _globals['_SETCONTROLLEDPEDESTRIANSREQUEST']._serialized_start = 2471
    _globals['_SETCONTROLLEDPEDESTRIANSREQUEST']._serialized_end = 2543
    _globals['_SETCONTROLLEDPEDESTRIANSRESPONSE']._serialized_start = 2545
    _globals['_SETCONTROLLEDPEDESTRIANSRESPONSE']._serialized_end = 2579
    _globals['_FETCHCONTROLLEDPEDESTRIANSENVSREQUEST']._serialized_start = 2581
    _globals['_FETCHCONTROLLEDPEDESTRIANSENVSREQUEST']._serialized_end = 2620
    _globals['_FETCHCONTROLLEDPEDESTRIANSENVSRESPONSE']._serialized_start = 2622
    _globals['_FETCHCONTROLLEDPEDESTRIANSENVSRESPONSE']._serialized_end = 2734
    _globals['_SETCONTROLLEDPEDESTRIANSACTIONSREQUEST']._serialized_start = 2736
    _globals['_SETCONTROLLEDPEDESTRIANSACTIONSREQUEST']._serialized_end = 2857
    _globals['_SETCONTROLLEDPEDESTRIANSACTIONSRESPONSE']._serialized_start = 2859
    _globals['_SETCONTROLLEDPEDESTRIANSACTIONSRESPONSE']._serialized_end = 2900
    _globals['_GETCONTROLLEDTAXIORDERALLOCATIONPLANREQUEST']._serialized_start = 2902
    _globals['_GETCONTROLLEDTAXIORDERALLOCATIONPLANREQUEST']._serialized_end = 2947
    _globals['_GETCONTROLLEDTAXIORDERALLOCATIONPLANRESPONSE']._serialized_start = 2950
    _globals['_GETCONTROLLEDTAXIORDERALLOCATIONPLANRESPONSE']._serialized_end = 3242
    _globals['_GETCONTROLLEDTAXIORDERALLOCATIONPLANRESPONSE_ORDERALLOCATIONSMAPENTRY']._serialized_start = 3138
    _globals['_GETCONTROLLEDTAXIORDERALLOCATIONPLANRESPONSE_ORDERALLOCATIONSMAPENTRY']._serialized_end = 3242
    _globals['_SETCONTROLLEDTAXIORDERALLOCATIONPLANREQUEST']._serialized_start = 3245
    _globals['_SETCONTROLLEDTAXIORDERALLOCATIONPLANREQUEST']._serialized_end = 3535
    _globals['_SETCONTROLLEDTAXIORDERALLOCATIONPLANREQUEST_ORDERALLOCATIONSMAPENTRY']._serialized_start = 3138
    _globals['_SETCONTROLLEDTAXIORDERALLOCATIONPLANREQUEST_ORDERALLOCATIONSMAPENTRY']._serialized_end = 3242
    _globals['_SETCONTROLLEDTAXIORDERALLOCATIONPLANRESPONSE']._serialized_start = 3537
    _globals['_SETCONTROLLEDTAXIORDERALLOCATIONPLANRESPONSE']._serialized_end = 3583
    _globals['_PERSONSERVICE']._serialized_start = 3586
    _globals['_PERSONSERVICE']._serialized_end = 5856