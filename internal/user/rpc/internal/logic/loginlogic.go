package logic

import (
	"context"
	"errors"

	"chatops-server/internal/user/rpc/internal/svc"
	"chatops-server/internal/user/rpc/user"
	md5secret "chatops-server/pkg/md5Secret"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 根据username查询password 和 userID
	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if errors.Is(err, sqlx.ErrNotFound) {
		return nil, errors.New("用户不存在")
	}
	if err != nil {
		logx.Errorw("user.rpc.UserModel.FindOneByUsername failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	// 判断密码是否正确
	// 用户传入的password加密后和数据库查出来的做比较
	secret := l.svcCtx.Config.Secret
	password := md5secret.MD5Secret(in.Password, secret)
	if password != userInfo.Password {
		return nil, errors.New("用户名或密码错误")
	}
	// 返回响应
	return &user.LoginResponse{
		Message:  "ok",
		UserID:   userInfo.Id,
		Username: userInfo.Username,
	}, nil
}
