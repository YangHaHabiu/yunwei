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
	sysLoginLogFieldNames          = builder.RawFieldNames(&SysLoginLog{})
	sysLoginLogRows                = strings.Join(sysLoginLogFieldNames, ",")
	sysLoginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysLoginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysLoginLogModel interface {
		Insert(ctx context.Context, data *SysLoginLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysLoginLog, error)
		Update(ctx context.Context, data *SysLoginLog) error
		Delete(ctx context.Context, id int64) error
		FindAll(ctx context.Context) (*[]SysLoginLog, error)
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysLoginLog, error)
		Count(ctx context.Context, filters ...interface{}) (int64, error)
		UpdateFieldStatusByUname(ctx context.Context, uName, status string) error
	}

	defaultSysLoginLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysLoginLog struct {
		Id         int64     `db:"id"`          // 编号
		UserName   string    `db:"user_name"`   // 用户名
		Status     string    `db:"status"`      // 登录状态
		Ip         string    `db:"ip"`          // IP地址
		CreateTime time.Time `db:"create_time"` // 创建时间
	}
)

func NewSysLoginLogModel(conn sqlx.SqlConn) SysLoginLogModel {
	return &defaultSysLoginLogModel{
		conn:  conn,
		table: "`sys_login_log`",
	}
}

func (m *defaultSysLoginLogModel) Insert(ctx context.Context, data *SysLoginLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, sysLoginLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserName, data.Status, data.Ip)
	return ret, err
}

func (m *defaultSysLoginLogModel) FindOne(ctx context.Context, id int64) (*SysLoginLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLoginLogRows, m.table)
	var resp SysLoginLog
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

func (m *defaultSysLoginLogModel) FindAll(ctx context.Context) (*[]SysLoginLog, error) {
	query := fmt.Sprintf("select %s from %s ", sysLoginLogRows, m.table)
	var resp []SysLoginLog
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
func (m *defaultSysLoginLogModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysLoginLog, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " where " + filter
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s %s order by create_time desc limit ? offset ?", sysLoginLogRows, m.table, condition)
	fmt.Println(query)
	var resp []SysLoginLog
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

func (m *defaultSysLoginLogModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
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

func (m *defaultSysLoginLogModel) Update(ctx context.Context, data *SysLoginLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLoginLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserName, data.Status, data.Ip, data.Id)
	return err
}

func (m *defaultSysLoginLogModel) UpdateFieldStatusByUname(ctx context.Context, uName, status string) error {
	query := fmt.Sprintf("update %s set `status` = ? where `user_name` = ? and `status` = 'online'", m.table)
	_, err := m.conn.ExecCtx(ctx, query, status, uName)
	return err
}

func (m *defaultSysLoginLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
