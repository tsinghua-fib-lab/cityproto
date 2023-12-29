"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%city/cognition/output/v1/output.proto\x12\x18city.cognition.output.v1".\n\x04Node\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n\x06status\x18\x03 \x01(\x01R\x06status"a\n\tInfluence\x12\x1b\n\tsource_id\x18\x01 \x01(\x05R\x08sourceId\x12\x1b\n\ttarget_id\x18\x02 \x01(\x05R\x08targetId\x12\x1a\n\x08strength\x18\x03 \x01(\x01R\x08strength"a\n\x07Heatmap\x12\x19\n\x08num_rows\x18\x01 \x01(\x05R\x07numRows\x12\x1f\n\x0bnum_columns\x18\x02 \x01(\x05R\nnumColumns\x12\x1a\n\x08strength\x18\x03 \x03(\x01R\x08strength"e\n\x04Stat\x12\x1d\n\ncrowd_cnts\x18\x01 \x03(\x05R\tcrowdCnts\x12!\n\x0ccrowd_ratios\x18\x02 \x03(\x01R\x0bcrowdRatios\x12\x1b\n\tkey_nodes\x18\x03 \x03(\x05R\x08keyNodes"-\n\x07Content\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n\x04text\x18\x02 \x01(\tR\x04text"T\n\x08NodeMeta\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x10\n\x03lng\x18\x02 \x01(\x01R\x03lng\x12\x10\n\x03lat\x18\x03 \x01(\x01R\x03lat\x12\x14\n\x05level\x18\x04 \x01(\x05R\x05level"E\n\tNodesMeta\x128\n\x05nodes\x18\x01 \x03(\x0b2".city.cognition.output.v1.NodeMetaR\x05nodes"E\n\x05Group\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n\x04size\x18\x02 \x01(\x05R\x04size\x12\x18\n\x07changes\x18\x03 \x03(\x05R\x07changesB\xee\x01\n\x1ccom.city.cognition.output.v1B\x0bOutputProtoP\x01Z>git.fiblab.net/sim/protos/go/city/cognition/output/v1;outputv1\xa2\x02\x03CCO\xaa\x02\x18City.Cognition.Output.V1\xca\x02\x18City\\Cognition\\Output\\V1\xe2\x02$City\\Cognition\\Output\\V1\\GPBMetadata\xea\x02\x1bCity::Cognition::Output::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.cognition.output.v1.output_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x1ccom.city.cognition.output.v1B\x0bOutputProtoP\x01Z>git.fiblab.net/sim/protos/go/city/cognition/output/v1;outputv1\xa2\x02\x03CCO\xaa\x02\x18City.Cognition.Output.V1\xca\x02\x18City\\Cognition\\Output\\V1\xe2\x02$City\\Cognition\\Output\\V1\\GPBMetadata\xea\x02\x1bCity::Cognition::Output::V1'
    _globals['_NODE']._serialized_start = 67
    _globals['_NODE']._serialized_end = 113
    _globals['_INFLUENCE']._serialized_start = 115
    _globals['_INFLUENCE']._serialized_end = 212
    _globals['_HEATMAP']._serialized_start = 214
    _globals['_HEATMAP']._serialized_end = 311
    _globals['_STAT']._serialized_start = 313
    _globals['_STAT']._serialized_end = 414
    _globals['_CONTENT']._serialized_start = 416
    _globals['_CONTENT']._serialized_end = 461
    _globals['_NODEMETA']._serialized_start = 463
    _globals['_NODEMETA']._serialized_end = 547
    _globals['_NODESMETA']._serialized_start = 549
    _globals['_NODESMETA']._serialized_end = 618
    _globals['_GROUP']._serialized_start = 620
    _globals['_GROUP']._serialized_end = 689