"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from ....city.social.v1 import message_pb2 as city_dot_social_dot_v1_dot_message__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#city/social/v1/social_service.proto\x12\x0ecity.social.v1\x1a\x1ccity/social/v1/message.proto"B\n\x0bSendRequest\x123\n\x08messages\x18\x01 \x03(\x0b2\x17.city.social.v1.MessageR\x08messages"\x0e\n\x0cSendResponse" \n\x0eReceiveRequest\x12\x0e\n\x02id\x18\x01 \x01(\x05R\x02id"F\n\x0fReceiveResponse\x123\n\x08messages\x18\x01 \x03(\x0b2\x17.city.social.v1.MessageR\x08messages2\x9e\x01\n\rSocialService\x12A\n\x04Send\x12\x1b.city.social.v1.SendRequest\x1a\x1c.city.social.v1.SendResponse\x12J\n\x07Receive\x12\x1e.city.social.v1.ReceiveRequest\x1a\x1f.city.social.v1.ReceiveResponseB\xb8\x01\n\x12com.city.social.v1B\x12SocialServiceProtoP\x01Z4git.fiblab.net/sim/protos/go/city/social/v1;socialv1\xa2\x02\x03CSX\xaa\x02\x0eCity.Social.V1\xca\x02\x0eCity\\Social\\V1\xe2\x02\x1aCity\\Social\\V1\\GPBMetadata\xea\x02\x10City::Social::V1b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'city.social.v1.social_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals['DESCRIPTOR']._options = None
    _globals['DESCRIPTOR']._serialized_options = b'\n\x12com.city.social.v1B\x12SocialServiceProtoP\x01Z4git.fiblab.net/sim/protos/go/city/social/v1;socialv1\xa2\x02\x03CSX\xaa\x02\x0eCity.Social.V1\xca\x02\x0eCity\\Social\\V1\xe2\x02\x1aCity\\Social\\V1\\GPBMetadata\xea\x02\x10City::Social::V1'
    _globals['_SENDREQUEST']._serialized_start = 85
    _globals['_SENDREQUEST']._serialized_end = 151
    _globals['_SENDRESPONSE']._serialized_start = 153
    _globals['_SENDRESPONSE']._serialized_end = 167
    _globals['_RECEIVEREQUEST']._serialized_start = 169
    _globals['_RECEIVEREQUEST']._serialized_end = 201
    _globals['_RECEIVERESPONSE']._serialized_start = 203
    _globals['_RECEIVERESPONSE']._serialized_end = 273
    _globals['_SOCIALSERVICE']._serialized_start = 276
    _globals['_SOCIALSERVICE']._serialized_end = 434