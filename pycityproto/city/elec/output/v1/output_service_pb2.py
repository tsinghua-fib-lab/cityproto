"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.elec.output.v1 import output_pb2 as city_dot_elec_dot_output_dot_v1_dot_output__pb2
from .....city.event.v1 import event_pb2 as city_dot_event_dot_v1_dot_event__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(city/elec/output/v1/output_service.proto\x12\x13city.elec.output.v1\x1a city/elec/output/v1/output.proto\x1a\x19city/event/v1/event.proto"\xf4\x03\n\rOutputRequest\x12\x12\n\x04step\x18\x01 \x01(\x05R\x04step\x12\x1d\n\nruined_ids\x18\x02 \x03(\x05R\truinedIds\x12\x1f\n\x0bstopped_ids\x18\x03 \x03(\x05R\nstoppedIds\x12,\n\x04aois\x18\x04 \x03(\x0b2\x18.city.elec.output.v1.AoiR\x04aois\x12-\n\x06events\x18\x05 \x01(\x0b2\x15.city.event.v1.EventsR\x06events\x12<\n\x1aresident_unsatisfied_ratio\x18\x06 \x01(\x01R\x18residentUnsatisfiedRatio\x12\'\n\x0fresident_demand\x18\x07 \x01(\x01R\x0eresidentDemand\x122\n\x15aoi_unsatisfied_ratio\x18\x08 \x01(\x01R\x13aoiUnsatisfiedRatio\x12.\n\x13aoi_unsatisfied_num\x18\t \x01(\x05R\x11aoiUnsatisfiedNum\x12\x1d\n\naoi_demand\x18\n \x01(\x01R\taoiDemand\x12H\n\x0eaverage_powers\x18\x0b \x01(\x0b2!.city.elec.output.v1.AveragePowerR\raveragePowers"\x10\n\x0eOutputResponse2b\n\rOutputService\x12Q\n\x06Output\x12".city.elec.output.v1.OutputRequest\x1a#.city.elec.output.v1.OutputResponseB\xd7\x01\n\x17com.city.elec.output.v1B\x12OutputServiceProtoP\x01Z9git.fiblab.net/sim/protos/go/city/elec/output/v1;outputv1\xa2\x02\x03CEO\xaa\x02\x13City.Elec.Output.V1\xca\x02\x13City\\Elec\\Output\\V1\xe2\x02\x1fCity\\Elec\\Output\\V1\\GPBMetadata\xea\x02\x16City::Elec::Output::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.elec.output.v1.output_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x17com.city.elec.output.v1B\x12OutputServiceProtoP\x01Z9git.fiblab.net/sim/protos/go/city/elec/output/v1;outputv1\xa2\x02\x03CEO\xaa\x02\x13City.Elec.Output.V1\xca\x02\x13City\\Elec\\Output\\V1\xe2\x02\x1fCity\\Elec\\Output\\V1\\GPBMetadata\xea\x02\x16City::Elec::Output::V1'
    _globals['_OUTPUTREQUEST']._serialized_start = 127
    _globals['_OUTPUTREQUEST']._serialized_end = 627
    _globals['_OUTPUTRESPONSE']._serialized_start = 629
    _globals['_OUTPUTRESPONSE']._serialized_end = 645
    _globals['_OUTPUTSERVICE']._serialized_start = 647
    _globals['_OUTPUTSERVICE']._serialized_end = 745