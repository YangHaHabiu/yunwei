package inws

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/inws/api/internal/logic/inws"
	"ywadmin-v3/service/inws/api/internal/svc"
	"ywadmin-v3/service/inws/api/internal/types"
)

func GetInsideTasksLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInsideTasksLogReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := inws.NewGetInsideTasksLogLogic(r.Context(), svcCtx, w, r)
		l.GetInsideTasksLog(&req)
		//result.HttpResult(r, w, nil, err)
	}
}
