// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: alert.proto

package pbgen

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

// CaterAlertRequestClient is the client API for CaterAlertRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaterAlertRequestClient interface {
	CaterAlert(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*AlertResponse, error)
	DescribeAlert(ctx context.Context, in *DescriptionRequest, opts ...grpc.CallOption) (*AlertRequest, error)
	RetrieveAlert(ctx context.Context, in *RetrieveAlertRequest, opts ...grpc.CallOption) (*AlertList, error)
}

type caterAlertRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewCaterAlertRequestClient(cc grpc.ClientConnInterface) CaterAlertRequestClient {
	return &caterAlertRequestClient{cc}
}

func (c *caterAlertRequestClient) CaterAlert(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*AlertResponse, error) {
	out := new(AlertResponse)
	err := c.cc.Invoke(ctx, "/alerter.CaterAlertRequest/CaterAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caterAlertRequestClient) DescribeAlert(ctx context.Context, in *DescriptionRequest, opts ...grpc.CallOption) (*AlertRequest, error) {
	out := new(AlertRequest)
	err := c.cc.Invoke(ctx, "/alerter.CaterAlertRequest/DescribeAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caterAlertRequestClient) RetrieveAlert(ctx context.Context, in *RetrieveAlertRequest, opts ...grpc.CallOption) (*AlertList, error) {
	out := new(AlertList)
	err := c.cc.Invoke(ctx, "/alerter.CaterAlertRequest/RetrieveAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaterAlertRequestServer is the server API for CaterAlertRequest service.
// All implementations must embed UnimplementedCaterAlertRequestServer
// for forward compatibility
type CaterAlertRequestServer interface {
	CaterAlert(context.Context, *AlertRequest) (*AlertResponse, error)
	DescribeAlert(context.Context, *DescriptionRequest) (*AlertRequest, error)
	RetrieveAlert(context.Context, *RetrieveAlertRequest) (*AlertList, error)
	mustEmbedUnimplementedCaterAlertRequestServer()
}

// UnimplementedCaterAlertRequestServer must be embedded to have forward compatible implementations.
type UnimplementedCaterAlertRequestServer struct {
}

func (UnimplementedCaterAlertRequestServer) CaterAlert(context.Context, *AlertRequest) (*AlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CaterAlert not implemented")
}
func (UnimplementedCaterAlertRequestServer) DescribeAlert(context.Context, *DescriptionRequest) (*AlertRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAlert not implemented")
}
func (UnimplementedCaterAlertRequestServer) RetrieveAlert(context.Context, *RetrieveAlertRequest) (*AlertList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveAlert not implemented")
}
func (UnimplementedCaterAlertRequestServer) mustEmbedUnimplementedCaterAlertRequestServer() {}

// UnsafeCaterAlertRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaterAlertRequestServer will
// result in compilation errors.
type UnsafeCaterAlertRequestServer interface {
	mustEmbedUnimplementedCaterAlertRequestServer()
}

func RegisterCaterAlertRequestServer(s grpc.ServiceRegistrar, srv CaterAlertRequestServer) {
	s.RegisterService(&CaterAlertRequest_ServiceDesc, srv)
}

func _CaterAlertRequest_CaterAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaterAlertRequestServer).CaterAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alerter.CaterAlertRequest/CaterAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaterAlertRequestServer).CaterAlert(ctx, req.(*AlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaterAlertRequest_DescribeAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaterAlertRequestServer).DescribeAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alerter.CaterAlertRequest/DescribeAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaterAlertRequestServer).DescribeAlert(ctx, req.(*DescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaterAlertRequest_RetrieveAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaterAlertRequestServer).RetrieveAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alerter.CaterAlertRequest/RetrieveAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaterAlertRequestServer).RetrieveAlert(ctx, req.(*RetrieveAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaterAlertRequest_ServiceDesc is the grpc.ServiceDesc for CaterAlertRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaterAlertRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alerter.CaterAlertRequest",
	HandlerType: (*CaterAlertRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CaterAlert",
			Handler:    _CaterAlertRequest_CaterAlert_Handler,
		},
		{
			MethodName: "DescribeAlert",
			Handler:    _CaterAlertRequest_DescribeAlert_Handler,
		},
		{
			MethodName: "RetrieveAlert",
			Handler:    _CaterAlertRequest_RetrieveAlert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alert.proto",
}