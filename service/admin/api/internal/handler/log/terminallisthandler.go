package log

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/admin/api/internal/logic/log"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"ywadmin-v3/common/result"
)

func TerminalListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListTerminalLogReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := log.NewTerminalListLogic(r.Context(), svcCtx)
		resp, err := l.TerminalList(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
		result.HttpResult(r, w, resp, err)

	}
}
