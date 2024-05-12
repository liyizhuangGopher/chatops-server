package logic

import (
	"context"

	"chatops-server/internal/user/model"
	"chatops-server/internal/user/rpc/internal/svc"
	"chatops-server/internal/user/rpc/user"
	md5secret "chatops-server/pkg/md5Secret"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error) {
	// 更新用户信息
	// 新密码加密
	newPassword := md5secret.MD5Secret(in.Password, l.svcCtx.Config.Secret)
	userInfoNew := &model.User{
		Id:              in.Id,
		Password:        newPassword,
		RealName:        in.RealName,
		PhoneNumber:     in.PhoneNumber,
		DepartmentId:    in.DepartmentId,
		Email:           in.Email,
		JenkinsAccount:  in.JenkinsAccount,
		JenkinsPassword: in.JenkinsPassword,
	}
	err := l.svcCtx.UserModel.Update(l.ctx, userInfoNew)
	// 返回响应
	if err != nil {
		logx.Errorw("user.rpc UserModel.Update failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	return &user.UpdateUserInfoResponse{
		Message: "ok",
	}, nil
}
