package svc

import (
	"chatops-server/internal/task/model"
	"chatops-server/internal/task/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	TaskModel  model.TaskModel
	UserModel  model.UserModel
	GroupModel model.GroupsModel
	Conn       sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		TaskModel:  model.NewTaskModel(conn),
		UserModel:  model.NewUserModel(conn),
		GroupModel: model.NewGroupsModel(conn),
		Conn:       conn,
	}
}
