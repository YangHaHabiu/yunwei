package insideTasks

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/intranet/api/internal/logic/insideTasks"
	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"
)

func InsideTasksDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteInsideTasksReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := insideTasks.NewInsideTasksDeleteLogic(r.Context(), svcCtx)
		err := l.InsideTasksDelete(&req)
		result.HttpResult(r, w, nil, err)
	}
}
