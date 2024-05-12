package logic

import (
	"context"
	"errors"

	"chatops-server/internal/user/rpc/internal/svc"
	"chatops-server/internal/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type userDbInfo struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
}

func (l *UserListLogic) UserList(in *user.UserListRequest) (*user.UserListResponse, error) {
	// 查询用户信息
	// 由于自动生成的model代码没有对应的查询语句，需要执行原生sql
	var userListFromDb []userDbInfo
	querySQL := `select id, username from user;`
	err := l.svcCtx.Conn.QueryRowsCtx(l.ctx, &userListFromDb, querySQL)
	if errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorw("user.rpc.Conn.Exec failed", logx.LogField{Key: "err", Value: sqlx.ErrNotFound.Error()})
		return nil, errors.New("查询用户列表失败")
	}
	if err != nil {
		logx.Errorw("user.rpc.Conn.Exec failed", logx.LogField{Key: "err", Value: err})
		return nil, err
	}
	logx.Infow("user.rpc.Conn.Exec success")
	var userList []*user.UserInfo
	for _, userFromDb := range userListFromDb {
		userResponse := user.UserInfo{
			UserID:   userFromDb.ID,
			Username: userFromDb.Username,
		}
		userList = append(userList, &userResponse)
	}
	return &user.UserListResponse{
		UserList: userList,
	}, nil
}
