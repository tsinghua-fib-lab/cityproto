// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/map/v2/lane_service.proto

#include "city/map/v2/lane_service.pb.h"
#include "city/map/v2/lane_service.grpc.pb.h"

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
namespace map {
namespace v2 {

static const char* LaneService_method_names[] = {
  "/city.map.v2.LaneService/SetLaneMaxV",
  "/city.map.v2.LaneService/GetLane",
  "/city.map.v2.LaneService/GetLaneByLongLatBBox",
};

std::unique_ptr< LaneService::Stub> LaneService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< LaneService::Stub> stub(new LaneService::Stub(channel, options));
  return stub;
}

LaneService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_SetLaneMaxV_(LaneService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetLane_(LaneService_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetLaneByLongLatBBox_(LaneService_method_names[2], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status LaneService::Stub::SetLaneMaxV(::grpc::ClientContext* context, const ::city::map::v2::SetLaneMaxVRequest& request, ::city::map::v2::SetLaneMaxVResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::map::v2::SetLaneMaxVRequest, ::city::map::v2::SetLaneMaxVResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SetLaneMaxV_, context, request, response);
}

void LaneService::Stub::async::SetLaneMaxV(::grpc::ClientContext* context, const ::city::map::v2::SetLaneMaxVRequest* request, ::city::map::v2::SetLaneMaxVResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::map::v2::SetLaneMaxVRequest, ::city::map::v2::SetLaneMaxVResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetLaneMaxV_, context, request, response, std::move(f));
}

void LaneService::Stub::async::SetLaneMaxV(::grpc::ClientContext* context, const ::city::map::v2::SetLaneMaxVRequest* request, ::city::map::v2::SetLaneMaxVResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetLaneMaxV_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::map::v2::SetLaneMaxVResponse>* LaneService::Stub::PrepareAsyncSetLaneMaxVRaw(::grpc::ClientContext* context, const ::city::map::v2::SetLaneMaxVRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::map::v2::SetLaneMaxVResponse, ::city::map::v2::SetLaneMaxVRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SetLaneMaxV_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::map::v2::SetLaneMaxVResponse>* LaneService::Stub::AsyncSetLaneMaxVRaw(::grpc::ClientContext* context, const ::city::map::v2::SetLaneMaxVRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSetLaneMaxVRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status LaneService::Stub::GetLane(::grpc::ClientContext* context, const ::city::map::v2::GetLaneRequest& request, ::city::map::v2::GetLaneResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::map::v2::GetLaneRequest, ::city::map::v2::GetLaneResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetLane_, context, request, response);
}

void LaneService::Stub::async::GetLane(::grpc::ClientContext* context, const ::city::map::v2::GetLaneRequest* request, ::city::map::v2::GetLaneResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::map::v2::GetLaneRequest, ::city::map::v2::GetLaneResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetLane_, context, request, response, std::move(f));
}

void LaneService::Stub::async::GetLane(::grpc::ClientContext* context, const ::city::map::v2::GetLaneRequest* request, ::city::map::v2::GetLaneResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetLane_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::map::v2::GetLaneResponse>* LaneService::Stub::PrepareAsyncGetLaneRaw(::grpc::ClientContext* context, const ::city::map::v2::GetLaneRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::map::v2::GetLaneResponse, ::city::map::v2::GetLaneRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetLane_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::map::v2::GetLaneResponse>* LaneService::Stub::AsyncGetLaneRaw(::grpc::ClientContext* context, const ::city::map::v2::GetLaneRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetLaneRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status LaneService::Stub::GetLaneByLongLatBBox(::grpc::ClientContext* context, const ::city::map::v2::GetLaneByLongLatBBoxRequest& request, ::city::map::v2::GetLaneByLongLatBBoxResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::map::v2::GetLaneByLongLatBBoxRequest, ::city::map::v2::GetLaneByLongLatBBoxResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetLaneByLongLatBBox_, context, request, response);
}

void LaneService::Stub::async::GetLaneByLongLatBBox(::grpc::ClientContext* context, const ::city::map::v2::GetLaneByLongLatBBoxRequest* request, ::city::map::v2::GetLaneByLongLatBBoxResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::map::v2::GetLaneByLongLatBBoxRequest, ::city::map::v2::GetLaneByLongLatBBoxResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetLaneByLongLatBBox_, context, request, response, std::move(f));
}

void LaneService::Stub::async::GetLaneByLongLatBBox(::grpc::ClientContext* context, const ::city::map::v2::GetLaneByLongLatBBoxRequest* request, ::city::map::v2::GetLaneByLongLatBBoxResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetLaneByLongLatBBox_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::map::v2::GetLaneByLongLatBBoxResponse>* LaneService::Stub::PrepareAsyncGetLaneByLongLatBBoxRaw(::grpc::ClientContext* context, const ::city::map::v2::GetLaneByLongLatBBoxRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::map::v2::GetLaneByLongLatBBoxResponse, ::city::map::v2::GetLaneByLongLatBBoxRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetLaneByLongLatBBox_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::map::v2::GetLaneByLongLatBBoxResponse>* LaneService::Stub::AsyncGetLaneByLongLatBBoxRaw(::grpc::ClientContext* context, const ::city::map::v2::GetLaneByLongLatBBoxRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetLaneByLongLatBBoxRaw(context, request, cq);
  result->StartCall();
  return result;
}

LaneService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      LaneService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< LaneService::Service, ::city::map::v2::SetLaneMaxVRequest, ::city::map::v2::SetLaneMaxVResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](LaneService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::map::v2::SetLaneMaxVRequest* req,
             ::city::map::v2::SetLaneMaxVResponse* resp) {
               return service->SetLaneMaxV(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      LaneService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< LaneService::Service, ::city::map::v2::GetLaneRequest, ::city::map::v2::GetLaneResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](LaneService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::map::v2::GetLaneRequest* req,
             ::city::map::v2::GetLaneResponse* resp) {
               return service->GetLane(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      LaneService_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< LaneService::Service, ::city::map::v2::GetLaneByLongLatBBoxRequest, ::city::map::v2::GetLaneByLongLatBBoxResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](LaneService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::map::v2::GetLaneByLongLatBBoxRequest* req,
             ::city::map::v2::GetLaneByLongLatBBoxResponse* resp) {
               return service->GetLaneByLongLatBBox(ctx, req, resp);
             }, this)));
}

LaneService::Service::~Service() {
}

::grpc::Status LaneService::Service::SetLaneMaxV(::grpc::ServerContext* context, const ::city::map::v2::SetLaneMaxVRequest* request, ::city::map::v2::SetLaneMaxVResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status LaneService::Service::GetLane(::grpc::ServerContext* context, const ::city::map::v2::GetLaneRequest* request, ::city::map::v2::GetLaneResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status LaneService::Service::GetLaneByLongLatBBox(::grpc::ServerContext* context, const ::city::map::v2::GetLaneByLongLatBBoxRequest* request, ::city::map::v2::GetLaneByLongLatBBoxResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace city
}  // namespace map
}  // namespace v2

