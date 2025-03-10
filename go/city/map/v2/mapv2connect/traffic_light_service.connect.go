// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/map/v2/traffic_light_service.proto

package mapv2connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2 "git.fiblab.net/sim/protos/v2/go/city/map/v2"
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
	// TrafficLightServiceName is the fully-qualified name of the TrafficLightService service.
	TrafficLightServiceName = "city.map.v2.TrafficLightService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TrafficLightServiceGetTrafficLightProcedure is the fully-qualified name of the
	// TrafficLightService's GetTrafficLight RPC.
	TrafficLightServiceGetTrafficLightProcedure = "/city.map.v2.TrafficLightService/GetTrafficLight"
	// TrafficLightServiceSetTrafficLightProcedure is the fully-qualified name of the
	// TrafficLightService's SetTrafficLight RPC.
	TrafficLightServiceSetTrafficLightProcedure = "/city.map.v2.TrafficLightService/SetTrafficLight"
	// TrafficLightServiceSetTrafficLightPhaseProcedure is the fully-qualified name of the
	// TrafficLightService's SetTrafficLightPhase RPC.
	TrafficLightServiceSetTrafficLightPhaseProcedure = "/city.map.v2.TrafficLightService/SetTrafficLightPhase"
	// TrafficLightServiceSetTrafficLightStatusProcedure is the fully-qualified name of the
	// TrafficLightService's SetTrafficLightStatus RPC.
	TrafficLightServiceSetTrafficLightStatusProcedure = "/city.map.v2.TrafficLightService/SetTrafficLightStatus"
)

// TrafficLightServiceClient is a client for the city.map.v2.TrafficLightService service.
type TrafficLightServiceClient interface {
	// 获取路口的红绿灯信息
	// Get traffic light information
	GetTrafficLight(context.Context, *connect.Request[v2.GetTrafficLightRequest]) (*connect.Response[v2.GetTrafficLightResponse], error)
	// 设置路口的红绿灯信息
	// Set traffic light information
	SetTrafficLight(context.Context, *connect.Request[v2.SetTrafficLightRequest]) (*connect.Response[v2.SetTrafficLightResponse], error)
	// 设置路口的红绿灯相位
	// Set traffic light phase
	SetTrafficLightPhase(context.Context, *connect.Request[v2.SetTrafficLightPhaseRequest]) (*connect.Response[v2.SetTrafficLightPhaseResponse], error)
	// 设置路口的红绿灯状态
	// Set traffic light status
	SetTrafficLightStatus(context.Context, *connect.Request[v2.SetTrafficLightStatusRequest]) (*connect.Response[v2.SetTrafficLightStatusResponse], error)
}

// NewTrafficLightServiceClient constructs a client for the city.map.v2.TrafficLightService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTrafficLightServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TrafficLightServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	trafficLightServiceMethods := v2.File_city_map_v2_traffic_light_service_proto.Services().ByName("TrafficLightService").Methods()
	return &trafficLightServiceClient{
		getTrafficLight: connect.NewClient[v2.GetTrafficLightRequest, v2.GetTrafficLightResponse](
			httpClient,
			baseURL+TrafficLightServiceGetTrafficLightProcedure,
			connect.WithSchema(trafficLightServiceMethods.ByName("GetTrafficLight")),
			connect.WithClientOptions(opts...),
		),
		setTrafficLight: connect.NewClient[v2.SetTrafficLightRequest, v2.SetTrafficLightResponse](
			httpClient,
			baseURL+TrafficLightServiceSetTrafficLightProcedure,
			connect.WithSchema(trafficLightServiceMethods.ByName("SetTrafficLight")),
			connect.WithClientOptions(opts...),
		),
		setTrafficLightPhase: connect.NewClient[v2.SetTrafficLightPhaseRequest, v2.SetTrafficLightPhaseResponse](
			httpClient,
			baseURL+TrafficLightServiceSetTrafficLightPhaseProcedure,
			connect.WithSchema(trafficLightServiceMethods.ByName("SetTrafficLightPhase")),
			connect.WithClientOptions(opts...),
		),
		setTrafficLightStatus: connect.NewClient[v2.SetTrafficLightStatusRequest, v2.SetTrafficLightStatusResponse](
			httpClient,
			baseURL+TrafficLightServiceSetTrafficLightStatusProcedure,
			connect.WithSchema(trafficLightServiceMethods.ByName("SetTrafficLightStatus")),
			connect.WithClientOptions(opts...),
		),
	}
}

// trafficLightServiceClient implements TrafficLightServiceClient.
type trafficLightServiceClient struct {
	getTrafficLight       *connect.Client[v2.GetTrafficLightRequest, v2.GetTrafficLightResponse]
	setTrafficLight       *connect.Client[v2.SetTrafficLightRequest, v2.SetTrafficLightResponse]
	setTrafficLightPhase  *connect.Client[v2.SetTrafficLightPhaseRequest, v2.SetTrafficLightPhaseResponse]
	setTrafficLightStatus *connect.Client[v2.SetTrafficLightStatusRequest, v2.SetTrafficLightStatusResponse]
}

// GetTrafficLight calls city.map.v2.TrafficLightService.GetTrafficLight.
func (c *trafficLightServiceClient) GetTrafficLight(ctx context.Context, req *connect.Request[v2.GetTrafficLightRequest]) (*connect.Response[v2.GetTrafficLightResponse], error) {
	return c.getTrafficLight.CallUnary(ctx, req)
}

// SetTrafficLight calls city.map.v2.TrafficLightService.SetTrafficLight.
func (c *trafficLightServiceClient) SetTrafficLight(ctx context.Context, req *connect.Request[v2.SetTrafficLightRequest]) (*connect.Response[v2.SetTrafficLightResponse], error) {
	return c.setTrafficLight.CallUnary(ctx, req)
}

// SetTrafficLightPhase calls city.map.v2.TrafficLightService.SetTrafficLightPhase.
func (c *trafficLightServiceClient) SetTrafficLightPhase(ctx context.Context, req *connect.Request[v2.SetTrafficLightPhaseRequest]) (*connect.Response[v2.SetTrafficLightPhaseResponse], error) {
	return c.setTrafficLightPhase.CallUnary(ctx, req)
}

// SetTrafficLightStatus calls city.map.v2.TrafficLightService.SetTrafficLightStatus.
func (c *trafficLightServiceClient) SetTrafficLightStatus(ctx context.Context, req *connect.Request[v2.SetTrafficLightStatusRequest]) (*connect.Response[v2.SetTrafficLightStatusResponse], error) {
	return c.setTrafficLightStatus.CallUnary(ctx, req)
}

// TrafficLightServiceHandler is an implementation of the city.map.v2.TrafficLightService service.
type TrafficLightServiceHandler interface {
	// 获取路口的红绿灯信息
	// Get traffic light information
	GetTrafficLight(context.Context, *connect.Request[v2.GetTrafficLightRequest]) (*connect.Response[v2.GetTrafficLightResponse], error)
	// 设置路口的红绿灯信息
	// Set traffic light information
	SetTrafficLight(context.Context, *connect.Request[v2.SetTrafficLightRequest]) (*connect.Response[v2.SetTrafficLightResponse], error)
	// 设置路口的红绿灯相位
	// Set traffic light phase
	SetTrafficLightPhase(context.Context, *connect.Request[v2.SetTrafficLightPhaseRequest]) (*connect.Response[v2.SetTrafficLightPhaseResponse], error)
	// 设置路口的红绿灯状态
	// Set traffic light status
	SetTrafficLightStatus(context.Context, *connect.Request[v2.SetTrafficLightStatusRequest]) (*connect.Response[v2.SetTrafficLightStatusResponse], error)
}

// NewTrafficLightServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTrafficLightServiceHandler(svc TrafficLightServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	trafficLightServiceMethods := v2.File_city_map_v2_traffic_light_service_proto.Services().ByName("TrafficLightService").Methods()
	trafficLightServiceGetTrafficLightHandler := connect.NewUnaryHandler(
		TrafficLightServiceGetTrafficLightProcedure,
		svc.GetTrafficLight,
		connect.WithSchema(trafficLightServiceMethods.ByName("GetTrafficLight")),
		connect.WithHandlerOptions(opts...),
	)
	trafficLightServiceSetTrafficLightHandler := connect.NewUnaryHandler(
		TrafficLightServiceSetTrafficLightProcedure,
		svc.SetTrafficLight,
		connect.WithSchema(trafficLightServiceMethods.ByName("SetTrafficLight")),
		connect.WithHandlerOptions(opts...),
	)
	trafficLightServiceSetTrafficLightPhaseHandler := connect.NewUnaryHandler(
		TrafficLightServiceSetTrafficLightPhaseProcedure,
		svc.SetTrafficLightPhase,
		connect.WithSchema(trafficLightServiceMethods.ByName("SetTrafficLightPhase")),
		connect.WithHandlerOptions(opts...),
	)
	trafficLightServiceSetTrafficLightStatusHandler := connect.NewUnaryHandler(
		TrafficLightServiceSetTrafficLightStatusProcedure,
		svc.SetTrafficLightStatus,
		connect.WithSchema(trafficLightServiceMethods.ByName("SetTrafficLightStatus")),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.map.v2.TrafficLightService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TrafficLightServiceGetTrafficLightProcedure:
			trafficLightServiceGetTrafficLightHandler.ServeHTTP(w, r)
		case TrafficLightServiceSetTrafficLightProcedure:
			trafficLightServiceSetTrafficLightHandler.ServeHTTP(w, r)
		case TrafficLightServiceSetTrafficLightPhaseProcedure:
			trafficLightServiceSetTrafficLightPhaseHandler.ServeHTTP(w, r)
		case TrafficLightServiceSetTrafficLightStatusProcedure:
			trafficLightServiceSetTrafficLightStatusHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTrafficLightServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTrafficLightServiceHandler struct{}

func (UnimplementedTrafficLightServiceHandler) GetTrafficLight(context.Context, *connect.Request[v2.GetTrafficLightRequest]) (*connect.Response[v2.GetTrafficLightResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.map.v2.TrafficLightService.GetTrafficLight is not implemented"))
}

func (UnimplementedTrafficLightServiceHandler) SetTrafficLight(context.Context, *connect.Request[v2.SetTrafficLightRequest]) (*connect.Response[v2.SetTrafficLightResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.map.v2.TrafficLightService.SetTrafficLight is not implemented"))
}

func (UnimplementedTrafficLightServiceHandler) SetTrafficLightPhase(context.Context, *connect.Request[v2.SetTrafficLightPhaseRequest]) (*connect.Response[v2.SetTrafficLightPhaseResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.map.v2.TrafficLightService.SetTrafficLightPhase is not implemented"))
}

func (UnimplementedTrafficLightServiceHandler) SetTrafficLightStatus(context.Context, *connect.Request[v2.SetTrafficLightStatusRequest]) (*connect.Response[v2.SetTrafficLightStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.map.v2.TrafficLightService.SetTrafficLightStatus is not implemented"))
}
