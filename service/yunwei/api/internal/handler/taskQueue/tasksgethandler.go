package taskQueue

import (
	"net/http"
	"ywadmin-v3/service/yunwei/api/internal/logic/taskQueue"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"ywadmin-v3/common/result"
)

func TasksGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTasksReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := taskQueue.NewTasksGetLogic(r.Context(), svcCtx)
		resp, err := l.TasksGet(&req)
		result.HttpResult(r, w, resp, err)

	}
}
