package graph

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/monitor/api/internal/logic/graph"
	"ywadmin-v3/service/monitor/api/internal/svc"
	"ywadmin-v3/service/monitor/api/internal/types"
)

func GraphListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListGraphReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := graph.NewGraphListLogic(r.Context(), svcCtx)
		resp, err := l.GraphList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
