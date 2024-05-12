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
	systemFieldNames          = builder.RawFieldNames(&System{})
	systemRows                = strings.Join(systemFieldNames, ",")
	systemRowsExpectAutoSet   = strings.Join(stringx.Remove(systemFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	systemRowsWithPlaceHolder = strings.Join(stringx.Remove(systemFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	systemModel interface {
		Insert(ctx context.Context, data *System) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*System, error)
		Update(ctx context.Context, data *System) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSystemModel struct {
		conn  sqlx.SqlConn
		table string
	}

	System struct {
		Id               int64 `db:"id"`
		DetectDuration   int64 `db:"detect_duration"`   // 探测时间
		HistoryDurantion int64 `db:"history_durantion"` // 历史记录保存时间
	}
)

func newSystemModel(conn sqlx.SqlConn) *defaultSystemModel {
	return &defaultSystemModel{
		conn:  conn,
		table: "`system`",
	}
}

func (m *defaultSystemModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSystemModel) FindOne(ctx context.Context, id int64) (*System, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", systemRows, m.table)
	var resp System
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

func (m *defaultSystemModel) Insert(ctx context.Context, data *System) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, systemRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DetectDuration, data.HistoryDurantion)
	return ret, err
}

func (m *defaultSystemModel) Update(ctx context.Context, data *System) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, systemRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DetectDuration, data.HistoryDurantion, data.Id)
	return err
}

func (m *defaultSystemModel) tableName() string {
	return m.table
}
