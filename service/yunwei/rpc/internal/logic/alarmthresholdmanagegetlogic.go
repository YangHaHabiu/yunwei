package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmThresholdManageGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlarmThresholdManageGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmThresholdManageGetLogic {
	return &AlarmThresholdManageGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlarmThresholdManageGetLogic) AlarmThresholdManageGet(in *yunweiclient.GetAlarmThresholdManageReq) (*yunweiclient.ListAlarmThresholdManageData, error) {
	var (
		one *model.AlarmThresholdManageList
		err error
	)

	if in.AlarmThresholdManageTypes == "" {
		in.AlarmThresholdManageTypes = "1"
	}

	if in.AlarmThresholdManageId == 0 {
		one = new(model.AlarmThresholdManageList)
		if in.AlarmThresholdManageTypes == "1" {
			one.Config = `{
    "machine":{
        "disk_use_percentage_alarm":80,
        "disk_iowait_alarm":7.5,
        "disk_util_alarm":40,
        "mem_free_alarm":2048,
        "mem_use_percentage_alarm":80
    },
    "business":{
        "memory_alarm_threshold":3,
        "cpu_alarm_threshold":150
    }
}
`
		} else if in.AlarmThresholdManageTypes == "2" {
			one.Config = `{
"single":{
	"GAME_RESERVE_MEM":null,
	"GAME_AVG_MEM":null,
	"MIN_DISK_PERCENTAGE":null,
	"MAX_INSTALL":null,
	"MAX_INSTALL_6":null,
	"MAX_INSTALL_9":null,
	"MAX_INSTALL_12":null
},
"cross":{
	"GAME_RESERVE_MEM":null,
	"GAME_AVG_MEM":null,
	"MIN_DISK_PERCENTAGE":null,
	"MAX_INSTALL":null
}
}
`
		}
	} else {
		one, err = l.svcCtx.AlarmThresholdManageModel.FindOne(l.ctx, in.AlarmThresholdManageId)
		if err != nil {
			return nil, xerr.NewErrMsg("查询单条数据失败")
		}
	}

	var tmp yunweiclient.ListAlarmThresholdManageData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
