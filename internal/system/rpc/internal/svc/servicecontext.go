package svc

import (
	"chatops-server/internal/system/model"
	"chatops-server/internal/system/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RobotModel  model.RobotModel
	SystemModel model.SystemModel
	Conn        sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		RobotModel:  model.NewRobotModel(conn),
		SystemModel: model.NewSystemModel(conn),
		Conn:        conn,
	}
}
