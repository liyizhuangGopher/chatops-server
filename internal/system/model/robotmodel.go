package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RobotModel = (*customRobotModel)(nil)

type (
	// RobotModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRobotModel.
	RobotModel interface {
		robotModel
		withSession(session sqlx.Session) RobotModel
	}

	customRobotModel struct {
		*defaultRobotModel
	}
)

// NewRobotModel returns a model for the database table.
func NewRobotModel(conn sqlx.SqlConn) RobotModel {
	return &customRobotModel{
		defaultRobotModel: newRobotModel(conn),
	}
}

func (m *customRobotModel) withSession(session sqlx.Session) RobotModel {
	return NewRobotModel(sqlx.NewSqlConnFromSession(session))
}
