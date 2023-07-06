package resource

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/admin/api/internal/types"

	"ywadmin-v3/service/admin/api/internal/logic/resource"
	"ywadmin-v3/service/admin/api/internal/svc"
)

func ResourceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListResourceReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := resource.NewResourceListLogic(r.Context(), svcCtx)
		resp, err := l.ResourceList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
