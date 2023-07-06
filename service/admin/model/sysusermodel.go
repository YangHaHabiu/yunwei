package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/common/xfilters"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysUserFieldNames          = builder.RawFieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`last_update_time`"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`name`", "`del_flag`", "`salt`", "`password`", "`create_by`", "`create_time`", "`last_update_time`"), "=?,") + "=?"
)

type (
	SysUserModel interface {
		Insert(ctx context.Context, data *SysUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysUser, error)
		FindOneByName(ctx context.Context, name string) (*SysUser, error)
		Update(ctx context.Context, data *SysUser) error
		Delete(ctx context.Context, id int64) error
		FindPageListByPage(ctx context.Context, page, pageSize int64, in *adminclient.UserListReq, filters ...interface{}) (*[]SysUserList, error)
		Count(ctx context.Context, in *adminclient.UserListReq, filters ...interface{}) (int64, error)
		DeleteSoft(ctx context.Context, data *SysUser) error
		UpdateStatus(ctx context.Context, data *SysUser) error
		UpdatePassword(ctx context.Context, data *SysUser) error
		UpdatePersonalInfo(ctx context.Context, data *SysUser) error
		TransactInsert(ctx context.Context, data *adminclient.UserAddReq) error
		SelectStrategyInfoByUname(ctx context.Context, uname string) (*[]StrategyList, error)
		FindAll(ctx context.Context, in *adminclient.UserListReq, filters ...interface{}) (*[]SysUserList, error)
		TransactDelete(ctx context.Context, data *SysUser) error
		FindOneOnlyWithName(ctx context.Context, name string) (*SysUser, error)
		UserBatchEditItems(ctx context.Context, in *adminclient.UserBatchEditItemsReq) error
	}

	defaultSysUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUser struct {
		Id             int64  `db:"id"`               // 编号
		Name           string `db:"name"`             // 用户名
		NickName       string `db:"nick_name"`        // 昵称
		Avatar         string `db:"avatar"`           // 头像
		Password       string `db:"password"`         // 密码
		Salt           string `db:"salt"`             // 加密盐
		Email          string `db:"email"`            // 邮箱
		Mobile         string `db:"mobile"`           // 手机号
		Status         int64  `db:"status"`           // 状态 1：正常 2：禁用
		DeptId         int64  `db:"dept_id"`          // 机构ID
		CreateBy       string `db:"create_by"`        // 创建人
		CreateTime     string `db:"create_time"`      // 创建时间
		LastUpdateBy   string `db:"last_update_by"`   // 更新人
		LastUpdateTime string `db:"last_update_time"` // 更新时间
		DelFlag        int64  `db:"del_flag"`         // 是否删除  1：已删除  0：正常

	}
	SysUserList struct {
		Id             int64  `db:"id"`               // 编号
		Name           string `db:"name"`             // 用户名
		NickName       string `db:"nick_name"`        // 昵称
		Avatar         string `db:"avatar"`           // 头像
		Password       string `db:"password"`         // 密码
		Salt           string `db:"salt"`             // 加密盐
		Email          string `db:"email"`            // 邮箱
		Mobile         string `db:"mobile"`           // 手机号
		Status         int64  `db:"status"`           // 状态  0：禁用   1：正常
		DeptId         int64  `db:"dept_id"`          // 机构ID
		CreateBy       string `db:"create_by"`        // 创建人
		CreateTime     string `db:"create_time"`      // 创建时间
		LastUpdateBy   string `db:"last_update_by"`   // 更新人
		LastUpdateTime string `db:"last_update_time"` // 更新时间
		DelFlag        int64  `db:"del_flag"`         // 是否删除  1：已删除  0：正常
		DeptName       string `db:"dept_name"`
		RoleName       string `db:"role_name"`
		RoleIds        string `db:"role_ids"`
		UgroupIds      string `db:"ugroup_ids"`
		ProjectIds     string `db:"project_ids"`
		UgroupNames    string `db:"ugroup_names"`
	}

	StrategyList struct {
		SysUserId     int64  `db:"sys_user_id"`
		SysUserName   string `db:"sys_user_name"`
		StgroupStJson string `db:"stgroup_st_json"`
	}
)

func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &defaultSysUserModel{
		conn:  conn,
		table: "`sys_user`",
	}
}

var commonSql = `SELECT
    su.id,
	su.name,
	su.nick_name,
	su.avatar,
	su.password,
	su.salt,
	su.email,
	su.mobile,
	su.status,
	su.dept_id,
	su.create_by,
	date_format(su.create_time,'%%Y-%%m-%%d %%H:%%i:%%s') create_time ,
	su.last_update_by,
	date_format(su.last_update_time,'%%Y-%%m-%%d %%H:%%i:%%s') last_update_time ,
	su.del_flag,
	ifnull(sd.NAME, '') AS dept_name,
	GROUP_CONCAT(DISTINCT ifnull(sr.id,'')) as role_ids,
	GROUP_CONCAT(DISTINCT ifnull(sup.project_id,'')) as project_ids,
	GROUP_CONCAT(DISTINCT ifnull(sug.id,'')) as ugroup_ids,
	GROUP_CONCAT(DISTINCT ifnull(sug.ug_name,'')) as ugroup_names,
	GROUP_CONCAT(DISTINCT ifnull(sr.remark,'')) as role_name
FROM
	sys_user su
	LEFT JOIN sys_user_role sur ON su.id = sur.user_id
	LEFT JOIN sys_role sr ON sur.role_id = sr.id
	LEFT JOIN sys_dept sd ON su.dept_id = sd.id
	LEFT JOIN sys_user_ugroup suug ON  suug.user_id=su.id
	LEFT JOIN sys_ugroup sug ON  sug.id=suug.ugroup_id
	LEFT JOIN sys_user_project sup ON sup.user_id=su.id
WHERE
	su.del_flag = 0
%s
GROUP BY su.id

`

//批量新增（删除）用户项目操作
func (m *defaultSysUserModel) UserBatchEditItems(ctx context.Context, in *adminclient.UserBatchEditItemsReq) error {
	if err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		var (
			sql       string
			sqlList   = make([]string, 0)
			tmps      string
			condition string
		)
		if in.Operate == "add" {
			sql = "INSERT INTO sys_user_project(user_id,project_id) VALUES "
			tmps = "(%s,%s)"
			condition = " , "

		} else if in.Operate == "del" {
			sql = "delete from sys_user_project where "
			tmps = "( user_id=%s and project_id=%s )"
			condition = " or "
		}

		for _, v := range strings.Split(in.UserIds, ",") {
			if in.ProjectIds != "" {
				for _, v1 := range strings.Split(in.ProjectIds, ",") {
					sqlList = append(sqlList, fmt.Sprintf(tmps, v, v1))
				}
			}
		}
		sql += strings.Join(sqlList, condition)
		_, err := m.conn.ExecCtx(ctx, sql)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

//特殊处理条件having
func (m *defaultSysUserModel) handleListInfo(in *adminclient.UserListReq) string {
	total := make([]string, 0)
	for _, v := range []string{
		in.ProjectIds + ":" + "project_ids", in.RoleIds + ":" + "role_ids", in.UgroupIds + ":" + "ugroup_ids",
	} {
		tmpIds := make([]string, 0)
		tmpObj := strings.Split(v, ":")
		if tmpObj[0] != "" {
			for _, v1 := range strings.Split(tmpObj[0], ",") {
				tmpIds = append(tmpIds, fmt.Sprintf("FIND_IN_SET(%s,%s)>=1", v1, tmpObj[1]))
			}
		}
		if len(tmpIds) != 0 {
			total = append(total, "("+strings.Join(tmpIds, " or ")+")")
		}
	}
	return strings.Join(total, " and ")
}

func (m *defaultSysUserModel) Insert(ctx context.Context, data *SysUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysUserRowsExpectAutoSet)

	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.NickName, data.Avatar, data.Password, data.Salt, data.Email, data.Mobile, data.Status, data.DeptId, data.CreateBy, data.LastUpdateBy, data.DelFlag)
	return ret, err
}

func (m *defaultSysUserModel) FindOne(ctx context.Context, id int64) (*SysUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_flag = 0 limit 1", sysUserRows, m.table)
	var resp SysUser
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

func (m *defaultSysUserModel) FindOneOnlyWithName(ctx context.Context, name string) (*SysUser, error) {
	var resp SysUser
	query := fmt.Sprintf("select %s from %s where `name` = ?  limit 1", sysUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) FindOneByName(ctx context.Context, name string) (*SysUser, error) {
	var resp SysUser
	query := fmt.Sprintf("select %s from %s where `name` = ? and del_flag = %d and status = 1 limit 1", sysUserRows, m.table, globalkey.DelStateNo)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) FindPageListByPage(ctx context.Context, page, pageSize int64, in *adminclient.UserListReq, filters ...interface{}) (*[]SysUserList, error) {
	sql := commonSql + `%s
limit ? offset ?
`
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * pageSize

	var (
		condition string
		extend    string
	)
	filter := xfilters.Xfilters(filters...)
	if len(filter) != 0 {
		condition = " and " + filter
	}
	info := m.handleListInfo(in)
	if info != "" {
		extend = "having " + info
	}
	query := fmt.Sprintf(sql, condition, extend)
	//fmt.Println(query)
	var resp []SysUserList
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

//根据用户返回策略信息
func (m *defaultSysUserModel) SelectStrategyInfoByUname(ctx context.Context, uname string) (*[]StrategyList, error) {
	query := `select sys_user_id,sys_user_name,concat('[',group_concat(stgroup_st_json),']') stgroup_st_json
from (
		select sys_user_id,sys_user_name,sys_user_nick_name,sys_ugroup_id,sys_ugroup_name,
			concat(ifnull(ug_st.sys_stgroup_st_name,''),
				if(length(trim(ifnull(u_st.sys_stgroup_st_name,'')))=0,'',
					if(length(trim(ifnull(ug_st.sys_stgroup_st_name,'')))=0,u_st.sys_stgroup_st_name,concat('+',u_st.sys_stgroup_st_name))
				)
			) as stgroup_st_name,
			concat(ifnull(ug_st.sys_stgroup_st_json,''),
				if(length(trim(ifnull(u_st.sys_stgroup_st_json,'')))=0,'',
					if(length(trim(ifnull(ug_st.sys_stgroup_st_json,'')))=0,u_st.sys_stgroup_st_json,concat(',',u_st.sys_stgroup_st_json))
				)
			) as stgroup_st_json,
			concat(ifnull(ug_st.st_identify,''),
				if(length(trim(ifnull(u_st.st_identify,'')))=0,'',
					if(length(trim(ifnull(ug_st.st_identify,'')))=0,u_st.st_identify,concat('+',u_st.st_identify))
				)
			) as st_identify
		from
			(
				select sys_user.id as sys_user_id,sys_user.name as sys_user_name,sys_user.nick_name as sys_user_nick_name,
					sys_ugroup.id as sys_ugroup_id,sys_ugroup.ug_name as sys_ugroup_name
				from sys_user left join sys_user_ugroup on sys_user.id=sys_user_ugroup.user_id 
					left join sys_ugroup on sys_user_ugroup.ugroup_id=sys_ugroup.id
			) ug left join
			(
				select sys_stgroup.id as sys_stgroup_id,sys_stgroup.st_name as sys_stgroup_st_name,
					sys_stgroup.st_json as sys_stgroup_st_json,sys_stgroup.st_remark as sys_stgroup_st_remark,
					sys_stgroup_ugroup.ugroup_id as sys_stgroup_ugroup_id,'ugroup' as st_identify
				from sys_stgroup left join sys_stgroup_ugroup on sys_stgroup.id=sys_stgroup_ugroup.stgroup_id
			) ug_st on sys_ugroup_id=sys_stgroup_ugroup_id
			left join
			(
				select sys_stgroup.id as sys_stgroup_id,sys_stgroup.st_name as sys_stgroup_st_name,
					sys_stgroup.st_json as sys_stgroup_st_json,sys_stgroup.st_remark as sys_stgroup_st_remark,
					sys_stgroup_user.user_id as sys_stgroup_user_id,'user' as st_identify
				from sys_stgroup left join sys_stgroup_user on sys_stgroup.id=sys_stgroup_user.stgroup_id
			) u_st on sys_user_id=sys_stgroup_user_id
) A
where  sys_user_name = ?
group by sys_user_id
`
	var resp []StrategyList
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uname)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) Count(ctx context.Context, in *adminclient.UserListReq, filters ...interface{}) (int64, error) {
	var (
		condition string
		extend    string
	)
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	query := `SELECT count(*) FROM (
SELECT
	su.*,
	ifnull(sd.NAME, '') AS dept_name,
	GROUP_CONCAT(DISTINCT ifnull(sr.id,'')) as role_ids,
	GROUP_CONCAT(DISTINCT ifnull(sup.project_id,'')) as project_ids,
	GROUP_CONCAT(DISTINCT ifnull(sug.id,'')) as ugroup_ids,
	GROUP_CONCAT(DISTINCT ifnull(sug.ug_name,'')) as ugroup_names,
	GROUP_CONCAT(DISTINCT ifnull(sr.remark,'')) as role_name
FROM
	sys_user su
	LEFT JOIN sys_user_role sur ON su.id = sur.user_id
	LEFT JOIN sys_role sr ON sur.role_id = sr.id
	LEFT JOIN sys_dept sd ON su.dept_id = sd.id
	LEFT JOIN sys_user_ugroup suug ON  suug.user_id=su.id
	LEFT JOIN sys_ugroup sug ON  sug.id=suug.ugroup_id
	LEFT JOIN sys_user_project sup ON sup.user_id=su.id
WHERE

	su.del_flag =0
%s
GROUP BY su.id
%s
) b
`
	info := m.handleListInfo(in)
	if info != "" {
		extend = "having " + info
	}
	query = fmt.Sprintf(query, condition, extend)
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

func (m *defaultSysUserModel) FindAll(ctx context.Context, in *adminclient.UserListReq, filters ...interface{}) (*[]SysUserList, error) {
	var (
		condition string
		extend    string
	)
	if len(xfilters.Xfilters(filters...)) != 0 {
		condition = " and " + xfilters.Xfilters(filters...)
	}
	sql := commonSql + "%s"
	info := m.handleListInfo(in)
	if info != "" {
		extend = "having " + info
	}
	query := fmt.Sprintf(sql, condition, extend)

	var resp []SysUserList
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

func (m *defaultSysUserModel) Update(ctx context.Context, data *SysUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.NickName, data.Avatar, data.Email, data.Mobile, data.Status, data.DeptId, data.LastUpdateBy, data.Id)
	return err
}

func (m *defaultSysUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysUserModel) DeleteSoft(ctx context.Context, data *SysUser) error {
	data.DelFlag = globalkey.DelStateYes

	query := fmt.Sprintf("update %s set `last_update_by`=?,`del_flag`=? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.LastUpdateBy, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysUserModel) TransactDelete(ctx context.Context, data *SysUser) error {

	updatesql := fmt.Sprintf("update %s set `last_update_by`=?,`del_flag`=? where `id` = ?", m.table)
	err := m.conn.Transact(func(session sqlx.Session) error {
		stmt, err := session.Prepare(updatesql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		// 返回任何错误都会回滚事务
		if _, err := stmt.ExecCtx(ctx, data.LastUpdateBy, globalkey.DelStateYes, data.Id); err != nil {
			logx.Errorf("update userinfo stmt exec: %s", err)
			return err
		}

		// delete sys_stgroup_user user_id
		_, err = m.conn.ExecCtx(ctx, "delete from sys_stgroup_user where `user_id` = ?", data.Id)
		if err != nil {
			logx.Errorf("delete sys_stgroup_user user_id: %s", err)
			return err
		}
		// delete sys_user_project user_id
		_, err = m.conn.ExecCtx(ctx, "delete from sys_user_project where `user_id` = ?", data.Id)
		if err != nil {
			logx.Errorf("delete sys_user_project user_id: %s", err)
			return err
		}
		// delete sys_user_role user_id
		_, err = m.conn.ExecCtx(ctx, "delete from sys_user_role where `user_id` = ?", data.Id)
		if err != nil {
			logx.Errorf("delete sys_user_role user_id: %s", err)
			return err
		}
		// delete sys_user_ugroup user_id
		_, err = m.conn.ExecCtx(ctx, "delete from sys_user_ugroup where `user_id` = ?", data.Id)
		if err != nil {
			logx.Errorf("delete sys_user_ugroup user_id: %s", err)
			return err
		}
		return nil
	})
	return err
}

func (m *defaultSysUserModel) UpdateStatus(ctx context.Context, data *SysUser) error {

	query := fmt.Sprintf("update %s set `last_update_by`=?,`status`=? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.LastUpdateBy, data.Status, data.Id)
	return err
}

func (m *defaultSysUserModel) UpdatePassword(ctx context.Context, data *SysUser) error {
	query := fmt.Sprintf("update %s set `password` = ? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.Password, data.Id)
	return err
}

func (m *defaultSysUserModel) UpdatePersonalInfo(ctx context.Context, data *SysUser) error {
	query := fmt.Sprintf("update %s set `nick_name` = ?,avatar=?,email=?,mobile=? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.NickName, data.Avatar, data.Email, data.Mobile, data.Id)
	return err
}

func (m *defaultSysUserModel) TransactInsert(ctx context.Context, data *adminclient.UserAddReq) error {
	salts := tool.Krand(6, tool.KC_RAND_KIND_ALL)
	insertsql := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,  ?, ?)", m.table, sysUserRowsExpectAutoSet)
	if err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		stmt, err := session.Prepare(insertsql)
		if err != nil {
			return err
		}
		defer stmt.Close()
		// 返回任何错误都会回滚事务
		obj, err := stmt.ExecCtx(ctx, data.Name, data.NickName, data.Avatar, tool.Md5ByString("123456"+salts+data.Name), salts, data.Email, data.Mobile, 1, data.DeptId, data.CreateBy, data.CreateBy, 0)
		if err != nil {
			return xerr.NewErrMsg("新增用户失败" + err.Error())
		}
		lastid, err := obj.LastInsertId()
		if err != nil || lastid < 0 {
			return xerr.NewErrMsg("新增用户失败")
		}

		//删除用户id相关角色id和用户组id对应信息
		_, err = m.conn.ExecCtx(ctx, "delete from sys_user_role where user_id = ?", lastid)
		if err != nil {
			return xerr.NewErrMsg("删除用户及角色关联失败")
		}
		_, err = m.conn.ExecCtx(ctx, "delete from sys_user_ugroup where user_id = ?", lastid)
		if err != nil {
			return xerr.NewErrMsg("删除用户及用户组关联失败")
		}
		_, err = m.conn.ExecCtx(ctx, "delete from sys_user_project where user_id = ?", lastid)
		if err != nil {
			return xerr.NewErrMsg("删除用户及项目关联失败")
		}
		//批量用户相关的用户组和角色id
		if data.UgroupIds != "" {

			ugroupObj := strings.Split(data.UgroupIds, ",")
			for _, v := range ugroupObj {
				if gconv.Int64(v) == 0 {
					return xerr.NewErrMsg("用户组id不存在")
				}
				_, err := m.conn.ExecCtx(ctx, "insert into sys_user_ugroup (user_id,ugroup_id) values (?, ?)", lastid, gconv.Int64(v))
				if err != nil {
					return xerr.NewErrMsg("插入用户及用户组关联失败")
				}
			}
		}
		if data.RoleIds != "" {
			roleObj := strings.Split(data.RoleIds, ",")
			for _, v := range roleObj {
				if gconv.Int64(v) == 0 {
					return xerr.NewErrMsg("角色id不存在")
				}
				_, err := m.conn.ExecCtx(ctx, "insert into sys_user_role (user_id,role_id) values (?, ?)", lastid, gconv.Int64(v))
				if err != nil {
					return xerr.NewErrMsg("插入用户及角色关联失败")
				}
			}
		}

		if data.ProjectIds != "" {

			projectObj := strings.Split(data.ProjectIds, ",")
			for _, v := range projectObj {
				if gconv.Int64(v) == 0 {
					return xerr.NewErrMsg("项目id不存在")
				}
				_, err := m.conn.ExecCtx(ctx, "insert into sys_user_project (user_id,project_id) values (?, ?)", lastid, gconv.Int64(v))
				if err != nil {
					return xerr.NewErrMsg("插入用户及项目关联失败")
				}
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
