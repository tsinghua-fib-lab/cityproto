// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/social/v1/social_service.proto

#include "city/social/v1/social_service.pb.h"
#include "city/social/v1/social_service.grpc.pb.h"

#include <functional>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/impl/channel_interface.h>
#include <grpcpp/impl/client_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/rpc_service_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/sync_stream.h>
namespace city {
namespace social {
namespace v1 {

static const char* SocialService_method_names[] = {
  "/city.social.v1.SocialService/Send",
  "/city.social.v1.SocialService/Receive",
};

std::unique_ptr< SocialService::Stub> SocialService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< SocialService::Stub> stub(new SocialService::Stub(channel, options));
  return stub;
}

SocialService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_Send_(SocialService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Receive_(SocialService_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status SocialService::Stub::Send(::grpc::ClientContext* context, const ::city::social::v1::SendRequest& request, ::city::social::v1::SendResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::social::v1::SendRequest, ::city::social::v1::SendResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Send_, context, request, response);
}

void SocialService::Stub::async::Send(::grpc::ClientContext* context, const ::city::social::v1::SendRequest* request, ::city::social::v1::SendResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::social::v1::SendRequest, ::city::social::v1::SendResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Send_, context, request, response, std::move(f));
}

void SocialService::Stub::async::Send(::grpc::ClientContext* context, const ::city::social::v1::SendRequest* request, ::city::social::v1::SendResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Send_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::social::v1::SendResponse>* SocialService::Stub::PrepareAsyncSendRaw(::grpc::ClientContext* context, const ::city::social::v1::SendRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::social::v1::SendResponse, ::city::social::v1::SendRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Send_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::social::v1::SendResponse>* SocialService::Stub::AsyncSendRaw(::grpc::ClientContext* context, const ::city::social::v1::SendRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSendRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status SocialService::Stub::Receive(::grpc::ClientContext* context, const ::city::social::v1::ReceiveRequest& request, ::city::social::v1::ReceiveResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::social::v1::ReceiveRequest, ::city::social::v1::ReceiveResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Receive_, context, request, response);
}

void SocialService::Stub::async::Receive(::grpc::ClientContext* context, const ::city::social::v1::ReceiveRequest* request, ::city::social::v1::ReceiveResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::social::v1::ReceiveRequest, ::city::social::v1::ReceiveResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Receive_, context, request, response, std::move(f));
}

void SocialService::Stub::async::Receive(::grpc::ClientContext* context, const ::city::social::v1::ReceiveRequest* request, ::city::social::v1::ReceiveResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Receive_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::social::v1::ReceiveResponse>* SocialService::Stub::PrepareAsyncReceiveRaw(::grpc::ClientContext* context, const ::city::social::v1::ReceiveRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::social::v1::ReceiveResponse, ::city::social::v1::ReceiveRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Receive_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::social::v1::ReceiveResponse>* SocialService::Stub::AsyncReceiveRaw(::grpc::ClientContext* context, const ::city::social::v1::ReceiveRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncReceiveRaw(context, request, cq);
  result->StartCall();
  return result;
}

SocialService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SocialService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SocialService::Service, ::city::social::v1::SendRequest, ::city::social::v1::SendResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SocialService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::social::v1::SendRequest* req,
             ::city::social::v1::SendResponse* resp) {
               return service->Send(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SocialService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SocialService::Service, ::city::social::v1::ReceiveRequest, ::city::social::v1::ReceiveResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SocialService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::social::v1::ReceiveRequest* req,
             ::city::social::v1::ReceiveResponse* resp) {
               return service->Receive(ctx, req, resp);
             }, this)));
}

SocialService::Service::~Service() {
}

::grpc::Status SocialService::Service::Send(::grpc::ServerContext* context, const ::city::social::v1::SendRequest* request, ::city::social::v1::SendResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SocialService::Service::Receive(::grpc::ServerContext* context, const ::city::social::v1::ReceiveRequest* request, ::city::social::v1::ReceiveResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace city
}  // namespace social
}  // namespace v1
