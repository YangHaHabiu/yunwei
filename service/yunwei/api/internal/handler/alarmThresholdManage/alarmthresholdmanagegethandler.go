package alarmThresholdManage

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ywadmin-v3/common/result"
	"ywadmin-v3/service/yunwei/api/internal/logic/alarmThresholdManage"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
)

func AlarmThresholdManageGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAlarmThresholdManageReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := alarmThresholdManage.NewAlarmThresholdManageGetLogic(r.Context(), svcCtx)
		resp, err := l.AlarmThresholdManageGet(&req)
		result.HttpResult(r, w, resp, err)
	}
}
