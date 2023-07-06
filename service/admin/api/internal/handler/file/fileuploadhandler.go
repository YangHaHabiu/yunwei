package file

import (
	"net/http"
	"ywadmin-v3/common/result"

	"ywadmin-v3/service/admin/api/internal/logic/file"
	"ywadmin-v3/service/admin/api/internal/svc"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewFileUploadLogic(r, r.Context(), svcCtx)
		resp, err := l.FileUpload()
		result.HttpResult(r, w, resp, err)
	}
}
