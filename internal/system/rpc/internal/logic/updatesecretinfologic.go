package logic

import (
	"context"

	"chatops-server/internal/system/model"
	"chatops-server/internal/system/rpc/internal/svc"
	"chatops-server/internal/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSecretInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSecretInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSecretInfoLogic {
	return &UpdateSecretInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSecretInfoLogic) UpdateSecretInfo(in *system.UpdateSecretInfoRequest) (*system.UpdateSecretInfoResponse, error) {
	// 更新密钥设置
	updateSecretInfo := model.Robot{
		Id:        1,
		Name:      in.Name,
		Appkey:    in.AppKey,
		Appsecret: in.AppSecret,
		AgentId:   in.AgentID,
	}
	err := l.svcCtx.RobotModel.Update(l.ctx, &updateSecretInfo)
	if err != nil {
		logx.Errorw("system.rpc.RobotModel.Update failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	logx.Infow("system.rpc.Conn.RobotModel.Update success")
	return &system.UpdateSecretInfoResponse{
		Message: "ok",
	}, nil
}
