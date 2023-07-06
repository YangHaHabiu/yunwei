package taskLog

import (
	"net/http"
	"ywadmin-v3/service/yunwei/api/internal/logic/taskLog"

	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/svc"
)

func HotLogHistoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := taskLog.NewHotLogHistoryListLogic(r.Context(), svcCtx)
		resp, err := l.HotLogHistoryList()
		result.HttpResult(r, w, resp, err)

	}
}
