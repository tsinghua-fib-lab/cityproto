// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/map/v2/junction_service.proto

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
	// JunctionServiceName is the fully-qualified name of the JunctionService service.
	JunctionServiceName = "city.map.v2.JunctionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// JunctionServiceGetJunctionProcedure is the fully-qualified name of the JunctionService's
	// GetJunction RPC.
	JunctionServiceGetJunctionProcedure = "/city.map.v2.JunctionService/GetJunction"
)

// JunctionServiceClient is a client for the city.map.v2.JunctionService service.
type JunctionServiceClient interface {
	// 查询路口信息
	// Get junction information
	GetJunction(context.Context, *connect.Request[v2.GetJunctionRequest]) (*connect.Response[v2.GetJunctionResponse], error)
}

// NewJunctionServiceClient constructs a client for the city.map.v2.JunctionService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewJunctionServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) JunctionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	junctionServiceMethods := v2.File_city_map_v2_junction_service_proto.Services().ByName("JunctionService").Methods()
	return &junctionServiceClient{
		getJunction: connect.NewClient[v2.GetJunctionRequest, v2.GetJunctionResponse](
			httpClient,
			baseURL+JunctionServiceGetJunctionProcedure,
			connect.WithSchema(junctionServiceMethods.ByName("GetJunction")),
			connect.WithClientOptions(opts...),
		),
	}
}

// junctionServiceClient implements JunctionServiceClient.
type junctionServiceClient struct {
	getJunction *connect.Client[v2.GetJunctionRequest, v2.GetJunctionResponse]
}

// GetJunction calls city.map.v2.JunctionService.GetJunction.
func (c *junctionServiceClient) GetJunction(ctx context.Context, req *connect.Request[v2.GetJunctionRequest]) (*connect.Response[v2.GetJunctionResponse], error) {
	return c.getJunction.CallUnary(ctx, req)
}

// JunctionServiceHandler is an implementation of the city.map.v2.JunctionService service.
type JunctionServiceHandler interface {
	// 查询路口信息
	// Get junction information
	GetJunction(context.Context, *connect.Request[v2.GetJunctionRequest]) (*connect.Response[v2.GetJunctionResponse], error)
}

// NewJunctionServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewJunctionServiceHandler(svc JunctionServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	junctionServiceMethods := v2.File_city_map_v2_junction_service_proto.Services().ByName("JunctionService").Methods()
	junctionServiceGetJunctionHandler := connect.NewUnaryHandler(
		JunctionServiceGetJunctionProcedure,
		svc.GetJunction,
		connect.WithSchema(junctionServiceMethods.ByName("GetJunction")),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.map.v2.JunctionService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case JunctionServiceGetJunctionProcedure:
			junctionServiceGetJunctionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedJunctionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedJunctionServiceHandler struct{}

func (UnimplementedJunctionServiceHandler) GetJunction(context.Context, *connect.Request[v2.GetJunctionRequest]) (*connect.Response[v2.GetJunctionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.map.v2.JunctionService.GetJunction is not implemented"))
}
