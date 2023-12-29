// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/elec/interaction/v1/elec_service.proto

package interactionv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "git.fiblab.net/sim/protos/go/city/elec/interaction/v1"
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
	// ElecServiceName is the fully-qualified name of the ElecService service.
	ElecServiceName = "city.elec.interaction.v1.ElecService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ElecServiceSetStatusProcedure is the fully-qualified name of the ElecService's SetStatus RPC.
	ElecServiceSetStatusProcedure = "/city.elec.interaction.v1.ElecService/SetStatus"
	// ElecServiceGetPowerProcedure is the fully-qualified name of the ElecService's GetPower RPC.
	ElecServiceGetPowerProcedure = "/city.elec.interaction.v1.ElecService/GetPower"
	// ElecServiceGetPowerStatusProcedure is the fully-qualified name of the ElecService's
	// GetPowerStatus RPC.
	ElecServiceGetPowerStatusProcedure = "/city.elec.interaction.v1.ElecService/GetPowerStatus"
	// ElecServiceGetNoPowerAOIProcedure is the fully-qualified name of the ElecService's GetNoPowerAOI
	// RPC.
	ElecServiceGetNoPowerAOIProcedure = "/city.elec.interaction.v1.ElecService/GetNoPowerAOI"
	// ElecServiceGetRuinInfoProcedure is the fully-qualified name of the ElecService's GetRuinInfo RPC.
	ElecServiceGetRuinInfoProcedure = "/city.elec.interaction.v1.ElecService/GetRuinInfo"
	// ElecServiceGetEdgeStatusProcedure is the fully-qualified name of the ElecService's GetEdgeStatus
	// RPC.
	ElecServiceGetEdgeStatusProcedure = "/city.elec.interaction.v1.ElecService/GetEdgeStatus"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	elecServiceServiceDescriptor              = v1.File_city_elec_interaction_v1_elec_service_proto.Services().ByName("ElecService")
	elecServiceSetStatusMethodDescriptor      = elecServiceServiceDescriptor.Methods().ByName("SetStatus")
	elecServiceGetPowerMethodDescriptor       = elecServiceServiceDescriptor.Methods().ByName("GetPower")
	elecServiceGetPowerStatusMethodDescriptor = elecServiceServiceDescriptor.Methods().ByName("GetPowerStatus")
	elecServiceGetNoPowerAOIMethodDescriptor  = elecServiceServiceDescriptor.Methods().ByName("GetNoPowerAOI")
	elecServiceGetRuinInfoMethodDescriptor    = elecServiceServiceDescriptor.Methods().ByName("GetRuinInfo")
	elecServiceGetEdgeStatusMethodDescriptor  = elecServiceServiceDescriptor.Methods().ByName("GetEdgeStatus")
)

// ElecServiceClient is a client for the city.elec.interaction.v1.ElecService service.
type ElecServiceClient interface {
	SetStatus(context.Context, *connect.Request[v1.SetStatusRequest]) (*connect.Response[v1.SetStatusResponse], error)
	GetPower(context.Context, *connect.Request[v1.GetPowerRequest]) (*connect.Response[v1.GetPowerResponse], error)
	GetPowerStatus(context.Context, *connect.Request[v1.GetPowerStatusRequest]) (*connect.Response[v1.GetPowerStatusResponse], error)
	GetNoPowerAOI(context.Context, *connect.Request[v1.GetNoPowerAOIRequest]) (*connect.Response[v1.GetNoPowerAOIResponse], error)
	GetRuinInfo(context.Context, *connect.Request[v1.GetRuinInfoRequest]) (*connect.Response[v1.GetRuinInfoResponse], error)
	GetEdgeStatus(context.Context, *connect.Request[v1.GetEdgeStatusRequest]) (*connect.Response[v1.GetEdgeStatusResponse], error)
}

// NewElecServiceClient constructs a client for the city.elec.interaction.v1.ElecService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewElecServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ElecServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &elecServiceClient{
		setStatus: connect.NewClient[v1.SetStatusRequest, v1.SetStatusResponse](
			httpClient,
			baseURL+ElecServiceSetStatusProcedure,
			connect.WithSchema(elecServiceSetStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPower: connect.NewClient[v1.GetPowerRequest, v1.GetPowerResponse](
			httpClient,
			baseURL+ElecServiceGetPowerProcedure,
			connect.WithSchema(elecServiceGetPowerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPowerStatus: connect.NewClient[v1.GetPowerStatusRequest, v1.GetPowerStatusResponse](
			httpClient,
			baseURL+ElecServiceGetPowerStatusProcedure,
			connect.WithSchema(elecServiceGetPowerStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getNoPowerAOI: connect.NewClient[v1.GetNoPowerAOIRequest, v1.GetNoPowerAOIResponse](
			httpClient,
			baseURL+ElecServiceGetNoPowerAOIProcedure,
			connect.WithSchema(elecServiceGetNoPowerAOIMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getRuinInfo: connect.NewClient[v1.GetRuinInfoRequest, v1.GetRuinInfoResponse](
			httpClient,
			baseURL+ElecServiceGetRuinInfoProcedure,
			connect.WithSchema(elecServiceGetRuinInfoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getEdgeStatus: connect.NewClient[v1.GetEdgeStatusRequest, v1.GetEdgeStatusResponse](
			httpClient,
			baseURL+ElecServiceGetEdgeStatusProcedure,
			connect.WithSchema(elecServiceGetEdgeStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// elecServiceClient implements ElecServiceClient.
type elecServiceClient struct {
	setStatus      *connect.Client[v1.SetStatusRequest, v1.SetStatusResponse]
	getPower       *connect.Client[v1.GetPowerRequest, v1.GetPowerResponse]
	getPowerStatus *connect.Client[v1.GetPowerStatusRequest, v1.GetPowerStatusResponse]
	getNoPowerAOI  *connect.Client[v1.GetNoPowerAOIRequest, v1.GetNoPowerAOIResponse]
	getRuinInfo    *connect.Client[v1.GetRuinInfoRequest, v1.GetRuinInfoResponse]
	getEdgeStatus  *connect.Client[v1.GetEdgeStatusRequest, v1.GetEdgeStatusResponse]
}

// SetStatus calls city.elec.interaction.v1.ElecService.SetStatus.
func (c *elecServiceClient) SetStatus(ctx context.Context, req *connect.Request[v1.SetStatusRequest]) (*connect.Response[v1.SetStatusResponse], error) {
	return c.setStatus.CallUnary(ctx, req)
}

// GetPower calls city.elec.interaction.v1.ElecService.GetPower.
func (c *elecServiceClient) GetPower(ctx context.Context, req *connect.Request[v1.GetPowerRequest]) (*connect.Response[v1.GetPowerResponse], error) {
	return c.getPower.CallUnary(ctx, req)
}

// GetPowerStatus calls city.elec.interaction.v1.ElecService.GetPowerStatus.
func (c *elecServiceClient) GetPowerStatus(ctx context.Context, req *connect.Request[v1.GetPowerStatusRequest]) (*connect.Response[v1.GetPowerStatusResponse], error) {
	return c.getPowerStatus.CallUnary(ctx, req)
}

// GetNoPowerAOI calls city.elec.interaction.v1.ElecService.GetNoPowerAOI.
func (c *elecServiceClient) GetNoPowerAOI(ctx context.Context, req *connect.Request[v1.GetNoPowerAOIRequest]) (*connect.Response[v1.GetNoPowerAOIResponse], error) {
	return c.getNoPowerAOI.CallUnary(ctx, req)
}

// GetRuinInfo calls city.elec.interaction.v1.ElecService.GetRuinInfo.
func (c *elecServiceClient) GetRuinInfo(ctx context.Context, req *connect.Request[v1.GetRuinInfoRequest]) (*connect.Response[v1.GetRuinInfoResponse], error) {
	return c.getRuinInfo.CallUnary(ctx, req)
}

// GetEdgeStatus calls city.elec.interaction.v1.ElecService.GetEdgeStatus.
func (c *elecServiceClient) GetEdgeStatus(ctx context.Context, req *connect.Request[v1.GetEdgeStatusRequest]) (*connect.Response[v1.GetEdgeStatusResponse], error) {
	return c.getEdgeStatus.CallUnary(ctx, req)
}

// ElecServiceHandler is an implementation of the city.elec.interaction.v1.ElecService service.
type ElecServiceHandler interface {
	SetStatus(context.Context, *connect.Request[v1.SetStatusRequest]) (*connect.Response[v1.SetStatusResponse], error)
	GetPower(context.Context, *connect.Request[v1.GetPowerRequest]) (*connect.Response[v1.GetPowerResponse], error)
	GetPowerStatus(context.Context, *connect.Request[v1.GetPowerStatusRequest]) (*connect.Response[v1.GetPowerStatusResponse], error)
	GetNoPowerAOI(context.Context, *connect.Request[v1.GetNoPowerAOIRequest]) (*connect.Response[v1.GetNoPowerAOIResponse], error)
	GetRuinInfo(context.Context, *connect.Request[v1.GetRuinInfoRequest]) (*connect.Response[v1.GetRuinInfoResponse], error)
	GetEdgeStatus(context.Context, *connect.Request[v1.GetEdgeStatusRequest]) (*connect.Response[v1.GetEdgeStatusResponse], error)
}

// NewElecServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewElecServiceHandler(svc ElecServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	elecServiceSetStatusHandler := connect.NewUnaryHandler(
		ElecServiceSetStatusProcedure,
		svc.SetStatus,
		connect.WithSchema(elecServiceSetStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	elecServiceGetPowerHandler := connect.NewUnaryHandler(
		ElecServiceGetPowerProcedure,
		svc.GetPower,
		connect.WithSchema(elecServiceGetPowerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	elecServiceGetPowerStatusHandler := connect.NewUnaryHandler(
		ElecServiceGetPowerStatusProcedure,
		svc.GetPowerStatus,
		connect.WithSchema(elecServiceGetPowerStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	elecServiceGetNoPowerAOIHandler := connect.NewUnaryHandler(
		ElecServiceGetNoPowerAOIProcedure,
		svc.GetNoPowerAOI,
		connect.WithSchema(elecServiceGetNoPowerAOIMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	elecServiceGetRuinInfoHandler := connect.NewUnaryHandler(
		ElecServiceGetRuinInfoProcedure,
		svc.GetRuinInfo,
		connect.WithSchema(elecServiceGetRuinInfoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	elecServiceGetEdgeStatusHandler := connect.NewUnaryHandler(
		ElecServiceGetEdgeStatusProcedure,
		svc.GetEdgeStatus,
		connect.WithSchema(elecServiceGetEdgeStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.elec.interaction.v1.ElecService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ElecServiceSetStatusProcedure:
			elecServiceSetStatusHandler.ServeHTTP(w, r)
		case ElecServiceGetPowerProcedure:
			elecServiceGetPowerHandler.ServeHTTP(w, r)
		case ElecServiceGetPowerStatusProcedure:
			elecServiceGetPowerStatusHandler.ServeHTTP(w, r)
		case ElecServiceGetNoPowerAOIProcedure:
			elecServiceGetNoPowerAOIHandler.ServeHTTP(w, r)
		case ElecServiceGetRuinInfoProcedure:
			elecServiceGetRuinInfoHandler.ServeHTTP(w, r)
		case ElecServiceGetEdgeStatusProcedure:
			elecServiceGetEdgeStatusHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedElecServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedElecServiceHandler struct{}

func (UnimplementedElecServiceHandler) SetStatus(context.Context, *connect.Request[v1.SetStatusRequest]) (*connect.Response[v1.SetStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.elec.interaction.v1.ElecService.SetStatus is not implemented"))
}

func (UnimplementedElecServiceHandler) GetPower(context.Context, *connect.Request[v1.GetPowerRequest]) (*connect.Response[v1.GetPowerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.elec.interaction.v1.ElecService.GetPower is not implemented"))
}

func (UnimplementedElecServiceHandler) GetPowerStatus(context.Context, *connect.Request[v1.GetPowerStatusRequest]) (*connect.Response[v1.GetPowerStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.elec.interaction.v1.ElecService.GetPowerStatus is not implemented"))
}

func (UnimplementedElecServiceHandler) GetNoPowerAOI(context.Context, *connect.Request[v1.GetNoPowerAOIRequest]) (*connect.Response[v1.GetNoPowerAOIResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.elec.interaction.v1.ElecService.GetNoPowerAOI is not implemented"))
}

func (UnimplementedElecServiceHandler) GetRuinInfo(context.Context, *connect.Request[v1.GetRuinInfoRequest]) (*connect.Response[v1.GetRuinInfoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.elec.interaction.v1.ElecService.GetRuinInfo is not implemented"))
}

func (UnimplementedElecServiceHandler) GetEdgeStatus(context.Context, *connect.Request[v1.GetEdgeStatusRequest]) (*connect.Response[v1.GetEdgeStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.elec.interaction.v1.ElecService.GetEdgeStatus is not implemented"))
}
