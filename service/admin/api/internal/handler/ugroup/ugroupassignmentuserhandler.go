package ugroup

import (
	"net/http"
	"ywadmin-v3/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/admin/api/internal/logic/ugroup"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
)

func UgroupAssignmentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UgroupAssignmentUserReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := ugroup.NewUgroupAssignmentUserLogic(r.Context(), svcCtx)
		err := l.UgroupAssignmentUser(&req)
		result.HttpResult(r, w, nil, err)
	}
}
