// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/economy/v1/org_service.proto

package economyv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "git.fiblab.net/sim/protos/go/city/economy/v1"
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
	// OrgServiceName is the fully-qualified name of the OrgService service.
	OrgServiceName = "city.economy.v1.OrgService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OrgServiceGetOrgProcedure is the fully-qualified name of the OrgService's GetOrg RPC.
	OrgServiceGetOrgProcedure = "/city.economy.v1.OrgService/GetOrg"
	// OrgServiceUpdateOrgMoneyProcedure is the fully-qualified name of the OrgService's UpdateOrgMoney
	// RPC.
	OrgServiceUpdateOrgMoneyProcedure = "/city.economy.v1.OrgService/UpdateOrgMoney"
	// OrgServiceUpdateOrgGoodsProcedure is the fully-qualified name of the OrgService's UpdateOrgGoods
	// RPC.
	OrgServiceUpdateOrgGoodsProcedure = "/city.economy.v1.OrgService/UpdateOrgGoods"
	// OrgServiceUpdateOrgEmployeeProcedure is the fully-qualified name of the OrgService's
	// UpdateOrgEmployee RPC.
	OrgServiceUpdateOrgEmployeeProcedure = "/city.economy.v1.OrgService/UpdateOrgEmployee"
	// OrgServiceUpdateOrgJobProcedure is the fully-qualified name of the OrgService's UpdateOrgJob RPC.
	OrgServiceUpdateOrgJobProcedure = "/city.economy.v1.OrgService/UpdateOrgJob"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	orgServiceServiceDescriptor                 = v1.File_city_economy_v1_org_service_proto.Services().ByName("OrgService")
	orgServiceGetOrgMethodDescriptor            = orgServiceServiceDescriptor.Methods().ByName("GetOrg")
	orgServiceUpdateOrgMoneyMethodDescriptor    = orgServiceServiceDescriptor.Methods().ByName("UpdateOrgMoney")
	orgServiceUpdateOrgGoodsMethodDescriptor    = orgServiceServiceDescriptor.Methods().ByName("UpdateOrgGoods")
	orgServiceUpdateOrgEmployeeMethodDescriptor = orgServiceServiceDescriptor.Methods().ByName("UpdateOrgEmployee")
	orgServiceUpdateOrgJobMethodDescriptor      = orgServiceServiceDescriptor.Methods().ByName("UpdateOrgJob")
)

// OrgServiceClient is a client for the city.economy.v1.OrgService service.
type OrgServiceClient interface {
	// 批量查询组织的经济情况（员工、岗位、资金、货物）
	GetOrg(context.Context, *connect.Request[v1.GetOrgRequest]) (*connect.Response[v1.GetOrgResponse], error)
	// 批量修改组织的资金
	UpdateOrgMoney(context.Context, *connect.Request[v1.UpdateOrgMoneyRequest]) (*connect.Response[v1.UpdateOrgMoneyResponse], error)
	// 批量修改组织的货物
	UpdateOrgGoods(context.Context, *connect.Request[v1.UpdateOrgGoodsRequest]) (*connect.Response[v1.UpdateOrgGoodsResponse], error)
	// 批量修改组织的员工
	UpdateOrgEmployee(context.Context, *connect.Request[v1.UpdateOrgEmployeeRequest]) (*connect.Response[v1.UpdateOrgEmployeeResponse], error)
	// 批量修改组织的岗位
	UpdateOrgJob(context.Context, *connect.Request[v1.UpdateOrgJobRequest]) (*connect.Response[v1.UpdateOrgJobResponse], error)
}

// NewOrgServiceClient constructs a client for the city.economy.v1.OrgService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrgServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OrgServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &orgServiceClient{
		getOrg: connect.NewClient[v1.GetOrgRequest, v1.GetOrgResponse](
			httpClient,
			baseURL+OrgServiceGetOrgProcedure,
			connect.WithSchema(orgServiceGetOrgMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateOrgMoney: connect.NewClient[v1.UpdateOrgMoneyRequest, v1.UpdateOrgMoneyResponse](
			httpClient,
			baseURL+OrgServiceUpdateOrgMoneyProcedure,
			connect.WithSchema(orgServiceUpdateOrgMoneyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateOrgGoods: connect.NewClient[v1.UpdateOrgGoodsRequest, v1.UpdateOrgGoodsResponse](
			httpClient,
			baseURL+OrgServiceUpdateOrgGoodsProcedure,
			connect.WithSchema(orgServiceUpdateOrgGoodsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateOrgEmployee: connect.NewClient[v1.UpdateOrgEmployeeRequest, v1.UpdateOrgEmployeeResponse](
			httpClient,
			baseURL+OrgServiceUpdateOrgEmployeeProcedure,
			connect.WithSchema(orgServiceUpdateOrgEmployeeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateOrgJob: connect.NewClient[v1.UpdateOrgJobRequest, v1.UpdateOrgJobResponse](
			httpClient,
			baseURL+OrgServiceUpdateOrgJobProcedure,
			connect.WithSchema(orgServiceUpdateOrgJobMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// orgServiceClient implements OrgServiceClient.
type orgServiceClient struct {
	getOrg            *connect.Client[v1.GetOrgRequest, v1.GetOrgResponse]
	updateOrgMoney    *connect.Client[v1.UpdateOrgMoneyRequest, v1.UpdateOrgMoneyResponse]
	updateOrgGoods    *connect.Client[v1.UpdateOrgGoodsRequest, v1.UpdateOrgGoodsResponse]
	updateOrgEmployee *connect.Client[v1.UpdateOrgEmployeeRequest, v1.UpdateOrgEmployeeResponse]
	updateOrgJob      *connect.Client[v1.UpdateOrgJobRequest, v1.UpdateOrgJobResponse]
}

// GetOrg calls city.economy.v1.OrgService.GetOrg.
func (c *orgServiceClient) GetOrg(ctx context.Context, req *connect.Request[v1.GetOrgRequest]) (*connect.Response[v1.GetOrgResponse], error) {
	return c.getOrg.CallUnary(ctx, req)
}

// UpdateOrgMoney calls city.economy.v1.OrgService.UpdateOrgMoney.
func (c *orgServiceClient) UpdateOrgMoney(ctx context.Context, req *connect.Request[v1.UpdateOrgMoneyRequest]) (*connect.Response[v1.UpdateOrgMoneyResponse], error) {
	return c.updateOrgMoney.CallUnary(ctx, req)
}

// UpdateOrgGoods calls city.economy.v1.OrgService.UpdateOrgGoods.
func (c *orgServiceClient) UpdateOrgGoods(ctx context.Context, req *connect.Request[v1.UpdateOrgGoodsRequest]) (*connect.Response[v1.UpdateOrgGoodsResponse], error) {
	return c.updateOrgGoods.CallUnary(ctx, req)
}

// UpdateOrgEmployee calls city.economy.v1.OrgService.UpdateOrgEmployee.
func (c *orgServiceClient) UpdateOrgEmployee(ctx context.Context, req *connect.Request[v1.UpdateOrgEmployeeRequest]) (*connect.Response[v1.UpdateOrgEmployeeResponse], error) {
	return c.updateOrgEmployee.CallUnary(ctx, req)
}

// UpdateOrgJob calls city.economy.v1.OrgService.UpdateOrgJob.
func (c *orgServiceClient) UpdateOrgJob(ctx context.Context, req *connect.Request[v1.UpdateOrgJobRequest]) (*connect.Response[v1.UpdateOrgJobResponse], error) {
	return c.updateOrgJob.CallUnary(ctx, req)
}

// OrgServiceHandler is an implementation of the city.economy.v1.OrgService service.
type OrgServiceHandler interface {
	// 批量查询组织的经济情况（员工、岗位、资金、货物）
	GetOrg(context.Context, *connect.Request[v1.GetOrgRequest]) (*connect.Response[v1.GetOrgResponse], error)
	// 批量修改组织的资金
	UpdateOrgMoney(context.Context, *connect.Request[v1.UpdateOrgMoneyRequest]) (*connect.Response[v1.UpdateOrgMoneyResponse], error)
	// 批量修改组织的货物
	UpdateOrgGoods(context.Context, *connect.Request[v1.UpdateOrgGoodsRequest]) (*connect.Response[v1.UpdateOrgGoodsResponse], error)
	// 批量修改组织的员工
	UpdateOrgEmployee(context.Context, *connect.Request[v1.UpdateOrgEmployeeRequest]) (*connect.Response[v1.UpdateOrgEmployeeResponse], error)
	// 批量修改组织的岗位
	UpdateOrgJob(context.Context, *connect.Request[v1.UpdateOrgJobRequest]) (*connect.Response[v1.UpdateOrgJobResponse], error)
}

// NewOrgServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrgServiceHandler(svc OrgServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	orgServiceGetOrgHandler := connect.NewUnaryHandler(
		OrgServiceGetOrgProcedure,
		svc.GetOrg,
		connect.WithSchema(orgServiceGetOrgMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orgServiceUpdateOrgMoneyHandler := connect.NewUnaryHandler(
		OrgServiceUpdateOrgMoneyProcedure,
		svc.UpdateOrgMoney,
		connect.WithSchema(orgServiceUpdateOrgMoneyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orgServiceUpdateOrgGoodsHandler := connect.NewUnaryHandler(
		OrgServiceUpdateOrgGoodsProcedure,
		svc.UpdateOrgGoods,
		connect.WithSchema(orgServiceUpdateOrgGoodsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orgServiceUpdateOrgEmployeeHandler := connect.NewUnaryHandler(
		OrgServiceUpdateOrgEmployeeProcedure,
		svc.UpdateOrgEmployee,
		connect.WithSchema(orgServiceUpdateOrgEmployeeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orgServiceUpdateOrgJobHandler := connect.NewUnaryHandler(
		OrgServiceUpdateOrgJobProcedure,
		svc.UpdateOrgJob,
		connect.WithSchema(orgServiceUpdateOrgJobMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.economy.v1.OrgService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OrgServiceGetOrgProcedure:
			orgServiceGetOrgHandler.ServeHTTP(w, r)
		case OrgServiceUpdateOrgMoneyProcedure:
			orgServiceUpdateOrgMoneyHandler.ServeHTTP(w, r)
		case OrgServiceUpdateOrgGoodsProcedure:
			orgServiceUpdateOrgGoodsHandler.ServeHTTP(w, r)
		case OrgServiceUpdateOrgEmployeeProcedure:
			orgServiceUpdateOrgEmployeeHandler.ServeHTTP(w, r)
		case OrgServiceUpdateOrgJobProcedure:
			orgServiceUpdateOrgJobHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOrgServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOrgServiceHandler struct{}

func (UnimplementedOrgServiceHandler) GetOrg(context.Context, *connect.Request[v1.GetOrgRequest]) (*connect.Response[v1.GetOrgResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.economy.v1.OrgService.GetOrg is not implemented"))
}

func (UnimplementedOrgServiceHandler) UpdateOrgMoney(context.Context, *connect.Request[v1.UpdateOrgMoneyRequest]) (*connect.Response[v1.UpdateOrgMoneyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.economy.v1.OrgService.UpdateOrgMoney is not implemented"))
}

func (UnimplementedOrgServiceHandler) UpdateOrgGoods(context.Context, *connect.Request[v1.UpdateOrgGoodsRequest]) (*connect.Response[v1.UpdateOrgGoodsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.economy.v1.OrgService.UpdateOrgGoods is not implemented"))
}

func (UnimplementedOrgServiceHandler) UpdateOrgEmployee(context.Context, *connect.Request[v1.UpdateOrgEmployeeRequest]) (*connect.Response[v1.UpdateOrgEmployeeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.economy.v1.OrgService.UpdateOrgEmployee is not implemented"))
}

func (UnimplementedOrgServiceHandler) UpdateOrgJob(context.Context, *connect.Request[v1.UpdateOrgJobRequest]) (*connect.Response[v1.UpdateOrgJobResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.economy.v1.OrgService.UpdateOrgJob is not implemented"))
}
