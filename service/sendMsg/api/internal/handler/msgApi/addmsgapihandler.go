package msgApi

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/sendMsg/api/internal/logic/msgApi"
	"ywadmin-v3/service/sendMsg/api/internal/svc"
	"ywadmin-v3/service/sendMsg/api/internal/types"
)

func AddMsgApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddMsgApiReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := msgApi.NewAddMsgApiLogic(r.Context(), svcCtx, r)
		err := l.AddMsgApi(&req)
		result.HttpResult(r, w, nil, err)
	}
}
