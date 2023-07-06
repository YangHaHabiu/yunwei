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
	sysDictFieldNames                     = builder.RawFieldNames(&SysDict{})
	sysDictRows                           = strings.Join(sysDictFieldNames, ",")
	sysDictRowsExpectAutoSet              = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysDictRowsWithPlaceHolder            = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
	sysDictRowsWithPlaceHolderWitoutTypes = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`types`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysDictModel interface {
		Insert(ctx context.Context, data *SysDict) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysDict, error)
		Update(ctx context.Context, data *SysDict) error
		Delete(ctx context.Context, id int64) error
		FindAll(ctx context.Context, filters ...interface{}) (*[]SysDict, error)
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysDict, error)
		Count(ctx context.Context, filters ...interface{}) (int64, error)
	}

	defaultSysDictModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysDict struct {
		Id          int64  `db:"id"`          // 编号
		Value       string `db:"value"`       // 数据值
		Label       string `db:"label"`       // 标签名
		Types       string `db:"types"`       // 类型
		Pid         int64  `db:"pid"`         //pid
		Description string `db:"description"` // 描述
		Sort        int64  `db:"sort"`        // 排序（升序）

	}
)

func NewSysDictModel(conn sqlx.SqlConn) SysDictModel {
	return &defaultSysDictModel{
		conn:  conn,
		table: "`sys_dict`",
	}
}

func (m *defaultSysDictModel) Insert(ctx context.Context, data *SysDict) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?,?, ?, ?, ?)", m.table, sysDictRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Value, data.Label, data.Types, data.Pid, data.Description, data.Sort)
	return ret, err
}

func (m *defaultSysDictModel) FindOne(ctx context.Context, id int64) (*SysDict, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDictRows, m.table)
	var resp SysDict
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

func (m *defaultSysDictModel) FindAll(ctx context.Context, filters ...interface{}) (*[]SysDict, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " where " + filter
	}
	query := fmt.Sprintf("select %s from %s %s order by sort,id desc", sysDictRows, m.table, condition)
	var resp []SysDict
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
func (m *defaultSysDictModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysDict, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " where " + filter
	}
	query := fmt.Sprintf("select %s from %s %s order by sort,id desc limit ? offset ?", sysDictRows, m.table, condition)
	var resp []SysDict
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

func (m *defaultSysDictModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
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

func (m *defaultSysDictModel) Update(ctx context.Context, data *SysDict) error {
	one, err2 := m.FindOne(ctx, data.Id)
	if err2 != nil {
		return err2
	}

	if one.Types != data.Types {
		_, err := m.conn.ExecCtx(ctx, fmt.Sprintf("update %s set `types` = '%s' where `types` = ?", m.table, data.Types), one.Types)
		if err != nil {
			return err
		}
	}

	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysDictRowsWithPlaceHolderWitoutTypes)
	_, err := m.conn.ExecCtx(ctx, query, data.Value, data.Label, data.Pid, data.Description, data.Sort, data.Id)

	return err
}

func (m *defaultSysDictModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
