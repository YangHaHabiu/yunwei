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
	sysUgroupFieldNames          = builder.RawFieldNames(&SysUgroup{})
	sysUgroupRows                = strings.Join(sysUgroupFieldNames, ",")
	sysUgroupRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUgroupFieldNames, "`id`", "`create_time`", "`last_update_time`"), ",")
	sysUgroupRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUgroupFieldNames, "`id`", "`create_by`", "`create_time`", "`last_update_time`"), "=?,") + "=?"
)

type (
	SysUgroupModel interface {
		Insert(ctx context.Context, data *SysUgroup) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysUgroup, error)
		Update(ctx context.Context, data *SysUgroup) error
		Delete(ctx context.Context, id int64) error
		Count(ctx context.Context) (int64, error)
		FindPageListByPage(ctx context.Context, page, pageSize int64) (*[]SysUgroup, error)
		FindAll(ctx context.Context) (*[]SysUgroup, error)
		DeleteSoft(ctx context.Context, data *SysUgroup) error
		TransactDelete(ctx context.Context, data *SysUgroup) error
	}

	defaultSysUgroupModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUgroup struct {
		Id             int64     `db:"id"`      // 组id
		Ugname         string    `db:"ug_name"` // 组名
		UgJson         string    `db:"ug_json"` // 组json值
		CreateBy       string    `db:"create_by"`
		CreateTime     time.Time `db:"create_time"`
		LastUpdateBy   string    `db:"last_update_by"`
		LastUpdateTime time.Time `db:"last_update_time"`
		DelFlag        int64     `db:"del_flag"`
	}
)

func NewSysUgroupModel(conn sqlx.SqlConn) SysUgroupModel {
	return &defaultSysUgroupModel{
		conn:  conn,
		table: "`sys_ugroup`",
	}
}

func (m *defaultSysUgroupModel) Insert(ctx context.Context, data *SysUgroup) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, sysUgroupRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Ugname, data.UgJson, data.CreateBy, data.LastUpdateBy, data.DelFlag)
	return ret, err
}

func (m *defaultSysUgroupModel) FindOne(ctx context.Context, id int64) (*SysUgroup, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUgroupRows, m.table)
	var resp SysUgroup
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

func (m *defaultSysUgroupModel) Update(ctx context.Context, data *SysUgroup) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUgroupRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Ugname, data.UgJson, data.LastUpdateBy, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysUgroupModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysUgroupModel) FindAll(ctx context.Context) (*[]SysUgroup, error) {

	query := fmt.Sprintf("select %s from %s where del_flag = %d ", sysUgroupRows, m.table, globalkey.DelStateNo)
	var resp []SysUgroup
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
func (m *defaultSysUgroupModel) FindPageListByPage(ctx context.Context, page, pageSize int64) (*[]SysUgroup, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s where del_flag = %d limit ? offset ?", sysUgroupRows, m.table, globalkey.DelStateNo)
	var resp []SysUgroup
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

func (m *defaultSysUgroupModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where del_flag = %d", m.table, globalkey.DelStateNo)
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
func (m *defaultSysUgroupModel) DeleteSoft(ctx context.Context, data *SysUgroup) error {
	data.DelFlag = globalkey.DelStateYes
	data.LastUpdateTime = time.Now()
	query := fmt.Sprintf("update %s set last_update_by=?,last_update_time=?,del_flag=? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysUgroupModel) TransactDelete(ctx context.Context, data *SysUgroup) error {

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

		// delete sys_stgroup_ugroup ugroup_id
		_, err = m.conn.ExecCtx(ctx, "delete from sys_stgroup_ugroup where `ugroup_id` = ?", data.Id)
		if err != nil {
			logx.Errorf("delete sys_stgroup_ugroup ugroup_id: %s", err)
			return err
		}

		return nil
	})
	return err
}
