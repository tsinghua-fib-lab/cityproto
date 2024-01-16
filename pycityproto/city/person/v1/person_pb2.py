"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bcity/person/v1/person.proto\x12\x0ecity.person.v1\x1a\x15city/geo/v2/geo.proto\x1a\x17city/trip/v2/trip.proto"\xae\x02\n\x0fPersonAttribute\x12\x16\n\x06length\x18\x02 \x01(\x01R\x06length\x12\x14\n\x05width\x18\x03 \x01(\x01R\x05width\x12\x1b\n\tmax_speed\x18\x04 \x01(\x01R\x08maxSpeed\x12)\n\x10max_acceleration\x18\x05 \x01(\x01R\x0fmaxAcceleration\x128\n\x18max_braking_acceleration\x18\x06 \x01(\x01R\x16maxBrakingAcceleration\x12-\n\x12usual_acceleration\x18\x07 \x01(\x01R\x11usualAcceleration\x12<\n\x1ausual_braking_acceleration\x18\x08 \x01(\x01R\x18usualBrakingAcceleration"Y\n\x10VehicleAttribute\x12,\n\x12lane_change_length\x18\x01 \x01(\x01R\x10laneChangeLength\x12\x17\n\x07min_gap\x18\x02 \x01(\x01R\x06minGap"C\n\x0cBusAttribute\x12\x17\n\x07line_id\x18\x01 \x01(\x05R\x06lineId\x12\x1a\n\x08capacity\x18\x02 \x01(\x05R\x08capacity"\x0f\n\rBikeAttribute"\xd1\x04\n\x06Person\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12=\n\tattribute\x18\x02 \x01(\x0b2\x1f.city.person.v1.PersonAttributeR\tattribute\x12)\n\x04home\x18\x03 \x01(\x0b2\x15.city.geo.v2.PositionR\x04home\x124\n\tschedules\x18\x04 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules\x12R\n\x11vehicle_attribute\x18\x07 \x01(\x0b2 .city.person.v1.VehicleAttributeH\x00R\x10vehicleAttribute\x88\x01\x01\x12F\n\rbus_attribute\x18\x08 \x01(\x0b2\x1c.city.person.v1.BusAttributeH\x01R\x0cbusAttribute\x88\x01\x01\x12I\n\x0ebike_attribute\x18\t \x01(\x0b2\x1d.city.person.v1.BikeAttributeH\x02R\rbikeAttribute\x88\x01\x01\x12:\n\x06labels\x18\n \x03(\x0b2".city.person.v1.Person.LabelsEntryR\x06labels\x1a9\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\x14\n\x12_vehicle_attributeB\x10\n\x0e_bus_attributeB\x11\n\x0f_bike_attribute";\n\x07Persons\x120\n\x07persons\x18\x01 \x03(\x0b2\x16.city.person.v1.PersonR\x07personsB\xb1\x01\n\x12com.city.person.v1B\x0bPersonProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v1.person_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v1B\x0bPersonProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1'
    _globals['_PERSON_LABELSENTRY']._options = None
    _globals['_PERSON_LABELSENTRY']._serialized_options = b'8\x01'
    _globals['_PERSONATTRIBUTE']._serialized_start = 96
    _globals['_PERSONATTRIBUTE']._serialized_end = 398
    _globals['_VEHICLEATTRIBUTE']._serialized_start = 400
    _globals['_VEHICLEATTRIBUTE']._serialized_end = 489
    _globals['_BUSATTRIBUTE']._serialized_start = 491
    _globals['_BUSATTRIBUTE']._serialized_end = 558
    _globals['_BIKEATTRIBUTE']._serialized_start = 560
    _globals['_BIKEATTRIBUTE']._serialized_end = 575
    _globals['_PERSON']._serialized_start = 578
    _globals['_PERSON']._serialized_end = 1171
    _globals['_PERSON_LABELSENTRY']._serialized_start = 1055
    _globals['_PERSON_LABELSENTRY']._serialized_end = 1112
    _globals['_PERSONS']._serialized_start = 1173
    _globals['_PERSONS']._serialized_end = 1232