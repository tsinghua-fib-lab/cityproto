// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/comm/interaction/aoi/v1/aoi_service.proto

#include "city/comm/interaction/aoi/v1/aoi_service.pb.h"
#include "city/comm/interaction/aoi/v1/aoi_service.grpc.pb.h"

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
namespace comm {
namespace interaction {
namespace aoi {
namespace v1 {

static const char* AoiService_method_names[] = {
  "/city.comm.interaction.aoi.v1.AoiService/GetBadAoiID",
};

std::unique_ptr< AoiService::Stub> AoiService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< AoiService::Stub> stub(new AoiService::Stub(channel, options));
  return stub;
}

AoiService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_GetBadAoiID_(AoiService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status AoiService::Stub::GetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetBadAoiID_, context, request, response);
}

void AoiService::Stub::async::GetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetBadAoiID_, context, request, response, std::move(f));
}

void AoiService::Stub::async::GetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetBadAoiID_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* AoiService::Stub::PrepareAsyncGetBadAoiIDRaw(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse, ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetBadAoiID_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* AoiService::Stub::AsyncGetBadAoiIDRaw(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetBadAoiIDRaw(context, request, cq);
  result->StartCall();
  return result;
}

AoiService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      AoiService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< AoiService::Service, ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](AoiService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* req,
             ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* resp) {
               return service->GetBadAoiID(ctx, req, resp);
             }, this)));
}

AoiService::Service::~Service() {
}

::grpc::Status AoiService::Service::GetBadAoiID(::grpc::ServerContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace city
}  // namespace comm
}  // namespace interaction
}  // namespace aoi
}  // namespace v1

