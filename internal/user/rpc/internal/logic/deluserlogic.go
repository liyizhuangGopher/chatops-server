package logic

import (
	"context"
	"errors"

	"chatops-server/internal/user/rpc/internal/svc"
	"chatops-server/internal/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLogic) DelUser(in *user.DelUserRequest) (*user.DelUserResponse, error) {
	// 先根据username来获取userID
	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if errors.Is(err, sqlx.ErrNotFound) {
		return nil, errors.New("用户不存在")
	}
	if err != nil {
		logx.Errorw("user.rpc.UserModel.FindOneByUsername failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	// 再根据用户userID删除用户
	err = l.svcCtx.UserModel.Delete(l.ctx, userInfo.Id)
	if err != nil {
		logx.Errorw("user.rpc.UserModel.Delete failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	logx.Infow("user.rpc.UserModel.Delete success")
	// 返回响应
	return &user.DelUserResponse{
		Message:  "ok",
		UserID:   userInfo.Id,
		Username: in.Username,
	}, nil
}
