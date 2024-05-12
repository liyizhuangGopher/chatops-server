package logic

import (
	"context"

	"chatops-server/internal/task/rpc/internal/svc"
	"chatops-server/internal/task/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskLogic {
	return &AddTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTaskLogic) AddTask(in *task.AddTaskRequest) (*task.AddTaskResponse, error) {
	// todo: add your logic here and delete this line

	return &task.AddTaskResponse{}, nil
}
