package insideInstallPlan

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/intranet/api/internal/logic/insideInstallPlan"
	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"
)

func InsideInstallPlanAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddInsideInstallPlanReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := insideInstallPlan.NewInsideInstallPlanAddLogic(r.Context(), svcCtx)
		err := l.InsideInstallPlanAdd(&req)
		result.HttpResult(r, w, nil, err)
	}
}
