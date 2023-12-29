"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from .....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1ecity/elec/input/v1/input.proto\x12\x12city.elec.input.v1\x1a\x15city/geo/v2/geo.proto"`\n\rRepairStation\x12\x15\n\x06aoi_id\x18\x01 \x01(\x05R\x05aoiId\x128\n\x08position\x18\x02 \x01(\x0b2\x1c.city.geo.v2.LongLatPositionR\x08position"\xc2\x02\n\x08Facility\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x128\n\x08position\x18\x02 \x01(\x0b2\x1c.city.geo.v2.LongLatPositionR\x08position\x124\n\x04type\x18\x03 \x01(\x0e2 .city.elec.input.v1.FacilityTypeR\x04type\x12\x1a\n\x08relation\x18\x04 \x03(\x05R\x08relation\x12"\n\nforeign_id\x18\x05 \x01(\x05H\x00R\tforeignId\x88\x01\x01\x12\x1a\n\x06aoi_id\x18\x06 \x01(\x05H\x01R\x05aoiId\x88\x01\x01\x12,\n\x0fnum_transformer\x18\x07 \x01(\x05H\x02R\x0enumTransformer\x88\x01\x01B\r\n\x0b_foreign_idB\t\n\x07_aoi_idB\x12\n\x10_num_transformer"\x96\x01\n\nFacilities\x12<\n\nfacilities\x18\x01 \x03(\x0b2\x1c.city.elec.input.v1.FacilityR\nfacilities\x12J\n\x0frepair_stations\x18\x02 \x03(\x0b2!.city.elec.input.v1.RepairStationR\x0erepairStations*\x8c\x03\n\x0cFacilityType\x12\x1d\n\x19FACILITY_TYPE_UNSPECIFIED\x10\x00\x12\x1f\n\x1bFACILITY_TYPE_POWER_STATION\x10\x01\x12!\n\x1dFACILITY_TYPE_TRANSFORMER_500\x10\x02\x12!\n\x1dFACILITY_TYPE_TRANSFORMER_220\x10\x03\x12!\n\x1dFACILITY_TYPE_TRANSFORMER_110\x10\x04\x12 \n\x1cFACILITY_TYPE_TRANSFORMER_10\x10\x05\x12\x1e\n\x1aFACILITY_TYPE_BASE_STATION\x10\x06\x12\x19\n\x15FACILITY_TYPE_GATEWAY\x10\x07\x12\x1f\n\x1bFACILITY_TYPE_DRAINAGE_PUMP\x10\x08\x12\x1f\n\x1bFACILITY_TYPE_TRAFFIC_LIGHT\x10\t\x12\x15\n\x11FACILITY_TYPE_AOI\x10\n\x12\x1d\n\x19FACILITY_TYPE_SUPPLY_PUMP\x10\x0bB\xc8\x01\n\x16com.city.elec.input.v1B\nInputProtoP\x01Z7git.fiblab.net/sim/protos/go/city/elec/input/v1;inputv1\xa2\x02\x03CEI\xaa\x02\x12City.Elec.Input.V1\xca\x02\x12City\\Elec\\Input\\V1\xe2\x02\x1eCity\\Elec\\Input\\V1\\GPBMetadata\xea\x02\x15City::Elec::Input::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.elec.input.v1.input_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x16com.city.elec.input.v1B\nInputProtoP\x01Z7git.fiblab.net/sim/protos/go/city/elec/input/v1;inputv1\xa2\x02\x03CEI\xaa\x02\x12City.Elec.Input.V1\xca\x02\x12City\\Elec\\Input\\V1\xe2\x02\x1eCity\\Elec\\Input\\V1\\GPBMetadata\xea\x02\x15City::Elec::Input::V1'
    _globals['_FACILITYTYPE']._serialized_start = 654
    _globals['_FACILITYTYPE']._serialized_end = 1050
    _globals['_REPAIRSTATION']._serialized_start = 77
    _globals['_REPAIRSTATION']._serialized_end = 173
    _globals['_FACILITY']._serialized_start = 176
    _globals['_FACILITY']._serialized_end = 498
    _globals['_FACILITIES']._serialized_start = 501
    _globals['_FACILITIES']._serialized_end = 651