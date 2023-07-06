package openPlan

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/openPlan"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func OpenPlanBatchModifyOpenTimeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchModifyOpenTimeReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := openPlan.NewOpenPlanBatchModifyOpenTimeLogic(r.Context(), svcCtx)
		err := l.OpenPlanBatchModifyOpenTime(&req)
		result.HttpResult(r, w, nil, err)
	}
}
