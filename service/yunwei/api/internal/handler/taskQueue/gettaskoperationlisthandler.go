package taskQueue

import (
	"net/http"
	"ywadmin-v3/service/yunwei/api/internal/logic/taskQueue"

	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/svc"
)

func GetTaskOperationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := taskQueue.NewGetTaskOperationListLogic(r.Context(), svcCtx)
		resp, err := l.GetTaskOperationList()
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
		result.HttpResult(r, w, resp, err)

	}
}
