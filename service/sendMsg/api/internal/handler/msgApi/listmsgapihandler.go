package msgApi

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/sendMsg/api/internal/logic/msgApi"
	"ywadmin-v3/service/sendMsg/api/internal/svc"
	"ywadmin-v3/service/sendMsg/api/internal/types"
)

func ListMsgApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMsgApiReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := msgApi.NewListMsgApiLogic(r.Context(), svcCtx)
		resp, err := l.ListMsgApi(&req)
		result.HttpResult(r, w, resp, err)
	}
}
