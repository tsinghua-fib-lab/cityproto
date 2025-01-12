// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/pause/v1/pause_service.proto

package pausev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "git.fiblab.net/sim/protos/v2/go/city/pause/v1"
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
	// PauseServiceName is the fully-qualified name of the PauseService service.
	PauseServiceName = "city.pause.v1.PauseService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PauseServicePauseProcedure is the fully-qualified name of the PauseService's Pause RPC.
	PauseServicePauseProcedure = "/city.pause.v1.PauseService/Pause"
	// PauseServiceResumeProcedure is the fully-qualified name of the PauseService's Resume RPC.
	PauseServiceResumeProcedure = "/city.pause.v1.PauseService/Resume"
)

// PauseServiceClient is a client for the city.pause.v1.PauseService service.
type PauseServiceClient interface {
	// 暂停程序
	// Pause the program
	Pause(context.Context, *connect.Request[v1.PauseRequest]) (*connect.Response[v1.PauseResponse], error)
	// 恢复程序
	// Resume the program
	Resume(context.Context, *connect.Request[v1.ResumeRequest]) (*connect.Response[v1.ResumeResponse], error)
}

// NewPauseServiceClient constructs a client for the city.pause.v1.PauseService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPauseServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PauseServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	pauseServiceMethods := v1.File_city_pause_v1_pause_service_proto.Services().ByName("PauseService").Methods()
	return &pauseServiceClient{
		pause: connect.NewClient[v1.PauseRequest, v1.PauseResponse](
			httpClient,
			baseURL+PauseServicePauseProcedure,
			connect.WithSchema(pauseServiceMethods.ByName("Pause")),
			connect.WithClientOptions(opts...),
		),
		resume: connect.NewClient[v1.ResumeRequest, v1.ResumeResponse](
			httpClient,
			baseURL+PauseServiceResumeProcedure,
			connect.WithSchema(pauseServiceMethods.ByName("Resume")),
			connect.WithClientOptions(opts...),
		),
	}
}

// pauseServiceClient implements PauseServiceClient.
type pauseServiceClient struct {
	pause  *connect.Client[v1.PauseRequest, v1.PauseResponse]
	resume *connect.Client[v1.ResumeRequest, v1.ResumeResponse]
}

// Pause calls city.pause.v1.PauseService.Pause.
func (c *pauseServiceClient) Pause(ctx context.Context, req *connect.Request[v1.PauseRequest]) (*connect.Response[v1.PauseResponse], error) {
	return c.pause.CallUnary(ctx, req)
}

// Resume calls city.pause.v1.PauseService.Resume.
func (c *pauseServiceClient) Resume(ctx context.Context, req *connect.Request[v1.ResumeRequest]) (*connect.Response[v1.ResumeResponse], error) {
	return c.resume.CallUnary(ctx, req)
}

// PauseServiceHandler is an implementation of the city.pause.v1.PauseService service.
type PauseServiceHandler interface {
	// 暂停程序
	// Pause the program
	Pause(context.Context, *connect.Request[v1.PauseRequest]) (*connect.Response[v1.PauseResponse], error)
	// 恢复程序
	// Resume the program
	Resume(context.Context, *connect.Request[v1.ResumeRequest]) (*connect.Response[v1.ResumeResponse], error)
}

// NewPauseServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPauseServiceHandler(svc PauseServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	pauseServiceMethods := v1.File_city_pause_v1_pause_service_proto.Services().ByName("PauseService").Methods()
	pauseServicePauseHandler := connect.NewUnaryHandler(
		PauseServicePauseProcedure,
		svc.Pause,
		connect.WithSchema(pauseServiceMethods.ByName("Pause")),
		connect.WithHandlerOptions(opts...),
	)
	pauseServiceResumeHandler := connect.NewUnaryHandler(
		PauseServiceResumeProcedure,
		svc.Resume,
		connect.WithSchema(pauseServiceMethods.ByName("Resume")),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.pause.v1.PauseService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PauseServicePauseProcedure:
			pauseServicePauseHandler.ServeHTTP(w, r)
		case PauseServiceResumeProcedure:
			pauseServiceResumeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPauseServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPauseServiceHandler struct{}

func (UnimplementedPauseServiceHandler) Pause(context.Context, *connect.Request[v1.PauseRequest]) (*connect.Response[v1.PauseResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.pause.v1.PauseService.Pause is not implemented"))
}

func (UnimplementedPauseServiceHandler) Resume(context.Context, *connect.Request[v1.ResumeRequest]) (*connect.Response[v1.ResumeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.pause.v1.PauseService.Resume is not implemented"))
}
