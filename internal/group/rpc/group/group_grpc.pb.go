// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: group.proto

package group

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

const (
	Group_AddGroup_FullMethodName    = "/group.Group/AddGroup"
	Group_DeleteGroup_FullMethodName = "/group.Group/DeleteGroup"
	Group_GroupList_FullMethodName   = "/group.Group/GroupList"
	Group_GroupDetail_FullMethodName = "/group.Group/GroupDetail"
)

// GroupClient is the client API for Group service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupClient interface {
	AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error)
	GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error)
	GroupDetail(ctx context.Context, in *GroupDetailRequest, opts ...grpc.CallOption) (*GroupDetailResponse, error)
}

type groupClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupClient(cc grpc.ClientConnInterface) GroupClient {
	return &groupClient{cc}
}

func (c *groupClient) AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error) {
	out := new(AddGroupResponse)
	err := c.cc.Invoke(ctx, Group_AddGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error) {
	out := new(DeleteGroupResponse)
	err := c.cc.Invoke(ctx, Group_DeleteGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error) {
	out := new(GroupListResponse)
	err := c.cc.Invoke(ctx, Group_GroupList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GroupDetail(ctx context.Context, in *GroupDetailRequest, opts ...grpc.CallOption) (*GroupDetailResponse, error) {
	out := new(GroupDetailResponse)
	err := c.cc.Invoke(ctx, Group_GroupDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServer is the server API for Group service.
// All implementations must embed UnimplementedGroupServer
// for forward compatibility
type GroupServer interface {
	AddGroup(context.Context, *AddGroupRequest) (*AddGroupResponse, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)
	GroupList(context.Context, *GroupListRequest) (*GroupListResponse, error)
	GroupDetail(context.Context, *GroupDetailRequest) (*GroupDetailResponse, error)
	mustEmbedUnimplementedGroupServer()
}

// UnimplementedGroupServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServer struct {
}

func (UnimplementedGroupServer) AddGroup(context.Context, *AddGroupRequest) (*AddGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroup not implemented")
}
func (UnimplementedGroupServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedGroupServer) GroupList(context.Context, *GroupListRequest) (*GroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupList not implemented")
}
func (UnimplementedGroupServer) GroupDetail(context.Context, *GroupDetailRequest) (*GroupDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupDetail not implemented")
}
func (UnimplementedGroupServer) mustEmbedUnimplementedGroupServer() {}

// UnsafeGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServer will
// result in compilation errors.
type UnsafeGroupServer interface {
	mustEmbedUnimplementedGroupServer()
}

func RegisterGroupServer(s grpc.ServiceRegistrar, srv GroupServer) {
	s.RegisterService(&Group_ServiceDesc, srv)
}

func _Group_AddGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).AddGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_AddGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).AddGroup(ctx, req.(*AddGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_DeleteGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_GroupList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GroupList(ctx, req.(*GroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GroupDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GroupDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_GroupDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GroupDetail(ctx, req.(*GroupDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Group_ServiceDesc is the grpc.ServiceDesc for Group service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Group_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "group.Group",
	HandlerType: (*GroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGroup",
			Handler:    _Group_AddGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _Group_DeleteGroup_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _Group_GroupList_Handler,
		},
		{
			MethodName: "GroupDetail",
			Handler:    _Group_GroupDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}