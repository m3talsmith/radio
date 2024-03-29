// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: radio.proto

#include "radio.pb.h"
#include "radio.grpc.pb.h"

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
namespace radio {

static const char* RadioAPI_method_names[] = {
  "/radio.RadioAPI/Station",
};

std::unique_ptr< RadioAPI::Stub> RadioAPI::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< RadioAPI::Stub> stub(new RadioAPI::Stub(channel, options));
  return stub;
}

RadioAPI::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_Station_(RadioAPI_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::BIDI_STREAMING, channel)
  {}

::grpc::ClientReaderWriter< ::radio::Request, ::radio::Broadcast>* RadioAPI::Stub::StationRaw(::grpc::ClientContext* context) {
  return ::grpc::internal::ClientReaderWriterFactory< ::radio::Request, ::radio::Broadcast>::Create(channel_.get(), rpcmethod_Station_, context);
}

void RadioAPI::Stub::async::Station(::grpc::ClientContext* context, ::grpc::ClientBidiReactor< ::radio::Request,::radio::Broadcast>* reactor) {
  ::grpc::internal::ClientCallbackReaderWriterFactory< ::radio::Request,::radio::Broadcast>::Create(stub_->channel_.get(), stub_->rpcmethod_Station_, context, reactor);
}

::grpc::ClientAsyncReaderWriter< ::radio::Request, ::radio::Broadcast>* RadioAPI::Stub::AsyncStationRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncReaderWriterFactory< ::radio::Request, ::radio::Broadcast>::Create(channel_.get(), cq, rpcmethod_Station_, context, true, tag);
}

::grpc::ClientAsyncReaderWriter< ::radio::Request, ::radio::Broadcast>* RadioAPI::Stub::PrepareAsyncStationRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncReaderWriterFactory< ::radio::Request, ::radio::Broadcast>::Create(channel_.get(), cq, rpcmethod_Station_, context, false, nullptr);
}

RadioAPI::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      RadioAPI_method_names[0],
      ::grpc::internal::RpcMethod::BIDI_STREAMING,
      new ::grpc::internal::BidiStreamingHandler< RadioAPI::Service, ::radio::Request, ::radio::Broadcast>(
          [](RadioAPI::Service* service,
             ::grpc::ServerContext* ctx,
             ::grpc::ServerReaderWriter<::radio::Broadcast,
             ::radio::Request>* stream) {
               return service->Station(ctx, stream);
             }, this)));
}

RadioAPI::Service::~Service() {
}

::grpc::Status RadioAPI::Service::Station(::grpc::ServerContext* context, ::grpc::ServerReaderWriter< ::radio::Broadcast, ::radio::Request>* stream) {
  (void) context;
  (void) stream;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace radio

