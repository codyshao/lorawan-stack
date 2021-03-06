// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/user_services.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("lorawan-stack/api/user_services.proto", fileDescriptor_82df9ba9356987c4)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/user_services.proto", fileDescriptor_82df9ba9356987c4)
}

var fileDescriptor_82df9ba9356987c4 = []byte{
	// 892 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x4d, 0x6c, 0xdc, 0x44,
	0x18, 0xf5, 0x84, 0xb0, 0x82, 0x69, 0x14, 0xd1, 0xa1, 0x34, 0xd4, 0x89, 0xbe, 0x52, 0x93, 0x36,
	0x62, 0xe9, 0x8e, 0x45, 0x52, 0x09, 0xa9, 0xb7, 0xf2, 0xa3, 0xa8, 0x02, 0x89, 0xaa, 0x3f, 0x17,
	0x90, 0x88, 0xbc, 0xbb, 0x13, 0x67, 0xb4, 0x89, 0x6d, 0x3c, 0xb3, 0x89, 0x4c, 0xd4, 0xaa, 0xf4,
	0x54, 0x89, 0x0b, 0x12, 0x48, 0x70, 0x44, 0x9c, 0x7a, 0xec, 0xb1, 0xc7, 0x1c, 0x7b, 0xac, 0xc4,
	0xa5, 0xc7, 0xae, 0xcd, 0xa1, 0xc7, 0x1e, 0x2b, 0x0e, 0x08, 0x79, 0xc6, 0xf6, 0x66, 0xbd, 0x76,
	0xd6, 0x37, 0x7b, 0xe6, 0xcd, 0x7b, 0xef, 0x7b, 0xfe, 0xbe, 0xd9, 0xc5, 0x17, 0x77, 0xfd, 0xd0,
	0x39, 0x70, 0xbc, 0x8e, 0x90, 0x4e, 0x6f, 0x60, 0x3b, 0x01, 0xb7, 0x87, 0x82, 0x85, 0x5b, 0x82,
	0x85, 0xfb, 0xbc, 0xc7, 0x04, 0x0d, 0x42, 0x5f, 0xfa, 0x64, 0x51, 0x4a, 0x8f, 0x66, 0x50, 0xba,
	0xbf, 0x61, 0x76, 0x5c, 0x2e, 0x77, 0x86, 0x5d, 0xda, 0xf3, 0xf7, 0x6c, 0xd7, 0x77, 0x7d, 0x5b,
	0xc1, 0xba, 0xc3, 0x6d, 0xf5, 0xa6, 0x5e, 0xd4, 0x93, 0x3e, 0x6e, 0xae, 0xb8, 0xbe, 0xef, 0xee,
	0x32, 0x45, 0xef, 0x78, 0x9e, 0x2f, 0x1d, 0xc9, 0x7d, 0x2f, 0x23, 0x37, 0x97, 0xb3, 0xdd, 0x82,
	0x83, 0xed, 0x05, 0x32, 0xca, 0x36, 0x3f, 0x9c, 0x36, 0xc8, 0xfb, 0xcc, 0x93, 0x7c, 0x9b, 0xb3,
	0x30, 0x67, 0x80, 0x69, 0x50, 0xc8, 0xdd, 0x1d, 0x99, 0xef, 0xaf, 0x54, 0x57, 0xa9, 0x77, 0xd7,
	0xff, 0x7d, 0x13, 0x2f, 0xdc, 0x11, 0x2c, 0xbc, 0xc9, 0x5c, 0x2e, 0x64, 0x18, 0x91, 0xdb, 0xb8,
	0xf5, 0x79, 0xc8, 0x1c, 0xc9, 0xc8, 0x05, 0x3a, 0x59, 0x38, 0xd5, 0xeb, 0x1a, 0xfd, 0xc3, 0x90,
	0x09, 0x69, 0x9e, 0x29, 0x43, 0xd2, 0x4d, 0xeb, 0xf4, 0x83, 0xbf, 0xff, 0xf9, 0x75, 0xee, 0x94,
	0xd5, 0x52, 0x42, 0xe2, 0x2a, 0x6a, 0x93, 0xef, 0xf1, 0x1b, 0x9b, 0x4c, 0x12, 0x28, 0xe3, 0x37,
	0x99, 0x9c, 0xcd, 0x77, 0x41, 0xf1, 0x2d, 0x93, 0x73, 0x9a, 0xcf, 0x3e, 0x54, 0x5f, 0x89, 0xf7,
	0x05, 0xcd, 0x1e, 0xee, 0x92, 0x6f, 0xf0, 0xfc, 0xd7, 0x5c, 0x48, 0xf2, 0x41, 0x99, 0x20, 0x5d,
	0x4d, 0x49, 0x44, 0x2e, 0xf1, 0x5e, 0x95, 0x84, 0xb0, 0x16, 0x95, 0xc6, 0x5b, 0x24, 0xf3, 0x4c,
	0x5c, 0xdc, 0xba, 0x13, 0xf4, 0x2b, 0x63, 0xd0, 0xeb, 0xb3, 0x6d, 0xaf, 0x2a, 0x4a, 0x30, 0x27,
	0x6c, 0xd3, 0xe3, 0xb6, 0xd3, 0x64, 0x7e, 0x47, 0x78, 0x49, 0x07, 0x7b, 0x9b, 0xed, 0x05, 0x7e,
	0xe8, 0x84, 0xd1, 0x0d, 0x47, 0x88, 0x03, 0x3f, 0xec, 0x13, 0x5a, 0xfd, 0x05, 0xa6, 0x80, 0xb9,
	0x8f, 0xb3, 0x54, 0x77, 0x13, 0xcd, 0xbb, 0x89, 0x7e, 0x99, 0x76, 0x93, 0x75, 0x45, 0x39, 0xa1,
	0xd6, 0xe5, 0xda, 0x00, 0x6d, 0x99, 0x73, 0x6e, 0x05, 0xb9, 0xfa, 0x03, 0x84, 0x17, 0x75, 0xad,
	0x85, 0xa1, 0x8f, 0xea, 0xb3, 0x68, 0xea, 0xa5, 0xa3, 0xbc, 0xac, 0x99, 0x56, 0xbd, 0x97, 0xdc,
	0x41, 0x1a, 0xcf, 0x77, 0xb8, 0xf5, 0x05, 0xdb, 0x65, 0x92, 0x91, 0xf3, 0x55, 0x21, 0x5f, 0x1f,
	0x8f, 0x43, 0xad, 0xe2, 0xfb, 0x4a, 0x91, 0xb4, 0xdf, 0x29, 0x29, 0xde, 0x5d, 0xff, 0x6f, 0x1e,
	0xe3, 0x94, 0xe5, 0x5a, 0xaf, 0xc7, 0x84, 0x20, 0xdb, 0x18, 0xa7, 0xed, 0x72, 0x53, 0x4d, 0x4f,
	0x13, 0xbd, 0x12, 0x40, 0x1f, 0xb4, 0xce, 0x2b, 0xbd, 0x73, 0x64, 0xa9, 0xac, 0x97, 0xcd, 0x25,
	0xb9, 0x87, 0x17, 0xf4, 0x87, 0xbc, 0x76, 0xe3, 0xfa, 0x57, 0x2c, 0x22, 0x6b, 0xf5, 0x83, 0xa6,
	0x11, 0xe3, 0x4c, 0x4b, 0x40, 0xbd, 0x9d, 0x67, 0x6a, 0x9d, 0x90, 0xa9, 0x13, 0xf0, 0xce, 0x80,
	0x45, 0x6a, 0x18, 0x7f, 0xc4, 0xa7, 0xd2, 0x3a, 0xf5, 0x61, 0x41, 0x2e, 0xd5, 0xcd, 0x4c, 0x06,
	0xc8, 0xd5, 0x97, 0xaa, 0xd5, 0x85, 0xd5, 0x56, 0xf2, 0xab, 0xa4, 0x81, 0x3c, 0xb9, 0x87, 0xdf,
	0xde, 0x64, 0x99, 0x34, 0x59, 0xad, 0xb9, 0x0e, 0x9a, 0x55, 0xbd, 0xa1, 0x64, 0x3b, 0xe4, 0xe3,
	0xd9, 0xb2, 0xf6, 0xe1, 0x80, 0x45, 0xea, 0xa2, 0xf8, 0x19, 0xe1, 0x05, 0xdd, 0xb4, 0x75, 0xe1,
	0x8f, 0x5b, 0xba, 0x99, 0x8d, 0xab, 0xca, 0xc6, 0x15, 0xd3, 0x6e, 0x62, 0xc3, 0x09, 0xf8, 0xd6,
	0x80, 0x45, 0x54, 0x0f, 0xff, 0xfa, 0xd1, 0x1c, 0x3e, 0xab, 0xda, 0xca, 0xdb, 0xe7, 0xfa, 0x77,
	0xa1, 0xb8, 0x87, 0xbb, 0x78, 0xfe, 0x16, 0xf3, 0xfa, 0xe4, 0x62, 0x59, 0x36, 0x5d, 0x3d, 0x8e,
	0xd7, 0xee, 0xcc, 0x32, 0x6c, 0x0c, 0xb1, 0x96, 0x94, 0xc3, 0xd3, 0xd6, 0x82, 0xcd, 0x8b, 0x45,
	0xd5, 0x08, 0x4e, 0x76, 0x6b, 0x56, 0x76, 0xc0, 0x98, 0xa0, 0xe8, 0x80, 0xe5, 0x7a, 0x11, 0x61,
	0x9d, 0x51, 0x2a, 0x8b, 0x64, 0x42, 0x85, 0x6c, 0x15, 0xf3, 0x3b, 0x15, 0xb4, 0x5e, 0x9f, 0x2e,
	0xa5, 0x6e, 0x8e, 0x33, 0x81, 0xf6, 0x84, 0xc0, 0xfa, 0x6f, 0x73, 0xf8, 0xdd, 0x34, 0xc2, 0x5b,
	0x4c, 0x88, 0xe3, 0xf9, 0x45, 0x59, 0x6d, 0x6b, 0x75, 0xdd, 0x9d, 0x1d, 0x28, 0x8a, 0x5b, 0xa9,
	0x9a, 0xf7, 0x1c, 0xd4, 0xa4, 0xc7, 0x45, 0x86, 0x25, 0x3f, 0xa1, 0xa2, 0xe8, 0x4b, 0x27, 0x90,
	0x36, 0xb9, 0xbb, 0x3e, 0x55, 0xb2, 0x9f, 0xb4, 0xed, 0xd9, 0xb2, 0xf6, 0x61, 0xf6, 0x94, 0xae,
	0x7e, 0xf6, 0x17, 0x7a, 0x3a, 0x02, 0xf4, 0x6c, 0x04, 0xe8, 0xf9, 0x08, 0x8c, 0x17, 0x23, 0x30,
	0x5e, 0x8e, 0xc0, 0x78, 0x35, 0x02, 0xe3, 0xf5, 0x08, 0xd0, 0xfd, 0x18, 0xd0, 0xc3, 0x18, 0x8c,
	0x47, 0x31, 0xa0, 0xc7, 0x31, 0x18, 0x4f, 0x62, 0x30, 0x8e, 0x62, 0x30, 0x9e, 0xc6, 0x80, 0x9e,
	0xc5, 0x80, 0x9e, 0xc7, 0x60, 0xbc, 0x88, 0x01, 0xbd, 0x8c, 0xc1, 0x78, 0x15, 0x03, 0x7a, 0x1d,
	0x83, 0x71, 0x3f, 0x01, 0xe3, 0x61, 0x02, 0xe8, 0x97, 0x04, 0x8c, 0x3f, 0x12, 0x40, 0x7f, 0x26,
	0x60, 0x3c, 0x4a, 0xc0, 0x78, 0x9c, 0x00, 0x7a, 0x92, 0x00, 0x3a, 0x4a, 0x00, 0x7d, 0x7b, 0xd9,
	0xf5, 0xa9, 0xdc, 0x61, 0x72, 0x87, 0x7b, 0xae, 0xa0, 0x1e, 0x93, 0x07, 0x7e, 0x38, 0xb0, 0x27,
	0xff, 0x81, 0x04, 0x03, 0xd7, 0x96, 0xd2, 0x0b, 0xba, 0xdd, 0x96, 0xaa, 0x76, 0xe3, 0xff, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x91, 0xfb, 0x3c, 0x07, 0x89, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserRegistryClient is the client API for UserRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserRegistryClient interface {
	// Register a new user. This method may be restricted by network settings.
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Get the user with the given identifiers, selecting the fields given by the
	// field mask. The method may return more or less fields, depending on the rights
	// of the caller.
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	List(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*Users, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Create a temporary password that can be used for updating a forgotten password.
	// The generated password is sent to the user's email address.
	CreateTemporaryPassword(ctx context.Context, in *CreateTemporaryPasswordRequest, opts ...grpc.CallOption) (*types.Empty, error)
	UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Delete(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type userRegistryClient struct {
	cc *grpc.ClientConn
}

func NewUserRegistryClient(cc *grpc.ClientConn) UserRegistryClient {
	return &userRegistryClient{cc}
}

func (c *userRegistryClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) List(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) CreateTemporaryPassword(ctx context.Context, in *CreateTemporaryPasswordRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/CreateTemporaryPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) Delete(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRegistryServer is the server API for UserRegistry service.
type UserRegistryServer interface {
	// Register a new user. This method may be restricted by network settings.
	Create(context.Context, *CreateUserRequest) (*User, error)
	// Get the user with the given identifiers, selecting the fields given by the
	// field mask. The method may return more or less fields, depending on the rights
	// of the caller.
	Get(context.Context, *GetUserRequest) (*User, error)
	List(context.Context, *ListUsersRequest) (*Users, error)
	Update(context.Context, *UpdateUserRequest) (*User, error)
	// Create a temporary password that can be used for updating a forgotten password.
	// The generated password is sent to the user's email address.
	CreateTemporaryPassword(context.Context, *CreateTemporaryPasswordRequest) (*types.Empty, error)
	UpdatePassword(context.Context, *UpdateUserPasswordRequest) (*types.Empty, error)
	Delete(context.Context, *UserIdentifiers) (*types.Empty, error)
}

// UnimplementedUserRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedUserRegistryServer struct {
}

func (*UnimplementedUserRegistryServer) Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserRegistryServer) Get(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUserRegistryServer) List(ctx context.Context, req *ListUsersRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedUserRegistryServer) Update(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserRegistryServer) CreateTemporaryPassword(ctx context.Context, req *CreateTemporaryPasswordRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemporaryPassword not implemented")
}
func (*UnimplementedUserRegistryServer) UpdatePassword(ctx context.Context, req *UpdateUserPasswordRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (*UnimplementedUserRegistryServer) Delete(ctx context.Context, req *UserIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserRegistryServer(s *grpc.Server, srv UserRegistryServer) {
	s.RegisterService(&_UserRegistry_serviceDesc, srv)
}

func _UserRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).Get(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).List(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_CreateTemporaryPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemporaryPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).CreateTemporaryPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/CreateTemporaryPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).CreateTemporaryPassword(ctx, req.(*CreateTemporaryPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).UpdatePassword(ctx, req.(*UpdateUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).Delete(ctx, req.(*UserIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.UserRegistry",
	HandlerType: (*UserRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserRegistry_Update_Handler,
		},
		{
			MethodName: "CreateTemporaryPassword",
			Handler:    _UserRegistry_CreateTemporaryPassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserRegistry_UpdatePassword_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/user_services.proto",
}

// UserAccessClient is the client API for UserAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserAccessClient interface {
	ListRights(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	CreateAPIKey(ctx context.Context, in *CreateUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	ListAPIKeys(ctx context.Context, in *ListUserAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error)
	GetAPIKey(ctx context.Context, in *GetUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Update the rights of an existing user API key. To generate an API key,
	// the CreateAPIKey should be used. To delete an API key, update it
	// with zero rights. It is required for the caller to have all assigned or/and removed rights.
	UpdateAPIKey(ctx context.Context, in *UpdateUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
}

type userAccessClient struct {
	cc *grpc.ClientConn
}

func NewUserAccessClient(cc *grpc.ClientConn) UserAccessClient {
	return &userAccessClient{cc}
}

func (c *userAccessClient) ListRights(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessClient) CreateAPIKey(ctx context.Context, in *CreateUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserAccess/CreateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessClient) ListAPIKeys(ctx context.Context, in *ListUserAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error) {
	out := new(APIKeys)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserAccess/ListAPIKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessClient) GetAPIKey(ctx context.Context, in *GetUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserAccess/GetAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessClient) UpdateAPIKey(ctx context.Context, in *UpdateUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserAccess/UpdateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAccessServer is the server API for UserAccess service.
type UserAccessServer interface {
	ListRights(context.Context, *UserIdentifiers) (*Rights, error)
	CreateAPIKey(context.Context, *CreateUserAPIKeyRequest) (*APIKey, error)
	ListAPIKeys(context.Context, *ListUserAPIKeysRequest) (*APIKeys, error)
	GetAPIKey(context.Context, *GetUserAPIKeyRequest) (*APIKey, error)
	// Update the rights of an existing user API key. To generate an API key,
	// the CreateAPIKey should be used. To delete an API key, update it
	// with zero rights. It is required for the caller to have all assigned or/and removed rights.
	UpdateAPIKey(context.Context, *UpdateUserAPIKeyRequest) (*APIKey, error)
}

// UnimplementedUserAccessServer can be embedded to have forward compatible implementations.
type UnimplementedUserAccessServer struct {
}

func (*UnimplementedUserAccessServer) ListRights(ctx context.Context, req *UserIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (*UnimplementedUserAccessServer) CreateAPIKey(ctx context.Context, req *CreateUserAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPIKey not implemented")
}
func (*UnimplementedUserAccessServer) ListAPIKeys(ctx context.Context, req *ListUserAPIKeysRequest) (*APIKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAPIKeys not implemented")
}
func (*UnimplementedUserAccessServer) GetAPIKey(ctx context.Context, req *GetUserAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIKey not implemented")
}
func (*UnimplementedUserAccessServer) UpdateAPIKey(ctx context.Context, req *UpdateUserAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPIKey not implemented")
}

func RegisterUserAccessServer(s *grpc.Server, srv UserAccessServer) {
	s.RegisterService(&_UserAccess_serviceDesc, srv)
}

func _UserAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessServer).ListRights(ctx, req.(*UserIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccess_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserAccess/CreateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessServer).CreateAPIKey(ctx, req.(*CreateUserAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccess_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserAPIKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserAccess/ListAPIKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessServer).ListAPIKeys(ctx, req.(*ListUserAPIKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccess_GetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessServer).GetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserAccess/GetAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessServer).GetAPIKey(ctx, req.(*GetUserAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccess_UpdateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessServer).UpdateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserAccess/UpdateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessServer).UpdateAPIKey(ctx, req.(*UpdateUserAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.UserAccess",
	HandlerType: (*UserAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _UserAccess_ListRights_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _UserAccess_CreateAPIKey_Handler,
		},
		{
			MethodName: "ListAPIKeys",
			Handler:    _UserAccess_ListAPIKeys_Handler,
		},
		{
			MethodName: "GetAPIKey",
			Handler:    _UserAccess_GetAPIKey_Handler,
		},
		{
			MethodName: "UpdateAPIKey",
			Handler:    _UserAccess_UpdateAPIKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/user_services.proto",
}

// UserInvitationRegistryClient is the client API for UserInvitationRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInvitationRegistryClient interface {
	Send(ctx context.Context, in *SendInvitationRequest, opts ...grpc.CallOption) (*Invitation, error)
	List(ctx context.Context, in *ListInvitationsRequest, opts ...grpc.CallOption) (*Invitations, error)
	Delete(ctx context.Context, in *DeleteInvitationRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type userInvitationRegistryClient struct {
	cc *grpc.ClientConn
}

func NewUserInvitationRegistryClient(cc *grpc.ClientConn) UserInvitationRegistryClient {
	return &userInvitationRegistryClient{cc}
}

func (c *userInvitationRegistryClient) Send(ctx context.Context, in *SendInvitationRequest, opts ...grpc.CallOption) (*Invitation, error) {
	out := new(Invitation)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserInvitationRegistry/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInvitationRegistryClient) List(ctx context.Context, in *ListInvitationsRequest, opts ...grpc.CallOption) (*Invitations, error) {
	out := new(Invitations)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserInvitationRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInvitationRegistryClient) Delete(ctx context.Context, in *DeleteInvitationRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserInvitationRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInvitationRegistryServer is the server API for UserInvitationRegistry service.
type UserInvitationRegistryServer interface {
	Send(context.Context, *SendInvitationRequest) (*Invitation, error)
	List(context.Context, *ListInvitationsRequest) (*Invitations, error)
	Delete(context.Context, *DeleteInvitationRequest) (*types.Empty, error)
}

// UnimplementedUserInvitationRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedUserInvitationRegistryServer struct {
}

func (*UnimplementedUserInvitationRegistryServer) Send(ctx context.Context, req *SendInvitationRequest) (*Invitation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedUserInvitationRegistryServer) List(ctx context.Context, req *ListInvitationsRequest) (*Invitations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedUserInvitationRegistryServer) Delete(ctx context.Context, req *DeleteInvitationRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserInvitationRegistryServer(s *grpc.Server, srv UserInvitationRegistryServer) {
	s.RegisterService(&_UserInvitationRegistry_serviceDesc, srv)
}

func _UserInvitationRegistry_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInvitationRegistryServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserInvitationRegistry/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInvitationRegistryServer).Send(ctx, req.(*SendInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInvitationRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInvitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInvitationRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserInvitationRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInvitationRegistryServer).List(ctx, req.(*ListInvitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInvitationRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInvitationRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserInvitationRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInvitationRegistryServer).Delete(ctx, req.(*DeleteInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInvitationRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.UserInvitationRegistry",
	HandlerType: (*UserInvitationRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _UserInvitationRegistry_Send_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserInvitationRegistry_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserInvitationRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/user_services.proto",
}

// UserSessionRegistryClient is the client API for UserSessionRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSessionRegistryClient interface {
	List(ctx context.Context, in *ListUserSessionsRequest, opts ...grpc.CallOption) (*UserSessions, error)
	Delete(ctx context.Context, in *UserSessionIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type userSessionRegistryClient struct {
	cc *grpc.ClientConn
}

func NewUserSessionRegistryClient(cc *grpc.ClientConn) UserSessionRegistryClient {
	return &userSessionRegistryClient{cc}
}

func (c *userSessionRegistryClient) List(ctx context.Context, in *ListUserSessionsRequest, opts ...grpc.CallOption) (*UserSessions, error) {
	out := new(UserSessions)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserSessionRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSessionRegistryClient) Delete(ctx context.Context, in *UserSessionIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserSessionRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSessionRegistryServer is the server API for UserSessionRegistry service.
type UserSessionRegistryServer interface {
	List(context.Context, *ListUserSessionsRequest) (*UserSessions, error)
	Delete(context.Context, *UserSessionIdentifiers) (*types.Empty, error)
}

// UnimplementedUserSessionRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedUserSessionRegistryServer struct {
}

func (*UnimplementedUserSessionRegistryServer) List(ctx context.Context, req *ListUserSessionsRequest) (*UserSessions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedUserSessionRegistryServer) Delete(ctx context.Context, req *UserSessionIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserSessionRegistryServer(s *grpc.Server, srv UserSessionRegistryServer) {
	s.RegisterService(&_UserSessionRegistry_serviceDesc, srv)
}

func _UserSessionRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSessionRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserSessionRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSessionRegistryServer).List(ctx, req.(*ListUserSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSessionRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSessionIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSessionRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserSessionRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSessionRegistryServer).Delete(ctx, req.(*UserSessionIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSessionRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.UserSessionRegistry",
	HandlerType: (*UserSessionRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserSessionRegistry_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserSessionRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/user_services.proto",
}
