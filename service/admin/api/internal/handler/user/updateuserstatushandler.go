package user

import (
	"net/http"
	"ywadmin-v3/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/admin/api/internal/logic/user"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
)

func UpdateUserStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewUpdateUserStatusLogic(r.Context(), svcCtx)
		err := l.UpdateUserStatus(&req)
		result.HttpResult(r, w, nil, err)
	}
}
