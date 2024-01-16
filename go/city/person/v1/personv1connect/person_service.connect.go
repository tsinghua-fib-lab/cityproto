// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/person/v1/person_service.proto

package personv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "git.fiblab.net/sim/protos/go/city/person/v1"
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
	// PersonServiceName is the fully-qualified name of the PersonService service.
	PersonServiceName = "city.person.v1.PersonService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PersonServiceGetPersonProcedure is the fully-qualified name of the PersonService's GetPerson RPC.
	PersonServiceGetPersonProcedure = "/city.person.v1.PersonService/GetPerson"
	// PersonServiceAddPersonProcedure is the fully-qualified name of the PersonService's AddPerson RPC.
	PersonServiceAddPersonProcedure = "/city.person.v1.PersonService/AddPerson"
	// PersonServiceSetScheduleProcedure is the fully-qualified name of the PersonService's SetSchedule
	// RPC.
	PersonServiceSetScheduleProcedure = "/city.person.v1.PersonService/SetSchedule"
	// PersonServiceGetPersonsByLongLatAreaProcedure is the fully-qualified name of the PersonService's
	// GetPersonsByLongLatArea RPC.
	PersonServiceGetPersonsByLongLatAreaProcedure = "/city.person.v1.PersonService/GetPersonsByLongLatArea"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	personServiceServiceDescriptor                       = v1.File_city_person_v1_person_service_proto.Services().ByName("PersonService")
	personServiceGetPersonMethodDescriptor               = personServiceServiceDescriptor.Methods().ByName("GetPerson")
	personServiceAddPersonMethodDescriptor               = personServiceServiceDescriptor.Methods().ByName("AddPerson")
	personServiceSetScheduleMethodDescriptor             = personServiceServiceDescriptor.Methods().ByName("SetSchedule")
	personServiceGetPersonsByLongLatAreaMethodDescriptor = personServiceServiceDescriptor.Methods().ByName("GetPersonsByLongLatArea")
)

// PersonServiceClient is a client for the city.person.v1.PersonService service.
type PersonServiceClient interface {
	// 获取person信息
	GetPerson(context.Context, *connect.Request[v1.GetPersonRequest]) (*connect.Response[v1.GetPersonResponse], error)
	// 新增person 传入person初始位置、目的地表、属性 返回personid
	AddPerson(context.Context, *connect.Request[v1.AddPersonRequest]) (*connect.Response[v1.AddPersonResponse], error)
	// 修改person的schedule 传入personid、目的地表
	SetSchedule(context.Context, *connect.Request[v1.SetScheduleRequest]) (*connect.Response[v1.SetScheduleResponse], error)
	// 获取特定区域内的person
	GetPersonsByLongLatArea(context.Context, *connect.Request[v1.GetPersonsByLongLatAreaRequest]) (*connect.Response[v1.GetPersonsByLongLatAreaResponse], error)
}

// NewPersonServiceClient constructs a client for the city.person.v1.PersonService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPersonServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PersonServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &personServiceClient{
		getPerson: connect.NewClient[v1.GetPersonRequest, v1.GetPersonResponse](
			httpClient,
			baseURL+PersonServiceGetPersonProcedure,
			connect.WithSchema(personServiceGetPersonMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		addPerson: connect.NewClient[v1.AddPersonRequest, v1.AddPersonResponse](
			httpClient,
			baseURL+PersonServiceAddPersonProcedure,
			connect.WithSchema(personServiceAddPersonMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setSchedule: connect.NewClient[v1.SetScheduleRequest, v1.SetScheduleResponse](
			httpClient,
			baseURL+PersonServiceSetScheduleProcedure,
			connect.WithSchema(personServiceSetScheduleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPersonsByLongLatArea: connect.NewClient[v1.GetPersonsByLongLatAreaRequest, v1.GetPersonsByLongLatAreaResponse](
			httpClient,
			baseURL+PersonServiceGetPersonsByLongLatAreaProcedure,
			connect.WithSchema(personServiceGetPersonsByLongLatAreaMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// personServiceClient implements PersonServiceClient.
type personServiceClient struct {
	getPerson               *connect.Client[v1.GetPersonRequest, v1.GetPersonResponse]
	addPerson               *connect.Client[v1.AddPersonRequest, v1.AddPersonResponse]
	setSchedule             *connect.Client[v1.SetScheduleRequest, v1.SetScheduleResponse]
	getPersonsByLongLatArea *connect.Client[v1.GetPersonsByLongLatAreaRequest, v1.GetPersonsByLongLatAreaResponse]
}

// GetPerson calls city.person.v1.PersonService.GetPerson.
func (c *personServiceClient) GetPerson(ctx context.Context, req *connect.Request[v1.GetPersonRequest]) (*connect.Response[v1.GetPersonResponse], error) {
	return c.getPerson.CallUnary(ctx, req)
}

// AddPerson calls city.person.v1.PersonService.AddPerson.
func (c *personServiceClient) AddPerson(ctx context.Context, req *connect.Request[v1.AddPersonRequest]) (*connect.Response[v1.AddPersonResponse], error) {
	return c.addPerson.CallUnary(ctx, req)
}

// SetSchedule calls city.person.v1.PersonService.SetSchedule.
func (c *personServiceClient) SetSchedule(ctx context.Context, req *connect.Request[v1.SetScheduleRequest]) (*connect.Response[v1.SetScheduleResponse], error) {
	return c.setSchedule.CallUnary(ctx, req)
}

// GetPersonsByLongLatArea calls city.person.v1.PersonService.GetPersonsByLongLatArea.
func (c *personServiceClient) GetPersonsByLongLatArea(ctx context.Context, req *connect.Request[v1.GetPersonsByLongLatAreaRequest]) (*connect.Response[v1.GetPersonsByLongLatAreaResponse], error) {
	return c.getPersonsByLongLatArea.CallUnary(ctx, req)
}

// PersonServiceHandler is an implementation of the city.person.v1.PersonService service.
type PersonServiceHandler interface {
	// 获取person信息
	GetPerson(context.Context, *connect.Request[v1.GetPersonRequest]) (*connect.Response[v1.GetPersonResponse], error)
	// 新增person 传入person初始位置、目的地表、属性 返回personid
	AddPerson(context.Context, *connect.Request[v1.AddPersonRequest]) (*connect.Response[v1.AddPersonResponse], error)
	// 修改person的schedule 传入personid、目的地表
	SetSchedule(context.Context, *connect.Request[v1.SetScheduleRequest]) (*connect.Response[v1.SetScheduleResponse], error)
	// 获取特定区域内的person
	GetPersonsByLongLatArea(context.Context, *connect.Request[v1.GetPersonsByLongLatAreaRequest]) (*connect.Response[v1.GetPersonsByLongLatAreaResponse], error)
}

// NewPersonServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPersonServiceHandler(svc PersonServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	personServiceGetPersonHandler := connect.NewUnaryHandler(
		PersonServiceGetPersonProcedure,
		svc.GetPerson,
		connect.WithSchema(personServiceGetPersonMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	personServiceAddPersonHandler := connect.NewUnaryHandler(
		PersonServiceAddPersonProcedure,
		svc.AddPerson,
		connect.WithSchema(personServiceAddPersonMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	personServiceSetScheduleHandler := connect.NewUnaryHandler(
		PersonServiceSetScheduleProcedure,
		svc.SetSchedule,
		connect.WithSchema(personServiceSetScheduleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	personServiceGetPersonsByLongLatAreaHandler := connect.NewUnaryHandler(
		PersonServiceGetPersonsByLongLatAreaProcedure,
		svc.GetPersonsByLongLatArea,
		connect.WithSchema(personServiceGetPersonsByLongLatAreaMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/city.person.v1.PersonService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PersonServiceGetPersonProcedure:
			personServiceGetPersonHandler.ServeHTTP(w, r)
		case PersonServiceAddPersonProcedure:
			personServiceAddPersonHandler.ServeHTTP(w, r)
		case PersonServiceSetScheduleProcedure:
			personServiceSetScheduleHandler.ServeHTTP(w, r)
		case PersonServiceGetPersonsByLongLatAreaProcedure:
			personServiceGetPersonsByLongLatAreaHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPersonServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPersonServiceHandler struct{}

func (UnimplementedPersonServiceHandler) GetPerson(context.Context, *connect.Request[v1.GetPersonRequest]) (*connect.Response[v1.GetPersonResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.GetPerson is not implemented"))
}

func (UnimplementedPersonServiceHandler) AddPerson(context.Context, *connect.Request[v1.AddPersonRequest]) (*connect.Response[v1.AddPersonResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.AddPerson is not implemented"))
}

func (UnimplementedPersonServiceHandler) SetSchedule(context.Context, *connect.Request[v1.SetScheduleRequest]) (*connect.Response[v1.SetScheduleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.SetSchedule is not implemented"))
}

func (UnimplementedPersonServiceHandler) GetPersonsByLongLatArea(context.Context, *connect.Request[v1.GetPersonsByLongLatAreaRequest]) (*connect.Response[v1.GetPersonsByLongLatAreaResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.GetPersonsByLongLatArea is not implemented"))
}
