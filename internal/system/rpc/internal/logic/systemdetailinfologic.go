package logic

import (
	"context"
	"errors"

	"chatops-server/internal/system/rpc/internal/svc"
	"chatops-server/internal/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SystemDetailInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSystemDetailInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemDetailInfoLogic {
	return &SystemDetailInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SystemDetailInfoLogic) SystemDetailInfo(in *system.SystemDetailInfoRequest) (*system.SystemDetailInfoResponse, error) {
	SystemInfo, err := l.svcCtx.SystemModel.FindOne(l.ctx, 1)
	if errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorw("system.rpc.SystemModel.FindOne failed", logx.LogField{Key: "err", Value: err})
		return nil, errors.New("未查询到系统配置")
	}
	logx.Infow("system.rpc.Conn.SystemModel.FindOne success")
	// 返回响应
	return &system.SystemDetailInfoResponse{
		DetectDuration:   SystemInfo.DetectDuration,
		HistoryDurantion: SystemInfo.HistoryDurantion,
	}, nil
}
