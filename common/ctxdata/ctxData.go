package ctxdata

import (
	"context"
)

// 从ctx获取uid
var (
	CtxKeyJwtUserId       = "jwtUserId"
	CtxKeyJwtUserName     = "jwtUserName"
	CtxKeyJwtNickName     = "jwtNickName"
	CtxKeyJwtProjectIds   = "jwtProjectIds"
	GlobalMiddleWtiteList = []string{
		"/admin/user/login",
		"/admin/captcha/getCaptcha",
		"/admin/captcha/getWordCaptcha",
		"/admin/captcha/checkCaptcha",
		"/monitor/report/add",
		"/sendMsg/msgApi/add",
	}
)

// 从ctx获取uid
func GetUidFromCtx(ctx context.Context) int64 {
	uid, _ := ctx.Value(CtxKeyJwtUserId).(int64)
	return uid
}

// 从ctx中获取uname
func GetUnameFromCtx(ctx context.Context) string {
	uid, _ := ctx.Value(CtxKeyJwtUserName).(string)
	return uid
}

// 从ctx中获取nickname
func GetNickNameFromCtx(ctx context.Context) string {
	urlencode, _ := ctx.Value(CtxKeyJwtNickName).(string)
	//unescape, _ := url.QueryUnescape(urlencode)
	return urlencode
}

// 从ctx中获取前端传递项目id组
func GetProjectIdsFromCtx(ctx context.Context) string {
	projects, _ := ctx.Value(CtxKeyJwtProjectIds).(string)
	return projects
}
