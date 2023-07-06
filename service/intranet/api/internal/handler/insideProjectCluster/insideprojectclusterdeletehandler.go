package insideProjectCluster

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/intranet/api/internal/logic/insideProjectCluster"
	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"
)

func InsideProjectClusterDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteInsideProjectClusterReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := insideProjectCluster.NewInsideProjectClusterDeleteLogic(r.Context(), svcCtx)
		err := l.InsideProjectClusterDelete(&req)
		result.HttpResult(r, w, nil, err)
	}
}
