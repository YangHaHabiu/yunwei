package asset

import (
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/asset"

	"ywadmin-v3/service/yunwei/api/internal/svc"
)

func OwnerProjectDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := asset.NewOwnerProjectDataLogic(r.Context(), svcCtx)
		resp, err := l.OwnerProjectData()
		result.HttpResult(r, w, resp, err)
	}
}
