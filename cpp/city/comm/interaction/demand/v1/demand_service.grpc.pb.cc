// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/comm/interaction/demand/v1/demand_service.proto

#include "city/comm/interaction/demand/v1/demand_service.pb.h"
#include "city/comm/interaction/demand/v1/demand_service.grpc.pb.h"

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
namespace demand {
namespace v1 {

static const char* DemandService_method_names[] = {
  "/city.comm.interaction.demand.v1.DemandService/SetDemandStatus",
};

std::unique_ptr< DemandService::Stub> DemandService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< DemandService::Stub> stub(new DemandService::Stub(channel, options));
  return stub;
}

DemandService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_SetDemandStatus_(DemandService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status DemandService::Stub::SetDemandStatus(::grpc::ClientContext* context, const ::city::comm::interaction::demand::v1::SetDemandStatusRequest& request, ::city::comm::interaction::demand::v1::SetDemandStatusResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::comm::interaction::demand::v1::SetDemandStatusRequest, ::city::comm::interaction::demand::v1::SetDemandStatusResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SetDemandStatus_, context, request, response);
}

void DemandService::Stub::async::SetDemandStatus(::grpc::ClientContext* context, const ::city::comm::interaction::demand::v1::SetDemandStatusRequest* request, ::city::comm::interaction::demand::v1::SetDemandStatusResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::comm::interaction::demand::v1::SetDemandStatusRequest, ::city::comm::interaction::demand::v1::SetDemandStatusResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetDemandStatus_, context, request, response, std::move(f));
}

void DemandService::Stub::async::SetDemandStatus(::grpc::ClientContext* context, const ::city::comm::interaction::demand::v1::SetDemandStatusRequest* request, ::city::comm::interaction::demand::v1::SetDemandStatusResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetDemandStatus_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::comm::interaction::demand::v1::SetDemandStatusResponse>* DemandService::Stub::PrepareAsyncSetDemandStatusRaw(::grpc::ClientContext* context, const ::city::comm::interaction::demand::v1::SetDemandStatusRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::comm::interaction::demand::v1::SetDemandStatusResponse, ::city::comm::interaction::demand::v1::SetDemandStatusRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SetDemandStatus_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::comm::interaction::demand::v1::SetDemandStatusResponse>* DemandService::Stub::AsyncSetDemandStatusRaw(::grpc::ClientContext* context, const ::city::comm::interaction::demand::v1::SetDemandStatusRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSetDemandStatusRaw(context, request, cq);
  result->StartCall();
  return result;
}

DemandService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DemandService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DemandService::Service, ::city::comm::interaction::demand::v1::SetDemandStatusRequest, ::city::comm::interaction::demand::v1::SetDemandStatusResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](DemandService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::comm::interaction::demand::v1::SetDemandStatusRequest* req,
             ::city::comm::interaction::demand::v1::SetDemandStatusResponse* resp) {
               return service->SetDemandStatus(ctx, req, resp);
             }, this)));
}

DemandService::Service::~Service() {
}

::grpc::Status DemandService::Service::SetDemandStatus(::grpc::ServerContext* context, const ::city::comm::interaction::demand::v1::SetDemandStatusRequest* request, ::city::comm::interaction::demand::v1::SetDemandStatusResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace city
}  // namespace comm
}  // namespace interaction
}  // namespace demand
}  // namespace v1
