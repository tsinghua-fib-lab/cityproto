"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
from ....city.trip.v2 import trip_pb2 as city_dot_trip_dot_v2_dot_trip__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bcity/person/v2/person.proto\x12\x0ecity.person.v2\x1a\x15city/geo/v2/geo.proto\x1a\x17city/trip/v2/trip.proto"\x11\n\x0fPersonAttribute"n\n\x17VehicleEngineEfficiency\x12@\n\x1cenergy_conversion_efficiency\x18\x01 \x01(\x01R\x1aenergyConversionEfficiency\x12\x11\n\x04c_ef\x18\x02 \x01(\x01R\x03cEf"\xad\x03\n\x11EmissionAttribute\x12\x16\n\x06weight\x18\x01 \x01(\x01R\x06weight\x125\n\x04type\x18\x02 \x01(\x0e2!.city.person.v2.VehicleEngineTypeR\x04type\x12)\n\x10coefficient_drag\x18\x03 \x01(\x01R\x0fcoefficientDrag\x12\x19\n\x08lambda_s\x18\x04 \x01(\x01R\x07lambdaS\x12!\n\x0cfrontal_area\x18\x05 \x01(\x01R\x0bfrontalArea\x12U\n\x0ffuel_efficiency\x18\x06 \x01(\x0b2\'.city.person.v2.VehicleEngineEfficiencyH\x00R\x0efuelEfficiency\x88\x01\x01\x12]\n\x13electric_efficiency\x18\x07 \x01(\x0b2\'.city.person.v2.VehicleEngineEfficiencyH\x01R\x12electricEfficiency\x88\x01\x01B\x12\n\x10_fuel_efficiencyB\x16\n\x14_electric_efficiency"\xf3\x04\n\x10VehicleAttribute\x12\x16\n\x06length\x18\x01 \x01(\x01R\x06length\x12\x14\n\x05width\x18\x02 \x01(\x01R\x05width\x12\x1b\n\tmax_speed\x18\x03 \x01(\x01R\x08maxSpeed\x12)\n\x10max_acceleration\x18\x04 \x01(\x01R\x0fmaxAcceleration\x128\n\x18max_braking_acceleration\x18\x05 \x01(\x01R\x16maxBrakingAcceleration\x12-\n\x12usual_acceleration\x18\x06 \x01(\x01R\x11usualAcceleration\x12<\n\x1ausual_braking_acceleration\x18\x07 \x01(\x01R\x18usualBrakingAcceleration\x12,\n\x12lane_change_length\x18\x08 \x01(\x01R\x10laneChangeLength\x12\x17\n\x07min_gap\x18\t \x01(\x01R\x06minGap\x12\x18\n\x07headway\x18\n \x01(\x01R\x07headway\x12\x19\n\x05model\x18\x0b \x01(\tH\x00R\x05model\x88\x01\x01\x12N\n$lane_max_speed_recognition_deviation\x18\x0c \x01(\x01R laneMaxSpeedRecognitionDeviation\x12P\n\x12emission_attribute\x18\r \x01(\x0b2!.city.person.v2.EmissionAttributeR\x11emissionAttribute\x12\x1a\n\x08capacity\x18\x0e \x01(\x05R\x08capacityB\x08\n\x06_model"v\n\x0cBusAttribute\x12\x1d\n\nsubline_id\x18\x01 \x01(\x05R\tsublineId\x12\x1a\n\x08capacity\x18\x02 \x01(\x05R\x08capacity\x12+\n\x04type\x18\x05 \x01(\x0e2\x17.city.person.v2.BusTypeR\x04type"P\n\x13PedestrianAttribute\x12\x14\n\x05speed\x18\x01 \x01(\x01R\x05speed\x12\x19\n\x05model\x18\x02 \x01(\tH\x00R\x05model\x88\x01\x01B\x08\n\x06_model"J\n\rBikeAttribute\x12\x14\n\x05speed\x18\x01 \x01(\x01R\x05speed\x12\x19\n\x05model\x18\x02 \x01(\tH\x00R\x05model\x88\x01\x01B\x08\n\x06_model"\xe4\x01\n\rPersonProfile\x12\x10\n\x03age\x18\x01 \x01(\x05R\x03age\x127\n\teducation\x18\x02 \x01(\x0e2\x19.city.person.v2.EducationR\teducation\x12.\n\x06gender\x18\x03 \x01(\x0e2\x16.city.person.v2.GenderR\x06gender\x12=\n\x0bconsumption\x18\x04 \x01(\x0e2\x1b.city.person.v2.ConsumptionR\x0bconsumption\x12\x19\n\x08house_id\x18\x05 \x01(\x05R\x07houseId"\xc1\x07\n\x06Person\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12=\n\tattribute\x18\x02 \x01(\x0b2\x1f.city.person.v2.PersonAttributeR\tattribute\x12)\n\x04home\x18\x03 \x01(\x0b2\x15.city.geo.v2.PositionR\x04home\x124\n\tschedules\x18\x04 \x03(\x0b2\x16.city.trip.v2.ScheduleR\tschedules\x12R\n\x11vehicle_attribute\x18\x07 \x01(\x0b2 .city.person.v2.VehicleAttributeH\x00R\x10vehicleAttribute\x88\x01\x01\x12F\n\rbus_attribute\x18\x08 \x01(\x0b2\x1c.city.person.v2.BusAttributeH\x01R\x0cbusAttribute\x88\x01\x01\x12[\n\x14pedestrian_attribute\x18\x0c \x01(\x0b2#.city.person.v2.PedestrianAttributeH\x02R\x13pedestrianAttribute\x88\x01\x01\x12I\n\x0ebike_attribute\x18\t \x01(\x0b2\x1d.city.person.v2.BikeAttributeH\x03R\rbikeAttribute\x88\x01\x01\x12:\n\x06labels\x18\n \x03(\x0b2".city.person.v2.Person.LabelsEntryR\x06labels\x12<\n\x07profile\x18\x0b \x01(\x0b2\x1d.city.person.v2.PersonProfileH\x04R\x07profile\x88\x01\x01\x12.\n\x04work\x18\r \x01(\x0b2\x15.city.geo.v2.PositionH\x05R\x04work\x88\x01\x01\x12/\n\x11output_when_sleep\x18\x0e \x01(\x08H\x06R\x0foutputWhenSleep\x88\x01\x01\x12.\n\x04type\x18\x0f \x01(\x0e2\x1a.city.person.v2.PersonTypeR\x04type\x1a9\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\x14\n\x12_vehicle_attributeB\x10\n\x0e_bus_attributeB\x17\n\x15_pedestrian_attributeB\x11\n\x0f_bike_attributeB\n\n\x08_profileB\x07\n\x05_workB\x14\n\x12_output_when_sleep";\n\x07Persons\x120\n\x07persons\x18\x01 \x03(\x0b2\x16.city.person.v2.PersonR\x07persons*\x98\x01\n\x11VehicleEngineType\x12#\n\x1fVEHICLE_ENGINE_TYPE_UNSPECIFIED\x10\x00\x12\x1c\n\x18VEHICLE_ENGINE_TYPE_FUEL\x10\x01\x12 \n\x1cVEHICLE_ENGINE_TYPE_ELECTRIC\x10\x02\x12\x1e\n\x1aVEHICLE_ENGINE_TYPE_HYBRID\x10\x03*J\n\x07BusType\x12\x18\n\x14BUS_TYPE_UNSPECIFIED\x10\x00\x12\x10\n\x0cBUS_TYPE_BUS\x10\x01\x12\x13\n\x0fBUS_TYPE_SUBWAY\x10\x02*\xdc\x01\n\tEducation\x12\x19\n\x15EDUCATION_UNSPECIFIED\x10\x00\x12\x14\n\x10EDUCATION_DOCTOR\x10\x01\x12\x14\n\x10EDUCATION_MASTER\x10\x02\x12\x16\n\x12EDUCATION_BACHELOR\x10\x03\x12\x19\n\x15EDUCATION_HIGH_SCHOOL\x10\x04\x12 \n\x1cEDUCATION_JUNIOR_HIGH_SCHOOL\x10\x05\x12\x1c\n\x18EDUCATION_PRIMARY_SCHOOL\x10\x06\x12\x15\n\x11EDUCATION_COLLEGE\x10\x07*D\n\x06Gender\x12\x16\n\x12GENDER_UNSPECIFIED\x10\x00\x12\x0f\n\x0bGENDER_MALE\x10\x01\x12\x11\n\rGENDER_FEMALE\x10\x02*\xae\x01\n\x0bConsumption\x12\x1b\n\x17CONSUMPTION_UNSPECIFIED\x10\x00\x12\x13\n\x0fCONSUMPTION_LOW\x10\x01\x12\x1e\n\x1aCONSUMPTION_RELATIVELY_LOW\x10\x02\x12\x16\n\x12CONSUMPTION_MEDIUM\x10\x03\x12\x1f\n\x1bCONSUMPTION_RELATIVELY_HIGH\x10\x04\x12\x14\n\x10CONSUMPTION_HIGH\x10\x05*W\n\nPersonType\x12\x1b\n\x17PERSON_TYPE_UNSPECIFIED\x10\x00\x12\x14\n\x10PERSON_TYPE_TAXI\x10\x01\x12\x16\n\x12PERSON_TYPE_NORMAL\x10\x02B\xb4\x01\n\x12com.city.person.v2B\x0bPersonProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v2.person_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v2B\x0bPersonProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2'
    _globals['_PERSON_LABELSENTRY']._options = None
    _globals['_PERSON_LABELSENTRY']._serialized_options = b'8\x01'
    _globals['_VEHICLEENGINETYPE']._serialized_start = 2823
    _globals['_VEHICLEENGINETYPE']._serialized_end = 2975
    _globals['_BUSTYPE']._serialized_start = 2977
    _globals['_BUSTYPE']._serialized_end = 3051
    _globals['_EDUCATION']._serialized_start = 3054
    _globals['_EDUCATION']._serialized_end = 3274
    _globals['_GENDER']._serialized_start = 3276
    _globals['_GENDER']._serialized_end = 3344
    _globals['_CONSUMPTION']._serialized_start = 3347
    _globals['_CONSUMPTION']._serialized_end = 3521
    _globals['_PERSONTYPE']._serialized_start = 3523
    _globals['_PERSONTYPE']._serialized_end = 3610
    _globals['_PERSONATTRIBUTE']._serialized_start = 95
    _globals['_PERSONATTRIBUTE']._serialized_end = 112
    _globals['_VEHICLEENGINEEFFICIENCY']._serialized_start = 114
    _globals['_VEHICLEENGINEEFFICIENCY']._serialized_end = 224
    _globals['_EMISSIONATTRIBUTE']._serialized_start = 227
    _globals['_EMISSIONATTRIBUTE']._serialized_end = 656
    _globals['_VEHICLEATTRIBUTE']._serialized_start = 659
    _globals['_VEHICLEATTRIBUTE']._serialized_end = 1286
    _globals['_BUSATTRIBUTE']._serialized_start = 1288
    _globals['_BUSATTRIBUTE']._serialized_end = 1406
    _globals['_PEDESTRIANATTRIBUTE']._serialized_start = 1408
    _globals['_PEDESTRIANATTRIBUTE']._serialized_end = 1488
    _globals['_BIKEATTRIBUTE']._serialized_start = 1490
    _globals['_BIKEATTRIBUTE']._serialized_end = 1564
    _globals['_PERSONPROFILE']._serialized_start = 1567
    _globals['_PERSONPROFILE']._serialized_end = 1795
    _globals['_PERSON']._serialized_start = 1798
    _globals['_PERSON']._serialized_end = 2759
    _globals['_PERSON_LABELSENTRY']._serialized_start = 2575
    _globals['_PERSON_LABELSENTRY']._serialized_end = 2632
    _globals['_PERSONS']._serialized_start = 2761
    _globals['_PERSONS']._serialized_end = 2820