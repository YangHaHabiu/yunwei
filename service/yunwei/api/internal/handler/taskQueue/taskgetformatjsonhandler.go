package taskQueue

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/taskQueue"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func TaskGetFormatJsonHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskGetFormatJsonReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := taskQueue.NewTaskGetFormatJsonLogic(r.Context(), svcCtx)
		resp, err := l.TaskGetFormatJson(&req)
		result.HttpResult(r, w, resp, err)
	}
}
