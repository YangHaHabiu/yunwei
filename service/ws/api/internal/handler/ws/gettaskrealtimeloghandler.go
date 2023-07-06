package ws

import (
	"net/http"
	"ywadmin-v3/service/ws/api/internal/logic/ws"

	"ywadmin-v3/service/ws/api/internal/svc"
)

func GetTaskRealTimeLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(w, r, svcCtx)
	}
}
