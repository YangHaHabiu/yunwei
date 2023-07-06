package dept

import (
	"net/http"
	"ywadmin-v3/common/result"

	"ywadmin-v3/service/admin/api/internal/logic/dept"
	"ywadmin-v3/service/admin/api/internal/svc"
)

func TreeselectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dept.NewTreeselectLogic(r.Context(), svcCtx)
		resp, err := l.Treeselect()
		result.HttpResult(r, w, resp, err)
	}
}
