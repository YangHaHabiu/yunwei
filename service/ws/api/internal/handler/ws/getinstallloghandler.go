package ws

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/ws/api/internal/logic/ws"
	"ywadmin-v3/service/ws/api/internal/svc"
	"ywadmin-v3/service/ws/api/internal/types"
)

func GetInstallLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInstallLogReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := ws.NewGetInstallLogLogic(r.Context(), svcCtx)
		l.GetInstallLog(&req, w, r)

	}
}
