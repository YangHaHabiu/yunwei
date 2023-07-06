package report

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/qqGroup/api/internal/logic/report"
	"ywadmin-v3/service/qqGroup/api/internal/svc"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


		l := report.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
