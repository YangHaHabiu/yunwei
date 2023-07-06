package insideOperation

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/intranet/api/internal/logic/insideOperation"
	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"
)

func InsideOperationDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteInsideOperationReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := insideOperation.NewInsideOperationDeleteLogic(r.Context(), svcCtx)
		err := l.InsideOperationDelete(&req)
		result.HttpResult(r, w, nil, err)
	}
}
