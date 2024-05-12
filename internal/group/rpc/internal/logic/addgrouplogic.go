package logic

import (
	"context"
	"errors"

	"chatops-server/internal/group/model"
	"chatops-server/internal/group/rpc/group"
	"chatops-server/internal/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupLogic {
	return &AddGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGroupLogic) AddGroup(in *group.AddGroupRequest) (*group.AddGroupResponse, error) {
	// 往数据库中插入新数据
	newGroup := model.Groups{
		Name:         in.Name,
		ConversionId: in.ConversionID,
	}
	_, err := l.svcCtx.GroupModel.Insert(l.ctx, &newGroup)
	if err != nil {
		logx.Errorw("group.rpc.GroupModel.Insert failed", logx.LogField{Key: "err", Value: err})
		return nil, errors.New("新增group失败")
	}
	// 返回响应

	return &group.AddGroupResponse{
		Message: "ok",
	}, nil
}
