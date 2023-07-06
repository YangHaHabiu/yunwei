// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xfilters"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	insideTasksLogsFieldNames          = builder.RawFieldNames(&InsideTasksLogs{})
	insideTasksLogsRows                = strings.Join(insideTasksLogsFieldNames, ",")
	insideTasksLogsRowsExpectAutoSet   = strings.Join(stringx.Remove(insideTasksLogsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	insideTasksLogsRowsWithPlaceHolder = strings.Join(stringx.Remove(insideTasksLogsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	insideTasksLogsModel interface {
		Insert(ctx context.Context, data *InsideTasksLogs) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*InsideTasksLogs, error)
		Update(ctx context.Context, data *InsideTasksLogs) error
		Delete(ctx context.Context, id int64) error
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]InsideTasksLogs, error)
		DeleteSoft(ctx context.Context, id int64) error
		Count(ctx context.Context, filters ...interface{}) (int64, error)
		FindAll(ctx context.Context, filters ...interface{}) (*[]InsideTasksLogs, error)
	}

	defaultInsideTasksLogsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	InsideTasksLogs struct {
		Id      int64  `db:"id"`
		TasksId int64  `db:"tasks_id"` // 任务id
		Content string `db:"content"`  // 任务日志
	}
)

func newInsideTasksLogsModel(conn sqlx.SqlConn) *defaultInsideTasksLogsModel {
	return &defaultInsideTasksLogsModel{
		conn:  conn,
		table: "`inside_tasks_logs`",
	}
}

func (m *defaultInsideTasksLogsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultInsideTasksLogsModel) FindOne(ctx context.Context, id int64) (*InsideTasksLogs, error) {
	query := fmt.Sprintf("select %s from %s where `tasks_id` = ? limit 1", insideTasksLogsRows, m.table)
	var resp InsideTasksLogs
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInsideTasksLogsModel) Insert(ctx context.Context, data *InsideTasksLogs) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, insideTasksLogsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.TasksId, data.Content)
	return ret, err
}

func (m *defaultInsideTasksLogsModel) Update(ctx context.Context, data *InsideTasksLogs) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, insideTasksLogsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.TasksId, data.Content, data.Id)
	return err
}

func (m *defaultInsideTasksLogsModel) tableName() string {
	return m.table
}

//分页条件查询
func (m *defaultInsideTasksLogsModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]InsideTasksLogs, error) {
	query := `SELECT * from %s where 1=1
%s
limit ? offset ?
`
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query = fmt.Sprintf(query, m.table, condition)

	var resp []InsideTasksLogs
	err := m.conn.QueryRowsCtx(ctx, &resp, query, pageSize, offset)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//条件查询所有
func (m *defaultInsideTasksLogsModel) FindAll(ctx context.Context, filters ...interface{}) (*[]InsideTasksLogs, error) {
	query := "select * from %s where 1=1 %s"
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query = fmt.Sprintf(query, m.table, condition)
	var resp []InsideTasksLogs
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//条件统计
func (m *defaultInsideTasksLogsModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf("select count(*) as count from %s where del_flag = %d %s", m.table, globalkey.DelStateNo, condition)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

//软删除
func (m *defaultInsideTasksLogsModel) DeleteSoft(ctx context.Context, id int64) error {
	query := fmt.Sprintf("update %s set `del_flag`=? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, globalkey.DelStateYes, id)
	return err
}
