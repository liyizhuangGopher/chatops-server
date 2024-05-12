package logic

import (
	"context"
	"errors"

	"chatops-server/internal/group/rpc/group"
	"chatops-server/internal/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GroupDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDetailLogic {
	return &GroupDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupDetailLogic) GroupDetail(in *group.GroupDetailRequest) (*group.GroupDetailResponse, error) {
	// 根据name来查询group详情
	groupInfo, err := l.svcCtx.GroupModel.FindOneByName(l.ctx, in.Name)
	if errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorw("group.rpc.GroupModel.FindOneByName failed", logx.LogField{Key: "err", Value: err})
		return nil, errors.New("未查询到指定组")
	}
	// 返回响应
	return &group.GroupDetailResponse{
		Message:      "ok",
		Name:         groupInfo.Name,
		ConversionID: groupInfo.ConversionId,
	}, nil
}
