"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.event.v1 import event_pb2 as city_dot_event_dot_v1_dot_event__pb2
from .....city.water.output.v1 import output_pb2 as city_dot_water_dot_output_dot_v1_dot_output__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)city/water/output/v1/output_service.proto\x12\x14city.water.output.v1\x1a\x19city/event/v1/event.proto\x1a!city/water/output/v1/output.proto"\xc9\x05\n\rOutputRequest\x12\x12\n\x04step\x18\x01 \x01(\x05R\x04step\x120\n\x05roads\x18\x02 \x03(\x0b2\x1a.city.water.output.v1.RoadR\x05roads\x12I\n\x0edetailed_roads\x18\x03 \x03(\x0b2".city.water.output.v1.DetailedRoadR\rdetailedRoads\x12A\n\x0edrainage_nodes\x18\x04 \x03(\x0b2\x1a.city.water.output.v1.NodeR\rdrainageNodes\x12A\n\x0edrainage_links\x18\x05 \x03(\x0b2\x1a.city.water.output.v1.LinkR\rdrainageLinks\x12=\n\x0csupply_nodes\x18\x06 \x03(\x0b2\x1a.city.water.output.v1.NodeR\x0bsupplyNodes\x12=\n\x0csupply_links\x18\x07 \x03(\x0b2\x1a.city.water.output.v1.LinkR\x0bsupplyLinks\x12-\n\x04aois\x18\x08 \x03(\x0b2\x19.city.water.output.v1.AoiR\x04aois\x12\'\n\x0fdrainage_metric\x18\t \x01(\x01R\x0edrainageMetric\x12-\n\x06events\x18\n \x01(\x0b2\x15.city.event.v1.EventsR\x06events\x12P\n\x10drainage_metrics\x18\x0b \x01(\x0b2%.city.water.output.v1.DrainageMetricsR\x0fdrainageMetrics\x12J\n\x0esupply_metrics\x18\x0c \x01(\x0b2#.city.water.output.v1.SupplyMetricsR\rsupplyMetrics"\x10\n\x0eOutputResponse2d\n\rOutputService\x12S\n\x06Output\x12#.city.water.output.v1.OutputRequest\x1a$.city.water.output.v1.OutputResponseB\xdd\x01\n\x18com.city.water.output.v1B\x12OutputServiceProtoP\x01Z:git.fiblab.net/sim/protos/go/city/water/output/v1;outputv1\xa2\x02\x03CWO\xaa\x02\x14City.Water.Output.V1\xca\x02\x14City\\Water\\Output\\V1\xe2\x02 City\\Water\\Output\\V1\\GPBMetadata\xea\x02\x17City::Water::Output::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.water.output.v1.output_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x18com.city.water.output.v1B\x12OutputServiceProtoP\x01Z:git.fiblab.net/sim/protos/go/city/water/output/v1;outputv1\xa2\x02\x03CWO\xaa\x02\x14City.Water.Output.V1\xca\x02\x14City\\Water\\Output\\V1\xe2\x02 City\\Water\\Output\\V1\\GPBMetadata\xea\x02\x17City::Water::Output::V1'
    _globals['_OUTPUTREQUEST']._serialized_start = 130
    _globals['_OUTPUTREQUEST']._serialized_end = 843
    _globals['_OUTPUTRESPONSE']._serialized_start = 845
    _globals['_OUTPUTRESPONSE']._serialized_end = 861
    _globals['_OUTPUTSERVICE']._serialized_start = 863
    _globals['_OUTPUTSERVICE']._serialized_end = 963