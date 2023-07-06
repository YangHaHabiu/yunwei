package captcha

import (
	"net/http"

	"ywadmin-v3/common/result"
	"ywadmin-v3/service/admin/api/internal/logic/captcha"
	"ywadmin-v3/service/admin/api/internal/svc"
)

func GetWordCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := captcha.NewGetWordCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetWordCaptcha()
		result.HttpResult(r, w, resp, err)
	}
}
