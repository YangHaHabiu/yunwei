package configFileDelivery

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/configFileDelivery"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func ConfigFileDeliveryUpdateTemplateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateConfigFileDeliveryTemplateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := configFileDelivery.NewConfigFileDeliveryUpdateTemplateLogic(r.Context(), svcCtx)
		err := l.ConfigFileDeliveryUpdateTemplate(&req)
		result.HttpResult(r, w, nil, err)
	}
}
