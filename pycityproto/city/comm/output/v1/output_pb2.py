"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n city/comm/output/v1/output.proto\x12\x13city.comm.output.v1"\xa5\x01\n\x04Node\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x127\n\x06status\x18\x02 \x01(\x0e2\x1f.city.comm.output.v1.NodeStatusR\x06status\x129\n\x16battery_remaining_time\x18\x03 \x01(\x01H\x00R\x14batteryRemainingTime\x88\x01\x01B\x19\n\x17_battery_remaining_time"\xcd\x02\n\x0bBaseStation\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x1f\n\x0bdemand_flow\x18\x02 \x01(\x01R\ndemandFlow\x12\x1f\n\x0bactual_flow\x18\x03 \x01(\x01R\nactualFlow\x12\x1d\n\nnum_agents\x18\x04 \x01(\x05R\tnumAgents\x12\x1a\n\x08overload\x18\x05 \x01(\x08R\x08overload\x12\'\n\x0funsatisfied_num\x18\x06 \x01(\x05R\x0eunsatisfiedNum\x12#\n\rsatisfied_num\x18\x07 \x01(\x05R\x0csatisfiedNum\x12\x1d\n\noutage_num\x18\x08 \x01(\x05R\toutageNum\x12\x1d\n\nactive_num\x18\t \x01(\x05R\tactiveNum\x12%\n\x0etransmit_power\x18\n \x01(\x01R\rtransmitPower"\xac\x01\n\x06Signal\x12\x19\n\x08num_rows\x18\x01 \x01(\x05R\x07numRows\x12\x1f\n\x0bnum_columns\x18\x02 \x01(\x05R\nnumColumns\x12\x1a\n\x08strength\x18\x03 \x03(\x01R\x08strength\x12&\n\x0fbase_station_id\x18\x04 \x03(\x05R\rbaseStationId\x12"\n\rfreq_range_id\x18\x05 \x03(\x05R\x0bfreqRangeId"\x8a\x03\n\x06Person\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x1f\n\x0bdemand_rate\x18\x02 \x01(\x01R\ndemandRate\x12\x1f\n\x0bactual_rate\x18\x03 \x01(\x01R\nactualRate\x12O\n\x0econnect_status\x18\x04 \x01(\x0e2(.city.comm.output.v1.PersonConnectStatusR\rconnectStatus\x12L\n\rdemand_status\x18\t \x01(\x0e2\'.city.comm.output.v1.PersonDemandStatusR\x0cdemandStatus\x12\x1a\n\x08strength\x18\x05 \x01(\x01R\x08strength\x12&\n\x0fbase_station_id\x18\x06 \x01(\x05R\rbaseStationId\x12$\n\x0efreq_range_ids\x18\x07 \x03(\x05R\x0cfreqRangeIds\x12%\n\x0ereceived_power\x18\x08 \x01(\x01R\rreceivedPower"\xec\x01\n\x03Aoi\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id\x12\x1f\n\x0bdemand_flow\x18\x02 \x01(\x01R\ndemandFlow\x12\x1f\n\x0bactual_flow\x18\x03 \x01(\x01R\nactualFlow\x12\x1d\n\noutage_num\x18\x04 \x01(\x05R\toutageNum\x12#\n\rsatisfied_num\x18\x05 \x01(\x05R\x0csatisfiedNum\x12\'\n\x0funsatisfied_num\x18\x06 \x01(\x05R\x0eunsatisfiedNum\x12&\n\x0factive_user_num\x18\x07 \x01(\x05R\ractiveUserNum"\xce\x06\n\nStatistics\x120\n\x14num_satisfied_agents\x18\x01 \x01(\x05R\x12numSatisfiedAgents\x124\n\x16num_unsatisfied_agents\x18\x02 \x01(\x05R\x14numUnsatisfiedAgents\x12*\n\x11num_outage_agents\x18\x03 \x01(\x05R\x0fnumOutageAgents\x12*\n\x11num_active_agents\x18\x04 \x01(\x05R\x0fnumActiveAgents\x12\x1f\n\x0bdemand_flow\x18\x05 \x01(\x01R\ndemandFlow\x12\x1f\n\x0bactual_flow\x18\x06 \x01(\x01R\nactualFlow\x12(\n\x10num_base_station\x18\x07 \x01(\x05R\x0enumBaseStation\x12-\n\x13num_ok_base_station\x18\x08 \x01(\x05R\x10numOkBaseStation\x125\n\x17num_ruined_base_station\x18\t \x01(\x05R\x14numRuinedBaseStation\x127\n\x18num_stopped_base_station\x18\n \x01(\x05R\x15numStoppedBaseStation\x12=\n\x1bnum_overloaded_base_station\x18\x0b \x01(\x05R\x18numOverloadedBaseStation\x12\x1f\n\x0bnum_gateway\x18\x0c \x01(\x05R\nnumGateway\x12$\n\x0enum_ok_gateway\x18\r \x01(\x05R\x0cnumOkGateway\x12,\n\x12num_ruined_gateway\x18\x0e \x01(\x05R\x10numRuinedGateway\x12.\n\x13num_stopped_gateway\x18\x0f \x01(\x05R\x11numStoppedGateway\x124\n\x16num_overloaded_gateway\x18\x10 \x01(\x05R\x14numOverloadedGateway\x12.\n\x13num_battery_gateway\x18\x11 \x01(\x05R\x11numBatteryGateway\x12+\n\x11power_consumption\x18\x12 \x01(\x01R\x10powerConsumption*\x87\x01\n\nNodeStatus\x12\x1b\n\x17NODE_STATUS_UNSPECIFIED\x10\x00\x12\x12\n\x0eNODE_STATUS_OK\x10\x01\x12\x17\n\x13NODE_STATUS_BATTERY\x10\x02\x12\x17\n\x13NODE_STATUS_FAILURE\x10\x03\x12\x16\n\x12NODE_STATUS_RUINED\x10\x04*|\n\x13PersonConnectStatus\x12%\n!PERSON_CONNECT_STATUS_UNSPECIFIED\x10\x00\x12\x1c\n\x18PERSON_CONNECT_STATUS_OK\x10\x01\x12 \n\x1cPERSON_CONNECT_STATUS_OUTAGE\x10\x02*\xa1\x01\n\x12PersonDemandStatus\x12$\n PERSON_DEMAND_STATUS_UNSPECIFIED\x10\x00\x12"\n\x1ePERSON_DEMAND_STATUS_SATISFIED\x10\x01\x12$\n PERSON_DEMAND_STATUS_UNSATISFIED\x10\x02\x12\x1b\n\x17PERSON_DEMAND_STATUS_NO\x10\x03B\xd3\x01\n\x17com.city.comm.output.v1B\x0bOutputProtoP\x01Z<git.fiblab.net/sim/protos/go/v2/city/comm/output/v1;outputv1\xa2\x02\x03CCO\xaa\x02\x13City.Comm.Output.V1\xca\x02\x13City\\Comm\\Output\\V1\xe2\x02\x1fCity\\Comm\\Output\\V1\\GPBMetadata\xea\x02\x16City::Comm::Output::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.comm.output.v1.output_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x17com.city.comm.output.v1B\x0bOutputProtoP\x01Z<git.fiblab.net/sim/protos/go/v2/city/comm/output/v1;outputv1\xa2\x02\x03CCO\xaa\x02\x13City.Comm.Output.V1\xca\x02\x13City\\Comm\\Output\\V1\xe2\x02\x1fCity\\Comm\\Output\\V1\\GPBMetadata\xea\x02\x16City::Comm::Output::V1'
    _globals['_NODESTATUS']._serialized_start = 2222
    _globals['_NODESTATUS']._serialized_end = 2357
    _globals['_PERSONCONNECTSTATUS']._serialized_start = 2359
    _globals['_PERSONCONNECTSTATUS']._serialized_end = 2483
    _globals['_PERSONDEMANDSTATUS']._serialized_start = 2486
    _globals['_PERSONDEMANDSTATUS']._serialized_end = 2647
    _globals['_NODE']._serialized_start = 58
    _globals['_NODE']._serialized_end = 223
    _globals['_BASESTATION']._serialized_start = 226
    _globals['_BASESTATION']._serialized_end = 559
    _globals['_SIGNAL']._serialized_start = 562
    _globals['_SIGNAL']._serialized_end = 734
    _globals['_PERSON']._serialized_start = 737
    _globals['_PERSON']._serialized_end = 1131
    _globals['_AOI']._serialized_start = 1134
    _globals['_AOI']._serialized_end = 1370
    _globals['_STATISTICS']._serialized_start = 1373
    _globals['_STATISTICS']._serialized_end = 2219