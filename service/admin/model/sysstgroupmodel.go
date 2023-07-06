package model

import (
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xfilters"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysStgroupFieldNames          = builder.RawFieldNames(&SysStgroup{})
	sysStgroupRows                = strings.Join(sysStgroupFieldNames, ",")
	sysStgroupRowsExpectAutoSet   = strings.Join(stringx.Remove(sysStgroupFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysStgroupRowsWithPlaceHolder = strings.Join(stringx.Remove(sysStgroupFieldNames, "`id`", "`create_by`", "`create_time`", "`update_time`", "`del_flag`"), "=?,") + "=?"
)

type (
	SysStgroupModel interface {
		Insert(ctx context.Context, data *SysStgroup) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysStgroup, error)
		Update(ctx context.Context, data *SysStgroup) error
		Delete(ctx context.Context, id int64) error
		Count(ctx context.Context, filters ...interface{}) (int64, error)
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysStgroup, error)
		FindAll(ctx context.Context) (*[]SysStgroup, error)
		DeleteSoft(ctx context.Context, data *SysStgroup) error
	}

	defaultSysStgroupModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysStgroup struct {
		Id             int64     `db:"id"`
		StName         string    `db:"st_name"` // 策略名称
		StJson         string    `db:"st_json"` // 策略json
		StRemark       string    `db:"st_remark"`
		CreateBy       string    `db:"create_by"`
		CreateTime     time.Time `db:"create_time"`
		LastUpdateBy   string    `db:"last_update_by"`
		LastUpdateTime time.Time `db:"last_update_time"`
		DelFlag        int64     `db:"del_flag"` // 是否删除  1：已删除  0：正常
	}
)

func NewSysStgroupModel(conn sqlx.SqlConn) SysStgroupModel {
	return &defaultSysStgroupModel{
		conn:  conn,
		table: "`sys_stgroup`",
	}
}

func (m *defaultSysStgroupModel) Insert(ctx context.Context, data *SysStgroup) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysStgroupRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.StName, data.StJson, data.StRemark, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag)
	return ret, err
}

func (m *defaultSysStgroupModel) FindOne(ctx context.Context, id int64) (*SysStgroup, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysStgroupRows, m.table)
	var resp SysStgroup
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

func (m *defaultSysStgroupModel) FindAll(ctx context.Context) (*[]SysStgroup, error) {

	query := fmt.Sprintf("select %s from %s where del_flag = %d ", sysStgroupRows, m.table, globalkey.DelStateNo)
	var resp []SysStgroup
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
func (m *defaultSysStgroupModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysStgroup, error) {
	filter := xfilters.Xfilters(filters...)
	var condition string
	if len(filter) != 0 {
		condition = "and " + filter
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s  where del_flag = %d %s limit ? offset ?", sysStgroupRows, m.table, globalkey.DelStateNo, condition)
	var resp []SysStgroup
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

func (m *defaultSysStgroupModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
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
func (m *defaultSysStgroupModel) DeleteSoft(ctx context.Context, data *SysStgroup) error {
	data.DelFlag = globalkey.DelStateYes
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		data.LastUpdateBy = md.Get("uname")[0]
	}
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join([]string{
		"last_update_by",
		"del_flag",
	}, "=?,")+"=?")
	_, err := m.conn.ExecCtx(ctx, query, data.LastUpdateBy, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysStgroupModel) Update(ctx context.Context, data *SysStgroup) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysStgroupRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.StName, data.StJson, data.StRemark, data.LastUpdateBy, data.LastUpdateTime, data.Id)
	return err
}

func (m *defaultSysStgroupModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
