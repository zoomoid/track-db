// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: internal/protobuf/track.proto

package waveman

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

// WavemanClient is the client API for Waveman service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WavemanClient interface {
	Waveform(ctx context.Context, in *TrackRequest, opts ...grpc.CallOption) (*TrackResponse, error)
}

type wavemanClient struct {
	cc grpc.ClientConnInterface
}

func NewWavemanClient(cc grpc.ClientConnInterface) WavemanClient {
	return &wavemanClient{cc}
}

func (c *wavemanClient) Waveform(ctx context.Context, in *TrackRequest, opts ...grpc.CallOption) (*TrackResponse, error) {
	out := new(TrackResponse)
	err := c.cc.Invoke(ctx, "/waveman.Waveman/Waveform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WavemanServer is the server API for Waveman service.
// All implementations must embed UnimplementedWavemanServer
// for forward compatibility
type WavemanServer interface {
	Waveform(context.Context, *TrackRequest) (*TrackResponse, error)
	mustEmbedUnimplementedWavemanServer()
}

// UnimplementedWavemanServer must be embedded to have forward compatible implementations.
type UnimplementedWavemanServer struct {
}

func (UnimplementedWavemanServer) Waveform(context.Context, *TrackRequest) (*TrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Waveform not implemented")
}
func (UnimplementedWavemanServer) mustEmbedUnimplementedWavemanServer() {}

// UnsafeWavemanServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WavemanServer will
// result in compilation errors.
type UnsafeWavemanServer interface {
	mustEmbedUnimplementedWavemanServer()
}

func RegisterWavemanServer(s grpc.ServiceRegistrar, srv WavemanServer) {
	s.RegisterService(&Waveman_ServiceDesc, srv)
}

func _Waveman_Waveform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WavemanServer).Waveform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/waveman.Waveman/Waveform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WavemanServer).Waveform(ctx, req.(*TrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Waveman_ServiceDesc is the grpc.ServiceDesc for Waveman service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Waveman_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "waveman.Waveman",
	HandlerType: (*WavemanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Waveform",
			Handler:    _Waveman_Waveform_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/protobuf/track.proto",
}
