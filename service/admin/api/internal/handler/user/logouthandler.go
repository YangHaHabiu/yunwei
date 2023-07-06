package user

import (
	"net/http"
	"ywadmin-v3/common/result"

	"ywadmin-v3/service/admin/api/internal/logic/user"
	"ywadmin-v3/service/admin/api/internal/svc"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout(r)
		result.HttpResult(r, w, nil, err)
	}
}
