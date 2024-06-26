// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/water/input/v1/input_service.proto

#include "city/water/input/v1/input_service.pb.h"
#include "city/water/input/v1/input_service.grpc.pb.h"

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
namespace water {
namespace input {
namespace v1 {

static const char* InputService_method_names[] = {
  "/city.water.input.v1.InputService/Init",
};

std::unique_ptr< InputService::Stub> InputService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< InputService::Stub> stub(new InputService::Stub(channel, options));
  return stub;
}

InputService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_Init_(InputService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status InputService::Stub::Init(::grpc::ClientContext* context, const ::city::water::input::v1::InitRequest& request, ::city::water::input::v1::InitResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::water::input::v1::InitRequest, ::city::water::input::v1::InitResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Init_, context, request, response);
}

void InputService::Stub::async::Init(::grpc::ClientContext* context, const ::city::water::input::v1::InitRequest* request, ::city::water::input::v1::InitResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::water::input::v1::InitRequest, ::city::water::input::v1::InitResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Init_, context, request, response, std::move(f));
}

void InputService::Stub::async::Init(::grpc::ClientContext* context, const ::city::water::input::v1::InitRequest* request, ::city::water::input::v1::InitResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Init_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::water::input::v1::InitResponse>* InputService::Stub::PrepareAsyncInitRaw(::grpc::ClientContext* context, const ::city::water::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::water::input::v1::InitResponse, ::city::water::input::v1::InitRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Init_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::water::input::v1::InitResponse>* InputService::Stub::AsyncInitRaw(::grpc::ClientContext* context, const ::city::water::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncInitRaw(context, request, cq);
  result->StartCall();
  return result;
}

InputService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      InputService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< InputService::Service, ::city::water::input::v1::InitRequest, ::city::water::input::v1::InitResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](InputService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::water::input::v1::InitRequest* req,
             ::city::water::input::v1::InitResponse* resp) {
               return service->Init(ctx, req, resp);
             }, this)));
}

InputService::Service::~Service() {
}

::grpc::Status InputService::Service::Init(::grpc::ServerContext* context, const ::city::water::input::v1::InitRequest* request, ::city::water::input::v1::InitResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace city
}  // namespace water
}  // namespace input
}  // namespace v1

