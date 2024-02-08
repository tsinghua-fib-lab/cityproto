// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/map/v2/lane_service.proto

package mapv2connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2 "git.fiblab.net/sim/protos/go/city/map/v2"
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
	// LaneServiceName is the fully-qualified name of the LaneService service.
	LaneServiceName = "city.map.v2.LaneService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LaneServiceSetLaneMaxVProcedure is the fully-qualified name of the LaneService's SetLaneMaxV RPC.
	LaneServiceSetLaneMaxVProcedure = "/city.map.v2.LaneService/SetLaneMaxV"
	// LaneServiceGetLaneProcedure is the fully-qualified name of the LaneService's GetLane RPC.
	LaneServiceGetLaneProcedure = "/city.map.v2.LaneService/GetLane"
	// LaneServiceGetLaneByLongLatBBoxProcedure is the fully-qualified name of the LaneService's
	// GetLaneByLongLatBBox RPC.
	LaneServiceGetLaneByLongLatBBoxProcedure = "/city.map.v2.LaneService/GetLaneByLongLatBBox"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	laneServiceServiceDescriptor                    = v2.File_city_map_v2_lane_service_proto.Services().ByName("LaneService")
	laneServiceSetLaneMaxVMethodDescriptor          = laneServiceServiceDescriptor.Methods().ByName("SetLaneMaxV")
	laneServiceGetLaneMethodDescriptor              = laneServiceServiceDescriptor.Methods().ByName("GetLane")
	laneServiceGetLaneByLongLatBBoxMethodDescriptor = laneServiceServiceDescriptor.Methods().ByName("GetLaneByLongLatBBox")
)

// LaneServiceClient is a client for the city.map.v2.LaneService service.
type LaneServiceClient interface {
	// 设置Lane的最大速度（限速）
	// Set Lane's maximum speed (speed limit)
	SetLaneMaxV(context.Context, *connect.Request[v2.SetLaneMaxVRequest]) (*connect.Response[v2.SetLaneMaxVResponse], error)
	// 获取Lane的信息
	// Get Lane information
	GetLane(context.Context, *connect.Request[v2.GetLaneRequest]) (*connect.Response[v2.GetLaneResponse], error)
	// 获取特定区域内的Lane的信息
	// Get Lane information in a specific region
	GetLaneByLongLatBBox(context.Context, *connect.Request[v2.GetLaneByLongLatBBoxRequest]) (*connect.Response[v2.GetLaneByLongLatBBoxResponse], error)
}

// NewLaneServiceClient constructs a client for the city.map.v2.LaneService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLaneServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LaneServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &laneServiceClient{
		setLaneMaxV: connect.NewClient[v2.SetLaneMaxVRequest, v2.SetLaneMaxVResponse](
			httpClient,
			baseURL+LaneServiceSetLaneMaxVProcedure,
			connect.WithSchema(laneServiceSetLaneMaxVMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getLane: connect.NewClient[v2.GetLaneRequest, v2.GetLaneResponse](
			httpClient,
			baseURL+LaneServiceGetLaneProcedure,
			connect.WithSchema(laneServiceGetLaneMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getLaneByLongLatBBox: connect.NewClient[v2.GetLaneByLongLatBBoxRequest, v2.GetLaneByLongLatBBoxResponse](
			httpClient,
			baseURL+LaneServiceGetLaneByLongLatBBoxProcedure,
			connect.WithSchema(laneServiceGetLaneByLongLatBBoxMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// laneServiceClient implements LaneServiceClient.
type laneServiceClient struct {
	setLaneMaxV          *connect.Client[v2.SetLaneMaxVRequest, v2.SetLaneMaxVResponse]
	getLane              *connect.Client[v2.GetLaneRequest, v2.GetLaneResponse]
	getLaneByLongLatBBox *connect.Client[v2.GetLaneByLongLatBBoxRequest, v2.GetLaneByLongLatBBoxResponse]
}

// SetLaneMaxV calls city.map.v2.LaneService.SetLaneMaxV.
func (c *laneServiceClient) SetLaneMaxV(ctx context.Context, req *connect.Request[v2.SetLaneMaxVRequest]) (*connect.Response[v2.SetLaneMaxVResponse], error) {
	return c.setLaneMaxV.CallUnary(ctx, req)
}

// GetLane calls city.map.v2.LaneService.GetLane.
func (c *laneServiceClient) GetLane(ctx context.Context, req *connect.Request[v2.GetLaneRequest]) (*connect.Response[v2.GetLaneResponse], error) {
	return c.getLane.CallUnary(ctx, req)
}

// GetLaneByLongLatBBox calls city.map.v2.LaneService.GetLaneByLongLatBBox.
func (c *laneServiceClient) GetLaneByLongLatBBox(ctx context.Context, req *connect.Request[v2.GetLaneByLongLatBBoxRequest]) (*connect.Response[v2.GetLaneByLongLatBBoxResponse], error) {
	return c.getLaneByLongLatBBox.CallUnary(ctx, req)
}

// LaneServiceHandler is an implementation of the city.map.v2.LaneService service.
type LaneServiceHandler interface {
	// 设置Lane的最大速度（限速）
	// Set Lane's maximum speed (speed limit)
	SetLaneMaxV(context.Context, *connect.Request[v2.SetLaneMaxVRequest]) (*connect.Response[v2.SetLaneMaxVResponse], error)
	// 获取Lane的信息
	// Get Lane information
	GetLane(context.Context, *connect.Request[v2.GetLaneRequest]) (*connect.Response[v2.GetLaneResponse], error)
	// 获取特定区域内的Lane的信息
	// Get Lane information in a specific region
	GetLaneByLongLatBBox(context.Context, *connect.Request[v2.GetLaneByLongLatBBoxRequest]) (*connect.Response[v2.GetLaneByLongLatBBoxResponse], error)
}

// NewLaneServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLaneServiceHandler(svc LaneServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	laneServiceSetLaneMaxVHandler := connect.NewUnaryHandler(
		LaneServiceSetLaneMaxVProcedure,
		svc.SetLaneMaxV,
		connect.WithSchema(laneServiceSetLaneMaxVMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	laneServiceGetLaneHandler := connect.NewUnaryHandler(
		LaneServiceGetLaneProcedure,
		svc.GetLane,
		connect.WithSchema(laneServiceGetLaneMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	laneServiceGetLaneByLongLatBBoxHandler := connect.NewUnaryHandler(
		LaneServiceGetLaneByLongLatBBoxProcedure,
		svc.GetLaneByLongLatBBox,
		connect.WithSchema(laneServiceGetLaneByLongLatBBoxMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.map.v2.LaneService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LaneServiceSetLaneMaxVProcedure:
			laneServiceSetLaneMaxVHandler.ServeHTTP(w, r)
		case LaneServiceGetLaneProcedure:
			laneServiceGetLaneHandler.ServeHTTP(w, r)
		case LaneServiceGetLaneByLongLatBBoxProcedure:
			laneServiceGetLaneByLongLatBBoxHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLaneServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLaneServiceHandler struct{}

func (UnimplementedLaneServiceHandler) SetLaneMaxV(context.Context, *connect.Request[v2.SetLaneMaxVRequest]) (*connect.Response[v2.SetLaneMaxVResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.map.v2.LaneService.SetLaneMaxV is not implemented"))
}

func (UnimplementedLaneServiceHandler) GetLane(context.Context, *connect.Request[v2.GetLaneRequest]) (*connect.Response[v2.GetLaneResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.map.v2.LaneService.GetLane is not implemented"))
}

func (UnimplementedLaneServiceHandler) GetLaneByLongLatBBox(context.Context, *connect.Request[v2.GetLaneByLongLatBBoxRequest]) (*connect.Response[v2.GetLaneByLongLatBBoxResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.map.v2.LaneService.GetLaneByLongLatBBox is not implemented"))
}
