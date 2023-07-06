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
	sysDeptFieldNames          = builder.RawFieldNames(&SysDept{})
	sysDeptRows                = strings.Join(sysDeptFieldNames, ",")
	sysDeptRowsExpectAutoSet   = strings.Join(stringx.Remove(sysDeptFieldNames, "`id`", "`create_time`", "`last_update_time`"), ",")
	sysDeptRowsWithPlaceHolder = strings.Join(stringx.Remove(sysDeptFieldNames, "`id`", "`create_by`", "`create_time`", "`last_update_time`"), "=?,") + "=?"
)

type (
	SysDeptModel interface {
		//新增数据
		Insert(ctx context.Context, data *SysDept) (sql.Result, error)
		//根据id查询
		FindOne(ctx context.Context, id int64) (*SysDept, error)
		//更新
		Update(ctx context.Context, data *SysDept) error
		//删除数据
		Delete(ctx context.Context, id int64) error
		//软删除数据
		DeleteSoft(ctx context.Context, data *SysDept) error
		//查询所有数据
		FindAll(ctx context.Context) (*[]SysDept, error)
		//分页查询数据
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysDept, error)
		//统计
		Count(ctx context.Context, filters ...interface{}) (int64, error)
	}

	defaultSysDeptModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysDept struct {
		Id             int64     `db:"id"`               // 编号
		Name           string    `db:"name"`             // 机构名称
		ParentId       int64     `db:"parent_id"`        // 上级机构ID，一级机构为0
		OrderNum       int64     `db:"order_num"`        // 排序
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  1：已删除  0：正常
	}
)

func NewSysDeptModel(conn sqlx.SqlConn) SysDeptModel {
	return &defaultSysDeptModel{
		conn:  conn,
		table: "`sys_dept`",
	}
}

func (m *defaultSysDeptModel) Insert(ctx context.Context, data *SysDept) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysDeptRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.ParentId, data.OrderNum, data.CreateBy, data.LastUpdateBy, data.DelFlag)
	return ret, err
}

func (m *defaultSysDeptModel) FindOne(ctx context.Context, id int64) (*SysDept, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDeptRows, m.table)
	var resp SysDept
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

//查询所有数据
func (m *defaultSysDeptModel) FindAll(ctx context.Context) (*[]SysDept, error) {
	query := fmt.Sprintf("select %s from %s where del_flag = %d order by order_num", sysDeptRows, m.table, globalkey.DelStateNo)
	var resp []SysDept
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
func (m *defaultSysDeptModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]SysDept, error) {

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf("select %s from %s where del_flag = %d %s limit ? offset ?", sysDeptRows, m.table, globalkey.DelStateNo, condition)
	var resp []SysDept
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

func (m *defaultSysDeptModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " where " + xfilters.Xfilters(filters...)
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

func (m *defaultSysDeptModel) Update(ctx context.Context, data *SysDept) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysDeptRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.ParentId, data.OrderNum, data.LastUpdateBy, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysDeptModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysDeptModel) DeleteSoft(ctx context.Context, data *SysDept) error {
	data.DelFlag = globalkey.DelStateYes
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		data.LastUpdateBy = md.Get("uname")[0]
	}
	query := fmt.Sprintf("update %s set del_flag=? ,last_update_by=? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.DelFlag, data.LastUpdateBy, data.Id)
	return err
}
