"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#city/streetview/v1/streetview.proto\x12\x12city.streetview.v1\x1a\x15city/geo/v2/geo.proto"C\n\x0fStreetViewImage\x12\x18\n\x07heading\x18\x01 \x01(\x01R\x07heading\x12\x16\n\x06object\x18\x02 \x01(\tR\x06object"\x7f\n\nStreetView\x124\n\x06lnglat\x18\x01 \x01(\x0b2\x1c.city.geo.v2.LongLatPositionR\x06lnglat\x12;\n\x06images\x18\x02 \x03(\x0b2#.city.streetview.v1.StreetViewImageR\x06images"P\n\x0bStreetViews\x12A\n\x0cstreet_views\x18\x01 \x03(\x0b2\x1e.city.streetview.v1.StreetViewR\x0bstreetViewsB\xd1\x01\n\x16com.city.streetview.v1B\x0fStreetviewProtoP\x01Z<git.fiblab.net/sim/protos/go/city/streetview/v1;streetviewv1\xa2\x02\x03CSX\xaa\x02\x12City.Streetview.V1\xca\x02\x12City\\Streetview\\V1\xe2\x02\x1eCity\\Streetview\\V1\\GPBMetadata\xea\x02\x14City::Streetview::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.streetview.v1.streetview_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x16com.city.streetview.v1B\x0fStreetviewProtoP\x01Z<git.fiblab.net/sim/protos/go/city/streetview/v1;streetviewv1\xa2\x02\x03CSX\xaa\x02\x12City.Streetview.V1\xca\x02\x12City\\Streetview\\V1\xe2\x02\x1eCity\\Streetview\\V1\\GPBMetadata\xea\x02\x14City::Streetview::V1'
    _globals['_STREETVIEWIMAGE']._serialized_start = 82
    _globals['_STREETVIEWIMAGE']._serialized_end = 149
    _globals['_STREETVIEW']._serialized_start = 151
    _globals['_STREETVIEW']._serialized_end = 278
    _globals['_STREETVIEWS']._serialized_start = 280
    _globals['_STREETVIEWS']._serialized_end = 360