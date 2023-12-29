"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.cognition.output.v1 import output_pb2 as city_dot_cognition_dot_output_dot_v1_dot_output__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n-city/cognition/output/v1/output_service.proto\x12\x18city.cognition.output.v1\x1a%city/cognition/output/v1/output.proto"\xd7\x03\n\rOutputRequest\x12\x12\n\x04step\x18\x01 \x01(\x05R\x04step\x124\n\x05nodes\x18\x02 \x03(\x0b2\x1e.city.cognition.output.v1.NodeR\x05nodes\x12C\n\ninfluences\x18\x03 \x03(\x0b2#.city.cognition.output.v1.InfluenceR\ninfluences\x12;\n\x07heatmap\x18\x04 \x01(\x0b2!.city.cognition.output.v1.HeatmapR\x07heatmap\x122\n\x04stat\x18\x05 \x01(\x0b2\x1e.city.cognition.output.v1.StatR\x04stat\x12=\n\x08contents\x18\x06 \x03(\x0b2!.city.cognition.output.v1.ContentR\x08contents\x12N\n\x10group_influences\x18\x07 \x03(\x0b2#.city.cognition.output.v1.InfluenceR\x0fgroupInfluences\x127\n\x06groups\x18\x08 \x03(\x0b2\x1f.city.cognition.output.v1.GroupR\x06groups"\x10\n\x0eOutputResponse2l\n\rOutputService\x12[\n\x06Output\x12\'.city.cognition.output.v1.OutputRequest\x1a(.city.cognition.output.v1.OutputResponseB\xf5\x01\n\x1ccom.city.cognition.output.v1B\x12OutputServiceProtoP\x01Z>git.fiblab.net/sim/protos/go/city/cognition/output/v1;outputv1\xa2\x02\x03CCO\xaa\x02\x18City.Cognition.Output.V1\xca\x02\x18City\\Cognition\\Output\\V1\xe2\x02$City\\Cognition\\Output\\V1\\GPBMetadata\xea\x02\x1bCity::Cognition::Output::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.cognition.output.v1.output_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x1ccom.city.cognition.output.v1B\x12OutputServiceProtoP\x01Z>git.fiblab.net/sim/protos/go/city/cognition/output/v1;outputv1\xa2\x02\x03CCO\xaa\x02\x18City.Cognition.Output.V1\xca\x02\x18City\\Cognition\\Output\\V1\xe2\x02$City\\Cognition\\Output\\V1\\GPBMetadata\xea\x02\x1bCity::Cognition::Output::V1'
    _globals['_OUTPUTREQUEST']._serialized_start = 115
    _globals['_OUTPUTREQUEST']._serialized_end = 586
    _globals['_OUTPUTRESPONSE']._serialized_start = 588
    _globals['_OUTPUTRESPONSE']._serialized_end = 604
    _globals['_OUTPUTSERVICE']._serialized_start = 606
    _globals['_OUTPUTSERVICE']._serialized_end = 714