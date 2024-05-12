package logic

import (
	"context"
	"errors"

	"chatops-server/internal/user/rpc/internal/svc"
	"chatops-server/internal/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDetailLogic) UserDetail(in *user.UserDetailRequest) (*user.UserDetailResponse, error) {
	// todo: add your logic here and delete this line
	// 根据username查询用户详情
	userDetailInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorw("user.rpc.UserModel.FindOneByUsername failed", logx.LogField{Key: "err", Value: sqlx.ErrNotFound})
		return nil, errors.New("请输入正确的用户名")
	}
	if err != nil {
		logx.Errorw("user.rpc.UserModel.FindOneByUsername failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	// 根据departmentId查询departmentName
	departmentInfo, err := l.svcCtx.DepartmentModel.FindOne(l.ctx, userDetailInfo.DepartmentId)
	if errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorw("user.rpc.DepartmentModel.FindOne failed", logx.LogField{Key: "err", Value: sqlx.ErrNotFound})
		return nil, errors.New("请输入正确的用户名")
	}
	if err != nil {
		logx.Errorw("user.rpc.DepartmentModel.FindOne failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	return &user.UserDetailResponse{
		UserID:         userDetailInfo.Id,
		Username:       userDetailInfo.Username,
		RealName:       userDetailInfo.RealName,
		PhoneNumber:    userDetailInfo.PhoneNumber,
		DepartmentName: departmentInfo.Name,
		Email:          userDetailInfo.Email,
		JenkinsAccount: userDetailInfo.JenkinsAccount,
		IsAdmin:        userDetailInfo.IsAdmin,
	}, nil
}
