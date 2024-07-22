// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/sync/v2/sync_service.proto

#include "city/sync/v2/sync_service.pb.h"
#include "city/sync/v2/sync_service.grpc.pb.h"

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
namespace sync {
namespace v2 {

static const char* SyncService_method_names[] = {
  "/city.sync.v2.SyncService/SetURL",
  "/city.sync.v2.SyncService/GetURL",
  "/city.sync.v2.SyncService/EnterStepSync",
  "/city.sync.v2.SyncService/ExitStepSync",
};

std::unique_ptr< SyncService::Stub> SyncService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< SyncService::Stub> stub(new SyncService::Stub(channel, options));
  return stub;
}

SyncService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_SetURL_(SyncService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetURL_(SyncService_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_EnterStepSync_(SyncService_method_names[2], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_ExitStepSync_(SyncService_method_names[3], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status SyncService::Stub::SetURL(::grpc::ClientContext* context, const ::city::sync::v2::SetURLRequest& request, ::city::sync::v2::SetURLResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::sync::v2::SetURLRequest, ::city::sync::v2::SetURLResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SetURL_, context, request, response);
}

void SyncService::Stub::async::SetURL(::grpc::ClientContext* context, const ::city::sync::v2::SetURLRequest* request, ::city::sync::v2::SetURLResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::sync::v2::SetURLRequest, ::city::sync::v2::SetURLResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetURL_, context, request, response, std::move(f));
}

void SyncService::Stub::async::SetURL(::grpc::ClientContext* context, const ::city::sync::v2::SetURLRequest* request, ::city::sync::v2::SetURLResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetURL_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::sync::v2::SetURLResponse>* SyncService::Stub::PrepareAsyncSetURLRaw(::grpc::ClientContext* context, const ::city::sync::v2::SetURLRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::sync::v2::SetURLResponse, ::city::sync::v2::SetURLRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SetURL_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::sync::v2::SetURLResponse>* SyncService::Stub::AsyncSetURLRaw(::grpc::ClientContext* context, const ::city::sync::v2::SetURLRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSetURLRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status SyncService::Stub::GetURL(::grpc::ClientContext* context, const ::city::sync::v2::GetURLRequest& request, ::city::sync::v2::GetURLResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::sync::v2::GetURLRequest, ::city::sync::v2::GetURLResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetURL_, context, request, response);
}

void SyncService::Stub::async::GetURL(::grpc::ClientContext* context, const ::city::sync::v2::GetURLRequest* request, ::city::sync::v2::GetURLResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::sync::v2::GetURLRequest, ::city::sync::v2::GetURLResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetURL_, context, request, response, std::move(f));
}

void SyncService::Stub::async::GetURL(::grpc::ClientContext* context, const ::city::sync::v2::GetURLRequest* request, ::city::sync::v2::GetURLResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetURL_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::sync::v2::GetURLResponse>* SyncService::Stub::PrepareAsyncGetURLRaw(::grpc::ClientContext* context, const ::city::sync::v2::GetURLRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::sync::v2::GetURLResponse, ::city::sync::v2::GetURLRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetURL_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::sync::v2::GetURLResponse>* SyncService::Stub::AsyncGetURLRaw(::grpc::ClientContext* context, const ::city::sync::v2::GetURLRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetURLRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status SyncService::Stub::EnterStepSync(::grpc::ClientContext* context, const ::city::sync::v2::EnterStepSyncRequest& request, ::city::sync::v2::EnterStepSyncResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::sync::v2::EnterStepSyncRequest, ::city::sync::v2::EnterStepSyncResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_EnterStepSync_, context, request, response);
}

void SyncService::Stub::async::EnterStepSync(::grpc::ClientContext* context, const ::city::sync::v2::EnterStepSyncRequest* request, ::city::sync::v2::EnterStepSyncResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::sync::v2::EnterStepSyncRequest, ::city::sync::v2::EnterStepSyncResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_EnterStepSync_, context, request, response, std::move(f));
}

void SyncService::Stub::async::EnterStepSync(::grpc::ClientContext* context, const ::city::sync::v2::EnterStepSyncRequest* request, ::city::sync::v2::EnterStepSyncResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_EnterStepSync_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::sync::v2::EnterStepSyncResponse>* SyncService::Stub::PrepareAsyncEnterStepSyncRaw(::grpc::ClientContext* context, const ::city::sync::v2::EnterStepSyncRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::sync::v2::EnterStepSyncResponse, ::city::sync::v2::EnterStepSyncRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_EnterStepSync_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::sync::v2::EnterStepSyncResponse>* SyncService::Stub::AsyncEnterStepSyncRaw(::grpc::ClientContext* context, const ::city::sync::v2::EnterStepSyncRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncEnterStepSyncRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status SyncService::Stub::ExitStepSync(::grpc::ClientContext* context, const ::city::sync::v2::ExitStepSyncRequest& request, ::city::sync::v2::ExitStepSyncResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::sync::v2::ExitStepSyncRequest, ::city::sync::v2::ExitStepSyncResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_ExitStepSync_, context, request, response);
}

void SyncService::Stub::async::ExitStepSync(::grpc::ClientContext* context, const ::city::sync::v2::ExitStepSyncRequest* request, ::city::sync::v2::ExitStepSyncResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::sync::v2::ExitStepSyncRequest, ::city::sync::v2::ExitStepSyncResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_ExitStepSync_, context, request, response, std::move(f));
}

void SyncService::Stub::async::ExitStepSync(::grpc::ClientContext* context, const ::city::sync::v2::ExitStepSyncRequest* request, ::city::sync::v2::ExitStepSyncResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_ExitStepSync_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::sync::v2::ExitStepSyncResponse>* SyncService::Stub::PrepareAsyncExitStepSyncRaw(::grpc::ClientContext* context, const ::city::sync::v2::ExitStepSyncRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::sync::v2::ExitStepSyncResponse, ::city::sync::v2::ExitStepSyncRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_ExitStepSync_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::sync::v2::ExitStepSyncResponse>* SyncService::Stub::AsyncExitStepSyncRaw(::grpc::ClientContext* context, const ::city::sync::v2::ExitStepSyncRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncExitStepSyncRaw(context, request, cq);
  result->StartCall();
  return result;
}

SyncService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SyncService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SyncService::Service, ::city::sync::v2::SetURLRequest, ::city::sync::v2::SetURLResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SyncService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::sync::v2::SetURLRequest* req,
             ::city::sync::v2::SetURLResponse* resp) {
               return service->SetURL(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SyncService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SyncService::Service, ::city::sync::v2::GetURLRequest, ::city::sync::v2::GetURLResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SyncService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::sync::v2::GetURLRequest* req,
             ::city::sync::v2::GetURLResponse* resp) {
               return service->GetURL(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SyncService_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SyncService::Service, ::city::sync::v2::EnterStepSyncRequest, ::city::sync::v2::EnterStepSyncResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SyncService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::sync::v2::EnterStepSyncRequest* req,
             ::city::sync::v2::EnterStepSyncResponse* resp) {
               return service->EnterStepSync(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SyncService_method_names[3],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SyncService::Service, ::city::sync::v2::ExitStepSyncRequest, ::city::sync::v2::ExitStepSyncResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SyncService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::sync::v2::ExitStepSyncRequest* req,
             ::city::sync::v2::ExitStepSyncResponse* resp) {
               return service->ExitStepSync(ctx, req, resp);
             }, this)));
}

SyncService::Service::~Service() {
}

::grpc::Status SyncService::Service::SetURL(::grpc::ServerContext* context, const ::city::sync::v2::SetURLRequest* request, ::city::sync::v2::SetURLResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SyncService::Service::GetURL(::grpc::ServerContext* context, const ::city::sync::v2::GetURLRequest* request, ::city::sync::v2::GetURLResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SyncService::Service::EnterStepSync(::grpc::ServerContext* context, const ::city::sync::v2::EnterStepSyncRequest* request, ::city::sync::v2::EnterStepSyncResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SyncService::Service::ExitStepSync(::grpc::ServerContext* context, const ::city::sync::v2::ExitStepSyncRequest* request, ::city::sync::v2::ExitStepSyncResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace city
}  // namespace sync
}  // namespace v2
