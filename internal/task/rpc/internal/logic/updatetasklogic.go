package logic

import (
	"context"

	"chatops-server/internal/task/rpc/internal/svc"
	"chatops-server/internal/task/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogic {
	return &UpdateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskLogic) UpdateTask(in *task.UpdateTaskRequest) (*task.UpdateTaskResponse, error) {
	// todo: add your logic here and delete this line

	return &task.UpdateTaskResponse{}, nil
}
