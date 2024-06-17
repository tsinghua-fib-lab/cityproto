// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/economy/v1/person_service.proto

#include "city/economy/v1/person_service.pb.h"
#include "city/economy/v1/person_service.grpc.pb.h"

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
namespace economy {
namespace v1 {

static const char* PersonService_method_names[] = {
  "/city.economy.v1.PersonService/GetPerson",
  "/city.economy.v1.PersonService/UpdatePersonMoney",
};

std::unique_ptr< PersonService::Stub> PersonService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< PersonService::Stub> stub(new PersonService::Stub(channel, options));
  return stub;
}

PersonService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_GetPerson_(PersonService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_UpdatePersonMoney_(PersonService_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status PersonService::Stub::GetPerson(::grpc::ClientContext* context, const ::city::economy::v1::GetPersonRequest& request, ::city::economy::v1::GetPersonResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::economy::v1::GetPersonRequest, ::city::economy::v1::GetPersonResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetPerson_, context, request, response);
}

void PersonService::Stub::async::GetPerson(::grpc::ClientContext* context, const ::city::economy::v1::GetPersonRequest* request, ::city::economy::v1::GetPersonResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::economy::v1::GetPersonRequest, ::city::economy::v1::GetPersonResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPerson_, context, request, response, std::move(f));
}

void PersonService::Stub::async::GetPerson(::grpc::ClientContext* context, const ::city::economy::v1::GetPersonRequest* request, ::city::economy::v1::GetPersonResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPerson_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::economy::v1::GetPersonResponse>* PersonService::Stub::PrepareAsyncGetPersonRaw(::grpc::ClientContext* context, const ::city::economy::v1::GetPersonRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::economy::v1::GetPersonResponse, ::city::economy::v1::GetPersonRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetPerson_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::economy::v1::GetPersonResponse>* PersonService::Stub::AsyncGetPersonRaw(::grpc::ClientContext* context, const ::city::economy::v1::GetPersonRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetPersonRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::UpdatePersonMoney(::grpc::ClientContext* context, const ::city::economy::v1::UpdatePersonMoneyRequest& request, ::city::economy::v1::UpdatePersonMoneyResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::economy::v1::UpdatePersonMoneyRequest, ::city::economy::v1::UpdatePersonMoneyResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_UpdatePersonMoney_, context, request, response);
}

void PersonService::Stub::async::UpdatePersonMoney(::grpc::ClientContext* context, const ::city::economy::v1::UpdatePersonMoneyRequest* request, ::city::economy::v1::UpdatePersonMoneyResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::economy::v1::UpdatePersonMoneyRequest, ::city::economy::v1::UpdatePersonMoneyResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_UpdatePersonMoney_, context, request, response, std::move(f));
}

void PersonService::Stub::async::UpdatePersonMoney(::grpc::ClientContext* context, const ::city::economy::v1::UpdatePersonMoneyRequest* request, ::city::economy::v1::UpdatePersonMoneyResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_UpdatePersonMoney_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::economy::v1::UpdatePersonMoneyResponse>* PersonService::Stub::PrepareAsyncUpdatePersonMoneyRaw(::grpc::ClientContext* context, const ::city::economy::v1::UpdatePersonMoneyRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::economy::v1::UpdatePersonMoneyResponse, ::city::economy::v1::UpdatePersonMoneyRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_UpdatePersonMoney_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::economy::v1::UpdatePersonMoneyResponse>* PersonService::Stub::AsyncUpdatePersonMoneyRaw(::grpc::ClientContext* context, const ::city::economy::v1::UpdatePersonMoneyRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncUpdatePersonMoneyRaw(context, request, cq);
  result->StartCall();
  return result;
}

PersonService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::economy::v1::GetPersonRequest, ::city::economy::v1::GetPersonResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::economy::v1::GetPersonRequest* req,
             ::city::economy::v1::GetPersonResponse* resp) {
               return service->GetPerson(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::economy::v1::UpdatePersonMoneyRequest, ::city::economy::v1::UpdatePersonMoneyResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::economy::v1::UpdatePersonMoneyRequest* req,
             ::city::economy::v1::UpdatePersonMoneyResponse* resp) {
               return service->UpdatePersonMoney(ctx, req, resp);
             }, this)));
}

PersonService::Service::~Service() {
}

::grpc::Status PersonService::Service::GetPerson(::grpc::ServerContext* context, const ::city::economy::v1::GetPersonRequest* request, ::city::economy::v1::GetPersonResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::UpdatePersonMoney(::grpc::ServerContext* context, const ::city::economy::v1::UpdatePersonMoneyRequest* request, ::city::economy::v1::UpdatePersonMoneyResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace city
}  // namespace economy
}  // namespace v1
