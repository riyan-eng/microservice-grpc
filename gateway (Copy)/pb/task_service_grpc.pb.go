// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/task_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TaskService_List_FullMethodName   = "/proto.TaskService/List"
	TaskService_Create_FullMethodName = "/proto.TaskService/Create"
	TaskService_Detail_FullMethodName = "/proto.TaskService/Detail"
	TaskService_Put_FullMethodName    = "/proto.TaskService/Put"
	TaskService_Delete_FullMethodName = "/proto.TaskService/Delete"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	List(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error)
	Create(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskCreateResponse, error)
	Detail(ctx context.Context, in *TaskDetailRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error)
	Put(ctx context.Context, in *TaskPutRequest, opts ...grpc.CallOption) (*TaskPutResponse, error)
	Delete(ctx context.Context, in *TaskDeleteRequest, opts ...grpc.CallOption) (*TaskDeleteResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) List(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, TaskService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Create(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskCreateResponse)
	err := c.cc.Invoke(ctx, TaskService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Detail(ctx context.Context, in *TaskDetailRequest, opts ...grpc.CallOption) (*TaskDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskDetailResponse)
	err := c.cc.Invoke(ctx, TaskService_Detail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Put(ctx context.Context, in *TaskPutRequest, opts ...grpc.CallOption) (*TaskPutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskPutResponse)
	err := c.cc.Invoke(ctx, TaskService_Put_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Delete(ctx context.Context, in *TaskDeleteRequest, opts ...grpc.CallOption) (*TaskDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskDeleteResponse)
	err := c.cc.Invoke(ctx, TaskService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility.
type TaskServiceServer interface {
	List(context.Context, *TaskListRequest) (*TaskListResponse, error)
	Create(context.Context, *TaskCreateRequest) (*TaskCreateResponse, error)
	Detail(context.Context, *TaskDetailRequest) (*TaskDetailResponse, error)
	Put(context.Context, *TaskPutRequest) (*TaskPutResponse, error)
	Delete(context.Context, *TaskDeleteRequest) (*TaskDeleteResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskServiceServer struct{}

func (UnimplementedTaskServiceServer) List(context.Context, *TaskListRequest) (*TaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTaskServiceServer) Create(context.Context, *TaskCreateRequest) (*TaskCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTaskServiceServer) Detail(context.Context, *TaskDetailRequest) (*TaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedTaskServiceServer) Put(context.Context, *TaskPutRequest) (*TaskPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedTaskServiceServer) Delete(context.Context, *TaskDeleteRequest) (*TaskDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}
func (UnimplementedTaskServiceServer) testEmbeddedByValue()                     {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	// If the following call pancis, it indicates UnimplementedTaskServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).List(ctx, req.(*TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Create(ctx, req.(*TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_Detail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Detail(ctx, req.(*TaskDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Put(ctx, req.(*TaskPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Delete(ctx, req.(*TaskDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TaskService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TaskService_Create_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _TaskService_Detail_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _TaskService_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TaskService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/task_service.proto",
}
