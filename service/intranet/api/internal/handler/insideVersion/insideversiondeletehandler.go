package insideVersion

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/intranet/api/internal/logic/insideVersion"
	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"
)

func InsideVersionDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteInsideVersionReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := insideVersion.NewInsideVersionDeleteLogic(r.Context(), svcCtx)
		err := l.InsideVersionDelete(&req)
		result.HttpResult(r, w, nil, err)
	}
}
