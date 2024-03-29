// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: protocol/v1/saturnsync.proto

package protocolv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TaskService_ExecuteActions_FullMethodName = "/protocol.v1.TaskService/ExecuteActions"
	TaskService_ExecuteFilters_FullMethodName = "/protocol.v1.TaskService/ExecuteFilters"
	TaskService_ListTasks_FullMethodName      = "/protocol.v1.TaskService/ListTasks"
	TaskService_OnPrClosed_FullMethodName     = "/protocol.v1.TaskService/OnPrClosed"
	TaskService_OnPrCreated_FullMethodName    = "/protocol.v1.TaskService/OnPrCreated"
	TaskService_OnPrMerged_FullMethodName     = "/protocol.v1.TaskService/OnPrMerged"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	ExecuteActions(ctx context.Context, in *ExecuteActionsRequest, opts ...grpc.CallOption) (*ExecuteActionsResponse, error)
	ExecuteFilters(ctx context.Context, in *ExecuteFiltersRequest, opts ...grpc.CallOption) (*ExecuteFiltersResponse, error)
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error)
	OnPrClosed(ctx context.Context, in *OnPrClosedRequest, opts ...grpc.CallOption) (*OnPrClosedResponse, error)
	OnPrCreated(ctx context.Context, in *OnPrCreatedRequest, opts ...grpc.CallOption) (*OnPrCreatedResponse, error)
	OnPrMerged(ctx context.Context, in *OnPrMergedRequest, opts ...grpc.CallOption) (*OnPrMergedResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) ExecuteActions(ctx context.Context, in *ExecuteActionsRequest, opts ...grpc.CallOption) (*ExecuteActionsResponse, error) {
	out := new(ExecuteActionsResponse)
	err := c.cc.Invoke(ctx, TaskService_ExecuteActions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ExecuteFilters(ctx context.Context, in *ExecuteFiltersRequest, opts ...grpc.CallOption) (*ExecuteFiltersResponse, error) {
	out := new(ExecuteFiltersResponse)
	err := c.cc.Invoke(ctx, TaskService_ExecuteFilters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error) {
	out := new(ListTasksResponse)
	err := c.cc.Invoke(ctx, TaskService_ListTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) OnPrClosed(ctx context.Context, in *OnPrClosedRequest, opts ...grpc.CallOption) (*OnPrClosedResponse, error) {
	out := new(OnPrClosedResponse)
	err := c.cc.Invoke(ctx, TaskService_OnPrClosed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) OnPrCreated(ctx context.Context, in *OnPrCreatedRequest, opts ...grpc.CallOption) (*OnPrCreatedResponse, error) {
	out := new(OnPrCreatedResponse)
	err := c.cc.Invoke(ctx, TaskService_OnPrCreated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) OnPrMerged(ctx context.Context, in *OnPrMergedRequest, opts ...grpc.CallOption) (*OnPrMergedResponse, error) {
	out := new(OnPrMergedResponse)
	err := c.cc.Invoke(ctx, TaskService_OnPrMerged_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	ExecuteActions(context.Context, *ExecuteActionsRequest) (*ExecuteActionsResponse, error)
	ExecuteFilters(context.Context, *ExecuteFiltersRequest) (*ExecuteFiltersResponse, error)
	ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	OnPrClosed(context.Context, *OnPrClosedRequest) (*OnPrClosedResponse, error)
	OnPrCreated(context.Context, *OnPrCreatedRequest) (*OnPrCreatedResponse, error)
	OnPrMerged(context.Context, *OnPrMergedRequest) (*OnPrMergedResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) ExecuteActions(context.Context, *ExecuteActionsRequest) (*ExecuteActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteActions not implemented")
}
func (UnimplementedTaskServiceServer) ExecuteFilters(context.Context, *ExecuteFiltersRequest) (*ExecuteFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteFilters not implemented")
}
func (UnimplementedTaskServiceServer) ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedTaskServiceServer) OnPrClosed(context.Context, *OnPrClosedRequest) (*OnPrClosedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnPrClosed not implemented")
}
func (UnimplementedTaskServiceServer) OnPrCreated(context.Context, *OnPrCreatedRequest) (*OnPrCreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnPrCreated not implemented")
}
func (UnimplementedTaskServiceServer) OnPrMerged(context.Context, *OnPrMergedRequest) (*OnPrMergedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnPrMerged not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_ExecuteActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ExecuteActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_ExecuteActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ExecuteActions(ctx, req.(*ExecuteActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ExecuteFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ExecuteFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_ExecuteFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ExecuteFilters(ctx, req.(*ExecuteFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_ListTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListTasks(ctx, req.(*ListTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_OnPrClosed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnPrClosedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).OnPrClosed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_OnPrClosed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).OnPrClosed(ctx, req.(*OnPrClosedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_OnPrCreated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnPrCreatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).OnPrCreated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_OnPrCreated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).OnPrCreated(ctx, req.(*OnPrCreatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_OnPrMerged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnPrMergedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).OnPrMerged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_OnPrMerged_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).OnPrMerged(ctx, req.(*OnPrMergedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.v1.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteActions",
			Handler:    _TaskService_ExecuteActions_Handler,
		},
		{
			MethodName: "ExecuteFilters",
			Handler:    _TaskService_ExecuteFilters_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _TaskService_ListTasks_Handler,
		},
		{
			MethodName: "OnPrClosed",
			Handler:    _TaskService_OnPrClosed_Handler,
		},
		{
			MethodName: "OnPrCreated",
			Handler:    _TaskService_OnPrCreated_Handler,
		},
		{
			MethodName: "OnPrMerged",
			Handler:    _TaskService_OnPrMerged_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/v1/saturnsync.proto",
}
