// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/clock/v1/clock_service.proto

package clockv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "git.fiblab.net/sim/protos/go/city/clock/v1"
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
	// ClockServiceName is the fully-qualified name of the ClockService service.
	ClockServiceName = "city.clock.v1.ClockService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ClockServiceNowProcedure is the fully-qualified name of the ClockService's Now RPC.
	ClockServiceNowProcedure = "/city.clock.v1.ClockService/Now"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	clockServiceServiceDescriptor   = v1.File_city_clock_v1_clock_service_proto.Services().ByName("ClockService")
	clockServiceNowMethodDescriptor = clockServiceServiceDescriptor.Methods().ByName("Now")
)

// ClockServiceClient is a client for the city.clock.v1.ClockService service.
type ClockServiceClient interface {
	// 获取当前的模拟时间
	// get current simulation clock
	Now(context.Context, *connect.Request[v1.NowRequest]) (*connect.Response[v1.NowResponse], error)
}

// NewClockServiceClient constructs a client for the city.clock.v1.ClockService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewClockServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ClockServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &clockServiceClient{
		now: connect.NewClient[v1.NowRequest, v1.NowResponse](
			httpClient,
			baseURL+ClockServiceNowProcedure,
			connect.WithSchema(clockServiceNowMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// clockServiceClient implements ClockServiceClient.
type clockServiceClient struct {
	now *connect.Client[v1.NowRequest, v1.NowResponse]
}

// Now calls city.clock.v1.ClockService.Now.
func (c *clockServiceClient) Now(ctx context.Context, req *connect.Request[v1.NowRequest]) (*connect.Response[v1.NowResponse], error) {
	return c.now.CallUnary(ctx, req)
}

// ClockServiceHandler is an implementation of the city.clock.v1.ClockService service.
type ClockServiceHandler interface {
	// 获取当前的模拟时间
	// get current simulation clock
	Now(context.Context, *connect.Request[v1.NowRequest]) (*connect.Response[v1.NowResponse], error)
}

// NewClockServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewClockServiceHandler(svc ClockServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	clockServiceNowHandler := connect.NewUnaryHandler(
		ClockServiceNowProcedure,
		svc.Now,
		connect.WithSchema(clockServiceNowMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.clock.v1.ClockService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ClockServiceNowProcedure:
			clockServiceNowHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedClockServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedClockServiceHandler struct{}

func (UnimplementedClockServiceHandler) Now(context.Context, *connect.Request[v1.NowRequest]) (*connect.Response[v1.NowResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.clock.v1.ClockService.Now is not implemented"))
}
