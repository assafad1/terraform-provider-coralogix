// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: com/coralogixapis/apm/services/v1/apm_service_service.proto

package __

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

// ApmServiceServiceClient is the client API for ApmServiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApmServiceServiceClient interface {
	GetApmService(ctx context.Context, in *GetApmServiceRequest, opts ...grpc.CallOption) (*GetApmServiceResponse, error)
	BatchGetApmServices(ctx context.Context, in *BatchGetApmServicesRequest, opts ...grpc.CallOption) (*BatchGetApmServicesResponse, error)
	DeleteApmService(ctx context.Context, in *DeleteApmServiceRequest, opts ...grpc.CallOption) (*DeleteApmServiceResponse, error)
	ListApmServices(ctx context.Context, in *ListApmServicesRequest, opts ...grpc.CallOption) (*ListApmServicesResponse, error)
}

type apmServiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApmServiceServiceClient(cc grpc.ClientConnInterface) ApmServiceServiceClient {
	return &apmServiceServiceClient{cc}
}

func (c *apmServiceServiceClient) GetApmService(ctx context.Context, in *GetApmServiceRequest, opts ...grpc.CallOption) (*GetApmServiceResponse, error) {
	out := new(GetApmServiceResponse)
	err := c.cc.Invoke(ctx, "/com.coralogixapis.apm.services.v1.ApmServiceService/GetApmService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apmServiceServiceClient) BatchGetApmServices(ctx context.Context, in *BatchGetApmServicesRequest, opts ...grpc.CallOption) (*BatchGetApmServicesResponse, error) {
	out := new(BatchGetApmServicesResponse)
	err := c.cc.Invoke(ctx, "/com.coralogixapis.apm.services.v1.ApmServiceService/BatchGetApmServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apmServiceServiceClient) DeleteApmService(ctx context.Context, in *DeleteApmServiceRequest, opts ...grpc.CallOption) (*DeleteApmServiceResponse, error) {
	out := new(DeleteApmServiceResponse)
	err := c.cc.Invoke(ctx, "/com.coralogixapis.apm.services.v1.ApmServiceService/DeleteApmService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apmServiceServiceClient) ListApmServices(ctx context.Context, in *ListApmServicesRequest, opts ...grpc.CallOption) (*ListApmServicesResponse, error) {
	out := new(ListApmServicesResponse)
	err := c.cc.Invoke(ctx, "/com.coralogixapis.apm.services.v1.ApmServiceService/ListApmServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApmServiceServiceServer is the server API for ApmServiceService service.
// All implementations must embed UnimplementedApmServiceServiceServer
// for forward compatibility
type ApmServiceServiceServer interface {
	GetApmService(context.Context, *GetApmServiceRequest) (*GetApmServiceResponse, error)
	BatchGetApmServices(context.Context, *BatchGetApmServicesRequest) (*BatchGetApmServicesResponse, error)
	DeleteApmService(context.Context, *DeleteApmServiceRequest) (*DeleteApmServiceResponse, error)
	ListApmServices(context.Context, *ListApmServicesRequest) (*ListApmServicesResponse, error)
	mustEmbedUnimplementedApmServiceServiceServer()
}

// UnimplementedApmServiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApmServiceServiceServer struct {
}

func (UnimplementedApmServiceServiceServer) GetApmService(context.Context, *GetApmServiceRequest) (*GetApmServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApmService not implemented")
}
func (UnimplementedApmServiceServiceServer) BatchGetApmServices(context.Context, *BatchGetApmServicesRequest) (*BatchGetApmServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetApmServices not implemented")
}
func (UnimplementedApmServiceServiceServer) DeleteApmService(context.Context, *DeleteApmServiceRequest) (*DeleteApmServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApmService not implemented")
}
func (UnimplementedApmServiceServiceServer) ListApmServices(context.Context, *ListApmServicesRequest) (*ListApmServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApmServices not implemented")
}
func (UnimplementedApmServiceServiceServer) mustEmbedUnimplementedApmServiceServiceServer() {}

// UnsafeApmServiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApmServiceServiceServer will
// result in compilation errors.
type UnsafeApmServiceServiceServer interface {
	mustEmbedUnimplementedApmServiceServiceServer()
}

func RegisterApmServiceServiceServer(s grpc.ServiceRegistrar, srv ApmServiceServiceServer) {
	s.RegisterService(&ApmServiceService_ServiceDesc, srv)
}

func _ApmServiceService_GetApmService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApmServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmServiceServiceServer).GetApmService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.coralogixapis.apm.services.v1.ApmServiceService/GetApmService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmServiceServiceServer).GetApmService(ctx, req.(*GetApmServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApmServiceService_BatchGetApmServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetApmServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmServiceServiceServer).BatchGetApmServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.coralogixapis.apm.services.v1.ApmServiceService/BatchGetApmServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmServiceServiceServer).BatchGetApmServices(ctx, req.(*BatchGetApmServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApmServiceService_DeleteApmService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApmServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmServiceServiceServer).DeleteApmService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.coralogixapis.apm.services.v1.ApmServiceService/DeleteApmService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmServiceServiceServer).DeleteApmService(ctx, req.(*DeleteApmServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApmServiceService_ListApmServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApmServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmServiceServiceServer).ListApmServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.coralogixapis.apm.services.v1.ApmServiceService/ListApmServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmServiceServiceServer).ListApmServices(ctx, req.(*ListApmServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApmServiceService_ServiceDesc is the grpc.ServiceDesc for ApmServiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApmServiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.apm.services.v1.ApmServiceService",
	HandlerType: (*ApmServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApmService",
			Handler:    _ApmServiceService_GetApmService_Handler,
		},
		{
			MethodName: "BatchGetApmServices",
			Handler:    _ApmServiceService_BatchGetApmServices_Handler,
		},
		{
			MethodName: "DeleteApmService",
			Handler:    _ApmServiceService_DeleteApmService_Handler,
		},
		{
			MethodName: "ListApmServices",
			Handler:    _ApmServiceService_ListApmServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/apm/services/v1/apm_service_service.proto",
}
