package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
	"ywadmin-v3/common/xfilters"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysLogFieldNames          = builder.RawFieldNames(&SysLog{})
	sysLogRows                = strings.Join(sysLogFieldNames, ",")
	sysLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysLogModel interface {
		Insert(ctx context.Context, data *SysLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysLog, error)
		Update(ctx context.Context, data *SysLog) error
		Delete(ctx context.Context, id int64) error
		FindAll(ctx context.Context) (*[]SysLog, error)
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysLog, error)
		Count(ctx context.Context, filters ...interface{}) (int64, error)
	}

	defaultSysLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysLog struct {
		Id         int64     `db:"id"`          // 编号
		UserName   string    `db:"user_name"`   // 用户名
		Operation  string    `db:"operation"`   // 用户操作
		Method     string    `db:"method"`      // 请求方法
		Params     string    `db:"params"`      // 请求参数
		Time       float32   `db:"time"`        // 执行时长(秒)
		Ip         string    `db:"ip"`          // IP地址
		CreateTime time.Time `db:"create_time"` // 创建时间
	}
)

func NewSysLogModel(conn sqlx.SqlConn) SysLogModel {
	return &defaultSysLogModel{
		conn:  conn,
		table: "`sys_log`",
	}
}

func (m *defaultSysLogModel) Insert(ctx context.Context, data *SysLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserName, data.Operation, data.Method, data.Params, data.Time, data.Ip)
	return ret, err
}

func (m *defaultSysLogModel) FindOne(ctx context.Context, id int64) (*SysLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLogRows, m.table)
	var resp SysLog
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

func (m *defaultSysLogModel) FindAll(ctx context.Context) (*[]SysLog, error) {

	query := fmt.Sprintf("select %s from %s", sysLogRows, m.table)
	var resp []SysLog
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

//根据页码分页查询数据
func (m *defaultSysLogModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysLog, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " where " + filter
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s %s order by create_time desc limit ? offset ?", sysLogRows, m.table, condition)

	var resp []SysLog
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

func (m *defaultSysLogModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " where " + filter
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

func (m *defaultSysLogModel) Update(ctx context.Context, data *SysLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserName, data.Operation, data.Method, data.Params, data.Time, data.Ip, data.Id)
	return err
}

func (m *defaultSysLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
