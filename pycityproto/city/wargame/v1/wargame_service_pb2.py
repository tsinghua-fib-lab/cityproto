"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.wargame.v1 import wargame_pb2 as city_dot_wargame_dot_v1_dot_wargame__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%city/wargame/v1/wargame_service.proto\x12\x0fcity.wargame.v1\x1a\x1dcity/wargame/v1/wargame.proto"n\n\x11PickPointsRequest\x12)\n\x04camp\x18\x01 \x01(\x0e2\x15.city.wargame.v1.CampR\x04camp\x12.\n\x06points\x18\x02 \x03(\x0b2\x16.city.wargame.v1.PointR\x06points"\x14\n\x12PickPointsResponse"A\n\x14GetPickPointsRequest\x12)\n\x04camp\x18\x01 \x01(\x0e2\x15.city.wargame.v1.CampR\x04camp"G\n\x15GetPickPointsResponse\x12.\n\x06points\x18\x01 \x03(\x0b2\x16.city.wargame.v1.PointR\x06points"\xf2\x04\n\x15SetScoreWeightRequest\x12\x14\n\x05money\x18\x01 \x01(\x01R\x05money\x12\'\n\x0fpopulation_loss\x18\x02 \x01(\x01R\x0epopulationLoss\x12\x1d\n\nelec_power\x18\x03 \x01(\x01R\telecPower\x12$\n\x0eelec_distory_1\x18\x04 \x01(\x01R\x0celecDistory1\x12&\n\x0fwater_distory_1\x18\x05 \x01(\x01R\rwaterDistory1\x12 \n\x0cbs_distory_1\x18\x06 \x01(\x01R\nbsDistory1\x12*\n\x11traffic_distory_1\x18\x07 \x01(\x01R\x0ftrafficDistory1\x12$\n\x0eelec_distory_2\x18\x08 \x01(\x01R\x0celecDistory2\x12&\n\x0fwater_distory_2\x18\t \x01(\x01R\rwaterDistory2\x12 \n\x0cbs_distory_2\x18\n \x01(\x01R\nbsDistory2\x12*\n\x11traffic_distory_2\x18\x0b \x01(\x01R\x0ftrafficDistory2\x12$\n\x0eelec_distory_3\x18\x0c \x01(\x01R\x0celecDistory3\x12&\n\x0fwater_distory_3\x18\r \x01(\x01R\rwaterDistory3\x12 \n\x0cbs_distory_3\x18\x0e \x01(\x01R\nbsDistory3\x12*\n\x11traffic_distory_3\x18\x0f \x01(\x01R\x0ftrafficDistory3\x12\'\n\x0fdefense_success\x18\x10 \x01(\x01R\x0edefenseSuccess"\x18\n\x16SetScoreWeightResponse"\xc2\x01\n\x17GiveDefenseOrderRequest\x12#\n\rweight_radius\x18\x01 \x01(\x01R\x0cweightRadius\x12\x1d\n\nweight_550\x18\x02 \x01(\x01R\tweight550\x12\x1d\n\nweight_220\x18\x03 \x01(\x01R\tweight220\x12\x1d\n\nweight_110\x18\x04 \x01(\x01R\tweight110\x12%\n\x0eprob_threshold\x18\x05 \x01(\x01R\rprobThreshold"\x1a\n\x18GiveDefenseOrderResponse"\x16\n\x14GetHitHistoryRequest"R\n\x15GetHitHistoryResponse\x129\n\thistories\x18\x01 \x03(\x0b2\x1b.city.wargame.v1.HitHistoryR\thistories"t\n\x14GetRecoPointsRequest\x12)\n\x04camp\x18\x01 \x01(\x0e2\x15.city.wargame.v1.CampR\x04camp\x121\n\x04type\x18\x02 \x01(\x0e2\x1d.city.wargame.v1.RecoAlgoTypeR\x04type"K\n\x15GetRecoPointsResponse\x122\n\x06points\x18\x01 \x03(\x0b2\x1a.city.wargame.v1.RecoPointR\x06points"\x10\n\x0eGetStepRequest"\xb1\x01\n\x0fGetStepResponse\x12\x12\n\x04step\x18\x01 \x01(\x05R\x04step\x12$\n\x0ered_pick_ready\x18\x02 \x01(\x08R\x0credPickReady\x12&\n\x0fblue_pick_ready\x18\x03 \x01(\x08R\rbluePickReady\x12\x14\n\x05round\x18\x04 \x01(\x05R\x05round\x12&\n\x0fis_game_started\x18\x05 \x01(\x08R\risGameStarted"\x16\n\x14GetCasualtiesRequest"R\n\x15GetCasualtiesResponse\x129\n\ncasualties\x18\x01 \x03(\x0b2\x19.city.wargame.v1.CasualtyR\ncasualties2\x81\x06\n\x0eWarGameService\x12U\n\nPickPoints\x12".city.wargame.v1.PickPointsRequest\x1a#.city.wargame.v1.PickPointsResponse\x12^\n\rGetPickPoints\x12%.city.wargame.v1.GetPickPointsRequest\x1a&.city.wargame.v1.GetPickPointsResponse\x12g\n\x10GiveDefenseOrder\x12(.city.wargame.v1.GiveDefenseOrderRequest\x1a).city.wargame.v1.GiveDefenseOrderResponse\x12a\n\x0eSetScoreWeight\x12&.city.wargame.v1.SetScoreWeightRequest\x1a\'.city.wargame.v1.SetScoreWeightResponse\x12^\n\rGetHitHistory\x12%.city.wargame.v1.GetHitHistoryRequest\x1a&.city.wargame.v1.GetHitHistoryResponse\x12^\n\rGetRecoPoints\x12%.city.wargame.v1.GetRecoPointsRequest\x1a&.city.wargame.v1.GetRecoPointsResponse\x12L\n\x07GetStep\x12\x1f.city.wargame.v1.GetStepRequest\x1a .city.wargame.v1.GetStepResponse\x12^\n\rGetCasualties\x12%.city.wargame.v1.GetCasualtiesRequest\x1a&.city.wargame.v1.GetCasualtiesResponseB\xc0\x01\n\x13com.city.wargame.v1B\x13WargameServiceProtoP\x01Z6git.fiblab.net/sim/protos/go/city/wargame/v1;wargamev1\xa2\x02\x03CWX\xaa\x02\x0fCity.Wargame.V1\xca\x02\x0fCity\\Wargame\\V1\xe2\x02\x1bCity\\Wargame\\V1\\GPBMetadata\xea\x02\x11City::Wargame::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.wargame.v1.wargame_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x13com.city.wargame.v1B\x13WargameServiceProtoP\x01Z6git.fiblab.net/sim/protos/go/city/wargame/v1;wargamev1\xa2\x02\x03CWX\xaa\x02\x0fCity.Wargame.V1\xca\x02\x0fCity\\Wargame\\V1\xe2\x02\x1bCity\\Wargame\\V1\\GPBMetadata\xea\x02\x11City::Wargame::V1'
    _globals['_PICKPOINTSREQUEST']._serialized_start = 89
    _globals['_PICKPOINTSREQUEST']._serialized_end = 199
    _globals['_PICKPOINTSRESPONSE']._serialized_start = 201
    _globals['_PICKPOINTSRESPONSE']._serialized_end = 221
    _globals['_GETPICKPOINTSREQUEST']._serialized_start = 223
    _globals['_GETPICKPOINTSREQUEST']._serialized_end = 288
    _globals['_GETPICKPOINTSRESPONSE']._serialized_start = 290
    _globals['_GETPICKPOINTSRESPONSE']._serialized_end = 361
    _globals['_SETSCOREWEIGHTREQUEST']._serialized_start = 364
    _globals['_SETSCOREWEIGHTREQUEST']._serialized_end = 990
    _globals['_SETSCOREWEIGHTRESPONSE']._serialized_start = 992
    _globals['_SETSCOREWEIGHTRESPONSE']._serialized_end = 1016
    _globals['_GIVEDEFENSEORDERREQUEST']._serialized_start = 1019
    _globals['_GIVEDEFENSEORDERREQUEST']._serialized_end = 1213
    _globals['_GIVEDEFENSEORDERRESPONSE']._serialized_start = 1215
    _globals['_GIVEDEFENSEORDERRESPONSE']._serialized_end = 1241
    _globals['_GETHITHISTORYREQUEST']._serialized_start = 1243
    _globals['_GETHITHISTORYREQUEST']._serialized_end = 1265
    _globals['_GETHITHISTORYRESPONSE']._serialized_start = 1267
    _globals['_GETHITHISTORYRESPONSE']._serialized_end = 1349
    _globals['_GETRECOPOINTSREQUEST']._serialized_start = 1351
    _globals['_GETRECOPOINTSREQUEST']._serialized_end = 1467
    _globals['_GETRECOPOINTSRESPONSE']._serialized_start = 1469
    _globals['_GETRECOPOINTSRESPONSE']._serialized_end = 1544
    _globals['_GETSTEPREQUEST']._serialized_start = 1546
    _globals['_GETSTEPREQUEST']._serialized_end = 1562
    _globals['_GETSTEPRESPONSE']._serialized_start = 1565
    _globals['_GETSTEPRESPONSE']._serialized_end = 1742
    _globals['_GETCASUALTIESREQUEST']._serialized_start = 1744
    _globals['_GETCASUALTIESREQUEST']._serialized_end = 1766
    _globals['_GETCASUALTIESRESPONSE']._serialized_start = 1768
    _globals['_GETCASUALTIESRESPONSE']._serialized_end = 1850
    _globals['_WARGAMESERVICE']._serialized_start = 1853
    _globals['_WARGAMESERVICE']._serialized_end = 2622