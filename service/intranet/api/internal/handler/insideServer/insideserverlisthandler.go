package insideServer

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/intranet/api/internal/logic/insideServer"
	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"
)

func InsideServerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListInsideServerReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := insideServer.NewInsideServerListLogic(r.Context(), svcCtx)
		resp, err := l.InsideServerList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
