// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xfilters"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	insideOperationFieldNames          = builder.RawFieldNames(&InsideOperation{})
	insideOperationRows                = strings.Join(insideOperationFieldNames, ",")
	insideOperationRowsExpectAutoSet   = strings.Join(stringx.Remove(insideOperationFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`", "`del_flag`"), ",")
	insideOperationRowsWithPlaceHolder = strings.Join(stringx.Remove(insideOperationFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`", "`del_flag`"), "=?,") + "=?"
)

type (
	insideOperationModel interface {
		Insert(ctx context.Context, data *InsideOperation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*InsideOperation, error)
		Update(ctx context.Context, data *InsideOperation) error
		Delete(ctx context.Context, id int64) error
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]InsideOperationNew, error)
		DeleteSoft(ctx context.Context, id int64) error
		Count(ctx context.Context, filters ...interface{}) (int64, error)
		FindAll(ctx context.Context, filters ...interface{}) (*[]InsideOperationNew, error)
	}

	defaultInsideOperationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	InsideOperation struct {
		Id        int64  `db:"id"`
		ProjectId int64  `db:"project_id"` // 项目id
		OperCn    string `db:"oper_cn"`    // 操作内容
		OperEn    string `db:"oper_en"`    // 操作内容
		OperType  string `db:"oper_type"`  // 操作类型 1：server 2：client
		Sort      int64  `db:"sort"`       // 排序
		DelFlag   int64  `db:"del_flag"`   // 0：使用中 1：已删除
	}
	InsideOperationNew struct {
		Id        int64  `db:"id"`
		ProjectId int64  `db:"project_id"` // 项目id
		ProjectCn string `db:"project_cn"` // 项目id
		OperCn    string `db:"oper_cn"`    // 操作内容
		OperEn    string `db:"oper_en"`    // 操作内容
		OperType  string `db:"oper_type"`  // 操作类型 1：server 2：client
		Sort      int64  `db:"sort"`       // 排序
		DelFlag   int64  `db:"del_flag"`   // 0：使用中 1：已删除
	}
)

func newInsideOperationModel(conn sqlx.SqlConn) *defaultInsideOperationModel {
	return &defaultInsideOperationModel{
		conn:  conn,
		table: "`inside_operation`",
	}
}

var operationCommonSQL = `
SELECT
	%s 
FROM
	(
	SELECT
		inside_operation.*,
		project.project_cn 
	FROM
		inside_operation
		LEFT JOIN project ON inside_operation.project_id = project.project_id 
	WHERE
		project.del_flag = 0 
		AND inside_operation.del_flag = 0 
	) A 
WHERE
	1 = 1
%s
ORDER BY oper_type,sort,project_id
%s
`

func (m *defaultInsideOperationModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultInsideOperationModel) FindOne(ctx context.Context, id int64) (*InsideOperation, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", insideOperationRows, m.table)
	var resp InsideOperation
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

func (m *defaultInsideOperationModel) Insert(ctx context.Context, data *InsideOperation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ? )", m.table, insideOperationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProjectId, data.OperCn, data.OperEn, data.OperType, data.Sort)
	return ret, err
}

func (m *defaultInsideOperationModel) Update(ctx context.Context, data *InsideOperation) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, insideOperationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ProjectId, data.OperCn, data.OperEn, data.OperType, data.Sort, data.Id)
	return err
}

func (m *defaultInsideOperationModel) tableName() string {
	return m.table
}

//分页条件查询
func (m *defaultInsideOperationModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]InsideOperationNew, error) {

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query := fmt.Sprintf(operationCommonSQL, "*", condition, "limit ? offset ?")

	var resp []InsideOperationNew
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

//条件查询所有
func (m *defaultInsideOperationModel) FindAll(ctx context.Context, filters ...interface{}) (*[]InsideOperationNew, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query := fmt.Sprintf(operationCommonSQL, "*", condition, "")
	var resp []InsideOperationNew
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

//条件统计
func (m *defaultInsideOperationModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf(operationCommonSQL, "count(*)", condition, "")
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

//软删除
func (m *defaultInsideOperationModel) DeleteSoft(ctx context.Context, id int64) error {
	query := fmt.Sprintf("update %s set `del_flag`=? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, globalkey.DelStateYes, id)
	return err
}