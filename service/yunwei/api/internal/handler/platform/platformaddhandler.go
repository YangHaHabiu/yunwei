package platform

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/platform"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func PlatformAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddPlatformReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := platform.NewPlatformAddLogic(r.Context(), svcCtx)
		err := l.PlatformAdd(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.Ok(w)
		//}
		result.HttpResult(r, w, nil, err)

	}
}
