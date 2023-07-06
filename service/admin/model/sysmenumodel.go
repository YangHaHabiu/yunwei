package model

import (
	"context"
	"database/sql"
	"fmt"
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
	sysMenuFieldNames          = builder.RawFieldNames(&SysMenu{})
	sysMenuRows                = strings.Join(sysMenuFieldNames, ",")
	sysMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysMenuModel interface {
		Insert(ctx context.Context, data *SysMenu) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysMenu, error)
		Update(ctx context.Context, data *SysMenu) error
		Delete(ctx context.Context, id int64) error
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, filters ...interface{}) (*[]SysMenu, error)
		FindAllByUserId(ctx context.Context, userId int64) (*[]SysMenu, error)
		DeleteSoft(ctx context.Context, data *SysMenu) error
		FindAllOperationList(ctx context.Context, filters ...interface{}) (*[]SysMenuList, error)
	}

	defaultSysMenuModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysMenu struct {
		Id             int64     `db:"id"`        // 编号
		Name           string    `db:"name"`      // 菜单名称
		ParentId       int64     `db:"parent_id"` // 父菜单ID，一级菜单为0
		Url            string    `db:"url"`
		Perms          string    `db:"perms"`            // 授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)
		Tp             int64     `db:"tp"`               // 类型   0：目录   1：菜单   2：按钮
		Icon           string    `db:"icon"`             // 菜单图标
		OrderNum       int64     `db:"order_num"`        // 排序
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  1：已删除  0：正常
		VuePath        string    `db:"vue_path"`         // vue系统的path
		VueComponent   string    `db:"vue_component"`    // vue的页面
		VueIcon        string    `db:"vue_icon"`         // vue的图标
		VueRedirect    string    `db:"vue_redirect"`     // vue的路由重定向
		TableName      string    `db:"table_name"`
		KeepAlive      string    `db:"keep_alive"`
		IsShow         string    `db:"is_show"`
	}

	SysMenuList struct {
		ParentId       int64  `db:"parent_id"`   // 父菜单ID，一级菜单为0
		ParentName     string `db:"parent_name"` // 父菜单ID，一级菜单为0
		ParentUrl      string `db:"parent_url"`
		ParentTp       int64  `db:"parent_tp"`        // 类型   0：目录   1：菜单   2：按钮
		ParentOrderNum int64  `db:"parent_order_num"` // 排序
		Id             int64  `db:"id"`               // 编号
		Name           string `db:"name"`             // 菜单名称
		Url            string `db:"url"`
		Tp             int64  `db:"tp"`        // 类型   0：目录   1：菜单   2：按钮
		OrderNum       int64  `db:"order_num"` // 排序

	}
)

func NewSysMenuModel(conn sqlx.SqlConn) SysMenuModel {
	return &defaultSysMenuModel{
		conn:  conn,
		table: "`sys_menu`",
	}
}

func (m *defaultSysMenuModel) Insert(ctx context.Context, data *SysMenu) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysMenuRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.ParentId, data.Url, data.Perms, data.Tp, data.Icon, data.OrderNum, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.VuePath, data.VueComponent, data.VueIcon, data.VueRedirect)
	return ret, err
}

func (m *defaultSysMenuModel) FindOne(ctx context.Context, id int64) (*SysMenu, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_flag = 0 limit 1", sysMenuRows, m.table)
	var resp SysMenu
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

func (m *defaultSysMenuModel) FindAll(ctx context.Context, filters ...interface{}) (*[]SysMenu, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	query := fmt.Sprintf("select %s from %s where del_flag = %d %s", sysMenuRows, m.table, globalkey.DelStateNo, condition)
	var resp []SysMenu
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

var operationSql = `
select parent_id,parent_name,parent_url,parent_tp,parent_order_num,id,tp,order_num,GROUP_CONCAT(name) as name,GROUP_CONCAT(url) as url from (
select A.parent_id,parent_name,parent_url,parent_tp,parent_order_num,id,name,url,tp,order_num
from (
	select id as parent_id,name as parent_name,url as parent_url,tp as parent_tp,order_num as parent_order_num 
	from sys_menu 
	where tp=2
	) A,(
	select id,name,parent_id,url,tp,order_num
	from sys_menu
	where tp=3
	) B
where A.parent_id=B.parent_id
order by parent_order_num,order_num
) C
%s
`

func (m *defaultSysMenuModel) FindAllOperationList(ctx context.Context, filters ...interface{}) (*[]SysMenuList, error) {
	var condition string
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " where " + filter
	}
	query := fmt.Sprintf(operationSql, condition)
	var resp []SysMenuList
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

func (m *defaultSysMenuModel) FindAllByUserId(ctx context.Context, userId int64) (*[]SysMenu, error) {
	query := "select sm.* from sys_user_role sur inner join sys_role sr on sur.role_id = sr.id inner join sys_role_menu srm on sr.id = srm.role_id inner join sys_menu sm on srm.menu_id = sm.id where sm.tp in (0,1) and sur.user_id=? GROUP BY sm.id"
	var resp []SysMenu
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysMenuModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where del_flag = 0", m.table)

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

func (m *defaultSysMenuModel) Update(ctx context.Context, data *SysMenu) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysMenuRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.ParentId, data.Url, data.Perms, data.Tp, data.Icon, data.OrderNum, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.VuePath, data.VueComponent, data.VueIcon, data.VueRedirect, data.Id)
	return err
}

func (m *defaultSysMenuModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysMenuModel) DeleteSoft(ctx context.Context, data *SysMenu) error {
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
