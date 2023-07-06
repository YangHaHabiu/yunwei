package company

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/admin/api/internal/logic/company"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
)

func UpdateSupplyCompanyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSupplyCompanyReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := company.NewUpdateSupplyCompanyLogic(r.Context(), svcCtx)
		err := l.UpdateSupplyCompany(&req)
		result.HttpResult(r, w, nil, err)
	}
}
