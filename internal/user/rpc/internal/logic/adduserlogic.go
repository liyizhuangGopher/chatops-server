package logic

import (
	"context"

	"chatops-server/internal/user/model"
	"chatops-server/internal/user/rpc/internal/svc"
	"chatops-server/internal/user/rpc/user"
	md5secret "chatops-server/pkg/md5Secret"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserRequest) (*user.AddUserResponse, error) {
	// 参数校验在apis层实现
	// 加密密码
	secret := l.svcCtx.Config.Secret
	password := md5secret.MD5Secret(in.Password, secret)
	// 将新增用户插入user表
	userInfo := &model.User{
		Username:        in.Username,
		Password:        password,
		RealName:        in.RealName,
		PhoneNumber:     in.PhoneNumber,
		DepartmentId:    in.DepartmentId,
		Email:           in.Email,
		JenkinsAccount:  in.JenkinsAccount,
		JenkinsPassword: in.JenkinsPassword, // md5加密不可逆，该密码还要用于登录Jenkins，所以不能加密
		Avatar:          in.Avatar,
		IsAdmin:         in.IsAdmin,
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, userInfo)
	if err != nil {
		logx.Errorw("user.rpc.UserModel.Insert failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	logx.Infow("user.rpc.UserModel.Insert success", logx.LogField{Key: "username", Value: in.Username})
	// 返回响应
	userID, _ := result.LastInsertId()
	return &user.AddUserResponse{
		Message:  "ok",
		UserID:   userID,
		Username: in.Username,
	}, nil
}
