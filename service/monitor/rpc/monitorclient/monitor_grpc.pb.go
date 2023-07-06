// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.8.0
// source: monitor.proto

package monitorclient

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

// MonitorClient is the client API for Monitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitorClient interface {
	ReportAdd(ctx context.Context, in *ReportAddReq, opts ...grpc.CallOption) (*ReportAddResp, error)
	GraphList(ctx context.Context, in *GraphListReq, opts ...grpc.CallOption) (*GraphListResp, error)
	SelectReport(ctx context.Context, in *SelectReportReq, opts ...grpc.CallOption) (*SelectReportResp, error)
}

type monitorClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorClient(cc grpc.ClientConnInterface) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) ReportAdd(ctx context.Context, in *ReportAddReq, opts ...grpc.CallOption) (*ReportAddResp, error) {
	out := new(ReportAddResp)
	err := c.cc.Invoke(ctx, "/monitorclient.Monitor/ReportAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) GraphList(ctx context.Context, in *GraphListReq, opts ...grpc.CallOption) (*GraphListResp, error) {
	out := new(GraphListResp)
	err := c.cc.Invoke(ctx, "/monitorclient.Monitor/GraphList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) SelectReport(ctx context.Context, in *SelectReportReq, opts ...grpc.CallOption) (*SelectReportResp, error) {
	out := new(SelectReportResp)
	err := c.cc.Invoke(ctx, "/monitorclient.Monitor/SelectReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorServer is the server API for Monitor service.
// All implementations must embed UnimplementedMonitorServer
// for forward compatibility
type MonitorServer interface {
	ReportAdd(context.Context, *ReportAddReq) (*ReportAddResp, error)
	GraphList(context.Context, *GraphListReq) (*GraphListResp, error)
	SelectReport(context.Context, *SelectReportReq) (*SelectReportResp, error)
	mustEmbedUnimplementedMonitorServer()
}

// UnimplementedMonitorServer must be embedded to have forward compatible implementations.
type UnimplementedMonitorServer struct {
}

func (UnimplementedMonitorServer) ReportAdd(context.Context, *ReportAddReq) (*ReportAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportAdd not implemented")
}
func (UnimplementedMonitorServer) GraphList(context.Context, *GraphListReq) (*GraphListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GraphList not implemented")
}
func (UnimplementedMonitorServer) SelectReport(context.Context, *SelectReportReq) (*SelectReportResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectReport not implemented")
}
func (UnimplementedMonitorServer) mustEmbedUnimplementedMonitorServer() {}

// UnsafeMonitorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitorServer will
// result in compilation errors.
type UnsafeMonitorServer interface {
	mustEmbedUnimplementedMonitorServer()
}

func RegisterMonitorServer(s grpc.ServiceRegistrar, srv MonitorServer) {
	s.RegisterService(&Monitor_ServiceDesc, srv)
}

func _Monitor_ReportAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).ReportAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitorclient.Monitor/ReportAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).ReportAdd(ctx, req.(*ReportAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_GraphList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).GraphList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitorclient.Monitor/GraphList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).GraphList(ctx, req.(*GraphListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_SelectReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).SelectReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitorclient.Monitor/SelectReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).SelectReport(ctx, req.(*SelectReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Monitor_ServiceDesc is the grpc.ServiceDesc for Monitor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monitor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitorclient.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportAdd",
			Handler:    _Monitor_ReportAdd_Handler,
		},
		{
			MethodName: "GraphList",
			Handler:    _Monitor_GraphList_Handler,
		},
		{
			MethodName: "SelectReport",
			Handler:    _Monitor_SelectReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor.proto",
}