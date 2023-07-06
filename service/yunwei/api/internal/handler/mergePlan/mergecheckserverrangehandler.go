package mergePlan

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/mergePlan"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func MergeCheckServerRangeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MergeCheckServerRangeReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := mergePlan.NewMergeCheckServerRangeLogic(r.Context(), svcCtx)
		resp, err := l.MergeCheckServerRange(&req)
		result.HttpResult(r, w, resp, err)
	}
}
