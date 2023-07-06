package taskQueue

import (
	"net/http"
	"ywadmin-v3/service/yunwei/api/internal/logic/taskQueue"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"ywadmin-v3/common/result"
)

func TasksStopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTasksReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := taskQueue.NewTasksStopLogic(r.Context(), svcCtx)
		err := l.TasksStop(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.Ok(w)
		//}
		result.HttpResult(r, w, nil, err)

	}
}
