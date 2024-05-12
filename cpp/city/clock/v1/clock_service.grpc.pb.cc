// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/clock/v1/clock_service.proto

#include "city/clock/v1/clock_service.pb.h"
#include "city/clock/v1/clock_service.grpc.pb.h"

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
namespace clock {
namespace v1 {

static const char* ClockService_method_names[] = {
  "/city.clock.v1.ClockService/Now",
};

std::unique_ptr< ClockService::Stub> ClockService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< ClockService::Stub> stub(new ClockService::Stub(channel, options));
  return stub;
}

ClockService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_Now_(ClockService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status ClockService::Stub::Now(::grpc::ClientContext* context, const ::city::clock::v1::NowRequest& request, ::city::clock::v1::NowResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::clock::v1::NowRequest, ::city::clock::v1::NowResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Now_, context, request, response);
}

void ClockService::Stub::async::Now(::grpc::ClientContext* context, const ::city::clock::v1::NowRequest* request, ::city::clock::v1::NowResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::clock::v1::NowRequest, ::city::clock::v1::NowResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Now_, context, request, response, std::move(f));
}

void ClockService::Stub::async::Now(::grpc::ClientContext* context, const ::city::clock::v1::NowRequest* request, ::city::clock::v1::NowResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Now_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::clock::v1::NowResponse>* ClockService::Stub::PrepareAsyncNowRaw(::grpc::ClientContext* context, const ::city::clock::v1::NowRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::clock::v1::NowResponse, ::city::clock::v1::NowRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Now_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::clock::v1::NowResponse>* ClockService::Stub::AsyncNowRaw(::grpc::ClientContext* context, const ::city::clock::v1::NowRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncNowRaw(context, request, cq);
  result->StartCall();
  return result;
}

ClockService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      ClockService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< ClockService::Service, ::city::clock::v1::NowRequest, ::city::clock::v1::NowResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](ClockService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::clock::v1::NowRequest* req,
             ::city::clock::v1::NowResponse* resp) {
               return service->Now(ctx, req, resp);
             }, this)));
}

ClockService::Service::~Service() {
}

::grpc::Status ClockService::Service::Now(::grpc::ServerContext* context, const ::city::clock::v1::NowRequest* request, ::city::clock::v1::NowResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace city
}  // namespace clock
}  // namespace v1

