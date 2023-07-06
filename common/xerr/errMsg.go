package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	//全局的错误
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[DB_DATA_ADD_ERROR] = "数据新增失败"
	message[DB_DATA_DELETE_ERROR] = "数据删除失败"
	message[GET_REDIS_USER_ERROR] = "获取redis用户的urls地址失败"
	message[ACCESS_ADDRESS_ERROR] = "没有该地址的访问权限"
	message[TOKEN_NOTEXIT_ERROR] = "token不存在"
	message[TOKEN_ERROR] = "token验证失败"

	//用户模块错误类型
	message[ADMIN_USERNAME_ERROR] = "用户名不存在或者已停用"
	message[ADMIN_USERNAMEPWD_ERROR] = "用户名的密码错误"
	message[ADMIN_NOTFOUNDUID_ERROR] = "找不到该用户的UID"
	message[ADMIN_UPDATEPERSON_ERROR] = "更新用户信息失败"
	message[ADMIN_UPDATEPERSON_PASSWORD_ERROR] = "更新用户密码失败"

	//菜单模块错误类型
	message[ADMIN_MENUSELECT_ERROR] = "获取菜单列表信息失败"

	//角色模块错误类型
	message[ADMIN_ROLESELECT_ERROR] = "获取角色列表信息失败"

	//部门模块错误类型
	message[ADMIN_DEPTSELECT_ERROR] = "获取部门列表信息失败"

	//角色模块错误类型
	message[ADMIN_ROLESELECT_ERROR] = "获取角色列表信息失败"

	//配置模块错误类型
	message[ADMIN_CONFIGSELECT_ERROR] = "获取配置列表信息失败"

	//职位模块错误类型
	message[ADMIN_JOBSELECT_ERROR] = "获取职位列表信息失败"

	//字典模块错误类型
	message[ADMIN_DICTSELECT_ERROR] = "获取字典列表信息失败"

	//系统日志模块错误类型
	message[ADMIN_SYSLOGSELECT_ERROR] = "获取系统日志列表信息失败"

	//登录日志模块错误类型
	message[ADMIN_LOGINLOGSELECT_ERROR] = "获取登录日志列表信息失败"

	//标签模块错误类型
	message[ADMIN_LABELSELECT_ERROR] = "获取标签列表信息失败"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
