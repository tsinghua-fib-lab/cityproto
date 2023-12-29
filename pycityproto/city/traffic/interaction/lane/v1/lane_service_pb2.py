"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ......city.traffic.interaction.lane.v1 import lane_pb2 as city_dot_traffic_dot_interaction_dot_lane_dot_v1_dot_lane__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n3city/traffic/interaction/lane/v1/lane_service.proto\x12 city.traffic.interaction.lane.v1\x1a+city/traffic/interaction/lane/v1/lane.proto">\n\x0eSetMaxVRequest\x12\x17\n\x07lane_id\x18\x01 \x01(\x05R\x06laneId\x12\x13\n\x05max_v\x18\x02 \x01(\x01R\x04maxV"\x11\n\x0fSetMaxVResponse"+\n\x0eGetLaneRequest\x12\x19\n\x08lane_ids\x18\x01 \x03(\x05R\x07laneIds"R\n\x0fGetLaneResponse\x12?\n\x06states\x18\x01 \x03(\x0b2\'.city.traffic.interaction.lane.v1.StateR\x06states2\xed\x01\n\x0bLaneService\x12n\n\x07SetMaxV\x120.city.traffic.interaction.lane.v1.SetMaxVRequest\x1a1.city.traffic.interaction.lane.v1.SetMaxVResponse\x12n\n\x07GetLane\x120.city.traffic.interaction.lane.v1.GetLaneRequest\x1a1.city.traffic.interaction.lane.v1.GetLaneResponseB\xa3\x02\n$com.city.traffic.interaction.lane.v1B\x10LaneServiceProtoP\x01ZDgit.fiblab.net/sim/protos/go/city/traffic/interaction/lane/v1;lanev1\xa2\x02\x04CTIL\xaa\x02 City.Traffic.Interaction.Lane.V1\xca\x02 City\\Traffic\\Interaction\\Lane\\V1\xe2\x02,City\\Traffic\\Interaction\\Lane\\V1\\GPBMetadata\xea\x02$City::Traffic::Interaction::Lane::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.traffic.interaction.lane.v1.lane_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n$com.city.traffic.interaction.lane.v1B\x10LaneServiceProtoP\x01ZDgit.fiblab.net/sim/protos/go/city/traffic/interaction/lane/v1;lanev1\xa2\x02\x04CTIL\xaa\x02 City.Traffic.Interaction.Lane.V1\xca\x02 City\\Traffic\\Interaction\\Lane\\V1\xe2\x02,City\\Traffic\\Interaction\\Lane\\V1\\GPBMetadata\xea\x02$City::Traffic::Interaction::Lane::V1'
    _globals['_SETMAXVREQUEST']._serialized_start = 134
    _globals['_SETMAXVREQUEST']._serialized_end = 196
    _globals['_SETMAXVRESPONSE']._serialized_start = 198
    _globals['_SETMAXVRESPONSE']._serialized_end = 215
    _globals['_GETLANEREQUEST']._serialized_start = 217
    _globals['_GETLANEREQUEST']._serialized_end = 260
    _globals['_GETLANERESPONSE']._serialized_start = 262
    _globals['_GETLANERESPONSE']._serialized_end = 344
    _globals['_LANESERVICE']._serialized_start = 347
    _globals['_LANESERVICE']._serialized_end = 584