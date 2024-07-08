"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.event.v2 import event_pb2 as city_dot_event_dot_v2_dot_event__pb2
from ....city.person.v1 import motion_pb2 as city_dot_person_dot_v1_dot_motion__pb2
from ....city.person.v1 import person_pb2 as city_dot_person_dot_v1_dot_person__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#city/person/v1/person_runtime.proto\x12\x0ecity.person.v1\x1a\x19city/event/v2/event.proto\x1a\x1bcity/person/v1/motion.proto\x1a\x1bcity/person/v1/person.proto"\xad\x01\n\rPersonRuntime\x12/\n\x04base\x18\x01 \x01(\x0b2\x16.city.person.v1.PersonH\x00R\x04base\x88\x01\x01\x124\n\x06motion\x18\x02 \x01(\x0b2\x1c.city.person.v1.PersonMotionR\x06motion\x12,\n\x06events\x18\x03 \x03(\x0b2\x14.city.event.v2.EventR\x06eventsB\x07\n\x05_baseB\xb8\x01\n\x12com.city.person.v1B\x12PersonRuntimeProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.person.v1.person_runtime_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.person.v1B\x12PersonRuntimeProtoP\x01Z4git.fiblab.net/sim/protos/go/city/person/v1;personv1\xa2\x02\x03CPX\xaa\x02\x0eCity.Person.V1\xca\x02\x0eCity\\Person\\V1\xe2\x02\x1aCity\\Person\\V1\\GPBMetadata\xea\x02\x10City::Person::V1'
    _globals['_PERSONRUNTIME']._serialized_start = 141
    _globals['_PERSONRUNTIME']._serialized_end = 314