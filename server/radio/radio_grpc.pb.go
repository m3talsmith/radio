// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: radio.proto

package radio

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

// RadioAPIClient is the client API for RadioAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RadioAPIClient interface {
	Station(ctx context.Context, opts ...grpc.CallOption) (RadioAPI_StationClient, error)
}

type radioAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewRadioAPIClient(cc grpc.ClientConnInterface) RadioAPIClient {
	return &radioAPIClient{cc}
}

func (c *radioAPIClient) Station(ctx context.Context, opts ...grpc.CallOption) (RadioAPI_StationClient, error) {
	stream, err := c.cc.NewStream(ctx, &RadioAPI_ServiceDesc.Streams[0], "/radio.RadioAPI/Station", opts...)
	if err != nil {
		return nil, err
	}
	x := &radioAPIStationClient{stream}
	return x, nil
}

type RadioAPI_StationClient interface {
	Send(*Request) error
	Recv() (*Broadcast, error)
	grpc.ClientStream
}

type radioAPIStationClient struct {
	grpc.ClientStream
}

func (x *radioAPIStationClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *radioAPIStationClient) Recv() (*Broadcast, error) {
	m := new(Broadcast)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RadioAPIServer is the server API for RadioAPI service.
// All implementations must embed UnimplementedRadioAPIServer
// for forward compatibility
type RadioAPIServer interface {
	Station(RadioAPI_StationServer) error
	mustEmbedUnimplementedRadioAPIServer()
}

// UnimplementedRadioAPIServer must be embedded to have forward compatible implementations.
type UnimplementedRadioAPIServer struct {
}

func (UnimplementedRadioAPIServer) Station(RadioAPI_StationServer) error {
	return status.Errorf(codes.Unimplemented, "method Station not implemented")
}
func (UnimplementedRadioAPIServer) mustEmbedUnimplementedRadioAPIServer() {}

// UnsafeRadioAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RadioAPIServer will
// result in compilation errors.
type UnsafeRadioAPIServer interface {
	mustEmbedUnimplementedRadioAPIServer()
}

func RegisterRadioAPIServer(s grpc.ServiceRegistrar, srv RadioAPIServer) {
	s.RegisterService(&RadioAPI_ServiceDesc, srv)
}

func _RadioAPI_Station_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RadioAPIServer).Station(&radioAPIStationServer{stream})
}

type RadioAPI_StationServer interface {
	Send(*Broadcast) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type radioAPIStationServer struct {
	grpc.ServerStream
}

func (x *radioAPIStationServer) Send(m *Broadcast) error {
	return x.ServerStream.SendMsg(m)
}

func (x *radioAPIStationServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RadioAPI_ServiceDesc is the grpc.ServiceDesc for RadioAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RadioAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "radio.RadioAPI",
	HandlerType: (*RadioAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Station",
			Handler:       _RadioAPI_Station_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "radio.proto",
}
