package report

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"ywadmin-v3/common/crawler"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"

	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/qqmanger/bot_send"
	"ywadmin-v3/common/qqmanger/discuss"
	"ywadmin-v3/common/qqmanger/group"
	"ywadmin-v3/common/xcmd"
	"ywadmin-v3/service/qqGroupv3/api/internal/svc"
	"ywadmin-v3/service/qqGroupv3/api/internal/types"
	"ywadmin-v3/service/qqGroupv3/model"
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
	l.Logger.Error(globalkey.QqMsgKey[data.MessageType+"_qq"])
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
				"combinex": {"收到一条批量合服任务操作，正在合服中，请等待", "该任务已正确添加合服任务，请勿重复合服"},
				"otherx":   {"不存在回复操作类型，目前仅支持[计划执行、计划维护、计划装服]的口令", ""},
			}
			//dir := l.svcCtx.Config.Project.MaintainFilePath
			username, ok := globalkey.QqDefaltManger[data.UserID]

			var (
				types       string
				location    time.Time
				projectId   int64
				taskTypex   string
				startTimex  string
				montainIdx  int64
				titlex      string
				uidx        string
				inLocationx time.Time
			)

			rs, _ := regexp.Compile(`\[CQ.*\]|[ ]+`)
			dataTmp := rs.ReplaceAllString(data.Message, "")
			if len(dataTmp) > 0 && strings.Contains(data.Message, "CQ:reply,id=") {

				if dataTmp == "计划执行" {
					types = "updatex"
				} else if dataTmp == "计划维护" {
					types = "crondx"
				} else if dataTmp == "计划装服" {
					types = "installx"
				} else if dataTmp == "计划合服" {
					types = "combinex"
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

					// //根据msgid 查询记录
					msgObj, err := l.svcCtx.QqMessageHistoryModel.FindOneByMsgId(mm[0], data.NewMessageID)

					//msgObj, err := l.svcCtx.QqMessageHistoryModel.FindOneBySeqId(data.Source.UserID, data.Source.Seq, data.Source.Rand)
					if err != nil || msgObj.Content == "" {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "查询不到回复内容", data.MessageType)
						return nil, nil
					}
					if msgObj.Executed == 1 {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), typeMap[types][1], data.MessageType)
						return nil, nil
					}
					compilex, _ := regexp.Compile("计划执行|计划维护|计划装服|计划合服")
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
					//compiley, _ := regexp.Compile(`仅停服|只更程序|关进程`)
					//if compiley.MatchString(allString) {
					if types != "installx" {
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
								//format := time.Now().Format("01-02 15:04")
								//inLocation, _ := time.ParseInLocation("01-02 15:04", format, time.Local)
								location, err = time.ParseInLocation("01-02 15:04", all, time.Local)
								if err != nil {
									bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "格式化日期错误:"+err.Error(), data.MessageType)
									return nil, nil
								}
								//sub := inLocation.Sub(location)

								//if sub.Seconds() < 0 && types == "updatex" {
								//	bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "未到更新执行时间，请稍后在执行", data.MessageType)
								//	return nil, nil
								//}
							} else {
								bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "更新时间格式错误", data.MessageType)
								return nil, nil
							}

						} else {
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "没检测更新时间", data.MessageType)
							return nil, nil
						}

						//}

						inLocation, _ := time.ParseInLocation("01-02 15:04", time.Now().Format("01-02 15:04"), time.Local)
						now := time.Now()
						if inLocation.Month() > location.Month() {
							now = now.AddDate(1, 0, 0)
						}
						startTime := fmt.Sprintf("%d-%s", now.Year(), location.Format("01-02 15:04"))
						inLocationx, _ = time.ParseInLocation("2006-01-02 15:04", startTime, time.Local)
					}
					//判断是否有正在执行的锁文件
					var (
						fileLocks string
						tips      string
						lockName  string
					)
					//fileLocks = fmt.Sprintf("%s/%s", l.svcCtx.Config.Project.LockFilePath, lockName)
					if types == "installx" {
						tips = "批量装服存在锁文件，请稍后再试"
						lockName = types + ".lock"
						fileLocks = fmt.Sprintf("%s/%s", l.svcCtx.Config.Project.InstallFilePath, lockName)
					} else {
						lockName = types + "_" + gameName + ".lock"
						tips = gameName + "执行存在锁文件，请稍后执行"
						fileLocks = fmt.Sprintf("%s/%s", l.svcCtx.Config.Project.LockFilePath, lockName)
					}

					// _, err = os.Stat(fileLocks)
					// if err == nil {
					// 	bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), tips, data.MessageType)
					// 	return nil, nil
					// }
					if tool.IsExist(fileLocks) {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), tips, data.MessageType)
						return nil, nil
					}

					if !tool.IsExist(l.svcCtx.Config.Project.LockFilePath) {
						tool.CreateMutiDir(l.svcCtx.Config.Project.LockFilePath)
					}

					//开始执行
					os.Create(fileLocks)
					list, _ := l.svcCtx.AdminRpc.ProjectList(l.ctx, &adminclient.ProjectListReq{ProjectEn: gameName})
					if len(list.List) == 1 {
						projectId = list.List[0].ViewProjectId
					}

					userList, _ := l.svcCtx.AdminRpc.UserList(l.ctx, &adminclient.UserListReq{Name: username})
					if len(userList.List) == 1 {
						uidx = gconv.String(userList.List[0].Id)
					}
					if types == "installx" {
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), typeMap[types][0], data.MessageType)
					}
					l.Logger.Error(msgObj.MessageId.String)
					if types == "updatex" || types == "crondx" {
						if strings.Contains(allString, "添加新平台") {
							dir := l.svcCtx.Config.Project.MaintainFilePath
							file := time.Now().Format("20060102150405")
							fileName := fmt.Sprintf("%s/format_files/reply_%s_%s.txt", dir, gameName, file)
							err = ioutil.WriteFile(fileName, []byte(allString), 0666)
							if err != nil {
								os.Remove(fileLocks)
								bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "生成文件失败，请检查", data.MessageType)
								return nil, nil
							}
							cmdsExec := fmt.Sprintf("source /etc/profile;cd %s;sh cmd_entrance.sh -g %s -f %s -op all -t QWEB -r 0 -u %s", dir, gameName, fileName, username)
							job := xcmd.NewCommandJob(10*time.Minute, cmdsExec)
							if !job.IsOk {
								os.Remove(fileLocks)
								bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "添加新平台失败"+job.ErrMsg, data.MessageType)
								return nil, nil
							}
							os.Remove(fileLocks)
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "添加新平台成功", data.MessageType)
							return nil, nil
						}
						formatJson, err := l.svcCtx.YunWeiRpc.TaskGetFormatJson(l.ctx, &yunweiclient.TaskGetFormatJsonReq{
							Game:    gameName,
							Content: allString,
							Uname:   username,
						})
						if err != nil {
							os.Remove(fileLocks)
							l.Logger.Error(fmt.Sprintf("格式化错误:%v", err))
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "格式化文件失败，请检查", data.MessageType)
							return nil, nil
						}

						if strings.Contains(allString, "仅停服更新") && types == "updatex" {
							os.Remove(fileLocks)
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "此任务需要关联维护计划，请执行[计划维护]口令", data.MessageType)
							return nil, nil
						}
						//获取平台
						var (
							platforms string
							datax     string
						)
						taskListObj := make([]model.OperationList, 0)
						for i, v := range formatJson.OperationListM {
							datax += fmt.Sprintf("==>任务[%d]", i+1)
							operlistForm := model.OperationList{}
							ListFormObj := make([]model.TaskCommonJson, 0)
							for _, v1 := range v.OperationListForm {
								tmp := model.TaskCommonJson{
									Operation:      v1.Operation,
									Stable:         v1.Stable,
									OuterIp:        v1.OuterIp,
									DbUpdate:       v1.DbUpdate,
									FileList:       v1.FileList,
									MaintainRange:  v1.MaintainRange,
									CmdList:        v1.CmdList,
									DbType:         v1.DbType,
									SQLCmd:         v1.SQLCmd,
									Merge:          v1.Merge,
									ExportFileName: v1.ExportFileName,
									PlatName:       v1.PlatName,
									CheckSt:        v1.CheckSt,
									InitSetTime:    v1.InitSetTime,
									ExecuteSQL:     v1.ExecuteSQL,
									AddRestartGame: v1.AddRestartGame,
									ExecuteFlag:    v1.ExecuteFlag,
								}
								ListFormObj = append(ListFormObj, tmp)
								switch v1.Operation {
								case "client_hot":
									datax += fmt.Sprintf(">%s(%s)", v1.OperationCn, v1.PlatName)
								default:
									datax += fmt.Sprintf(">%s(%s)", v1.OperationCn, v1.MaintainRange)
								}

								split := strings.Split(v1.MaintainRange, ",")
								for _, v2 := range split {
									split2 := strings.Split(v2, " ")
									platforms += split2[0] + ","
								}
							}
							operlistForm.OperationListForm = ListFormObj
							taskListObj = append(taskListObj, operlistForm)
						}

						idata := strings.Split(datax, "==>")
						content := strings.Join(idata[1:len(idata)], "==>")
						split := strings.Split(platforms, ",")
						newPlatforms := strings.Join(split[0:len(split)-1], ",")

						marshal, _ := json.Marshal(taskListObj)
						//fmt.Println(newPlatforms)
						//根据平台查集群
						info, err := l.svcCtx.YunWeiRpc.PlatformGetClusterInfo(l.ctx, &yunweiclient.GetClusterByPlatformReq{
							ProjectId:   projectId,
							PlatformEns: newPlatforms,
						})
						if err != nil {
							os.Remove(fileLocks)
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "根据平台查集群信息失败", data.MessageType)
							return nil, nil
						}
						tmpIp := make([]string, 0)
						tmpLabelId := make([]string, 0)
						for _, v := range info.Data {
							outerIp := strings.TrimSpace(v.OuterIp)
							if outerIp != "" {
								tmpIp = append(tmpIp, outerIp)
							}
							if v.ClusterLabelId != "" {
								tmpLabelId = append(tmpLabelId, v.ClusterLabelId)
							}
						}
						duplicateIps := tool.RemoveDuplicate(tmpIp)
						duplicateLabelIds := tool.RemoveDuplicate(tmpLabelId)

						switch types {
						case "updatex":
							taskTypex = "1"
							//startTimex = gconv.String(time.Now().Unix())
							titlex = "小助手临时维护"
						case "crondx":
							taskTypex = "2"
							//根据维护时间查维护计划
							listx, err := l.svcCtx.YunWeiRpc.MaintainPlanList(l.ctx, &yunweiclient.ListMaintainPlanReq{
								ProjectId:     projectId,
								PlanStartTime: inLocationx.Unix(),
								Current:       0,
								PageSize:      0,
								//MaintainType:  "1",
								TaskId: -1,
							})
							if err != nil || len(listx.Rows) != 1 {
								l.Logger.Error(err)
								os.Remove(fileLocks)
								bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "查询维护计划失败，请检查", data.MessageType)
								return nil, nil
							}
							titlex = listx.Rows[0].Title
							montainIdx = listx.Rows[0].Id
						}
						if inLocationx.Unix() < time.Now().Unix() {
							startTimex = gconv.String(time.Now().Unix())
						} else {
							startTimex = gconv.String(inLocationx.Unix())
						}
						inputObj := new(yunweiclient.AddTasksReq)
						inputObj.Content = content
						inputObj.ProjectId = projectId
						inputObj.StartTime = startTimex
						inputObj.TaskType = taskTypex
						inputObj.TaskListForm = string(marshal)
						inputObj.MaintainId = montainIdx
						inputObj.Title = titlex
						inputObj.OuterIp = strings.Join(duplicateIps, ",")
						inputObj.ClusterId = strings.Join(duplicateLabelIds, ",")
						inputObj.Uid = uidx
						//添加计划任务
						_, err = l.svcCtx.YunWeiRpc.TasksAdd(l.ctx, inputObj)
						if err != nil {
							os.Remove(fileLocks)
							bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), fmt.Sprintf("新增任务队列失败，原因：%v", err), data.MessageType)
							return nil, nil
						}
						//修改qq信息，以防止二次执行
						//只有成功了才可以修改状态
						if err := l.svcCtx.QqMessageHistoryModel.UpdateExecuted(msgObj.MessageId.String); err != nil {
							l.Logger.Error("更新修改状态失败--->", msgObj.MessageId)
						}
						os.Remove(fileLocks)
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "新增操作队列成功", data.MessageType)
						return nil, nil
					} else if types == "installx" {
						//if l.svcCtx.Config.YwQQGroup != "629049295" {
						replyUser = strings.ReplaceAll(replyUser, "]", ",message_id="+msgObj.MessageId.String+"]")
						//}
						var senders string
						if msgObj.UserId == data.UserID {
							senders = fmt.Sprintf("[CQ:at,qq=%d]", data.UserID)
						} else {
							senders = fmt.Sprintf("[CQ:at,qq=%d][CQ:at,qq=%d]", data.UserID, msgObj.UserId)
						}
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

						cmdsExec := fmt.Sprintf("sh %s/main.sh -s all %s -w QWEB -u %s -g %s %s %s",
							l.svcCtx.Config.Project.InstallFilePath, installTyps, username, gameName, installPlat, endTime)
						job := xcmd.NewCommandJob(48*time.Hour, cmdsExec)
						var (
							status int64
							errorx string
						)

						if job.IsOk {
							//只有成功了才可以修改状态
							if err := l.svcCtx.QqMessageHistoryModel.UpdateExecuted(msgObj.MessageId.String); err != nil {
								l.Logger.Error("装服修改状态失败--->", msgObj.MessageId)
							}
							status = 0
						} else {
							status = 1
							errorx = "[ERROR]"
							l.Logger.Error("语句："+cmdsExec+"\n装服错误信息-->:", err)
						}
						cmdsExecx := fmt.Sprintf(" cat %s/Log/tips_run_main.log |sed '/^$/d'", l.svcCtx.Config.Project.InstallFilePath)
						jobx := xcmd.NewCommandJob(48*time.Hour, cmdsExecx)
						bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), replyUser+senders+"\n"+jobx.OutMsg, data.MessageType)
						//插入日志
						title := fmt.Sprintf("%s%s-安装新服-%s", errorx, gameName, time.Now().Format("150405"))
						hotLogHistory := new(yunweiclient.ListHotLogHistoryData)
						hotLogHistory.HotTitle = title
						hotLogHistory.OperStatus = status
						hotLogHistory.OperContent = job.OutMsg + job.ErrMsg
						hotLogHistory.ProjectId = projectId
						hotLogHistory.CreateBy = uidx
						hotLogHistory.OperType = "install_game"
						ctx, _ := context.WithCancel(context.Background())
						_, err = l.svcCtx.YunWeiRpc.HotLogHistoryAdd(ctx, &yunweiclient.AddHotLogHistoryReq{Data: hotLogHistory, Uid: uidx})
						if err != nil {
							l.Logger.Error("插入日志失败:" + err.Error())
						}
						os.Remove(fileLocks)
					}

				}
			} else if data.Message == "你好小助手" {
				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "你好，欢迎使用v3小助手", data.MessageType)
				return nil, nil
			} else if strings.Contains(data.Message, "获取服务器信息") || strings.Contains(data.Message, "获取帮助") {
				all := bytes.ReplaceAll(bytes.ReplaceAll([]byte(data.Message), []byte("\r\n"), []byte("\n")), []byte("\r"), []byte("\n"))
				allString := string(all)
				//fmt.Println(allString)
				file := time.Now().Format("20060102150405")
				dir := l.svcCtx.Config.Project.MaintainFilePath
				username, _ := globalkey.QqDefaltManger[data.UserID]
				re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z|\t`)
				allString = re.ReplaceAllString(allString, "")
				if strings.Contains(allString, "获取帮助") {
					cmdsExec := fmt.Sprintf("cd %s;cat README", dir)
					job := xcmd.NewCommandJob(48*time.Hour, cmdsExec)
					senders := fmt.Sprintf("[CQ:at,qq=%d]", data.UserID)
					result := fmt.Sprintf(`%s
%s
`, senders, job.OutMsg)
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), result, data.MessageType)
					return nil, nil
				}
				gameSlice := strings.Split(allString, "\n")
				gameName := strings.Split(strings.TrimSpace(gameSlice[0]), " ")[0]
				//判断游戏是否存在
				if !group.Contains(strings.Split(l.svcCtx.Config.Project.SupportGame, ","), gameName) {
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "游戏名错误，不允许执行", data.MessageType)
					return nil, nil
				}

				fileName := fmt.Sprintf("%s/format_files/reply_%s_%s.txt", dir, gameName, file)
				//fmt.Println(fileName)
				ioutil.WriteFile(fileName, []byte(allString), 0666)
				cmdsExec := fmt.Sprintf("cd %s;source /etc/profile && sh cmd_entrance.sh -g %s -f %s -op all -t QWEB -r 0 -u %s", dir, gameName, fileName, username)
				job := xcmd.NewCommandJob(48*time.Hour, cmdsExec)
				senders := fmt.Sprintf("[CQ:at,qq=%d]", data.UserID)
				var (
					operStatus string
				)
				finishTime := time.Now().Format("2006-01-02 15:04:05")

				if job.IsOk {
					operStatus = "执行完成"
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
`, senders, operStatus, gameName, username, finishTime, strings.Join(v, ""))
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
			} else if data.Message == "切换小助手" {
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
				return nil, nil
			} else if data.Message == "切换线路" {
				if !ok {
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "非管理员禁止回复小助手执行操作", data.MessageType)
					return nil, nil
				}
				if l.svcCtx.Config.YwQQGroup != "629049295" {
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "切换线路功能只适用于黑石", data.MessageType)
					return nil, nil
				}

				url := "http://244.mengzuofang.com:8081"
				setCookie, err := AikuaiLoginFunc(url)
				if err != nil {
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "请求不到爱快后台地址", data.MessageType)
					return nil, nil
				}
				msg, err := AikuaiSwitchLineFunc(url, setCookie)
				if err != nil {
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "切换VPN线路失败", data.MessageType)
					return nil, nil
				}
				if msg != "" {
					bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), msg, data.MessageType)
					time.Sleep(2 * time.Second)
				}

				bot_send.BotSend(data.GroupID, globalkey.QqMsgKey["group_qqapi"].(string), "切换VPN线路成功", data.MessageType)
				return nil, nil
			}

		}
	}
	return nil, nil
}

type AikuaiResp struct {
	Result int    `json:"Result"`
	ErrMsg string `json:"ErrMsg"`
	Data   struct {
		Total int `json:"total"`
		Data  []struct {
			Interface string `json:"interface"`
			Name      string `json:"name"`
			Timestamp int    `json:"timestamp"`
		} `json:"data"`
		Interface [][]string `json:"interface"`
	}
}
type NewAikuaiResp struct {
	Result int    `json:"Result"`
	ErrMsg string `json:"ErrMsg"`
	Data   struct {
		IfaceCheck []struct {
			ID              int    `json:"id"`
			Interface       string `json:"interface"`
			ParentInterface string `json:"parent_interface"`
			IPAddr          string `json:"ip_addr"`
			Gateway         string `json:"gateway"`
			Internet        string `json:"internet"`
			Updatetime      string `json:"updatetime"`
			AutoSwitch      string `json:"auto_switch"`
			Result          string `json:"result"`
			Errmsg          string `json:"errmsg"`
			Comment         string `json:"comment"`
		} `json:"iface_check"`
	} `json:"Data"`
}

// 爱快登录函数
func AikuaiLoginFunc(url string) (string, error) {
	datas := `{
		"username": "admin",
		"passwd": "c63e88cbfc879f3610a220a36f34dd3f",
		"pass": "c2FsdF8xMXlteGgoMTIzKQ==",
		"remember_password": "true"
	  }`
	err, s := crawler.PostNew(url+"/Action/login", datas, "", true)
	if err != nil {
		return "", err
	}
	var tmp AikuaiResp
	header := strings.Split(s, "@@")[1]
	dataHeader := strings.Split(s, "@@")[0]
	json.Unmarshal([]byte(dataHeader), &tmp)
	if tmp.ErrMsg != "Succeess" {
		return "", err
	}
	setCookie := "username=admin; login=1; " + header
	return setCookie, nil
}

// 切换线路
func AikuaiSwitchLineFunc(url, setCookie string) (string, error) {
	datas := `{
		"func_name": "l2tp_client",
		"action": "show",
		"param": {
		  "TYPE": "total,data,interface"
		}
	  }`
	err, s2 := crawler.PostNew(url+"/Action/call", datas, setCookie, false)
	if err != nil {
		return "", err
	}
	var tmp2 AikuaiResp
	json.Unmarshal([]byte(s2), &tmp2)
	if tmp2.ErrMsg != "Success" {
		return "", errors.New("切换失败")
	}

	lineList := make([]string, 0)
	for _, v := range tmp2.Data.Interface {
		if !group.Contains([]string{"auto", "vwan100", "adsl5"}, v[0]) {
			lineList = append(lineList, v[0])
		}
	}
	index := 0
	for i, v := range lineList {
		if v == tmp2.Data.Data[1].Interface {
			index = i
		}
	}
	index = (index + 1) % len(lineList)
	datas = fmt.Sprintf(`
	{
		"func_name": "l2tp_client",
		"action": "edit",
		"param": {
		  "name": "l2tp_tfgame2",
		  "comment": "",
		  "server": "office.gztfgame.com",
		  "gateway": "172.30.30.1",
		  "server_port": 1701,
		  "username": "heystone",
		  "passwd": "yw@dqw.com",
		  "ipsec_secret": "TFgame",
		  "interface": "%s",
		  "leftid": "",
		  "rightid": "",
		  "mru": 1400,
		  "timing_rst_switch": 0,
		  "timing_rst_week": "1234567",
		  "timing_rst_time": "00:00,,",
		  "cycle_rst_time": 0,
		  "qos_switch": 0,
		  "updatetime": %d,
		  "dns2": "",
		  "dns1": "",
		  "mppe": "",
		  "ip_addr": "172.30.30.4",
		  "id": 2,
		  "enabled": "yes",
		  "mtu": 1400,
		  "week": "1234567",
		  "mon9": 0,
		  "date1": "00:00",
		  "date2": "",
		  "date3": ""
		}
	  }
	`, lineList[index], time.Now().Unix())
	err, s3 := crawler.PostNew(url+"/Action/call", datas, setCookie, false)
	if err != nil {
		return "", err
	}
	var tmp3 AikuaiResp
	json.Unmarshal([]byte(s3), &tmp3)
	if tmp3.ErrMsg != "Success" {
		return "", errors.New("切换失败")
	}
	return fmt.Sprintf("VPN线路由原%s，切换到%s", tmp2.Data.Data[1].Interface, lineList[index]), nil
}

// 利用日志判断切换线路的条件之一
func AikuaiJudgeLine(url, setCookie string) bool {
	datas := `{
		"func_name": "syslog-wanpppoe",
		"action": "show",
		"param": {
		  "TYPE": "total,data,interface",
		  "ORDER_BY": "id",
		  "ORDER": "desc",
		  "limit": "0,20",
		  "FINDS": "content",
		  "KEYWORDS": "%s",
		  "FILTER1": "interface,==,l2tp_tfgame2"
		}
	  }`
	var noramlTime, exitTime int
	for i, v := range []string{"Connect:%20l2tp_tfgame2", "Exit."} {
		data := fmt.Sprintf(datas, v)
		err, s2 := crawler.PostNew(url+"/Action/call", data, setCookie, false)
		if err != nil {
			return false
		}
		var tmp2 AikuaiResp
		json.Unmarshal([]byte(s2), &tmp2)
		if tmp2.ErrMsg != "Success" {
			return false
		}
		time.Sleep(time.Second)
		if i == 0 {
			noramlTime = tmp2.Data.Data[0].Timestamp
		} else {
			exitTime = tmp2.Data.Data[0].Timestamp
		}
	}
	if noramlTime < exitTime {
		return true
	}
	return false
}

// 利用拨号判断切换线路的条件之一
func AikuaiDialJudgeLine(url, setCookie string) bool {
	datas := `{
		"func_name": "monitor_iface",
		"action": "show",
		"param": {
		  "TYPE": "iface_check,iface_stream,ether_info,snapshoot"
		}
	  }`
	err, s2 := crawler.PostNew(url+"/Action/call", datas, setCookie, false)
	if err != nil {
		return false
	}
	var tmp2 NewAikuaiResp
	err = json.Unmarshal([]byte(s2), &tmp2)
	if err != nil {
		return false
	}
	if tmp2.ErrMsg != "Success" {
		return false
	}
	for _, v := range tmp2.Data.IfaceCheck {
		if v.Interface == "l2tp_tfgame2" && v.Errmsg != "线路检测成功" {
			return true
		}
	}
	return false
}
