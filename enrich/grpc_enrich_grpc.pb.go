// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package enrich

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

// EnrichClient is the client API for Enrich service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnrichClient interface {
	NLP(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type enrichClient struct {
	cc grpc.ClientConnInterface
}

func NewEnrichClient(cc grpc.ClientConnInterface) EnrichClient {
	return &enrichClient{cc}
}

func (c *enrichClient) NLP(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/enrich.Enrich/NLP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnrichServer is the server API for Enrich service.
// All implementations must embed UnimplementedEnrichServer
// for forward compatibility
type EnrichServer interface {
	NLP(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedEnrichServer()
}

// UnimplementedEnrichServer must be embedded to have forward compatible implementations.
type UnimplementedEnrichServer struct {
}

func (UnimplementedEnrichServer) NLP(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NLP not implemented")
}
func (UnimplementedEnrichServer) mustEmbedUnimplementedEnrichServer() {}

// UnsafeEnrichServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnrichServer will
// result in compilation errors.
type UnsafeEnrichServer interface {
	mustEmbedUnimplementedEnrichServer()
}

func RegisterEnrichServer(s grpc.ServiceRegistrar, srv EnrichServer) {
	s.RegisterService(&Enrich_ServiceDesc, srv)
}

func _Enrich_NLP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnrichServer).NLP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enrich.Enrich/NLP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnrichServer).NLP(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Enrich_ServiceDesc is the grpc.ServiceDesc for Enrich service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Enrich_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "enrich.Enrich",
	HandlerType: (*EnrichServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NLP",
			Handler:    _Enrich_NLP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "enrich/grpc_enrich.proto",
}
