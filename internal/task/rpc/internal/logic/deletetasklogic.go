package logic

import (
	"context"

	"chatops-server/internal/task/rpc/internal/svc"
	"chatops-server/internal/task/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTaskLogic) DeleteTask(in *task.DeleteTaskRequest) (*task.DeleteTaskResponse, error) {
	// todo: add your logic here and delete this line

	return &task.DeleteTaskResponse{}, nil
}
