package tasksScheduleQueue

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/tasksScheduleQueue"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func TasksScheduleQueueUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTasksScheduleQueueReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := tasksScheduleQueue.NewTasksScheduleQueueUpdateLogic(r.Context(), svcCtx)
		err := l.TasksScheduleQueueUpdate(&req)
		result.HttpResult(r, w, nil, err)
	}
}
