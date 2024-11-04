"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.person.v2 import carbon_pb2 as city_dot_person_dot_v2_dot_carbon__pb2
from ....city.person.v2 import motion_pb2 as city_dot_person_dot_v2_dot_motion__pb2
from ....city.routing.v2 import routing_pb2 as city_dot_routing_dot_v2_dot_routing__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ccity/person/v2/vehicle.proto\x12\x0ecity.person.v2\x1a\x1bcity/person/v2/carbon.proto\x1a\x1bcity/person/v2/motion.proto\x1a\x1dcity/routing/v2/routing.proto"\x84\x01\n\x02LC\x12$\n\x0eshadow_lane_id\x18\x01 \x01(\x05R\x0cshadowLaneId\x12\x19\n\x08shadow_s\x18\x02 \x01(\x01R\x07shadowS\x12\x14\n\x05angle\x18\x03 \x01(\x01R\x05angle\x12\'\n\x0fcompleted_ratio\x18\x04 \x01(\x01R\x0ecompletedRatio"\x7f\n\rVehicleAction\x12\x0e\n\x02id\x18\x04 \x01(\x05R\x02id\x12\x10\n\x03acc\x18\x01 \x01(\x01R\x03acc\x12%\n\x0clc_target_id\x18\x02 \x01(\x05H\x00R\nlcTargetId\x88\x01\x01\x12\x14\n\x05angle\x18\x03 \x01(\x01R\x05angleB\x0f\n\r_lc_target_id"X\n\x12VehicleRouteAction\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x122\n\x07journey\x18\x02 \x01(\x0b2\x18.city.routing.v2.JourneyR\x07journey"\xb2\x03\n\x0eVehicleRuntime\x120\n\x04base\x18\x01 \x01(\x0b2\x1c.city.person.v2.PersonMotionR\x04base\x12\'\n\x02lc\x18\x04 \x01(\x0b2\x12.city.person.v2.LCH\x00R\x02lc\x88\x01\x01\x12:\n\x06action\x18\x05 \x01(\x0b2\x1d.city.person.v2.VehicleActionH\x01R\x06action\x88\x01\x01\x12)\n\x10running_distance\x18\x06 \x01(\x01R\x0frunningDistance\x12(\n\x10num_going_astray\x18\x07 \x01(\x05R\x0enumGoingAstray\x12%\n\x0edeparture_time\x18\x08 \x01(\x01R\rdepartureTime\x12\x10\n\x03eta\x18\t \x01(\x01R\x03eta\x12"\n\reta_free_flow\x18\n \x01(\x01R\x0betaFreeFlow\x12:\n\x06carbon\x18\x0b \x01(\x0b2\x1d.city.person.v2.VehicleCarbonH\x02R\x06carbon\x88\x01\x01B\x05\n\x03_lcB\t\n\x07_actionB\t\n\x07_carbon"\xc1\x01\n\x0fObservedVehicle\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x124\n\x06motion\x18\x02 \x01(\x0b2\x1c.city.person.v2.PersonMotionR\x06motion\x12+\n\x11relative_distance\x18\x03 \x01(\x01R\x10relativeDistance\x12;\n\x08relation\x18\x04 \x01(\x0e2\x1f.city.person.v2.VehicleRelationR\x08relation"\xaf\x01\n\x0cObservedLane\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12 \n\x0brestriction\x18\x02 \x01(\x08R\x0brestriction\x12;\n\x0blight_state\x18\x03 \x01(\x0e2\x1a.city.person.v2.LightStateR\nlightState\x120\n\x14light_remaining_time\x18\x04 \x01(\x01R\x12lightRemainingTime"\x9d\x02\n\nVehicleEnv\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x128\n\x07runtime\x18\x02 \x01(\x0b2\x1e.city.person.v2.VehicleRuntimeR\x07runtime\x122\n\x07journey\x18\x03 \x01(\x0b2\x18.city.routing.v2.JourneyR\x07journey\x12L\n\x11observed_vehicles\x18\x04 \x03(\x0b2\x1f.city.person.v2.ObservedVehicleR\x10observedVehicles\x12C\n\x0eobserved_lanes\x18\x05 \x03(\x0b2\x1c.city.person.v2.ObservedLaneR\robservedLanes*\xbb\x02\n\x0fVehicleRelation\x12 \n\x1cVEHICLE_RELATION_UNSPECIFIED\x10\x00\x12\x1a\n\x16VEHICLE_RELATION_AHEAD\x10\x01\x12\x1b\n\x17VEHICLE_RELATION_BEHIND\x10\x02\x12!\n\x1dVEHICLE_RELATION_SHADOW_AHEAD\x10\x03\x12"\n\x1eVEHICLE_RELATION_SHADOW_BEHIND\x10\x04\x12\x1f\n\x1bVEHICLE_RELATION_LEFT_AHEAD\x10\x05\x12 \n\x1cVEHICLE_RELATION_RIGHT_AHEAD\x10\x06\x12 \n\x1cVEHICLE_RELATION_LEFT_BEHIND\x10\x07\x12!\n\x1dVEHICLE_RELATION_RIGHT_BEHIND\x10\x08*m\n\nLightState\x12\x1b\n\x17LIGHT_STATE_UNSPECIFIED\x10\x00\x12\x13\n\x0fLIGHT_STATE_RED\x10\x01\x12\x15\n\x11LIGHT_STATE_GREEN\x10\x02\x12\x16\n\x12LIGHT_STATE_YELLOW\x10\x03B\xb5\x01\n\x12com.city.person.v2B\x0cVehicleProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v2.vehicle_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v2B\x0cVehicleProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2'
    _globals['_VEHICLERELATION']._serialized_start = 1591
    _globals['_VEHICLERELATION']._serialized_end = 1906
    _globals['_LIGHTSTATE']._serialized_start = 1908
    _globals['_LIGHTSTATE']._serialized_end = 2017
    _globals['_LC']._serialized_start = 138
    _globals['_LC']._serialized_end = 270
    _globals['_VEHICLEACTION']._serialized_start = 272
    _globals['_VEHICLEACTION']._serialized_end = 399
    _globals['_VEHICLEROUTEACTION']._serialized_start = 401
    _globals['_VEHICLEROUTEACTION']._serialized_end = 489
    _globals['_VEHICLERUNTIME']._serialized_start = 492
    _globals['_VEHICLERUNTIME']._serialized_end = 926
    _globals['_OBSERVEDVEHICLE']._serialized_start = 929
    _globals['_OBSERVEDVEHICLE']._serialized_end = 1122
    _globals['_OBSERVEDLANE']._serialized_start = 1125
    _globals['_OBSERVEDLANE']._serialized_end = 1300
    _globals['_VEHICLEENV']._serialized_start = 1303
    _globals['_VEHICLEENV']._serialized_end = 1588