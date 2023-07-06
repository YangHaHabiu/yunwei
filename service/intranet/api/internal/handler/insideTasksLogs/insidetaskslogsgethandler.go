package insideTasksLogs

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/intranet/api/internal/logic/insideTasksLogs"
	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"
)

func InsideTasksLogsGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInsideTasksLogsReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := insideTasksLogs.NewInsideTasksLogsGetLogic(r.Context(), svcCtx)
		resp, err := l.InsideTasksLogsGet(&req)
		result.HttpResult(r, w, resp, err)
	}
}
