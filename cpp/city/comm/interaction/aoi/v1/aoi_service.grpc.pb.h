// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/comm/interaction/aoi/v1/aoi_service.proto
#ifndef GRPC_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto__INCLUDED
#define GRPC_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto__INCLUDED

#include "city/comm/interaction/aoi/v1/aoi_service.pb.h"

#include <functional>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/client_context.h>
#include <grpcpp/completion_queue.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/proto_utils.h>
#include <grpcpp/impl/rpc_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/status.h>
#include <grpcpp/support/stub_options.h>
#include <grpcpp/support/sync_stream.h>

namespace city {
namespace comm {
namespace interaction {
namespace aoi {
namespace v1 {

class AoiService final {
 public:
  static constexpr char const* service_full_name() {
    return "city.comm.interaction.aoi.v1.AoiService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status GetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>> AsyncGetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>>(AsyncGetBadAoiIDRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>> PrepareAsyncGetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>>(PrepareAsyncGetBadAoiIDRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void GetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* AsyncGetBadAoiIDRaw(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* PrepareAsyncGetBadAoiIDRaw(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status GetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>> AsyncGetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>>(AsyncGetBadAoiIDRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>> PrepareAsyncGetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>>(PrepareAsyncGetBadAoiIDRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void GetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response, std::function<void(::grpc::Status)>) override;
      void GetBadAoiID(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* AsyncGetBadAoiIDRaw(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* PrepareAsyncGetBadAoiIDRaw(::grpc::ClientContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_GetBadAoiID_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status GetBadAoiID(::grpc::ServerContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_GetBadAoiID : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_GetBadAoiID() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_GetBadAoiID() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetBadAoiID(::grpc::ServerContext* /*context*/, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* /*request*/, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetBadAoiID(::grpc::ServerContext* context, ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::grpc::ServerAsyncResponseWriter< ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_GetBadAoiID<Service > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_GetBadAoiID : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_GetBadAoiID() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* request, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* response) { return this->GetBadAoiID(context, request, response); }));}
    void SetMessageAllocatorFor_GetBadAoiID(
        ::grpc::MessageAllocator< ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_GetBadAoiID() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetBadAoiID(::grpc::ServerContext* /*context*/, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* /*request*/, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* GetBadAoiID(
      ::grpc::CallbackServerContext* /*context*/, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* /*request*/, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_GetBadAoiID<Service > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_GetBadAoiID : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_GetBadAoiID() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_GetBadAoiID() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetBadAoiID(::grpc::ServerContext* /*context*/, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* /*request*/, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_GetBadAoiID : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_GetBadAoiID() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_GetBadAoiID() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetBadAoiID(::grpc::ServerContext* /*context*/, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* /*request*/, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetBadAoiID(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_GetBadAoiID : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_GetBadAoiID() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->GetBadAoiID(context, request, response); }));
    }
    ~WithRawCallbackMethod_GetBadAoiID() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetBadAoiID(::grpc::ServerContext* /*context*/, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* /*request*/, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* GetBadAoiID(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetBadAoiID : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_GetBadAoiID() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* streamer) {
                       return this->StreamedGetBadAoiID(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_GetBadAoiID() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetBadAoiID(::grpc::ServerContext* /*context*/, const ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest* /*request*/, ::city::comm::interaction::aoi::v1::GetBadAoiIDResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetBadAoiID(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::city::comm::interaction::aoi::v1::GetBadAoiIDRequest,::city::comm::interaction::aoi::v1::GetBadAoiIDResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_GetBadAoiID<Service > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_GetBadAoiID<Service > StreamedService;
};

}  // namespace v1
}  // namespace aoi
}  // namespace interaction
}  // namespace comm
}  // namespace city


#endif  // GRPC_city_2fcomm_2finteraction_2faoi_2fv1_2faoi_5fservice_2eproto__INCLUDED
