package logic

import (
	"context"
	"errors"

	"chatops-server/internal/group/rpc/group"
	"chatops-server/internal/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DeleteGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGroupLogic) DeleteGroup(in *group.DeleteGroupRequest) (*group.DeleteGroupResponse, error) {
	// 根据name查出要删除的id
	delGroupInfo, err := l.svcCtx.GroupModel.FindOneByName(l.ctx, in.Name)
	if errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorw("group.rpc.GroupModel.FindOneByName failed", logx.LogField{Key: "err", Value: err})
		return nil, errors.New("未查询到指定组")
	}
	// 根据id删除对应的组信息
	err = l.svcCtx.GroupModel.Delete(l.ctx, delGroupInfo.Id)
	if err != nil {
		logx.Errorw("group.rpc.GroupModel.Delete failed", logx.LogField{Key: "err", Value: err})
		return nil, errors.New("删除组失败")
	}
	logx.Info("group.rpc.GroupModel.Delete success")
	// 返回响应

	return &group.DeleteGroupResponse{
		Message: "ok",
	}, nil
}
