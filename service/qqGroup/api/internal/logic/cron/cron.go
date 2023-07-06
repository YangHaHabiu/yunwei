package cron

import (
	"encoding/json"
	"fmt"
	"time"
	"ywadmin-v3/common/crawler"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/qqmanger/discuss"
	"ywadmin-v3/service/qqGroup/api/internal/svc"

	"github.com/robfig/cron/v3"
)

var url string

type cronTask struct {
	//logx.Logger
	svcCtx *svc.ServiceContext
}

func NewCronTask(svcCtx *svc.ServiceContext) *cronTask {
	url = fmt.Sprintf("%s/sendDiscussMsg", globalkey.QqMsgKey["discuss_qqapi"].(string))
	return &cronTask{
		svcCtx: svcCtx,
	}
}

func (l *cronTask) Start() {
	c := cron.New(cron.WithSeconds())
	// 每天23点执行一次
	c.AddFunc("0 0 23 * * ?", l.cronTask)
	//每天9点30执行一次
	c.AddFunc("0 30 9 * * ?", l.sendTodayNews)
	//每天11点00执行一次
	c.AddFunc("0 0 11 * * ?", l.sendTodayYiqing)
	//每10分钟执行一次
	c.AddFunc("0 */10 * * * ?", l.testOnceEveryHour)
	c.Start()
	defer c.Stop()
	select {}
}
func (l *cronTask) Stop() {
}

//凌晨23点开始执行清理一个月前的数据
func (l *cronTask) cronTask() {
	fmt.Println("cron start")
	err := l.svcCtx.QqMessageHistoryModel.DeleteMessageHistory(31)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cron end")
}

//每天9点30发送一次新闻，及假期信息
func (l *cronTask) sendTodayNews() {
	crawler.Post(url, map[string]string{
		"discuss_id":  "1544553049",
		"message":     discuss.TodayNews() + "------------------------------\n",
		"auto_escape": "false",
	})
	time.Sleep(10 * time.Second)
	l.sendMoyu()
}

func (l *cronTask) sendMoyu() {
	crawler.Post(url, map[string]string{
		"discuss_id":  "1544553049",
		"message":     discuss.FishingReminder(),
		"auto_escape": "false",
	})
}

func (l *cronTask) sendTodayYiqing() {
	crawler.Post(url, map[string]string{
		"discuss_id":  "1544553049",
		"message":     discuss.TodayYiqing(),
		"auto_escape": "false",
	})
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
	if !qqTmpStatus.Data.Online || qqTmpStatus.Retcode != 0 {
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
			"group_id":    "324113338",
			"message":     fmt.Sprintf("因上个QQ号异常，现自动切换成新QQ号：%d进行工作", globalkey.QqGroupList[m]),
			"auto_escape": "false",
		})
	} else {
		fmt.Printf("%s--> QQ：%d暂无异常\n", time.Now().Format("2006-01-02 15:04"), globalkey.QqMsgKey["group_qq"].(int64))
	}

}
