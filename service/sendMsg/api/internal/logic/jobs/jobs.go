package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/service/sendMsg/api/internal/logic/notify"
	"ywadmin-v3/service/sendMsg/api/internal/svc"
	"ywadmin-v3/service/sendMsg/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

// 消息队列
func RunSendMsgJobs(svcCtx *svc.ServiceContext) {
	var (
		wechatCount   int64
		emailCount    int64
		dingdingCount int64
		feishuCount   int64
	)
	total := svcCtx.Config.ApiAuthKey.Limit

	for {
		ctx, _ := context.WithCancel(context.Background())
		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		nums := r1.Intn(10)
		//fmt.Println("开始发消息随机等待", nums, "秒")
		time.Sleep(time.Duration(nums) * time.Second)
		all, err := svcCtx.SendMsgRecordModel.FindAll(ctx,
			"status__=", "1",
		)
		if err != nil {
			logx.Errorf("查询消息任务失败，原因：" + err.Error())
			return
		}
		if len(*all) > 0 {
			for _, v := range *all {
				svcCtx.SendMsgRecordModel.UpdateStatusById(ctx, v.Id, 2, "")
				var (
					masterAccount *model.SendAccount
					findAllX      *[]model.SendAccount
					status        int64
					failureLog    string
				)
				status = 3
				findAllX, err = svcCtx.SendAccountModel.FindAll(ctx,
					"send_channel__=", v.MsgType,
					"app_key__=", v.AppKey,
				)
				if err != nil {
					logx.Errorf("查询帐号信息失败，原因：" + err.Error())
					return
				}
				if len(*findAllX) != 0 {
					//查询master
					for _, list := range *findAllX {
						if list.IsMaster == 1 {
							masterAccount = &list
							break
						}
					}
					if masterAccount == nil {
						masterAccount = &(*findAllX)[0]
						logx.Error("修改默认的master-->", masterAccount)
						svcCtx.SendAccountModel.UpdateMasterById(ctx,
							masterAccount.SendChannel,
							masterAccount.AppKey,
							masterAccount.Id)
					}

					logx.Error("当前账户信息：--->", masterAccount)

					switch v.MsgType {
					case "wechat":
						if masterAccount.Config != "" {
							err := notify.SendToWechatChan(v.MsgTo, masterAccount.Config, map[string]string{"content": v.MsgContent})
							if err != nil {
								status = 4
								b, _ := json.Marshal(err.Error())
								failureLog = string(b)

							}
							wechatCount++
							if wechatCount > total {
								wechatCount = 0
								FoundNextUpdateMaster(findAllX, svcCtx, ctx)
							}
						} else {
							status = 4
							failureLog = "微信帐号信息为空"
						}
					case "email":
						if masterAccount.Config != "" {
							err := notify.SendToEmailChan(v.MsgTo, v.MsgTitle, v.MsgContent, masterAccount.Config)
							if err != nil {
								status = 4
								b, _ := json.Marshal(err.Error())
								failureLog = string(b)

							}
							emailCount++
							if emailCount > total {
								emailCount = 0
								FoundNextUpdateMaster(findAllX, svcCtx, ctx)
							}
						} else {
							status = 4
							failureLog = "邮件帐号信息为空"
						}
					case "dingding":
						if masterAccount.Config != "" {
							dingChan, err := notify.SendToDingDingChan(masterAccount.Config, false, false)
							err = dingChan.SendMessage(v.MsgContent)
							if err != nil {
								status = 4
								b, _ := json.Marshal(err.Error())
								failureLog = string(b)

							}
							dingdingCount++
							if dingdingCount > total {
								dingdingCount = 0
								FoundNextUpdateMaster(findAllX, svcCtx, ctx)

							}
						} else {
							status = 4
							failureLog = "钉钉帐号信息为空"
						}
					case "feishu":
						if masterAccount.Config != "" {

							fsChan, err := notify.SendToFeishuChan(masterAccount.Config, false, false, svcCtx.Config.ApiAuthKey.CardSending)
							err = fsChan.SendMsg(v.MsgContent, v.MsgLevel, v.MsgTitle)
							if err != nil {
								status = 4
								b, _ := json.Marshal(err.Error())
								failureLog = string(b)
							}
							feishuCount++
							if feishuCount > total {
								feishuCount = 0
								FoundNextUpdateMaster(findAllX, svcCtx, ctx)
							}
						} else {
							status = 4
							failureLog = "飞书帐号信息为空"
						}

					}
					svcCtx.SendMsgRecordModel.UpdateStatusById(ctx, v.Id, status, failureLog)
					time.Sleep(1 * time.Second)
					logx.Error("结束...")

				} else {
					svcCtx.SendMsgRecordModel.UpdateStatusById(ctx, v.Id, 4, "该类型帐号信息为空")

				}
				time.Sleep(1 * time.Second)
			}

		}
	}

}

// 修改当前渠道列表的下一条数据，并修改为master
func FoundNextUpdateMaster(all *[]model.SendAccount, svcCtx *svc.ServiceContext, ctx context.Context) error {

	var one *model.SendAccount
	for i, v := range *all {
		if v.IsMaster == 1 {
			if len(*all) > 1 {
				if len(*all)-1 == i {
					one = &(*all)[0]
				} else {
					one = &(*all)[i+1]
				}

			} else {
				one = &(*all)[0]
			}
		}
	}
	if one == nil {
		return errors.New("修改失败")
	}
	//fmt.Println("修改下一条数据为master，信息：-->", one)
	err := svcCtx.SendAccountModel.UpdateMasterById(ctx, one.SendChannel, one.AppKey, one.Id)
	if err != nil {
		return err
	}
	return nil
}

// 计划提醒队列
func RemindJobs(svcCtx *svc.ServiceContext) {
	for {
		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		nums := r1.Intn(5)
		//fmt.Println("开始计划提醒随机等待", nums, "秒")
		time.Sleep(time.Duration(nums) * time.Second)

		var (
			timex string
		)
		totalMsgMap := make(map[string]string, 0)
		maintainList := make([]string, 0)
		mergeList := make([]string, 0)
		mergeGameList := make([]string, 0)

		ctx, _ := context.WithCancel(context.Background())
		//查询维护计划和合服计划信息
		msg, err := svcCtx.SendUserModel.FindALLCallMsg(ctx, 30)
		if err != nil {
			fmt.Println("查询提醒消息队列失败，原因：" + err.Error())
			return
		}
		//fmt.Println(msg)
		//处理数据生成对应计划的标题和内容
		for _, v := range *msg {
			if v.ProjectEn.String != "" {
				if v.Tags.String == "maintain_plan" {
					maintainList = append(maintainList, v.ProjectEn.String)
				} else if v.Tags.String == "merge_plan" {
					mergeList = append(mergeList, fmt.Sprintf("%s:%d", v.ProjectEn.String, v.Counts.Int64))
					mergeGameList = append(mergeGameList, v.ProjectEn.String)
				}
				timex = time.Unix(v.StartTime.Int64, 0).Format("2006-01-02 15:04:05")
			}
		}
		if len(maintainList) != 0 {
			totalMsgMap["maintain_plan"] = strings.Join(maintainList, "\n")
		}
		if len(mergeList) != 0 {
			totalMsgMap["merge_plan"] = strings.Join(mergeList, "\n")
		}
		if totalMsgMap != nil {
			flag := false
			for k, v := range totalMsgMap {
				fmt.Println("开始计划提醒任务", totalMsgMap)
				var (
					title   string
					content string
				)
				if k == "maintain_plan" {
					title = fmt.Sprintf("[提醒]%s有维护", timex)
					content = fmt.Sprintf(`需要维护的项目信息如下：
%s
`, v)

				} else if k == "merge_plan" {
					title = fmt.Sprintf("[提醒]%s有合服", timex)
					content = fmt.Sprintf(`需要合服的项目信息如下：
%s
`, v)
				}
				//循环加入消息队列
				for _, msgType := range []string{
					"wechat", "email", "feishu", "dingding", "wechat1",
				} {
					var contentx string
					if msgType == "wechat" || msgType == "dingding" || msgType == "wechat1" {
						contentx = fmt.Sprintf(`%s
==========================
%s
`, title, content)
					} else {
						contentx = content
					}

					list := make([]string, 0)
					appkey := "10011"
					types := msgType
					users := "yw"
					if msgType == "wechat1" {
						appkey = "10012"
						types = "wechat"
						users = "yy"
					}
					all, err := svcCtx.SendUserModel.FindAllByNames(ctx, users, types)
					if err != nil || len(*all) == 0 {
						fmt.Println("查询报警用户失败，原因：", err, all)
						return
					}
					for _, v1 := range *all {
						list = append(list, v1.Result)
					}
					duplicate := tool.RemoveDuplicate(list)
					msgx := new(model.SendMsgRecord)

					msgx.SendType = "3"
					msgx.MsgType = types
					msgx.Status = "1"
					msgx.MsgContent = contentx
					msgx.MsgTitle = title
					msgx.MsgTo = strings.Join(duplicate, ",")
					msgx.CreateDate = time.Now().Unix()
					msgx.AppKey = appkey
					msgx.MsgLevel = "remind"
					_, err = svcCtx.SendMsgRecordModel.Insert(ctx, msgx)
					if err != nil {
						fmt.Println("写入消息队列失败，原因：" + err.Error())
						return
					}
				}

				flag = true
			}
			if flag {
				time.Sleep(40 * time.Second)
			}
		}
	}

}
