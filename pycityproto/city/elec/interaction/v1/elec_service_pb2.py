"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+city/elec/interaction/v1/elec_service.proto\x12\x18city.elec.interaction.v1":\n\x10SetStatusRequest\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n\x06status\x18\x02 \x01(\x08R\x06status"\x13\n\x11SetStatusResponse"!\n\x0fGetPowerRequest\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id"(\n\x10GetPowerResponse\x12\x14\n\x05power\x18\x01 \x01(\x01R\x05power"+\n\x15GetPowerStatusRequest\x12\x12\n\x04flag\x18\x01 \x01(\x05R\x04flag"\xbe\x01\n\x16GetPowerStatusResponse\x12d\n\x0cpower_status\x18\x01 \x03(\x0b2A.city.elec.interaction.v1.GetPowerStatusResponse.PowerStatusEntryR\x0bpowerStatus\x1a>\n\x10PowerStatusEntry\x12\x10\n\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n\x05value\x18\x02 \x01(\x01R\x05value:\x028\x01"*\n\x14GetNoPowerAOIRequest\x12\x12\n\x04flag\x18\x01 \x01(\x05R\x04flag")\n\x15GetNoPowerAOIResponse\x12\x10\n\x03aoi\x18\x01 \x03(\x05R\x03aoi"\x14\n\x12GetRuinInfoRequest"2\n\x08RuinInfo\x12\x10\n\x03num\x18\x01 \x01(\x05R\x03num\x12\x14\n\x05ratio\x18\x02 \x01(\x01R\x05ratio"\xbb\x01\n\x13GetRuinInfoResponse\x124\n\x03one\x18\x01 \x01(\x0b2".city.elec.interaction.v1.RuinInfoR\x03one\x124\n\x03two\x18\x02 \x01(\x0b2".city.elec.interaction.v1.RuinInfoR\x03two\x128\n\x05three\x18\x03 \x01(\x0b2".city.elec.interaction.v1.RuinInfoR\x05three"\x16\n\x14GetEdgeStatusRequest"e\n\x15GetEdgeStatusResponse\x12\x18\n\x07reason1\x18\x01 \x03(\tR\x07reason1\x12\x18\n\x07reason2\x18\x02 \x03(\tR\x07reason2\x12\x18\n\x07reason3\x18\x03 \x03(\tR\x07reason32\x9b\x05\n\x0bElecService\x12d\n\tSetStatus\x12*.city.elec.interaction.v1.SetStatusRequest\x1a+.city.elec.interaction.v1.SetStatusResponse\x12a\n\x08GetPower\x12).city.elec.interaction.v1.GetPowerRequest\x1a*.city.elec.interaction.v1.GetPowerResponse\x12s\n\x0eGetPowerStatus\x12/.city.elec.interaction.v1.GetPowerStatusRequest\x1a0.city.elec.interaction.v1.GetPowerStatusResponse\x12p\n\rGetNoPowerAOI\x12..city.elec.interaction.v1.GetNoPowerAOIRequest\x1a/.city.elec.interaction.v1.GetNoPowerAOIResponse\x12j\n\x0bGetRuinInfo\x12,.city.elec.interaction.v1.GetRuinInfoRequest\x1a-.city.elec.interaction.v1.GetRuinInfoResponse\x12p\n\rGetEdgeStatus\x12..city.elec.interaction.v1.GetEdgeStatusRequest\x1a/.city.elec.interaction.v1.GetEdgeStatusResponseB\xf8\x01\n\x1ccom.city.elec.interaction.v1B\x10ElecServiceProtoP\x01ZCgit.fiblab.net/sim/protos/go/city/elec/interaction/v1;interactionv1\xa2\x02\x03CEI\xaa\x02\x18City.Elec.Interaction.V1\xca\x02\x18City\\Elec\\Interaction\\V1\xe2\x02$City\\Elec\\Interaction\\V1\\GPBMetadata\xea\x02\x1bCity::Elec::Interaction::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.elec.interaction.v1.elec_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x1ccom.city.elec.interaction.v1B\x10ElecServiceProtoP\x01ZCgit.fiblab.net/sim/protos/go/city/elec/interaction/v1;interactionv1\xa2\x02\x03CEI\xaa\x02\x18City.Elec.Interaction.V1\xca\x02\x18City\\Elec\\Interaction\\V1\xe2\x02$City\\Elec\\Interaction\\V1\\GPBMetadata\xea\x02\x1bCity::Elec::Interaction::V1'
    _globals['_GETPOWERSTATUSRESPONSE_POWERSTATUSENTRY']._options = None
    _globals['_GETPOWERSTATUSRESPONSE_POWERSTATUSENTRY']._serialized_options = b'8\x01'
    _globals['_SETSTATUSREQUEST']._serialized_start = 73
    _globals['_SETSTATUSREQUEST']._serialized_end = 131
    _globals['_SETSTATUSRESPONSE']._serialized_start = 133
    _globals['_SETSTATUSRESPONSE']._serialized_end = 152
    _globals['_GETPOWERREQUEST']._serialized_start = 154
    _globals['_GETPOWERREQUEST']._serialized_end = 187
    _globals['_GETPOWERRESPONSE']._serialized_start = 189
    _globals['_GETPOWERRESPONSE']._serialized_end = 229
    _globals['_GETPOWERSTATUSREQUEST']._serialized_start = 231
    _globals['_GETPOWERSTATUSREQUEST']._serialized_end = 274
    _globals['_GETPOWERSTATUSRESPONSE']._serialized_start = 277
    _globals['_GETPOWERSTATUSRESPONSE']._serialized_end = 467
    _globals['_GETPOWERSTATUSRESPONSE_POWERSTATUSENTRY']._serialized_start = 405
    _globals['_GETPOWERSTATUSRESPONSE_POWERSTATUSENTRY']._serialized_end = 467
    _globals['_GETNOPOWERAOIREQUEST']._serialized_start = 469
    _globals['_GETNOPOWERAOIREQUEST']._serialized_end = 511
    _globals['_GETNOPOWERAOIRESPONSE']._serialized_start = 513
    _globals['_GETNOPOWERAOIRESPONSE']._serialized_end = 554
    _globals['_GETRUININFOREQUEST']._serialized_start = 556
    _globals['_GETRUININFOREQUEST']._serialized_end = 576
    _globals['_RUININFO']._serialized_start = 578
    _globals['_RUININFO']._serialized_end = 628
    _globals['_GETRUININFORESPONSE']._serialized_start = 631
    _globals['_GETRUININFORESPONSE']._serialized_end = 818
    _globals['_GETEDGESTATUSREQUEST']._serialized_start = 820
    _globals['_GETEDGESTATUSREQUEST']._serialized_end = 842
    _globals['_GETEDGESTATUSRESPONSE']._serialized_start = 844
    _globals['_GETEDGESTATUSRESPONSE']._serialized_end = 945
    _globals['_ELECSERVICE']._serialized_start = 948
    _globals['_ELECSERVICE']._serialized_end = 1615