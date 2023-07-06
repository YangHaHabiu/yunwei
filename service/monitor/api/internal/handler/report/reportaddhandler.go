package report

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/monitor/api/internal/logic/report"
	"ywadmin-v3/service/monitor/api/internal/svc"
	"ywadmin-v3/service/monitor/api/internal/types"
)

func ReportAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReportAddReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := report.NewReportAddLogic(r.Context(), svcCtx)
		err := l.ReportAdd(&req)
		result.HttpResult(r, w, nil, err)
	}
}
