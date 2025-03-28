// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/social/v1/social_service.proto

package socialv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "git.fiblab.net/sim/protos/v2/go/city/social/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SocialServiceName is the fully-qualified name of the SocialService service.
	SocialServiceName = "city.social.v1.SocialService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SocialServiceSendProcedure is the fully-qualified name of the SocialService's Send RPC.
	SocialServiceSendProcedure = "/city.social.v1.SocialService/Send"
	// SocialServiceReceiveProcedure is the fully-qualified name of the SocialService's Receive RPC.
	SocialServiceReceiveProcedure = "/city.social.v1.SocialService/Receive"
)

// SocialServiceClient is a client for the city.social.v1.SocialService service.
type SocialServiceClient interface {
	// 发送消息
	// Send message
	Send(context.Context, *connect.Request[v1.SendRequest]) (*connect.Response[v1.SendResponse], error)
	// 接收消息，并清空该用户的消息队列
	// Receive messages and clear the user's message queue
	Receive(context.Context, *connect.Request[v1.ReceiveRequest]) (*connect.Response[v1.ReceiveResponse], error)
}

// NewSocialServiceClient constructs a client for the city.social.v1.SocialService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSocialServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SocialServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	socialServiceMethods := v1.File_city_social_v1_social_service_proto.Services().ByName("SocialService").Methods()
	return &socialServiceClient{
		send: connect.NewClient[v1.SendRequest, v1.SendResponse](
			httpClient,
			baseURL+SocialServiceSendProcedure,
			connect.WithSchema(socialServiceMethods.ByName("Send")),
			connect.WithClientOptions(opts...),
		),
		receive: connect.NewClient[v1.ReceiveRequest, v1.ReceiveResponse](
			httpClient,
			baseURL+SocialServiceReceiveProcedure,
			connect.WithSchema(socialServiceMethods.ByName("Receive")),
			connect.WithClientOptions(opts...),
		),
	}
}

// socialServiceClient implements SocialServiceClient.
type socialServiceClient struct {
	send    *connect.Client[v1.SendRequest, v1.SendResponse]
	receive *connect.Client[v1.ReceiveRequest, v1.ReceiveResponse]
}

// Send calls city.social.v1.SocialService.Send.
func (c *socialServiceClient) Send(ctx context.Context, req *connect.Request[v1.SendRequest]) (*connect.Response[v1.SendResponse], error) {
	return c.send.CallUnary(ctx, req)
}

// Receive calls city.social.v1.SocialService.Receive.
func (c *socialServiceClient) Receive(ctx context.Context, req *connect.Request[v1.ReceiveRequest]) (*connect.Response[v1.ReceiveResponse], error) {
	return c.receive.CallUnary(ctx, req)
}

// SocialServiceHandler is an implementation of the city.social.v1.SocialService service.
type SocialServiceHandler interface {
	// 发送消息
	// Send message
	Send(context.Context, *connect.Request[v1.SendRequest]) (*connect.Response[v1.SendResponse], error)
	// 接收消息，并清空该用户的消息队列
	// Receive messages and clear the user's message queue
	Receive(context.Context, *connect.Request[v1.ReceiveRequest]) (*connect.Response[v1.ReceiveResponse], error)
}

// NewSocialServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSocialServiceHandler(svc SocialServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	socialServiceMethods := v1.File_city_social_v1_social_service_proto.Services().ByName("SocialService").Methods()
	socialServiceSendHandler := connect.NewUnaryHandler(
		SocialServiceSendProcedure,
		svc.Send,
		connect.WithSchema(socialServiceMethods.ByName("Send")),
		connect.WithHandlerOptions(opts...),
	)
	socialServiceReceiveHandler := connect.NewUnaryHandler(
		SocialServiceReceiveProcedure,
		svc.Receive,
		connect.WithSchema(socialServiceMethods.ByName("Receive")),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.social.v1.SocialService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SocialServiceSendProcedure:
			socialServiceSendHandler.ServeHTTP(w, r)
		case SocialServiceReceiveProcedure:
			socialServiceReceiveHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSocialServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSocialServiceHandler struct{}

func (UnimplementedSocialServiceHandler) Send(context.Context, *connect.Request[v1.SendRequest]) (*connect.Response[v1.SendResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.social.v1.SocialService.Send is not implemented"))
}

func (UnimplementedSocialServiceHandler) Receive(context.Context, *connect.Request[v1.ReceiveRequest]) (*connect.Response[v1.ReceiveResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.social.v1.SocialService.Receive is not implemented"))
}
