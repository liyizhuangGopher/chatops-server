package logic

import (
	"context"
	"errors"

	"chatops-server/internal/group/rpc/group"
	"chatops-server/internal/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type groupDbInfo struct {
	Name         string `db:"name"`
	ConversionID string `db:"conversion_id"`
}

func (l *GroupListLogic) GroupList(in *group.GroupListRequest) (*group.GroupListResponse, error) {
	// 查询群组信息
	var groupListFromDb []groupDbInfo
	querySQL := "select name, conversion_id from `groups`"
	err := l.svcCtx.Conn.QueryRowsCtx(l.ctx, &groupListFromDb, querySQL)
	if errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorw("group.rpc.Conn.QueryRowsCtx failed", logx.LogField{Key: "err", Value: sqlx.ErrNotFound.Error()})
		return nil, errors.New("查询群组列表失败")
	}
	if err != nil {
		logx.Errorw("group.rpc.Conn.QueryRowsCtx failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	logx.Infow("group.rpc.Conn.QueryRowsCtx success")
	var groupList []*group.GroupInfo
	for _, groupFromDb := range groupListFromDb {
		groupResponse := group.GroupInfo{
			Name:         groupFromDb.Name,
			ConversionID: groupFromDb.ConversionID,
		}
		groupList = append(groupList, &groupResponse)
	}
	return &group.GroupListResponse{
		Message:   "ok",
		GroupList: groupList,
	}, nil
}
