package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"
	"ywadmin-v3/common/globalkey"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysRoleFieldNames          = builder.RawFieldNames(&SysRole{})
	sysRoleRows                = strings.Join(sysRoleFieldNames, ",")
	sysRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_time`", "`last_update_time`"), ",")
	sysRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_by`", "`create_time`", "`last_update_time`"), "=?,") + "=?"
)

type (
	SysRoleModel interface {
		Insert(ctx context.Context, data *SysRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysRole, error)
		Update(ctx context.Context, data *SysRole) error
		Delete(ctx context.Context, id int64) error
		FindAll(ctx context.Context) (*[]SysRole, error)
		FindPageListByPage(ctx context.Context, page, pageSize int64) (*[]SysRole, error)
		Count(ctx context.Context) (int64, error)
		DeleteSoft(ctx context.Context, data *SysRole) error
		TransactDelete(ctx context.Context, data *SysRole) error
	}

	defaultSysRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysRole struct {
		Id             int64     `db:"id"`               // 编号
		Name           string    `db:"name"`             // 角色名称
		Remark         string    `db:"remark"`           // 备注
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  1：已删除  0：正常
	}
)

func NewSysRoleModel(conn sqlx.SqlConn) SysRoleModel {
	return &defaultSysRoleModel{
		conn:  conn,
		table: "`sys_role`",
	}
}

func (m *defaultSysRoleModel) Insert(ctx context.Context, data *SysRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, sysRoleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Remark, data.CreateBy, data.LastUpdateBy, data.DelFlag)
	return ret, err
}

func (m *defaultSysRoleModel) FindOne(ctx context.Context, id int64) (*SysRole, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleRows, m.table)
	var resp SysRole
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

func (m *defaultSysRoleModel) FindAll(ctx context.Context) (*[]SysRole, error) {

	query := fmt.Sprintf("select %s from %s where del_flag = %d ", sysRoleRows, m.table, globalkey.DelStateNo)
	var resp []SysRole
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
func (m *defaultSysRoleModel) FindPageListByPage(ctx context.Context, page, pageSize int64) (*[]SysRole, error) {

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s where del_flag = %d limit ? offset ?", sysRoleRows, m.table, globalkey.DelStateNo)
	var resp []SysRole
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

func (m *defaultSysRoleModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

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

func (m *defaultSysRoleModel) Update(ctx context.Context, data *SysRole) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Remark, data.LastUpdateBy, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysRoleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysRoleModel) DeleteSoft(ctx context.Context, data *SysRole) error {
	data.DelFlag = globalkey.DelStateYes
	data.LastUpdateTime = time.Now()
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, strings.Join([]string{
		"last_update_by",
		"last_update_time",
		"del_flag",
	}, "=?,")+"=?")
	_, err := m.conn.ExecCtx(ctx, query, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysRoleModel) TransactDelete(ctx context.Context, data *SysRole) error {

	updatesql := fmt.Sprintf("update %s set `last_update_by`=?,`last_update_time`=?,`del_flag`=? where `id` = ?", m.table)
	err := m.conn.Transact(func(session sqlx.Session) error {
		stmt, err := session.Prepare(updatesql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		// 返回任何错误都会回滚事务
		if _, err := stmt.ExecCtx(ctx, data.LastUpdateBy, time.Now(), globalkey.DelStateYes, data.Id); err != nil {
			logx.Errorf("update ugroup stmt exec: %s", err)
			return err
		}
		// delete sys_role_menu role_id
		_, err = m.conn.ExecCtx(ctx, "delete from sys_role_menu where `role_id` = ?", data.Id)
		if err != nil {
			logx.Errorf("delete sys_role_menu role_id: %s", err)
			return err
		}
		return nil
	})
	return err
}
