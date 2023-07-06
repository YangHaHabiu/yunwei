package maintainPlan

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/yunwei/api/internal/logic/maintainPlan"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"ywadmin-v3/common/result"
)

func MaintainPlanUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMaintainPlanReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := maintainPlan.NewMaintainPlanUpdateLogic(r.Context(), svcCtx)
		err := l.MaintainPlanUpdate(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.Ok(w)
		//}
		result.HttpResult(r, w, nil, err)

	}
}
