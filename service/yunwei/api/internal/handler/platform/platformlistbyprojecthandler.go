package platform

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/yunwei/api/internal/logic/platform"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"ywadmin-v3/common/result"
)

func PlatformListByProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlatformListByProjectReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := platform.NewPlatformListByProjectLogic(r.Context(), svcCtx)
		resp, err := l.PlatformListByProject(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
		result.HttpResult(r, w, resp, err)

	}
}
