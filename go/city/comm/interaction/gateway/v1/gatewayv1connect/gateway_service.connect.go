// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/comm/interaction/gateway/v1/gateway_service.proto

package gatewayv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "git.fiblab.net/sim/protos/v2/go/city/comm/interaction/gateway/v1"
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
	// GatewayServiceName is the fully-qualified name of the GatewayService service.
	GatewayServiceName = "city.comm.interaction.gateway.v1.GatewayService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GatewayServiceSetGatewayPowerStatusProcedure is the fully-qualified name of the GatewayService's
	// SetGatewayPowerStatus RPC.
	GatewayServiceSetGatewayPowerStatusProcedure = "/city.comm.interaction.gateway.v1.GatewayService/SetGatewayPowerStatus"
	// GatewayServiceSetGatewayRuinStatusProcedure is the fully-qualified name of the GatewayService's
	// SetGatewayRuinStatus RPC.
	GatewayServiceSetGatewayRuinStatusProcedure = "/city.comm.interaction.gateway.v1.GatewayService/SetGatewayRuinStatus"
	// GatewayServiceGetAllStatusProcedure is the fully-qualified name of the GatewayService's
	// GetAllStatus RPC.
	GatewayServiceGetAllStatusProcedure = "/city.comm.interaction.gateway.v1.GatewayService/GetAllStatus"
	// GatewayServiceGetRuinInfoProcedure is the fully-qualified name of the GatewayService's
	// GetRuinInfo RPC.
	GatewayServiceGetRuinInfoProcedure = "/city.comm.interaction.gateway.v1.GatewayService/GetRuinInfo"
	// GatewayServiceGetEventsProcedure is the fully-qualified name of the GatewayService's GetEvents
	// RPC.
	GatewayServiceGetEventsProcedure = "/city.comm.interaction.gateway.v1.GatewayService/GetEvents"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	gatewayServiceServiceDescriptor                     = v1.File_city_comm_interaction_gateway_v1_gateway_service_proto.Services().ByName("GatewayService")
	gatewayServiceSetGatewayPowerStatusMethodDescriptor = gatewayServiceServiceDescriptor.Methods().ByName("SetGatewayPowerStatus")
	gatewayServiceSetGatewayRuinStatusMethodDescriptor  = gatewayServiceServiceDescriptor.Methods().ByName("SetGatewayRuinStatus")
	gatewayServiceGetAllStatusMethodDescriptor          = gatewayServiceServiceDescriptor.Methods().ByName("GetAllStatus")
	gatewayServiceGetRuinInfoMethodDescriptor           = gatewayServiceServiceDescriptor.Methods().ByName("GetRuinInfo")
	gatewayServiceGetEventsMethodDescriptor             = gatewayServiceServiceDescriptor.Methods().ByName("GetEvents")
)

// GatewayServiceClient is a client for the city.comm.interaction.gateway.v1.GatewayService service.
type GatewayServiceClient interface {
	SetGatewayPowerStatus(context.Context, *connect.Request[v1.SetGatewayPowerStatusRequest]) (*connect.Response[v1.SetGatewayPowerStatusResponse], error)
	SetGatewayRuinStatus(context.Context, *connect.Request[v1.SetGatewayRuinStatusRequest]) (*connect.Response[v1.SetGatewayRuinStatusResponse], error)
	GetAllStatus(context.Context, *connect.Request[v1.GetAllStatusRequest]) (*connect.Response[v1.GetAllStatusResponse], error)
	GetRuinInfo(context.Context, *connect.Request[v1.GetRuinInfoRequest]) (*connect.Response[v1.GetRuinInfoResponse], error)
	GetEvents(context.Context, *connect.Request[v1.GetEventsRequest]) (*connect.Response[v1.GetEventsResponse], error)
}

// NewGatewayServiceClient constructs a client for the
// city.comm.interaction.gateway.v1.GatewayService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGatewayServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GatewayServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &gatewayServiceClient{
		setGatewayPowerStatus: connect.NewClient[v1.SetGatewayPowerStatusRequest, v1.SetGatewayPowerStatusResponse](
			httpClient,
			baseURL+GatewayServiceSetGatewayPowerStatusProcedure,
			connect.WithSchema(gatewayServiceSetGatewayPowerStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setGatewayRuinStatus: connect.NewClient[v1.SetGatewayRuinStatusRequest, v1.SetGatewayRuinStatusResponse](
			httpClient,
			baseURL+GatewayServiceSetGatewayRuinStatusProcedure,
			connect.WithSchema(gatewayServiceSetGatewayRuinStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAllStatus: connect.NewClient[v1.GetAllStatusRequest, v1.GetAllStatusResponse](
			httpClient,
			baseURL+GatewayServiceGetAllStatusProcedure,
			connect.WithSchema(gatewayServiceGetAllStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getRuinInfo: connect.NewClient[v1.GetRuinInfoRequest, v1.GetRuinInfoResponse](
			httpClient,
			baseURL+GatewayServiceGetRuinInfoProcedure,
			connect.WithSchema(gatewayServiceGetRuinInfoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getEvents: connect.NewClient[v1.GetEventsRequest, v1.GetEventsResponse](
			httpClient,
			baseURL+GatewayServiceGetEventsProcedure,
			connect.WithSchema(gatewayServiceGetEventsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// gatewayServiceClient implements GatewayServiceClient.
type gatewayServiceClient struct {
	setGatewayPowerStatus *connect.Client[v1.SetGatewayPowerStatusRequest, v1.SetGatewayPowerStatusResponse]
	setGatewayRuinStatus  *connect.Client[v1.SetGatewayRuinStatusRequest, v1.SetGatewayRuinStatusResponse]
	getAllStatus          *connect.Client[v1.GetAllStatusRequest, v1.GetAllStatusResponse]
	getRuinInfo           *connect.Client[v1.GetRuinInfoRequest, v1.GetRuinInfoResponse]
	getEvents             *connect.Client[v1.GetEventsRequest, v1.GetEventsResponse]
}

// SetGatewayPowerStatus calls
// city.comm.interaction.gateway.v1.GatewayService.SetGatewayPowerStatus.
func (c *gatewayServiceClient) SetGatewayPowerStatus(ctx context.Context, req *connect.Request[v1.SetGatewayPowerStatusRequest]) (*connect.Response[v1.SetGatewayPowerStatusResponse], error) {
	return c.setGatewayPowerStatus.CallUnary(ctx, req)
}

// SetGatewayRuinStatus calls city.comm.interaction.gateway.v1.GatewayService.SetGatewayRuinStatus.
func (c *gatewayServiceClient) SetGatewayRuinStatus(ctx context.Context, req *connect.Request[v1.SetGatewayRuinStatusRequest]) (*connect.Response[v1.SetGatewayRuinStatusResponse], error) {
	return c.setGatewayRuinStatus.CallUnary(ctx, req)
}

// GetAllStatus calls city.comm.interaction.gateway.v1.GatewayService.GetAllStatus.
func (c *gatewayServiceClient) GetAllStatus(ctx context.Context, req *connect.Request[v1.GetAllStatusRequest]) (*connect.Response[v1.GetAllStatusResponse], error) {
	return c.getAllStatus.CallUnary(ctx, req)
}

// GetRuinInfo calls city.comm.interaction.gateway.v1.GatewayService.GetRuinInfo.
func (c *gatewayServiceClient) GetRuinInfo(ctx context.Context, req *connect.Request[v1.GetRuinInfoRequest]) (*connect.Response[v1.GetRuinInfoResponse], error) {
	return c.getRuinInfo.CallUnary(ctx, req)
}

// GetEvents calls city.comm.interaction.gateway.v1.GatewayService.GetEvents.
func (c *gatewayServiceClient) GetEvents(ctx context.Context, req *connect.Request[v1.GetEventsRequest]) (*connect.Response[v1.GetEventsResponse], error) {
	return c.getEvents.CallUnary(ctx, req)
}

// GatewayServiceHandler is an implementation of the city.comm.interaction.gateway.v1.GatewayService
// service.
type GatewayServiceHandler interface {
	SetGatewayPowerStatus(context.Context, *connect.Request[v1.SetGatewayPowerStatusRequest]) (*connect.Response[v1.SetGatewayPowerStatusResponse], error)
	SetGatewayRuinStatus(context.Context, *connect.Request[v1.SetGatewayRuinStatusRequest]) (*connect.Response[v1.SetGatewayRuinStatusResponse], error)
	GetAllStatus(context.Context, *connect.Request[v1.GetAllStatusRequest]) (*connect.Response[v1.GetAllStatusResponse], error)
	GetRuinInfo(context.Context, *connect.Request[v1.GetRuinInfoRequest]) (*connect.Response[v1.GetRuinInfoResponse], error)
	GetEvents(context.Context, *connect.Request[v1.GetEventsRequest]) (*connect.Response[v1.GetEventsResponse], error)
}

// NewGatewayServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGatewayServiceHandler(svc GatewayServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	gatewayServiceSetGatewayPowerStatusHandler := connect.NewUnaryHandler(
		GatewayServiceSetGatewayPowerStatusProcedure,
		svc.SetGatewayPowerStatus,
		connect.WithSchema(gatewayServiceSetGatewayPowerStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gatewayServiceSetGatewayRuinStatusHandler := connect.NewUnaryHandler(
		GatewayServiceSetGatewayRuinStatusProcedure,
		svc.SetGatewayRuinStatus,
		connect.WithSchema(gatewayServiceSetGatewayRuinStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gatewayServiceGetAllStatusHandler := connect.NewUnaryHandler(
		GatewayServiceGetAllStatusProcedure,
		svc.GetAllStatus,
		connect.WithSchema(gatewayServiceGetAllStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gatewayServiceGetRuinInfoHandler := connect.NewUnaryHandler(
		GatewayServiceGetRuinInfoProcedure,
		svc.GetRuinInfo,
		connect.WithSchema(gatewayServiceGetRuinInfoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gatewayServiceGetEventsHandler := connect.NewUnaryHandler(
		GatewayServiceGetEventsProcedure,
		svc.GetEvents,
		connect.WithSchema(gatewayServiceGetEventsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.comm.interaction.gateway.v1.GatewayService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GatewayServiceSetGatewayPowerStatusProcedure:
			gatewayServiceSetGatewayPowerStatusHandler.ServeHTTP(w, r)
		case GatewayServiceSetGatewayRuinStatusProcedure:
			gatewayServiceSetGatewayRuinStatusHandler.ServeHTTP(w, r)
		case GatewayServiceGetAllStatusProcedure:
			gatewayServiceGetAllStatusHandler.ServeHTTP(w, r)
		case GatewayServiceGetRuinInfoProcedure:
			gatewayServiceGetRuinInfoHandler.ServeHTTP(w, r)
		case GatewayServiceGetEventsProcedure:
			gatewayServiceGetEventsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGatewayServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGatewayServiceHandler struct{}

func (UnimplementedGatewayServiceHandler) SetGatewayPowerStatus(context.Context, *connect.Request[v1.SetGatewayPowerStatusRequest]) (*connect.Response[v1.SetGatewayPowerStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.comm.interaction.gateway.v1.GatewayService.SetGatewayPowerStatus is not implemented"))
}

func (UnimplementedGatewayServiceHandler) SetGatewayRuinStatus(context.Context, *connect.Request[v1.SetGatewayRuinStatusRequest]) (*connect.Response[v1.SetGatewayRuinStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.comm.interaction.gateway.v1.GatewayService.SetGatewayRuinStatus is not implemented"))
}

func (UnimplementedGatewayServiceHandler) GetAllStatus(context.Context, *connect.Request[v1.GetAllStatusRequest]) (*connect.Response[v1.GetAllStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.comm.interaction.gateway.v1.GatewayService.GetAllStatus is not implemented"))
}

func (UnimplementedGatewayServiceHandler) GetRuinInfo(context.Context, *connect.Request[v1.GetRuinInfoRequest]) (*connect.Response[v1.GetRuinInfoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.comm.interaction.gateway.v1.GatewayService.GetRuinInfo is not implemented"))
}

func (UnimplementedGatewayServiceHandler) GetEvents(context.Context, *connect.Request[v1.GetEventsRequest]) (*connect.Response[v1.GetEventsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.comm.interaction.gateway.v1.GatewayService.GetEvents is not implemented"))
}
