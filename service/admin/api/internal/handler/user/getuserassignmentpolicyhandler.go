package user

import (
	"net/http"
	"ywadmin-v3/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/admin/api/internal/logic/user"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
)

func GetUserAssignmentPolicyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserAssignmentPolicyReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewGetUserAssignmentPolicyLogic(r.Context(), svcCtx)
		resp, err := l.GetUserAssignmentPolicy(&req)
		result.HttpResult(r, w, resp, err)
	}
}
