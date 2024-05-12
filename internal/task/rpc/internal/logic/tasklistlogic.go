package logic

import (
	"context"

	"chatops-server/internal/task/rpc/internal/svc"
	"chatops-server/internal/task/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskListLogic) TaskList(in *task.TaskListRequest) (*task.TaskListResponse, error) {
	// todo: add your logic here and delete this line

	return &task.TaskListResponse{}, nil
}
