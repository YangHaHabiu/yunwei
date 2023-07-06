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
)

var (
	sysRoleMenuFieldNames          = builder.RawFieldNames(&SysRoleMenu{})
	sysRoleMenuRows                = strings.Join(sysRoleMenuFieldNames, ",")
	sysRoleMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysRoleMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysRoleMenuModel interface {
		Insert(ctx context.Context, data *SysRoleMenu) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysRoleMenu, error)
		Update(ctx context.Context, data *SysRoleMenu) error
		Delete(ctx context.Context, id int64) error
		FindByRoleId(ctx context.Context, RoleId int64) (*[]SysRoleMenu, error)
	}

	defaultSysRoleMenuModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysRoleMenu struct {
		Id     int64 `db:"id"`      // 编号
		RoleId int64 `db:"role_id"` // 角色ID
		MenuId int64 `db:"menu_id"` // 菜单ID
	}
)

func NewSysRoleMenuModel(conn sqlx.SqlConn) SysRoleMenuModel {
	return &defaultSysRoleMenuModel{
		conn:  conn,
		table: "`sys_role_menu`",
	}
}

func (m *defaultSysRoleMenuModel) FindByRoleId(ctx context.Context, RoleId int64) (*[]SysRoleMenu, error) {

	query := fmt.Sprintf("select %s from %s where role_id = ?", sysRoleMenuRows, m.table)
	var resp []SysRoleMenu
	err := m.conn.QueryRowsCtx(ctx, &resp, query, RoleId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysRoleMenuModel) Insert(ctx context.Context, data *SysRoleMenu) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysRoleMenuRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RoleId, data.MenuId)
	return ret, err
}

func (m *defaultSysRoleMenuModel) FindOne(ctx context.Context, id int64) (*SysRoleMenu, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleMenuRows, m.table)
	var resp SysRoleMenu
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

func (m *defaultSysRoleMenuModel) Update(ctx context.Context, data *SysRoleMenu) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleMenuRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.RoleId, data.MenuId, data.Id)
	return err
}

func (m *defaultSysRoleMenuModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `role_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
