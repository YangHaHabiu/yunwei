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
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/common/xfilters"
	"ywadmin-v3/service/admin/rpc/adminclient"
)

var (
	sysUserRoleFieldNames          = builder.RawFieldNames(&SysUserRole{})
	sysUserRoleRows                = strings.Join(sysUserRoleFieldNames, ",")
	sysUserRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserRoleFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysUserRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserRoleFieldNames, "`id`", "`create_by`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysUserRoleModel interface {
		Insert(ctx context.Context, data *SysUserRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysUserRole, error)
		Update(ctx context.Context, data *SysUserRole) error
		//根据字段删除对应数据
		DeleteByCustomFiled(ctx context.Context, filter ...interface{}) error
		//批量插入
		BulkInserter(data []*SysUserRole) error
		FindAll(ctx context.Context, filters ...interface{}) (*[]SysUserRole, error)
		TransactInsert(ctx context.Context, data *adminclient.RoleAssignmentUserReq, opts ...string) error
	}

	defaultSysUserRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUserRole struct {
		Id     int64 `db:"id"`      // 编号
		UserId int64 `db:"user_id"` // 用户ID
		RoleId int64 `db:"role_id"` // 角色ID
	}
)

func NewSysUserRoleModel(conn sqlx.SqlConn) SysUserRoleModel {
	return &defaultSysUserRoleModel{
		conn:  conn,
		table: "`sys_user_role`",
	}
}

func (m *defaultSysUserRoleModel) TransactInsert(ctx context.Context, data *adminclient.RoleAssignmentUserReq, opts ...string) error {
	insertsql := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysUserRoleRowsExpectAutoSet)
	if err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		stmt, err := session.Prepare(insertsql)
		if err != nil {
			return err
		}
		defer stmt.Close()
		tmp := fmt.Sprintf("delete from sys_user_role where user_id = %d", data.RoleId)
		if len(opts) != 0 {
			tmp = fmt.Sprintf("delete from sys_user_role where role_id = %d", data.RoleId)
		}
		_, err = m.conn.ExecCtx(ctx, tmp)
		if err != nil {
			return xerr.NewErrMsg("删除角色与用户关联失败")
		}
		for _, v := range strings.Split(data.UserCheck, ",") {
			if v != "" {
				if len(opts) != 0 {
					_, err = stmt.ExecCtx(ctx, v, data.RoleId)
				} else {
					_, err = stmt.ExecCtx(ctx, data.RoleId, v)
				}

				if err != nil {
					return xerr.NewErrMsg("新增角色与用户关联失败")
				}
			}
		}
		return nil
	}); err != nil {
		return xerr.NewErrMsg(fmt.Sprintf("%+v", err.Error()))
	}
	return nil
}

func (m *defaultSysUserRoleModel) Insert(ctx context.Context, data *SysUserRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysUserRoleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.RoleId)
	return ret, err
}

func (m *defaultSysUserRoleModel) BulkInserter(data []*SysUserRole) error {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysUserRoleRowsExpectAutoSet)
	bulkInserter, err := sqlx.NewBulkInserter(m.conn, query)
	if err != nil {
		return err
	}
	for _, v := range data {
		if err = bulkInserter.Insert(v.UserId, v.RoleId); err != nil {
			return err
		}
	}
	bulkInserter.Flush()
	return err
}

func (m *defaultSysUserRoleModel) FindOne(ctx context.Context, id int64) (*SysUserRole, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRoleRows, m.table)
	var resp SysUserRole
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

func (m *defaultSysUserRoleModel) Update(ctx context.Context, data *SysUserRole) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRoleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.RoleId, data.Id)
	return err
}

func (m *defaultSysUserRoleModel) DeleteByCustomFiled(ctx context.Context, filters ...interface{}) error {
	query := fmt.Sprintf("delete from %s where %s", m.table, xfilters.Xfilters(filters...))
	_, err := m.conn.ExecCtx(ctx, query)
	return err
}

func (m *defaultSysUserRoleModel) FindAll(ctx context.Context, filters ...interface{}) (*[]SysUserRole, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " where " + filter
	}
	query := fmt.Sprintf("select * from %s %s", m.table, condition)

	var resp []SysUserRole
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
