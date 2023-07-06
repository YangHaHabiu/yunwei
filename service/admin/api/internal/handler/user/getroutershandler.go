package user

import (
	"net/http"
	"ywadmin-v3/common/result"

	"ywadmin-v3/service/admin/api/internal/logic/user"
	"ywadmin-v3/service/admin/api/internal/svc"
)

func GetRoutersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetRoutersLogic(r.Context(), svcCtx)
		resp, err := l.GetRouters()
		result.HttpResult(r, w, resp, err)
	}
}
