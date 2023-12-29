"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!city/water/output/v1/output.proto\x12\x14city.water.output.v1\x1a\x15city/geo/v2/geo.proto",\n\x04Road\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n\x05depth\x18\x02 \x01(\x01R\x05depth"[\n\tRoadFlood\x128\n\x08position\x18\x01 \x01(\x0b2\x1c.city.geo.v2.LongLatPositionR\x08position\x12\x14\n\x05depth\x18\x02 \x01(\x01R\x05depth"W\n\x0cDetailedRoad\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x127\n\x06depths\x18\x02 \x03(\x0b2\x1f.city.water.output.v1.RoadFloodR\x06depths"*\n\x04Node\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04head\x18\x02 \x01(\x01R\x04head"n\n\x04Link\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x122\n\x04type\x18\x02 \x01(\x0e2\x1e.city.water.output.v1.LinkTypeR\x04type\x12\x12\n\x04flow\x18\x03 \x01(\x01R\x04flow\x12\x0e\n\x02ok\x18\x04 \x01(\x08R\x02ok"\x9b\x01\n\x03Aoi\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\'\n\x0funsatisfied_num\x18\x02 \x01(\x05R\x0eunsatisfiedNum\x12+\n\x11unsatisfied_ratio\x18\x03 \x01(\x01R\x10unsatisfiedRatio\x12\x16\n\x06demand\x18\x04 \x01(\x01R\x06demand\x12\x16\n\x06supply\x18\x05 \x01(\x01R\x06supply"\xd4\x01\n\x11DrainageBasicInfo\x12#\n\raverage_power\x18\x01 \x01(\x01R\x0caveragePower\x12)\n\x10undrained_volume\x18\x02 \x01(\x01R\x0fundrainedVolume\x12%\n\x0edrained_volume\x18\x03 \x01(\x01R\rdrainedVolume\x12!\n\x0caverage_flow\x18\x04 \x01(\x01R\x0baverageFlow\x12%\n\x0eflooded_volume\x18\x05 \x01(\x01R\rfloodedVolume"Y\n\x0fSupplyBasicInfo\x12#\n\raverage_power\x18\x01 \x01(\x01R\x0caveragePower\x12!\n\x0caverage_flow\x18\x02 \x01(\x01R\x0baverageFlow"\xae\x02\n\x16SupplyDemandStatistics\x12%\n\x0epersons_demand\x18\x01 \x01(\x01R\rpersonsDemand\x12/\n\x13unsatisfied_persons\x18\x02 \x01(\x05R\x12unsatisfiedPersons\x12:\n\x19unsatisfied_persons_ratio\x18\x03 \x01(\x01R\x17unsatisfiedPersonsRatio\x12\x1f\n\x0baois_demand\x18\x04 \x01(\x01R\naoisDemand\x12)\n\x10unsatisfied_aois\x18\x05 \x01(\x05R\x0funsatisfiedAois\x124\n\x16unsatisfied_aois_ratio\x18\x06 \x01(\x01R\x14unsatisfiedAoisRatio"x\n\x11FailureStatistics\x12\x1f\n\x0bfailure_num\x18\x01 \x01(\x05R\nfailureNum\x12\x1d\n\nnormal_num\x18\x02 \x01(\x05R\tnormalNum\x12#\n\rfailure_ratio\x18\x03 \x01(\x01R\x0cfailureRatio"\xe1\x01\n\x0fDrainageMetrics\x12W\n\x13drainage_basic_info\x18\x01 \x01(\x0b2\'.city.water.output.v1.DrainageBasicInfoR\x11drainageBasicInfo\x12\x1d\n\nload_ratio\x18\x02 \x01(\x01R\tloadRatio\x12V\n\x12failure_statistics\x18\x03 \x01(\x0b2\'.city.water.output.v1.FailureStatisticsR\x11failureStatistics"\xc1\x02\n\rSupplyMetrics\x12Q\n\x11supply_basic_info\x18\x01 \x01(\x0b2%.city.water.output.v1.SupplyBasicInfoR\x0fsupplyBasicInfo\x12f\n\x18supply_demand_statistics\x18\x02 \x01(\x0b2,.city.water.output.v1.SupplyDemandStatisticsR\x16supplyDemandStatistics\x12\x1d\n\nload_ratio\x18\x03 \x01(\x01R\tloadRatio\x12V\n\x12failure_statistics\x18\x04 \x01(\x0b2\'.city.water.output.v1.FailureStatisticsR\x11failureStatistics*M\n\x08LinkType\x12\x19\n\x15LINK_TYPE_UNSPECIFIED\x10\x00\x12\x12\n\x0eLINK_TYPE_PIPE\x10\x01\x12\x12\n\x0eLINK_TYPE_PUMP\x10\x02B\xd6\x01\n\x18com.city.water.output.v1B\x0bOutputProtoP\x01Z:git.fiblab.net/sim/protos/go/city/water/output/v1;outputv1\xa2\x02\x03CWO\xaa\x02\x14City.Water.Output.V1\xca\x02\x14City\\Water\\Output\\V1\xe2\x02 City\\Water\\Output\\V1\\GPBMetadata\xea\x02\x17City::Water::Output::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.water.output.v1.output_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x18com.city.water.output.v1B\x0bOutputProtoP\x01Z:git.fiblab.net/sim/protos/go/city/water/output/v1;outputv1\xa2\x02\x03CWO\xaa\x02\x14City.Water.Output.V1\xca\x02\x14City\\Water\\Output\\V1\xe2\x02 City\\Water\\Output\\V1\\GPBMetadata\xea\x02\x17City::Water::Output::V1'
    _globals['_LINKTYPE']._serialized_start = 1909
    _globals['_LINKTYPE']._serialized_end = 1986
    _globals['_ROAD']._serialized_start = 82
    _globals['_ROAD']._serialized_end = 126
    _globals['_ROADFLOOD']._serialized_start = 128
    _globals['_ROADFLOOD']._serialized_end = 219
    _globals['_DETAILEDROAD']._serialized_start = 221
    _globals['_DETAILEDROAD']._serialized_end = 308
    _globals['_NODE']._serialized_start = 310
    _globals['_NODE']._serialized_end = 352
    _globals['_LINK']._serialized_start = 354
    _globals['_LINK']._serialized_end = 464
    _globals['_AOI']._serialized_start = 467
    _globals['_AOI']._serialized_end = 622
    _globals['_DRAINAGEBASICINFO']._serialized_start = 625
    _globals['_DRAINAGEBASICINFO']._serialized_end = 837
    _globals['_SUPPLYBASICINFO']._serialized_start = 839
    _globals['_SUPPLYBASICINFO']._serialized_end = 928
    _globals['_SUPPLYDEMANDSTATISTICS']._serialized_start = 931
    _globals['_SUPPLYDEMANDSTATISTICS']._serialized_end = 1233
    _globals['_FAILURESTATISTICS']._serialized_start = 1235
    _globals['_FAILURESTATISTICS']._serialized_end = 1355
    _globals['_DRAINAGEMETRICS']._serialized_start = 1358
    _globals['_DRAINAGEMETRICS']._serialized_end = 1583
    _globals['_SUPPLYMETRICS']._serialized_start = 1586
    _globals['_SUPPLYMETRICS']._serialized_end = 1907