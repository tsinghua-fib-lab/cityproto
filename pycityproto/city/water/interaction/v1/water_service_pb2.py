"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n-city/water/interaction/v1/water_service.proto\x12\x19city.water.interaction.v1"\x85\x01\n\x19SetPumpPowerStatusRequest\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n\x06status\x18\x02 \x01(\x08R\x06status\x12@\n\x04type\x18\x03 \x01(\x0e2,.city.water.interaction.v1.WaterFacilityTypeR\x04type"\x1c\n\x1aSetPumpPowerStatusResponse"\x87\x01\n\x1bSetPumpNetworkStatusRequest\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n\x06status\x18\x02 \x01(\x08R\x06status\x12@\n\x04type\x18\x03 \x01(\x0e2,.city.water.interaction.v1.WaterFacilityTypeR\x04type"\x1e\n\x1cSetPumpNetworkStatusResponse"\x80\x01\n\x14SetPumpStatusRequest\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n\x06status\x18\x02 \x01(\x08R\x06status\x12@\n\x04type\x18\x03 \x01(\x0e2,.city.water.interaction.v1.WaterFacilityTypeR\x04type"\x17\n\x15SetPumpStatusResponse"*\n\x14GetPumpStatusRequest\x12\x12\n\x04flag\x18\x01 \x01(\x05R\x04flag"\xb9\x01\n\x15GetPumpStatusResponse\x12a\n\x0bpump_status\x18\x01 \x03(\x0b2@.city.water.interaction.v1.GetPumpStatusResponse.PumpStatusEntryR\npumpStatus\x1a=\n\x0fPumpStatusEntry\x12\x10\n\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n\x05value\x18\x02 \x01(\x05R\x05value:\x028\x01"*\n\x14GetNoWaterAOIRequest\x12\x12\n\x04flag\x18\x01 \x01(\x05R\x04flag")\n\x15GetNoWaterAOIResponse\x12\x10\n\x03aoi\x18\x01 \x03(\x05R\x03aoi"\x14\n\x12GetRuinInfoRequest"2\n\x08RuinInfo\x12\x10\n\x03num\x18\x01 \x01(\x05R\x03num\x12\x14\n\x05ratio\x18\x02 \x01(\x01R\x05ratio"\xbe\x01\n\x13GetRuinInfoResponse\x125\n\x03one\x18\x01 \x01(\x0b2#.city.water.interaction.v1.RuinInfoR\x03one\x125\n\x03two\x18\x02 \x01(\x0b2#.city.water.interaction.v1.RuinInfoR\x03two\x129\n\x05three\x18\x03 \x01(\x0b2#.city.water.interaction.v1.RuinInfoR\x05three*z\n\x11WaterFacilityType\x12#\n\x1fWATER_FACILITY_TYPE_UNSPECIFIED\x10\x00\x12\x1e\n\x1aWATER_FACILITY_TYPE_SUPPLY\x10\x01\x12 \n\x1cWATER_FACILITY_TYPE_DRAINAGE\x10\x022\xf2\x05\n\x0cWaterService\x12\x83\x01\n\x12SetPumpPowerStatus\x124.city.water.interaction.v1.SetPumpPowerStatusRequest\x1a5.city.water.interaction.v1.SetPumpPowerStatusResponse"\x00\x12\x89\x01\n\x14SetPumpNetworkStatus\x126.city.water.interaction.v1.SetPumpNetworkStatusRequest\x1a7.city.water.interaction.v1.SetPumpNetworkStatusResponse"\x00\x12t\n\rSetPumpStatus\x12/.city.water.interaction.v1.SetPumpStatusRequest\x1a0.city.water.interaction.v1.SetPumpStatusResponse"\x00\x12t\n\rGetPumpStatus\x12/.city.water.interaction.v1.GetPumpStatusRequest\x1a0.city.water.interaction.v1.GetPumpStatusResponse"\x00\x12t\n\rGetNoWaterAOI\x12/.city.water.interaction.v1.GetNoWaterAOIRequest\x1a0.city.water.interaction.v1.GetNoWaterAOIResponse"\x00\x12n\n\x0bGetRuinInfo\x12-.city.water.interaction.v1.GetRuinInfoRequest\x1a..city.water.interaction.v1.GetRuinInfoResponse"\x00B\xff\x01\n\x1dcom.city.water.interaction.v1B\x11WaterServiceProtoP\x01ZDgit.fiblab.net/sim/protos/go/city/water/interaction/v1;interactionv1\xa2\x02\x03CWI\xaa\x02\x19City.Water.Interaction.V1\xca\x02\x19City\\Water\\Interaction\\V1\xe2\x02%City\\Water\\Interaction\\V1\\GPBMetadata\xea\x02\x1cCity::Water::Interaction::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.water.interaction.v1.water_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x1dcom.city.water.interaction.v1B\x11WaterServiceProtoP\x01ZDgit.fiblab.net/sim/protos/go/city/water/interaction/v1;interactionv1\xa2\x02\x03CWI\xaa\x02\x19City.Water.Interaction.V1\xca\x02\x19City\\Water\\Interaction\\V1\xe2\x02%City\\Water\\Interaction\\V1\\GPBMetadata\xea\x02\x1cCity::Water::Interaction::V1'
    _globals['_GETPUMPSTATUSRESPONSE_PUMPSTATUSENTRY']._options = None
    _globals['_GETPUMPSTATUSRESPONSE_PUMPSTATUSENTRY']._serialized_options = b'8\x01'
    _globals['_WATERFACILITYTYPE']._serialized_start = 1154
    _globals['_WATERFACILITYTYPE']._serialized_end = 1276
    _globals['_SETPUMPPOWERSTATUSREQUEST']._serialized_start = 77
    _globals['_SETPUMPPOWERSTATUSREQUEST']._serialized_end = 210
    _globals['_SETPUMPPOWERSTATUSRESPONSE']._serialized_start = 212
    _globals['_SETPUMPPOWERSTATUSRESPONSE']._serialized_end = 240
    _globals['_SETPUMPNETWORKSTATUSREQUEST']._serialized_start = 243
    _globals['_SETPUMPNETWORKSTATUSREQUEST']._serialized_end = 378
    _globals['_SETPUMPNETWORKSTATUSRESPONSE']._serialized_start = 380
    _globals['_SETPUMPNETWORKSTATUSRESPONSE']._serialized_end = 410
    _globals['_SETPUMPSTATUSREQUEST']._serialized_start = 413
    _globals['_SETPUMPSTATUSREQUEST']._serialized_end = 541
    _globals['_SETPUMPSTATUSRESPONSE']._serialized_start = 543
    _globals['_SETPUMPSTATUSRESPONSE']._serialized_end = 566
    _globals['_GETPUMPSTATUSREQUEST']._serialized_start = 568
    _globals['_GETPUMPSTATUSREQUEST']._serialized_end = 610
    _globals['_GETPUMPSTATUSRESPONSE']._serialized_start = 613
    _globals['_GETPUMPSTATUSRESPONSE']._serialized_end = 798
    _globals['_GETPUMPSTATUSRESPONSE_PUMPSTATUSENTRY']._serialized_start = 737
    _globals['_GETPUMPSTATUSRESPONSE_PUMPSTATUSENTRY']._serialized_end = 798
    _globals['_GETNOWATERAOIREQUEST']._serialized_start = 800
    _globals['_GETNOWATERAOIREQUEST']._serialized_end = 842
    _globals['_GETNOWATERAOIRESPONSE']._serialized_start = 844
    _globals['_GETNOWATERAOIRESPONSE']._serialized_end = 885
    _globals['_GETRUININFOREQUEST']._serialized_start = 887
    _globals['_GETRUININFOREQUEST']._serialized_end = 907
    _globals['_RUININFO']._serialized_start = 909
    _globals['_RUININFO']._serialized_end = 959
    _globals['_GETRUININFORESPONSE']._serialized_start = 962
    _globals['_GETRUININFORESPONSE']._serialized_end = 1152
    _globals['_WATERSERVICE']._serialized_start = 1279
    _globals['_WATERSERVICE']._serialized_end = 2033