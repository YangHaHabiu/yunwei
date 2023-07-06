package asset

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ywadmin-v3/service/yunwei/api/internal/logic/asset"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"ywadmin-v3/common/result"
)

func AssetFileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetFileReq
		if err := httpx.Parse(r, &req); err != nil {

			result.ParamErrorResult(r, w, err)
			return
		}

		l := asset.NewAssetFileUploadLogic(r.Context(), svcCtx, r)
		err := l.AssetFileUpload(&req)
		result.HttpResult(r, w, nil, err)
	}
}
