package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SystemModel = (*customSystemModel)(nil)

type (
	// SystemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemModel.
	SystemModel interface {
		systemModel
		withSession(session sqlx.Session) SystemModel
	}

	customSystemModel struct {
		*defaultSystemModel
	}
)

// NewSystemModel returns a model for the database table.
func NewSystemModel(conn sqlx.SqlConn) SystemModel {
	return &customSystemModel{
		defaultSystemModel: newSystemModel(conn),
	}
}

func (m *customSystemModel) withSession(session sqlx.Session) SystemModel {
	return NewSystemModel(sqlx.NewSqlConnFromSession(session))
}
