// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/person/v2/person_service.proto

#include "city/person/v2/person_service.pb.h"
#include "city/person/v2/person_service.grpc.pb.h"

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
namespace person {
namespace v2 {

static const char* PersonService_method_names[] = {
  "/city.person.v2.PersonService/GetPerson",
  "/city.person.v2.PersonService/AddPerson",
  "/city.person.v2.PersonService/SetSchedule",
  "/city.person.v2.PersonService/GetPersons",
  "/city.person.v2.PersonService/GetPersonByLongLatBBox",
  "/city.person.v2.PersonService/GetAllVehicles",
  "/city.person.v2.PersonService/ResetPersonPosition",
  "/city.person.v2.PersonService/SetControlledVehicleIDs",
  "/city.person.v2.PersonService/FetchControlledVehicleEnvs",
  "/city.person.v2.PersonService/SetControlledVehicleActions",
};

std::unique_ptr< PersonService::Stub> PersonService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< PersonService::Stub> stub(new PersonService::Stub(channel, options));
  return stub;
}

PersonService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_GetPerson_(PersonService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_AddPerson_(PersonService_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SetSchedule_(PersonService_method_names[2], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetPersons_(PersonService_method_names[3], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetPersonByLongLatBBox_(PersonService_method_names[4], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetAllVehicles_(PersonService_method_names[5], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_ResetPersonPosition_(PersonService_method_names[6], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SetControlledVehicleIDs_(PersonService_method_names[7], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_FetchControlledVehicleEnvs_(PersonService_method_names[8], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SetControlledVehicleActions_(PersonService_method_names[9], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status PersonService::Stub::GetPerson(::grpc::ClientContext* context, const ::city::person::v2::GetPersonRequest& request, ::city::person::v2::GetPersonResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::GetPersonRequest, ::city::person::v2::GetPersonResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetPerson_, context, request, response);
}

void PersonService::Stub::async::GetPerson(::grpc::ClientContext* context, const ::city::person::v2::GetPersonRequest* request, ::city::person::v2::GetPersonResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::GetPersonRequest, ::city::person::v2::GetPersonResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPerson_, context, request, response, std::move(f));
}

void PersonService::Stub::async::GetPerson(::grpc::ClientContext* context, const ::city::person::v2::GetPersonRequest* request, ::city::person::v2::GetPersonResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPerson_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::GetPersonResponse>* PersonService::Stub::PrepareAsyncGetPersonRaw(::grpc::ClientContext* context, const ::city::person::v2::GetPersonRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::GetPersonResponse, ::city::person::v2::GetPersonRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetPerson_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::GetPersonResponse>* PersonService::Stub::AsyncGetPersonRaw(::grpc::ClientContext* context, const ::city::person::v2::GetPersonRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetPersonRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::AddPerson(::grpc::ClientContext* context, const ::city::person::v2::AddPersonRequest& request, ::city::person::v2::AddPersonResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::AddPersonRequest, ::city::person::v2::AddPersonResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_AddPerson_, context, request, response);
}

void PersonService::Stub::async::AddPerson(::grpc::ClientContext* context, const ::city::person::v2::AddPersonRequest* request, ::city::person::v2::AddPersonResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::AddPersonRequest, ::city::person::v2::AddPersonResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_AddPerson_, context, request, response, std::move(f));
}

void PersonService::Stub::async::AddPerson(::grpc::ClientContext* context, const ::city::person::v2::AddPersonRequest* request, ::city::person::v2::AddPersonResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_AddPerson_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::AddPersonResponse>* PersonService::Stub::PrepareAsyncAddPersonRaw(::grpc::ClientContext* context, const ::city::person::v2::AddPersonRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::AddPersonResponse, ::city::person::v2::AddPersonRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_AddPerson_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::AddPersonResponse>* PersonService::Stub::AsyncAddPersonRaw(::grpc::ClientContext* context, const ::city::person::v2::AddPersonRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncAddPersonRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::SetSchedule(::grpc::ClientContext* context, const ::city::person::v2::SetScheduleRequest& request, ::city::person::v2::SetScheduleResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::SetScheduleRequest, ::city::person::v2::SetScheduleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SetSchedule_, context, request, response);
}

void PersonService::Stub::async::SetSchedule(::grpc::ClientContext* context, const ::city::person::v2::SetScheduleRequest* request, ::city::person::v2::SetScheduleResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::SetScheduleRequest, ::city::person::v2::SetScheduleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetSchedule_, context, request, response, std::move(f));
}

void PersonService::Stub::async::SetSchedule(::grpc::ClientContext* context, const ::city::person::v2::SetScheduleRequest* request, ::city::person::v2::SetScheduleResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetSchedule_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::SetScheduleResponse>* PersonService::Stub::PrepareAsyncSetScheduleRaw(::grpc::ClientContext* context, const ::city::person::v2::SetScheduleRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::SetScheduleResponse, ::city::person::v2::SetScheduleRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SetSchedule_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::SetScheduleResponse>* PersonService::Stub::AsyncSetScheduleRaw(::grpc::ClientContext* context, const ::city::person::v2::SetScheduleRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSetScheduleRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::GetPersons(::grpc::ClientContext* context, const ::city::person::v2::GetPersonsRequest& request, ::city::person::v2::GetPersonsResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::GetPersonsRequest, ::city::person::v2::GetPersonsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetPersons_, context, request, response);
}

void PersonService::Stub::async::GetPersons(::grpc::ClientContext* context, const ::city::person::v2::GetPersonsRequest* request, ::city::person::v2::GetPersonsResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::GetPersonsRequest, ::city::person::v2::GetPersonsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPersons_, context, request, response, std::move(f));
}

void PersonService::Stub::async::GetPersons(::grpc::ClientContext* context, const ::city::person::v2::GetPersonsRequest* request, ::city::person::v2::GetPersonsResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPersons_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::GetPersonsResponse>* PersonService::Stub::PrepareAsyncGetPersonsRaw(::grpc::ClientContext* context, const ::city::person::v2::GetPersonsRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::GetPersonsResponse, ::city::person::v2::GetPersonsRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetPersons_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::GetPersonsResponse>* PersonService::Stub::AsyncGetPersonsRaw(::grpc::ClientContext* context, const ::city::person::v2::GetPersonsRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetPersonsRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::GetPersonByLongLatBBox(::grpc::ClientContext* context, const ::city::person::v2::GetPersonByLongLatBBoxRequest& request, ::city::person::v2::GetPersonByLongLatBBoxResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::GetPersonByLongLatBBoxRequest, ::city::person::v2::GetPersonByLongLatBBoxResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetPersonByLongLatBBox_, context, request, response);
}

void PersonService::Stub::async::GetPersonByLongLatBBox(::grpc::ClientContext* context, const ::city::person::v2::GetPersonByLongLatBBoxRequest* request, ::city::person::v2::GetPersonByLongLatBBoxResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::GetPersonByLongLatBBoxRequest, ::city::person::v2::GetPersonByLongLatBBoxResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPersonByLongLatBBox_, context, request, response, std::move(f));
}

void PersonService::Stub::async::GetPersonByLongLatBBox(::grpc::ClientContext* context, const ::city::person::v2::GetPersonByLongLatBBoxRequest* request, ::city::person::v2::GetPersonByLongLatBBoxResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPersonByLongLatBBox_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::GetPersonByLongLatBBoxResponse>* PersonService::Stub::PrepareAsyncGetPersonByLongLatBBoxRaw(::grpc::ClientContext* context, const ::city::person::v2::GetPersonByLongLatBBoxRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::GetPersonByLongLatBBoxResponse, ::city::person::v2::GetPersonByLongLatBBoxRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetPersonByLongLatBBox_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::GetPersonByLongLatBBoxResponse>* PersonService::Stub::AsyncGetPersonByLongLatBBoxRaw(::grpc::ClientContext* context, const ::city::person::v2::GetPersonByLongLatBBoxRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetPersonByLongLatBBoxRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::GetAllVehicles(::grpc::ClientContext* context, const ::city::person::v2::GetAllVehiclesRequest& request, ::city::person::v2::GetAllVehiclesResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::GetAllVehiclesRequest, ::city::person::v2::GetAllVehiclesResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetAllVehicles_, context, request, response);
}

void PersonService::Stub::async::GetAllVehicles(::grpc::ClientContext* context, const ::city::person::v2::GetAllVehiclesRequest* request, ::city::person::v2::GetAllVehiclesResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::GetAllVehiclesRequest, ::city::person::v2::GetAllVehiclesResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetAllVehicles_, context, request, response, std::move(f));
}

void PersonService::Stub::async::GetAllVehicles(::grpc::ClientContext* context, const ::city::person::v2::GetAllVehiclesRequest* request, ::city::person::v2::GetAllVehiclesResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetAllVehicles_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::GetAllVehiclesResponse>* PersonService::Stub::PrepareAsyncGetAllVehiclesRaw(::grpc::ClientContext* context, const ::city::person::v2::GetAllVehiclesRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::GetAllVehiclesResponse, ::city::person::v2::GetAllVehiclesRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetAllVehicles_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::GetAllVehiclesResponse>* PersonService::Stub::AsyncGetAllVehiclesRaw(::grpc::ClientContext* context, const ::city::person::v2::GetAllVehiclesRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetAllVehiclesRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::ResetPersonPosition(::grpc::ClientContext* context, const ::city::person::v2::ResetPersonPositionRequest& request, ::city::person::v2::ResetPersonPositionResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::ResetPersonPositionRequest, ::city::person::v2::ResetPersonPositionResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_ResetPersonPosition_, context, request, response);
}

void PersonService::Stub::async::ResetPersonPosition(::grpc::ClientContext* context, const ::city::person::v2::ResetPersonPositionRequest* request, ::city::person::v2::ResetPersonPositionResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::ResetPersonPositionRequest, ::city::person::v2::ResetPersonPositionResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_ResetPersonPosition_, context, request, response, std::move(f));
}

void PersonService::Stub::async::ResetPersonPosition(::grpc::ClientContext* context, const ::city::person::v2::ResetPersonPositionRequest* request, ::city::person::v2::ResetPersonPositionResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_ResetPersonPosition_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::ResetPersonPositionResponse>* PersonService::Stub::PrepareAsyncResetPersonPositionRaw(::grpc::ClientContext* context, const ::city::person::v2::ResetPersonPositionRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::ResetPersonPositionResponse, ::city::person::v2::ResetPersonPositionRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_ResetPersonPosition_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::ResetPersonPositionResponse>* PersonService::Stub::AsyncResetPersonPositionRaw(::grpc::ClientContext* context, const ::city::person::v2::ResetPersonPositionRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncResetPersonPositionRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::SetControlledVehicleIDs(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleIDsRequest& request, ::city::person::v2::SetControlledVehicleIDsResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::SetControlledVehicleIDsRequest, ::city::person::v2::SetControlledVehicleIDsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SetControlledVehicleIDs_, context, request, response);
}

void PersonService::Stub::async::SetControlledVehicleIDs(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleIDsRequest* request, ::city::person::v2::SetControlledVehicleIDsResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::SetControlledVehicleIDsRequest, ::city::person::v2::SetControlledVehicleIDsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetControlledVehicleIDs_, context, request, response, std::move(f));
}

void PersonService::Stub::async::SetControlledVehicleIDs(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleIDsRequest* request, ::city::person::v2::SetControlledVehicleIDsResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetControlledVehicleIDs_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::SetControlledVehicleIDsResponse>* PersonService::Stub::PrepareAsyncSetControlledVehicleIDsRaw(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleIDsRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::SetControlledVehicleIDsResponse, ::city::person::v2::SetControlledVehicleIDsRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SetControlledVehicleIDs_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::SetControlledVehicleIDsResponse>* PersonService::Stub::AsyncSetControlledVehicleIDsRaw(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleIDsRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSetControlledVehicleIDsRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::FetchControlledVehicleEnvs(::grpc::ClientContext* context, const ::city::person::v2::FetchControlledVehicleEnvsRequest& request, ::city::person::v2::FetchControlledVehicleEnvsResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::FetchControlledVehicleEnvsRequest, ::city::person::v2::FetchControlledVehicleEnvsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_FetchControlledVehicleEnvs_, context, request, response);
}

void PersonService::Stub::async::FetchControlledVehicleEnvs(::grpc::ClientContext* context, const ::city::person::v2::FetchControlledVehicleEnvsRequest* request, ::city::person::v2::FetchControlledVehicleEnvsResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::FetchControlledVehicleEnvsRequest, ::city::person::v2::FetchControlledVehicleEnvsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_FetchControlledVehicleEnvs_, context, request, response, std::move(f));
}

void PersonService::Stub::async::FetchControlledVehicleEnvs(::grpc::ClientContext* context, const ::city::person::v2::FetchControlledVehicleEnvsRequest* request, ::city::person::v2::FetchControlledVehicleEnvsResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_FetchControlledVehicleEnvs_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::FetchControlledVehicleEnvsResponse>* PersonService::Stub::PrepareAsyncFetchControlledVehicleEnvsRaw(::grpc::ClientContext* context, const ::city::person::v2::FetchControlledVehicleEnvsRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::FetchControlledVehicleEnvsResponse, ::city::person::v2::FetchControlledVehicleEnvsRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_FetchControlledVehicleEnvs_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::FetchControlledVehicleEnvsResponse>* PersonService::Stub::AsyncFetchControlledVehicleEnvsRaw(::grpc::ClientContext* context, const ::city::person::v2::FetchControlledVehicleEnvsRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncFetchControlledVehicleEnvsRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status PersonService::Stub::SetControlledVehicleActions(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleActionsRequest& request, ::city::person::v2::SetControlledVehicleActionsResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::city::person::v2::SetControlledVehicleActionsRequest, ::city::person::v2::SetControlledVehicleActionsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SetControlledVehicleActions_, context, request, response);
}

void PersonService::Stub::async::SetControlledVehicleActions(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleActionsRequest* request, ::city::person::v2::SetControlledVehicleActionsResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::city::person::v2::SetControlledVehicleActionsRequest, ::city::person::v2::SetControlledVehicleActionsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetControlledVehicleActions_, context, request, response, std::move(f));
}

void PersonService::Stub::async::SetControlledVehicleActions(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleActionsRequest* request, ::city::person::v2::SetControlledVehicleActionsResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetControlledVehicleActions_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::SetControlledVehicleActionsResponse>* PersonService::Stub::PrepareAsyncSetControlledVehicleActionsRaw(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleActionsRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::city::person::v2::SetControlledVehicleActionsResponse, ::city::person::v2::SetControlledVehicleActionsRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SetControlledVehicleActions_, context, request);
}

::grpc::ClientAsyncResponseReader< ::city::person::v2::SetControlledVehicleActionsResponse>* PersonService::Stub::AsyncSetControlledVehicleActionsRaw(::grpc::ClientContext* context, const ::city::person::v2::SetControlledVehicleActionsRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSetControlledVehicleActionsRaw(context, request, cq);
  result->StartCall();
  return result;
}

PersonService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::GetPersonRequest, ::city::person::v2::GetPersonResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::GetPersonRequest* req,
             ::city::person::v2::GetPersonResponse* resp) {
               return service->GetPerson(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::AddPersonRequest, ::city::person::v2::AddPersonResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::AddPersonRequest* req,
             ::city::person::v2::AddPersonResponse* resp) {
               return service->AddPerson(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::SetScheduleRequest, ::city::person::v2::SetScheduleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::SetScheduleRequest* req,
             ::city::person::v2::SetScheduleResponse* resp) {
               return service->SetSchedule(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[3],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::GetPersonsRequest, ::city::person::v2::GetPersonsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::GetPersonsRequest* req,
             ::city::person::v2::GetPersonsResponse* resp) {
               return service->GetPersons(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[4],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::GetPersonByLongLatBBoxRequest, ::city::person::v2::GetPersonByLongLatBBoxResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::GetPersonByLongLatBBoxRequest* req,
             ::city::person::v2::GetPersonByLongLatBBoxResponse* resp) {
               return service->GetPersonByLongLatBBox(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[5],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::GetAllVehiclesRequest, ::city::person::v2::GetAllVehiclesResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::GetAllVehiclesRequest* req,
             ::city::person::v2::GetAllVehiclesResponse* resp) {
               return service->GetAllVehicles(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[6],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::ResetPersonPositionRequest, ::city::person::v2::ResetPersonPositionResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::ResetPersonPositionRequest* req,
             ::city::person::v2::ResetPersonPositionResponse* resp) {
               return service->ResetPersonPosition(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[7],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::SetControlledVehicleIDsRequest, ::city::person::v2::SetControlledVehicleIDsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::SetControlledVehicleIDsRequest* req,
             ::city::person::v2::SetControlledVehicleIDsResponse* resp) {
               return service->SetControlledVehicleIDs(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[8],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::FetchControlledVehicleEnvsRequest, ::city::person::v2::FetchControlledVehicleEnvsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::FetchControlledVehicleEnvsRequest* req,
             ::city::person::v2::FetchControlledVehicleEnvsResponse* resp) {
               return service->FetchControlledVehicleEnvs(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[9],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::city::person::v2::SetControlledVehicleActionsRequest, ::city::person::v2::SetControlledVehicleActionsResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::city::person::v2::SetControlledVehicleActionsRequest* req,
             ::city::person::v2::SetControlledVehicleActionsResponse* resp) {
               return service->SetControlledVehicleActions(ctx, req, resp);
             }, this)));
}

PersonService::Service::~Service() {
}

::grpc::Status PersonService::Service::GetPerson(::grpc::ServerContext* context, const ::city::person::v2::GetPersonRequest* request, ::city::person::v2::GetPersonResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::AddPerson(::grpc::ServerContext* context, const ::city::person::v2::AddPersonRequest* request, ::city::person::v2::AddPersonResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::SetSchedule(::grpc::ServerContext* context, const ::city::person::v2::SetScheduleRequest* request, ::city::person::v2::SetScheduleResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::GetPersons(::grpc::ServerContext* context, const ::city::person::v2::GetPersonsRequest* request, ::city::person::v2::GetPersonsResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::GetPersonByLongLatBBox(::grpc::ServerContext* context, const ::city::person::v2::GetPersonByLongLatBBoxRequest* request, ::city::person::v2::GetPersonByLongLatBBoxResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::GetAllVehicles(::grpc::ServerContext* context, const ::city::person::v2::GetAllVehiclesRequest* request, ::city::person::v2::GetAllVehiclesResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::ResetPersonPosition(::grpc::ServerContext* context, const ::city::person::v2::ResetPersonPositionRequest* request, ::city::person::v2::ResetPersonPositionResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::SetControlledVehicleIDs(::grpc::ServerContext* context, const ::city::person::v2::SetControlledVehicleIDsRequest* request, ::city::person::v2::SetControlledVehicleIDsResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::FetchControlledVehicleEnvs(::grpc::ServerContext* context, const ::city::person::v2::FetchControlledVehicleEnvsRequest* request, ::city::person::v2::FetchControlledVehicleEnvsResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status PersonService::Service::SetControlledVehicleActions(::grpc::ServerContext* context, const ::city::person::v2::SetControlledVehicleActionsRequest* request, ::city::person::v2::SetControlledVehicleActionsResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace city
}  // namespace person
}  // namespace v2

