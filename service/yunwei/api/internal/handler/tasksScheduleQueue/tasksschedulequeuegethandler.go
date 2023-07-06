package tasksScheduleQueue

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/tasksScheduleQueue"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func TasksScheduleQueueGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTasksScheduleQueueReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := tasksScheduleQueue.NewTasksScheduleQueueGetLogic(r.Context(), svcCtx)
		resp, err := l.TasksScheduleQueueGet(&req)
		result.HttpResult(r, w, resp, err)
	}
}
