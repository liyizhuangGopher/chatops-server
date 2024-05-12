package svc

import (
	"chatops-server/internal/group/model"
	"chatops-server/internal/group/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	GroupModel model.GroupsModel
	Conn       sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		GroupModel: model.NewGroupsModel(conn),
		Conn:       conn,
	}
}
