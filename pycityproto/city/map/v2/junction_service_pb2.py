"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.map.v2 import lane_state_pb2 as city_dot_map_dot_v2_dot_lane__state__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n"city/map/v2/junction_service.proto\x12\x0bcity.map.v2\x1a\x1ccity/map/v2/lane_state.proto"\x81\x01\n\x12GetJunctionRequest\x12!\n\x0cjunction_ids\x18\x01 \x03(\x05R\x0bjunctionIds\x12!\n\x0cexclude_lane\x18\x02 \x01(\x08R\x0bexcludeLane\x12%\n\x0eexclude_person\x18\x03 \x01(\x08R\rexcludePerson"I\n\x13GetJunctionResponse\x122\n\x06states\x18\x01 \x03(\x0b2\x1a.city.map.v2.JunctionStateR\x06states"\xe2\x04\n\rJunctionState\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12$\n\x0ein_vehicle_cnt\x18\x02 \x01(\x05R\x0cinVehicleCnt\x12&\n\x0fout_vehicle_cnt\x18\x03 \x01(\x05R\routVehicleCnt\x12\x1f\n\x0bvehicle_cnt\x18\x04 \x01(\x05R\nvehicleCnt\x12+\n\x12cum_in_vehicle_cnt\x18\x05 \x01(\x05R\x0fcumInVehicleCnt\x12-\n\x13cum_out_vehicle_cnt\x18\x06 \x01(\x05R\x10cumOutVehicleCnt\x12,\n\x05lanes\x18\x07 \x03(\x0b2\x16.city.map.v2.LaneStateR\x05lanes\x12R\n\x19predecessor_driving_lanes\x18\x08 \x03(\x0b2\x16.city.map.v2.LaneStateR\x17predecessorDrivingLanes\x129\n\x19total_queuing_vehicle_cnt\x18\t \x01(\x05R\x16totalQueuingVehicleCnt\x12,\n\x12total_queuing_time\x18\n \x01(\x01R\x10totalQueuingTime\x12(\n\x10avg_queuing_time\x18\x0b \x01(\x01R\x0eavgQueuingTime\x125\n\x17max_queuing_vehicle_cnt\x18\x0c \x01(\x05R\x14maxQueuingVehicleCnt\x12*\n\x11has_traffic_light\x18\r \x01(\x08R\x0fhasTrafficLight2c\n\x0fJunctionService\x12P\n\x0bGetJunction\x12\x1f.city.map.v2.GetJunctionRequest\x1a .city.map.v2.GetJunctionResponseB\xa8\x01\n\x0fcom.city.map.v2B\x14JunctionServiceProtoP\x01Z1git.fiblab.net/sim/protos/v2/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.map.v2.junction_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.map.v2B\x14JunctionServiceProtoP\x01Z1git.fiblab.net/sim/protos/v2/go/city/map/v2;mapv2\xa2\x02\x03CMX\xaa\x02\x0bCity.Map.V2\xca\x02\x0bCity\\Map\\V2\xe2\x02\x17City\\Map\\V2\\GPBMetadata\xea\x02\rCity::Map::V2'
    _globals['_GETJUNCTIONREQUEST']._serialized_start = 82
    _globals['_GETJUNCTIONREQUEST']._serialized_end = 211
    _globals['_GETJUNCTIONRESPONSE']._serialized_start = 213
    _globals['_GETJUNCTIONRESPONSE']._serialized_end = 286
    _globals['_JUNCTIONSTATE']._serialized_start = 289
    _globals['_JUNCTIONSTATE']._serialized_end = 899
    _globals['_JUNCTIONSERVICE']._serialized_start = 901
    _globals['_JUNCTIONSERVICE']._serialized_end = 1000