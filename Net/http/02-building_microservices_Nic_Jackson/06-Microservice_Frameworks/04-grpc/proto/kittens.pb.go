// Code generated by protoc-gen-go.
// source: kittens.proto
// DO NOT EDIT!

/*
Package bmigo_grpc is a generated protocol buffer package.

It is generated from these files:

	kittens.proto

It has these top-level messages:

	Request
	Response
*/
package bmigo_grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/04-grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "bmigo.04-grpc.Request")
	proto.RegisterType((*Response)(nil), "bmigo.04-grpc.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the 04-grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Kittens service

type KittensClient interface {
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type kittensClient struct {
	cc *grpc.ClientConn
}

func NewKittensClient(cc *grpc.ClientConn) KittensClient {
	return &kittensClient{cc}
}

func (c *kittensClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/bmigo.04-grpc.Kittens/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Kittens service

type KittensServer interface {
	Hello(context.Context, *Request) (*Response, error)
}

func RegisterKittensServer(s *grpc.Server, srv KittensServer) {
	s.RegisterService(&_Kittens_serviceDesc, srv)
}

func _Kittens_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KittensServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bmigo.04-grpc.Kittens/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KittensServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kittens_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bmigo.04-grpc.Kittens",
	HandlerType: (*KittensServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Kittens_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("kittens.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xce, 0x2c, 0x29,
	0x49, 0xcd, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xca, 0xcd, 0x4c, 0xcf,
	0xd7, 0x4b, 0x2f, 0x2a, 0x48, 0x56, 0x92, 0xe5, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0xb3, 0x95, 0x64, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x04, 0xb8,
	0x98, 0x73, 0x8b, 0xd3, 0xa1, 0xd2, 0x20, 0xa6, 0x91, 0x3d, 0x17, 0xbb, 0x37, 0xc4, 0x64, 0x21,
	0x13, 0x2e, 0x56, 0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0x21, 0x61, 0x3d, 0x84, 0xe9, 0x7a, 0x50, 0xa3,
	0xa5, 0x44, 0x50, 0x05, 0x21, 0x06, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x1d, 0x64, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x26, 0xe4, 0x36, 0x55, 0xa1, 0x00, 0x00, 0x00,
}
