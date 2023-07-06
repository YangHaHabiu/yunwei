/*
@Time : 2022-3-24 18:31
@Author : acool
@File : setUidToCtxMiddleware.go
*/
package middleware

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/result"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/identity/rpc/identity"
	"ywadmin-v3/service/sendMsg/api/internal/svc"
)

type SetUidToCtxMiddleware struct {
	svcCtx *svc.ServiceContext
}

func NewSetUidToCtxMiddleware(svc *svc.ServiceContext) *SetUidToCtxMiddleware {
	return &SetUidToCtxMiddleware{
		svcCtx: svc,
	}
}

func (m *SetUidToCtxMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		urlx := strings.Split(r.RequestURI, "?")[0]
		// if urlx == "/sendMsg/msgApi/add" {
		// 	next(w, r)
		// } else {
		if tool.StrInArr(urlx, ctxdata.GlobalMiddleWtiteList) {
			next(w, r)

		} else {
			// 获取头部token
			token := r.Header.Get("authorization")
			if token == "" {
				result.HttpResult(r, w, nil, xerr.NewErrCode(xerr.TOKEN_NOTEXIT_ERROR))
				return
			}

			// 解析token获取对应的值
			parseToken, err := m.svcCtx.IdentityRpc.ParseToken(r.Context(), &identity.ParseTokenReq{
				Token:  token,
				Secret: m.svcCtx.Config.Auth.AccessSecret,
			})
			if err != nil {
				result.HttpResult(r, w, nil, xerr.NewErrCode(xerr.TOKEN_NOTEXIT_ERROR))
				return
			}

			// 检测token是否有效
			validateToken, err := m.svcCtx.IdentityRpc.ValidateToken(r.Context(), &identity.ValidateTokenReq{
				UserId: parseToken.UserId,
				Token:  token,
			})
			if err != nil || !validateToken.Ok {
				result.HttpResult(r, w, nil, xerr.NewErrCode(xerr.TOKEN_NOTEXIT_ERROR))
				return
			}

			// 需要鉴权url
			if !urlNoAuth(urlx, m.svcCtx, r) {
				compile := regexp.MustCompile(`[/|?]`)
				realRequestPathList := compile.Split(urlx, -1)
				if _, ok := globalkey.NoAuthGroup[parseToken.UserName]; !ok {
					urls, err := m.svcCtx.RedisClient.Get(fmt.Sprintf(globalkey.CacheUserAuthKey, parseToken.UserId))
					if err != nil {
						result.HttpResult(r, w, nil, xerr.NewErrCode(xerr.GET_REDIS_USER_ERROR))

						return
					}
					b := false
					for _, url := range strings.Split(urls, ",") {
						urlList := compile.Split(url, -1)
						if len(urlList) >= 3 {
							if url == strings.Join(realRequestPathList[:4], "/") {
								b = true
								break
							}
						}

					}
					if !b {
						result.HttpResult(r, w, nil, xerr.NewErrCode(xerr.ACCESS_ADDRESS_ERROR))
						return
					}
				}
			}
			getHeaderProjects := r.Header.Get("GlobalProjectIds")
			ctx := r.Context()
			ctx = context.WithValue(ctx, ctxdata.CtxKeyJwtUserId, parseToken.UserId)
			ctx = context.WithValue(ctx, ctxdata.CtxKeyJwtUserName, parseToken.UserName)
			ctx = context.WithValue(ctx, ctxdata.CtxKeyJwtNickName, parseToken.NickName)
			ctx = context.WithValue(ctx, ctxdata.CtxKeyJwtProjectIds, getHeaderProjects)
			next(w, r.WithContext(ctx))
		}

	}
}

// 当前url是否需要授权验证
func urlNoAuth(path string, svc *svc.ServiceContext, r *http.Request) bool {
	split := strings.Split(path, "/")
	if len(split) < 3 {
		return false
	}
	path = "/" + split[1] + "/" + split[2] + "/" + split[3]
	list, err := svc.AdminRpc.StrategyList(r.Context(), &adminclient.StrategyListReq{
		StLevel:  3,
		StIsAuth: 2,
	})
	if err != nil {
		return false
	}
	for _, val := range list.List {
		if val.StUrls == path {
			return true
		}
	}
	return false
}
