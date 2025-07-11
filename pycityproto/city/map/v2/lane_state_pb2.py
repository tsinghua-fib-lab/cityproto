"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.map.v2 import light_pb2 as city_dot_map_dot_v2_dot_light__pb2
from ....city.person.v2 import motion_pb2 as city_dot_person_dot_v2_dot_motion__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ccity/map/v2/lane_state.proto\x12\x0bcity.map.v2\x1a\x17city/map/v2/light.proto\x1a\x1bcity/person/v2/motion.proto"\xc6\x03\n\tLaneState\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x126\n\x07persons\x18\x02 \x03(\x0b2\x1c.city.person.v2.PersonMotionR\x07persons\x12\x13\n\x05avg_v\x18\x03 \x01(\x01R\x04avgV\x12 \n\x0brestriction\x18\x04 \x01(\x08R\x0brestriction\x128\n\x0blight_state\x18\x05 \x01(\x0e2\x17.city.map.v2.LightStateR\nlightState\x12$\n\x0ein_vehicle_cnt\x18\x06 \x01(\x05R\x0cinVehicleCnt\x12&\n\x0fout_vehicle_cnt\x18\x07 \x01(\x05R\routVehicleCnt\x12\x1f\n\x0bvehicle_cnt\x18\x08 \x01(\x05R\nvehicleCnt\x129\n\x19total_queuing_vehicle_cnt\x18\t \x01(\x05R\x16totalQueuingVehicleCnt\x12,\n\x12total_queuing_time\x18\n \x01(\x01R\x10totalQueuingTime\x12(\n\x10avg_queuing_time\x18\x0b \x01(\x01R\x0eavgQueuingTimeB\xa2\x01\n\x0fcom.city.map.v2B\x0eLaneStateProtoP\x01Z1git.fiblab.net/sim/protos/v2/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.lane_state_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\x0eLaneStateProtoP\x01Z1git.fiblab.net/sim/protos/v2/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_LANESTATE']._serialized_start = 100
    _globals['_LANESTATE']._serialized_end = 554