// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: coverage_service.proto

package v1

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
	CoverageService_GetTestSuite_FullMethodName = "/coverage.v1.CoverageService/GetTestSuite"
	CoverageService_InsertFile_FullMethodName   = "/coverage.v1.CoverageService/InsertFile"
	CoverageService_DeleteFile_FullMethodName   = "/coverage.v1.CoverageService/DeleteFile"
)

// CoverageServiceClient is the client API for CoverageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoverageServiceClient interface {
	GetTestSuite(ctx context.Context, in *CoverageRequest, opts ...grpc.CallOption) (*CoverageListResponse, error)
	InsertFile(ctx context.Context, in *CoverageRequest, opts ...grpc.CallOption) (*CoverageChangeResponse, error)
	DeleteFile(ctx context.Context, in *CoverageRequest, opts ...grpc.CallOption) (*CoverageChangeResponse, error)
}

type coverageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoverageServiceClient(cc grpc.ClientConnInterface) CoverageServiceClient {
	return &coverageServiceClient{cc}
}

func (c *coverageServiceClient) GetTestSuite(ctx context.Context, in *CoverageRequest, opts ...grpc.CallOption) (*CoverageListResponse, error) {
	out := new(CoverageListResponse)
	err := c.cc.Invoke(ctx, CoverageService_GetTestSuite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverageServiceClient) InsertFile(ctx context.Context, in *CoverageRequest, opts ...grpc.CallOption) (*CoverageChangeResponse, error) {
	out := new(CoverageChangeResponse)
	err := c.cc.Invoke(ctx, CoverageService_InsertFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coverageServiceClient) DeleteFile(ctx context.Context, in *CoverageRequest, opts ...grpc.CallOption) (*CoverageChangeResponse, error) {
	out := new(CoverageChangeResponse)
	err := c.cc.Invoke(ctx, CoverageService_DeleteFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoverageServiceServer is the server API for CoverageService service.
// All implementations must embed UnimplementedCoverageServiceServer
// for forward compatibility
type CoverageServiceServer interface {
	GetTestSuite(context.Context, *CoverageRequest) (*CoverageListResponse, error)
	InsertFile(context.Context, *CoverageRequest) (*CoverageChangeResponse, error)
	DeleteFile(context.Context, *CoverageRequest) (*CoverageChangeResponse, error)
	mustEmbedUnimplementedCoverageServiceServer()
}

// UnimplementedCoverageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoverageServiceServer struct {
}

func (UnimplementedCoverageServiceServer) GetTestSuite(context.Context, *CoverageRequest) (*CoverageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestSuite not implemented")
}
func (UnimplementedCoverageServiceServer) InsertFile(context.Context, *CoverageRequest) (*CoverageChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertFile not implemented")
}
func (UnimplementedCoverageServiceServer) DeleteFile(context.Context, *CoverageRequest) (*CoverageChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedCoverageServiceServer) mustEmbedUnimplementedCoverageServiceServer() {}

// UnsafeCoverageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoverageServiceServer will
// result in compilation errors.
type UnsafeCoverageServiceServer interface {
	mustEmbedUnimplementedCoverageServiceServer()
}

func RegisterCoverageServiceServer(s grpc.ServiceRegistrar, srv CoverageServiceServer) {
	s.RegisterService(&CoverageService_ServiceDesc, srv)
}

func _CoverageService_GetTestSuite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverageServiceServer).GetTestSuite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoverageService_GetTestSuite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverageServiceServer).GetTestSuite(ctx, req.(*CoverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoverageService_InsertFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverageServiceServer).InsertFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoverageService_InsertFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverageServiceServer).InsertFile(ctx, req.(*CoverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoverageService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverageServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoverageService_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverageServiceServer).DeleteFile(ctx, req.(*CoverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoverageService_ServiceDesc is the grpc.ServiceDesc for CoverageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoverageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coverage.v1.CoverageService",
	HandlerType: (*CoverageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTestSuite",
			Handler:    _CoverageService_GetTestSuite_Handler,
		},
		{
			MethodName: "InsertFile",
			Handler:    _CoverageService_InsertFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _CoverageService_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coverage_service.proto",
}