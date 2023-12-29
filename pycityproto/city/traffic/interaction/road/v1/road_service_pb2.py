"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ......city.event.v1 import event_pb2 as city_dot_event_dot_v1_dot_event__pb2
from ......city.traffic.interaction.road.v1 import road_pb2 as city_dot_traffic_dot_interaction_dot_road_dot_v1_dot_road__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n3city/traffic/interaction/road/v1/road_service.proto\x12 city.traffic.interaction.road.v1\x1a\x19city/event/v1/event.proto\x1a+city/traffic/interaction/road/v1/road.proto"+\n\x0eGetRoadRequest\x12\x19\n\x08road_ids\x18\x01 \x03(\x05R\x07roadIds"R\n\x0fGetRoadResponse\x12?\n\x06states\x18\x01 \x03(\x0b2\'.city.traffic.interaction.road.v1.StateR\x06states"\x14\n\x12GetRuinInfoRequest"2\n\x08RuinInfo\x12\x10\n\x03num\x18\x01 \x01(\x05R\x03num\x12\x14\n\x05ratio\x18\x02 \x01(\x01R\x05ratio"\xd3\x01\n\x13GetRuinInfoResponse\x12<\n\x03one\x18\x01 \x01(\x0b2*.city.traffic.interaction.road.v1.RuinInfoR\x03one\x12<\n\x03two\x18\x02 \x01(\x0b2*.city.traffic.interaction.road.v1.RuinInfoR\x03two\x12@\n\x05three\x18\x03 \x01(\x0b2*.city.traffic.interaction.road.v1.RuinInfoR\x05three"\x12\n\x10GetEventsRequest"B\n\x11GetEventsResponse\x12-\n\x06events\x18\x01 \x01(\x0b2\x15.city.event.v1.EventsR\x06events2\xef\x02\n\x0bRoadService\x12n\n\x07GetRoad\x120.city.traffic.interaction.road.v1.GetRoadRequest\x1a1.city.traffic.interaction.road.v1.GetRoadResponse\x12z\n\x0bGetRuinInfo\x124.city.traffic.interaction.road.v1.GetRuinInfoRequest\x1a5.city.traffic.interaction.road.v1.GetRuinInfoResponse\x12t\n\tGetEvents\x122.city.traffic.interaction.road.v1.GetEventsRequest\x1a3.city.traffic.interaction.road.v1.GetEventsResponseB\xa3\x02\n$com.city.traffic.interaction.road.v1B\x10RoadServiceProtoP\x01ZDgit.fiblab.net/sim/protos/go/city/traffic/interaction/road/v1;roadv1\xa2\x02\x04CTIR\xaa\x02 City.Traffic.Interaction.Road.V1\xca\x02 City\\Traffic\\Interaction\\Road\\V1\xe2\x02,City\\Traffic\\Interaction\\Road\\V1\\GPBMetadata\xea\x02$City::Traffic::Interaction::Road::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.traffic.interaction.road.v1.road_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n$com.city.traffic.interaction.road.v1B\x10RoadServiceProtoP\x01ZDgit.fiblab.net/sim/protos/go/city/traffic/interaction/road/v1;roadv1\xa2\x02\x04CTIR\xaa\x02 City.Traffic.Interaction.Road.V1\xca\x02 City\\Traffic\\Interaction\\Road\\V1\xe2\x02,City\\Traffic\\Interaction\\Road\\V1\\GPBMetadata\xea\x02$City::Traffic::Interaction::Road::V1'
    _globals['_GETROADREQUEST']._serialized_start = 161
    _globals['_GETROADREQUEST']._serialized_end = 204
    _globals['_GETROADRESPONSE']._serialized_start = 206
    _globals['_GETROADRESPONSE']._serialized_end = 288
    _globals['_GETRUININFOREQUEST']._serialized_start = 290
    _globals['_GETRUININFOREQUEST']._serialized_end = 310
    _globals['_RUININFO']._serialized_start = 312
    _globals['_RUININFO']._serialized_end = 362
    _globals['_GETRUININFORESPONSE']._serialized_start = 365
    _globals['_GETRUININFORESPONSE']._serialized_end = 576
    _globals['_GETEVENTSREQUEST']._serialized_start = 578
    _globals['_GETEVENTSREQUEST']._serialized_end = 596
    _globals['_GETEVENTSRESPONSE']._serialized_start = 598
    _globals['_GETEVENTSRESPONSE']._serialized_end = 664
    _globals['_ROADSERVICE']._serialized_start = 667
    _globals['_ROADSERVICE']._serialized_end = 1034