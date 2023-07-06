package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
	"ywadmin-v3/common/xfilters"
)

var (
	sysStrategyFieldNames          = builder.RawFieldNames(&SysStrategy{})
	sysStrategyRows                = strings.Join(sysStrategyFieldNames, ",")
	sysStrategyRowsExpectAutoSet   = strings.Join(stringx.Remove(sysStrategyFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysStrategyRowsWithPlaceHolder = strings.Join(stringx.Remove(sysStrategyFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysStrategyModel interface {
		Insert(ctx context.Context, data *SysStrategy) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysStrategy, error)
		Update(ctx context.Context, data *SysStrategy) error
		Delete(ctx context.Context, id int64) error
		Count(ctx context.Context, filters ...interface{}) (int64, error)
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysStrategy, error)
		FindAll(ctx context.Context, filters ...interface{}) (*[]SysStrategy, error)
	}

	defaultSysStrategyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysStrategy struct {
		Id       int64  `db:"id"`
		StName   string `db:"st_name"`   // 名称
		StSort   int64  `db:"st_sort"`   // 排序
		StPid    int64  `db:"st_pid"`    // 上级pid
		StLevel  int64  `db:"st_level"`  // 授权值
		StUrls   string `db:"st_urls"`   // urls
		StRemark string `db:"st_remark"` // 备注
		StIsAuth int64  `db:"st_is_auth"`
	}
)

func NewSysStrategyModel(conn sqlx.SqlConn) SysStrategyModel {
	return &defaultSysStrategyModel{
		conn:  conn,
		table: "`sys_strategy`",
	}
}

func (m *defaultSysStrategyModel) Insert(ctx context.Context, data *SysStrategy) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysStrategyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.StName, data.StSort, data.StPid, data.StLevel, data.StUrls, data.StRemark)
	return ret, err
}

func (m *defaultSysStrategyModel) FindOne(ctx context.Context, id int64) (*SysStrategy, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysStrategyRows, m.table)
	var resp SysStrategy
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

func (m *defaultSysStrategyModel) FindAll(ctx context.Context, filters ...interface{}) (*[]SysStrategy, error) {
	filter := xfilters.Xfilters(filters...)
	var conditions string
	if len(filter) != 0 {
		conditions = "where " + filter
	}
	query := fmt.Sprintf("select %s from %s %s", sysStrategyRows, m.table, conditions)
	var resp []SysStrategy
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
func (m *defaultSysStrategyModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " where " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf("select count(*) as count from %s %s", m.table, condition)
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

//根据页码分页查询数据
func (m *defaultSysStrategyModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysStrategy, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = "where " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf("select %s from %s %s limit ? offset ?", sysDeptRows, m.table, condition)
	var resp []SysStrategy
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

func (m *defaultSysStrategyModel) Update(ctx context.Context, data *SysStrategy) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysStrategyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.StName, data.StSort, data.StPid, data.StLevel, data.StUrls, data.StRemark, data.Id)
	return err
}

func (m *defaultSysStrategyModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
