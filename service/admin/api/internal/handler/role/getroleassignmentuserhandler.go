package role

import (
	"net/http"
	"ywadmin-v3/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/admin/api/internal/logic/role"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
)

func GetRoleAssignmentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRoleAssignmentUserReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := role.NewGetRoleAssignmentUserLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleAssignmentUser(&req)
		result.HttpResult(r, w, resp, err)
	}
}
