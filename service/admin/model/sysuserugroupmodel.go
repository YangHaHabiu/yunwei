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
	sysUserUgroupFieldNames          = builder.RawFieldNames(&SysUserUgroup{})
	sysUserUgroupRows                = strings.Join(sysUserUgroupFieldNames, ",")
	sysUserUgroupRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserUgroupFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysUserUgroupRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserUgroupFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysUserUgroupModel interface {
		Insert(ctx context.Context, data *SysUserUgroup) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysUserUgroup, error)
		Update(ctx context.Context, data *SysUserUgroup) error
		DeleteByCustomFiled(ctx context.Context, filters ...interface{}) error
		BulkInserter(data []*SysUserUgroup) error
		FindAll(ctx context.Context, filters ...interface{}) (*[]SysUserUgroup, error)
		TransactInsert(ctx context.Context, data *adminclient.UgroupAssignmentUserReq, opts ...string) error
	}

	defaultSysUserUgroupModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUserUgroup struct {
		Id       int64 `db:"id"`
		UserId   int64 `db:"user_id"`   // 用户id
		UgroupId int64 `db:"ugroup_id"` // 用户组id
	}
)

func NewSysUserUgroupModel(conn sqlx.SqlConn) SysUserUgroupModel {
	return &defaultSysUserUgroupModel{
		conn:  conn,
		table: "`sys_user_ugroup`",
	}
}

func (m *defaultSysUserUgroupModel) Insert(ctx context.Context, data *SysUserUgroup) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysUserUgroupRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UgroupId)
	return ret, err
}

func (m *defaultSysUserUgroupModel) FindOne(ctx context.Context, id int64) (*SysUserUgroup, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserUgroupRows, m.table)
	var resp SysUserUgroup
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

func (m *defaultSysUserUgroupModel) TransactInsert(ctx context.Context, data *adminclient.UgroupAssignmentUserReq, opts ...string) error {
	insertsql := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysUserUgroupRowsExpectAutoSet)
	if err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		stmt, err := session.Prepare(insertsql)
		if err != nil {
			return err
		}
		defer stmt.Close()
		tmp := fmt.Sprintf("delete from sys_user_ugroup where user_id = %d", data.UgroupId)
		if len(opts) != 0 {
			tmp = fmt.Sprintf("delete from sys_user_ugroup where ugroup_id = %d", data.UgroupId)
		}
		_, err = m.conn.ExecCtx(ctx, tmp)
		if err != nil {
			return xerr.NewErrMsg("删除用户组与用户关联失败")
		}
		for _, v := range strings.Split(data.UserCheck, ",") {
			if v != "" {
				if len(opts) != 0 {
					_, err = stmt.ExecCtx(ctx, v, data.UgroupId)
				} else {
					_, err = stmt.ExecCtx(ctx, data.UgroupId, v)
				}

				if err != nil {
					return xerr.NewErrMsg("新增用户组与用户关联失败")
				}
			}
		}
		return nil
	}); err != nil {
		return xerr.NewErrMsg(fmt.Sprintf("%+v", err.Error()))
	}
	return nil
}

func (m *defaultSysUserUgroupModel) FindAll(ctx context.Context, filters ...interface{}) (*[]SysUserUgroup, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " where " + filter
	}
	query := fmt.Sprintf("select * from %s %s", m.table, condition)

	var resp []SysUserUgroup
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

func (m *defaultSysUserUgroupModel) Update(ctx context.Context, data *SysUserUgroup) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserUgroupRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UgroupId, data.UserId, data.Id)
	return err
}

func (m *defaultSysUserUgroupModel) DeleteByCustomFiled(ctx context.Context, filters ...interface{}) error {
	query := fmt.Sprintf("delete from %s where 1=1 and %s", m.table, xfilters.Xfilters(filters...))
	fmt.Println(query)
	_, err := m.conn.ExecCtx(ctx, query)
	return err
}

func (m *defaultSysUserUgroupModel) BulkInserter(data []*SysUserUgroup) error {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysUserUgroupRowsExpectAutoSet)
	bulkInserter, err := sqlx.NewBulkInserter(m.conn, query)
	if err != nil {
		return err
	}
	for _, v := range data {
		if err = bulkInserter.Insert(v.UserId, v.UgroupId); err != nil {
			return err
		}
	}
	bulkInserter.Flush()
	return err
}
