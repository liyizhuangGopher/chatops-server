// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package systemclient

import (
	"context"

	"chatops-server/internal/system/rpc/system"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SecretDetailInfoRequest  = system.SecretDetailInfoRequest
	SecretDetailInfoResponse = system.SecretDetailInfoResponse
	SystemDetailInfoRequest  = system.SystemDetailInfoRequest
	SystemDetailInfoResponse = system.SystemDetailInfoResponse
	UpdateSecretInfoRequest  = system.UpdateSecretInfoRequest
	UpdateSecretInfoResponse = system.UpdateSecretInfoResponse
	UpdateSystemInfoRequest  = system.UpdateSystemInfoRequest
	UpdateSystemInfoResponse = system.UpdateSystemInfoResponse

	System interface {
		SystemDetailInfo(ctx context.Context, in *SystemDetailInfoRequest, opts ...grpc.CallOption) (*SystemDetailInfoResponse, error)
		UpdateSystemInfo(ctx context.Context, in *UpdateSystemInfoRequest, opts ...grpc.CallOption) (*UpdateSystemInfoResponse, error)
		SecretDetailInfo(ctx context.Context, in *SecretDetailInfoRequest, opts ...grpc.CallOption) (*SecretDetailInfoResponse, error)
		UpdateSecretInfo(ctx context.Context, in *UpdateSecretInfoRequest, opts ...grpc.CallOption) (*UpdateSecretInfoResponse, error)
	}

	defaultSystem struct {
		cli zrpc.Client
	}
)

func NewSystem(cli zrpc.Client) System {
	return &defaultSystem{
		cli: cli,
	}
}

func (m *defaultSystem) SystemDetailInfo(ctx context.Context, in *SystemDetailInfoRequest, opts ...grpc.CallOption) (*SystemDetailInfoResponse, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.SystemDetailInfo(ctx, in, opts...)
}

func (m *defaultSystem) UpdateSystemInfo(ctx context.Context, in *UpdateSystemInfoRequest, opts ...grpc.CallOption) (*UpdateSystemInfoResponse, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateSystemInfo(ctx, in, opts...)
}

func (m *defaultSystem) SecretDetailInfo(ctx context.Context, in *SecretDetailInfoRequest, opts ...grpc.CallOption) (*SecretDetailInfoResponse, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.SecretDetailInfo(ctx, in, opts...)
}

func (m *defaultSystem) UpdateSecretInfo(ctx context.Context, in *UpdateSecretInfoRequest, opts ...grpc.CallOption) (*UpdateSecretInfoResponse, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateSecretInfo(ctx, in, opts...)
}
