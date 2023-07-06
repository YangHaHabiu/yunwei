package taskQueue

import (
	"net/http"
	"ywadmin-v3/service/yunwei/api/internal/logic/taskQueue"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"ywadmin-v3/common/result"
)

func TasksListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListTasksReq
		if err := httpx.Parse(r, &req); err != nil {

			result.ParamErrorResult(r, w, err)
			return
		}

		l := taskQueue.NewTasksListLogic(r.Context(), svcCtx)
		resp, err := l.TasksList(&req)

		result.HttpResult(r, w, resp, err)

	}
}
