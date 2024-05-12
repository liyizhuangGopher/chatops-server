package logic

import (
	"context"

	"chatops-server/internal/system/model"
	"chatops-server/internal/system/rpc/internal/svc"
	"chatops-server/internal/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSystemInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSystemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSystemInfoLogic {
	return &UpdateSystemInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSystemInfoLogic) UpdateSystemInfo(in *system.UpdateSystemInfoRequest) (*system.UpdateSystemInfoResponse, error) {
	// 更新系统设置
	updateSystemInfo := model.System{
		Id:               1,
		DetectDuration:   in.DetectDuration,
		HistoryDurantion: in.HistoryDurantion,
	}
	err := l.svcCtx.SystemModel.Update(l.ctx, &updateSystemInfo)
	if err != nil {
		logx.Errorw("system.rpc.SystemModel.Update failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	logx.Infow("system.rpc.Conn.SystemModel.Update success")
	// 返回响应
	return &system.UpdateSystemInfoResponse{
		Message: "ok",
	}, nil
}
