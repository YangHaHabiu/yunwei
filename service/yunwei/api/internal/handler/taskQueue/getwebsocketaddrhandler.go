package taskQueue

import (
	"net/http"
	"ywadmin-v3/service/yunwei/api/internal/logic/taskQueue"

	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/svc"
)

func GetWebSocketAddrHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := taskQueue.NewGetWebSocketAddrLogic(r.Context(), svcCtx)
		resp, err := l.GetWebSocketAddr()
		result.HttpResult(r, w, resp, err)

	}
}
