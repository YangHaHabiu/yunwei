package mergePlan

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/yunwei/api/internal/logic/mergePlan"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"ywadmin-v3/common/result"
)

func MergePlanRangeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMergeListTreeByPlatIdAndProIdReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := mergePlan.NewMergePlanRangeListLogic(r.Context(), svcCtx)
		resp, err := l.MergePlanRangeList(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
		result.HttpResult(r, w, resp, err)

	}
}
