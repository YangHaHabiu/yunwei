package cron

import (
	"encoding/json"
	"fmt"
	"time"
	"ywadmin-v3/common/crawler"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/service/qqGroupv3/api/internal/config"
	"ywadmin-v3/service/qqGroupv3/api/internal/logic/report"
	"ywadmin-v3/service/qqGroupv3/api/internal/svc"

	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	url       string
	setCookie string
	urlx      = "http://244.mengzuofang.com:8081"
)

type cronTask struct {
	//logx.Logger
	svcCtx *svc.ServiceContext
	config.Config
}

func NewCronTask(svcCtx *svc.ServiceContext) *cronTask {

	if svcCtx.Config.IsOpenCheckVpn {
		s, err := report.AikuaiLoginFunc(urlx)
		if err != nil {
			fmt.Println(err)
		}
		setCookie = s
	}

	return &cronTask{
		svcCtx: svcCtx,
	}
}

func (l *cronTask) Start() {
	c := cron.New(cron.WithSeconds())
	// 每天23点执行一次
	c.AddFunc("0 0 23 * * ?", l.cronTask)
	//每天9点30执行一次
	//c.AddFunc("0 30 9 * * ?", l.sendTodayNews)
	//每天11点00执行一次
	//c.AddFunc("0 0 11 * * ?", l.sendTodayYiqing)
	//每10分钟执行一次
	c.AddFunc("0 */10 * * * ?", l.testOnceEveryHour)
	//只有黑石才执行每分钟线路检测并切换
	if l.svcCtx.Config.IsOpenCheckVpn {
		c.AddFunc("1 */1 * * * ?", l.erveryMinuteCheckLine)
	}
	c.Start()
	defer c.Stop()
	select {}
}
func (l *cronTask) Stop() {
}

// 凌晨23点开始执行清理一个月前的数据
func (l *cronTask) cronTask() {
	fmt.Println("cron start")
	err := l.svcCtx.QqMessageHistoryModel.DeleteMessageHistory(31)
	if err != nil {
		fmt.Println(err)
		logx.Error(err)
	}
	fmt.Println("cron end")
}

type qqMsgStatus struct {
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
	Data    struct {
		Online bool `json:"online"`
		Status int  `json:"status"`
	} `json:"data"`
}

func (l *cronTask) testOnceEveryHour() {

	content, errx := crawler.GetHttpResponse(fmt.Sprintf("%s/getStatus", globalkey.QqMsgKey["group_qqapi"].(string)), false)
	var qqTmpStatus qqMsgStatus
	if errx == nil {
		err := json.Unmarshal(content, &qqTmpStatus)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if qqTmpStatus.Retcode != 0 || !qqTmpStatus.Data.Online {
		msg := fmt.Sprintf("qq异常不在线，在线状态：%v 状态值：%d", qqTmpStatus.Data.Online, qqTmpStatus.Data.Status)
		allList, _ := l.svcCtx.QqLoadBalanceModel.FindAll("group_type__=", "group")
		if len(*allList) == 1 {
			fmt.Println(msg, "，因只有一个qq，无法自动切换")
			return
		}
		m := 0
		for i := 0; i < len(globalkey.QqGroupList); i++ {
			if globalkey.QqGroupList[i] == globalkey.QqMsgKey["group_qq"] {
				m = i
				break
			}
		}
		m = (m + 1) % len(globalkey.QqGroupList)
		err := l.svcCtx.QqLoadBalanceModel.UpdateIsMaster("group", globalkey.QqMsgKey["group_qq"].(int64), 0)
		err = l.svcCtx.QqLoadBalanceModel.UpdateIsMaster("group", globalkey.QqGroupList[m], 1)
		if err != nil {
			fmt.Printf("异常错误%v\n", err)
			return
		}
		filters, _ := l.svcCtx.QqLoadBalanceModel.FindMasterByFilters("qq__=", globalkey.QqGroupList[m], "group_type__=", "group")
		globalkey.QqMsgKey["group_qq"] = (*filters)[0].Qq
		globalkey.QqMsgKey["group_qqapi"] = (*filters)[0].QqApi
		url = fmt.Sprintf("%s/sendGroupMsg", globalkey.QqMsgKey["group_qqapi"].(string))
		crawler.Post(url, map[string]string{
			"group_id":    l.svcCtx.Config.YwQQGroup,
			"message":     fmt.Sprintf("因上个QQ号异常，现自动切换成新QQ号：%d进行工作", globalkey.QqGroupList[m]),
			"auto_escape": "false",
		})
	}

}

func (l *cronTask) erveryMinuteCheckLine() {
	//每天登陆下后台保持cookie最新
	t := time.Now()
	if t.Hour() == 23 && ( t.Minute() == 59 || t.Minute()==58 ){
		s, err := report.AikuaiLoginFunc(urlx)
		if err != nil {
			fmt.Println("登陆路由后台失败：", err)
			return
		}
		setCookie = s
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")+"开始登录路由后台了--->", setCookie)
	}
	//利用日志或者拨号记录判断
	if report.AikuaiJudgeLine(urlx, setCookie) || report.AikuaiDialJudgeLine(urlx, setCookie) {
		//开始切换
		coookie, _ := report.AikuaiLoginFunc(urlx)
		var msg string
		s, err := report.AikuaiSwitchLineFunc(urlx, coookie)
		if err != nil {
			msg = "切换线路失败,原因：" + err.Error()
		} else {
			msg = s
		}
		url = fmt.Sprintf("%s/sendGroupMsg", globalkey.QqMsgKey["group_qqapi"].(string))
		err, _ = crawler.Post(url, map[string]string{
			"group_id":    l.svcCtx.Config.YwQQGroup,
			"message":     msg,
			"auto_escape": "false",
		})
		if err != nil {
			fmt.Println("发送qq消息失败，原因：" + err.Error())
		}
	}

}
