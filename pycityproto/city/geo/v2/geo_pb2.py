"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15city/geo/v2/geo.proto\x12\x0bcity.geo.v2"K\n\x0fLongLatPosition\x12\x1c\n\tlongitude\x18\x01 \x01(\x01R\tlongitude\x12\x1a\n\x08latitude\x18\x02 \x01(\x01R\x08latitude"(\n\nXYPosition\x12\x0c\n\x01x\x18\x01 \x01(\x01R\x01x\x12\x0c\n\x01y\x18\x02 \x01(\x01R\x01y"5\n\x0cLanePosition\x12\x17\n\x07lane_id\x18\x01 \x01(\x05R\x06laneId\x12\x0c\n\x01s\x18\x02 \x01(\x01R\x01s"K\n\x0bAoiPosition\x12\x15\n\x06aoi_id\x18\x01 \x01(\x05R\x05aoiId\x12\x1a\n\x06poi_id\x18\x02 \x01(\x05H\x00R\x05poiId\x88\x01\x01B\t\n\x07_poi_id"\xe6\x02\n\x08Position\x12C\n\rlane_position\x18\x01 \x01(\x0b2\x19.city.geo.v2.LanePositionH\x00R\x0clanePosition\x88\x01\x01\x12@\n\x0caoi_position\x18\x02 \x01(\x0b2\x18.city.geo.v2.AoiPositionH\x01R\x0baoiPosition\x88\x01\x01\x12L\n\x10longlat_position\x18\x03 \x01(\x0b2\x1c.city.geo.v2.LongLatPositionH\x02R\x0flonglatPosition\x88\x01\x01\x12=\n\x0bxy_position\x18\x04 \x01(\x0b2\x17.city.geo.v2.XYPositionH\x03R\nxyPosition\x88\x01\x01B\x10\n\x0e_lane_positionB\x0f\n\r_aoi_positionB\x13\n\x11_longlat_positionB\x0e\n\x0c_xy_position"m\n\x0fLongLatRectArea\x12,\n\x02ne\x18\x01 \x01(\x0b2\x1c.city.geo.v2.LongLatPositionR\x02ne\x12,\n\x02sw\x18\x02 \x01(\x0b2\x1c.city.geo.v2.LongLatPositionR\x02swB\x99\x01\n\x0fcom.city.geo.v2B\x08GeoProtoP\x01Z.git.fiblab.net/sim/protos/go/city/geo/v2;geov2\xa2\x02\x03CGX\xaa\x02\x0bCity.Geo.V2\xca\x02\x0bCity\\Geo\\V2\xe2\x02\x17City\\Geo\\V2\\GPBMetadata\xea\x02\rCity::Geo::V2b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.geo.v2.geo_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x0fcom.city.geo.v2B\x08GeoProtoP\x01Z.git.fiblab.net/sim/protos/go/city/geo/v2;geov2\xa2\x02\x03CGX\xaa\x02\x0bCity.Geo.V2\xca\x02\x0bCity\\Geo\\V2\xe2\x02\x17City\\Geo\\V2\\GPBMetadata\xea\x02\rCity::Geo::V2'
    _globals['_LONGLATPOSITION']._serialized_start = 38
    _globals['_LONGLATPOSITION']._serialized_end = 113
    _globals['_XYPOSITION']._serialized_start = 115
    _globals['_XYPOSITION']._serialized_end = 155
    _globals['_LANEPOSITION']._serialized_start = 157
    _globals['_LANEPOSITION']._serialized_end = 210
    _globals['_AOIPOSITION']._serialized_start = 212
    _globals['_AOIPOSITION']._serialized_end = 287
    _globals['_POSITION']._serialized_start = 290
    _globals['_POSITION']._serialized_end = 648
    _globals['_LONGLATRECTAREA']._serialized_start = 650
    _globals['_LONGLATRECTAREA']._serialized_end = 759