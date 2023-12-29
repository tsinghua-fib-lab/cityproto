"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n city/elec/output/v1/output.proto\x12\x13city.elec.output.v1"\xb5\x01\n\x03Aoi\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x128\n\x18unsatisfied_demand_ratio\x18\x02 \x01(\x01R\x16unsatisfiedDemandRatio\x124\n\x16unsatisfied_demand_num\x18\x03 \x01(\x05R\x14unsatisfiedDemandNum\x12\x16\n\x06demand\x18\x04 \x01(\x01R\x06demand\x12\x16\n\x06supply\x18\x05 \x01(\x01R\x06supply"\xb0\x01\n\x0cAveragePower\x12\'\n\x0ftransformer_500\x18\x01 \x01(\x01R\x0etransformer500\x12\'\n\x0ftransformer_220\x18\x02 \x01(\x01R\x0etransformer220\x12\'\n\x0ftransformer_110\x18\x03 \x01(\x01R\x0etransformer110\x12%\n\x0etransformer_10\x18\x04 \x01(\x01R\rtransformer10B\xd0\x01\n\x17com.city.elec.output.v1B\x0bOutputProtoP\x01Z9git.fiblab.net/sim/protos/go/city/elec/output/v1;outputv1\xa2\x02\x03CEO\xaa\x02\x13City.Elec.Output.V1\xca\x02\x13City\\Elec\\Output\\V1\xe2\x02\x1fCity\\Elec\\Output\\V1\\GPBMetadata\xea\x02\x16City::Elec::Output::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.elec.output.v1.output_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x17com.city.elec.output.v1B\x0bOutputProtoP\x01Z9git.fiblab.net/sim/protos/go/city/elec/output/v1;outputv1\xa2\x02\x03CEO\xaa\x02\x13City.Elec.Output.V1\xca\x02\x13City\\Elec\\Output\\V1\xe2\x02\x1fCity\\Elec\\Output\\V1\\GPBMetadata\xea\x02\x16City::Elec::Output::V1'
    _globals['_AOI']._serialized_start = 58
    _globals['_AOI']._serialized_end = 239
    _globals['_AVERAGEPOWER']._serialized_start = 242
    _globals['_AVERAGEPOWER']._serialized_end = 418