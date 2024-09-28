// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/elec/output/v1/output_service.proto

package outputv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "git.fiblab.net/sim/protos/v2/go/city/elec/output/v1"
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
	// OutputServiceName is the fully-qualified name of the OutputService service.
	OutputServiceName = "city.elec.output.v1.OutputService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OutputServiceOutputProcedure is the fully-qualified name of the OutputService's Output RPC.
	OutputServiceOutputProcedure = "/city.elec.output.v1.OutputService/Output"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	outputServiceServiceDescriptor      = v1.File_city_elec_output_v1_output_service_proto.Services().ByName("OutputService")
	outputServiceOutputMethodDescriptor = outputServiceServiceDescriptor.Methods().ByName("Output")
)

// OutputServiceClient is a client for the city.elec.output.v1.OutputService service.
type OutputServiceClient interface {
	Output(context.Context, *connect.Request[v1.OutputRequest]) (*connect.Response[v1.OutputResponse], error)
}

// NewOutputServiceClient constructs a client for the city.elec.output.v1.OutputService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOutputServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OutputServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &outputServiceClient{
		output: connect.NewClient[v1.OutputRequest, v1.OutputResponse](
			httpClient,
			baseURL+OutputServiceOutputProcedure,
			connect.WithSchema(outputServiceOutputMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// outputServiceClient implements OutputServiceClient.
type outputServiceClient struct {
	output *connect.Client[v1.OutputRequest, v1.OutputResponse]
}

// Output calls city.elec.output.v1.OutputService.Output.
func (c *outputServiceClient) Output(ctx context.Context, req *connect.Request[v1.OutputRequest]) (*connect.Response[v1.OutputResponse], error) {
	return c.output.CallUnary(ctx, req)
}

// OutputServiceHandler is an implementation of the city.elec.output.v1.OutputService service.
type OutputServiceHandler interface {
	Output(context.Context, *connect.Request[v1.OutputRequest]) (*connect.Response[v1.OutputResponse], error)
}

// NewOutputServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOutputServiceHandler(svc OutputServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	outputServiceOutputHandler := connect.NewUnaryHandler(
		OutputServiceOutputProcedure,
		svc.Output,
		connect.WithSchema(outputServiceOutputMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.elec.output.v1.OutputService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OutputServiceOutputProcedure:
			outputServiceOutputHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOutputServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOutputServiceHandler struct{}

func (UnimplementedOutputServiceHandler) Output(context.Context, *connect.Request[v1.OutputRequest]) (*connect.Response[v1.OutputResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.elec.output.v1.OutputService.Output is not implemented"))
}
