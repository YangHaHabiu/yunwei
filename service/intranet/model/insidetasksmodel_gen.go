// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/util/gconv"
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
	insideTasksFieldNames          = builder.RawFieldNames(&InsideTasks{})
	insideTasksRows                = strings.Join(insideTasksFieldNames, ",")
	insideTasksRowsExpectAutoSet   = strings.Join(stringx.Remove(insideTasksFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`", "`del_flag`"), ",")
	insideTasksRowsWithPlaceHolder = strings.Join(stringx.Remove(insideTasksFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`", "`del_flag`"), "=?,") + "=?"
)

type (
	insideTasksModel interface {
		Insert(ctx context.Context, data *InsideTasks) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*InsideTasks, error)
		Update(ctx context.Context, data *InsideTasks) error
		Delete(ctx context.Context, id int64) error
		FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]InsideTasksNew, error)
		DeleteSoft(ctx context.Context, id int64) error
		Count(ctx context.Context, filters ...interface{}) (int64, error)
		FindAll(ctx context.Context, filters ...interface{}) (*[]InsideTasksNew, error)
	}

	defaultInsideTasksModel struct {
		conn  sqlx.SqlConn
		table string
	}

	InsideTasks struct {
		Id          int64  `db:"id"`
		ProjectId   int64  `db:"project_id"`   // 项目id
		ClusterId   int64  `db:"cluster_id"`   // 集群id
		ServerId    int64  `db:"server_id"`    // 集群id
		VersionId   int64  `db:"version_id"`   // 集群id
		OperationId int64  `db:"operation_id"` //操作id
		TasksType   string `db:"tasks_type"`   // 任务类型：1：client 2：server 3：php
		Version     string `db:"version"`
		Cmd         string `db:"cmd"`        // 命令
		StartTime   int64  `db:"start_time"` // 开始时间
		EndTime     int64  `db:"end_time"`   // 结束时间
		Status      string `db:"status"`     // 状态 1：等待开始 2：执行中 3：执行失败 4：执行成功
		CreateBy    int64  `db:"create_by"`  // 创建人
		DelFlag     int64  `db:"del_flag"`   // 0：使用中 1：已删除
	}
	InsideTasksNew struct {
		Id           int64  `db:"id"`
		ProjectId    int64  `db:"project_id"`   // 项目id
		ClusterId    int64  `db:"cluster_id"`   // 集群id
		ServerId     int64  `db:"server_id"`    // 集群id
		VersionId    int64  `db:"version_id"`   // 集群id
		OperationId  int64  `db:"operation_id"` //操作id
		TasksType    string `db:"tasks_type"`   // 任务类型：1：client 2：server 3：php
		Version      string `db:"version"`
		Cmd          string `db:"cmd"`            // 命令
		StartTime    int64  `db:"start_time"`     // 开始时间
		EndTime      int64  `db:"end_time"`       // 结束时间
		Status       string `db:"status"`         // 状态 1：等待开始 2：执行中 3：执行失败 4：执行成功
		CreateBy     int64  `db:"create_by"`      // 创建人
		DelFlag      int64  `db:"del_flag"`       // 0：使用中 1：已删除
		ProjectCn    string `db:"project_cn"`     // 项目cn
		ClusterCn    string `db:"cluster_cn"`     // 集群cn
		ServerTitle  string `db:"server_title"`   // 服务器名
		ServerPath   string `db:"server_path"`    // 服务器地址信息
		ServerDescDb string `db:"server_desc_db"` // 服务器目标库
		VersionName  string `db:"version_name"`   // 版本名称
		VersionData  string `db:"version_data"`   // 版本数据
		VersionConf  string `db:"version_conf"`   // 版本配置
		VersionType  string `db:"version_type"`   // 版本类型 1：svn 2：git
		OperCn       string `db:"oper_cn"`        // 操作内容
		OperEn       string `db:"oper_en"`        // 操作内容
		OperType     string `db:"oper_type"`      // 操作类型 1：server 2：client
		NickName     string `db:"nick_name"`
		ProjectEn    string `db:"project_en"`
	}
)

func newInsideTasksModel(conn sqlx.SqlConn) *defaultInsideTasksModel {
	return &defaultInsideTasksModel{
		conn:  conn,
		table: "`inside_tasks`",
	}
}

var tasksCommonSQL = `SELECT
	%s
FROM
	(
	SELECT
		inside_tasks.*,
		project.project_cn,
		project.project_en,
 		label.label_name AS cluster_cn,
		inside_server.server_path,
		inside_server.server_title,
		inside_server.server_desc_db,
		inside_version.source_db,
		inside_version.version_name,
		inside_version.version_data,
		inside_version.version_conf,
		inside_version.version_type,
		inside_operation.oper_cn,
		inside_operation.oper_en,
		inside_operation.oper_type,
		sys_user.nick_name 
	FROM
		inside_tasks
		LEFT JOIN  project ON project.project_id = inside_tasks.project_id
 		LEFT JOIN (select label_id,label_name,del_flag from label union all select '0','0','0') label ON label_id = inside_tasks.cluster_id
		LEFT JOIN (select id,server_path,server_title,server_desc_db,del_flag from inside_server union all select '0','0','0','0','0') inside_server ON inside_tasks.server_id = inside_server.id
		LEFT JOIN (select id,source_db,version_name,version_data,version_conf,version_type,del_flag from inside_version union all select '0','0','0','0','0','0','0') inside_version ON inside_version.id = inside_tasks.version_id
		LEFT JOIN  (select id,oper_cn,oper_en,oper_type,del_flag from inside_operation union all select '0','0','0','0','0') inside_operation ON inside_tasks.operation_id = inside_operation.id
		LEFT JOIN sys_user ON sys_user.id = inside_tasks.create_by 
	WHERE
		 project.del_flag = 0 
 		 AND 	label.del_flag = 0 
		AND inside_server.del_flag = 0 
		AND inside_version.del_flag = 0 
		AND inside_tasks.del_flag = 0 
		AND inside_operation.del_flag = 0 
	) A 
	WHERE 1=1
	%s
	ORDER BY  start_time desc,tasks_type,project_id
	%s
`

func (m *defaultInsideTasksModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultInsideTasksModel) FindOne(ctx context.Context, id int64) (*InsideTasks, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", insideTasksRows, m.table)
	var resp InsideTasks
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

func (m *defaultInsideTasksModel) Insert(ctx context.Context, data *InsideTasks) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?,?,?, ?,?,?, ?, ?, ?, ?, ?, ? )", m.table, insideTasksRowsExpectAutoSet)
	var uid string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		uid = md.Get("uid")[0]
	}
	data.Status = "1"
	data.CreateBy = gconv.Int64(uid)
	data.StartTime = time.Now().Unix()

	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.ProjectId, data.ClusterId, data.ServerId, data.VersionId, data.OperationId, data.TasksType, data.Version, data.Cmd, data.StartTime, data.EndTime, data.Status, data.CreateBy)
	return ret, err
}

func (m *defaultInsideTasksModel) Update(ctx context.Context, data *InsideTasks) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, insideTasksRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ProjectId, data.ClusterId, data.ServerId, data.VersionId, data.OperationId, data.TasksType, data.Cmd, data.StartTime, data.EndTime, data.Status, data.CreateBy, data.Id)
	return err
}

func (m *defaultInsideTasksModel) tableName() string {
	return m.table
}

//分页条件查询
func (m *defaultInsideTasksModel) FindPageListByPage(ctx context.Context, page, pageSize int64, filters ...interface{}) (*[]InsideTasksNew, error) {

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query := fmt.Sprintf(tasksCommonSQL, "*", condition, "limit ? offset ?")
	var resp []InsideTasksNew
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
func (m *defaultInsideTasksModel) FindAll(ctx context.Context, filters ...interface{}) (*[]InsideTasksNew, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query := fmt.Sprintf(tasksCommonSQL, "*", condition, "")
	var resp []InsideTasksNew
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
func (m *defaultInsideTasksModel) Count(ctx context.Context, filters ...interface{}) (int64, error) {
	var condition string
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	query := fmt.Sprintf(tasksCommonSQL, "count(*)", condition, "")
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
func (m *defaultInsideTasksModel) DeleteSoft(ctx context.Context, id int64) error {
	query := fmt.Sprintf("update %s set `del_flag`=? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, globalkey.DelStateYes, id)
	return err
}
