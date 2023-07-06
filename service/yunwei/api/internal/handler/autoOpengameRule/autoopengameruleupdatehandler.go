package autoOpengameRule

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/autoOpengameRule"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func AutoOpengameRuleUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAutoOpengameRuleReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := autoOpengameRule.NewAutoOpengameRuleUpdateLogic(r.Context(), svcCtx)
		err := l.AutoOpengameRuleUpdate(&req)
		result.HttpResult(r, w, nil, err)
	}
}
