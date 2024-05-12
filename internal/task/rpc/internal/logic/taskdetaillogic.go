package logic

import (
	"context"

	"chatops-server/internal/task/rpc/internal/svc"
	"chatops-server/internal/task/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskDetailLogic {
	return &TaskDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskDetailLogic) TaskDetail(in *task.TaskDetailRequest) (*task.TaskDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &task.TaskDetailResponse{}, nil
}
