// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: city/pause/v1/pause_service.proto
#ifndef GRPC_city_2fpause_2fv1_2fpause_5fservice_2eproto__INCLUDED
#define GRPC_city_2fpause_2fv1_2fpause_5fservice_2eproto__INCLUDED

#include "city/pause/v1/pause_service.pb.h"

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
namespace pause {
namespace v1 {

class PauseService final {
 public:
  static constexpr char const* service_full_name() {
    return "city.pause.v1.PauseService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    // 暂停程序
    // Pause the program
    virtual ::grpc::Status Pause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::city::pause::v1::PauseResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::PauseResponse>> AsyncPause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::PauseResponse>>(AsyncPauseRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::PauseResponse>> PrepareAsyncPause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::PauseResponse>>(PrepareAsyncPauseRaw(context, request, cq));
    }
    // 恢复程序
    // Resume the program
    virtual ::grpc::Status Resume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::city::pause::v1::ResumeResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::ResumeResponse>> AsyncResume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::ResumeResponse>>(AsyncResumeRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::ResumeResponse>> PrepareAsyncResume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::ResumeResponse>>(PrepareAsyncResumeRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      // 暂停程序
      // Pause the program
      virtual void Pause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest* request, ::city::pause::v1::PauseResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Pause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest* request, ::city::pause::v1::PauseResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      // 恢复程序
      // Resume the program
      virtual void Resume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest* request, ::city::pause::v1::ResumeResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Resume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest* request, ::city::pause::v1::ResumeResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::PauseResponse>* AsyncPauseRaw(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::PauseResponse>* PrepareAsyncPauseRaw(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::ResumeResponse>* AsyncResumeRaw(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::city::pause::v1::ResumeResponse>* PrepareAsyncResumeRaw(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status Pause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::city::pause::v1::PauseResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::pause::v1::PauseResponse>> AsyncPause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::pause::v1::PauseResponse>>(AsyncPauseRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::pause::v1::PauseResponse>> PrepareAsyncPause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::pause::v1::PauseResponse>>(PrepareAsyncPauseRaw(context, request, cq));
    }
    ::grpc::Status Resume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::city::pause::v1::ResumeResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::pause::v1::ResumeResponse>> AsyncResume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::pause::v1::ResumeResponse>>(AsyncResumeRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::pause::v1::ResumeResponse>> PrepareAsyncResume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::city::pause::v1::ResumeResponse>>(PrepareAsyncResumeRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void Pause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest* request, ::city::pause::v1::PauseResponse* response, std::function<void(::grpc::Status)>) override;
      void Pause(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest* request, ::city::pause::v1::PauseResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void Resume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest* request, ::city::pause::v1::ResumeResponse* response, std::function<void(::grpc::Status)>) override;
      void Resume(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest* request, ::city::pause::v1::ResumeResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
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
    ::grpc::ClientAsyncResponseReader< ::city::pause::v1::PauseResponse>* AsyncPauseRaw(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::city::pause::v1::PauseResponse>* PrepareAsyncPauseRaw(::grpc::ClientContext* context, const ::city::pause::v1::PauseRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::city::pause::v1::ResumeResponse>* AsyncResumeRaw(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::city::pause::v1::ResumeResponse>* PrepareAsyncResumeRaw(::grpc::ClientContext* context, const ::city::pause::v1::ResumeRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_Pause_;
    const ::grpc::internal::RpcMethod rpcmethod_Resume_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    // 暂停程序
    // Pause the program
    virtual ::grpc::Status Pause(::grpc::ServerContext* context, const ::city::pause::v1::PauseRequest* request, ::city::pause::v1::PauseResponse* response);
    // 恢复程序
    // Resume the program
    virtual ::grpc::Status Resume(::grpc::ServerContext* context, const ::city::pause::v1::ResumeRequest* request, ::city::pause::v1::ResumeResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_Pause : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_Pause() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_Pause() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Pause(::grpc::ServerContext* /*context*/, const ::city::pause::v1::PauseRequest* /*request*/, ::city::pause::v1::PauseResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestPause(::grpc::ServerContext* context, ::city::pause::v1::PauseRequest* request, ::grpc::ServerAsyncResponseWriter< ::city::pause::v1::PauseResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_Resume : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_Resume() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_Resume() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Resume(::grpc::ServerContext* /*context*/, const ::city::pause::v1::ResumeRequest* /*request*/, ::city::pause::v1::ResumeResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestResume(::grpc::ServerContext* context, ::city::pause::v1::ResumeRequest* request, ::grpc::ServerAsyncResponseWriter< ::city::pause::v1::ResumeResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_Pause<WithAsyncMethod_Resume<Service > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_Pause : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_Pause() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::city::pause::v1::PauseRequest, ::city::pause::v1::PauseResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::city::pause::v1::PauseRequest* request, ::city::pause::v1::PauseResponse* response) { return this->Pause(context, request, response); }));}
    void SetMessageAllocatorFor_Pause(
        ::grpc::MessageAllocator< ::city::pause::v1::PauseRequest, ::city::pause::v1::PauseResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::city::pause::v1::PauseRequest, ::city::pause::v1::PauseResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_Pause() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Pause(::grpc::ServerContext* /*context*/, const ::city::pause::v1::PauseRequest* /*request*/, ::city::pause::v1::PauseResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Pause(
      ::grpc::CallbackServerContext* /*context*/, const ::city::pause::v1::PauseRequest* /*request*/, ::city::pause::v1::PauseResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_Resume : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_Resume() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::city::pause::v1::ResumeRequest, ::city::pause::v1::ResumeResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::city::pause::v1::ResumeRequest* request, ::city::pause::v1::ResumeResponse* response) { return this->Resume(context, request, response); }));}
    void SetMessageAllocatorFor_Resume(
        ::grpc::MessageAllocator< ::city::pause::v1::ResumeRequest, ::city::pause::v1::ResumeResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(1);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::city::pause::v1::ResumeRequest, ::city::pause::v1::ResumeResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_Resume() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Resume(::grpc::ServerContext* /*context*/, const ::city::pause::v1::ResumeRequest* /*request*/, ::city::pause::v1::ResumeResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Resume(
      ::grpc::CallbackServerContext* /*context*/, const ::city::pause::v1::ResumeRequest* /*request*/, ::city::pause::v1::ResumeResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_Pause<WithCallbackMethod_Resume<Service > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_Pause : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_Pause() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_Pause() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Pause(::grpc::ServerContext* /*context*/, const ::city::pause::v1::PauseRequest* /*request*/, ::city::pause::v1::PauseResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_Resume : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_Resume() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_Resume() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Resume(::grpc::ServerContext* /*context*/, const ::city::pause::v1::ResumeRequest* /*request*/, ::city::pause::v1::ResumeResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_Pause : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_Pause() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_Pause() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Pause(::grpc::ServerContext* /*context*/, const ::city::pause::v1::PauseRequest* /*request*/, ::city::pause::v1::PauseResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestPause(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_Resume : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_Resume() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_Resume() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Resume(::grpc::ServerContext* /*context*/, const ::city::pause::v1::ResumeRequest* /*request*/, ::city::pause::v1::ResumeResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestResume(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_Pause : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_Pause() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->Pause(context, request, response); }));
    }
    ~WithRawCallbackMethod_Pause() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Pause(::grpc::ServerContext* /*context*/, const ::city::pause::v1::PauseRequest* /*request*/, ::city::pause::v1::PauseResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Pause(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_Resume : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_Resume() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->Resume(context, request, response); }));
    }
    ~WithRawCallbackMethod_Resume() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Resume(::grpc::ServerContext* /*context*/, const ::city::pause::v1::ResumeRequest* /*request*/, ::city::pause::v1::ResumeResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Resume(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Pause : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_Pause() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::city::pause::v1::PauseRequest, ::city::pause::v1::PauseResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::city::pause::v1::PauseRequest, ::city::pause::v1::PauseResponse>* streamer) {
                       return this->StreamedPause(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_Pause() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Pause(::grpc::ServerContext* /*context*/, const ::city::pause::v1::PauseRequest* /*request*/, ::city::pause::v1::PauseResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedPause(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::city::pause::v1::PauseRequest,::city::pause::v1::PauseResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Resume : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_Resume() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler<
          ::city::pause::v1::ResumeRequest, ::city::pause::v1::ResumeResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::city::pause::v1::ResumeRequest, ::city::pause::v1::ResumeResponse>* streamer) {
                       return this->StreamedResume(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_Resume() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Resume(::grpc::ServerContext* /*context*/, const ::city::pause::v1::ResumeRequest* /*request*/, ::city::pause::v1::ResumeResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedResume(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::city::pause::v1::ResumeRequest,::city::pause::v1::ResumeResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_Pause<WithStreamedUnaryMethod_Resume<Service > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_Pause<WithStreamedUnaryMethod_Resume<Service > > StreamedService;
};

}  // namespace v1
}  // namespace pause
}  // namespace city


#endif  // GRPC_city_2fpause_2fv1_2fpause_5fservice_2eproto__INCLUDED
