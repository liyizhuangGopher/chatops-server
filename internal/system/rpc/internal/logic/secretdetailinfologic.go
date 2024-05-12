package logic

import (
	"context"
	"errors"

	"chatops-server/internal/system/rpc/internal/svc"
	"chatops-server/internal/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SecretDetailInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSecretDetailInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SecretDetailInfoLogic {
	return &SecretDetailInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SecretDetailInfoLogic) SecretDetailInfo(in *system.SecretDetailInfoRequest) (*system.SecretDetailInfoResponse, error) {
	// 从数据库获取机器人的密钥信息
	robotInfo, err := l.svcCtx.RobotModel.FindOne(l.ctx, 1)
	if errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorw("system.rpc.RobotModel.FindOne failed", logx.LogField{Key: "err", Value: err})
		return nil, errors.New("未查找到机器人信息")
	}
	logx.Info("system.rpc.RobotModel.FindOne success")
	// 返回响应
	return &system.SecretDetailInfoResponse{
		Message:   "ok",
		AgentID:   robotInfo.AgentId,
		AppKey:    robotInfo.Appkey,
		AppSecret: robotInfo.Appsecret,
	}, nil
}
