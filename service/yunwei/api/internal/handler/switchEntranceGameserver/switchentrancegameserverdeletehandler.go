package switchEntranceGameserver

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/switchEntranceGameserver"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func SwitchEntranceGameserverDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSwitchEntranceGameserverReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := switchEntranceGameserver.NewSwitchEntranceGameserverDeleteLogic(r.Context(), svcCtx)
		err := l.SwitchEntranceGameserverDelete(&req)
		result.HttpResult(r, w, nil, err)
	}
}
