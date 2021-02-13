// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ThemeManagerClient is the client API for ThemeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThemeManagerClient interface {
	Create(ctx context.Context, in *ThemeRequest, opts ...grpc.CallOption) (*ThemeResponse, error)
}

type themeManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewThemeManagerClient(cc grpc.ClientConnInterface) ThemeManagerClient {
	return &themeManagerClient{cc}
}

func (c *themeManagerClient) Create(ctx context.Context, in *ThemeRequest, opts ...grpc.CallOption) (*ThemeResponse, error) {
	out := new(ThemeResponse)
	err := c.cc.Invoke(ctx, "/himo.v1.ThemeManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThemeManagerServer is the server API for ThemeManager service.
// All implementations must embed UnimplementedThemeManagerServer
// for forward compatibility
type ThemeManagerServer interface {
	Create(context.Context, *ThemeRequest) (*ThemeResponse, error)
	mustEmbedUnimplementedThemeManagerServer()
}

// UnimplementedThemeManagerServer must be embedded to have forward compatible implementations.
type UnimplementedThemeManagerServer struct {
}

func (UnimplementedThemeManagerServer) Create(context.Context, *ThemeRequest) (*ThemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedThemeManagerServer) mustEmbedUnimplementedThemeManagerServer() {}

// UnsafeThemeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThemeManagerServer will
// result in compilation errors.
type UnsafeThemeManagerServer interface {
	mustEmbedUnimplementedThemeManagerServer()
}

func RegisterThemeManagerServer(s grpc.ServiceRegistrar, srv ThemeManagerServer) {
	s.RegisterService(&_ThemeManager_serviceDesc, srv)
}

func _ThemeManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/himo.v1.ThemeManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeManagerServer).Create(ctx, req.(*ThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ThemeManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "himo.v1.ThemeManager",
	HandlerType: (*ThemeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ThemeManager_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/v1/himo/proto/theme.proto",
}
