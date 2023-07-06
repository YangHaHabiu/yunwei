package report

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"ywadmin-v3/common/crawler"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/qqmanger/bot_send"
	"ywadmin-v3/common/qqmanger/discuss"
	"ywadmin-v3/common/qqmanger/group"
	"ywadmin-v3/common/xcmd"
	"ywadmin-v3/service/qqGroup/api/internal/svc"
	"ywadmin-v3/service/qqGroup/api/internal/types"
	"ywadmin-v3/service/qqGroup/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(r *http.Request) (resp *types.ReportResp, err error) {

	out, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var data types.QqMessage
	l.Logger.Error(string(out))
	err = json.Unmarshal(out, &data)

	if err != nil {
		return nil, err
	}
	l.Logger.Error("全局master----------------->", globalkey.QqMsgKey)
	l.Logger.Error("全局管理员----------------->", globalkey.QqDefaltManger)
	//判断为master的qq信息进行存储
	if data.SelfID == globalkey.QqMsgKey[data.MessageType+"_qq"] && data.MessageType == "group" {
		_, err = l.svcCtx.QqMessageHistoryModel.Insert(&model.QqMessageHistory{
			MessageId: sql.NullString{
				String: data.MessageID,
				Valid:  false,
			},
			UserId:     data.UserID,
			Content:    data.Message,
			CreateTime: time.Now().Unix(),
			Executed:   0,
			GroupId: sql.NullInt64{
				Int64: data.GroupID,
				Valid: false,
			},
			GroupType: data.MessageType,
			SelfId:    data.SelfID,
			DiscussId: sql.NullInt64{
				Int64: data.DiscussId,
				Valid: false,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	//fmt.Println(data.Message)

	//任务开始
	if data.SelfID == globalkey.QqMsgKey[data.MessageType+"_qq"] {
		var data1 discuss.QqMessage
		copier.Copy(&data1, &data)
		switch data.MessageType {
		case "group":
			typeMap := map[string][]string{
				"updatex":  {"收到一条执行更新任务操作，正在执行中，请等待", "该任务已执行无误，请勿重复执行"},
				"crondx":   {"收到一条添加更新任务操作，正在操作中，请等待", "该任务已正确添加更新任务，请勿重复添加"},
				"installx": {"收到一条批量装服任务操作，正在装服中，请等待", "该任务已正确添加更新任务，请勿重复装服"},
				"otherx":   {"不存在回复操作类型，目前仅支持[执行、计划、装服]的口令", ""},
			}
			dir := l.svcCtx.Config.Project.MaintainFilePath
			username, ok := globalkey.QqDefaltManger[data.UserID]

			var (
				types    string
				location time.Time
			)

			rs, _ := regexp.Compile(`\[CQ.*\]|[ ]+`)
			dataTmp := rs.ReplaceAllString(data.Message, "")
			if len(dataTmp) > 0 && strings.Contains(data.Message, "CQ:reply,id=") {

				if dataTmp == "执行" {
					types = "updatex"
				} else if dataTmp == "计划" {
					types = "crondx"
				} else if dataTmp == "装服" {
					types = "installx"
				} else {
					types = ""
				}
				if types != "" {
					if !ok {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "非管理员禁止回复小助手执行操作", data.MessageType)
						return nil, nil
					}
					tmpSplit := strings.Split(data.Message, "reply,id=")
					mm := strings.Split(tmpSplit[1], "]")

					//根据msgid 查询记录
					msgObj, err := l.svcCtx.QqMessageHistoryModel.FindOneByMsgId(mm[0])
					if err != nil || msgObj.Content == "" {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "查询不到回复内容", data.MessageType)
						return nil, nil
					}
					if msgObj.Executed == 1 {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), typeMap[types][1], data.MessageType)
						return nil, nil
					}
					compilex, _ := regexp.Compile("计划|执行|装服")
					compileNew, _ := regexp.Compile(fmt.Sprintf(`\[CQ:at,qq=%d,.*\]`, data.SelfID))
					replyUser := compileNew.ReplaceAllString(data.Message, "")
					replyUser = compilex.ReplaceAllString(replyUser, "")
					compile2, _ := regexp.Compile(`\[CQ.*\]`)
					allString := compile2.ReplaceAllString(msgObj.Content, "")
					re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z|\t`)
					allString = re.ReplaceAllString(allString, "")
					gameSlice := strings.Split(allString, "\n")
					gameName := strings.Split(strings.TrimSpace(gameSlice[0]), " ")[0]
					//判断游戏是否存在
					if !group.Contains(strings.Split(l.svcCtx.Config.Project.SupportGame, ","), gameName) {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "游戏名错误，不允许执行", data.MessageType)
						return nil, nil
					}
					//判断仅停服需要检测开始时间
					compiley, _ := regexp.Compile(`仅停服|只更程序|关进程`)
					if compiley.MatchString(allString) {
						splits := strings.Split(allString, "\n")
						if strings.Contains(splits[1], "时间") {
							match, _ := regexp.Compile(`间:|间：`)
							i := match.Split(splits[1], -1)
							if len(i) > 1 {
								compile, _ := regexp.Compile(`月|日|:`)
								i2 := compile.Split(strings.ReplaceAll(i[1], " ", ""), -1)
								if len(i2) != 4 {
									bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "格式化日期错误", data.MessageType)
									return nil, nil

								}
								month, _ := strconv.Atoi(i2[0])
								day, _ := strconv.Atoi(i2[1])
								hour, _ := strconv.Atoi(i2[2])
								minute, _ := strconv.Atoi(i2[3])
								all := fmt.Sprintf("%.2d-%.2d %.2d:%.2d", month, day, hour, minute)
								format := time.Now().Format("01-02 15:04")
								inLocation, _ := time.ParseInLocation("01-02 15:04", format, time.Local)
								location, err = time.ParseInLocation("01-02 15:04", all, time.Local)
								if err != nil {
									bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "格式化日期错误:"+err.Error(), data.MessageType)
									return nil, nil
								}
								sub := inLocation.Sub(location)

								if sub.Seconds() < 0 && types == "updatex" {
									bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "未到更新执行时间，请稍后在执行", data.MessageType)
									return nil, nil
								}
							} else {
								bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "更新时间格式错误", data.MessageType)
								return nil, nil
							}

						} else {
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "没检测更新时间", data.MessageType)
							return nil, nil
						}

					}

					//操作开始
					file := time.Now().Format("20060102150405")
					lockFile := fmt.Sprintf("%s/format_files/reply.%s.lock", dir, gameName)
					fileName := fmt.Sprintf("%s/format_files/reply_%s_%s.txt", dir, gameName, file)
					mainLog := fmt.Sprintf("%s/format_files/main.%s.log", dir, gameName)
					installLockFile := fmt.Sprintf("%s/Log/main.lock", l.svcCtx.Config.Project.InstallFilePath)
					//判断是否有正在执行的锁文件
					var (
						fileLocks string
						tips      string
					)
					if types == "installx" {
						fileLocks = installLockFile
						tips = "批量装服存在锁文件，请稍后再试，路径：" + fileLocks
					} else {
						fileLocks = lockFile
						tips = gameName + "存在锁文件，请稍后执行，路径：" + fileLocks
					}
					_, err = os.Stat(fileLocks)
					if err == nil {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), tips, data.MessageType)
						return nil, nil
					}
					os.Remove(mainLog)
					//开始生成热更文件
					err = ioutil.WriteFile(fileName, []byte(allString), 0666)
					if err != nil {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "生成文件失败，请检查", data.MessageType)
						return nil, nil
					}
					//开始执行
					os.Create(fileLocks)
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), typeMap[types][0], data.MessageType)
					l.Logger.Error(msgObj.MessageId.String)
					if types == "updatex" {
						cmdsExec := fmt.Sprintf("cd %s;sh cmd_entrance.sh -g %s -f %s -op all -t QWEB -r 0 -u %s >> %s", dir, gameName, fileName, username, mainLog)
						job := xcmd.NewCommandJob(48*time.Hour, cmdsExec)
						senders := fmt.Sprintf("[CQ:at,qq=%d]", data.UserID)
						var (
							operStatus string
						)
						finishTime := time.Now().Format("2006-01-02 15:04:05")

						if job.IsOk {
							operStatus = "执行完成"
							//只有成功了才可以修改状态
							if err := l.svcCtx.QqMessageHistoryModel.UpdateExecuted(msgObj.MessageId.String); err != nil {
								fmt.Println("修改状态失败--->", msgObj.MessageId)
							}
						} else {
							operStatus = "执行失败"
						}
						_, allList := group.ReadLine(fmt.Sprintf("%s/tmp/format/%s/task_brief.txt", dir, gameName))
						for i, v := range allList {
							if i == 0 {
								result := fmt.Sprintf(`%s
操作状态：%s
游戏名：%s
操作者：%s
完成时间：%s
操作内容：
------------
%s
`, replyUser+senders, operStatus, gameName, username, finishTime, strings.Join(v, ""))
								bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), result, data.MessageType)
							} else {
								time.Sleep(1 * time.Second)
								if i == len(allList)-1 {
									bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), fmt.Sprintf("%s\n#####################\n本次通知消息共%d段，已发送完毕", strings.Join(v, ""), len(allList)), data.MessageType)
								} else {
									bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), strings.Join(v, ""), data.MessageType)
								}
							}
						}

					} else if types == "crondx" {
						inLocation, _ := time.ParseInLocation("01-02 15:04", time.Now().Format("01-02 15:04"), time.Local)
						now := time.Now()
						if inLocation.Month() > location.Month() {
							now = now.AddDate(1, 0, 0)
						}
						startTime := fmt.Sprintf("%d-%s", now.Year(), location.Format("01-02 15:04"))
						postData := url.Values{}
						postData.Set("operator", username)
						postData.Set("startTime", startTime)
						postData.Set("game", gameName)
						postData.Set("content", allString)
						response, err := crawler.PostHttpResponse(l.svcCtx.Config.Project.YwTaskAddApiUrl, postData.Encode(), false)
						if err != nil {
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "请求添加计划失败", data.MessageType)
							os.Remove(fileLocks)
							return nil, nil
						}
						fmt.Println(string(response))
						var msg group.GetResult
						err = json.Unmarshal(response, &msg)
						if err != nil {
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "解析请求信息失败"+err.Error(), data.MessageType)
							os.Remove(fileLocks)
							return nil, nil
						}
						if msg.Code != 0 {
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), msg.Msg, data.MessageType)
							os.Remove(fileLocks)
							return nil, nil
						}
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), msg.Msg, data.MessageType)

						//只有成功了才可以修改状态
						if err := l.svcCtx.QqMessageHistoryModel.UpdateExecuted(msgObj.MessageId.String); err != nil {
							l.Logger.Error("修改状态失败--->", msgObj.MessageId)

						}
					} else if types == "installx" {
						senders := fmt.Sprintf("[CQ:at,qq=%d]", data.UserID)
						var endTime, installPlat string
						split := strings.Split(allString, "\n")
						compile := regexp.MustCompile(`:|：`)
						installTyps := "-a all"
						for _, x := range split {
							i := strings.TrimSpace(x)
							if strings.Contains(i, "时间") {
								//校验时间
								_, err := time.Parse("2006-01-02", compile.Split(i, -1)[1])
								if err != nil {
									bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), replyUser+senders+"\n输入时间范围错误，例如：结束时间：2022-01-20", data.MessageType)
									os.Remove(fileLocks)
									return nil, nil
								}
								startTime := time.Now().Format("2006-01-02")
								endTime = fmt.Sprintf("-bt %s -et %s", startTime, compile.Split(i, -1)[1])
							} else if strings.Contains(i, "平台") {
								installPlat = fmt.Sprintf("-p %s", compile.Split(i, -1)[1])
							} else if strings.Contains(i, "跨服") {
								installTyps = "-a cross"
							} else if strings.Contains(i, "单服") {
								installTyps = "-a game"
							}
						}

						cmdsExec := fmt.Sprintf("sh %s/main.sh -s all %s -w -u %s -g %s %s %s", l.svcCtx.Config.Project.InstallFilePath, installTyps, username, gameName, installPlat, endTime)
						//fmt.Println(cmdsExec)
						job := xcmd.NewCommandJob(48*time.Hour, cmdsExec)
						var status, errors string

						if job.IsOk {
							//只有成功了才可以修改状态
							if err := l.svcCtx.QqMessageHistoryModel.UpdateExecuted(msgObj.MessageId.String); err != nil {
								l.Logger.Error("修改状态失败--->", msgObj.MessageId)
							}
							status = "0"
						} else {
							status = "1"
							errors = "[ERROR]"
							l.Logger.Error("语句："+cmdsExec+"\n装服错误信息-->:", err)
						}
						cmdsExecx := fmt.Sprintf(" cat %s/Log/tips_run_main.log |sed '/^$/d'", l.svcCtx.Config.Project.InstallFilePath)
						jobx := xcmd.NewCommandJob(48*time.Hour, cmdsExecx)
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), replyUser+senders+"\n"+jobx.OutMsg, data.MessageType)

						//插入日志
						postData := url.Values{}
						postData.Add("title", fmt.Sprintf("%s%s-安装新服-%s", errors, gameName, time.Now().Format("150405")))
						postData.Add("gameName", gameName)
						postData.Add("operType", "install_game")
						postData.Add("operator", username)
						postData.Add("operStatus", status)
						postData.Add("types", "1")
						postData.Add("operContent", job.OutMsg+job.ErrMsg)
						ts := fmt.Sprintf("%d", time.Now().Unix())
						_, err := crawler.PostHttpResponse(fmt.Sprintf("%s?sn=%s&ts=%s", l.svcCtx.Config.Project.HotUpdateApiUrl, group.Md5s(group.CreateSn(ts)), ts), postData.Encode(), false)
						if err != nil {
							l.Logger.Error("插入日志失败:" + err.Error())
						}

					}
					os.Remove(fileLocks)

				}

			} else if data.Message == "切换小助手旧" {
				_, ok := globalkey.QqDefaltManger[data.UserID]
				if !ok {
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "非管理员禁止回复小助手执行操作", data.MessageType)
					return nil, nil
				}
				if len(globalkey.QqGroupList) == 1 {
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "您就一个QQ小助手，不用再切啦", data.MessageType)
					return nil, nil
				}
				m := 0
				for i := 0; i < len(globalkey.QqGroupList); i++ {
					if globalkey.QqGroupList[i] == globalkey.QqMsgKey["group_qq"] {
						m = i
						break
					}
				}
				m = (m + 1) % len(globalkey.QqGroupList)
				l.svcCtx.QqLoadBalanceModel.UpdateIsMaster(data.MessageType, globalkey.QqMsgKey["group_qq"].(int64), 0)
				l.svcCtx.QqLoadBalanceModel.UpdateIsMaster(data.MessageType, globalkey.QqGroupList[m], 1)
				filters, _ := l.svcCtx.QqLoadBalanceModel.FindMasterByFilters("qq__=", globalkey.QqGroupList[m], "group_type__=", data.MessageType)
				globalkey.QqMsgKey["group_qq"] = (*filters)[0].Qq
				globalkey.QqMsgKey["group_qqapi"] = (*filters)[0].QqApi
				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), fmt.Sprintf("接下来由QQ号：%d为您服务", globalkey.QqMsgKey["group_qq"]), data.MessageType)
			} else if data.Message == "帮助" {
				help, _ := ioutil.ReadFile(fmt.Sprintf("%s/README", dir))
				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), string(help), data.MessageType)
				return nil, nil
			} else if data.Message == "平台格式" {
				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), `平台名与区服之间用空格分隔。例:37wan 1
平台与平台之间用英文逗号分隔。例:37wan 1,r2games 1
整个平台与平台之间。例:jswar,cjzz
区服之间,如果不连续，用/分隔。例:37wan 1/3/7
区服之间,如果连续，用-分隔。例:37wan 1-3
完整示例:37wan 1/3/5-8/10/13,r2games 2/5-10/12/15,jswar,cjzz
`, data.MessageType)
				return nil, nil
			} else if data.Message == "同步后台" && data.GroupID == 324113338 {
				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "正在同步中，请稍等...", data.MessageType)
				job := xcmd.NewCommandJob(48*time.Minute, "sh /root/zhangqingxian/yunwei_new/old_export_new/rsync_yunwei.sh >> /root/zhangqingxian/yunwei_new/old_export_new/log/rsync_yunwei.log")
				var msg string
				if job.IsOk {
					msg = "同步后台成功"
				} else {
					msg = "同步后台失败"
				}
				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), msg, data.MessageType)
				return nil, nil

			} else if data.Message == "小助手关键字" {
				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), `小助手常用指令：
帮助：获取更新关键字文档
同步后台：同步旧运维后台的平台、区服（仅在运维群执行）
切换小助手：发送切换小助手的命令
平台格式：展示常规的平台与服范围格式
游戏名 小助手更新模板：输出更新模板相关信息，不指定游戏名，则导出通用模板
小助手装服模板：输出装服模板相关信息
回复操作：
目前仅支持[执行、计划、装服]的口令（不用再艾特小助手啦）
`, data.MessageType)
			} else if strings.Contains(data.Message, "小助手更新模板") {
				split := strings.Split(data.Message, "小助手更新模板")
				game := strings.TrimSpace(strings.Join(split, ""))
				if game == "" {
					game = "all"
				}

				job := xcmd.NewCommandJob(20*time.Minute, fmt.Sprintf("sh %s/template.sh -g %s", l.svcCtx.Config.Project.MaintainFilePath, game))
				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), job.OutMsg+"\n"+job.ErrMsg, data.MessageType)
				return nil, nil

			} else if strings.Contains(data.Message, "小助手装服模板") {
				msg := `游戏英文名（必填）
安装新服|安装跨服|安装单服（必填3选1）
结束时间：2022-05-17（可选）
平台：xxx（可选）
---------------------
说明：
仅有单服计划：关键字 安装新服|安装单服 都可
仅有跨服计划：关键字只能是 安装跨服
跨服和单服计划都有：关键字只能是 安装新服
`
				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), msg, data.MessageType)
				return nil, nil
			}

		case "discuss":

			discuss.ReportDiscuss(data1, "discuss")
		}

	}
	return nil, nil
}
