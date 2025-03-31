// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: reporter.proto

package reporter

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
	Reporter_SendReport_FullMethodName = "/reporter.Reporter/SendReport"
)

// ReporterClient is the client API for Reporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReporterClient interface {
	SendReport(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type reporterClient struct {
	cc grpc.ClientConnInterface
}

func NewReporterClient(cc grpc.ClientConnInterface) ReporterClient {
	return &reporterClient{cc}
}

func (c *reporterClient) SendReport(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, Reporter_SendReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReporterServer is the server API for Reporter service.
// All implementations must embed UnimplementedReporterServer
// for forward compatibility.
type ReporterServer interface {
	SendReport(context.Context, *ReportRequest) (*ReportResponse, error)
	mustEmbedUnimplementedReporterServer()
}

// UnimplementedReporterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReporterServer struct{}

func (UnimplementedReporterServer) SendReport(context.Context, *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReport not implemented")
}
func (UnimplementedReporterServer) mustEmbedUnimplementedReporterServer() {}
func (UnimplementedReporterServer) testEmbeddedByValue()                  {}

// UnsafeReporterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReporterServer will
// result in compilation errors.
type UnsafeReporterServer interface {
	mustEmbedUnimplementedReporterServer()
}

func RegisterReporterServer(s grpc.ServiceRegistrar, srv ReporterServer) {
	// If the following call pancis, it indicates UnimplementedReporterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Reporter_ServiceDesc, srv)
}

func _Reporter_SendReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReporterServer).SendReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reporter_SendReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReporterServer).SendReport(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reporter_ServiceDesc is the grpc.ServiceDesc for Reporter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reporter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reporter.Reporter",
	HandlerType: (*ReporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendReport",
			Handler:    _Reporter_SendReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reporter.proto",
}
