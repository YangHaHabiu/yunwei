package ws

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/service/ws/api/internal/logic/ws"
	"ywadmin-v3/service/ws/api/internal/svc"
	"ywadmin-v3/service/ws/api/internal/types"
)

func GetWebSshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetWebSshReq
		if err := httpx.Parse(r, &req); err != nil {
			return
		}

		ws.WebSshHandler(w, r, svcCtx, &req)
	}
}
