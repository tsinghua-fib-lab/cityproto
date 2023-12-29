"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.geo.v2 import geo_pb2 as city_dot_geo_dot_v2_dot_geo__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dcity/wargame/v1/wargame.proto\x12\x0fcity.wargame.v1\x1a\x15city/geo/v2/geo.proto"j\n\tRecoPoint\x12.\n\x03pos\x18\x01 \x01(\x0b2\x1c.city.geo.v2.LongLatPositionR\x03pos\x12-\n\x04type\x18\x02 \x01(\x0e2\x19.city.wargame.v1.RecoTypeR\x04type"\xb8\x01\n\x06Weapon\x12>\n\x08red_type\x18\x02 \x01(\x0e2\x1e.city.wargame.v1.RedWeaponTypeH\x00R\x07redType\x88\x01\x01\x12A\n\tblue_type\x18\x03 \x01(\x0e2\x1f.city.wargame.v1.BlueWeaponTypeH\x01R\x08blueType\x88\x01\x01\x12\x10\n\x03num\x18\x04 \x01(\x05R\x03numB\x0b\n\t_red_typeB\x0c\n\n_blue_type"\x80\x01\n\x05Point\x12/\n\x06weapon\x18\x01 \x01(\x0b2\x17.city.wargame.v1.WeaponR\x06weapon\x12.\n\x03pos\x18\x02 \x01(\x0b2\x1c.city.geo.v2.LongLatPositionR\x03pos\x12\x16\n\x06radius\x18\x03 \x01(\x01R\x06radius"\x99\x01\n\x03Hit\x12.\n\x06attack\x18\x01 \x01(\x0b2\x16.city.wargame.v1.PointR\x06attack\x122\n\x08defenses\x18\x02 \x03(\x0b2\x16.city.wargame.v1.PointR\x08defenses\x12\x16\n\x06result\x18\x03 \x01(\x08R\x06result\x12\x16\n\x06arrive\x18\x04 \x01(\x05R\x06arrive"6\n\nHitHistory\x12(\n\x04hits\x18\x01 \x03(\x0b2\x14.city.wargame.v1.HitR\x04hits"3\n\x08Casualty\x12\x15\n\x06aoi_id\x18\x01 \x01(\x05R\x05aoiId\x12\x10\n\x03num\x18\x02 \x01(\x05R\x03num*\xc0\x01\n\rRedWeaponType\x12\x1f\n\x1bRED_WEAPON_TYPE_UNSPECIFIED\x10\x00\x12\x1b\n\x17RED_WEAPON_TYPE_AGM_158\x10\x01\x12\x1b\n\x17RED_WEAPON_TYPE_AGM_183\x10\x02\x12\x1a\n\x16RED_WEAPON_TYPE_AGM_88\x10\x03\x12\x1b\n\x17RED_WEAPON_TYPE_UGM_109\x10\x04\x12\x1b\n\x17RED_WEAPON_TYPE_AGM_142\x10\x05*\xa1\x01\n\x0eBlueWeaponType\x12 \n\x1cBLUE_WEAPON_TYPE_UNSPECIFIED\x10\x00\x12\x19\n\x15BLUE_WEAPON_TYPE_HQ_9\x10\x01\x12\x1a\n\x16BLUE_WEAPON_TYPE_HQ_12\x10\x02\x12\x1a\n\x16BLUE_WEAPON_TYPE_HQ_16\x10\x03\x12\x1a\n\x16BLUE_WEAPON_TYPE_HQ_22\x10\x04*9\n\x04Camp\x12\x14\n\x10CAMP_UNSPECIFIED\x10\x00\x12\x0c\n\x08CAMP_RED\x10\x01\x12\r\n\tCAMP_BLUE\x10\x02*P\n\x08RecoType\x12\x19\n\x15RECO_TYPE_UNSPECIFIED\x10\x00\x12\x15\n\x11RECO_TYPE_MISSILE\x10\x01\x12\x12\n\x0eRECO_TYPE_ELEC\x10\x02*\x89\x01\n\x0cRecoAlgoType\x12\x1e\n\x1aRECO_ALGO_TYPE_UNSPECIFIED\x10\x00\x12\x1e\n\x1aRECO_ALGO_TYPE_MAX_DESTORY\x10\x01\x12\x19\n\x15RECO_ALGO_TYPE_RANDOM\x10\x02\x12\x1e\n\x1aRECO_ALGO_TYPE_MIN_DESTORY\x10\x03B\xb9\x01\n\x13com.city.wargame.v1B\x0cWargameProtoP\x01Z6git.fiblab.net/sim/protos/go/city/wargame/v1;wargamev1\xa2\x02\x03CWX\xaa\x02\x0fCity.Wargame.V1\xca\x02\x0fCity\\Wargame\\V1\xe2\x02\x1bCity\\Wargame\\V1\\GPBMetadata\xea\x02\x11City::Wargame::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.wargame.v1.wargame_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.wargame.v1B\x0cWargameProtoP\x01Z6git.fiblab.net/sim/protos/go/city/wargame/v1;wargamev1\xa2\x02\x03CWX\xaa\x02\x0fCity.Wargame.V1\xca\x02\x0fCity\\Wargame\\V1\xe2\x02\x1bCity\\Wargame\\V1\\GPBMetadata\xea\x02\x11City::Wargame::V1'
    _globals['_REDWEAPONTYPE']._serialized_start = 765
    _globals['_REDWEAPONTYPE']._serialized_end = 957
    _globals['_BLUEWEAPONTYPE']._serialized_start = 960
    _globals['_BLUEWEAPONTYPE']._serialized_end = 1121
    _globals['_CAMP']._serialized_start = 1123
    _globals['_CAMP']._serialized_end = 1180
    _globals['_RECOTYPE']._serialized_start = 1182
    _globals['_RECOTYPE']._serialized_end = 1262
    _globals['_RECOALGOTYPE']._serialized_start = 1265
    _globals['_RECOALGOTYPE']._serialized_end = 1402
    _globals['_RECOPOINT']._serialized_start = 73
    _globals['_RECOPOINT']._serialized_end = 179
    _globals['_WEAPON']._serialized_start = 182
    _globals['_WEAPON']._serialized_end = 366
    _globals['_POINT']._serialized_start = 369
    _globals['_POINT']._serialized_end = 497
    _globals['_HIT']._serialized_start = 500
    _globals['_HIT']._serialized_end = 653
    _globals['_HITHISTORY']._serialized_start = 655
    _globals['_HITHISTORY']._serialized_end = 709
    _globals['_CASUALTY']._serialized_start = 711
    _globals['_CASUALTY']._serialized_end = 762