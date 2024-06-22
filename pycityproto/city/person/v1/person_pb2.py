"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bcity/person/v1/person.proto\x12\x0ecity.person.v1\x1a\x15city/geo/v2/geo.proto\x1a\x17city/trip/v2/trip.proto"\xae\x02\n\x0fPersonAttribute\x12\x16\n\x06length\x18\x02 \x01(\x01R\x06length\x12\x14\n\x05width\x18\x03 \x01(\x01R\x05width\x12\x1b\n\tmax_speed\x18\x04 \x01(\x01R\x08maxSpeed\x12)\n\x10max_acceleration\x18\x05 \x01(\x01R\x0fmaxAcceleration\x128\n\x18max_braking_acceleration\x18\x06 \x01(\x01R\x16maxBrakingAcceleration\x12-\n\x12usual_acceleration\x18\x07 \x01(\x01R\x11usualAcceleration\x12<\n\x1ausual_braking_acceleration\x18\x08 \x01(\x01R\x18usualBrakingAcceleration"~\n\x10VehicleAttribute\x12,\n\x12lane_change_length\x18\x01 \x01(\x01R\x10laneChangeLength\x12\x17\n\x07min_gap\x18\x02 \x01(\x01R\x06minGap\x12\x19\n\x05model\x18\x03 \x01(\tH\x00R\x05model\x88\x01\x01B\x08\n\x06_model"h\n\x0cBusAttribute\x12\x17\n\x07line_id\x18\x01 \x01(\x05R\x06lineId\x12\x1a\n\x08capacity\x18\x02 \x01(\x05R\x08capacity\x12\x19\n\x05model\x18\x03 \x01(\tH\x00R\x05model\x88\x01\x01B\x08\n\x06_model"P\n\x13PedestrianAttribute\x12\x14\n\x05speed\x18\x01 \x01(\x01R\x05speed\x12\x19\n\x05model\x18\x02 \x01(\tH\x00R\x05model\x88\x01\x01B\x08\n\x06_model"J\n\rBikeAttribute\x12\x14\n\x05speed\x18\x01 \x01(\x01R\x05speed\x12\x19\n\x05model\x18\x02 \x01(\tH\x00R\x05model\x88\x01\x01B\x08\n\x06_model"\xc9\x01\n\rPersonProfile\x12\x10\n\x03age\x18\x01 \x01(\x05R\x03age\x127\n\teducation\x18\x02 \x01(\x0e2\x19.city.person.v1.EducationR\teducation\x12.\n\x06gender\x18\x03 \x01(\x0e2\x16.city.person.v1.GenderR\x06gender\x12=\n\x0bconsumption\x18\x04 \x01(\x0e2\x1b.city.person.v1.ConsumptionR\x0bconsumption"\xca\x06\n\x06Person\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12=\n\tattribute\x18\x02 \x01(\x0b2\x1f.city.person.v1.PersonAttributeR\tattribute\x12)\n\x04home\x18\x03 \x01(\x0b2\x15.city.geo.v2.PositionR\x04home\x124\n\tschedules\x18\x04 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules\x12R\n\x11vehicle_attribute\x18\x07 \x01(\x0b2 .city.person.v1.VehicleAttributeH\x00R\x10vehicleAttribute\x88\x01\x01\x12F\n\rbus_attribute\x18\x08 \x01(\x0b2\x1c.city.person.v1.BusAttributeH\x01R\x0cbusAttribute\x88\x01\x01\x12[\n\x14pedestrian_attribute\x18\x0c \x01(\x0b2#.city.person.v1.PedestrianAttributeH\x02R\x13pedestrianAttribute\x88\x01\x01\x12I\n\x0ebike_attribute\x18\t \x01(\x0b2\x1d.city.person.v1.BikeAttributeH\x03R\rbikeAttribute\x88\x01\x01\x12:\n\x06labels\x18\n \x03(\x0b2".city.person.v1.Person.LabelsEntryR\x06labels\x12<\n\x07profile\x18\x0b \x01(\x0b2\x1d.city.person.v1.PersonProfileH\x04R\x07profile\x88\x01\x01\x12.\n\x04work\x18\r \x01(\x0b2\x15.city.geo.v2.PositionH\x05R\x04work\x88\x01\x01\x1a9\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\x14\n\x12_vehicle_attributeB\x10\n\x0e_bus_attributeB\x17\n\x15_pedestrian_attributeB\x11\n\x0f_bike_attributeB\n\n\x08_profileB\x07\n\x05_work";\n\x07Persons\x120\n\x07persons\x18\x01 \x03(\x0b2\x16.city.person.v1.PersonR\x07persons*\xdc\x01\n\tEducation\x12\x19\n\x15EDUCATION_UNSPECIFIED\x10\x00\x12\x14\n\x10EDUCATION_DOCTOR\x10\x01\x12\x14\n\x10EDUCATION_MASTER\x10\x02\x12\x16\n\x12EDUCATION_BACHELOR\x10\x03\x12\x19\n\x15EDUCATION_HIGH_SCHOOL\x10\x04\x12 \n\x1cEDUCATION_JUNIOR_HIGH_SCHOOL\x10\x05\x12\x1c\n\x18EDUCATION_PRIMARY_SCHOOL\x10\x06\x12\x15\n\x11EDUCATION_COLLEGE\x10\x07*D\n\x06Gender\x12\x16\n\x12GENDER_UNSPECIFIED\x10\x00\x12\x0f\n\x0bGENDER_MALE\x10\x01\x12\x11\n\rGENDER_FEMALE\x10\x02*\xae\x01\n\x0bConsumption\x12\x1b\n\x17CONSUMPTION_UNSPECIFIED\x10\x00\x12\x13\n\x0fCONSUMPTION_LOW\x10\x01\x12\x1e\n\x1aCONSUMPTION_RELATIVELY_LOW\x10\x02\x12\x16\n\x12CONSUMPTION_MEDIUM\x10\x03\x12\x1f\n\x1bCONSUMPTION_RELATIVELY_HIGH\x10\x04\x12\x14\n\x10CONSUMPTION_HIGH\x10\x05B\xb1\x01\n\x12com.city.person.v1B\x0bPersonProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v1.person_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v1B\x0bPersonProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1'
    _globals['_PERSON_LABELSENTRY']._options = None
    _globals['_PERSON_LABELSENTRY']._serialized_options = b'8\x01'
    _globals['_EDUCATION']._serialized_start = 1903
    _globals['_EDUCATION']._serialized_end = 2123
    _globals['_GENDER']._serialized_start = 2125
    _globals['_GENDER']._serialized_end = 2193
    _globals['_CONSUMPTION']._serialized_start = 2196
    _globals['_CONSUMPTION']._serialized_end = 2370
    _globals['_PERSONATTRIBUTE']._serialized_start = 96
    _globals['_PERSONATTRIBUTE']._serialized_end = 398
    _globals['_VEHICLEATTRIBUTE']._serialized_start = 400
    _globals['_VEHICLEATTRIBUTE']._serialized_end = 526
    _globals['_BUSATTRIBUTE']._serialized_start = 528
    _globals['_BUSATTRIBUTE']._serialized_end = 632
    _globals['_PEDESTRIANATTRIBUTE']._serialized_start = 634
    _globals['_PEDESTRIANATTRIBUTE']._serialized_end = 714
    _globals['_BIKEATTRIBUTE']._serialized_start = 716
    _globals['_BIKEATTRIBUTE']._serialized_end = 790
    _globals['_PERSONPROFILE']._serialized_start = 793
    _globals['_PERSONPROFILE']._serialized_end = 994
    _globals['_PERSON']._serialized_start = 997
    _globals['_PERSON']._serialized_end = 1839
    _globals['_PERSON_LABELSENTRY']._serialized_start = 1677
    _globals['_PERSON_LABELSENTRY']._serialized_end = 1734
    _globals['_PERSONS']._serialized_start = 1841
    _globals['_PERSONS']._serialized_end = 1900