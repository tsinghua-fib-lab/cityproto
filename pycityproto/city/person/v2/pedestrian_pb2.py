"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.person.v2 import motion_pb2 as city_dot_person_dot_v2_dot_motion__pb2
from ....city.routing.v2 import routing_pb2 as city_dot_routing_dot_v2_dot_routing__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fcity/person/v2/pedestrian.proto\x12\x0ecity.person.v2\x1a\x1bcity/person/v2/motion.proto\x1a\x1dcity/routing/v2/routing.proto"\xf3\x01\n\rPedestrianEnv\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x124\n\x06motion\x18\x02 \x01(\x0b2\x1c.city.person.v2.PersonMotionR\x06motion\x122\n\x07journey\x18\x03 \x01(\x0b2\x18.city.routing.v2.JourneyR\x07journey\x126\n\x18is_current_lane_no_entry\x18\x04 \x01(\x08R\x14isCurrentLaneNoEntry\x120\n\x15is_next_lane_no_entry\x18\x05 \x01(\x08R\x11isNextLaneNoEntry"B\n\x10PedestrianAction\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x0e\n\x02vx\x18\x02 \x01(\x01R\x02vx\x12\x0e\n\x02vy\x18\x03 \x01(\x01R\x02vyB\xb8\x01\n\x12com.city.person.v2B\x0fPedestrianProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v2.pedestrian_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v2B\x0fPedestrianProtoP\x01Z7git.fiblab.net/sim/protos/v2/go/city/person/v2;personv2\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V2\xca\x02\x0eCity\\Person\\V2\xe2\x02\x1aCity\\Person\\V2\\GPBMetadata\xea\x02\x10City::Person::V2'
    _globals['_PEDESTRIANENV']._serialized_start = 112
    _globals['_PEDESTRIANENV']._serialized_end = 355
    _globals['_PEDESTRIANACTION']._serialized_start = 357
    _globals['_PEDESTRIANACTION']._serialized_end = 423