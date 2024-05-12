// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/elec/input/v1/input_service.proto
#ifndef GRPC_city_2felec_2finput_2fv1_2finput_5fservice_2eproto__INCLUDED
#define GRPC_city_2felec_2finput_2fv1_2finput_5fservice_2eproto__INCLUDED

#include "city/elec/input/v1/input_service.pb.h"

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
namespace elec {
namespace input {
namespace v1 {

class InputService final {
 public:
  static constexpr char const* service_full_name() {
    return "city.elec.input.v1.InputService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status Init(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::city::elec::input::v1::InitResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::elec::input::v1::InitResponse>> AsyncInit(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::elec::input::v1::InitResponse>>(AsyncInitRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::elec::input::v1::InitResponse>> PrepareAsyncInit(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::elec::input::v1::InitResponse>>(PrepareAsyncInitRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void Init(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest* request, ::city::elec::input::v1::InitResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Init(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest* request, ::city::elec::input::v1::InitResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::city::elec::input::v1::InitResponse>* AsyncInitRaw(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::city::elec::input::v1::InitResponse>* PrepareAsyncInitRaw(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status Init(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::city::elec::input::v1::InitResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::elec::input::v1::InitResponse>> AsyncInit(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::elec::input::v1::InitResponse>>(AsyncInitRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::elec::input::v1::InitResponse>> PrepareAsyncInit(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::elec::input::v1::InitResponse>>(PrepareAsyncInitRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void Init(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest* request, ::city::elec::input::v1::InitResponse* response, std::function<void(::grpc::Status)>) override;
      void Init(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest* request, ::city::elec::input::v1::InitResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
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
    ::grpc::ClientAsyncResponseReader< ::city::elec::input::v1::InitResponse>* AsyncInitRaw(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::city::elec::input::v1::InitResponse>* PrepareAsyncInitRaw(::grpc::ClientContext* context, const ::city::elec::input::v1::InitRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_Init_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status Init(::grpc::ServerContext* context, const ::city::elec::input::v1::InitRequest* request, ::city::elec::input::v1::InitResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_Init : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_Init() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_Init() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Init(::grpc::ServerContext* /*context*/, const ::city::elec::input::v1::InitRequest* /*request*/, ::city::elec::input::v1::InitResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestInit(::grpc::ServerContext* context, ::city::elec::input::v1::InitRequest* request, ::grpc::ServerAsyncResponseWriter< ::city::elec::input::v1::InitResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_Init<Service > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_Init : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_Init() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::city::elec::input::v1::InitRequest, ::city::elec::input::v1::InitResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::city::elec::input::v1::InitRequest* request, ::city::elec::input::v1::InitResponse* response) { return this->Init(context, request, response); }));}
    void SetMessageAllocatorFor_Init(
        ::grpc::MessageAllocator< ::city::elec::input::v1::InitRequest, ::city::elec::input::v1::InitResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::city::elec::input::v1::InitRequest, ::city::elec::input::v1::InitResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_Init() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Init(::grpc::ServerContext* /*context*/, const ::city::elec::input::v1::InitRequest* /*request*/, ::city::elec::input::v1::InitResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Init(
      ::grpc::CallbackServerContext* /*context*/, const ::city::elec::input::v1::InitRequest* /*request*/, ::city::elec::input::v1::InitResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_Init<Service > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_Init : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_Init() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_Init() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Init(::grpc::ServerContext* /*context*/, const ::city::elec::input::v1::InitRequest* /*request*/, ::city::elec::input::v1::InitResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_Init : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_Init() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_Init() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Init(::grpc::ServerContext* /*context*/, const ::city::elec::input::v1::InitRequest* /*request*/, ::city::elec::input::v1::InitResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestInit(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_Init : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_Init() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->Init(context, request, response); }));
    }
    ~WithRawCallbackMethod_Init() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Init(::grpc::ServerContext* /*context*/, const ::city::elec::input::v1::InitRequest* /*request*/, ::city::elec::input::v1::InitResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Init(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Init : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_Init() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::city::elec::input::v1::InitRequest, ::city::elec::input::v1::InitResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::city::elec::input::v1::InitRequest, ::city::elec::input::v1::InitResponse>* streamer) {
                       return this->StreamedInit(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_Init() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Init(::grpc::ServerContext* /*context*/, const ::city::elec::input::v1::InitRequest* /*request*/, ::city::elec::input::v1::InitResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedInit(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::city::elec::input::v1::InitRequest,::city::elec::input::v1::InitResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_Init<Service > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_Init<Service > StreamedService;
};

}  // namespace v1
}  // namespace input
}  // namespace elec
}  // namespace city


#endif  // GRPC_city_2felec_2finput_2fv1_2finput_5fservice_2eproto__INCLUDED
