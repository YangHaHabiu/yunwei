package configFileDelivery

import (
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/configFileDelivery"
	"ywadmin-v3/service/yunwei/api/internal/svc"
)

func ConfigFileDeliveryGetLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := configFileDelivery.NewConfigFileDeliveryGetLogLogic(r.Context(), svcCtx)
		err := l.ConfigFileDeliveryGetLog()
		result.HttpResult(r, w, nil, err)
	}
}
