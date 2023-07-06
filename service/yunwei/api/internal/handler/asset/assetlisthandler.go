package asset

import (
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/asset"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func AssetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListAssetReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := asset.NewAssetListLogic(r.Context(), svcCtx)
		resp, err := l.AssetList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
