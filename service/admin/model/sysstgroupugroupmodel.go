package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/common/xfilters"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysStgroupUgroupFieldNames          = builder.RawFieldNames(&SysStgroupUgroup{})
	sysStgroupUgroupRows                = strings.Join(sysStgroupUgroupFieldNames, ",")
	sysStgroupUgroupRowsExpectAutoSet   = strings.Join(stringx.Remove(sysStgroupUgroupFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysStgroupUgroupRowsWithPlaceHolder = strings.Join(stringx.Remove(sysStgroupUgroupFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysStgroupUgroupModel interface {
		Insert(ctx context.Context, data *SysStgroupUgroup) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysStgroupUgroup, error)
		Update(ctx context.Context, data *SysStgroupUgroup) error
		Delete(ctx context.Context, id int64) error
		TransactInsert(ctx context.Context, data *adminclient.PolicyAssociatedUsersReq, opts ...string) error
		FindAll(ctx context.Context, filters ...interface{}) (*[]SysStgroupUgroup, error)
	}

	defaultSysStgroupUgroupModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysStgroupUgroup struct {
		Id        int64 `db:"id"`
		StgroupId int64 `db:"stgroup_id"` // 决策组id
		UgroupId  int64 `db:"ugroup_id"`  // 用户组id
	}
)

func NewSysStgroupUgroupModel(conn sqlx.SqlConn) SysStgroupUgroupModel {
	return &defaultSysStgroupUgroupModel{
		conn:  conn,
		table: "`sys_stgroup_ugroup`",
	}
}

func (m *defaultSysStgroupUgroupModel) Insert(ctx context.Context, data *SysStgroupUgroup) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, sysStgroupUgroupRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.StgroupId, data.UgroupId)
	return ret, err
}

func (m *defaultSysStgroupUgroupModel) FindAll(ctx context.Context, filters ...interface{}) (*[]SysStgroupUgroup, error) {
	filter := xfilters.Xfilters(filters...)
	var conditions string
	if len(filter) != 0 {
		conditions = "where " + filter
	}
	query := fmt.Sprintf("select %s from %s %s", sysStgroupUgroupRows, m.table, conditions)
	var resp []SysStgroupUgroup
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

func (m *defaultSysStgroupUgroupModel) FindOne(ctx context.Context, id int64) (*SysStgroupUgroup, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysStgroupUgroupRows, m.table)
	var resp SysStgroupUgroup
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

func (m *defaultSysStgroupUgroupModel) Update(ctx context.Context, data *SysStgroupUgroup) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysStgroupUgroupRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.StgroupId, data.UgroupId, data.Id)
	return err
}

func (m *defaultSysStgroupUgroupModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysStgroupUgroupModel) TransactInsert(ctx context.Context, data *adminclient.PolicyAssociatedUsersReq, opts ...string) error {
	insertsql := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysStgroupUgroupRowsExpectAutoSet)
	if err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		stmt, err := session.Prepare(insertsql)
		if err != nil {
			return err
		}
		defer stmt.Close()
		tmp := fmt.Sprintf("delete from sys_stgroup_ugroup where stgroup_id = %d", data.StgroupId)
		if len(opts) != 0 {
			tmp = fmt.Sprintf("delete from sys_stgroup_ugroup where ugroup_id = %d", data.StgroupId)
		}
		_, err = m.conn.ExecCtx(ctx, tmp)
		if err != nil {
			return xerr.NewErrMsg("删除用户组与策略关联失败")
		}
		for _, v := range strings.Split(data.UgroupCheck, ",") {
			if v != "" {
				if len(opts) != 0 {
					_, err = stmt.ExecCtx(ctx, v, data.StgroupId)
				} else {
					_, err = stmt.ExecCtx(ctx, data.StgroupId, v)
				}

				if err != nil {
					return xerr.NewErrMsg("新增用户组与策略关联失败")
				}
			}
		}
		return nil
	}); err != nil {
		return xerr.NewErrMsg(fmt.Sprintf("%+v", err.Error()))
	}
	return nil
}
