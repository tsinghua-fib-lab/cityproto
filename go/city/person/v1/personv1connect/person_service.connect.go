// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: city/person/v1/person_service.proto

package personv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "git.fiblab.net/sim/protos/v2/go/city/person/v1"
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
	// PersonServiceGetPersonsProcedure is the fully-qualified name of the PersonService's GetPersons
	// RPC.
	PersonServiceGetPersonsProcedure = "/city.person.v1.PersonService/GetPersons"
	// PersonServiceGetPersonByLongLatBBoxProcedure is the fully-qualified name of the PersonService's
	// GetPersonByLongLatBBox RPC.
	PersonServiceGetPersonByLongLatBBoxProcedure = "/city.person.v1.PersonService/GetPersonByLongLatBBox"
	// PersonServiceGetAllVehiclesProcedure is the fully-qualified name of the PersonService's
	// GetAllVehicles RPC.
	PersonServiceGetAllVehiclesProcedure = "/city.person.v1.PersonService/GetAllVehicles"
	// PersonServiceResetPersonPositionProcedure is the fully-qualified name of the PersonService's
	// ResetPersonPosition RPC.
	PersonServiceResetPersonPositionProcedure = "/city.person.v1.PersonService/ResetPersonPosition"
	// PersonServiceSetControlledVehicleIDsProcedure is the fully-qualified name of the PersonService's
	// SetControlledVehicleIDs RPC.
	PersonServiceSetControlledVehicleIDsProcedure = "/city.person.v1.PersonService/SetControlledVehicleIDs"
	// PersonServiceFetchControlledVehicleEnvsProcedure is the fully-qualified name of the
	// PersonService's FetchControlledVehicleEnvs RPC.
	PersonServiceFetchControlledVehicleEnvsProcedure = "/city.person.v1.PersonService/FetchControlledVehicleEnvs"
	// PersonServiceSetControlledVehicleActionsProcedure is the fully-qualified name of the
	// PersonService's SetControlledVehicleActions RPC.
	PersonServiceSetControlledVehicleActionsProcedure = "/city.person.v1.PersonService/SetControlledVehicleActions"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	personServiceServiceDescriptor                           = v1.File_city_person_v1_person_service_proto.Services().ByName("PersonService")
	personServiceGetPersonMethodDescriptor                   = personServiceServiceDescriptor.Methods().ByName("GetPerson")
	personServiceAddPersonMethodDescriptor                   = personServiceServiceDescriptor.Methods().ByName("AddPerson")
	personServiceSetScheduleMethodDescriptor                 = personServiceServiceDescriptor.Methods().ByName("SetSchedule")
	personServiceGetPersonsMethodDescriptor                  = personServiceServiceDescriptor.Methods().ByName("GetPersons")
	personServiceGetPersonByLongLatBBoxMethodDescriptor      = personServiceServiceDescriptor.Methods().ByName("GetPersonByLongLatBBox")
	personServiceGetAllVehiclesMethodDescriptor              = personServiceServiceDescriptor.Methods().ByName("GetAllVehicles")
	personServiceResetPersonPositionMethodDescriptor         = personServiceServiceDescriptor.Methods().ByName("ResetPersonPosition")
	personServiceSetControlledVehicleIDsMethodDescriptor     = personServiceServiceDescriptor.Methods().ByName("SetControlledVehicleIDs")
	personServiceFetchControlledVehicleEnvsMethodDescriptor  = personServiceServiceDescriptor.Methods().ByName("FetchControlledVehicleEnvs")
	personServiceSetControlledVehicleActionsMethodDescriptor = personServiceServiceDescriptor.Methods().ByName("SetControlledVehicleActions")
)

// PersonServiceClient is a client for the city.person.v1.PersonService service.
type PersonServiceClient interface {
	// 获取person信息
	// Get person information
	GetPerson(context.Context, *connect.Request[v1.GetPersonRequest]) (*connect.Response[v1.GetPersonResponse], error)
	// 新增person 传入person初始位置、目的地表、属性 返回personid
	// Add a new person. Input person's initial location, destination table, and attributes, return personid
	AddPerson(context.Context, *connect.Request[v1.AddPersonRequest]) (*connect.Response[v1.AddPersonResponse], error)
	// 修改person的schedule 传入personid、目的地表
	// Set person's schedule. Input personid and destination table
	SetSchedule(context.Context, *connect.Request[v1.SetScheduleRequest]) (*connect.Response[v1.SetScheduleResponse], error)
	// 获取多个person信息
	// Get information of multiple persons
	GetPersons(context.Context, *connect.Request[v1.GetPersonsRequest]) (*connect.Response[v1.GetPersonsResponse], error)
	// 获取特定区域内的person
	// Get persons in a specific region
	GetPersonByLongLatBBox(context.Context, *connect.Request[v1.GetPersonByLongLatBBoxRequest]) (*connect.Response[v1.GetPersonByLongLatBBoxResponse], error)
	// 获取所有车辆
	// Get all vehicles
	GetAllVehicles(context.Context, *connect.Request[v1.GetAllVehiclesRequest]) (*connect.Response[v1.GetAllVehiclesResponse], error)
	// 重置人的位置（将停止当前正在进行的出行，转为sleep状态）
	// Reset person's position (stop the current trip and switch to sleep status)
	ResetPersonPosition(context.Context, *connect.Request[v1.ResetPersonPositionRequest]) (*connect.Response[v1.ResetPersonPositionResponse], error)
	// 设置由外部控制行为的vehicle
	// Set vehicle controlled by external behavior
	SetControlledVehicleIDs(context.Context, *connect.Request[v1.SetControlledVehicleIDsRequest]) (*connect.Response[v1.SetControlledVehicleIDsResponse], error)
	// 获取由外部控制行为的vehicle信息
	// Get information of vehicle controlled by external behavior
	FetchControlledVehicleEnvs(context.Context, *connect.Request[v1.FetchControlledVehicleEnvsRequest]) (*connect.Response[v1.FetchControlledVehicleEnvsResponse], error)
	// 设置由外部控制行为的vehicle的行为
	// Set behavior of vehicle controlled by external behavior
	SetControlledVehicleActions(context.Context, *connect.Request[v1.SetControlledVehicleActionsRequest]) (*connect.Response[v1.SetControlledVehicleActionsResponse], error)
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
		getPersons: connect.NewClient[v1.GetPersonsRequest, v1.GetPersonsResponse](
			httpClient,
			baseURL+PersonServiceGetPersonsProcedure,
			connect.WithSchema(personServiceGetPersonsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPersonByLongLatBBox: connect.NewClient[v1.GetPersonByLongLatBBoxRequest, v1.GetPersonByLongLatBBoxResponse](
			httpClient,
			baseURL+PersonServiceGetPersonByLongLatBBoxProcedure,
			connect.WithSchema(personServiceGetPersonByLongLatBBoxMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAllVehicles: connect.NewClient[v1.GetAllVehiclesRequest, v1.GetAllVehiclesResponse](
			httpClient,
			baseURL+PersonServiceGetAllVehiclesProcedure,
			connect.WithSchema(personServiceGetAllVehiclesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		resetPersonPosition: connect.NewClient[v1.ResetPersonPositionRequest, v1.ResetPersonPositionResponse](
			httpClient,
			baseURL+PersonServiceResetPersonPositionProcedure,
			connect.WithSchema(personServiceResetPersonPositionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setControlledVehicleIDs: connect.NewClient[v1.SetControlledVehicleIDsRequest, v1.SetControlledVehicleIDsResponse](
			httpClient,
			baseURL+PersonServiceSetControlledVehicleIDsProcedure,
			connect.WithSchema(personServiceSetControlledVehicleIDsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		fetchControlledVehicleEnvs: connect.NewClient[v1.FetchControlledVehicleEnvsRequest, v1.FetchControlledVehicleEnvsResponse](
			httpClient,
			baseURL+PersonServiceFetchControlledVehicleEnvsProcedure,
			connect.WithSchema(personServiceFetchControlledVehicleEnvsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setControlledVehicleActions: connect.NewClient[v1.SetControlledVehicleActionsRequest, v1.SetControlledVehicleActionsResponse](
			httpClient,
			baseURL+PersonServiceSetControlledVehicleActionsProcedure,
			connect.WithSchema(personServiceSetControlledVehicleActionsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// personServiceClient implements PersonServiceClient.
type personServiceClient struct {
	getPerson                   *connect.Client[v1.GetPersonRequest, v1.GetPersonResponse]
	addPerson                   *connect.Client[v1.AddPersonRequest, v1.AddPersonResponse]
	setSchedule                 *connect.Client[v1.SetScheduleRequest, v1.SetScheduleResponse]
	getPersons                  *connect.Client[v1.GetPersonsRequest, v1.GetPersonsResponse]
	getPersonByLongLatBBox      *connect.Client[v1.GetPersonByLongLatBBoxRequest, v1.GetPersonByLongLatBBoxResponse]
	getAllVehicles              *connect.Client[v1.GetAllVehiclesRequest, v1.GetAllVehiclesResponse]
	resetPersonPosition         *connect.Client[v1.ResetPersonPositionRequest, v1.ResetPersonPositionResponse]
	setControlledVehicleIDs     *connect.Client[v1.SetControlledVehicleIDsRequest, v1.SetControlledVehicleIDsResponse]
	fetchControlledVehicleEnvs  *connect.Client[v1.FetchControlledVehicleEnvsRequest, v1.FetchControlledVehicleEnvsResponse]
	setControlledVehicleActions *connect.Client[v1.SetControlledVehicleActionsRequest, v1.SetControlledVehicleActionsResponse]
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

// GetPersons calls city.person.v1.PersonService.GetPersons.
func (c *personServiceClient) GetPersons(ctx context.Context, req *connect.Request[v1.GetPersonsRequest]) (*connect.Response[v1.GetPersonsResponse], error) {
	return c.getPersons.CallUnary(ctx, req)
}

// GetPersonByLongLatBBox calls city.person.v1.PersonService.GetPersonByLongLatBBox.
func (c *personServiceClient) GetPersonByLongLatBBox(ctx context.Context, req *connect.Request[v1.GetPersonByLongLatBBoxRequest]) (*connect.Response[v1.GetPersonByLongLatBBoxResponse], error) {
	return c.getPersonByLongLatBBox.CallUnary(ctx, req)
}

// GetAllVehicles calls city.person.v1.PersonService.GetAllVehicles.
func (c *personServiceClient) GetAllVehicles(ctx context.Context, req *connect.Request[v1.GetAllVehiclesRequest]) (*connect.Response[v1.GetAllVehiclesResponse], error) {
	return c.getAllVehicles.CallUnary(ctx, req)
}

// ResetPersonPosition calls city.person.v1.PersonService.ResetPersonPosition.
func (c *personServiceClient) ResetPersonPosition(ctx context.Context, req *connect.Request[v1.ResetPersonPositionRequest]) (*connect.Response[v1.ResetPersonPositionResponse], error) {
	return c.resetPersonPosition.CallUnary(ctx, req)
}

// SetControlledVehicleIDs calls city.person.v1.PersonService.SetControlledVehicleIDs.
func (c *personServiceClient) SetControlledVehicleIDs(ctx context.Context, req *connect.Request[v1.SetControlledVehicleIDsRequest]) (*connect.Response[v1.SetControlledVehicleIDsResponse], error) {
	return c.setControlledVehicleIDs.CallUnary(ctx, req)
}

// FetchControlledVehicleEnvs calls city.person.v1.PersonService.FetchControlledVehicleEnvs.
func (c *personServiceClient) FetchControlledVehicleEnvs(ctx context.Context, req *connect.Request[v1.FetchControlledVehicleEnvsRequest]) (*connect.Response[v1.FetchControlledVehicleEnvsResponse], error) {
	return c.fetchControlledVehicleEnvs.CallUnary(ctx, req)
}

// SetControlledVehicleActions calls city.person.v1.PersonService.SetControlledVehicleActions.
func (c *personServiceClient) SetControlledVehicleActions(ctx context.Context, req *connect.Request[v1.SetControlledVehicleActionsRequest]) (*connect.Response[v1.SetControlledVehicleActionsResponse], error) {
	return c.setControlledVehicleActions.CallUnary(ctx, req)
}

// PersonServiceHandler is an implementation of the city.person.v1.PersonService service.
type PersonServiceHandler interface {
	// 获取person信息
	// Get person information
	GetPerson(context.Context, *connect.Request[v1.GetPersonRequest]) (*connect.Response[v1.GetPersonResponse], error)
	// 新增person 传入person初始位置、目的地表、属性 返回personid
	// Add a new person. Input person's initial location, destination table, and attributes, return personid
	AddPerson(context.Context, *connect.Request[v1.AddPersonRequest]) (*connect.Response[v1.AddPersonResponse], error)
	// 修改person的schedule 传入personid、目的地表
	// Set person's schedule. Input personid and destination table
	SetSchedule(context.Context, *connect.Request[v1.SetScheduleRequest]) (*connect.Response[v1.SetScheduleResponse], error)
	// 获取多个person信息
	// Get information of multiple persons
	GetPersons(context.Context, *connect.Request[v1.GetPersonsRequest]) (*connect.Response[v1.GetPersonsResponse], error)
	// 获取特定区域内的person
	// Get persons in a specific region
	GetPersonByLongLatBBox(context.Context, *connect.Request[v1.GetPersonByLongLatBBoxRequest]) (*connect.Response[v1.GetPersonByLongLatBBoxResponse], error)
	// 获取所有车辆
	// Get all vehicles
	GetAllVehicles(context.Context, *connect.Request[v1.GetAllVehiclesRequest]) (*connect.Response[v1.GetAllVehiclesResponse], error)
	// 重置人的位置（将停止当前正在进行的出行，转为sleep状态）
	// Reset person's position (stop the current trip and switch to sleep status)
	ResetPersonPosition(context.Context, *connect.Request[v1.ResetPersonPositionRequest]) (*connect.Response[v1.ResetPersonPositionResponse], error)
	// 设置由外部控制行为的vehicle
	// Set vehicle controlled by external behavior
	SetControlledVehicleIDs(context.Context, *connect.Request[v1.SetControlledVehicleIDsRequest]) (*connect.Response[v1.SetControlledVehicleIDsResponse], error)
	// 获取由外部控制行为的vehicle信息
	// Get information of vehicle controlled by external behavior
	FetchControlledVehicleEnvs(context.Context, *connect.Request[v1.FetchControlledVehicleEnvsRequest]) (*connect.Response[v1.FetchControlledVehicleEnvsResponse], error)
	// 设置由外部控制行为的vehicle的行为
	// Set behavior of vehicle controlled by external behavior
	SetControlledVehicleActions(context.Context, *connect.Request[v1.SetControlledVehicleActionsRequest]) (*connect.Response[v1.SetControlledVehicleActionsResponse], error)
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
	personServiceGetPersonsHandler := connect.NewUnaryHandler(
		PersonServiceGetPersonsProcedure,
		svc.GetPersons,
		connect.WithSchema(personServiceGetPersonsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	personServiceGetPersonByLongLatBBoxHandler := connect.NewUnaryHandler(
		PersonServiceGetPersonByLongLatBBoxProcedure,
		svc.GetPersonByLongLatBBox,
		connect.WithSchema(personServiceGetPersonByLongLatBBoxMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	personServiceGetAllVehiclesHandler := connect.NewUnaryHandler(
		PersonServiceGetAllVehiclesProcedure,
		svc.GetAllVehicles,
		connect.WithSchema(personServiceGetAllVehiclesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	personServiceResetPersonPositionHandler := connect.NewUnaryHandler(
		PersonServiceResetPersonPositionProcedure,
		svc.ResetPersonPosition,
		connect.WithSchema(personServiceResetPersonPositionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	personServiceSetControlledVehicleIDsHandler := connect.NewUnaryHandler(
		PersonServiceSetControlledVehicleIDsProcedure,
		svc.SetControlledVehicleIDs,
		connect.WithSchema(personServiceSetControlledVehicleIDsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	personServiceFetchControlledVehicleEnvsHandler := connect.NewUnaryHandler(
		PersonServiceFetchControlledVehicleEnvsProcedure,
		svc.FetchControlledVehicleEnvs,
		connect.WithSchema(personServiceFetchControlledVehicleEnvsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	personServiceSetControlledVehicleActionsHandler := connect.NewUnaryHandler(
		PersonServiceSetControlledVehicleActionsProcedure,
		svc.SetControlledVehicleActions,
		connect.WithSchema(personServiceSetControlledVehicleActionsMethodDescriptor),
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
		case PersonServiceGetPersonsProcedure:
			personServiceGetPersonsHandler.ServeHTTP(w, r)
		case PersonServiceGetPersonByLongLatBBoxProcedure:
			personServiceGetPersonByLongLatBBoxHandler.ServeHTTP(w, r)
		case PersonServiceGetAllVehiclesProcedure:
			personServiceGetAllVehiclesHandler.ServeHTTP(w, r)
		case PersonServiceResetPersonPositionProcedure:
			personServiceResetPersonPositionHandler.ServeHTTP(w, r)
		case PersonServiceSetControlledVehicleIDsProcedure:
			personServiceSetControlledVehicleIDsHandler.ServeHTTP(w, r)
		case PersonServiceFetchControlledVehicleEnvsProcedure:
			personServiceFetchControlledVehicleEnvsHandler.ServeHTTP(w, r)
		case PersonServiceSetControlledVehicleActionsProcedure:
			personServiceSetControlledVehicleActionsHandler.ServeHTTP(w, r)
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

func (UnimplementedPersonServiceHandler) GetPersons(context.Context, *connect.Request[v1.GetPersonsRequest]) (*connect.Response[v1.GetPersonsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.GetPersons is not implemented"))
}

func (UnimplementedPersonServiceHandler) GetPersonByLongLatBBox(context.Context, *connect.Request[v1.GetPersonByLongLatBBoxRequest]) (*connect.Response[v1.GetPersonByLongLatBBoxResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.GetPersonByLongLatBBox is not implemented"))
}

func (UnimplementedPersonServiceHandler) GetAllVehicles(context.Context, *connect.Request[v1.GetAllVehiclesRequest]) (*connect.Response[v1.GetAllVehiclesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.GetAllVehicles is not implemented"))
}

func (UnimplementedPersonServiceHandler) ResetPersonPosition(context.Context, *connect.Request[v1.ResetPersonPositionRequest]) (*connect.Response[v1.ResetPersonPositionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.ResetPersonPosition is not implemented"))
}

func (UnimplementedPersonServiceHandler) SetControlledVehicleIDs(context.Context, *connect.Request[v1.SetControlledVehicleIDsRequest]) (*connect.Response[v1.SetControlledVehicleIDsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.SetControlledVehicleIDs is not implemented"))
}

func (UnimplementedPersonServiceHandler) FetchControlledVehicleEnvs(context.Context, *connect.Request[v1.FetchControlledVehicleEnvsRequest]) (*connect.Response[v1.FetchControlledVehicleEnvsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.FetchControlledVehicleEnvs is not implemented"))
}

func (UnimplementedPersonServiceHandler) SetControlledVehicleActions(context.Context, *connect.Request[v1.SetControlledVehicleActionsRequest]) (*connect.Response[v1.SetControlledVehicleActionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("city.person.v1.PersonService.SetControlledVehicleActions is not implemented"))
}
