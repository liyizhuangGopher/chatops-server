// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	departmentFieldNames          = builder.RawFieldNames(&Department{})
	departmentRows                = strings.Join(departmentFieldNames, ",")
	departmentRowsExpectAutoSet   = strings.Join(stringx.Remove(departmentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	departmentRowsWithPlaceHolder = strings.Join(stringx.Remove(departmentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	departmentModel interface {
		Insert(ctx context.Context, data *Department) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Department, error)
		FindOneByName(ctx context.Context, name string) (*Department, error)
		Update(ctx context.Context, data *Department) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDepartmentModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Department struct {
		Id   int64  `db:"id"`
		Name string `db:"name"` // 部门名称
	}
)

func newDepartmentModel(conn sqlx.SqlConn) *defaultDepartmentModel {
	return &defaultDepartmentModel{
		conn:  conn,
		table: "`department`",
	}
}

func (m *defaultDepartmentModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDepartmentModel) FindOne(ctx context.Context, id int64) (*Department, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", departmentRows, m.table)
	var resp Department
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDepartmentModel) FindOneByName(ctx context.Context, name string) (*Department, error) {
	var resp Department
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", departmentRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDepartmentModel) Insert(ctx context.Context, data *Department) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, departmentRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name)
	return ret, err
}

func (m *defaultDepartmentModel) Update(ctx context.Context, newData *Department) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, departmentRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.Id)
	return err
}

func (m *defaultDepartmentModel) tableName() string {
	return m.table
}
