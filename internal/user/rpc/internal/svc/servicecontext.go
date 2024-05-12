package svc

import (
	"chatops-server/internal/user/model"
	"chatops-server/internal/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	UserModel       model.UserModel
	DepartmentModel model.DepartmentModel
	Conn            sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:          c,
		UserModel:       model.NewUserModel(conn),
		DepartmentModel: model.NewDepartmentModel(conn),
		Conn:            conn,
	}
}
