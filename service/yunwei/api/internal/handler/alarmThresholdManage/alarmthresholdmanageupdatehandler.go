package alarmThresholdManage

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/alarmThresholdManage"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func AlarmThresholdManageUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAlarmThresholdManageReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := alarmThresholdManage.NewAlarmThresholdManageUpdateLogic(r.Context(), svcCtx)
		err := l.AlarmThresholdManageUpdate(&req)
		result.HttpResult(r, w, nil, err)
	}
}
