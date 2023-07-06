package switchEntranceGameserver

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/switchEntranceGameserver"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func SwitchEntranceGameserverListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListSwitchEntranceGameserverReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := switchEntranceGameserver.NewSwitchEntranceGameserverListLogic(r.Context(), svcCtx)
		resp, err := l.SwitchEntranceGameserverList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
