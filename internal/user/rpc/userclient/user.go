// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"chatops-server/internal/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddUserRequest         = user.AddUserRequest
	AddUserResponse        = user.AddUserResponse
	DelUserRequest         = user.DelUserRequest
	DelUserResponse        = user.DelUserResponse
	LoginRequest           = user.LoginRequest
	LoginResponse          = user.LoginResponse
	UpdateUserInfoRequest  = user.UpdateUserInfoRequest
	UpdateUserInfoResponse = user.UpdateUserInfoResponse
	UserDetailRequest      = user.UserDetailRequest
	UserDetailResponse     = user.UserDetailResponse
	UserInfo               = user.UserInfo
	UserListRequest        = user.UserListRequest
	UserListResponse       = user.UserListResponse

	User interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
		UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
		UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
		UserDetail(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailResponse, error)
		DelUser(ctx context.Context, in *DelUserRequest, opts ...grpc.CallOption) (*DelUserResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}

func (m *defaultUser) UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultUser) UserDetail(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserDetail(ctx, in, opts...)
}

func (m *defaultUser) DelUser(ctx context.Context, in *DelUserRequest, opts ...grpc.CallOption) (*DelUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DelUser(ctx, in, opts...)
}
