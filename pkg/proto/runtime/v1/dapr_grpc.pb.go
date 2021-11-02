// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package runtime

import (
	context "context"
	v1 "github.com/dapr/dapr/pkg/proto/common/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DaprClient is the client API for Dapr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaprClient interface {
	// Invokes a method on a remote Dapr app.
	InvokeService(ctx context.Context, in *InvokeServiceRequest, opts ...grpc.CallOption) (*v1.InvokeResponse, error)
	// Gets the state for a specific key.
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error)
	// Gets a bulk of state items for a list of keys
	GetBulkState(ctx context.Context, in *GetBulkStateRequest, opts ...grpc.CallOption) (*GetBulkStateResponse, error)
	// Saves the state for a specific key.
	SaveState(ctx context.Context, in *SaveStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deletes the state for a specific key.
	DeleteState(ctx context.Context, in *DeleteStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deletes a bulk of state items for a list of keys
	DeleteBulkState(ctx context.Context, in *DeleteBulkStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Executes transactions for a specified store
	ExecuteStateTransaction(ctx context.Context, in *ExecuteStateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Publishes events to the specific topic.
	PublishEvent(ctx context.Context, in *PublishEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Invokes binding data to specific output bindings
	InvokeBinding(ctx context.Context, in *InvokeBindingRequest, opts ...grpc.CallOption) (*InvokeBindingResponse, error)
	// Gets secrets from secret stores.
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
	// Gets a bulk of secrets
	GetBulkSecret(ctx context.Context, in *GetBulkSecretRequest, opts ...grpc.CallOption) (*GetBulkSecretResponse, error)
	// Register an actor timer.
	RegisterActorTimer(ctx context.Context, in *RegisterActorTimerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Unregister an actor timer.
	UnregisterActorTimer(ctx context.Context, in *UnregisterActorTimerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Register an actor reminder.
	RegisterActorReminder(ctx context.Context, in *RegisterActorReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Unregister an actor reminder.
	UnregisterActorReminder(ctx context.Context, in *UnregisterActorReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Gets the state for a specific actor.
	GetActorState(ctx context.Context, in *GetActorStateRequest, opts ...grpc.CallOption) (*GetActorStateResponse, error)
	// Executes state transactions for a specified actor
	ExecuteActorStateTransaction(ctx context.Context, in *ExecuteActorStateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// InvokeActor calls a method on an actor.
	InvokeActor(ctx context.Context, in *InvokeActorRequest, opts ...grpc.CallOption) (*InvokeActorResponse, error)
	// GetConfiguration gets configuration from configuration store.
	GetConfigurationAlpha1(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
	// SubscribeConfiguration gets configuration from configuration store and subscribe the updates event by grpc stream
	SubscribeConfigurationAlpha1(ctx context.Context, in *SubscribeConfigurationRequest, opts ...grpc.CallOption) (Dapr_SubscribeConfigurationAlpha1Client, error)
	// Gets metadata of the sidecar
	GetMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMetadataResponse, error)
	// Sets value in extended metadata of the sidecar
	SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Shutdown the sidecar
	Shutdown(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type daprClient struct {
	cc grpc.ClientConnInterface
}

func NewDaprClient(cc grpc.ClientConnInterface) DaprClient {
	return &daprClient{cc}
}

func (c *daprClient) InvokeService(ctx context.Context, in *InvokeServiceRequest, opts ...grpc.CallOption) (*v1.InvokeResponse, error) {
	out := new(v1.InvokeResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/InvokeService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error) {
	out := new(GetStateResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) GetBulkState(ctx context.Context, in *GetBulkStateRequest, opts ...grpc.CallOption) (*GetBulkStateResponse, error) {
	out := new(GetBulkStateResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/GetBulkState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) SaveState(ctx context.Context, in *SaveStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/SaveState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) DeleteState(ctx context.Context, in *DeleteStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/DeleteState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) DeleteBulkState(ctx context.Context, in *DeleteBulkStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/DeleteBulkState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) ExecuteStateTransaction(ctx context.Context, in *ExecuteStateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/ExecuteStateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) PublishEvent(ctx context.Context, in *PublishEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/PublishEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) InvokeBinding(ctx context.Context, in *InvokeBindingRequest, opts ...grpc.CallOption) (*InvokeBindingResponse, error) {
	out := new(InvokeBindingResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/InvokeBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) GetBulkSecret(ctx context.Context, in *GetBulkSecretRequest, opts ...grpc.CallOption) (*GetBulkSecretResponse, error) {
	out := new(GetBulkSecretResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/GetBulkSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) RegisterActorTimer(ctx context.Context, in *RegisterActorTimerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/RegisterActorTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) UnregisterActorTimer(ctx context.Context, in *UnregisterActorTimerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/UnregisterActorTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) RegisterActorReminder(ctx context.Context, in *RegisterActorReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/RegisterActorReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) UnregisterActorReminder(ctx context.Context, in *UnregisterActorReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/UnregisterActorReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) GetActorState(ctx context.Context, in *GetActorStateRequest, opts ...grpc.CallOption) (*GetActorStateResponse, error) {
	out := new(GetActorStateResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/GetActorState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) ExecuteActorStateTransaction(ctx context.Context, in *ExecuteActorStateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/ExecuteActorStateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) InvokeActor(ctx context.Context, in *InvokeActorRequest, opts ...grpc.CallOption) (*InvokeActorResponse, error) {
	out := new(InvokeActorResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/InvokeActor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) GetConfigurationAlpha1(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/GetConfigurationAlpha1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) SubscribeConfigurationAlpha1(ctx context.Context, in *SubscribeConfigurationRequest, opts ...grpc.CallOption) (Dapr_SubscribeConfigurationAlpha1Client, error) {
	stream, err := c.cc.NewStream(ctx, &Dapr_ServiceDesc.Streams[0], "/dapr.proto.runtime.v1.Dapr/SubscribeConfigurationAlpha1", opts...)
	if err != nil {
		return nil, err
	}
	x := &daprSubscribeConfigurationAlpha1Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dapr_SubscribeConfigurationAlpha1Client interface {
	Recv() (*SubscribeConfigurationResponse, error)
	grpc.ClientStream
}

type daprSubscribeConfigurationAlpha1Client struct {
	grpc.ClientStream
}

func (x *daprSubscribeConfigurationAlpha1Client) Recv() (*SubscribeConfigurationResponse, error) {
	m := new(SubscribeConfigurationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daprClient) GetMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMetadataResponse, error) {
	out := new(GetMetadataResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/GetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/SetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) Shutdown(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaprServer is the server API for Dapr service.
// All implementations should embed UnimplementedDaprServer
// for forward compatibility
type DaprServer interface {
	// Invokes a method on a remote Dapr app.
	InvokeService(context.Context, *InvokeServiceRequest) (*v1.InvokeResponse, error)
	// Gets the state for a specific key.
	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)
	// Gets a bulk of state items for a list of keys
	GetBulkState(context.Context, *GetBulkStateRequest) (*GetBulkStateResponse, error)
	// Saves the state for a specific key.
	SaveState(context.Context, *SaveStateRequest) (*emptypb.Empty, error)
	// Deletes the state for a specific key.
	DeleteState(context.Context, *DeleteStateRequest) (*emptypb.Empty, error)
	// Deletes a bulk of state items for a list of keys
	DeleteBulkState(context.Context, *DeleteBulkStateRequest) (*emptypb.Empty, error)
	// Executes transactions for a specified store
	ExecuteStateTransaction(context.Context, *ExecuteStateTransactionRequest) (*emptypb.Empty, error)
	// Publishes events to the specific topic.
	PublishEvent(context.Context, *PublishEventRequest) (*emptypb.Empty, error)
	// Invokes binding data to specific output bindings
	InvokeBinding(context.Context, *InvokeBindingRequest) (*InvokeBindingResponse, error)
	// Gets secrets from secret stores.
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	// Gets a bulk of secrets
	GetBulkSecret(context.Context, *GetBulkSecretRequest) (*GetBulkSecretResponse, error)
	// Register an actor timer.
	RegisterActorTimer(context.Context, *RegisterActorTimerRequest) (*emptypb.Empty, error)
	// Unregister an actor timer.
	UnregisterActorTimer(context.Context, *UnregisterActorTimerRequest) (*emptypb.Empty, error)
	// Register an actor reminder.
	RegisterActorReminder(context.Context, *RegisterActorReminderRequest) (*emptypb.Empty, error)
	// Unregister an actor reminder.
	UnregisterActorReminder(context.Context, *UnregisterActorReminderRequest) (*emptypb.Empty, error)
	// Gets the state for a specific actor.
	GetActorState(context.Context, *GetActorStateRequest) (*GetActorStateResponse, error)
	// Executes state transactions for a specified actor
	ExecuteActorStateTransaction(context.Context, *ExecuteActorStateTransactionRequest) (*emptypb.Empty, error)
	// InvokeActor calls a method on an actor.
	InvokeActor(context.Context, *InvokeActorRequest) (*InvokeActorResponse, error)
	// GetConfiguration gets configuration from configuration store.
	GetConfigurationAlpha1(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)
	// SubscribeConfiguration gets configuration from configuration store and subscribe the updates event by grpc stream
	SubscribeConfigurationAlpha1(*SubscribeConfigurationRequest, Dapr_SubscribeConfigurationAlpha1Server) error
	// Gets metadata of the sidecar
	GetMetadata(context.Context, *emptypb.Empty) (*GetMetadataResponse, error)
	// Sets value in extended metadata of the sidecar
	SetMetadata(context.Context, *SetMetadataRequest) (*emptypb.Empty, error)
	// Shutdown the sidecar
	Shutdown(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

// UnimplementedDaprServer should be embedded to have forward compatible implementations.
type UnimplementedDaprServer struct {
}

func (UnimplementedDaprServer) InvokeService(context.Context, *InvokeServiceRequest) (*v1.InvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeService not implemented")
}
func (UnimplementedDaprServer) GetState(context.Context, *GetStateRequest) (*GetStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (UnimplementedDaprServer) GetBulkState(context.Context, *GetBulkStateRequest) (*GetBulkStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulkState not implemented")
}
func (UnimplementedDaprServer) SaveState(context.Context, *SaveStateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveState not implemented")
}
func (UnimplementedDaprServer) DeleteState(context.Context, *DeleteStateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteState not implemented")
}
func (UnimplementedDaprServer) DeleteBulkState(context.Context, *DeleteBulkStateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBulkState not implemented")
}
func (UnimplementedDaprServer) ExecuteStateTransaction(context.Context, *ExecuteStateTransactionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteStateTransaction not implemented")
}
func (UnimplementedDaprServer) PublishEvent(context.Context, *PublishEventRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishEvent not implemented")
}
func (UnimplementedDaprServer) InvokeBinding(context.Context, *InvokeBindingRequest) (*InvokeBindingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeBinding not implemented")
}
func (UnimplementedDaprServer) GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedDaprServer) GetBulkSecret(context.Context, *GetBulkSecretRequest) (*GetBulkSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulkSecret not implemented")
}
func (UnimplementedDaprServer) RegisterActorTimer(context.Context, *RegisterActorTimerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterActorTimer not implemented")
}
func (UnimplementedDaprServer) UnregisterActorTimer(context.Context, *UnregisterActorTimerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterActorTimer not implemented")
}
func (UnimplementedDaprServer) RegisterActorReminder(context.Context, *RegisterActorReminderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterActorReminder not implemented")
}
func (UnimplementedDaprServer) UnregisterActorReminder(context.Context, *UnregisterActorReminderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterActorReminder not implemented")
}
func (UnimplementedDaprServer) GetActorState(context.Context, *GetActorStateRequest) (*GetActorStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorState not implemented")
}
func (UnimplementedDaprServer) ExecuteActorStateTransaction(context.Context, *ExecuteActorStateTransactionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteActorStateTransaction not implemented")
}
func (UnimplementedDaprServer) InvokeActor(context.Context, *InvokeActorRequest) (*InvokeActorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeActor not implemented")
}
func (UnimplementedDaprServer) GetConfigurationAlpha1(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigurationAlpha1 not implemented")
}
func (UnimplementedDaprServer) SubscribeConfigurationAlpha1(*SubscribeConfigurationRequest, Dapr_SubscribeConfigurationAlpha1Server) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeConfigurationAlpha1 not implemented")
}
func (UnimplementedDaprServer) GetMetadata(context.Context, *emptypb.Empty) (*GetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedDaprServer) SetMetadata(context.Context, *SetMetadataRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMetadata not implemented")
}
func (UnimplementedDaprServer) Shutdown(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}

// UnsafeDaprServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaprServer will
// result in compilation errors.
type UnsafeDaprServer interface {
	mustEmbedUnimplementedDaprServer()
}

func RegisterDaprServer(s grpc.ServiceRegistrar, srv DaprServer) {
	s.RegisterService(&Dapr_ServiceDesc, srv)
}

func _Dapr_InvokeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).InvokeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/InvokeService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).InvokeService(ctx, req.(*InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_GetBulkState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBulkStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).GetBulkState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/GetBulkState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).GetBulkState(ctx, req.(*GetBulkStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_SaveState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).SaveState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/SaveState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).SaveState(ctx, req.(*SaveStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_DeleteState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).DeleteState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/DeleteState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).DeleteState(ctx, req.(*DeleteStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_DeleteBulkState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBulkStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).DeleteBulkState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/DeleteBulkState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).DeleteBulkState(ctx, req.(*DeleteBulkStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_ExecuteStateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteStateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).ExecuteStateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/ExecuteStateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).ExecuteStateTransaction(ctx, req.(*ExecuteStateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_PublishEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).PublishEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/PublishEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).PublishEvent(ctx, req.(*PublishEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_InvokeBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).InvokeBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/InvokeBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).InvokeBinding(ctx, req.(*InvokeBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/GetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_GetBulkSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBulkSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).GetBulkSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/GetBulkSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).GetBulkSecret(ctx, req.(*GetBulkSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_RegisterActorTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterActorTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).RegisterActorTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/RegisterActorTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).RegisterActorTimer(ctx, req.(*RegisterActorTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_UnregisterActorTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterActorTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).UnregisterActorTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/UnregisterActorTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).UnregisterActorTimer(ctx, req.(*UnregisterActorTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_RegisterActorReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterActorReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).RegisterActorReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/RegisterActorReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).RegisterActorReminder(ctx, req.(*RegisterActorReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_UnregisterActorReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterActorReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).UnregisterActorReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/UnregisterActorReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).UnregisterActorReminder(ctx, req.(*UnregisterActorReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_GetActorState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActorStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).GetActorState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/GetActorState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).GetActorState(ctx, req.(*GetActorStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_ExecuteActorStateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteActorStateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).ExecuteActorStateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/ExecuteActorStateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).ExecuteActorStateTransaction(ctx, req.(*ExecuteActorStateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_InvokeActor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeActorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).InvokeActor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/InvokeActor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).InvokeActor(ctx, req.(*InvokeActorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_GetConfigurationAlpha1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).GetConfigurationAlpha1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/GetConfigurationAlpha1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).GetConfigurationAlpha1(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_SubscribeConfigurationAlpha1_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeConfigurationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaprServer).SubscribeConfigurationAlpha1(m, &daprSubscribeConfigurationAlpha1Server{stream})
}

type Dapr_SubscribeConfigurationAlpha1Server interface {
	Send(*SubscribeConfigurationResponse) error
	grpc.ServerStream
}

type daprSubscribeConfigurationAlpha1Server struct {
	grpc.ServerStream
}

func (x *daprSubscribeConfigurationAlpha1Server) Send(m *SubscribeConfigurationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Dapr_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/GetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).GetMetadata(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_SetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).SetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/SetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).SetMetadata(ctx, req.(*SetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).Shutdown(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Dapr_ServiceDesc is the grpc.ServiceDesc for Dapr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dapr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.runtime.v1.Dapr",
	HandlerType: (*DaprServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InvokeService",
			Handler:    _Dapr_InvokeService_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _Dapr_GetState_Handler,
		},
		{
			MethodName: "GetBulkState",
			Handler:    _Dapr_GetBulkState_Handler,
		},
		{
			MethodName: "SaveState",
			Handler:    _Dapr_SaveState_Handler,
		},
		{
			MethodName: "DeleteState",
			Handler:    _Dapr_DeleteState_Handler,
		},
		{
			MethodName: "DeleteBulkState",
			Handler:    _Dapr_DeleteBulkState_Handler,
		},
		{
			MethodName: "ExecuteStateTransaction",
			Handler:    _Dapr_ExecuteStateTransaction_Handler,
		},
		{
			MethodName: "PublishEvent",
			Handler:    _Dapr_PublishEvent_Handler,
		},
		{
			MethodName: "InvokeBinding",
			Handler:    _Dapr_InvokeBinding_Handler,
		},
		{
			MethodName: "GetSecret",
			Handler:    _Dapr_GetSecret_Handler,
		},
		{
			MethodName: "GetBulkSecret",
			Handler:    _Dapr_GetBulkSecret_Handler,
		},
		{
			MethodName: "RegisterActorTimer",
			Handler:    _Dapr_RegisterActorTimer_Handler,
		},
		{
			MethodName: "UnregisterActorTimer",
			Handler:    _Dapr_UnregisterActorTimer_Handler,
		},
		{
			MethodName: "RegisterActorReminder",
			Handler:    _Dapr_RegisterActorReminder_Handler,
		},
		{
			MethodName: "UnregisterActorReminder",
			Handler:    _Dapr_UnregisterActorReminder_Handler,
		},
		{
			MethodName: "GetActorState",
			Handler:    _Dapr_GetActorState_Handler,
		},
		{
			MethodName: "ExecuteActorStateTransaction",
			Handler:    _Dapr_ExecuteActorStateTransaction_Handler,
		},
		{
			MethodName: "InvokeActor",
			Handler:    _Dapr_InvokeActor_Handler,
		},
		{
			MethodName: "GetConfigurationAlpha1",
			Handler:    _Dapr_GetConfigurationAlpha1_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _Dapr_GetMetadata_Handler,
		},
		{
			MethodName: "SetMetadata",
			Handler:    _Dapr_SetMetadata_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Dapr_Shutdown_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeConfigurationAlpha1",
			Handler:       _Dapr_SubscribeConfigurationAlpha1_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dapr/proto/runtime/v1/dapr.proto",
}
