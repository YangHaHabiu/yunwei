package dashboard

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/dashboard"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func GetGameTrendChartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.GetGameTrendChartListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := dashboard.NewGetGameTrendChartLogic(r.Context(), svcCtx)
		resp, err := l.GetGameTrendChart(&req)
		result.HttpResult(r, w, resp, err)
	}
}
