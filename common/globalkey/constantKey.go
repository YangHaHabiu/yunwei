package globalkey

//软删除
var (
	DelStateNo  int64 = 0 //未删除
	DelStateYes int64 = 1 //已删除
	//最高管理员用户名称
	SuperUserName = "admin"
	//不需要鉴权用户
	NoAuthGroup = map[string]bool{
		SuperUserName: true,
	}
)
