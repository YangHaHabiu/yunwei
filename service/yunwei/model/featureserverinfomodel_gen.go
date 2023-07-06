// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"ywadmin-v3/common/xfilters"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	featureServerInfoFieldNames          = builder.RawFieldNames(&FeatureServerInfo{})
	featureServerInfoRows                = strings.Join(featureServerInfoFieldNames, ",")
	featureServerInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(featureServerInfoFieldNames, "`feature_server_id`", "`create_time`", "`update_time`", "`del_flag`"), ",")
	featureServerInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(featureServerInfoFieldNames, "`feature_server_id`", "`create_time`", "`update_time`", "`del_flag`"), "=?,") + "=?"
)

type (
	featureServerInfoModel interface {
		Insert(ctx context.Context, data *FeatureServerInfo) (sql.Result, error)
		FindOne(ctx context.Context, featureServerId int64) (*FeatureServerInfo, error)
		Update(ctx context.Context, data *FeatureServerInfo) error
		Delete(ctx context.Context, featureServerId int64) error
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]FeatureServerInfoNew, error)
		DeleteSoft(ctx context.Context, id int64) error
		Count(ctx context.Context, filters ...interface{}) (int64, error)
	}

	defaultFeatureServerInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	FeatureServerInfo struct {
		FeatureServerId   int64  `db:"feature_server_id"`   // 功能服信息ID
		ProjectId         int64  `db:"project_id"`          // 项目ID
		FeatureServerInfo string `db:"feature_server_info"` // 功能服相关信息
		DelFlag           int64  `db:"del_flag"`            // 删除状态：0:未删除(数据使用中);1:已删除(回收)
		Remark            string `db:"remark"`
	}

	FeatureServerInfoNew struct {
		FeatureServerId   int64  `db:"feature_server_id"`   // 功能服信息ID
		ProjectId         int64  `db:"project_id"`          // 项目ID
		FeatureServerInfo string `db:"feature_server_info"` // 功能服相关信息
		DelFlag           int64  `db:"del_flag"`            // 删除状态：0:未删除(数据使用中);1:已删除(回收)
		ProjectCn         string `db:"project_cn"`
		ProjectEn         string `db:"project_en"`
		Remark            string `db:"remark"`
	}
)

func newFeatureServerInfoModel(conn sqlx.SqlConn) *defaultFeatureServerInfoModel {
	return &defaultFeatureServerInfoModel{
		conn:  conn,
		table: "`feature_server_info`",
	}
}

func (m *defaultFeatureServerInfoModel) Insert(ctx context.Context, data *FeatureServerInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, featureServerInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProjectId, data.FeatureServerInfo, data.Remark)
	return ret, err
}

func (m *defaultFeatureServerInfoModel) FindOne(ctx context.Context, featureServerId int64) (*FeatureServerInfo, error) {
	query := fmt.Sprintf("select %s from %s where `feature_server_id` = ? limit 1", featureServerInfoRows, m.table)
	var resp FeatureServerInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, featureServerId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFeatureServerInfoModel) Update(ctx context.Context, data *FeatureServerInfo) error {
	query := fmt.Sprintf("update %s set %s where `feature_server_id` = ?", m.table, featureServerInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ProjectId, data.FeatureServerInfo, data.Remark, data.FeatureServerId)
	return err
}

func (m *defaultFeatureServerInfoModel) Delete(ctx context.Context, featureServerId int64) error {
	query := fmt.Sprintf("delete from %s where `feature_server_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, featureServerId)
	return err
}

func (m *defaultFeatureServerInfoModel) tableName() string {
	return m.table
}

func (m *defaultFeatureServerInfoModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]FeatureServerInfoNew, error) {
	query := `SELECT
	fsi.*,
	p.project_cn,
	project_en 
FROM
	feature_server_info AS fsi,
	project p 
WHERE
	fsi.project_id = p.project_id 
	AND p.del_flag = 0
%s
order by del_flag,p.project_id desc,feature_server_id desc
limit ? offset ?
`

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query = fmt.Sprintf(query, condition)
	//fmt.Println(query)
	var resp []FeatureServerInfoNew
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

func (m *defaultFeatureServerInfoModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	query := `SELECT COUNT(*) count from (
SELECT
	fsi.*,
	p.project_cn,
	project_en 
FROM
	feature_server_info AS fsi,
	project p 
WHERE
	fsi.project_id = p.project_id 
	AND p.del_flag = 0
%s
order by feature_server_id desc
) C
`
	query = fmt.Sprintf(query, condition)
	//fmt.Println(query)
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
func (m *defaultFeatureServerInfoModel) DeleteSoft(ctx context.Context, id int64) error {
	status := 0
	one, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	if one.DelFlag == 0 {
		status = 1
	}

	query := fmt.Sprintf("update %s set `del_flag`=? where `feature_server_id` = ?", m.table)
	_, err = m.conn.ExecCtx(ctx, query, status, id)
	return err
}
