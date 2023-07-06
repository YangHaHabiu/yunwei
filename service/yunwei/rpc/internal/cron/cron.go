package cron

import (
	"context"
	"fmt"
	"log"
	"time"
	"ywadmin-v3/service/monitor/rpc/monitorclient"
	"ywadmin-v3/service/yunwei/model"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"

	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

type cronTask struct {
	svcCtx *svc.ServiceContext
}

func NewCronTask(svcCtx *svc.ServiceContext) *cronTask {
	return &cronTask{
		svcCtx: svcCtx,
	}
}

func (l *cronTask) Start() {
	c := cron.New(cron.WithSeconds())
	// 每半个月23点执行一次
	c.AddFunc("0 0 23 1,15 * ?", l.cronTask)
	// 每5分钟计划任务
	c.AddFunc("0 */5 * * * ?", l.fiveCrond)
	c.Start()
	defer c.Stop()
	select {}
}
func (l *cronTask) Stop() {
}

func (l *cronTask) fiveCrond() {
	if l.svcCtx.Config.IsOpenCall {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		format := "2006-01-02 15:04"
		timeNow := time.Now().Format(format)
		t, err := time.ParseInLocation(format, timeNow, time.Local)
		if err != nil {
			logx.Error(fmt.Sprintf("格式化日期出错,%v", err))
			return
		}
		nowMinute := t.Unix()

		filters := make([]interface{}, 0)
		filters = append(filters,
			"view_recycle_type__=", "2",
			"view_clean_type__=", "2",
			"view_init_type__=", "1",
			"view_user_project_type__=", "1",
		)
		va, err := l.svcCtx.AssetModel.FindAll(ctx, filters...)
		if err != nil {
			logx.Error(fmt.Sprintf("查询资产出错,%v", err))
			return
		}
		if len(*va) > 0 {
			for _, v := range *va {
				_, err = l.svcCtx.MonitorRpc.SelectReport(ctx, &monitorclient.SelectReportReq{
					AssetId:    v.ViewAssetId.Int64,
					AssetIp:    v.ViewOuterIp.String,
					Remark:     fmt.Sprintf("%s_%s", v.ViewUserProjectEn.String, v.ViewEnHostRole.String),
					ReportTime: nowMinute,
				})
				if err != nil {
					logx.Error(fmt.Sprintf("查询信息失败了%v", err))
					return
				}
				time.Sleep(200 * time.Millisecond)
			}
		}

	}

}

func (l *cronTask) cronTask() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	timeNow := time.Now().Unix()
	for _, v := range []string{
		"server", "game",
	} {
		trend, err := l.svcCtx.StatServerGameInfoModel.FindPageListByPageTrend(ctx, v)
		if err != nil {
			logx.Error(fmt.Sprintf("计划任务失败，原因:%v", err))
			return
		}
		for _, v1 := range *trend {
			_, err := l.svcCtx.StatServerGameInfoModel.Insert(ctx, &model.StatServerGameInfo{
				ProjectId:  v1.ProjectId,
				ProjectEn:  v1.ProjectEn,
				Counts:     v1.Counts,
				Detail:     v1.Detail,
				CreateTime: timeNow,
				CountType:  v,
			})
			if err != nil {
				logx.Error(fmt.Sprintf("插入统计任务失败，原因:%v", err))
				return
			}
		}

	}
	log.Println("cron end")
}
