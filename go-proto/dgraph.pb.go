// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dgraph.proto

package go_proto

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("dgraph.proto", fileDescriptor_4ef5a190f1ceed8f) }

var fileDescriptor_4ef5a190f1ceed8f = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x2f, 0x4a,
	0x2c, 0xc8, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x5c, 0x29, 0xa9,
	0x69, 0xc5, 0x10, 0x21, 0xa3, 0x45, 0xcc, 0x5c, 0x5c, 0x10, 0x35, 0xa1, 0xc5, 0xa9, 0x45, 0x42,
	0x2a, 0x5c, 0x6c, 0x9e, 0x79, 0x99, 0x25, 0x2e, 0x49, 0x42, 0x3c, 0x10, 0x05, 0x7a, 0xae, 0xb9,
	0x05, 0x25, 0x95, 0x52, 0x28, 0x3c, 0x21, 0x6d, 0x2e, 0x9e, 0xa0, 0xd4, 0xf4, 0xcc, 0xe2, 0x92,
	0xd4, 0x22, 0xb0, 0x2e, 0x3e, 0xa8, 0xac, 0x5f, 0x6a, 0x39, 0x88, 0x2f, 0xc5, 0x0d, 0xe5, 0x83,
	0x25, 0x35, 0xb9, 0x58, 0x7d, 0xf2, 0xd3, 0x33, 0xf3, 0x84, 0x84, 0xa1, 0xa2, 0x60, 0x5e, 0x50,
	0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09, 0xaa, 0x52, 0x45, 0x2e, 0xb6, 0xd0, 0xcc, 0x14, 0xf7, 0xd4,
	0x12, 0x21, 0x2e, 0xa8, 0x70, 0x70, 0x49, 0x11, 0xba, 0x12, 0xce, 0xd0, 0xcc, 0x14, 0x97, 0xd4,
	0x9c, 0xd4, 0x92, 0x54, 0x14, 0x55, 0x1c, 0x30, 0x37, 0x08, 0xa9, 0x71, 0x71, 0x83, 0x94, 0xe6,
	0x25, 0xe6, 0xa6, 0xe2, 0x35, 0x4a, 0x8d, 0x8b, 0x0f, 0xa6, 0x0e, 0xaf, 0x79, 0x4a, 0x5c, 0xec,
	0x20, 0x75, 0xc1, 0xa9, 0x25, 0x42, 0xc8, 0xfa, 0xa5, 0x90, 0x54, 0x0b, 0x29, 0x73, 0x71, 0x81,
	0xc4, 0xa0, 0xe6, 0xa0, 0x28, 0x43, 0x18, 0x64, 0xce, 0x25, 0x00, 0x0b, 0x36, 0xe7, 0xd2, 0xe2,
	0x92, 0xfc, 0xdc, 0xd4, 0x22, 0x78, 0xa0, 0x78, 0x24, 0xe6, 0xa5, 0x14, 0x67, 0x24, 0x66, 0xa7,
	0x06, 0xa5, 0x16, 0x4a, 0xc1, 0xc2, 0x33, 0x3c, 0x35, 0x27, 0x39, 0x3f, 0x37, 0xd5, 0x49, 0xe2,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x0a, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xce, 0xa0, 0xf2, 0x36, 0xe8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DgraphUserClient is the client API for DgraphUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DgraphUserClient interface {
	InitDb(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	RegisterUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*User, error)
	UidGet(ctx context.Context, in *Str, opts ...grpc.CallOption) (*User, error)
	UidDelete(ctx context.Context, in *Str, opts ...grpc.CallOption) (*N, error)
	UsernameGet(ctx context.Context, in *Str, opts ...grpc.CallOption) (*User, error)
	UsernameDelete(ctx context.Context, in *Str, opts ...grpc.CallOption) (*N, error)
	UserSet(ctx context.Context, in *User, opts ...grpc.CallOption) (*Str, error)
	UserDelete(ctx context.Context, in *User, opts ...grpc.CallOption) (*N, error)
	RegisterCustomer(ctx context.Context, in *HandshakeReq, opts ...grpc.CallOption) (*Welcome, error)
}

type dgraphUserClient struct {
	cc *grpc.ClientConn
}

func NewDgraphUserClient(cc *grpc.ClientConn) DgraphUserClient {
	return &dgraphUserClient{cc}
}

func (c *dgraphUserClient) InitDb(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/InitDb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) RegisterUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UidGet(ctx context.Context, in *Str, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/UidGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UidDelete(ctx context.Context, in *Str, opts ...grpc.CallOption) (*N, error) {
	out := new(N)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/UidDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UsernameGet(ctx context.Context, in *Str, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/UsernameGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UsernameDelete(ctx context.Context, in *Str, opts ...grpc.CallOption) (*N, error) {
	out := new(N)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/UsernameDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UserSet(ctx context.Context, in *User, opts ...grpc.CallOption) (*Str, error) {
	out := new(Str)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/UserSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) UserDelete(ctx context.Context, in *User, opts ...grpc.CallOption) (*N, error) {
	out := new(N)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dgraphUserClient) RegisterCustomer(ctx context.Context, in *HandshakeReq, opts ...grpc.CallOption) (*Welcome, error) {
	out := new(Welcome)
	err := c.cc.Invoke(ctx, "/proto.dgraphUser/RegisterCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DgraphUserServer is the server API for DgraphUser service.
type DgraphUserServer interface {
	InitDb(context.Context, *Empty) (*Empty, error)
	RegisterUser(context.Context, *NewUser) (*User, error)
	Login(context.Context, *LoginRequest) (*User, error)
	UidGet(context.Context, *Str) (*User, error)
	UidDelete(context.Context, *Str) (*N, error)
	UsernameGet(context.Context, *Str) (*User, error)
	UsernameDelete(context.Context, *Str) (*N, error)
	UserSet(context.Context, *User) (*Str, error)
	UserDelete(context.Context, *User) (*N, error)
	RegisterCustomer(context.Context, *HandshakeReq) (*Welcome, error)
}

// UnimplementedDgraphUserServer can be embedded to have forward compatible implementations.
type UnimplementedDgraphUserServer struct {
}

func (*UnimplementedDgraphUserServer) InitDb(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDb not implemented")
}
func (*UnimplementedDgraphUserServer) RegisterUser(ctx context.Context, req *NewUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (*UnimplementedDgraphUserServer) Login(ctx context.Context, req *LoginRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedDgraphUserServer) UidGet(ctx context.Context, req *Str) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UidGet not implemented")
}
func (*UnimplementedDgraphUserServer) UidDelete(ctx context.Context, req *Str) (*N, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UidDelete not implemented")
}
func (*UnimplementedDgraphUserServer) UsernameGet(ctx context.Context, req *Str) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameGet not implemented")
}
func (*UnimplementedDgraphUserServer) UsernameDelete(ctx context.Context, req *Str) (*N, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameDelete not implemented")
}
func (*UnimplementedDgraphUserServer) UserSet(ctx context.Context, req *User) (*Str, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSet not implemented")
}
func (*UnimplementedDgraphUserServer) UserDelete(ctx context.Context, req *User) (*N, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (*UnimplementedDgraphUserServer) RegisterCustomer(ctx context.Context, req *HandshakeReq) (*Welcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCustomer not implemented")
}

func RegisterDgraphUserServer(s *grpc.Server, srv DgraphUserServer) {
	s.RegisterService(&_DgraphUser_serviceDesc, srv)
}

func _DgraphUser_InitDb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).InitDb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/InitDb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).InitDb(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).RegisterUser(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UidGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Str)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UidGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/UidGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UidGet(ctx, req.(*Str))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UidDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Str)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UidDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/UidDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UidDelete(ctx, req.(*Str))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UsernameGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Str)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UsernameGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/UsernameGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UsernameGet(ctx, req.(*Str))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UsernameDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Str)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UsernameDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/UsernameDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UsernameDelete(ctx, req.(*Str))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UserSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UserSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/UserSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UserSet(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).UserDelete(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _DgraphUser_RegisterCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphUserServer).RegisterCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dgraphUser/RegisterCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphUserServer).RegisterCustomer(ctx, req.(*HandshakeReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DgraphUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.dgraphUser",
	HandlerType: (*DgraphUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDb",
			Handler:    _DgraphUser_InitDb_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _DgraphUser_RegisterUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _DgraphUser_Login_Handler,
		},
		{
			MethodName: "UidGet",
			Handler:    _DgraphUser_UidGet_Handler,
		},
		{
			MethodName: "UidDelete",
			Handler:    _DgraphUser_UidDelete_Handler,
		},
		{
			MethodName: "UsernameGet",
			Handler:    _DgraphUser_UsernameGet_Handler,
		},
		{
			MethodName: "UsernameDelete",
			Handler:    _DgraphUser_UsernameDelete_Handler,
		},
		{
			MethodName: "UserSet",
			Handler:    _DgraphUser_UserSet_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _DgraphUser_UserDelete_Handler,
		},
		{
			MethodName: "RegisterCustomer",
			Handler:    _DgraphUser_RegisterCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dgraph.proto",
}
