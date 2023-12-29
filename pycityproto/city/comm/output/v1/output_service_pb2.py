"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.comm.output.v1 import output_pb2 as city_dot_comm_dot_output_dot_v1_dot_output__pb2
from .....city.event.v1 import event_pb2 as city_dot_event_dot_v1_dot_event__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(city/comm/output/v1/output_service.proto\x12\x13city.comm.output.v1\x1a city/comm/output/v1/output.proto\x1a\x19city/event/v1/event.proto"\x83\x04\n\rOutputRequest\x12\x12\n\x04step\x18\x01 \x01(\x05R\x04step\x12/\n\x05nodes\x18\x02 \x03(\x0b2\x19.city.comm.output.v1.NodeR\x05nodes\x12E\n\rbase_stations\x18\x03 \x03(\x0b2 .city.comm.output.v1.BaseStationR\x0cbaseStations\x12B\n\x0esignal_heatmap\x18\x04 \x01(\x0b2\x1b.city.comm.output.v1.SignalR\rsignalHeatmap\x12M\n\x14small_signal_heatmap\x18\t \x01(\x0b2\x1b.city.comm.output.v1.SignalR\x12smallSignalHeatmap\x125\n\x07persons\x18\x05 \x03(\x0b2\x1b.city.comm.output.v1.PersonR\x07persons\x12,\n\x04aois\x18\x06 \x03(\x0b2\x18.city.comm.output.v1.AoiR\x04aois\x12-\n\x06events\x18\x07 \x01(\x0b2\x15.city.event.v1.EventsR\x06events\x12?\n\nstatistics\x18\x08 \x01(\x0b2\x1f.city.comm.output.v1.StatisticsR\nstatistics"\x10\n\x0eOutputResponse2b\n\rOutputService\x12Q\n\x06Output\x12".city.comm.output.v1.OutputRequest\x1a#.city.comm.output.v1.OutputResponseB\xd7\x01\n\x17com.city.comm.output.v1B\x12OutputServiceProtoP\x01Z9git.fiblab.net/sim/protos/go/city/comm/output/v1;outputv1\xa2\x02\x03CCO\xaa\x02\x13City.Comm.Output.V1\xca\x02\x13City\\Comm\\Output\\V1\xe2\x02\x1fCity\\Comm\\Output\\V1\\GPBMetadata\xea\x02\x16City::Comm::Output::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.comm.output.v1.output_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x17com.city.comm.output.v1B\x12OutputServiceProtoP\x01Z9git.fiblab.net/sim/protos/go/city/comm/output/v1;outputv1\xa2\x02\x03CCO\xaa\x02\x13City.Comm.Output.V1\xca\x02\x13City\\Comm\\Output\\V1\xe2\x02\x1fCity\\Comm\\Output\\V1\\GPBMetadata\xea\x02\x16City::Comm::Output::V1'
    _globals['_OUTPUTREQUEST']._serialized_start = 127
    _globals['_OUTPUTREQUEST']._serialized_end = 642
    _globals['_OUTPUTRESPONSE']._serialized_start = 644
    _globals['_OUTPUTRESPONSE']._serialized_end = 660
    _globals['_OUTPUTSERVICE']._serialized_start = 662
    _globals['_OUTPUTSERVICE']._serialized_end = 760