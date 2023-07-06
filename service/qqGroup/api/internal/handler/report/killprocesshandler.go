package report

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/qqGroup/api/internal/logic/report"
	"ywadmin-v3/service/qqGroup/api/internal/svc"
	"ywadmin-v3/service/qqGroup/api/internal/types"
)

func KillProcessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.KillProcessReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := report.NewKillProcessLogic(r.Context(), svcCtx)
		err := l.KillProcess(&req)
		result.HttpResult(r, w, nil, err)
	}
}
