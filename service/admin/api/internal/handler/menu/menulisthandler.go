package menu

import (
	"net/http"
	"ywadmin-v3/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/admin/api/internal/logic/menu"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
)

func MenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := menu.NewMenuListLogic(r.Context(), svcCtx)
		resp, err := l.MenuList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
