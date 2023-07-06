package dashboard

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/dashboard"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func GetSumOfCurrentInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSumOfCurrentInfoListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := dashboard.NewGetSumOfCurrentInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetSumOfCurrentInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
